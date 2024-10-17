import "../interfaces/IPermissionController.sol";

library PermissionControllerLib {
    /// @dev Thrown when the caller is not allowed to call a function on behalf of an account.
    error InvalidCaller();
    function canCall(IPermissionController permissionController, address account, address caller) internal view returns (bool) {
        require(permissionController.canCall(account, caller), InvalidCaller());
    }
}