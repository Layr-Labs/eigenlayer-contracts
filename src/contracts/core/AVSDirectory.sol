// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "../permissions/Pausable.sol";
import "../libraries/EIP1271SignatureUtils.sol";
import "./AVSDirectoryStorage.sol";

contract AVSDirectory is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AVSDirectoryStorage,
    ReentrancyGuardUpgradeable
{
    /// @dev Index for flag that pauses operator register/deregister to avs when set.
    uint8 internal constant PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS = 0;

    /// @dev Chain ID at the time of contract deployment
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    /// @notice Canonical, virtual beacon chain ETH strategy
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the immutable addresses of the strategy mananger, delegationManager, slasher,
     * and eigenpodManager contracts
     */
    constructor(
        IDelegationManager _delegation,
        IStrategyManager _strategyManager
    ) AVSDirectoryStorage(_delegation, _strategyManager) {
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
     *  @notice Called by AVSs to add an operator to an operator set.
     *
     *  @param operator The address of the operator to be added to the operator set.
     *  @param operatorSetIds The IDs of the operator sets.
     *  @param signature The signature of the operator on their intent to register.
     *
     *  @dev msg.sender is used as the AVS.
     *  @dev The operator must not have a pending deregistration from the operator set.
     *  @dev If this is the first operator set in the AVS that the operator is
     *  registering for, a OperatorAVSRegistrationStatusUpdated event is emitted with
     *  a REGISTERED status.
     */
    function registerOperatorToOperatorSets(
        address operator,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        // Assert `operator` is actually an operator.
        require(
            delegation.isOperator(operator),
            "AVSDirectory.registerOperatorToOperatorSets: operator not registered to EigenLayer yet"
        );
        // Assert operator's signature has not expired.
        require(
            operatorSignature.expiry >= block.timestamp,
            "AVSDirectory.registerOperatorToOperatorSets: operator signature expired"
        );
        // Assert operator's signature `salt` has not already been spent.
        require(
            !operatorSaltIsSpent[operator][operatorSignature.salt],
            "AVSDirectory.registerOperatorToOperatorSets: salt already spent"
        );

        // Assert signature provided by `operator` is valid.
        EIP1271SignatureUtils.checkSignature_EIP1271(
            operator,
            calculateOperatorSetRegistrationDigestHash({
                avs: msg.sender,
                operatorSetIds: operatorSetIds,
                salt: operatorSignature.salt,
                expiry: operatorSignature.expiry
            }),
            operatorSignature.signature
        );

        // Mutate `operatorSaltIsSpent` to `true` to prevent future respending.
        operatorSaltIsSpent[operator][operatorSignature.salt] = true;

        // Register `operator` if not already registered.
        if (avsOperatorStatus[msg.sender][operator] != OperatorAVSRegistrationStatus.REGISTERED) {
            avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.REGISTERED;
            emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.REGISTERED);
        }

        // Loop over `operatorSetIds` array and register `operator` for each item.
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            // Assert `operator` has not already been registered to `operatorSetIds[i]`.
            require(
                !isOperatorInOperatorSet[msg.sender][operator][operatorSetIds[i]],
                "AVSDirectory.registerOperatorToOperatorSets: operator already registered to operator set"
            );

            // Mutate calling AVS to operator set AVS status, preventing further legacy registrations.
            if (!isOperatorSetAVS[msg.sender]) isOperatorSetAVS[msg.sender] = true;

            // Mutate `isOperatorInOperatorSet` to `true` for `operatorSetIds[i]`.
            isOperatorInOperatorSet[msg.sender][operator][operatorSetIds[i]] = true;

            emit OperatorAddedToOperatorSet(operator, OperatorSet({avs: msg.sender, id: operatorSetIds[i]}));
        }

        // Increase `operatorAVSOperatorSetCount` by `operatorSetIds.length`.
        // You would have to add the operator to 2**256-2 operator sets before overflow is possible here.
        unchecked {
            operatorAVSOperatorSetCount[msg.sender][operator] += operatorSetIds.length;
        }
    }

    /**
     *  @notice Called by AVSs or operators to remove an operator from an operator set.
     *
     *  @param operator The address of the operator to be removed from the operator set.
     *  @param operatorSetIds The IDs of the operator sets.
     *
     *  @dev msg.sender is used as the AVS.
     *  @dev The operator must be registered for the msg.sender AVS and the given operator set.
     *  @dev If this call removes the operator from all operator sets for the msg.sender AVS,
     *  then an OperatorAVSRegistrationStatusUpdated event is emitted with a DEREGISTERED status.
     */
    function deregisterOperatorFromOperatorSets(
        address operator,
        uint32[] calldata operatorSetIds
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        // Loop over `operatorSetIds` array and deregister `operator` for each item.
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            // Assert `operator` is registered for this iterations operator set.
            require(
                isOperatorInOperatorSet[msg.sender][operator][operatorSetIds[i]],
                "AVSDirectory.deregisterOperatorFromOperatorSet: operator not registered for operator set"
            );

            // Mutate `isOperatorInOperatorSet` to `false` for `operatorSetIds[i]`.
            isOperatorInOperatorSet[msg.sender][operator][operatorSetIds[i]] = false;

            // Decrement `operatorAVSOperatorSetCount` by 1.
            // The above assertion makes underflow logically impossible here.
            unchecked {
                --operatorAVSOperatorSetCount[msg.sender][operator];
            }

            emit OperatorRemovedFromOperatorSet(operator, OperatorSet({avs: msg.sender, id: operatorSetIds[i]}));
        }

        // Set the operator as deregistered if no longer registered for any operator sets
        if (operatorAVSOperatorSetCount[msg.sender][operator] == 0) {
            avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.UNREGISTERED;
            emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.UNREGISTERED);
        }
    }

    /**
     *  @notice Called by the AVS's service manager contract to register an operator with the AVS.
     *
     *  @param operator The address of the operator to register.
     *  @param operatorSignature The signature, salt, and expiry of the operator's signature.
     *
     *  @dev msg.sender must be the AVS.
     *  @dev Only used by legacy M2 AVSs that have not integrated with operator sets.
     */
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        require(
            operatorSignature.expiry >= block.timestamp,
            "AVSDirectory.registerOperatorToAVS: operator signature expired"
        );
        require(
            avsOperatorStatus[msg.sender][operator] != OperatorAVSRegistrationStatus.REGISTERED,
            "AVSDirectory.registerOperatorToAVS: operator already registered"
        );
        require(
            !operatorSaltIsSpent[operator][operatorSignature.salt],
            "AVSDirectory.registerOperatorToAVS: salt already spent"
        );
        require(
            delegation.isOperator(operator),
            "AVSDirectory.registerOperatorToAVS: operator not registered to EigenLayer yet"
        );
        require(
            !isOperatorSetAVS[msg.sender],
            "AVSDirectory.registerOperatorToAVS: operator set AVS cannot register operators with legacy method"
        );

        // Calculate the digest hash
        bytes32 operatorRegistrationDigestHash = calculateOperatorAVSRegistrationDigestHash({
            operator: operator,
            avs: msg.sender,
            salt: operatorSignature.salt,
            expiry: operatorSignature.expiry
        });

        // Check that the signature is valid
        EIP1271SignatureUtils.checkSignature_EIP1271(
            operator, operatorRegistrationDigestHash, operatorSignature.signature
        );

        // Set the operator as registered
        avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.REGISTERED;

        // Mark the salt as spent
        operatorSaltIsSpent[operator][operatorSignature.salt] = true;

        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.REGISTERED);
    }

    /**
     *  @notice Called by an AVS to deregister an operator from the AVS.
     *
     *  @param operator The address of the operator to deregister.
     *
     *  @dev Only used by legacy M2 AVSs that have not integrated with operator sets.
     */
    function deregisterOperatorFromAVS(address operator)
        external
        onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS)
    {
        require(
            avsOperatorStatus[msg.sender][operator] == OperatorAVSRegistrationStatus.REGISTERED,
            "AVSDirectory.deregisterOperatorFromAVS: operator not registered"
        );

        // Set the operator as deregistered
        avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.UNREGISTERED;

        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.UNREGISTERED);
    }

    /**
     *  @notice Called by an AVS to emit an `AVSMetadataURIUpdated` event indicating the information has updated.
     *
     *  @param metadataURI The URI for metadata associated with an AVS.
     *
     *  @dev Note that the `metadataURI` is *never stored* and is only emitted in the `AVSMetadataURIUpdated` event.
     */
    function updateAVSMetadataURI(string calldata metadataURI) external {
        emit AVSMetadataURIUpdated(msg.sender, metadataURI);
    }

    /**
     *  @notice Called by AVSs to add strategies to its operator set.
     *
     *  @param operatorSetID The ID of the operator set.
     *  @param strategies The list of strategies to add to the operator set.
     *
     *  @dev msg.sender is used as the AVS.
     *  @dev No storage is updated, as the event is used by off-chain services.
     */
    function addStrategiesToOperatorSet(uint32 operatorSetID, IStrategy[] calldata strategies) external {
        for (uint256 i = 0; i < strategies.length; ++i) {
            // Require that the strategy is valid
            require(
                strategyManager.strategyIsWhitelistedForDeposit(strategies[i])
                    || strategies[i] == beaconChainETHStrategy,
                "AVSDirectory.addStrategiesToOperatorSet: invalid strategy considered"
            );
            require(
                !isOperatorSetStrategy[msg.sender][operatorSetID][strategies[i]],
                "AVSDirectory.addStrategiesToOperatorSet: strategy already added to operator set"
            );
            isOperatorSetStrategy[msg.sender][operatorSetID][strategies[i]] = true;
            emit OperatorSetStrategyAdded(OperatorSet({avs: msg.sender, id: operatorSetID}), strategies[i]);
        }
    }

    /**
     *  @notice Called by AVSs to remove strategies from its operator set.
     *
     *  @param operatorSetID The ID of the operator set.
     *  @param strategies The list of strategies to remove from the operator set.
     *
     *  @dev msg.sender is used as the AVS.
     *  @dev No storage is updated, as the event is used by off-chain services.
     */
    function removeStrategiesFromOperatorSet(uint32 operatorSetID, IStrategy[] calldata strategies) external {
        for (uint256 i = 0; i < strategies.length; ++i) {
            require(
                isOperatorSetStrategy[msg.sender][operatorSetID][strategies[i]],
                "AVSDirectory.removeStrategiesFromOperatorSet: strategy not a member of operator set"
            );
            isOperatorSetStrategy[msg.sender][operatorSetID][strategies[i]] = false;
            emit OperatorSetStrategyRemoved(OperatorSet({avs: msg.sender, id: operatorSetID}), strategies[i]);
        }
    }

    /**
     * @notice Called by an operator to cancel a salt that has been used to register with an AVS.
     *
     * @param salt A unique and single use value associated with the approver signature.
     */
    function cancelSalt(bytes32 salt) external {
        require(!operatorSaltIsSpent[msg.sender][salt], "AVSDirectory.cancelSalt: cannot cancel spent salt");
        operatorSaltIsSpent[msg.sender][salt] = true;
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     *  @notice Calculates the digest hash to be signed by an operator to register with an AVS.
     *
     *  @param operator The account registering as an operator.
     *  @param avs The AVS the operator is registering with.
     *  @param salt A unique and single-use value associated with the approver's signature.
     *  @param expiry The time after which the approver's signature becomes invalid.
     */
    function calculateOperatorAVSRegistrationDigestHash(
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        // calculate the struct hash
        bytes32 structHash = keccak256(abi.encode(OPERATOR_AVS_REGISTRATION_TYPEHASH, operator, avs, salt, expiry));
        // calculate the digest hash
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", domainSeparator(), structHash));
        return digestHash;
    }

    /**
     * @notice Calculates the digest hash to be signed by an operator to register with an operator set.
     *
     * @param avs The AVS that operator is registering to operator sets for.
     * @param operatorSetIds An array of operator set IDs the operator is registering to.
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid.
     */
    function calculateOperatorSetRegistrationDigestHash(
        address avs,
        uint32[] calldata operatorSetIds,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        // calculate the struct hash
        bytes32 structHash =
            keccak256(abi.encode(OPERATOR_SET_REGISTRATION_TYPEHASH, avs, operatorSetIds, salt, expiry));

        // calculate the digest hash
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", domainSeparator(), structHash));

        return digestHash;
    }

    /// @notice Getter function for the current EIP-712 domain separator for this contract.
    /// @dev The domain separator will change in the event of a fork that changes the ChainID.
    function domainSeparator() public view returns (bytes32) {
        if (block.chainid == ORIGINAL_CHAIN_ID) {
            return _DOMAIN_SEPARATOR;
        } else {
            return _calculateDomainSeparator();
        }
    }

    /// @notice Internal function for calculating the current domain separator of this contract
    function _calculateDomainSeparator() internal view returns (bytes32) {
        return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("EigenLayer")), block.chainid, address(this)));
    }
}
