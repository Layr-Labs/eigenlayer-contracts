// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

/**
 * @title Constants shared between 'EigenPod' and 'EigenPodManager' contracts.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
abstract contract EigenPodPausingConstants {
    /// @notice Index for flag that pauses creation of new EigenPods when set. See EigenPodManager code for details.
    uint8 internal constant PAUSED_NEW_EIGENPODS = 0;
    /**
     * @notice Index for flag that pauses all withdrawal-of-restaked ETH related functionality `
     * function *of the EigenPodManager* when set. See EigenPodManager code for details.
     */
    uint8 internal constant PAUSED_WITHDRAW_RESTAKED_ETH = 1;

    /// @notice Index for flag that pauses the deposit related functions *of the EigenPods* when set. see EigenPod code for details.
    uint8 internal constant PAUSED_EIGENPODS_VERIFY_CREDENTIALS = 2;

    // Deprecated
    // uint8 internal constant PAUSED_EIGENPODS_VERIFY_BALANCE_UPDATE = 3;

    // Deprecated
    // uint8 internal constant PAUSED_EIGENPODS_VERIFY_WITHDRAWAL = 4;

    /// @notice Pausability for EigenPod's "accidental transfer" withdrawal methods
    uint8 internal constant PAUSED_NON_PROOF_WITHDRAWALS = 5;

    uint8 internal constant PAUSED_START_CHECKPOINT = 6;

    /// @notice Index for flag that pauses the `verifyCheckpointProofs` function *of the EigenPods* when set. see EigenPod code for details.
    uint8 internal constant PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS = 7;

    uint8 internal constant PAUSED_VERIFY_STALE_BALANCE = 8;
}
