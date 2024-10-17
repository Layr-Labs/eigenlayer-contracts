// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "../interfaces/IPermissionController.sol";

library PermissionControllerLib {
    function canCall(IPermissionController permissionController, address account, address caller) internal view returns (bool) {
        return permissionController.canCall(account, caller);
    }
}