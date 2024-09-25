// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "../permissions/Pausable.sol";
import "../libraries/EIP1271SignatureUtils.sol";
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
    ReentrancyGuardUpgradeable
{
    using SlashingLib for *;

    // @dev Index for flag that pauses new delegations when set
    uint8 internal constant PAUSED_NEW_DELEGATION = 0;

    // @dev Index for flag that pauses queuing new withdrawals when set.
    uint8 internal constant PAUSED_ENTER_WITHDRAWAL_QUEUE = 1;

    // @dev Index for flag that pauses completing existing withdrawals when set.
    uint8 internal constant PAUSED_EXIT_WITHDRAWAL_QUEUE = 2;

    // @dev Chain ID at the time of contract deployment
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    /// @dev The minimum number of blocks to complete a withdrawal of a strategy. 50400 * 12 seconds = 1 week
    uint256 public constant LEGACY_MIN_WITHDRAWAL_DELAY_BLOCKS = 50_400;

    /// @dev Wed Jan 01 2025 17:00:00 GMT+0000, timestamp used to check whether a pending withdrawal
    /// should be processed as legacy M2 or with slashing considered.
    uint32 public constant LEGACY_WITHDRAWALS_TIMESTAMP = 1_735_750_800;

    /// @notice Canonical, virtual beacon chain ETH strategy
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

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

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the immutable addresses of the strategy mananger and slasher.
     */
    constructor(
        IStrategyManager _strategyManager,
        ISlasher _slasher,
        IEigenPodManager _eigenPodManager,
        IAVSDirectory _avsDirectory,
        IAllocationManager _allocationManager,
        uint32 _MIN_WITHDRAWAL_DELAY
    )
        DelegationManagerStorage(
            _strategyManager,
            _slasher,
            _eigenPodManager,
            _avsDirectory,
            _allocationManager,
            _MIN_WITHDRAWAL_DELAY
        )
    {
        _disableInitializers();
        ORIGINAL_CHAIN_ID = block.chainid;
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
        _DOMAIN_SEPARATOR = _calculateDomainSeparator();
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
        bytes32 stakerDigestHash =
            calculateStakerDelegationDigestHash(staker, currentStakerNonce, operator, stakerSignatureAndExpiry.expiry);
        unchecked {
            stakerNonce[staker] = currentStakerNonce + 1;
        }

        // actually check that the signature is valid
        EIP1271SignatureUtils.checkSignature_EIP1271(staker, stakerDigestHash, stakerSignatureAndExpiry.signature);

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

        // Gather strategies and shares from the staker. Calculate delegatedShares to remove from operator during undelegation
        // Undelegation removes ALL currently-active strategies and shares
        (IStrategy[] memory strategies, OwnedShares[] memory ownedShares) = getDelegatableShares(staker);

        // emit an event if this action was not initiated by the staker themselves
        if (msg.sender != staker) {
            emit StakerForceUndelegated(staker, operator);
        }

        // undelegate the staker
        emit StakerUndelegated(staker, operator);
        delegatedTo[staker] = address(0);

        // if no delegatable shares, return an empty array, and don't queue a withdrawal
        if (strategies.length == 0) {
            withdrawalRoots = new bytes32[](0);
        } else {
            withdrawalRoots = new bytes32[](strategies.length);
            uint64[] memory totalMagnitudes = allocationManager.getTotalMagnitudes(operator, strategies);

            for (uint256 i = 0; i < strategies.length; i++) {
                IStrategy[] memory singleStrategy = new IStrategy[](1);
                OwnedShares[] memory singleOwnedShare = new OwnedShares[](1);
                uint64[] memory singleTotalMagnitude = new uint64[](1);
                singleStrategy[0] = strategies[i];
                singleOwnedShare[0] = ownedShares[i];
                singleTotalMagnitude[0] = totalMagnitudes[i];

                withdrawalRoots[i] = _removeSharesAndQueueWithdrawal({
                    staker: staker,
                    operator: operator,
                    strategies: singleStrategy,
                    ownedSharesToWithdraw: singleOwnedShare,
                    totalMagnitudes: singleTotalMagnitude
                });

                // all shares and queued withdrawn and no delegated operator
                // reset staker's depositScalingFactor to default
                stakerScalingFactors[staker][strategies[i]].depositScalingFactor = WAD;
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
                ownedSharesToWithdraw: queuedWithdrawalParams[i].ownedShares,
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
     * @param existingShares The number of deposit shares the staker already has in the strategy. This is the shares amount stored in the
     * StrategyManager/EigenPodManager for the staker's shares.
     * @param addedOwnedShares The number of shares to added to the staker's shares in the strategy
     *
     * @dev *If the staker is actively delegated*, then increases the `staker`'s delegated delegatedShares in `strategy`.
     * Otherwise does nothing.
     * @dev Callable only by the StrategyManager or EigenPodManager.
     */
    function increaseDelegatedShares(
        address staker,
        IStrategy strategy,
        Shares existingShares,
        OwnedShares addedOwnedShares
    ) external onlyStrategyManagerOrEigenPodManager {
        // if the staker is delegated to an operator
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];
            uint64 totalMagnitude = allocationManager.getTotalMagnitude(operator, strategy);

            // add deposit shares to operator's stake shares and update the staker's depositScalingFactor
            _increaseDelegation({
                operator: operator,
                staker: staker,
                strategy: strategy,
                existingShares: existingShares,
                addedOwnedShares: addedOwnedShares,
                totalMagnitude: totalMagnitude
            });
        }
    }

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
        Shares existingShares,
        uint64 proportionOfOldBalance
    ) external onlyEigenPodManager {
        DelegatedShares delegatedSharesBefore = existingShares.toDelegatedShares(stakerScalingFactors[staker][beaconChainETHStrategy]);

        // decrease the staker's beaconChainScalingFactor proportionally
        // forgefmt: disable-next-item
        stakerScalingFactors[staker][beaconChainETHStrategy].decreaseBeaconChainScalingFactor(proportionOfOldBalance);

        DelegatedShares delegatedSharesAfter = existingShares.toDelegatedShares(stakerScalingFactors[staker][beaconChainETHStrategy]);

        // if the staker is delegated to an operators
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];

            // subtract strategy shares from delegated scaled shares
            _decreaseDelegation({
                operator: operator,
                staker: staker,
                strategy: beaconChainETHStrategy,
                delegatedShares: delegatedSharesBefore.sub(delegatedSharesAfter)
            });
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
        address _delegationApprover = _operatorDetails[operator].delegationApprover;
        /**
         * Check the `_delegationApprover`'s signature, if applicable.
         * If the `_delegationApprover` is the zero address, then the operator allows all stakers to delegate to them and this verification is skipped.
         * If the `_delegationApprover` or the `operator` themselves is the caller, then approval is assumed and signature verification is skipped as well.
         */
        if (_delegationApprover != address(0) && msg.sender != _delegationApprover && msg.sender != operator) {
            // check the signature expiry
            require(approverSignatureAndExpiry.expiry >= block.timestamp, SignatureExpired());
            // check that the salt hasn't been used previously, then mark the salt as spent
            require(!delegationApproverSaltIsSpent[_delegationApprover][approverSalt], SaltSpent());
            delegationApproverSaltIsSpent[_delegationApprover][approverSalt] = true;

            // forgefmt: disable-next-item
            // calculate the digest hash
            bytes32 approverDigestHash = calculateDelegationApprovalDigestHash(
                staker, 
                operator, 
                _delegationApprover, 
                approverSalt, 
                approverSignatureAndExpiry.expiry
            );

            // forgefmt: disable-next-item
            // actually check that the signature is valid
            EIP1271SignatureUtils.checkSignature_EIP1271(
                _delegationApprover, 
                approverDigestHash, 
                approverSignatureAndExpiry.signature
            );
        }

        // record the delegation relation between the staker and operator, and emit an event
        delegatedTo[staker] = operator;
        emit StakerDelegated(staker, operator);

        // read staker's delegatable shares and strategies to add to operator's delegatedShares
        // and also update the staker depositScalingFactor for each strategy
        (IStrategy[] memory strategies, OwnedShares[] memory ownedShares) = getDelegatableShares(staker);
        uint64[] memory totalMagnitudes = allocationManager.getTotalMagnitudes(operator, strategies);

        for (uint256 i = 0; i < strategies.length; ++i) {
            // forgefmt: disable-next-item
            _increaseDelegation({
                operator: operator, 
                staker: staker, 
                strategy: strategies[i],
                existingShares: uint256(0).wrapShares(),
                addedOwnedShares: ownedShares[i],
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
            OwnedShares ownedSharesToWithdraw = 
                withdrawal.delegatedShares[i]
                .scaleForCompleteWithdrawal(stakerScalingFactors[withdrawal.staker][withdrawal.strategies[i]])
                .toOwnedShares(totalMagnitudes[i]);

            if (receiveAsTokens) {
                // Withdraws `shares` in `strategy` to `withdrawer`. If the shares are virtual beaconChainETH shares,
                // then a call is ultimately forwarded to the `staker`s EigenPod; otherwise a call is ultimately forwarded
                // to the `strategy` with info on the `token`.
                shareManager.withdrawSharesAsTokens({
                    staker: withdrawal.staker,
                    strategy: withdrawal.strategies[i],
                    token: tokens[i],
                    ownedShares: ownedSharesToWithdraw
                });
            } else {
                // Award shares back in StrategyManager/EigenPodManager.
                shareManager.addOwnedShares({
                    staker: withdrawal.staker,
                    strategy: withdrawal.strategies[i],
                    token: tokens[i],
                    ownedShares: ownedSharesToWithdraw
                });
            }
        }

        // Remove `withdrawalRoot` from pending roots
        delete pendingWithdrawals[withdrawalRoot];
        emit WithdrawalCompleted(withdrawalRoot);
    }

    /**
     * @notice Increases `operator`s delegated delegatedShares in `strategy` based on staker's added shares and operator's totalMagnitude
     * and increases the staker's depositScalingFactor for the strategy.
     * @param operator The operator to increase the delegated delegatedShares for
     * @param staker The staker to increase the depositScalingFactor for
     * @param strategy The strategy to increase the delegated delegatedShares and the depositScalingFactor for
     * @param existingShares The number of deposit shares the staker already has in the strategy.
     * @param addedOwnedShares The shares added to the staker in the StrategyManager/EigenPodManager
     * @param totalMagnitude The current total magnitude of the operator for the strategy
     */
    function _increaseDelegation(
        address operator,
        address staker,
        IStrategy strategy,
        Shares existingShares,
        OwnedShares addedOwnedShares,
        uint64 totalMagnitude
    ) internal {
        _updateDepositScalingFactor({
            staker: staker,
            strategy: strategy,
            existingShares: existingShares,
            addedOwnedShares: addedOwnedShares,
            totalMagnitude: totalMagnitude
        });

        // based on total magnitude, update operators delegatedShares
        DelegatedShares delegatedShares = addedOwnedShares.toDelegatedShares(totalMagnitude);
        // forgefmt: disable-next-line
        operatorDelegatedShares[operator][strategy] = operatorDelegatedShares[operator][strategy].add(delegatedShares);

        // TODO: What to do about event wrt scaling?
        emit OperatorSharesIncreased(operator, staker, strategy, delegatedShares);
    }

    /**
     * @notice Decreases `operator`s delegated delegatedShares in `strategy` based on staker's removed shares and operator's totalMagnitude
     * @param operator The operator to decrease the delegated delegated shares for
     * @param staker The staker to decrease the delegated delegated shares for
     * @param strategy The strategy to decrease the delegated delegated shares for
     * @param delegatedShares The delegatedShares to remove from the operator's delegated shares
     */
    function _decreaseDelegation(
        address operator,
        address staker,
        IStrategy strategy,
        DelegatedShares delegatedShares
    ) internal {
        // based on total magnitude, decrement operator's delegatedShares
        operatorDelegatedShares[operator][strategy] = operatorDelegatedShares[operator][strategy].sub(delegatedShares);
        
        // TODO: What to do about event wrt scaling?
        emit OperatorSharesDecreased(operator, staker, strategy, delegatedShares);
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
        OwnedShares[] memory ownedSharesToWithdraw,
        uint64[] memory totalMagnitudes
    ) internal returns (bytes32) {
        require(staker != address(0), InputAddressZero());
        require(strategies.length != 0, InputArrayLengthZero());

        DelegatedShares[] memory delegatedSharesToWithdraw = new DelegatedShares[](strategies.length);
        // Remove shares from staker and operator
        // Each of these operations fail if we attempt to remove more shares than exist
        for (uint256 i = 0; i < strategies.length; ++i) {
            IShareManager shareManager = _getShareManager(strategies[i]);

            // delegatedShares for staker to place into queueWithdrawal
            DelegatedShares delegatedSharesToRemove = ownedSharesToWithdraw[i].toDelegatedShares(totalMagnitudes[i]);
            // TODO: should this include beaconChainScalingFactor?            
            Shares sharesToWithdraw = delegatedSharesToRemove.toShares(stakerScalingFactors[staker][strategies[i]]);
            // TODO: maybe have a getter to get shares for all strategies, like getDelegatableShares
            // check sharesToWithdraw is valid
            // but for inputted strategies
            Shares sharesWithdrawable = shareManager.stakerStrategyShares(staker, strategies[i]);
            require(sharesToWithdraw.unwrap() <= sharesWithdrawable.unwrap(), WithdrawalExeedsMax());

            // Similar to `isDelegated` logic
            if (operator != address(0)) {
                // forgefmt: disable-next-item
                _decreaseDelegation({
                    operator: operator, 
                    staker: staker, 
                    strategy: strategies[i], 
                    delegatedShares: delegatedSharesToRemove
                });
            }
            delegatedSharesToWithdraw[i] = delegatedSharesToRemove.scaleForQueueWithdrawal(stakerScalingFactors[staker][strategies[i]]);

            // Remove active shares from EigenPodManager/StrategyManager
            // EigenPodManager: this call will revert if it would reduce the Staker's virtual beacon chain ETH shares below zero
            // StrategyManager: this call will revert if `sharesToDecrement` exceeds the Staker's current deposit shares in `strategies[i]`
            shareManager.removeShares(staker, strategies[i], sharesToWithdraw);
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
            delegatedShares: delegatedSharesToWithdraw
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

    /**
     * @notice helper to calculate the new depositScalingFactor after adding shares. This is only used
     * when a staker is depositing through the StrategyManager or EigenPodManager. A staker's depositScalingFactor
     * is only updated when they have new deposits and their shares are being increased.
     */
    function _updateDepositScalingFactor(
        address staker,
        IStrategy strategy,
        uint64 totalMagnitude,
        Shares existingShares,
        OwnedShares addedOwnedShares
    ) internal {
        uint256 newDepositScalingFactor;
        if (existingShares.unwrap() == 0) {
            newDepositScalingFactor = WAD / totalMagnitude;
        } else {
            // otherwise since
            //
            // newShares
            //      = existingShares + addedShares
            //      = existingPrincipalShares.toDelegatedShares(stakerScalingFactors[staker][strategy).toOwnedShares(totalMagnitude) + addedShares
            //
            // and it also is
            //
            // newShares
            //     = newPrincipalShares.toDelegatedShares(stakerScalingFactors[staker][strategy).toOwnedShares(totalMagnitude)
            //     = newPrincipalShares * newDepositScalingFactor / WAD * beaonChainScalingFactor / WAD * totalMagnitude / WAD
            //     = (existingPrincipalShares + addedShares) * newDepositScalingFactor / WAD * beaonChainScalingFactor / WAD * totalMagnitude / WAD 
            //
            // we can solve for
            //
            OwnedShares existingOwnedShares =
                existingShares
                .toDelegatedShares(stakerScalingFactors[staker][strategy])
                .toOwnedShares(totalMagnitude);
            newDepositScalingFactor =
                existingOwnedShares
                .add(addedOwnedShares)
                .unwrap()
                .divWad(existingShares.unwrap() + addedOwnedShares.unwrap())
                .divWad(stakerScalingFactors[staker][strategy].getBeaconChainScalingFactor())
                .divWad(totalMagnitude);
        }

        // update the staker's depositScalingFactor
        stakerScalingFactors[staker][strategy].depositScalingFactor = newDepositScalingFactor;
    }

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
     * @notice Getter function for the current EIP-712 domain separator for this contract.
     *
     * @dev The domain separator will change in the event of a fork that changes the ChainID.
     * @dev By introducing a domain separator the DApp developers are guaranteed that there can be no signature collision.
     * for more detailed information please read EIP-712.
     */
    function domainSeparator() public view returns (bytes32) {
        if (block.chainid == ORIGINAL_CHAIN_ID) {
            return _DOMAIN_SEPARATOR;
        } else {
            return _calculateDomainSeparator();
        }
    }

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

    /// @notice a legacy function that returns the total delegated shares for an operator and strategy
    function operatorShares(address operator, IStrategy strategy) public view returns (uint256) {
        uint64 totalMagnitude = allocationManager.getTotalMagnitude(operator, strategy);
        return operatorDelegatedShares[operator][strategy].toOwnedShares(totalMagnitude).unwrap();
    }

    /**
     * @notice Given a staker and a set of strategies, return the shares they can queue for withdrawal.
     * This value depends on which operator the staker is delegated to.
     * The shares amount returned is the actual amount of Strategy shares the staker would receive (subject
     * to each strategy's underlying shares to token ratio).
     */
    function getDelegatableShares(
        address staker,
        IStrategy[] memory strategies
    ) public view returns (OwnedShares[] memory ownedShares) {
        address operator = delegatedTo[staker];
        for (uint256 i = 0; i < strategies.length; ++i) {
            IShareManager shareManager = _getShareManager(strategies[i]);
            // TODO: batch call for strategyManager shares?
            // 1. read strategy deposit shares
            
            // forgefmt: disable-next-item
            Shares shares = shareManager.stakerStrategyShares(staker, strategies[i]);

            // 2. if the staker is delegated, actual withdrawable shares can be different from what is stored
            // in the StrategyManager/EigenPodManager because they could have been slashed
            if (operator != address(0)) {
                uint64 totalMagnitude = allocationManager.getTotalMagnitude(operator, strategies[i]);
                
                // forgefmt: disable-next-item
                ownedShares[i] = shares
                    .toDelegatedShares(stakerScalingFactors[staker][strategies[i]])
                    .toOwnedShares(totalMagnitude);
            }
        }
        return ownedShares;
    }

    /**
     * @notice Returns the number of actively-delegatable shares a staker has across all strategies.
     * @dev Returns two empty arrays in the case that the Staker has no actively-delegateable shares.
     */
    function getDelegatableShares(
        address staker
    ) public view returns (IStrategy[] memory, OwnedShares[] memory) {
        // Get a list of all the strategies to check
        IStrategy[] memory strategies = strategyManager.getStakerStrategyList(staker);
        // resize and add beaconChainETH to the end
        assembly {
            mstore(strategies, add(mload(strategies), 1))
        }
        strategies[strategies.length - 1] = beaconChainETHStrategy;

        // get the delegatable shares for each strategy
        OwnedShares[] memory shares = getDelegatableShares(staker, strategies);

        // if the last shares are 0, remove them
        if (shares[strategies.length - 1].unwrap() == 0) {
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
        // fetch the staker's current nonce
        uint256 currentStakerNonce = stakerNonce[staker];
        // calculate the digest hash
        return calculateStakerDelegationDigestHash(staker, currentStakerNonce, operator, expiry);
    }

    /**
     * @notice Calculates the digest hash to be signed and used in the `delegateToBySignature` function
     * @param staker The signing staker
     * @param _stakerNonce The nonce of the staker. In practice we use the staker's current nonce, stored at `stakerNonce[staker]`
     * @param operator The operator who is being delegated to
     * @param expiry The desired expiry time of the staker's signature
     */
    function calculateStakerDelegationDigestHash(
        address staker,
        uint256 _stakerNonce,
        address operator,
        uint256 expiry
    ) public view returns (bytes32) {
        // calculate the struct hash
        bytes32 stakerStructHash =
            keccak256(abi.encode(STAKER_DELEGATION_TYPEHASH, staker, operator, _stakerNonce, expiry));
        // calculate the digest hash
        bytes32 stakerDigestHash = keccak256(abi.encodePacked("\x19\x01", domainSeparator(), stakerStructHash));
        return stakerDigestHash;
    }

    /**
     * @notice Calculates the digest hash to be signed by the operator's delegationApprove and used in the `delegateTo` and `delegateToBySignature` functions.
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
    ) public view returns (bytes32) {
        // calculate the struct hash
        bytes32 approverStructHash = keccak256(
            abi.encode(DELEGATION_APPROVAL_TYPEHASH, _delegationApprover, staker, operator, approverSalt, expiry)
        );
        // calculate the digest hash
        bytes32 approverDigestHash = keccak256(abi.encodePacked("\x19\x01", domainSeparator(), approverStructHash));
        return approverDigestHash;
    }

    /**
     * @dev Recalculates the domain separator when the chainid changes due to a fork.
     */
    function _calculateDomainSeparator() internal view returns (bytes32) {
        return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("EigenLayer")), block.chainid, address(this)));
    }
}
