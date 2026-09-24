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

// IEigenPodTypesCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodTypesCheckpoint struct {
	BeaconBlockRoot       [32]byte
	ProofsRemaining       *big.Int
	PodBalanceGwei        uint64
	BalanceDeltasGwei     int64
	PrevBeaconBalanceGwei uint64
}

// IEigenPodTypesConsolidationRequest is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodTypesConsolidationRequest struct {
	SrcPubkey    []byte
	TargetPubkey []byte
}

// IEigenPodTypesValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodTypesValidatorInfo struct {
	ValidatorIndex      uint64
	RestakedBalanceGwei uint64
	LastCheckpointedAt  uint64
	Status              uint8
}

// IEigenPodTypesWithdrawalRequest is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodTypesWithdrawalRequest struct {
	Pubkey     []byte
	AmountGwei uint64
}

// EigenPodMetaData contains all meta data concerning the EigenPod contract.
var EigenPodMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_ethPOS\",\"type\":\"address\",\"internalType\":\"contractIETHPOSDeposit\"},{\"name\":\"_eigenPodManager\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"activeValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointBalanceExitedGwei\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentCheckpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPodTypes.Checkpoint\",\"components\":[{\"name\":\"beaconBlockRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proofsRemaining\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"podBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"balanceDeltasGwei\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"prevBeaconBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentCheckpointTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableRestaking\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethPOS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIETHPOSDeposit\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConsolidationRequestFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParentBlockRoot\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawalRequestFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastCheckpointTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"podOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proofSubmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverTokens\",\"inputs\":[{\"name\":\"tokenList\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"amountsToWithdraw\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestConsolidation\",\"inputs\":[{\"name\":\"requests\",\"type\":\"tuple[]\",\"internalType\":\"structIEigenPodTypes.ConsolidationRequest[]\",\"components\":[{\"name\":\"srcPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"targetPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestWithdrawal\",\"inputs\":[{\"name\":\"requests\",\"type\":\"tuple[]\",\"internalType\":\"structIEigenPodTypes.WithdrawalRequest[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amountGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"restakingDisabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setProofSubmitter\",\"inputs\":[{\"name\":\"newProofSubmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"startCheckpoint\",\"inputs\":[{\"name\":\"revertIfNoBalance\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorPubkeyHashToInfo\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPodTypes.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastCheckpointedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPodTypes.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorPubkeyToInfo\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPodTypes.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastCheckpointedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPodTypes.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPodTypes.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPodTypes.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCheckpointProofs\",\"inputs\":[{\"name\":\"balanceContainerProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.BalanceContainerProof\",\"components\":[{\"name\":\"balanceContainerRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"proofs\",\"type\":\"tuple[]\",\"internalType\":\"structBeaconChainProofs.BalanceProof[]\",\"components\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"balanceRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyStaleBalance\",\"inputs\":[{\"name\":\"beaconTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.ValidatorProof\",\"components\":[{\"name\":\"validatorFields\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyWithdrawalCredentials\",\"inputs\":[{\"name\":\"beaconTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"validatorIndices\",\"type\":\"uint40[]\",\"internalType\":\"uint40[]\"},{\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawDisabledPodETH\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawRestakedBeaconChainETH\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawableRestakedExecutionLayerGwei\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CheckpointCreated\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"beaconBlockRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validatorCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointFinalized\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"totalShareDeltaWei\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConsolidationRequested\",\"inputs\":[{\"name\":\"sourcePubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"targetPubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisabledPodETHWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EigenPodStaked\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExitRequested\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonBeaconChainETHReceived\",\"inputs\":[{\"name\":\"amountReceived\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProofSubmitterUpdated\",\"inputs\":[{\"name\":\"prevProofSubmitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newProofSubmitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RestakedBeaconChainETHWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RestakingPermanentlyDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SwitchToCompoundingRequested\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorBalanceUpdated\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"balanceTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newValidatorBalanceGwei\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorCheckpointed\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRestaked\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorWithdrawn\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalRequested\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"withdrawalAmountGwei\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BeaconTimestampBeforeLatestCheckpoint\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BeaconTimestampTooFarInPast\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotCheckpointTwiceInSingleBlock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CheckpointAlreadyActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CredentialsAlreadyVerified\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeQueryFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForkTimestampZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFunds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientWithdrawableBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEIP4788Response\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPubKeyLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorFieldsLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LeavesNotPowerOfTwo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MsgValueNot32ETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoActiveCheckpoint\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoBalanceToCheckpoint\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughLeaves\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPodManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPodOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPodOwnerOrProofSubmitter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PredeployFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RestakingDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RestakingNotDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TimestampOutOfRange\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorInactiveOnBeaconChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorIsExitingBeaconChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorNotActiveInPod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorNotSlashedOnBeaconChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalCredentialsNotForEigenPod\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b50604051614b1f380380614b1f83398101604081905261002e91610123565b6001600160a01b03808316608052811660a052610049610050565b505061015b565b5f54610100900460ff16156100bb5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161461010a575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610120575f5ffd5b50565b5f5f60408385031215610134575f5ffd5b825161013f8161010c565b60208401519092506101508161010c565b809150509250929050565b60805160a0516149336101ec5f395f818161033f0152818161072c015281816107d401528181610a8201528181610e4d01528181611072015281816110e201528181611233015281816117040152818161189d0152818161194001528181611bf401528181611f6d015281816120a10152818161257701526134c801525f818161058301526119d401526149335ff3fe6080604052600436106101d3575f3560e01c80636c0d2d5a116100fd578063c44e30dc11610092578063d06d558711610062578063d06d558714610697578063dda3346c146106b6578063ee94d67c146106d5578063f074ba62146106f4575f5ffd5b8063c44e30dc14610615578063c452eeec14610629578063c490744214610659578063c4d66de814610678575f5ffd5b80638842752c116100cd5780638842752c146105a557806388676cad146105c45780639b4e4634146105e3578063b522538a146105f6575f5ffd5b80636c0d2d5a146104f25780636fcd0e53146105115780637439841f1461053d57806374cdd79814610572575f5ffd5b806342ecff2a11610173578063566f95fa11610143578063566f95fa14610480578063587533571461049457806358eaee79146104b35780636691954e146104df575f5ffd5b806342ecff2a146103085780634665bcda1461032e57806347d283721461036157806352396a591461044c575f5ffd5b80632340e8d3116101ae5780632340e8d3146102905780633474aa16146102a55780633f5fa57a146102d65780633f65cf19146102e9575f5ffd5b8063039157d2146102115780630b18ff66146102325780631e5155331461026e575f5ffd5b3661020d576040513481527f6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf499060200160405180910390a1005b5f5ffd5b34801561021c575f5ffd5b5061023061022b366004613dda565b610713565b005b34801561023d575f5ffd5b50603354610251906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b348015610279575f5ffd5b50610282610a48565b604051908152602001610265565b34801561029b575f5ffd5b5061028260395481565b3480156102b0575f5ffd5b506034546001600160401b03165b6040516001600160401b039091168152602001610265565b6102306102e4366004613e95565b610a69565b3480156102f4575f5ffd5b50610230610303366004613ed3565b610df4565b348015610313575f5ffd5b50603a546102be90600160401b90046001600160401b031681565b348015610339575f5ffd5b506102517f000000000000000000000000000000000000000000000000000000000000000081565b34801561036c575f5ffd5b506103f16040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152506040805160a081018252603c548152603d5462ffffff811660208301526001600160401b0363010000008204811693830193909352600160581b810460070b6060830152600160981b9004909116608082015290565b60405161026591905f60a0820190508251825262ffffff60208401511660208301526001600160401b036040840151166040830152606083015160070b60608301526001600160401b03608084015116608083015292915050565b348015610457575f5ffd5b506102be610466366004613faa565b603b6020525f90815260409020546001600160401b031681565b34801561048b575f5ffd5b506102306110d7565b34801561049f575f5ffd5b50603e54610251906001600160a01b031681565b3480156104be575f5ffd5b506104d26104cd366004614002565b6111b8565b6040516102659190614068565b6102306104ed366004613e95565b61121a565b3480156104fd575f5ffd5b5061028261050c366004613faa565b611507565b34801561051c575f5ffd5b5061053061052b366004614076565b611615565b604051610265919061408d565b348015610548575f5ffd5b506104d2610557366004614076565b5f90815260366020526040902054600160c01b900460ff1690565b34801561057d575f5ffd5b506102517f000000000000000000000000000000000000000000000000000000000000000081565b3480156105b0575f5ffd5b506102306105bf366004614104565b6116c0565b3480156105cf575f5ffd5b506102306105de36600461412c565b611844565b6102306105f1366004614147565b611935565b348015610601575f5ffd5b50610530610610366004614002565b611ade565b348015610620575f5ffd5b50610282611bcd565b348015610634575f5ffd5b50603e5461064990600160a01b900460ff1681565b6040519015158152602001610265565b348015610664575f5ffd5b506102306106733660046141b8565b611be9565b348015610683575f5ffd5b50610230610692366004614104565b611d4b565b3480156106a2575f5ffd5b506102306106b1366004614104565b611e95565b3480156106c1575f5ffd5b506102306106d03660046142b2565b611f29565b3480156106e0575f5ffd5b50603a546102be906001600160401b031681565b3480156106ff575f5ffd5b5061023061070e366004614384565b612088565b604051635ac86ab760e01b8152600660048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa158015610779573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061079d91906143eb565b156107bb5760405163840a48d560e01b815260040160405180910390fd5b604051635ac86ab760e01b8152600860048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa158015610821573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061084591906143eb565b156108635760405163840a48d560e01b815260040160405180910390fd5b5f6108a76108718580614406565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525061248492505050565b5f818152603660209081526040808320815160808101835281546001600160401b038082168352600160401b8204811695830195909552600160801b8104909416928101929092529394509192906060830190600160c01b900460ff16600281111561091557610915614034565b600281111561092657610926614034565b81525050905080604001516001600160401b0316876001600160401b031611610962576040516337e07ffd60e01b815260040160405180910390fd5b60018160600151600281111561097a5761097a614034565b146109985760405163d49e19a760e01b815260040160405180910390fd5b6109db6109a58680614406565b808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152506124a692505050565b6109f85760405163161ce5ed60e31b815260040160405180910390fd5b610a0a610a0488611507565b876124ce565b610a36610a1688612573565b8735610a228880614406565b610a2f60208b018b61444b565b875161264a565b610a3f5f61277e565b50505050505050565b5f610a6471bbddc7ce488642fb579f8b00f3a590007251612929565b905090565b604051635ac86ab760e01b8152600a60048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa158015610acf573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610af391906143eb565b15610b115760405163840a48d560e01b815260040160405180910390fd5b6033546001600160a01b0316331480610b345750603e546001600160a01b031633145b610b515760405163427a777960e01b815260040160405180910390fd5b603e54600160a01b900460ff165f610b67611bcd565b90505f610b7485836144a1565b905080341015610b975760405163356680b760e01b815260040160405180910390fd5b5f610ba282346144b8565b90505f5b86811015610de35736888883818110610bc157610bc16144cb565b9050602002810190610bd391906144df565b90505f610c1c610be3838061444b565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506129af92505050565b905086610c6c5760015f82815260366020526040902054600160c01b900460ff166002811115610c4e57610c4e614034565b14610c6c5760405163d49e19a760e01b815260040160405180910390fd5b5f610c77838061444b565b610c876040860160208701613faa565b604051602001610c99939291906144fd565b60405160208183030381529060405290505f710961ef480eb55e80d19ad83579a64c0070026001600160a01b03168883604051610cd69190614535565b5f6040518083038185875af1925050503d805f8114610d10576040519150601f19603f3d011682016040523d82523d5f602084013e610d15565b606091505b5050905080610d375760405163fc52d48360e01b815260040160405180910390fd5b610d476040850160208601613faa565b6001600160401b03165f03610d855760405183907f60d8ca014d4765a2b8b389e25714cb1cef83b574222911a01d90c1bd69d2d320905f90a2610dd3565b827f8b2737bb64ab2f2dc09552dfa1c250399e6a42c7ea9f0e1c658f5d65d708ec05610db76040870160208801613faa565b6040516001600160401b03909116815260200160405180910390a25b505060019092019150610ba69050565b508015610a3f57610a3f3382612a40565b6033546001600160a01b0316331480610e175750603e546001600160a01b031633145b610e345760405163427a777960e01b815260040160405180910390fd5b604051635ac86ab760e01b8152600260048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa158015610e9a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ebe91906143eb565b15610edc5760405163840a48d560e01b815260040160405180910390fd5b603e54600160a01b900460ff1615610f0757604051634549b2d160e11b815260040160405180910390fd5b8584148015610f1557508382145b610f32576040516343714afd60e01b815260040160405180910390fd5b603a546001600160401b03600160401b9091048116908a1611610f68576040516337e07ffd60e01b815260040160405180910390fd5b603a546001600160401b03908116908a1611610f97576040516335e7f6b760e01b815260040160405180910390fd5b610fa9610fa38a611507565b896124ce565b5f805b878110156110425761102e8b8b358b8b85818110610fcc57610fcc6144cb565b9050602002016020810190610fe19190614540565b8a8a86818110610ff357610ff36144cb565b9050602002810190611005919061444b565b8a8a88818110611017576110176144cb565b90506020028101906110299190614406565b612b55565b6110389083614564565b9150600101610fac565b5060335460405163a1ca780b60e01b81526001600160a01b0391821660048201525f6024820152604481018390527f00000000000000000000000000000000000000000000000000000000000000009091169063a1ca780b906064015f604051808303815f87803b1580156110b5575f5ffd5b505af11580156110c7573d5f5f3e3d5ffd5b5050505050505050505050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461112057604051633213a66160e21b815260040160405180910390fd5b603e54600160a01b900460ff161561114b57604051634549b2d160e11b815260040160405180910390fd5b603a54600160401b90046001600160401b03161561117b5760405162be9bc360e81b815260040160405180910390fd5b603e805460ff60a01b1916600160a01b1790556040517f4370e346659e8dc4bc97977bf89c631877baae18e45c9050029ca10247f68a69905f90a1565b5f5f6111f884848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506129af92505050565b5f90815260366020526040902054600160c01b900460ff169150505b92915050565b604051635ac86ab760e01b8152600960048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa158015611280573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112a491906143eb565b156112c25760405163840a48d560e01b815260040160405180910390fd5b6033546001600160a01b031633146112ed5760405163719f370360e11b815260040160405180910390fd5b603e54600160a01b900460ff166113175760405163db78bf0f60e01b815260040160405180910390fd5b5f611320610a48565b90505f61132d84836144a1565b9050803410156113505760405163356680b760e01b815260040160405180910390fd5b5f61135b82346144b8565b90505f5b858110156114ee573687878381811061137a5761137a6144cb565b905060200281019061138c91906144df565b90505f61139c610be3838061444b565b90505f6113af610be3602085018561444b565b90505f6113bc848061444b565b6113c9602087018761444b565b6040516020016113dc9493929190614577565b60405160208183030381529060405290505f71bbddc7ce488642fb579f8b00f3a5900072516001600160a01b031689836040516114199190614535565b5f6040518083038185875af1925050503d805f8114611453576040519150601f19603f3d011682016040523d82523d5f602084013e611458565b606091505b505090508061147a5760405163fc52d48360e01b815260040160405180910390fd5b8284036114b05760405184907fc97b965b92ae7fd20095fe8eb7b99f81f95f8c4adffb22a19116d8eb2846b016905f90a26114dd565b604051839085907f42f9c9db2ca443e9ec62f4588bd0c9b241065c02c2a8001ac164ae1282dc7b94905f90a35b50506001909301925061135f915050565b5080156114ff576114ff3382612a40565b505050505050565b5f611515611fff600c6144a1565b6115286001600160401b038416426144b8565b1061154657604051637944e66d60e11b815260040160405180910390fd5b604080516001600160401b03841660208201525f918291720f3df6d732807ef1319fb7b8bb8522d0beac02910160408051601f198184030181529082905261158d91614535565b5f60405180830381855afa9150503d805f81146115c5576040519150601f19603f3d011682016040523d82523d5f602084013e6115ca565b606091505b50915091508180156115dc57505f8151115b6115f95760405163558ad0a360e01b815260040160405180910390fd5b8080602001905181019061160d9190614596565b949350505050565b61163c604080516080810182525f8082526020820181905291810182905290606082015290565b5f82815260366020908152604091829020825160808101845281546001600160401b038082168352600160401b8204811694830194909452600160801b810490931693810193909352906060830190600160c01b900460ff1660028111156116a6576116a6614034565b60028111156116b7576116b7614034565b90525092915050565b6033546001600160a01b031633146116eb5760405163719f370360e11b815260040160405180910390fd5b604051635ac86ab760e01b8152600560048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa158015611751573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061177591906143eb565b156117935760405163840a48d560e01b815260040160405180910390fd5b61179b613056565b6001600160a01b0382166117c2576040516339b190bb60e11b815260040160405180910390fd5b603e54600160a01b900460ff166117ec5760405163db78bf0f60e01b815260040160405180910390fd5b60405147808252906001600160a01b038416907f2273c420d02f2b6ef92ed2c3d60848ab18c03c0ce20d6d2cf736030838f08f469060200160405180910390a26118368382612a40565b5061184060018055565b5050565b6033546001600160a01b03163314806118675750603e546001600160a01b031633145b6118845760405163427a777960e01b815260040160405180910390fd5b604051635ac86ab760e01b8152600660048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa1580156118ea573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061190e91906143eb565b1561192c5760405163840a48d560e01b815260040160405180910390fd5b6118408261277e565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461197e57604051633213a66160e21b815260040160405180910390fd5b603e54600160a01b900460ff16156119a957604051634549b2d160e11b815260040160405180910390fd5b346801bc16d674ec800000146119d25760405163049696b360e31b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663228951186801bc16d674ec8000008787611a156130af565b8888886040518863ffffffff1660e01b8152600401611a3996959493929190614603565b5f604051808303818588803b158015611a50575f5ffd5b505af1158015611a62573d5f5f3e3d5ffd5b50505050507fa01003766d3cd97cf2ade5429690bf5d206be7fb01ef9d3a0089ecf67bc11219611ac686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506129af92505050565b60405190815260200160405180910390a15050505050565b611b05604080516080810182525f8082526020820181905291810182905290606082015290565b60365f611b4685858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506129af92505050565b815260208082019290925260409081015f20815160808101835281546001600160401b038082168352600160401b8204811695830195909552600160801b81049094169281019290925290916060830190600160c01b900460ff166002811115611bb257611bb2614034565b6002811115611bc357611bc3614034565b9052509392505050565b5f610a64710961ef480eb55e80d19ad83579a64c007002612929565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611c3257604051633213a66160e21b815260040160405180910390fd5b603e54600160a01b900460ff1615611c5d57604051634549b2d160e11b815260040160405180910390fd5b5f611c6c633b9aca0083614665565b9050611c85633b9aca006001600160401b0383166144a1565b6034549092506001600160401b039081169082161115611cb8576040516302c6f54760e21b815260040160405180910390fd5b603480548291905f90611cd59084906001600160401b0316614678565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550826001600160a01b03167f8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e83604051611d3491815260200190565b60405180910390a2611d468383612a40565b505050565b5f54610100900460ff1615808015611d6957505f54600160ff909116105b80611d825750303b158015611d8257505f5460ff166001145b611dea5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015611e0b575f805461ff0019166101001790555b6001600160a01b038216611e32576040516339b190bb60e11b815260040160405180910390fd5b603380546001600160a01b0319166001600160a01b0384161790558015611840575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6033546001600160a01b03163314611ec05760405163719f370360e11b815260040160405180910390fd5b603e54604080516001600160a01b03928316815291831660208301527ffb8129080a19d34dceac04ba253fc50304dc86c729bd63cdca4a969ad19a5eac910160405180910390a1603e80546001600160a01b0319166001600160a01b0392909216919091179055565b6033546001600160a01b03163314611f545760405163719f370360e11b815260040160405180910390fd5b604051635ac86ab760e01b8152600560048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa158015611fba573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611fde91906143eb565b15611ffc5760405163840a48d560e01b815260040160405180910390fd5b825184511461201e576040516343714afd60e01b815260040160405180910390fd5b5f5b8451811015612081576120798385838151811061203f5761203f6144cb565b6020026020010151878481518110612059576120596144cb565b60200260200101516001600160a01b03166130df9092919063ffffffff16565b600101612020565b5050505050565b604051635ac86ab760e01b8152600760048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635ac86ab790602401602060405180830381865afa1580156120ee573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061211291906143eb565b156121305760405163840a48d560e01b815260040160405180910390fd5b603a54600160401b90046001600160401b03165f81900361216457604051631a544f4960e01b815260040160405180910390fd5b6040805160a081018252603c548152603d5462ffffff811660208301526001600160401b0363010000008204811693830193909352600160581b810460070b6060830152600160981b900490911660808201526121cb6121c383612573565b825188613131565b5f805b8581101561242b57368787838181106121e9576121e96144cb565b90506020028101906121fb9190614697565b80355f908152603660209081526040808320815160808101835281546001600160401b038082168352600160401b8204811695830195909552600160801b8104909416928101929092529394509192906060830190600160c01b900460ff16600281111561226b5761226b614034565b600281111561227c5761227c614034565b905250905060018160600151600281111561229957612299614034565b146122a5575050612423565b856001600160401b031681604001516001600160401b0316106122c9575050612423565b5f80806122d9848a8f35886131f2565b60208b01805193965091945092506122f0826146ab565b62ffffff1690525060808801805184919061230c9083906146c8565b6001600160401b031690525060608801805183919061232c9083906146e7565b60070b90525061233c81886146c8565b85355f908152603660209081526040918290208751815492890151938901516001600160401b03908116600160801b0267ffffffffffffffff60801b19958216600160401b026001600160801b0319909516919092161792909217928316821781556060880151939a50879390929091839160ff60c01b1990911668ffffffffffffffffff60801b1990911617600160c01b8360028111156123e0576123e0614034565b021790555050604051863591506001600160401b038b16907fe4866335761a51dcaff766448ab0af6064291ee5dc94e68492bb9cd757c1e350905f90a350505050505b6001016121ce565b506001600160401b038084165f908152603b6020526040812080548493919291612457918591166146c8565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550610a3f82613308565b5f815f81518110612497576124976144cb565b60200260200101519050919050565b5f816003815181106124ba576124ba6144cb565b60200260200101515f5f1b14159050919050565b6124da600360206144a1565b6124e7602083018361444b565b905014612507576040516313717da960e21b815260040160405180910390fd5b612556612517602083018361444b565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525086925050843590506003613528565b611840576040516309bde33960e01b815260040160405180910390fd5b5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632704351a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156125d1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125f59190614716565b9050806001600160401b03165f03612620576040516341a02cc960e01b815260040160405180910390fd5b806001600160401b0316836001600160401b03161115612641576001612643565b5f5b9392505050565b6008841461266b5760405163200591bd60e01b815260040160405180910390fd5b5f6126758861355d565b90508061268460286001614564565b61268e9190614564565b6126999060206144a1565b83146126b8576040516313717da960e21b815260040160405180910390fd5b5f6126f48787808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525061358592505050565b90505f64ffffffffff841661270b60286001614564565b600b901b17905061275586868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508d9250869150859050613528565b612772576040516309bde33960e01b815260040160405180910390fd5b50505050505050505050565b603e54600160a01b900460ff16156127a957604051634549b2d160e11b815260040160405180910390fd5b603a54600160401b90046001600160401b0316156127d95760405162be9bc360e81b815260040160405180910390fd5b603a546001600160401b03428116911603612807576040516367db5b8b60e01b815260040160405180910390fd5b6034545f906001600160401b0316612823633b9aca0047614665565b61282d9190614678565b905081801561284357506001600160401b038116155b15612861576040516332dea95960e21b815260040160405180910390fd5b5f6040518060a0016040528061287642611507565b815260395462ffffff1660208201526001600160401b0380851660408301525f60608301819052608090920191909152603a805442909216600160401b026fffffffffffffffff00000000000000001990921691909117905590506128da81613308565b805160208083015160405162ffffff90911681526001600160401b034216917f575796133bbed337e5b39aa49a30dc2556a91e0c6c2af4b7b886ae77ebef1076910160405180910390a3505050565b5f5f5f836001600160a01b03166040515f60405180830381855afa9150503d805f8114612971576040519150601f19603f3d011682016040523d82523d5f602084013e612976565b606091505b5091509150818015612989575080516020145b6129a65760405163c90158af60e01b815260040160405180910390fd5b61160d81614731565b5f81516030146129d257604051634f88323960e11b815260040160405180910390fd5b6040516002906129e89084905f90602001614754565b60408051601f1981840301815290829052612a0291614535565b602060405180830381855afa158015612a1d573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906112149190614596565b80471015612a905760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401611de1565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114612ad9576040519150601f19603f3d011682016040523d82523d5f602084013e612ade565b606091505b5050905080611d465760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401611de1565b5f5f612b928484808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525061248492505050565b5f818152603660209081526040808320815160808101835281546001600160401b038082168352600160401b8204811695830195909552600160801b8104909416928101929092529394509192906060830190600160c01b900460ff166002811115612c0057612c00614034565b6002811115612c1157612c11614034565b90525090505f81606001516002811115612c2d57612c2d614034565b14612c4b576040516335e09e9d60e01b815260040160405180910390fd5b6001600160401b038016612c908686808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525061385392505050565b6001600160401b031603612cb757604051631958236d60e21b815260040160405180910390fd5b6001600160401b038016612cfc8686808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525061387792505050565b6001600160401b031614612d2357604051632eade63760e01b815260040160405180910390fd5b612d2b6130af565b612d3490614731565b612d6f8686808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525061388e92505050565b1480612dc45750612d7e6138a2565b612d8790614731565b612dc28686808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525061388e92505050565b145b612de157604051633772dd5360e11b815260040160405180910390fd5b5f612e1d8686808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152506138be92505050565b9050612e36612e2b8c612573565b8b88888c8c8f61264a565b60398054905f612e4583614778565b9091555050603a545f90600160401b90046001600160401b031615612e7c57603a54600160401b90046001600160401b0316612e89565b603a546001600160401b03165b6040805160808101825264ffffffffff8d1681526001600160401b03858116602083015283169181019190915290915060608101600190525f858152603660209081526040918290208351815492850151938501516001600160401b03908116600160801b0267ffffffffffffffff60801b19958216600160401b026001600160801b031990951691909216179290921792831682178155606084015190929091839160ff60c01b1990911668ffffffffffffffffff60801b1990911617600160c01b836002811115612f5e57612f5e614034565b021790555050603d8054849250601390612f89908490600160981b90046001600160401b03166146c8565b92506101000a8154816001600160401b0302191690836001600160401b031602179055507f101790c2993f6a4d962bd17c786126823ba1c4cf04ff4cccb2659d50fb20aee884604051612fde91815260200190565b60405180910390a1604080518581526001600160401b03838116602083015284168183015290517fcdae700d7241bc027168c53cf6f889763b0a2c88a65d77fc13a8a9fef0d8605f9181900360600190a1613046633b9aca006001600160401b0384166144a1565b9c9b505050505050505050505050565b6002600154036130a85760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611de1565b6002600155565b6040516060906130cb90600160f81b905f903090602001614790565b604051602081830303815290604052905090565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052611d469084906138d5565b5f61313b8461355d565b9050613148816003614564565b6131539060206144a1565b613160602084018461444b565b905014613180576040516313717da960e21b815260040160405180910390fd5b6003811b600c176131d5613197602085018561444b565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250508635905084613528565b612081576040516309bde33960e01b815260040160405180910390fd5b5f5f5f866020015192505f61320b86895f0151876139a8565b9050836001600160401b0316816001600160401b03161461327f5761323084826147ce565b60408051873581526001600160401b038a8116602083015284168183015290519194507fcdae700d7241bc027168c53cf6f889763b0a2c88a65d77fc13a8a9fef0d8605f919081900360600190a15b6001600160401b0380821660208a0181905290881660408a01525f036132fd5760398054905f6132ae836147fd565b9091555050600260608901526132c383614812565b6040519092508535906001600160401b038916907f5ce0aa04ae51d52da6e680fbe0336d2e2432f7c3dc2d4f3193204c57b9072107905f90a35b509450945094915050565b8051603c556020810151603d805460408401516060850151608086015162ffffff9095166affffffffffffffffffffff19909316831763010000006001600160401b0393841602176fffffffffffffffffffffffffffffffff60581b1916600160581b9183169190910267ffffffffffffffff60981b191617600160981b9190941602929092179055156133995750565b60808101516034545f916133b5916001600160401b03166146c8565b90505f826060015183604001516133cc91906146e7565b60408401516034805492935090915f906133f09084906001600160401b03166146c8565b82546101009290920a6001600160401b03818102199093169183160217909155603a8054600160401b810483166001600160801b03199091161790555f915061344090633b9aca009085166144a1565b90505f613455633b9aca00600785900b614837565b603a546040518281529192506001600160401b0316907f525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e449060200160405180910390a260335460405163a1ca780b60e01b81526001600160a01b03918216600482015260248101849052604481018390527f00000000000000000000000000000000000000000000000000000000000000009091169063a1ca780b906064015f604051808303815f87803b15801561350b575f5ffd5b505af115801561351d573d5f5f3e3d5ffd5b505050505050505050565b5f83613547576040516329e7276760e11b815260040160405180910390fd5b83613553868585613a86565b1495945050505050565b5f8082600181111561357157613571614034565b1461357d576006611214565b600592915050565b5f60018251116135a85760405163f8ef036760e01b815260040160405180910390fd5b6135b28251613b79565b6135cf5760405163f6558f5160e01b815260040160405180910390fd5b5f600283516135de9190614665565b90505f816001600160401b038111156135f9576135f96141e2565b604051908082528060200260200182016040528015613622578160200160208202803683370190505b5090505f5b8281101561371c5760028561363c83836144a1565b8151811061364c5761364c6144cb565b60200260200101518683600261366291906144a1565b61366d906001614564565b8151811061367d5761367d6144cb565b602002602001015160405160200161369f929190918252602082015260400190565b60408051601f19818403018152908290526136b991614535565b602060405180830381855afa1580156136d4573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906136f79190614596565b828281518110613709576137096144cb565b6020908102919091010152600101613627565b505b8160011461383057613731600283614665565b91505f5b8281101561382a5760028261374a83836144a1565b8151811061375a5761375a6144cb565b60200260200101518383600261377091906144a1565b61377b906001614564565b8151811061378b5761378b6144cb565b60200260200101516040516020016137ad929190918252602082015260400190565b60408051601f19818403018152908290526137c791614535565b602060405180830381855afa1580156137e2573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906138059190614596565b828281518110613817576138176144cb565b6020908102919091010152600101613735565b5061371e565b805f81518110613842576138426144cb565b602002602001015192505050919050565b5f6112148260058151811061386a5761386a6144cb565b6020026020010151613b98565b5f6112148260068151811061386a5761386a6144cb565b5f81600181518110612497576124976144cb565b6040516060906130cb90600160f91b905f903090602001614790565b5f6112148260028151811061386a5761386a6144cb565b5f613929826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613bff9092919063ffffffff16565b905080515f148061394957508080602001905181019061394991906143eb565b611d465760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401611de1565b5f6139b560266001614564565b6139c09060206144a1565b6139cd604084018461444b565b9050146139ed576040516313717da960e21b815260040160405180910390fd5b5f6139f9600485614866565b64ffffffffff169050613a52613a12604085018561444b565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508992505050602086013584613528565b613a6f576040516309bde33960e01b815260040160405180910390fd5b613a7d836020013585613c0d565b95945050505050565b5f83515f14158015613aa3575060208451613aa1919061488f565b155b613ac0576040516313717da960e21b815260040160405180910390fd5b604080516020808201909252848152905b85518111613b5057613ae460028561488f565b5f03613b165781515f528086015160205260208260405f60026107d05a03fa613b0b575f5ffd5b600284049350613b3e565b808601515f52815160205260208260405f60026107d05a03fa613b37575f5ffd5b6002840493505b613b49602082614564565b9050613ad1565b508215613b70576040516363df817160e01b815260040160405180910390fd5b51949350505050565b5f81158015906112145750613b8f6001836144b8565b82161592915050565b60f881901c60e882901c61ff00161760d882901c62ff0000161760c882901c63ff000000161764ff0000000060b883901c161765ff000000000060a883901c161766ff000000000000609883901c161767ff0000000000000060889290921c919091161790565b606061160d84845f85613c39565b5f80613c1a6004846148a2565b613c259060406148cb565b64ffffffffff16905061160d84821b613b98565b606082471015613c9a5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401611de1565b5f5f866001600160a01b03168587604051613cb59190614535565b5f6040518083038185875af1925050503d805f8114613cef576040519150601f19603f3d011682016040523d82523d5f602084013e613cf4565b606091505b5091509150613d0587838387613d10565b979650505050505050565b60608315613d7e5782515f03613d77576001600160a01b0385163b613d775760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611de1565b508161160d565b61160d8383815115613d935781518083602001fd5b8060405162461bcd60e51b8152600401611de191906148eb565b6001600160401b0381168114613dc1575f5ffd5b50565b5f60408284031215613dd4575f5ffd5b50919050565b5f5f5f60608486031215613dec575f5ffd5b8335613df781613dad565b925060208401356001600160401b03811115613e11575f5ffd5b613e1d86828701613dc4565b92505060408401356001600160401b03811115613e38575f5ffd5b613e4486828701613dc4565b9150509250925092565b5f5f83601f840112613e5e575f5ffd5b5081356001600160401b03811115613e74575f5ffd5b6020830191508360208260051b8501011115613e8e575f5ffd5b9250929050565b5f5f60208385031215613ea6575f5ffd5b82356001600160401b03811115613ebb575f5ffd5b613ec785828601613e4e565b90969095509350505050565b5f5f5f5f5f5f5f5f60a0898b031215613eea575f5ffd5b8835613ef581613dad565b975060208901356001600160401b03811115613f0f575f5ffd5b613f1b8b828c01613dc4565b97505060408901356001600160401b03811115613f36575f5ffd5b613f428b828c01613e4e565b90975095505060608901356001600160401b03811115613f60575f5ffd5b613f6c8b828c01613e4e565b90955093505060808901356001600160401b03811115613f8a575f5ffd5b613f968b828c01613e4e565b999c989b5096995094979396929594505050565b5f60208284031215613fba575f5ffd5b813561264381613dad565b5f5f83601f840112613fd5575f5ffd5b5081356001600160401b03811115613feb575f5ffd5b602083019150836020828501011115613e8e575f5ffd5b5f5f60208385031215614013575f5ffd5b82356001600160401b03811115614028575f5ffd5b613ec785828601613fc5565b634e487b7160e01b5f52602160045260245ffd5b6003811061406457634e487b7160e01b5f52602160045260245ffd5b9052565b602081016112148284614048565b5f60208284031215614086575f5ffd5b5035919050565b5f6080820190506001600160401b0383511682526001600160401b0360208401511660208301526001600160401b03604084015116604083015260608301516140d96060840182614048565b5092915050565b6001600160a01b0381168114613dc1575f5ffd5b80356140ff816140e0565b919050565b5f60208284031215614114575f5ffd5b8135612643816140e0565b8015158114613dc1575f5ffd5b5f6020828403121561413c575f5ffd5b81356126438161411f565b5f5f5f5f5f6060868803121561415b575f5ffd5b85356001600160401b03811115614170575f5ffd5b61417c88828901613fc5565b90965094505060208601356001600160401b0381111561419a575f5ffd5b6141a688828901613fc5565b96999598509660400135949350505050565b5f5f604083850312156141c9575f5ffd5b82356141d4816140e0565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561421e5761421e6141e2565b604052919050565b5f6001600160401b0382111561423e5761423e6141e2565b5060051b60200190565b5f82601f830112614257575f5ffd5b813561426a61426582614226565b6141f6565b8082825260208201915060208360051b86010192508583111561428b575f5ffd5b602085015b838110156142a8578035835260209283019201614290565b5095945050505050565b5f5f5f606084860312156142c4575f5ffd5b83356001600160401b038111156142d9575f5ffd5b8401601f810186136142e9575f5ffd5b80356142f761426582614226565b8082825260208201915060208360051b850101925088831115614318575f5ffd5b6020840193505b82841015614343578335614332816140e0565b82526020938401939091019061431f565b955050505060208401356001600160401b03811115614360575f5ffd5b61436c86828701614248565b92505061437b604085016140f4565b90509250925092565b5f5f5f60408486031215614396575f5ffd5b83356001600160401b038111156143ab575f5ffd5b6143b786828701613dc4565b93505060208401356001600160401b038111156143d2575f5ffd5b6143de86828701613e4e565b9497909650939450505050565b5f602082840312156143fb575f5ffd5b81516126438161411f565b5f5f8335601e1984360301811261441b575f5ffd5b8301803591506001600160401b03821115614434575f5ffd5b6020019150600581901b3603821315613e8e575f5ffd5b5f5f8335601e19843603018112614460575f5ffd5b8301803591506001600160401b03821115614479575f5ffd5b602001915036819003821315613e8e575f5ffd5b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176112145761121461448d565b818103818111156112145761121461448d565b634e487b7160e01b5f52603260045260245ffd5b5f8235603e198336030181126144f3575f5ffd5b9190910192915050565b8284823760c09190911b6001600160c01b0319169101908152600801919050565b5f81518060208401855e5f93019283525090919050565b5f612643828461451e565b5f60208284031215614550575f5ffd5b813564ffffffffff81168114612643575f5ffd5b808201808211156112145761121461448d565b838582375f8482015f8152838582375f93019283525090949350505050565b5f602082840312156145a6575f5ffd5b5051919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b608081525f61461660808301888a6145ad565b828103602084015261462881886145d5565b9050828103604084015261463d8186886145ad565b915050826060830152979650505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f8261467357614673614651565b500490565b6001600160401b0382811682821603908111156112145761121461448d565b5f8235605e198336030181126144f3575f5ffd5b5f62ffffff8216806146bf576146bf61448d565b5f190192915050565b6001600160401b0381811683821601908111156112145761121461448d565b600781810b9083900b01677fffffffffffffff8113677fffffffffffffff19821217156112145761121461448d565b5f60208284031215614726575f5ffd5b815161264381613dad565b80516020808301519190811015613dd4575f1960209190910360031b1b16919050565b5f61475f828561451e565b6001600160801b03199390931683525050601001919050565b5f600182016147895761478961448d565b5060010190565b6001600160f81b03199390931683526001600160a81b031991909116600183015260601b6bffffffffffffffffffffffff1916600c82015260200190565b600782810b9082900b03677fffffffffffffff198112677fffffffffffffff821317156112145761121461448d565b5f8161480b5761480b61448d565b505f190190565b5f8160070b677fffffffffffffff19810361482f5761482f61448d565b5f0392915050565b8082025f8212600160ff1b841416156148525761485261448d565b81810583148215176112145761121461448d565b5f64ffffffffff83168061487c5761487c614651565b8064ffffffffff84160491505092915050565b5f8261489d5761489d614651565b500690565b5f64ffffffffff8316806148b8576148b8614651565b8064ffffffffff84160691505092915050565b64ffffffffff81811683821602908116908181146140d9576140d961448d565b602081525f61264360208301846145d556fea2646970667358221220a168ee7c11a1ba7d7a4465793dd362f00459d7dc456364cd2d27bd5de9bdeb2464736f6c634300081e0033",
}

