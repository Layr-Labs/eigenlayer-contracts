// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./IStrategy.sol";
import "./IPauserRegistry.sol";
import "./ISignatureUtils.sol";
import "../libraries/SlashingLib.sol";

interface IDelegationManagerErrors {
    /// @dev Thrown when msg.sender is not allowed to call a function
    error UnauthorizedCaller();
    /// @dev Thrown when msg.sender is not the EigenPodManager
    error OnlyEigenPodManager();
    /// @dev Throw when msg.sender is not the AllocationManager
    error OnlyAllocationManager();

    /// Delegation Status

    /// @dev Thrown when an operator attempts to undelegate.
    error OperatorsCannotUndelegate();
    /// @dev Thrown when an account is actively delegated.
    error ActivelyDelegated();
    /// @dev Thrown when an account is not actively delegated.
    error NotActivelyDelegated();
    /// @dev Thrown when `operator` is not a registered operator.
    error OperatorNotRegistered();

    /// Invalid Inputs

    /// @dev Thrown when attempting to execute an action that was not queued.
    error WithdrawalNotQueued();
    /// @dev Thrown when caller cannot undelegate on behalf of a staker.
    error CallerCannotUndelegate();
    /// @dev Thrown when two array parameters have mismatching lengths.
    error InputArrayLengthMismatch();
    /// @dev Thrown when input arrays length is zero.
    error InputArrayLengthZero();
    /// @dev Thrown when caller is neither the StrategyManager or EigenPodManager contract.
    error OnlyStrategyManagerOrEigenPodManager();

    /// Slashing

    /// @dev Thrown when an operator has been fully slashed(maxMagnitude is 0) for a strategy.
    /// or if the staker has had been natively slashed to the point of their beaconChainScalingFactor equalling 0.
    error FullySlashed();

    /// Signatures

    /// @dev Thrown when attempting to spend a spent eip-712 salt.
    error SaltSpent();

    /// Withdrawal Processing

    /// @dev Thrown when attempting to execute an action that was not queued.
    error WithdrawalDoesNotExist();
    /// @dev Thrown when attempting to withdraw before delay has elapsed.
    error WithdrawalDelayNotElapsed();
    /// @dev Thrown when provided delay exceeds maximum.
    error WithdrawalDelayExeedsMax();
    /// @dev Thrown when a withdraw amount larger than max is attempted.
    error WithdrawalExceedsMax();
    /// @dev Thrown when withdrawer is not the current caller.
    error WithdrawerNotCaller();
    /// @dev Thrown when `withdrawer` is not staker.
    error WithdrawerNotStaker();
}

interface IDelegationManagerTypes {
    // @notice Struct used for storing information about a single operator who has registered with EigenLayer
    struct OperatorDetails {
        /// @notice DEPRECATED -- this field is no longer used, payments are handled in RewardsCoordinator.sol
        address __deprecated_earningsReceiver;
        /**
         * @notice Address to verify signatures when a staker wishes to delegate to the operator, as well as controlling "forced undelegations".
         * @dev Signature verification follows these rules:
         * 1) If this address is left as address(0), then any staker will be free to delegate to the operator, i.e. no signature verification will be performed.
         * 2) If this address is an EOA (i.e. it has no code), then we follow standard ECDSA signature verification for delegations to the operator.
         * 3) If this address is a contract (i.e. it has code) then we forward a call to the contract and verify that it returns the correct EIP-1271 "magic value".
         */
        address delegationApprover;
        /// @notice DEPRECATED -- this field is no longer used. An analogous field is the `allocationDelay` stored in the AllocationManager
        uint32 __deprecated_stakerOptOutWindowBlocks;
    }

    /**
     * @notice Abstract struct used in calculating an EIP712 signature for an operator's delegationApprover to approve that a specific staker delegate to the operator.
     * @dev Used in computing the `DELEGATION_APPROVAL_TYPEHASH` and as a reference in the computation of the approverDigestHash in the `_delegate` function.
     */
    struct DelegationApproval {
        // the staker who is delegating
        address staker;
        // the operator being delegated to
        address operator;
        // the operator's provided salt
        bytes32 salt;
        // the expiration timestamp (UTC) of the signature
        uint256 expiry;
    }

