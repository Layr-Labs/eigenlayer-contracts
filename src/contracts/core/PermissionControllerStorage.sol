// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IPermissionController.sol";

abstract contract PermissionControllerStorage is IPermissionController {
    /// @notice Mapping (account => admin)
    mapping(address => address) accountToAdmin;
}