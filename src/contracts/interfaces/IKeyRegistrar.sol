// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import {BN254} from "../libraries/BN254.sol";
import "./ISemVerMixin.sol";

interface IKeyRegistrarErrors {
    /// Key Management
    error KeyAlreadyRegistered();
    error InvalidKeyFormat();
    error ZeroAddress();
    error ZeroPubkey();
    error InvalidCurveType();
    error InvalidKeypair();
    error ConfigurationAlreadySet();
    error OperatorSetNotConfigured();
    error KeyNotFound(OperatorSet operatorSet, address operator);
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
    event KeyRegistered(OperatorSet operatorSet, address indexed operator, CurveType curveType, bytes pubkey);
    event KeyDeregistered(OperatorSet operatorSet, address indexed operator, CurveType curveType);
    event AggregateBN254KeyUpdated(OperatorSet operatorSet, BN254.G1Point newAggregateKey);
    event OperatorSetConfigured(OperatorSet operatorSet, CurveType curveType);
}

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
     * @param pubkey Public key bytes
     * @param signature Signature proving ownership (only needed for BN254 keys)
     * @dev Can be called by operator directly or by addresses they've authorized via PermissionController
     * @dev Reverts if key is already registered
     * @dev There exist no restriction on the state of the operator with respect to the operatorSet. That is, an operator
     *      does not have to be registered for the operator in the `AllocationManager` to register a key for it
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
     * @dev Can be called by the operator directly or by addresses they've authorized via PermissionController
     * @dev Reverts if key was not registered
     * @dev Reverts if operator is still slashable for the operator set (prevents key rotation while slashable)
     * @dev Keys remain in global key registry to prevent reuse
     */
    function deregisterKey(address operator, OperatorSet memory operatorSet) external;

    /**
     * @notice Checks if an operator has a registered key
     * @param operatorSet The operator set to check and update
     * @param operator Address of the operator
     * @return whether the operator has a registered key
     * @dev This function is called by the AVSRegistrar when an operator registers for an AVS
     * @dev Only authorized callers for the AVS can call this function
     * @dev Reverts if operator doesn't have a registered key for this operator set
     */
    function checkKey(OperatorSet memory operatorSet, address operator) external view returns (bool);

    /**
     * @notice Checks if a key is registered for an operator with a specific operator set
     * @param operatorSet The operator set to check
     * @param operator Address of the operator
     * @return True if the key is registered
     */
    function isRegistered(OperatorSet memory operatorSet, address operator) external view returns (bool);

    /**
     * @notice Gets the configuration for an operator set
     * @param operatorSet The operator set to get configuration for
     * @return The operator set configuration
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
     */
    function getBN254Key(
        OperatorSet memory operatorSet,
        address operator
    ) external view returns (BN254.G1Point memory g1Point, BN254.G2Point memory g2Point);

    /**
     * @notice Gets the ECDSA public key for an operator with a specific operator set as bytes
     * @param operatorSet The operator set to get the key for
     * @param operator Address of the operator
     * @return pubkey The ECDSA public key
     */
    function getECDSAKey(OperatorSet memory operatorSet, address operator) external view returns (bytes memory);

    /**
     * @notice Gets the ECDSA public key for an operator with a specific operator set
     * @param operatorSet The operator set to get the key for
     * @param operator Address of the operator
     * @return pubkey The ECDSA public key
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
     * @notice Returns the message hash for ECDSA key registration
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
     * @notice Returns the message hash for BN254 key registration
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
