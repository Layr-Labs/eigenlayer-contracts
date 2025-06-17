# ReleaseManager

| File | Notes |
| -------- | -------- |
| [`ReleaseManager.sol`](../../src/contracts/core/ReleaseManager.sol) |  |
| [`ReleaseManagerStorage.sol`](../../src/contracts/core/ReleaseManagerStorage.sol) | state variables |
| [`IReleaseManager.sol`](../../src/contracts/interfaces/IReleaseManager.sol) | interface |

Libraries and Mixins:

| File | Notes |
| -------- | -------- |
| [`PermissionControllerMixin.sol`](../../src/contracts/mixins/PermissionControllerMixin.sol) | account delegation |
| [`SemVerMixin.sol`](../../src/contracts/mixins/SemVerMixin.sol) | semantic versioning |
| [`OperatorSetLib.sol`](../../src/contracts/libraries/OperatorSetLib.sol) | encode/decode operator sets |

## Overview

The `ReleaseManager` manages software releases for AVS operator sets. It provides a standardized way for AVSs to publish software artifacts (binaries, docker images, etc.) that operators in their operator sets should upgrade to by specified deadlines. This ensures operators run compatible and up-to-date software versions required by the AVS.

The `ReleaseManager's` responsibilities include:

* **Release Publishing**: AVSs can publish new releases containing one or more software artifacts for their operator sets.
* **Upgrade Deadlines**: Each release specifies a deadline by which operators must upgrade.
* **Release Tracking**: Maintains a history of all releases for each operator set.
* **Release Queries**: Provides view functions to query release information.

