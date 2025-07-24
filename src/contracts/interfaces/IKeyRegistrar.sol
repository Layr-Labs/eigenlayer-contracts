// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import {BN254} from "../libraries/BN254.sol";
import "./ISemVerMixin.sol";

interface IKeyRegistrarErrors {
    /// @notice Error thrown when a key is already registered
    error KeyAlreadyRegistered();
    /// @notice Error thrown when the key format is invalid
    error InvalidKeyFormat();
    /// @notice Error thrown when the address is zero
    error ZeroAddress();
    /// @notice Error thrown when the public key is zero
    error ZeroPubkey();
    /// @notice Error thrown when the curve type is invalid
    error InvalidCurveType();
    /// @notice Error thrown when the keypair is invalid
    error InvalidKeypair();
    /// @notice Error thrown when the configuration is already set
    error ConfigurationAlreadySet();
    /// @notice Error thrown when the operator set is not configured
    error OperatorSetNotConfigured();
    /// @notice Error thrown when the key is not found
    error KeyNotFound(OperatorSet operatorSet, address operator);
    /// @notice Error thrown when the operator is still slashable when trying to deregister a key
    error OperatorStillSlashable(OperatorSet operatorSet, address operator);
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
        bytes keyData; // Flexible storage for different curve types
    }
}

interface IKeyRegistrarEvents is IKeyRegistrarTypes {
    /// @notice Emitted when a key is registered
    event KeyRegistered(OperatorSet operatorSet, address indexed operator, CurveType curveType, bytes pubkey);
    /// @notice Emitted when a key is deregistered
    event KeyDeregistered(OperatorSet operatorSet, address indexed operator, CurveType curveType);
    /// @notice Emitted when the aggregate BN254 key is updated
    event AggregateBN254KeyUpdated(OperatorSet operatorSet, BN254.G1Point newAggregateKey);
    /// @notice Emitted when an operator set is configured
    event OperatorSetConfigured(OperatorSet operatorSet, CurveType curveType);
}

/// @notice The `KeyRegistrar` is used by AVSs to set their key type and by operators to register and deregister keys to operatorSets
/// @notice The integration pattern is as follows:
/// 1. The AVS calls `configureOperatorSet` to set the key type for their operatorSet
/// 2. Operators call `registerKey` to register their keys to the operatorSet
/// @dev This contract requires that keys are unique across all operatorSets, globally
/// @dev For the multichain protocol, the key type of the operatorSet must be set in the `KeyRegistrar`, but the
///      AVS is not required to use the KeyRegistrar for operator key management and can implement its own registry
interface IKeyRegistrar is IKeyRegistrarErrors, IKeyRegistrarEvents, ISemVerMixin {
    /**
     * @notice Configures an operator set with curve type
     * @param operatorSet The operator set to configure
     * @param curveType Type of curve (ECDSA, BN254)
     * @dev Only authorized callers for the AVS can configure operator sets
     */
    function configureOperatorSet(OperatorSet memory operatorSet, CurveType curveType) external;

    /**
     * @notice Registers a cryptographic key for an operator with a specific operator set
     * @param operator Address of the operator to register key for
     * @param operatorSet The operator set to register the key for
     * @param pubkey Public key bytes. For ECDSA, this is the address of the key. For BN254, this is the G1 and G2 key combined (see `encodeBN254KeyData`)
     * @param signature Signature proving ownership. For ECDSA this is a signature of the `getECDSAKeyRegistrationMessageHash`. For BN254 this is a signature of the `getBN254KeyRegistrationMessageHash`.
     * @dev Can be called by operator directly or by addresses they've authorized via the `PermissionController`
     * @dev Reverts if key is already registered
     * @dev There exist no restriction on the state of the operator with respect to the operatorSet. That is, an operator
     *      does not have to be registered for the operator in the `AllocationManager` to register a key for it
     * @dev For ECDSA, we allow a smart contract to be the pubkey (via ERC1271 signatures), but note that the multichain protocol DOES NOT support smart contract signatures
     */
    function registerKey(
        address operator,
        OperatorSet memory operatorSet,
        bytes calldata pubkey,
        bytes calldata signature
    ) external;

    /**
     * @notice Deregisters a cryptographic key for an operator with a specific operator set
     * @param operator Address of the operator to deregister key for
     * @param operatorSet The operator set to deregister the key from
     * @dev Can be called by the operator directly or by addresses they've authorized via the `PermissionController`
     * @dev Reverts if key was not registered
     * @dev Reverts if operator is still slashable for the operator set (prevents key rotation while slashable)
     * @dev Keys remain in global key registry to prevent reuse
     */
    function deregisterKey(address operator, OperatorSet memory operatorSet) external;

