// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import {BN254} from "../libraries/BN254.sol";

interface IKeyRegistrarErrors {
    /// @notice Error thrown when a key is already registered
    /// @dev Error code: 0x18f78402
    /// @dev We prevent duplicate key registrations to maintain global key uniqueness and avoid conflicting operator-key mappings
    error KeyAlreadyRegistered();

    /// @notice Error thrown when an operator is already registered
    /// @dev Error code: 0x42ee68b5
    /// @dev We prevent duplicate operator registrations to prevent re-registrations with a new key
    error OperatorAlreadyRegistered();

    /// @notice Error thrown when the key format is invalid
    /// @dev Error code: 0xd1091181
    /// @dev We enforce proper key formats (20 bytes for ECDSA, valid G1/G2 points for BN254) to ensure cryptographic validity and prevent malformed key data
    error InvalidKeyFormat();

    /// @notice Error thrown when the address is zero
    /// @dev Error code: 0xd92e233d
    error ZeroAddress();

    /// @notice Error thrown when the public key is zero
    /// @dev Error code: 0x4935505f
    error ZeroPubkey();

    /// @notice Error thrown when the curve type is invalid
    /// @dev Error code: 0xfdea7c09
    /// @dev We require valid curve types (ECDSA or BN254)
    error InvalidCurveType();

    /// @notice Error thrown when the keypair is invalid
    /// @dev Error code: 0x1b56a68b
    error InvalidKeypair();

    /// @notice Error thrown when the configuration is already set
    /// @dev Error code: 0x0081f09f
    /// @dev We prevent reconfiguration of operator sets to maintain consistency and avoid conflicting curve type settings
    error ConfigurationAlreadySet();

    /// @notice Error thrown when the operator set is not configured
    /// @dev Error code: 0xb9a620da
    /// @dev We require operator sets to be configured before key operations to ensure proper curve type validation and prevent operations on unconfigured sets
    error OperatorSetNotConfigured();

    /// @notice Error thrown when the key is not found
    /// @dev Error code: 0x2e40e187
    /// @dev We require existing key registrations for deregistration operations to ensure meaningful state changes and prevent operations on non-existent keys
    error KeyNotFound(OperatorSet operatorSet, address operator);

    /// @notice Error thrown when the operator is still slashable when trying to deregister a key
    /// @dev Error code: 0x10702879
    /// @dev We prevent key deregistration while operators are slashable to avoid race conditions and ensure operators cannot escape slashing by deregistering keys
    error OperatorStillSlashable(OperatorSet operatorSet, address operator);

    /// @notice Error thrown when a pending rotation already exists
    /// @dev Error code: 0x7a6a3c5b
    error PendingRotationExists();

    /// @notice Error thrown when key rotation is disabled for an operator set
    /// @dev Error code: 0x2052ff0e
    /// @dev This occurs when the operator set's minimum rotation delay is set to the maximum value
    error RotationDisabled();
}

interface IKeyRegistrarTypes {
    /// @dev Enum defining supported curve types
    enum CurveType {
        NONE,
        ECDSA,
        BN254
    }

    /// @dev Structure to store key information
    struct KeyInfo {
        bool isRegistered;
        /// @dev The currently active key bytes (20 for ECDSA, 192 for BN254)
        bytes currentKey;
        /// @dev The pending key bytes (same format as currentKey). Empty when no rotation scheduled
        bytes pendingKey;
        /// @dev Activation timestamp for pendingKey. Zero when no rotation scheduled
        uint64 pendingActivateAt;
    }
}

interface IKeyRegistrarEvents is IKeyRegistrarTypes {
    /// @notice Emitted when a key is registered
    event KeyRegistered(OperatorSet operatorSet, address indexed operator, CurveType curveType, bytes pubkey);
    /// @notice Emitted when a key is deregistered
    event KeyDeregistered(OperatorSet operatorSet, address indexed operator, CurveType curveType);
    /// @notice Emitted when a key rotation is scheduled
    event KeyRotationScheduled(
        OperatorSet operatorSet,
        address indexed operator,
        CurveType curveType,
        bytes oldPubkey,
        bytes newPubkey,
        uint64 activateAt
    );
    /// @notice Emitted when the aggregate BN254 key is updated
    event AggregateBN254KeyUpdated(OperatorSet operatorSet, BN254.G1Point newAggregateKey);
    /// @notice Emitted when an operator set is configured
    event OperatorSetConfigured(OperatorSet operatorSet, CurveType curveType);

