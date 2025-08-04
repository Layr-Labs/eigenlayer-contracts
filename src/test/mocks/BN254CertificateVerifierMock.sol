// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/interfaces/IBN254CertificateVerifier.sol";
import "src/contracts/interfaces/IOperatorTableUpdater.sol";
import "src/contracts/interfaces/IBaseCertificateVerifier.sol";

contract BN254CertificateVerifierMock is Test, IBN254CertificateVerifierTypes, ICrossChainRegistryTypes, IBaseCertificateVerifierErrors {
    using OperatorSetLib for OperatorSet;

    receive() external payable {}
    fallback() external payable {}

    mapping(bytes32 certificateHash => bool isValid) internal _isValidCertificate;

    mapping(bytes32 operatorSetKey => mapping(uint32 referenceTimestamp => BN254OperatorSetInfo operatorSetInfo)) internal _operatorSetInfos;
    mapping(bytes32 operatorSetKey => address owner) internal _operatorSetOwners;
    mapping(bytes32 operatorSetKey => uint32 maxStalenessPeriod) internal _maxStalenessPeriods;
    mapping(bytes32 operatorSetKey => uint32 latestReferenceTimestamp) internal _latestReferenceTimestamp;
    mapping(bytes32 operatorSetKey => mapping(uint32 referenceTimestamp => bool isSet)) internal _isReferenceTimestampSet;

    // Mock operator table updater for testing
    IOperatorTableUpdater public operatorTableUpdater;

    function setOperatorTableUpdater(IOperatorTableUpdater _operatorTableUpdater) public {
        operatorTableUpdater = _operatorTableUpdater;
    }

    function setIsValidCertificate(BN254Certificate memory certificate, bool isValid) public {
        bytes32 certificateHash = keccak256(abi.encode(certificate));
        _isValidCertificate[certificateHash] = isValid;
    }

    function updateOperatorTable(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp,
        BN254OperatorSetInfo memory operatorSetInfo,
        OperatorSetConfig memory operatorSetConfig
    ) public {
        _latestReferenceTimestamp[operatorSet.key()] = referenceTimestamp;
        _operatorSetInfos[operatorSet.key()][referenceTimestamp] = operatorSetInfo;
        _operatorSetOwners[operatorSet.key()] = operatorSetConfig.owner;
        _maxStalenessPeriods[operatorSet.key()] = operatorSetConfig.maxStalenessPeriod;
        _isReferenceTimestampSet[operatorSet.key()][referenceTimestamp] = true;
    }

    function verifyCertificateProportion(
        OperatorSet memory operatorSet,
        BN254Certificate memory certificate,
        uint16[] memory stakeProportionThresholds
    ) external view returns (bool) {
        // Always validate certificate timestamp like the real contract does
        _validateCertificateTimestamp(operatorSet.key(), certificate.referenceTimestamp);

        bytes32 certificateHash = keccak256(abi.encode(certificate));
        return _isValidCertificate[certificateHash];
    }

    function setLatestReferenceTimestamp(OperatorSet memory operatorSet, uint32 referenceTimestamp) public {
        _latestReferenceTimestamp[operatorSet.key()] = referenceTimestamp;
    }

    function latestReferenceTimestamp(OperatorSet memory operatorSet) external view returns (uint32) {
        return _latestReferenceTimestamp[operatorSet.key()];
    }

    function setIsReferenceTimestampSet(OperatorSet memory operatorSet, uint32 referenceTimestamp, bool isSet) public {
        _isReferenceTimestampSet[operatorSet.key()][referenceTimestamp] = isSet;
    }

    function isReferenceTimestampSet(OperatorSet memory operatorSet, uint32 referenceTimestamp) external view returns (bool) {
        return _isReferenceTimestampSet[operatorSet.key()][referenceTimestamp];
    }

    /**
     * @notice Validates certificate timestamp against staleness requirements
     * @param operatorSetKey The operator set key
     * @param referenceTimestamp The reference timestamp to validate
     */
    function _validateCertificateTimestamp(bytes32 operatorSetKey, uint32 referenceTimestamp) internal view {
        uint32 maxStaleness = _maxStalenessPeriods[operatorSetKey];
        require(maxStaleness == 0 || block.timestamp <= referenceTimestamp + maxStaleness, CertificateStale());

        // Assert that the root that corresponds to the reference timestamp is not disabled
        require(operatorTableUpdater.isRootValidByTimestamp(referenceTimestamp), RootDisabled());
    }

    /**
     * @notice Public function to validate certificate timestamp (for testing)
     * @param operatorSet The operator set
     * @param referenceTimestamp The reference timestamp to validate
     */
    function validateCertificateTimestamp(OperatorSet memory operatorSet, uint32 referenceTimestamp) external view {
        _validateCertificateTimestamp(operatorSet.key(), referenceTimestamp);
    }
}
