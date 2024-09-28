// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

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

    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

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

    /// @notice The minimum number of blocks to complete a withdrawal of a strategy. 50400 * 12 seconds = 1 week
    uint256 public constant LEGACY_MIN_WITHDRAWAL_DELAY_BLOCKS = 50_400;

    /// @notice Wed Jan 01 2025 17:00:00 GMT+0000, timestamp used to check whether a pending withdrawal
    /// should be processed as legacy M2 or with slashing considered.
    uint32 public constant LEGACY_WITHDRAWALS_TIMESTAMP = 1_735_750_800;

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

    /// @notice Minimum withdrawal delay in seconds until all queued withdrawals can be completed.
    uint32 public immutable MIN_WITHDRAWAL_DELAY;

    /// @dev Chain ID at the time of contract deployment
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    // Mutatables
    
    /**
     * @notice Original EIP-712 Domain separator for this contract.
     * @dev The domain separator may change in the event of a fork that modifies the ChainID.
     * Use the getter function `domainSeparator` to get the current domain separator for this contract.
     */
    bytes32 internal _DOMAIN_SEPARATOR;

    /**
     * @notice returns the total number of delegatedShares (i.e. shares divided by the `operator`'s
     * totalMagnitude) in `strategy` that are delegated to `operator`.
     * @notice Mapping: operator => strategy => total number of delegatedShares in the strategy delegated to the operator.
     * @dev By design, the following invariant should hold for each Strategy:
     * (operator's delegatedShares in delegation manager) = sum (delegatedShares above zero of all stakers delegated to operator)
     * = sum (delegateable delegatedShares of all stakers delegated to the operator)
     * @dev FKA `operatorShares`
     */
    mapping(address => mapping(IStrategy => DelegatedShares)) public operatorDelegatedShares;

    /**
     * @notice Mapping: operator => OperatorDetails struct
     * @dev This struct is internal with an external getter so we can return an `OperatorDetails memory` object
     */
    mapping(address => OperatorDetails) internal _operatorDetails;

    /**
     * @notice Mapping: staker => operator whom the staker is currently delegated to.
     * @dev Note that returning address(0) indicates that the staker is not actively delegated to any operator.
     */
    mapping(address => address) public delegatedTo;

    /// @notice Mapping: staker => number of signed messages (used in `delegateToBySignature`) from the staker that this contract has already checked.
    mapping(address => uint256) public stakerNonce;

    /**
     * @notice Mapping: delegationApprover => 32-byte salt => whether or not the salt has already been used by the delegationApprover.
     * @dev Salts are used in the `delegateTo` and `delegateToBySignature` functions. Note that these functions only process the delegationApprover's
     * signature + the provided salt if the operator being delegated to has specified a nonzero address as their `delegationApprover`.
     */
    mapping(address => mapping(bytes32 => bool)) public delegationApproverSaltIsSpent;

    /**
     * @notice Global minimum withdrawal delay for all strategy withdrawals.
     * In a prior Goerli release, we only had a global min withdrawal delay across all strategies.
     * In addition, we now also configure withdrawal delays on a per-strategy basis.
     * To withdraw from a strategy, max(minWithdrawalDelayBlocks, strategyWithdrawalDelayBlocks[strategy]) number of blocks must have passed.
     * See mapping strategyWithdrawalDelayBlocks below for per-strategy withdrawal delays.
     */
    uint256 private __deprecated_minWithdrawalDelayBlocks;

    /// @notice Mapping: hash of withdrawal inputs, aka 'withdrawalRoot' => whether the withdrawal is pending
    mapping(bytes32 => bool) public pendingWithdrawals;

    /// @notice Mapping: staker => cumulative number of queued withdrawals they have ever initiated.
    /// @dev This only increments (doesn't decrement), and is used to help ensure that otherwise identical withdrawals have unique hashes.
    mapping(address => uint256) public cumulativeWithdrawalsQueued;

    /// @notice Deprecated from an old Goerli release
    /// See conversation here: https://github.com/Layr-Labs/eigenlayer-contracts/pull/365/files#r1417525270
    address private __deprecated_stakeRegistry;

    /**
     * @notice Minimum delay enforced by this contract per Strategy for completing queued withdrawals. Measured in blocks, and adjustable by this contract's owner,
     * up to a maximum of `MAX_WITHDRAWAL_DELAY_BLOCKS`. Minimum value is 0 (i.e. no delay enforced).
     */
    mapping(IStrategy => uint256) private __deprecated_strategyWithdrawalDelayBlocks;

    /// @notice Mapping: staker => strategy =>
    ///    (
    ///       scaling factor used to calculate the staker's shares in the strategy,
    ///       beacon chain scaling factor used to calculate the staker's withdrawable shares in the strategy.
    ///    )
    /// Note that we don't need the beaconChainScalingFactor for non beaconChainETHStrategy strategies, but it's nicer syntactically to keep it.
    mapping(address => mapping(IStrategy => StakerScalingFactors)) public stakerScalingFactors;

    // Construction

    constructor(
        IAVSDirectory _avsDirectory,
        IStrategyManager _strategyManager,
        IEigenPodManager _eigenPodManager,
        IAllocationManager _allocationManager,
        uint32 _MIN_WITHDRAWAL_DELAY
    ) {
        avsDirectory = _avsDirectory;
        strategyManager = _strategyManager;
        eigenPodManager = _eigenPodManager;
        allocationManager = _allocationManager;
        MIN_WITHDRAWAL_DELAY = _MIN_WITHDRAWAL_DELAY;
        ORIGINAL_CHAIN_ID = block.chainid;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[38] private __gap;
}
