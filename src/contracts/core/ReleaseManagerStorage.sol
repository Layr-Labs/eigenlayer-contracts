// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "../interfaces/IReleaseManager.sol";
import "../interfaces/IAllocationManager.sol";

abstract contract ReleaseManagerStorage is IReleaseManager {
    // Immutables

    /// @notice The EigenLayer AllocationManager contract.
    IAllocationManager public immutable allocationManager;

    // Mutables

    mapping(bytes32 operatorSetKey => EnumerableSet.Bytes32Set versions) internal _versions;

    mapping(bytes32 operatorSetKey => mapping(bytes32 versionKey => Release release)) internal _releases;

    mapping(bytes32 operatorSetKey => mapping(uint16 major => uint32 upgradeByTime)) internal _upgradeByTimes;

    constructor(
        IAllocationManager _allocationManager
    ) {
        allocationManager = _allocationManager;
    }

    uint256[48] private __gap;
}