    /**
     * Struct type used to specify an existing queued withdrawal. Rather than storing the entire struct, only a hash is stored.
     * In functions that operate on existing queued withdrawals -- e.g. completeQueuedWithdrawal`, the data is resubmitted and the hash of the submitted
     * data is computed by `calculateWithdrawalRoot` and checked against the stored hash in order to confirm the integrity of the submitted data.
     */
    struct Withdrawal {
        // The address that originated the Withdrawal
        address staker;
        // The address that the staker was delegated to at the time that the Withdrawal was created
        address delegatedTo;
        // The address that can complete the Withdrawal + will receive funds when completing the withdrawal
        address withdrawer;
        // Nonce used to guarantee that otherwise identical withdrawals have unique hashes
        uint256 nonce;
        // Blocknumber when the Withdrawal was created.
        uint32 startBlock;
        // Array of strategies that the Withdrawal contains
        IStrategy[] strategies;
        // Array containing the amount of staker's scaledShares for withdrawal in each Strategy in the `strategies` array
        // Note that these scaledShares need to be multiplied by the operator's maxMagnitude and beaconChainScalingFactor at completion to include
        // slashing occurring during the queue withdrawal delay. This is because scaledShares = sharesToWithdraw / (maxMagnitude * beaconChainScalingFactor)
        // at queue time. beaconChainScalingFactor is simply equal to 1 if the strategy is not the beaconChainStrategy.
        // To account for slashing, we later multiply scaledShares * maxMagnitude * beaconChainScalingFactor at the earliest possible completion time
        // to get the withdrawn shares after applying slashing during the delay period.
        uint256[] scaledShares;
    }

    struct QueuedWithdrawalParams {
        // Array of strategies that the QueuedWithdrawal contains
        IStrategy[] strategies;
        // Array containing the amount of depositShares for withdrawal in each Strategy in the `strategies` array
        // Note that the actual shares received on completing withdrawal may be less than the depositShares if slashing occurred
        uint256[] depositShares;
        // The address of the withdrawer
        address withdrawer;
    }
}

interface IDelegationManagerEvents is IDelegationManagerTypes {
    // @notice Emitted when a new operator registers in EigenLayer and provides their OperatorDetails.
    event OperatorRegistered(address indexed operator, OperatorDetails operatorDetails);

    /// @notice Emitted when an operator updates their OperatorDetails to @param newOperatorDetails
    event OperatorDetailsModified(address indexed operator, OperatorDetails newOperatorDetails);

    /**
     * @notice Emitted when @param operator indicates that they are updating their MetadataURI string
     * @dev Note that these strings are *never stored in storage* and are instead purely emitted in events for off-chain indexing
     */
    event OperatorMetadataURIUpdated(address indexed operator, string metadataURI);

    /// @notice Emitted whenever an operator's shares are increased for a given strategy. Note that shares is the delta in the operator's shares.
    event OperatorSharesIncreased(address indexed operator, address staker, IStrategy strategy, uint256 shares);

    /// @notice Emitted whenever an operator's shares are decreased for a given strategy. Note that shares is the delta in the operator's shares.
    event OperatorSharesDecreased(address indexed operator, address staker, IStrategy strategy, uint256 shares);

    /// @notice Emitted when @param staker delegates to @param operator.
    event StakerDelegated(address indexed staker, address indexed operator);

    /// @notice Emitted when @param staker undelegates from @param operator.
    event StakerUndelegated(address indexed staker, address indexed operator);

    /// @notice Emitted when @param staker is undelegated via a call not originating from the staker themself
    event StakerForceUndelegated(address indexed staker, address indexed operator);

    /// @notice Emitted when a staker's depositScalingFactor is updated
    event DepositScalingFactorUpdated(address staker, IStrategy strategy, uint256 newDepositScalingFactor);

