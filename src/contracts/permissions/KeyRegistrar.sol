// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "../libraries/BN254.sol";
import "../mixins/PermissionControllerMixin.sol";
import "../mixins/SignatureUtilsMixin.sol";
import "../interfaces/IPermissionController.sol";
import "../interfaces/IAllocationManager.sol";
import "../libraries/OperatorSetLib.sol";
import "./KeyRegistrarStorage.sol";

/**
 * @title KeyRegistrar
 * @notice A core singleton contract that manages operator keys for different AVSs with global key uniqueness
 * @dev Provides registration and deregistration of keys with support for aggregate keys
 *      Keys must be unique globally across all AVSs and operator sets
 *      Operators call functions directly to manage their own keys
 *      Aggregate keys are updated via callback from AVSRegistrar on registration and deregistration
 */
contract KeyRegistrar is KeyRegistrarStorage, PermissionControllerMixin, SignatureUtilsMixin {
    using BN254 for BN254.G1Point;

    // EIP-712 type hashes
    bytes32 public constant ECDSA_KEY_REGISTRATION_TYPEHASH =
        keccak256("ECDSAKeyRegistration(address operator,address avs,uint32 operatorSetId,address keyAddress)");

    bytes32 public constant BN254_KEY_REGISTRATION_TYPEHASH =
        keccak256("BN254KeyRegistration(address operator,address avs,uint32 operatorSetId,bytes keyData)");

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
        SignatureUtilsMixin(_version)
    {}

    /// @inheritdoc IKeyRegistrar
    function configureOperatorSet(
        OperatorSet memory operatorSet,
        CurveType curveType
    ) external checkCanCall(operatorSet.avs) {
        require(curveType == CurveType.ECDSA || curveType == CurveType.BN254, InvalidCurveType());

        // Prevent overwriting existing configurations
        CurveType _curveType = operatorSetCurveTypes[operatorSet.key()];
        require(_curveType == CurveType.NONE, ConfigurationAlreadySet());

        operatorSetCurveTypes[operatorSet.key()] = curveType;

        emit OperatorSetConfigured(operatorSet, curveType);
    }

    /// @inheritdoc IKeyRegistrar
    function registerKey(
        address operator,
        OperatorSet memory operatorSet,
        bytes calldata keyData,
        bytes calldata signature
    ) external checkCanCall(operator) {
        CurveType curveType = operatorSetCurveTypes[operatorSet.key()];
        require(curveType != CurveType.NONE, OperatorSetNotConfigured());

        // Check if the key is already registered
        require(!operatorKeyInfo[operatorSet.key()][operator].isRegistered, KeyAlreadyRegistered());

        // Register key based on curve type - both now require signature verification
        if (curveType == CurveType.ECDSA) {
            _registerECDSAKey(operatorSet, operator, keyData, signature);
        } else if (curveType == CurveType.BN254) {
            _registerBN254Key(operatorSet, operator, keyData, signature);
        } else {
            revert InvalidCurveType();
        }

        emit KeyRegistered(operatorSet, operator, curveType, keyData);
    }

    /// @inheritdoc IKeyRegistrar
    function deregisterKey(address operator, OperatorSet memory operatorSet) external {
        require(address(allocationManager.getAVSRegistrar(operatorSet.avs)) == msg.sender, InvalidPermissions());

        CurveType curveType = operatorSetCurveTypes[operatorSet.key()];
        require(curveType != CurveType.NONE, OperatorSetNotConfigured());

        KeyInfo memory keyInfo = operatorKeyInfo[operatorSet.key()][operator];

        require(keyInfo.isRegistered, KeyNotFound(operatorSet, operator));

        // Clear key info
        delete operatorKeyInfo[operatorSet.key()][operator];

        emit KeyDeregistered(operatorSet, operator, curveType);
    }

    /// @inheritdoc IKeyRegistrar
    function checkKey(OperatorSet memory operatorSet, address operator) external view returns (bool) {
        CurveType curveType = operatorSetCurveTypes[operatorSet.key()];
        require(curveType != CurveType.NONE, OperatorSetNotConfigured());

        KeyInfo memory keyInfo = operatorKeyInfo[operatorSet.key()][operator];
        return keyInfo.isRegistered;
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Validates and registers an ECDSA address with EIP-712 signature verification
     * @param operatorSet The operator set to register the key for
     * @param operator Address of the operator
     * @param keyData The ECDSA address encoded as bytes (20 bytes)
     * @param signature EIP-712 signature over the registration message
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

        // Create EIP-712 compliant message hash
        bytes32 structHash = keccak256(
            abi.encode(ECDSA_KEY_REGISTRATION_TYPEHASH, operator, operatorSet.avs, operatorSet.id, keyAddress)
        );
        bytes32 signableDigest = _calculateSignableDigest(structHash);

        _checkIsValidSignatureNow(keyAddress, signableDigest, signature, type(uint256).max);

        // Store key data
        _storeKeyData(operatorSet, operator, keyData, keyHash);
    }

    /**
     * @notice Validates and registers a BN254 public key with proper signature verification
     * @param operatorSet The operator set to register the key for
     * @param operator Address of the operator
     * @param keyData The BN254 public key bytes (G1 and G2 components)
     * @param signature Signature proving key ownership
     * @dev Validates keypair, verifies signature using hash-to-G1, and ensures global uniqueness
     */
    function _registerBN254Key(
        OperatorSet memory operatorSet,
        address operator,
        bytes calldata keyData,
        bytes calldata signature
    ) internal {
        BN254.G1Point memory g1Point;
        BN254.G2Point memory g2Point;

        {
            // Decode BN254 G1 and G2 points from the keyData bytes
            (uint256 g1X, uint256 g1Y, uint256[2] memory g2X, uint256[2] memory g2Y) =
                abi.decode(keyData, (uint256, uint256, uint256[2], uint256[2]));

            // Validate G1 point
            g1Point = BN254.G1Point(g1X, g1Y);
            require(!(g1X == 0 && g1Y == 0), ZeroPubkey());

            // Construct BN254 G2 point from coordinates
            g2Point = BN254.G2Point(g2X, g2Y);
        }

        // Create EIP-712 compliant message hash
        bytes32 structHash = keccak256(
            abi.encode(BN254_KEY_REGISTRATION_TYPEHASH, operator, operatorSet.avs, operatorSet.id, keccak256(keyData))
        );
        bytes32 signableDigest = _calculateSignableDigest(structHash);

        // hash to g1
        BN254.G1Point memory messagePoint = BN254.hashToG1(signableDigest);
        _verifyBN254Signature(messagePoint, signature, g1Point, g2Point);

        // Calculate key hash and check global uniqueness
        bytes32 keyHash = _getKeyHashForKeyData(keyData, CurveType.BN254);
        require(!globalKeyRegistry[keyHash], KeyAlreadyRegistered());

        // Store key data
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
        operatorKeyInfo[operatorSet.key()][operator] = KeyInfo({isRegistered: true, keyData: pubkey});

        // Update global key registry
        globalKeyRegistry[keyHash] = true;
    }

    /**
     * @notice Internal helper to get key hash for pubkey data using consistent hashing
     * @param pubkey The public key data
     * @param curveType The curve type (ECDSA or BN254)
     * @return keyHash The key hash
     */
    function _getKeyHashForKeyData(bytes memory pubkey, CurveType curveType) internal pure returns (bytes32) {
        if (curveType == CurveType.ECDSA) {
            return keccak256(pubkey);
        } else if (curveType == CurveType.BN254) {
            (uint256 g1X, uint256 g1Y,,) = abi.decode(pubkey, (uint256, uint256, uint256[2], uint256[2]));
            return BN254.hashG1Point(BN254.G1Point(g1X, g1Y));
        }

        revert InvalidCurveType();
    }

    /**
     * @notice Verifies a BN254 signature
     * @param messagePoint The G1 point representing the hashed message
     * @param signature The signature bytes
     * @param pubkeyG1 The G1 component of the public key
     * @param pubkeyG2 The G2 component of the public key
     */
    function _verifyBN254Signature(
        BN254.G1Point memory messagePoint,
        bytes memory signature,
        BN254.G1Point memory pubkeyG1,
        BN254.G2Point memory pubkeyG2
    ) public view {
        // Decode signature
        BN254.G1Point memory sigPoint;
        {
            (uint256 sigX, uint256 sigY) = abi.decode(signature, (uint256, uint256));
            sigPoint = BN254.G1Point(sigX, sigY);
        }

        // gamma = h(sigma, P, P', H(m)) - exact same pattern as BLSApkRegistry
        uint256 gamma = uint256(
            keccak256(
                abi.encodePacked(
                    sigPoint.X,
                    sigPoint.Y,
                    pubkeyG1.X,
                    pubkeyG1.Y,
                    pubkeyG2.X[0],
                    pubkeyG2.X[1],
                    pubkeyG2.Y[0],
                    pubkeyG2.Y[1],
                    messagePoint.X,
                    messagePoint.Y
                )
            )
        ) % BN254.FR_MODULUS;

        // e(sigma + P * gamma, [1]_2) = e(H(m) + [1]_1 * gamma, P')
        require(
            BN254.pairing(
                sigPoint.plus(pubkeyG1.scalar_mul(gamma)),
                BN254.negGeneratorG2(),
                messagePoint.plus(BN254.generatorG1().scalar_mul(gamma)),
                pubkeyG2
            ),
            InvalidSignature()
        );
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IKeyRegistrar
    function isRegistered(OperatorSet memory operatorSet, address operator) external view returns (bool) {
        return operatorKeyInfo[operatorSet.key()][operator].isRegistered;
    }

    /// @inheritdoc IKeyRegistrar
    function getOperatorSetCurveType(
        OperatorSet memory operatorSet
    ) external view returns (CurveType) {
        return operatorSetCurveTypes[operatorSet.key()];
    }

    /// @inheritdoc IKeyRegistrar
    function getBN254Key(
        OperatorSet memory operatorSet,
        address operator
    ) external view returns (BN254.G1Point memory g1Point, BN254.G2Point memory g2Point) {
        // Validate operator set curve type
        CurveType curveType = operatorSetCurveTypes[operatorSet.key()];
        require(curveType == CurveType.BN254, InvalidCurveType());

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
    function getECDSAKey(OperatorSet memory operatorSet, address operator) public view returns (bytes memory) {
        // Validate operator set curve type
        CurveType curveType = operatorSetCurveTypes[operatorSet.key()];
        require(curveType == CurveType.ECDSA, InvalidCurveType());

        KeyInfo memory keyInfo = operatorKeyInfo[operatorSet.key()][operator];
        return keyInfo.keyData; // Returns the 20-byte address as bytes
    }

    /// @inheritdoc IKeyRegistrar
    function getECDSAAddress(OperatorSet memory operatorSet, address operator) external view returns (address) {
        return address(bytes20(getECDSAKey(operatorSet, operator)));
    }

    /// @inheritdoc IKeyRegistrar
    function isKeyGloballyRegistered(
        bytes32 keyHash
    ) external view returns (bool) {
        return globalKeyRegistry[keyHash];
    }

    /// @inheritdoc IKeyRegistrar
    function getKeyHash(OperatorSet memory operatorSet, address operator) external view returns (bytes32) {
        KeyInfo memory keyInfo = operatorKeyInfo[operatorSet.key()][operator];
        CurveType curveType = operatorSetCurveTypes[operatorSet.key()];

        if (!keyInfo.isRegistered) {
            return bytes32(0);
        }

        return _getKeyHashForKeyData(keyInfo.keyData, curveType);
    }

    /**
     * @notice Returns the message hash for ECDSA key registration
     * @param operator The operator address
     * @param operatorSet The operator set
     * @param keyAddress The ECDSA key address
     * @return The message hash for signing
     */
    function getECDSAKeyRegistrationMessageHash(
        address operator,
        OperatorSet memory operatorSet,
        address keyAddress
    ) external view returns (bytes32) {
        bytes32 structHash = keccak256(
            abi.encode(ECDSA_KEY_REGISTRATION_TYPEHASH, operator, operatorSet.avs, operatorSet.id, keyAddress)
        );
        return _calculateSignableDigest(structHash);
    }

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
    ) external view returns (bytes32) {
        bytes32 structHash = keccak256(
            abi.encode(BN254_KEY_REGISTRATION_TYPEHASH, operator, operatorSet.avs, operatorSet.id, keccak256(keyData))
        );
        return _calculateSignableDigest(structHash);
    }
}
