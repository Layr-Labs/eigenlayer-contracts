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
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using Checkpoints for Checkpoints.History;

    /// @dev Index for flag that pauses operator register/deregister to avs when set.
    uint8 internal constant PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS = 0;
    /// @dev Index for flag that pauses operator register/deregister to operator sets when set.
    uint8 internal constant PAUSER_OPERATOR_REGISTER_DEREGISTER_TO_OPERATOR_SETS = 1;

    /// @dev BIPS factor for slashable bips
    uint256 internal constant BIPS_FACTOR = 10_000;
    /// @dev Delay before allocations take effect and how long until deallocations are completable
    uint32 public constant ALLOCATION_DELAY = 21 days;
    /// @dev Maximum number of pending updates that can be queued for allocations/deallocations
    uint256 public constant MAX_PENDING_UPDATES = 1;

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
    ) external override onlyWhenNotPaused(PAUSER_OPERATOR_REGISTER_DEREGISTER_TO_OPERATOR_SETS) {
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
            _registerToOperatorSets(operators[i], msg.sender, operatorSetIds[i]);

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
    ) external override onlyWhenNotPaused(PAUSER_OPERATOR_REGISTER_DEREGISTER_TO_OPERATOR_SETS) {
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

        _registerToOperatorSets(operator, msg.sender, operatorSetIds);
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
    ) external override onlyWhenNotPaused(PAUSER_OPERATOR_REGISTER_DEREGISTER_TO_OPERATOR_SETS) {
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
    ) external override onlyWhenNotPaused(PAUSER_OPERATOR_REGISTER_DEREGISTER_TO_OPERATOR_SETS) {
        _deregisterFromOperatorSets(msg.sender, operator, operatorSetIds);
    }

    /**
     * @notice Allocates a set of magnitude adjustments to increase the slashable stake of an operator set for the
     * given operator for the given strategy.
     * free magnitude for each strategy will decrement by the sum of all
     * allocations for that strategy and the allocations will take effect 21 days from calling.
     *
     * @param operator address to increase allocations for
     * @param allocations array of magnitude adjustments for multiple strategies and corresponding operator sets
     * @param operatorSignature signature of the operator if msg.sender is not the operator
     */
    function allocate(
        address operator,
        MagnitudeAdjustment[] calldata allocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external {
        // TODO signature verification

        uint32 effectTimestamp = uint32(block.timestamp) + ALLOCATION_DELAY;
        for (uint256 i = 0; i < allocations.length; ++i) {
            _allocate(operator, allocations[i], effectTimestamp);
        }
    }

    /**
     * @notice Queues a set of magnitude adjustments to decrease the slashable stake of an operator set for the given operator for the given strategy.
     * The deallocations will take effect 21 days from calling. In order for the operator to have their nonslashable magnitude increased,
     * they must call the contract again to complete the deallocation. Stake deallocations are still subject to slashing until 21 days have passed since queuing.
     *
     * @param operator address to decrease allocations for
     * @param deallocations array of magnitude adjustments for multiple strategies and corresponding operator sets
     * @param operatorSignature signature of the operator if msg.sender is not the operator
     */
    function queueDeallocate(
        address operator,
        MagnitudeAdjustment[] calldata deallocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external {
        // TODO signature verification

        uint32 completableTimestamp = uint32(block.timestamp) + ALLOCATION_DELAY;
        // deallocate for each strategy
        for (uint256 i = 0; i < deallocations.length; ++i) {
            _queueDeallocate(operator, deallocations[i], completableTimestamp);
        }
    }

    /**
     * @notice Complete queued deallocations of slashable stake for an operator, permissionlessly called by anyone
     * Increments the free magnitude of the operator by the sum of all deallocation amounts for each strategy.
     * If the operator was slashed, this will be a smaller amount than during queuing.
     *
     * @param operator address to complete deallocations for
     * @param strategies a list of strategies to complete deallocations for
     * @param operatorSets a 2d list of operator sets to complete deallocations for, one list for each strategy
     *
     * @dev can be called permissionlessly by anyone
     */
    function completeDeallocations(
        address operator,
        IStrategy[] calldata strategies,
        OperatorSet[][] calldata operatorSets
    ) external {
        // complete all queued deallocations for strategies
        for (uint256 i = 0; i < strategies.length; ++i) {
            uint64 freeMagnitudeToAdd = 0;
            // complete all queued deallocations for specified operatorSets
            for (uint256 j = 0; j < operatorSets[i].length; ++j) {
                freeMagnitudeToAdd += _completeDeallocation(operator, strategies[i], operatorSets[i][j]);
            }
            // add to free available magnitude
            freeMagnitude[operator][strategies[i]] += freeMagnitudeToAdd;
        }
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
            // 1. calculate slashed magnitude from current allocation
            // update current and all following queued magnitude updates for (operator, strategy, operatorSetId) tuple
            uint64 slashedMagnitude;
            {
                uint64 currentMagnitude = uint64(
                    _magnitudeUpdate[operator][strategies[i]][msg.sender][operatorSetId].upperLookupRecent(
                        uint32(block.timestamp)
                    )
                );
                /// TODO: add wrapping library for rounding up for slashing accounting
                slashedMagnitude = uint64(uint256(bipsToSlash) * uint256(currentMagnitude) / BIPS_FACTOR);

                _magnitudeUpdate[operator][strategies[i]][msg.sender][operatorSetId].decrementAtAndFutureCheckpoints({
                    key: uint32(block.timestamp),
                    decrementValue: slashedMagnitude
                });
            }

            // 2. if there are any pending deallocations then need to update and decrement if they fall within slashable window
            {
                uint256 queuedDeallocationsLen =
                    _queuedDeallocations[operator][strategies[i]][msg.sender][operatorSetId].length;
                for (uint256 j = queuedDeallocationsLen; j > 0; --j) {
                    QueuedDeallocation storage queuedDeallocation =
                        _queuedDeallocations[operator][strategies[i]][msg.sender][operatorSetId][j - 1];

                    // if queued deallocation is still within slashable window, then slash and keep track of sum to decrement from totalMagnitude
                    if (uint32(block.timestamp) + ALLOCATION_DELAY > queuedDeallocation.completableTimestamp) {
                        uint64 slashedAmount =
                            uint64(uint256(bipsToSlash) * uint256(queuedDeallocation.magnitudeDiff) / BIPS_FACTOR);
                        queuedDeallocation.magnitudeDiff -= slashedAmount;
                        slashedMagnitude += slashedAmount;
                    } else {
                        break;
                    }
                }
            }

            // 3. update totalMagnitude, get total magnitude and subtract slashedMagnitude
            _totalMagnitudeUpdate[operator][strategies[i]].push({
                key: uint32(block.timestamp),
                value: _totalMagnitudeUpdate[operator][strategies[i]].latest() - slashedMagnitude
            });
        }
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
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Helper function used by migration & registration functions to register an operator to operator sets.
     * @param avs The AVS that the operator is registering to.
     * @param operator The operator to register.
     * @param operatorSetIds The IDs of the operator sets.
     */
    function _registerToOperatorSets(address operator, address avs, uint32[] calldata operatorSetIds) internal {
        // Loop over `operatorSetIds` array and register `operator` for each item.
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            OperatorSet memory operatorSet = OperatorSet(avs, operatorSetIds[i]);

            require(
                isOperatorSet[avs][operatorSetIds[i]],
                "AVSDirectory._registerOperatorToOperatorSets: invalid operator set"
            );

            require(
                !isMember(operator, operatorSet),
                "AVSDirectory._registerOperatorToOperatorSets: operator already registered to operator set"
            );

            ++operatorSetMemberCount[avs][operatorSetIds[i]];

            _operatorSetsMemberOf[operator].add(_encodeOperatorSet(operatorSet));

            OperatorSetRegistrationStatus storage registrationStatus =
                operatorSetStatus[avs][operator][operatorSetIds[i]];
            require(
                !registrationStatus.registered,
                "AVSDirectory._registerOperatorToOperatorSets: operator already registered for operator set"
            );
            registrationStatus.registered = true;

            emit OperatorAddedToOperatorSet(operator, operatorSet);
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
            OperatorSet memory operatorSet = OperatorSet(avs, operatorSetIds[i]);

            require(
                isMember(operator, operatorSet),
                "AVSDirectory._deregisterOperatorFromOperatorSet: operator not registered for operator set"
            );

            --operatorSetMemberCount[avs][operatorSetIds[i]];

            _operatorSetsMemberOf[operator].remove(_encodeOperatorSet(operatorSet));

            emit OperatorRemovedFromOperatorSet(operator, operatorSet);
        }
    }

    /**
     * @notice For a single strategy, allocate magnitude to a list of operator sets and then update freeMagnitude
     * for the strategy after all allocations
     * @param operator address to allocate magnitude for
     * @param allocation magnitude adjustment for a single strategy and multiple operator sets
     * @param effectTimestamp timestamp when the allocation will take effect
     */
    function _allocate(address operator, MagnitudeAdjustment calldata allocation, uint32 effectTimestamp) internal {
        IStrategy strategy = allocation.strategy;
        OperatorSet[] calldata operatorSets = allocation.operatorSets;
        uint64 freeAllocatableMagnitude = freeMagnitude[operator][strategy];

        for (uint256 i = 0; i < operatorSets.length; ++i) {
            // 1. check freeMagnitude available, that is the allocation of stake is backed and not slashable
            require(
                allocation.magnitudeDiffs[i] > 0,
                "AVSDirectory.queueAllocations: magnitudeDiff must be greater than 0"
            );
            require(
                freeAllocatableMagnitude >= allocation.magnitudeDiffs[i],
                "AVSDirectory.queueAllocations: insufficient available free magnitude to allocate"
            );
            // 2. rate limiting of 1 pending allocation at a time
            // read number of checkpoints after current timestamp
            (uint224 value, uint256 pos) = _magnitudeUpdate[operator][strategy][operatorSets[i].avs][operatorSets[i]
                .operatorSetId].upperLookupRecentWithPos(uint32(block.timestamp));
            // no checkpoint exists if value == 0 && pos == 0, so check the negation before checking if there is a
            // a pending allocation
            if (value != 0 || pos != 0) {
                require(
                    pos + MAX_PENDING_UPDATES
                        == _magnitudeUpdate[operator][strategy][operatorSets[i].avs][operatorSets[i].operatorSetId].length(),
                    "AVSDirectory.queueAllocations: exceed max pending allocations allowed for op, opSet, strategy"
                );
            }

            // 3. allocate magnitude which will take effect in the future 21 days from now
            _magnitudeUpdate[operator][strategy][operatorSets[i].avs][operatorSets[i].operatorSetId].push({
                key: effectTimestamp,
                value: value + uint224(allocation.magnitudeDiffs[i])
            });
            // 4. keep track of available freeMagnitude to update later
            freeAllocatableMagnitude -= allocation.magnitudeDiffs[i];
        }

        // 5. update freeMagnitude after all allocations
        freeMagnitude[operator][strategy] = freeAllocatableMagnitude;
    }

    /**
     * @notice For a single strategy, queue deallocate magnitude for a list of operator sets
     * @param operator address to deallocate magnitude for
     * @param deallocation magnitude adjustment for a single strategy and multiple operator sets
     * @param completableTimestamp timestamp when the queued deallocation can be completed and freeMagnitude updated
     */
    function _queueDeallocate(
        address operator,
        MagnitudeAdjustment calldata deallocation,
        uint32 completableTimestamp
    ) internal {
        IStrategy strategy = deallocation.strategy;
        require(
            deallocation.operatorSets.length == deallocation.magnitudeDiffs.length,
            "AVSDirectory._queueDeallocate: operatorSets and magnitudeDiffs length mismatch"
        );
        OperatorSet[] calldata operatorSets = deallocation.operatorSets;

        // deallocate from each operator set for this strategy
        for (uint256 i = 0; i < operatorSets.length; ++i) {
            // 1. ensure deallocation doesn't exceed current allocation
            uint64 currentMagnitude = uint64(
                _magnitudeUpdate[operator][strategy][operatorSets[i].avs][operatorSets[i].operatorSetId]
                    .upperLookupRecent(uint32(block.timestamp))
            );
            require(
                deallocation.magnitudeDiffs[i] > 0,
                "AVSDirectory._queueDeallocate: magnitudeDiff must be greater than 0"
            );
            require(
                deallocation.magnitudeDiffs[i] <= currentMagnitude,
                "AVSDirectory._queueDeallocate: cannot deallocate more than what is allocated"
            );

            // 2. update and decrement current and future queued amounts in case any pending allocations exist
            _magnitudeUpdate[operator][strategy][operatorSets[i].avs][operatorSets[i].operatorSetId]
                .decrementAtAndFutureCheckpoints({
                key: uint32(block.timestamp),
                decrementValue: deallocation.magnitudeDiffs[i]
            });

            // 3. ensure only queued deallocation per operator, operatorSet, strategy
            _checkPendingDeallocations(operator, strategy, operatorSets[i]);

            // 4. queue deallocation for (op, opSet, strategy) for magnitudeDiff amount
            _queuedDeallocations[operator][strategy][operatorSets[i].avs][operatorSets[i].operatorSetId].push(
                QueuedDeallocation({
                    magnitudeDiff: deallocation.magnitudeDiffs[i],
                    completableTimestamp: completableTimestamp
                })
            );
        }
    }

    /**
     * @notice Completes queued deallocations for a single strategy and multiple operator sets
     * @param operator address to complete deallocations for
     * @param strategy the strategy to complete deallocations for
     * @param operatorSet the operator set to complete deallocations for the given strategy
     */
    function _completeDeallocation(
        address operator,
        IStrategy strategy,
        OperatorSet calldata operatorSet
    ) internal returns (uint64 freeMagnitudeToAdd) {
        QueuedDeallocation[] memory queuedDeallocation =
            _queuedDeallocations[operator][strategy][operatorSet.avs][operatorSet.operatorSetId];
        // read last completed deallocation index in the queuedDeallocations array for the (op, opSet, Strategy) tuple
        uint256 i = _nextDeallocationIndex[operator][strategy][operatorSet.avs][operatorSet.operatorSetId];

        for (; i < queuedDeallocation.length; ++i) {
            // queued deallocations ordered by timestamp, if completableTimestamp is greater than completeUntilTimestamp, break
            if (queuedDeallocation[i].completableTimestamp > uint32(block.timestamp)) {
                break;
            }

            freeMagnitudeToAdd += queuedDeallocation[i].magnitudeDiff;
        }
        _nextDeallocationIndex[operator][strategy][operatorSet.avs][operatorSet.operatorSetId] = i;
        return freeMagnitudeToAdd;
    }

    /// @dev Check for max number of pending deallocations, ensuring <= MAX_PENDING_UPDATES
    function _checkPendingDeallocations(
        address operator,
        IStrategy strategy,
        OperatorSet calldata operatorSet
    ) internal view returns (uint64 freeMagnitudeToAdd) {
        QueuedDeallocation[] memory queuedDeallocations =
            _queuedDeallocations[operator][strategy][operatorSet.avs][operatorSet.operatorSetId];
        uint256 length = queuedDeallocations.length;

        for (uint256 i = length; i > 0; --i) {
            // If completableTimestamp is greater than completeUntilTimestamp, break
            if (queuedDeallocations[i - 1].completableTimestamp < uint32(block.timestamp)) {
                require(
                    length - i  + 1 < MAX_PENDING_UPDATES,
                    "AVSDirectory._checkPendingDeallocations: exceeds max pending deallocations"
                );
            } else {
                break;
            }
        }
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @notice Returns operator sets an operator is registered to in the order they were registered.
    /// @param operator The operator address to query.
    /// @param index The index of the enumerated list of operator sets.
    function operatorSetsMemberOf(address operator, uint256 index) public view returns (OperatorSet memory) {
        return _decodeOperatorSet(_operatorSetsMemberOf[operator].at(index));
    }

    /// @notice Returns an array of operator sets an operator is registered to.
    /// @param operator The operator address to query.
    /// @param start The starting index of the array to query.
    /// @param length The amount of items of the array to return.
    function operatorSetsMemberOf(
        address operator,
        uint256 start,
        uint256 length
    ) public view returns (OperatorSet[] memory operatorSets) {
        uint256 maxLength = _operatorSetsMemberOf[operator].length() - start;
        if (length > maxLength) length = maxLength;
        operatorSets = new OperatorSet[](length);
        for (uint256 i; i < length; ++i) {
            operatorSets[i] = _decodeOperatorSet(_operatorSetsMemberOf[operator].at(start + i));
        }
    }

    /// @notice Returns the total number of operator sets an operator is registered to.
    /// @param operator The operator address to query.
    function inTotalOperatorSets(address operator) public view returns (uint256) {
        return _operatorSetsMemberOf[operator].length();
    }

    /// @notice Returns whether or not an operator is registered to an operator set.
    /// @param operator The operator address to query.
    /// @param operatorSet The `OperatorSet` to query.
    function isMember(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        return _operatorSetsMemberOf[operator].contains(_encodeOperatorSet(operatorSet));
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
        return isMember(operator, operatorSet)
            || registrationStatus.lastDeregisteredTimestamp + ALLOCATION_DELAY >= block.timestamp;
    }

    // /**
    //  * @notice fetches the minimum slashable shares for a certain operator and operatorSet for a list of strategies
    //  * from the current timestamp until the given timestamp
    //  *
    //  * @param operator the operator to get the minimum slashable shares for
    //  * @param operatorSet the operatorSet to get the minimum slashable shares for
    //  * @param strategies the strategies to get the minimum slashable shares for
    //  * @param timestamp the timestamp to the minimum slashable shares before
    //  *
    //  * @dev used to get the slashable stakes of operators to weigh over a given slashability window
    //  *
    //  * @return the list of share amounts for each strategy
    //  */
    // function getMinimumSlashableSharesBefore(
    //     address operator,
    //     OperatorSet calldata operatorSet,
    //     IStrategy[] calldata strategies,
    //     uint32 timestamp
    // ) external view returns (uint256[] calldata) {}

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

    /// @notice Returns an EIP-712 encoded hash struct.
    function _calculateDigestHash(bytes32 structHash) internal view returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", _calculateDomainSeparator(), structHash));
    }

    /// @dev Returns an `OperatorSet` encoded into a 32-byte value.
    /// @param operatorSet The `OperatorSet` to encode.
    function _encodeOperatorSet(OperatorSet memory operatorSet) internal pure returns (bytes32) {
        return bytes32(abi.encodePacked(operatorSet.avs, uint96(operatorSet.operatorSetId)));
    }

    /// @dev Returns an `OperatorSet` decoded from an encoded 32-byte value.
    /// @param encoded The encoded `OperatorSet` to decode.
    /// @dev Assumes `encoded` is encoded via `_encodeOperatorSet(operatorSet)`.
    function _decodeOperatorSet(bytes32 encoded) internal pure returns (OperatorSet memory) {
        return OperatorSet({
            avs: address(uint160(uint256(encoded) >> 96)),
            operatorSetId: uint32(uint256(encoded) & type(uint96).max)
        });
    }
}
