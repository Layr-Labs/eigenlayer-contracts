// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/mixins/SemVerMixin.sol";
import "src/contracts/multichain/ECDSACertificateVerifierStorage.sol";

abstract contract ECDSACertificateVerifier is Initializable, ECDSACertificateVerifierStorage, SemVerMixin {
    using ECDSA for bytes32;
    using OperatorSetLib for OperatorSet;

    modifier onlyOperatorTableUpdater() {
        require(msg.sender == address(operatorTableUpdater), OnlyTableUpdater());
        _;
    }

    /// @dev Sets the immutable variables for the contract
    constructor(
        IOperatorTableUpdater _operatorTableUpdater,
        string memory _semVer
    ) ECDSACertificateVerifierStorage(_operatorTableUpdater) SemVerMixin(_semVer) {
        _disableInitializers();
    }

    /**
     *
     *                         TABLE UPDATING FUNCTIONS
     *
     */

    /// @inheritdoc IECDSACertificateVerifier
    function updateOperatorTable(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp,
        ECDSAOperatorInfo[] calldata operatorInfos,
        OperatorSetConfig calldata operatorSetConfig
    ) external onlyOperatorTableUpdater {
        // Check that the table update is greater than the latest reference timestamp
        require(referenceTimestamp >= latestReferenceTimestamp(operatorSet), TableUpdateStale());

        // Store the number of operators
        bytes32 operatorSetKey = operatorSet.key();
        _numOperators[operatorSetKey][referenceTimestamp] = operatorInfos.length;

        // Store each operatorInfo
        for (uint256 i = 0; i < operatorInfos.length; i++) {
            // Update the pubkey of the operator
            _operatorInfos[operatorSetKey][referenceTimestamp][i].pubkey = operatorInfos[i].pubkey;

            // Copy each weight
            for (uint256 j = 0; j < operatorInfos[i].weights.length; j++) {
                _operatorInfos[operatorSetKey][referenceTimestamp][i].weights[j] = operatorInfos[i].weights[j];
            }
        }

        // Update the latest reference timestamp
        _latestReferenceTimestamp[operatorSetKey] = referenceTimestamp;

        // Update the operatorSetConfig
        _updateOperatorSetConfig(operatorSet, operatorSetConfig);

        // Emit the event
        emit TableUpdated(operatorSet, referenceTimestamp, operatorInfos);
    }

    /**
     *
     *           CERTIFICATE VERIFICATION FUNCTIONS
     *
     */

    /// @inheritdoc IECDSACertificateVerifier
    function verifyCertificate(
        OperatorSet calldata operatorSet,
        ECDSACertificate calldata cert
    ) external view returns (uint96[] memory) {
        return _verifyECDSACertificate(operatorSet, cert);
    }

    /// @inheritdoc IECDSACertificateVerifier
    function verifyCertificateProportion(
        OperatorSet calldata operatorSet,
        ECDSACertificate calldata cert,
        uint16[] calldata totalStakeProportionThresholds
    ) external view returns (bool) {
        // Get signed stakes
        uint96[] memory signedStakes = _verifyECDSACertificate(operatorSet, cert);

        // Get total stakes
        uint96[] memory totalStakes = _getTotalStakes(operatorSet, cert.referenceTimestamp);

        // Verify that each stake meets the threshold
        require(signedStakes.length == totalStakeProportionThresholds.length, ArrayLengthMismatch());

        for (uint256 i = 0; i < signedStakes.length; i++) {
            // Calculate threshold as proportion of total stake
            // totalStakeProportionThresholds is a percentage with 2 decimal places (e.g. 6600 = 66%)
            uint96 threshold = uint96(uint256(totalStakes[i]) * uint256(totalStakeProportionThresholds[i]) / 10_000);

            // If signed stake doesn't meet threshold, return false
            if (signedStakes[i] < threshold) {
                return false;
            }
        }
    }

    /// @inheritdoc IECDSACertificateVerifier
    function verifyCertificateNominal(
        OperatorSet calldata operatorSet,
        ECDSACertificate calldata cert,
        uint96[] memory totalStakeNominalThresholds
    ) external view returns (bool) {
        // Get signed stakes
        uint96[] memory signedStakes = _verifyECDSACertificate(operatorSet, cert);

        // Verify that each stake meets the threshold
        require(signedStakes.length == totalStakeNominalThresholds.length, "Length mismatch");

        for (uint256 i = 0; i < signedStakes.length; i++) {
            // If signed stake doesn't meet nominal threshold, return false
            if (signedStakes[i] < totalStakeNominalThresholds[i]) {
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

    /// @dev Updates the operatorSetConfig for a given operatorSet
    function _updateOperatorSetConfig(
        OperatorSet calldata operatorSet,
        OperatorSetConfig calldata operatorSetConfig
    ) internal {
        _setOperatorSetOwner(operatorSet, operatorSetConfig.owner);
        _setMaxStalenessPeriod(operatorSet, operatorSetConfig.maxStalenessPeriod);
    }

    /// @dev Sets the owner of a given operatorSet
    function _setOperatorSetOwner(OperatorSet calldata operatorSet, address owner) internal {
        _owner[operatorSet.key()] = owner;
        emit OperatorSetOwnerUpdated(operatorSet, owner);
    }

    /// @dev Sets the max staleness period for a given operatorSet
    function _setMaxStalenessPeriod(OperatorSet calldata operatorSet, uint32 maxStalenessPeriod) internal {
        _maxStalenessPeriod[operatorSet.key()] = maxStalenessPeriod;
        emit MaxStalenessPeriodUpdated(operatorSet, maxStalenessPeriod);
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
        require(block.timestamp > cert.referenceTimestamp + _maxStalenessPeriod[operatorSetKey], CertificateStale());

        // Assert that the reference timestamp exists
        require(_latestReferenceTimestamp[operatorSetKey] == cert.referenceTimestamp, ReferenceTimestampDoesNotExist());

        // Get the total stakes
        uint96[] memory totalStakes = _getTotalStakes(operatorSet, cert.referenceTimestamp);
        uint96[] memory signedStakes = new uint96[](totalStakes.length);

        // Parse the signatures
        (address[] memory signers, bool validSignatures) = _parseSignatures(cert.messageHash, cert.sig);

        if (!validSignatures) {
            revert VerificationFailed();
        }

        // Process each operator to check if they signed
        uint256 operatorCount = _numOperators[operatorSetKey][cert.referenceTimestamp];
        for (uint256 i = 0; i < operatorCount; i++) {
            // Check if this operator is in the signers list
            bool isSigner = false;
            for (uint256 j = 0; j < signers.length; j++) {
                if (_operatorInfos[operatorSetKey][cert.referenceTimestamp][i].pubkey == signers[j]) {
                    isSigner = true;
                    break;
                }
            }

            if (isSigner) {
                // Add this operator's weights to the signed stakes
                uint96[] storage weights = _operatorInfos[operatorSetKey][cert.referenceTimestamp][i].weights;
                for (uint256 j = 0; j < weights.length && j < signedStakes.length; j++) {
                    signedStakes[j] += weights[j];
                }
            }
        }

        return signedStakes;
    }

    /**
     * @notice Calculate the total stakes for all operators at a given reference timestamp
     * @param referenceTimestamp The reference timestamp
     * @return totalStakes The total stakes for all operators
     */
    function _getTotalStakes(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp
    ) internal view returns (uint96[] memory totalStakes) {
        // Ensure the reference timestamp exists
        uint256 operatorCount = _numOperators[operatorSet.key()][referenceTimestamp];
        require(operatorCount > 0, ReferenceTimestampDoesNotExist());

        // Use the first operator to determine the number of stake types
        uint256 stakeTypesCount = _operatorInfos[operatorSet.key()][referenceTimestamp][0].weights.length;
        totalStakes = new uint96[](stakeTypesCount);

        // Sum up all stakes for all operators
        for (uint256 i = 0; i < operatorCount; i++) {
            uint96[] storage weights = _operatorInfos[operatorSet.key()][referenceTimestamp][i].weights;
            for (uint256 j = 0; j < weights.length && j < stakeTypesCount; j++) {
                totalStakes[j] += weights[j];
            }
        }

        return totalStakes;
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
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IBaseCertificateVerifier
    function latestReferenceTimestamp(
        OperatorSet calldata operatorSet
    ) public view returns (uint32) {
        return _latestReferenceTimestamp[operatorSet.key()];
    }

    /// @inheritdoc IBaseCertificateVerifier
    function getOperatorSetOwner(
        OperatorSet memory operatorSet
    ) external view returns (address) {
        return _owner[operatorSet.key()];
    }

    /// @inheritdoc IBaseCertificateVerifier
    function maxOperatorTableStaleness(
        OperatorSet memory operatorSet
    ) external view returns (uint32) {
        return _maxStalenessPeriod[operatorSet.key()];
    }
}
