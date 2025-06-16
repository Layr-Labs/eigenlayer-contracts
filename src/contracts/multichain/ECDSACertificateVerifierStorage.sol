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