    /// @notice Emitted when a staker's beaconChainScalingFactor is updated
    event BeaconChainScalingFactorDecreased(address staker, uint64 newBeaconChainScalingFactor);

    /**
     * @notice Emitted when a new withdrawal is queued.
     * @param withdrawalRoot Is the hash of the `withdrawal`.
     * @param withdrawal Is the withdrawal itself.
     * @param sharesToWithdraw Is an array of the expected shares that were queued for withdrawal corresponding to the strategies in the `withdrawal`.
     */
    event SlashingWithdrawalQueued(bytes32 withdrawalRoot, Withdrawal withdrawal, uint256[] sharesToWithdraw);

    /// @notice Emitted when a queued withdrawal is completed
    event SlashingWithdrawalCompleted(bytes32 withdrawalRoot);
}

/**
 * @title DelegationManager
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice  This is the contract for delegation in EigenLayer. The main functionalities of this contract are
 * - enabling anyone to register as an operator in EigenLayer
 * - allowing operators to specify parameters related to stakers who delegate to them
 * - enabling any staker to delegate its stake to the operator of its choice (a given staker can only delegate to a single operator at a time)
 * - enabling a staker to undelegate its assets from the operator it is delegated to (performed as part of the withdrawal process, initiated through the StrategyManager)
 */
interface IDelegationManager is ISignatureUtils, IDelegationManagerErrors, IDelegationManagerEvents {
    /**
     * @dev Initializes the initial owner and paused status.
     */
    function initialize(address initialOwner, uint256 initialPausedStatus) external;

    /**
     * @notice Registers the caller as an operator in EigenLayer.
     * @param registeringOperatorDetails is the `OperatorDetails` for the operator.
     * @param allocationDelay The delay before allocations take effect.
     * @param metadataURI is a URI for the operator's metadata, i.e. a link providing more details on the operator.
     *
     * @dev Once an operator is registered, they cannot 'deregister' as an operator, and they will forever be considered "delegated to themself".
     * @dev This function will revert if the caller is already delegated to an operator.
     * @dev Note that the `metadataURI` is *never stored * and is only emitted in the `OperatorMetadataURIUpdated` event
     */
    function registerAsOperator(
        OperatorDetails calldata registeringOperatorDetails,
        uint32 allocationDelay,
        string calldata metadataURI
    ) external;

    /**
     * @notice Updates an operator's stored `OperatorDetails`.
     * @param operator is the operator to update.
     * @param newOperatorDetails is the updated `OperatorDetails` for the operator, to replace their current OperatorDetails`.
     *
     * @dev The caller must have previously registered as an operator in EigenLayer.
     */
    function modifyOperatorDetails(address operator, OperatorDetails calldata newOperatorDetails) external;

    /**
     * @notice Called by an operator to emit an `OperatorMetadataURIUpdated` event indicating the information has updated.
     * @param operator The operator to update metadata for
     * @param metadataURI The URI for metadata associated with an operator
     * @dev Note that the `metadataURI` is *never stored * and is only emitted in the `OperatorMetadataURIUpdated` event
     */
    function updateOperatorMetadataURI(address operator, string calldata metadataURI) external;

    /**
     * @notice Caller delegates their stake to an operator.
     * @param operator The account (`msg.sender`) is delegating its assets to for use in serving applications built on EigenLayer.
     * @param approverSignatureAndExpiry Verifies the operator approves of this delegation
     * @param approverSalt A unique single use value tied to an individual signature.
     * @dev The approverSignatureAndExpiry is used in the event that:
     *          1) the operator's `delegationApprover` address is set to a non-zero value.
     *                  AND
     *          2) neither the operator nor their `delegationApprover` is the `msg.sender`, since in the event that the operator
     *             or their delegationApprover is the `msg.sender`, then approval is assumed.
     * @dev In the event that `approverSignatureAndExpiry` is not checked, its content is ignored entirely; it's recommended to use an empty input
     * in this case to save on complexity + gas costs
     * @dev If the staker delegating has shares in a strategy that the operator was slashed 100% for (the operator's maxMagnitude = 0),
     * then delegation is blocked and will revert.
     */
    function delegateTo(
        address operator,
        SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt
    ) external;

