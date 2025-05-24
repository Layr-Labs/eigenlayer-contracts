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

    /// @notice Reverts if the caller does not have permission to call on behalf of `account`.
    modifier checkCanCall(
        address account
    ) {
        _checkCanCall(account);
        _;
    }

    /**
     * @notice Checks if the caller is allowed to call a function on behalf of an account.
     * @param account The account to check permissions for.
     * @dev `msg.sender` is the caller to check that can call the function on behalf of `account`.
     * @dev Returns whether the caller has permission to call the function on behalf of the account.
     */
    function _canCall(
        address account
    ) internal returns (bool) {
        return permissionController.canCall(account, msg.sender, address(this), msg.sig);
    }

    /**
     * @notice Checks if the caller is allowed to call a function on behalf of an account.
     * @param account The account to check permissions for.
     * @dev `msg.sender` is the caller to check that can call the function on behalf of `account`.
     * @dev Reverts if the caller does not have permission.
     */
    function _checkCanCall(
        address account
    ) internal {
        require(_canCall(account), InvalidPermissions());
    }
}
