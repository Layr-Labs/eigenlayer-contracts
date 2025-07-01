## OperatorTableUpdater

| File | Type | Proxy |
| -------- | -------- |  -------- | 
| [`OperatorTableUpdater.sol`](../../../src/contracts/multichain/OperatorTableUpdater.sol) | Singleton | Transparent Proxy |

Libraries and Mixins:

| File | Notes |
| -------- | -------- |
| [`OwnableUpgradeable.sol`](https://github.com/OpenZeppelin/openzeppelin-contracts-upgradeable/blob/v4.9.0/contracts/access/OwnableUpgradeable.sol) | Access control |
| [`SemVerMixin.sol`](../../../src/contracts/mixins/SemVerMixin.sol) | Versioning |
| [`Merkle.sol`](../../../src/contracts/libraries/Merkle.sol) | Merkle proof verification |

## Overview

The `OperatorTableUpdater` is responsible for updating the `GlobalTableRoot` and updating operator tables from merkle proofs against the `GlobalTableRoot`. The contract is deployed on every destination chain. The contract maintains a set of valid global table roots that are confirmed by a designated generator, and allows updating individual operator tables by providing merkle proofs against these roots.

The contract supports both BN254 and ECDSA operator tables and routes updates to the appropriate certificate verifier based on the curve type.

## Parameterization
Upon initialization, the `generator` is updated. The `generator` is represented in storage as an operatorSet. The `generator` should be considered a ghost-operatorSet` since it does not exist in the core protocol, does not have stake backing it, and is not transported to other chains via the multichain protocol. It can only be updated upon initialization or by a [privileged role](#updategenerator). This entity is the same across all destination chains. 

The following values are set upon initialization: 

* `generator` is an EigenLabs-run entity that signs off on `globalTableRoots`. The operatorSet is of size 1. 
* `globalRootConfirmationThreshold`: 10000. The threshold in basis points required for global root confirmation. Since the operatorSet is of size 1 a single signature is needed.
* `referenceTimestamp`: A past timestamp at which the `generator` is set. We hardcode this value to 1 upon initialization. This value is also the `latestReferenceTimestamp`. Once roots are updated, the `latestReferenceTimestamp` will increase. *Note: the reference timestamp for the `generator`, given by `operatorTableUpdater.getGeneratorReferenceTimestamp` will remain 1 unless [`updateGenerator`](#updategenerator) is called*.
* `generatorInfo`: The key material needed to verify certificates of the `generator`
* `operatorSetConfig`: A configuration for the `generator` 
    * `maxStalenessPeriod`: 0. Set to zero to confirm `globalTableRoots` without updating the `generator` operatorSet. See [`CertificateVerifier`](./CertificateVerifier.md#overview) for specifics`OperatorTableUpdater`. It is the same across all destination chains, even for destination chains that are supported after the initial deployment. 
    * `owner`: Unused parameter for `Generator`


---

## Global Root Confirmation

Global table roots must be confirmed by the `generator` before operator tables can be updated against them.

### `confirmGlobalTableRoot`

```solidity
/**
 * @notice Confirms Global operator table root
 * @param globalTableRootCert certificate of the root
 * @param globalTableRoot merkle root of all operatorSet tables
 * @param referenceTimestamp timestamp of the root
 * @param referenceBlockNumber block number of the root
 * @dev Any entity can submit with a valid certificate signed off by the `generator`
 * @dev The `msgHash` in the `globalOperatorTableRootCert` is the hash of the `globalOperatorTableRoot`
 */
function confirmGlobalTableRoot(
    BN254Certificate calldata globalTableRootCert,
    bytes32 globalTableRoot,
    uint32 referenceTimestamp,
    uint32 referenceBlockNumber
) external;
```

Confirms a new global table root by verifying a BN254 certificate signed by the `generator`. See [`BN254CertificateVerifier`](./CertificateVerifier.md#bn254certificateverifier) for certificate verification. Roots are append only and cannot be overridden, only [disabled](#disableroot). 

*Effects*:
* Updates `_latestReferenceTimestamp` to the new `referenceTimestamp`
* Sets `_referenceBlockNumbers[referenceTimestamp]` to `referenceBlockNumber`
* Sets `_referenceTimestamps[referenceBlockNumber]` to `referenceTimestamp`
* Sets `_globalTableRoots[referenceTimestamp]` to `globalTableRoot`
* Sets `_isRootValid[globalTableRoot]` to `true`
* Emits a `NewGlobalTableRoot` event

*Requirements*:
* The contract MUST NOT be paused for global root updates
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
 * @notice Updates an operator table
 * @param referenceTimestamp the reference block number of the globalTableRoot
 * @param globalTableRoot the new globalTableRoot
 * @param operatorSetIndex the index of the given operatorSet being updated
 * @param proof the proof of the leaf at index against the globalTableRoot
 * @param operatorTableBytes the bytes of the operator table
 * @dev Depending on the decoded KeyType, the tableInfo will be decoded
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
* The contract MUST NOT be paused for operator table updates
* The `globalTableRoot` MUST be valid (not disabled)
* The `referenceTimestamp` MUST be greater than the latest timestamp for the operator set
* The merkle proof MUST verify the operator table's inclusion in the global root
* The `globalTableRoot` at `referenceTimestamp` MUST match the provided root
* Meets all requirements in [`ecdsaCertificateVerifier.updateOperatorTable`](./CertificateVerifier.md#updateoperatortable) or [`bn254CertificateVerifier.updateOperatorTable`](./CertificateVerifier.md#updateoperatortable-1)

---

## System Configuration

The `owner` can configure the `generator` and confirmation parameters.

### `setGenerator`

```solidity
/**
 * @notice Set the operatorSet which certifies against global roots
 * @param operatorSet the operatorSet which certifies against global roots
 * @dev The `operatorSet` is used to verify the certificate of the global table root
 * @dev Only callable by the owner of the contract
 */
function setGenerator(
    OperatorSet calldata operatorSet
) external;
```

Updates the operator set responsible for confirming global table roots.

*Effects*:
* Updates `_generator` to the new `operatorSet`
* Emits a `GeneratorUpdated` event

*Requirements*:
* Caller MUST be the `owner`

### `setGlobalRootConfirmationThreshold`

```solidity
/**
 * @notice The threshold, in bps, for a global root to be signed off on and updated
 * @param bps The threshold in basis points
 * @dev Only callable by the owner of the contract
 */
function setGlobalRootConfirmationThreshold(
    uint16 bps
) external;
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
 * @notice Disables a global table root
 * @param globalTableRoot the global table root to disable
 * @dev Only callable by the pauser
 */
function disableRoot(
    bytes32 globalTableRoot
) external;
```

Disables a global table root, preventing further operator table updates against it. This function also prevents the `CertificateVerifier` from verifying certificates. The function is intended to prevent a malicious or invalid root from being used by downstream consumers. Once a root is disabled, it cannot be re-enabled. 

*Effects*:
* Sets `_isRootValid[globalTableRoot]` to `false`
* Emits a `GlobalRootDisabled` event

*Requirements*:
* Caller MUST be the `pauser`
* The `globalTableRoot` MUST exist and be currently valid

### `updateGenerator`

```solidity
/**
 * @notice Updates the operator table for the generator
 * @param referenceTimestamp The reference timestamp of the operator table update
 * @param generatorInfo The operatorSetInfo for the generator
 * @param generatorConfig The operatorSetConfig for the generator
 * @dev We have a separate function for updating this operatorSet since it's not transported and updated
 *      in the same way as the other operatorSets
 * @dev Only callable by the owner of the contract
 */
function updateGenerator(
    uint32 referenceTimestamp,
    BN254OperatorSetInfo calldata generatorInfo,
    OperatorSetConfig calldata generatorConfig
) external;
```

Updates the operator table for the `generator` itself. This operatorSet is a ["shadow-operatorSet"](#parameterization), so it must be updated manually

*Effects*:
* Calls `bn254CertificateVerifier.updateOperatorTable` for the `generator`

*Requirements*:
* Caller MUST be the `owner`
* Meet all requirements in [`bn254CertificateVerifier.updateOperatorTable`](../destination/CertificateVerifier.md#updateoperatortable-1)