    /**
     * @notice Undelegates the staker from the operator who they are delegated to.
     * Queues withdrawals of all of the staker's withdrawable shares in the StrategyManager (to the staker) and/or EigenPodManager, if necessary.
     * @param staker The account to be undelegated.
     * @return withdrawalRoots The roots of the newly queued withdrawals, if a withdrawal was queued. Otherwise just bytes32(0).
     *
     * @dev Reverts if the `staker` is also an operator, since operators are not allowed to undelegate from themselves.
     * @dev Reverts if the caller is not the staker, nor the operator who the staker is delegated to, nor the operator's specified "delegationApprover"
     * @dev Reverts if the `staker` is already undelegated.
     */
    function undelegate(
        address staker
    ) external returns (bytes32[] memory withdrawalRoots);

    /**
     * @notice Allows a staker to withdraw some shares. Withdrawn shares/strategies are immediately removed
     * from the staker. If the staker is delegated, withdrawn shares/strategies are also removed from
     * their operator.
     *
     * All withdrawn shares/strategies are placed in a queue and can be withdrawn after a delay. Withdrawals
     * are still subject to slashing during the delay period so the amount withdrawn on completion may actually be less
     * than what was queued if slashing has occurred in that period.
     *
     * @dev To view what the staker is able to queue withdraw, see `getWithdrawableShares()`
     */
    function queueWithdrawals(
        QueuedWithdrawalParams[] calldata params
    ) external returns (bytes32[] memory);

    /**
     * @notice Used to complete the all queued withdrawals.
     * Used to complete the specified `withdrawals`. The function caller must match `withdrawals[...].withdrawer`
     * @param tokens Array of tokens for each Withdrawal. See `completeQueuedWithdrawal` for the usage of a single array.
     * @param receiveAsTokens Whether or not to complete each withdrawal as tokens. See `completeQueuedWithdrawal` for the usage of a single boolean.
     * @param numToComplete The number of withdrawals to complete. This must be less than or equal to the number of queued withdrawals.
     * @dev See `completeQueuedWithdrawal` for relevant dev tags
     */
    function completeQueuedWithdrawals(
        IERC20[][] calldata tokens,
        bool[] calldata receiveAsTokens,
        uint256 numToComplete
    ) external;

    /**
     * @notice Used to complete the lastest queued withdrawal.
     * @param withdrawal The withdrawal to complete.
     * @param tokens Array in which the i-th entry specifies the `token` input to the 'withdraw' function of the i-th Strategy in the `withdrawal.strategies` array.
     * @param receiveAsTokens If true, the shares calculated to be withdrawn will be withdrawn from the specified strategies themselves
     * and sent to the caller, through calls to `withdrawal.strategies[i].withdraw`. If false, then the shares in the specified strategies
     * will simply be transferred to the caller directly.
     * @dev beaconChainETHStrategy shares are non-transferrable, so if `receiveAsTokens = false` and `withdrawal.withdrawer != withdrawal.staker`, note that
     * any beaconChainETHStrategy shares in the `withdrawal` will be _returned to the staker_, rather than transferred to the withdrawer, unlike shares in
     * any other strategies, which will be transferred to the withdrawer.
     */
    function completeQueuedWithdrawal(
        Withdrawal calldata withdrawal,
        IERC20[] calldata tokens,
        bool receiveAsTokens
    ) external;

    /**
     * @notice Used to complete the all queued withdrawals.
     * Used to complete the specified `withdrawals`. The function caller must match `withdrawals[...].withdrawer`
     * @param withdrawals Array of Withdrawals to complete. See `completeQueuedWithdrawal` for the usage of a single Withdrawal.
     * @param tokens Array of tokens for each Withdrawal. See `completeQueuedWithdrawal` for the usage of a single array.
     * @param receiveAsTokens Whether or not to complete each withdrawal as tokens. See `completeQueuedWithdrawal` for the usage of a single boolean.
     * @dev See `completeQueuedWithdrawal` for relevant dev tags
     */
    function completeQueuedWithdrawals(
        Withdrawal[] calldata withdrawals,
        IERC20[][] calldata tokens,
        bool[] calldata receiveAsTokens
    ) external;