// EigenPodABI is the input ABI used to generate the binding from.
// Deprecated: Use EigenPodMetaData.ABI instead.
var EigenPodABI = EigenPodMetaData.ABI

// EigenPodBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EigenPodMetaData.Bin instead.
var EigenPodBin = EigenPodMetaData.Bin

// DeployEigenPod deploys a new Ethereum contract, binding an instance of EigenPod to it.
func DeployEigenPod(auth *bind.TransactOpts, backend bind.ContractBackend, _ethPOS common.Address, _eigenPodManager common.Address) (common.Address, *types.Transaction, *EigenPod, error) {
	parsed, err := EigenPodMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EigenPodBin), backend, _ethPOS, _eigenPodManager)
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
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int64,uint64))
func (_EigenPod *EigenPodCaller) CurrentCheckpoint(opts *bind.CallOpts) (IEigenPodTypesCheckpoint, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "currentCheckpoint")

	if err != nil {
		return *new(IEigenPodTypesCheckpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodTypesCheckpoint)).(*IEigenPodTypesCheckpoint)

	return out0, err

}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int64,uint64))
func (_EigenPod *EigenPodSession) CurrentCheckpoint() (IEigenPodTypesCheckpoint, error) {
	return _EigenPod.Contract.CurrentCheckpoint(&_EigenPod.CallOpts)
}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int64,uint64))
func (_EigenPod *EigenPodCallerSession) CurrentCheckpoint() (IEigenPodTypesCheckpoint, error) {
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

// GetConsolidationRequestFee is a free data retrieval call binding the contract method 0x1e515533.
//
// Solidity: function getConsolidationRequestFee() view returns(uint256)
func (_EigenPod *EigenPodCaller) GetConsolidationRequestFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "getConsolidationRequestFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsolidationRequestFee is a free data retrieval call binding the contract method 0x1e515533.
//
// Solidity: function getConsolidationRequestFee() view returns(uint256)
func (_EigenPod *EigenPodSession) GetConsolidationRequestFee() (*big.Int, error) {
	return _EigenPod.Contract.GetConsolidationRequestFee(&_EigenPod.CallOpts)
}

// GetConsolidationRequestFee is a free data retrieval call binding the contract method 0x1e515533.
//
// Solidity: function getConsolidationRequestFee() view returns(uint256)
func (_EigenPod *EigenPodCallerSession) GetConsolidationRequestFee() (*big.Int, error) {
	return _EigenPod.Contract.GetConsolidationRequestFee(&_EigenPod.CallOpts)
}

// GetParentBlockRoot is a free data retrieval call binding the contract method 0x6c0d2d5a.
//
// Solidity: function getParentBlockRoot(uint64 timestamp) view returns(bytes32)
func (_EigenPod *EigenPodCaller) GetParentBlockRoot(opts *bind.CallOpts, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "getParentBlockRoot", timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetParentBlockRoot is a free data retrieval call binding the contract method 0x6c0d2d5a.
//
// Solidity: function getParentBlockRoot(uint64 timestamp) view returns(bytes32)
func (_EigenPod *EigenPodSession) GetParentBlockRoot(timestamp uint64) ([32]byte, error) {
	return _EigenPod.Contract.GetParentBlockRoot(&_EigenPod.CallOpts, timestamp)
}

// GetParentBlockRoot is a free data retrieval call binding the contract method 0x6c0d2d5a.
//
// Solidity: function getParentBlockRoot(uint64 timestamp) view returns(bytes32)
func (_EigenPod *EigenPodCallerSession) GetParentBlockRoot(timestamp uint64) ([32]byte, error) {
	return _EigenPod.Contract.GetParentBlockRoot(&_EigenPod.CallOpts, timestamp)
}

// GetWithdrawalRequestFee is a free data retrieval call binding the contract method 0xc44e30dc.
//
// Solidity: function getWithdrawalRequestFee() view returns(uint256)
func (_EigenPod *EigenPodCaller) GetWithdrawalRequestFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "getWithdrawalRequestFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalRequestFee is a free data retrieval call binding the contract method 0xc44e30dc.
//
// Solidity: function getWithdrawalRequestFee() view returns(uint256)
func (_EigenPod *EigenPodSession) GetWithdrawalRequestFee() (*big.Int, error) {
	return _EigenPod.Contract.GetWithdrawalRequestFee(&_EigenPod.CallOpts)
}

// GetWithdrawalRequestFee is a free data retrieval call binding the contract method 0xc44e30dc.
//
// Solidity: function getWithdrawalRequestFee() view returns(uint256)
func (_EigenPod *EigenPodCallerSession) GetWithdrawalRequestFee() (*big.Int, error) {
	return _EigenPod.Contract.GetWithdrawalRequestFee(&_EigenPod.CallOpts)
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

// ProofSubmitter is a free data retrieval call binding the contract method 0x58753357.
//
// Solidity: function proofSubmitter() view returns(address)
func (_EigenPod *EigenPodCaller) ProofSubmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "proofSubmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProofSubmitter is a free data retrieval call binding the contract method 0x58753357.
//
// Solidity: function proofSubmitter() view returns(address)
func (_EigenPod *EigenPodSession) ProofSubmitter() (common.Address, error) {
	return _EigenPod.Contract.ProofSubmitter(&_EigenPod.CallOpts)
}

// ProofSubmitter is a free data retrieval call binding the contract method 0x58753357.
//
// Solidity: function proofSubmitter() view returns(address)
func (_EigenPod *EigenPodCallerSession) ProofSubmitter() (common.Address, error) {
	return _EigenPod.Contract.ProofSubmitter(&_EigenPod.CallOpts)
}

// RestakingDisabled is a free data retrieval call binding the contract method 0xc452eeec.
//
// Solidity: function restakingDisabled() view returns(bool)
func (_EigenPod *EigenPodCaller) RestakingDisabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "restakingDisabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RestakingDisabled is a free data retrieval call binding the contract method 0xc452eeec.
//
// Solidity: function restakingDisabled() view returns(bool)
func (_EigenPod *EigenPodSession) RestakingDisabled() (bool, error) {
	return _EigenPod.Contract.RestakingDisabled(&_EigenPod.CallOpts)
}

// RestakingDisabled is a free data retrieval call binding the contract method 0xc452eeec.
//
// Solidity: function restakingDisabled() view returns(bool)
func (_EigenPod *EigenPodCallerSession) RestakingDisabled() (bool, error) {
	return _EigenPod.Contract.RestakingDisabled(&_EigenPod.CallOpts)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_EigenPod *EigenPodCaller) ValidatorPubkeyHashToInfo(opts *bind.CallOpts, validatorPubkeyHash [32]byte) (IEigenPodTypesValidatorInfo, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "validatorPubkeyHashToInfo", validatorPubkeyHash)

	if err != nil {
		return *new(IEigenPodTypesValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodTypesValidatorInfo)).(*IEigenPodTypesValidatorInfo)

	return out0, err

}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_EigenPod *EigenPodSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodTypesValidatorInfo, error) {
	return _EigenPod.Contract.ValidatorPubkeyHashToInfo(&_EigenPod.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_EigenPod *EigenPodCallerSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodTypesValidatorInfo, error) {
	return _EigenPod.Contract.ValidatorPubkeyHashToInfo(&_EigenPod.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_EigenPod *EigenPodCaller) ValidatorPubkeyToInfo(opts *bind.CallOpts, validatorPubkey []byte) (IEigenPodTypesValidatorInfo, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "validatorPubkeyToInfo", validatorPubkey)

	if err != nil {
		return *new(IEigenPodTypesValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodTypesValidatorInfo)).(*IEigenPodTypesValidatorInfo)

	return out0, err

}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_EigenPod *EigenPodSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodTypesValidatorInfo, error) {
	return _EigenPod.Contract.ValidatorPubkeyToInfo(&_EigenPod.CallOpts, validatorPubkey)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_EigenPod *EigenPodCallerSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodTypesValidatorInfo, error) {
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

// DisableRestaking is a paid mutator transaction binding the contract method 0x566f95fa.
//
// Solidity: function disableRestaking() returns()
func (_EigenPod *EigenPodTransactor) DisableRestaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "disableRestaking")
}

