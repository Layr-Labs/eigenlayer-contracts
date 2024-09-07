// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "../permissions/Pausable.sol";
import "../libraries/EIP1271SignatureUtils.sol";
import "../libraries/ShareScalingLib.sol";
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

    // The max configurable withdrawal delay per strategy. Set to 30 days in seconds
    uint256 public constant MAX_WITHDRAWAL_DELAY = 30 days;

    // the number of 12-second blocks in 30 days (60 * 60 * 24 * 30 / 12 = 216,000)
    uint256 public constant MAX_WITHDRAWAL_DELAY_BLOCKS = MAX_WITHDRAWAL_DELAY / 12;

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
        ISlasher _slasher,
        IEigenPodManager _eigenPodManager,
        IAVSDirectory _avsDirectory
    ) DelegationManagerStorage(_strategyManager, _slasher, _eigenPodManager, _avsDirectory) {
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
     * @notice Registers the caller as an operator in EigenLayer.
     * @param registeringOperatorDetails is the `OperatorDetails` for the operator.
     * @param allocationDelay is a one-time configurable delay for the operator when performing slashable magnitude allocations.
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
        require(!isDelegated(msg.sender), "DelegationManager.registerAsOperator: caller is already actively delegated");
        _setOperatorDetails(msg.sender, registeringOperatorDetails);
        SignatureWithExpiry memory emptySignatureAndExpiry;
        // delegate from the operator to themselves
        _delegate(msg.sender, msg.sender, emptySignatureAndExpiry, bytes32(0));
        // set the allocation delay
        _initializeAllocationDelay(allocationDelay);
        // emit events
        emit OperatorRegistered(msg.sender, registeringOperatorDetails);
        emit OperatorMetadataURIUpdated(msg.sender, metadataURI);
    }

    /**
     * @notice Called by operators to set their allocation delay one time. Cannot be updated
     * after being set. This delay is required to be set for an operator to be able to allocate slashable magnitudes.
     * @param delay the allocation delay in seconds
     */
    function initializeAllocationDelay(uint32 delay) external {
        _initializeAllocationDelay(delay);
    }

    /**
     * @notice Updates an operator's stored `OperatorDetails`.
     * @param newOperatorDetails is the updated `OperatorDetails` for the operator, to replace their current OperatorDetails`.
     *
     * @dev The caller must have previously registered as an operator in EigenLayer.
     */
    function modifyOperatorDetails(OperatorDetails calldata newOperatorDetails) external {
        require(isOperator(msg.sender), "DelegationManager.modifyOperatorDetails: caller must be an operator");
        _setOperatorDetails(msg.sender, newOperatorDetails);
    }

    /**
     * @notice Called by an operator to emit an `OperatorMetadataURIUpdated` event indicating the information has updated.
     * @param metadataURI The URI for metadata associated with an operator
     */
    function updateOperatorMetadataURI(string calldata metadataURI) external {
        require(isOperator(msg.sender), "DelegationManager.updateOperatorMetadataURI: caller must be an operator");
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
        require(!isDelegated(msg.sender), "DelegationManager.delegateTo: staker is already actively delegated");
        require(isOperator(operator), "DelegationManager.delegateTo: operator is not registered in EigenLayer");
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
        require(
            stakerSignatureAndExpiry.expiry >= block.timestamp,
            "DelegationManager.delegateToBySignature: staker signature expired"
        );
        require(!isDelegated(staker), "DelegationManager.delegateToBySignature: staker is already actively delegated");
        require(
            isOperator(operator), "DelegationManager.delegateToBySignature: operator is not registered in EigenLayer"
        );

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
    function undelegate(address staker)
        external
        onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE)
        returns (bytes32[] memory withdrawalRoots)
    {
        require(isDelegated(staker), "DelegationManager.undelegate: staker must be delegated to undelegate");
        require(!isOperator(staker), "DelegationManager.undelegate: operators cannot be undelegated");
        require(staker != address(0), "DelegationManager.undelegate: cannot undelegate zero address");
        address operator = delegatedTo[staker];
        require(
            msg.sender == staker || msg.sender == operator
                || msg.sender == _operatorDetails[operator].delegationApprover,
            "DelegationManager.undelegate: caller cannot undelegate staker"
        );

        // Gather strategies and scaled shares to remove from staker/operator during undelegation
        // Undelegation removes ALL currently-active strategies and shares
        (IStrategy[] memory strategies, uint256[] memory scaledShares) = getDelegatableScaledShares(staker);

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
            for (uint256 i = 0; i < strategies.length; i++) {
                IStrategy[] memory singleStrategy = new IStrategy[](1);
                uint256[] memory singleScaledShare = new uint256[](1);
                singleStrategy[0] = strategies[i];
                singleScaledShare[0] = scaledShares[i];

                withdrawalRoots[i] = _removeSharesAndQueueWithdrawal({
                    staker: staker,
                    operator: operator,
                    withdrawer: staker,
                    strategies: singleStrategy,
                    scaledShares: singleScaledShare
                });
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
    function queueWithdrawals(QueuedWithdrawalParams[] calldata queuedWithdrawalParams)
        external
        onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE)
        returns (bytes32[] memory)
    {
        bytes32[] memory withdrawalRoots = new bytes32[](queuedWithdrawalParams.length);
        address operator = delegatedTo[msg.sender];

        for (uint256 i = 0; i < queuedWithdrawalParams.length; i++) {
            require(
                queuedWithdrawalParams[i].strategies.length == queuedWithdrawalParams[i].shares.length,
                "DelegationManager.queueWithdrawal: input length mismatch"
            );
            require(
                queuedWithdrawalParams[i].withdrawer == msg.sender,
                "DelegationManager.queueWithdrawal: withdrawer must be staker"
            );

            // Withdrawer inputs the number of real Strategy shares they expect to withdraw,
            // convert accordingly to scaledShares in storage
            uint256[] memory scaledShares = ShareScalingLib.scaleShares(
                avsDirectory,
                operator,
                queuedWithdrawalParams[i].strategies,
                queuedWithdrawalParams[i].shares
            );

            // Remove shares from staker's strategies and place strategies/shares in queue.
            // If the staker is delegated to an operator, the operator's delegated shares are also reduced
            // NOTE: This will fail if the staker doesn't have the shares implied by the input parameters
            withdrawalRoots[i] = _removeSharesAndQueueWithdrawal({
                staker: msg.sender,
                operator: operator,
                withdrawer: queuedWithdrawalParams[i].withdrawer,
                strategies: queuedWithdrawalParams[i].strategies,
                scaledShares: scaledShares
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
     * @notice Increases a staker's delegated share balance in a strategy.
     * @param staker The address to increase the delegated scaled shares for their operator.
     * @param strategy The strategy in which to increase the delegated scaled shares.
     * @param scaledShares The number of scaled shares to increase.
     *
     * @dev *If the staker is actively delegated*, then increases the `staker`'s delegated shares in `strategy` by `scaledShares`. Otherwise does nothing.
     * @dev Callable only by the StrategyManager or EigenPodManager.
     */
    function increaseDelegatedScaledShares(
        address staker,
        IStrategy strategy,
        uint256 scaledShares
    ) external onlyStrategyManagerOrEigenPodManager {
        // if the staker is delegated to an operator
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];

            // add strategy shares to delegate's shares
            _increaseOperatorScaledShares({operator: operator, staker: staker, strategy: strategy, scaledShares: scaledShares});
        }
    }

    /**
     * @notice Decreases a staker's delegated share balance in a strategy.
     * @param staker The address to increase the delegated scaled shares for their operator.
     * @param strategy The strategy in which to decrease the delegated scaled shares.
     * @param scaledShares The number of scaled shares to decrease.
     *
     * @dev *If the staker is actively delegated*, then decreases the `staker`'s delegated scaled shares in `strategy` by `scaledShares`. Otherwise does nothing.
     * @dev Callable only by the StrategyManager or EigenPodManager.
     */
    function decreaseDelegatedScaledShares(
        address staker,
        IStrategy strategy,
        uint256 scaledShares
    ) external onlyStrategyManagerOrEigenPodManager {
        // if the staker is delegated to an operator
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];

            // forgefmt: disable-next-item
            // subtract strategy shares from delegated scaled shares
            _decreaseOperatorScaledShares({
                operator: operator, 
                staker: staker, 
                strategy: strategy, 
                scaledShares: scaledShares
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
     * @notice Called by owner to set the minimum withdrawal delay for each passed in strategy
     * Note that the min number of blocks to complete a withdrawal of a strategy is
     * MAX(minWithdrawalDelay, strategyWithdrawalDelay[strategy])
     * @param strategies The strategies to set the minimum withdrawal delay for
     * @param withdrawalDelays The minimum withdrawal delay (in seconds) to set for each strategy
     */
    function setStrategyWithdrawalDelay(
        IStrategy[] calldata strategies,
        uint256[] calldata withdrawalDelays
    ) external onlyOwner {
        _setStrategyWithdrawalDelay(strategies, withdrawalDelays);
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
        require(
            newOperatorDetails.stakerOptOutWindowBlocks <= MAX_STAKER_OPT_OUT_WINDOW_BLOCKS,
            "DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be > MAX_STAKER_OPT_OUT_WINDOW_BLOCKS"
        );
        require(
            newOperatorDetails.stakerOptOutWindowBlocks >= _operatorDetails[operator].stakerOptOutWindowBlocks,
            "DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be decreased"
        );
        _operatorDetails[operator] = newOperatorDetails;
        emit OperatorDetailsModified(msg.sender, newOperatorDetails);
    }

    /**
     * @notice Called by operators to set their allocation delay one time. Cannot be updated
     * after being set. This delay is required to be set for an operator to be able to allocate slashable magnitudes.
     * @param delay the allocation delay in seconds
     */
    function _initializeAllocationDelay(uint32 delay) internal {
        require(
            isOperator(msg.sender),
            "DelegationManager._initializeAllocationDelay: operator not registered to EigenLayer yet"
        );
        require(
            !_operatorAllocationDelay[msg.sender].isSet,
            "DelegationManager._initializeAllocationDelay: allocation delay already set"
        );
        _operatorAllocationDelay[msg.sender] = AllocationDelayDetails({
            isSet: true,
            allocationDelay: delay
        });
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

        (IStrategy[] memory strategies, uint256[] memory shares) = getDelegatableScaledShares(staker);

        // NOTE: technically the function above returns scaled shares, so it needs to be descaled before scaling to the
        // the new operator according to their totalMagnitude. However, since the staker is not delegated,
        // scaledShares = shares so we can skip descaling and read these as descaled shares.
        uint256[] memory scaledShares = ShareScalingLib.scaleShares(avsDirectory, operator, strategies, shares);

        for (uint256 i = 0; i < strategies.length;) {
            // forgefmt: disable-next-item
            _increaseOperatorScaledShares({
                operator: operator, 
                staker: staker, 
                strategy: strategies[i], 
                scaledShares: scaledShares[i]
            });

            // stakerStrategyScaledShares(staker, strategies[i]) needs to be updated to be equal to newly scaled shares.
            // we take the difference between the shares and the scaledShares to update the staker's scaledShares. The key property
            // here is that scaledShares should be >= shares and can never be less than shares due to totalMagnitude being
            // a monotonically decreasing value (see ShareScalingLib.scaleShares for more info).
            // The exact same property applies to the EigenPodManager and podOwnerScaledShares where real ETH shares are scaled based on
            // the delegated operator (if any) prior to being stored as scaledShares in the EigenPodManager.
            if (shares[i] < scaledShares[i]) {
                if (strategies[i] == beaconChainETHStrategy) {
                    eigenPodManager.addScaledShares(staker, scaledShares[i] - shares[i]);
                } else {
                    strategyManager.addScaledShares(
                        staker,
                        strategies[i].underlyingToken(), //TODO possibly remove this from the interface
                        strategies[i],
                        scaledShares[i] - shares[i]
                    );
                }
            }

            unchecked {
                ++i;
            }
        }
    }

    /**
     * @dev This function "descales" the scaledShares according to the totalMagnitude at the time of withdrawal completion.
     * This will apply any slashing that has occurred since the scaledShares were initially set in storage. 
     * If receiveAsTokens is true, then these descaled shares will be withdrawn as tokens.
     * If receiveAsTokens is false, then they will be rescaled according to the current operator the staker is delegated to, and added back to the operator's scaledShares.
     */
    function _completeQueuedWithdrawal(
        Withdrawal calldata withdrawal,
        IERC20[] calldata tokens,
        bool receiveAsTokens
    ) internal {
        bytes32 withdrawalRoot = calculateWithdrawalRoot(withdrawal);

        require(
            pendingWithdrawals[withdrawalRoot], "DelegationManager._completeQueuedWithdrawal: action is not in queue"
        );
        require(
            withdrawal.startTimestamp + minWithdrawalDelay <= block.number,
            "DelegationManager._completeQueuedWithdrawal: minWithdrawalDelay period has not yet passed"
        );
        require(
            msg.sender == withdrawal.withdrawer,
            "DelegationManager._completeQueuedWithdrawal: only withdrawer can complete action"
        );
        require(
            tokens.length == withdrawal.strategies.length,
            "DelegationManager._completeQueuedWithdrawal: input length mismatch"
        );

        // read delegated operator for scaling and adding shares back if needed
        address currentOperator = delegatedTo[msg.sender];
        // descale shares to get the "real" strategy share values of the withdrawal
        uint256[] memory descaledShares = ShareScalingLib.descaleSharesAtTimestamp(
            avsDirectory,
            withdrawal.delegatedTo,
            withdrawal.strategies,
            withdrawal.scaledShares,
            withdrawal.startTimestamp + avsDirectory.DEALLOCATION_DELAY()
        );

        if (receiveAsTokens) {
            // complete the withdrawal by converting descaled shares to tokens
            _completeReceiveAsTokens(
                withdrawal,
                tokens,
                descaledShares
            );
        } else {
            // Award shares back in StrategyManager/EigenPodManager.
            _completeReceiveAsShares(
                withdrawal,
                tokens,
                descaledShares
            );
        }

        // Remove `withdrawalRoot` from pending roots
        delete pendingWithdrawals[withdrawalRoot];
        emit WithdrawalCompleted(withdrawalRoot);
    }

    /// TODO: natspec
    function _completeReceiveAsTokens(
        Withdrawal calldata withdrawal,
        IERC20[] calldata tokens,
        uint256[] memory descaledShares
    ) internal {
        // Finalize action by converting scaled shares to tokens for each strategy, or
        // by re-awarding shares in each strategy.
        for (uint256 i = 0; i < withdrawal.strategies.length;) {
            require(
                withdrawal.startTimestamp + strategyWithdrawalDelays[withdrawal.strategies[i]] <= block.timestamp,
                "DelegationManager._completeReceiveAsTokens: withdrawalDelay period has not yet passed for this strategy"
            );

            _withdrawSharesAsTokens({
                staker: withdrawal.staker,
                withdrawer: msg.sender,
                strategy: withdrawal.strategies[i],
                shares: descaledShares[i],
                token: tokens[i]
            });
            unchecked {
                ++i;
            }
        }
    }

    /// TODO: natspec
    function _completeReceiveAsShares(
        Withdrawal calldata withdrawal,
        IERC20[] calldata tokens,
        uint256[] memory descaledShares
    ) internal {
        // // read delegated operator for scaling and adding shares back if needed
        // address currentOperator = delegatedTo[msg.sender];
        // // We scale shares again to the new totalMagnitude of the currentOperator.
        // uint256[] memory scaledShares = ShareScalingLib.scaleShares(
        //     avsDirectory,
        //     currentOperator,
        //     withdrawal.strategies,
        //     descaledShares
        // );

        // for (uint256 i = 0; i < withdrawal.strategies.length;) {
        //     require(
        //         withdrawal.startTimestamp + strategyWithdrawalDelays[withdrawal.strategies[i]] <= block.number,
        //         "DelegationManager._completeReceiveAsShares: withdrawalDelay period has not yet passed for this strategy"
        //     );

        //     /**
        //      * When awarding podOwnerShares in EigenPodManager, we need to be sure to only give them back to the original podOwner.
        //      * Other strategy shares can + will be awarded to the withdrawer.
        //      */
        //     if (withdrawal.strategies[i] == beaconChainETHStrategy) {
        //         address staker = withdrawal.staker;
        //         /**
        //          * Update shares amount depending upon the returned value.
        //          * The return value will be lower than the input value in the case where the staker has an existing share deficit
        //          */
        //         uint256 increaseInDelegateableScaledShares =
        //             eigenPodManager.addScaledShares({podOwner: staker, scaledShares: scaledShares[i]});
        //         address podOwnerOperator = delegatedTo[staker];
        //         // Similar to `isDelegated` logic
        //         if (podOwnerOperator != address(0)) {
        //             _increaseOperatorScaledShares({
        //                 operator: podOwnerOperator,
        //                 // the 'staker' here is the address receiving new shares
        //                 staker: staker,
        //                 strategy: withdrawal.strategies[i],
        //                 scaledShares: increaseInDelegateableScaledShares
        //             });
        //         }
        //     } else {
        //         strategyManager.addScaledShares(msg.sender, tokens[i], withdrawal.strategies[i], scaledShares[i]);
        //         // Similar to `isDelegated` logic
        //         if (currentOperator != address(0)) {
        //             _increaseOperatorScaledShares({
        //                 operator: currentOperator,
        //                 // the 'staker' here is the address receiving new shares
        //                 staker: msg.sender,
        //                 strategy: withdrawal.strategies[i],
        //                 scaledShares: scaledShares[i]
        //             });
        //         }
        //     }
        //     unchecked {
        //         ++i;
        //     }
        // }
    }

    // @notice Increases `operator`s delegated scaled shares in `strategy` by `scaledShares` and emits an `OperatorSharesIncreased` event
    function _increaseOperatorScaledShares(address operator, address staker, IStrategy strategy, uint256 scaledShares) internal {
        operatorScaledShares[operator][strategy] += scaledShares;
        // TODO: What to do about event wrt scaling?
        emit OperatorSharesIncreased(operator, staker, strategy, scaledShares);
    }

    // @notice Decreases `operator`s delegated scaled shares in `strategy` by `scaledShares` and emits an `OperatorSharesDecreased` event
    function _decreaseOperatorScaledShares(address operator, address staker, IStrategy strategy, uint256 scaledShares) internal {
        // This will revert on underflow, so no check needed
        operatorScaledShares[operator][strategy] -= scaledShares;
        // TODO: What to do about event wrt scaling?
        emit OperatorSharesDecreased(operator, staker, strategy, scaledShares);
    }

    /**
     * @notice Removes `scaledShares` in `strategies` from `staker` who is currently delegated to `operator` and queues a withdrawal to the `withdrawer`.
     * @dev If the `operator` is indeed an operator, then the operator's delegated shares in the `strategies` are also decreased appropriately.
     * @dev If `withdrawer` is not the same address as `staker`, then thirdPartyTransfersForbidden[strategy] must be set to false in the StrategyManager.
     */
    function _removeSharesAndQueueWithdrawal(
        address staker,
        address operator,
        address withdrawer,
        IStrategy[] memory strategies,
        uint256[] memory scaledShares
    ) internal returns (bytes32) {
        return bytes32(0);
        // require(
        //     staker != address(0), "DelegationManager._removeSharesAndQueueWithdrawal: staker cannot be zero address"
        // );
        // require(strategies.length != 0, "DelegationManager._removeSharesAndQueueWithdrawal: strategies cannot be empty");

        // // Remove shares from staker and operator
        // // Each of these operations fail if we attempt to remove more shares than exist
        // for (uint256 i = 0; i < strategies.length;) {
        //     // Similar to `isDelegated` logic
        //     if (operator != address(0)) {
        //         // forgefmt: disable-next-item
        //         _decreaseOperatorScaledShares({
        //             operator: operator, 
        //             staker: staker, 
        //             strategy: strategies[i], 
        //             scaledShares: scaledShares[i]
        //         });
        //     }

        //     // Remove active shares from EigenPodManager/StrategyManager
        //     if (strategies[i] == beaconChainETHStrategy) {
        //         /**
        //          * This call will revert if it would reduce the Staker's virtual beacon chain ETH shares below zero.
        //          * This behavior prevents a Staker from queuing a withdrawal which improperly removes excessive
        //          * shares from the operator to whom the staker is delegated.
        //          * It will also revert if the share amount being withdrawn is not a whole Gwei amount.
        //          */
        //         eigenPodManager.removeScaledShares(staker, scaledShares[i]);
        //     } else {
        //         require(
        //             staker == withdrawer || !strategyManager.thirdPartyTransfersForbidden(strategies[i]),
        //             "DelegationManager._removeSharesAndQueueWithdrawal: withdrawer must be same address as staker if thirdPartyTransfersForbidden are set"
        //         );
        //         // this call will revert if `scaledShares[i]` exceeds the Staker's current shares in `strategies[i]`
        //         strategyManager.removeScaledShares(staker, strategies[i], scaledShares[i]);
        //     }

        //     unchecked {
        //         ++i;
        //     }
        // }

        // // Create queue entry and increment withdrawal nonce
        // uint256 nonce = cumulativeWithdrawalsQueued[staker];
        // cumulativeWithdrawalsQueued[staker]++;

        // Withdrawal memory withdrawal = Withdrawal({
        //     staker: staker,
        //     delegatedTo: operator,
        //     withdrawer: withdrawer,
        //     nonce: nonce,
        //     startTimestamp: uint32(block.timestamp),
        //     strategies: strategies,
        //     scaledShares: scaledShares
        // });

        // bytes32 withdrawalRoot = calculateWithdrawalRoot(withdrawal);

        // // Place withdrawal in queue
        // pendingWithdrawals[withdrawalRoot] = true;

        // emit WithdrawalQueued(withdrawalRoot, withdrawal);
        // return withdrawalRoot;
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
        // TODO: rename params
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
        // require(
        //     _strategies.length == _withdrawalDelayBlocks.length,
        //     "DelegationManager._setStrategyWithdrawalDelayBlocks: input length mismatch"
        // );
        // uint256 numStrats = _strategies.length;
        // for (uint256 i = 0; i < numStrats; ++i) {
        //     IStrategy strategy = _strategies[i];
        //     uint256 prevStrategyWithdrawalDelayBlocks = strategyWithdrawalDelayBlocks[strategy];
        //     uint256 newStrategyWithdrawalDelayBlocks = _withdrawalDelayBlocks[i];
        //     require(
        //         newStrategyWithdrawalDelayBlocks <= MAX_WITHDRAWAL_DELAY_BLOCKS,
        //         "DelegationManager._setStrategyWithdrawalDelayBlocks: _withdrawalDelayBlocks cannot be > MAX_WITHDRAWAL_DELAY_BLOCKS"
        //     );

        //     // set the new withdrawal delay blocks
        //     strategyWithdrawalDelayBlocks[strategy] = newStrategyWithdrawalDelayBlocks;
        //     emit StrategyWithdrawalDelayBlocksSet(
        //         strategy, prevStrategyWithdrawalDelayBlocks, newStrategyWithdrawalDelayBlocks
        //     );
        // }
    }

    /**
     * @notice Sets the withdrawal delay (in seconds) for each strategy in `_strategies` to `_withdrawalDelay`.
     * gets called when initializing contract or by calling `setStrategyWithdrawalDelay`
     */
    function _setStrategyWithdrawalDelay(
        IStrategy[] calldata _strategies,
        uint256[] calldata _withdrawalDelays
    ) internal {
        // require(
        //     _strategies.length == _withdrawalDelays.length,
        //     "DelegationManager._setStrategyWithdrawalDelay: input length mismatch"
        // );
        // uint256 numStrats = _strategies.length;
        // for (uint256 i = 0; i < numStrats; ++i) {
        //     IStrategy strategy = _strategies[i];
        //     uint256 prevStrategyWithdrawalDelay = strategyWithdrawalDelays[strategy];
        //     uint256 newStrategyWithdrawalDelay = _withdrawalDelays[i];
        //     require(
        //         newStrategyWithdrawalDelay <= MAX_WITHDRAWAL_DELAY,
        //         "DelegationManager._setStrategyWithdrawalDelay: _withdrawalDelay cannot be > MAX_WITHDRAWAL_DELAY"
        //     );

        //     // set the new withdrawal delay (in seconds)
        //     strategyWithdrawalDelays[strategy] = newStrategyWithdrawalDelay;
        //     emit StrategyWithdrawalDelaySet(
        //         strategy, prevStrategyWithdrawalDelay, newStrategyWithdrawalDelay
        //     );
        // }
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
        return (delegatedTo[staker] != address(0));
    }

    /**
     * @notice Returns true is an operator has previously registered for delegation.
     */
    function isOperator(address operator) public view returns (bool) {
        return operator != address(0) && delegatedTo[operator] == operator;
    }

    /**
     * @notice Returns the OperatorDetails struct associated with an `operator`.
     */
    function operatorDetails(address operator) external view returns (OperatorDetails memory) {
        return _operatorDetails[operator];
    }

    /**
     * @notice Returns the AllocationDelayDetails struct associated with an `operator`
     * @dev If the operator has not set an allocation delay, then the `isSet` field will be `false`.
     */
    function operatorAllocationDelay(address operator) external view returns (AllocationDelayDetails memory) {
        return _operatorAllocationDelay[operator];
    }

    /**
     * @notice Returns the delegationApprover account for an operator
     */
    function delegationApprover(address operator) external view returns (address) {
        return _operatorDetails[operator].delegationApprover;
    }

    /**
     * @notice Returns the stakerOptOutWindowBlocks for an operator
     */
    function stakerOptOutWindowBlocks(address operator) external view returns (uint256) {
        return _operatorDetails[operator].stakerOptOutWindowBlocks;
    }

    /// @notice a legacy function that returns the operatorShares for an operator and strategy
    function operatorShares(address operator, IStrategy strategy) public view returns (uint256) {
        return ShareScalingLib.descaleShares({
            avsDirectory: avsDirectory,
            operator: operator,
            strategy: strategy,
            scaledShares: operatorScaledShares[operator][strategy]
        });
    }

    /// @notice Given array of strategies, returns array of scaled shares for the operator
    function getOperatorScaledShares(
        address operator,
        IStrategy[] memory strategies
    ) public view returns (uint256[] memory) {
        uint256[] memory scaledShares = new uint256[](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            scaledShares[i] = operatorScaledShares[operator][strategies[i]];
        }
        return scaledShares;
    }

    /// @notice Given array of strategies, returns array of shares for the operator
    function getOperatorShares(
        address operator,
        IStrategy[] memory strategies
    ) public view returns (uint256[] memory) {
        return ShareScalingLib.descaleShares({
            avsDirectory: avsDirectory,
            operator: operator,
            strategies: strategies,
            scaledShares: getOperatorScaledShares(operator, strategies)
        });
    }

    /**
     * @notice Given a staker and shares amounts of deposits, return the scaled shares calculated if
     * the staker were to deposit. Depends on which operator the staker is delegated to.
     */
    function getStakerScaledShares(
        address staker,
        IStrategy strategy,
        uint256 shares
    ) public view returns (uint256 scaledShares) {
        address operator = delegatedTo[staker];
        if (operator == address(0)) {
            // if the staker is not delegated to an operator, return the shares as is
            // as no slashing and scaling applied
            scaledShares = shares;
        } else {
            // if the staker is delegated to an operator, scale the shares accordingly
            scaledShares = ShareScalingLib.scaleShares({
                avsDirectory: avsDirectory,
                operator: operator,
                strategy: strategy,
                shares: shares
            });
        }
    }

    /**
     * @notice Given a staker and scaled shares amounts of deposits, return the shares calculated if
     * the staker were to withdraw. This value depends on which operator the staker is delegated to.
     * The shares amount returned is the actual amount of Strategy shares the staker would receive (subject
     * to each strategy's underlying shares to token ratio).
     */
    function getStakerShares(
        address staker,
        IStrategy strategy,
        uint256 scaledShares
    ) public view returns (uint256 shares) {
        address operator = delegatedTo[staker];
        if (operator == address(0)) {
            // if the staker is not delegated to an operator, return the shares as is
            // as no slashing and scaling applied
            shares = scaledShares;
        } else {
            // if the staker is delegated to an operator, descale the shares accordingly
            shares = ShareScalingLib.descaleShares({
                avsDirectory: avsDirectory,
                operator: operator,
                strategy: strategy,
                scaledShares: scaledShares
            });
        }
    }

    /**
     * @notice Returns the number of actively-delegatable scaled shares a staker has across all strategies.
     * @dev Returns two empty arrays in the case that the Staker has no actively-delegateable shares.
     */
    function getDelegatableScaledShares(address staker) public view returns (IStrategy[] memory, uint256[] memory) {
        // Get currently active scaled shares and strategies for `staker`
        int256 scaledPodShares = eigenPodManager.podOwnerScaledShares(staker);
        (IStrategy[] memory strategyManagerStrats, uint256[] memory strategyManagerShares) =
            strategyManager.getDeposits(staker);

        // Has no shares in EigenPodManager, but potentially some in StrategyManager
        if (scaledPodShares <= 0) {
            return (strategyManagerStrats, strategyManagerShares);
        }

        IStrategy[] memory strategies;
        uint256[] memory scaledShares;

        if (strategyManagerStrats.length == 0) {
            // Has shares in EigenPodManager, but not in StrategyManager
            strategies = new IStrategy[](1);
            scaledShares = new uint256[](1);
            strategies[0] = beaconChainETHStrategy;
            scaledShares[0] = uint256(scaledPodShares);
        } else {
            // Has shares in both
// TODO: make more efficient by resizing array
            // 1. Allocate return arrays
            strategies = new IStrategy[](strategyManagerStrats.length + 1);
            scaledShares = new uint256[](strategies.length);

            // 2. Place StrategyManager strats/shares in return arrays
            for (uint256 i = 0; i < strategyManagerStrats.length;) {
                strategies[i] = strategyManagerStrats[i];
                scaledShares[i] = strategyManagerShares[i];

                unchecked {
                    ++i;
                }
            }

            // 3. Place EigenPodManager strat/shares in return arrays
            strategies[strategies.length - 1] = beaconChainETHStrategy;
            scaledShares[strategies.length - 1] = uint256(scaledPodShares);
        }
        
        return (strategies, scaledShares);
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
