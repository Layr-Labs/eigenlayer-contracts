// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "../permissions/Pausable.sol";
import "../libraries/EIP1271SignatureUtils.sol";
import "../libraries/SlashingAccountingUtils.sol";
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
     *                         SLASHING FUNCTIONS POC
     *
     */

    /// For queue/complete reallocations
    enum MagnitudeAdjustmentType {
        ALLOCATION,
        DEALLOCATION
    }

    struct MagnitudeAdjustment {
        OperatorSet operatorSet;
        uint64 magnitudeDiff;
    }

    struct AllocatorOperatorDetails {
        bool isAllocatorFor;
        uint32 reallocationNonce;
    }

    struct QueuedReallocation {
        MagnitudeAdjustmentType magnitudeAdjustmentType;
        address allocator;
        address operator;
        OperatorSet operatorSet;
        IStrategy strategy;
        uint64 magnitudeDiff;
        uint32 nonce;
        uint64 queuedScalingFactor;
        uint32 timestamp;
    }

    struct ScalingFactorUpdate {
        uint32 timestamp;
        uint64 scalingFactor;
    }

    struct TotalAndNonslashableUpdate {
        uint32 timestamp;
        uint64 totalMagnitude;
        uint64 nonslashableMagnitude;
    }

    event OperatorSlashed(
        address operator,
        IStrategy strategy,
        OperatorSet opSet,
        address allocator,
        uint32 bipsToSlash
    );

    uint256 constant MAX_ALLOCATORS_PER_OPERATOR = 10; 
    uint256 constant QUEUED_REALLOCATION_DELAY = 21 days; // placeholder val for now

    uint64 constant DEFAULT_TOTAL_MAGNITUDE = 1e18; // 1e18 = 1


    /**
     *
     *                         SLASHING UTILS ACCOUNTING
     *
     */

    /// mapping: allocator => operator => AllocatorOperatorDetails
    mapping(address => mapping(address => AllocatorOperatorDetails)) public allocatorOperatorDetails;

    /// @notice queued reallocation updates
    /// mapping: hash(operator, strategy, avs, operatorSetId, allocator, allocatorNonce) => allocator => bool
    mapping(bytes32 => bool) public queuedReallocationHashes;

    /// Note: that we reuse hash of (strategy, avs, operatorSetId) as the key for the mappings below for
    /// to reduce number of key lookups and mapping hashes
    /// @notice all the allocators for a given (operator, strategy, OperatorSet) 
    /// length cannot be more than MAX_ALLOCATORS_PER_OPERATOR
    /// mapping: operator => hash(strategy, avs, operatorSetId) => address[] (operator's allocators for a strategy)
    mapping(address => mapping(bytes32 => address[])) private _operatorAllocators;


    /// @notice Updated whenever the operator gets slashed and their scaling factor is reduced.
    /// Scaling factors lie in the range [0,1]
    /// mapping: operator => hash(strategy, avs, operatorSetId) => ScalingFactorUpdate
    mapping(address => mapping(bytes32 => ScalingFactorUpdate[])) private _scalingFactorUpdates;

    /// mapping: operator => hash(strategy, avs, operatorSetId) => allocator => nonNormalizedMagnitude
    mapping(address => mapping(bytes32 => mapping(address => uint64))) private _nonNormalizedMagnitudes;

    /// mapping: operator => IStrategy => allocator => TotalAndNonslashableUpdate[]
    mapping(address => mapping(IStrategy => mapping(address => TotalAndNonslashableUpdate[]))) private _totalMagnitudeUpdates;


    /**
     * @notice Queues magnitude adjustment updates of type ALLOCATION or DEALLOCATION
     * The magnitude allocation takes 21 days from time when it is queued to take effect.
     *
     * @param allocator the allocator who is adjusting the magnitudes
     * @param operator the operator whom the magnitude parameters are being adjusted
     * @param strategy the strategy for which the magnitudes are being adjusted
     * @param magnitudeAdjustmentType the type of adjustment, either ALLOCATION or DEALLOCATION
     * @param magnitudeAdjustments an array of magnitude adjustments to be made for each operatorSet
     * @param allocatorSignature if non-empty is the signature of the allocator on
     * the modification. if empty, the msg.sender must be the allocator for the
     * operator
     *
     * @dev pushes a MagnitudeUpdate and TotalAndNonslashableUpdate to take effect in 21 days
     * @dev If ALLOCATION adjustment type:
     * reverts if sum of magnitudeDiffs > nonslashable magnitude for the latest pending update - sum of all pending allocations
     * since one cannot allocate more than is nonslahsable
     * @dev if DEALLOCATION adjustment type:
     * reverts if magnitudeDiff > allocated magnitude for the latest pending update
     * since one cannot deallocate more than is already allocated
     * @dev reverts if there are more than 3 pending allocations/deallocations for the given (operator, strategy, operatorSet) tuple
     * @dev emits events MagnitudeUpdated, TotalAndNonslashableMagnitudeUpdated
     */
    function queueReallocation(
        address allocator,
        address operator,
        IStrategy strategy,
        MagnitudeAdjustmentType magnitudeAdjustmentType,
        MagnitudeAdjustment[] calldata magnitudeAdjustments,
        SignatureWithSaltAndExpiry calldata allocatorSignature
    ) external returns (QueuedReallocation[] memory queuedReallocations) {
        AllocatorOperatorDetails storage allocatorDetails = allocatorOperatorDetails[allocator][operator];
        require(
            allocatorDetails.isAllocatorFor,
            "AVSDirectory.queueReallocation: allocator is not an allocator for operator"
        );
        if (msg.sender != allocator) {
            _verifyAllocatorSignature(allocator, operator, allocatorSignature);
        }
        uint32 reallocationNonce = allocatorDetails.reallocationNonce;

        if (magnitudeAdjustmentType == MagnitudeAdjustmentType.ALLOCATION) {
            queuedReallocations = _queueAllocation(allocator, operator, strategy, magnitudeAdjustments, allocatorDetails);
        } else {
            queuedReallocations = _queueDeallocation(allocator, operator, strategy, magnitudeAdjustments, allocatorDetails);
        }
    }

    function completeReallocation(QueuedReallocation[] memory queuedReallocations) external {
        for (uint256 i = 0; i < queuedReallocations.length; ++i) {
            bytes32 reallocationHash = calculateQueuedReallocationHash(queuedReallocations[i]);
            require(
                queuedReallocationHashes[reallocationHash],
                "AVSDirectory.completeReallocation: reallocation hash was not queued"
            );
            require(
                queuedReallocations[i].timestamp + QUEUED_REALLOCATION_DELAY < block.timestamp,
                "AVSDirectory.completeReallocation: queued reallocation not completable yet"
            );
            delete queuedReallocationHashes[reallocationHash];

            if (queuedReallocations[i].magnitudeAdjustmentType == MagnitudeAdjustmentType.ALLOCATION) {
                _completeAllocation(queuedReallocations[i]);
            } else {
                _completeDeallocation(queuedReallocations[i]);
            }
        }
    }

    /**
     * @notice Called by an AVS to slash an operator for a given operatorSetId, list of strategies, and bipsToSlash
     * @param operator the operator to slash
     * @param operatorSetId which operator set the operator is being slashed from
     * @param strategies list of strategies to slash the operator for
     * @param bipsToSlash the amount of bips to slash the operator by
     */
    function slashOperator(
        address operator,
        uint32 operatorSetId,
        IStrategy[] calldata strategies,
        uint32 bipsToSlash
    ) external {
        require(
            0 < bipsToSlash && bipsToSlash < SlashingAccountingUtils.BIPS_FACTOR,
            "AVSDirectory.slashOperator: invalid bipsToSlash"
        );
        // Slash the operator for each strategy
        for (uint256 i = 0; i < strategies.length; ++i) {
            OperatorSet memory opSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
            bytes32 opSetKey = _hashStrategyAndOperatorSet(strategies[i], opSet);
            address[] memory allocators = _operatorAllocators[operator][opSetKey];

            // For each strategy, slash each of its allocators and reduce their total magnitudes
            for (uint256 j = 0; j < allocators.length; ++j) {
                _slashAllocator(operator, opSetKey, allocators[j], bipsToSlash);
                emit OperatorSlashed(operator, strategies[i], opSet, allocators[j], bipsToSlash);
            }
        }
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    function _queueAllocation(
        address allocator,
        address operator,
        IStrategy strategy,
        MagnitudeAdjustment[] calldata magnitudeAdjustments,
        AllocatorOperatorDetails storage allocatorDetails
    ) internal returns (QueuedReallocation[] memory) {
        TotalAndNonslashableUpdate storage totalMagnitudeUpdate = _getLatestTotalAndNonslashableUpdate(operator, strategy, allocator);
        uint64 nonslashableMagnitude = totalMagnitudeUpdate.nonslashableMagnitude;
        uint32 nonce = allocatorDetails.reallocationNonce;

        QueuedReallocation[] memory queuedAllocations = new QueuedReallocation[](magnitudeAdjustments.length);

        for (uint256 i = 0; i < magnitudeAdjustments.length; ++i) {
            uint64 scalingFactor = _getCurrentScalingFactor(
                operator,
                _hashStrategyAndOperatorSet(strategy, magnitudeAdjustments[i].operatorSet)
            );
            uint64 normalizedMagnitudeIncrement = SlashingAccountingUtils.normalizeMagnitude({
                nonNormalizedMagnitude: magnitudeAdjustments[i].magnitudeDiff,
                scalingFactor: scalingFactor
            });
            require(
                normalizedMagnitudeIncrement <= nonslashableMagnitude,
                "AVSDirectory._queueAllocation: normalized magnitude increment exceeds nonslashable magnitude"
            );
            
            queuedAllocations[i] = QueuedReallocation({
                magnitudeAdjustmentType: MagnitudeAdjustmentType.ALLOCATION,
                allocator: allocator,
                operator: operator,
                operatorSet: magnitudeAdjustments[i].operatorSet,
                strategy: strategy,
                magnitudeDiff: magnitudeAdjustments[i].magnitudeDiff,
                nonce: nonce,
                queuedScalingFactor: scalingFactor,
                timestamp: uint32(block.timestamp)
            });

            // set queued allocation hash struct in mapping as pending
            bytes32 reallocationHash = calculateQueuedReallocationHash(queuedAllocations[i]);
            queuedReallocationHashes[reallocationHash] = true;

            // decrement from non slashable magnitude
            nonslashableMagnitude -= normalizedMagnitudeIncrement;
            ++nonce;
        }
        allocatorDetails.reallocationNonce = nonce;

        return queuedAllocations;
    }

    function _queueDeallocation(
        address allocator,
        address operator,
        IStrategy strategy,
        MagnitudeAdjustment[] calldata magnitudeAdjustments,
        AllocatorOperatorDetails storage allocatorDetails
    ) internal returns (QueuedReallocation[] memory) {
        uint32 nonce = allocatorDetails.reallocationNonce;
        QueuedReallocation[] memory queuedDeallocations = new QueuedReallocation[](magnitudeAdjustments.length);

        for (uint256 i = 0; i < magnitudeAdjustments.length; ++i) {
            bytes32 opSetKey = _hashStrategyAndOperatorSet(strategy, magnitudeAdjustments[i].operatorSet);
            uint64 scalingFactor = _getCurrentScalingFactor(
                operator,
                _hashStrategyAndOperatorSet(strategy, magnitudeAdjustments[i].operatorSet)
            );
            uint64 normalizedMagnitudeDecrement = SlashingAccountingUtils.normalizeMagnitude({
                nonNormalizedMagnitude: magnitudeAdjustments[i].magnitudeDiff,
                scalingFactor: scalingFactor
            });
            require(
                normalizedMagnitudeDecrement <= _nonNormalizedMagnitudes[operator][opSetKey][allocator],
                "AVSDirectory._queueDeallocation: normalized magnitude decrement exceeds allocated magnitude"
            );
            
            queuedDeallocations[i] = QueuedReallocation({
                magnitudeAdjustmentType: MagnitudeAdjustmentType.ALLOCATION,
                allocator: allocator,
                operator: operator,
                operatorSet: magnitudeAdjustments[i].operatorSet,
                strategy: strategy,
                magnitudeDiff: magnitudeAdjustments[i].magnitudeDiff,
                nonce: nonce,
                queuedScalingFactor: scalingFactor,
                timestamp: uint32(block.timestamp)
            });

            // set queued allocation hash struct in mapping as pending
            bytes32 reallocationHash = calculateQueuedReallocationHash(queuedDeallocations[i]);
            queuedReallocationHashes[reallocationHash] = true;
            ++nonce;
        }
        allocatorDetails.reallocationNonce = nonce;

        return queuedDeallocations;
    }

    function _completeAllocation(QueuedReallocation memory allocation) internal {}

    function _completeDeallocation(QueuedReallocation memory deallocation) internal {}


    function _verifyAllocatorSignature(
        address allocator,
        address operator,
        SignatureWithSaltAndExpiry calldata allocatorSignature
    ) internal {
        require(
            allocatorSignature.expiry >= block.timestamp,
            "AVSDirectory._verifyAllocatorSignature: allocator signature expired"
        );
        // TODO with added storage
    }

    /// @notice slash an operator's allocator
    function _slashAllocator(
        address operator,
        bytes32 opSetKey,
        address allocator,
        uint32 bipsToSlash
    ) internal {
        /// TODO
    }

    /// @notice Return latest total and nonslashable update for the given operator and strategy
    /// if empty array, push an empty update
    function _getLatestTotalAndNonslashableUpdate(
        address operator,
        IStrategy strategy,
        address allocator
    ) internal returns (TotalAndNonslashableUpdate storage) {
        uint256 totalMagnitudeUpdatesLength = _totalMagnitudeUpdates[operator][strategy][allocator].length;
        if (totalMagnitudeUpdatesLength == 0) {
            // used for very first reallocation so use default values
            _totalMagnitudeUpdates[operator][strategy][allocator].push(TotalAndNonslashableUpdate({
                timestamp: 0,
                totalMagnitude: DEFAULT_TOTAL_MAGNITUDE,
                nonslashableMagnitude: DEFAULT_TOTAL_MAGNITUDE
            }));
        }
        return _totalMagnitudeUpdates[operator][strategy][allocator][totalMagnitudeUpdatesLength - 1];
    }

    function getCurrentScalingFactor(
        address operator,
        IStrategy strategy,
        OperatorSet calldata operatorSet
    ) public returns (uint64) {
        bytes32 opSetKey = _hashStrategyAndOperatorSet(strategy, operatorSet);
        return _getCurrentScalingFactor(operator, opSetKey);
    }

    function _getCurrentScalingFactor(address operator, bytes32 opSetKey) internal returns (uint64 scalingFactor) {
        uint256 scalingFactorUpdatesLength = _scalingFactorUpdates[operator][opSetKey].length;
        if (scalingFactorUpdatesLength == 0) {
            return SlashingAccountingUtils.SHARE_CONVERSION_SCALE;
        } else {
            return _scalingFactorUpdates[operator][opSetKey][scalingFactorUpdatesLength - 1].scalingFactor;
        }
    }

    function calculateQueuedReallocationHash(QueuedReallocation memory reallocation) internal pure returns (bytes32) {
        return keccak256(abi.encode(reallocation));
    }

    function _hashStrategyAndOperatorSet(
        IStrategy strategy,
        OperatorSet memory opSet
    ) internal pure returns (bytes32) {
        return keccak256(abi.encode(strategy, opSet));
    }

    function _hashReallocation() internal pure returns (bytes32) {

    }
}