    /**
     * @notice Increases a staker's delegated share balance in a strategy. Note that before adding to operator shares,
     * the delegated delegatedShares. The staker's depositScalingFactor is updated here.
     * @param staker The address to increase the delegated shares for their operator.
     * @param strategy The strategy in which to increase the delegated shares.
     * @param existingDepositShares The number of deposit shares the staker already has in the strategy. This is the shares amount stored in the
     * StrategyManager/EigenPodManager for the staker's shares.
     * @param addedShares The number of shares added to the staker's shares in the strategy
     *
     * @dev *If the staker is actively delegated*, then increases the `staker`'s delegated delegatedShares in `strategy`.
     * Otherwise does nothing.
     * @dev If the operator was slashed 100% for the strategy (the operator's maxMagnitude = 0), then increasing delegated shares is blocked and will revert.
     * @dev Callable only by the StrategyManager or EigenPodManager.
     */
    function increaseDelegatedShares(
        address staker,
        IStrategy strategy,
        uint256 existingDepositShares,
        uint256 addedShares
    ) external;

    /**
     * @notice Decreases a native restaker's delegated share balance in a strategy due to beacon chain slashing. This updates their beaconChainScalingFactor.
     * Their operator's stakeShares are also updated (if they are delegated).
     * @param staker The address to increase the delegated stakeShares for their operator.
     * @param existingShares The number of shares the staker already has in the EPM. This does not change upon decreasing shares.
     * @param proportionOfOldBalance The current pod owner shares proportion of the previous pod owner shares
     *
     * @dev *If the staker is actively delegated*, then decreases the `staker`'s delegated stakeShares in `strategy` by `proportionPodBalanceDecrease` proportion. Otherwise does nothing.
     * @dev Callable only by the EigenPodManager.
     */
    function decreaseBeaconChainScalingFactor(
        address staker,
        uint256 existingShares,
        uint64 proportionOfOldBalance
    ) external;

    /**
     * @notice Decreases the operators shares in storage after a slash
     * @param operator The operator to decrease shares for
     * @param strategy The strategy to decrease shares for
     * @param wadSlashed The proportion of 1e18 slashed
     * @dev Callable only by the AllocationManager
     */
    function decreaseOperatorShares(address operator, IStrategy strategy, uint256 wadSlashed) external;

    /**
     *
     *                         BACKWARDS COMPATIBLE LEGACY FUNCTIONS
     *                         TO BE DEPRECATED IN FUTURE
     *
     */

    /// @notice Overloaded version of `completeQueuedWithdrawal` that includes a `middlewareTimesIndex` parameter
    /// for backwards compatibility with the M2 release. To be deprecated in a future release.
    function completeQueuedWithdrawal(
        Withdrawal calldata withdrawal,
        IERC20[] calldata tokens,
        uint256 middlewareTimesIndex,
        bool receiveAsTokens
    ) external;

    /// @notice Overloaded version of `completeQueuedWithdrawals` that includes a `middlewareTimesIndexes` parameter
    /// for backwards compatibility with the M2 release. To be deprecated in a future release.
    function completeQueuedWithdrawals(
        Withdrawal[] calldata withdrawals,
        IERC20[][] calldata tokens,
        uint256[] calldata middlewareTimesIndexes,
        bool[] calldata receiveAsTokens
    ) external;

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice returns the address of the operator that `staker` is delegated to.
     * @notice Mapping: staker => operator whom the staker is currently delegated to.
     * @dev Note that returning address(0) indicates that the staker is not actively delegated to any operator.
     */
    function delegatedTo(
        address staker
    ) external view returns (address);

    /**
     * @notice Mapping: delegationApprover => 32-byte salt => whether or not the salt has already been used by the delegationApprover.
     * @dev Salts are used in the `delegateTo` function. Note that this function only processes the delegationApprover's
     * signature + the provided salt if the operator being delegated to has specified a nonzero address as their `delegationApprover`.
     */
    function delegationApproverSaltIsSpent(address _delegationApprover, bytes32 salt) external view returns (bool);

