// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../mixins/SignatureUtils.sol";
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
    SignatureUtils
{
    using SlashingLib for *;

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
        uint32 _MIN_WITHDRAWAL_DELAY
    )
        DelegationManagerStorage(
            _avsDirectory,
            _strategyManager,
            _eigenPodManager,
            _allocationManager,
            _MIN_WITHDRAWAL_DELAY
        )
    {
        _disableInitializers();
    }

    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, and paused status.
     * minWithdrawalDelayBlocks is set only once here
     */
    function initialize(
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 initialPausedStatus
    ) external initializer {
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _transferOwnership(initialOwner);
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

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
    ) external {
        require(!isDelegated(msg.sender), ActivelyDelegated());
        allocationManager.setAllocationDelay(msg.sender, allocationDelay);
        _setOperatorDetails(msg.sender, registeringOperatorDetails);
        SignatureWithExpiry memory emptySignatureAndExpiry;
        // delegate from the operator to themselves
        _delegate(msg.sender, msg.sender, emptySignatureAndExpiry, bytes32(0));
        // emit events
        emit OperatorRegistered(msg.sender, registeringOperatorDetails);
        emit OperatorMetadataURIUpdated(msg.sender, metadataURI);
    }

    /**
     * @notice Updates an operator's stored `OperatorDetails`.
     * @param newOperatorDetails is the updated `OperatorDetails` for the operator, to replace their current OperatorDetails`.
     *
     * @dev The caller must have previously registered as an operator in EigenLayer.
     */
    function modifyOperatorDetails(
        OperatorDetails calldata newOperatorDetails
    ) external {
        require(isOperator(msg.sender), OperatorNotRegistered());
        _setOperatorDetails(msg.sender, newOperatorDetails);
    }

    /**
     * @notice Called by an operator to emit an `OperatorMetadataURIUpdated` event indicating the information has updated.
     * @param metadataURI The URI for metadata associated with an operator
     */
    function updateOperatorMetadataURI(
        string calldata metadataURI
    ) external {
        require(isOperator(msg.sender), OperatorNotRegistered());
        emit OperatorMetadataURIUpdated(msg.sender, metadataURI);
    }

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
     */
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

    /**
     * @notice Caller delegates a staker's stake to an operator with valid signatures from both parties.
     * @param staker The account delegating stake to an `operator` account
     * @param operator The account (`staker`) is delegating its assets to for use in serving applications built on EigenLayer.
     * @param stakerSignatureAndExpiry Signed data from the staker authorizing delegating stake to an operator
     * @param approverSignatureAndExpiry is a parameter that will be used for verifying that the operator approves of this delegation action in the event that:
     * @param approverSalt Is a salt used to help guarantee signature uniqueness. Each salt can only be used once by a given approver.
     *
     * @dev If `staker` is an EOA, then `stakerSignature` is verified to be a valid ECDSA stakerSignature from `staker`, indicating their intention for this action.
     * @dev If `staker` is a contract, then `stakerSignature` will be checked according to EIP-1271.
     * @dev the operator's `delegationApprover` address is set to a non-zero value.
     * @dev neither the operator nor their `delegationApprover` is the `msg.sender`, since in the event that the operator or their delegationApprover
     * is the `msg.sender`, then approval is assumed.
     * @dev This function will revert if the current `block.timestamp` is equal to or exceeds the expiry
     * @dev In the case that `approverSignatureAndExpiry` is not checked, its content is ignored entirely; it's recommended to use an empty input
     * in this case to save on complexity + gas costs
     */
    function delegateToBySignature(
        address staker,
        address operator,
        SignatureWithExpiry memory stakerSignatureAndExpiry,
        SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt
    ) external {
        // check the signature expiry
        require(stakerSignatureAndExpiry.expiry >= block.timestamp, SignatureExpired());
        require(!isDelegated(staker), ActivelyDelegated());
        require(isOperator(operator), OperatorNotRegistered());

        // calculate the digest hash, then increment `staker`'s nonce
        uint256 currentStakerNonce = stakerNonce[staker];

        // actually check that the signature is valid
        _checkIsValidSignatureNow({
            signer: staker, 
            signableDigest: calculateStakerDelegationDigestHash({
                staker: staker, 
                nonce: currentStakerNonce, 
                operator: operator, 
                expiry: stakerSignatureAndExpiry.expiry
            }), 
            signature: stakerSignatureAndExpiry.signature
        });

        unchecked {
            stakerNonce[staker] = currentStakerNonce + 1;
        }

        // go through the internal delegation flow, checking the `approverSignatureAndExpiry` if applicable
        _delegate(staker, operator, approverSignatureAndExpiry, approverSalt);
    }

    /**
     * Allows the staker, the staker's operator, or that operator's delegationApprover to undelegate
     * a staker from their operator. Undelegation immediately removes ALL active shares/strategies from
     * both the staker and operator, and places the shares and strategies in the withdrawal queue
     */
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

        // emit an event if this action was not initiated by the staker themselves
        if (msg.sender != staker) {
            emit StakerForceUndelegated(staker, operator);
        }

        // undelegate the staker
        emit StakerUndelegated(staker, operator);
        delegatedTo[staker] = address(0);

        // if no deposited shares, return an empty array, and don't queue a withdrawal
        if (strategies.length == 0) {
            withdrawalRoots = new bytes32[](0);
        } else {
            withdrawalRoots = new bytes32[](strategies.length);
            uint64[] memory totalMagnitudes = allocationManager.getTotalMagnitudes(operator, strategies);

            for (uint256 i = 0; i < strategies.length; i++) {
                IStrategy[] memory singleStrategy = new IStrategy[](1);
                uint256[] memory singleShares = new uint256[](1);
                uint64[] memory singleTotalMagnitude = new uint64[](1);
                singleStrategy[0] = strategies[i];
                // TODO: this part is a bit gross, can we make it better?
                singleShares[0] =
                    depositedShares[i].toShares(stakerScalingFactor[staker][strategies[i]], totalMagnitudes[i]);
                singleTotalMagnitude[0] = totalMagnitudes[i];

                withdrawalRoots[i] = _removeSharesAndQueueWithdrawal({
                    staker: staker,
                    operator: operator,
                    strategies: singleStrategy,
                    sharesToWithdraw: singleShares,
                    totalMagnitudes: singleTotalMagnitude
                });

                // all shares and queued withdrawn and no delegated operator
                // reset staker's depositScalingFactor to default
                stakerScalingFactor[staker][strategies[i]].depositScalingFactor = WAD;
                emit DepositScalingFactorUpdated(staker, strategies[i], WAD);
            }
        }

        return withdrawalRoots;
    }

    /**
     * Allows a staker to withdraw some shares. Withdrawn shares/strategies are immediately removed
     * from the staker. If the staker is delegated, withdrawn shares/strategies are also removed from
     * their operator.
     *
     * All withdrawn shares/strategies are placed in a queue and can be fully withdrawn after a delay.
     */
    function queueWithdrawals(
        QueuedWithdrawalParams[] calldata queuedWithdrawalParams
    ) external onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE) returns (bytes32[] memory) {
        bytes32[] memory withdrawalRoots = new bytes32[](queuedWithdrawalParams.length);
        address operator = delegatedTo[msg.sender];

        for (uint256 i = 0; i < queuedWithdrawalParams.length; i++) {
            require(
                queuedWithdrawalParams[i].strategies.length == queuedWithdrawalParams[i].ownedShares.length,
                InputArrayLengthMismatch()
            );
            require(queuedWithdrawalParams[i].withdrawer == msg.sender, WithdrawerNotStaker());

            uint64[] memory totalMagnitudes =
                allocationManager.getTotalMagnitudes(operator, queuedWithdrawalParams[i].strategies);

            // Remove shares from staker's strategies and place strategies/shares in queue.
            // If the staker is delegated to an operator, the operator's delegated shares are also reduced
            // NOTE: This will fail if the staker doesn't have the shares implied by the input parameters
            withdrawalRoots[i] = _removeSharesAndQueueWithdrawal({
                staker: msg.sender,
                operator: operator,
                strategies: queuedWithdrawalParams[i].strategies,
                sharesToWithdraw: queuedWithdrawalParams[i].ownedShares,
                totalMagnitudes: totalMagnitudes
            });
        }
        return withdrawalRoots;
    }

    /**
     * @notice Used to complete the specified `withdrawal`. The caller must match `withdrawal.withdrawer`
     * @param withdrawal The Withdrawal to complete.
     * @param tokens Array in which the i-th entry specifies the `token` input to the 'withdraw' function of the i-th Strategy in the `withdrawal.strategies` array.
     * @param receiveAsTokens If true, the shares specified in the withdrawal will be withdrawn from the specified strategies themselves
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
    ) external onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE) nonReentrant {
        _completeQueuedWithdrawal(withdrawal, tokens, receiveAsTokens);
    }

    /**
     * @notice Array-ified version of `completeQueuedWithdrawal`.
     * Used to complete the specified `withdrawals`. The function caller must match `withdrawals[...].withdrawer`
     * @param withdrawals The Withdrawals to complete.
     * @param tokens Array of tokens for each Withdrawal. See `completeQueuedWithdrawal` for the usage of a single array.
     * @param receiveAsTokens Whether or not to complete each withdrawal as tokens. See `completeQueuedWithdrawal` for the usage of a single boolean.
     * @dev See `completeQueuedWithdrawal` for relevant dev tags
     */
    function completeQueuedWithdrawals(
        Withdrawal[] calldata withdrawals,
        IERC20[][] calldata tokens,
        bool[] calldata receiveAsTokens
    ) external onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE) nonReentrant {
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            _completeQueuedWithdrawal(withdrawals[i], tokens[i], receiveAsTokens[i]);
        }
    }

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
     * @dev Callable only by the StrategyManager or EigenPodManager.
     */
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
            uint64[] memory totalMagnitudes = allocationManager.getTotalMagnitudes(operator, strategies);

            // add deposit shares to operator's stake shares and update the staker's depositScalingFactor
            _increaseDelegation({
                operator: operator,
                staker: staker,
                strategy: strategy,
                existingDepositShares: existingDepositShares,
                addedShares: addedShares,
                totalMagnitude: totalMagnitudes[0]
            });
        }
    }

    /**
     * @notice Decreases a native restaker's delegated share balance in a strategy due to beacon chain slashing. This updates their beaconChainScalingFactor.
     * Their operator's stakeShares are also updated (if they are delegated).
     * @param staker The address to increase the delegated stakeShares for their operator.
     * @param existingDepositShares The number of shares the staker already has in the EPM. This does not change upon decreasing shares.
     * @param proportionOfOldBalance The current pod owner shares proportion of the previous pod owner shares
     *
     * @dev *If the staker is actively delegated*, then decreases the `staker`'s delegated stakeShares in `strategy` by `proportionPodBalanceDecrease` proportion. Otherwise does nothing.
     * @dev Callable only by the EigenPodManager.
     */
    function decreaseBeaconChainScalingFactor(
        address staker,
        uint256 existingDepositShares,
        uint64 proportionOfOldBalance
    ) external onlyEigenPodManager {
        // decrease the staker's beaconChainScalingFactor proportionally
        address operator = delegatedTo[staker];
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = beaconChainETHStrategy;
        uint64[] memory totalMagnitudes = allocationManager.getTotalMagnitudes(operator, strategies);

        uint256 sharesBefore =
            existingDepositShares.toShares(stakerScalingFactor[staker][beaconChainETHStrategy], totalMagnitudes[0]);
        StakerScalingFactors storage ssf = stakerScalingFactor[staker][beaconChainETHStrategy];
        ssf.decreaseBeaconChainScalingFactor(proportionOfOldBalance);
        emit BeaconChainScalingFactorDecreased(staker, ssf.beaconChainScalingFactor);
        uint256 sharesAfter =
            existingDepositShares.toShares(stakerScalingFactor[staker][beaconChainETHStrategy], totalMagnitudes[0]);

        // if the staker is delegated to an operators
        if (isDelegated(staker)) {
            // subtract strategy shares from delegated scaled shares
            _decreaseDelegation({
                operator: operator,
                staker: staker,
                strategy: beaconChainETHStrategy,
                // TODO: fix this
                operatorSharesToDecrease: sharesBefore - sharesAfter
            });
        }
    }

    /**
     * @notice Decreases the operators shares in storage after a slash
     * @param operator The operator to decrease shares for
     * @param strategy The strategy to decrease shares for
     * @param previousTotalMagnitude The total magnitude before the slash
     * @param newTotalMagnitude The total magnitude after the slash
     * @dev Callable only by the AllocationManager
     */
    function decreaseOperatorShares(
        address operator,
        IStrategy strategy,
        uint64 previousTotalMagnitude,
        uint64 newTotalMagnitude
    ) external onlyAllocationManager {
        uint256 operatorSharesToDecrease =
            operatorShares[operator][strategy].getOperatorSharesToDecrease(previousTotalMagnitude, newTotalMagnitude);

        _decreaseDelegation({
            operator: operator,
            staker: address(0), // we treat this as a decrease for the zero address staker
            strategy: strategy,
            operatorSharesToDecrease: operatorSharesToDecrease
        });
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
            // check the signature expiry
            require(approverSignatureAndExpiry.expiry >= block.timestamp, SignatureExpired());
            // check that the salt hasn't been used previously, then mark the salt as spent
            require(!delegationApproverSaltIsSpent[approver][approverSalt], SaltSpent());
            // actually check that the signature is valid
            _checkIsValidSignatureNow({
                signer: approver, 
                signableDigest:  calculateDelegationApprovalDigestHash(
                    staker, 
                    operator, 
                    approver, 
                    approverSalt, 
                    approverSignatureAndExpiry.expiry
                ), 
                signature: approverSignatureAndExpiry.signature
            });

            delegationApproverSaltIsSpent[approver][approverSalt] = true;
        }

        // record the delegation relation between the staker and operator, and emit an event
        delegatedTo[staker] = operator;
        emit StakerDelegated(staker, operator);

        // read staker's deposited shares and strategies to add to operator's shares
        // and also update the staker depositScalingFactor for each strategy
        (IStrategy[] memory strategies, uint256[] memory depositedShares) = getDepositedShares(staker);
        uint64[] memory totalMagnitudes = allocationManager.getTotalMagnitudes(operator, strategies);

        for (uint256 i = 0; i < strategies.length; ++i) {
            // forgefmt: disable-next-item
            _increaseDelegation({
                operator: operator, 
                staker: staker, 
                strategy: strategies[i],
                existingDepositShares: uint256(0),
                addedShares: depositedShares[i],
                totalMagnitude: totalMagnitudes[i]
            });
        }
    }

    /**
     * @dev This function completes a queued withdrawal for a staker.
     * This will apply any slashing that has occurred since the the withdrawal was queued. By multiplying the withdrawl's
     * delegatedShares by the operator's total magnitude for each strategy
     * If receiveAsTokens is true, then these shares will be withdrawn as tokens.
     * If receiveAsTokens is false, then they will be redeposited according to the current operator the staker is delegated to,
     * and added back to the operator's delegatedShares.
     */
    function _completeQueuedWithdrawal(
        Withdrawal calldata withdrawal,
        IERC20[] calldata tokens,
        bool receiveAsTokens
    ) internal {
        require(tokens.length == withdrawal.strategies.length, InputArrayLengthMismatch());
        require(msg.sender == withdrawal.withdrawer, WithdrawerNotCaller());
        bytes32 withdrawalRoot = calculateWithdrawalRoot(withdrawal);
        require(pendingWithdrawals[withdrawalRoot], WithdrawalNotQueued());

        // TODO: is there a cleaner way to do this?
        uint32 completableTimestamp = _checkCompletability(withdrawal.startTimestamp);
        // read delegated operator's totalMagnitudes at time of withdrawal to convert the delegatedShares to shared
        // factoring in slashing that occured during withdrawal delay
        uint64[] memory totalMagnitudes = allocationManager.getTotalMagnitudesAtTimestamp({
            operator: withdrawal.delegatedTo,
            strategies: withdrawal.strategies,
            timestamp: completableTimestamp
        });

        for (uint256 i = 0; i < withdrawal.strategies.length; i++) {
            IShareManager shareManager = _getShareManager(withdrawal.strategies[i]);
            uint256 sharesToWithdraw = withdrawal.scaledSharesToWithdraw[i].scaleSharesForCompleteWithdrawal(
                stakerScalingFactor[withdrawal.staker][withdrawal.strategies[i]], totalMagnitudes[i]
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

        // Remove `withdrawalRoot` from pending roots
        delete pendingWithdrawals[withdrawalRoot];
        emit WithdrawalCompleted(withdrawalRoot);
    }

    /**
     * @notice Increases `operator`s depositedShares in `strategy` based on staker's addedDepositShares
     * and increases the staker's depositScalingFactor for the strategy.
     * @param operator The operator to increase the delegated delegatedShares for
     * @param staker The staker to increase the depositScalingFactor for
     * @param strategy The strategy to increase the delegated delegatedShares and the depositScalingFactor for
     * @param existingDepositShares The number of deposit shares the staker already has in the strategy.
     * @param addedShares The shares added to the staker in the StrategyManager/EigenPodManager
     * @param totalMagnitude The current total magnitude of the operator for the strategy
     */
    function _increaseDelegation(
        address operator,
        address staker,
        IStrategy strategy,
        uint256 existingDepositShares,
        uint256 addedShares,
        uint64 totalMagnitude
    ) internal {
        //TODO: double check ordering here
        operatorShares[operator][strategy] += addedShares;
        emit OperatorSharesIncreased(operator, staker, strategy, addedShares);

        // update the staker's depositScalingFactor
        StakerScalingFactors storage ssf = stakerScalingFactor[staker][strategy];
        ssf.updateDepositScalingFactor(existingDepositShares, addedShares, totalMagnitude);
        emit DepositScalingFactorUpdated(staker, strategy, ssf.depositScalingFactor);
    }

    /**
     * @notice Decreases `operator`s shares in `strategy` based on staker's removed shares and operator's totalMagnitude
     * @param operator The operator to decrease the delegated delegated shares for
     * @param staker The staker to decrease the delegated delegated shares for
     * @param strategy The strategy to decrease the delegated delegated shares for
     * @param operatorSharesToDecrease The delegatedShares to remove from the operator's delegated shares
     */
    function _decreaseDelegation(
        address operator,
        address staker,
        IStrategy strategy,
        uint256 operatorSharesToDecrease
    ) internal {
        // Decrement operator shares
        operatorShares[operator][strategy] -= operatorSharesToDecrease;

        emit OperatorSharesDecreased(operator, staker, strategy, operatorSharesToDecrease);
    }

    /**
     * @notice Removes `sharesToWithdraw` in `strategies` from `staker` who is currently delegated to `operator` and queues a withdrawal to the `withdrawer`.
     * @dev If the `operator` is indeed an operator, then the operator's delegated shares in the `strategies` are also decreased appropriately.
     * @dev If `withdrawer` is not the same address as `staker`
     */
    function _removeSharesAndQueueWithdrawal(
        address staker,
        address operator,
        IStrategy[] memory strategies,
        uint256[] memory sharesToWithdraw,
        uint64[] memory totalMagnitudes
    ) internal returns (bytes32) {
        require(staker != address(0), InputAddressZero());
        require(strategies.length != 0, InputArrayLengthZero());

        uint256[] memory scaledSharesToWithdraw = new uint256[](strategies.length);
        // Remove shares from staker and operator
        // Each of these operations fail if we attempt to remove more shares than exist
        for (uint256 i = 0; i < strategies.length; ++i) {
            IShareManager shareManager = _getShareManager(strategies[i]);

            uint256 depositSharesToRemove =
                sharesToWithdraw[i].toDepositShares(stakerScalingFactor[staker][strategies[i]], totalMagnitudes[i]);
            // TODO: maybe have a getter to get shares for all strategies,
            // check sharesToWithdraw is valid
            // but for inputted strategies
            uint256 depositSharesWithdrawable = shareManager.stakerDepositShares(staker, strategies[i]);
            require(depositSharesToRemove <= depositSharesWithdrawable, WithdrawalExeedsMax());

            // Similar to `isDelegated` logic
            if (operator != address(0)) {
                // forgefmt: disable-next-item
                _decreaseDelegation({
                    operator: operator, 
                    staker: staker, 
                    strategy: strategies[i], 
                    operatorSharesToDecrease: sharesToWithdraw[i]
                });
            }
            scaledSharesToWithdraw[i] = sharesToWithdraw[i].scaleSharesForQueuedWithdrawal(
                stakerScalingFactor[staker][strategies[i]], totalMagnitudes[i]
            );

            // Remove active shares from EigenPodManager/StrategyManager
            // EigenPodManager: this call will revert if it would reduce the Staker's virtual beacon chain ETH shares below zero
            // StrategyManager: this call will revert if `sharesToDecrement` exceeds the Staker's current deposit shares in `strategies[i]`
            shareManager.removeDepositShares(staker, strategies[i], depositSharesToRemove);
        }

        // Create queue entry and increment withdrawal nonce
        uint256 nonce = cumulativeWithdrawalsQueued[staker];
        cumulativeWithdrawalsQueued[staker]++;

        Withdrawal memory withdrawal = Withdrawal({
            staker: staker,
            delegatedTo: operator,
            withdrawer: staker,
            nonce: nonce,
            startTimestamp: uint32(block.timestamp),
            strategies: strategies,
            scaledSharesToWithdraw: scaledSharesToWithdraw
        });

        bytes32 withdrawalRoot = calculateWithdrawalRoot(withdrawal);

        // Place withdrawal in queue
        pendingWithdrawals[withdrawalRoot] = true;

        emit WithdrawalQueued(withdrawalRoot, withdrawal);
        return withdrawalRoot;
    }

    /// @dev check whether the withdrawal delay has elapsed (handles both legacy and post-slashing-release withdrawals) and returns the completable timestamp
    function _checkCompletability(
        uint32 startTimestamp
    ) internal view returns (uint32 completableTimestamp) {
        if (startTimestamp < LEGACY_WITHDRAWALS_TIMESTAMP) {
            // this is a legacy M2 withdrawal using blocknumbers.
            // It would take up to 600+ years for the blocknumber to reach the LEGACY_WITHDRAWALS_TIMESTAMP, so this is a safe check.
            require(startTimestamp + LEGACY_MIN_WITHDRAWAL_DELAY_BLOCKS <= block.number, WithdrawalDelayNotElapsed());
            // sourcing the magnitudes from time=0, will always give us WAD, which doesn't factor in slashing
            completableTimestamp = 0;
        } else {
            // this is a post Slashing release withdrawal using timestamps
            require(startTimestamp + MIN_WITHDRAWAL_DELAY <= block.timestamp, WithdrawalDelayNotElapsed());
            // source magnitudes from the time of completability
            completableTimestamp = startTimestamp + MIN_WITHDRAWAL_DELAY;
        }
    }

    /**
     *
     *                              SHARES CONVERSION FUNCTIONS
     *
     */
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

    /**
     * @notice Returns 'true' if `staker` *is* actively delegated, and 'false' otherwise.
     */
    function isDelegated(
        address staker
    ) public view returns (bool) {
        return (delegatedTo[staker] != address(0));
    }

    /**
     * @notice Returns true is an operator has previously registered for delegation.
     */
    function isOperator(
        address operator
    ) public view returns (bool) {
        return operator != address(0) && delegatedTo[operator] == operator;
    }

    /**
     * @notice Returns the OperatorDetails struct associated with an `operator`.
     */
    function operatorDetails(
        address operator
    ) external view returns (OperatorDetails memory) {
        return _operatorDetails[operator];
    }

    /**
     * @notice Returns the delegationApprover account for an operator
     */
    function delegationApprover(
        address operator
    ) external view returns (address) {
        return _operatorDetails[operator].delegationApprover;
    }

    /**
     * @notice Returns the stakerOptOutWindowBlocks for an operator
     */
    function stakerOptOutWindowBlocks(
        address operator
    ) external view returns (uint256) {
        return _operatorDetails[operator].stakerOptOutWindowBlocks;
    }

    /**
     * @notice Given a staker and a set of strategies, return the shares they can queue for withdrawal.
     * This value depends on which operator the staker is delegated to.
     * The shares amount returned is the actual amount of Strategy shares the staker would receive (subject
     * to each strategy's underlying shares to token ratio).
     */
    function getWithdrawableShares(
        address staker,
        IStrategy[] memory strategies
    ) public view returns (uint256[] memory withdrawableShares) {
        address operator = delegatedTo[staker];
        uint64[] memory totalMagnitudes = allocationManager.getTotalMagnitudes(operator, strategies);

        for (uint256 i = 0; i < strategies.length; ++i) {
            IShareManager shareManager = _getShareManager(strategies[i]);
            // TODO: batch call for strategyManager shares?
            // 1. read strategy deposit shares

            // forgefmt: disable-next-item
            uint256 depositShares = shareManager.stakerDepositShares(staker, strategies[i]);

            // 2. if the staker is delegated, actual withdrawable shares can be different from what is stored
            // in the StrategyManager/EigenPodManager because they could have been slashed
            if (operator != address(0)) {
                // forgefmt: disable-next-item
                withdrawableShares[i] = depositShares.toShares(
                    stakerScalingFactor[staker][strategies[i]],
                    totalMagnitudes[i]
                );
            } else {
                withdrawableShares[i] = depositShares;
            }
        }
        return withdrawableShares;
    }

    /**
     * @notice Returns the number of shares in storage for a staker and all their strategies
     * TODO: make cleaner, get rid of assembly
     */
    function getDepositedShares(
        address staker
    ) public view returns (IStrategy[] memory, uint256[] memory) {
        // Get a list of all the strategies to check
        IStrategy[] memory strategies = strategyManager.getStakerStrategyList(staker);

        // resize and add beaconChainETH to the end
        assembly {
            mstore(strategies, add(mload(strategies), 1))
        }
        strategies[strategies.length - 1] = beaconChainETHStrategy;
        uint256[] memory shares = new uint256[](strategies.length);

        // Get owned shares for each strategy
        for (uint256 i = 0; i < strategies.length; ++i) {
            IShareManager shareManager = _getShareManager(strategies[i]);
            shares[i] = shareManager.stakerDepositShares(staker, strategies[i]);
        }

        // if the last shares are 0, remove them
        if (shares[strategies.length - 1] == 0) {
            // resize the arrays
            assembly {
                mstore(strategies, sub(mload(strategies), 1))
                mstore(shares, sub(mload(shares), 1))
            }
        }

        return (strategies, shares);
    }

    /// @notice Returns the keccak256 hash of `withdrawal`.
    function calculateWithdrawalRoot(
        Withdrawal memory withdrawal
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(withdrawal));
    }

    /**
     * @notice Calculates the digestHash for a `staker` to sign to delegate to an `operator`
     * @param staker The signing staker
     * @param operator The operator who is being delegated to
     * @param expiry The desired expiry time of the staker's signature
     */
    function calculateCurrentStakerDelegationDigestHash(
        address staker,
        address operator,
        uint256 expiry
    ) external view returns (bytes32) {
        return calculateStakerDelegationDigestHash(staker, stakerNonce[staker], operator, expiry);
    }

    /**
     * @notice Calculates the digest hash to be signed and used in the `delegateToBySignature` function
     * @param staker The signing staker
     * @param nonce The nonce of the staker. In practice we use the staker's current nonce, stored at `stakerNonce[staker]`
     * @param operator The operator who is being delegated to
     * @param expiry The desired expiry time of the staker's signature
     */
    function calculateStakerDelegationDigestHash(
        address staker,
        uint256 nonce,
        address operator,
        uint256 expiry
    ) public view returns (bytes32) {
        /// forgefmt: disable-next-item
        return _calculateSignableDigest(
            keccak256(
                abi.encode(
                    STAKER_DELEGATION_TYPEHASH, 
                    staker, 
                    operator, 
                    nonce, 
                    expiry
                )
            )
        );
    }

    /**
     * @notice Calculates the digest hash to be signed by the operator's delegationApprove and used in the `delegateTo` and `delegateToBySignature` functions.
     * @param staker The account delegating their stake
     * @param operator The account receiving delegated stake
     * @param approver the operator's `delegationApprover` who will be signing the delegationHash (in general)
     * @param approverSalt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid
     */
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