    /// @notice Emitted when the minimum key rotation delay is set for an operator set
    event MinKeyRotationDelaySet(OperatorSet operatorSet, uint64 minDelay);
}

/// @notice The `KeyRegistrar` is used by AVSs to set their key type and by operators to register and deregister keys to operatorSets
/// @notice The integration pattern is as follows:
/// 1. The AVS calls `configureOperatorSet` to set the key type for their operatorSet
/// 2. Operators call `registerKey` to register their keys to the operatorSet
/// @dev This contract requires that keys are unique across all operatorSets, globally
/// @dev For the multichain protocol, the key type of the operatorSet must be set in the `KeyRegistrar`, but the
///      AVS is not required to use the KeyRegistrar for operator key management and can implement its own registry
interface IKeyRegistrar is IKeyRegistrarErrors, IKeyRegistrarEvents {
    /**
     * @notice Configures an operator set with curve type
     * @param operatorSet The operator set to configure
     * @param curveType Type of curve (ECDSA, BN254)
     * @dev Only authorized callers for the AVS can configure operator sets
     * @dev This function sets the minimum rotation delay to type(uint64).max, effectively disabling key rotation.
     *      If key rotation is desired, use `configureOperatorSetWithMinDelay` instead.
     * @dev Consider using `configureOperatorSetWithMinDelay` for new integrations to enable key rotation flexibility.
     * @dev Reverts for:
     *      - InvalidPermissions: Caller is not authorized for the AVS (via the PermissionController)
     *      - InvalidCurveType: The curve type is not ECDSA or BN254
     *      - ConfigurationAlreadySet: The operator set is already configured
     * @dev Emits the following events:
     *      - OperatorSetConfigured: When the operator set is successfully configured with a curve type
     *      - MinKeyRotationDelaySet: With delay set to type(uint64).max (rotation disabled)
     */
    function configureOperatorSet(OperatorSet memory operatorSet, CurveType curveType) external;

    /**
     * @notice Configures an operator set with curve type and minimum rotation delay
     * @param operatorSet The operator set to configure
     * @param curveType Type of curve (ECDSA, BN254)
     * @param minDelaySeconds Minimum delay in seconds before a rotation can activate. Set to type(uint64).max to disable rotation
     * @dev Only authorized callers for the AVS can configure operator sets
     * @dev Reverts for:
     *      - InvalidPermissions: Caller is not authorized for the AVS (via the PermissionController)
     *      - InvalidCurveType: The curve type is not ECDSA or BN254
     *      - ConfigurationAlreadySet: The operator set is already configured
     * @dev Emits the following events:
     *      - OperatorSetConfigured: When the operator set is successfully configured with a curve type
     *      - MinKeyRotationDelaySet: When the minimum rotation delay is set
     */
    function configureOperatorSetWithMinDelay(
        OperatorSet memory operatorSet,
        CurveType curveType,
        uint64 minDelaySeconds
    ) external;

    /**
     * @notice Registers a cryptographic key for an operator with a specific operator set
     * @param operator Address of the operator to register key for
     * @param operatorSet The operator set to register the key for
     * @param pubkey Public key bytes. For ECDSA, this is the address of the key. For BN254, this is the G1 and G2 key combined (see `encodeBN254KeyData`)
     * @param signature Signature proving ownership. For ECDSA this is a signature of the `getECDSAKeyRegistrationMessageHash`. For BN254 this is a signature of the `getBN254KeyRegistrationMessageHash`.
     * @dev Can be called by operator directly or by addresses they've authorized via the `PermissionController`
     * @dev There exist no restriction on the state of the operator with respect to the operatorSet. That is, an operator
     *      does not have to be registered for the operator in the `AllocationManager` to register a key for it
     * @dev For ECDSA, we allow a smart contract to be the pubkey (via ERC1271 signatures), but note that the multichain protocol DOES NOT support smart contract signatures
     * @dev Reverts for:
     *      - InvalidPermissions: Caller is not the operator or authorized via the PermissionController
     *      - OperatorSetNotConfigured: The operator set is not configured
     *      - OperatorAlreadyRegistered: The operator is already registered for the operatorSet in the KeyRegistrar
     *      - InvalidKeyFormat: For ECDSA: The key is not exactly 20 bytes
     *      - ZeroPubkey: For ECDSA: The key is the zero address
     *      - KeyAlreadyRegistered: For ECDSA: The key is already registered globally by hash
     *      - InvalidSignature: For ECDSA: The signature is not valid
     *      - InvalidKeyFormat: For BN254: The key data is not exactly 192 bytes
     *      - InvalidSignature: For BN254: The signature is not exactly 64 bytes
     *      - ZeroPubkey: For BN254: The G1 point is the zero point
     *      - InvalidSignature: For BN254: The signature is not valid
     *      - KeyAlreadyRegistered: For BN254: The key is already registered globally by hash
     * @dev Emits the following events:
     *      - KeyRegistered: When the key is successfully registered for the operator and operatorSet
     */
    function registerKey(
        address operator,
        OperatorSet memory operatorSet,
        bytes calldata pubkey,
        bytes calldata signature
    ) external;

