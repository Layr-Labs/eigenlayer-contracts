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
     * @param _version The version string for the SignatureUtilsMixin
     */
    constructor(
        IOperatorTableUpdater _operatorTableUpdater,
        string memory _version
    ) ECDSACertificateVerifierStorage(_operatorTableUpdater) SignatureUtilsMixin(_version) {
        _disableInitializers();
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

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
    ) external view returns (uint256[] memory, address[] memory) {
        (uint256[] memory signedStakes, address[] memory signers) = _verifyECDSACertificate(operatorSet, cert);
        return (signedStakes, signers);
    }

    ///@inheritdoc IECDSACertificateVerifier
    function verifyCertificateProportion(
        OperatorSet calldata operatorSet,
        ECDSACertificate calldata cert,
        uint16[] calldata totalStakeProportionThresholds
    ) external view returns (bool, address[] memory) {
        (uint256[] memory signedStakes, address[] memory signers) = _verifyECDSACertificate(operatorSet, cert);
        uint256[] memory totalStakes = getTotalStakes(operatorSet, cert.referenceTimestamp);
        require(signedStakes.length == totalStakeProportionThresholds.length, ArrayLengthMismatch());
        for (uint256 i = 0; i < signedStakes.length; i++) {
            uint256 threshold = (totalStakes[i] * totalStakeProportionThresholds[i]) / 10_000;
            if (signedStakes[i] < threshold) {
                return (false, signers);
            }
        }
        return (true, signers);
    }

    ///@inheritdoc IECDSACertificateVerifier
    function verifyCertificateNominal(
        OperatorSet calldata operatorSet,
        ECDSACertificate calldata cert,
        uint256[] memory totalStakeNominalThresholds
    ) external view returns (bool, address[] memory) {
        (uint256[] memory signedStakes, address[] memory signers) = _verifyECDSACertificate(operatorSet, cert);
        require(signedStakes.length == totalStakeNominalThresholds.length, ArrayLengthMismatch());
        for (uint256 i = 0; i < signedStakes.length; i++) {
            if (signedStakes[i] < totalStakeNominalThresholds[i]) {
                return (false, signers);
            }
        }
        return (true, signers);
    }

    /**
     * @notice Internal function to verify a certificate
     * @param cert The certificate to verify
     * @return signedStakes The amount of stake that signed the certificate for each stake type
     */
    function _verifyECDSACertificate(
        OperatorSet calldata operatorSet,
        ECDSACertificate calldata cert
    ) internal view returns (uint256[] memory, address[] memory) {
        bytes32 operatorSetKey = operatorSet.key();

        // Assert that reference timestamp is not stale
        require(block.timestamp <= cert.referenceTimestamp + _maxStalenessPeriods[operatorSetKey], CertificateStale());

        // Assert that the reference timestamp exists
        require(_latestReferenceTimestamps[operatorSetKey] == cert.referenceTimestamp, ReferenceTimestampDoesNotExist());

        // Assert that the root that corresponds to the reference timestamp is not disabled
        require(operatorTableUpdater.isRootValidByTimestamp(cert.referenceTimestamp), RootDisabled());

        // Get the total stakes
        uint256[] memory totalStakes = getTotalStakes(operatorSet, cert.referenceTimestamp);
        uint256[] memory signedStakes = new uint256[](totalStakes.length);

        // Compute the EIP-712 digest for signature recovery
        bytes32 signableDigest = calculateCertificateDigest(cert.referenceTimestamp, cert.messageHash);

        // Parse the signatures
        address[] memory signers = _parseSignatures(signableDigest, cert.sig);

        // Process each recovered signer
        for (uint256 i = 0; i < signers.length; i++) {
            address signer = signers[i];

            // Check if this signer is an operator
            bool isOperator = false;
            ECDSAOperatorInfo memory operatorInfo;

            for (uint256 j = 0; j < _numOperators[operatorSetKey][cert.referenceTimestamp]; j++) {
                operatorInfo = _operatorInfos[operatorSetKey][cert.referenceTimestamp][uint32(j)];
                if (operatorInfo.pubkey == signer) {
                    isOperator = true;
                    break;
                }
            }

            // If not an operator, the certificate is invalid
            if (!isOperator) {
                revert VerificationFailed();
            }

            // Add this operator's weights to the signed stakes
            uint256[] memory weights = operatorInfo.weights;
            for (uint256 j = 0; j < weights.length && j < signedStakes.length; j++) {
                signedStakes[j] += weights[j];
            }
        }

        return (signedStakes, signers);
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Parse signatures from the concatenated signature bytes
     * @param signableDigest The signable digest that was signed
     * @param signatures The concatenated signatures
     * @return signers Array of addresses that signed the message
     * @dev Signatures must be ordered by signer address (ascending)
     * @dev This does not support smart contract based signatures for multichain
     */
    function _parseSignatures(
        bytes32 signableDigest,
        bytes memory signatures
    ) internal pure returns (address[] memory signers) {
        // Each ECDSA signature is 65 bytes: r (32 bytes) + s (32 bytes) + v (1 byte)
        require(signatures.length > 0 && signatures.length % 65 == 0, InvalidSignatureLength());

        uint256 signatureCount = signatures.length / 65;
        signers = new address[](signatureCount);

        for (uint256 i = 0; i < signatureCount; i++) {
            bytes memory signature = new bytes(65);
            for (uint256 j = 0; j < 65; j++) {
                signature[j] = signatures[i * 65 + j];
            }

            // Recover the signer
            (address recovered, ECDSA.RecoverError err) = ECDSA.tryRecover(signableDigest, signature);
            require(err == ECDSA.RecoverError.NoError, InvalidSignature());

            // Check that signatures are ordered by signer address
            require(i == 0 || recovered > signers[i - 1], SignersNotOrdered());

            signers[i] = recovered;
        }

        return signers;
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

    /// @inheritdoc IECDSACertificateVerifier
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

    /// @inheritdoc IECDSACertificateVerifier
    function getOperatorInfo(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp,
        uint32 operatorIndex
    ) external view returns (ECDSAOperatorInfo memory) {
        bytes32 operatorSetKey = operatorSet.key();
        require(operatorIndex < _numOperators[operatorSetKey][referenceTimestamp], "Operator index out of bounds");
        return _operatorInfos[operatorSetKey][referenceTimestamp][operatorIndex];
    }

    /// @inheritdoc IECDSACertificateVerifier
    function getOperatorCount(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp
    ) external view returns (uint32) {
        bytes32 operatorSetKey = operatorSet.key();
        return uint32(_numOperators[operatorSetKey][referenceTimestamp]);
    }

    /// @inheritdoc IECDSACertificateVerifier
    function getTotalStakes(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp
    ) public view returns (uint256[] memory) {
        bytes32 operatorSetKey = operatorSet.key();
        require(_latestReferenceTimestamps[operatorSetKey] == referenceTimestamp, ReferenceTimestampDoesNotExist());
        uint256 operatorCount = _numOperators[operatorSetKey][referenceTimestamp];
        require(operatorCount > 0, ReferenceTimestampDoesNotExist());
        uint256 stakeTypesCount = _operatorInfos[operatorSetKey][referenceTimestamp][0].weights.length;
        uint256[] memory totalStakes = new uint256[](stakeTypesCount);
        for (uint256 i = 0; i < operatorCount; i++) {
            uint256[] memory weights = _operatorInfos[operatorSetKey][referenceTimestamp][uint32(i)].weights;
            for (uint256 j = 0; j < weights.length && j < stakeTypesCount; j++) {
                totalStakes[j] += weights[j];
            }
        }
        return totalStakes;
    }

    /// @inheritdoc IECDSACertificateVerifier
    function domainSeparator() public view override(IECDSACertificateVerifier, SignatureUtilsMixin) returns (bytes32) {
        return keccak256(
            abi.encode(
                EIP712_DOMAIN_TYPEHASH_NO_CHAINID,
                keccak256(bytes("EigenLayer")),
                keccak256(bytes(_majorVersion())),
                address(this)
            )
        );
    }

    /// @inheritdoc IECDSACertificateVerifier
    function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) public view returns (bytes32) {
        bytes32 structHash = keccak256(abi.encode(ECDSA_CERTIFICATE_TYPEHASH, referenceTimestamp, messageHash));
        return _calculateSignableDigest(structHash);
    }
}
