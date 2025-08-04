// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/interfaces/IECDSACertificateVerifier.sol";

contract ECDSACertificateVerifierMock is Test, IECDSACertificateVerifierTypes {
    using OperatorSetLib for OperatorSet;

    receive() external payable {}
    fallback() external payable {}

    mapping(bytes32 certificateHash => bool isValid) internal _isValidCertificate;

    mapping(bytes32 operatorSetKey => uint32 latestReferenceTimestamp) internal _latestReferenceTimestamp;
    mapping(bytes32 operatorSetKey => mapping(uint32 referenceTimestamp => bool isSet)) internal _isReferenceTimestampSet;

    function setIsValidCertificate(ECDSACertificate memory certificate, bool isValid) public {
        bytes32 certificateHash = keccak256(abi.encode(certificate));
        _isValidCertificate[certificateHash] = isValid;
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
}
