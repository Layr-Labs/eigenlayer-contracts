// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./ISignatureUtils.sol";
import "./IPauserRegistry.sol";
import "./IStrategy.sol";

interface IAVSDirectoryErrors {
    /// Operator Status

    /// @dev Thrown when an operator does not exist in the DelegationManager
    error OperatorNotRegisteredToEigenLayer();
    /// @dev Thrown when an operator is already registered to an AVS.
    error OperatorNotRegisteredToAVS();
    /// @dev Thrown when `operator` is already registered to the AVS.
    error OperatorAlreadyRegisteredToAVS();

    /// @notice Enum representing the status of an operator's registration with an AVS
    /// @dev Thrown when an invalid AVS is provided.
    error InvalidAVS();
    /// @dev Thrown when an invalid operator is provided.
    error InvalidOperator();
    /// @dev Thrown when an invalid operator set is provided.
    error InvalidOperatorSet();
    /// @dev Thrown when a strategy is already added to an operator set.
    error StrategyAlreadyInOperatorSet();
    /// @dev Thrown when a strategy is not in an operator set.
    error StrategyNotInOperatorSet();

    /// @dev Thrown when attempting to spend a spent eip-712 salt.
    error SaltSpent();
    /// @dev Thrown when attempting to use an expired eip-712 signature.
    error SignatureExpired();
}

interface IAVSDirectoryTypes {
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
}

interface IAVSDirectoryEvents is IAVSDirectoryTypes {
    /**
     *  @notice Emitted when an operator's registration status with an AVS id udpated
     *  @notice Only used by legacy M2 AVSs that have not integrated with operatorSets.
     */
    event OperatorAVSRegistrationStatusUpdated(
        address indexed operator, address indexed avs, OperatorAVSRegistrationStatus status
    );

    /// @notice Emitted when an AVS updates their metadata URI (Uniform Resource Identifier).
    /// @dev The URI is never stored; it is simply emitted through an event for off-chain indexing.
    event AVSMetadataURIUpdated(address indexed avs, string metadataURI);

    /// @notice Emitted when an AVS migrates to using operator sets.
    event AVSMigratedToOperatorSets(address indexed avs);

    /// @notice Emitted when an operator is migrated from M2 registration to operator sets.
    event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds);
}

interface IAVSDirectory is IAVSDirectoryEvents, IAVSDirectoryErrors, ISignatureUtils {
    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, and paused status.
     */
    function initialize(address initialOwner, IPauserRegistry _pauserRegistry, uint256 initialPausedStatus) external;

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
    function deregisterOperatorFromAVS(
        address operator
    ) external;

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */
    function operatorSaltIsSpent(address operator, bytes32 salt) external view returns (bool);

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

    /// @notice The EIP-712 typehash for the Registration struct used by the contract.
    function OPERATOR_AVS_REGISTRATION_TYPEHASH() external view returns (bytes32);

    /// @notice The EIP-712 typehash for the OperatorSetRegistration struct used by the contract.
    function OPERATOR_SET_REGISTRATION_TYPEHASH() external view returns (bytes32);
}
