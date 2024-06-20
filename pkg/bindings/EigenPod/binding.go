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
	Bin: "0x60e060405260638054600160ff1991821681179092556067805490911690911790553480156200002e57600080fd5b50604051620053863803806200538683398101604081905262000051916200015f565b6001600160a01b03808416608052821660a0526001600160401b03811660c0526200007b62000084565b505050620001be565b600054610100900460ff1615620000f15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000144576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015c57600080fd5b50565b6000806000606084860312156200017557600080fd5b8351620001828162000146565b6020850151909350620001958162000146565b60408501519092506001600160401b0381168114620001b357600080fd5b809150509250925092565b60805160a05160c05161515262000234600039600061073c0152600081816103c20152818161079101528181610c2501528181610f6b0152818161134901528181611534015281816119d101528181611e1d01528181612026015261375401526000818161059f01526115ff01526151526000f3fe6080604052600436106101f25760003560e01c80637439841f1161010d578063ba414fa6116100a0578063e20c9f711161006f578063e20c9f71146106d5578063ee94d67c146106ea578063f074ba621461070a578063f28824611461072a578063fa7626d41461075e57600080fd5b8063ba414fa614610660578063c490744214610675578063c4d66de814610695578063dda3346c146106b557600080fd5b8063916a17c6116100dc578063916a17c6146106035780639b4e463414610618578063b522538a1461062b578063b5508aa91461064b57600080fd5b80637439841f1461055657806374cdd7981461058d57806385226c81146105c157806388676cad146105e357600080fd5b80633f7286f41161018557806352396a591161015457806352396a59146104a457806358eaee79146104da57806366d9a9a0146105075780636fcd0e531461052957600080fd5b80633f7286f41461037457806342ecff2a146103895780634665bcda146103b057806347d28372146103e457600080fd5b80633106ab53116101c15780633106ab53146102d65780633474aa16146103075780633e5e3c231461033f5780633f65cf191461035457600080fd5b8063039157d2146102315780630b18ff66146102535780631ed7831c146102905780632340e8d3146102b257600080fd5b3661022c576040513481527f6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf499060200160405180910390a1005b600080fd5b34801561023d57600080fd5b5061025161024c366004614277565b610778565b005b34801561025f57600080fd5b50603354610273906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561029c57600080fd5b506102a5610b20565b60405161028791906142ec565b3480156102be57600080fd5b506102c860395481565b604051908152602001610287565b3480156102e257600080fd5b506034546102f790600160401b900460ff1681565b6040519015158152602001610287565b34801561031357600080fd5b50603454610327906001600160401b031681565b6040516001600160401b039091168152602001610287565b34801561034b57600080fd5b506102a5610b82565b34801561036057600080fd5b5061025161036f366004614384565b610be2565b34801561038057600080fd5b506102a5610fd5565b34801561039557600080fd5b50603a5461032790600160401b90046001600160401b031681565b3480156103bc57600080fd5b506102737f000000000000000000000000000000000000000000000000000000000000000081565b3480156103f057600080fd5b506104606040805160808101825260008082526020820181905291810182905260608101919091525060408051608081018252603c548152603d5462ffffff811660208301526001600160401b03630100000082041692820192909252600160581b909104600f0b606082015290565b60405161028791908151815260208083015162ffffff16908201526040808301516001600160401b031690820152606091820151600f0b9181019190915260800190565b3480156104b057600080fd5b506103276104bf366004614454565b603b602052600090815260409020546001600160401b031681565b3480156104e657600080fd5b506104fa6104f53660046144b2565b611035565b604051610287919061452b565b34801561051357600080fd5b5061051c61109a565b6040516102879190614539565b34801561053557600080fd5b506105496105443660046145ec565b611189565b6040516102879190614605565b34801561056257600080fd5b506104fa6105713660046145ec565b600090815260366020526040902054600160c01b900460ff1690565b34801561059957600080fd5b506102737f000000000000000000000000000000000000000000000000000000000000000081565b3480156105cd57600080fd5b506105d6611236565b60405161028791906146a9565b3480156105ef57600080fd5b506102516105fe366004614719565b611306565b34801561060f57600080fd5b5061051c611443565b610251610626366004614736565b611529565b34801561063757600080fd5b506105496106463660046144b2565b6116d6565b34801561065757600080fd5b506105d66117c9565b34801561066c57600080fd5b506102f7611899565b34801561068157600080fd5b506102516106903660046147c9565b6119c6565b3480156106a157600080fd5b506102516106b03660046147f5565b611c03565b3480156106c157600080fd5b506102516106d03660046148e6565b611dda565b3480156106e157600080fd5b506102a5611fad565b3480156106f657600080fd5b50603a54610327906001600160401b031681565b34801561071657600080fd5b506102516107253660046149b7565b61200d565b34801561073657600080fd5b506103277f000000000000000000000000000000000000000000000000000000000000000081565b34801561076a57600080fd5b506063546102f79060ff1681565b604051635ac86ab760e01b8152600860048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa1580156107e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108049190614a1f565b1561082a5760405162461bcd60e51b815260040161082190614a3c565b60405180910390fd5b60006108706108398480614a99565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506123a192505050565b6000818152603660209081526040808320815160808101835281546001600160401b038082168352600160401b8204811695830195909552600160801b8104909416928101929092529394509192906060830190600160c01b900460ff1660028111156108df576108df6144f3565b60028111156108f0576108f06144f3565b8152505090506212750081604001516001600160401b03166109129190614af8565b866001600160401b03161161098f5760405162461bcd60e51b815260206004820152603f60248201527f456967656e506f642e7665726966795374616c6542616c616e63653a2076616c60448201527f696461746f722062616c616e6365206973206e6f74207374616c6520796574006064820152608401610821565b6001816060015160028111156109a7576109a76144f3565b14610a115760405162461bcd60e51b815260206004820152603460248201527f456967656e506f642e7665726966795374616c6542616c616e63653a2076616c604482015273696461746f72206973206e6f742061637469766560601b6064820152608401610821565b610a55610a1e8580614a99565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506123c592505050565b610ad95760405162461bcd60e51b815260206004820152604960248201527f456967656e506f642e7665726966795374616c6542616c616e63653a2076616c60448201527f696461746f72206d75737420626520736c617368656420746f206265206d61726064820152686b6564207374616c6560b81b608482015260a401610821565b610aeb610ae5876123ef565b866125a5565b610b0e8535610afa8680614a99565b610b076020890189614b10565b8651612700565b610b186000612917565b505050505050565b60606070805480602002602001604051908101604052809291908181526020018280548015610b7857602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610b5a575b5050505050905090565b60606072805480602002602001604051908101604052809291908181526020018280548015610b78576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610b5a575050505050905090565b6033546001600160a01b03163314610c0c5760405162461bcd60e51b815260040161082190614b56565b604051635ac86ab760e01b8152600260048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa158015610c74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c989190614a1f565b15610cb55760405162461bcd60e51b815260040161082190614a3c565b603454600160401b900460ff16610d345760405162461bcd60e51b815260206004820152603a60248201527f456967656e506f642e7665726966795769746864726177616c43726564656e7460448201527f69616c733a2072657374616b696e67206e6f74206163746976650000000000006064820152608401610821565b8584148015610d4257508382145b610dd25760405162461bcd60e51b815260206004820152605560248201527f456967656e506f642e7665726966795769746864726177616c43726564656e7460448201527f69616c733a2076616c696461746f72496e646963657320616e642070726f6f666064820152740e640daeae6e840c4ca40e6c2daca40d8cadccee8d605b1b608482015260a401610821565b603a546001600160401b03908116908a16118015610e055750603a546001600160401b03600160401b9091048116908a16115b610e8c5760405162461bcd60e51b815260206004820152604c60248201527f456967656e506f642e7665726966795769746864726177616c43726564656e7460448201527f69616c733a207370656369666965642074696d657374616d7020697320746f6f60648201526b0819985c881a5b881c185cdd60a21b608482015260a401610821565b610e9e610e988a6123ef565b896125a5565b6000805b87811015610f4157610f238a358a8a84818110610ec157610ec1614b9e565b9050602002016020810190610ed69190614bb4565b898985818110610ee857610ee8614b9e565b9050602002810190610efa9190614b10565b898987818110610f0c57610f0c614b9e565b9050602002810190610f1e9190614a99565b612b99565b610f2d9083614af8565b915080610f3981614bdb565b915050610ea2565b5060335460405163030b147160e61b81526001600160a01b039182166004820152602481018390527f00000000000000000000000000000000000000000000000000000000000000009091169063c2c51c4090604401600060405180830381600087803b158015610fb157600080fd5b505af1158015610fc5573d6000803e3d6000fd5b5050505050505050505050505050565b60606071805480602002602001604051908101604052809291908181526020018280548015610b78576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610b5a575050505050905090565b60008061107784848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061315492505050565b600090815260366020526040902054600160c01b900460ff169150505b92915050565b60606075805480602002602001604051908101604052809291908181526020016000905b828210156111805760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561116857602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841161112a5790505b505050505081525050815260200190600101906110be565b50505050905090565b6111b16040805160808101825260008082526020820181905291810182905290606082015290565b600082815260366020908152604091829020825160808101845281546001600160401b038082168352600160401b8204811694830194909452600160801b810490931693810193909352906060830190600160c01b900460ff16600281111561121c5761121c6144f3565b600281111561122d5761122d6144f3565b90525092915050565b60606074805480602002602001604051908101604052809291908181526020016000905b8282101561118057838290600052602060002001805461127990614bf6565b80601f01602080910402602001604051908101604052809291908181526020018280546112a590614bf6565b80156112f25780601f106112c7576101008083540402835291602001916112f2565b820191906000526020600020905b8154815290600101906020018083116112d557829003601f168201915b50505050508152602001906001019061125a565b6033546001600160a01b031633146113305760405162461bcd60e51b815260040161082190614b56565b604051635ac86ab760e01b8152600660048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa158015611398573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113bc9190614a1f565b156113d95760405162461bcd60e51b815260040161082190614a3c565b6113e282612917565b603454600160401b900460ff1661143f576034805460ff60401b1916600160401b1790556033546040516001600160a01b03909116907fca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd90600090a25b5050565b60606076805480602002602001604051908101604052809291908181526020016000905b828210156111805760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561151157602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116114d35790505b50505050508152505081526020019060010190611467565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146115715760405162461bcd60e51b815260040161082190614c2b565b346801bc16d674ec800000146115fd5760405162461bcd60e51b8152602060048201526044602482018190527f456967656e506f642e7374616b653a206d75737420696e697469616c6c792073908201527f74616b6520666f7220616e792076616c696461746f72207769746820333220656064820152633a3432b960e11b608482015260a401610821565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663228951186801bc16d674ec800000878761164061324e565b8888886040518863ffffffff1660e01b815260040161166496959493929190614ca5565b6000604051808303818588803b15801561167d57600080fd5b505af1158015611691573d6000803e3d6000fd5b50505050507f606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e2385856040516116c7929190614cf4565b60405180910390a15050505050565b6116fe6040805160808101825260008082526020820181905291810182905290606082015290565b6036600061174185858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061315492505050565b81526020808201929092526040908101600020815160808101835281546001600160401b038082168352600160401b8204811695830195909552600160801b81049094169281019290925290916060830190600160c01b900460ff1660028111156117ae576117ae6144f3565b60028111156117bf576117bf6144f3565b9052509392505050565b60606073805480602002602001604051908101604052809291908181526020016000905b8282101561118057838290600052602060002001805461180c90614bf6565b80601f016020809104026020016040519081016040528092919081815260200182805461183890614bf6565b80156118855780601f1061185a57610100808354040283529160200191611885565b820191906000526020600020905b81548152906001019060200180831161186857829003601f168201915b5050505050815260200190600101906117ed565b606354600090610100900460ff16156118bb5750606354610100900460ff1690565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d3b156119c15760408051737109709ecfa91a80626ff3989d68f67f5b1dd12d602082018190526519985a5b195960d21b82840152825180830384018152606083019093526000929091611949917f667f9d70ca411d70ead50d8d5c22070dafc36ad75f3dcf5e7237b22ade9aecc491608001614d08565b60408051601f198184030181529082905261196391614d39565b6000604051808303816000865af19150503d80600081146119a0576040519150601f19603f3d011682016040523d82523d6000602084013e6119a5565b606091505b50915050808060200190518101906119bd9190614a1f565b9150505b919050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611a0e5760405162461bcd60e51b815260040161082190614c2b565b611a1c633b9aca0082614d6b565b15611aa65760405162461bcd60e51b815260206004820152604e60248201527f456967656e506f642e776974686472617752657374616b6564426561636f6e4360448201527f6861696e4554483a20616d6f756e74576569206d75737420626520612077686f60648201526d1b194811ddd95a48185b5bdd5b9d60921b608482015260a401610821565b6000611ab6633b9aca0083614d7f565b6034549091506001600160401b039081169082161115611b6f5760405162461bcd60e51b815260206004820152606260248201527f456967656e506f642e776974686472617752657374616b6564426561636f6e4360448201527f6861696e4554483a20616d6f756e74477765692065786365656473207769746860648201527f6472617761626c6552657374616b6564457865637574696f6e4c617965724777608482015261656960f01b60a482015260c401610821565b60348054829190600090611b8d9084906001600160401b0316614d93565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550826001600160a01b03167f8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e83604051611bec91815260200190565b60405180910390a2611bfe8383613293565b505050565b600054610100900460ff1615808015611c235750600054600160ff909116105b80611c3d5750303b158015611c3d575060005460ff166001145b611ca05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610821565b6000805460ff191660011790558015611cc3576000805461ff0019166101001790555b6001600160a01b038216611d365760405162461bcd60e51b815260206004820152603460248201527f456967656e506f642e696e697469616c697a653a20706f644f776e65722063616044820152736e6e6f74206265207a65726f206164647265737360601b6064820152608401610821565b603380546001600160a01b0384166001600160a01b031990911681179091556034805460ff60401b1916600160401b1790556040517fca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd90600090a2801561143f576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6033546001600160a01b03163314611e045760405162461bcd60e51b815260040161082190614b56565b604051635ac86ab760e01b8152600560048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa158015611e6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e909190614a1f565b15611ead5760405162461bcd60e51b815260040161082190614a3c565b8251845114611f385760405162461bcd60e51b815260206004820152604b60248201527f456967656e506f642e7265636f766572546f6b656e733a20746f6b656e4c697360448201527f7420616e6420616d6f756e7473546f5769746864726177206d7573742062652060648201526a0e6c2daca40d8cadccee8d60ab1b608482015260a401610821565b60005b8451811015611fa657611f9483858381518110611f5a57611f5a614b9e565b6020026020010151878481518110611f7457611f74614b9e565b60200260200101516001600160a01b031661329d9092919063ffffffff16565b80611f9e81614bdb565b915050611f3b565b5050505050565b6060606f805480602002602001604051908101604052809291908181526020018280548015610b78576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610b5a575050505050905090565b604051635ac86ab760e01b8152600760048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa158015612075573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120999190614a1f565b156120b65760405162461bcd60e51b815260040161082190614a3c565b603a54600160401b90046001600160401b0316806121625760405162461bcd60e51b815260206004820152605860248201527f456967656e506f642e766572696679436865636b706f696e7450726f6f66733a60448201527f206d75737420686176652061637469766520636865636b706f696e7420746f2060648201527f706572666f726d20636865636b706f696e742070726f6f660000000000000000608482015260a401610821565b60408051608081018252603c54808252603d5462ffffff811660208401526001600160401b03630100000082041693830193909352600160581b909204600f0b6060820152906121b290876132ef565b6000805b8581101561233e576000603660008989858181106121d6576121d6614b9e565b90506020028101906121e89190614dbb565b3581526020808201929092526040908101600020815160808101835281546001600160401b038082168352600160401b8204811695830195909552600160801b81049094169281019290925290916060830190600160c01b900460ff166002811115612256576122566144f3565b6002811115612267576122676144f3565b9052509050600181606001516002811115612284576122846144f3565b1461228f575061232c565b846001600160401b031681604001516001600160401b0316106122b2575061232c565b6000806122e583888d358d8d898181106122ce576122ce614b9e565b90506020028101906122e09190614dbb565b61346b565b6020880180519294509092506122fa82614dd1565b62ffffff16905250606086018051839190612316908390614df0565b600f0b9052506123268186614e3f565b94505050505b8061233681614bdb565b9150506121b6565b506001600160401b038084166000908152603b602052604081208054849391929161236b91859116614e3f565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555061239882613671565b50505050505050565b6000816000815181106123b6576123b6614b9e565b60200260200101519050919050565b6000816003815181106123da576123da614b9e565b60200260200101516000801b14159050919050565b60006123fe611fff600c614e6a565b6124116001600160401b03841642614e89565b1061247b5760405162461bcd60e51b815260206004820152603460248201527f456967656e506f642e5f676574506172656e74426c6f636b526f6f743a2074696044820152736d657374616d70206f7574206f662072616e676560601b6064820152608401610821565b604080516001600160401b03841660208201526000918291720f3df6d732807ef1319fb7b8bb8522d0beac02910160408051601f19818403018152908290526124c391614d39565b600060405180830381855afa9150503d80600081146124fe576040519150601f19603f3d011682016040523d82523d6000602084013e612503565b606091505b5091509150818015612516575060008151115b15612537578080602001905181019061252f9190614ea0565b949350505050565b60405162461bcd60e51b815260206004820152603960248201527f456967656e506f642e5f676574506172656e74426c6f636b526f6f743a20696e60448201527f76616c696420626c6f636b20726f6f742072657475726e6564000000000000006064820152608401610821565b6125b160036020614e6a565b6125be6020830183614b10565b9050146126335760405162461bcd60e51b815260206004820152603d60248201527f426561636f6e436861696e50726f6f66732e7665726966795374617465526f6f60448201527f743a2050726f6f662068617320696e636f7272656374206c656e6774680000006064820152608401610821565b6126836126436020830183614b10565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525086925050843590506003613878565b61143f5760405162461bcd60e51b815260206004820152604260248201527f426561636f6e436861696e50726f6f66732e7665726966795374617465526f6f60448201527f743a20496e76616c696420737461746520726f6f74206d65726b6c652070726f60648201526137b360f11b608482015260a401610821565b6008841461277b5760405162461bcd60e51b815260206004820152604e60248201526000805160206150fd83398151915260448201527f724669656c64733a2056616c696461746f72206669656c64732068617320696e60648201526d0c6dee4e4cac6e840d8cadccee8d60931b608482015260a401610821565b600561278960286001614af8565b6127939190614af8565b61279e906020614e6a565b821461280c5760405162461bcd60e51b815260206004820152604360248201526000805160206150fd83398151915260448201527f724669656c64733a2050726f6f662068617320696e636f7272656374206c656e6064820152620cee8d60eb1b608482015260a401610821565b600061284a86868080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061389092505050565b9050600064ffffffffff831661286260286001614af8565b600b901b1790506128ad85858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250869150859050613878565b61290d5760405162461bcd60e51b815260206004820152603d60248201526000805160206150fd83398151915260448201527f724669656c64733a20496e76616c6964206d65726b6c652070726f6f660000006064820152608401610821565b5050505050505050565b603a54600160401b90046001600160401b0316156129b85760405162461bcd60e51b815260206004820152605260248201527f456967656e506f642e5f7374617274436865636b706f696e743a206d7573742060448201527f66696e6973682070726576696f757320636865636b706f696e74206265666f72606482015271329039ba30b93a34b7339030b737ba3432b960711b608482015260a401610821565b603a54426001600160401b0390811691161415612a3d5760405162461bcd60e51b815260206004820152603f60248201527f456967656e506f642e5f7374617274436865636b706f696e743a2063616e6e6f60448201527f7420636865636b706f696e7420747769636520696e206f6e6520626c6f636b006064820152608401610821565b6034546000906001600160401b0316612a5a633b9aca0047614d7f565b612a649190614d93565b9050818015612a7a57506001600160401b038116155b15612aed5760405162461bcd60e51b815260206004820152603d60248201527f456967656e506f642e5f7374617274436865636b706f696e743a206e6f20626160448201527f6c616e636520617661696c61626c6520746f20636865636b706f696e740000006064820152608401610821565b60006040518060800160405280612b03426123ef565b815260200160395462ffffff168152602001836001600160401b031681526020016000600f0b815250905042603a60086101000a8154816001600160401b0302191690836001600160401b03160217905550612b5e81613671565b80516040516001600160401b034216907f17a875d60c621f9c2f4d68122fe3a5bd359233b6fd8fcb79a0f65a6e433c6bf690600090a3505050565b600080612bd88484808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506123a192505050565b6000818152603660209081526040808320815160808101835281546001600160401b038082168352600160401b8204811695830195909552600160801b8104909416928101929092529394509192906060830190600160c01b900460ff166002811115612c4757612c476144f3565b6002811115612c5857612c586144f3565b9052509050600081606001516002811115612c7557612c756144f3565b14612d185760405162461bcd60e51b815260206004820152606160248201527f456967656e506f642e5f7665726966795769746864726177616c43726564656e60448201527f7469616c733a2076616c696461746f72206d75737420626520696e616374697660648201527f6520746f2070726f7665207769746864726177616c2063726564656e7469616c6084820152607360f81b60a482015260c401610821565b6001600160401b038016612d5e868680806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250613b3d92505050565b6001600160401b031614612de85760405162461bcd60e51b8152602060048201526044602482018190527f456967656e506f642e5f7665726966795769746864726177616c43726564656e908201527f7469616c733a2076616c696461746f72206d757374206e6f742062652065786960648201526374696e6760e01b608482015260a401610821565b612df061324e565b612df990614eb9565b612e35868680806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250613b6292505050565b14612eb65760405162461bcd60e51b815260206004820152604560248201527f456967656e506f642e5f7665726966795769746864726177616c43726564656e60448201527f7469616c733a2070726f6f66206973206e6f7420666f72207468697320456967606482015264195b941bd960da1b608482015260a401610821565b600085858080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604051634534711b60e01b815273__$4050e451cf207daf724f57201af415e6ef$__93634534711b9350612f21925090600401614edd565b602060405180830381865af4158015612f3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f629190614f15565b9050612f728a87878b8b8e612700565b60398054906000612f8283614bdb565b9091555050603a54600090600160401b90046001600160401b031615612fba57603a54600160401b90046001600160401b0316612fc7565b603a546001600160401b03165b6040805160808101825264ffffffffff8d1681526001600160401b03858116602083015283169181019190915290915060608101600190526000858152603660209081526040918290208351815492850151938501516001600160401b03908116600160801b0267ffffffffffffffff60801b19958216600160401b026001600160801b031990951691909216179290921792831682178155606084015190929091839160ff60c01b1990911668ffffffffffffffffff60801b1990911617600160c01b83600281111561309d5761309d6144f3565b02179055505060405164ffffffffff8c1681527f2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449915060200160405180910390a16040805164ffffffffff8c1681526001600160401b03838116602083015284168183015290517f0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df9181900360600190a1613145633b9aca006001600160401b038416614e6a565b9b9a5050505050505050505050565b600081516030146131dd5760405162461bcd60e51b815260206004820152604760248201527f456967656e506f642e5f63616c63756c61746556616c696461746f725075626b60448201527f657948617368206d75737420626520612034382d6279746520424c53207075626064820152666c6963206b657960c81b608482015260a401610821565b6040516002906131f4908490600090602001614f32565b60408051601f198184030181529082905261320e91614d39565b602060405180830381855afa15801561322b573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906110949190614ea0565b60408051600160f81b60208201526000602182015230606090811b6bffffffffffffffffffffffff1916602c8301529101604051602081830303815290604052905090565b61143f8282613b77565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052611bfe908490613c90565b6132fb60056003614af8565b613306906020614e6a565b6133136020830183614b10565b9050146133965760405162461bcd60e51b8152602060048201526044602482018190527f426561636f6e436861696e50726f6f66732e76657269667942616c616e636543908201527f6f6e7461696e65723a2050726f6f662068617320696e636f7272656374206c656064820152630dccee8d60e31b608482015260a401610821565b606c6133e76133a86020840184614b10565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250879250508535905084613878565b611bfe5760405162461bcd60e51b815260206004820152604960248201527f426561636f6e436861696e50726f6f66732e76657269667942616c616e63654360448201527f6f6e7461696e65723a20696e76616c69642062616c616e636520636f6e7461696064820152683732b910383937b7b360b91b608482015260a401610821565b83516020850151600091829182613483878488613d62565b9050816001600160401b0316816001600160401b0316146134fd576134a88183613ed9565b6040805164ffffffffff861681526001600160401b038b8116602083015284168183015290519196507f0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df919081900360600190a15b6001600160401b0380821660208b0181905290891660408b0152613581576039805490600061352b83614f61565b9091555050600260608a015261354085614f78565b93508264ffffffffff16886001600160401b03167f2a02361ffa66cf2c2da4682c2355a6adcaa9f6c227b6e6563e68480f9587626a60405160405180910390a35b85356000908152603660209081526040918290208b518154928d0151938d01516001600160401b03908116600160801b0267ffffffffffffffff60801b19958216600160401b026001600160801b03199095169190921617929092179283168217815560608c01518c9391929091839160ff60c01b191668ffffffffffffffffff60801b1990911617600160c01b836002811115613621576136216144f3565b02179055505060405164ffffffffff851691506001600160401b038a16907fa91c59033c3423e18b54d0acecebb4972f9ea95aedf5f4cae3b677b02eaf3a3f90600090a350505094509492505050565b602081015162ffffff166137f8576000633b9aca00826060015183604001516001600160401b03166136a39190614df0565b600f0b6136b09190614f9f565b60408301516034805492935090916000906136d59084906001600160401b0316614e3f565b82546101009290920a6001600160401b03818102199093169183160217909155603a8054600160401b81049092166001600160801b0319909216919091179055506000603c55603d80546001600160d81b031916905560335460405163030b147160e61b81526001600160a01b039182166004820152602481018390527f00000000000000000000000000000000000000000000000000000000000000009091169063c2c51c4090604401600060405180830381600087803b15801561379a57600080fd5b505af11580156137ae573d6000803e3d6000fd5b5050603a546040518481526001600160401b0390911692507f525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e44915060200160405180910390a25050565b8051603c556020810151603d8054604084015160608501516fffffffffffffffffffffffffffffffff16600160581b026fffffffffffffffffffffffffffffffff60581b196001600160401b039092166301000000026affffffffffffffffffffff1990931662ffffff9095169490941791909117169190911790555b50565b600083613886868585613ef1565b1495945050505050565b600080600283516138a19190614d7f565b90506000816001600160401b038111156138bd576138bd614812565b6040519080825280602002602001820160405280156138e6578160200160208202803683370190505b50905060005b828110156139ed576002856139018383614e6a565b8151811061391157613911614b9e565b6020026020010151868360026139279190614e6a565b613932906001614af8565b8151811061394257613942614b9e565b6020026020010151604051602001613964929190918252602082015260400190565b60408051601f198184030181529082905261397e91614d39565b602060405180830381855afa15801561399b573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906139be9190614ea0565b8282815181106139d0576139d0614b9e565b6020908102919091010152806139e581614bdb565b9150506138ec565b506139f9600283614d7f565b91505b8115613b195760005b82811015613b0657600282613a1a8383614e6a565b81518110613a2a57613a2a614b9e565b602002602001015183836002613a409190614e6a565b613a4b906001614af8565b81518110613a5b57613a5b614b9e565b6020026020010151604051602001613a7d929190918252602082015260400190565b60408051601f1981840301815290829052613a9791614d39565b602060405180830381855afa158015613ab4573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190613ad79190614ea0565b828281518110613ae957613ae9614b9e565b602090810291909101015280613afe81614bdb565b915050613a05565b50613b12600283614d7f565b91506139fc565b80600081518110613b2c57613b2c614b9e565b602002602001015192505050919050565b600061109482600681518110613b5557613b55614b9e565b602002602001015161403d565b6000816001815181106123b6576123b6614b9e565b80471015613bc75760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610821565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114613c14576040519150601f19603f3d011682016040523d82523d6000602084013e613c19565b606091505b5050905080611bfe5760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610821565b6000613ce5826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166140a49092919063ffffffff16565b805190915015611bfe5780806020019051810190613d039190614a1f565b611bfe5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610821565b6000613d7060266001614af8565b613d7b906020614e6a565b613d886040840184614b10565b905014613df95760405162461bcd60e51b8152602060048201526044602482018190526000805160206150fd833981519152908201527f7242616c616e63653a2050726f6f662068617320696e636f7272656374206c656064820152630dccee8d60e31b608482015260a401610821565b6000613e06600485615024565b64ffffffffff169050613e60613e1f6040850185614b10565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508992505050602086013584613878565b613ec05760405162461bcd60e51b815260206004820152603e60248201526000805160206150fd83398151915260448201527f7242616c616e63653a20496e76616c6964206d65726b6c652070726f6f6600006064820152608401610821565b613ece8360200135856140b3565b9150505b9392505050565b6000613ed26001600160401b03808416908516615048565b60008351600014158015613f10575060208451613f0e9190614d6b565b155b613f9f5760405162461bcd60e51b815260206004820152605460248201527f4d65726b6c652e70726f63657373496e636c7573696f6e50726f6f665368613260448201527f35363a2070726f6f66206c656e6774682073686f756c642062652061206e6f6e60648201527316bd32b9379036bab63a34b836329037b310199960611b608482015260a401610821565b604080516020808201909252848152905b8551811161403357613fc3600285614d6b565b613ff6578151600052808601516020526020826040600060026107d05a03fa613feb57600080fd5b600284049350614021565b8086015160005281516020526020826040600060026107d05a03fa61401a57600080fd5b6002840493505b61402c602082614af8565b9050613fb0565b5051949350505050565b60f881901c60e882901c61ff00161760d882901c62ff0000161760c882901c63ff000000161764ff0000000060b883901c161765ff000000000060a883901c161766ff000000000000609883901c161767ff0000000000000060889290921c919091161790565b606061252f84846000856140e0565b6000806140c1600484615098565b6140cc9060406150bc565b64ffffffffff16905061252f84821b61403d565b6060824710156141415760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610821565b6001600160a01b0385163b6141985760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610821565b600080866001600160a01b031685876040516141b49190614d39565b60006040518083038185875af1925050503d80600081146141f1576040519150601f19603f3d011682016040523d82523d6000602084013e6141f6565b606091505b5091509150614206828286614211565b979650505050505050565b60608315614220575081613ed2565b8251156142305782518084602001fd5b8160405162461bcd60e51b815260040161082191906150e9565b6001600160401b038116811461387557600080fd5b60006040828403121561427157600080fd5b50919050565b60008060006060848603121561428c57600080fd5b83356142978161424a565b925060208401356001600160401b03808211156142b357600080fd5b6142bf8783880161425f565b935060408601359150808211156142d557600080fd5b506142e28682870161425f565b9150509250925092565b6020808252825182820181905260009190848201906040850190845b8181101561432d5783516001600160a01b031683529284019291840191600101614308565b50909695505050505050565b60008083601f84011261434b57600080fd5b5081356001600160401b0381111561436257600080fd5b6020830191508360208260051b850101111561437d57600080fd5b9250929050565b60008060008060008060008060a0898b0312156143a057600080fd5b88356143ab8161424a565b975060208901356001600160401b03808211156143c757600080fd5b6143d38c838d0161425f565b985060408b01359150808211156143e957600080fd5b6143f58c838d01614339565b909850965060608b013591508082111561440e57600080fd5b61441a8c838d01614339565b909650945060808b013591508082111561443357600080fd5b506144408b828c01614339565b999c989b5096995094979396929594505050565b60006020828403121561446657600080fd5b8135613ed28161424a565b60008083601f84011261448357600080fd5b5081356001600160401b0381111561449a57600080fd5b60208301915083602082850101111561437d57600080fd5b600080602083850312156144c557600080fd5b82356001600160401b038111156144db57600080fd5b6144e785828601614471565b90969095509350505050565b634e487b7160e01b600052602160045260246000fd5b6003811061452757634e487b7160e01b600052602160045260246000fd5b9052565b602081016110948284614509565b60006020808301818452808551808352604092508286019150828160051b8701018488016000805b848110156145dd57898403603f19018652825180516001600160a01b03168552880151888501889052805188860181905290890190839060608701905b808310156145c85783516001600160e01b0319168252928b019260019290920191908b019061459e565b50978a01979550505091870191600101614561565b50919998505050505050505050565b6000602082840312156145fe57600080fd5b5035919050565b60006080820190506001600160401b038084511683528060208501511660208401528060408501511660408401525060608301516146466060840182614509565b5092915050565b60005b83811015614668578181015183820152602001614650565b83811115614677576000848401525b50505050565b6000815180845261469581602086016020860161464d565b601f01601f19169290920160200192915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156146fe57603f198886030184526146ec85835161467d565b945092850192908501906001016146d0565b5092979650505050505050565b801515811461387557600080fd5b60006020828403121561472b57600080fd5b8135613ed28161470b565b60008060008060006060868803121561474e57600080fd5b85356001600160401b038082111561476557600080fd5b61477189838a01614471565b9097509550602088013591508082111561478a57600080fd5b5061479788828901614471565b96999598509660400135949350505050565b6001600160a01b038116811461387557600080fd5b80356119c1816147a9565b600080604083850312156147dc57600080fd5b82356147e7816147a9565b946020939093013593505050565b60006020828403121561480757600080fd5b8135613ed2816147a9565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171561485057614850614812565b604052919050565b60006001600160401b0382111561487157614871614812565b5060051b60200190565b600082601f83011261488c57600080fd5b813560206148a161489c83614858565b614828565b82815260059290921b840181019181810190868411156148c057600080fd5b8286015b848110156148db57803583529183019183016148c4565b509695505050505050565b6000806000606084860312156148fb57600080fd5b83356001600160401b038082111561491257600080fd5b818601915086601f83011261492657600080fd5b8135602061493661489c83614858565b82815260059290921b8401810191818101908a84111561495557600080fd5b948201945b8386101561497c57853561496d816147a9565b8252948201949082019061495a565b9750508701359250508082111561499257600080fd5b5061499f8682870161487b565b9250506149ae604085016147be565b90509250925092565b6000806000604084860312156149cc57600080fd5b83356001600160401b03808211156149e357600080fd5b6149ef8783880161425f565b94506020860135915080821115614a0557600080fd5b50614a1286828701614339565b9497909650939450505050565b600060208284031215614a3157600080fd5b8151613ed28161470b565b6020808252603e908201527f456967656e506f642e6f6e6c795768656e4e6f745061757365643a20696e646560408201527f782069732070617573656420696e20456967656e506f644d616e616765720000606082015260800190565b6000808335601e19843603018112614ab057600080fd5b8301803591506001600160401b03821115614aca57600080fd5b6020019150600581901b360382131561437d57600080fd5b634e487b7160e01b600052601160045260246000fd5b60008219821115614b0b57614b0b614ae2565b500190565b6000808335601e19843603018112614b2757600080fd5b8301803591506001600160401b03821115614b4157600080fd5b60200191503681900382131561437d57600080fd5b60208082526028908201527f456967656e506f642e6f6e6c79456967656e506f644f776e65723a206e6f74206040820152673837b227bbb732b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600060208284031215614bc657600080fd5b813564ffffffffff81168114613ed257600080fd5b6000600019821415614bef57614bef614ae2565b5060010190565b600181811c90821680614c0a57607f821691505b6020821081141561427157634e487b7160e01b600052602260045260246000fd5b60208082526031908201527f456967656e506f642e6f6e6c79456967656e506f644d616e616765723a206e6f6040820152703a1032b4b3b2b72837b226b0b730b3b2b960791b606082015260800190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b608081526000614cb960808301888a614c7c565b8281036020840152614ccb818861467d565b90508281036040840152614ce0818688614c7c565b915050826060830152979650505050505050565b60208152600061252f602083018486614c7c565b6001600160e01b0319831681528151600090614d2b81600485016020870161464d565b919091016004019392505050565b60008251614d4b81846020870161464d565b9190910192915050565b634e487b7160e01b600052601260045260246000fd5b600082614d7a57614d7a614d55565b500690565b600082614d8e57614d8e614d55565b500490565b60006001600160401b0383811690831681811015614db357614db3614ae2565b039392505050565b60008235605e19833603018112614d4b57600080fd5b600062ffffff821680614de657614de6614ae2565b6000190192915050565b600081600f0b83600f0b600082128260016001607f1b0303821381151615614e1a57614e1a614ae2565b8260016001607f1b0319038212811615614e3657614e36614ae2565b50019392505050565b60006001600160401b03808316818516808303821115614e6157614e61614ae2565b01949350505050565b6000816000190483118215151615614e8457614e84614ae2565b500290565b600082821015614e9b57614e9b614ae2565b500390565b600060208284031215614eb257600080fd5b5051919050565b805160208083015191908110156142715760001960209190910360031b1b16919050565b6020808252825182820181905260009190848201906040850190845b8181101561432d57835183529284019291840191600101614ef9565b600060208284031215614f2757600080fd5b8151613ed28161424a565b60008351614f4481846020880161464d565b6001600160801b0319939093169190920190815260100192915050565b600081614f7057614f70614ae2565b506000190190565b600081600f0b60016001607f1b0319811415614f9657614f96614ae2565b60000392915050565b60006001600160ff1b0381841382841380821686840486111615614fc557614fc5614ae2565b600160ff1b6000871282811687830589121615614fe457614fe4614ae2565b6000871292508782058712848416161561500057615000614ae2565b8785058712818416161561501657615016614ae2565b505050929093029392505050565b600064ffffffffff8084168061503c5761503c614d55565b92169190910492915050565b600081600f0b83600f0b600081128160016001607f1b03190183128115161561507357615073614ae2565b8160016001607f1b0301831381161561508e5761508e614ae2565b5090039392505050565b600064ffffffffff808416806150b0576150b0614d55565b92169190910692915050565b600064ffffffffff808316818516818304811182151516156150e0576150e0614ae2565b02949350505050565b602081526000613ed2602083018461467d56fe426561636f6e436861696e50726f6f66732e76657269667956616c696461746fa2646970667358221220a391ad2b9906874edfa8ae59f593983d91c803ca0b5da2b2000c8f562bedcbad64736f6c634300080c0033",
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
