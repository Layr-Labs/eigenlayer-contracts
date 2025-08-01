// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IOperatorTableUpdater.sol";
import "../interfaces/IBN254CertificateVerifier.sol";
import "../interfaces/IBaseCertificateVerifier.sol";

abstract contract BN254CertificateVerifierStorage is IBN254CertificateVerifier {
    // Constants

    /// @dev Gas limit for pairing operations to prevent DoS attacks
    uint256 internal constant PAIRING_EQUALITY_CHECK_GAS = 400_000;

    /// @dev Basis point unit denominator for division
    uint256 internal constant BPS_DENOMINATOR = 10_000;

    // OPERATOR_INFO_LEAF_SALT is now inherited from LeafCalculatorMixin

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

    /// @dev Mapping from operatorSet key to reference timestamp to operator set info
    mapping(bytes32 => mapping(uint32 => BN254OperatorSetInfo)) internal _operatorSetInfos;

    /// @dev Mapping from operatorSet key to reference timestamp to operator index to operator info
    /// This is used to cache operator info that has been proven against a tree root
    mapping(bytes32 => mapping(uint32 => mapping(uint256 => BN254OperatorInfo))) internal _operatorInfos;

    /// @dev Mapping from operatorSet key to reference timestamp to whether it has been updated
    mapping(bytes32 operatorSetKey => mapping(uint32 referenceTimestamp => bool set)) internal _referenceTimestampsSet;

    // Construction

    constructor(
        IOperatorTableUpdater _operatorTableUpdater
    ) {
        operatorTableUpdater = _operatorTableUpdater;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[44] private __gap;
}
