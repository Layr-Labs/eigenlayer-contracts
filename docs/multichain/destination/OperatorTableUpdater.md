## OperatorTableUpdater

| File | Notes |
| -------- | -------- |
| [`OperatorTableUpdater.sol`](../../../src/contracts/multichain/OperatorTableUpdater.sol) | Updates operator tables on destination chains |
| [`OperatorTableUpdaterStorage.sol`](../../../src/contracts/multichain/OperatorTableUpdaterStorage.sol) | Storage layout and state variables |
| [`IOperatorTableUpdater.sol`](../../../src/contracts/interfaces/IOperatorTableUpdater.sol) | Interface |

Libraries and Mixins:

| File | Notes |
| -------- | -------- |
| [`OwnableUpgradeable.sol`](https://github.com/OpenZeppelin/openzeppelin-contracts-upgradeable/blob/v4.9.0/contracts/access/OwnableUpgradeable.sol) | Access control |
| [`SemVerMixin.sol`](../../../src/contracts/mixins/SemVerMixin.sol) | Versioning |
| [`Merkle.sol`](../../../src/contracts/libraries/Merkle.sol) | Merkle proof verification |

## Overview

The `OperatorTableUpdater` is responsible for updating the `GlobalTableRoot` and updating operator tables from merkle proofs against the `GlobalTableRoot`. The contract is deployed on every destination chain. The contract maintains a set of valid global table roots that are confirmed by a designated global root confirmer set, and allows updating individual operator tables by providing merkle proofs against these roots.

The contract supports both BN254 and ECDSA operator tables and routes updates to the appropriate certificate verifier based on the curve type.

## Parameterization
* `GlobalRootConfirmerSet`, also known as the `Generator`, is an EigenLabs-run entity that signs off on `GlobalTableRoots`. 
    * `maxStalenessPeriod`: 0. Set to zero to confirm roots without updating the operatorSet. See [`CertificateVerifier`](./CertificateVerifier.md#bn254certificateverifier) for specifics
    * `globalRootConfirmationThreshold`: 10000. The threshold in basis points required for global root confirmation

---

## Global Root Confirmation

Global table roots must be confirmed by the `globalRootConfirmerSet` (ie. `Generator`) before operator tables can be updated against them.

### `confirmGlobalTableRoot`

```solidity
/**
 * @notice Confirms a global table root with a BN254 certificate
 * @param globalTableRootCert The BN254 certificate confirming the global table root
 * @param globalTableRoot The global table root being confirmed
 * @param referenceTimestamp The reference timestamp for the global table root
 * @param referenceBlockNumber The reference block number for the global table root
 */
function confirmGlobalTableRoot(
    BN254Certificate calldata globalTableRootCert,
    bytes32 globalTableRoot,
    uint32 referenceTimestamp,
    uint32 referenceBlockNumber
) external;
```

Confirms a new global table root by verifying a BN254 certificate signed by the `globalRootConfirmerSet`. See [`BN254CertificateVerifier`](./CertificateVerifier.md#bn254certificateverifier) for certificate verification. Roots are append only and cannot be overridden, only [disabled](#disableroot). 

*Effects*:
* Updates `_latestReferenceTimestamp` to the new `referenceTimestamp`
* Sets `_referenceBlockNumbers[referenceTimestamp]` to `referenceBlockNumber`
* Sets `_referenceTimestamps[referenceBlockNumber]` to `referenceTimestamp`
* Sets `_globalTableRoots[referenceTimestamp]` to `globalTableRoot`
* Sets `_isRootValid[globalTableRoot]` to `true`
* Emits a `NewGlobalTableRoot` event

*Requirements*:
* The `referenceTimestamp` MUST NOT be in the future
* The `referenceTimestamp` MUST be greater than `_latestReferenceTimestamp`
* The certificate's `messageHash` MUST match the expected EIP-712 hash
* The certificate MUST be valid according to the `globalRootConfirmationThreshold`

---

## Operator Table Updates

Once a global root is confirmed, individual operator tables can be updated by providing merkle proofs against the root. 

### `updateOperatorTable`

```solidity
/**
 * @notice Updates an operator table for a specific operator set
 * @param referenceTimestamp The reference timestamp of the operator table
 * @param globalTableRoot The global table root to verify against
 * @param operatorSetIndex The index of the operator set in the global table
 * @param proof The merkle proof for the operator table
 * @param operatorTableBytes The encoded operator table data
 */
function updateOperatorTable(
    uint32 referenceTimestamp,
    bytes32 globalTableRoot,
    uint32 operatorSetIndex,
    bytes calldata proof,
    bytes calldata operatorTableBytes
) external;
```

Updates an operator table by verifying its inclusion in a confirmed global table root via merkle proof. The function decodes the operator table data and routes the update to the appropriate certificate verifier based on the curve type.

*Effects*:
* For BN254 operator sets:
  * Calls `bn254CertificateVerifier.updateOperatorTable` with the decoded operator info
* For ECDSA operator sets:
  * Calls `ecdsaCertificateVerifier.updateOperatorTable` with the decoded operator info

*Requirements*:
* The `globalTableRoot` MUST be valid (not disabled)
* The `referenceTimestamp` MUST be greater than the latest timestamp for the operator set
* The merkle proof MUST verify the operator table's inclusion in the global root
* The `globalTableRoot` at `referenceTimestamp` MUST match the provided root

---

## System Configuration

The `owner` can configure the `globalRootConfirmerSet` and confirmation parameters.

### `setGlobalRootConfirmerSet`

```solidity
/**
 * @notice Sets the operator set that confirms global table roots
 * @param operatorSet The new global root confirmer set
 */
function setGlobalRootConfirmerSet(
    OperatorSet calldata operatorSet
) external onlyOwner;
```

Updates the operator set responsible for confirming global table roots.

*Effects*:
* Updates `_globalRootConfirmerSet` to the new `operatorSet`
* Emits a `GlobalRootConfirmerSetUpdated` event

*Requirements*:
* Caller MUST be the `owner`

### `setGlobalRootConfirmationThreshold`

```solidity
/**
 * @notice Sets the threshold for global root confirmation
 * @param bps The threshold in basis points
 */
function setGlobalRootConfirmationThreshold(
    uint16 bps
) external onlyOwner;
```

Sets the stake proportion threshold required for confirming global table roots.

*Effects*:
* Updates `globalRootConfirmationThreshold` to `bps`
* Emits a `GlobalRootConfirmationThresholdUpdated` event

*Requirements*:
* Caller MUST be the `owner`
* `bps` MUST NOT exceed `MAX_BPS` (10000)

### `disableRoot`

```solidity
/**
 * @notice Disables a previously confirmed global table root
 * @param globalTableRoot The root to disable
 */
function disableRoot(
    bytes32 globalTableRoot
) external onlyOwner;
```

Disables a global table root, preventing further operator table updates against it. This function also prevents the `CertificateVerifier` from verifying certificates. The function is intended to prevent a malicious or invalid root from being used by downstream consumers. 

*Effects*:
* Sets `_isRootValid[globalTableRoot]` to `false`
* Emits a `GlobalRootDisabled` event

*Requirements*:
* Caller MUST be the `owner`
* The `globalTableRoot` MUST exist and be currently valid

### `updateGlobalRootConfirmerSet`

```solidity
/**
 * @notice Updates the operator table for the global root confirmer set
 * @param referenceTimestamp The reference timestamp for the update
 * @param globalRootConfirmerSetInfo The BN254 operator set info
 * @param globalRootConfirmerSetConfig The operator set config
 */
function updateGlobalRootConfirmerSet(
    uint32 referenceTimestamp,
    BN254OperatorSetInfo calldata globalRootConfirmerSetInfo,
    OperatorSetConfig calldata globalRootConfirmerSetConfig
) external onlyOwner;
```

Updates the operator table for the `globalRootConfirmerSet` itself, enabling it to sign future global roots.

*Effects*:
* Calls `bn254CertificateVerifier.updateOperatorTable` for the `globalRootConfirmerSet`

*Requirements*:
* Caller MUST be the `owner`
