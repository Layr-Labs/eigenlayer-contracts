// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";

import "../interfaces/IPermissionController.sol";

abstract contract PermissionControllerStorage is IPermissionController {
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.AddressSet;
    using EnumerableMap for EnumerableMap.AddressToUintMap;
    using EnumerableMap for EnumerableMap.Bytes32ToUintMap;

    struct AccountPermissions {
        /// @notice The pending admins of the account
        EnumerableSet.AddressSet pendingAdmins;
        /// @notice The admins of the account
        EnumerableSet.AddressSet admins;
        /// @notice Mapping from an appointee to the list of encoded target & selectors
        mapping(address appointee => EnumerableSet.Bytes32Set) appointeePermissions;
        /// @notice Mapping from encoded target & selector to the list of appointees
        mapping(bytes32 targetSelector => EnumerableSet.AddressSet) permissionAppointees;
        /// @notice Mapping from target & selector to the effectBlock of a delayed target selector
        mapping(bytes32 targetSelector => EnumerableMap.AddressToUintMap) appointeeEffectBlock;
    }

    /// @notice Mapping from an account to its permission
    mapping(address account => AccountPermissions) internal _permissions;

    /// @notice Map of delayed selectors to the their delay
    EnumerableMap.Bytes32ToUintMap internal _permissionEffectBlock;

    /// @notice Map of target & selector to whether the admin cannot call it
    mapping(bytes32 targetSelector => bool) internal _adminCannotCall;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[46] private __gap;
}
