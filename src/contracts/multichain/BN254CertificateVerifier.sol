// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import {BN254} from "../libraries/BN254.sol";
import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import {
    IBN254TableCalculator, IBN254TableCalculatorTypes
} from "../interfaces/IBN254TableCalculator.sol";
import {
    IBN254CertificateVerifier,
    IBN254CertificateVerifierTypes
} from "../interfaces/IBN254CertificateVerifier.sol";
import {IBaseCertificateVerifier} from "../interfaces/IBaseCertificateVerifier.sol";
import {Merkle} from "../libraries/Merkle.sol";
import "./BN254CertificateVerifierStorage.sol";

/**
 * @title BN254CertificateVerifier
 * @notice Singleton verifier for BN254 certificates across multiple operator sets
 * @dev This contract uses BN254 curves for signature verification and
 *      caches operator information for efficient verification
 */
contract BN254CertificateVerifier is 
    Initializable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable,
    BN254CertificateVerifierStorage 
{
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

    // Modifier to restrict access to the operator table updater
    modifier onlyTableUpdater() {
        if (msg.sender != _operatorTableUpdater) revert OnlyTableUpdater();
        _;
    }

    /**
     * @notice Constructor for the certificate verifier
     * @dev Disables initializers to prevent implementation initialization
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize the contract
     * @param __operatorTableUpdater The address that can update operator tables
     * @param __owner The initial owner of the contract
     */
    function initialize(
        address __operatorTableUpdater,
        address __owner
    ) external initializer {
        __Ownable_init();
        __ReentrancyGuard_init();
        
        _operatorTableUpdater = __operatorTableUpdater;
        _transferOwnership(__owner);
    }

    /**
     * @inheritdoc IBaseCertificateVerifier
     */
    function getOperatorSetOwner(
        OperatorSet memory operatorSet
    ) external view returns (address) {
        bytes32 operatorSetKey = operatorSet.key();
        return _operatorSetOwners[operatorSetKey];
    }

    /**
     * @inheritdoc IBaseCertificateVerifier
     */
    function maxOperatorTableStaleness(
        OperatorSet memory operatorSet
    ) external view returns (uint32) {
        bytes32 operatorSetKey = operatorSet.key();
        return _maxStalenessPeriods[operatorSetKey];
    }

    /**
     * @inheritdoc IBaseCertificateVerifier
     */
    function latestReferenceTimestamp(
        OperatorSet memory operatorSet
    ) external view returns (uint32) {
        bytes32 operatorSetKey = operatorSet.key();
        return _latestReferenceTimestamps[operatorSetKey];
    }

    /**
     * @notice Set the operator table updater address
     * @param newOperatorTableUpdater The new operator table updater address
     */
    function setOperatorTableUpdater(address newOperatorTableUpdater) external onlyOwner {
        _operatorTableUpdater = newOperatorTableUpdater;
    }

    /**
     * @inheritdoc IBN254CertificateVerifier
     */
    function updateOperatorTable(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp,
        BN254OperatorSetInfo memory operatorSetInfo,
        OperatorSetConfig calldata operatorSetConfig
    ) external onlyTableUpdater {
        bytes32 operatorSetKey = operatorSet.key();
        
        // Require that the new timestamp is greater than the latest reference timestamp
        require(referenceTimestamp > _latestReferenceTimestamps[operatorSetKey], TableUpdateStale());

        // Store the operator set info
        _operatorSetInfos[operatorSetKey][referenceTimestamp] = operatorSetInfo;
        _latestReferenceTimestamps[operatorSetKey] = referenceTimestamp;
        _operatorSetOwners[operatorSetKey] = operatorSetConfig.owner;
        _maxStalenessPeriods[operatorSetKey] = operatorSetConfig.maxStalenessPeriod;

        // Emit event
        emit TableUpdated(operatorSet, referenceTimestamp, operatorSetInfo);
    }

    /**
     * @inheritdoc IBN254CertificateVerifier
     */
    function verifyCertificate(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert
    ) external returns (uint256[] memory signedStakes) {
        return _verifyCertificate(operatorSet, cert);
    }

    /**
     * @inheritdoc IBN254CertificateVerifier
     */
    function verifyCertificateProportion(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert,
        uint16[] memory totalStakeProportionThresholds
    ) external returns (bool) {
        // Get signed stakes
        uint256[] memory signedStakes = _verifyCertificate(operatorSet, cert);

        // Get total stakes from the operator set info
        bytes32 operatorSetKey = operatorSet.key();
        BN254OperatorSetInfo memory operatorSetInfo = 
            _operatorSetInfos[operatorSetKey][cert.referenceTimestamp];
        uint256[] memory totalStakes = operatorSetInfo.totalWeights;

        // Verify that each stake meets the threshold
        if (signedStakes.length != totalStakeProportionThresholds.length) {
            revert ArrayLengthMismatch();
        }

        for (uint256 i = 0; i < signedStakes.length; i++) {
            // Calculate threshold as proportion of total stake
            // totalStakeProportionThresholds is in basis points (e.g. 6600 = 66%)
            uint256 threshold = (totalStakes[i] * totalStakeProportionThresholds[i]) / 10000;

            // If signed stake doesn't meet threshold, return false
            if (signedStakes[i] < threshold) {
                return false;
            }
        }

        return true;
    }

    /**
     * @inheritdoc IBN254CertificateVerifier
     */
    function verifyCertificateNominal(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert,
        uint256[] memory totalStakeNominalThresholds
    ) external returns (bool) {
        // Get signed stakes
        uint256[] memory signedStakes = _verifyCertificate(operatorSet, cert);

        // Verify that each stake meets the threshold
        if (signedStakes.length != totalStakeNominalThresholds.length) {
            revert ArrayLengthMismatch();
        }

        for (uint256 i = 0; i < signedStakes.length; i++) {
            // If signed stake doesn't meet nominal threshold, return false
            if (signedStakes[i] < totalStakeNominalThresholds[i]) {
                return false;
            }
        }

        return true;
    }

    /**
     * @notice Try signature verification with gas limit for safety
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
        
        // Check staleness and get operator set info
        _validateCertificateTimestamp(ctx.operatorSetKey, cert.referenceTimestamp);
        ctx.operatorSetInfo = _operatorSetInfos[ctx.operatorSetKey][cert.referenceTimestamp];
        
        // Check that this reference timestamp exists
        if (ctx.operatorSetInfo.operatorInfoTreeRoot == bytes32(0)) {
            revert ReferenceTimestampDoesNotExist();
        }

        // Initialize signed stakes with total stakes
        ctx.signedStakes = new uint256[](ctx.operatorSetInfo.totalWeights.length);
        for (uint256 i = 0; i < ctx.operatorSetInfo.totalWeights.length; i++) {
            ctx.signedStakes[i] = ctx.operatorSetInfo.totalWeights[i];
        }

        // Process non-signers
        ctx.nonSignerApk = _processNonSigners(ctx, cert);

        // Verify signature
        _verifySignature(ctx, cert);

        return ctx.signedStakes;
    }

    /**
     * @notice Validates certificate timestamp against staleness requirements
     */
    function _validateCertificateTimestamp(
        bytes32 operatorSetKey,
        uint32 referenceTimestamp
    ) internal view {
        uint32 maxStaleness = _maxStalenessPeriods[operatorSetKey];
        if (maxStaleness > 0 && block.timestamp > referenceTimestamp + maxStaleness) {
            revert CertificateStale();
        }
    }

    /**
     * @notice Processes non-signer witnesses and returns aggregate non-signer public key
     */
    function _processNonSigners(
        VerificationContext memory ctx,
        BN254Certificate memory cert
    ) internal returns (BN254.G1Point memory nonSignerApk) {
        nonSignerApk = BN254.G1Point(0, 0);

        for (uint256 i = 0; i < cert.nonSignerWitnesses.length; i++) {
            BN254OperatorInfoWitness memory witness = cert.nonSignerWitnesses[i];
            
            // Validate index
            if (witness.operatorIndex >= ctx.operatorSetInfo.numOperators) {
                revert InvalidOperatorIndex();
            }

            // Get or cache operator info
            BN254OperatorInfo memory operatorInfo = _getOrCacheOperatorInfo(
                ctx.operatorSetKey,
                cert.referenceTimestamp,
                witness
            );

            // Aggregate non-signer public key
            nonSignerApk = nonSignerApk.plus(operatorInfo.pubkey);

            for (uint256 j = 0; j < operatorInfo.weights.length; j++) {
                if (j < ctx.signedStakes.length) {
                    ctx.signedStakes[j] -= operatorInfo.weights[j];
                }
            }
        }
    }

    /**
     * @notice Gets operator info from cache or verifies and caches it
     */
    function _getOrCacheOperatorInfo(
        bytes32 operatorSetKey,
        uint32 referenceTimestamp,
        BN254OperatorInfoWitness memory witness
    ) internal returns (BN254OperatorInfo memory operatorInfo) {
        // Check cache first
        BN254OperatorInfo storage cachedInfo = 
            _operatorInfos[operatorSetKey][referenceTimestamp][witness.operatorIndex];

        // weights can be 0, so check if operator has been cached using bn254 pubkey
        bool isInfoCached = (cachedInfo.pubkey.X != 0 || cachedInfo.pubkey.Y != 0);

        if (!isInfoCached) {
            // Verify merkle proof
            bool verified = _verifyOperatorInfoMerkleProof(
                operatorSetKey,
                referenceTimestamp,
                witness.operatorIndex,
                witness.operatorInfo,
                witness.operatorInfoProof
            );

            if (!verified) {
                revert VerificationFailed();
            }

            // Cache the operator info
            _operatorInfos[operatorSetKey][referenceTimestamp][witness.operatorIndex] = witness.operatorInfo;
            operatorInfo = witness.operatorInfo;
        } else {
            operatorInfo = cachedInfo;
        }
    }

    /**
     * @notice Verifies the BLS signature
     */
    function _verifySignature(
        VerificationContext memory ctx,
        BN254Certificate memory cert
    ) internal view {
        // Calculate signer aggregate public key
        BN254.G1Point memory signerApk = ctx.operatorSetInfo.aggregatePubkey.plus(ctx.nonSignerApk.negate());

        // Verify the BLS signature
        (bool pairingSuccessful, bool signatureValid) =
            trySignatureVerification(cert.messageHash, signerApk, cert.apk, cert.signature);

        require(pairingSuccessful && signatureValid, VerificationFailed());
    }

    /**
     * @notice Verifies a merkle proof for an operator info
     * @param operatorSetKey The operatorSet key
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