// DisableRestaking is a paid mutator transaction binding the contract method 0x566f95fa.
//
// Solidity: function disableRestaking() returns()
func (_EigenPod *EigenPodSession) DisableRestaking() (*types.Transaction, error) {
	return _EigenPod.Contract.DisableRestaking(&_EigenPod.TransactOpts)
}

// DisableRestaking is a paid mutator transaction binding the contract method 0x566f95fa.
//
// Solidity: function disableRestaking() returns()
func (_EigenPod *EigenPodTransactorSession) DisableRestaking() (*types.Transaction, error) {
	return _EigenPod.Contract.DisableRestaking(&_EigenPod.TransactOpts)
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

// RequestConsolidation is a paid mutator transaction binding the contract method 0x6691954e.
//
// Solidity: function requestConsolidation((bytes,bytes)[] requests) payable returns()
func (_EigenPod *EigenPodTransactor) RequestConsolidation(opts *bind.TransactOpts, requests []IEigenPodTypesConsolidationRequest) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "requestConsolidation", requests)
}

// RequestConsolidation is a paid mutator transaction binding the contract method 0x6691954e.
//
// Solidity: function requestConsolidation((bytes,bytes)[] requests) payable returns()
func (_EigenPod *EigenPodSession) RequestConsolidation(requests []IEigenPodTypesConsolidationRequest) (*types.Transaction, error) {
	return _EigenPod.Contract.RequestConsolidation(&_EigenPod.TransactOpts, requests)
}

