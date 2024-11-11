// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

interface IPermissionControllerErrors {
    /// @notice Thrown when the caller is not the admin
    error NotAdmin();
    /// @notice Thrown when an admin is set to the zero address
    error AdminAlreadySet();
    /// @notice Thrown when the admin to remove is not an admin
    error AdminNotSet();
    /// @notice Thrown when a delegate is already set for the account's function
    error DelegateAlreadySet();
    /// @notice Thrown when a delegate is not set for the account's function
    error DelegateNotSet();
}

interface IPermissionControllerEvents {
    /// @notice Emitted when a delegate is set
    event DelegateSet(address indexed account, address indexed delegate, address target, bytes4 selector);

    /// @notice Emitted when a delegate is revoked
    event DelegateRemoved(address indexed account, address indexed delegate, address target, bytes4 selector);

    /// @notice Emitted when an admin is set for a given account
    event AdminSet(address indexed account, address admin);

    /// @notice Emitted when an admin is removed for a given account
    event AdminRemoved(address indexed account, address admin);
}

interface IPermissionController is IPermissionControllerErrors, IPermissionControllerEvents {
    /**
     * @notice Set the admin of an account
     * @param account to set admin for
     * @param admin to set
     * @dev Multiple admins can be set for an account
     */
    function setAdmin(address account, address admin) external;

    /**
     * @notice Remove an admin of an account
     * @param account to remove admin for
     * @param admin to remove
     * @dev Only the admin of the account can remove an admin
     */
    function removeAdmin(address account, address admin) external;

    /**
     * @notice Set a delegate for a given account
     * @param account to set delegate for
     * @param delegate to set
     * @param target to set delegate for
     * @param selector to set delegate for
     * @dev Only the admin of the account can set a delegate
     */
    function setDelegate(address account, address delegate, address target, bytes4 selector) external;

    /**
     * Removes a delegate for a given account
     * @param account to remove delegate for
     * @param delegate to remove
     * @param target to remove delegate for
     * @param selector to remove delegate for
     * @dev Only the admin of the account can remove a delegate
     * @dev If all admins are removed, the original account is now the admin
     */
    function removeDelegate(address account, address delegate, address target, bytes4 selector) external;

    /**
     * @notice Checks if the given caller is an admin of the account
     * @dev If the account has no admin, the caller is checked to be the account itself
     */
    function isAdmin(address account, address caller) external view returns (bool);

    /**
     * @notice Get the admins of an account
     * @param account The account to get the admin of
     * @dev If the account has no admin, the account itself is returned
     */
    function getAdmins(
        address account
    ) external view returns (address[] memory);

    /**
     * @notice Checks if the given caller has permissions to call the fucntion
     * @param account to check
     * @param caller to check permission for
     * @param target to check permission for
     * @param selector to check permission for
     * @dev Returns `true` if the admin OR the delegate is the caller
     */
    function canCall(address account, address caller, address target, bytes4 selector) external returns (bool);

    /**
     * @notice Gets the list of permissions of a delegate for a given account
     * @param account to get delegate permissions for
     * @param delegate to get permissions
     */
    function getDelegatePermissions(
        address account,
        address delegate
    ) external returns (address[] memory, bytes4[] memory);

    /**
     * @notice Returns the list of delegates for a given account and function
     * @param account to get delegates for
     * @param target to get delegates for
     * @param selector to get delegates for
     * @dev Does NOT include admin as a delegate, even though it can call
     */
    function getDelegates(address account, address target, bytes4 selector) external returns (address[] memory);
}
