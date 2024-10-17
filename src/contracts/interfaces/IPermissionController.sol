// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

interface IPermissionControllerErrors {
    error CannotSetAdmin();
}

interface IPermissionControllerEvents {
    /// @notice Emitted when an admin is set
    event AdminSet(address indexed account, address oldAdmin, address newAdmin);
}

interface IPermissionController is IPermissionControllerErrors, IPermissionControllerEvents {
    /**
     * @notice Sets the admin for a given account
     * @param account the account to set the admin for
     * @param newAdmin the admin to set for the account
     * @dev Only callable by the contract admin OR the account itself if the admin is not set
     */
    function setAdmin(address account, address newAdmin) external;

    /**
     * @notice Checks if a given admin can call a function on behalf of an account
     * @param account The account to check
     * @param caller The address to check if it can call on behalf of the account 
     * @return true if the admin can call a function on behalf of the account, false otherwise
     */
    function canCall(address account, address caller) external returns (bool);
}