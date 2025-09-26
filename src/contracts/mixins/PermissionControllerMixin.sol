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
        _checkCanCall(account);
        _;
    }

    /**
     * @notice Checks if the caller (msg.sender) is permitted to call the current function on behalf of the given account.
     * @param account The account on whose behalf the function is being called.
     * @dev Reverts if the caller is not permitted to call the current function on behalf of the given account.
     * @dev This function queries the permissionController to determine if msg.sender is authorized
     *      to call the current function (identified by msg.sig) on behalf of `account`.
     */
    function _checkCanCall(
        address account
    ) internal view {
        require(_canCall(account), InvalidPermissions());
    }

    /**
     * @notice Checks if the caller (msg.sender) is permitted to call the current function on behalf of the given account.
     * @param account The account on whose behalf the function is being called.
     * @return allowed True if the caller is permitted, false otherwise.
     */
    function _canCall(
        address account
    ) internal view returns (bool allowed) {
        return permissionController.canCall(account, msg.sender, address(this), msg.sig);
    }
}
