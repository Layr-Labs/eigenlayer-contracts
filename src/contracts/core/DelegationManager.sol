// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "../interfaces/ISlasher.sol";
import "./DelegationManagerStorage.sol";
import "../permissions/Pausable.sol";
import "../libraries/EIP1271SignatureUtils.sol";

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
contract DelegationManager is Initializable, OwnableUpgradeable, Pausable, DelegationManagerStorage, ReentrancyGuardUpgradeable {
    // @dev Index for flag that pauses new delegations when set
    uint8 internal constant PAUSED_NEW_DELEGATION = 0;

    // @dev Index for flag that pauses queuing new withdrawals when set.
    uint8 internal constant PAUSED_ENTER_WITHDRAWAL_QUEUE = 1;

    // @dev Index for flag that pauses completing existing withdrawals when set.
    uint8 internal constant PAUSED_EXIT_WITHDRAWAL_QUEUE = 2;

    // @dev Index for flag that pauses operator register/deregister to avs when set.
    uint8 internal constant PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS = 3;

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

    /*******************************************************************************
                            INITIALIZING FUNCTIONS
    *******************************************************************************/

    /**
     * @dev Initializes the immutable addresses of the strategy mananger and slasher.
     */
    constructor(
        IStrategyManager _strategyManager,
        ISlasher _slasher,
        IEigenPodManager _eigenPodManager
    ) DelegationManagerStorage(_strategyManager, _slasher, _eigenPodManager) {
        _disableInitializers();
        ORIGINAL_CHAIN_ID = block.chainid;
    }

    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, and paused status.
     * withdrawalDelayBlocks is set only once here
     */
    function initialize(
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 initialPausedStatus,
        uint256 _withdrawalDelayBlocks
    ) external initializer {
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _DOMAIN_SEPARATOR = _calculateDomainSeparator();
        _transferOwnership(initialOwner);
        _initializeWithdrawalDelayBlocks(_withdrawalDelayBlocks);
    }

    /*******************************************************************************
                            EXTERNAL FUNCTIONS 
    *******************************************************************************/

    /**
     * @notice Registers the caller as an operator in EigenLayer.
     * @param registeringOperatorDetails is the `OperatorDetails` for the operator.
     * @param metadataURI is a URI for the operator's metadata, i.e. a link providing more details on the operator.
     *
     * @dev Once an operator is registered, they cannot 'deregister' as an operator, and they will forever be considered "delegated to themself".
     * @dev This function will revert if the caller attempts to set their `earningsReceiver` to address(0).
     * @dev Note that the `metadataURI` is *never stored * and is only emitted in the `OperatorMetadataURIUpdated` event
     */
    function registerAsOperator(
        OperatorDetails calldata registeringOperatorDetails,
        string calldata metadataURI
    ) external {
        require(
            _operatorDetails[msg.sender].earningsReceiver == address(0),
            "DelegationManager.registerAsOperator: operator has already registered"
        );
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
     * @dev This function will revert if the caller attempts to set their `earningsReceiver` to address(0).
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
     * @notice Called by an avs to emit an `AVSMetadataURIUpdated` event indicating the information has updated.
     * @param metadataURI The URI for metadata associated with an avs
     */
    function updateAVSMetadataURI(string calldata metadataURI) external {
        emit AVSMetadataURIUpdated(msg.sender, metadataURI);
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

        // calculate the digest hash, then increment `staker`'s nonce
        uint256 currentStakerNonce = stakerNonce[staker];
        bytes32 stakerDigestHash = calculateStakerDelegationDigestHash(
            staker,
            currentStakerNonce,
            operator,
            stakerSignatureAndExpiry.expiry
        );
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
    function undelegate(address staker) external onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE) returns (bytes32) {
        require(isDelegated(staker), "DelegationManager.undelegate: staker must be delegated to undelegate");
        require(!isOperator(staker), "DelegationManager.undelegate: operators cannot be undelegated");
        require(staker != address(0), "DelegationManager.undelegate: cannot undelegate zero address");
        address operator = delegatedTo[staker];
        require(
            msg.sender == staker ||
                msg.sender == operator ||
                msg.sender == _operatorDetails[operator].delegationApprover,
            "DelegationManager.undelegate: caller cannot undelegate staker"
        );

        // Gather strategies and shares to remove from staker/operator during undelegation
        // Undelegation removes ALL currently-active strategies and shares
        (IStrategy[] memory strategies, uint256[] memory shares)
            = getDelegatableShares(staker);

        // emit an event if this action was not initiated by the staker themselves
        if (msg.sender != staker) {
            emit StakerForceUndelegated(staker, operator);
        }

        // undelegate the staker
        emit StakerUndelegated(staker, operator);
        delegatedTo[staker] = address(0);

        // if no delegatable shares, return zero root, and don't queue a withdrawal
        if (strategies.length == 0) {
            return bytes32(0);
        } else {
            // Remove all strategies/shares from staker and operator and place into queue
            return _removeSharesAndQueueWithdrawal({
                staker: staker,
                operator: operator,
                withdrawer: staker,
                strategies: strategies,
                shares: shares
            });
        }
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

        for (uint256 i = 0; i < queuedWithdrawalParams.length; i++) {
            require(queuedWithdrawalParams[i].strategies.length == queuedWithdrawalParams[i].shares.length, "DelegationManager.queueWithdrawal: input length mismatch");
            require(queuedWithdrawalParams[i].withdrawer != address(0), "DelegationManager.queueWithdrawal: must provide valid withdrawal address");

            address operator = delegatedTo[msg.sender];

            // Remove shares from staker's strategies and place strategies/shares in queue.
            // If the staker is delegated to an operator, the operator's delegated shares are also reduced
            // NOTE: This will fail if the staker doesn't have the shares implied by the input parameters
            withdrawalRoots[i] = _removeSharesAndQueueWithdrawal({
                staker: msg.sender,
                operator: operator,
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

    /// @notice Migrates an array of queued withdrawals from the StrategyManager contract to this contract.
    /// @dev This function is expected to be removed in the next upgrade, after all queued withdrawals have been migrated.
    function migrateQueuedWithdrawals(IStrategyManager.DeprecatedStruct_QueuedWithdrawal[] memory withdrawalsToMigrate) external {
        for(uint256 i = 0; i < withdrawalsToMigrate.length;) {
            IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory withdrawalToMigrate = withdrawalsToMigrate[i];
            // Delete withdrawal root from strateyManager
            (bool isDeleted, bytes32 oldWithdrawalRoot) = strategyManager.migrateQueuedWithdrawal(withdrawalToMigrate);
            // If old storage is deleted from strategyManager
            if (isDeleted) {
                address staker = withdrawalToMigrate.staker;
                // Create queue entry and increment withdrawal nonce
                uint256 nonce = cumulativeWithdrawalsQueued[staker];
                cumulativeWithdrawalsQueued[staker]++;

                Withdrawal memory migratedWithdrawal = Withdrawal({
                    staker: staker,
                    delegatedTo: withdrawalToMigrate.delegatedAddress,
                    withdrawer: withdrawalToMigrate.withdrawerAndNonce.withdrawer,
                    nonce: nonce,
                    startBlock: withdrawalToMigrate.withdrawalStartBlock,
                    strategies: withdrawalToMigrate.strategies,
                    shares: withdrawalToMigrate.shares
                });

                // create the new storage
                bytes32 newRoot = calculateWithdrawalRoot(migratedWithdrawal);
                // safety check to ensure that root doesn't exist already -- this should *never* be hit
                require(!pendingWithdrawals[newRoot], "DelegationManager.migrateQueuedWithdrawals: withdrawal already exists");
                pendingWithdrawals[newRoot] = true;

                emit WithdrawalQueued(newRoot, migratedWithdrawal);

                emit WithdrawalMigrated(oldWithdrawalRoot, newRoot);
            }
            unchecked {
                ++i;
            }
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
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];

            // add strategy shares to delegate's shares
            _increaseOperatorShares({operator: operator, staker: staker, strategy: strategy, shares: shares});
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
        // if the staker is delegated to an operator
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];

            // subtract strategy shares from delegate's shares
            _decreaseOperatorShares({
                operator: operator,
                staker: staker,
                strategy: strategy,
                shares: shares
            });
        }
    }

    /**
     * @notice Called by an avs to register an operator with the avs.
     * @param operator The address of the operator to register.
     * @param operatorSignature The signature, salt, and expiry of the operator's signature.
     */
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {

        require(
            operatorSignature.expiry >= block.timestamp,
            "DelegationManager.registerOperatorToAVS: operator signature expired"
        );
        require(
            avsOperatorStatus[msg.sender][operator] != OperatorAVSRegistrationStatus.REGISTERED,
            "DelegationManager.registerOperatorToAVS: operator already registered"
        );
        require(
            !operatorSaltIsSpent[operator][operatorSignature.salt],
            "DelegationManager.registerOperatorToAVS: salt already spent"
        );
        require(
            isOperator(operator),
            "DelegationManager.registerOperatorToAVS: operator not registered to EigenLayer yet");

        // Calculate the digest hash
        bytes32 operatorRegistrationDigestHash = calculateOperatorAVSRegistrationDigestHash({
            operator: operator,
            avs: msg.sender,
            salt: operatorSignature.salt,
            expiry: operatorSignature.expiry
        });

        // Check that the signature is valid
        EIP1271SignatureUtils.checkSignature_EIP1271(
            operator,
            operatorRegistrationDigestHash,
            operatorSignature.signature
        );

        // Set the operator as registered
        avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.REGISTERED;

        // Mark the salt as spent
        operatorSaltIsSpent[operator][operatorSignature.salt] = true;

        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.REGISTERED);
    }

    /**
     * @notice Called by an avs to deregister an operator with the avs.
     * @param operator The address of the operator to deregister.
     */
    function deregisterOperatorFromAVS(address operator) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        require(
            avsOperatorStatus[msg.sender][operator] == OperatorAVSRegistrationStatus.REGISTERED,
            "DelegationManager.deregisterOperatorFromAVS: operator not registered"
        );

        // Set the operator as deregistered
        avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.UNREGISTERED;

        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.UNREGISTERED);
    }

    /*******************************************************************************
                            INTERNAL FUNCTIONS
    *******************************************************************************/

    /**
     * @notice Sets operator parameters in the `_operatorDetails` mapping.
     * @param operator The account registered as an operator updating their operatorDetails
     * @param newOperatorDetails The new parameters for the operator
     *
     * @dev This function will revert if the operator attempts to set their `earningsReceiver` to address(0).
     */
    function _setOperatorDetails(address operator, OperatorDetails calldata newOperatorDetails) internal {
        require(
            newOperatorDetails.earningsReceiver != address(0),
            "DelegationManager._setOperatorDetails: cannot set `earningsReceiver` to zero address"
        );
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
     * @notice Delegates *from* a `staker` *to* an `operator`.
     * @param staker The address to delegate *from* -- this address is delegating control of its own assets.
     * @param operator The address to delegate *to* -- this address is being given power to place the `staker`'s assets at risk on services
     * @param approverSignatureAndExpiry Verifies the operator approves of this delegation
     * @param approverSalt Is a salt used to help guarantee signature uniqueness. Each salt can only be used once by a given approver.
     * @dev Ensures that:
     *          1) the `staker` is not already delegated to an operator
     *          2) the `operator` has indeed registered as an operator in EigenLayer
     *          3) if applicable, that the approver signature is valid and non-expired
     */
    function _delegate(
        address staker,
        address operator,
        SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt
    ) internal onlyWhenNotPaused(PAUSED_NEW_DELEGATION) {
        require(!isDelegated(staker), "DelegationManager._delegate: staker is already actively delegated");
        require(isOperator(operator), "DelegationManager._delegate: operator is not registered in EigenLayer");

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

            // calculate the digest hash
            bytes32 approverDigestHash = calculateDelegationApprovalDigestHash(
                staker,
                operator,
                _delegationApprover,
                approverSalt,
                approverSignatureAndExpiry.expiry
            );

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

        (IStrategy[] memory strategies, uint256[] memory shares)
            = getDelegatableShares(staker);

        for (uint256 i = 0; i < strategies.length;) {
            _increaseOperatorShares({
                operator: operator,
                staker: staker,
                strategy: strategies[i],
                shares: shares[i]
            });

            unchecked { ++i; }
        }
    }

    /**
     * @dev commented-out param (middlewareTimesIndex) is the index in the operator that the staker who triggered the withdrawal was delegated to's middleware times array
     * This param is intended to be passed on to the Slasher contract, but is unused in the M2 release of these contracts, and is thus commented-out.
     */
    function _completeQueuedWithdrawal(
        Withdrawal calldata withdrawal,
        IERC20[] calldata tokens,
        uint256 /*middlewareTimesIndex*/,
        bool receiveAsTokens
    ) internal {
        bytes32 withdrawalRoot = calculateWithdrawalRoot(withdrawal);

        require(
            pendingWithdrawals[withdrawalRoot], 
            "DelegationManager.completeQueuedAction: action is not in queue"
        );

        require(
            withdrawal.startBlock + withdrawalDelayBlocks <= block.number, 
            "DelegationManager.completeQueuedAction: withdrawalDelayBlocks period has not yet passed"
        );

        require(
            msg.sender == withdrawal.withdrawer, 
            "DelegationManager.completeQueuedAction: only withdrawer can complete action"
        );

        if (receiveAsTokens) {
            require(
                tokens.length == withdrawal.strategies.length, 
                "DelegationManager.completeQueuedAction: input length mismatch"
            );
        }

        // Remove `withdrawalRoot` from pending roots
        delete pendingWithdrawals[withdrawalRoot];

        // Finalize action by converting shares to tokens for each strategy, or
        // by re-awarding shares in each strategy.
        if (receiveAsTokens) {
            for (uint256 i = 0; i < withdrawal.strategies.length; ) {
                _withdrawSharesAsTokens({
                    staker: withdrawal.staker,
                    withdrawer: msg.sender,
                    strategy: withdrawal.strategies[i],
                    shares: withdrawal.shares[i],
                    token: tokens[i]
                });
                unchecked { ++i; }
            }
        // Award shares back in StrategyManager/EigenPodManager. If withdrawer is delegated, increase the shares delegated to the operator
        } else {
            address currentOperator = delegatedTo[msg.sender];
            for (uint256 i = 0; i < withdrawal.strategies.length; ) {
                /** When awarding podOwnerShares in EigenPodManager, we need to be sure to only give them back to the original podOwner.
                 * Other strategy shares can + will be awarded to the withdrawer.
                 */
                if (withdrawal.strategies[i] == beaconChainETHStrategy) {
                    address staker = withdrawal.staker;
                    /**
                    * Update shares amount depending upon the returned value.
                    * The return value will be lower than the input value in the case where the staker has an existing share deficit
                    */
                    uint256 increaseInDelegateableShares = eigenPodManager.addShares({
                        podOwner: staker,
                        shares: withdrawal.shares[i]
                    });
                    address podOwnerOperator = delegatedTo[staker];
                    // Similar to `isDelegated` logic
                    if (podOwnerOperator != address(0)) {
                        _increaseOperatorShares({
                            operator: podOwnerOperator,
                            // the 'staker' here is the address receiving new shares
                            staker: staker,
                            strategy: withdrawal.strategies[i],
                            shares: increaseInDelegateableShares
                        });
                    }
                } else {
                    strategyManager.addShares(msg.sender, withdrawal.strategies[i], withdrawal.shares[i]);
                    // Similar to `isDelegated` logic
                    if (currentOperator != address(0)) {
                        _increaseOperatorShares({
                            operator: currentOperator,
                            // the 'staker' here is the address receiving new shares
                            staker: msg.sender,
                            strategy: withdrawal.strategies[i],
                            shares: withdrawal.shares[i]
                        });
                    }
                }
                unchecked { ++i; }
            }
        }

        emit WithdrawalCompleted(withdrawalRoot);
    }

    // @notice Increases `operator`s delegated shares in `strategy` by `shares` and emits an `OperatorSharesIncreased` event
    function _increaseOperatorShares(address operator, address staker, IStrategy strategy, uint256 shares) internal {
        operatorShares[operator][strategy] += shares;
        emit OperatorSharesIncreased(operator, staker, strategy, shares);
    }

    // @notice Decreases `operator`s delegated shares in `strategy` by `shares` and emits an `OperatorSharesDecreased` event
    function _decreaseOperatorShares(address operator, address staker, IStrategy strategy, uint256 shares) internal {
        // This will revert on underflow, so no check needed
        operatorShares[operator][strategy] -= shares;
        emit OperatorSharesDecreased(operator, staker, strategy, shares);
    }

    /**
     * @notice Removes `shares` in `strategies` from `staker` who is currently delegated to `operator` and queues a withdrawal to the `withdrawer`.
     * @dev If the `operator` is indeed an operator, then the operator's delegated shares in the `strategies` are also decreased appropriately.
     */
    function _removeSharesAndQueueWithdrawal(
        address staker, 
        address operator,
        address withdrawer,
        IStrategy[] memory strategies, 
        uint256[] memory shares
    ) internal returns (bytes32) {
        require(staker != address(0), "DelegationManager._removeSharesAndQueueWithdrawal: staker cannot be zero address");
        require(strategies.length != 0, "DelegationManager._removeSharesAndQueueWithdrawal: strategies cannot be empty");
    
        // Remove shares from staker and operator
        // Each of these operations fail if we attempt to remove more shares than exist
        for (uint256 i = 0; i < strategies.length;) {
            // Similar to `isDelegated` logic
            if (operator != address(0)) {
                _decreaseOperatorShares({
                    operator: operator,
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
                // this call will revert if `shares[i]` exceeds the Staker's current shares in `strategies[i]`
                strategyManager.removeShares(staker, strategies[i], shares[i]);
            }

            unchecked { ++i; }
        }

        // Create queue entry and increment withdrawal nonce
        uint256 nonce = cumulativeWithdrawalsQueued[staker];
        cumulativeWithdrawalsQueued[staker]++;

        Withdrawal memory withdrawal = Withdrawal({
            staker: staker,
            delegatedTo: operator,
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
    function _withdrawSharesAsTokens(address staker, address withdrawer, IStrategy strategy, uint256 shares, IERC20 token) internal {
        if (strategy == beaconChainETHStrategy) {
            eigenPodManager.withdrawSharesAsTokens({
                podOwner: staker,
                destination: withdrawer,
                shares: shares
            });
        } else {
            strategyManager.withdrawSharesAsTokens(withdrawer, strategy, shares, token);
        }
    }

    function _initializeWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) internal {
        require(
            _withdrawalDelayBlocks <= MAX_WITHDRAWAL_DELAY_BLOCKS,
            "DelegationManager._initializeWithdrawalDelayBlocks: _withdrawalDelayBlocks cannot be > MAX_WITHDRAWAL_DELAY_BLOCKS"
        );
        emit WithdrawalDelayBlocksSet(withdrawalDelayBlocks, _withdrawalDelayBlocks);
        withdrawalDelayBlocks = _withdrawalDelayBlocks;
    }

    /*******************************************************************************
                            VIEW FUNCTIONS
    *******************************************************************************/

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
        return (_operatorDetails[operator].earningsReceiver != address(0));
    }

    /**
     * @notice Returns the OperatorDetails struct associated with an `operator`.
     */
    function operatorDetails(address operator) external view returns (OperatorDetails memory) {
        return _operatorDetails[operator];
    }

    /*
     * @notice Returns the earnings receiver address for an operator
     */
    function earningsReceiver(address operator) external view returns (address) {
        return _operatorDetails[operator].earningsReceiver;
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

    /**
     * @notice Returns the number of actively-delegatable shares a staker has across all strategies.
     * @dev Returns two empty arrays in the case that the Staker has no actively-delegateable shares.
     */
    function getDelegatableShares(address staker) public view returns (IStrategy[] memory, uint256[] memory) {
        // Get currently active shares and strategies for `staker`
        int256 podShares = eigenPodManager.podOwnerShares(staker);
        (IStrategy[] memory strategyManagerStrats, uint256[] memory strategyManagerShares) 
            = strategyManager.getDeposits(staker);

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
            for (uint256 i = 0; i < strategyManagerStrats.length; ) {
                strategies[i] = strategyManagerStrats[i];
                shares[i] = strategyManagerShares[i];

                unchecked { ++i; }
            }

            // 3. Place EigenPodManager strat/shares in return arrays
            strategies[strategies.length - 1] = beaconChainETHStrategy;
            shares[strategies.length - 1] = uint256(podShares);
        }

        return (strategies, shares);
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
        bytes32 stakerStructHash = keccak256(
            abi.encode(STAKER_DELEGATION_TYPEHASH, staker, operator, _stakerNonce, expiry)
        );
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
     * @notice Calculates the digest hash to be signed by an operator to register with an AVS
     * @param operator The account registering as an operator
     * @param avs The AVS the operator is registering to
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid
     */
    function calculateOperatorAVSRegistrationDigestHash(
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        // calculate the struct hash
        bytes32 structHash = keccak256(
            abi.encode(OPERATOR_AVS_REGISTRATION_TYPEHASH, operator, avs, salt, expiry)
        );
        // calculate the digest hash
        bytes32 digestHash = keccak256(
            abi.encodePacked("\x19\x01", domainSeparator(), structHash)
        );
        return digestHash;
    }

    /**
     * @dev Recalculates the domain separator when the chainid changes due to a fork.
     */
    function _calculateDomainSeparator() internal view returns (bytes32) {
        return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("EigenLayer")), block.chainid, address(this)));
    }
}