// RequestConsolidation is a paid mutator transaction binding the contract method 0x6691954e.
//
// Solidity: function requestConsolidation((bytes,bytes)[] requests) payable returns()
func (_EigenPod *EigenPodTransactorSession) RequestConsolidation(requests []IEigenPodTypesConsolidationRequest) (*types.Transaction, error) {
	return _EigenPod.Contract.RequestConsolidation(&_EigenPod.TransactOpts, requests)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x3f5fa57a.
//
// Solidity: function requestWithdrawal((bytes,uint64)[] requests) payable returns()
func (_EigenPod *EigenPodTransactor) RequestWithdrawal(opts *bind.TransactOpts, requests []IEigenPodTypesWithdrawalRequest) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "requestWithdrawal", requests)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x3f5fa57a.
//
// Solidity: function requestWithdrawal((bytes,uint64)[] requests) payable returns()
func (_EigenPod *EigenPodSession) RequestWithdrawal(requests []IEigenPodTypesWithdrawalRequest) (*types.Transaction, error) {
	return _EigenPod.Contract.RequestWithdrawal(&_EigenPod.TransactOpts, requests)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x3f5fa57a.
//
// Solidity: function requestWithdrawal((bytes,uint64)[] requests) payable returns()
func (_EigenPod *EigenPodTransactorSession) RequestWithdrawal(requests []IEigenPodTypesWithdrawalRequest) (*types.Transaction, error) {
	return _EigenPod.Contract.RequestWithdrawal(&_EigenPod.TransactOpts, requests)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address newProofSubmitter) returns()
func (_EigenPod *EigenPodTransactor) SetProofSubmitter(opts *bind.TransactOpts, newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "setProofSubmitter", newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address newProofSubmitter) returns()
func (_EigenPod *EigenPodSession) SetProofSubmitter(newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EigenPod.Contract.SetProofSubmitter(&_EigenPod.TransactOpts, newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address newProofSubmitter) returns()
func (_EigenPod *EigenPodTransactorSession) SetProofSubmitter(newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EigenPod.Contract.SetProofSubmitter(&_EigenPod.TransactOpts, newProofSubmitter)
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

// WithdrawDisabledPodETH is a paid mutator transaction binding the contract method 0x8842752c.
//
// Solidity: function withdrawDisabledPodETH(address recipient) returns()
func (_EigenPod *EigenPodTransactor) WithdrawDisabledPodETH(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "withdrawDisabledPodETH", recipient)
}

// WithdrawDisabledPodETH is a paid mutator transaction binding the contract method 0x8842752c.
//
// Solidity: function withdrawDisabledPodETH(address recipient) returns()
func (_EigenPod *EigenPodSession) WithdrawDisabledPodETH(recipient common.Address) (*types.Transaction, error) {
	return _EigenPod.Contract.WithdrawDisabledPodETH(&_EigenPod.TransactOpts, recipient)
}

// WithdrawDisabledPodETH is a paid mutator transaction binding the contract method 0x8842752c.
//
// Solidity: function withdrawDisabledPodETH(address recipient) returns()
func (_EigenPod *EigenPodTransactorSession) WithdrawDisabledPodETH(recipient common.Address) (*types.Transaction, error) {
	return _EigenPod.Contract.WithdrawDisabledPodETH(&_EigenPod.TransactOpts, recipient)
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
	ValidatorCount      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCheckpointCreated is a free log retrieval operation binding the contract event 0x575796133bbed337e5b39aa49a30dc2556a91e0c6c2af4b7b886ae77ebef1076.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot, uint256 validatorCount)
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

// WatchCheckpointCreated is a free log subscription operation binding the contract event 0x575796133bbed337e5b39aa49a30dc2556a91e0c6c2af4b7b886ae77ebef1076.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot, uint256 validatorCount)
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

// ParseCheckpointCreated is a log parse operation binding the contract event 0x575796133bbed337e5b39aa49a30dc2556a91e0c6c2af4b7b886ae77ebef1076.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot, uint256 validatorCount)
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

// EigenPodConsolidationRequestedIterator is returned from FilterConsolidationRequested and is used to iterate over the raw logs and unpacked data for ConsolidationRequested events raised by the EigenPod contract.
type EigenPodConsolidationRequestedIterator struct {
	Event *EigenPodConsolidationRequested // Event containing the contract specifics and raw log

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
func (it *EigenPodConsolidationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodConsolidationRequested)
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
		it.Event = new(EigenPodConsolidationRequested)
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
func (it *EigenPodConsolidationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodConsolidationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodConsolidationRequested represents a ConsolidationRequested event raised by the EigenPod contract.
type EigenPodConsolidationRequested struct {
	SourcePubkeyHash [32]byte
	TargetPubkeyHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConsolidationRequested is a free log retrieval operation binding the contract event 0x42f9c9db2ca443e9ec62f4588bd0c9b241065c02c2a8001ac164ae1282dc7b94.
//
// Solidity: event ConsolidationRequested(bytes32 indexed sourcePubkeyHash, bytes32 indexed targetPubkeyHash)
func (_EigenPod *EigenPodFilterer) FilterConsolidationRequested(opts *bind.FilterOpts, sourcePubkeyHash [][32]byte, targetPubkeyHash [][32]byte) (*EigenPodConsolidationRequestedIterator, error) {

	var sourcePubkeyHashRule []interface{}
	for _, sourcePubkeyHashItem := range sourcePubkeyHash {
		sourcePubkeyHashRule = append(sourcePubkeyHashRule, sourcePubkeyHashItem)
	}
	var targetPubkeyHashRule []interface{}
	for _, targetPubkeyHashItem := range targetPubkeyHash {
		targetPubkeyHashRule = append(targetPubkeyHashRule, targetPubkeyHashItem)
	}

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "ConsolidationRequested", sourcePubkeyHashRule, targetPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodConsolidationRequestedIterator{contract: _EigenPod.contract, event: "ConsolidationRequested", logs: logs, sub: sub}, nil
}

// WatchConsolidationRequested is a free log subscription operation binding the contract event 0x42f9c9db2ca443e9ec62f4588bd0c9b241065c02c2a8001ac164ae1282dc7b94.
//
// Solidity: event ConsolidationRequested(bytes32 indexed sourcePubkeyHash, bytes32 indexed targetPubkeyHash)
func (_EigenPod *EigenPodFilterer) WatchConsolidationRequested(opts *bind.WatchOpts, sink chan<- *EigenPodConsolidationRequested, sourcePubkeyHash [][32]byte, targetPubkeyHash [][32]byte) (event.Subscription, error) {

	var sourcePubkeyHashRule []interface{}
	for _, sourcePubkeyHashItem := range sourcePubkeyHash {
		sourcePubkeyHashRule = append(sourcePubkeyHashRule, sourcePubkeyHashItem)
	}
	var targetPubkeyHashRule []interface{}
	for _, targetPubkeyHashItem := range targetPubkeyHash {
		targetPubkeyHashRule = append(targetPubkeyHashRule, targetPubkeyHashItem)
	}

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "ConsolidationRequested", sourcePubkeyHashRule, targetPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodConsolidationRequested)
				if err := _EigenPod.contract.UnpackLog(event, "ConsolidationRequested", log); err != nil {
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

// ParseConsolidationRequested is a log parse operation binding the contract event 0x42f9c9db2ca443e9ec62f4588bd0c9b241065c02c2a8001ac164ae1282dc7b94.
//
// Solidity: event ConsolidationRequested(bytes32 indexed sourcePubkeyHash, bytes32 indexed targetPubkeyHash)
func (_EigenPod *EigenPodFilterer) ParseConsolidationRequested(log types.Log) (*EigenPodConsolidationRequested, error) {
	event := new(EigenPodConsolidationRequested)
	if err := _EigenPod.contract.UnpackLog(event, "ConsolidationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodDisabledPodETHWithdrawnIterator is returned from FilterDisabledPodETHWithdrawn and is used to iterate over the raw logs and unpacked data for DisabledPodETHWithdrawn events raised by the EigenPod contract.
type EigenPodDisabledPodETHWithdrawnIterator struct {
	Event *EigenPodDisabledPodETHWithdrawn // Event containing the contract specifics and raw log

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
func (it *EigenPodDisabledPodETHWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodDisabledPodETHWithdrawn)
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
		it.Event = new(EigenPodDisabledPodETHWithdrawn)
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
func (it *EigenPodDisabledPodETHWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodDisabledPodETHWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodDisabledPodETHWithdrawn represents a DisabledPodETHWithdrawn event raised by the EigenPod contract.
type EigenPodDisabledPodETHWithdrawn struct {
	Recipient common.Address
	AmountWei *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDisabledPodETHWithdrawn is a free log retrieval operation binding the contract event 0x2273c420d02f2b6ef92ed2c3d60848ab18c03c0ce20d6d2cf736030838f08f46.
//
// Solidity: event DisabledPodETHWithdrawn(address indexed recipient, uint256 amountWei)
func (_EigenPod *EigenPodFilterer) FilterDisabledPodETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*EigenPodDisabledPodETHWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "DisabledPodETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodDisabledPodETHWithdrawnIterator{contract: _EigenPod.contract, event: "DisabledPodETHWithdrawn", logs: logs, sub: sub}, nil
}

// WatchDisabledPodETHWithdrawn is a free log subscription operation binding the contract event 0x2273c420d02f2b6ef92ed2c3d60848ab18c03c0ce20d6d2cf736030838f08f46.
//
// Solidity: event DisabledPodETHWithdrawn(address indexed recipient, uint256 amountWei)
func (_EigenPod *EigenPodFilterer) WatchDisabledPodETHWithdrawn(opts *bind.WatchOpts, sink chan<- *EigenPodDisabledPodETHWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "DisabledPodETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodDisabledPodETHWithdrawn)
				if err := _EigenPod.contract.UnpackLog(event, "DisabledPodETHWithdrawn", log); err != nil {
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

// ParseDisabledPodETHWithdrawn is a log parse operation binding the contract event 0x2273c420d02f2b6ef92ed2c3d60848ab18c03c0ce20d6d2cf736030838f08f46.
//
// Solidity: event DisabledPodETHWithdrawn(address indexed recipient, uint256 amountWei)
func (_EigenPod *EigenPodFilterer) ParseDisabledPodETHWithdrawn(log types.Log) (*EigenPodDisabledPodETHWithdrawn, error) {
	event := new(EigenPodDisabledPodETHWithdrawn)
	if err := _EigenPod.contract.UnpackLog(event, "DisabledPodETHWithdrawn", log); err != nil {
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
	PubkeyHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEigenPodStaked is a free log retrieval operation binding the contract event 0xa01003766d3cd97cf2ade5429690bf5d206be7fb01ef9d3a0089ecf67bc11219.
//
// Solidity: event EigenPodStaked(bytes32 pubkeyHash)
func (_EigenPod *EigenPodFilterer) FilterEigenPodStaked(opts *bind.FilterOpts) (*EigenPodEigenPodStakedIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "EigenPodStaked")
	if err != nil {
		return nil, err
	}
	return &EigenPodEigenPodStakedIterator{contract: _EigenPod.contract, event: "EigenPodStaked", logs: logs, sub: sub}, nil
}

// WatchEigenPodStaked is a free log subscription operation binding the contract event 0xa01003766d3cd97cf2ade5429690bf5d206be7fb01ef9d3a0089ecf67bc11219.
//
// Solidity: event EigenPodStaked(bytes32 pubkeyHash)
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

// ParseEigenPodStaked is a log parse operation binding the contract event 0xa01003766d3cd97cf2ade5429690bf5d206be7fb01ef9d3a0089ecf67bc11219.
//
// Solidity: event EigenPodStaked(bytes32 pubkeyHash)
func (_EigenPod *EigenPodFilterer) ParseEigenPodStaked(log types.Log) (*EigenPodEigenPodStaked, error) {
	event := new(EigenPodEigenPodStaked)
	if err := _EigenPod.contract.UnpackLog(event, "EigenPodStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodExitRequestedIterator is returned from FilterExitRequested and is used to iterate over the raw logs and unpacked data for ExitRequested events raised by the EigenPod contract.
type EigenPodExitRequestedIterator struct {
	Event *EigenPodExitRequested // Event containing the contract specifics and raw log

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
func (it *EigenPodExitRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodExitRequested)
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
		it.Event = new(EigenPodExitRequested)
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
func (it *EigenPodExitRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodExitRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodExitRequested represents a ExitRequested event raised by the EigenPod contract.
type EigenPodExitRequested struct {
	ValidatorPubkeyHash [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterExitRequested is a free log retrieval operation binding the contract event 0x60d8ca014d4765a2b8b389e25714cb1cef83b574222911a01d90c1bd69d2d320.
//
// Solidity: event ExitRequested(bytes32 indexed validatorPubkeyHash)
func (_EigenPod *EigenPodFilterer) FilterExitRequested(opts *bind.FilterOpts, validatorPubkeyHash [][32]byte) (*EigenPodExitRequestedIterator, error) {

	var validatorPubkeyHashRule []interface{}
	for _, validatorPubkeyHashItem := range validatorPubkeyHash {
		validatorPubkeyHashRule = append(validatorPubkeyHashRule, validatorPubkeyHashItem)
	}

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "ExitRequested", validatorPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodExitRequestedIterator{contract: _EigenPod.contract, event: "ExitRequested", logs: logs, sub: sub}, nil
}

// WatchExitRequested is a free log subscription operation binding the contract event 0x60d8ca014d4765a2b8b389e25714cb1cef83b574222911a01d90c1bd69d2d320.
//
// Solidity: event ExitRequested(bytes32 indexed validatorPubkeyHash)
func (_EigenPod *EigenPodFilterer) WatchExitRequested(opts *bind.WatchOpts, sink chan<- *EigenPodExitRequested, validatorPubkeyHash [][32]byte) (event.Subscription, error) {

	var validatorPubkeyHashRule []interface{}
	for _, validatorPubkeyHashItem := range validatorPubkeyHash {
		validatorPubkeyHashRule = append(validatorPubkeyHashRule, validatorPubkeyHashItem)
	}

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "ExitRequested", validatorPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodExitRequested)
				if err := _EigenPod.contract.UnpackLog(event, "ExitRequested", log); err != nil {
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

// ParseExitRequested is a log parse operation binding the contract event 0x60d8ca014d4765a2b8b389e25714cb1cef83b574222911a01d90c1bd69d2d320.
//
// Solidity: event ExitRequested(bytes32 indexed validatorPubkeyHash)
func (_EigenPod *EigenPodFilterer) ParseExitRequested(log types.Log) (*EigenPodExitRequested, error) {
	event := new(EigenPodExitRequested)
	if err := _EigenPod.contract.UnpackLog(event, "ExitRequested", log); err != nil {
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

// EigenPodProofSubmitterUpdatedIterator is returned from FilterProofSubmitterUpdated and is used to iterate over the raw logs and unpacked data for ProofSubmitterUpdated events raised by the EigenPod contract.
type EigenPodProofSubmitterUpdatedIterator struct {
	Event *EigenPodProofSubmitterUpdated // Event containing the contract specifics and raw log

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
func (it *EigenPodProofSubmitterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodProofSubmitterUpdated)
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
		it.Event = new(EigenPodProofSubmitterUpdated)
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
func (it *EigenPodProofSubmitterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodProofSubmitterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodProofSubmitterUpdated represents a ProofSubmitterUpdated event raised by the EigenPod contract.
type EigenPodProofSubmitterUpdated struct {
	PrevProofSubmitter common.Address
	NewProofSubmitter  common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProofSubmitterUpdated is a free log retrieval operation binding the contract event 0xfb8129080a19d34dceac04ba253fc50304dc86c729bd63cdca4a969ad19a5eac.
//
// Solidity: event ProofSubmitterUpdated(address prevProofSubmitter, address newProofSubmitter)
func (_EigenPod *EigenPodFilterer) FilterProofSubmitterUpdated(opts *bind.FilterOpts) (*EigenPodProofSubmitterUpdatedIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "ProofSubmitterUpdated")
	if err != nil {
		return nil, err
	}
	return &EigenPodProofSubmitterUpdatedIterator{contract: _EigenPod.contract, event: "ProofSubmitterUpdated", logs: logs, sub: sub}, nil
}

// WatchProofSubmitterUpdated is a free log subscription operation binding the contract event 0xfb8129080a19d34dceac04ba253fc50304dc86c729bd63cdca4a969ad19a5eac.
//
// Solidity: event ProofSubmitterUpdated(address prevProofSubmitter, address newProofSubmitter)
func (_EigenPod *EigenPodFilterer) WatchProofSubmitterUpdated(opts *bind.WatchOpts, sink chan<- *EigenPodProofSubmitterUpdated) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "ProofSubmitterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodProofSubmitterUpdated)
				if err := _EigenPod.contract.UnpackLog(event, "ProofSubmitterUpdated", log); err != nil {
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

// ParseProofSubmitterUpdated is a log parse operation binding the contract event 0xfb8129080a19d34dceac04ba253fc50304dc86c729bd63cdca4a969ad19a5eac.
//
// Solidity: event ProofSubmitterUpdated(address prevProofSubmitter, address newProofSubmitter)
func (_EigenPod *EigenPodFilterer) ParseProofSubmitterUpdated(log types.Log) (*EigenPodProofSubmitterUpdated, error) {
	event := new(EigenPodProofSubmitterUpdated)
	if err := _EigenPod.contract.UnpackLog(event, "ProofSubmitterUpdated", log); err != nil {
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

// EigenPodRestakingPermanentlyDisabledIterator is returned from FilterRestakingPermanentlyDisabled and is used to iterate over the raw logs and unpacked data for RestakingPermanentlyDisabled events raised by the EigenPod contract.
type EigenPodRestakingPermanentlyDisabledIterator struct {
	Event *EigenPodRestakingPermanentlyDisabled // Event containing the contract specifics and raw log

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
func (it *EigenPodRestakingPermanentlyDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodRestakingPermanentlyDisabled)
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
		it.Event = new(EigenPodRestakingPermanentlyDisabled)
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
func (it *EigenPodRestakingPermanentlyDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodRestakingPermanentlyDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodRestakingPermanentlyDisabled represents a RestakingPermanentlyDisabled event raised by the EigenPod contract.
type EigenPodRestakingPermanentlyDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRestakingPermanentlyDisabled is a free log retrieval operation binding the contract event 0x4370e346659e8dc4bc97977bf89c631877baae18e45c9050029ca10247f68a69.
//
// Solidity: event RestakingPermanentlyDisabled()
func (_EigenPod *EigenPodFilterer) FilterRestakingPermanentlyDisabled(opts *bind.FilterOpts) (*EigenPodRestakingPermanentlyDisabledIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "RestakingPermanentlyDisabled")
	if err != nil {
		return nil, err
	}
	return &EigenPodRestakingPermanentlyDisabledIterator{contract: _EigenPod.contract, event: "RestakingPermanentlyDisabled", logs: logs, sub: sub}, nil
}

// WatchRestakingPermanentlyDisabled is a free log subscription operation binding the contract event 0x4370e346659e8dc4bc97977bf89c631877baae18e45c9050029ca10247f68a69.
//
// Solidity: event RestakingPermanentlyDisabled()
func (_EigenPod *EigenPodFilterer) WatchRestakingPermanentlyDisabled(opts *bind.WatchOpts, sink chan<- *EigenPodRestakingPermanentlyDisabled) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "RestakingPermanentlyDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodRestakingPermanentlyDisabled)
				if err := _EigenPod.contract.UnpackLog(event, "RestakingPermanentlyDisabled", log); err != nil {
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

// ParseRestakingPermanentlyDisabled is a log parse operation binding the contract event 0x4370e346659e8dc4bc97977bf89c631877baae18e45c9050029ca10247f68a69.
//
// Solidity: event RestakingPermanentlyDisabled()
func (_EigenPod *EigenPodFilterer) ParseRestakingPermanentlyDisabled(log types.Log) (*EigenPodRestakingPermanentlyDisabled, error) {
	event := new(EigenPodRestakingPermanentlyDisabled)
	if err := _EigenPod.contract.UnpackLog(event, "RestakingPermanentlyDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodSwitchToCompoundingRequestedIterator is returned from FilterSwitchToCompoundingRequested and is used to iterate over the raw logs and unpacked data for SwitchToCompoundingRequested events raised by the EigenPod contract.
type EigenPodSwitchToCompoundingRequestedIterator struct {
	Event *EigenPodSwitchToCompoundingRequested // Event containing the contract specifics and raw log

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
func (it *EigenPodSwitchToCompoundingRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodSwitchToCompoundingRequested)
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
		it.Event = new(EigenPodSwitchToCompoundingRequested)
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
func (it *EigenPodSwitchToCompoundingRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodSwitchToCompoundingRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodSwitchToCompoundingRequested represents a SwitchToCompoundingRequested event raised by the EigenPod contract.
type EigenPodSwitchToCompoundingRequested struct {
	ValidatorPubkeyHash [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSwitchToCompoundingRequested is a free log retrieval operation binding the contract event 0xc97b965b92ae7fd20095fe8eb7b99f81f95f8c4adffb22a19116d8eb2846b016.
//
// Solidity: event SwitchToCompoundingRequested(bytes32 indexed validatorPubkeyHash)
func (_EigenPod *EigenPodFilterer) FilterSwitchToCompoundingRequested(opts *bind.FilterOpts, validatorPubkeyHash [][32]byte) (*EigenPodSwitchToCompoundingRequestedIterator, error) {

	var validatorPubkeyHashRule []interface{}
	for _, validatorPubkeyHashItem := range validatorPubkeyHash {
		validatorPubkeyHashRule = append(validatorPubkeyHashRule, validatorPubkeyHashItem)
	}

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "SwitchToCompoundingRequested", validatorPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodSwitchToCompoundingRequestedIterator{contract: _EigenPod.contract, event: "SwitchToCompoundingRequested", logs: logs, sub: sub}, nil
}

// WatchSwitchToCompoundingRequested is a free log subscription operation binding the contract event 0xc97b965b92ae7fd20095fe8eb7b99f81f95f8c4adffb22a19116d8eb2846b016.
//
// Solidity: event SwitchToCompoundingRequested(bytes32 indexed validatorPubkeyHash)
func (_EigenPod *EigenPodFilterer) WatchSwitchToCompoundingRequested(opts *bind.WatchOpts, sink chan<- *EigenPodSwitchToCompoundingRequested, validatorPubkeyHash [][32]byte) (event.Subscription, error) {

	var validatorPubkeyHashRule []interface{}
	for _, validatorPubkeyHashItem := range validatorPubkeyHash {
		validatorPubkeyHashRule = append(validatorPubkeyHashRule, validatorPubkeyHashItem)
	}

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "SwitchToCompoundingRequested", validatorPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodSwitchToCompoundingRequested)
				if err := _EigenPod.contract.UnpackLog(event, "SwitchToCompoundingRequested", log); err != nil {
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

// ParseSwitchToCompoundingRequested is a log parse operation binding the contract event 0xc97b965b92ae7fd20095fe8eb7b99f81f95f8c4adffb22a19116d8eb2846b016.
//
// Solidity: event SwitchToCompoundingRequested(bytes32 indexed validatorPubkeyHash)
func (_EigenPod *EigenPodFilterer) ParseSwitchToCompoundingRequested(log types.Log) (*EigenPodSwitchToCompoundingRequested, error) {
	event := new(EigenPodSwitchToCompoundingRequested)
	if err := _EigenPod.contract.UnpackLog(event, "SwitchToCompoundingRequested", log); err != nil {
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
	PubkeyHash              [32]byte
	BalanceTimestamp        uint64
	NewValidatorBalanceGwei uint64
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterValidatorBalanceUpdated is a free log retrieval operation binding the contract event 0xcdae700d7241bc027168c53cf6f889763b0a2c88a65d77fc13a8a9fef0d8605f.
//
// Solidity: event ValidatorBalanceUpdated(bytes32 pubkeyHash, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_EigenPod *EigenPodFilterer) FilterValidatorBalanceUpdated(opts *bind.FilterOpts) (*EigenPodValidatorBalanceUpdatedIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "ValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &EigenPodValidatorBalanceUpdatedIterator{contract: _EigenPod.contract, event: "ValidatorBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorBalanceUpdated is a free log subscription operation binding the contract event 0xcdae700d7241bc027168c53cf6f889763b0a2c88a65d77fc13a8a9fef0d8605f.
//
// Solidity: event ValidatorBalanceUpdated(bytes32 pubkeyHash, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
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

// ParseValidatorBalanceUpdated is a log parse operation binding the contract event 0xcdae700d7241bc027168c53cf6f889763b0a2c88a65d77fc13a8a9fef0d8605f.
//
// Solidity: event ValidatorBalanceUpdated(bytes32 pubkeyHash, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
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
	PubkeyHash          [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorCheckpointed is a free log retrieval operation binding the contract event 0xe4866335761a51dcaff766448ab0af6064291ee5dc94e68492bb9cd757c1e350.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash)
func (_EigenPod *EigenPodFilterer) FilterValidatorCheckpointed(opts *bind.FilterOpts, checkpointTimestamp []uint64, pubkeyHash [][32]byte) (*EigenPodValidatorCheckpointedIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "ValidatorCheckpointed", checkpointTimestampRule, pubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodValidatorCheckpointedIterator{contract: _EigenPod.contract, event: "ValidatorCheckpointed", logs: logs, sub: sub}, nil
}

// WatchValidatorCheckpointed is a free log subscription operation binding the contract event 0xe4866335761a51dcaff766448ab0af6064291ee5dc94e68492bb9cd757c1e350.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash)
func (_EigenPod *EigenPodFilterer) WatchValidatorCheckpointed(opts *bind.WatchOpts, sink chan<- *EigenPodValidatorCheckpointed, checkpointTimestamp []uint64, pubkeyHash [][32]byte) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "ValidatorCheckpointed", checkpointTimestampRule, pubkeyHashRule)
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

// ParseValidatorCheckpointed is a log parse operation binding the contract event 0xe4866335761a51dcaff766448ab0af6064291ee5dc94e68492bb9cd757c1e350.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash)
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
	PubkeyHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorRestaked is a free log retrieval operation binding the contract event 0x101790c2993f6a4d962bd17c786126823ba1c4cf04ff4cccb2659d50fb20aee8.
//
// Solidity: event ValidatorRestaked(bytes32 pubkeyHash)
func (_EigenPod *EigenPodFilterer) FilterValidatorRestaked(opts *bind.FilterOpts) (*EigenPodValidatorRestakedIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "ValidatorRestaked")
	if err != nil {
		return nil, err
	}
	return &EigenPodValidatorRestakedIterator{contract: _EigenPod.contract, event: "ValidatorRestaked", logs: logs, sub: sub}, nil
}

// WatchValidatorRestaked is a free log subscription operation binding the contract event 0x101790c2993f6a4d962bd17c786126823ba1c4cf04ff4cccb2659d50fb20aee8.
//
// Solidity: event ValidatorRestaked(bytes32 pubkeyHash)
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

// ParseValidatorRestaked is a log parse operation binding the contract event 0x101790c2993f6a4d962bd17c786126823ba1c4cf04ff4cccb2659d50fb20aee8.
//
// Solidity: event ValidatorRestaked(bytes32 pubkeyHash)
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
	PubkeyHash          [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorWithdrawn is a free log retrieval operation binding the contract event 0x5ce0aa04ae51d52da6e680fbe0336d2e2432f7c3dc2d4f3193204c57b9072107.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash)
func (_EigenPod *EigenPodFilterer) FilterValidatorWithdrawn(opts *bind.FilterOpts, checkpointTimestamp []uint64, pubkeyHash [][32]byte) (*EigenPodValidatorWithdrawnIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "ValidatorWithdrawn", checkpointTimestampRule, pubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodValidatorWithdrawnIterator{contract: _EigenPod.contract, event: "ValidatorWithdrawn", logs: logs, sub: sub}, nil
}

// WatchValidatorWithdrawn is a free log subscription operation binding the contract event 0x5ce0aa04ae51d52da6e680fbe0336d2e2432f7c3dc2d4f3193204c57b9072107.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash)
func (_EigenPod *EigenPodFilterer) WatchValidatorWithdrawn(opts *bind.WatchOpts, sink chan<- *EigenPodValidatorWithdrawn, checkpointTimestamp []uint64, pubkeyHash [][32]byte) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "ValidatorWithdrawn", checkpointTimestampRule, pubkeyHashRule)
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

// ParseValidatorWithdrawn is a log parse operation binding the contract event 0x5ce0aa04ae51d52da6e680fbe0336d2e2432f7c3dc2d4f3193204c57b9072107.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash)
func (_EigenPod *EigenPodFilterer) ParseValidatorWithdrawn(log types.Log) (*EigenPodValidatorWithdrawn, error) {
	event := new(EigenPodValidatorWithdrawn)
	if err := _EigenPod.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the EigenPod contract.
type EigenPodWithdrawalRequestedIterator struct {
	Event *EigenPodWithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *EigenPodWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodWithdrawalRequested)
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
		it.Event = new(EigenPodWithdrawalRequested)
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
func (it *EigenPodWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodWithdrawalRequested represents a WithdrawalRequested event raised by the EigenPod contract.
type EigenPodWithdrawalRequested struct {
	ValidatorPubkeyHash  [32]byte
	WithdrawalAmountGwei uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0x8b2737bb64ab2f2dc09552dfa1c250399e6a42c7ea9f0e1c658f5d65d708ec05.
//
// Solidity: event WithdrawalRequested(bytes32 indexed validatorPubkeyHash, uint64 withdrawalAmountGwei)
func (_EigenPod *EigenPodFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts, validatorPubkeyHash [][32]byte) (*EigenPodWithdrawalRequestedIterator, error) {

	var validatorPubkeyHashRule []interface{}
	for _, validatorPubkeyHashItem := range validatorPubkeyHash {
		validatorPubkeyHashRule = append(validatorPubkeyHashRule, validatorPubkeyHashItem)
	}

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "WithdrawalRequested", validatorPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodWithdrawalRequestedIterator{contract: _EigenPod.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0x8b2737bb64ab2f2dc09552dfa1c250399e6a42c7ea9f0e1c658f5d65d708ec05.
//
// Solidity: event WithdrawalRequested(bytes32 indexed validatorPubkeyHash, uint64 withdrawalAmountGwei)
func (_EigenPod *EigenPodFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *EigenPodWithdrawalRequested, validatorPubkeyHash [][32]byte) (event.Subscription, error) {

	var validatorPubkeyHashRule []interface{}
	for _, validatorPubkeyHashItem := range validatorPubkeyHash {
		validatorPubkeyHashRule = append(validatorPubkeyHashRule, validatorPubkeyHashItem)
	}

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "WithdrawalRequested", validatorPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodWithdrawalRequested)
				if err := _EigenPod.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
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

// ParseWithdrawalRequested is a log parse operation binding the contract event 0x8b2737bb64ab2f2dc09552dfa1c250399e6a42c7ea9f0e1c658f5d65d708ec05.
//
// Solidity: event WithdrawalRequested(bytes32 indexed validatorPubkeyHash, uint64 withdrawalAmountGwei)
func (_EigenPod *EigenPodFilterer) ParseWithdrawalRequested(log types.Log) (*EigenPodWithdrawalRequested, error) {
	event := new(EigenPodWithdrawalRequested)
	if err := _EigenPod.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
