// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./IDelegationManager.sol";
import "./ISignatureUtils.sol";

interface IAVSDirectory is ISignatureUtils {
    /// @notice Enum representing the registration status of an operator with an AVS.
    /// @notice Only used by legacy M2 AVSs that have not integrated with operatorSets.
    enum OperatorAVSRegistrationStatus {
        UNREGISTERED, // Operator not registered to AVS
        REGISTERED // Operator registered to AVS
    }

    /**
     * @notice Struct representing the registration status of an operator with an operator set.
     * Keeps track of last deregistered timestamp for slashability concerns.
     * @param registered whether the operator is registered with the operator set
     * @param lastDeregisteredTimestamp the timestamp at which the operator was last deregistered
     */
    struct OperatorSetRegistrationStatus {
        bool registered;
        uint32 lastDeregisteredTimestamp;
    }

    /// @notice Struct representing an operator set
    struct OperatorSet {
        address avs;
        uint32 operatorSetId;
    }

    /**
     * @notice struct used to modify the allocation of slashable magnitude to list of operatorSets
     * @param strategy the strategy to allocate magnitude for
     * @param expectedTotalMagnitude the expected total magnitude of the operator used to combat against race conditions with slashing
     * @param operatorSets the operatorSets to allocate magnitude for
     * @param magnitudes the magnitudes to allocate for each operatorSet
     */
    struct MagnitudeAllocation {
        IStrategy strategy;
        uint64 expectedTotalMagnitude;
        OperatorSet[] operatorSets;
        uint64[] magnitudes;
    }

    /**
     * @notice struct used for pending free magnitude. Stored in (operator, strategy, operatorSet) mapping
     * to be used in completeDeallocations.
     * @param magnitudeDiff the amount of magnitude to deallocate
     * @param completableTimestamp the timestamp at which the deallocation can be completed, 21 days from when queued
     */
    struct PendingFreeMagnitude {
        uint64 magnitudeDiff;
        uint32 completableTimestamp;
    }

    /**
     * @notice Struct containing info regarding free allocatable magnitude.
     * @param nextPendingFreeMagnitudeIndex The next available update index.
     * @param freeMagnitude The total amount of free allocatable magnitude.
     */
    struct OperatorMagnitudeInfo {
        uint192 nextPendingFreeMagnitudeIndex;
        uint64 freeMagnitude;
    }

    /// @notice Emitted when an operator set is created by an AVS.
    event OperatorSetCreated(OperatorSet operatorSet);

    /**
     *  @notice Emitted when an operator's registration status with an AVS id udpated
     *  @notice Only used by legacy M2 AVSs that have not integrated with operatorSets.
     */
    event OperatorAVSRegistrationStatusUpdated(
        address indexed operator, address indexed avs, OperatorAVSRegistrationStatus status
    );

    /// @notice Emitted when an operator is added to an operator set.
    event OperatorAddedToOperatorSet(address indexed operator, OperatorSet operatorSet);

    /// @notice Emitted when an operator is removed from an operator set.
    event OperatorRemovedFromOperatorSet(address indexed operator, OperatorSet operatorSet);

    /// @notice Emitted when an AVS updates their metadata URI (Uniform Resource Identifier).
    /// @dev The URI is never stored; it is simply emitted through an event for off-chain indexing.
    event AVSMetadataURIUpdated(address indexed avs, string metadataURI);

    /// @notice Emitted when an AVS migrates to using operator sets.
    event AVSMigratedToOperatorSets(address indexed avs);

