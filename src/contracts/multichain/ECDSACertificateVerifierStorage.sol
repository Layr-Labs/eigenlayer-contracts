// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import "../interfaces/IOperatorTableUpdater.sol";
import "../interfaces/IECDSATableCalculator.sol";
import "../interfaces/IECDSACertificateVerifier.sol";
import "../interfaces/IBaseCertificateVerifier.sol";

abstract contract ECDSACertificateVerifierStorage is IECDSACertificateVerifier {
    // Constants

    /// @dev Basis point unit denominator for division
    uint256 internal constant BPS_DENOMINATOR = 10_000;

    /// @dev EIP-712 type hash for certificate verification
    bytes32 internal constant ECDSA_CERTIFICATE_TYPEHASH =
        keccak256("ECDSACertificate(uint32 referenceTimestamp,bytes32 messageHash)");

    /// @dev The EIP-712 domain type hash used for computing the domain separator without chainId
    bytes32 internal constant EIP712_DOMAIN_TYPEHASH_NO_CHAINID =
        keccak256("EIP712Domain(string name,string version,address verifyingContract)");

    // Immutables

    /// @dev The address that can update operator tables
    IOperatorTableUpdater public immutable operatorTableUpdater;

    // Mutatables

    /// @dev Mapping from operatorSet key to owner address
    mapping(bytes32 => address) internal _operatorSetOwners;

    /// @dev Mapping from operatorSet key to maximum staleness period
    mapping(bytes32 => uint32) internal _maxStalenessPeriods;

    /// @dev Mapping from operatorSet key to latest reference timestamp
    mapping(bytes32 => uint32) internal _latestReferenceTimestamps;

    /// @dev Mapping from referenceTimestamp to the number of operators
    mapping(bytes32 operatorSetKey => mapping(uint32 referenceTimestamp => uint256 numOperators)) internal _numOperators;

    /// @dev Mapping from operatorSetKey to referenceTimestamp to operatorInfos
    mapping(bytes32 operatorSetKey => mapping(uint32 referenceTimestamp => mapping(uint32 => ECDSAOperatorInfo)))
        internal _operatorInfos;

    // Construction

    constructor(
        IOperatorTableUpdater _operatorTableUpdater
    ) {
        operatorTableUpdater = _operatorTableUpdater;
    }
}
