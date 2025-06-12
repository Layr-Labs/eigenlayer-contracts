// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../libraries/Snapshots.sol";
import "../interfaces/IReleaseManager.sol";
import "../interfaces/IAllocationManager.sol";

abstract contract ReleaseManagerStorage is IReleaseManager {
    // Immutables

    /// @notice The EigenLayer AllocationManager contract.
    IAllocationManager public immutable allocationManager;

    // Mutables

    /// @notice A set of all release digests for an operator set.
    /// @dev operatorSet, releaseId => (releaseDigest => (uint32(deprecationTimestamp), uint16(major), uint16(minor), uint16(patch)))
    mapping(bytes32 operatorSetKey => EnumerableMap.Bytes32ToUintMap releaseDigests) internal _releaseDigests;

    /// @notice The deployment deadline for a release digest.
    /// @dev operatorSet => (uint224(releaseId))
    mapping(bytes32 operatorSetKey => Snapshots.DefaultZeroHistory) internal _releaseSnapshots;

    /// @notice The registry URL for a release digest.
    mapping(bytes32 operatorSetKey => mapping(bytes32 releaseDigest => string registryUrl)) internal
        _releaseRegistryUrls;

    constructor(
        IAllocationManager _allocationManager
    ) {
        allocationManager = _allocationManager;
    }

    // TODO: modify gap to account for odd variable sizes.

    uint256[46] private __gap;
}