    /**
     * @notice Checks if a key is registered for an operator with a specific operator set
     * @param operatorSet The operator set to check
     * @param operator Address of the operator
     * @return True if the key is registered, false otherwise
     * @dev If the operatorSet is not configured, this function will return false
     */
    function isRegistered(OperatorSet memory operatorSet, address operator) external view returns (bool);

    /**
     * @notice Gets the curve type for an operator set
     * @param operatorSet The operator set to get the curve type for
     * @return The curve type, either ECDSA, BN254, or NONE
     */
    function getOperatorSetCurveType(
        OperatorSet memory operatorSet
    ) external view returns (CurveType);

    /**
     * @notice Gets the BN254 public key for an operator with a specific operator set
     * @param operatorSet The operator set to get the key for
     * @param operator Address of the operator
     * @return g1Point The BN254 G1 public key
     * @return g2Point The BN254 G2 public key
     * @dev Reverts if the operatorSet is not configured for BN254
     */
    function getBN254Key(
        OperatorSet memory operatorSet,
        address operator
    ) external view returns (BN254.G1Point memory g1Point, BN254.G2Point memory g2Point);

    /**
     * @notice Gets the ECDSA public key for an operator with a specific operator set as bytes
     * @param operatorSet The operator set to get the key for
     * @param operator Address of the operator
     * @return pubkey The ECDSA public key in bytes format
     * @dev Reverts if the operatorSet is not configured for ECDSA
     */
    function getECDSAKey(OperatorSet memory operatorSet, address operator) external view returns (bytes memory);

    /**
     * @notice Gets the ECDSA public key for an operator with a specific operator set
     * @param operatorSet The operator set to get the key for
     * @param operator Address of the operator
     * @return pubkey The ECDSA public key in address format
     * @dev Reverts if the operatorSet is not configured for ECDSA
     */
    function getECDSAAddress(OperatorSet memory operatorSet, address operator) external view returns (address);

    /**
     * @notice Checks if a key hash is globally registered
     * @param keyHash Hash of the key
     * @return True if the key is globally registered
     */
    function isKeyGloballyRegistered(
        bytes32 keyHash
    ) external view returns (bool);

    /**
     * @notice Gets the key hash for an operator with a specific operator set
     * @param operatorSet The operator set to get the key hash for
     * @param operator Address of the operator
     * @return keyHash The key hash
     * @dev Reverts if the operatorSet is not configured
     */
    function getKeyHash(OperatorSet memory operatorSet, address operator) external view returns (bytes32);

    /**
     * @notice Gets the operator from signing key
     * @param operatorSet The operator set to get the operator for
     * @param keyData The key data. For ECDSA, this is the signing key address. For BN254, this can be either the G1 key or the G1 and G2 key combined.
     * @return operator. Returns 0x0 if the key is not registered
     * @return status registration status. Returns false if the key is not registered
     * @dev This function decodes the key data based on the curve type of the operator set
     * @dev This function will return the operator address even if the operator is not registered for the operator set
     */
    function getOperatorFromSigningKey(
        OperatorSet memory operatorSet,
        bytes memory keyData
    ) external view returns (address, bool);

    /**
     * @notice Returns the message hash for ECDSA key registration, which must be signed by the operator when registering an ECDSA key
     * @param operator The operator address
     * @param operatorSet The operator set
     * @param keyAddress The address of the key
     * @return The message hash for signing
     */
    function getECDSAKeyRegistrationMessageHash(
        address operator,
        OperatorSet memory operatorSet,
        address keyAddress
    ) external view returns (bytes32);

    /**
     * @notice Returns the message hash for BN254 key registration, which must be signed by the operator when registering a BN254 key
     * @param operator The operator address
     * @param operatorSet The operator set
     * @param keyData The BN254 key data
     * @return The message hash for signing
     */
    function getBN254KeyRegistrationMessageHash(
        address operator,
        OperatorSet memory operatorSet,
        bytes calldata keyData
    ) external view returns (bytes32);

    /**
     * @notice Encodes the BN254 key data into a bytes array
     * @param g1Point The BN254 G1 public key
     * @param g2Point The BN254 G2 public key
     * @return The encoded key data
     */
    function encodeBN254KeyData(
        BN254.G1Point memory g1Point,
        BN254.G2Point memory g2Point
    ) external pure returns (bytes memory);
}
