// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import "../interfaces/IPermissionController.sol";

abstract contract PermissionControllerStorage is IPermissionController {
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.AddressSet;

    struct AccountPermissions {
        /// @notice The admins of the account
        EnumerableSet.AddressSet admins;
        /// @notice Mapping from a delegate to the list of encoded target & selectors
        mapping(address delegate => EnumerableSet.Bytes32Set) delegatePermissions;
        /// @notice Mapping from encoded target & selector to the list of delegates
        mapping(bytes32 targetSelector => EnumerableSet.AddressSet) permissionDelegates;
    }

    /// @notice Mapping from an account to its permission
    mapping(address account => AccountPermissions) internal _permissions;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}
