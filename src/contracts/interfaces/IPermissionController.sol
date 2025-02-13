// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

interface IPermissionControllerErrors {
    /// @notice Thrown when the caller is not an admin of the account.
    error NotAdmin();
    /// @notice Thrown when attempting to remove an admin that is not set.
    error AdminNotSet();
    /// @notice Thrown when attempting to set an appointee that is already set for the account's function.
    error AppointeeAlreadySet();
    /// @notice Thrown when attempting to remove an appointee that is not set for the account's function.
    error AppointeeNotSet();
    /// @notice Thrown when attempting to remove the last admin from an account.
    error CannotHaveZeroAdmins();
    /// @notice Thrown when attempting to add an admin that is already set.
    error AdminAlreadySet();
    /// @notice Thrown when attempting to accept admin role without being a pending admin.
    error AdminNotPending();
    /// @notice Thrown when attempting to add a pending admin that is already pending.
    error AdminAlreadyPending();
}

interface IPermissionControllerEvents {
    /// @notice Emitted when an appointee is granted permission to call a function for an account.
    event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector);

    /// @notice Emitted when an appointee's permission to call a function is revoked for an account.
    event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector);

    /// @notice Emitted when an address is added as a pending admin for an account.
    event PendingAdminAdded(address indexed account, address admin);

    /// @notice Emitted when an address is removed from pending admins for an account.
    event PendingAdminRemoved(address indexed account, address admin);

    /// @notice Emitted when a pending admin accepts their role and becomes an admin for an account.
    event AdminSet(address indexed account, address admin);

    /// @notice Emitted when an admin is removed from an account.
    event AdminRemoved(address indexed account, address admin);
}

interface IPermissionController is IPermissionControllerErrors, IPermissionControllerEvents {
    /**
     * @notice Sets a pending admin for an account.
     * @param account The account to set the pending admin for.
     * @param admin The address to set as pending admin.
     * @dev Multiple admins can be set for an account.
     */
    function addPendingAdmin(address account, address admin) external;

    /**
     * @notice Removes a pending admin from an account.
     * @param account The account to remove the pending admin from.
     * @param admin The pending admin address to remove.
     * @dev Only the admin of the account can remove a pending admin.
     */
    function removePendingAdmin(address account, address admin) external;

    /**
     * @notice Accepts the admin role for an account.
     * @param account The account to accept the admin role for.
     * @dev Only a pending admin for the account can become an admin.
     */
    function acceptAdmin(
        address account
    ) external;

    /**
     * @notice Removes an admin from an account.
     * @param account The account to remove the admin from.
     * @param admin The admin address to remove.
     * @dev Only the admin of the account can remove another admin.
     * @dev Reverts if removing this admin would leave the account with no admins.
     */
    function removeAdmin(address account, address admin) external;

    /**
     * @notice Sets an appointee's permission to call a specific function for an account.
     * @param account The account to set the appointee for.
     * @param appointee The address to grant permission to.
     * @param target The contract address containing the function.
     * @param selector The function selector to grant permission for.
     * @dev Only the admin of the account can set an appointee.
     */
    function setAppointee(address account, address appointee, address target, bytes4 selector) external;

    /**
     * @notice Removes an appointee's permission to call a specific function for an account.
     * @param account The account to remove the appointee from.
     * @param appointee The address to remove permission from.
     * @param target The contract address containing the function.
     * @param selector The function selector to remove permission for.
     * @dev Only the admin of the account can remove an appointee.
     */
    function removeAppointee(address account, address appointee, address target, bytes4 selector) external;

    /**
     * @notice Checks if a given caller is an admin of the account.
     * @param account The account to check admin status for.
     * @param caller The address to check admin status of.
     * @return Returns true if the caller is an admin, false otherwise.
     * @dev If the account has no admin, the caller must be the account itself.
     */
    function isAdmin(address account, address caller) external view returns (bool);

    /**
     * @notice Checks if an address is a pending admin for an account.
     * @param account The account to check pending admin status for.
     * @param pendingAdmin The address to check pending admin status of.
     * @return Returns true if the address is a pending admin, false otherwise.
     */
    function isPendingAdmin(address account, address pendingAdmin) external view returns (bool);

    /**
     * @notice Gets the list of admins for an account.
     * @param account The account to get the admins for.
     * @return Returns an array of admin addresses.
     * @dev If the account has no admin, returns an array containing only the account address.
     */
    function getAdmins(
        address account
    ) external view returns (address[] memory);

    /**
     * @notice Gets the list of pending admins for an account.
     * @param account The account to get the pending admins for.
     * @return Returns an array of pending admin addresses.
     */
    function getPendingAdmins(
        address account
    ) external view returns (address[] memory);

    /**
     * @notice Checks if a caller has permission to call a specific function for an account.
     * @param account The account to check permissions for.
     * @param caller The address to check permissions of.
     * @param target The contract address containing the function.
     * @param selector The function selector to check permissions for.
     * @return Returns true if the caller has permission, false otherwise.
     * @dev Returns true if the caller is either an admin or an appointed caller for the function.
     * @dev Function upgrades that modify parameters will invalidate existing permissions.
     */
    function canCall(address account, address caller, address target, bytes4 selector) external returns (bool);

    /**
     * @notice Gets the list of functions an appointee has permission to call for an account.
     * @param account The account to get appointee permissions for.
     * @param appointee The appointee address to get permissions for.
     * @return Returns two arrays: target contract addresses and their corresponding function selectors.
     */
    function getAppointeePermissions(
        address account,
        address appointee
    ) external returns (address[] memory, bytes4[] memory);

    /**
     * @notice Gets the list of appointees that can call a specific function for an account.
     * @param account The account to get appointees for.
     * @param target The contract address containing the function.
     * @param selector The function selector to get appointees for.
     * @return Returns an array of appointee addresses.
     * @dev Does not include admins in the returned list, even though they have permission to call the function.
     */
    function getAppointees(address account, address target, bytes4 selector) external returns (address[] memory);
}
