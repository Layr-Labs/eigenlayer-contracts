// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IEigenPod.sol";
import "../../contracts/mixins/SemVerMixin.sol";

contract EigenPodMock is IEigenPod, SemVerMixin, Test {
    constructor() SemVerMixin("v9.9.9") {}

    function nonBeaconChainETHBalanceWei() external view returns (uint) {}

    /// @notice the amount of execution layer ETH in this contract that is staked in EigenLayer (i.e. withdrawn from beaconchain but not EigenLayer),
    function withdrawableRestakedExecutionLayerGwei() external view returns (uint64) {}

    /// @notice Used to initialize the pointers to contracts crucial to the pod's functionality, in beacon proxy construction from EigenPodManager
    function initialize(address owner) external {}

    /// @notice Called by EigenPodManager when the owner wants to create another ETH validator.
    function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable {}

    /**
     * @notice Transfers `amountWei` in ether from this contract to the specified `recipient` address
     * @notice Called by EigenPodManager to withdrawBeaconChainETH that has been added to the EigenPod's balance due to a withdrawal from the beacon chain.
     * @dev Called during withdrawal or slashing.
     * @dev Note that this function is marked as non-reentrant to prevent the recipient calling back into it
     */
    function withdrawRestakedBeaconChainETH(address recipient, uint amount) external {}

    /// @notice The single EigenPodManager for EigenLayer
    function eigenPodManager() external view returns (IEigenPodManager) {}

    /// @notice The owner of this EigenPod
    function podOwner() external view returns (address) {}

    /// @notice an indicator of whether or not the podOwner has ever "fully restaked" by successfully calling `verifyCorrectWithdrawalCredentials`.
    function hasRestaked() external view returns (bool) {}

    /// @notice block timestamp of the most recent withdrawal
    function mostRecentWithdrawalTimestamp() external view returns (uint64) {}

    /// @notice Returns the validatorInfo struct for the provided pubkeyHash
    function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) external view returns (ValidatorInfo memory) {}

    /// @notice This returns the status of a given validator
    function validatorStatus(bytes32 pubkeyHash) external view returns (VALIDATOR_STATUS) {}

    /// @notice Number of validators with proven withdrawal credentials, who do not have proven full withdrawals
    function activeValidatorCount() external view returns (uint) {}

    /// @notice The timestamp of the last checkpoint finalized
    function lastCheckpointTimestamp() external view returns (uint64) {}

    /// @notice The timestamp of the currently-active checkpoint. Will be 0 if there is not active checkpoint
    function currentCheckpointTimestamp() external view returns (uint64) {}

    /// @notice Returns the currently-active checkpoint
    function currentCheckpoint() external view returns (Checkpoint memory) {}

    function checkpointBalanceExitedGwei(uint64) external view returns (uint64) {}

    function startCheckpoint(bool revertIfNoBalance) external {}

    function verifyCheckpointProofs(
        BeaconChainProofs.BalanceContainerProof calldata balanceContainerProof,
        BeaconChainProofs.BalanceProof[] calldata proofs
    ) external {}

    function verifyStaleBalance(
        uint64 beaconTimestamp,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        BeaconChainProofs.ValidatorProof calldata proof
    ) external {}

    function verifyWithdrawalCredentials(
        uint64 oracleTimestamp,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        uint40[] calldata validatorIndices,
        bytes[] calldata withdrawalCredentialProofs,
        bytes32[][] calldata validatorFields
    ) external {}

    /// @notice Called by the pod owner to withdraw the balance of the pod when `hasRestaked` is set to false
    function activateRestaking() external {}

    /// @notice Called by the pod owner to withdraw the balance of the pod when `hasRestaked` is set to false
    function withdrawBeforeRestaking() external {}

    /// @notice Called by the pod owner to withdraw the nonBeaconChainETHBalanceWei
    function withdrawNonBeaconChainETHBalanceWei(address recipient, uint amountToWithdraw) external {}

    /// @notice called by owner of a pod to remove any ERC20s deposited in the pod
    function recoverTokens(IERC20[] memory tokenList, uint[] memory amountsToWithdraw, address recipient) external {}

    function setProofSubmitter(address newProofSubmitter) external {}

    function proofSubmitter() external view returns (address) {}

    function validatorStatus(bytes calldata pubkey) external view returns (VALIDATOR_STATUS) {}
    function validatorPubkeyToInfo(bytes calldata validatorPubkey) external view returns (ValidatorInfo memory) {}

    /// @notice Query the 4788 oracle to get the parent block root of the slot with the given `timestamp`
    /// @param timestamp of the block for which the parent block root will be returned. MUST correspond
    /// to an existing slot within the last 24 hours. If the slot at `timestamp` was skipped, this method
    /// will revert.
    function getParentBlockRoot(uint64 timestamp) external view returns (bytes32) {}
}
