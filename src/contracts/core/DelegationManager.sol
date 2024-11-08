// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../mixins/SignatureUtils.sol";
import "../mixins/PermissionControllerMixin.sol";
import "../permissions/Pausable.sol";
import "../libraries/SlashingLib.sol";
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
    SignatureUtils,
    PermissionControllerMixin
{
    using SlashingLib for *;
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
     * @dev Initializes the immutable addresses of the strategy mananger, eigenpod manager, and allocation manager.
     */
    constructor(
        IAVSDirectory _avsDirectory,
        IStrategyManager _strategyManager,
        IEigenPodManager _eigenPodManager,
        IAllocationManager _allocationManager,
        IPauserRegistry _pauserRegistry,
        IPermissionController _permissionController,
        uint32 _MIN_WITHDRAWAL_DELAY
    )
        DelegationManagerStorage(
            _avsDirectory,
            _strategyManager,
            _eigenPodManager,
            _allocationManager,
            _MIN_WITHDRAWAL_DELAY
        )
        Pausable(_pauserRegistry)
        PermissionControllerMixin(_permissionController)
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
        OperatorDetails calldata registeringOperatorDetails,
        uint32 allocationDelay,
        string calldata metadataURI
    ) external {
        require(!isDelegated(msg.sender), ActivelyDelegated());

        allocationManager.setAllocationDelay(msg.sender, allocationDelay);
        _setOperatorDetails(msg.sender, registeringOperatorDetails);

        // delegate from the operator to themselves
        SignatureWithExpiry memory emptySignatureAndExpiry;
        _delegate(msg.sender, msg.sender, emptySignatureAndExpiry, bytes32(0));

        emit OperatorRegistered(msg.sender, registeringOperatorDetails);
        emit OperatorMetadataURIUpdated(msg.sender, metadataURI);
    }

    /// @inheritdoc IDelegationManager
    function modifyOperatorDetails(
        address operator,
        OperatorDetails calldata newOperatorDetails
    ) external checkCanCall(operator) {
        require(isOperator(operator), OperatorNotRegistered());
        _setOperatorDetails(operator, newOperatorDetails);
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
    ) external {
        require(!isDelegated(msg.sender), ActivelyDelegated());
        require(isOperator(operator), OperatorNotRegistered());

        // go through the internal delegation flow, checking the `approverSignatureAndExpiry` if applicable
        _delegate(msg.sender, operator, approverSignatureAndExpiry, approverSalt);
    }

    /// @inheritdoc IDelegationManager
    function undelegate(
        address staker
    ) external onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE) returns (bytes32[] memory withdrawalRoots) {
        require(isDelegated(staker), NotActivelyDelegated());
        require(!isOperator(staker), OperatorsCannotUndelegate());
        require(staker != address(0), InputAddressZero());
        address operator = delegatedTo[staker];
        require(
            msg.sender == staker || _checkCanCall(operator)
                || msg.sender == _operatorDetails[operator].delegationApprover,
            CallerCannotUndelegate()
        );

        // Gather strategies and shares from the staker. Calculate depositedShares to remove from operator during undelegation
        // Undelegation removes ALL currently-active strategies and shares
        (IStrategy[] memory strategies, uint256[] memory depositedShares) = getDepositedShares(staker);

        // Undelegate the staker
        delegatedTo[staker] = address(0);
        emit StakerUndelegated(staker, operator);
        // Emit an event if this action was not initiated by the staker themselves
        if (msg.sender != staker) {
            emit StakerForceUndelegated(staker, operator);
        }

        // if no deposited shares, return an empty array, and don't queue a withdrawal
        if (strategies.length == 0) {
            return withdrawalRoots;
        }

        withdrawalRoots = new bytes32[](strategies.length);
        uint64[] memory maxMagnitudes = allocationManager.getMaxMagnitudes(operator, strategies);

        for (uint256 i = 0; i < strategies.length; i++) {
            StakerScalingFactors storage ssf = stakerScalingFactor[staker][strategies[i]];

            // If the operator was not slashed 100% for the strategy and the staker has not been fully slashed
            // for native restaking (if the strategy is beaconChainStrategy) then handle a normal queued withdrawal.
            // Otherwise if the operator has been slashed 100% for the strategy, it implies
            // the staker has no available shares to withdraw and we simply decrement their entire depositShares amount.
            // Note the returned withdrawal root will be 0x0 in this scenario but is not actually a valid null root.
            if (!ssf.isFullySlashed(maxMagnitudes[i])) {
                IStrategy[] memory singleStrategy = new IStrategy[](1);
                uint256[] memory singleDepositShares = new uint256[](1);
                uint64[] memory singleMaxMagnitude = new uint64[](1);
                singleStrategy[0] = strategies[i];
                singleDepositShares[0] = depositedShares[i];
                singleMaxMagnitude[0] = maxMagnitudes[i];

                withdrawalRoots[i] = _removeSharesAndQueueWithdrawal({
                    staker: staker,
                    operator: operator,
                    strategies: singleStrategy,
                    depositSharesToWithdraw: singleDepositShares,
                    maxMagnitudes: singleMaxMagnitude
                });
            } else {
                IShareManager shareManager = _getShareManager(strategies[i]);

                // Remove active shares from EigenPodManager/StrategyManager
                // This is to ensure that all shares are removed entirely and cannot be withdrawn
                // or redelegated.
                shareManager.removeDepositShares(staker, strategies[i], depositedShares[i]);
            }

            // all shares are queued withdrawn with no delegated operator, so
            // reset staker's depositScalingFactor back to WAD default.
            // If this is not reset, the depositScalingFactor would be incorrect
            // when the staker deposits and queue withdraws in the future.
            ssf.depositScalingFactor = WAD;
            emit DepositScalingFactorUpdated(staker, strategies[i], WAD);
        }

        return withdrawalRoots;
    }

    /// @inheritdoc IDelegationManager
    function queueWithdrawals(
        QueuedWithdrawalParams[] calldata params
    ) external onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE) returns (bytes32[] memory) {
        bytes32[] memory withdrawalRoots = new bytes32[](params.length);
        address operator = delegatedTo[msg.sender];

        for (uint256 i = 0; i < params.length; i++) {
            require(params[i].strategies.length == params[i].depositShares.length, InputArrayLengthMismatch());
            require(params[i].withdrawer == msg.sender, WithdrawerNotStaker());

            uint64[] memory maxMagnitudes = allocationManager.getMaxMagnitudes(operator, params[i].strategies);

            // Remove shares from staker's strategies and place strategies/shares in queue.
            // If the staker is delegated to an operator, the operator's delegated shares are also reduced
            // NOTE: This will fail if the staker doesn't have the shares implied by the input parameters.
            // The view function getWithdrawableShares() can be used to check what shares are available for withdrawal.
            withdrawalRoots[i] = _removeSharesAndQueueWithdrawal({
                staker: msg.sender,
                operator: operator,
                strategies: params[i].strategies,
                depositSharesToWithdraw: params[i].depositShares,
                maxMagnitudes: maxMagnitudes
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
    function completeQueuedWithdrawals(
        IERC20[][] calldata tokens,
        bool[] calldata receiveAsTokens,
        uint256 numToComplete
    ) external onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE) nonReentrant {
        EnumerableSet.Bytes32Set storage withdrawalRoots = _stakerQueuedWithdrawalRoots[msg.sender];
        uint256 totalQueued = withdrawalRoots.length();
        numToComplete = numToComplete > totalQueued ? totalQueued : numToComplete;
        for (uint256 i; i < numToComplete; ++i) {
            _completeQueuedWithdrawal(queuedWithdrawals[withdrawalRoots.at(i)], tokens[i], receiveAsTokens[i]);
        }
    }

    /// @inheritdoc IDelegationManager
    function increaseDelegatedShares(
        address staker,
        IStrategy strategy,
        uint256 existingDepositShares,
        uint256 addedShares
    ) external onlyStrategyManagerOrEigenPodManager {
        // if the staker is delegated to an operator
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];
            IStrategy[] memory strategies = new IStrategy[](1);
            strategies[0] = strategy;
            uint64[] memory maxMagnitudes = allocationManager.getMaxMagnitudes(operator, strategies);

            // add deposit shares to operator's stake shares and update the staker's depositScalingFactor
            _increaseDelegation({
                operator: operator,
                staker: staker,
                strategy: strategy,
                existingDepositShares: existingDepositShares,
                addedShares: addedShares,
                maxMagnitude: maxMagnitudes[0]
            });
        }
    }

    /// @inheritdoc IDelegationManager
    function decreaseBeaconChainScalingFactor(
        address staker,
        uint256 existingDepositShares,
        uint64 proportionOfOldBalance
    ) external onlyEigenPodManager {
        // decrease the staker's beaconChainScalingFactor proportionally
        address operator = delegatedTo[staker];
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = beaconChainETHStrategy;
        uint64[] memory maxMagnitudes = allocationManager.getMaxMagnitudes(operator, strategies);

        StakerScalingFactors storage ssf = stakerScalingFactor[staker][beaconChainETHStrategy];
        uint256 sharesBefore = existingDepositShares.toShares(ssf, maxMagnitudes[0]);
        ssf.decreaseBeaconChainScalingFactor(proportionOfOldBalance);
        emit BeaconChainScalingFactorDecreased(staker, ssf.beaconChainScalingFactor);
        uint256 sharesAfter = existingDepositShares.toShares(ssf, maxMagnitudes[0]);

        // if the staker is delegated to an operators
        if (isDelegated(staker)) {
            // subtract strategy shares from delegated scaled shares
            _decreaseDelegation({
                operator: operator,
                staker: staker,
                strategy: beaconChainETHStrategy,
                sharesToDecrease: sharesBefore - sharesAfter
            });
        }
    }

    /// @inheritdoc IDelegationManager
    function decreaseOperatorShares(
        address operator,
        IStrategy strategy,
        uint256 wadSlashed
    ) external onlyAllocationManager {
        /// forgefmt: disable-next-item
        uint256 amountSlashed = SlashingLib.calcSlashedAmount({
            operatorShares: operatorShares[operator][strategy], 
            wadSlashed: wadSlashed
        });

        _decreaseDelegation({
            operator: operator,
            staker: address(0), // we treat this as a decrease for the zero address staker
            strategy: strategy,
            sharesToDecrease: amountSlashed
        });
    }

    /**
     *
     *                         BACKWARDS COMPATIBLE LEGACY FUNCTIONS
     *                         TO BE DEPRECATED IN FUTURE
     *
     */

    /// @inheritdoc IDelegationManager
    function completeQueuedWithdrawal(
        Withdrawal calldata withdrawal,
        IERC20[] calldata tokens,
        uint256, // middlewareTimesIndex
        bool receiveAsTokens
    ) external onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE) nonReentrant {
        _completeQueuedWithdrawal(withdrawal, tokens, receiveAsTokens);
    }

    /// @inheritdoc IDelegationManager
    function completeQueuedWithdrawals(
        Withdrawal[] calldata withdrawals,
        IERC20[][] calldata tokens,
        uint256[] calldata, // middlewareTimesIndexes
        bool[] calldata receiveAsTokens
    ) external onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE) nonReentrant {
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            _completeQueuedWithdrawal(withdrawals[i], tokens[i], receiveAsTokens[i]);
        }
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Sets operator parameters in the `_operatorDetails` mapping.
     * @param operator The account registered as an operator updating their operatorDetails
     * @param newOperatorDetails The new parameters for the operator
     */
    function _setOperatorDetails(address operator, OperatorDetails calldata newOperatorDetails) internal {
        _operatorDetails[operator] = newOperatorDetails;
        emit OperatorDetailsModified(msg.sender, newOperatorDetails);
    }

    /**
     * @notice Delegates *from* a `staker` *to* an `operator`.
     * @param staker The address to delegate *from* -- this address is delegating control of its own assets.
     * @param operator The address to delegate *to* -- this address is being given power to place the `staker`'s assets at risk on services
     * @param approverSignatureAndExpiry Verifies the operator approves of this delegation
     * @param approverSalt Is a salt used to help guarantee signature uniqueness. Each salt can only be used once by a given approver.
     * @dev Assumes the following is checked before calling this function:
     *          1) the `staker` is not already delegated to an operator
     *          2) the `operator` has indeed registered as an operator in EigenLayer
     * Ensures that:
     *          1) if applicable, that the approver signature is valid and non-expired
     *          2) new delegations are not paused (PAUSED_NEW_DELEGATION)
     */
    function _delegate(
        address staker,
        address operator,
        SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt
    ) internal onlyWhenNotPaused(PAUSED_NEW_DELEGATION) {
        // fetch the operator's `delegationApprover` address and store it in memory in case we need to use it multiple times
        address approver = _operatorDetails[operator].delegationApprover;
        /**
         * Check the `approver`'s signature, if applicable.
         * If the `approver` is the zero address, then the operator allows all stakers to delegate to them and this verification is skipped.
         * If the `approver` or the `operator` themselves is the caller, then approval is assumed and signature verification is skipped as well.
         */
        if (approver != address(0) && msg.sender != approver && msg.sender != operator) {
            // check that the salt hasn't been used previously, then mark the salt as spent
            require(!delegationApproverSaltIsSpent[approver][approverSalt], SaltSpent());
            // actually check that the signature is valid
            _checkIsValidSignatureNow({
                signer: approver,
                signableDigest: calculateDelegationApprovalDigestHash(
                    staker, operator, approver, approverSalt, approverSignatureAndExpiry.expiry
                ),
                signature: approverSignatureAndExpiry.signature,
                expiry: approverSignatureAndExpiry.expiry
            });

            delegationApproverSaltIsSpent[approver][approverSalt] = true;
        }

        // record the delegation relation between the staker and operator, and emit an event
        delegatedTo[staker] = operator;
        emit StakerDelegated(staker, operator);

        // read staker's deposited shares and strategies to add to operator's shares
        // and also update the staker depositScalingFactor for each strategy
        (IStrategy[] memory strategies, uint256[] memory depositedShares) = getDepositedShares(staker);
        uint64[] memory maxMagnitudes = allocationManager.getMaxMagnitudes(operator, strategies);

        for (uint256 i = 0; i < strategies.length; ++i) {
            // forgefmt: disable-next-item
            _increaseDelegation({
                operator: operator, 
                staker: staker, 
                strategy: strategies[i],
                existingDepositShares: uint256(0),
                addedShares: depositedShares[i],
                maxMagnitude: maxMagnitudes[i]
            });
        }
    }

    /**
     * @dev This function completes a queued withdrawal for a staker.
     * This will apply any slashing that has occurred since the the withdrawal was queued by multiplying the withdrawal's
     * scaledShares by the operator's maxMagnitude for each strategy. This ensures that any slashing that has occurred
     * during the period the withdrawal was queued until its completable timestamp is applied to the withdrawal amount.
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

        uint32 completableBlock = withdrawal.startBlock + MIN_WITHDRAWAL_DELAY_BLOCKS;
        require(completableBlock <= uint32(block.number), WithdrawalDelayNotElapsed());

        // read delegated operator's maxMagnitudes at the earliest time that the withdrawal could be completed
        // to convert the delegatedShares to shares factoring in slashing that occured during withdrawal delay
        uint64[] memory maxMagnitudes = allocationManager.getMaxMagnitudesAtBlock({
            operator: withdrawal.delegatedTo,
            strategies: withdrawal.strategies,
            blockNumber: completableBlock
        });

        for (uint256 i = 0; i < withdrawal.strategies.length; i++) {
            IShareManager shareManager = _getShareManager(withdrawal.strategies[i]);
            uint256 sharesToWithdraw = withdrawal.scaledShares[i].scaleSharesForCompleteWithdrawal(
                stakerScalingFactor[withdrawal.staker][withdrawal.strategies[i]], maxMagnitudes[i]
            );

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
                shareManager.addShares({
                    staker: withdrawal.staker,
                    strategy: withdrawal.strategies[i],
                    token: tokens[i],
                    shares: sharesToWithdraw
                });
            }
        }

        _stakerQueuedWithdrawalRoots[withdrawal.staker].remove(withdrawalRoot);

        delete queuedWithdrawals[withdrawalRoot];

        delete pendingWithdrawals[withdrawalRoot];

        emit SlashingWithdrawalCompleted(withdrawalRoot);
    }

    /**
     * @notice Increases `operator`s depositedShares in `strategy` based on staker's addedDepositShares
     * and updates the staker's depositScalingFactor for the strategy.
     * @param operator The operator to increase the delegated delegatedShares for
     * @param staker The staker to increase the depositScalingFactor for
     * @param strategy The strategy to increase the delegated delegatedShares and the depositScalingFactor for
     * @param existingDepositShares The number of deposit shares the staker already has in the strategy.
     * @param addedShares The shares added to the staker in the StrategyManager/EigenPodManager
     * @param maxMagnitude The current max magnitude of the operator for the strategy
     */
    function _increaseDelegation(
        address operator,
        address staker,
        IStrategy strategy,
        uint256 existingDepositShares,
        uint256 addedShares,
        uint64 maxMagnitude
    ) internal {
        StakerScalingFactors storage ssf = stakerScalingFactor[staker][strategy];

        // Ensure that the operator has not been fully slashed for a strategy
        // and that the staker has not been fully slashed if its the beaconChainStrategy
        require(!ssf.isFullySlashed(maxMagnitude), FullySlashed());

        // Increment operator shares
        operatorShares[operator][strategy] += addedShares;
        emit OperatorSharesIncreased(operator, staker, strategy, addedShares);

        // update the staker's depositScalingFactor
        ssf.updateDepositScalingFactor(existingDepositShares, addedShares, maxMagnitude);
        emit DepositScalingFactorUpdated(staker, strategy, ssf.depositScalingFactor);
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

    /**
     * @notice Removes `sharesToWithdraw` in `strategies` from `staker` who is currently delegated to `operator` and queues a withdrawal to the `withdrawer`.
     * @param staker The staker queuing a withdrawal
     * @param operator The operator the staker is delegated to
     * @param strategies The strategies to queue a withdrawal for
     * @param depositSharesToWithdraw The amount of deposit shares the staker wishes to withdraw, must be less than staker's depositShares in storage
     * @param maxMagnitudes The corresponding maxMagnitudes of the operator for the respective strategies
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
        uint64[] memory maxMagnitudes
    ) internal returns (bytes32) {
        require(staker != address(0), InputAddressZero());
        require(strategies.length != 0, InputArrayLengthZero());

        uint256[] memory scaledShares = new uint256[](strategies.length);
        uint256[] memory sharesToWithdraw = new uint256[](strategies.length);

        // Remove shares from staker and operator
        // Each of these operations fail if we attempt to remove more shares than exist
        for (uint256 i = 0; i < strategies.length; ++i) {
            IShareManager shareManager = _getShareManager(strategies[i]);
            StakerScalingFactors memory ssf = stakerScalingFactor[staker][strategies[i]];

            // Ensure that the operator has not been fully slashed for a strategy
            // and that the staker has not been slashed fully if its the beaconChainStrategy
            require(!ssf.isFullySlashed(maxMagnitudes[i]), FullySlashed());
            // Check withdrawing deposit shares amount doesn't exceed balance
            require(
                depositSharesToWithdraw[i] <= shareManager.stakerDepositShares(staker, strategies[i]),
                WithdrawalExceedsMax()
            );

            // Calculate the shares to withdraw
            sharesToWithdraw[i] = depositSharesToWithdraw[i].toShares(ssf, maxMagnitudes[i]);

            // Remove delegated shares from the operator
            if (operator != address(0)) {
                // forgefmt: disable-next-item
                _decreaseDelegation({
                    operator: operator,
                    staker: staker,
                    strategy: strategies[i],
                    sharesToDecrease: sharesToWithdraw[i]
                });
            }

            scaledShares[i] = sharesToWithdraw[i].scaleSharesForQueuedWithdrawal(ssf, maxMagnitudes[i]);

            // Remove active shares from EigenPodManager/StrategyManager
            shareManager.removeDepositShares(staker, strategies[i], depositSharesToWithdraw[i]);
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

        _stakerQueuedWithdrawalRoots[staker].add(withdrawalRoot);

        queuedWithdrawals[withdrawalRoot] = withdrawal;

        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, sharesToWithdraw);
        return withdrawalRoot;
    }

    /**
     *
     *                         SHARES CONVERSION FUNCTIONS
     *
     */

    /// @notice Depending on the strategy used, determine which ShareManager contract to make external calls to
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
    function operatorDetails(
        address operator
    ) external view returns (OperatorDetails memory) {
        return _operatorDetails[operator];
    }

    /// @inheritdoc IDelegationManager
    function delegationApprover(
        address operator
    ) external view returns (address) {
        return _operatorDetails[operator].delegationApprover;
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
    function getWithdrawableShares(
        address staker,
        IStrategy[] memory strategies
    ) public view returns (uint256[] memory withdrawableShares, uint256[] memory depositShares) {
        address operator = delegatedTo[staker];
        uint64[] memory maxMagnitudes = allocationManager.getMaxMagnitudes(operator, strategies);
        withdrawableShares = new uint256[](strategies.length);
        depositShares = new uint256[](strategies.length);

        for (uint256 i = 0; i < strategies.length; ++i) {
            IShareManager shareManager = _getShareManager(strategies[i]);
            // TODO: batch call for strategyManager shares?
            // 1. read strategy deposit shares
            // forgefmt: disable-next-item
            depositShares[i] = shareManager.stakerDepositShares(staker, strategies[i]);

            // 2. Calculate the withdrawable shares
            withdrawableShares[i] =
                depositShares[i].toShares(stakerScalingFactor[staker][strategies[i]], maxMagnitudes[i]);
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

    /// @inheritdoc IDelegationManager
    function getQueuedWithdrawals(
        address staker
    ) external view returns (Withdrawal[] memory withdrawals, uint256[][] memory shares) {
        bytes32[] memory withdrawalRoots = _stakerQueuedWithdrawalRoots[staker].values();
        uint256 totalQueued = withdrawalRoots.length;

        withdrawals = new Withdrawal[](totalQueued);
        shares = new uint256[][](totalQueued);

        address operator = delegatedTo[staker];

        for (uint256 i; i < totalQueued; ++i) {
            withdrawals[i] = queuedWithdrawals[withdrawalRoots[i]];

            uint64[] memory operatorMagnitudes = allocationManager.getMaxMagnitudes(operator, withdrawals[i].strategies);

            for (uint256 j; j < withdrawals[i].strategies.length; ++j) {
                StakerScalingFactors memory ssf = stakerScalingFactor[staker][withdrawals[i].strategies[j]];

                shares[i][j] =
                    withdrawals[i].scaledShares[j].scaleSharesForCompleteWithdrawal(ssf, operatorMagnitudes[i]);
            }
        }
    }

    /// @inheritdoc IDelegationManager
    function calculateWithdrawalRoot(
        Withdrawal memory withdrawal
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(withdrawal));
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
