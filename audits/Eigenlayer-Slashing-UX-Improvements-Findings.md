# Audit Findings: EigenLayer - Slashing UX Improvements

**Auditor:** Certora
**Date:** December 2025
**Status:** Draft
**Total findings:** 19

## Summary

| Severity | Discovered | Confirmed | Fixed |
|----------|------------|-----------|-------|
| Critical | 0 | - | - |
| High | 0 | - | - |
| Medium | 1 | | |
| Low | 3 | | |
| Informational | 15 | | |
| **Total** | **19** | | |

## Action Required (4 items pending)

- [ ] **M-01** (Medium): MigrateSlashers may assign an incompatible slasher address
- [ ] **L-01** (Low): Instant slasher setting leaves stale slasher field in storage
- [ ] **L-01** (Low): Major version parsing truncates multi-digit majors
- [ ] **L-01** (Low): State inconsistency for new operator allocation delay

---

## All Findings

### (PR-1643) Allocation Manager Space Savings

#### I-01. getSlasher is implemented twice
| Severity | Impact | Likelihood | Status |
|----------|--------|------------|--------|
| Informational | - | - | Pending |

**Files:** `AllocationManager.sol`, `AllocationManagerView.sol`

**Description:** The external view methods from the AllocationManagerView are going to be invoked by the AllocationManager using the `_delegateView` internal function. However, `getSlasher()` is a public view function that executes logic directly without delegating a call to AllocationManagerView. Furthermore, `getSlasher()` is also implemented in AllocationManagerView but is never used by AllocationManager.

**Recommendation:** Consider removing `getSlasher()` from AllocationManagerView.

---

### (PR-1645, PR-544) Slashing Commitments

#### M-01. MigrateSlashers may assign an incompatible slasher address
| Severity | Impact | Likelihood | Status |
|----------|--------|------------|--------|
| Medium | High | Low | Pending |

**Files:** `AllocationManager.sol`

**Description:** The previous implementation allowed the AVS to set multiple appointees enabled for the slashing call. An AVS can have different operator sets with different strategies and many AVS implementations could have split their slasher appointees with distinct implementations for each operator set. `migrateSlashers()` can be invoked by anyone and will automatically set the slasher of the AVS as the first appointee in the PermissionController. There is no safe way on-chain to ensure all AVSes have only one slasher.

**Exploit Scenario:**
1. An AVS has specified two slashers as appointees: SlasherA (for operatorSetA) and SlasherB (for operatorSetB)
2. `migrateSlashers()` is invoked. Now SlasherA is assigned as the slasher for both operator sets
3. Due to implementation limitations, SlasherA cannot slash OperatorSetB. Slashing will not be penalized for ~17.5 days until the slasher is updated

**Recommendation:** Allow only the AVS or its appointees to call `migrateSlashers()` to migrate its own operator sets (this could be done during a small window, before forceful migration).

---

#### L-01. Instant slasher setting leaves stale slasher field in storage
| Severity | Impact | Likelihood | Status |
|----------|--------|------------|--------|
| Low | Low | Low | Pending |

**Files:** `AllocationManager.sol`

**Description:** When a slasher is set with immediate effect (via `createOperatorSets` or `migrateSlashers` calling `_updateSlasher(..., instantEffectBlock=true)`), `_updateSlasher` only writes `pendingSlasher` and sets `effectBlock` to `block.number`. It does not directly update `params.slasher` in storage, creating an inconsistent on-chain representation where `_slashers[operatorSetKey]` can store `slasher = address(0)` while `pendingSlasher` is non-zero and `effectBlock` is the current block.

**Note:** `AllocationManager::getSlasher()` returns the correct slasher because it applies the pending value in-memory.

**Recommendation:** When `instantEffectBlock` is used, consider updating storage so `slasher` reflects the effective value immediately.

---

#### I-01. Re-proposing the same pending slasher restarts the delay
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** `AllocationManager::_updateSlasher()` will always overwrite `pendingSlasher` and reset the `effectBlock`. This happens even when the AVS proposes the exact same slasher address that is already stored as `pendingSlasher`. A no-op "re-proposal" restarts the `ALLOCATION_CONFIGURATION_DELAY` countdown without changing the effective outcome.

**Recommendation:** In `_updateSlasher`, treat "propose the same pendingSlasher again" as a no-op and avoid resetting `effectBlock`.

---

#### I-02. getPendingSlasher and getSlasher do not validate if the operator set exists
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** Both `getPendingSlasher()` and `getSlasher()` do not validate if an operator set exists. If an invalid operator set is passed, it will return `address(0)` as slasher and `effectBlock` set to 0.