An AVS in the context of `ReleaseManager` is defined as the `address` of the contract that implements the AVS logic. For `PermissionController` purposes, this AVS address is also the AVS [account](https://github.com/Layr-Labs/eigenlayer-contracts/blob/dev/docs/permissions/PermissionController.md#accounts).

### Important Notes on Release Management

* **Latest Release Validity**: Only the latest release for an operator set is considered valid. Previous releases become obsolete as soon as a new release is published.
* **Upgrade Deadlines**: The `upgradeByTime` timestamp (in Unix time) is a suggested deadline and is not enforced on-chain or off-chain. It serves as a communication mechanism for AVSs to indicate when operators should complete their upgrades.
* **Multiple Releases in Same Block**: If multiple releases are published in the same block with the same `upgradeByTime`, the last transaction processed in that block will determine the latest valid release.

---

## Releases

A release represents a collection of software artifacts that operators in an operator set must run. Each release is associated with a specific operator set and contains:

* **Artifacts**: An array of software artifacts, each with a digest (hash) and registry URL
* **Upgrade By Time**: A Unix timestamp by which operators should complete the upgrade

The release structure is defined as:

```solidity
/**
 * @notice Represents a software artifact with its digest and registry URL.
 * @param digest The hash digest of the artifact.
 * @param registryUrl The URL where the artifact can be found.
 */
struct Artifact {
    bytes32 digest;
    string registryUrl;
}

/**
 * @notice Represents a release containing multiple artifacts and an upgrade deadline.
 * @param artifacts Array of artifacts included in this release.
 * @param upgradeByTime Timestamp by which operators must upgrade to this release.
 */
struct Release {
    Artifact[] artifacts;
    uint32 upgradeByTime;
}
```

**State Management:**

```solidity
/// @dev Mapping from operator set key to array of releases for that operator set
mapping(bytes32 operatorSetKey => Release[]) internal _operatorSetReleases;
```

**Methods:**
* [`publishRelease`](#publishrelease)
* [`getTotalReleases`](#gettotalreleases)
* [`getRelease`](#getrelease)
* [`getLatestRelease`](#getlatestrelease)
* [`getLatestUpgradeByTime`](#getlatestupgradebytime)
* [`isValidRelease`](#isValidRelease)

---

### Write Functions

#### `publishRelease`

```solidity
/**
 * @notice Publishes a new release for an operator set.
 * @param operatorSet The operator set this release is for.
 * @param release The release that was published.
 * @return releaseId The index of the newly published release.
 */
function publishRelease(
    OperatorSet calldata operatorSet,
    Release calldata release
) 
    external 
    checkCanCall(operatorSet.avs) 
    returns (uint256 releaseId)
```

_Note: this method can be called directly by an AVS, or by a caller authorized by the AVS. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

AVSs use this method to publish new software releases for their operator sets. Each release contains one or more artifacts that represent the software components operators must run (e.g., validator clients, network monitors, etc.). The AVS specifies a deadline (`upgradeByTime`) by which all operators in the operator set must upgrade to the new release.

The `releaseId` returned is the zero-based index of the release in the operator set's release array. This ID can be used to query the release later using [`getRelease`](#getrelease).

*Effects*:
* Appends the release to `_operatorSetReleases[operatorSet.key()]`
* The new release is assigned a `releaseId` equal to the previous array length
* All artifact information (digest and registryUrl) is copied to storage
* Emits a `ReleasePublished` event with the operator set, release ID, and release details

*Requirements*:
* Caller MUST be authorized, either as the AVS itself or an admin/appointee (see [`PermissionController.md`](../permissions/PermissionController.md))
* `release.upgradeByTime` MUST be greater than or equal to the current block timestamp
---

### View Functions

#### `getTotalReleases`

```solidity
/**
 * @notice Returns the total number of releases for an operator set.
 * @param operatorSet The operator set to query.
 * @return The number of releases.
 */
function getTotalReleases(
    OperatorSet memory operatorSet
) 
    public 
    view 
    returns (uint256)
```

Returns the total number of releases that have been published for the specified operator set. This can be used to iterate through all releases or to determine if any releases exist.

*Returns*:
* The length of the releases array for the given operator set
* Returns 0 if no releases have been published

#### `getRelease`

```solidity
/**
 * @notice Returns a specific release by index.
 * @param operatorSet The operator set to query.
 * @param releaseId The id of the release to get.
 * @return The release at the specified index.
 */
function getRelease(
    OperatorSet memory operatorSet, 
    uint256 releaseId
) 
    external 
    view 
    returns (Release memory)
```

Retrieves a specific release by its ID for a given operator set. The `releaseId` is the zero-based index of the release in the operator set's release history.

*Returns*:
* The complete `Release` struct including all artifacts and the upgrade deadline
* Reverts if `releaseId` is out of bounds

#### `getLatestRelease`

```solidity
/**
 * @notice Returns the latest release for an operator set.
 * @param operatorSet The operator set to query.
 * @return The id of the latest release.
 * @return The latest release.
 */
function getLatestRelease(
    OperatorSet memory operatorSet
) 
    public 
    view 
    returns (uint256, Release memory)
```

Retrieves the most recently published release for an operator set. This is typically the release that operators should be running or upgrading to.

*Returns*:
* The latest `Release` struct from the operator set's release array
* Reverts if no releases have been published for the operator set

#### `getLatestUpgradeByTime`

```solidity
/**
 * @notice Returns the upgrade by time for the latest release.
 * @param operatorSet The operator set to query.
 * @return The upgrade by time for the latest release.
 */
function getLatestUpgradeByTime(
    OperatorSet memory operatorSet
) 
    external 
    view 
    returns (uint256)
```

A convenience function that returns just the upgrade deadline from the latest release. This can be useful for quickly checking when operators must complete their upgrades.

*Returns*:
* The `upgradeByTime` timestamp from the latest release
* Reverts if no releases have been published for the operator set

#### `isValidRelease`

```solidity
/**
 * @notice Returns true if the release is the latest release, false otherwise.
 * @param operatorSet The operator set to query.
 * @param releaseId The id of the release to check.
 * @return True if the release is the latest release, false otherwise.
 */
function isValidRelease(
    OperatorSet memory operatorSet, 
    uint256 releaseId
) 
    external 
    view 
    returns (bool)
```

Checks whether a given release ID corresponds to the latest release for an operator set. This can be useful for operators to verify they are running the most current software version.

**Note**: Only the latest release is considered valid. All previous releases are considered obsolete and should not be used by operators.

*Returns*:
* `true` if the `releaseId` matches the latest release index
* `false` if the `releaseId` refers to an older release
* Reverts if the operator set has no releases

---
