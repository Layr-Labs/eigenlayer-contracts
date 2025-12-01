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

/// @title KeyRegistrar
/// @notice A core singleton contract that manages operator keys for different AVSs with global key uniqueness
/// @dev Provides registration and deregistration of keys with support for aggregate keys
///      Keys must be unique globally across all AVSs and operator sets
///      Operators call functions directly to manage their own keys
///      Aggregate keys are updated via callback from AVSRegistrar on registration and deregistration
contract KeyRegistrar is KeyRegistrarStorage, PermissionControllerMixin, SignatureUtilsMixin {
    using BN254 for BN254.G1Point;

    // EIP-712 type hashes
    bytes32 public constant ECDSA_KEY_REGISTRATION_TYPEHASH =
        keccak256("ECDSAKeyRegistration(address operator,address avs,uint32 operatorSetId,address keyAddress)");

    bytes32 public constant BN254_KEY_REGISTRATION_TYPEHASH =
        keccak256("BN254KeyRegistration(address operator,address avs,uint32 operatorSetId,bytes keyData)");

    /// @dev Constructor for the KeyRegistrar contract
    /// @param _permissionController The permission controller contract
    /// @param _allocationManager The allocation manager contract
    /// @param _version The version string for the contract
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
        _setMinRotationDelay(operatorSet, type(uint64).max);
    }

    /// @inheritdoc IKeyRegistrar
    function configureOperatorSetWithMinDelay(
        OperatorSet memory operatorSet,
        CurveType curveType,
        uint64 minDelaySeconds
    ) external checkCanCall(operatorSet.avs) {
        require(curveType == CurveType.ECDSA || curveType == CurveType.BN254, InvalidCurveType());

        // Prevent overwriting existing configurations
        CurveType _curveType = _operatorSetCurveTypes[operatorSet.key()];
        require(_curveType == CurveType.NONE, ConfigurationAlreadySet());

        _operatorSetCurveTypes[operatorSet.key()] = curveType;
        emit OperatorSetConfigured(operatorSet, curveType);
        _setMinRotationDelay(operatorSet, minDelaySeconds);
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
        require(!_operatorKeyInfo[operatorSet.key()][operator].isRegistered, OperatorAlreadyRegistered());

        // Validate and reserve the key globally, then store
        _validateAndReserveKey(operatorSet, operator, keyData, signature, curveType);
        _operatorKeyInfo[operatorSet.key()][operator] =
            KeyInfo({isRegistered: true, currentKey: keyData, pendingKey: bytes(""), pendingActivateAt: 0});

        emit KeyRegistered(operatorSet, operator, curveType, keyData);
    }

    /// @inheritdoc IKeyRegistrar
    function deregisterKey(
        address operator,
        OperatorSet memory operatorSet
    ) external checkCanCall(operator) {
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

    /// @inheritdoc IKeyRegistrar
    function rotateKey(
        address operator,
        OperatorSet memory operatorSet,
        bytes calldata newPubkey,
        bytes calldata signature
    ) external checkCanCall(operator) {
        CurveType curveType = _operatorSetCurveTypes[operatorSet.key()];
        require(curveType != CurveType.NONE, OperatorSetNotConfigured());

        // Finalize any past-due scheduled rotation so we can schedule a new one
        _finalizeRotationIfActive(operatorSet, operator);

        KeyInfo storage info = _operatorKeyInfo[operatorSet.key()][operator];
        require(info.isRegistered, KeyNotFound(operatorSet, operator));

        // Ensure no pending rotation exists
        require(info.pendingActivateAt == 0, PendingRotationExists());

        // Calculate activation time based on AVS-configured minimum delay
        uint64 minDelay = _minRotationDelayByOperatorSet[operatorSet.key()];
        // If minDelay is max, rotation is disabled for this operator set
        require(minDelay != type(uint64).max, RotationDisabled());
        uint64 activateAt = uint64(block.timestamp) + minDelay;

        // Validate new key and reserve it globally
        _validateAndReserveKey(operatorSet, operator, newPubkey, signature, curveType);

        // Store scheduled rotation in-place
        info.pendingKey = newPubkey;
        info.pendingActivateAt = activateAt;

        emit KeyRotationScheduled(operatorSet, operator, curveType, info.currentKey, newPubkey, activateAt);
    }

    /// @inheritdoc IKeyRegistrar
    function finalizeScheduledRotation(
        address operator,
        OperatorSet memory operatorSet
    ) external returns (bool success) {
        return _finalizeRotationIfActive(operatorSet, operator);
    }

    /// @inheritdoc IKeyRegistrar
    function setMinKeyRotationDelay(
        OperatorSet memory operatorSet,
        uint64 minDelaySeconds
    ) external checkCanCall(operatorSet.avs) {
        _setMinRotationDelay(operatorSet, minDelaySeconds);
    }

    ///
    ///                         INTERNAL FUNCTIONS
    ///

    /// @notice Validate a key + signature and atomically reserve the canonical key hash.
    /// @dev For ECDSA, enforces 20-byte address format and verifies an EIP-712 signature
    ///      from the key address. For BN254, enforces the 192-byte encoding and verifies a
    ///      pairing-based signature over the EIP-712 digest. Then enforces global uniqueness
    ///      by marking the key hash as used and mapping it to the operator.
    /// @param operatorSet Operator set context bound into the signed message
    /// @param operator Operator address bound into the signed message
    /// @param keyData Raw key bytes (20 bytes for ECDSA or 192 bytes for BN254)
    /// @param signature Signature proving control of the key
    /// @param curveType The curve to use for validation
    /// @return keyHash Canonical key hash that is now reserved globally
    function _validateAndReserveKey(
        OperatorSet memory operatorSet,
        address operator,
        bytes calldata keyData,
        bytes calldata signature,
        CurveType curveType
    ) internal returns (bytes32) {
        bytes32 keyHash;
        if (curveType == CurveType.ECDSA) {
            keyHash = _validateECDSAKey(operatorSet, operator, keyData, signature);
        } else if (curveType == CurveType.BN254) {
            keyHash = _validateBN254Key(operatorSet, operator, keyData, signature);
        } else {
            revert InvalidCurveType();
        }

        require(!_globalKeyRegistry[keyHash], KeyAlreadyRegistered());
        _globalKeyRegistry[keyHash] = true;
        _keyHashToOperator[keyHash] = operator;
        return keyHash;
    }

    /// @notice If a scheduled rotation has passed activation, collapse storage to the new current key
    function _finalizeRotationIfActive(
        OperatorSet memory operatorSet,
        address operator
    ) internal returns (bool) {
        KeyInfo storage keyInfoStorage = _operatorKeyInfo[operatorSet.key()][operator];
        if (!keyInfoStorage.isRegistered) return false;
        if (keyInfoStorage.pendingActivateAt != 0 && block.timestamp >= keyInfoStorage.pendingActivateAt) {
            keyInfoStorage.currentKey = keyInfoStorage.pendingKey;
            keyInfoStorage.pendingKey = bytes("");
            keyInfoStorage.pendingActivateAt = 0;
            return true;
        }
        return false;
    }

    /// @notice Validate an ECDSA key and signature and return its key hash.
    /// @dev Ensures `keyData` is a 20-byte non-zero address, derives the EIP-712
    ///      digest with `operatorSet` and `operator`, and verifies the signature
    ///      from the key address. View-only; does not mutate state.
    /// @param operatorSet Operator set context bound into the signed message
    /// @param operator Operator address bound into the signed message
    /// @param keyData Raw ECDSA key bytes (20-byte address encoded as bytes)
    /// @param signature Signature produced by the key address over the EIP-712 digest
    /// @return keyHash Key hash for global uniqueness tracking
    function _validateECDSAKey(
        OperatorSet memory operatorSet,
        address operator,
        bytes calldata keyData,
        bytes calldata signature
    ) internal view returns (bytes32) {
        require(keyData.length == 20, InvalidKeyFormat());
        address keyAddress = address(bytes20(keyData));
        require(keyAddress != address(0), ZeroPubkey());

        bytes32 keyHash = _getKeyHashForKeyData(keyData, CurveType.ECDSA);

        bytes32 signableDigest = getECDSAKeyRegistrationMessageHash(operator, operatorSet, keyAddress);
        _checkIsValidSignatureNow(keyAddress, signableDigest, signature, type(uint256).max);

        return keyHash;
    }

    /// @notice Validate a BN254 key and signature and return its key hash.
    /// @dev Ensures `keyData` encodes (g1X,g1Y,g2X[2],g2Y[2]) totaling 192 bytes and
    ///      `signature` is 64 bytes encoding a G1 point, checks non-zero G1, computes
    ///      the EIP-712 digest with `operatorSet` and `operator`, and verifies the
    ///      pairing. View-only; does not mutate state.
    /// @param operatorSet Operator set context bound into the signed message
    /// @param operator Operator address bound into the signed message
    /// @param keyData Raw BN254 key bytes (G1 and G2 components)
    /// @param signature BN254 signature over the EIP-712 digest (G1 point: (sigX, sigY))
    /// @return keyHash Key hash for global uniqueness tracking
    function _validateBN254Key(
        OperatorSet memory operatorSet,
        address operator,
        bytes calldata keyData,
        bytes calldata signature
    ) internal view returns (bytes32) {
        require(keyData.length == 192, InvalidKeyFormat());
        require(signature.length == 64, InvalidSignature());

        bytes32 signableDigest = getBN254KeyRegistrationMessageHash(operator, operatorSet, keyData);

        bool pairingSuccessful;
        {
            (uint256 g1X, uint256 g1Y, uint256[2] memory g2X, uint256[2] memory g2Y) =
                abi.decode(keyData, (uint256, uint256, uint256[2], uint256[2]));
            require(!(g1X == 0 && g1Y == 0), ZeroPubkey());

            (uint256 sigX, uint256 sigY) = abi.decode(signature, (uint256, uint256));
            BN254.G1Point memory signaturePoint = BN254.G1Point(sigX, sigY);
            BN254.G1Point memory g1Point = BN254.G1Point(g1X, g1Y);
            BN254.G2Point memory g2Point = BN254.G2Point(g2X, g2Y);

            (, pairingSuccessful) =
                BN254SignatureVerifier.verifySignature(signableDigest, signaturePoint, g1Point, g2Point, false, 0);
        }
        require(pairingSuccessful, InvalidSignature());

        return _getKeyHashForKeyData(keyData, CurveType.BN254);
    }

    /// @notice Gets the key hash for pubkey data using consistent hashing
    /// @param pubkey The public key data
    /// @param curveType The curve type (ECDSA or BN254)
    /// @return keyHash The key hash
    function _getKeyHashForKeyData(
        bytes memory pubkey,
        CurveType curveType
    ) internal pure returns (bytes32) {
        if (curveType == CurveType.ECDSA) {
            return keccak256(pubkey);
        } else if (curveType == CurveType.BN254) {
            (uint256 g1X, uint256 g1Y,,) = abi.decode(pubkey, (uint256, uint256, uint256[2], uint256[2]));
            return BN254.hashG1Point(BN254.G1Point(g1X, g1Y));
        }

        revert InvalidCurveType();
    }

    /// @dev Internal helper to set and emit the minimum key rotation delay for an operator set
    /// @param operatorSet The operator set being configured
    /// @param minDelaySeconds The minimum rotation delay in seconds
    function _setMinRotationDelay(
        OperatorSet memory operatorSet,
        uint64 minDelaySeconds
    ) internal {
        _minRotationDelayByOperatorSet[operatorSet.key()] = minDelaySeconds;
        emit MinKeyRotationDelaySet(operatorSet, minDelaySeconds);
    }

    ///
    ///                         VIEW FUNCTIONS
    ///

    /// @inheritdoc IKeyRegistrar
    function isRegistered(
        OperatorSet memory operatorSet,
        address operator
    ) public view returns (bool) {
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

        bytes memory active = _getActiveKey(keyInfo);
        (uint256 g1X, uint256 g1Y, uint256[2] memory g2X, uint256[2] memory g2Y) =
            abi.decode(active, (uint256, uint256, uint256[2], uint256[2]));

        return (BN254.G1Point(g1X, g1Y), BN254.G2Point(g2X, g2Y));
    }

    /// @inheritdoc IKeyRegistrar
    function getECDSAKey(
        OperatorSet memory operatorSet,
        address operator
    ) public view returns (bytes memory) {
        // Validate operator set curve type
        CurveType curveType = _operatorSetCurveTypes[operatorSet.key()];
        require(curveType == CurveType.ECDSA, InvalidCurveType());

        KeyInfo memory keyInfo = _operatorKeyInfo[operatorSet.key()][operator];
        if (!keyInfo.isRegistered) return bytes("");
        bytes memory active = _getActiveKey(keyInfo);
        // Note: this is a view; we don't mutate storage here. Use finalizeScheduledRotation or rotateKey to collapse.
        return active; // 20-byte address as bytes
    }

    /// @inheritdoc IKeyRegistrar
    function getECDSAAddress(
        OperatorSet memory operatorSet,
        address operator
    ) external view returns (address) {
        return address(bytes20(getECDSAKey(operatorSet, operator)));
    }

    /// @inheritdoc IKeyRegistrar
    function isKeyGloballyRegistered(
        bytes32 keyHash
    ) external view returns (bool) {
        return _globalKeyRegistry[keyHash];
    }

    /// @inheritdoc IKeyRegistrar
    function getKeyHash(
        OperatorSet memory operatorSet,
        address operator
    ) external view returns (bytes32) {
        KeyInfo memory keyInfo = _operatorKeyInfo[operatorSet.key()][operator];
        CurveType curveType = _operatorSetCurveTypes[operatorSet.key()];

        if (!keyInfo.isRegistered) {
            return bytes32(0);
        }

        bytes memory active = _getActiveKey(keyInfo);
        return _getKeyHashForKeyData(active, curveType);
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

    /// @notice Returns the active key based on current timestamp
    /// @dev Returns `pendingKey` if a rotation is scheduled and the activation time has passed.
    ///      Otherwise returns `currentKey`.
    /// @param keyInfo The key information to check
    /// @return The active key bytes
    function _getActiveKey(
        KeyInfo memory keyInfo
    ) internal view returns (bytes memory) {
        if (keyInfo.pendingActivateAt != 0 && block.timestamp >= keyInfo.pendingActivateAt) {
            return keyInfo.pendingKey;
        }
        return keyInfo.currentKey;
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
