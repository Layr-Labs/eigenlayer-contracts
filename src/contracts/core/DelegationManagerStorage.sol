// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import "../libraries/SlashingLib.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IEigenPodManager.sol";
import "../interfaces/IAllocationManager.sol";

/**
 * @title Storage variables for the `DelegationManager` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract DelegationManagerStorage is IDelegationManager {
    // Constants

    /// @notice The EIP-712 typehash for the `StakerDelegation` struct used by the contract
    bytes32 public constant STAKER_DELEGATION_TYPEHASH =
        keccak256("StakerDelegation(address staker,address operator,uint256 nonce,uint256 expiry)");

    /// @notice The EIP-712 typehash for the `DelegationApproval` struct used by the contract
    bytes32 public constant DELEGATION_APPROVAL_TYPEHASH = keccak256(
        "DelegationApproval(address delegationApprover,address staker,address operator,bytes32 salt,uint256 expiry)"
    );

    /// @dev Index for flag that pauses new delegations when set
    uint8 internal constant PAUSED_NEW_DELEGATION = 0;

    /// @dev Index for flag that pauses queuing new withdrawals when set.
    uint8 internal constant PAUSED_ENTER_WITHDRAWAL_QUEUE = 1;

    /// @dev Index for flag that pauses completing existing withdrawals when set.
    uint8 internal constant PAUSED_EXIT_WITHDRAWAL_QUEUE = 2;

    /// @notice Canonical, virtual beacon chain ETH strategy
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    // Immutables

    /// @notice The AVSDirectory contract for EigenLayer
    IAVSDirectory public immutable avsDirectory;

    // TODO: Switch these to ShareManagers, but this breaks a lot of tests

    /// @notice The StrategyManager contract for EigenLayer
    IStrategyManager public immutable strategyManager;

    /// @notice The EigenPodManager contract for EigenLayer
    IEigenPodManager public immutable eigenPodManager;

    /// @notice The AllocationManager contract for EigenLayer
    IAllocationManager public immutable allocationManager;

    /// @notice Minimum withdrawal delay in blocks until a queued withdrawal can be completed.
    uint32 public immutable MIN_WITHDRAWAL_DELAY_BLOCKS;

    // Mutatables

    /// @dev Do not remove, deprecated storage.
    bytes32 internal __deprecated_DOMAIN_SEPARATOR;

    /**
     * @notice Returns the total number of shares owned by an `operator` for a given `strategy`.
     *
     * @dev By design, the following invariant should hold for each Strategy:
     *
     * (operator's delegatedShares in delegation manager) = sum (delegatedShares above zero of all stakers delegated to operator)
     * = sum (delegateable delegatedShares of all stakers delegated to the operator)
     */
    mapping(address operator => mapping(IStrategy strategy => uint256 shares)) public operatorShares;

    /// @notice Returns the operator details for a given `operator`.
    mapping(address operator => OperatorDetails) internal _operatorDetails;

    /// @notice Returns the `operator` a `staker` is delgated to, address(0) if not delegated.
    mapping(address staker => address operator) public delegatedTo;

    /// @notice Returns the number of EIP-712 signatures validated via `delegateToBySignature` for a given `staker`.
    mapping(address staker => uint256 nonce) public stakerNonce;

    /// @notice Returns whether `delegationApprover` has already used the given `salt`.
    mapping(address delegationApprover => mapping(bytes32 salt => bool spent)) public delegationApproverSaltIsSpent;

    /// @dev Do not remove, deprecated storage.
    uint256 private __deprecated_minWithdrawalDelayBlocks;

    /// @dev Returns whether a withdrawal is pending for a given `withdrawalRoot`.
    /// @dev This variable will be deprecated in the future, values should only be read or deleted.
    mapping(bytes32 withdrawalRoot => bool pending) public pendingWithdrawals;

    /// @notice Returns the total number of withdrawals that have been queued for a given `staker`.
    /// @dev This only increments (doesn't decrement), and is used to help ensure that otherwise identical withdrawals have unique hashes.
    mapping(address staker => uint256 totalQueued) public cumulativeWithdrawalsQueued;

    /// @dev Do not remove, deprecated storage.
    /// See conversation here: https://github.com/Layr-Labs/eigenlayer-contracts/pull/365/files#r1417525270
    address private __deprecated_stakeRegistry;

    /// @dev Do not remove, deprecated storage.
    mapping(IStrategy strategy => uint256 delayBlocks) private __deprecated_strategyWithdrawalDelayBlocks;

    /// @notice Returns the scaling factors for a `staker` for a given `strategy`.
    /// @dev We do not need the `beaconChainScalingFactor` for non-beaconchain strategies, but it's nicer syntactically to keep it.
    mapping(address staker => mapping(IStrategy strategy => StakerScalingFactors)) public stakerScalingFactor;

    /// @notice Returns a list of queued withdrawals for a given `staker`.
    /// @dev Entrys are removed when the withdrawal is completed.
    /// @dev This variable only reflects withdrawals that were made after the slashing release.
    mapping(address staker => EnumerableSet.Bytes32Set withdrawalRoots) internal _stakerQueuedWithdrawalRoots;

    /// @notice Returns the details of a queued withdrawal for a given `staker` and `withdrawalRoot`.
    /// @dev This variable only reflects withdrawals that were made after the slashing release.
    mapping(bytes32 withdrawalRoot => Withdrawal withdrawal) public queuedWithdrawals;

    // Construction

    constructor(
        IAVSDirectory _avsDirectory,
        IStrategyManager _strategyManager,
        IEigenPodManager _eigenPodManager,
        IAllocationManager _allocationManager,
        uint32 _MIN_WITHDRAWAL_DELAY_BLOCKS
    ) {
        avsDirectory = _avsDirectory;
        strategyManager = _strategyManager;
        eigenPodManager = _eigenPodManager;
        allocationManager = _allocationManager;
        MIN_WITHDRAWAL_DELAY_BLOCKS = _MIN_WITHDRAWAL_DELAY_BLOCKS;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[36] private __gap;
}
