// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/interfaces/IBN254CertificateVerifier.sol";

contract BN254CertificateVerifierMock is Test, IBN254CertificateVerifierTypes {
    using OperatorSetLib for OperatorSet;

    receive() external payable {}
    fallback() external payable {}

    mapping(bytes32 certificateHash => bool isValid) internal _isValidCertificate;

    mapping(bytes32 operatorSetKey => uint32 latestReferenceTimestamp) internal _latestReferenceTimestamp;

    function setIsValidCertificate(BN254Certificate memory certificate, bool isValid) public {
        bytes32 certificateHash = keccak256(abi.encode(certificate));
        _isValidCertificate[certificateHash] = isValid;
    }

    function verifyCertificateProportion(
        OperatorSet memory operatorSet,
        BN254Certificate memory certificate,
        uint16[] memory stakeProportionThresholds
    ) external view returns (bool) {
        bytes32 certificateHash = keccak256(abi.encode(certificate));
        return _isValidCertificate[certificateHash];
    }

    function setLatestReferenceTimestamp(OperatorSet memory operatorSet, uint32 referenceTimestamp) public {
        _latestReferenceTimestamp[operatorSet.key()] = referenceTimestamp;
    }

    function latestReferenceTimestamp(OperatorSet memory operatorSet) external view returns (uint32) {
        return _latestReferenceTimestamp[operatorSet.key()];
    }
}
