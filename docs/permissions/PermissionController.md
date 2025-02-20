# PermissionController

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`PermissionController.sol`](../../src/contracts/permissions/PermissionController.sol) | Singleton | Transparent proxy |

The `PermissionController` handles user permissions for protocol contracts which explicitly integrate it. Note that "users" in the context of the `PermissionController` refers to **AVSs** and **Operators**; it does *not* refer to **Stakers**.

The `PermissionController` defines three different roles:
* [Accounts](#accounts)
* [Admins](#admins)
* [Appointees](#appointees)

The core contracts using the `PermissionController` as a dependency are the:
* `DelegationManager`
* `AllocationManager`
* `RewardsCoordinator`

Note that the `AVSDirectory` will soon be deprecated and thus does not support this contract.

---

## Accounts

**Accounts** refer to the Ethereum address through which one interacts with the protocol. Accounts have the ability to set and remove **admins** and/or **appointees**, which can take actions on an account's behalf depending on their role.

Note that an account cannot appoint an address to *register* on its behalf. In other words, the address a user registers with is their account, and they can then transfer admin of that account in a subsequent tx. Similarly, the account of an AVS is the address they use to initialize state in the `AllocationManager`.

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

Admins are able to take any action on behalf of an original account -- including adding or removing admins. This enables operations like key rotation for operators, or creating a backup admin which is stored on a cold key.

### Adding an Admin

The relevant functions for admin addition are:

* [`addPendingAdmin`](#addpendingadmin)
* [`removePendingAdmin`](#removependingadmin)
* [`acceptAdmin`](#acceptadmin)

#### `addPendingAdmin`

```solidity
function addPendingAdmin(address account, address admin) external onlyAdmin(account);
```

When adding an admin, an account must call `addPendingAdmin()`. Then, the pending admin must call `acceptAdmin()` for the given account, completing the process. An account cannot force an admin role upon another account.

*Effects*:
* An address is added to the `pendingAdmins` set for the account
* A `PendingAdminAdded` event is emitted specifying the account for which a pending admin was added

*Requirements*:
* The proposed admin MUST NOT already be an admin for the account
* The proposed admin MUST NOT be a pending admin for the account
* Caller MUST be an admin for the account, or the account's address itself if no admin is set

#### `removePendingAdmin`

```solidity
function removePendingAdmin(address account, address admin) external onlyAdmin(account);
```

An account, if it wishes to retract the pending admin offer, can call `removePendingAdmin()` to prevent the potential admin from accepting their role. However, this will only work if `acceptAdmin()` has not yet been called by the potential admin.

*Effects*:
* An address is removed from the `pendingAdmins` set for the account
* A `PendingAdminRemoved` event is emitted specifying the account for which a pending admin was removed

*Requirements*:
* The proposed admin MUST be a pending admin for the account
* Caller MUST be an admin for the account, or the account's address itself if no admin is set

#### `acceptAdmin`

```solidity
function acceptAdmin(address account) external;
```

To claim the admin role for an account, the pending admin must call `acceptAdmin()` passing in the relevant account.

Note that an account has successfully added an admin (i.e. the pending admin has called `acceptAdmin()`), **the account's address itself no longer has admin privileges**. This behavior benefits operators seeking to perform a *key rotation*, as adding an admin allows them to remove permissions from their original, potentially compromised, key. If an account wants to retain admin privileges for its own address, it is recommended to first add itself as an admin, then add any other admins as desired.

*Effects*:
* The caller is removed from the `pendingAdmins` set for the account
* The caller is added to the `admins` set for the account
* A `AdminSet` event is emitted specifying the account for which an admin was added

*Requirements*:
* Caller MUST be a pending admin for the account

### Removing an Admin

#### `removeAdmin`

```solidity
function removeAdmin(address account, address admin) external onlyAdmin(account);
```

An admin of an account can call `removeAdmin()` to remove any other admins of the same account. However, one admin must always remain for any given account. In other words, once an account has added an admin, it must always have at least one admin in perpetuity.

Note that once an admin has been added to an account, at least one must remain on the account.

*Effects*:
* The specified admin is removed from the `admins` set for the account
* An `AdminRemoved` event is emitted specifying the accuont for which an admin was removed

*Requirements*:
* `admins.length()` MUST be greater than 1, such that removing the admin does not remove all admins for the account
* The address to remove MUST be an admin for the account
* Caller MUST be an admin for the account, or the account's address itself if no admin is set

## Appointees

Appointees are able to act as another account *for a specific function for a specific contract*, granting accounts granular access control.

Specifically, an account (or its admins) can grant an appointee access to a specific `selector` (i.e [function](https://solidity-by-example.org/function-selector/)) on a given `target` (i.e. contract). The `target` and `selector` are combined in the form of the `targetSelector` and serve to uniquely identify a permissioned function.

Appointees can be appointed more than once so that they can access additional functions on a given contract and/or additional functions on *other* contracts. Each new `targetSelector` permission granted requires setting the appointee from scratch, and revoking the appointee's permission requires revoking each individual `targetSelector` permission, as described below.

### Adding an Appointee

#### `setAppointee`

```solidity
function setAppointee(
    address account,
    address appointee,
    address target,
    bytes4 selector
) external onlyAdmin(account);
```

An account (or its admins) can call `setAppointee()` to give another address the ability to call a specific function on a given contract. That address is then only able to call that specific function on that specific contract.

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
function removeAppointee(
    address account,
    address appointee,
    address target,
    bytes4 selector
) external onlyAdmin(account);
```

Similarly, an account (or its admins) can call `removeAppointee()` to remove that address's permissions for a given contract/function pair. Note that there does not exist any way currently to remove all permissions for a given appointee, or all appointees for a given function selector.

*Effects*:
* The `targetSelector` is removed from the specified `appointee` set within the  `appointeePermissions` mapping
* The `appointee` is removed from the specified `targetSelector` set within the  `permissionAppointees` mapping
* The `AppointeeRemoved` event is emitted, specifying the account, appointee, target contract, and function selector

*Requirements*:
* Caller MUST be an admin for the account, or the account's address itself if no admin is set
* The proposed appointee MUST already have permissions for the given `targetSelector`

---

## Functions using the PermissionController

Below are the functions for which the PermissionController can be used to specify access controls.

### DelegationManager

* [`modifyOperatorDetails`](../core/DelegationManager.md#modifyOperatorDetails)
* [`updateOperatorMetadataURI`](../core/DelegationManager.md#updateOperatorMetadataURI)
* [`undelegate`](../core/DelegationManager.md#undelegate)

### AllocationManager

* [`slashOperator`](../core/AllocationManager.md#slashOperator)
* [`modifyAllocations`](../core/AllocationManager.md#modifyAllocations)
* [`registerForOperatorSets`](../core/AllocationManager.md#registerForOperatorSets)
* [`deregisterFromOperatorSets`](../core/AllocationManager.md#deregisterFromOperatorSets)
* [`setAllocationDelay`](../core/AllocationManager.md#setAllocationDelay)
* [`setAVSRegistrar`](../core/AllocationManager.md#setAVSRegistrar)
* [`updateAVSMetadataURI`](../core/AllocationManager.md#updateAVSMetadataURI)
* [`createOperatorSets`](../core/AllocationManager.md#createOperatorSets)
* [`addStrategiesToOperatorSet`](../core/AllocationManager.md#addStrategiesToOperatorSet)
* [`removeStrategiesFromOperatorSet`](../core/AllocationManager.md#removeStrategiesFromOperatorSet)

### RewardsCoordinator

* [`createAVSRewardsSubmission`](../core/RewardsCoordinator.md#createAVSRewardsSubmission)
* [`setClaimerFor`](../core/RewardsCoordinator.md#setClaimerFor)
