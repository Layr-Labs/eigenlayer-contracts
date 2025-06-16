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

    /// @notice Returns an array of releases for a given operator set.
    mapping(bytes32 operatorSetKey => Release[]) internal _operatorSetReleases;

    constructor(
        IAllocationManager _allocationManager
    ) {
        allocationManager = _allocationManager;
    }

    uint256[49] private __gap;
}
