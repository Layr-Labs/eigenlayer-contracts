// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import "../interfaces/IPermissionController.sol";

abstract contract PermissionControllerStorage is IPermissionController {
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.AddressSet;

    bytes32 public constant PERMIT_APPOINTEE_TYPEHASH = keccak256(
        "PermitAppointee(address appointee,address target,bytes4 selector,bool permitted,bytes32 salt,uint32 expiry)"
    );

    struct AccountPermissions {
        /// @notice The pending admins of the account
        EnumerableSet.AddressSet pendingAdmins;
        /// @notice The admins of the account
        EnumerableSet.AddressSet admins;
        /// @notice Mapping from an appointee to the list of encoded target & selectors
        mapping(address appointee => EnumerableSet.Bytes32Set) appointeePermissions;
        /// @notice Mapping from encoded target & selector to the list of appointees
        mapping(bytes32 targetSelector => EnumerableSet.AddressSet) permissionAppointees;
    }

    /// @notice Mapping from an account to its permission
    mapping(address account => AccountPermissions) internal _permissions;

    /// @notice Mapping from an account to its nonce
    mapping(address account => mapping(bytes32 => bool)) public isSaltSpent;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}
