// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

import "../libraries/BN254.sol";
import "../libraries/BN254SignatureVerifier.sol";
import "../libraries/Merkle.sol";
import "../libraries/OperatorSetLib.sol";
import "../mixins/SemVerMixin.sol";
import "../mixins/LeafCalculatorMixin.sol";
import "./BN254CertificateVerifierStorage.sol";

/**
 * @title BN254CertificateVerifier
 * @notice Singleton verifier for BN254 certificates across multiple operator sets
 * @dev This contract uses BN254 curves for signature verification and
 *      caches operator information for efficient verification
 */
contract BN254CertificateVerifier is
    Initializable,
    BN254CertificateVerifierStorage,
    SemVerMixin,
    LeafCalculatorMixin
{
    using Merkle for bytes;
    using BN254 for BN254.G1Point;

    /**
     * @notice Struct to hold verification context and reduce stack depth
     */
    struct VerificationContext {
        bytes32 operatorSetKey;
        BN254OperatorSetInfo operatorSetInfo;
        uint256[] totalSignedStakeWeights;
        BN254.G1Point nonSignerApk;
    }

    /**
     * @notice Restricts access to the operator table updater
     */
    modifier onlyTableUpdater() {
        require(msg.sender == address(operatorTableUpdater), OnlyTableUpdater());
        _;
    }

    /**
     * @notice Constructor for the certificate verifier
     * @dev Disables initializers to prevent implementation initialization
     * @param _operatorTableUpdater Address authorized to update operator tables
     * @param _version The semantic version of the contract
     */
    constructor(
        IOperatorTableUpdater _operatorTableUpdater,
        string memory _version
    ) BN254CertificateVerifierStorage(_operatorTableUpdater) SemVerMixin(_version) {
        _disableInitializers();
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

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
        _referenceTimestampsSet[operatorSetKey][referenceTimestamp] = true;

        emit TableUpdated(operatorSet, referenceTimestamp, operatorSetInfo);
    }

    ///@inheritdoc IBN254CertificateVerifier
    function verifyCertificate(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert
    ) external returns (uint256[] memory totalSignedStakeWeights) {
        return _verifyCertificate(operatorSet, cert);
    }

    ///@inheritdoc IBN254CertificateVerifier
    function verifyCertificateProportion(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert,
        uint16[] memory totalStakeProportionThresholds
    ) external returns (bool) {
        uint256[] memory totalSignedStakeWeights = _verifyCertificate(operatorSet, cert);

        bytes32 operatorSetKey = operatorSet.key();
        BN254OperatorSetInfo memory operatorSetInfo = _operatorSetInfos[operatorSetKey][cert.referenceTimestamp];
        uint256[] memory totalStakes = operatorSetInfo.totalWeights;

        require(totalSignedStakeWeights.length == totalStakeProportionThresholds.length, ArrayLengthMismatch());

        for (uint256 i = 0; i < totalSignedStakeWeights.length; i++) {
            // Calculate threshold as proportion of total signed stake weight
            // totalStakeProportionThresholds is in basis points (e.g. 6600 = 66%)
            uint256 threshold = (totalStakes[i] * totalStakeProportionThresholds[i]) / BPS_DENOMINATOR;

            if (totalSignedStakeWeights[i] < threshold) {
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
        uint256[] memory totalSignedStakeWeights = _verifyCertificate(operatorSet, cert);

        require(totalSignedStakeWeights.length == totalStakeNominalThresholds.length, ArrayLengthMismatch());

        for (uint256 i = 0; i < totalSignedStakeWeights.length; i++) {
            if (totalSignedStakeWeights[i] < totalStakeNominalThresholds[i]) {
                return false;
            }
        }

        return true;
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Internal function to verify a certificate
     * @param operatorSet The operator set the certificate is for
     * @param cert The certificate to verify
     * @return totalSignedStakeWeights The amount of stake that signed the certificate for each stake type
     */
    function _verifyCertificate(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert
    ) internal returns (uint256[] memory totalSignedStakeWeights) {
        VerificationContext memory ctx;
        ctx.operatorSetKey = operatorSet.key();

        _validateCertificateTimestamp(ctx.operatorSetKey, cert.referenceTimestamp);
        ctx.operatorSetInfo = _operatorSetInfos[ctx.operatorSetKey][cert.referenceTimestamp];

        // Initialize signed stakes with total stake weights
        ctx.totalSignedStakeWeights = new uint256[](ctx.operatorSetInfo.totalWeights.length);
        for (uint256 i = 0; i < ctx.operatorSetInfo.totalWeights.length; i++) {
            ctx.totalSignedStakeWeights[i] = ctx.operatorSetInfo.totalWeights[i];
        }

        ctx.nonSignerApk = _processNonSigners(ctx, cert);

        _verifySignature(ctx, cert);

        return ctx.totalSignedStakeWeights;
    }

    /**
     * @notice Validates certificate timestamp against staleness requirements
     * @param operatorSetKey The operator set key
     * @param referenceTimestamp The reference timestamp to validate
     */
    function _validateCertificateTimestamp(bytes32 operatorSetKey, uint32 referenceTimestamp) internal view {
        // Assert that the certificate is not stale
        uint32 maxStaleness = _maxStalenessPeriods[operatorSetKey];
        require(maxStaleness == 0 || block.timestamp <= referenceTimestamp + maxStaleness, CertificateStale());

        // Assert that the reference timestamp has been set
        require(_referenceTimestampsSet[operatorSetKey][referenceTimestamp], ReferenceTimestampDoesNotExist());

        // Assert that the root that corresponds to the reference timestamp is not disabled
        require(operatorTableUpdater.isRootValidByTimestamp(referenceTimestamp), RootDisabled());
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
                _getOrCacheNonsignerOperatorInfo(ctx.operatorSetKey, cert.referenceTimestamp, witness);

            nonSignerApk = nonSignerApk.plus(operatorInfo.pubkey);

            // Subtract non-signer stakes from total signed stakes
            for (uint256 j = 0; j < operatorInfo.weights.length; j++) {
                if (j < ctx.totalSignedStakeWeights.length) {
                    ctx.totalSignedStakeWeights[j] -= operatorInfo.weights[j];
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
    function _getOrCacheNonsignerOperatorInfo(
        bytes32 operatorSetKey,
        uint32 referenceTimestamp,
        BN254OperatorInfoWitness memory witness
    ) internal returns (BN254OperatorInfo memory operatorInfo) {
        BN254OperatorInfo memory cachedInfo = _operatorInfos[operatorSetKey][referenceTimestamp][witness.operatorIndex];

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

        // Compute EIP-712 digest that binds referenceTimestamp and messageHash
        bytes32 signableDigest = calculateCertificateDigest(cert.referenceTimestamp, cert.messageHash);

        (bool pairingSuccessful, bool signatureValid) =
            trySignatureVerification(signableDigest, signerApk, cert.apk, cert.signature);

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
        bytes32 leaf = calculateOperatorInfoLeaf(operatorInfo);
        bytes32 root = _operatorSetInfos[operatorSetKey][referenceTimestamp].operatorInfoTreeRoot;
        return proof.verifyInclusionKeccak(root, leaf, operatorIndex);
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

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

    ///@inheritdoc IBaseCertificateVerifier
    function isReferenceTimestampSet(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp
    ) external view returns (bool) {
        bytes32 operatorSetKey = operatorSet.key();
        return _referenceTimestampsSet[operatorSetKey][referenceTimestamp];
    }

    ///@inheritdoc IBaseCertificateVerifier
    function getTotalStakeWeights(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp
    ) external view returns (uint256[] memory) {
        bytes32 operatorSetKey = operatorSet.key();
        return _operatorSetInfos[operatorSetKey][referenceTimestamp].totalWeights;
    }

    /// @inheritdoc IBaseCertificateVerifier
    function getOperatorCount(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp
    ) external view returns (uint256) {
        bytes32 operatorSetKey = operatorSet.key();
        return _operatorSetInfos[operatorSetKey][referenceTimestamp].numOperators;
    }

    ///@inheritdoc IBN254CertificateVerifier
    function trySignatureVerification(
        bytes32 msgHash,
        BN254.G1Point memory aggPubkey,
        BN254.G2Point memory apkG2,
        BN254.G1Point memory signature
    ) public view returns (bool pairingSuccessful, bool signatureValid) {
        return BN254SignatureVerifier.verifySignature(
            msgHash,
            signature,
            aggPubkey,
            apkG2,
            true, // use gas limit
            PAIRING_EQUALITY_CHECK_GAS
        );
    }

    ///@inheritdoc IBN254CertificateVerifier
    function getNonsignerOperatorInfo(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp,
        uint256 operatorIndex
    ) external view returns (BN254OperatorInfo memory) {
        bytes32 operatorSetKey = operatorSet.key();
        return _operatorInfos[operatorSetKey][referenceTimestamp][operatorIndex];
    }

    ///@inheritdoc IBN254CertificateVerifier
    function isNonsignerCached(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp,
        uint256 operatorIndex
    ) external view returns (bool) {
        bytes32 operatorSetKey = operatorSet.key();
        BN254OperatorInfo memory operatorInfo = _operatorInfos[operatorSetKey][referenceTimestamp][operatorIndex];
        // Check if operator info is cached using pubkey existence (weights can be 0)
        return operatorInfo.pubkey.X != 0 && operatorInfo.pubkey.Y != 0;
    }

    ///@inheritdoc IBN254CertificateVerifier
    function getOperatorSetInfo(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp
    ) external view returns (BN254OperatorSetInfo memory) {
        bytes32 operatorSetKey = operatorSet.key();
        return _operatorSetInfos[operatorSetKey][referenceTimestamp];
    }

    /// @inheritdoc IBN254CertificateVerifier
    function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) public pure returns (bytes32) {
        return keccak256(abi.encode(BN254_CERTIFICATE_TYPEHASH, referenceTimestamp, messageHash));
    }
}