**Recommendation:** Consider validating if a slasher exists or documenting this behavior in the NatSpec.

---

#### I-03. ALLOCATION_CONFIGURATION_DELAY is used both when setting the allocation delay info and changing the slasher
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** The `ALLOCATION_CONFIGURATION_DELAY` constant was originally used in `_setAllocationDelay()`. However, the newly introduced `_updateSlasher` function also uses this constant, even though both actions are not related to each other.

**Recommendation:** Consider using an alternative constant for the update slasher delay.

---

#### I-04. SlashingRegistryCoordinator::QuorumCreated event omits slasher parameter
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** `SlashingRegistryCoordinator::_createQuorum` was updated to take a slasher address (used when calling `AllocationManager::createOperatorSets`), but `QuorumCreated` event is still emitted without this new parameter. Off-chain indexers and monitoring cannot reconstruct the full quorum/operator-set configuration from event logs.

**Recommendation:** Extend `ISlashingRegistryCoordinator::QuorumCreated` event to include the slasher address.

---

#### I-05. Slasher migration can be gas-griefed by large appointee sets
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** `AllocationManager::migrateSlashers` calls `PermissionController.getAppointees(...)` for each operator set, which returns the full appointee list by copying an EnumerableSet to memory. Since the AVS controls how many appointees exist for the slashing selector, an AVS can make this migration step unexpectedly expensive and potentially push the transaction over the block gas limit.

**Recommendation:** Document that `migrateSlashers` can be expensive for AVSs with many slashing appointees.

---

### (PR-1655) Protocol Registry

#### L-01. Major version parsing truncates multi-digit majors
| Severity | Impact | Likelihood | Status |
|----------|--------|------------|--------|
| Low | Low | Low | Pending |

**Files:** `ProtocolRegistry.sol`

**Description:** `ProtocolRegistry::majorVersion` returns only the first byte of the stored semantic version string. This works only for single-character majors (e.g., "1.2.3" -> "1"), but produces incorrect results for multi-digit majors (e.g., "10.0.0" -> "1"). If the version string is prefixed (e.g., "v1.2.3"), it returns "v". This can break on-chain or off-chain logic relying on `majorVersion` for version gating, upgrade coordination, or signature/version-domain selection. It can also revert if `_semanticVersion` is empty.

**Recommendation:** Parse and return the full major component (all digits up to the first `.`), and handle unexpected formats safely (e.g., optional `v` prefix and empty/invalid strings).

---

#### I-01. ship lacks array length and zero-address validation
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** `ProtocolRegistry::ship` takes three parallel arrays (addresses, configs, names) but does not validate that they have the same length before iterating. If configs or names are shorter, the call will revert with "array out-of-bounds" panic. If configs/names are longer, the extra entries are silently ignored.

**Recommendation:** Add input validation to (1) require `addresses.length == configs.length == names.length`, and (2) reject `address(0)` in addresses.

---

#### I-02. Overwriting deployment names leaves orphaned configs
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** In `ProtocolRegistry::ship`, when an existing name is re-shipped with a new address, the previous address keeps its `_deploymentConfigs[oldAddr]` entry, but there is no longer any name pointing to `oldAddr`. The old config becomes effectively unreachable and permanently stored, creating silent storage bloat.

**Recommendation:** In `_appendDeployment`, if name already exists and the stored address is being overwritten, delete `_deploymentConfigs[oldAddr]`.

---

#### I-03. configure can set configs for unshipped addresses
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** `ProtocolRegistry::configure` allows the admin to write a `DeploymentConfig` for any `addr` without checking that the address is currently registered in `_deployments`. This makes it easy to accidentally configure a wrong/typo address and silently create a config entry that is not tied to any deployment name.

**Recommendation:** Restrict `configure` to only allow configuration of addresses that are already shipped/registered.

---

#### I-04. ship cannot change the contract name, as mentioned in the NatSpec
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** The `ship()` function has the comment `@dev Contract names can be overridden any number of times.` This statement is not accurate as contract names are used as keys which point to the contract address. There is no way to change the contract name using `ship`.

**Recommendation:** Update the misleading comment.

---

#### I-05. pauseAll can be blocked by one misconfigured deployment
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** `ProtocolRegistry::pauseAll` iterates over all shipped deployments and calls `IPausable(addr).pauseAll()` for those marked pausable. If any entry is misconfigured as pausable but does not implement `pauseAll()`, the entire transaction reverts and none of the deployments are paused.

