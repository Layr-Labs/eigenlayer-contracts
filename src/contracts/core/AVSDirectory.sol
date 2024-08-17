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
    uint32 public constant DEFAULT_ALLOCATION_DELAY = 21 days;

    /// @dev Delay before allocations take effect and how long until deallocations are completable
    uint32 public constant DEALLOCATION_DELAY = 17.5 days;


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
     * @notice Modifies the propotions of slashable stake allocated to a list of operatorSets for a set of strategies
     * @param operator address to modify allocations for
     * @param allocations array of magnitude adjustments for multiple strategies and corresponding operator sets
     * @param operatorSignature signature of the operator if msg.sender is not the operator
     * @dev updates freeMagnitude for the updated strategies
     * @dev must be called by the operator
     */
    function modifyAllocations(
        address operator,
        MagnitudeAllocation[] calldata allocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external {
        // TODO: allow signature on behalf of operator
        require(msg.sender == operator, "AVSDirectory.modifyAllocations: only operator can modify allocations");
        // completable timestamp for deallocations
        uint32 completableTimestamp = uint32(block.timestamp) + DEALLOCATION_DELAY;
        // effect timestamp for allocations to take effect. This is configurable by operators
        uint32 effectTimestamp = uint32(block.timestamp) + getAllocationDelay(operator);
        for (uint256 i = 0; i < allocations.length; ++i) {
            // 1. For the given (operator,strategy) clear all pending free magnitude for the strategy and update freeMagnitude
            // numToComplete = 0 defaults to completing all pending deallocations up to uint8 max (256)
            _updateFreeMagnitude({
                operator: operator, 
                strategy: allocations[i].strategy,
                numToComplete: 0
            });

            // 2. check current totalMagnitude matches expected value
            uint64 currentTotalMagnitude = uint64(_totalMagnitudeUpdate[operator][allocations[i].strategy].latest());
            require(
                currentTotalMagnitude == allocations[i].expectedTotalMagnitude,
                "AVSDirectory.setAllocations: current total magnitude does not match expected"
            );

            // 3. set allocations for the strategy after updating freeMagnitude
            _modifyAllocations({
                operator: operator, 
                allocation: allocations[i],
                allocationEffectTimestamp: effectTimestamp,
                deallocationCompletableTimestamp: completableTimestamp
            }); 
        }
    }

    function _modifyAllocations(
        address operator,
        MagnitudeAllocation calldata allocation,
        uint32 allocationEffectTimestamp,
        uint32 deallocationCompletableTimestamp
    ) internal {
        uint64 newFreeMagnitude = freeMagnitude[operator][allocation.strategy];
        OperatorSet[] calldata opSets = allocation.operatorSets;

        for (uint256 i = 0; i < opSets.length; ++i) {
            /// TODO??? ensure ordered in ascending bytes32 hash of opSets, can we have duplicate opSets, duplicate strategies?

            // Read current magnitude allocation including its respective array index and length.
            // We'll use these values later to check the number of pending allocations/deallocations.
            (
                uint224 currentMagnitude,
                uint256 pos,
                uint256 length
            ) = _magnitudeUpdate[operator][allocation.strategy][opSets[i].avs][opSets[i].operatorSetId]
                .upperLookupRecentWithPos(uint32(block.timestamp));

            if (allocation.magnitudes[i] < uint64(currentMagnitude)) {
                // Newly configured magnitude is less than current value. 
                // Therefore we handle this as a deallocation

                // 1. ensure only pending queued deallocation per operator, operatorSet, strategy
                _checkQueuedDeallocations(operator, allocation.strategy, opSets[i]);
 
                // 2. update and decrement current and future queued amounts in case any pending allocations exist
                _magnitudeUpdate[operator][allocation.strategy][opSets[i].avs][opSets[i].operatorSetId]
                    .decrementAtAndFutureCheckpoints({
                        key: uint32(block.timestamp),
                        decrementValue: uint64(currentMagnitude) - allocation.magnitudes[i]
                    });
                
                // 3. push PendingFreeMagnitude and respective array index into (op,opSet,Strategy) queued deallocations
                uint256 index = _pendingFreeMagnitude[operator][allocation.strategy].length;
                _pendingFreeMagnitude[operator][allocation.strategy].push(
                    PendingFreeMagnitude({
                        magnitudeDiff: uint64(currentMagnitude) - allocation.magnitudes[i],
                        completableTimestamp: deallocationCompletableTimestamp
                    })
                );
                _queuedDeallocationIndices[operator][allocation.strategy][opSets[i].avs][opSets[i].operatorSetId].push(
                    index
                );                
            } else if (allocation.magnitudes[i] > uint64(currentMagnitude)) {
                // Newly configured magnitude is greater than current value. 
                // Therefore we handle this as an allocation

                // 1. ensure only 1 pending allocation at a time
                // read number of checkpoints after current timestamp
                // no checkpoint exists if value == 0 && pos == 0, so check the negation before checking if there is a
                // a pending allocation
                if (currentMagnitude != 0 || pos != 0) {
                    require(
                        pos + MAX_PENDING_UPDATES <= length,
                        "AVSDirectory.queueAllocations: exceed max pending allocations allowed for op, opSet, strategy"
                    );
                }
                // 2. allocate magnitude which will take effect in the future 21 days from now
                _magnitudeUpdate[operator][allocation.strategy][opSets[i].avs][opSets[i].operatorSetId].push({
                    key: allocationEffectTimestamp,
                    value: allocation.magnitudes[i]
                });
                // 3. decrement free magnitude by incremented amount
                require(
                    newFreeMagnitude >= allocation.magnitudes[i] - uint64(currentMagnitude),
                    "AVSDirectory._setAllocations: insufficient available free magnitude to allocate"
                );
                newFreeMagnitude -= allocation.magnitudes[i] - uint64(currentMagnitude);
            }
        }

        // update freeMagnitude after all allocations.
        // if provided allocations only resulted in deallocating, then this value would be unchanged
        freeMagnitude[operator][allocation.strategy] = newFreeMagnitude;
    }

    /**
     * @notice For all pending deallocations that have become completable, their pending free magnitude can be
     * added back to the free magnitude of the (operator, strategy) amount. This function takes a list of strategies
     * and adds all completable deallocations for each strategy, updating the freeMagnitudes of the operator
     *
     * @param operator address to complete deallocations for
     * @param strategies a list of strategies to complete deallocations for
     * @param numToComplete a list of number of pending free magnitude deallocations to complete for each strategy
     *
     * @dev can be called permissionlessly by anyone
     */
    function updateFreeMagnitude(
        address operator,
        IStrategy[] calldata strategies,
        uint8[] calldata numToComplete
    ) external {
        for (uint256 i = 0; i < strategies.length; ++i) {
            _updateFreeMagnitude({
                operator: operator,
                strategy: strategies[i],
                numToComplete: numToComplete[i]
            });
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
            // loop backwards through _queuedDeallocationIndices, each element contains an array index to
            // corresponding deallocation to access in pendingFreeMagnitude
            // if completable, then break
            //      (since ordered by completableTimestamps, older deallocations will also be completable and outside slashable window)
            // if NOT completable, then slash
            {
                uint256 queuedDeallocationIndicesLen =
                    _queuedDeallocationIndices[operator][strategies[i]][msg.sender][operatorSetId].length;
                for (uint256 j = queuedDeallocationIndicesLen; j > 0; --j) {
                    // index of pendingFreeMagnitude/deallocation to check for slashing
                    uint256 index =
                        _queuedDeallocationIndices[operator][strategies[i]][msg.sender][operatorSetId][j - 1];
                    PendingFreeMagnitude storage pendingFreeMagnitude =
                        _pendingFreeMagnitude[operator][strategies[i]][index];

                    // Reached pendingFreeMagnitude/deallocation that is completable and not within slashability window,
                    // therefore older deallocations will also be completable. Since this is ordered by completableTimestamps break loop now
                    if (pendingFreeMagnitude.completableTimestamp >= uint32(block.timestamp)) {
                        break;
                    }

                    // pending deallocation is still within slashable window, slash magnitudeDiff and add to slashedMagnitude
                    uint64 slashedAmount =
                        uint64(uint256(bipsToSlash) * uint256(pendingFreeMagnitude.magnitudeDiff) / BIPS_FACTOR);
                    pendingFreeMagnitude.magnitudeDiff -= slashedAmount;
                    slashedMagnitude += slashedAmount;
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
     * @notice Called by operators to set their allocation delay one time
     * @param operator address to set allocation delay for
     * @param delay the allocation delay in seconds
     * @dev this is expected to be updatable in a future release
     */
    function initializeAllocationDelay(
        address operator,
        uint32 delay
    ) external {
        require(
            msg.sender == operator,
            "AVSDirectory.initializeAllocationDelay: only operator can set allocation delay"
        );
        require(
            delegation.isOperator(operator),
            "AVSDirectory.initializeAllocationDelay: operator not registered to EigenLayer yet"
        );
        require(
            !allocationDelay[operator].isSet,
            "AVSDirectory.initializeAllocationDelay: allocation delay already set"
        );
        allocationDelay[operator] = AllocationDelayDetails({
            isSet: true,
            allocationDelay: delay
        });
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

    /// @dev read through pending free magnitudes and add to freeMagnitude if completableTimestamp is >= block timestamp
    /// In additiona to updating freeMagnitude, updates next starting index to read from for pending free magnitudes after completing

    /**
     * @notice For a single strategy, update freeMagnitude by adding completable pending free magnitudes
     * @param operator address to update freeMagnitude for
     * @param strategy the strategy to update freeMagnitude for
     * @param numToComplete the number of pending free magnitudes deallocations to complete, 0 to complete all (uint8 max 256)
     * @dev read through pending free magnitudes and add to freeMagnitude if completableTimestamp is >= block timestamp
     * In addition to updating freeMagnitude, updates next starting index to read from for pending free magnitudes after completing
     */
    function _updateFreeMagnitude(address operator, IStrategy strategy, uint8 numToComplete) internal {
        uint256 nextIndex = _nextPendingFreeMagnitudeIndex[operator][strategy];
        uint256 pendingFreeMagnitudeLength = _pendingFreeMagnitude[operator][strategy].length;
        uint8 completed = numToComplete == 0 ? type(uint8).max : 0;
        while (nextIndex < pendingFreeMagnitudeLength && completed < numToComplete) {
            PendingFreeMagnitude memory pendingFreeMagnitude = _pendingFreeMagnitude[operator][strategy][nextIndex];
            // pendingFreeMagnitude is ordered by completableTimestamp. If we reach one that is not completable yet, then break
            // loop until completableTimestamp is < block.timestamp
            if (pendingFreeMagnitude.completableTimestamp < uint32(block.timestamp)) {
                break;
            }

            // pending free magnitude can be added to freeMagnitude
            freeMagnitude[operator][strategy] += pendingFreeMagnitude.magnitudeDiff;
            ++nextIndex;
            ++completed;
        }
        // update next pending free magnitude index to start from after adding all completable magnitudes
        _nextPendingFreeMagnitudeIndex[operator][strategy] = nextIndex;
    }

    /// @dev Check for max number of pending queued deallocations, ensuring <= MAX_PENDING_UPDATES
    function _checkQueuedDeallocations(
        address operator,
        IStrategy strategy,
        OperatorSet calldata operatorSet
    ) internal view {
        uint256 length =
            _queuedDeallocationIndices[operator][strategy][operatorSet.avs][operatorSet.operatorSetId].length;

        for (uint256 i = length; i > 0; --i) {
            // index of pendingFreeMagnitude/deallocation to check for slashing
            uint256 index =
                _queuedDeallocationIndices[operator][strategy][operatorSet.avs][operatorSet.operatorSetId][i - 1];
            PendingFreeMagnitude memory pendingFreeMagnitude = _pendingFreeMagnitude[operator][strategy][index];

            // If completableTimestamp is greater than completeUntilTimestamp, break
            if (pendingFreeMagnitude.completableTimestamp < uint32(block.timestamp)) {
                require(
                    length - i + 1 < MAX_PENDING_UPDATES,
                    "AVSDirectory._checkQueuedDeallocations: exceeds max pending deallocations"
                );
            } else {
                break;
            }
        }
    }

    // /// @dev Verify allocator's signature and spend salt
    // function _verifyAllocatorSignature(
    //     address allocator,
    //     address operator,
    //     MagnitudeAdjustment[] calldata magnitudeAdjustments,
    //     SignatureWithSaltAndExpiry calldata allocatorSignature
    // ) internal {
    //     // check the signature expiry
    //     require(
    //         allocatorSignature.expiry >= block.timestamp,
    //         "AVSDirectory._verifyAllocatorSignature: allocator signature expired"
    //     );
    //     // Assert allocator's signature cannot be replayed.
    //     require(
    //         !allocatorSaltIsSpent[allocator][allocatorSignature.salt],
    //         "AVSDirectory._verifyAllocatorSignature: salt spent"
    //     );

    //     bytes32 digestHash = calculateMagnitudeAdjustmentDigestHash(
    //         operator, magnitudeAdjustments, allocatorSignature.salt, allocatorSignature.expiry
    //     );

    //     // Assert allocator's signature is valid.
    //     EIP1271SignatureUtils.checkSignature_EIP1271(allocator, digestHash, allocatorSignature.signature);
    //     // Spend salt.
    //     allocatorSaltIsSpent[allocator][allocatorSignature.salt] = true;
    // }

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
     * @param linear whether the search should be linear (from the most recent) or binary
     *
     * @return slashableBips the slashable bips of the given strategy owned by
     * the given OperatorSet for the given operator and timestamp
     */
    function getSlashableBips(
        address operator,
        OperatorSet calldata operatorSet,
        IStrategy strategy,
        uint32 timestamp,
        bool linear
    ) public view returns (uint16) {
        uint64 totalMagnitude;
        if (linear) {
            totalMagnitude = uint64(_totalMagnitudeUpdate[operator][strategy].upperLookupLinear(timestamp));
        } else {
            totalMagnitude = uint64(_totalMagnitudeUpdate[operator][strategy].upperLookup(timestamp));
        }
        // return early if totalMagnitude is 0
        if (totalMagnitude == 0) {
            return 0;
        }

        uint64 currentMagnitude;
        if (linear) {
            currentMagnitude = uint64(
                _magnitudeUpdate[operator][strategy][operatorSet.avs][operatorSet.operatorSetId].upperLookupLinear(
                    timestamp
                )
            );
        } else {
            currentMagnitude = uint64(
                _magnitudeUpdate[operator][strategy][operatorSet.avs][operatorSet.operatorSetId].upperLookup(timestamp)
            );
        }

        return uint16(currentMagnitude * BIPS_FACTOR / totalMagnitude);
    }

    function getAllocationDelay(address operator) public view returns (uint32) {
        AllocationDelayDetails memory details = allocationDelay[operator];
        return details.isSet ? details.allocationDelay : DEFAULT_ALLOCATION_DELAY;
    }

    /// @notice operator is slashable by operatorSet if currently registered OR last deregistered within 21 days
    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        OperatorSetRegistrationStatus memory registrationStatus =
            operatorSetStatus[operatorSet.avs][operator][operatorSet.operatorSetId];
        return isMember(operator, operatorSet)
            || registrationStatus.lastDeregisteredTimestamp + DEALLOCATION_DELAY >= block.timestamp;
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

    // /**
    //  * @notice Calculates the digest hash to be signed by an allocator to allocate or deallocate magnitude for an operator
    //  * @param operator The operator to allocate or deallocate magnitude for.
    //  * @param adjustments The magnitude allocations/deallocations to be made.
    //  * @param salt A unique and single use value associated with the approver signature.
    //  * @param expiry Time after which the approver's signature becomes invalid.
    //  */
    // function calculateMagnitudeAdjustmentDigestHash(
    //     address operator,
    //     MagnitudeAdjustment[] calldata adjustments,
    //     bytes32 salt,
    //     uint256 expiry
    // ) public view returns (bytes32) {
    //     return _calculateDigestHash(
    //         keccak256(abi.encode(MAGNITUDE_ADJUSTMENT_TYPEHASH, operator, adjustments, salt, expiry))
    //     );
    // }

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
