// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

interface IOperatorControllerEvents {
    event DelegateSet(address indexed delegate, address indexed target, bytes4 indexed selector);

    event DelegateRevoked(address indexed delegate, address indexed target, bytes4 indexed selector);
}

interface IOperatorController is IOperatorControllerEvents {
    /// @notice Set a delegated role for a target contract
    function setDelegatedRole(address delegate, address target, bytes4 selector) external;

    /// @notice Revoke a delegated role for a target contract
    function revokeDelegatedRole(address delegate, address target, bytes4 selector) external;

    /// @notice Checks if a delegate has permission to call a function
    function hasPermission(address delegate, address target, bytes4 selector) external returns (bool);

    /// @notice Returns the list of permissions of a delegate
    function getDelegatePermissions(
        address delegate
    ) external returns (address[] memory, bytes4[] memory);

    /// @notice Returns the list of delegates of a given (target, selector)
    function getDelegates(address target, bytes4 selector) external returns (address[] memory);
}
