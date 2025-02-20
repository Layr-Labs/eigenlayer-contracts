# PermissionController

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`PermissionController.sol`](../../src/contracts/permissions/PermissionController.sol) | Singleton | Transparent proxy |

The `PermissionController` handles user permissions for protocol contracts which explicitly integrate it. Note that "users" in the context of the `PermissionController` refers to **AVSs** and **operators**; it does *not* refer to **stakers**.

The `PermissionController` is integrated into other core contracts, enabling (for specific methods) AVSs and operators to designate _other accounts_ ("appointees") that can call these methods on their behalf. The core contracts using the `PermissionController` as a dependency are the:
* `DelegationManager`
* `AllocationManager`
* `RewardsCoordinator`

The `PermissionController` defines three different roles:
* [Accounts](#accounts)
* [Admins](#admins)
* [Appointees](#appointees)

---

## Accounts

**Accounts** refer to the Ethereum address through which one interacts with the protocol _if no appointees are set_. From the core contracts' perspective, accounts are the "state holder," i.e. the address referenced in storage when a contract method interacts with state. For example, in the `DelegationManager`, the `operator` address that holds shares in the `operatorShares` mapping is an "account." In the `AllocationManager`, an AVS's "account" is the address under which operator sets are created.

The `PermissionController` allows an account to designate **admins** and/or **appointees** to take certain actions on its behalf. Note that setting up admins/appointees is _optional_, and carries with it a significant responsibility to **ensure the designated actors are intentionally being granted authority**.

Both admins AND appointees can be granted authority to act on an account's behalf. Admins are granted full reign over any `PermissionController`-enabled functions, while appointees must be granted authority to call specific functions on specific contracts. The list of methods that are `PermissionController`-enabled follow.

For operators:
* `AllocationManager.modifyAllocations`
* `AllocationManager.registerForOperatorSets`
* `AllocationManager.deregisterFromOperatorSets`
* `AllocationManager.setAllocationDelay`
* `DelegationManager.modifyOperatorDetails`
* `DelegationManager.updateOperatorMetadataURI`
* `DelegationManager.undelegate`
* `RewardsCoordinator.setClaimerFor`
* `RewardsCoordinator.setClaimerFor`
* `RewardsCoordinator.setOperatorAVSSplit`
* `RewardsCoordinator.setOperatorPISplit`

For AVSs:
* `AllocationManager.slashOperator`
* `AllocationManager.deregisterFromOperatorSets`
* `AllocationManager.setAVSRegistrar`
* `AllocationManager.updateAVSMetadataURI`
* `AllocationManager.createOperatorSets`
* `AllocationManager.addStrategiesToOperatorSet`
* `AllocationManager.removeStrategiesFromOperatorSet`
* `RewardsCoordinator.createOperatorDirectedAVSRewardsSubmission`
* `RewardsCoordinator.setClaimerFor`

### Account Permissions

Account permissions are stored within a struct defined as follows:

```solidity
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
```

These structs are then stored within a mapping defined as follows, allowing for fetching account permissions for a given account with ease:

```solidity
mapping(address account => AccountPermissions) internal _permissions;
```

By default, no other address can perform an action on behalf of a given account. However, accounts can add admins and/or appointees to give other addresses the ability to act on their behalf.

## Admins

Admins are able to take ANY action on behalf of an original account -- including adding or removing admins. This enables operations like key rotation for operators, or creating a backup admin which is stored on a cold key.

**Note:** by default, an account is its own admin. However, once an admin is added, this is no longer the case; only the admins listed in `_permissions.admins` are admins. If an account wants to both add admins AND continue acting as its own admin, _it must be added to the admins list_.

### Adding an Admin

The relevant functions for admin addition are:

* [`addPendingAdmin`](#addpendingadmin)
* [`removePendingAdmin`](#removependingadmin)
* [`acceptAdmin`](#acceptadmin)

#### `addPendingAdmin`

```solidity
/**
 * @notice Sets a pending admin of an account
 * @param account to set pending admin for
 * @param admin to set
 * @dev Multiple admins can be set for an account
 */
function addPendingAdmin(address account, address admin) external onlyAdmin(account);
```

When adding a new admin, an account or admin must first call `addPendingAdmin()`. Then, the pending admin must call `acceptAdmin()` to complete the process. An account cannot force an admin role upon another account.

Pending admins do not have any particular authority, but are granted the full authority of an admin once they call `acceptAdmin()`.

*Effects*:
* An address is added to the `pendingAdmins` set for the account
* A `PendingAdminAdded` event is emitted specifying the account for which a pending admin was added

*Requirements*:
* The proposed admin MUST NOT already be an admin for the `account`
* The proposed admin MUST NOT be a pending admin for the `account`
* Caller MUST be an admin for the `account`, or the `account` itself if no admin is set

#### `removePendingAdmin`

```solidity
/**
 * @notice Removes a pending admin of an account
 * @param account to remove pending admin for
 * @param admin to remove
 * @dev Only the admin of the account can remove a pending admin
 */
function removePendingAdmin(address account, address admin) external onlyAdmin(account);
```

An account or admin can call `removePendingAdmin()` to prevent a pending admin from accepting their role. However, this will only work if the pending admin has not already called `acceptAdmin()`. If this occurs, an admin can call `removeAdmin` to remove the unwanted admin.

*Effects*:
* An address is removed from the `pendingAdmins` set for the account
* A `PendingAdminRemoved` event is emitted specifying the account for which a pending admin was removed

*Requirements*:
* The proposed admin MUST be a pending admin for the account
* Caller MUST be an admin for the account, or the account's address itself if no admin is set

#### `acceptAdmin`

```solidity
/**
 * @notice Accepts the admin role of an account
 * @param account to accept admin for
 * @dev Only a pending admin for the account can become an admin
 */
function acceptAdmin(address account) external;
```

Called by a pending admin to claim the admin role for an account. The caller must have been previously added as a pending admin.

Note that once an account has successfully added an admin (i.e. the pending admin has called `acceptAdmin()`), **the account's address itself no longer has its default admin privileges**. This behavior benefits accounts seeking to perform a *key rotation*, as adding an admin allows them to remove permissions from their original, potentially compromised, key. If an account wants to retain admin privileges for its own address, it is recommended to first add itself as an admin, then add any other admins as desired.

*Effects*:
* The caller is removed from the `pendingAdmins` set for the account
* The caller is added to the `admins` set for the account
* A `AdminSet` event is emitted specifying the account for which an admin was added

*Requirements*:
* Caller MUST be a pending admin for the account

### Removing an Admin

#### `removeAdmin`

```solidity
/**
 * @notice Remove an admin of an account
 * @param account to remove admin for
 * @param admin to remove
 * @dev Only the admin of the account can remove an admin
 * @dev Reverts when an admin is removed such that no admins are remaining
 */
function removeAdmin(address account, address admin) external onlyAdmin(account);
```

An admin of an account can call `removeAdmin()` to remove any other admins of the same account. However, one admin must always remain for any given account. In other words, once an account has added an admin, it must always have at least one admin in perpetuity.

*Effects*:
* The specified admin is removed from the `admins` set for the account
* An `AdminRemoved` event is emitted specifying the accuont for which an admin was removed

*Requirements*:
* `admins.length()` MUST be greater than 1, such that removing the admin does not remove all admins for the account
* The address to remove MUST be an admin for the account
* Caller MUST be an admin for the account, or the account's address itself if no admin is set

## Appointees

Appointees are able to act as another account *for a specific function for a specific contract*, granting accounts granular access control.

Specifically, an account (or its admins) can grant an appointee access to a specific `selector` (i.e [function](https://solidity-by-example.org/function-selector/)) on a given `target` (i.e. contract). The `target` and `selector` are combined in the form of the `targetSelector` and serve to uniquely identify a permissioned function on a specific contract.

Appointees can be granted access to multiple functions/contracts. Each new `targetSelector` permission granted requires setting the appointee from scratch, and revoking the appointee's permission requires revoking each individual `targetSelector` permission, as described below.

### Adding an Appointee

#### `setAppointee`

```solidity
/**
 * @notice Set an appointee for a given account
 * @param account to set appointee for
 * @param appointee to set
 * @param target to set appointee for
 * @param selector to set appointee for
 * @dev Only the admin of the account can set an appointee
 */
function setAppointee(
    address account,
    address appointee,
    address target,
    bytes4 selector
) external onlyAdmin(account);
```

An account (or its admins) can call `setAppointee()` to give another address the ability to call a specific function on a given contract. That address is then only able to call that specific function on that specific contract on behalf of `account`.

Note that unlike the process to become an admin, there is no requirement for the `appointee` to accept the appointment.

*Effects*:
* The `targetSelector` is added to the specified `appointee` set within the  `appointeePermissions` mapping
* The `appointee` is added to the specified `targetSelector` set within the  `permissionAppointees` mapping
* The `AppointeeSet` event is emitted, specifying the account, appointee, target contract, and function selector

*Requirements*:
* Caller MUST be an admin for the account, or the account's address itself if no admin is set
* The proposed appointee MUST NOT already have permissions for the given `targetSelector`

### Removing an Appointee

#### `removeAppointee`

```solidity
/**
 * Removes an appointee for a given account
 * @param account to remove appointee for
 * @param appointee to remove
 * @param target to remove appointee for
 * @param selector to remove appointee for
 * @dev Only the admin of the account can remove an appointee
 */
function removeAppointee(
    address account,
    address appointee,
    address target,
    bytes4 selector
) external onlyAdmin(account);
```

An account (or its admins) can call `removeAppointee()` to remove an `appointee's` permissions for a given contract/function pair. Note that there does not exist any way currently to atomically remove all permissions for a given appointee, or all appointees for a given function selector - each permission must be revoked individually.

Also note that permissions to specific functions/contracts cannot be revoked for _admins_. Admins always have full access, unless another admin removes them from the admin list.

*Effects*:
* The `targetSelector` is removed from the specified `appointee` set within the  `appointeePermissions` mapping
* The `appointee` is removed from the specified `targetSelector` set within the  `permissionAppointees` mapping
* The `AppointeeRemoved` event is emitted, specifying the account, appointee, target contract, and function selector

*Requirements*:
* Caller MUST be an admin for the account, or the account's address itself if no admin is set
* The proposed appointee MUST already have permissions for the given `targetSelector`