    /// @notice Deregisters a cryptographic key for an operator with a specific operator set
    /// @param operator Address of the operator to deregister key for
    /// @param operatorSet The operator set to deregister the key from
    /// @dev Can be called by the operator directly or by addresses they've authorized via the `PermissionController`
    /// @dev Keys remain in global key registry to prevent reuse
    /// @dev Reverts for:
    ///      - InvalidPermissions: Caller is not authorized for the operator (via the PermissionController)
    ///      - OperatorStillSlashable: The operator is still slashable for the AVS
    ///      - OperatorSetNotConfigured: The operator set is not configured
    ///      - KeyNotFound: The operator does not have a registered key for this operator set
    /// @dev Emits the following events:
    ///      - KeyDeregistered: When the key is successfully deregistered for the operator and operatorSet
    function deregisterKey(
        address operator,
        OperatorSet memory operatorSet
    ) external;

    /**
     * @notice Rotates an operator's key for an operator set, replacing the current key with a new key
     * @param operator Address of the operator whose key is being rotated
     * @param operatorSet The operator set for which the key is being rotated
     * @param newPubkey New public key bytes. For ECDSA, this is the address of the key. For BN254, this is the G1 and G2 key combined (see `encodeBN254KeyData`)
     * @param signature Signature from the new key proving ownership over the appropriate registration message hash
     * @dev The new key will activate at block.timestamp + the minimum rotation delay configured for the operator set
     * @dev Keys remain in the global key registry to prevent reuse
     * @dev There is no slashability restriction for rotation; operators may rotate while slashable
     * @dev Reverts for:
     *      - InvalidPermissions: Caller is not authorized for the operator (via the PermissionController)
     *      - OperatorSetNotConfigured: The operator set is not configured
     *      - KeyNotFound: The operator does not have a registered key for this operator set
     *      - PendingRotationExists: A rotation is already scheduled and has not yet activated
     *      - RotationDisabled: Key rotation is disabled for this operator set (minDelay set to type(uint64).max)
     *      - InvalidKeyFormat / ZeroPubkey / InvalidSignature: New key data/signature invalid per curve type
     *      - KeyAlreadyRegistered: New key is already globally registered
     * @dev Emits the following event:
     *      - KeyRotationScheduled: When the rotation is successfully scheduled
     */
    function rotateKey(
        address operator,
        OperatorSet memory operatorSet,
        bytes calldata newPubkey,
        bytes calldata signature
    ) external;

    /**
     * @notice Sets the minimum allowed rotation delay for an operator set
     * @param operatorSet The operator set to configure
     * @param minDelaySeconds The minimum rotation delay in seconds
     * @dev Only callable by the AVS or its authorized caller via the PermissionController
     */
    function setMinKeyRotationDelay(OperatorSet memory operatorSet, uint64 minDelaySeconds) external;

    /**
     * @notice Finalizes a scheduled rotation if activation time has passed, compacting storage
     * @notice This is optional, getters already return the correct active key based on time
     * @param operator The operator address
     * @param operatorSet The operator set
     * @return success True if a pending rotation was finalized
     */
    function finalizeScheduledRotation(
        address operator,
        OperatorSet memory operatorSet
    ) external returns (bool success);

    /**
     * @notice Checks if a key is registered for an operator with a specific operator set
     * @param operatorSet The operator set to check
     * @param operator Address of the operator
     * @return True if the key is registered, false otherwise
     * @dev If the operatorSet is not configured, this function will return false
     */
    function isRegistered(OperatorSet memory operatorSet, address operator) external view returns (bool);

    /// @notice Gets the curve type for an operator set
    /// @param operatorSet The operator set to get the curve type for
    /// @return The curve type, either ECDSA, BN254, or NONE
    function getOperatorSetCurveType(
        OperatorSet memory operatorSet
    ) external view returns (CurveType);

