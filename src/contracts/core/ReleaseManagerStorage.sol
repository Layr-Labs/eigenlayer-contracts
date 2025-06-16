// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IReleaseManager.sol";
import "../interfaces/IAllocationManager.sol";

abstract contract ReleaseManagerStorage is IReleaseManager {
    // Mutables

    /// @notice Returns an array of releases for a given operator set.
    mapping(bytes32 operatorSetKey => Release[]) internal _operatorSetReleases;

    uint256[49] private __gap;
}
