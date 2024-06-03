// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IEigenPod.sol";

abstract contract EigenPodStorage is IEigenPod {

    /// @notice The owner of this EigenPod
    address public podOwner;

    /**
     * @notice The latest timestamp at which the pod owner withdrew the balance of the pod, via calling `withdrawBeforeRestaking`.
     * @dev This variable is only updated when the `withdrawBeforeRestaking` function is called, which can only occur before `hasRestaked` is set to true for this pod.
     * Proofs for this pod are only valid against Beacon Chain state roots corresponding to timestamps after the stored `mostRecentWithdrawalTimestamp`.
     */
    uint64 internal __deprecated_mostRecentWithdrawalTimestamp;

    /// @notice the amount of execution layer ETH in this contract that is staked in EigenLayer (i.e. withdrawn from the Beacon Chain but not from EigenLayer),
    uint64 public withdrawableRestakedExecutionLayerGwei;

    /// @notice an indicator of whether or not the podOwner has ever "fully restaked" by successfully calling `verifyCorrectWithdrawalCredentials`.
    bool public hasRestaked;

    /// @notice This is a mapping of validatorPubkeyHash to timestamp to whether or not they have proven a withdrawal for that timestamp
    mapping(bytes32 => mapping(uint64 => bool)) internal __deprecated_provenWithdrawal;

    /// @notice This is a mapping that tracks a validator's information by their pubkey hash
    mapping(bytes32 => ValidatorInfo) internal _validatorPubkeyHashToInfo;

    /// @notice This variable tracks any ETH deposited into this contract via the `receive` fallback function
    uint256 internal __deprecated_nonBeaconChainETHBalanceWei;

    /// @notice This variable tracks the total amount of partial withdrawals claimed via merkle proofs prior to a switch to ZK proofs for claiming partial withdrawals
    uint64 __deprecated_sumOfPartialWithdrawalsClaimedGwei;

    /// @notice Number of validators with proven withdrawal credentials, who do not have proven full withdrawals
    uint256 public activeValidatorCount;

    /// @notice The timestamp of the last checkpoint finalized
    uint64 public lastCheckpointTimestamp;

    /// @notice The timestamp of the currently-active checkpoint. Will be 0 if there is not active checkpoint
    uint64 public currentCheckpointTimestamp;

    /// @notice The current checkpoint, if there is one active
    Checkpoint internal _currentCheckpoint;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[38] private __gap;
}