    /// @notice Gets the BN254 public key for an operator with a specific operator set
    /// @param operatorSet The operator set to get the key for
    /// @param operator Address of the operator
    /// @return g1Point The BN254 G1 public key
    /// @return g2Point The BN254 G2 public key
    /// @dev Reverts for:
    ///      - InvalidCurveType: The operatorSet is not configured for BN254
    /// @dev Returns empty points if the operator has not registered a key for the operatorSet. We
    ///      recommend calling `isRegistered` first to check if the operator has a key registered
    function getBN254Key(
        OperatorSet memory operatorSet,
        address operator
    ) external view returns (BN254.G1Point memory g1Point, BN254.G2Point memory g2Point);

    /// @notice Gets the ECDSA public key for an operator with a specific operator set as bytes
    /// @param operatorSet The operator set to get the key for
    /// @param operator Address of the operator
    /// @return pubkey The ECDSA public key in bytes format
    /// @dev Reverts for:
    ///      - InvalidCurveType: The operatorSet is not configured for ECDSA
    /// @dev Returns 0x0 if the operator has not registered a key for the operatorSet. We
    ///      recommend calling `isRegistered` first to check if the operator has a key registered
    function getECDSAKey(
        OperatorSet memory operatorSet,
        address operator
    ) external view returns (bytes memory);

    /// @notice Gets the ECDSA public key for an operator with a specific operator set
    /// @param operatorSet The operator set to get the key for
    /// @param operator Address of the operator
    /// @return pubkey The ECDSA public key in address format
    /// @dev Reverts for:
    ///      - InvalidCurveType: The operatorSet is not configured for ECDSA
    /// @dev Returns 0x0 if the operator has not registered a key for the operatorSet. We
    ///      recommend calling `isRegistered` first to check if the operator has a key registered
    function getECDSAAddress(
        OperatorSet memory operatorSet,
        address operator
    ) external view returns (address);

    /// @notice Checks if a key hash is globally registered
    /// @param keyHash Hash of the key
    /// @return True if the key is globally registered
    function isKeyGloballyRegistered(
        bytes32 keyHash
    ) external view returns (bool);

    /// @notice Gets the key hash for an operator with a specific operator set
    /// @param operatorSet The operator set to get the key hash for
    /// @param operator Address of the operator
    /// @return keyHash The key hash
    function getKeyHash(
        OperatorSet memory operatorSet,
        address operator
    ) external view returns (bytes32);

    /// @notice Gets the operator from signing key
    /// @param operatorSet The operator set to get the operator for
    /// @param keyData The key data. For ECDSA, this is the signing key address. For BN254, this can be either the G1 key or the G1 and G2 key combined.
    /// @return operator. Returns 0x0 if the key is not registered
    /// @return status registration status. Returns false if the key is not registered
    /// @dev This function decodes the key data based on the curve type of the operator set
    /// @dev This function will return the operator address even if the operator is not registered for the operator set
    /// @dev Reverts for:
    ///      - InvalidCurveType: The CurveType is not configured
    function getOperatorFromSigningKey(
        OperatorSet memory operatorSet,
        bytes memory keyData
    ) external view returns (address, bool);

    /// @notice Returns the message hash for ECDSA key registration, which must be signed by the key when registering an ECDSA key
    /// @param operator The operator address
    /// @param operatorSet The operator set
    /// @param keyAddress The address of the key
    /// @return The message hash for signing
    function getECDSAKeyRegistrationMessageHash(
        address operator,
        OperatorSet memory operatorSet,
        address keyAddress
    ) external view returns (bytes32);

    /// @notice Returns the message hash for BN254 key registration, which must be signed by the key when registering a BN254 key
    /// @param operator The operator address
    /// @param operatorSet The operator set
    /// @param keyData The BN254 key data
    /// @return The message hash for signing
    function getBN254KeyRegistrationMessageHash(
        address operator,
        OperatorSet memory operatorSet,
        bytes calldata keyData
    ) external view returns (bytes32);

    /// @notice Encodes the BN254 key data into a bytes array
    /// @param g1Point The BN254 G1 public key
    /// @param g2Point The BN254 G2 public key
    /// @return The encoded key data
    function encodeBN254KeyData(
        BN254.G1Point memory g1Point,
        BN254.G2Point memory g2Point
    ) external pure returns (bytes memory);
}
