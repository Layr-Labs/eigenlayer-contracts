# ProtocolRegistry

| File | Notes |
| --- | --- |
| [`ProtocolRegistry.sol`](../../src/contracts/core/ProtocolRegistry.sol) | core logic |
| [`ProtocolRegistryStorage.sol`](../../src/contracts/core/storage/ProtocolRegistryStorage.sol) | state variables |
| [`IProtocolRegistry.sol`](../../src/contracts/interfaces/IProtocolRegistry.sol) | interface, events, and types |

Libraries and Mixins:

| File | Notes |
| --- | --- |
| [`AccessControlEnumerableUpgradeable`](https://docs.openzeppelin.com/contracts/5.x/api/access#AccessControlEnumerableUpgradeable) | role-based access control |
| [`Initializable`](https://docs.openzeppelin.com/contracts/5.x/api/proxy#Initializable) | protects initializer |
| [`ShortStringsUpgradeable`](https://docs.openzeppelin.com/contracts/5.x/api/utils#ShortStringsUpgradeable) | compact semantic version storage |
| [`EnumerableMap`](https://docs.openzeppelin.com/contracts/5.x/api/utils#EnumerableMap) | iterable mapping for deployment catalog |
| [`IPausable`](../../src/contracts/interfaces/IPausable.sol) | pause hook invoked during emergencies |

## Overview

`ProtocolRegistry` is the canonical catalog of EigenLayer protocol deployments. It maps human-readable deployment names to contract addresses, tracks per-contract configuration flags, emits semantic-version updates every time a new protocol shipment occurs, and exposes functionality to pause the the entire core protocol. This contract is deployed on all EigenLayer source and destination chains.

### Roles and Permissions

* `DEFAULT_ADMIN_ROLE`: Full control, required for `initialize`, `ship`, and `configure`.
* `PAUSER_ROLE`: Defined as `hex"01"` in storage. Addresses with this role can invoke `pauseAll()`.

### Deployment Config Structure

```solidity
struct DeploymentConfig {
    bool pausable;   // deployment supports IPausable.pauseAll()
    bool deprecated; // deployment should no longer be interacted with
}
```

* `pausable` gates whether `pauseAll()` targets the deployment.
* `deprecated` prevents new pauses from being attempted against sunset contracts.

---

## Write Functions

### `initialize`

```solidity
function initialize(address initialAdmin, address pauserMultisig) external initializer
```

Initializes the proxy once by granting:

* `DEFAULT_ADMIN_ROLE` to `initialAdmin`.
* `PAUSER_ROLE` to `pauserMultisig`.

The constructor disables further initializers; therefore `initialize` must be called exactly once after deployment. Upon deployment the `executorMultisig` will be the default admin and the `pauserMultisig` will hold the `PAUSER_ROLE`.

### `ship`

```solidity
function ship(
    address[] calldata addresses,
    DeploymentConfig[] calldata configs,
    string[] calldata names,
    string calldata semanticVersion
) external onlyRole(DEFAULT_ADMIN_ROLE)
```

Ships a new semantic version and batch-registers deployments:

*Effects*:
* Calls `_updateSemanticVersion(semanticVersion)` and emits `SemanticVersionUpdated`.
* For each address to ship:
    * Calls `_appendDeployment`, which stores the name→address mapping, records the config, and emits `DeploymentShipped`.

*Requirements*:
* Caller must hold `DEFAULT_ADMIN_ROLE`.

In practice, for upgrades that do not deploy net new contracts, only the `semanticVersion` parameter will be populated, with the rest left empty.

A contract name may be passed to this function repeatedly; each time, the previous address mapping for that name is overwritten with the new one. In general, this should be uncommon, as changing a core contract's address is an exceptional event.

### `configure`

```solidity
function configure(string calldata name, DeploymentConfig calldata config) external onlyRole(DEFAULT_ADMIN_ROLE)
```

Updates the stored `DeploymentConfig` for a single deployment:

*Effects*:
* Looks up the address for `name` from `_deployments`.
* Overwrites `_deploymentConfigs[addr]` with the supplied configuration.
* Emits `DeploymentConfigured(addr, config)`.

*Requirements*:
* Caller must hold `DEFAULT_ADMIN_ROLE`.
* `name` must have been previously shipped.

### `pauseAll`

```solidity
function pauseAll() external onlyRole(PAUSER_ROLE)
```

Triggers emergency pausing across all tracked deployments:

*Effects*:
* Iterates over `_deployments`, invoking `IPausable(addr).pauseAll()` on every entry marked `pausable` and not `deprecated`.
* Silently skips deployments that fail the interface call (propagates no revert) to maximize coverage in emergencies.

*Requirements*:
* Caller must hold `PAUSER_ROLE`.
* Deployments targeted must implement `IPausable`; registry does not enforce this at compile time, so operators should ensure configs reflect actual contract capabilities.

Intended for crisis response (e.g., discovered exploit) where pauser multisig needs to freeze protocol components rapidly.

---

## View Functions

### `version`

Returns the full semantic version string (e.g., `"1.9.0"`) by decoding `_semanticVersion`.

### `majorVersion`

Extracts and returns only the major component (characters before the first dot) of the semantic version.

### `getAddress`

Resolves a deployment name to its registered address. Useful for on-chain lookups by other contracts that only need the address.

### `getDeployment`

Returns both the address and `DeploymentConfig` for a given name. Preferred by off-chain services that need to inspect configuration flags alongside the address.

### `getAllDeployments`

Materializes the entire registry into parallel arrays of names, addresses, and configs. Primarily intended for off-chain introspection.

### `totalDeployments`

Returns the length of `_deployments`, which doubles as the number of rows the registry will iterate over during `pauseAll()`.