**Recommendation:** Document this behaviour and cautiously maintain the registry.

---

### (PR-1646) Instant Allocation for New Operators

#### L-01. State inconsistency for new operator allocation delay
| Severity | Impact | Likelihood | Status |
|----------|--------|------------|--------|
| Low | Low | Low | Pending |

**Files:** `AllocationManager.sol`

**Description:** When a new operator registers through DelegationManager, their allocation delay is set via `AllocationManager::setAllocationDelay`, which calls `_setAllocationDelay(..., newlyRegistered=true)` and sets `effectBlock = block.number`. However, `_setAllocationDelay` does not immediately update `delay`/`isSet`; it only updates `pendingDelay` and relies on "apply pending if block.number >= effectBlock" logic. This creates a temporary storage inconsistency where `_allocationDelayInfo[operator]` can have `delay = 0` and `isSet = false`, while `pendingDelay` is non-zero and `effectBlock` is already the current block.

**Note:** `getAllocationDelay` returns the correct values because it applies the pending delay in-memory.

**Recommendation:** Update the stored `AllocationDelayInfo` immediately for newly registered operators so that `delay` and `isSet` reflect the effective value.

---

#### I-01. Operators may flood operator sets and disturb quorums and table updates
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** Apart from instant allocations for new operators, the change may indirectly introduce instant registration for new operators in the context of some AVS registrars (e.g., registrars that only require a non-zero allocation to register an operator). Instantly registering operators without a delay may introduce a risk for some AVS implementations of being flooded with operators. In the context of middleware contracts, this may DoS operator table upgrades.

**Recommendation:** Clearly communicate this potential attack vector to AVSes so they can apply suitable defence mechanisms.

---

### (PR-547) Middleware

#### I-01. getNonSignerWitnessesAndApk will only work for freshly updated tables
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** `getNonSignerWitnessesAndApk()` computes the `nonSignerWitnesses` array using the latest operator weights stored on-chain. However, `verifyCertificate()` validates certificates against the operator set state at `referenceTimestamp`, not at `block.timestamp`. The output is only valid when the operator table has been freshly updated and the current state exactly matches the state at `referenceTimestamp`.

**Recommendation:** Document this limitation.

---

#### I-02. getNonSignerWitnessesAndApk is too gas intensive
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** `BN254TableCalculatorBase::getNonSignerWitnessesAndApk` rebuilds the registered operator table, does a nested scan to check whether each operator is in `signingOperators`, and generates a Merkle proof for every non-signer. Gas grows quickly with table size; in a benchmark with 70 registered operators and no signers (all non-signers), it used ~11.3M gas.

**Recommendation:** Document in NatSpec that this function is intended for off-chain `eth_call` usage and on-chain transactions calling it may drain callers' ether balance due to gas costs.

---

#### I-03. Operator info indices can be mixed up
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** `BN254TableCalculatorBase::getOperatorIndex` returns the compacted 0-based index used by `calculateOperatorTable` (only key-registered operators, no holes). However, `getOperatorInfos` returns an array sized to the full candidate list and can contain "holes" for unregistered operators, so its array positions do not match `operatorIndex`. Integrators may mix these two index conventions and build invalid operatorIndex/Merkle proofs.

**Recommendation:** Align the functions indexing, or explicitly document that `getOperatorInfos` is sparse and must not be used to derive `operatorIndex`.

---

## Deployment Script Review

### 1. AllocationManagerView Registration
| Status |
|--------|
| Pending |

**Details:** `AllocationManagerView` is deployed in Script 5 but is not registered in `ProtocolRegistry`. As a core EigenLayer split-view contract, it can be registered with `{ pausable: false, deprecated: false }` since it is not Pausable. Furthermore, some additional implementations and beacons are included in the script.

---

### 2. StrategyBaseTVLLimits Count Source-of-Truth
| Status |
|--------|
| Pending |

**Details:** Script 5 iterates over TVL limits using `strategyBaseTVLLimits_Count()` as the loop bound. The source-of-truth for this count is unclear. If externally inflatable, this may cause upgrade-time self-DoS via unbounded iteration.

---

### 3. Pauser Grant Timing Divergence Across Chains
| Status |
|--------|
| Pending |

**Details:** Pauser assignment timing differs by chain. Destination Script 5 grants `ProtocolRegistry` pauser status immediately via `setIsPauser` (outside timelock), while source/mainnet grants it via the timelocked executor bundle. In v1.9.0-slashing-ux, `setIsPauser` occurs during scheduling, whereas in v1.9.0-slashing-ux-destination it occurs immediately.
