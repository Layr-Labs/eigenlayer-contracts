// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import "../mixins/SignatureUtilsMixin.sol";
import "./ECDSACertificateVerifierStorage.sol";

/**
 * @title ECDSACertificateVerifier
 * @notice Verifies ECDSA certificates across multiple operator sets
 * @dev Implements ECDSA signature verification with operator information caching
 */
contract ECDSACertificateVerifier is Initializable, ECDSACertificateVerifierStorage, SignatureUtilsMixin {
    using ECDSA for bytes32;

    // EIP-712 type hash for certificate verification
    bytes32 public constant ECDSA_CERTIFICATE_TYPEHASH =
        keccak256("ECDSACertificate(uint32 referenceTimestamp,bytes32 messageHash)");

    /**
     * @notice Struct to hold verification context and reduce stack depth
     */
    struct VerificationContext {
        bytes32 operatorSetKey;
        ECDSAOperatorInfo[] operatorInfos;
        uint256[] signedStakes;
        address[] nonSigners;
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
     */
    constructor(
        IOperatorTableUpdater _operatorTableUpdater
    ) ECDSACertificateVerifierStorage(_operatorTableUpdater) SignatureUtilsMixin("1.0.0") {
        _disableInitializers();
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

    ///@inheritdoc IECDSACertificateVerifier
    function updateOperatorTable(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp,
        ECDSAOperatorInfo[] calldata operatorInfos,
        OperatorSetConfig calldata operatorSetConfig
    ) external onlyTableUpdater {
        bytes32 operatorSetKey = operatorSet.key();

        // Validate that the new timestamp is greater than the latest reference timestamp
        require(referenceTimestamp > _latestReferenceTimestamps[operatorSetKey], TableUpdateStale());

        // Store the number of operators
        _numOperators[operatorSetKey][referenceTimestamp] = operatorInfos.length;

        // Store each operator info in the indexed mapping
        for (uint256 i = 0; i < operatorInfos.length; i++) {
            _operatorInfos[operatorSetKey][referenceTimestamp][uint32(i)] = operatorInfos[i];
        }

        _latestReferenceTimestamps[operatorSetKey] = referenceTimestamp;
        _operatorSetOwners[operatorSetKey] = operatorSetConfig.owner;
        _maxStalenessPeriods[operatorSetKey] = operatorSetConfig.maxStalenessPeriod;

        emit TableUpdated(operatorSet, referenceTimestamp, operatorInfos);
    }

    ///@inheritdoc IECDSACertificateVerifier
    function verifyCertificate(
        OperatorSet calldata operatorSet,
        ECDSACertificate calldata cert
    ) external view returns (uint256[] memory) {
        uint96[] memory signedStakes96 = _verifyECDSACertificate(operatorSet, cert);
        uint256[] memory signedStakes = new uint256[](signedStakes96.length);
        for (uint256 i = 0; i < signedStakes96.length; i++) {
            signedStakes[i] = uint256(signedStakes96[i]);
        }
        return signedStakes;
    }

    ///@inheritdoc IECDSACertificateVerifier
    function verifyCertificateProportion(
        OperatorSet calldata operatorSet,
        ECDSACertificate calldata cert,
        uint16[] calldata totalStakeProportionThresholds
    ) external view returns (bool) {
        uint96[] memory signedStakes96 = _verifyECDSACertificate(operatorSet, cert);
        uint256[] memory signedStakes = new uint256[](signedStakes96.length);
        for (uint256 i = 0; i < signedStakes96.length; i++) {
            signedStakes[i] = uint256(signedStakes96[i]);
        }
        uint96[] memory totalStakes96 = _getTotalStakes(operatorSet, cert.referenceTimestamp);
        uint256[] memory totalStakes = new uint256[](totalStakes96.length);
        for (uint256 i = 0; i < totalStakes96.length; i++) {
            totalStakes[i] = uint256(totalStakes96[i]);
        }
        require(signedStakes.length == totalStakeProportionThresholds.length, ArrayLengthMismatch());
        for (uint256 i = 0; i < signedStakes.length; i++) {
            uint256 threshold = (totalStakes[i] * totalStakeProportionThresholds[i]) / 10_000;
            if (signedStakes[i] < threshold) {
                return false;
            }
        }
        return true;
    }

    ///@inheritdoc IECDSACertificateVerifier
    function verifyCertificateNominal(
        OperatorSet calldata operatorSet,
        ECDSACertificate calldata cert,
        uint256[] memory totalStakeNominalThresholds
    ) external view returns (bool) {
        uint96[] memory signedStakes96 = _verifyECDSACertificate(operatorSet, cert);
        uint256[] memory signedStakes = new uint256[](signedStakes96.length);
        for (uint256 i = 0; i < signedStakes96.length; i++) {
            signedStakes[i] = uint256(signedStakes96[i]);
        }
        require(signedStakes.length == totalStakeNominalThresholds.length, "Length mismatch");
        for (uint256 i = 0; i < signedStakes.length; i++) {
            if (signedStakes[i] < totalStakeNominalThresholds[i]) {
                return false;
            }
        }
        return true;
    }

    /**
     * @notice Internal function to verify a certificate
     * @param cert The certificate to verify
     * @return signedStakes The amount of stake that signed the certificate for each stake type
     */
    function _verifyECDSACertificate(
        OperatorSet calldata operatorSet,
        ECDSACertificate calldata cert
    ) internal view returns (uint96[] memory) {
        bytes32 operatorSetKey = operatorSet.key();

        // Assert that reference timestamp is not stale
        require(block.timestamp <= cert.referenceTimestamp + _maxStalenessPeriods[operatorSetKey], CertificateStale());

        // Assert that the reference timestamp exists
        require(_latestReferenceTimestamps[operatorSetKey] == cert.referenceTimestamp, ReferenceTimestampDoesNotExist());

        // Get the total stakes
        uint96[] memory totalStakes = _getTotalStakes(operatorSet, cert.referenceTimestamp);
        uint96[] memory signedStakes = new uint96[](totalStakes.length);

        // Compute the EIP-712 digest for signature recovery
        bytes32 structHash =
            keccak256(abi.encode(ECDSA_CERTIFICATE_TYPEHASH, cert.referenceTimestamp, cert.messageHash));
        bytes32 signableDigest = keccak256(abi.encodePacked("\x19\x01", domainSeparator(), structHash));

        // Parse the signatures
        (address[] memory signers, bool validSignatures) = _parseSignatures(signableDigest, cert.sig);

        require(validSignatures, VerificationFailed());

        // Process each operator to check if they signed
        uint256 operatorCount = _numOperators[operatorSetKey][cert.referenceTimestamp];
        for (uint256 i = 0; i < operatorCount; i++) {
            // Check if this operator is in the signers list
            bool isSigner = false;
            for (uint256 j = 0; j < signers.length; j++) {
                if (_operatorInfos[operatorSetKey][cert.referenceTimestamp][uint32(i)].pubkey == signers[j]) {
                    isSigner = true;
                    break;
                }
            }

            if (isSigner) {
                // Add this operator's weights to the signed stakes
                uint256[] storage weights = _operatorInfos[operatorSetKey][cert.referenceTimestamp][uint32(i)].weights;
                for (uint256 j = 0; j < weights.length && j < signedStakes.length; j++) {
                    signedStakes[j] += uint96(weights[j]);
                }
            }
        }

        // After processing, check if all signed stakes are zero
        bool anySigned = false;
        for (uint256 i = 0; i < signedStakes.length; i++) {
            if (signedStakes[i] > 0) {
                anySigned = true;
                break;
            }
        }
        require(anySigned, VerificationFailed());

        return signedStakes;
    }

    /**
     * @notice Parse signatures from the concatenated signature bytes
     * @param messageHash The message hash that was signed
     * @param signatures The concatenated signatures
     * @return signers Array of addresses that signed the message
     * @return valid Whether all signatures are valid
     */
    function _parseSignatures(
        bytes32 messageHash,
        bytes memory signatures
    ) internal pure returns (address[] memory signers, bool valid) {
        // Each ECDSA signature is 65 bytes: r (32 bytes) + s (32 bytes) + v (1 byte)
        require(signatures.length % 65 == 0, "Invalid signature length");

        uint256 signatureCount = signatures.length / 65;
        signers = new address[](signatureCount);

        for (uint256 i = 0; i < signatureCount; i++) {
            bytes memory signature = new bytes(65);
            for (uint256 j = 0; j < 65; j++) {
                signature[j] = signatures[i * 65 + j];
            }

            // Recover the signer
            address signer = messageHash.recover(signature);

            // If any signature is invalid (returns address(0)), the whole certificate is invalid
            if (signer == address(0)) {
                return (signers, false);
            }

            // Check for duplicate signers
            for (uint256 j = 0; j < i; j++) {
                if (signers[j] == signer) {
                    return (signers, false);
                }
            }

            signers[i] = signer;
        }

        return (signers, true);
    }

    /**
     * @notice Get operator infos for a timestamp
     * @param operatorSet The operator set
     * @param referenceTimestamp The reference timestamp
     * @return The operator infos
     */
    function getOperatorInfos(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp
    ) external view returns (ECDSAOperatorInfo[] memory) {
        bytes32 operatorSetKey = operatorSet.key();
        uint32 numOperators = uint32(_numOperators[operatorSetKey][referenceTimestamp]);
        ECDSAOperatorInfo[] memory operatorInfos = new ECDSAOperatorInfo[](numOperators);

        for (uint32 i = 0; i < numOperators; i++) {
            operatorInfos[i] = _operatorInfos[operatorSetKey][referenceTimestamp][i];
        }

        return operatorInfos;
    }

    /**
     * @notice Calculate the total stakes for all operators at a given reference timestamp
     * @param operatorSet The operator set to calculate stakes for
     * @param referenceTimestamp The reference timestamp
     * @return totalStakes The total stakes for all operators
     */
    function _getTotalStakes(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp
    ) internal view returns (uint96[] memory totalStakes) {
        bytes32 operatorSetKey = operatorSet.key();
        require(_latestReferenceTimestamps[operatorSetKey] == referenceTimestamp, ReferenceTimestampDoesNotExist());
        uint256 operatorCount = _numOperators[operatorSetKey][referenceTimestamp];
        require(operatorCount > 0, ReferenceTimestampDoesNotExist());
        uint256 stakeTypesCount = _operatorInfos[operatorSetKey][referenceTimestamp][0].weights.length;
        totalStakes = new uint96[](stakeTypesCount);
        for (uint256 i = 0; i < operatorCount; i++) {
            uint256[] storage weights = _operatorInfos[operatorSetKey][referenceTimestamp][uint32(i)].weights;
            for (uint256 j = 0; j < weights.length && j < stakeTypesCount; j++) {
                totalStakes[j] += uint96(weights[j]);
            }
        }
        return totalStakes;
    }
}
