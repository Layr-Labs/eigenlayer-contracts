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

    /// @notice Maps operator set keys to their version history.
    /// @param operatorSetKey The key identifying the operator set.
    /// @return versions The set of version keys for this operator set.
    mapping(bytes32 operatorSetKey => EnumerableSet.Bytes32Set versions) internal _versions;

    /// @notice Maps operator set and version keys to their release details.
    /// @param operatorSetKey The key identifying the operator set.
    /// @param versionKey The key identifying the version.
    /// @return release The release details including artifacts and deprecation time.
    mapping(bytes32 operatorSetKey => mapping(bytes32 versionKey => Release release)) internal _releases;

    /// @notice Maps operator set keys and major versions to their upgrade deadlines.
    /// @param operatorSetKey The key identifying the operator set.
    /// @param major The major version number.
    /// @return upgradeByTime The timestamp by which operators must upgrade to this major version.
    mapping(bytes32 operatorSetKey => mapping(uint16 major => uint32 upgradeByTime)) internal _upgradeByTimes;

    constructor(
        IAllocationManager _allocationManager
    ) {
        allocationManager = _allocationManager;
    }

    uint256[47] private __gap;
}
