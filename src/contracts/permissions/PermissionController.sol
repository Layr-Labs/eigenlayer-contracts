// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "./PermissionControllerStorage.sol";

contract PermissionController is Initializable, PermissionControllerStorage {
    using EnumerableSet for *;

    modifier onlyAdmin(
        address account
    ) {
        require(isAdmin(account, msg.sender), NotAdmin());
        _;
    }

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */
    constructor() {
        _disableInitializers();
    }

    function initialize() external initializer {}

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /// @inheritdoc IPermissionController
    function addAdmin(address account, address admin) external onlyAdmin(account) {
        // Add the admin to the account's admins
        // If the admin is already set, the set will fail
        EnumerableSet.AddressSet storage admins = _permissions[account].admins;
        require(admins.add(admin), AdminAlreadySet());

        emit AdminSet(account, admin);
    }

    /// @inheritdoc IPermissionController
    function removeAdmin(address account, address admin) external onlyAdmin(account) {
        EnumerableSet.AddressSet storage admins = _permissions[account].admins;

        require(admins.length() > 1, CannotHaveZeroAdmins());

        // Remove the admin from the account's admins
        // If the admin is not set, the remove will fail
        require(admins.remove(admin), AdminNotSet());

        emit AdminRemoved(account, admin);
    }

    /// @inheritdoc IPermissionController
    function setAppointee(
        address account,
        address appointee,
        address target,
        bytes4 selector
    ) external onlyAdmin(account) {
        AccountPermissions storage permissions = _permissions[account];

        bytes32 targetSelector = _encodeTargetSelector(target, selector);
        require(!permissions.appointeePermissions[appointee].contains(targetSelector), AppointeeAlreadySet());

        // Add the appointee to the account's permissions
        permissions.appointeePermissions[appointee].add(targetSelector);
        permissions.permissionAppointees[targetSelector].add(appointee);

        emit AppointeeSet(account, appointee, target, selector);
    }

    /// @inheritdoc IPermissionController
    function removeAppointee(
        address account,
        address appointee,
        address target,
        bytes4 selector
    ) external onlyAdmin(account) {
        AccountPermissions storage permissions = _permissions[account];

        bytes32 targetSelector = _encodeTargetSelector(target, selector);
        require(permissions.appointeePermissions[appointee].contains(targetSelector), AppointeeNotSet());

        // Remove the appointee from the account's permissions
        permissions.appointeePermissions[appointee].remove(targetSelector);
        permissions.permissionAppointees[targetSelector].remove(appointee);

        emit AppointeeRemoved(account, appointee, target, selector);
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /// @dev Encodes the target and selector into a single bytes32 values
    /// @dev Encoded Format: [160 bits target][32 bits selector][64 bits padding],
    function _encodeTargetSelector(address target, bytes4 selector) internal pure returns (bytes32) {
        // Reserve 96 bits for the target
        uint256 shiftedTarget = uint256(uint160(target)) << 96;
        // Reserve 32 bits for the selector
        uint256 shiftedSelector = uint256(uint32(selector)) << 64;
        // Combine the target and selector
        return bytes32(shiftedTarget | shiftedSelector);
    }

    /// @dev Decodes the target and selector from a single bytes32 value
    /// @dev Encoded Format: [160 bits target][32 bits selector][64 bits padding], 
    function _decodeTargetSelector(
        bytes32 targetSelector
    ) internal pure returns (address, bytes4) {
        // The target is in the upper 160 bits of the targetSelector
        address target = address(uint160(uint256(targetSelector) >> 96));
        // The selector is in the lower 32 bits after the padding is removed
        bytes4 selector = bytes4(uint32(uint256(targetSelector) >> 64));

        return (target, selector);
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IPermissionController
    function isAdmin(address account, address caller) public view returns (bool) {
        if (_permissions[account].admins.length() == 0) {
            // If the account does not have an admin, the caller must be the account
            return account == caller;
        } else {
            // If the account has an admin, the caller must be an admin
            return _permissions[account].admins.contains(caller);
        }
    }

    /// @inheritdoc IPermissionController
    function getAdmins(
        address account
    ) external view returns (address[] memory) {
        if (_permissions[account].admins.length() == 0) {
            address[] memory admin = new address[](1);
            admin[0] = account;
            return admin;
        } else {
            return _permissions[account].admins.values();
        }
    }

    /// @inheritdoc IPermissionController
    function canCall(address account, address caller, address target, bytes4 selector) external view returns (bool) {
        return isAdmin(account, caller)
            || _permissions[account].appointeePermissions[caller].contains(_encodeTargetSelector(target, selector));
    }

    /// @inheritdoc IPermissionController
    function getAppointeePermissions(
        address account,
        address appointee
    ) external view returns (address[] memory, bytes4[] memory) {
        EnumerableSet.Bytes32Set storage appointeePermissions = _permissions[account].appointeePermissions[appointee];

        uint256 length = appointeePermissions.length();

        address[] memory targets = new address[](length);
        bytes4[] memory selectors = new bytes4[](length);

        for (uint256 i = 0; i < length; i++) {
            (address target, bytes4 selector) = _decodeTargetSelector(appointeePermissions.at(i));
            targets[i] = target;
            selectors[i] = selector;
        }

        return (targets, selectors);
    }

    /// @inheritdoc IPermissionController
    function getAppointees(address account, address target, bytes4 selector) external view returns (address[] memory) {
        bytes32 targetSelector = _encodeTargetSelector(target, selector);
        return _permissions[account].permissionAppointees[targetSelector].values();
    }
}
