// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../permissions/Pausable.sol";
import "../libraries/EIP1271SignatureUtils.sol";
import "../libraries/ShareScalingLib.sol";
import "./AVSDirectoryStorage.sol";

contract AVSDirectory is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AVSDirectoryStorage,
    ReentrancyGuardUpgradeable
{
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.AddressSet;
    using Checkpoints for Checkpoints.History;

    /// @dev Index for flag that pauses operator register/deregister to avs when set.
    uint8 internal constant PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS = 0;
    /// @dev Index for flag that pauses operator register/deregister to operator sets when set.
    uint8 internal constant PAUSER_OPERATOR_REGISTER_DEREGISTER_TO_OPERATOR_SETS = 1;

    /// @dev BIPS factor for slashable bips
    uint256 internal constant BIPS_FACTOR = 10_000;

    /// @dev Delay before deallocations are completable and can be added back into freeMagnitude
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
    constructor(
        IDelegationManager _delegation
    ) AVSDirectoryStorage(_delegation) {
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
    function createOperatorSets(
        uint32[] calldata operatorSetIds
    ) external {
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
        // // Assert that the AVS is an operator set AVS.
        // require(
        //     isOperatorSetAVS[msg.sender], "AVSDirectory.migrateOperatorsToOperatorSets: AVS is not an operator set AVS"
        // );

        // for (uint256 i = 0; i < operators.length; i++) {
        //     // Assert that the operator is registered & has not been migrated.
        //     require(
        //         avsOperatorStatus[msg.sender][operators[i]] == OperatorAVSRegistrationStatus.REGISTERED,
        //         "AVSDirectory.migrateOperatorsToOperatorSets: operator already migrated or not a legacy registered operator"
        //     );

        //     // Migrate operator to operator sets.
        //     _registerToOperatorSets(operators[i], msg.sender, operatorSetIds[i]);

        //     // Deregister operator from AVS - this prevents the operator from being migrated again since
        //     // the AVS can no longer use the legacy M2 registration path
        //     avsOperatorStatus[msg.sender][operators[i]] = OperatorAVSRegistrationStatus.UNREGISTERED;
        //     emit OperatorAVSRegistrationStatusUpdated(
        //         operators[i], msg.sender, OperatorAVSRegistrationStatus.UNREGISTERED
        //     );
        //     emit OperatorMigratedToOperatorSets(operators[i], msg.sender, operatorSetIds[i]);
        // }
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
        // COMMENTED FOR CODESIZE
        // if (operatorSignature.signature.length == 0) {
        //     require(msg.sender == operator, "AVSDirectory.forceDeregisterFromOperatorSets: caller must be operator");
        // } else {
        //     // Assert operator's signature has not expired.
        //     require(
        //         operatorSignature.expiry >= block.timestamp,
        //         "AVSDirectory.forceDeregisterFromOperatorSets: operator signature expired"
        //     );
        //     // Assert operator's signature `salt` has not already been spent.
        //     require(
        //         !operatorSaltIsSpent[operator][operatorSignature.salt],
        //         "AVSDirectory.forceDeregisterFromOperatorSets: salt already spent"
        //     );

        //     // Assert that `operatorSignature.signature` is a valid signature for operator set deregistrations.
        //     EIP1271SignatureUtils.checkSignature_EIP1271(
        //         operator,
        //         calculateOperatorSetForceDeregistrationTypehash({
        //             avs: avs,
        //             operatorSetIds: operatorSetIds,
        //             salt: operatorSignature.salt,
        //             expiry: operatorSignature.expiry
        //         }),
        //         operatorSignature.signature
        //     );

        //     // Mutate `operatorSaltIsSpent` to `true` to prevent future respending.
        //     operatorSaltIsSpent[operator][operatorSignature.salt] = true;
        // }
        // _deregisterFromOperatorSets(avs, operator, operatorSetIds);
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
     * @dev Updates freeMagnitude for the updated strategies
     * @dev Must be called by the operator or with a valid operator signature
     * @dev For each allocation, allocation.operatorSets MUST be ordered in ascending order according to the
     * encoding of the operatorSet. This is to prevent duplicate operatorSets being passed in. The easiest way to ensure
     * ordering is to sort allocated operatorSets by address first, and then sort for each avs by ascending operatorSetIds.
     */
    function modifyAllocations(
        address operator,
        MagnitudeAllocation[] calldata allocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external {
        if (msg.sender != operator) {
            _verifyOperatorSignature(operator, allocations, operatorSignature);
        }
        require(
            delegation.isOperator(operator), "AVSDirectory.modifyAllocations: operator not registered to EigenLayer yet"
        );
        IDelegationManager.AllocationDelayDetails memory details = delegation.operatorAllocationDelay(operator);
        require(
            details.isSet,
            "AVSDirectory.modifyAllocations: operator must initialize allocation delay before modifying allocations"
        );
        // effect timestamp for allocations to take effect. This is configurable by operators
        uint32 effectTimestamp = uint32(block.timestamp) + details.allocationDelay;
        // completable timestamp for deallocations
        uint32 completableTimestamp = uint32(block.timestamp) + DEALLOCATION_DELAY;

        for (uint256 i = 0; i < allocations.length; ++i) {
            // 1. For the given (operator,strategy) clear all pending free magnitude for the strategy and update freeMagnitude
            _updateFreeMagnitude({
                operator: operator,
                strategy: allocations[i].strategy,
                numToComplete: type(uint16).max
            });

            // 2. Check current totalMagnitude matches expected value. This is to check for slashing race conditions
            // where an operator gets slashed from an operatorSet and as a result all the configured allocations have larger
            // proprtional magnitudes relative to eachother. This check prevents any surprising behavior.
            uint64 currentTotalMagnitude = _getLatestTotalMagnitude(operator, allocations[i].strategy);
            require(
                currentTotalMagnitude == allocations[i].expectedTotalMagnitude,
                "AVSDirectory.modifyAllocations: current total magnitude does not match expected"
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
        uint16[] calldata numToComplete
    ) external {
        require(
            delegation.isOperator(operator),
            "AVSDirectory.updateFreeMagnitude: operator not registered to EigenLayer yet"
        );
        require(strategies.length == numToComplete.length, "AVSDirectory.updateFreeMagnitude: array length mismatch");
        for (uint256 i = 0; i < strategies.length; ++i) {
            _updateFreeMagnitude({operator: operator, strategy: strategies[i], numToComplete: numToComplete[i]});
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
        // OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        // bytes32 operatorSetKey = _encodeOperatorSet(operatorSet);
        // require(
        //     isOperatorSlashable(operator, operatorSet),
        //     "AVSDirectory.slashOperator: operator not slashable for operatorSet"
        // );

        // for (uint256 i = 0; i < strategies.length; ++i) {
        //     // 1. calculate slashed magnitude from current allocation
        //     // update current and all following queued magnitude updates for (operator, strategy, operatorSetId) tuple
        //     uint64 slashedMagnitude;
        //     {
        //         uint64 currentMagnitude = uint64(
        //             _magnitudeUpdate[operator][strategies[i]][operatorSetKey].upperLookupRecent(uint32(block.timestamp))
        //         );
        //         // TODO: if we don't continue here we get into weird "total/free magnitude" not initialized cases. Is this ok?
        //         if (currentMagnitude == 0) {
        //             continue;
        //         }

        //         /// TODO: add wrapping library for rounding up for slashing accounting
        //         slashedMagnitude = uint64(uint256(bipsToSlash) * uint256(currentMagnitude) / BIPS_FACTOR);

        //         _magnitudeUpdate[operator][strategies[i]][operatorSetKey].decrementAtAndFutureCheckpoints({
        //             key: uint32(block.timestamp),
        //             decrementValue: slashedMagnitude
        //         });
        //     }

        //     // 2. if there are any pending deallocations then need to update and decrement if they fall within slashable window
        //     // loop backwards through _queuedDeallocationIndices, each element contains an array index to
        //     // corresponding deallocation to access in pendingFreeMagnitude
        //     // if completable, then break
        //     //      (since ordered by completableTimestamps, older deallocations will also be completable and outside slashable window)
        //     // if NOT completable, then add to slashed magnitude
        //     {
        //         uint256 queuedDeallocationIndicesLen =
        //             _queuedDeallocationIndices[operator][strategies[i]][operatorSetKey].length;
        //         for (uint256 j = queuedDeallocationIndicesLen; j > 0; --j) {
        //             // index of pendingFreeMagnitude/deallocation to check for slashing
        //             uint256 index = _queuedDeallocationIndices[operator][strategies[i]][operatorSetKey][j - 1];
        //             PendingFreeMagnitude storage pendingFreeMagnitude =
        //                 _pendingFreeMagnitude[operator][strategies[i]][index];

        //             // Reached pendingFreeMagnitude/deallocation that is completable and not within slashability window,
        //             // therefore older deallocations will also be completable. Since this is ordered by completableTimestamps break loop now
        //             if (pendingFreeMagnitude.completableTimestamp >= uint32(block.timestamp)) {
        //                 break;
        //             }

        //             // pending deallocation is still within slashable window, slash magnitudeDiff and add to slashedMagnitude
        //             uint64 slashedAmount =
        //                 uint64(uint256(bipsToSlash) * uint256(pendingFreeMagnitude.magnitudeDiff) / BIPS_FACTOR);
        //             pendingFreeMagnitude.magnitudeDiff -= slashedAmount;
        //             slashedMagnitude += slashedAmount;
        //         }
        //     }

        //     // 3. update totalMagnitude, get total magnitude and subtract slashedMagnitude
        //     _totalMagnitudeUpdate[operator][strategies[i]].push({
        //         key: uint32(block.timestamp),
        //         value: _getLatestTotalMagnitude(operator, strategies[i]) - slashedMagnitude
        //     });
        // }
    }

    /**
     *  @notice Called by an AVS to emit an `AVSMetadataURIUpdated` event indicating the information has updated.
     *
     *  @param metadataURI The URI for metadata associated with an AVS.
     *
     *  @dev Note that the `metadataURI` is *never stored* and is only emitted in the `AVSMetadataURIUpdated` event.
     */
    function updateAVSMetadataURI(
        string calldata metadataURI
    ) external override {
        // emit AVSMetadataURIUpdated(msg.sender, metadataURI);
    }

    /**
     * @notice Called by an operator to cancel a salt that has been used to register with an AVS.
     *
     * @param salt A unique and single use value associated with the approver signature.
     */
    function cancelSalt(
        bytes32 salt
    ) external override {
        // Mutate `operatorSaltIsSpent` to `true` to prevent future spending.
        // operatorSaltIsSpent[msg.sender][salt] = true;
    }

    /**
     *
     *        LEGACY EXTERNAL FUNCTIONS - SUPPORT DEPRECATED IN FUTURE RELEASE AFTER SLASHING RELEASE
     *
     */

    /**
     *  @notice Legacy function called by the AVS's service manager contract
     * to register an operator with the AVS. NOTE: this function will be deprecated in a future release
     * after the slashing release. New AVSs should use `registerOperatorToOperatorSets` instead.
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
        // // Assert `operatorSignature.expiry` has not elapsed.
        // require(
        //     operatorSignature.expiry >= block.timestamp,
        //     "AVSDirectory.registerOperatorToAVS: operator signature expired"
        // );

        // // Assert that the AVS is not an operator set AVS.
        // require(!isOperatorSetAVS[msg.sender], "AVSDirectory.registerOperatorToAVS: AVS is an operator set AVS");

        // // Assert that the `operator` is not actively registered to the AVS.
        // require(
        //     avsOperatorStatus[msg.sender][operator] != OperatorAVSRegistrationStatus.REGISTERED,
        //     "AVSDirectory.registerOperatorToAVS: operator already registered"
        // );

        // // Assert `operator` has not already spent `operatorSignature.salt`.
        // require(
        //     !operatorSaltIsSpent[operator][operatorSignature.salt],
        //     "AVSDirectory.registerOperatorToAVS: salt already spent"
        // );

        // // Assert `operator` is a registered operator.
        // require(
        //     delegation.isOperator(operator),
        //     "AVSDirectory.registerOperatorToAVS: operator not registered to EigenLayer yet"
        // );

        // // Assert that `operatorSignature.signature` is a valid signature for the operator AVS registration.
        // EIP1271SignatureUtils.checkSignature_EIP1271({
        //     signer: operator,
        //     digestHash: calculateOperatorAVSRegistrationDigestHash({
        //         operator: operator,
        //         avs: msg.sender,
        //         salt: operatorSignature.salt,
        //         expiry: operatorSignature.expiry
        //     }),
        //     signature: operatorSignature.signature
        // });

        // // Mutate `operatorSaltIsSpent` to `true` to prevent future respending.
        // operatorSaltIsSpent[operator][operatorSignature.salt] = true;

        // // Set the operator as registered
        // avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.REGISTERED;

        // emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.REGISTERED);
    }

    /**
     *  @notice Legacy function called by an AVS to deregister an operator from the AVS.
     * NOTE: this function will be deprecated in a future release after the slashing release.
     * New AVSs integrating should use `deregisterOperatorFromOperatorSets` instead.
     *
     *  @param operator The address of the operator to deregister.
     *
     *  @dev Only used by legacy M2 AVSs that have not integrated with operator sets.
     */
    function deregisterOperatorFromAVS(
        address operator
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        // require(
        //     avsOperatorStatus[msg.sender][operator] == OperatorAVSRegistrationStatus.REGISTERED,
        //     "AVSDirectory.deregisterOperatorFromAVS: operator not registered"
        // );

        // // Assert that the AVS is not an operator set AVS.
        // require(!isOperatorSetAVS[msg.sender], "AVSDirectory.deregisterOperatorFromAVS: AVS is an operator set AVS");

        // // Set the operator as deregistered
        // avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.UNREGISTERED;

        // emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.UNREGISTERED);
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

            bytes32 encodedOperatorSet = _encodeOperatorSet(operatorSet);

            require(
                _operatorSetsMemberOf[operator].add(encodedOperatorSet),
                "AVSDirectory._registerOperatorToOperatorSets: operator already registered to operator set"
            );

            _operatorSetMembers[encodedOperatorSet].add(operator);

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

            bytes32 encodedOperatorSet = _encodeOperatorSet(operatorSet);

            require(
                _operatorSetsMemberOf[operator].remove(encodedOperatorSet),
                "AVSDirectory._deregisterOperatorFromOperatorSet: operator not registered for operator set"
            );

            _operatorSetMembers[encodedOperatorSet].remove(operator);

            emit OperatorRemovedFromOperatorSet(operator, operatorSet);
        }
    }

    /**
     * @notice For a single strategy, update freeMagnitude by adding completable pending free magnitudes
     * @param operator address to update freeMagnitude for
     * @param strategy the strategy to update freeMagnitude for
     * @param numToComplete the number of pending free magnitudes deallocations to complete
     * @dev read through pending free magnitudes and add to freeMagnitude if completableTimestamp is >= block timestamp
     * In addition to updating freeMagnitude, updates next starting index to read from for pending free magnitudes after completing
     */
    function _updateFreeMagnitude(address operator, IStrategy strategy, uint16 numToComplete) internal {
        OperatorMagnitudeInfo storage info = operatorMagnitudeInfo[operator][strategy];
        (uint64 freeMagnitudeToAdd, uint192 completed) =
            _getPendingFreeMagnitude(operator, strategy, numToComplete, info.nextPendingFreeMagnitudeIndex);
        info.freeMagnitude += freeMagnitudeToAdd;
        info.nextPendingFreeMagnitudeIndex += completed;
    }

    /**
     * @notice For a single strategy, modify magnitude allocations for each of the specified operatorSets
     * @param operator address to modify allocations for
     * @param allocation the magnitude allocations to modify for a single strategy
     * @param allocationEffectTimestamp the timestamp when the allocations will take effect
     * @param deallocationCompletableTimestamp the timestamp when the deallocations will be completable
     * @dev For each allocation, allocation.operatorSets MUST be ordered in ascending order according to the
     * encoding of the operatorSet. This is to prevent duplicate operatorSets being passed in. The easiest way to ensure
     * ordering is to sort allocated operatorSets by address first, and then sort for each avs by ascending operatorSetIds.
     */
    function _modifyAllocations(
        address operator,
        MagnitudeAllocation calldata allocation,
        uint32 allocationEffectTimestamp,
        uint32 deallocationCompletableTimestamp
    ) internal {
        require(
            allocation.operatorSets.length == allocation.magnitudes.length,
            "AVSDirectory._modifyAllocations: operatorSets and magnitudes length mismatch"
        );

        // OperatorSet[] calldata opSets = allocation.operatorSets;

        bytes32 prevOperatorSet = bytes32(0);

        for (uint256 i = 0; i < allocation.operatorSets.length; ++i) {
            require(
                isOperatorSet[allocation.operatorSets[i].avs][allocation.operatorSets[i].operatorSetId],
                "AVSDirectory._modifyAllocations: operatorSet does not exist"
            );
            // use encoding of operatorSet to ensure ordering and also used to use OperatorSet struct as key in mappings
            bytes32 operatorSetKey = _encodeOperatorSet(allocation.operatorSets[i]);
            require(prevOperatorSet < operatorSetKey, "AVSDirectory._modifyAllocations: operatorSets not ordered");
            prevOperatorSet = operatorSetKey;

            // Read current magnitude allocation including its respective array index and length.
            // We'll use these values later to check the number of pending allocations/deallocations.
            (uint224 currentMagnitude, uint256 pos, uint256 length) = _magnitudeUpdate[operator][allocation.strategy][operatorSetKey]
                .upperLookupRecentWithPos(uint32(block.timestamp));

            // Check that there is at MOST `MAX_PENDING_UPDATES` combined allocations & deallocations for the operator, operatorSet, strategy
            {
                uint256 numPendingAllocations;
                // if no lookup found (currentMagnitude == 0 && pos == 0), then we are at the beginning of the array
                // the number of pending allocations is simply length
                if (currentMagnitude == 0 && pos == 0) {
                    numPendingAllocations = length;
                    // if lookup found, then we take the difference between length-1 and pos
                } else {
                    numPendingAllocations = length - pos - 1;
                }
                uint256 numPendingDeallocations =
                    _getNumQueuedDeallocations(operator, allocation.strategy, operatorSetKey);

                require(
                    numPendingAllocations + numPendingDeallocations < MAX_PENDING_UPDATES,
                    "AVSDirectory._modifyAllocations: Cannot set magnitude with a pending allocation or deallocation"
                );
            }

            if (allocation.magnitudes[i] < uint64(currentMagnitude)) {
                // Newly configured magnitude is less than current value.
                // Therefore we handle this as a deallocation

                // Note: MAX_PENDING_UPDATES == 1, so we do not have to decrement any allocations

                // 1. push PendingFreeMagnitude and respective array index into (op,opSet,Strategy) queued deallocations
                uint256 index = _pendingFreeMagnitude[operator][allocation.strategy].length;
                _pendingFreeMagnitude[operator][allocation.strategy].push(
                    PendingFreeMagnitude({
                        magnitudeDiff: uint64(currentMagnitude) - allocation.magnitudes[i],
                        completableTimestamp: deallocationCompletableTimestamp
                    })
                );
                _queuedDeallocationIndices[operator][allocation.strategy][operatorSetKey].push(index);
            } else if (allocation.magnitudes[i] > uint64(currentMagnitude)) {
                // Newly configured magnitude is greater than current value.
                // Therefore we handle this as an allocation

                // 1. allocate magnitude which will take effect in the future 21 days from now
                _magnitudeUpdate[operator][allocation.strategy][operatorSetKey].push({
                    key: allocationEffectTimestamp,
                    value: allocation.magnitudes[i]
                });

                // 2. decrement free magnitude by incremented amount
                OperatorMagnitudeInfo storage info = operatorMagnitudeInfo[operator][allocation.strategy];
                require(
                    info.freeMagnitude >= allocation.magnitudes[i] - uint64(currentMagnitude),
                    "AVSDirectory._modifyAllocations: insufficient available free magnitude to allocate"
                );
                info.freeMagnitude -= allocation.magnitudes[i] - uint64(currentMagnitude);
            }
        }
    }

    /**
     * @notice Get the number of queued dealloations for the given (operator, strategy, operatorSetKey) tuple
     * @param operator address to get queued deallocations for
     * @param strategy the strategy to get queued deallocations for
     * @param operatorSetKey the encoded operatorSet to get queued deallocations for
     */
    function _getNumQueuedDeallocations(
        address operator,
        IStrategy strategy,
        bytes32 operatorSetKey
    ) internal view returns (uint256) {
        uint256 numQueuedDeallocations;

        uint256 length = _queuedDeallocationIndices[operator][strategy][operatorSetKey].length;

        for (uint256 i = length; i > 0; --i) {
            // index of pendingFreeMagnitude/deallocation to check for slashing
            uint256 index = _queuedDeallocationIndices[operator][strategy][operatorSetKey][i - 1];
            PendingFreeMagnitude memory pendingFreeMagnitude = _pendingFreeMagnitude[operator][strategy][index];

            // If completableTimestamp is greater than completeUntilTimestamp, break
            if (pendingFreeMagnitude.completableTimestamp < uint32(block.timestamp)) {
                ++numQueuedDeallocations;
            } else {
                break;
            }
        }

        return numQueuedDeallocations;
    }

    /// @dev gets the latest total magnitude or overwrites it if it is not set
    function _getLatestTotalMagnitude(address operator, IStrategy strategy) internal returns (uint64) {
        (bool exists,, uint224 totalMagnitude) = _totalMagnitudeUpdate[operator][strategy].latestCheckpoint();
        if (!exists) {
            totalMagnitude = ShareScalingLib.INITIAL_TOTAL_MAGNITUDE;
            _totalMagnitudeUpdate[operator][strategy].push({key: uint32(block.timestamp), value: totalMagnitude});
            operatorMagnitudeInfo[operator][strategy].freeMagnitude = ShareScalingLib.INITIAL_TOTAL_MAGNITUDE;
        }

        return uint64(totalMagnitude);
    }

    /// @dev gets the latest total magnitude or overwrites it if it is not set
    function _getLatestTotalMagnitudeView(address operator, IStrategy strategy) internal view returns (uint64) {
        (bool exists,, uint224 totalMagnitude) = _totalMagnitudeUpdate[operator][strategy].latestCheckpoint();
        if (!exists) {
            return ShareScalingLib.INITIAL_TOTAL_MAGNITUDE;
        }

        return uint64(totalMagnitude);
    }

    /// @dev gets the pending free magnitude available to add by completing numToComplete pending deallocations
    /// and returns the next index to start from if completed.
    function _getPendingFreeMagnitude(
        address operator,
        IStrategy strategy,
        uint16 numToComplete,
        uint192 nextIndex
    ) internal view returns (uint64 freeMagnitudeToAdd, uint192 completed) {
        uint256 pendingFreeMagnitudeLength = _pendingFreeMagnitude[operator][strategy].length;
        uint16 completed = 0;
        freeMagnitudeToAdd = 0;
        while (nextIndex < pendingFreeMagnitudeLength && completed < numToComplete) {
            PendingFreeMagnitude memory pendingFreeMagnitude = _pendingFreeMagnitude[operator][strategy][nextIndex];
            // pendingFreeMagnitude is ordered by completableTimestamp. If we reach one that is not completable yet, then break
            // loop until completableTimestamp is < block.timestamp
            if (pendingFreeMagnitude.completableTimestamp < uint32(block.timestamp)) {
                break;
            }

            // pending free magnitude can be added to freeMagnitude
            freeMagnitudeToAdd += pendingFreeMagnitude.magnitudeDiff;
            ++nextIndex;
            ++completed;
        }
        return (freeMagnitudeToAdd, nextIndex);
    }

    /// @dev Verify operator's signature and spend salt
    function _verifyOperatorSignature(
        address operator,
        MagnitudeAllocation[] calldata allocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) internal {
        // check the signature expiry
        // require(
        //     operatorSignature.expiry >= block.timestamp,
        //     "AVSDirectory._verifyOperatorSignature: operator signature expired"
        // );
        // // Assert operator's signature cannot be replayed.
        // require(
        //     !operatorSaltIsSpent[operator][operatorSignature.salt], "AVSDirectory._verifyOperatorSignature: salt spent"
        // );

        // bytes32 digestHash = calculateMagnitudeAllocationDigestHash(
        //     operator, allocations, operatorSignature.salt, operatorSignature.expiry
        // );

        // // Assert operator's signature is valid.
        // EIP1271SignatureUtils.checkSignature_EIP1271(operator, digestHash, operatorSignature.signature);
        // // Spend salt.
        // operatorSaltIsSpent[operator][operatorSignature.salt] = true;
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice Returns operatorSet an operator is registered to in the order they were registered.
     * @param operator The operator address to query.
     * @param index The index of the enumerated list of operator sets.
     */
    function operatorSetsMemberOfAtIndex(address operator, uint256 index) external view returns (OperatorSet memory) {
        return _decodeOperatorSet(_operatorSetsMemberOf[operator].at(index));
    }

    /**
     * @notice Returns the operator registered to an operatorSet in the order that it was registered.
     * @param operatorSet The operatorSet to query.
     * @param index The index of the enumerated list of operators.
     */
    function operatorSetMemberAtIndex(OperatorSet memory operatorSet, uint256 index) external view returns (address) {
        return _operatorSetMembers[_encodeOperatorSet(operatorSet)].at(index);
    }

    /**
     * @notice Returns an array of operator sets an operator is registered to.
     * @param operator The operator address to query.
     * @param start The starting index of the array to query.
     *  @param length The amount of items of the array to return.
     */
    function getOperatorSetsOfOperator(
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

    /**
     * @notice Returns an array of operators registered to the operatorSet.
     * @param operatorSet The operatorSet to query.
     * @param start The starting index of the array to query.
     * @param length The amount of items of the array to return.
     */
    function getOperatorsInOperatorSet(
        OperatorSet memory operatorSet,
        uint256 start,
        uint256 length
    ) external view returns (address[] memory operators) {
        bytes32 encodedOperatorSet = _encodeOperatorSet(operatorSet);
        uint256 maxLength = _operatorSetMembers[encodedOperatorSet].length() - start;
        if (length > maxLength) length = maxLength;
        operators = new address[](length);
        for (uint256 i; i < length; ++i) {
            operators[i] = _operatorSetMembers[encodedOperatorSet].at(start + i);
        }
    }

    /**
     * @notice Returns the number of operators registered to an operatorSet.
     * @param operatorSet The operatorSet to get the member count for
     */
    function getNumOperatorsInOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (uint256) {
        return _operatorSetMembers[_encodeOperatorSet(operatorSet)].length();
    }

    /**
     *  @notice Returns the total number of operator sets an operator is registered to.
     *  @param operator The operator address to query.
     */
    function inTotalOperatorSets(
        address operator
    ) external view returns (uint256) {
        return _operatorSetsMemberOf[operator].length();
    }

    /**
     * @notice Returns whether or not an operator is registered to an operator set.
     * @param operator The operator address to query.
     *  @param operatorSet The `OperatorSet` to query.
     */
    function isMember(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        return _operatorSetsMemberOf[operator].contains(_encodeOperatorSet(operatorSet));
    }

    /**
     * @param operator the operator to get the slashable magnitude for
     * @param strategies the strategies to get the slashable magnitude for
     *
     * @return operatorSets the operator sets the operator is a member of and the current slashable magnitudes for each strategy
     */
    function getCurrentSlashableMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (OperatorSet[] memory, uint64[][] memory) {
        OperatorSet[] memory operatorSets =
            getOperatorSetsOfOperator(operator, 0, _operatorSetsMemberOf[operator].length());
        uint64[][] memory slashableMagnitudes = new uint64[][](strategies.length);
        // for (uint256 i = 0; i < strategies.length; ++i) {
        //     slashableMagnitudes[i] = new uint64[](operatorSets.length);
        //     for (uint256 j = 0; j < operatorSets.length; ++j) {
        //         slashableMagnitudes[i][j] = uint64(
        //             _magnitudeUpdate[operator][strategies[i]][_encodeOperatorSet(operatorSets[j])].upperLookupLinear(
        //                 uint32(block.timestamp)
        //             )
        //         );
        //     }
        // }
        return (operatorSets, slashableMagnitudes);
    }

    /**
     * @param operator the operator to get the slashable magnitude for
     * @param strategies the strategies to get the slashable magnitude for
     * @param timestamp the timestamp to get the slashable magnitude for
     *
     * @return operatorSets the operator sets the operator is a member of and the slashable magnitudes for each strategy
     */
    function getSlashableMagnitudes(
        address operator,
        IStrategy[] calldata strategies,
        uint32 timestamp
    ) external view returns (OperatorSet[] memory, uint64[][] memory) {
        OperatorSet[] memory operatorSets =
            getOperatorSetsOfOperator(operator, 0, _operatorSetsMemberOf[operator].length());
        uint64[][] memory slashableMagnitudes = new uint64[][](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            slashableMagnitudes[i] = new uint64[](operatorSets.length);
            for (uint256 j = 0; j < operatorSets.length; ++j) {
                slashableMagnitudes[i][j] = uint64(
                    _magnitudeUpdate[operator][strategies[i]][_encodeOperatorSet(operatorSets[j])].upperLookupRecent(
                        timestamp
                    )
                );
            }
        }
        return (operatorSets, slashableMagnitudes);
    }

    function _getTotalAndAllocatedMagnitude(
        address operator,
        OperatorSet calldata operatorSet,
        IStrategy strategy
    ) internal view returns (uint64, uint64) {
        uint64 totalMagnitude = _getLatestTotalMagnitudeView(operator, strategy);

        bytes32 operatorSetKey = _encodeOperatorSet(operatorSet);
        uint64 currentMagnitude = uint64(
                _magnitudeUpdate[operator][strategy][operatorSetKey].upperLookupLinear(
                    uint32(block.timestamp)
                )
            );

        return (totalMagnitude, currentMagnitude);
    }

    /**
     * @notice Get the allocatable magnitude for an operator and strategy based on number of pending deallocations
     * that could be completed at the same time. This is the sum of freeMagnitude and the sum of all pending completable deallocations.
     * @param operator the operator to get the allocatable magnitude for
     * @param strategy the strategy to get the allocatable magnitude for
     * @param numToComplete the number of pending free magnitudes deallocations to complete
     */
    function getAllocatableMagnitude(
        address operator,
        IStrategy strategy,
        uint16 numToComplete
    ) external view returns (uint64) {
        OperatorMagnitudeInfo storage info = operatorMagnitudeInfo[operator][strategy];
        (uint64 freeMagnitudeToAdd, ) =
            _getPendingFreeMagnitude(operator, strategy, numToComplete, info.nextPendingFreeMagnitudeIndex);
        return info.freeMagnitude + freeMagnitudeToAdd;
    }

    /// @notice operator is slashable by operatorSet if currently registered OR last deregistered within 21 days
    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        OperatorSetRegistrationStatus memory registrationStatus =
            operatorSetStatus[operatorSet.avs][operator][operatorSet.operatorSetId];
        return isMember(operator, operatorSet)
            || registrationStatus.lastDeregisteredTimestamp + DEALLOCATION_DELAY >= block.timestamp;
    }

    /// @notice Returns the total magnitude of an operator for a given set of strategies
    function getTotalMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (uint64[] memory) {
        uint64[] memory totalMagnitudes = new uint64[](strategies.length);
        for (uint256 i = 0; i < strategies.length;) {
            (bool exists, uint32 key, uint224 value) = _totalMagnitudeUpdate[operator][strategies[i]].latestCheckpoint();
            if (!exists) {
                totalMagnitudes[i] = ShareScalingLib.INITIAL_TOTAL_MAGNITUDE;
            } else {
                totalMagnitudes[i] = uint64(value);
            }

            unchecked {
                ++i;
            }
        }
        return totalMagnitudes;
    }

    /// @notice Returns the total magnitude of an operator for a given set of strategies at a given timestamp
    function getTotalMagnitudesAtTimestamp(
        address operator,
        IStrategy[] calldata strategies,
        uint32 timestamp
    ) external view returns (uint64[] memory) {
        uint64[] memory totalMagnitudes = new uint64[](strategies.length);
        for (uint256 i = 0; i < strategies.length;) {
            (uint224 value, uint256 pos, uint256 length) =
                _totalMagnitudeUpdate[operator][strategies[i]].upperLookupRecentWithPos(timestamp);

            // if there is no existing total magnitude checkpoint
            if (value != 0 || pos != 0) {
                totalMagnitudes[i] = ShareScalingLib.INITIAL_TOTAL_MAGNITUDE;
            } else {
                totalMagnitudes[i] = uint64(value);
            }

            unchecked {
                ++i;
            }
        }
        return totalMagnitudes;
    }

    /**
     * @param operator the operator to get the total and allocated magnitudes for
     * @param operatorSet the operatorSet to get the total and allocated magnitudes for
     * @param strategies the strategies to get the total and allocated magnitudes for
     *
     * @return the list of total magnitudes for each strategy and the list of allocated magnitudes for each strategy
     */
    function getTotalAndAllocatedMagnitudes(
        address operator,
        OperatorSet calldata operatorSet,
        IStrategy[] calldata strategies
    ) public view returns (uint64[] memory, uint64[] memory) {
        uint64[] memory totalMagnitude = new uint64[](strategies.length);
        uint64[] memory allocatedMagnitude = new uint64[](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            (totalMagnitude[i], allocatedMagnitude[i]) = _getTotalAndAllocatedMagnitude(operator, operatorSet, strategies[i]);
        }
        return (totalMagnitude, allocatedMagnitude);
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

    /**
     * @notice Calculates the digest hash to be signed by an operator to modify magnitude allocations
     * @param operator The operator to allocate or deallocate magnitude for.
     * @param allocations The magnitude allocations/deallocations to be made.
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid.
     */
    function calculateMagnitudeAllocationDigestHash(
        address operator,
        MagnitudeAllocation[] calldata allocations,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        return _calculateDigestHash(
            keccak256(abi.encode(MAGNITUDE_ADJUSTMENT_TYPEHASH, operator, allocations, salt, expiry))
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
    function _calculateDigestHash(
        bytes32 structHash
    ) internal view returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", _calculateDomainSeparator(), structHash));
    }

    /// @dev Returns an `OperatorSet` encoded into a 32-byte value.
    /// @param operatorSet The `OperatorSet` to encode.
    function _encodeOperatorSet(
        OperatorSet memory operatorSet
    ) internal pure returns (bytes32) {
        return bytes32(abi.encodePacked(operatorSet.avs, uint96(operatorSet.operatorSetId)));
    }

    /// @dev Returns an `OperatorSet` decoded from an encoded 32-byte value.
    /// @param encoded The encoded `OperatorSet` to decode.
    /// @dev Assumes `encoded` is encoded via `_encodeOperatorSet(operatorSet)`.
    function _decodeOperatorSet(
        bytes32 encoded
    ) internal pure returns (OperatorSet memory) {
        return OperatorSet({
            avs: address(uint160(uint256(encoded) >> 96)),
            operatorSetId: uint32(uint256(encoded) & type(uint96).max)
        });
    }
}
