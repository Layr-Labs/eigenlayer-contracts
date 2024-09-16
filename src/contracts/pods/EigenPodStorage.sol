// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IEigenPod.sol";

abstract contract EigenPodStorage is IEigenPod {
    /// @notice The owner of this EigenPod
    address public podOwner;

    /// @notice DEPRECATED: previously used to track the time when restaking was activated
    uint64 internal __deprecated_mostRecentWithdrawalTimestamp;

    /// @notice the amount of execution layer ETH in this contract that is staked in EigenLayer (i.e. withdrawn from the Beacon Chain but not from EigenLayer),
    uint64 internal restakedExecutionLayerGwei;

    /// @notice DEPRECATED: previously used to track whether a pod had activated restaking
    bool internal __deprecated_hasRestaked;

    /// @notice DEPRECATED: previously tracked withdrawals proven per validator
    mapping(bytes32 => mapping(uint64 => bool)) internal __deprecated_provenWithdrawal;

    /// @notice This is a mapping that tracks a validator's information by their pubkey hash
    mapping(bytes32 => ValidatorInfo) internal _validatorPubkeyHashToInfo;

    /// @notice DEPRECATED: previously used to track ETH sent to the fallback function
    uint256 internal __deprecated_nonBeaconChainETHBalanceWei;

    /// @notice DEPRECATED: previously used to track claimed partial withdrawals
    uint64 __deprecated_sumOfPartialWithdrawalsClaimedGwei;

    /// @notice Number of validators with proven withdrawal credentials, who do not have proven full withdrawals
    uint256 public activeValidatorCount;

    /// @notice The timestamp of the last checkpoint finalized
    uint64 public lastCheckpointTimestamp;

    /// @notice The timestamp of the currently-active checkpoint. Will be 0 if there is not active checkpoint
    uint64 public currentCheckpointTimestamp;

    /// @notice For each checkpoint, the total balance attributed to exited validators, in gwei
    ///
    /// NOTE that the values added to this mapping are NOT guaranteed to capture the entirety of a validator's
    /// exit - rather, they capture the total change in a validator's balance when a checkpoint shows their
    /// balance change from nonzero to zero. While a change from nonzero to zero DOES guarantee that a validator
    /// has been fully exited, it is possible that the magnitude of this change does not capture what is
    /// typically thought of as a "full exit."
    ///
    /// For example:
    /// 1. Consider a validator was last checkpointed at 32 ETH before exiting. Once the exit has been processed,
    /// it is expected that the validator's exited balance is calculated to be `32 ETH`.
    /// 2. However, before `startCheckpoint` is called, a deposit is made to the validator for 1 ETH. The beacon
    /// chain will automatically withdraw this ETH, but not until the withdrawal sweep passes over the validator
    /// again. Until this occurs, the validator's current balance (used for checkpointing) is 1 ETH.
    /// 3. If `startCheckpoint` is called at this point, the balance delta calculated for this validator will be
    /// `-31 ETH`, and because the validator has a nonzero balance, it is not marked WITHDRAWN.
    /// 4. After the exit is processed by the beacon chain, a subsequent `startCheckpoint` and checkpoint proof
    /// will calculate a balance delta of `-1 ETH` and attribute a 1 ETH exit to the validator.
    ///
    /// If this edge case impacts your usecase, it should be possible to mitigate this by monitoring for deposits
    /// to your exited validators, and waiting to call `startCheckpoint` until those deposits have been automatically
    /// exited.
    ///
    /// Additional edge cases this mapping does not cover:
    /// - If a validator is slashed, their balance exited will reflect their original balance rather than the slashed amount
    /// - The final partial withdrawal for an exited validator will be likely be included in this mapping.
    ///   i.e. if a validator was last checkpointed at 32.1 ETH before exiting, the next checkpoint will calculate their
    ///   "exited" amount to be 32.1 ETH rather than 32 ETH.
    mapping(uint64 => uint64) public checkpointBalanceExitedGwei;

    /// @notice The current checkpoint, if there is one active
    Checkpoint internal _currentCheckpoint;

    /// @notice An address with permissions to call `startCheckpoint` and `verifyWithdrawalCredentials`, set
    /// by the podOwner. This role exists to allow a podOwner to designate a hot wallet that can call
    /// these methods, allowing the podOwner to remain a cold wallet that is only used to manage funds.
    /// @dev If this address is NOT set, only the podOwner can call `startCheckpoint` and `verifyWithdrawalCredentials`
    address public proofSubmitter;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[35] private __gap;
}
