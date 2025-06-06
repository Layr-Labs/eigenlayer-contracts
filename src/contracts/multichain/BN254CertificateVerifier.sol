// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

import {BN254} from "../libraries/BN254.sol";
import {Merkle} from "../libraries/Merkle.sol";
import {OperatorSet} from "../libraries/OperatorSetLib.sol";

import "../interfaces/IBN254TableCalculator.sol";
import "../interfaces/IBN254CertificateVerifier.sol";
import "../interfaces/IBaseCertificateVerifier.sol";
import "./BN254CertificateVerifierStorage.sol";

/**
 * @title BN254CertificateVerifier
 * @notice Singleton verifier for BN254 certificates across multiple operator sets
 * @dev This contract uses BN254 curves for signature verification and
 *      caches operator information for efficient verification
 */
contract BN254CertificateVerifier is Initializable, OwnableUpgradeable, BN254CertificateVerifierStorage {
    using Merkle for bytes;
    using BN254 for BN254.G1Point;

    /**
     * @notice Struct to hold verification context and reduce stack depth
     */
    struct VerificationContext {
        bytes32 operatorSetKey;
        BN254OperatorSetInfo operatorSetInfo;
        uint256[] signedStakes;
        BN254.G1Point nonSignerApk;
    }

    /**
     * @notice Restricts access to the operator table updater
     */
    modifier onlyTableUpdater() {
        require(msg.sender == _operatorTableUpdater, OnlyTableUpdater());
        _;
    }

    /**
     * @notice Constructor for the certificate verifier
     * @dev Disables initializers to prevent implementation initialization
     * @param __operatorTableUpdater Address authorized to update operator tables
     */
    constructor(
        address __operatorTableUpdater
    ) BN254CertificateVerifierStorage(__operatorTableUpdater) {
        _disableInitializers();
    }

    /**
     * @notice Initialize the contract
     * @param __owner The initial owner of the contract
     */
    function initialize(
        address __owner
    ) external initializer {
        __Ownable_init();
        _transferOwnership(__owner);
    }

    ///@inheritdoc IBaseCertificateVerifier
    function getOperatorSetOwner(
        OperatorSet memory operatorSet
    ) external view returns (address) {
        bytes32 operatorSetKey = operatorSet.key();
        return _operatorSetOwners[operatorSetKey];
    }

    ///@inheritdoc IBaseCertificateVerifier
    function maxOperatorTableStaleness(
        OperatorSet memory operatorSet
    ) external view returns (uint32) {
        bytes32 operatorSetKey = operatorSet.key();
        return _maxStalenessPeriods[operatorSetKey];
    }

    ///@inheritdoc IBaseCertificateVerifier
    function latestReferenceTimestamp(
        OperatorSet memory operatorSet
    ) external view returns (uint32) {
        bytes32 operatorSetKey = operatorSet.key();
        return _latestReferenceTimestamps[operatorSetKey];
    }

    ///@inheritdoc IBN254CertificateVerifier
    function updateOperatorTable(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp,
        BN254OperatorSetInfo memory operatorSetInfo,
        OperatorSetConfig calldata operatorSetConfig
    ) external onlyTableUpdater {
        bytes32 operatorSetKey = operatorSet.key();

        // Validate that the new timestamp is greater than the latest reference timestamp
        require(referenceTimestamp > _latestReferenceTimestamps[operatorSetKey], TableUpdateStale());

        // Store the operator set info
        _operatorSetInfos[operatorSetKey][referenceTimestamp] = operatorSetInfo;
        _latestReferenceTimestamps[operatorSetKey] = referenceTimestamp;
        _operatorSetOwners[operatorSetKey] = operatorSetConfig.owner;
        _maxStalenessPeriods[operatorSetKey] = operatorSetConfig.maxStalenessPeriod;

        emit TableUpdated(operatorSet, referenceTimestamp, operatorSetInfo);
    }

    ///@inheritdoc IBN254CertificateVerifier
    function verifyCertificate(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert
    ) external returns (uint256[] memory signedStakes) {
        return _verifyCertificate(operatorSet, cert);
    }

    ///@inheritdoc IBN254CertificateVerifier
    function verifyCertificateProportion(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert,
        uint16[] memory totalStakeProportionThresholds
    ) external returns (bool) {
        uint256[] memory signedStakes = _verifyCertificate(operatorSet, cert);

        bytes32 operatorSetKey = operatorSet.key();
        BN254OperatorSetInfo memory operatorSetInfo = _operatorSetInfos[operatorSetKey][cert.referenceTimestamp];
        uint256[] memory totalStakes = operatorSetInfo.totalWeights;

        require(signedStakes.length == totalStakeProportionThresholds.length, ArrayLengthMismatch());

        for (uint256 i = 0; i < signedStakes.length; i++) {
            // Calculate threshold as proportion of total stake
            // totalStakeProportionThresholds is in basis points (e.g. 6600 = 66%)
            uint256 threshold = (totalStakes[i] * totalStakeProportionThresholds[i]) / BPS_DENOMINATOR;

            if (signedStakes[i] < threshold) {
                return false;
            }
        }

        return true;
    }

    ///@inheritdoc IBN254CertificateVerifier
    function verifyCertificateNominal(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert,
        uint256[] memory totalStakeNominalThresholds
    ) external returns (bool) {
        uint256[] memory signedStakes = _verifyCertificate(operatorSet, cert);

        require(signedStakes.length == totalStakeNominalThresholds.length, ArrayLengthMismatch());

        for (uint256 i = 0; i < signedStakes.length; i++) {
            if (signedStakes[i] < totalStakeNominalThresholds[i]) {
                return false;
            }
        }

        return true;
    }

    /**
     * @notice Attempts signature verification with gas limit for safety
     * @param msgHash The message hash that was signed
     * @param aggPubkey The aggregate public key of signers
     * @param apkG2 The G2 point representation of the aggregate public key
     * @param signature The BLS signature to verify
     * @return pairingSuccessful Whether the pairing operation completed successfully
     * @return signatureValid Whether the signature is valid
     */
    function trySignatureVerification(
        bytes32 msgHash,
        BN254.G1Point memory aggPubkey,
        BN254.G2Point memory apkG2,
        BN254.G1Point memory signature
    ) internal view returns (bool pairingSuccessful, bool signatureValid) {
        uint256 gamma = uint256(
            keccak256(
                abi.encodePacked(
                    msgHash,
                    aggPubkey.X,
                    aggPubkey.Y,
                    apkG2.X[0],
                    apkG2.X[1],
                    apkG2.Y[0],
                    apkG2.Y[1],
                    signature.X,
                    signature.Y
                )
            )
        ) % BN254.FR_MODULUS;

        (pairingSuccessful, signatureValid) = BN254.safePairing(
            signature.plus(aggPubkey.scalar_mul(gamma)), // sigma + apk*gamma
            BN254.negGeneratorG2(), // -G2
            BN254.hashToG1(msgHash).plus(BN254.generatorG1().scalar_mul(gamma)), // H(m) + g1*gamma
            apkG2, // apkG2
            PAIRING_EQUALITY_CHECK_GAS
        );
    }

    /**
     * @notice Internal function to verify a certificate
     * @param operatorSet The operator set the certificate is for
     * @param cert The certificate to verify
     * @return signedStakes The amount of stake that signed the certificate for each stake type
     */
    function _verifyCertificate(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert
    ) internal returns (uint256[] memory signedStakes) {
        VerificationContext memory ctx;
        ctx.operatorSetKey = operatorSet.key();

        _validateCertificateTimestamp(ctx.operatorSetKey, cert.referenceTimestamp);
        ctx.operatorSetInfo = _operatorSetInfos[ctx.operatorSetKey][cert.referenceTimestamp];

        require(ctx.operatorSetInfo.operatorInfoTreeRoot != bytes32(0), ReferenceTimestampDoesNotExist());

        // Initialize signed stakes with total stakes
        ctx.signedStakes = new uint256[](ctx.operatorSetInfo.totalWeights.length);
        for (uint256 i = 0; i < ctx.operatorSetInfo.totalWeights.length; i++) {
            ctx.signedStakes[i] = ctx.operatorSetInfo.totalWeights[i];
        }

        ctx.nonSignerApk = _processNonSigners(ctx, cert);

        _verifySignature(ctx, cert);

        return ctx.signedStakes;
    }

    /**
     * @notice Validates certificate timestamp against staleness requirements
     * @param operatorSetKey The operator set key
     * @param referenceTimestamp The reference timestamp to validate
     */
    function _validateCertificateTimestamp(bytes32 operatorSetKey, uint32 referenceTimestamp) internal view {
        uint32 maxStaleness = _maxStalenessPeriods[operatorSetKey];
        require(maxStaleness == 0 || block.timestamp <= referenceTimestamp + maxStaleness, CertificateStale());
    }

    /**
     * @notice Processes non-signer witnesses and returns aggregate non-signer public key
     * @param ctx The verification context
     * @param cert The certificate being verified
     * @return nonSignerApk The aggregate public key of non-signers
     */
    function _processNonSigners(
        VerificationContext memory ctx,
        BN254Certificate memory cert
    ) internal returns (BN254.G1Point memory nonSignerApk) {
        nonSignerApk = BN254.G1Point(0, 0);

        for (uint256 i = 0; i < cert.nonSignerWitnesses.length; i++) {
            BN254OperatorInfoWitness memory witness = cert.nonSignerWitnesses[i];

            require(witness.operatorIndex < ctx.operatorSetInfo.numOperators, InvalidOperatorIndex());

            BN254OperatorInfo memory operatorInfo =
                _getOrCacheOperatorInfo(ctx.operatorSetKey, cert.referenceTimestamp, witness);

            nonSignerApk = nonSignerApk.plus(operatorInfo.pubkey);

            // Subtract non-signer stakes from total signed stakes
            for (uint256 j = 0; j < operatorInfo.weights.length; j++) {
                if (j < ctx.signedStakes.length) {
                    ctx.signedStakes[j] -= operatorInfo.weights[j];
                }
            }
        }
    }

    /**
     * @notice Gets operator info from cache or verifies and caches it
     * @param operatorSetKey The operator set key
     * @param referenceTimestamp The reference timestamp
     * @param witness The operator info witness containing proof data
     * @return operatorInfo The verified operator information
     */
    function _getOrCacheOperatorInfo(
        bytes32 operatorSetKey,
        uint32 referenceTimestamp,
        BN254OperatorInfoWitness memory witness
    ) internal returns (BN254OperatorInfo memory operatorInfo) {
        BN254OperatorInfo storage cachedInfo = _operatorInfos[operatorSetKey][referenceTimestamp][witness.operatorIndex];

        // Check if operator info is cached using pubkey existence (weights can be 0)
        bool isInfoCached = (cachedInfo.pubkey.X != 0 || cachedInfo.pubkey.Y != 0);

        if (!isInfoCached) {
            bool verified = _verifyOperatorInfoMerkleProof(
                operatorSetKey,
                referenceTimestamp,
                witness.operatorIndex,
                witness.operatorInfo,
                witness.operatorInfoProof
            );

            require(verified, VerificationFailed());

            _operatorInfos[operatorSetKey][referenceTimestamp][witness.operatorIndex] = witness.operatorInfo;
            operatorInfo = witness.operatorInfo;
        } else {
            operatorInfo = cachedInfo;
        }
    }

    /**
     * @notice Verifies the BLS signature
     * @param ctx The verification context
     * @param cert The certificate containing the signature to verify
     */
    function _verifySignature(VerificationContext memory ctx, BN254Certificate memory cert) internal view {
        // Calculate signer aggregate public key by subtracting non-signers from total
        BN254.G1Point memory signerApk = ctx.operatorSetInfo.aggregatePubkey.plus(ctx.nonSignerApk.negate());

        (bool pairingSuccessful, bool signatureValid) =
            trySignatureVerification(cert.messageHash, signerApk, cert.apk, cert.signature);

        require(pairingSuccessful && signatureValid, VerificationFailed());
    }

    /**
     * @notice Verifies a merkle proof for an operator info
     * @param operatorSetKey The operator set key
     * @param referenceTimestamp The reference timestamp
     * @param operatorIndex The index of the operator
     * @param operatorInfo The operator info
     * @param proof The merkle proof as bytes
     * @return verified Whether the proof is valid
     */
    function _verifyOperatorInfoMerkleProof(
        bytes32 operatorSetKey,
        uint32 referenceTimestamp,
        uint32 operatorIndex,
        BN254OperatorInfo memory operatorInfo,
        bytes memory proof
    ) internal view returns (bool verified) {
        bytes32 leaf = keccak256(abi.encode(operatorInfo));
        bytes32 root = _operatorSetInfos[operatorSetKey][referenceTimestamp].operatorInfoTreeRoot;
        return proof.verifyInclusionKeccak(root, leaf, operatorIndex);
    }

    /**
     * @notice Get cached operator info
     * @param operatorSet The operator set
     * @param referenceTimestamp The reference timestamp
     * @param operatorIndex The operator index
     * @return The cached operator info
     */
    function getOperatorInfo(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp,
        uint256 operatorIndex
    ) external view returns (BN254OperatorInfo memory) {
        bytes32 operatorSetKey = operatorSet.key();
        return _operatorInfos[operatorSetKey][referenceTimestamp][operatorIndex];
    }

    /**
     * @notice Get operator set info for a timestamp
     * @param operatorSet The operator set
     * @param referenceTimestamp The reference timestamp
     * @return The operator set info
     */
    function getOperatorSetInfo(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp
    ) external view returns (BN254OperatorSetInfo memory) {
        bytes32 operatorSetKey = operatorSet.key();
        return _operatorSetInfos[operatorSetKey][referenceTimestamp];
    }

    /**
     * @notice Get the current operator table updater address
     * @return The operator table updater address
     */
    function getOperatorTableUpdater() external view returns (address) {
        return _operatorTableUpdater;
    }
}
