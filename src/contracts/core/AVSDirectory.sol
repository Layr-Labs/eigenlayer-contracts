// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import {CheckpointsUpgradeable} from "@openzeppelin-upgrades-v4.9.0/contracts/utils/CheckpointsUpgradeable.sol";

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

    /// @dev Returns the chain ID from the time the contract was deployed.
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the immutable addresses of the strategy mananger, delegationManager, slasher,
     * and eigenpodManager contracts
     */
    constructor(IDelegationManager _delegation) AVSDirectoryStorage(_delegation) {
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
     *                    EXTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Called by an AVS to create a list of new operatorSets.
     *
     * @param operatorSetIds The IDs of the operator set to initialize.
     *
     * @dev msg.sender must be the AVS.
     * @dev The AVS may create operator sets before it becomes an operator set AVS.
     */
    function createOperatorSets(uint32[] calldata operatorSetIds) external {
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            require(
                !isOperatorSet[msg.sender][operatorSetIds[i]],
                "AVSDirectory.createOperatorSet: operator set already exists"
            );
            isOperatorSet[msg.sender][operatorSetIds[i]] = true;
            emit OperatorSetCreated(OperatorSet({avs: msg.sender, operatorSetId: operatorSetIds[i]}));
        }
    }

    /**
     * @notice Sets the AVS as an operator set AVS, preventing legacy M2 operator registrations.
     *
     * @dev msg.sender must be the AVS.
     */
    function becomeOperatorSetAVS() external {
        require(!isOperatorSetAVS[msg.sender], "AVSDirectory.becomeOperatorSetAVS: already an operator set AVS");
        isOperatorSetAVS[msg.sender] = true;
        emit AVSMigratedToOperatorSets(msg.sender);
    }

    /**
     * @notice Called by an AVS to migrate operators that have a legacy M2 registration to operator sets.
     *
     * @param operators The list of operators to migrate
     * @param operatorSetIds The list of operatorSets to migrate the operators to
     *
     * @dev The msg.sender used is the AVS
     * @dev The operator can only be migrated at most once per AVS
     * @dev The AVS can no longer register operators via the legacy M2 registration path once it begins migration
     * @dev The operator is deregistered from the M2 legacy AVS once migrated
     */
    function migrateOperatorsToOperatorSets(
        address[] calldata operators,
        uint32[][] calldata operatorSetIds
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        // Assert that the AVS is an operator set AVS.
        require(
            isOperatorSetAVS[msg.sender], "AVSDirectory.migrateOperatorsToOperatorSets: AVS is not an operator set AVS"
        );

        for (uint256 i = 0; i < operators.length; i++) {
            // Assert that the operator is registered & has not been migrated.
            require(
                avsOperatorStatus[msg.sender][operators[i]] == OperatorAVSRegistrationStatus.REGISTERED,
                "AVSDirectory.migrateOperatorsToOperatorSets: operator already migrated or not a legacy registered operator"
            );

            // Migrate operator to operator sets.
            _registerToOperatorSets(msg.sender, operators[i], operatorSetIds[i]);

            // Deregister operator from AVS - this prevents the operator from being migrated again since
            // the AVS can no longer use the legacy M2 registration path
            avsOperatorStatus[msg.sender][operators[i]] = OperatorAVSRegistrationStatus.UNREGISTERED;
            emit OperatorAVSRegistrationStatusUpdated(
                operators[i], msg.sender, OperatorAVSRegistrationStatus.UNREGISTERED
            );
            emit OperatorMigratedToOperatorSets(operators[i], msg.sender, operatorSetIds[i]);
        }
    }

    /**
     *  @notice Called by AVSs to add an operator to a list of operatorSets.
     *
     *  @param operator The address of the operator to be added to the operator set.
     *  @param operatorSetIds The IDs of the operator sets.
     *  @param operatorSignature The signature of the operator on their intent to register.
     *
     *  @dev msg.sender is used as the AVS.
     *  @dev The operator must not have a pending deregistration from the operator set.
     */
    function registerOperatorToOperatorSets(
        address operator,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        // Assert operator's signature has not expired.
        require(
            operatorSignature.expiry >= block.timestamp,
            "AVSDirectory.registerOperatorToOperatorSets: operator signature expired"
        );
        // Assert `operator` is actually an operator.
        require(
            delegation.isOperator(operator),
            "AVSDirectory.registerOperatorToOperatorSets: operator not registered to EigenLayer yet"
        );
        // Assert that the AVS is an operator set AVS.
        require(
            isOperatorSetAVS[msg.sender], "AVSDirectory.registerOperatorToOperatorSets: AVS is not an operator set AVS"
        );
        // Assert operator's signature `salt` has not already been spent.
        require(
            !operatorSaltIsSpent[operator][operatorSignature.salt],
            "AVSDirectory.registerOperatorToOperatorSets: salt already spent"
        );

        // Assert that `operatorSignature.signature` is a valid signature for operator set registrations.
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

        _registerToOperatorSets(msg.sender, operator, operatorSetIds);
    }

    /**
     * @notice Called by an operator to deregister from an operator set
     *
     * @param operator The operator to deregister from the operatorSets.
     * @param avs The address of the AVS to deregister the operator from.
     * @param operatorSetIds The IDs of the operator sets.
     * @param operatorSignature the signature of the operator on their intent to deregister or empty if the operator itself is calling
     *
     * @dev if the operatorSignature is empty, the caller must be the operator
     * @dev this will likely only be called in case the AVS contracts are in a state that prevents operators from deregistering
     */
    function forceDeregisterFromOperatorSets(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        if (operatorSignature.signature.length == 0) {
            require(msg.sender == operator, "AVSDirectory.forceDeregisterFromOperatorSets: caller must be operator");
        } else {
            // Assert operator's signature has not expired.
            require(
                operatorSignature.expiry >= block.timestamp,
                "AVSDirectory.forceDeregisterFromOperatorSets: operator signature expired"
            );
            // Assert operator's signature `salt` has not already been spent.
            require(
                !operatorSaltIsSpent[operator][operatorSignature.salt],
                "AVSDirectory.forceDeregisterFromOperatorSets: salt already spent"
            );

            // Assert that `operatorSignature.signature` is a valid signature for operator set deregistrations.
            EIP1271SignatureUtils.checkSignature_EIP1271(
                operator,
                calculateOperatorSetForceDeregistrationTypehash({
                    avs: avs,
                    operatorSetIds: operatorSetIds,
                    salt: operatorSignature.salt,
                    expiry: operatorSignature.expiry
                }),
                operatorSignature.signature
            );

            // Mutate `operatorSaltIsSpent` to `true` to prevent future respending.
            operatorSaltIsSpent[operator][operatorSignature.salt] = true;
        }
        _deregisterFromOperatorSets(avs, operator, operatorSetIds);
    }

    /**
     *  @notice Called by AVSs to remove an operator from an operator set.
     *
     *  @param operator The address of the operator to be removed from the operator set.
     *  @param operatorSetIds The IDs of the operator sets.
     *
     *  @dev msg.sender is used as the AVS.
     */
    function deregisterOperatorFromOperatorSets(
        address operator,
        uint32[] calldata operatorSetIds
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        _deregisterFromOperatorSets(msg.sender, operator, operatorSetIds);
    }

    /**
     *  @notice Called by an AVS to emit an `AVSMetadataURIUpdated` event indicating the information has updated.
     *
     *  @param metadataURI The URI for metadata associated with an AVS.
     *
     *  @dev Note that the `metadataURI` is *never stored* and is only emitted in the `AVSMetadataURIUpdated` event.
     */
    function updateAVSMetadataURI(string calldata metadataURI) external override {
        emit AVSMetadataURIUpdated(msg.sender, metadataURI);
    }

    /**
     * @notice Called by an operator to cancel a salt that has been used to register with an AVS.
     *
     * @param salt A unique and single use value associated with the approver signature.
     */
    function cancelSalt(bytes32 salt) external override {
        // Mutate `operatorSaltIsSpent` to `true` to prevent future spending.
        operatorSaltIsSpent[msg.sender][salt] = true;
    }

    /**
     *
     *        LEGACY EXTERNAL FUNCTIONS - SUPPORT DEPRECATED BY SLASHING RELEASE
     *
     */

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
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        // Assert `operatorSignature.expiry` has not elapsed.
        require(
            operatorSignature.expiry >= block.timestamp,
            "AVSDirectory.registerOperatorToAVS: operator signature expired"
        );

        // Assert that the AVS is not an operator set AVS.
        require(!isOperatorSetAVS[msg.sender], "AVSDirectory.registerOperatorToAVS: AVS is an operator set AVS");

        // Assert that the `operator` is not actively registered to the AVS.
        require(
            avsOperatorStatus[msg.sender][operator] != OperatorAVSRegistrationStatus.REGISTERED,
            "AVSDirectory.registerOperatorToAVS: operator already registered"
        );

        // Assert `operator` has not already spent `operatorSignature.salt`.
        require(
            !operatorSaltIsSpent[operator][operatorSignature.salt],
            "AVSDirectory.registerOperatorToAVS: salt already spent"
        );

        // Assert `operator` is a registered operator.
        require(
            delegation.isOperator(operator),
            "AVSDirectory.registerOperatorToAVS: operator not registered to EigenLayer yet"
        );

        // Assert that `operatorSignature.signature` is a valid signature for the operator AVS registration.
        EIP1271SignatureUtils.checkSignature_EIP1271({
            signer: operator,
            digestHash: calculateOperatorAVSRegistrationDigestHash({
                operator: operator,
                avs: msg.sender,
                salt: operatorSignature.salt,
                expiry: operatorSignature.expiry
            }),
            signature: operatorSignature.signature
        });

        // Mutate `operatorSaltIsSpent` to `true` to prevent future respending.
        operatorSaltIsSpent[operator][operatorSignature.salt] = true;

        // Set the operator as registered
        avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.REGISTERED;

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
        override
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
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Helper function used by migration & registration functions to register an operator to operator sets.
     * @param avs The AVS that the operator is registering to.
     * @param operator The operator to register.
     * @param operatorSetIds The IDs of the operator sets.
     */
    function _registerToOperatorSets(address avs, address operator, uint32[] calldata operatorSetIds) internal {
        // Loop over `operatorSetIds` array and register `operator` for each item.
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            require(
                isOperatorSet[avs][operatorSetIds[i]],
                "AVSDirectory._registerOperatorToOperatorSets: invalid operator set"
            );

            // Assert `operator` has not already been registered to `operatorSetIds[i]`.
            require(
                !isMember[avs][operator][operatorSetIds[i]],
                "AVSDirectory._registerOperatorToOperatorSets: operator already registered to operator set"
            );

            // Mutate `isMember` to `true`.
            isMember[avs][operator][operatorSetIds[i]] = true;

            OperatorSetRegistrationStatus storage registrationStatus =
                operatorSetStatus[avs][operator][operatorSetIds[i]];
            require(
                !registrationStatus.registered,
                "AVSDirectory._registerOperatorToOperatorSets: operator already registered for operator set"
            );
            registrationStatus.registered = true;

            emit OperatorAddedToOperatorSet(operator, OperatorSet({avs: avs, operatorSetId: operatorSetIds[i]}));
        }
    }

    /**
     * @notice Internal function to deregister an operator from an operator set.
     *
     * @param avs The AVS that the operator is deregistering from.
     * @param operator The operator to deregister.
     * @param operatorSetIds The IDs of the operator sets.
     */
    function _deregisterFromOperatorSets(address avs, address operator, uint32[] calldata operatorSetIds) internal {
        // Loop over `operatorSetIds` array and deregister `operator` for each item.
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            // Assert `operator` is registered for this iterations operator set.
            require(
                isMember[avs][operator][operatorSetIds[i]],
                "AVSDirectory._deregisterOperatorFromOperatorSet: operator not registered for operator set"
            );

            // Mutate `isMember` to `false`.
            isMember[avs][operator][operatorSetIds[i]] = false;

            require(
                operatorSetStatus[avs][operator][operatorSetIds[i]].registered,
                "AVSDirectory._deregisterOperatorFromOperatorSet: operator not registered for operator set"
            );
            operatorSetStatus[avs][operator][operatorSetIds[i]] =
                OperatorSetRegistrationStatus({registered: true, lastDeregisteredTimestamp: uint32(block.timestamp)});

            emit OperatorRemovedFromOperatorSet(operator, OperatorSet({avs: avs, operatorSetId: operatorSetIds[i]}));
        }
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
    ) public view override returns (bytes32) {
        return
            _calculateDigestHash(keccak256(abi.encode(OPERATOR_AVS_REGISTRATION_TYPEHASH, operator, avs, salt, expiry)));
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
    ) public view override returns (bytes32) {
        return _calculateDigestHash(
            keccak256(abi.encode(OPERATOR_SET_REGISTRATION_TYPEHASH, avs, operatorSetIds, salt, expiry))
        );
    }

    /**
     * @notice Calculates the digest hash to be signed by an operator to force deregister from an operator set.
     *
     * @param avs The AVS that operator is deregistering from.
     * @param operatorSetIds An array of operator set IDs the operator is deregistering from.
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid.
     */
    function calculateOperatorSetForceDeregistrationTypehash(
        address avs,
        uint32[] calldata operatorSetIds,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        return _calculateDigestHash(
            keccak256(abi.encode(OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH, avs, operatorSetIds, salt, expiry))
        );
    }

    /// @notice Getter function for the current EIP-712 domain separator for this contract.
    /// @dev The domain separator will change in the event of a fork that changes the ChainID.
    function domainSeparator() public view override returns (bytes32) {
        return _calculateDomainSeparator();
    }

    /// @notice Internal function for calculating the current domain separator of this contract
    function _calculateDomainSeparator() internal view returns (bytes32) {
        if (block.chainid == ORIGINAL_CHAIN_ID) {
            return _DOMAIN_SEPARATOR;
        } else {
            return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("EigenLayer")), block.chainid, address(this)));
        }
    }

    function _calculateDigestHash(bytes32 structHash) internal view returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", _calculateDomainSeparator(), structHash));
    }

    /**
     *
     *                         ALLOCATOR AND SLASHING FUNCTIONS
     *
     */
    using CheckpointsUpgradeable for CheckpointsUpgradeable.Trace224;

    enum MagnitudeAdjustmentType {
        ALLOCATE,
        DEALLOCATE
    }

    /**
     * @notice this struct is used in allocate and queueDeallocation in order to specify an operator's slashability for a certain operator set
     *
     * @param strategy the strategy to adjust slashable stake for
     * @param operatorSets the operator sets to adjust slashable stake for
     * @param magnitudeDiff magnitude difference; the difference in proportional parts of the operator's slashable stake
     * that is slashable by the operatorSet.
     * Slashable stake for an operator set is (magnitude / sum of all magnitudes for the strategy/operator + nonSlashableMagnitude) of
     * an operator's delegated stake.
     */
    struct MagnitudeAdjustment {
        MagnitudeAdjustmentType adjustmentType;
        IStrategy strategy;
        OperatorSet[] operatorSets;
        uint64[] magnitudeDiffs;
    }

    /**
     * @notice struct used for queued deallocations. Hash of struct is set in storage to be referenced later when completing deallocations.
     */
    struct QueuedDeallocation {
        address operator;
        uint16 nonce;
        uint32 timestamp;
        IStrategy strategy;
        OperatorSet[] operatorSets;
        uint64[] magnitudeDiffs;
    }

    uint256 internal constant BIPS_FACTOR = 10_000;

    uint32 public constant ALLOCATION_DELAY = 21 days;

    /// (operator, strategy, timestamp) => checkpointed totalMagnitude
    /// Note that totalMagnitude is monotonically decreasing and only gets updated upon slashing
    /// mapping: operator => strategy => checkpointed totalMagnitude
    mapping(address => mapping(IStrategy => CheckpointsUpgradeable.Trace224)) private _totalMagnitudeUpdate;

    /// (operator, strategy, timestamp) => nonslashable magnitude
    /// Note that nonslashable magnitude updates in the future are only due to allocations. Deallocations
    /// require a complete tx step to update the nonslashableMagnitude with a checkpointed blocknumber at current block.
    /// mapping: operator => strategy => checkpointed nonSlashableMagnitude
    mapping(address => mapping(IStrategy => CheckpointsUpgradeable.Trace224)) private _nonSlashableMagnitudeUpdate;

    /// (allocator, strategy, OperatorSet) => checkpointed decrementable magnitudes
    /// mapping: operator => strategy => avs => operatorSetId => checkpointed magnitude
    mapping(address => mapping(IStrategy => mapping(address => mapping(uint32 => CheckpointsUpgradeable.Trace224))))
        private _magnitudeUpdate;

    function queueReallocation(
        address operator,
        MagnitudeAdjustment[] calldata reallocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external {
        // TODO signature verification

        for (uint256 i = 0; i < reallocations.length; ++i) {
            if (reallocations[i].adjustmentType == MagnitudeAdjustmentType.ALLOCATE) {
                _queueAllocation(operator, reallocations[i]);
            } else {
                _queueDeallocation(operator, reallocations[i]);
            }
        }
    }

    function _queueAllocation(address operator, MagnitudeAdjustment calldata allocation) internal {
        IStrategy strategy = allocation.strategy;
        OperatorSet[] calldata operatorSets = allocation.operatorSets;
        uint32 effectTimestamp = uint32(block.timestamp) + ALLOCATION_DELAY;
        uint64 nonslashableMagnitude = _getLatestNonslashableMagnitude(operator, strategy);

        for (uint256 i = 0; i < operatorSets.length; ++i) {
            require(
                nonslashableMagnitude >= allocation.magnitudeDiffs[i],
                "AVSDirectory.queueAllocations: insufficient available nonslashable magnitude"
            );

            nonslashableMagnitude -= allocation.magnitudeDiffs[i];

            (bool exists, uint32 timestamp, uint224 value) = _magnitudeUpdate[operator][strategy][operatorSets[i].avs][operatorSets[i]
                .operatorSetId].latestCheckpoint();
            if (exists) {
                require(
                    timestamp <= block.timestamp,
                    "AVSDirectory.queueAllocations: only one pending allocation allowed for op, opSet, strategy"
                );
            }

            // Push queued allocation update
            _magnitudeUpdate[operator][strategy][operatorSets[i].avs][operatorSets[i].operatorSetId].push(
                effectTimestamp, value + uint224(allocation.magnitudeDiffs[i])
            );
            // decrement nonslashableMagnitude
            nonslashableMagnitude -= allocation.magnitudeDiffs[i];
        }

        // Push a new nonSlashableMagnitude update
        _nonSlashableMagnitudeUpdate[operator][strategy].push(effectTimestamp, uint224(nonslashableMagnitude));
    }

    function _queueDeallocation(address operator, MagnitudeAdjustment calldata deallocation) internal {
        IStrategy strategy = deallocation.strategy;
        OperatorSet[] calldata operatorSets = deallocation.operatorSets;
        uint32 effectTimestamp = uint32(block.timestamp) + ALLOCATION_DELAY;
        uint64 deallocationSum = 0;

        for (uint256 i = 0; i < operatorSets.length; ++i) {
            (bool exists, uint32 timestamp, uint224 value) = _magnitudeUpdate[operator][strategy][operatorSets[i].avs][operatorSets[i]
                .operatorSetId].latestCheckpoint();
            if (exists) {
                require(
                    timestamp <= block.timestamp,
                    "AVSDirectory.queueAllocations: only one pending allocation allowed for op, opSet, strategy"
                );
            }

            require(
                deallocation.magnitudeDiffs[i] <= uint64(value),
                "AVSDirectory.queueDeallocation: cannot deallocate more than what is allocated"
            );

            // Push queued deallocation update
            _magnitudeUpdate[operator][strategy][operatorSets[i].avs][operatorSets[i].operatorSetId].push(
                effectTimestamp, value - uint224(deallocation.magnitudeDiffs[i])
            );

            // keep track of running sum to add back to nonslashableMagnitude
            deallocationSum += deallocation.magnitudeDiffs[i];
        }

        // Push a new nonSlashableMagnitude update
        uint64 nonslashableMagnitude = _getLatestNonslashableMagnitude(operator, strategy);
        _nonSlashableMagnitudeUpdate[operator][strategy].push(
            effectTimestamp, nonslashableMagnitude + uint224(deallocationSum)
        );
    }

    /**
     * @notice Called by an AVS to slash an operator for given operatorSetId, list of strategies, and bipsToSlash.
     * For each given (operator, operatorSetId, strategy) tuple, bipsToSlash
     * bips of the operatorSet's slashable stake allocation will be slashed
     *
     * @param operator the address to slash
     * @param operatorSetId the ID of the operatorSet the operator is being slashed on behalf of
     * @param strategies the set of strategies to slash
     * @param bipsToSlash the number of bips to slash, this will be proportional to the
     * operator's slashable stake allocation for the operatorSet
     */
    function slashOperator(
        address operator,
        uint32 operatorSetId,
        IStrategy[] calldata strategies,
        uint16 bipsToSlash
    ) external {
        require(
            isOperatorSlashable(operator, OperatorSet({avs: msg.sender, operatorSetId: operatorSetId})),
            "AVSDirectory.slashOperator: operator not slashable for operatorSet"
        );

        for (uint256 i = 0; i < strategies.length; ++i) {
            // IStrategy strategy = strategies[i];

            uint64 currentMagnitude = uint64(
                _magnitudeUpdate[operator][strategies[i]][msg.sender][operatorSetId].upperLookupRecent(
                    uint32(block.timestamp)
                )
            );
            uint64 currentTotalMagnitude = uint64(
                _magnitudeUpdate[operator][strategies[i]][msg.sender][operatorSetId].upperLookupRecent(
                    uint32(block.timestamp)
                )
            );

            uint64 slashedMagnitude = uint64(uint256(bipsToSlash) * uint256(currentMagnitude) / BIPS_FACTOR);

            // 1. decrement slashedMagnitude from current magnitude and totalMagnitude
            _totalMagnitudeUpdate[operator][strategies[i]].push({
                key: uint32(block.timestamp),
                value: uint224(currentTotalMagnitude - slashedMagnitude)
            });
            _magnitudeUpdate[operator][strategies[i]][msg.sender][operatorSetId].push({
                key: uint32(block.timestamp),
                value: uint224(currentMagnitude - slashedMagnitude)
            });

            // 2. handle if there is a pending magnitude update in the future, decrement as well
            // we know there can at most be only 1 pending update
            (, uint32 timestamp, uint224 value) =
                _magnitudeUpdate[operator][strategies[i]][msg.sender][operatorSetId].latestCheckpoint();
            // if there is a pending magitude update in the future, decrement as well
            if (timestamp > block.timestamp) {
                // slashed amount is greater than future allocation value, requires decrementing from nonslashableMagnitude as well
                if (slashedMagnitude > uint64(value)) {
                    _magnitudeUpdate[operator][strategies[i]][msg.sender][operatorSetId].push({key: timestamp, value: 0});

                    // remainder to subtract = slashedMagnitude - uint64(value)
                    _nonSlashableMagnitudeUpdate[operator][strategies[i]].push({
                        key: timestamp,
                        value: _getLatestNonslashableMagnitude(operator, strategies[i]) + value - uint224(slashedMagnitude)
                    });
                } else {
                    _magnitudeUpdate[operator][strategies[i]][msg.sender][operatorSetId].push({
                        key: timestamp,
                        value: value - slashedMagnitude
                    });
                }
            }
        }
    }

    /**
     * @param operator the operator to get the slashable bips for
     * @param operatorSet the operatorSet to get the slashable bips for
     * @param strategy the strategy to get the slashable bips for
     * @param timestamp the timestamp to get the slashable bips for for
     *
     * @return slashableBips the slashable bips of the given strategy owned by
     * the given OperatorSet for the given operator and timestamp
     */
    function getSlashableBips(
        address operator,
        OperatorSet calldata operatorSet,
        IStrategy strategy,
        uint32 timestamp
    ) public view returns (uint16) {
        uint64 totalMagnitude = uint64(_totalMagnitudeUpdate[operator][strategy].upperLookup(timestamp));
        uint64 currentMagnitude = uint64(
            _magnitudeUpdate[operator][strategy][operatorSet.avs][operatorSet.operatorSetId].upperLookup(timestamp)
        );

        return uint16(currentMagnitude * BIPS_FACTOR / totalMagnitude);
    }

    /// @notice operator is slashable by operatorSet if currently registered OR last deregistered within 21 days
    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        OperatorSetRegistrationStatus memory registrationStatus =
            operatorSetStatus[operatorSet.avs][operator][operatorSet.operatorSetId];
        return registrationStatus.registered
            || registrationStatus.lastDeregisteredTimestamp + ALLOCATION_DELAY >= block.timestamp;
    }

    function hashQueuedDeallocation(QueuedDeallocation memory deallocation) public pure returns (bytes32) {
        return keccak256(abi.encode(deallocation));
    }

    function _getCurrentTotalMagnitude(address operator, IStrategy strategy) internal view returns (uint64) {
        uint224 latestValue = _totalMagnitudeUpdate[operator][strategy].latest();

        // TODO handle if latestValue or currValue is 0, i.e no checkpoints exist

        return uint64(latestValue);
    }

    function _getLatestNonslashableMagnitude(address operator, IStrategy strategy) internal view returns (uint64) {
        uint224 latestValue = _nonSlashableMagnitudeUpdate[operator][strategy].latest();
        // TODO handle if latestValue or currValue is 0, i.e no checkpoints exist

        return uint64(latestValue);
    }
}
