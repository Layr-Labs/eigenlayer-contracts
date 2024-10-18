// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "./PermissionControllerStorage.sol";

contract PermissionController is Initializable, PermissionControllerStorage {
    constructor() {
        _disableInitializers();
    }

    /// @inheritdoc IPermissionController
    function setAdmin(address account, address newAdmin) external {
        address currentAdmin = accountToAdmin[account];
        if (currentAdmin == address(0)) {
            require(msg.sender == account, CannotSetAdmin());
        } else {
            require(currentAdmin == msg.sender, CannotSetAdmin());
        }

        require(newAdmin != address(0), InvalidAdmin());

        accountToAdmin[account] = newAdmin;
        emit AdminSet(account, currentAdmin, newAdmin);
    }

    /// @inheritdoc IPermissionController
    function canCall(address account, address caller) external view returns (bool) {
        address admin = accountToAdmin[account];
        if (admin == address(0)) {
            return account == caller;
        } else {
            return admin == caller;
        }
    }

    /// @inheritdoc IPermissionController
    function getAdmin(
        address account
    ) external view returns (address) {
        address admin = accountToAdmin[account];
        return admin == address(0) ? account : admin;
    }
}
