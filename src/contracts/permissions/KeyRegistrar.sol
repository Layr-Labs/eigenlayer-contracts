// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "../libraries/BN254.sol";
import "../libraries/BN254SignatureVerifier.sol";
import "../mixins/PermissionControllerMixin.sol";
import "../mixins/SignatureUtilsMixin.sol";
import "../interfaces/IPermissionController.sol";
import "../interfaces/IAllocationManager.sol";
import "../interfaces/IKeyRegistrar.sol";
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
        CurveType _curveType = _operatorSetCurveTypes[operatorSet.key()];
        require(_curveType == CurveType.NONE, ConfigurationAlreadySet());

        _operatorSetCurveTypes[operatorSet.key()] = curveType;

        emit OperatorSetConfigured(operatorSet, curveType);
    }

    /// @inheritdoc IKeyRegistrar
    function registerKey(
        address operator,
        OperatorSet memory operatorSet,
        bytes calldata keyData,
        bytes calldata signature
    ) external checkCanCall(operator) {
        CurveType curveType = _operatorSetCurveTypes[operatorSet.key()];
        require(curveType != CurveType.NONE, OperatorSetNotConfigured());

        // Check if the operator is already registered to the operatorSet
        require(!_operatorKeyInfo[operatorSet.key()][operator].isRegistered, KeyAlreadyRegistered());

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
    function deregisterKey(address operator, OperatorSet memory operatorSet) external checkCanCall(operator) {
        // Operators can only deregister if they are not slashable for this operator set
        require(
            !allocationManager.isOperatorSlashable(operator, operatorSet), OperatorStillSlashable(operatorSet, operator)
        );

        CurveType curveType = _operatorSetCurveTypes[operatorSet.key()];
        require(curveType != CurveType.NONE, OperatorSetNotConfigured());

        KeyInfo memory keyInfo = _operatorKeyInfo[operatorSet.key()][operator];

        require(keyInfo.isRegistered, KeyNotFound(operatorSet, operator));

        // Clear key info
        delete _operatorKeyInfo[operatorSet.key()][operator];

        emit KeyDeregistered(operatorSet, operator, curveType);
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
        require(!_globalKeyRegistry[keyHash], KeyAlreadyRegistered());

        // Get the signable digest for the ECDSA key registration message
        bytes32 signableDigest = getECDSAKeyRegistrationMessageHash(operator, operatorSet, keyAddress);

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
        require(keyData.length == 192, InvalidKeyFormat());
        require(signature.length == 64, InvalidSignature());

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
        bytes32 signableDigest = getBN254KeyRegistrationMessageHash(operator, operatorSet, keyData);

        // Decode signature from bytes to G1 point
        (uint256 sigX, uint256 sigY) = abi.decode(signature, (uint256, uint256));
        BN254.G1Point memory signaturePoint = BN254.G1Point(sigX, sigY);

        // Verify signature
        (, bool pairingSuccessful) =
            BN254SignatureVerifier.verifySignature(signableDigest, signaturePoint, g1Point, g2Point, false, 0);
        require(pairingSuccessful, InvalidSignature());

        // Calculate key hash and check global uniqueness
        bytes32 keyHash = _getKeyHashForKeyData(keyData, CurveType.BN254);
        require(!_globalKeyRegistry[keyHash], KeyAlreadyRegistered());

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
        _operatorKeyInfo[operatorSet.key()][operator] = KeyInfo({isRegistered: true, keyData: pubkey});

        // Mark the key hash as spent
        _globalKeyRegistry[keyHash] = true;

        // Store the operator for the key hash
        _keyHashToOperator[keyHash] = operator;
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
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IKeyRegistrar
    function isRegistered(OperatorSet memory operatorSet, address operator) public view returns (bool) {
        return _operatorKeyInfo[operatorSet.key()][operator].isRegistered;
    }

    /// @inheritdoc IKeyRegistrar
    function getOperatorSetCurveType(
        OperatorSet memory operatorSet
    ) external view returns (CurveType) {
        return _operatorSetCurveTypes[operatorSet.key()];
    }

    /// @inheritdoc IKeyRegistrar
    function getBN254Key(
        OperatorSet memory operatorSet,
        address operator
    ) external view returns (BN254.G1Point memory g1Point, BN254.G2Point memory g2Point) {
        // Validate operator set curve type
        CurveType curveType = _operatorSetCurveTypes[operatorSet.key()];
        require(curveType == CurveType.BN254, InvalidCurveType());

        KeyInfo memory keyInfo = _operatorKeyInfo[operatorSet.key()][operator];

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
        CurveType curveType = _operatorSetCurveTypes[operatorSet.key()];
        require(curveType == CurveType.ECDSA, InvalidCurveType());

        KeyInfo memory keyInfo = _operatorKeyInfo[operatorSet.key()][operator];
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
        return _globalKeyRegistry[keyHash];
    }

    /// @inheritdoc IKeyRegistrar
    function getKeyHash(OperatorSet memory operatorSet, address operator) external view returns (bytes32) {
        KeyInfo memory keyInfo = _operatorKeyInfo[operatorSet.key()][operator];
        CurveType curveType = _operatorSetCurveTypes[operatorSet.key()];

        if (!keyInfo.isRegistered) {
            return bytes32(0);
        }

        return _getKeyHashForKeyData(keyInfo.keyData, curveType);
    }

    /// @inheritdoc IKeyRegistrar
    function getOperatorFromSigningKey(
        OperatorSet memory operatorSet,
        bytes memory keyData
    ) external view returns (address, bool) {
        CurveType curveType = _operatorSetCurveTypes[operatorSet.key()];

        // We opt to not use _getKeyHashForKeyData here because it expects the G1 and G2 key encoded together for BN254
        bytes32 keyHash;
        if (curveType == CurveType.ECDSA) {
            keyHash = keccak256(keyData);
        } else if (curveType == CurveType.BN254) {
            /// We cannot use _getKeyHashForKeyData here because it expects the G1 and G2 key encoded together
            (uint256 g1X, uint256 g1Y) = abi.decode(keyData, (uint256, uint256));
            keyHash = BN254.hashG1Point(BN254.G1Point(g1X, g1Y));
        } else {
            revert InvalidCurveType();
        }

        address operator = _keyHashToOperator[keyHash];
        return (operator, isRegistered(operatorSet, operator));
    }

    /// @inheritdoc IKeyRegistrar
    function getECDSAKeyRegistrationMessageHash(
        address operator,
        OperatorSet memory operatorSet,
        address keyAddress
    ) public view returns (bytes32) {
        bytes32 structHash = keccak256(
            abi.encode(ECDSA_KEY_REGISTRATION_TYPEHASH, operator, operatorSet.avs, operatorSet.id, keyAddress)
        );
        return _calculateSignableDigest(structHash);
    }

    /// @inheritdoc IKeyRegistrar
    function getBN254KeyRegistrationMessageHash(
        address operator,
        OperatorSet memory operatorSet,
        bytes calldata keyData
    ) public view returns (bytes32) {
        bytes32 structHash = keccak256(
            abi.encode(BN254_KEY_REGISTRATION_TYPEHASH, operator, operatorSet.avs, operatorSet.id, keccak256(keyData))
        );
        return _calculateSignableDigest(structHash);
    }

    /// @inheritdoc IKeyRegistrar
    function encodeBN254KeyData(
        BN254.G1Point memory g1Point,
        BN254.G2Point memory g2Point
    ) public pure returns (bytes memory) {
        return abi.encode(g1Point.X, g1Point.Y, g2Point.X[0], g2Point.X[1], g2Point.Y[0], g2Point.Y[1]);
    }
}
