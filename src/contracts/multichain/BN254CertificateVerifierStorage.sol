// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import "../interfaces/IBN254TableCalculator.sol";
import "../interfaces/IBN254CertificateVerifier.sol";
import "../interfaces/IBaseCertificateVerifier.sol";

abstract contract BN254CertificateVerifierStorage is IBN254CertificateVerifier {
    // Constants

    /// @dev Gas limit for pairing operations to prevent DoS attacks
    uint256 internal constant PAIRING_EQUALITY_CHECK_GAS = 400_000;

    /// @dev Basis point unit denominator for division
    uint256 internal constant BPS_DENOMINATOR = 10_000;

    // Immutables - None in this case, but could be added if needed

    // Mutatables

    /// @dev The address that can update operator tables
    address immutable _operatorTableUpdater;

    /// @dev Mapping from operatorSet key to owner address
    mapping(bytes32 => address) internal _operatorSetOwners;

    /// @dev Mapping from operatorSet key to maximum staleness period
    mapping(bytes32 => uint32) internal _maxStalenessPeriods;

    /// @dev Mapping from operatorSet key to latest reference timestamp
    mapping(bytes32 => uint32) internal _latestReferenceTimestamps;

    /// @dev Mapping from operatorSet key to reference timestamp to operator set info
    mapping(bytes32 => mapping(uint32 => IBN254TableCalculatorTypes.BN254OperatorSetInfo)) internal _operatorSetInfos;

    /// @dev Mapping from operatorSet key to reference timestamp to operator index to operator info
    /// This is used to cache operator info that has been proven against a tree root
    mapping(bytes32 => mapping(uint32 => mapping(uint256 => IBN254TableCalculatorTypes.BN254OperatorInfo))) internal
        _operatorInfos;

    // Construction

    constructor(
        address __operatorTableUpdater
    ) {
        _operatorTableUpdater = __operatorTableUpdater;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[45] private __gap;
}
