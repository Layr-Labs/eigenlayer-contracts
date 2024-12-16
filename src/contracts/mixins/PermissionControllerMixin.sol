// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "../interfaces/IPermissionController.sol";

abstract contract PermissionControllerMixin {
    /// @dev Thrown when the caller is not allowed to call a function on behalf of an account.
    error InvalidPermissions();

    /// @notice Pointer to the permission controller contract.
    IPermissionController public immutable permissionController;

    constructor(
        IPermissionController _permissionController
    ) {
        permissionController = _permissionController;
    }

    /// @notice Checks if the caller (msg.sender) can call on behalf of an account.
    modifier checkCanCall(
        address account
    ) {
        require(_checkCanCall(account), InvalidPermissions());
        _;
    }

    /**
     * @notice Checks if the caller is allowed to call a function on behalf of an account.
     * @param account the account to check
     * @dev `msg.sender` is the caller to check that can call the function on behalf of `account`.
     * @dev Returns a bool, instead of reverting
     */
    function _checkCanCall(
        address account
    ) internal returns (bool) {
        return permissionController.canCall(account, msg.sender, address(this), msg.sig);
    }
}
