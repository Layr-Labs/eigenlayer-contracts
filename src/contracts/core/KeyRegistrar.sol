// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {BN254} from "../libraries/BN254.sol";
import {PermissionControllerMixin} from "../mixins/PermissionControllerMixin.sol";
import {SemVerMixin} from "../mixins/SemVerMixin.sol";
import {IPermissionController} from "../interfaces/IPermissionController.sol";
import {IAllocationManager} from "../interfaces/IAllocationManager.sol";
import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import "./KeyRegistrarStorage.sol";

/**
 * @title KeyRegistrar
 * @notice A core singleton contract that manages operator keys for different AVSs with global key uniqueness
 * @dev Provides registration and deregistration of keys with support for aggregate keys
 *      Keys must be unique globally across all AVSs and operator sets
 *      Operators call functions directly to manage their own keys
 *      Aggregate keys are updated via callback from AVSRegistrar on registration and deregistration
 */
contract KeyRegistrar is 
    Initializable,
    OwnableUpgradeable,
    KeyRegistrarStorage,
    PermissionControllerMixin,
    SemVerMixin
{
    using BN254 for BN254.G1Point;

    /**
     * @dev Constructor for the KeyRegistrar contract
     * @param _permissionController The permission controller contract
     * @param _allocationManager The allocation manager contract
     * @param _version The version string for the contract
     */
    constructor(
        IPermissionController _permissionController,
        IAllocationManager _allocationManager,
        string memory _version
    ) 
        KeyRegistrarStorage(_allocationManager)
        PermissionControllerMixin(_permissionController) 
        SemVerMixin(_version)
    {
        _disableInitializers();
    }

    /// @inheritdoc IKeyRegistrar
    function initialize(address initialOwner) external initializer {
        require(initialOwner != address(0), ZeroAddress());
        _transferOwnership(initialOwner);
    }

    /// @inheritdoc IKeyRegistrar
    function configureOperatorSet(
        OperatorSet memory operatorSet,
        CurveType curveType
    ) external {
        require(address(allocationManager.getAVSRegistrar(operatorSet.avs)) == msg.sender, "Unauthorized");
        require(curveType == CurveType.ECDSA || curveType == CurveType.BN254, InvalidCurveType());
        
        // Prevent overwriting existing configurations
        OperatorSetConfig storage config = operatorSetConfigs[operatorSet.key()];
        require(!config.isActive, ConfigurationAlreadySet());
        
        operatorSetConfigs[operatorSet.key()] = OperatorSetConfig({
            curveType: curveType,
            isActive: true
        });

        emit OperatorSetConfigured(operatorSet, curveType);
    }

    /// @inheritdoc IKeyRegistrar
    function registerKey(
        address operator,
        OperatorSet memory operatorSet,
        bytes calldata keyData,
        bytes calldata signature
    ) external checkCanCall(operator) {
        
        OperatorSetConfig memory config = operatorSetConfigs[operatorSet.key()];
        require(config.isActive, OperatorSetNotConfigured());

        // Check if the key is already registered
        require(!operatorKeyInfo[operatorSet.key()][operator].isRegistered, KeyAlreadyRegistered());

        // Register key based on curve type - both now require signature verification
        if (config.curveType == CurveType.ECDSA) {
            _registerECDSAKey(operatorSet, operator, keyData, signature);
        } 
        else if (config.curveType == CurveType.BN254) {
            _registerBN254Key(operatorSet, operator, keyData, signature);
        } 
        else {
            revert InvalidCurveType();
        }

        emit KeyRegistered(operatorSet, operator, config.curveType, keyData);
    }

    /// @inheritdoc IKeyRegistrar
    function deregisterKey(
        address operator,
        OperatorSet memory operatorSet
    ) external {
        require(address(allocationManager.getAVSRegistrar(operatorSet.avs)) == msg.sender, "Unauthorized");
        
        OperatorSetConfig memory config = operatorSetConfigs[operatorSet.key()];
        require(config.isActive, OperatorSetNotConfigured());
        
        KeyInfo memory keyInfo = operatorKeyInfo[operatorSet.key()][operator];
        
        require(keyInfo.isRegistered, KeyNotFound(operatorSet, operator));

        // Clear key info
        delete operatorKeyInfo[operatorSet.key()][operator];

        emit KeyDeregistered(operatorSet, operator, config.curveType);
    }

    /// @inheritdoc IKeyRegistrar
    function checkKey(
        OperatorSet memory operatorSet,
        address operator
    ) external view returns (bool) {
        require(address(allocationManager.getAVSRegistrar(operatorSet.avs)) == msg.sender, "Unauthorized");
        
        OperatorSetConfig memory config = operatorSetConfigs[operatorSet.key()];
        require(config.isActive, OperatorSetNotConfigured());
        
        KeyInfo memory keyInfo = operatorKeyInfo[operatorSet.key()][operator];
        return keyInfo.isRegistered;
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Validates and registers an ECDSA address with signature verification
     * @param operatorSet The operator set to register the key for
     * @param operator Address of the operator
     * @param keyData The ECDSA address encoded as bytes (20 bytes)
     * @param signature Signature over a message that includes the operator's address
     * @dev Validates address format, verifies signature ownership, and ensures global uniqueness
     */
    function _registerECDSAKey(
        OperatorSet memory operatorSet,
        address operator,
        bytes calldata keyData,
        bytes calldata signature
    ) internal {
        // Validate ECDSA address format
        require(keyData.length == 20, InvalidKeyFormat());
        
        // Decode address from bytes
        address keyAddress = address(bytes20(keyData));
        require(keyAddress != address(0), ZeroPubkey());
        
        // Calculate key hash using the address
        bytes32 keyHash = _getKeyHashForKeyData(keyData, CurveType.ECDSA);
        
        // Check global uniqueness
        require(!globalKeyRegistry[keyHash], KeyAlreadyRegistered());

        // Verify signature to prevent frontrunning attacks
        bytes32 messageHash = keccak256(abi.encodePacked(
            "EigenLayer.KeyRegistrar.ECDSA.v1", // Domain separator
            address(this), // Contract address for additional separation
            operatorSet.avs, 
            operatorSet.id, 
            operator, 
            keyAddress
        ));

        // Verify that the signature was created by the key address
        address recoveredSigner = ECDSA.recover(messageHash, signature);
        require(recoveredSigner == keyAddress, InvalidSignature());

        // Store key data using internal helper
        _storeKeyData(operatorSet, operator, keyData, keyHash);
    }

    /**
     * @notice Validates and registers a BN254 public key
     * @param operatorSet The operator set to register the key for
     * @param operator Address of the operator
     * @param keyData The BN254 public key bytes (G1 and G2 components)
     * @param signature Signature proving key ownership
     * @dev Validates keypair, verifies signature, and ensures global uniqueness
     */
    function _registerBN254Key(
        OperatorSet memory operatorSet,
        address operator,
        bytes calldata keyData,
        bytes calldata signature
    ) internal {
        BN254.G1Point memory g1Point;

        {
        // Decode BN254 G1 and G2 points from the keyData bytes
        (uint256 g1X, uint256 g1Y, uint256[2] memory g2X, uint256[2] memory g2Y) = 
            abi.decode(keyData, (uint256, uint256, uint256[2], uint256[2]));

        // Validate G1 point
        g1Point = BN254.G1Point(g1X, g1Y);
        require(!(g1X == 0 && g1Y == 0), ZeroPubkey());

        // Construct BN254 G2 point from coordinates
        BN254.G2Point memory g2Point = BN254.G2Point(g2X, g2Y);

        // Verify the signature to prevent rogue key attacks with domain separation
        bytes32 messageHash = keccak256(abi.encodePacked(
            "EigenLayer.KeyRegistrar.v1", // Domain separator
            address(this), // Contract address for additional separation
            operatorSet.avs, 
            operatorSet.id, 
            operator, 
            keyData
        ));
        _verifyBN254Signature(messageHash, signature, g1Point, g2Point);
        }

        // Calculate key hash and check global uniqueness
        bytes32 keyHash = _getKeyHashForKeyData(keyData, CurveType.BN254);
        require(!globalKeyRegistry[keyHash], KeyAlreadyRegistered());

        // Store key data using internal helper
        _storeKeyData(operatorSet, operator, keyData, keyHash);
    }

    /**
     * @notice Internal helper to store key data and update global registry
     * @param operatorSet The operator set
     * @param operator The operator address
     * @param pubkey The public key data
     * @param keyHash The key hash
     */
    function _storeKeyData(
        OperatorSet memory operatorSet,
        address operator,
        bytes memory pubkey,
        bytes32 keyHash
    ) internal {
        // Store key data
        operatorKeyInfo[operatorSet.key()][operator] = KeyInfo({
            isRegistered: true,
            keyData: pubkey
        });

        // Update global key registry
        globalKeyRegistry[keyHash] = true;
    }

    /**
     * @notice Internal helper to get key hash for pubkey data using consistent hashing
     * @param pubkey The public key data
     * @param curveType The curve type (ECDSA or BN254)
     * @return keyHash The key hash
     */
    function _getKeyHashForKeyData(
        bytes memory pubkey,
        CurveType curveType
    ) internal pure returns (bytes32) {
        if (curveType == CurveType.ECDSA) {
            return keccak256(pubkey);
        } else if (curveType == CurveType.BN254) {
            (uint256 g1X, uint256 g1Y, , ) = abi.decode(pubkey, (uint256, uint256, uint256[2], uint256[2]));
            return BN254.hashG1Point(BN254.G1Point(g1X, g1Y));
        }
        
        revert InvalidCurveType();
    }

    /**
     * @notice Verifies a BN254 signature using Fiat-Shamir challenge to prevent rogue key attacks
     * @param messageHash Hash of the message being signed
     * @param signature The signature bytes
     * @param pubkeyG1 The G1 component of the public key
     * @param pubkeyG2 The G2 component of the public key
     * @dev Uses gamma challenge value derived from message and public key components
     */
    function _verifyBN254Signature(
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
        
        require(signatureValid, InvalidSignature());
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IKeyRegistrar
    function isRegistered(
        OperatorSet memory operatorSet,
        address operator
    ) external view returns (bool) {
        return operatorKeyInfo[operatorSet.key()][operator].isRegistered;
    }

    /// @inheritdoc IKeyRegistrar
    function getOperatorSetConfig(
        OperatorSet memory operatorSet
    ) external view returns (OperatorSetConfig memory) {
        return operatorSetConfigs[operatorSet.key()];
    }

    /// @inheritdoc IKeyRegistrar
    function getBN254Key(
        OperatorSet memory operatorSet, 
        address operator
    ) external view returns (BN254.G1Point memory g1Point, BN254.G2Point memory g2Point) {
        // Validate operator set curve type
        OperatorSetConfig memory config = operatorSetConfigs[operatorSet.key()];
        require(config.curveType == CurveType.BN254, InvalidCurveType());
        
        KeyInfo memory keyInfo = operatorKeyInfo[operatorSet.key()][operator];
        
        if (!keyInfo.isRegistered) {
            // Create default values for an empty key
            uint256[2] memory zeroArray = [uint256(0), uint256(0)];
            return (BN254.G1Point(0, 0), BN254.G2Point(zeroArray, zeroArray));
        }
        
        (uint256 g1X, uint256 g1Y, uint256[2] memory g2X, uint256[2] memory g2Y) = 
            abi.decode(keyInfo.keyData, (uint256, uint256, uint256[2], uint256[2]));
        
        return (BN254.G1Point(g1X, g1Y), BN254.G2Point(g2X, g2Y));
    }

    /// @inheritdoc IKeyRegistrar
    function getECDSAKey(
        OperatorSet memory operatorSet, 
        address operator
    ) public view returns (bytes memory) {
        // Validate operator set curve type
        OperatorSetConfig memory config = operatorSetConfigs[operatorSet.key()];
        require(config.curveType == CurveType.ECDSA, InvalidCurveType());
        
        KeyInfo memory keyInfo = operatorKeyInfo[operatorSet.key()][operator];
        return keyInfo.keyData; // Returns the 20-byte address as bytes
    }

    /// @inheritdoc IKeyRegistrar
    function getECDSAAddress(
        OperatorSet memory operatorSet, 
        address operator
    ) external view returns (address) {    
        return address(bytes20(getECDSAKey(operatorSet, operator)));
    }

    /// @inheritdoc IKeyRegistrar
    function isKeyGloballyRegistered(bytes32 keyHash) external view returns (bool) {
        return globalKeyRegistry[keyHash];
    }

    /// @inheritdoc IKeyRegistrar
    function getKeyHash(
        OperatorSet memory operatorSet,
        address operator
    ) external view returns (bytes32) {
        KeyInfo memory keyInfo = operatorKeyInfo[operatorSet.key()][operator];
        OperatorSetConfig memory config = operatorSetConfigs[operatorSet.key()];
        
        if (!keyInfo.isRegistered) {
            return bytes32(0);
        }
        
        return _getKeyHashForKeyData(keyInfo.keyData, config.curveType);
    }

    /// @inheritdoc IKeyRegistrar
    function verifyBN254Signature(
        bytes32 messageHash,
        bytes memory signature,
        BN254.G1Point memory pubkeyG1,
        BN254.G2Point memory pubkeyG2
    ) external view {
        _verifyBN254Signature(messageHash, signature, pubkeyG1, pubkeyG2);
    }
}