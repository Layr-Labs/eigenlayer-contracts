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

    /**
     * @notice Modifier that checks if the caller can call on behalf of an account, reverts if not permitted.
     * @param account The account on whose behalf the function is being called.
     * @dev Use this modifier when the entire function requires authorization.
     * @dev This is the most common pattern - prefer this over `_checkCanCall` when possible.
     */
    modifier checkCanCall(
        address account
    ) {
        _checkCanCall(account);
        _;
    }

    /**
     * @notice Checks if the caller is permitted to call the current function on behalf of the given account.
     * @param account The account on whose behalf the function is being called.
     * @dev Reverts with `InvalidPermissions()` if the caller is not permitted.
     * @dev Use this function instead of the modifier when:
     *      - You need to avoid "stack too deep" errors (e.g., when combining multiple modifiers)
     *      - You need more control over when the check occurs in your function logic
     */
    function _checkCanCall(
        address account
    ) internal view {
        require(_canCall(account), InvalidPermissions());
    }

    /**
     * @notice Checks if the caller is permitted to call the current function on behalf of the given account.
     * @param account The account on whose behalf the function is being called.
     * @return allowed True if the caller is permitted, false otherwise.
     * @dev Unlike `_checkCanCall`, this function returns a boolean instead of reverting.
     * @dev Use this function when you need conditional logic based on permissions, such as:
     *      - OR conditions: `require(_canCall(operator) || _canCall(avs), InvalidCaller());`
     *      - If-else branches: `if (_canCall(account)) { ... } else { ... }`
     *      - Multiple authorization paths in the same function
     * @dev This function queries the permissionController to determine if msg.sender is authorized
     *      to call the current function (identified by msg.sig) on behalf of `account`.
     */
    function _canCall(
        address account
    ) internal view returns (bool allowed) {
        return permissionController.canCall(account, msg.sender, address(this), msg.sig);
    }
}
