// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./ISignatureUtils.sol";

/// @notice Struct representing an operator set
struct OperatorSet {
    uint32 avs; // Identifier of an avs
    uint32 operatorSetId;
}

interface IAVSDirectory is ISignatureUtils {
    /// Operator Status

    /// @dev Thrown when an operator does not exist in the DelegationManager
    error OperatorDoesNotExist();
    /// @dev Thrown when `operator` is already registered to the AVS.
    error OperatorAlreadyRegistered();

    /// @notice Enum representing the status of an operator's registration with an AVS
    /// @dev Thrown when an invalid AVS is provided.
    error InvalidAVS();
    /// @dev Thrown when an invalid operator is provided.
    error InvalidOperator();
    /// @dev Thrown when an invalid operator set is provided.
    error InvalidOperatorSet();
    /// @dev Thrown when `operator` is not a registered operator.
    error OperatorNotRegistered();
    /// @dev Thrown when attempting to spend a spent eip-712 salt.
    error SaltSpent();
    /// @dev Thrown when attempting to use an expired eip-712 signature.
    error SignatureExpired();
    /// @dev Thrown when a identifier is already claimed by an AVS
    error IdentifierAlreadyClaimed();
    /// @dev Thrown when an AVS already has a identifier
    error AVSIdentifierNotSet();

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
    event AVSMetadataURIUpdated(uint32 indexed avs, string metadataURI);

    /// @notice Emitted when an AVS migrates to using operator sets.
    event AVSMigratedToOperatorSets(uint32 indexed avs, address indexed dispatcher);

    /// @notice Emitted when an operator is migrated from M2 registration to operator sets.
    event OperatorMigratedToOperatorSets(address indexed operator, uint32 indexed avs, uint32[] operatorSetIds);

    /// @notice Emitted when an AVS claims a identifier
    event AVSIdentifierClaimed(uint32 indexed avs, address dispatcher);

    /// @notice Emitted when an AVS has relinquished their identifier
    event AVSIdentifierTransferred(uint32 indexed avs, address oldDispatcher, address newDispatcher);

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
    function createOperatorSets(
        uint32[] calldata operatorSetIds
    ) external;

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
     * @param avs The avs to deregister the operator from.
     * @param operatorSetIds The IDs of the operator sets.
     * @param operatorSignature the signature of the operator on their intent to deregister or empty if the operator itself is calling
     *
     * @dev if the operatorSignature is empty, the caller must be the operator
     * @dev this will likely only be called in case the AVS contracts are in a state that prevents operators from deregistering
     */
    function forceDeregisterFromOperatorSets(
        address operator,
        uint32 avs,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external;

    /**
     *  @notice Legacy function called by the AVS's service manager (dispatcher) contract
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
    function deregisterOperatorFromAVS(
        address operator
    ) external;


    /**
     * @notice Claims an identiier for the avs
     */
    function claimIdentifier() external;

    /**
     * @notice Transfers the identifier associated with the caller to a new dispatcher
     * @param newDispatcher The new dispatcher to transfer the identifier
     * @dev This function should be used in tandem with `claimidentifier` for atomic
     *      claiming of a identifier when an AVS's dispatcher contract is migrated
     */
    function transferIdentifier(address newDispatcher) external;   

    /**
     *  @notice Called by an AVS to emit an `AVSMetadataURIUpdated` event indicating the information has updated.
     *
     *  @param metadataURI The URI for metadata associated with an AVS.
     *
     *  @dev Note that the `metadataURI` is *never stored* and is only emitted in the `AVSMetadataURIUpdated` event.
     */
    function updateAVSMetadataURI(
        string calldata metadataURI
    ) external;

    /**
     * @notice Called by an operator to cancel a salt that has been used to register with an AVS.
     *
     * @param salt A unique and single use value associated with the approver signature.
     */
    function cancelSalt(
        bytes32 salt
    ) external;

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */
    function operatorSaltIsSpent(address operator, bytes32 salt) external view returns (bool);

    function isMember(address operator, OperatorSet memory operatorSet) external view returns (bool);

    function isOperatorSetAVS(
        uint32 avs
    ) external view returns (bool);

    function isOperatorSet(bytes32 encodedOperatorSet) external view returns (bool);

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
    function getNumOperatorsInOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (uint256);

    /**
     *  @notice Returns the total number of operator sets an operator is registered to.
     *  @param operator The operator address to query.
     */
    function inTotalOperatorSets(
        address operator
    ) external view returns (uint256);

    /**
     *  @notice Calculates the digest hash to be signed by an operator to register with an AVS.
     *
     *  @param operator The account registering as an operator.
     *  @param avs The AVS the operator is registering to.
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
        uint32 avs,
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
        uint32 avs,
        uint32[] calldata operatorSetIds,
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

    function operatorSetStatus(
        address operator,
        bytes32 encodedOperatorSet
    ) external view returns (bool registered, uint32 lastDeregisteredTimestamp);

    /**
     * @notice Method to get the status of an operator set by dispatcher contract
     * @param dispatcher The dispatcher contract of the avs
     * @param operator The operator to query
     * @param operatorSetId The operator set id to query
     */
    function operatorSetStatusByDispatcher(
        address dispatcher,
        address operator,
        uint32 operatorSetId
    ) external view returns (bool registered, uint32 lastDeregisteredTimestamp);

    /**
     * @notice Returns the dispatcher associated with an AVS Identifier
     */
    function avsToDispatcher(uint32 avs) external view returns (address);

    /**
     * @notice Returns the avs associated with a dispatcher
     */
    function dispatcherToAVS(address dispatcher) external view returns (uint32);
}