    /// @notice Mapping: staker => cumulative number of queued withdrawals they have ever initiated.
    /// @dev This only increments (doesn't decrement), and is used to help ensure that otherwise identical withdrawals have unique hashes.
    function cumulativeWithdrawalsQueued(
        address staker
    ) external view returns (uint256);

    /**
     * @notice Returns 'true' if `staker` *is* actively delegated, and 'false' otherwise.
     */
    function isDelegated(
        address staker
    ) external view returns (bool);

    /**
     * @notice Returns true is an operator has previously registered for delegation.
     */
    function isOperator(
        address operator
    ) external view returns (bool);

    /**
     * @notice Returns the OperatorDetails struct associated with an `operator`.
     */
    function operatorDetails(
        address operator
    ) external view returns (OperatorDetails memory);

    /**
     * @notice Returns the delegationApprover account for an operator
     */
    function delegationApprover(
        address operator
    ) external view returns (address);

    /**
     * @notice Returns the shares that an operator has delegated to them in a set of strategies
     * @param operator the operator to get shares for
     * @param strategies the strategies to get shares for
     */
    function getOperatorShares(
        address operator,
        IStrategy[] memory strategies
    ) external view returns (uint256[] memory);

    /**
     * @notice Returns the shares that a set of operators have delegated to them in a set of strategies
     * @param operators the operators to get shares for
     * @param strategies the strategies to get shares for
     */
    function getOperatorsShares(
        address[] memory operators,
        IStrategy[] memory strategies
    ) external view returns (uint256[][] memory);

    /**
     * @notice Given a staker and a set of strategies, return the shares they can queue for withdrawal and the
     * corresponding depositShares.
     * This value depends on which operator the staker is delegated to.
     * The shares amount returned is the actual amount of Strategy shares the staker would receive (subject
     * to each strategy's underlying shares to token ratio).
     */
    function getWithdrawableShares(
        address staker,
        IStrategy[] memory strategies
    ) external view returns (uint256[] memory withdrawableShares, uint256[] memory depositShares);

    /**
     * @notice Returns the number of shares in storage for a staker and all their strategies
     */
    function getDepositedShares(
        address staker
    ) external view returns (IStrategy[] memory, uint256[] memory);

    /**
     * @notice Returns the minimum withdrawal delay in blocks to pass for withdrawals queued to be completable.
     * Also applies to legacy withdrawals so any withdrawals not completed prior to the slashing upgrade will be subject
     * to this longer delay.
     */
    function MIN_WITHDRAWAL_DELAY_BLOCKS() external view returns (uint32);

    /// @notice Returns a list of pending queued withdrawals for a `staker`, and the `shares` to be withdrawn.
    function getQueuedWithdrawals(
        address staker
    ) external view returns (Withdrawal[] memory withdrawals, uint256[][] memory shares);

    /// @notice Returns the keccak256 hash of `withdrawal`.
    function calculateWithdrawalRoot(
        Withdrawal memory withdrawal
    ) external pure returns (bytes32);

    /**
     * @notice Calculates the digest hash to be signed by the operator's delegationApprove and used in the `delegateTo` function.
     * @param staker The account delegating their stake
     * @param operator The account receiving delegated stake
     * @param _delegationApprover the operator's `delegationApprover` who will be signing the delegationHash (in general)
     * @param approverSalt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid
     */
    function calculateDelegationApprovalDigestHash(
        address staker,
        address operator,
        address _delegationApprover,
        bytes32 approverSalt,
        uint256 expiry
    ) external view returns (bytes32);

    /// @notice return address of the beaconChainETHStrategy
    function beaconChainETHStrategy() external view returns (IStrategy);

    /// @notice The EIP-712 typehash for the DelegationApproval struct used by the contract
    function DELEGATION_APPROVAL_TYPEHASH() external view returns (bytes32);
}
