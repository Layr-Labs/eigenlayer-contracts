// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/interfaces/IECDSACertificateVerifier.sol";
import "src/contracts/interfaces/IOperatorTableUpdater.sol";

abstract contract ECDSACertificateVerifierStorage is IECDSACertificateVerifier {
    /// @notice The operatorTableUpdater contract
    IOperatorTableUpdater public immutable operatorTableUpdater;

    constructor(
        IOperatorTableUpdater _operatorTableUpdater
    ) {
        operatorTableUpdater = _operatorTableUpdater;
    }

    /// @dev Mapping from operatorSetKey to latest referenceTimestamp
    mapping(bytes32 operatorSetKey => uint32 latestReferenceTimestamp) internal _latestReferenceTimestamp;

    /// @dev Mapping from operatorSetKey to max staleness period
    mapping(bytes32 operatorSetKey => uint32 maxStalenessPeriod) internal _maxStalenessPeriod;

    /// @dev Mapping from operatorSetKey to owner
    mapping(bytes32 operatorSetKey => address owner) internal _owner;

    /// @dev Mapping from referenceTimestamp to the number of operators
    mapping(bytes32 operatorSetKey => mapping(uint32 referenceTimestamp => uint256 numOperators)) internal _numOperators;

    /// @dev Mapping from operatorSetKey to referenceTimestamp to operatorInfos
    mapping(bytes32 operatorSetKey => mapping(uint32 referenceTimestamp => ECDSAOperatorInfo[] operatorInfos)) internal
        _operatorInfos;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[45] private __gap;
}
