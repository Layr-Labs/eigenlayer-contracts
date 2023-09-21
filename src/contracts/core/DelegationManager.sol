// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
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
contract DelegationManager is Initializable, OwnableUpgradeable, Pausable, DelegationManagerStorage {

    /**
     * @dev Index for flag that pauses new delegations when set
     */
    uint8 internal constant PAUSED_NEW_DELEGATION = 0;

    /**
     * @dev Chain ID at the time of contract deployment
     */
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    /**
     * @dev Maximum Value for stakerOptOutWindowApproximately that is approximately equivalent to 6 months in blocks.
     */
    uint256 public constant MAX_STAKER_OPT_OUT_WINDOW_BLOCKS = (180 days) / 12;

    /// @notice Canonical, virtual beacon chain ETH strategy
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    /// @notice Simple permission for functions that are only callable by the StrategyManager contract.
    modifier onlyStrategyManager() {
        require(msg.sender == address(strategyManager), "onlyStrategyManager");
        _;
    }

    // @notice Simple permission for functions that are only callable by the StrategyManager contract OR by the EigenPodManagerContract
    modifier onlyStrategyManagerOrEigenPodManager() {
        require(msg.sender == address(strategyManager) || msg.sender == address(eigenPodManager),
            "DelegationManager: onlyStrategyManagerOrEigenPodManager");
        _;
    }

    /*******************************************************************************
                            INITIALIZING FUNCTIONS
    *******************************************************************************/

    /**
     * @dev Initializes the immutable addresses of the strategy mananger and slasher.
     */
    constructor(IStrategyManager _strategyManager, ISlasher _slasher, IEigenPodManager _eigenPodManager)
        DelegationManagerStorage(_strategyManager, _slasher, _eigenPodManager)
    {
        _disableInitializers();
        ORIGINAL_CHAIN_ID = block.chainid;
    }

    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, and paused status.
     */
    function initialize(address initialOwner, IPauserRegistry _pauserRegistry, uint256 initialPausedStatus)
        external
        initializer
    {
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _DOMAIN_SEPARATOR = _calculateDomainSeparator();
        _transferOwnership(initialOwner);
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
    function registerAsOperator(OperatorDetails calldata registeringOperatorDetails, string calldata metadataURI) external {
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
    function delegateTo(address operator, SignatureWithExpiry memory approverSignatureAndExpiry, bytes32 approverSalt) external {
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
        require(stakerSignatureAndExpiry.expiry >= block.timestamp, "DelegationManager.delegateToBySignature: staker signature expired");

        // calculate the digest hash, then increment `staker`'s nonce
        uint256 currentStakerNonce = stakerNonce[staker];
        bytes32 stakerDigestHash = calculateStakerDelegationDigestHash(staker, currentStakerNonce, operator, stakerSignatureAndExpiry.expiry);
        unchecked {
            stakerNonce[staker] = currentStakerNonce + 1;
        }

        // actually check that the signature is valid
        EIP1271SignatureUtils.checkSignature_EIP1271(staker, stakerDigestHash, stakerSignatureAndExpiry.signature);

        // go through the internal delegation flow, checking the `approverSignatureAndExpiry` if applicable
        _delegate(staker, operator, approverSignatureAndExpiry, approverSalt);
    }

    /**
     * @notice Undelegates the staker from the operator who they are delegated to. Puts the staker into the "undelegation limbo" mode of the EigenPodManager
     * and queues a withdrawal of all of the staker's shares in the StrategyManager (to the staker), if necessary.
     * @param staker The account to be undelegated.
     * @return withdrawalRoot The root of the newly queued withdrawal, if a withdrawal was queued. Otherwise just bytes32(0).
     *
     * @dev Reverts if the `staker` is also an operator, since operators are not allowed to undelegate from themselves.
     * @dev Reverts if the caller is not the staker, nor the operator who the staker is delegated to, nor the operator's specified "delegationApprover"
     * @dev Does nothing (but should not revert) if the staker is already undelegated.
     */
    function undelegate(address staker) external returns (bytes32 withdrawalRoot) {
        address operator = delegatedTo[staker];
        require(!isOperator(staker), "DelegationManager.undelegate: operators cannot be undelegated");
        require(staker != address(0), "DelegationManager.undelegate: cannot undelegate zero address");
        require(
            msg.sender == staker ||
            msg.sender == operator ||
            msg.sender == _operatorDetails[operator].delegationApprover,
            "DelegationManager.undelegate: caller cannot undelegate staker"
        );
        
        if (!stakerCanUndelegate(staker)) {
            // force the staker into "undelegation limbo" in the EigenPodManager if necessary
            uint256 podShares = eigenPodManager.forceIntoUndelegationLimbo(staker, operator);

            IStrategy[] memory strategies;
            uint256[] memory strategyShares;

            // force a withdrawal of all of the staker's shares from the StrategyManager
            (strategies, strategyShares, withdrawalRoot)
                = strategyManager.forceTotalWithdrawal(staker);

            _decreaseOperatorShares(operator, beaconChainETHStrategy, podShares);
            for (uint i = 0; i < strategies.length; ) {
                _decreaseOperatorShares(operator, strategies[i], strategyShares[i]);

                unchecked {
                    ++i;
                }
            }
        }

        // actually undelegate the staker -- only make storage changes + emit an event if the staker is actively delegated, otherwise do nothing
        if (isDelegated(staker)) {
            emit StakerUndelegated(staker, operator);
            delegatedTo[staker] = address(0);

            // emit an event if this action was not initiated by the staker themselves
            if (msg.sender != staker) {
                emit StakerForceUndelegated(staker, operator);
            }
        }

        return withdrawalRoot;
    }

    /**
     * @notice Increases a staker's delegated share balance in a strategy.
     * @param staker The address to increase the delegated shares for their operator.
     * @param strategy The strategy in which to increase the delegated shares.
     * @param shares The number of shares to increase.
     * 
     * @dev *If the staker is actively delegated*, then increases the `staker`'s delegated shares in `strategy` by `shares`. Otherwise does nothing.
     * @dev Callable only by the StrategyManager.
     */
    function increaseDelegatedShares(address staker, IStrategy strategy, uint256 shares)
        external
        onlyStrategyManagerOrEigenPodManager
    {
        // if the staker is delegated to an operator
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];

            // add strategy shares to delegate's shares
            operatorShares[operator][strategy] += shares;
        }
    }

    /**
     * @notice Decreases a staker's delegated share balance in a strategy.
     * @param staker The address to decrease the delegated shares for their operator.
     * @param strategies An array of strategies to crease the delegated shares.
     * @param shares An array of the number of shares to decrease for a operator and strategy.
     * 
     * @dev *If the staker is actively delegated*, then decreases the `staker`'s delegated shares in each entry of `strategies` by its respective `shares[i]`. Otherwise does nothing.
     * @dev Callable only by the StrategyManager or EigenPodManager.
     */
    function decreaseDelegatedShares(address staker, IStrategy[] calldata strategies, uint256[] calldata shares)
        external
        onlyStrategyManagerOrEigenPodManager
    {
        // if the staker is delegated to an operator
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];

            // subtract strategy shares from delegate's shares
            uint256 stratsLength = strategies.length;
            for (uint256 i = 0; i < stratsLength;) {
                _decreaseOperatorShares(operator, strategies[i], shares[i]);
                unchecked {
                    ++i;
                }
            }
        }
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
            "DelegationManager._setOperatorDetails: cannot set `earningsReceiver` to zero address");
        require(newOperatorDetails.stakerOptOutWindowBlocks <= MAX_STAKER_OPT_OUT_WINDOW_BLOCKS,
            "DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be > MAX_STAKER_OPT_OUT_WINDOW_BLOCKS");
        require(newOperatorDetails.stakerOptOutWindowBlocks >= _operatorDetails[operator].stakerOptOutWindowBlocks,
            "DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be decreased");
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
     *          3) the `operator` is not actively frozen
     *          4) if applicable, that the approver signature is valid and non-expired
     */ 
    function _delegate(
        address staker,
        address operator,
        SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt
    ) internal onlyWhenNotPaused(PAUSED_NEW_DELEGATION) {
        require(!isDelegated(staker), "DelegationManager._delegate: staker is already actively delegated");
        require(isOperator(operator), "DelegationManager._delegate: operator is not registered in EigenLayer");
        require(!slasher.isFrozen(operator), "DelegationManager._delegate: cannot delegate to a frozen operator");

        // fetch the operator's `delegationApprover` address and store it in memory in case we need to use it multiple times
        address _delegationApprover = _operatorDetails[operator].delegationApprover;
        /**
         * Check the `_delegationApprover`'s signature, if applicable.
         * If the `_delegationApprover` is the zero address, then the operator allows all stakers to delegate to them and this verification is skipped.
         * If the `_delegationApprover` or the `operator` themselves is the caller, then approval is assumed and signature verification is skipped as well.
         */
        if (_delegationApprover != address(0) && msg.sender != _delegationApprover && msg.sender != operator) {
            // check the signature expiry
            require(approverSignatureAndExpiry.expiry >= block.timestamp, "DelegationManager._delegate: approver signature expired");
            // check that the salt hasn't been used previously, then mark the salt as spent
            require(!delegationApproverSaltIsSpent[_delegationApprover][approverSalt], "DelegationManager._delegate: approverSalt already spent");
            delegationApproverSaltIsSpent[_delegationApprover][approverSalt] = true;

            // calculate the digest hash
            bytes32 approverDigestHash =
                calculateDelegationApprovalDigestHash(staker, operator, _delegationApprover, approverSalt, approverSignatureAndExpiry.expiry);

            // actually check that the signature is valid
            EIP1271SignatureUtils.checkSignature_EIP1271(_delegationApprover, approverDigestHash, approverSignatureAndExpiry.signature);
        }

        // retrieve any beacon chain ETH shares the staker might have
        uint256 beaconChainETHShares = eigenPodManager.podOwnerShares(staker);

        // increase the operator's shares in the canonical 'beaconChainETHStrategy' *if* the staker is not in "undelegation limbo"
        if (beaconChainETHShares != 0 && !eigenPodManager.isInUndelegationLimbo(staker)) {
            operatorShares[operator][beaconChainETHStrategy] += beaconChainETHShares;
        }

        // retrieve `staker`'s list of strategies and the staker's shares in each strategy from the StrategyManager
        (IStrategy[] memory strategies, uint256[] memory shares) = strategyManager.getDeposits(staker);

        // update the share amounts for each of the `operator`'s strategies
        for (uint256 i = 0; i < strategies.length;) {
            operatorShares[operator][strategies[i]] += shares[i];
            unchecked {
                ++i;
            }
        }

        // record the delegation relation between the staker and operator, and emit an event
        delegatedTo[staker] = operator;
        emit StakerDelegated(staker, operator);
    }

    function _decreaseOperatorShares(address operator, IStrategy strategy, uint shares) internal {
        // This will revert on underflow, so no check needed
        operatorShares[operator][strategy] -= shares;
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
        }
        else {
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
     * @notice Calculates the digestHash for a `staker` to sign to delegate to an `operator`
     * @param staker The signing staker
     * @param operator The operator who is being delegated to
     * @param expiry The desired expiry time of the staker's signature
     */
    function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) external view returns (bytes32) {
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
    function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) public view returns (bytes32) {
        // calculate the struct hash
        bytes32 stakerStructHash = keccak256(abi.encode(STAKER_DELEGATION_TYPEHASH, staker, operator, _stakerNonce, expiry));
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
        bytes32 approverStructHash = keccak256(abi.encode(DELEGATION_APPROVAL_TYPEHASH, _delegationApprover, staker, operator, approverSalt, expiry));
        // calculate the digest hash
        bytes32 approverDigestHash = keccak256(abi.encodePacked("\x19\x01", domainSeparator(), approverStructHash));
        return approverDigestHash;
    }

    /** 
     * @notice Returns 'true' if the staker can immediately undelegate without queuing a new withdrawal OR if the staker is already undelegated, and 'false' otherwise
     * @dev A staker can only undelegate if they have no "active" shares in EigenLayer and are not themselves an operator
     */
    function stakerCanUndelegate(address staker) public view returns (bool) {
        return (!isOperator(staker) &&
            strategyManager.stakerStrategyListLength(staker) == 0 &&
            eigenPodManager.podOwnerHasNoDelegatedShares(staker));
    }

    /** 
      * @dev Recalculates the domain separator when the chainid changes due to a fork.
      */
    function _calculateDomainSeparator() internal view returns (bytes32) {
        return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("EigenLayer")), block.chainid, address(this)));
    }
}
