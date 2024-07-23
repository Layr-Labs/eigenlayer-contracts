// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "../permissions/Pausable.sol";
import "../libraries/EIP1271SignatureUtils.sol";
import "./AllocatorManagerStorage.sol";

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
contract AllocationManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AllocatorManagerStorage,
    ReentrancyGuardUpgradeable
{
    // @dev Index for flag that pauses new delegations when set
    uint8 internal constant PAUSED_NEW_DELEGATION = 0;

    // @dev Index for flag that pauses queuing new withdrawals when set.
    uint8 internal constant PAUSED_ENTER_WITHDRAWAL_QUEUE = 1;

    // @dev Index for flag that pauses completing existing withdrawals when set.
    uint8 internal constant PAUSED_EXIT_WITHDRAWAL_QUEUE = 2;

    // @dev Chain ID at the time of contract deployment
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    // @dev Maximum Value for `stakerOptOutWindowBlocks`. Approximately equivalent to 6 months in blocks.
    uint256 public constant MAX_STAKER_OPT_OUT_WINDOW_BLOCKS = (180 days) / 12;

    /// @notice Canonical, virtual beacon chain ETH strategy
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    // @notice Simple permission for functions that are only callable by the StrategyManager contract OR by the EigenPodManagerContract
    modifier onlyStrategyManagerOrEigenPodManager() {
        require(
            msg.sender == address(strategyManager) || msg.sender == address(eigenPodManager),
            "DelegationManager: onlyStrategyManagerOrEigenPodManager"
        );
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
        IDelegationManager _delegationManager, 
        IEigenPodManager _eigenPodManager, 
        IAVSDirectory _avsDirectory
    ) AllocatorManagerStorage(_strategyManager, _delegationManager, _eigenPodManager, _avsDirectory) {
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
        uint256 initialPausedStatus,
        uint256 _minWithdrawalDelayBlocks,
        IStrategy[] calldata _strategies,
        uint256[] calldata _withdrawalDelayBlocks
    ) external initializer {
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _DOMAIN_SEPARATOR = _calculateDomainSeparator();
        _transferOwnership(initialOwner);
        _setMinWithdrawalDelayBlocks(_minWithdrawalDelayBlocks);
        _setStrategyWithdrawalDelayBlocks(_strategies, _withdrawalDelayBlocks);
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /**
     * Allows the staker, the staker's operator, or that operator's delegationApprover to undelegate
     * a staker from their operator. Undelegation immediately removes ALL active shares/strategies from
     * both the staker and operator, and places the shares and strategies in the withdrawal queue
     */
    function undelegate(address staker)
        external
        onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE)
        returns (bytes32[] memory withdrawalRoots)
    {
        require(isDelegated(staker), "DelegationManager.undelegate: staker must be delegated to undelegate");
        require(!isAllocator(staker), "DelegationManager.undelegate: operators cannot be undelegated");
        require(staker != address(0), "DelegationManager.undelegate: cannot undelegate zero address");
        address allocator = delegatedTo(staker);
        require(
            msg.sender == staker || msg.sender == allocator
                || msg.sender == _allocatorDetails[allocator].delegationApprover,
            "DelegationManager.undelegate: caller cannot undelegate staker"
        );

        // Gather strategies and shares to remove from staker/operator during undelegation
        // Undelegation removes ALL currently-active strategies and shares
        (IStrategy[] memory strategies, uint256[] memory shares) = getDelegatableShares(staker);

        // emit an event if this action was not initiated by the staker themselves
        if (msg.sender != staker) {
            emit StakerForceUndelegated(staker, allocator);
        }

        // undelegate the staker
        emit StakerUndelegated(staker, allocator);
        _delegatedTo[staker] = address(0);

        // if no delegatable shares, return an empty array, and don't queue a withdrawal
        if (strategies.length == 0) {
            withdrawalRoots = new bytes32[](0);
        } else {
            withdrawalRoots = new bytes32[](strategies.length);
            for (uint256 i = 0; i < strategies.length; i++) {
                IStrategy[] memory singleStrategy = new IStrategy[](1);
                uint256[] memory singleShare = new uint256[](1);
                singleStrategy[0] = strategies[i];
                singleShare[0] = shares[i];

                withdrawalRoots[i] = _removeSharesAndQueueWithdrawal({
                    staker: staker,
                    allocator: allocator,
                    withdrawer: staker,
                    strategies: singleStrategy,
                    shares: singleShare
                });
            }
        }

        return withdrawalRoots;
    }

    /**
     * @notice Registers the caller as an allocator in EigenLayer.
     * @param delegationApprover is the address that must approve of the delegation of a staker's stake to the allocator. If set to 0, no approval is 
	 * required.
     * @param metadataURI is a URI for the allocator's metadata, i.e. a link providing more details on the allocator.
     *
     * @dev Once an allocator is registered, they cannot 'allocator' as an operator, and they will forever be considered "delegated themself".
     * @dev Note that the `metadataURI` is *never stored * and is only emitted in the `AllocatorMetadataURIUpdated` event
     * @dev reverts if the caller is delegated to a pre-SDA operator
     */
    function registerAsAllocator(
        address delegationApprover,
        string calldata metadataURI
    ) external{
        require(delegatedTo(msg.sender) == address(0), "DelegationManager.registerAsAllocator: caller is an operator or a staker");
        _allocatorDetails[msg.sender] = AllocatorDetails({
            delegationApprover: delegationApprover,
            isAllocator: true
        });
        _setDelegationApprover(msg.sender, delegationApprover);
        // TODO: events
    }

    /**
     * @notice Called by an allocator to emit an `AllocatorMetadataURIUpdated` event indicating the information has updated.
     * @param metadataURI The URI for metadata associated with an allocator
     * @dev Note that the `metadataURI` is *never stored * and is only emitted in the `AllocatorMetadataURIUpdated` event
     */
    function updateAllocatorMetadataURI(string calldata metadataURI) external {
        require(isAllocator(msg.sender), "DelegationManager.updateOperatorMetadataURI: caller must be an operator");
        emit AllocatorMetadataURIUpdated(msg.sender, metadataURI);
    }

	/**
	 * @notice Updates the delegation approver for the calling allocator.
	 *
	 * @param delegationApprover is the address that must approve of the delegations of a staker's stake to the allocator. If set to 0, no approval is
	 * required.
	 */
    function setDelegationApprover(address delegationApprover) external {
        _setDelegationApprover(msg.sender, delegationApprover);
    }

    function _setDelegationApprover(address allocator, address delegationApprover) internal {
        require(isAllocator(allocator), "DelegationManager._setDelegationApprover: allocator is not registered in EigenLayer");
        _allocatorDetails[allocator].delegationApprover = delegationApprover;
        // TODO: events
    }

    /**
	 * @notice Completes a handoff queued via queueHandoff.
	 *
	 * @param operator the operator in the queued handoff
	 * @param allocator the allocator in the queued handoff
	 * @param strategies the strategies to be handed off
	 * 
	 * @dev must be called 14 days after the handoff was queued
	 * @dev the allocator's shares are incremented by the operator's shares for each strategy and the operator's shares are decremented by the operator's
	 * shares for each strategy.
	 * @dev if all strategies are not handed off, this function can be called by anyone else to 
	 * complete the handoff for different strategies
	 */
	function completeAllocatorHandoff(address operator, address allocator, IStrategy[] calldata strategies) external {
        uint256[] memory shares = delegationManager.getMigratableOperatorShares(operator, allocator, strategies);
        for (uint256 i = 0; i < strategies.length; i++) {
            // increment the allocator's shares by the operator's shares for each strategy
            // TODO: staker address?
            _increaseAllocatorShares({allocator: allocator, staker: address(0), strategy: strategies[i], shares: shares[i]});
        }
    }

    /**
	 * @notice Delegates to an allocator for the calling staker
	 *
	 * @param allocator the allocator delegated to by the calling staker
	 * @param approverSignatureAndExpiry Verifies the operator approves of this delegation
     * @param approverSalt A unique single use value tied to an individual signature.
     * @dev The approverSignatureAndExpiry is used in the event that:
     *          1) the allocator's `delegationApprover` address is set to a non-zero value.
     *                  AND
     *          2) neither the allocator nor their `delegationApprover` is the `msg.sender`, since in the event that the allocator
     *             or their delegationApprover is the `msg.sender`, then approval is assumed.
	 *
	 * @dev Reverts if the staker is delegated to an allocator or if the staker's M2 operator has handed off to an allocator
	 */
	function delegateTo(
		address allocator,
		SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt
	) external {
        require(delegationManager.isDelegated(msg.sender), "DelegationManager.delegateTo: staker is already actively delegated in pre-SDA");
        require(delegatedTo(msg.sender) == address(0), "DelegationManager.delegateTo: staker is already actively delegated");
        require(isAllocator(allocator), "DelegationManager.delegateTo: allocator is not registered in EigenLayer");
        // go through the internal delegation flow, checking the `approverSignatureAndExpiry` if applicable
        _delegate(msg.sender, allocator, approverSignatureAndExpiry, approverSalt);
    }

    /**
     * Allows a staker to withdraw some shares. Withdrawn shares/strategies are immediately removed
     * from the staker. If the staker is delegated, withdrawn shares/strategies are also removed from
     * their operator.
     *
     * All withdrawn shares/strategies are placed in a queue and can be fully withdrawn after a delay.
     */
    function queueWithdrawals(QueuedWithdrawalParams[] calldata queuedWithdrawalParams)
        external
        onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE)
        returns (bytes32[] memory)
    {
        bytes32[] memory withdrawalRoots = new bytes32[](queuedWithdrawalParams.length);
        address allocator = delegatedTo(msg.sender);

        for (uint256 i = 0; i < queuedWithdrawalParams.length; i++) {
            require(
                queuedWithdrawalParams[i].strategies.length == queuedWithdrawalParams[i].shares.length,
                "DelegationManager.queueWithdrawal: input length mismatch"
            );
            require(
                queuedWithdrawalParams[i].withdrawer == msg.sender,
                "DelegationManager.queueWithdrawal: withdrawer must be staker"
            );

            // Remove shares from staker's strategies and place strategies/shares in queue.
            // If the staker is delegated to an allocator, the allocator's delegated shares are also reduced
            // NOTE: This will fail if the staker doesn't have the shares implied by the input parameters
            withdrawalRoots[i] = _removeSharesAndQueueWithdrawal({
                staker: msg.sender,
                allocator: allocator,
                withdrawer: queuedWithdrawalParams[i].withdrawer,
                strategies: queuedWithdrawalParams[i].strategies,
                shares: queuedWithdrawalParams[i].shares
            });
        }
        return withdrawalRoots;
    }

    /**
     * @notice Used to complete the specified `withdrawal`. The caller must match `withdrawal.withdrawer`
     * @param withdrawal The Withdrawal to complete.
     * @param tokens Array in which the i-th entry specifies the `token` input to the 'withdraw' function of the i-th Strategy in the `withdrawal.strategies` array.
     * This input can be provided with zero length if `receiveAsTokens` is set to 'false' (since in that case, this input will be unused)
     * @param middlewareTimesIndex is the index in the operator that the staker who triggered the withdrawal was delegated to's middleware times array
     * @param receiveAsTokens If true, the shares specified in the withdrawal will be withdrawn from the specified strategies themselves
     * and sent to the caller, through calls to `withdrawal.strategies[i].withdraw`. If false, then the shares in the specified strategies
     * will simply be transferred to the caller directly.
     * @dev middlewareTimesIndex is unused, but will be used in the Slasher eventually
     * @dev beaconChainETHStrategy shares are non-transferrable, so if `receiveAsTokens = false` and `withdrawal.withdrawer != withdrawal.staker`, note that
     * any beaconChainETHStrategy shares in the `withdrawal` will be _returned to the staker_, rather than transferred to the withdrawer, unlike shares in
     * any other strategies, which will be transferred to the withdrawer.
     */
    function completeQueuedWithdrawal(
        Withdrawal calldata withdrawal,
        IERC20[] calldata tokens,
        uint256 middlewareTimesIndex,
        bool receiveAsTokens
    ) external onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE) nonReentrant {
        _completeQueuedWithdrawal(withdrawal, tokens, middlewareTimesIndex, receiveAsTokens);
    }

    /**
     * @notice Array-ified version of `completeQueuedWithdrawal`.
     * Used to complete the specified `withdrawals`. The function caller must match `withdrawals[...].withdrawer`
     * @param withdrawals The Withdrawals to complete.
     * @param tokens Array of tokens for each Withdrawal. See `completeQueuedWithdrawal` for the usage of a single array.
     * @param middlewareTimesIndexes One index to reference per Withdrawal. See `completeQueuedWithdrawal` for the usage of a single index.
     * @param receiveAsTokens Whether or not to complete each withdrawal as tokens. See `completeQueuedWithdrawal` for the usage of a single boolean.
     * @dev See `completeQueuedWithdrawal` for relevant dev tags
     */
    function completeQueuedWithdrawals(
        Withdrawal[] calldata withdrawals,
        IERC20[][] calldata tokens,
        uint256[] calldata middlewareTimesIndexes,
        bool[] calldata receiveAsTokens
    ) external onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE) nonReentrant {
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            _completeQueuedWithdrawal(withdrawals[i], tokens[i], middlewareTimesIndexes[i], receiveAsTokens[i]);
        }
    }

    /**
     * @notice Increases a staker's delegated share balance in a strategy.
     * @param staker The address to increase the delegated shares for their operator.
     * @param strategy The strategy in which to increase the delegated shares.
     * @param shares The number of shares to increase.
     *
     * @dev *If the staker is actively delegated*, then increases the `staker`'s delegated shares in `strategy` by `shares`. Otherwise does nothing.
     * @dev Callable only by the StrategyManager or EigenPodManager.
     */
    function increaseDelegatedShares(
        address staker,
        IStrategy strategy,
        uint256 shares
    ) external onlyStrategyManagerOrEigenPodManager {
        // if the staker is delegated to an operator
        address allocator = delegatedTo(staker);
        if (allocator != address(0)) {
            // add strategy shares to delegate's shares
            _increaseAllocatorShares({allocator: allocator, staker: staker, strategy: strategy, shares: shares});
        }
    }

    /**
     * @notice Decreases a staker's delegated share balance in a strategy.
     * @param staker The address to increase the delegated shares for their operator.
     * @param strategy The strategy in which to decrease the delegated shares.
     * @param shares The number of shares to decrease.
     *
     * @dev *If the staker is actively delegated*, then decreases the `staker`'s delegated shares in `strategy` by `shares`. Otherwise does nothing.
     * @dev Callable only by the StrategyManager or EigenPodManager.
     */
    function decreaseDelegatedShares(
        address staker,
        IStrategy strategy,
        uint256 shares
    ) external onlyStrategyManagerOrEigenPodManager {
        // if the staker is delegated to an allocator
        address allocator = delegatedTo(staker);
        if (allocator != address(0)) {
            // forgefmt: disable-next-item
            // subtract strategy shares from delegate's shares
            _decreaseAllocatorShares({
                allocator: allocator, 
                staker: staker, 
                strategy: strategy, 
                shares: shares
            });
        }
    }

    /**
     * @notice Owner-only function for modifying the value of the `minWithdrawalDelayBlocks` variable.
     * @param newMinWithdrawalDelayBlocks new value of `minWithdrawalDelayBlocks`.
     */
    function setMinWithdrawalDelayBlocks(uint256 newMinWithdrawalDelayBlocks) external onlyOwner {
        _setMinWithdrawalDelayBlocks(newMinWithdrawalDelayBlocks);
    }

    /**
     * @notice Called by owner to set the minimum withdrawal delay blocks for each passed in strategy
     * Note that the min number of blocks to complete a withdrawal of a strategy is
     * MAX(minWithdrawalDelayBlocks, strategyWithdrawalDelayBlocks[strategy])
     * @param strategies The strategies to set the minimum withdrawal delay blocks for
     * @param withdrawalDelayBlocks The minimum withdrawal delay blocks to set for each strategy
     */
    function setStrategyWithdrawalDelayBlocks(
        IStrategy[] calldata strategies,
        uint256[] calldata withdrawalDelayBlocks
    ) external onlyOwner {
        _setStrategyWithdrawalDelayBlocks(strategies, withdrawalDelayBlocks);
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Delegates *from* a `staker` *to* an `allocator`.
     * @param staker The address to delegate *from* -- this address is delegating control of its own assets.
     * @param allocator The address to delegate *to* -- this address is being given power to place the `staker`'s assets at risk on services
     * @param approverSignatureAndExpiry Verifies the operator approves of this delegation
     * @param approverSalt Is a salt used to help guarantee signature uniqueness. Each salt can only be used once by a given approver.
     * @dev Assumes the following is checked before calling this function:
     *          1) the `staker` is not already delegated to an operator
     *          2) the `allocator` has indeed registered as an operator in EigenLayer
     * Ensures that:
     *          1) if applicable, that the approver signature is valid and non-expired
     *          2) new delegations are not paused (PAUSED_NEW_DELEGATION)
     */
    function _delegate(
        address staker,
        address allocator,
        SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt
    ) internal onlyWhenNotPaused(PAUSED_NEW_DELEGATION) {
        // fetch the allocator's `delegationApprover` address and store it in memory in case we need to use it multiple times
        address _delegationApprover = _allocatorDetails[allocator].delegationApprover;
        /**
         * Check the `_delegationApprover`'s signature, if applicable.
         * If the `_delegationApprover` is the zero address, then the allocator allows all stakers to delegate to them and this verification is skipped.
         * If the `_delegationApprover` or the `allocators` themselves is the caller, then approval is assumed and signature verification is skipped as well.
         */
        if (_delegationApprover != address(0) && msg.sender != _delegationApprover && msg.sender != allocator) {
            // check the signature expiry
            require(
                approverSignatureAndExpiry.expiry >= block.timestamp,
                "DelegationManager._delegate: approver signature expired"
            );
            // check that the salt hasn't been used previously, then mark the salt as spent
            require(
                !delegationApproverSaltIsSpent[_delegationApprover][approverSalt],
                "DelegationManager._delegate: approverSalt already spent"
            );
            delegationApproverSaltIsSpent[_delegationApprover][approverSalt] = true;

            // forgefmt: disable-next-item
            // calculate the digest hash
            bytes32 approverDigestHash = calculateDelegationApprovalDigestHash(
                staker, 
                allocator, 
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

        // record the delegation relation between the staker and allocators, and emit an event
        _delegatedTo[staker] = allocator;
        emit StakerDelegated(staker, allocator);

        (IStrategy[] memory strategies, uint256[] memory shares) = getDelegatableShares(staker);

        for (uint256 i = 0; i < strategies.length;) {
            // forgefmt: disable-next-item
            _increaseAllocatorShares({
                allocator: allocator, 
                staker: staker, 
                strategy: strategies[i], 
                shares: shares[i]
            });

            unchecked {
                ++i;
            }
        }
    }

    /**
     * @dev commented-out param (middlewareTimesIndex) is the index in the operator that the staker who triggered the withdrawal was delegated to's middleware times array
     * This param is intended to be passed on to the Slasher contract, but is unused in the M2 release of these contracts, and is thus commented-out.
     */
    function _completeQueuedWithdrawal(
        Withdrawal calldata withdrawal,
        IERC20[] calldata tokens,
        uint256, /*middlewareTimesIndex*/
        bool receiveAsTokens
    ) internal {
        bytes32 withdrawalRoot = calculateWithdrawalRoot(withdrawal);

        require(
            pendingWithdrawals[withdrawalRoot], "DelegationManager._completeQueuedWithdrawal: action is not in queue"
        );

        require(
            withdrawal.startBlock + minWithdrawalDelayBlocks <= block.number,
            "DelegationManager._completeQueuedWithdrawal: minWithdrawalDelayBlocks period has not yet passed"
        );

        require(
            msg.sender == withdrawal.withdrawer,
            "DelegationManager._completeQueuedWithdrawal: only withdrawer can complete action"
        );

        if (receiveAsTokens) {
            require(
                tokens.length == withdrawal.strategies.length,
                "DelegationManager._completeQueuedWithdrawal: input length mismatch"
            );
        }

        // Remove `withdrawalRoot` from pending roots
        delete pendingWithdrawals[withdrawalRoot];

        if (receiveAsTokens) {
            // Finalize action by converting shares to tokens for each strategy, or
            // by re-awarding shares in each strategy.
            for (uint256 i = 0; i < withdrawal.strategies.length;) {
                require(
                    withdrawal.startBlock + strategyWithdrawalDelayBlocks[withdrawal.strategies[i]] <= block.number,
                    "DelegationManager._completeQueuedWithdrawal: withdrawalDelayBlocks period has not yet passed for this strategy"
                );

                _withdrawSharesAsTokens({
                    staker: withdrawal.staker,
                    withdrawer: msg.sender,
                    strategy: withdrawal.strategies[i],
                    shares: withdrawal.shares[i],
                    token: tokens[i]
                });
                unchecked {
                    ++i;
                }
            }
            // Award shares back in StrategyManager/EigenPodManager. If withdrawer is delegated, increase the shares delegated to the operator
        } else {
            // Award shares back in StrategyManager/EigenPodManager.
            // If withdrawer is delegated, increase the shares delegated to the allocator.
            address currentAllocator = delegatedTo(msg.sender);
            for (uint256 i = 0; i < withdrawal.strategies.length;) {
                require(
                    withdrawal.startBlock + strategyWithdrawalDelayBlocks[withdrawal.strategies[i]] <= block.number,
                    "DelegationManager._completeQueuedWithdrawal: withdrawalDelayBlocks period has not yet passed for this strategy"
                );

                /**
                 * When awarding podOwnerShares in EigenPodManager, we need to be sure to only give them back to the original podOwner.
                 * Other strategy shares can + will be awarded to the withdrawer.
                 */
                if (withdrawal.strategies[i] == beaconChainETHStrategy) {
                    address staker = withdrawal.staker;
                    /**
                     * Update shares amount depending upon the returned value.
                     * The return value will be lower than the input value in the case where the staker has an existing share deficit
                     */
                    uint256 increaseInDelegateableShares =
                        eigenPodManager.addShares({podOwner: staker, shares: withdrawal.shares[i]});
                    address podOwnerAllocator = delegatedTo(staker);
                    // Similar to `isDelegated` logic
                    if (podOwnerAllocator != address(0)) {
                        _increaseAllocatorShares({
                            allocator: podOwnerAllocator,
                            // the 'staker' here is the address receiving new shares
                            staker: staker,
                            strategy: withdrawal.strategies[i],
                            shares: increaseInDelegateableShares
                        });
                    }
                } else {
                    strategyManager.addShares(msg.sender, tokens[i], withdrawal.strategies[i], withdrawal.shares[i]);
                    // Similar to `isDelegated` logic
                    if (currentAllocator != address(0)) {
                        _increaseAllocatorShares({
                            allocator: currentAllocator,
                            // the 'staker' here is the address receiving new shares
                            staker: msg.sender,
                            strategy: withdrawal.strategies[i],
                            shares: withdrawal.shares[i]
                        });
                    }
                }
                unchecked {
                    ++i;
                }
            }
        }

        emit WithdrawalCompleted(withdrawalRoot);
    }

    /// @notice Increases `allocator`s delegated shares in `strategy` by `shares` and emits an `OperatorSharesIncreased` event
    function _increaseAllocatorShares(address allocator, address staker, IStrategy strategy, uint256 shares) internal {
        scaledDelegatedShares[allocator][strategy] += shares * BIG_NUMBER / avsDirectory.totalMagnitude(allocator, strategy);
    }

    /// @notice Decreases `operator`s delegated shares in `strategy` by `shares` and emits an `OperatorSharesDecreased` event
    function _decreaseAllocatorShares(address allocator, address staker, IStrategy strategy, uint256 shares) internal {
        // This will revert on underflow, so no check needed
        scaledDelegatedShares[allocator][strategy] -= shares; // TODO: Decide scaling
    }

    /**
     * @notice Removes `shares` in `strategies` from `staker` who is currently delegated to `operator` and queues a withdrawal to the `withdrawer`.
     * @dev If the `operator` is indeed an operator, then the operator's delegated shares in the `strategies` are also decreased appropriately.
     * @dev If `withdrawer` is not the same address as `staker`, then thirdPartyTransfersForbidden[strategy] must be set to false in the StrategyManager.
     */
    function _removeSharesAndQueueWithdrawal(
        address staker,
        address allocator,
        address withdrawer,
        IStrategy[] memory strategies,
        uint256[] memory shares
    ) internal returns (bytes32) {
        require(
            staker != address(0), "DelegationManager._removeSharesAndQueueWithdrawal: staker cannot be zero address"
        );
        require(strategies.length != 0, "DelegationManager._removeSharesAndQueueWithdrawal: strategies cannot be empty");

        // Remove shares from staker and allocator
        // Each of these operations fail if we attempt to remove more shares than exist
        for (uint256 i = 0; i < strategies.length;) {
            // Similar to `isDelegated` logic
            if (allocator != address(0)) {
                // forgefmt: disable-next-item
                _decreaseAllocatorShares({
                    allocator: allocator, 
                    staker: staker, 
                    strategy: strategies[i], 
                    shares: shares[i]
                });
            }

            // Remove active shares from EigenPodManager/StrategyManager
            if (strategies[i] == beaconChainETHStrategy) {
                /**
                 * This call will revert if it would reduce the Staker's virtual beacon chain ETH shares below zero.
                 * This behavior prevents a Staker from queuing a withdrawal which improperly removes excessive
                 * shares from the operator to whom the staker is delegated.
                 * It will also revert if the share amount being withdrawn is not a whole Gwei amount.
                 */
                eigenPodManager.removeShares(staker, shares[i]);
            } else {
                require(
                    staker == withdrawer || !strategyManager.thirdPartyTransfersForbidden(strategies[i]),
                    "DelegationManager._removeSharesAndQueueWithdrawal: withdrawer must be same address as staker if thirdPartyTransfersForbidden are set"
                );
                // this call will revert if `shares[i]` exceeds the Staker's current shares in `strategies[i]`
                strategyManager.removeShares(staker, strategies[i], shares[i]);
            }

            unchecked {
                ++i;
            }
        }

        // Create queue entry and increment withdrawal nonce
        uint256 nonce = cumulativeWithdrawalsQueued[staker];
        cumulativeWithdrawalsQueued[staker]++;

        Withdrawal memory withdrawal = Withdrawal({
            staker: staker,
            delegatedTo: allocator,
            withdrawer: withdrawer,
            nonce: nonce,
            startBlock: uint32(block.number),
            strategies: strategies,
            shares: shares
        });

        bytes32 withdrawalRoot = calculateWithdrawalRoot(withdrawal);

        // Place withdrawal in queue
        pendingWithdrawals[withdrawalRoot] = true;

        emit WithdrawalQueued(withdrawalRoot, withdrawal);
        return withdrawalRoot;
    }

    /**
     * @notice Withdraws `shares` in `strategy` to `withdrawer`. If the shares are virtual beaconChainETH shares, then a call is ultimately forwarded to the
     * `staker`s EigenPod; otherwise a call is ultimately forwarded to the `strategy` with info on the `token`.
     */
    function _withdrawSharesAsTokens(
        address staker,
        address withdrawer,
        IStrategy strategy,
        uint256 shares,
        IERC20 token
    ) internal {
        if (strategy == beaconChainETHStrategy) {
            eigenPodManager.withdrawSharesAsTokens({podOwner: staker, destination: withdrawer, shares: shares});
        } else {
            strategyManager.withdrawSharesAsTokens(withdrawer, strategy, shares, token);
        }
    }

    function _setMinWithdrawalDelayBlocks(uint256 _minWithdrawalDelayBlocks) internal {
        require(
            _minWithdrawalDelayBlocks <= MAX_WITHDRAWAL_DELAY_BLOCKS,
            "DelegationManager._setMinWithdrawalDelayBlocks: _minWithdrawalDelayBlocks cannot be > MAX_WITHDRAWAL_DELAY_BLOCKS"
        );
        emit MinWithdrawalDelayBlocksSet(minWithdrawalDelayBlocks, _minWithdrawalDelayBlocks);
        minWithdrawalDelayBlocks = _minWithdrawalDelayBlocks;
    }

    /**
     * @notice Sets the withdrawal delay blocks for each strategy in `_strategies` to `_withdrawalDelayBlocks`.
     * gets called when initializing contract or by calling `setStrategyWithdrawalDelayBlocks`
     */
    function _setStrategyWithdrawalDelayBlocks(
        IStrategy[] calldata _strategies,
        uint256[] calldata _withdrawalDelayBlocks
    ) internal {
        require(
            _strategies.length == _withdrawalDelayBlocks.length,
            "DelegationManager._setStrategyWithdrawalDelayBlocks: input length mismatch"
        );
        uint256 numStrats = _strategies.length;
        for (uint256 i = 0; i < numStrats; ++i) {
            IStrategy strategy = _strategies[i];
            uint256 prevStrategyWithdrawalDelayBlocks = strategyWithdrawalDelayBlocks[strategy];
            uint256 newStrategyWithdrawalDelayBlocks = _withdrawalDelayBlocks[i];
            require(
                newStrategyWithdrawalDelayBlocks <= MAX_WITHDRAWAL_DELAY_BLOCKS,
                "DelegationManager._setStrategyWithdrawalDelayBlocks: _withdrawalDelayBlocks cannot be > MAX_WITHDRAWAL_DELAY_BLOCKS"
            );

            // set the new withdrawal delay blocks
            strategyWithdrawalDelayBlocks[strategy] = newStrategyWithdrawalDelayBlocks;
            emit StrategyWithdrawalDelayBlocksSet(
                strategy, prevStrategyWithdrawalDelayBlocks, newStrategyWithdrawalDelayBlocks
            );
        }
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
    function isDelegated(address staker) public view returns (bool) {
        return (delegatedToView(staker) != address(0));
    }

    function getDelegationApprover(address allocator) public view returns (address) {
        return _allocatorDetails[allocator].delegationApprover;
    }
    
    function isAllocator(address allocator) public view returns (bool) {
        return _allocatorDetails[allocator].isAllocator;
    }

    function delegatedToView(address staker) public view returns (address) {
        address allocator = _delegatedTo[staker];
        // If the staker is not actively delegated, check if they have a handoff
        if (allocator == address(0) && !isPostSDAStaker[staker]) {
            IDelegationManager.Handoff memory handoff = delegationManager.getHandoff(staker);
            if (handoff.completableTimestamp != 0 && handoff.completableTimestamp <= block.timestamp) {
                // If the handoff is completable, return the allocator and overwrite the allocator
                allocator = handoff.allocator;
            }
        }
        return allocator;
    }

    function delegatedTo(address staker) public returns (address) {
        address allocator = _delegatedTo[staker];
        // If the staker is not actively delegated, check if they have a handoff
        if (allocator == address(0) && !isPostSDAStaker[staker]) {
            IDelegationManager.Handoff memory handoff = delegationManager.getHandoff(staker);
            if (handoff.completableTimestamp != 0 && handoff.completableTimestamp <= block.timestamp) {
                // If the handoff is completable, return the allocator and overwrite the allocator
                isPostSDAStaker[staker] = true;
                _delegatedTo[staker] = handoff.allocator;
                allocator = handoff.allocator;
            }
        }

        isPostSDAStaker[staker] = true;
        return allocator;
    }

    /// @notice Given array of strategies, returns array of shares for the allocator
    function getScaledDelegatedShares(
        address allocator,
        IStrategy[] memory strategies
    ) public view returns (uint256[] memory) {
        uint256[] memory shares = new uint256[](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            shares[i] = scaledDelegatedShares[allocator][strategies[i]];
        }
        return shares;
    }

    /**
     * @notice Returns the number of actively-delegatable shares a staker has across all strategies.
     * @dev Returns two empty arrays in the case that the Staker has no actively-delegateable shares.
     */
    function getDelegatableShares(address staker) public view returns (IStrategy[] memory, uint256[] memory) {
        // Get currently active shares and strategies for `staker`
        int256 podShares = eigenPodManager.podOwnerShares(staker);
        (IStrategy[] memory strategyManagerStrats, uint256[] memory strategyManagerShares) =
            strategyManager.getDeposits(staker);

        // Has no shares in EigenPodManager, but potentially some in StrategyManager
        if (podShares <= 0) {
            return (strategyManagerStrats, strategyManagerShares);
        }

        IStrategy[] memory strategies;
        uint256[] memory shares;

        if (strategyManagerStrats.length == 0) {
            // Has shares in EigenPodManager, but not in StrategyManager
            strategies = new IStrategy[](1);
            shares = new uint256[](1);
            strategies[0] = beaconChainETHStrategy;
            shares[0] = uint256(podShares);
        } else {
            // Has shares in both

            // 1. Allocate return arrays
            strategies = new IStrategy[](strategyManagerStrats.length + 1);
            shares = new uint256[](strategies.length);

            // 2. Place StrategyManager strats/shares in return arrays
            for (uint256 i = 0; i < strategyManagerStrats.length;) {
                strategies[i] = strategyManagerStrats[i];
                shares[i] = strategyManagerShares[i];

                unchecked {
                    ++i;
                }
            }

            // 3. Place EigenPodManager strat/shares in return arrays
            strategies[strategies.length - 1] = beaconChainETHStrategy;
            shares[strategies.length - 1] = uint256(podShares);
        }

        return (strategies, shares);
    }

    /**
     * @notice Given a list of strategies, return the minimum number of blocks that must pass to withdraw
     * from all the inputted strategies. Return value is >= minWithdrawalDelayBlocks as this is the global min withdrawal delay.
     * @param strategies The strategies to check withdrawal delays for
     */
    function getWithdrawalDelay(IStrategy[] calldata strategies) public view returns (uint256) {
        uint256 withdrawalDelay = minWithdrawalDelayBlocks;
        for (uint256 i = 0; i < strategies.length; ++i) {
            uint256 currWithdrawalDelay = strategyWithdrawalDelayBlocks[strategies[i]];
            if (currWithdrawalDelay > withdrawalDelay) {
                withdrawalDelay = currWithdrawalDelay;
            }
        }
        return withdrawalDelay;
    }

    /// @notice Returns the keccak256 hash of `withdrawal`.
    function calculateWithdrawalRoot(Withdrawal memory withdrawal) public pure returns (bytes32) {
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
