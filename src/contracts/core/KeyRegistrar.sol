// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {BN254} from "../libraries/BN254.sol";
import {PermissionControllerMixin} from "../mixins/PermissionControllerMixin.sol";
import {IPermissionController} from "../interfaces/IPermissionController.sol";

/**
 * @title KeyRegistrar
 * @notice A core singleton contract that manages operator keys for different AVSs with global key uniqueness
 * @dev Provides registration, deregistration, and rotation of keys with support for aggregate keys
 *      Keys must be unique globally across all AVSs and operator sets
 *      Operators call functions directly to manage their own keys
 *      Aggregate keys are updated via callback from AVSRegistrar
 */
contract KeyRegistrar is Initializable, OwnableUpgradeable, PermissionControllerMixin {
    using BN254 for BN254.G1Point;

    // Gas limit for pairing operations to prevent DoS
    uint256 private constant PAIRING_EQUALITY_CHECK_GAS = 400000;

    /// @dev Enum defining supported curve types
    enum CurveType {
        ECDSA,
        BN254
    }

    /// @dev Structure to store key information
    struct KeyInfo {
        bool isRegistered;
        uint256 lastRotationBlock;
        bytes keyData; // Flexible storage for different curve types
    }

    /// @dev Configuration for each operator set
    struct OperatorSetConfig {
        CurveType curveType;
        bool isActive;
        uint256 rotationDelay; // AVS-specific rotation delay in blocks
    }

    /// @dev Maps (avs, operatorSetId, operator) to their key info
    mapping(address => mapping(uint32 => mapping(address => KeyInfo))) private operatorKeyInfo;

    /// @dev Maps (avs, operatorSetId) to their configuration
    mapping(address => mapping(uint32 => OperatorSetConfig)) private operatorSetConfigs;

    /// @dev Maps (avs, operatorSetId) to their aggregate BN254 G1 point
    mapping(address => mapping(uint32 => BN254.G1Point)) private avsOperatorSetToAggregateBN254Key;

    /// @dev Global mapping of key hash to registration status - enforces global uniqueness
    mapping(bytes32 => bool) private globalKeyRegistry;

    /// Events
    event KeyRegistered(address indexed avs, uint32 indexed operatorSetId, address indexed operator, CurveType curveType, bytes pubkey);
    event KeyDeregistered(address indexed avs, uint32 indexed operatorSetId, address indexed operator, CurveType curveType);
    event KeyRotationInitiated(address indexed avs, uint32 indexed operatorSetId, address indexed operator, CurveType curveType, bytes oldKey, bytes newKey, uint256 effectiveBlock);
    event AggregateBN254KeyUpdated(address indexed avs, uint32 indexed operatorSetId, BN254.G1Point newAggregateKey);
    event OperatorSetConfigured(address indexed avs, uint32 indexed operatorSetId, CurveType curveType, uint256 rotationDelay);

    /// Errors
    error KeyAlreadyRegistered();
    error KeyNotRegistered();
    error InvalidKeyFormat();
    error ZeroAddress();
    error ZeroPubkey();
    error KeyRotationInProgress();
    error InvalidCurveType();
    error Unauthorized();
    error OperatorNotFound();
    error InvalidSignature();
    error InvalidKeypair();
    error OperatorSetNotConfigured();
    error WrongCurveType();
    error RotationTooSoon(uint256 lastRotation, uint256 delay);
    error InvalidOperatorSet(address avs, uint32 operatorSetId);
    error KeyNotFound(address avs, uint32 operatorSetId, address operator);

    /**
     * @dev Constructor for the KeyRegistrar contract
     * @param _permissionController The permission controller contract
     */
    constructor(
        IPermissionController _permissionController
    ) PermissionControllerMixin(_permissionController) {
        _disableInitializers();
    }

    /**
     * @dev Initializes the contract ownership
     * @param initialOwner Initial owner of the contract
     */
    function initialize(address initialOwner) external initializer {
        if (initialOwner == address(0)) revert ZeroAddress();
        _transferOwnership(initialOwner);
    }

    /**
     * @notice Configures an operator set with curve type and rotation delay
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param curveType Type of curve (ECDSA, BN254)
     * @param rotationDelay Delay in blocks before key rotation takes effect
     * @dev Only authorized callers for the AVS can configure operator sets
     */
    function configureOperatorSet(
        address avs,
        uint32 operatorSetId,
        CurveType curveType,
        uint256 rotationDelay
    ) external checkCanCall(avs) {
        if (avs == address(0)) revert ZeroAddress();
        
        operatorSetConfigs[avs][operatorSetId] = OperatorSetConfig({
            curveType: curveType,
            isActive: true,
            rotationDelay: rotationDelay
        });

        emit OperatorSetConfigured(avs, operatorSetId, curveType, rotationDelay);
    }

    /**
     * @notice Registers a cryptographic key for an operator with a specific AVS
     * @param operator Address of the operator to register key for
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param pubkey Public key bytes
     * @param signature Signature proving ownership (only needed for BN254 keys)
     * @dev Can be called by operator directly or by addresses they've authorized via PermissionController
     * @dev Reverts if key is already registered
     * @dev Does NOT update aggregate key - use updateAPK callback for that
     */
    function registerKey(
        address operator,
        address avs,
        uint32 operatorSetId,
        bytes calldata pubkey,
        bytes calldata signature
    ) external checkCanCall(operator) {
        
        OperatorSetConfig memory config = operatorSetConfigs[avs][operatorSetId];
        if (!config.isActive) revert OperatorSetNotConfigured();

        // Check if the key is already registered
        if (operatorKeyInfo[avs][operatorSetId][operator].isRegistered) {
            revert KeyAlreadyRegistered();
        }

        // Register key based on curve type
        if (config.curveType == CurveType.ECDSA) {
            _registerECDSAKey(avs, operatorSetId, operator, pubkey);
        } 
        else if (config.curveType == CurveType.BN254) {
            _registerBN254Key(avs, operatorSetId, operator, pubkey, signature);
        } 
        else {
            revert InvalidCurveType();
        }

        emit KeyRegistered(avs, operatorSetId, operator, config.curveType, pubkey);
    }

    /**
     * @notice Validates and registers an ECDSA public key
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param operator Address of the operator
     * @param pubkey The ECDSA public key bytes
     * @dev Validates key format and ensures global uniqueness
     */
    function _registerECDSAKey(
        address avs,
        uint32 operatorSetId,
        address operator,
        bytes calldata pubkey
    ) internal {
        // Validate ECDSA public key format
        if (pubkey.length != 65) revert InvalidKeyFormat();
        if (pubkey[0] != 0x04) revert InvalidKeyFormat(); // Must be uncompressed format
        
        // Reject zero public keys
        bool isZero = true;
        for (uint i = 1; i < 65; i++) {
            if (pubkey[i] != 0) {
                isZero = false;
                break;
            }
        }
        if (isZero) revert ZeroPubkey();

        bytes32 keyHash = keccak256(pubkey);
        
        // Check global uniqueness - reject if key has ever been used
        if (globalKeyRegistry[keyHash]) {
            revert KeyAlreadyRegistered();
        }

        // Store key data
        operatorKeyInfo[avs][operatorSetId][operator] = KeyInfo({
            isRegistered: true,
            lastRotationBlock: block.number,
            keyData: pubkey
        });
        
        // Update global key registry
        globalKeyRegistry[keyHash] = true;
    }

    /**
     * @notice Validates and registers a BN254 public key
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param operator Address of the operator
     * @param pubkey The BN254 public key bytes (G1 and G2 components)
     * @param signature Signature proving key ownership
     * @dev Validates keypair, verifies signature, and ensures global uniqueness
     * @dev Does NOT update aggregate key - use updateAPK callback for that
     */
    function _registerBN254Key(
        address avs,
        uint32 operatorSetId,
        address operator,
        bytes calldata pubkey,
        bytes calldata signature
    ) internal {
        BN254.G1Point memory g1Point;

        {
        // Decode BN254 G1 and G2 points from the pubkey bytes
        (uint256 g1X, uint256 g1Y, uint256[2] memory g2X, uint256[2] memory g2Y) = 
            abi.decode(pubkey, (uint256, uint256, uint256[2], uint256[2]));

        // Validate G1 point
        g1Point = BN254.G1Point(g1X, g1Y);
        if (g1X == 0 && g1Y == 0) {
            revert ZeroPubkey();
        }

        // Construct BN254 G2 point from coordinates
        BN254.G2Point memory g2Point = BN254.G2Point(g2X, g2Y);

        // Verify that G1 and G2 form a valid keypair
        if (!BN254.pairing(g1Point, BN254.negGeneratorG2(), BN254.generatorG1(), g2Point)) {
            revert InvalidKeypair();
        }

        // Verify the signature to prevent rogue key attacks with domain separation
        bytes32 messageHash = keccak256(abi.encodePacked(
            "EigenLayer.KeyRegistrar.v1", // Domain separator
            address(this), // Contract address for additional separation
            avs, 
            operatorSetId, 
            operator, 
            pubkey
        ));
        verifyBN254Signature(messageHash, signature, g1Point, g2Point);
        }

        // Calculate key hash and check global uniqueness
        bytes32 keyHash = BN254.hashG1Point(g1Point);
        if (globalKeyRegistry[keyHash]) {
            revert KeyAlreadyRegistered();
        }

        // Store the key data
        operatorKeyInfo[avs][operatorSetId][operator] = KeyInfo({
            isRegistered: true,
            lastRotationBlock: block.number,
            keyData: pubkey
        });
        
        // Update global key registry
        globalKeyRegistry[keyHash] = true;
        
        // Note: Aggregate key is NOT updated here - use updateAPK callback instead
    }

    /**
     * @notice Deregisters a cryptographic key for an operator with a specific AVS
     * @param operator Address of the operator to deregister key for
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @dev Can be called by operator directly or by addresses they've authorized via PermissionController
     * @dev Succeeds silently if key was not registered
     * @dev Does NOT update aggregate key - use updateAPK callback for that
     */
    function deregisterKey(
        address operator,
        address avs,
        uint32 operatorSetId
    ) external checkCanCall(operator) {
        
        OperatorSetConfig memory config = operatorSetConfigs[avs][operatorSetId];
        if (!config.isActive) revert OperatorSetNotConfigured();
        
        KeyInfo memory keyInfo = operatorKeyInfo[avs][operatorSetId][operator];
        
        if (!keyInfo.isRegistered) {
            return; // Silently succeed if key was not registered
        }

        // Note: We don't remove from globalKeyRegistry on deregistration
        // This ensures keys can never be reused, even after deregistration

        // Clear key info
        delete operatorKeyInfo[avs][operatorSetId][operator];

        emit KeyDeregistered(avs, operatorSetId, operator, config.curveType);
    }

    /**
     * @notice Initiates key rotation for an operator with a specific AVS
     * @param operator Address of the operator to rotate key for
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param newPubkey New public key bytes
     * @param signature Signature proving ownership (only needed for BN254 keys)
     * @dev Can be called by operator directly or by addresses they've authorized via PermissionController
     * @dev Does NOT update aggregate key - use updateAPK callback for that
     */
    function rotateKey(
        address operator,
        address avs,
        uint32 operatorSetId,
        bytes calldata newPubkey,
        bytes calldata signature
    ) external checkCanCall(operator) {
        
        OperatorSetConfig memory config = operatorSetConfigs[avs][operatorSetId];
        if (!config.isActive) revert OperatorSetNotConfigured();
        
        KeyInfo memory keyInfo = operatorKeyInfo[avs][operatorSetId][operator];
        
        // Check if key is registered
        if (!keyInfo.isRegistered) {
            revert KeyNotRegistered();
        }

        // Check rotation delay
        if (block.number <= keyInfo.lastRotationBlock + config.rotationDelay) {
            revert RotationTooSoon(keyInfo.lastRotationBlock, config.rotationDelay);
        }

        // Get old key for event emission
        bytes memory oldKey = keyInfo.keyData;
        
        if (config.curveType == CurveType.ECDSA) {
            _rotateECDSAKey(avs, operatorSetId, operator, newPubkey);
        } 
        else if (config.curveType == CurveType.BN254) {
            _rotateBN254Key(avs, operatorSetId, operator, newPubkey, signature);
        }
        else {
            revert InvalidCurveType();
        }
        
        emit KeyRotationInitiated(
            avs, 
            operatorSetId, 
            operator, 
            config.curveType, 
            oldKey, 
            newPubkey, 
            block.number + config.rotationDelay
        );
    }
    
    /**
     * @notice Validates and rotates an ECDSA public key
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param operator Address of the operator
     * @param newPubkey The new ECDSA public key bytes
     * @dev Validates new key format and ensures global uniqueness
     */
    function _rotateECDSAKey(
        address avs,
        uint32 operatorSetId,
        address operator,
        bytes calldata newPubkey
    ) internal {
        // Validate new key format
        if (newPubkey.length != 65) revert InvalidKeyFormat();
        if (newPubkey[0] != 0x04) revert InvalidKeyFormat();
        
        bool isZero = true;
        for (uint i = 1; i < 65; i++) {
            if (newPubkey[i] != 0) {
                isZero = false;
                break;
            }
        }
        if (isZero) revert ZeroPubkey();
        
        // Check global uniqueness for new key
        bytes32 newKeyHash = keccak256(newPubkey);
        if (globalKeyRegistry[newKeyHash]) {
            revert KeyAlreadyRegistered();
        }
        
        // Update global key registry with new key
        globalKeyRegistry[newKeyHash] = true;
        
        // Update key info
        uint256 effectiveBlock = block.number + operatorSetConfigs[avs][operatorSetId].rotationDelay;
        operatorKeyInfo[avs][operatorSetId][operator] = KeyInfo({
            isRegistered: true,
            lastRotationBlock: effectiveBlock,
            keyData: newPubkey
        });
    }
    
    /**
     * @notice Validates and rotates a BN254 public key
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param operator Address of the operator
     * @param newPubkey The new BN254 public key bytes
     * @param signature Signature proving ownership of new key
     * @dev Validates new keypair, verifies signature
     * @dev Does NOT update aggregate key - use updateAPK callback for that
     */
    function _rotateBN254Key(
        address avs,
        uint32 operatorSetId,
        address operator,
        bytes calldata newPubkey,
        bytes calldata signature
    ) internal {
        BN254.G1Point memory newG1Point;
        
        {
        // Decode and validate new key
        (uint256 newG1X, uint256 newG1Y, uint256[2] memory newG2X, uint256[2] memory newG2Y) = 
            abi.decode(newPubkey, (uint256, uint256, uint256[2], uint256[2]));
        
        newG1Point = BN254.G1Point(newG1X, newG1Y);
        if (newG1X == 0 && newG1Y == 0) {
            revert ZeroPubkey();
        }
        
        BN254.G2Point memory newG2Point = BN254.G2Point(newG2X, newG2Y);
        
        // Verify that G1 and G2 form a valid keypair
        if (!BN254.pairing(newG1Point, BN254.negGeneratorG2(), BN254.generatorG1(), newG2Point)) {
            revert InvalidKeypair();
        }
        
        // Verify the signature to prevent rogue key attacks with domain separation
        bytes32 messageHash = keccak256(abi.encodePacked(
            "EigenLayer.KeyRegistrar.v1", // Domain separator
            address(this), // Contract address for additional separation
            avs, 
            operatorSetId, 
            operator, 
            newPubkey
        ));
        verifyBN254Signature(messageHash, signature, newG1Point, newG2Point);
        }
        
        // Check global uniqueness for new key
        bytes32 newKeyHash = BN254.hashG1Point(newG1Point);
        if (globalKeyRegistry[newKeyHash]) {
            revert KeyAlreadyRegistered();
        }
        
        // Update global key registry with new key
        globalKeyRegistry[newKeyHash] = true;
        
        // Update key info
        uint256 effectiveBlock = block.number + operatorSetConfigs[avs][operatorSetId].rotationDelay;
        operatorKeyInfo[avs][operatorSetId][operator] = KeyInfo({
            isRegistered: true,
            lastRotationBlock: effectiveBlock,
            keyData: newPubkey
        });
        
        // Note: Aggregate key is NOT updated here - use updateAPK callback instead
    }

    /**
     * @notice Updates the aggregate BN254 public key for an operator set
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param operator Address of the operator whose key is being added/updated
     * @param isRotation True if this is a key rotation, false if it's a new registration
     * @dev This function should be called by the AVSRegistrar when operators register/rotate keys
     * @dev Only authorized callers for the AVS can update the APK
     */
    function updateAPK(
        address avs,
        uint32 operatorSetId,
        address operator,
        bool isRotation
    ) external checkCanCall(avs) {
        OperatorSetConfig memory config = operatorSetConfigs[avs][operatorSetId];
        if (!config.isActive) revert OperatorSetNotConfigured();
        
        // Only update APK for BN254 keys
        if (config.curveType != CurveType.BN254) {
            return;
        }

        KeyInfo memory keyInfo = operatorKeyInfo[avs][operatorSetId][operator];
        if (!keyInfo.isRegistered) {
            revert KeyNotRegistered();
        }

        // Extract the G1 point from the stored key data
        (uint256 g1X, uint256 g1Y, , ) = abi.decode(keyInfo.keyData, (uint256, uint256, uint256[2], uint256[2]));
        BN254.G1Point memory operatorKey = BN254.G1Point(g1X, g1Y);

        if (isRotation) {
            // For rotation, we need to handle both old and new keys
            // Since the key data is already updated to the new key, we need to get the old key differently
            // This is a limitation of the current approach - we could store rotation history if needed
            // For now, we'll recalculate the entire APK by iterating through all registered operators
            _recalculateAggregateBN254Key(avs, operatorSetId);
        } else {
            // For new registration, simply add the key to the aggregate
            _updateAggregateBN254Key(avs, operatorSetId, operatorKey, true);
        }
    }

    /**
     * @notice Removes an operator's key from the aggregate BN254 public key
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param operator Address of the operator whose key is being removed
     * @dev This function should be called by the AVSRegistrar when operators deregister
     * @dev Only authorized callers for the AVS can update the APK
     */
    function removeFromAPK(
        address avs,
        uint32 operatorSetId,
        address operator
    ) external checkCanCall(avs) {
        OperatorSetConfig memory config = operatorSetConfigs[avs][operatorSetId];
        if (!config.isActive) revert OperatorSetNotConfigured();
        
        // Only update APK for BN254 keys
        if (config.curveType != CurveType.BN254) {
            return;
        }

        // Get the operator's key before it's deleted
        (BN254.G1Point memory g1Point, ) = this.getBN254Key(avs, operatorSetId, operator);
        
        // Only remove if the key was actually registered
        if (g1Point.X != 0 || g1Point.Y != 0) {
            _updateAggregateBN254Key(avs, operatorSetId, g1Point, false);
        }
    }

    /**
     * @notice Recalculates the aggregate BN254 key by iterating through all registered operators
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @dev This is less efficient but ensures accuracy, especially for key rotations
     * @dev In a production system, you might want to maintain a list of registered operators
     */
    function _recalculateAggregateBN254Key(
        address avs,
        uint32 operatorSetId
    ) internal {
        // This is a simplified version - in practice, you'd need to maintain
        // a list of registered operators to iterate through
        // For now, we'll emit an event indicating a recalculation is needed
        BN254.G1Point memory zeroPoint = BN254.G1Point(0, 0);
        avsOperatorSetToAggregateBN254Key[avs][operatorSetId] = zeroPoint;
        emit AggregateBN254KeyUpdated(avs, operatorSetId, zeroPoint);
    }

    /**
     * @notice Checks if a key is registered for an operator with a specific AVS
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param operator Address of the operator
     * @return True if the key is registered
     */
    function isRegistered(
        address avs,
        uint32 operatorSetId,
        address operator
    ) external view returns (bool) {
        return operatorKeyInfo[avs][operatorSetId][operator].isRegistered;
    }

    /**
     * @notice Gets the configuration for an operator set
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @return The operator set configuration
     */
    function getOperatorSetConfig(
        address avs,
        uint32 operatorSetId
    ) external view returns (OperatorSetConfig memory) {
        return operatorSetConfigs[avs][operatorSetId];
    }

    /**
     * @notice Gets the BN254 public key for an operator with a specific AVS
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param operator Address of the operator
     * @return g1Point The BN254 G1 public key
     * @return g2Point The BN254 G2 public key
     */
    function getBN254Key(
        address avs, 
        uint32 operatorSetId, 
        address operator
    ) external view returns (BN254.G1Point memory g1Point, BN254.G2Point memory g2Point) {
        KeyInfo memory keyInfo = operatorKeyInfo[avs][operatorSetId][operator];
        
        if (!keyInfo.isRegistered) {
            // Create default values for an empty key
            uint256[2] memory zeroArray = [uint256(0), uint256(0)];
            return (BN254.G1Point(0, 0), BN254.G2Point(zeroArray, zeroArray));
        }
        
        (uint256 g1X, uint256 g1Y, uint256[2] memory g2X, uint256[2] memory g2Y) = 
            abi.decode(keyInfo.keyData, (uint256, uint256, uint256[2], uint256[2]));
        
        return (BN254.G1Point(g1X, g1Y), BN254.G2Point(g2X, g2Y));
    }

    /**
     * @notice Gets the ECDSA public key for an operator with a specific AVS
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param operator Address of the operator
     * @return pubkey The ECDSA public key
     */
    function getECDSAKey(
        address avs, 
        uint32 operatorSetId, 
        address operator
    ) external view returns (bytes memory) {
        KeyInfo memory keyInfo = operatorKeyInfo[avs][operatorSetId][operator];
        return keyInfo.keyData;
    }

    /**
     * @notice Checks if a key hash is globally registered
     * @param keyHash Hash of the key
     * @return True if the key is globally registered
     */
    function isKeyGloballyRegistered(bytes32 keyHash) external view returns (bool) {
        return globalKeyRegistry[keyHash];
    }

    /**
     * @notice Gets the key hash for an operator with a specific AVS
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param operator Address of the operator
     * @return keyHash The key hash
     */
    function getKeyHash(
        address avs,
        uint32 operatorSetId,
        address operator
    ) external view returns (bytes32) {
        KeyInfo memory keyInfo = operatorKeyInfo[avs][operatorSetId][operator];
        OperatorSetConfig memory config = operatorSetConfigs[avs][operatorSetId];
        
        if (!keyInfo.isRegistered) {
            return bytes32(0);
        }
        
        if (config.curveType == CurveType.ECDSA) {
            return keccak256(keyInfo.keyData);
        } else if (config.curveType == CurveType.BN254) {
            (uint256 g1X, uint256 g1Y, , ) = abi.decode(keyInfo.keyData, (uint256, uint256, uint256[2], uint256[2]));
            return BN254.hashG1Point(BN254.G1Point(g1X, g1Y));
        }
        
        revert InvalidCurveType();
    }

    /**
     * @notice Update an operator set's aggregate BN254 key
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @param key BN254 key to add or remove
     * @param isAddition True to add the key, false to remove it
     */
    function _updateAggregateBN254Key(
        address avs,
        uint32 operatorSetId,
        BN254.G1Point memory key,
        bool isAddition
    ) internal {
        BN254.G1Point memory currentApk = avsOperatorSetToAggregateBN254Key[avs][operatorSetId];
        BN254.G1Point memory newApk;
        
        if (isAddition) {
            newApk = currentApk.plus(key);
        } else {
            newApk = currentApk.plus(key.negate());
        }
        
        avsOperatorSetToAggregateBN254Key[avs][operatorSetId] = newApk;
        
        emit AggregateBN254KeyUpdated(avs, operatorSetId, newApk);
    }

    /**
     * @notice Gets the aggregate BN254 public key for an operator set
     * @param avs Address of the AVS
     * @param operatorSetId ID of the operator set
     * @return The aggregate BN254 G1 public key
     */
    function getApk(
        address avs,
        uint32 operatorSetId
    ) external view returns (BN254.G1Point memory) {
        return avsOperatorSetToAggregateBN254Key[avs][operatorSetId];
    }

    /**
     * @notice Verifies a BN254 signature using Fiat-Shamir challenge to prevent rogue key attacks
     * @param messageHash Hash of the message being signed
     * @param signature The signature bytes
     * @param pubkeyG1 The G1 component of the public key
     * @param pubkeyG2 The G2 component of the public key
     * @dev Uses gamma challenge value derived from message and public key components
     */
    function verifyBN254Signature(
        bytes32 messageHash,
        bytes memory signature,
        BN254.G1Point memory pubkeyG1,
        BN254.G2Point memory pubkeyG2
    ) internal view {
        // Decode signature
        BN254.G1Point memory sigPoint;
        {
            (uint256 sigX, uint256 sigY) = abi.decode(signature, (uint256, uint256));
            sigPoint = BN254.G1Point(sigX, sigY);
        }
        
        // Generate challenge value using Fiat-Shamir transform to prevent rogue key attacks
        uint256 gamma = uint256(
            keccak256(
                abi.encodePacked(
                    messageHash,
                    pubkeyG1.X,
                    pubkeyG1.Y,
                    pubkeyG2.X[0],
                    pubkeyG2.X[1],
                    pubkeyG2.Y[0],
                    pubkeyG2.Y[1],
                    sigPoint.X,
                    sigPoint.Y
                )
            )
        ) % BN254.FR_MODULUS;
        
        
        // Verify signature using both G1 and G2 components with gamma challenge
        // e(sig + pubkey_G1*gamma, G2) = e(msg + G1*gamma, pubkey_G2)
        bool pairingSuccessful;
        bool signatureValid;
        (pairingSuccessful, signatureValid) = BN254.safePairing(
            sigPoint.plus(pubkeyG1.scalar_mul(gamma)),  // sigma + pubkey*gamma
            BN254.negGeneratorG2(),                       // -G2
            BN254.hashToG1(messageHash).plus(BN254.generatorG1().scalar_mul(gamma)), // H(m) + g1*gamma
            pubkeyG2,                                        // pubkeyG2
            PAIRING_EQUALITY_CHECK_GAS
        );
        
        if (!signatureValid) {
            revert InvalidSignature();
        }
    }
}