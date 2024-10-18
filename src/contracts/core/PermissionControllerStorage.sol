// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IPermissionController.sol";

abstract contract PermissionControllerStorage is IPermissionController {
    /// @notice Mapping (account => admin)
    mapping(address => address) accountToAdmin;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}
