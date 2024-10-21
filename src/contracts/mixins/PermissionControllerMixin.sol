// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "../interfaces/IPermissionController.sol";

abstract contract PermissionControllerMixin {
    /// @dev Thrown when the caller is not allowed to call a function on behalf of an account.
    error InvalidCaller();

    /// @notice Pointer to the permission controller contract.
    IPermissionController public immutable permissionController;

    constructor(
        IPermissionController _permissionController
    ) {
        permissionController = _permissionController;
    }

    modifier checkCanCall(address caller) {
        _checkCanCall(msg.sender, caller);
        _;
    }

    /**
     * @notice Checks if the caller is allowed to call a function on behalf of an account.
     * @param account the account to check
     * @param caller the caller to check admin status for
     */
    function _checkCanCall(address account, address caller) internal {
        require(permissionController.canCall(account, caller), InvalidCaller());
    }
}
