// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "./ISemVerMixin.sol";

interface IPermissionControllerErrors {
    /// @notice Thrown when a non-admin caller attempts to perform an admin-only action.
    error NotAdmin();
    /// @notice Thrown when attempting to remove an admin that does not exist.
    error AdminNotSet();
    /// @notice Thrown when attempting to set an appointee for a function that already has one.
    error AppointeeAlreadySet();
    /// @notice Thrown when attempting to interact with a non-existent appointee.
    error AppointeeNotSet();
    /// @notice Thrown when attempting to remove the last remaining admin.
    error CannotHaveZeroAdmins();
    /// @notice Thrown when attempting to set an admin that is already registered.
    error AdminAlreadySet();
    /// @notice Thrown when attempting to interact with an admin that is not in pending status.
    error AdminNotPending();
    /// @notice Thrown when attempting to add an admin that is already pending.
    error AdminAlreadyPending();
}

interface IPermissionControllerEvents {
    /// @notice Emitted when an appointee is set for an account to handle specific function calls.
    event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector);

    /// @notice Emitted when an appointee's permission to handle function calls for an account is revoked.
    event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector);

    /// @notice Emitted when an address is set as a pending admin for an account, requiring acceptance.
    event PendingAdminAdded(address indexed account, address admin);

    /// @notice Emitted when a pending admin status is removed for an account before acceptance.
    event PendingAdminRemoved(address indexed account, address admin);

    /// @notice Emitted when an address accepts and becomes an active admin for an account.
    event AdminSet(address indexed account, address admin);

    /// @notice Emitted when an admin's permissions are removed from an account.
    event AdminRemoved(address indexed account, address admin);
}

interface IPermissionController is IPermissionControllerErrors, IPermissionControllerEvents, ISemVerMixin {
    /**
     * @notice Sets a pending admin for an account.
     * @param account The account to set the pending admin for.
     * @param admin The address to set as pending admin.
     * @dev The pending admin must accept the role before becoming an active admin.
     * @dev Multiple admins can be set for a single account.
     */
    function addPendingAdmin(address account, address admin) external;

    /**
     * @notice Removes a pending admin from an account before they have accepted the role.
     * @param account The account to remove the pending admin from.
     * @param admin The pending admin address to remove.
     * @dev Only an existing admin of the account can remove a pending admin.
     */
    function removePendingAdmin(address account, address admin) external;

    /**
     * @notice Allows a pending admin to accept their admin role for an account.
     * @param account The account to accept the admin role for.
     * @dev Only addresses that were previously set as pending admins can accept the role.
     */
    function acceptAdmin(
        address account
    ) external;

    /**
     * @notice Removes an active admin from an account.
     * @param account The account to remove the admin from.
     * @param admin The admin address to remove.
     * @dev Only an existing admin of the account can remove another admin.
     * @dev Will revert if removing this admin would leave the account with zero admins.
     */
    function removeAdmin(address account, address admin) external;

    /**
     * @notice Sets an appointee who can call specific functions on behalf of an account.
     * @param account The account to set the appointee for.
     * @param appointee The address to be given permission.
     * @param target The contract address the appointee can interact with.
     * @param selector The function selector the appointee can call.
     * @dev Only an admin of the account can set appointees.
     */
    function setAppointee(address account, address appointee, address target, bytes4 selector) external;

    /**
     * @notice Removes an appointee's permission to call a specific function.
     * @param account The account to remove the appointee from.
     * @param appointee The appointee address to remove.
     * @param target The contract address to remove permissions for.
     * @param selector The function selector to remove permissions for.
     * @dev Only an admin of the account can remove appointees.
     */
    function removeAppointee(address account, address appointee, address target, bytes4 selector) external;

    /**
     * @notice Checks if a given address is an admin of an account.
     * @param account The account to check admin status for.
     * @param caller The address to check.
     * @dev If the account has no admins, returns true only if the caller is the account itself.
     * @return Returns true if the caller is an admin, false otherwise.
     */
    function isAdmin(address account, address caller) external view returns (bool);

    /**
     * @notice Checks if an address is currently a pending admin for an account.
     * @param account The account to check pending admin status for.
     * @param pendingAdmin The address to check.
     * @return Returns true if the address is a pending admin, false otherwise.
     */
    function isPendingAdmin(address account, address pendingAdmin) external view returns (bool);

    /**
     * @notice Retrieves all active admins for an account.
     * @param account The account to get the admins for.
     * @dev If the account has no admins, returns an array containing only the account address.
     * @return An array of admin addresses.
     */
    function getAdmins(
        address account
    ) external view returns (address[] memory);

    /**
     * @notice Retrieves all pending admins for an account.
     * @param account The account to get the pending admins for.
     * @return An array of pending admin addresses.
     */
    function getPendingAdmins(
        address account
    ) external view returns (address[] memory);

    /**
     * @notice Checks if a caller has permission to call a specific function.
     * @param account The account to check permissions for.
     * @param caller The address attempting to make the call.
     * @param target The contract address being called.
     * @param selector The function selector being called.
     * @dev Returns true if the caller is either an admin or an appointed caller.
     * @dev Be mindful that upgrades to the contract may invalidate the appointee's permissions.
     * This is only possible if a function's selector changes (e.g. if a function's parameters are modified).
     * @return Returns true if the caller has permission, false otherwise.
     */
    function canCall(address account, address caller, address target, bytes4 selector) external returns (bool);

    /**
     * @notice Retrieves all permissions granted to an appointee for a given account.
     * @param account The account to check appointee permissions for.
     * @param appointee The appointee address to check.
     * @return Two arrays: target contract addresses and their corresponding function selectors.
     */
    function getAppointeePermissions(
        address account,
        address appointee
    ) external returns (address[] memory, bytes4[] memory);

    /**
     * @notice Retrieves all appointees that can call a specific function for an account.
     * @param account The account to get appointees for.
     * @param target The contract address to check.
     * @param selector The function selector to check.
     * @dev Does not include admins in the returned list, even though they have calling permission.
     * @return An array of appointee addresses.
     */
    function getAppointees(address account, address target, bytes4 selector) external returns (address[] memory);
}
