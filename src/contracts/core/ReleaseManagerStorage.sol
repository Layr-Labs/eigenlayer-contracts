// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IReleaseManager.sol";

abstract contract ReleaseManagerStorage is IReleaseManager {
    // Mutables

    /// @notice Returns an array of releases for a given operator set.
    mapping(bytes32 operatorSetKey => Release[]) internal _operatorSetReleases;

    /// @notice Returns the metadata URI for a given operator set.
    mapping(bytes32 operatorSetKey => string metadataURI) internal _operatorSetMetadataURI;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[48] private __gap;
}
