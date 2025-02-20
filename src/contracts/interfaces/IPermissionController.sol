// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

interface IPermissionControllerErrors {
    /// @notice Thrown when the caller is not the admin
    error NotAdmin();
    /// @notice Thrown when the admin to remove is not an admin
    error AdminNotSet();
    /// @notice Thrown when an appointee is already set for the account's function
    error AppointeeAlreadySet();
    /// @notice Thrown when an appointee is not set for the account's function
    error AppointeeNotSet();
    /// @notice Thrown when the account attempts to remove the only admin
    error CannotHaveZeroAdmins();
    /// @notice Thrown when an admin is already set
    error AdminAlreadySet();
    /// @notice Thrown when an admin is not pending
    error AdminNotPending();
    /// @notice Thrown when an admin is already pending
    error AdminAlreadyPending();
}

interface IPermissionControllerEvents {
    /// @notice Emitted when an appointee is set
    event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector);

    /// @notice Emitted when an appointee is revoked
    event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector);

    /// @notice Emitted when an admin is set as pending for an account
    event PendingAdminAdded(address indexed account, address admin);

    /// @notice Emitted when an admin is removed as pending for an account
    event PendingAdminRemoved(address indexed account, address admin);

    /// @notice Emitted when an admin is set for a given account
    event AdminSet(address indexed account, address admin);

    /// @notice Emitted when an admin is removed for a given account
    event AdminRemoved(address indexed account, address admin);
}

interface IPermissionController is IPermissionControllerErrors, IPermissionControllerEvents {
    /**
     * @notice Sets a pending admin of an account
     * @param account to set pending admin for
     * @param admin to set
     * @dev Multiple admins can be set for an account
     */
    function addPendingAdmin(address account, address admin) external;

    /**
     * @notice Removes a pending admin of an account
     * @param account to remove pending admin for
     * @param admin to remove
     * @dev Only the admin of the account can remove a pending admin
     */
    function removePendingAdmin(address account, address admin) external;

    /**
     * @notice Accepts the admin role of an account
     * @param account to accept admin for
     * @dev Only a pending admin for the account can become an admin
     */
    function acceptAdmin(
        address account
    ) external;

    /**
     * @notice Remove an admin of an account
     * @param account to remove admin for
     * @param admin to remove
     * @dev Only the admin of the account can remove an admin
     * @dev Reverts when an admin is removed such that no admins are remaining
     */
    function removeAdmin(address account, address admin) external;

    /**
     * @notice Set an appointee for a given account
     * @param account to set appointee for
     * @param appointee to set
     * @param target to set appointee for
     * @param selector to set appointee for
     * @dev Only the admin of the account can set an appointee
     */
    function setAppointee(address account, address appointee, address target, bytes4 selector) external;

    /**
     * Removes an appointee for a given account
     * @param account to remove appointee for
     * @param appointee to remove
     * @param target to remove appointee for
     * @param selector to remove appointee for
     * @dev Only the admin of the account can remove an appointee
     */
    function removeAppointee(address account, address appointee, address target, bytes4 selector) external;

    /**
     * @notice Checks if the given caller is an admin of the account
     * @dev If the account has no admin, the caller is checked to be the account itself
     */
    function isAdmin(address account, address caller) external view returns (bool);

    /**
     * @notice Checks if the `pendingAdmin` is a pending admin of the `account`
     */
    function isPendingAdmin(address account, address pendingAdmin) external view returns (bool);

    /**
     * @notice Get the admins of an account
     * @param account The account to get the admin of
     * @dev If the account has no admin, the account itself is returned
     */
    function getAdmins(
        address account
    ) external view returns (address[] memory);

    /**
     * @notice Get the pending admins of an account
     * @param account The account to get the pending admin of
     */
    function getPendingAdmins(
        address account
    ) external view returns (address[] memory);

    /**
     * @notice Checks if the given caller has permissions to call the fucntion
     * @param account to check
     * @param caller to check permission for
     * @param target to check permission for
     * @param selector to check permission for
     * @dev Returns `true` if the admin OR the appointee is the caller
     */
    function canCall(address account, address caller, address target, bytes4 selector) external returns (bool);

    /**
     * @notice Gets the list of permissions of an appointee for a given account
     * @param account to get appointee permissions for
     * @param appointee to get permissions
     */
    function getAppointeePermissions(
        address account,
        address appointee
    ) external returns (address[] memory, bytes4[] memory);

    /**
     * @notice Returns the list of appointees for a given account and function
     * @param account to get appointees for
     * @param target to get appointees for
     * @param selector to get appointees for
     * @dev Does NOT include admin as an appointee, even though it can call
     */
    function getAppointees(address account, address target, bytes4 selector) external returns (address[] memory);
}