    /// @notice Emitted when an operator is migrated from M2 registration to operator sets.
    event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds);

    /// @notice Emitted when an operator allocates slashable magnitude to an operator set
    event MagnitudeAllocated(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint64 magnitudeToAllocate,
        uint32 effectTimestamp
    );

    /// @notice Emitted when an operator queues deallocations of slashable magnitude from an operator set
    event MagnitudeQueueDeallocated(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint64 magnitudeToDeallocate,
        uint32 completableTimestamp
    );

    /// @notice Emitted when an operator completes deallocations of slashable magnitude from an operator set
    /// and adds back magnitude to free allocatable magnitude
    event MagnitudeDeallocationCompleted(
        address operator, IStrategy strategy, OperatorSet operatorSet, uint64 freeMagnitudeAdded
    );

    /// @notice Emitted when an operator is slashed by an operator set for a strategy
    event OperatorSlashed(address operator, uint32 operatorSetId, IStrategy strategy, uint16 bipsToSlash);

    /**
     *
     *                         EXTERNAL FUNCTIONS
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
    function createOperatorSets(uint32[] calldata operatorSetIds) external;

    /**
     * @notice Sets the AVS as an operator set AVS, preventing legacy M2 operator registrations.
     *
     * @dev msg.sender must be the AVS.
     */
    function becomeOperatorSetAVS() external;

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
    ) external;

    /**
     *  @notice Called by AVSs to add an operator to list of operatorSets.
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
    ) external;

    /**
     *  @notice Called by AVSs to remove an operator from an operator set.
     *
     *  @param operator The address of the operator to be removed from the operator set.
     *  @param operatorSetIds The IDs of the operator sets.
     *
     *  @dev msg.sender is used as the AVS.
     */
    function deregisterOperatorFromOperatorSets(address operator, uint32[] calldata operatorSetIds) external;

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
    ) external;

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
    ) external;

    /**
     *  @notice Legacy function called by an AVS to deregister an operator from the AVS.
     * NOTE: this function will be deprecated in a future release after the slashing release.
     * New AVSs integrating should use `deregisterOperatorFromOperatorSets` instead.
     *
     *  @param operator The address of the operator to deregister.
     *
     *  @dev Only used by legacy M2 AVSs that have not integrated with operator sets.
     */
    function deregisterOperatorFromAVS(address operator) external;

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
    ) external;

    /**
     * @notice For all pending deallocations that have become completable, their pending free magnitude can be
     * added back to the free magnitude of the (operator, strategy) amount. This function takes a list of strategies
     * and adds all completable deallocations for each strategy, updating the freeMagnitudes of the operator
     *
     * @param operator address to complete deallocations for
     * @param strategies a list of strategies to complete deallocations for
     *
     * @dev can be called permissionlessly by anyone
     */
    function updateFreeMagnitude(
        address operator,
        IStrategy[] calldata strategies,
        uint16[] calldata numToComplete
    ) external;

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
    ) external;

    /**
     *  @notice Called by an AVS to emit an `AVSMetadataURIUpdated` event indicating the information has updated.
     *
     *  @param metadataURI The URI for metadata associated with an AVS.
     *
     *  @dev Note that the `metadataURI` is *never stored* and is only emitted in the `AVSMetadataURIUpdated` event.
     */
    function updateAVSMetadataURI(string calldata metadataURI) external;

    /**
     * @notice Called by an operator to cancel a salt that has been used to register with an AVS.
     *
     * @param salt A unique and single use value associated with the approver signature.
     */
    function cancelSalt(bytes32 salt) external;

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice Get the allocatable magnitude for an operator and strategy based on number of pending deallocations
     * that could be completed at the same time. This is the sum of freeMagnitude and the sum of all pending completable deallocations.
     * @param operator the operator to get the allocatable magnitude for
     * @param strategy the strategy to get the allocatable magnitude for
     * @param numToComplete the number of pending free magnitudes deallocations to complete, 0 to complete all (uint8 max 256)
     */
    function getAllocatableMagnitude(
        address operator,
        IStrategy strategy,
        uint16 numToComplete
    ) external view returns (uint64);

    /// @dev Delay before deallocations take effect and are added back into freeMagnitude
    function DEALLOCATION_DELAY() external view returns (uint32);

    /**
     * @notice operator is slashable by operatorSet if currently registered OR last deregistered within 21 days
     * @param operator the operator to check slashability for
     * @param operatorSet the operatorSet to check slashability for
     * @return bool if the operator is slashable by the operatorSet
     */
    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) external view returns (bool);

    /**
     * @param operator the operator to get the slashable magnitude for
     * @param strategies the strategies to get the slashable magnitude for
     * 
     * @return operatorSets the operator sets the operator is a member of and the current slashable magnitudes for each strategy
     */
    function getCurrentSlashableMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (OperatorSet[] memory, uint64[][] memory);

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
    ) external view returns (OperatorSet[] memory, uint64[][] memory);

    /**
     * @param operator the operator to get the slashable ppm for
     * @param operatorSet the operatorSet to get the slashable ppm for
     * @param strategies the strategies to get the slashable ppm for
     *
     * @return the list of total magnitudes for each strategy and the list of allocated magnitudes for each strategy
     */
    function getTotalAndAllocatedMagnitudes(
        address operator,
        OperatorSet calldata operatorSet,
        IStrategy[] calldata strategies
    ) external view returns (uint64[] memory, uint64[] memory);

    /// @notice Returns the total magnitude of an operator for a given set of strategies
    /// TODO: finish natspec
    function getTotalMagnitudes(address operator, IStrategy[] calldata strategies) external view returns (uint64[] memory);

    /// @notice Returns the total magnitude of an operator for a given set of strategies at a given timestamp
    /// TODO: finish natspec
    function getTotalMagnitudesAtTimestamp(
        address operator,
        IStrategy[] calldata strategies,
        uint32 timestamp
    ) external view returns (uint64[] memory);

    function operatorSaltIsSpent(address operator, bytes32 salt) external view returns (bool);

    function isMember(address operator, OperatorSet memory operatorSet) external view returns (bool);

    function isOperatorSetAVS(address avs) external view returns (bool);

    function isOperatorSet(address avs, uint32 operatorSetId) external view returns (bool);

    /**
     * @notice Returns operator set an operator is registered to in the order they were registered.
     * @param operator The operator address to query.
     * @param index The index of the enumerated list of operator sets.
     */
    function operatorSetsMemberOfAtIndex(address operator, uint256 index) external view returns (OperatorSet memory);

    /**
     * @notice Retursn the operator registered to an operatorSet in the order that it was registered.
     *  @param operatorSet The operatorSet to query.
     *  @param index The index of the enumerated list of operators.
     */
    function operatorSetMemberAtIndex(OperatorSet memory operatorSet, uint256 index) external view returns (address);

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
    ) external view returns (OperatorSet[] memory operatorSets);

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
    ) external view returns (address[] memory operators);

    /**
     * @notice Returns the number of operators registered to an operatorSet.
     * @param operatorSet The operatorSet to get the member count for
     */
    function getNumOperatorsInOperatorSet(OperatorSet memory operatorSet) external view returns (uint256);

    /**
     *  @notice Returns the total number of operator sets an operator is registered to.
     *  @param operator The operator address to query.
     */
    function inTotalOperatorSets(address operator) external view returns (uint256);

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
    ) external view returns (bytes32);

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
    ) external view returns (bytes32);

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
    ) external view returns (bytes32);

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
    ) external view returns (bytes32);

    /// @notice Getter function for the current EIP-712 domain separator for this contract.
    /// @dev The domain separator will change in the event of a fork that changes the ChainID.
    function domainSeparator() external view returns (bytes32);

    /// @notice The EIP-712 typehash for the Registration struct used by the contract.
    function OPERATOR_AVS_REGISTRATION_TYPEHASH() external view returns (bytes32);

    /// @notice The EIP-712 typehash for the OperatorSetRegistration struct used by the contract.
    function OPERATOR_SET_REGISTRATION_TYPEHASH() external view returns (bytes32);
}
