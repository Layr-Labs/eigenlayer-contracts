// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../mixins/SignatureUtils.sol";
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
    SignatureUtils
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
     * @dev Initializes the immutable addresses of the strategy mananger, eigenpod manager, and allocation manager.
     */
    constructor(
        IAVSDirectory _avsDirectory,
        IStrategyManager _strategyManager,
        IEigenPodManager _eigenPodManager,
        IAllocationManager _allocationManager,
        IPauserRegistry _pauserRegistry,
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
        _delegate(msg.sender, msg.sender);

        emit OperatorRegistered(msg.sender, registeringOperatorDetails);
        emit OperatorMetadataURIUpdated(msg.sender, metadataURI);
    }

    /// @inheritdoc IDelegationManager
    function modifyOperatorDetails(
        OperatorDetails calldata newOperatorDetails
    ) external {
        require(isOperator(msg.sender), OperatorNotRegistered());
        _setOperatorDetails(msg.sender, newOperatorDetails);
    }

    /// @inheritdoc IDelegationManager
    function updateOperatorMetadataURI(
        string calldata metadataURI
    ) external {
        require(isOperator(msg.sender), OperatorNotRegistered());
        emit OperatorMetadataURIUpdated(msg.sender, metadataURI);
    }

    /// @inheritdoc IDelegationManager
    function delegateTo(
        address operator,
        SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt
    ) external {
        require(!isDelegated(msg.sender), ActivelyDelegated());
        require(isOperator(operator), OperatorNotRegistered());

        // Checking the `approverSignatureAndExpiry` if applicable
        address approver = _operatorDetails[operator].delegationApprover;
        if (approver != address(0) && msg.sender != approver && msg.sender != operator) {
            // check that the salt hasn't been used previously, then mark the salt as spent
            require(!delegationApproverSaltIsSpent[approver][approverSalt], SaltSpent());
            // actually check that the signature is valid
            _checkIsValidSignatureNow({
                signer: approver,
                signableDigest: calculateDelegationApprovalDigestHash(
                    msg.sender, operator, approver, approverSalt, approverSignatureAndExpiry.expiry
                ),
                signature: approverSignatureAndExpiry.signature,
                expiry: approverSignatureAndExpiry.expiry
            });

            delegationApproverSaltIsSpent[approver][approverSalt] = true;
        }

        // Delegate msg.sender to the operator
        _delegate(msg.sender, operator);
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
            msg.sender == staker || msg.sender == operator
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
        uint256[] memory slashingFactors = _getSlashingFactors(staker, operator, strategies);

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

    /// @inheritdoc IDelegationManager
    function queueWithdrawals(
        QueuedWithdrawalParams[] calldata params
    ) external onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE) returns (bytes32[] memory) {
        bytes32[] memory withdrawalRoots = new bytes32[](params.length);
        address operator = delegatedTo[msg.sender];

        for (uint256 i = 0; i < params.length; i++) {
            require(params[i].strategies.length == params[i].depositShares.length, InputArrayLengthMismatch());
            require(params[i].withdrawer == msg.sender, WithdrawerNotStaker());

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

    // /// @inheritdoc IDelegationManager
    // function completeQueuedWithdrawals(
    //     IERC20[][] calldata tokens,
    //     bool[] calldata receiveAsTokens,
    //     uint256 numToComplete
    // ) external onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE) nonReentrant {
    //     EnumerableSet.Bytes32Set storage withdrawalRoots = _stakerQueuedWithdrawalRoots[msg.sender];
    //     uint256 length = withdrawalRoots.length();
    //     numToComplete = numToComplete > length ? length : numToComplete;

    //     // Read withdrawals to complete. We use 2 seperate loops here because the second
    //     // loop will remove elements by index from `withdrawalRoots`.
    //     Withdrawal[] memory withdrawals = new Withdrawal[](numToComplete);
    //     for (uint256 i; i < withdrawals.length; ++i) {
    //         withdrawals[i] = queuedWithdrawals[withdrawalRoots.at(i)];
    //     }

    //     for (uint256 i; i < withdrawals.length; ++i) {
    //         _completeQueuedWithdrawal(withdrawals[i], tokens[i], receiveAsTokens[i]);
    //     }
    // }

    /// @inheritdoc IDelegationManager
    function increaseDelegatedShares(
        address staker,
        IStrategy strategy,
        uint256 existingDepositShares,
        uint256 addedShares
    ) external onlyStrategyManagerOrEigenPodManager {
        address operator = delegatedTo[staker];
        uint64 maxMagnitude = allocationManager.getMaxMagnitude(operator, strategy);
        uint256 slashingFactor = _getSlashingFactor(staker, strategy, maxMagnitude);

        // Increase the staker's deposit scaling factor and delegate shares to the operator
        _increaseDelegation({
            operator: operator,
            staker: staker,
            strategy: strategy,
            existingDepositShares: existingDepositShares,
            addedShares: addedShares,
            slashingFactor: slashingFactor
        });
    }

    /// @inheritdoc IDelegationManager
    function decreaseBeaconChainScalingFactor(
        address staker,
        uint256 existingDepositShares,
        uint64 proportionOfOldBalance
    ) external onlyEigenPodManager {
        // decrease the staker's beaconChainScalingFactor proportionally
        address operator = delegatedTo[staker];
        uint64 maxMagnitude = allocationManager.getMaxMagnitude(operator, beaconChainETHStrategy);

        DepositScalingFactor memory dsf = _depositScalingFactor[staker][beaconChainETHStrategy];

        uint256 slashingFactor = _getSlashingFactor(staker, beaconChainETHStrategy, maxMagnitude);
        uint256 withdrawableBefore = dsf.calcWithdrawable(existingDepositShares, slashingFactor);

        _decreaseBeaconChainSlashingFactor(staker, proportionOfOldBalance);

        slashingFactor = _getSlashingFactor(staker, beaconChainETHStrategy, maxMagnitude);
        uint256 withdrawableAfter = dsf.calcWithdrawable(existingDepositShares, slashingFactor);

        // if the staker is delegated to an operators
        if (isDelegated(staker)) {
            // subtract strategy shares from delegated scaled shares
            _decreaseDelegation({
                operator: operator,
                staker: staker,
                strategy: beaconChainETHStrategy,
                sharesToDecrease: withdrawableBefore - withdrawableAfter
            });
        }
    }

    /// @inheritdoc IDelegationManager
    function burnOperatorShares(
        address operator,
        IStrategy strategy,
        uint64 prevMaxMagnitude,
        uint64 newMaxMagnitude
    ) external onlyAllocationManager {
        require(newMaxMagnitude < prevMaxMagnitude, MaxMagnitudeCantIncrease());

        /// forgefmt: disable-next-item
        uint256 sharesToDecrement = SlashingLib.calcSlashedAmount({
            operatorShares: operatorShares[operator][strategy],
            prevMaxMagnitude: prevMaxMagnitude,
            newMaxMagnitude: newMaxMagnitude
        });

        // While `sharesToDecrement` describes the amount we should directly remove from the operator's delegated
        // shares, `sharesToBurn` also includes any shares that have been queued for withdrawal and are still
        // slashable given the withdrawal delay.
        uint256 sharesToBurn =
            sharesToDecrement + _getSlashedSharesInQueue(operator, strategy, prevMaxMagnitude, newMaxMagnitude);

        // Remove shares from operator
        _decreaseDelegation({
            operator: operator,
            staker: address(0), // we treat this as a decrease for the zero address staker
            strategy: strategy,
            sharesToDecrease: sharesToDecrement
        });

        /// TODO: implement EPM.burnShares interface. Likely requires more complex interface than just shares
        /// so not adding a burnShares method in IShareManager
        if (strategy != beaconChainETHStrategy) {
            strategyManager.burnShares(strategy, sharesToBurn);
            emit OperatorSharesBurned(operator, strategy, sharesToBurn);
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
     * @dev Assumes the following is checked before calling this function:
     *          1) the `staker` is not already delegated to an operator
     *          2) the `operator` has indeed registered as an operator in EigenLayer
     * Ensures that:
     *          1) if applicable, that the approver signature is valid and non-expired
     *          2) new delegations are not paused (PAUSED_NEW_DELEGATION)
     */
    function _delegate(address staker, address operator) internal onlyWhenNotPaused(PAUSED_NEW_DELEGATION) {
        // record the delegation relation between the staker and operator, and emit an event
        delegatedTo[staker] = operator;
        emit StakerDelegated(staker, operator);

        // read staker's deposited shares and strategies to add to operator's shares
        // and also update the staker depositScalingFactor for each strategy
        (IStrategy[] memory strategies, uint256[] memory depositedShares) = getDepositedShares(staker);
        uint256[] memory slashingFactors = _getSlashingFactors(staker, operator, strategies);

        for (uint256 i = 0; i < strategies.length; ++i) {
            // forgefmt: disable-next-item
            _increaseDelegation({
                operator: operator, 
                staker: staker, 
                strategy: strategies[i],
                existingDepositShares: uint256(0),
                addedShares: depositedShares[i],
                slashingFactor: slashingFactors[i]
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

        uint256[] memory prevSlashingFactors;
        {
            uint32 completableBlock = withdrawal.startBlock + MIN_WITHDRAWAL_DELAY_BLOCKS;
            require(completableBlock <= uint32(block.number), WithdrawalDelayNotElapsed());

            // Given the max magnitudes of the operator the staker was originally delegated to, calculate
            // the slashing factors for each of the withdrawal's strategies.
            prevSlashingFactors = _getSlashingFactorsAtBlock({
                staker: withdrawal.staker,
                operator: withdrawal.delegatedTo,
                strategies: withdrawal.strategies,
                blockNumber: completableBlock
            });
        }

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
                (uint256 existingDepositShares, uint256 addedShares) = shareManager.addShares({
                    staker: withdrawal.staker,
                    strategy: withdrawal.strategies[i],
                    token: tokens[i],
                    shares: sharesToWithdraw
                });

                // Update the staker's deposit scaling factor and delegate shares to their operator
                _increaseDelegation({
                    operator: newOperator,
                    staker: withdrawal.staker,
                    strategy: withdrawal.strategies[i],
                    existingDepositShares: existingDepositShares,
                    addedShares: addedShares,
                    slashingFactor: newSlashingFactors[i]
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
     * @param slashingFactor The current slashing factor for the staker/operator/strategy
     */
    function _increaseDelegation(
        address operator,
        address staker,
        IStrategy strategy,
        uint256 existingDepositShares,
        uint256 addedShares,
        uint256 slashingFactor
    ) internal {
        // Ensure that the operator has not been fully slashed for a strategy
        // and that the staker has not been fully slashed if it is the beaconChainStrategy
        require(slashingFactor != 0, FullySlashed());

        // Update the staker's depositScalingFactor. This only results in an update
        // if the slashing factor has changed for this strategy.
        DepositScalingFactor storage dsf = _depositScalingFactor[staker][strategy];
        dsf.update(existingDepositShares, addedShares, slashingFactor);
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
        uint256[] memory sharesToWithdraw = new uint256[](strategies.length);

        // Remove shares from staker and operator
        // Each of these operations fail if we attempt to remove more shares than exist
        for (uint256 i = 0; i < strategies.length; ++i) {
            IShareManager shareManager = _getShareManager(strategies[i]);
            DepositScalingFactor memory dsf = _depositScalingFactor[staker][strategies[i]];

            // Check withdrawing deposit shares amount doesn't exceed balance
            require(
                depositSharesToWithdraw[i] <= shareManager.stakerDepositShares(staker, strategies[i]),
                WithdrawalExceedsMax()
            );

            // Calculate how many shares can be withdrawn after factoring in slashing
            sharesToWithdraw[i] = dsf.calcWithdrawable(depositSharesToWithdraw[i], slashingFactors[i]);

            // Apply slashing. If the staker or operator has been fully slashed, this will return 0
            scaledShares[i] = SlashingLib.scaleForQueueWithdrawal({
                sharesToWithdraw: sharesToWithdraw[i],
                slashingFactor: slashingFactors[i]
            });

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
                    sharesToDecrease: sharesToWithdraw[i]
                });
            }

            // Remove deposit shares from EigenPodManager/StrategyManager
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
        queuedWithdrawals[withdrawalRoot] = withdrawal;
        _stakerQueuedWithdrawalRoots[staker].add(withdrawalRoot);

        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, sharesToWithdraw);
        return withdrawalRoot;
    }

    /// @dev Calculate the amount of slashing to apply to the staker's shares
    function _getSlashingFactor(
        address staker,
        IStrategy strategy,
        uint64 operatorMaxMagnitude
    ) internal view returns (uint256) {
        if (strategy == beaconChainETHStrategy) {
            uint64 beaconChainSlashingFactor = getBeaconChainSlashingFactor(staker);
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

    function _decreaseBeaconChainSlashingFactor(address staker, uint64 proportionOfOldBalance) internal {
        BeaconChainSlashingFactor memory bsf = _beaconChainSlashingFactor[staker];
        bsf.slashingFactor = uint64(uint256(getBeaconChainSlashingFactor(staker)).mulWad(proportionOfOldBalance));
        bsf.isSet = true;

        emit BeaconChainScalingFactorDecreased(staker, bsf.slashingFactor);
        _beaconChainSlashingFactor[staker] = bsf;
    }

    /**
     * @dev Calculate amount of slashable shares that would be slashed from the queued withdrawals from an operator for a strategy
     * given the previous maxMagnitude and the new maxMagnitude.
     * Note: To get the total amount of slashable shares in the queue withdrawable, set newMaxMagnitude to 0 and prevMaxMagnitude
     * is the current maxMagnitude of the operator.
     */
    function _getSlashedSharesInQueue(
        address operator,
        IStrategy strategy,
        uint64 prevMaxMagnitude,
        uint64 newMaxMagnitude
    ) internal view returns (uint256) {
        // Fetch the cumulative scaled shares sitting in the withdrawal queue both now and before
        // the withdrawal delay.
        uint256 curCumulativeScaledShares = _cumulativeScaledSharesHistory[operator][strategy].latest();
        uint256 prevCumulativeScaledShares = _cumulativeScaledSharesHistory[operator][strategy].upperLookup({
            key: uint32(block.number) - MIN_WITHDRAWAL_DELAY_BLOCKS
        });

        // The difference between these values represents the number of scaled shares that entered the
        // withdrawal queue less than `MIN_WITHDRAWAL_DELAY_BLOCKS` ago. These shares are still slashable,
        // so we use them to calculate the number of slashable shares in the withdrawal queue.
        uint256 slashableScaledShares = curCumulativeScaledShares - prevCumulativeScaledShares;

        return SlashingLib.scaleForBurning({
            scaledShares: slashableScaledShares,
            prevMaxMagnitude: prevMaxMagnitude,
            newMaxMagnitude: newMaxMagnitude
        });
    }

    /// @dev Add to the cumulative withdrawn scaled shares from an operator for a given strategy
    function _addQueuedSlashableShares(address operator, IStrategy strategy, uint256 scaledShares) internal {
        if (strategy != beaconChainETHStrategy) {
            uint256 currCumulativeScaledShares = _cumulativeScaledSharesHistory[operator][strategy].latest();
            _cumulativeScaledSharesHistory[operator][strategy].push({
                key: uint32(block.number),
                value: currCumulativeScaledShares + scaledShares
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
    function depositScalingFactor(address staker, IStrategy strategy) public view returns (uint256) {
        return _depositScalingFactor[staker][strategy].scalingFactor();
    }

    /// @inheritdoc IDelegationManager
    function getBeaconChainSlashingFactor(
        address staker
    ) public view returns (uint64) {
        BeaconChainSlashingFactor memory bsf = _beaconChainSlashingFactor[staker];
        return bsf.isSet ? bsf.slashingFactor : WAD;
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
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategy;
        uint64 maxMagnitude = allocationManager.getMaxMagnitudes(operator, strategies)[0];
        // Return amount of shares slashed if all remaining magnitude were to be slashed
        return _getSlashedSharesInQueue({
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
            shares[i] = new uint256[](withdrawals[i].strategies.length);

            uint256[] memory slashingFactors = _getSlashingFactors(staker, operator, withdrawals[i].strategies);

            for (uint256 j; j < withdrawals[i].strategies.length; ++j) {
                shares[i][j] = SlashingLib.scaleForCompleteWithdrawal({
                    scaledShares: withdrawals[i].scaledShares[j],
                    slashingFactor: slashingFactors[i]
                });
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
