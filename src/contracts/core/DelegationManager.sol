// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../mixins/SignatureUtilsMixin.sol";
import "../mixins/PermissionControllerMixin.sol";
import "../permissions/Pausable.sol";
import "../libraries/SlashingLib.sol";
import "../libraries/Snapshots.sol";
import "./DelegationManagerStorage.sol";

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
contract DelegationManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    DelegationManagerStorage,
    ReentrancyGuardUpgradeable,
    PermissionControllerMixin,
    SignatureUtilsMixin
{
    using SlashingLib for *;
    using Snapshots for Snapshots.DefaultZeroHistory;
    using EnumerableSet for EnumerableSet.Bytes32Set;

    // @notice Simple permission for functions that are only callable by the StrategyManager contract OR by the EigenPodManagerContract
    modifier onlyStrategyManagerOrEigenPodManager() {
        require(
            (msg.sender == address(strategyManager) || msg.sender == address(eigenPodManager)),
            OnlyStrategyManagerOrEigenPodManager()
        );
        _;
    }

    modifier onlyEigenPodManager() {
        require(msg.sender == address(eigenPodManager), OnlyEigenPodManager());
        _;
    }

    modifier onlyAllocationManager() {
        require(msg.sender == address(allocationManager), OnlyAllocationManager());
        _;
    }

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the immutable addresses of the strategy manager, eigenpod manager, and allocation manager.
     */
    constructor(
        IStrategyManager _strategyManager,
        IEigenPodManager _eigenPodManager,
        IAllocationManager _allocationManager,
        IPauserRegistry _pauserRegistry,
        IPermissionController _permissionController,
        uint32 _MIN_WITHDRAWAL_DELAY,
        string memory _version
    )
        DelegationManagerStorage(_strategyManager, _eigenPodManager, _allocationManager, _MIN_WITHDRAWAL_DELAY)
        Pausable(_pauserRegistry)
        PermissionControllerMixin(_permissionController)
        SignatureUtilsMixin(_version)
    {
        _disableInitializers();
    }

    function initialize(address initialOwner, uint256 initialPausedStatus) external initializer {
        _setPausedStatus(initialPausedStatus);
        _transferOwnership(initialOwner);
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /// @inheritdoc IDelegationManager
    function registerAsOperator(
        address initDelegationApprover,
        uint32 allocationDelay,
        string calldata metadataURI
    ) external nonReentrant {
        require(!isDelegated(msg.sender), ActivelyDelegated());

        allocationManager.setAllocationDelay(msg.sender, allocationDelay);
        _setDelegationApprover(msg.sender, initDelegationApprover);

        // delegate from the operator to themselves
        _delegate(msg.sender, msg.sender);

        emit OperatorRegistered(msg.sender, initDelegationApprover);
        emit OperatorMetadataURIUpdated(msg.sender, metadataURI);
    }

    /// @inheritdoc IDelegationManager
    function modifyOperatorDetails(
        address operator,
        address newDelegationApprover
    ) external checkCanCall(operator) nonReentrant {
        require(isOperator(operator), OperatorNotRegistered());
        _setDelegationApprover(operator, newDelegationApprover);
    }

    /// @inheritdoc IDelegationManager
    function updateOperatorMetadataURI(address operator, string calldata metadataURI) external checkCanCall(operator) {
        require(isOperator(operator), OperatorNotRegistered());
        emit OperatorMetadataURIUpdated(operator, metadataURI);
    }

    /// @inheritdoc IDelegationManager
    function delegateTo(
        address operator,
        SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt
    ) public nonReentrant {
        require(!isDelegated(msg.sender), ActivelyDelegated());
        require(isOperator(operator), OperatorNotRegistered());

        // If the operator has a `delegationApprover`, check the provided signature
        _checkApproverSignature({
            staker: msg.sender,
            operator: operator,
            signature: approverSignatureAndExpiry,
            salt: approverSalt
        });

        // Delegate msg.sender to the operator
        _delegate(msg.sender, operator);
    }

    /// @inheritdoc IDelegationManager
    function undelegate(
        address staker
    ) public nonReentrant returns (bytes32[] memory withdrawalRoots) {
        // Check that the `staker` can undelegate
        require(isDelegated(staker), NotActivelyDelegated());
        require(!isOperator(staker), OperatorsCannotUndelegate());

        // If the action is not being initiated by the staker, validate that it is initiated
        // by the operator or their delegationApprover.
        if (msg.sender != staker) {
            address operator = delegatedTo[staker];

            require(_checkCanCall(operator) || msg.sender == delegationApprover(operator), CallerCannotUndelegate());
            emit StakerForceUndelegated(staker, operator);
        }

        return _undelegate(staker);
    }

    /// @inheritdoc IDelegationManager
    function redelegate(
        address newOperator,
        SignatureWithExpiry memory newOperatorApproverSig,
        bytes32 approverSalt
    ) external returns (bytes32[] memory withdrawalRoots) {
        withdrawalRoots = undelegate(msg.sender);
        // delegateTo uses msg.sender as staker
        delegateTo(newOperator, newOperatorApproverSig, approverSalt);
    }

    /// @inheritdoc IDelegationManager
    function queueWithdrawals(
        QueuedWithdrawalParams[] calldata params
    ) external onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE) nonReentrant returns (bytes32[] memory) {
        bytes32[] memory withdrawalRoots = new bytes32[](params.length);
        address operator = delegatedTo[msg.sender];

        for (uint256 i = 0; i < params.length; i++) {
            require(params[i].strategies.length == params[i].depositShares.length, InputArrayLengthMismatch());

            uint256[] memory slashingFactors = _getSlashingFactors(msg.sender, operator, params[i].strategies);

            // Remove shares from staker's strategies and place strategies/shares in queue.
            // If the staker is delegated to an operator, the operator's delegated shares are also reduced
            // NOTE: This will fail if the staker doesn't have the shares implied by the input parameters.
            // The view function getWithdrawableShares() can be used to check what shares are available for withdrawal.
            withdrawalRoots[i] = _removeSharesAndQueueWithdrawal({
                staker: msg.sender,
                operator: operator,
                strategies: params[i].strategies,
                depositSharesToWithdraw: params[i].depositShares,
                slashingFactors: slashingFactors
            });
        }

        return withdrawalRoots;
    }

    /// @inheritdoc IDelegationManager
    function completeQueuedWithdrawal(
        Withdrawal calldata withdrawal,
        IERC20[] calldata tokens,
        bool receiveAsTokens
    ) external onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE) nonReentrant {
        _completeQueuedWithdrawal(withdrawal, tokens, receiveAsTokens);
    }

    /// @inheritdoc IDelegationManager
    function completeQueuedWithdrawals(
        Withdrawal[] calldata withdrawals,
        IERC20[][] calldata tokens,
        bool[] calldata receiveAsTokens
    ) external onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE) nonReentrant {
        uint256 n = withdrawals.length;
        for (uint256 i; i < n; ++i) {
            _completeQueuedWithdrawal(withdrawals[i], tokens[i], receiveAsTokens[i]);
        }
    }

    /// @inheritdoc IDelegationManager
    function increaseDelegatedShares(
        address staker,
        IStrategy strategy,
        uint256 prevDepositShares,
        uint256 addedShares
    ) external onlyStrategyManagerOrEigenPodManager nonReentrant {
        /// Note: Unlike `decreaseDelegatedShares`, we don't return early if the staker has no operator.
        /// This is because `_increaseDelegation` updates the staker's deposit scaling factor, which we
        /// need to do even if not delegated.
        address operator = delegatedTo[staker];
        uint64 maxMagnitude = allocationManager.getMaxMagnitude(operator, strategy);
        uint256 slashingFactor = _getSlashingFactor(staker, strategy, maxMagnitude);

        // Increase the staker's deposit scaling factor and delegate shares to the operator
        _increaseDelegation({
            operator: operator,
            staker: staker,
            strategy: strategy,
            prevDepositShares: prevDepositShares,
            addedShares: addedShares,
            slashingFactor: slashingFactor
        });
    }

    /// @inheritdoc IDelegationManager
    function decreaseDelegatedShares(
        address staker,
        uint256 curDepositShares,
        uint64 beaconChainSlashingFactorDecrease
    ) external onlyEigenPodManager nonReentrant {
        if (!isDelegated(staker)) {
            return;
        }
        address operator = delegatedTo[staker];

        // Calculate the shares to remove from the operator by calculating difference in shares
        // from the newly updated beaconChainSlashingFactor
        uint64 maxMagnitude = allocationManager.getMaxMagnitude(operator, beaconChainETHStrategy);
        DepositScalingFactor memory dsf = _depositScalingFactor[staker][beaconChainETHStrategy];
        uint256 sharesToRemove = dsf.calcWithdrawable({
            depositShares: curDepositShares,
            slashingFactor: maxMagnitude.mulWad(beaconChainSlashingFactorDecrease)
        });

        // Decrease the operator's shares
        _decreaseDelegation({
            operator: operator,
            staker: staker,
            strategy: beaconChainETHStrategy,
            sharesToDecrease: sharesToRemove
        });
    }

    /// @inheritdoc IDelegationManager
    function slashOperatorShares(
        address operator,
        IStrategy strategy,
        uint64 prevMaxMagnitude,
        uint64 newMaxMagnitude
    ) external onlyAllocationManager nonReentrant {
        /// forgefmt: disable-next-item
        uint256 operatorSharesSlashed = SlashingLib.calcSlashedAmount({
            operatorShares: operatorShares[operator][strategy],
            prevMaxMagnitude: prevMaxMagnitude,
            newMaxMagnitude: newMaxMagnitude
        });

        uint256 scaledSharesSlashedFromQueue = _getSlashableSharesInQueue({
            operator: operator,
            strategy: strategy,
            prevMaxMagnitude: prevMaxMagnitude,
            newMaxMagnitude: newMaxMagnitude
        });

        // Calculate the total deposit shares to burn - slashed operator shares plus still-slashable
        // shares sitting in the withdrawal queue.
        uint256 totalDepositSharesToBurn = operatorSharesSlashed + scaledSharesSlashedFromQueue;

        // Remove shares from operator
        _decreaseDelegation({
            operator: operator,
            staker: address(0), // we treat this as a decrease for the 0-staker (only used for events)
            strategy: strategy,
            sharesToDecrease: operatorSharesSlashed
        });

        // Emit event for operator shares being slashed
        emit OperatorSharesSlashed(operator, strategy, totalDepositSharesToBurn);

        IShareManager shareManager = _getShareManager(strategy);
        // NOTE: for beaconChainETHStrategy, increased burnable shares currently have no mechanism for burning
        shareManager.increaseBurnableShares(strategy, totalDepositSharesToBurn);
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Sets operator parameters in the `_operatorDetails` mapping.
     * @param operator The account registered as an operator updating their operatorDetails
     * @param newDelegationApprover The new parameters for the operator
     */
    function _setDelegationApprover(address operator, address newDelegationApprover) internal {
        _operatorDetails[operator].delegationApprover = newDelegationApprover;
        emit DelegationApproverUpdated(operator, newDelegationApprover);
    }

    /**
     * @notice Delegates *from* a `staker` *to* an `operator`.
     * @param staker The address to delegate *from* -- this address is delegating control of its own assets.
     * @param operator The address to delegate *to* -- this address is being given power to place the `staker`'s assets at risk on services
     * @dev Assumes the following is checked before calling this function:
     *          1) the `staker` is not already delegated to an operator
     *          2) the `operator` has indeed registered as an operator in EigenLayer
     *          3) if applicable, the `operator's` `delegationApprover` signed off on delegation
     * Ensures that:
     *          1) new delegations are not paused (PAUSED_NEW_DELEGATION)
     */
    function _delegate(address staker, address operator) internal onlyWhenNotPaused(PAUSED_NEW_DELEGATION) {
        // When a staker is not delegated to an operator, their deposit shares are equal to their
        // withdrawable shares -- except for the beaconChainETH strategy, which is handled below
        (IStrategy[] memory strategies, uint256[] memory withdrawableShares) = getDepositedShares(staker);

        // Retrieve the amount of slashing experienced by the operator in each strategy so far.
        // When delegating, we "forgive" the staker for this slashing by adjusting their
        // deposit scaling factor.
        uint256[] memory operatorSlashingFactors = _getSlashingFactors(address(0), operator, strategies);

        // Delegate to the operator
        delegatedTo[staker] = operator;
        emit StakerDelegated(staker, operator);

        for (uint256 i = 0; i < strategies.length; ++i) {
            // Special case for beacon chain slashing - ensure the staker's beacon chain slashing is
            // reflected in the number of shares they delegate.
            if (strategies[i] == beaconChainETHStrategy) {
                uint64 stakerBeaconChainSlashing = eigenPodManager.beaconChainSlashingFactor(staker);

                DepositScalingFactor memory dsf = _depositScalingFactor[staker][strategies[i]];
                withdrawableShares[i] = dsf.calcWithdrawable(withdrawableShares[i], stakerBeaconChainSlashing);
            }

            // forgefmt: disable-next-item
            _increaseDelegation({
                operator: operator, 
                staker: staker, 
                strategy: strategies[i],
                prevDepositShares: uint256(0),
                addedShares: withdrawableShares[i],
                slashingFactor: operatorSlashingFactors[i]
            });
        }
    }

    /**
     * @dev Undelegates `staker` from their operator, queueing a withdrawal for all
     * their deposited shares in the process.
     * @dev Assumes the following is checked before calling this function:
     *          1) the `staker` is currently delegated to an operator
     *          2) the `staker` is not an operator themselves
     * Ensures that:
     *          1) the withdrawal queue is not paused (PAUSED_ENTER_WITHDRAWAL_QUEUE)
     */
    function _undelegate(
        address staker
    ) internal onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE) returns (bytes32[] memory withdrawalRoots) {
        // Undelegate the staker
        address operator = delegatedTo[staker];
        delegatedTo[staker] = address(0);
        emit StakerUndelegated(staker, operator);

        // Get all of the staker's deposited strategies/shares. These will be removed from the operator
        // and queued for withdrawal.
        (IStrategy[] memory strategies, uint256[] memory depositedShares) = getDepositedShares(staker);
        if (strategies.length == 0) {
            return withdrawalRoots;
        }

        // For the operator and each of the staker's strategies, get the slashing factors to apply
        // when queueing for withdrawal
        withdrawalRoots = new bytes32[](strategies.length);
        uint256[] memory slashingFactors = _getSlashingFactors(staker, operator, strategies);

        // Queue a withdrawal for each strategy independently. This is done for UX reasons.
        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy[] memory singleStrategy = new IStrategy[](1);
            uint256[] memory singleDepositShares = new uint256[](1);
            uint256[] memory singleSlashingFactor = new uint256[](1);
            singleStrategy[0] = strategies[i];
            singleDepositShares[0] = depositedShares[i];
            singleSlashingFactor[0] = slashingFactors[i];

            // Remove shares from staker's strategies and place strategies/shares in queue.
            // The operator's delegated shares are also reduced.
            withdrawalRoots[i] = _removeSharesAndQueueWithdrawal({
                staker: staker,
                operator: operator,
                strategies: singleStrategy,
                depositSharesToWithdraw: singleDepositShares,
                slashingFactors: singleSlashingFactor
            });
        }

        return withdrawalRoots;
    }

    /**
     * @notice Removes `sharesToWithdraw` in `strategies` from `staker` who is currently delegated to `operator` and queues a withdrawal to the `withdrawer`.
     * @param staker The staker queuing a withdrawal
     * @param operator The operator the staker is delegated to
     * @param strategies The strategies to queue a withdrawal for
     * @param depositSharesToWithdraw The amount of deposit shares the staker wishes to withdraw, must be less than staker's depositShares in storage
     * @param slashingFactors The corresponding slashing factor for the staker/operator for each strategy
     *
     * @dev The amount withdrawable by the staker may not actually be the same as the depositShares that are in storage in the StrategyManager/EigenPodManager.
     * This is a result of any slashing that has occurred during the time the staker has been delegated to an operator. So the proportional amount that is withdrawn
     * out of the amount withdrawable for the staker has to also be decremented from the staker's deposit shares.
     * So the amount of depositShares withdrawn out has to be proportionally scaled down depending on the slashing that has occurred.
     * Ex. Suppose as a staker, I have 100 depositShares for a strategy thats sitting in the StrategyManager in the `stakerDepositShares` mapping but I actually have been slashed 50%
     * and my real withdrawable amount is 50 shares.
     * Now when I go to withdraw 40 depositShares, I'm proportionally withdrawing 40% of my withdrawable shares. We calculate below the actual shares withdrawn via the `toShares()` function to
     * get 20 shares to queue withdraw. The end state is that I have 60 depositShares and 30 withdrawable shares now, this still accurately reflects a 50% slashing that has occurred on my existing stake.
     * @dev depositSharesToWithdraw are converted to sharesToWithdraw using the `toShares` library function. sharesToWithdraw are then divided by the current maxMagnitude of the operator (at queue time)
     * and this value is stored in the Withdrawal struct as `scaledShares.
     * Upon completion the `scaledShares` are then multiplied by the maxMagnitude of the operator at completion time. This is how we factor in any slashing events
     * that occurred during the withdrawal delay period. Shares in a withdrawal are no longer slashable once the withdrawal is completable.
     * @dev If the `operator` is indeed an operator, then the operator's delegated shares in the `strategies` are also decreased appropriately.
     */
    function _removeSharesAndQueueWithdrawal(
        address staker,
        address operator,
        IStrategy[] memory strategies,
        uint256[] memory depositSharesToWithdraw,
        uint256[] memory slashingFactors
    ) internal returns (bytes32) {
        require(staker != address(0), InputAddressZero());
        require(strategies.length != 0, InputArrayLengthZero());

        uint256[] memory scaledShares = new uint256[](strategies.length);
        uint256[] memory withdrawableShares = new uint256[](strategies.length);

        // Remove shares from staker and operator
        // Each of these operations fail if we attempt to remove more shares than exist
        for (uint256 i = 0; i < strategies.length; ++i) {
            IShareManager shareManager = _getShareManager(strategies[i]);
            DepositScalingFactor memory dsf = _depositScalingFactor[staker][strategies[i]];

            // Calculate how many shares can be withdrawn after factoring in slashing
            withdrawableShares[i] = dsf.calcWithdrawable(depositSharesToWithdraw[i], slashingFactors[i]);

            // Scale shares for queue withdrawal
            scaledShares[i] = dsf.scaleForQueueWithdrawal(depositSharesToWithdraw[i]);

            // Remove delegated shares from the operator
            if (operator != address(0)) {
                // Staker was delegated and remains slashable during the withdrawal delay period
                // Cumulative withdrawn scaled shares are updated for the strategy, this is for accounting
                // purposes for burning shares if slashed
                _addQueuedSlashableShares(operator, strategies[i], scaledShares[i]);

                // forgefmt: disable-next-item
                _decreaseDelegation({
                    operator: operator,
                    staker: staker,
                    strategy: strategies[i],
                    sharesToDecrease: withdrawableShares[i]
                });
            }

            // Remove deposit shares from EigenPodManager/StrategyManager
            uint256 sharesAfter = shareManager.removeDepositShares(staker, strategies[i], depositSharesToWithdraw[i]);

            if (sharesAfter == 0) {
                _depositScalingFactor[staker][strategies[i]].reset();
            }
        }

        // Create queue entry and increment withdrawal nonce
        uint256 nonce = cumulativeWithdrawalsQueued[staker];
        cumulativeWithdrawalsQueued[staker]++;

        Withdrawal memory withdrawal = Withdrawal({
            staker: staker,
            delegatedTo: operator,
            withdrawer: staker,
            nonce: nonce,
            startBlock: uint32(block.number),
            strategies: strategies,
            scaledShares: scaledShares
        });

        bytes32 withdrawalRoot = calculateWithdrawalRoot(withdrawal);

        pendingWithdrawals[withdrawalRoot] = true;
        _queuedWithdrawals[withdrawalRoot] = withdrawal;
        _stakerQueuedWithdrawalRoots[staker].add(withdrawalRoot);

        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, withdrawableShares);
        return withdrawalRoot;
    }

    /**
     * @dev This function completes a queued withdrawal for a staker.
     * This will apply any slashing that has occurred since the the withdrawal was queued by multiplying the withdrawal's
     * scaledShares by the operator's maxMagnitude for each strategy. This ensures that any slashing that has occurred
     * during the period the withdrawal was queued until its slashableUntil block is applied to the withdrawal amount.
     * If receiveAsTokens is true, then these shares will be withdrawn as tokens.
     * If receiveAsTokens is false, then they will be redeposited according to the current operator the staker is delegated to,
     * and added back to the operator's delegatedShares.
     */
    function _completeQueuedWithdrawal(
        Withdrawal memory withdrawal,
        IERC20[] calldata tokens,
        bool receiveAsTokens
    ) internal {
        require(tokens.length == withdrawal.strategies.length, InputArrayLengthMismatch());
        require(msg.sender == withdrawal.withdrawer, WithdrawerNotCaller());
        bytes32 withdrawalRoot = calculateWithdrawalRoot(withdrawal);
        require(pendingWithdrawals[withdrawalRoot], WithdrawalNotQueued());

        uint256[] memory prevSlashingFactors;
        {
            // slashableUntil is block inclusive so we need to check if the current block is strictly greater than the slashableUntil block
            // meaning the withdrawal can be completed.
            uint32 slashableUntil = withdrawal.startBlock + MIN_WITHDRAWAL_DELAY_BLOCKS;
            require(uint32(block.number) > slashableUntil, WithdrawalDelayNotElapsed());

            // Given the max magnitudes of the operator the staker was originally delegated to, calculate
            // the slashing factors for each of the withdrawal's strategies.
            prevSlashingFactors = _getSlashingFactorsAtBlock({
                staker: withdrawal.staker,
                operator: withdrawal.delegatedTo,
                strategies: withdrawal.strategies,
                blockNumber: slashableUntil
            });
        }

        // Remove the withdrawal from the queue. Note that for legacy withdrawals, the removals
        // from `_stakerQueuedWithdrawalRoots` and `queuedWithdrawals` will no-op.
        _stakerQueuedWithdrawalRoots[withdrawal.staker].remove(withdrawalRoot);
        delete _queuedWithdrawals[withdrawalRoot];
        delete pendingWithdrawals[withdrawalRoot];
        emit SlashingWithdrawalCompleted(withdrawalRoot);

        // Given the max magnitudes of the operator the staker is now delegated to, calculate the current
        // slashing factors to apply to each withdrawal if it is received as shares.
        address newOperator = delegatedTo[withdrawal.staker];
        uint256[] memory newSlashingFactors = _getSlashingFactors(withdrawal.staker, newOperator, withdrawal.strategies);

        for (uint256 i = 0; i < withdrawal.strategies.length; i++) {
            IShareManager shareManager = _getShareManager(withdrawal.strategies[i]);

            // Calculate how much slashing to apply, as well as shares to withdraw
            uint256 sharesToWithdraw = SlashingLib.scaleForCompleteWithdrawal({
                scaledShares: withdrawal.scaledShares[i],
                slashingFactor: prevSlashingFactors[i]
            });

            //Do nothing if 0 shares to withdraw
            if (sharesToWithdraw == 0) {
                continue;
            }

            if (receiveAsTokens) {
                // Withdraws `shares` in `strategy` to `withdrawer`. If the shares are virtual beaconChainETH shares,
                // then a call is ultimately forwarded to the `staker`s EigenPod; otherwise a call is ultimately forwarded
                // to the `strategy` with info on the `token`.
                shareManager.withdrawSharesAsTokens({
                    staker: withdrawal.staker,
                    strategy: withdrawal.strategies[i],
                    token: tokens[i],
                    shares: sharesToWithdraw
                });
            } else {
                // Award shares back in StrategyManager/EigenPodManager.
                (uint256 prevDepositShares, uint256 addedShares) = shareManager.addShares({
                    staker: withdrawal.staker,
                    strategy: withdrawal.strategies[i],
                    shares: sharesToWithdraw
                });

                // Update the staker's deposit scaling factor and delegate shares to their operator
                _increaseDelegation({
                    operator: newOperator,
                    staker: withdrawal.staker,
                    strategy: withdrawal.strategies[i],
                    prevDepositShares: prevDepositShares,
                    addedShares: addedShares,
                    slashingFactor: newSlashingFactors[i]
                });
            }
        }
    }

    /**
     * @notice Increases `operator`s depositedShares in `strategy` based on staker's addedDepositShares
     * and updates the staker's depositScalingFactor for the strategy.
     * @param operator The operator to increase the delegated delegatedShares for
     * @param staker The staker to increase the depositScalingFactor for
     * @param strategy The strategy to increase the delegated delegatedShares and the depositScalingFactor for
     * @param prevDepositShares The number of delegated deposit shares the staker had in the strategy prior to the increase
     * @param addedShares The shares added to the staker in the StrategyManager/EigenPodManager
     * @param slashingFactor The current slashing factor for the staker/operator/strategy
     */
    function _increaseDelegation(
        address operator,
        address staker,
        IStrategy strategy,
        uint256 prevDepositShares,
        uint256 addedShares,
        uint256 slashingFactor
    ) internal {
        // Ensure that the operator has not been fully slashed for a strategy
        // and that the staker has not been fully slashed if it is the beaconChainStrategy
        // This is to prevent a divWad by 0 when updating the depositScalingFactor
        require(slashingFactor != 0, FullySlashed());

        // If `addedShares` is 0, do nothing
        if (addedShares == 0) {
            return;
        }

        // Update the staker's depositScalingFactor. This only results in an update
        // if the slashing factor has changed for this strategy.
        DepositScalingFactor storage dsf = _depositScalingFactor[staker][strategy];
        dsf.update(prevDepositShares, addedShares, slashingFactor);
        emit DepositScalingFactorUpdated(staker, strategy, dsf.scalingFactor());

        // If the staker is delegated to an operator, update the operator's shares
        if (isDelegated(staker)) {
            operatorShares[operator][strategy] += addedShares;
            emit OperatorSharesIncreased(operator, staker, strategy, addedShares);
        }
    }

    /**
     * @notice Decreases `operator`s shares in `strategy` based on staker's removed shares
     * @param operator The operator to decrease the delegated delegated shares for
     * @param staker The staker to decrease the delegated delegated shares for
     * @param strategy The strategy to decrease the delegated delegated shares for
     * @param sharesToDecrease The shares to remove from the operator's delegated shares
     */
    function _decreaseDelegation(
        address operator,
        address staker,
        IStrategy strategy,
        uint256 sharesToDecrease
    ) internal {
        // Decrement operator shares
        operatorShares[operator][strategy] -= sharesToDecrease;
        emit OperatorSharesDecreased(operator, staker, strategy, sharesToDecrease);
    }

    /// @dev If `operator` has configured a `delegationApprover`, check that `signature` and `salt`
    /// are a valid approval for `staker` delegating to `operator`.
    function _checkApproverSignature(
        address staker,
        address operator,
        SignatureWithExpiry memory signature,
        bytes32 salt
    ) internal {
        address approver = _operatorDetails[operator].delegationApprover;
        if (approver == address(0)) {
            return;
        }

        // Check that the salt hasn't been used previously, then mark the salt as spent
        require(!delegationApproverSaltIsSpent[approver][salt], SaltSpent());
        delegationApproverSaltIsSpent[approver][salt] = true;

        // Validate the signature
        _checkIsValidSignatureNow({
            signer: approver,
            signableDigest: calculateDelegationApprovalDigestHash(staker, operator, approver, salt, signature.expiry),
            signature: signature.signature,
            expiry: signature.expiry
        });
    }

    /// @dev Calculate the amount of slashing to apply to the staker's shares.
    /// @dev Be mindful of rounding in `mulWad()`, it's possible for the slashing factor to round down to 0
    /// even when both operatorMaxMagnitude and beaconChainSlashingFactor are non-zero. This is only possible
    /// in an edge case where the operator has a very low maxMagnitude.
    function _getSlashingFactor(
        address staker,
        IStrategy strategy,
        uint64 operatorMaxMagnitude
    ) internal view returns (uint256) {
        if (strategy == beaconChainETHStrategy) {
            uint64 beaconChainSlashingFactor = eigenPodManager.beaconChainSlashingFactor(staker);
            return operatorMaxMagnitude.mulWad(beaconChainSlashingFactor);
        }

        return operatorMaxMagnitude;
    }

    /// @dev Calculate the amount of slashing to apply to the staker's shares across multiple strategies
    function _getSlashingFactors(
        address staker,
        address operator,
        IStrategy[] memory strategies
    ) internal view returns (uint256[] memory) {
        uint256[] memory slashingFactors = new uint256[](strategies.length);
        uint64[] memory maxMagnitudes = allocationManager.getMaxMagnitudes(operator, strategies);

        for (uint256 i = 0; i < strategies.length; i++) {
            slashingFactors[i] = _getSlashingFactor(staker, strategies[i], maxMagnitudes[i]);
        }

        return slashingFactors;
    }

    /// @dev Calculate the amount of slashing to apply to the staker's shares across multiple strategies
    /// Note: specifically checks the operator's magnitude at a prior block, used for completing withdrawals
    function _getSlashingFactorsAtBlock(
        address staker,
        address operator,
        IStrategy[] memory strategies,
        uint32 blockNumber
    ) internal view returns (uint256[] memory) {
        uint256[] memory slashingFactors = new uint256[](strategies.length);
        uint64[] memory maxMagnitudes = allocationManager.getMaxMagnitudesAtBlock({
            operator: operator,
            strategies: strategies,
            blockNumber: blockNumber
        });

        for (uint256 i = 0; i < strategies.length; i++) {
            slashingFactors[i] = _getSlashingFactor(staker, strategies[i], maxMagnitudes[i]);
        }

        return slashingFactors;
    }

    /**
     * @dev Calculate amount of slashable shares that would be slashed from the queued withdrawals from an operator for a strategy
     * given the previous maxMagnitude and the new maxMagnitude.
     * Note: To get the total amount of slashable shares in the queue withdrawable, set newMaxMagnitude to 0 and prevMaxMagnitude
     * is the current maxMagnitude of the operator.
     */
    function _getSlashableSharesInQueue(
        address operator,
        IStrategy strategy,
        uint64 prevMaxMagnitude,
        uint64 newMaxMagnitude
    ) internal view returns (uint256) {
        // We want ALL shares added to the withdrawal queue in the window [block.number - MIN_WITHDRAWAL_DELAY_BLOCKS, block.number]
        //
        // To get this, we take the current shares in the withdrawal queue and subtract the number of shares
        // that were in the queue before MIN_WITHDRAWAL_DELAY_BLOCKS.
        uint256 curQueuedScaledShares = _cumulativeScaledSharesHistory[operator][strategy].latest();
        uint256 prevQueuedScaledShares = _cumulativeScaledSharesHistory[operator][strategy].upperLookup({
            key: uint32(block.number) - MIN_WITHDRAWAL_DELAY_BLOCKS - 1
        });

        // The difference between these values is the number of scaled shares that entered the withdrawal queue
        // less than or equal to MIN_WITHDRAWAL_DELAY_BLOCKS ago. These shares are still slashable.
        uint256 scaledSharesAdded = curQueuedScaledShares - prevQueuedScaledShares;

        return SlashingLib.scaleForBurning({
            scaledShares: scaledSharesAdded,
            prevMaxMagnitude: prevMaxMagnitude,
            newMaxMagnitude: newMaxMagnitude
        });
    }

    /// @dev Add to the cumulative withdrawn scaled shares from an operator for a given strategy
    function _addQueuedSlashableShares(address operator, IStrategy strategy, uint256 scaledShares) internal {
        uint256 currCumulativeScaledShares = _cumulativeScaledSharesHistory[operator][strategy].latest();
        _cumulativeScaledSharesHistory[operator][strategy].push({
            key: uint32(block.number),
            value: currCumulativeScaledShares + scaledShares
        });
    }

    /// @dev Get the shares from a queued withdrawal.
    function _getSharesByWithdrawalRoot(
        bytes32 withdrawalRoot
    ) internal view returns (Withdrawal memory withdrawal, uint256[] memory shares) {
        withdrawal = _queuedWithdrawals[withdrawalRoot];
        shares = new uint256[](withdrawal.strategies.length);

        uint32 slashableUntil = withdrawal.startBlock + MIN_WITHDRAWAL_DELAY_BLOCKS;

        // If the slashableUntil block is in the past, read the slashing factors at that block.
        // Otherwise, read the current slashing factors. Note that if the slashableUntil block is the current block
        // or in the future, then the slashing factors are still subject to change before the withdrawal is completable,
        // which may result in fewer shares being withdrawn.
        uint256[] memory slashingFactors = slashableUntil < uint32(block.number)
            ? _getSlashingFactorsAtBlock({
                staker: withdrawal.staker,
                operator: withdrawal.delegatedTo,
                strategies: withdrawal.strategies,
                blockNumber: slashableUntil
            })
            : _getSlashingFactors({
                staker: withdrawal.staker,
                operator: withdrawal.delegatedTo,
                strategies: withdrawal.strategies
            });

        for (uint256 j; j < withdrawal.strategies.length; ++j) {
            shares[j] = SlashingLib.scaleForCompleteWithdrawal({
                scaledShares: withdrawal.scaledShares[j],
                slashingFactor: slashingFactors[j]
            });
        }
    }

    /// @dev Depending on the strategy used, determine which ShareManager contract to make external calls to
    function _getShareManager(
        IStrategy strategy
    ) internal view returns (IShareManager) {
        return strategy == beaconChainETHStrategy
            ? IShareManager(address(eigenPodManager))
            : IShareManager(address(strategyManager));
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IDelegationManager
    function isDelegated(
        address staker
    ) public view returns (bool) {
        return (delegatedTo[staker] != address(0));
    }

    /// @inheritdoc IDelegationManager
    function isOperator(
        address operator
    ) public view returns (bool) {
        return operator != address(0) && delegatedTo[operator] == operator;
    }

    /// @inheritdoc IDelegationManager
    function delegationApprover(
        address operator
    ) public view returns (address) {
        return _operatorDetails[operator].delegationApprover;
    }

    /// @inheritdoc IDelegationManager
    function depositScalingFactor(address staker, IStrategy strategy) external view returns (uint256) {
        return _depositScalingFactor[staker][strategy].scalingFactor();
    }

    /// @inheritdoc IDelegationManager
    function getOperatorShares(
        address operator,
        IStrategy[] memory strategies
    ) public view returns (uint256[] memory) {
        uint256[] memory shares = new uint256[](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            shares[i] = operatorShares[operator][strategies[i]];
        }
        return shares;
    }

    /// @inheritdoc IDelegationManager
    function getOperatorsShares(
        address[] memory operators,
        IStrategy[] memory strategies
    ) public view returns (uint256[][] memory) {
        uint256[][] memory shares = new uint256[][](operators.length);
        for (uint256 i = 0; i < operators.length; ++i) {
            shares[i] = getOperatorShares(operators[i], strategies);
        }
        return shares;
    }

    /// @inheritdoc IDelegationManager
    function getSlashableSharesInQueue(address operator, IStrategy strategy) public view returns (uint256) {
        uint64 maxMagnitude = allocationManager.getMaxMagnitude(operator, strategy);

        // Return amount of slashable scaled shares remaining
        return _getSlashableSharesInQueue({
            operator: operator,
            strategy: strategy,
            prevMaxMagnitude: maxMagnitude,
            newMaxMagnitude: 0
        });
    }

    /// @inheritdoc IDelegationManager
    function getWithdrawableShares(
        address staker,
        IStrategy[] memory strategies
    ) public view returns (uint256[] memory withdrawableShares, uint256[] memory depositShares) {
        withdrawableShares = new uint256[](strategies.length);
        depositShares = new uint256[](strategies.length);

        // Get the slashing factors for the staker/operator/strategies
        address operator = delegatedTo[staker];
        uint256[] memory slashingFactors = _getSlashingFactors(staker, operator, strategies);

        for (uint256 i = 0; i < strategies.length; ++i) {
            IShareManager shareManager = _getShareManager(strategies[i]);
            depositShares[i] = shareManager.stakerDepositShares(staker, strategies[i]);

            // Calculate the withdrawable shares based on the slashing factor
            DepositScalingFactor memory dsf = _depositScalingFactor[staker][strategies[i]];
            withdrawableShares[i] = dsf.calcWithdrawable(depositShares[i], slashingFactors[i]);
        }

        return (withdrawableShares, depositShares);
    }

    /// @inheritdoc IDelegationManager
    function getDepositedShares(
        address staker
    ) public view returns (IStrategy[] memory, uint256[] memory) {
        // Get a list of the staker's deposited strategies/shares in the strategy manager
        (IStrategy[] memory tokenStrategies, uint256[] memory tokenDeposits) = strategyManager.getDeposits(staker);

        // If the staker has no beacon chain ETH shares, return any shares from the strategy manager
        uint256 podOwnerShares = eigenPodManager.stakerDepositShares(staker, beaconChainETHStrategy);
        if (podOwnerShares == 0) {
            return (tokenStrategies, tokenDeposits);
        }

        // Allocate extra space for beaconChainETHStrategy and shares
        IStrategy[] memory strategies = new IStrategy[](tokenStrategies.length + 1);
        uint256[] memory shares = new uint256[](tokenStrategies.length + 1);

        strategies[tokenStrategies.length] = beaconChainETHStrategy;
        shares[tokenStrategies.length] = podOwnerShares;

        // Copy any strategy manager shares to complete array
        for (uint256 i = 0; i < tokenStrategies.length; i++) {
            strategies[i] = tokenStrategies[i];
            shares[i] = tokenDeposits[i];
        }

        return (strategies, shares);
    }

    function queuedWithdrawals(
        bytes32 withdrawalRoot
    ) external view returns (Withdrawal memory withdrawal) {
        return _queuedWithdrawals[withdrawalRoot];
    }

    /// @inheritdoc IDelegationManager
    function getQueuedWithdrawal(
        bytes32 withdrawalRoot
    ) external view returns (Withdrawal memory withdrawal, uint256[] memory shares) {
        (withdrawal, shares) = _getSharesByWithdrawalRoot(withdrawalRoot);
    }

    /// @inheritdoc IDelegationManager
    function getQueuedWithdrawals(
        address staker
    ) external view returns (Withdrawal[] memory withdrawals, uint256[][] memory shares) {
        bytes32[] memory withdrawalRoots = getQueuedWithdrawalRoots(staker);

        uint256 totalQueued = withdrawalRoots.length;
        withdrawals = new Withdrawal[](totalQueued);
        shares = new uint256[][](totalQueued);

        for (uint256 i; i < totalQueued; ++i) {
            (withdrawals[i], shares[i]) = _getSharesByWithdrawalRoot(withdrawalRoots[i]);
        }
    }

    /// @inheritdoc IDelegationManager
    function getQueuedWithdrawalRoots(
        address staker
    ) public view returns (bytes32[] memory) {
        return _stakerQueuedWithdrawalRoots[staker].values();
    }

    /// @inheritdoc IDelegationManager
    function convertToDepositShares(
        address staker,
        IStrategy[] memory strategies,
        uint256[] memory withdrawableShares
    ) external view returns (uint256[] memory) {
        // Get the slashing factors for the staker/operator/strategies
        address operator = delegatedTo[staker];
        uint256[] memory slashingFactors = _getSlashingFactors(staker, operator, strategies);

        // Calculate the deposit shares based on the slashing factor
        uint256[] memory depositShares = new uint256[](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            DepositScalingFactor memory dsf = _depositScalingFactor[staker][strategies[i]];
            depositShares[i] = dsf.calcDepositShares(withdrawableShares[i], slashingFactors[i]);
        }
        return depositShares;
    }

    /// @inheritdoc IDelegationManager
    function calculateWithdrawalRoot(
        Withdrawal memory withdrawal
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(withdrawal));
    }

    /// @inheritdoc IDelegationManager
    function minWithdrawalDelayBlocks() external view returns (uint32) {
        return MIN_WITHDRAWAL_DELAY_BLOCKS;
    }

    /// @inheritdoc IDelegationManager
    function calculateDelegationApprovalDigestHash(
        address staker,
        address operator,
        address approver,
        bytes32 approverSalt,
        uint256 expiry
    ) public view returns (bytes32) {
        /// forgefmt: disable-next-item
        return _calculateSignableDigest(
            keccak256(
                abi.encode(
                    DELEGATION_APPROVAL_TYPEHASH, 
                    approver, 
                    staker, 
                    operator, 
                    approverSalt, 
                    expiry
                )
            )
        );
    }
}
