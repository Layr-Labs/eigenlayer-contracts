## CertificateVerifier

| File | Type | Notes |
| -------- | -------- | -------- |
| [`BN254CertificateVerifier.sol`](../../../src/contracts/multichain/BN254CertificateVerifier.sol) | Implementation | BN254 signature verification with merkle proofs |
| [`ECDSACertificateVerifier.sol`](../../../src/contracts/multichain/ECDSACertificateVerifier.sol) | Implementation | ECDSA signature verification |

Interfaces:

| File | Notes |
| -------- | -------- |
| [`IBaseCertificateVerifier.sol`](../../../src/contracts/interfaces/IBaseCertificateVerifier.sol) | Base interface for all verifiers |
| [`IBN254CertificateVerifier.sol`](../../../src/contracts/interfaces/IBN254CertificateVerifier.sol) | BN254-specific interface |
| [`IECDSACertificateVerifier.sol`](../../../src/contracts/interfaces/IECDSACertificateVerifier.sol) | ECDSA-specific interface |

---

## Overview

The CertificateVerifier contracts are responsible for verifying signatures from operator sets on destination chains. They work in conjunction with the [`OperatorTableUpdater`](./OperatorTableUpdater.md) to maintain operator information and verify certificates against the most recent operator tables. These contracts support two signature schemes: BN254 for aggregated signatures and ECDSA for individual signatures.

Both verifiers implement staleness checks based on a `maxStalenessPeriod` to ensure certificates are not verified against outdated operator information. They also support various threshold verification methods including proportion-based and nominal stake thresholds.

---

## BN254CertificateVerifier

The `BN254CertificateVerifier` implements BN254 signature verification with support for aggregated signatures and merkle proof-based non-signer tracking. It uses signature aggregation to efficiently verify certificates with a single pairing check.

### Core Functions

#### `updateOperatorTable`

```solidity
/**
 * @notice Updates the operator table for a specific operator set
 * @param operatorSet The operator set to update
 * @param referenceTimestamp The reference timestamp for the update
 * @param operatorSetInfo The BN254 operator set information
 * @param operatorSetConfig The operator set configuration
 */
function updateOperatorTable(
    OperatorSet calldata operatorSet,
    uint32 referenceTimestamp,
    BN254OperatorSetInfo memory operatorSetInfo,
    OperatorSetConfig calldata operatorSetConfig
) external onlyTableUpdater;
```

Updates the operator table with new operator information. This function can only be called by the authorized `OperatorTableUpdater`.

*Effects*:
* Stores `operatorSetInfo` at `_operatorSetInfos[operatorSetKey][referenceTimestamp]`
* Updates `_latestReferenceTimestamps[operatorSetKey]` to `referenceTimestamp`
* Updates `_operatorSetOwners[operatorSetKey]` to `operatorSetConfig.owner`
* Updates `_maxStalenessPeriods[operatorSetKey]` to `operatorSetConfig.maxStalenessPeriod`
* Emits a `TableUpdated` event

*Requirements*:
* Caller MUST be the `operatorTableUpdater`
* The `referenceTimestamp` MUST be greater than the latest reference timestamp

#### `verifyCertificate`

```solidity
/**
 * @notice A struct containing a BN254 signature and metadata
 * @param referenceTimestamp The timestamp the certificate references
 * @param messageHash The hash that was signed
 * @param signature The aggregated BN254 signature
 * @param apk The aggregate public key in G2
 * @param nonSignerWitnesses Array of non-signer information with merkle proofs
 */
struct BN254Certificate {
    uint32 referenceTimestamp;
    bytes32 messageHash;
    BN254.G1Point signature;
    BN254.G2Point apk;
    BN254OperatorInfoWitness[] nonSignerWitnesses;
}

/**
 * @notice Verifies a BN254 certificate
 * @param operatorSet The operator set that created the certificate
 * @param cert The certificate to verify
 * @return signedStakes The stakes that signed for each stake type
 */
function verifyCertificate(
    OperatorSet memory operatorSet,
    BN254Certificate memory cert
) external returns (uint256[] memory signedStakes);
```

Verifies a BN254 certificate by checking the aggregated signature against the operator set's public keys.

*Process*:
* Validates the certificate timestamp against staleness requirements
* Initializes signed stakes with total stakes from the operator set
* Processes non-signer witnesses:
  * Verifies merkle proofs for non-signers (or uses cached data)
  * Subtracts non-signer stakes from total signed stakes
  * Aggregates non-signer public keys
* Calculates signer aggregate public key by subtracting non-signers from total
* Verifies the BLS signature using pairing checks

*Requirements*:
* The certificate MUST NOT be stale (based on `maxStalenessPeriod`)
* The root at `referenceTimestamp` MUST be valid (not disabled)
* The operator set info MUST exist for the `referenceTimestamp`
* All merkle proofs MUST be valid
* The BLS signature MUST verify correctly

#### `verifyCertificateProportion`

```solidity
/**
 * @notice Verifies a certificate meets proportion thresholds
 * @param operatorSet The operator set that created the certificate
 * @param cert The certificate to verify
 * @param totalStakeProportionThresholds Thresholds in basis points (10000 = 100%)
 * @return bool Whether all thresholds are met
 */
function verifyCertificateProportion(
    OperatorSet memory operatorSet,
    BN254Certificate memory cert,
    uint16[] memory totalStakeProportionThresholds
) external returns (bool);
```

Verifies that a certificate meets specified proportion thresholds for each stake type.

*Effects*:
* Same as `verifyCertificate` (caches non-signer data)

*Requirements*:
* All requirements from `verifyCertificate`
* For each stake type: `signedStakes[i] >= (totalStakes[i] * threshold[i]) / 10000`

### Caching Mechanism

The BN254CertificateVerifier implements a caching mechanism for non-signer operator information:

```solidity
/**
 * @notice Witness data for a non-signing operator
 * @param operatorIndex The index of the operator in the operator set
 * @param operatorInfo The operator's information
 * @param operatorInfoProof Merkle proof for the operator info
 */
struct BN254OperatorInfoWitness {
    uint32 operatorIndex;
    BN254OperatorInfo operatorInfo;
    bytes operatorInfoProof;
}
```

When processing non-signers:
1. First checks if operator info is cached at `_operatorInfos[operatorSetKey][referenceTimestamp][operatorIndex]`
2. If not cached, verifies the merkle proof and caches the data
3. If cached, uses the cached data without re-verifying

This optimization significantly reduces gas costs for repeated verifications with the same non-signers.

---

## ECDSACertificateVerifier

The `ECDSACertificateVerifier` implements ECDSA signature verification where each operator signs individually. It stores complete operator information on-chain rather than using merkle trees.

### Core Functions

#### `updateOperatorTable`

```solidity
/**
 * @notice Updates the operator table for a specific operator set
 * @param operatorSet The operator set to update
 * @param referenceTimestamp The reference timestamp for the update
 * @param operatorInfos Array of ECDSA operator information
 * @param operatorSetConfig The operator set configuration
 */
function updateOperatorTable(
    OperatorSet calldata operatorSet,
    uint32 referenceTimestamp,
    ECDSAOperatorInfo[] calldata operatorInfos,
    OperatorSetConfig calldata operatorSetConfig
) external onlyTableUpdater;
```

Updates the operator table with new operator information. Unlike BN254, this stores individual operator info rather than aggregated data.

*Effects*:
* Stores the number of operators at `_numOperators[operatorSetKey][referenceTimestamp]`
* Stores each operator info at `_operatorInfos[operatorSetKey][referenceTimestamp][index]`
* Updates `_latestReferenceTimestamps[operatorSetKey]` to `referenceTimestamp`
* Updates `_operatorSetOwners[operatorSetKey]` to `operatorSetConfig.owner`
* Updates `_maxStalenessPeriods[operatorSetKey]` to `operatorSetConfig.maxStalenessPeriod`
* Emits a `TableUpdated` event

*Requirements*:
* Caller MUST be the `operatorTableUpdater`
* The `referenceTimestamp` MUST be greater than the latest reference timestamp

#### `verifyCertificate`

```solidity
/**
 * @notice A struct containing ECDSA signatures and metadata
 * @param referenceTimestamp The timestamp the certificate references
 * @param messageHash The hash that was signed
 * @param sig Concatenated ECDSA signatures (65 bytes each)
 */
struct ECDSACertificate {
    uint32 referenceTimestamp;
    bytes32 messageHash;
    bytes sig;
}

/**
 * @notice Verifies an ECDSA certificate
 * @param operatorSet The operator set that created the certificate
 * @param cert The certificate to verify
 * @return signedStakes The stakes that signed for each stake type
 */
function verifyCertificate(
    OperatorSet calldata operatorSet,
    ECDSACertificate calldata cert
) external view returns (uint256[] memory);
```

Verifies an ECDSA certificate by checking individual signatures from operators.

*Process*:
* Validates the certificate timestamp against staleness requirements
* Computes the EIP-712 digest for the certificate
* Parses concatenated signatures and recovers signers
* For each recovered signer:
  * Verifies the signer is a registered operator
  * Adds the operator's weights to signed stakes
* Returns the total signed stakes

*Requirements*:
* The certificate MUST NOT be stale (based on `maxStalenessPeriod`)
* The root at `referenceTimestamp` MUST be valid (not disabled)
* The operator table MUST exist for the `referenceTimestamp`
* Signatures MUST be ordered by signer address (ascending)
* All signers MUST be registered operators
* Each signature MUST be valid

#### `calculateCertificateDigest`

```solidity
/**
 * @notice Calculates the EIP-712 digest for a certificate
 * @param referenceTimestamp The reference timestamp
 * @param messageHash The message hash
 * @return The EIP-712 digest
 */
function calculateCertificateDigest(
    uint32 referenceTimestamp, 
    bytes32 messageHash
) public view returns (bytes32);
```

Computes the EIP-712 structured data hash for ECDSA certificate signing.

*Returns*:
* EIP-712 digest using domain separator and certificate type hash

### View Functions

Both verifiers provide view functions for querying operator information:

**BN254CertificateVerifier**:
- `getNonsignerOperatorInfo(operatorSet, referenceTimestamp, operatorIndex)` - Get cached non-signer info
- `isNonsignerCached(operatorSet, referenceTimestamp, operatorIndex)` - Check if non-signer is cached
- `getOperatorSetInfo(operatorSet, referenceTimestamp)` - Get complete operator set info

**ECDSACertificateVerifier**:
- `getOperatorInfos(operatorSet, referenceTimestamp)` - Get all operator infos
- `getOperatorInfo(operatorSet, referenceTimestamp, operatorIndex)` - Get specific operator info
- `getOperatorCount(operatorSet, referenceTimestamp)` - Get number of operators
- `getTotalStakes(operatorSet, referenceTimestamp)` - Get total stakes for all stake types

**Common View Functions**:
- `getOperatorSetOwner(operatorSet)` - Get the owner address for an operator set
- `maxOperatorTableStaleness(operatorSet)` - Get the maximum staleness period
- `latestReferenceTimestamp(operatorSet)` - Get the latest reference timestamp

---

## Staleness and Validity Checks

Both verifiers implement staleness checks to ensure certificates are not verified against outdated operator information:

1. **Staleness Check**: If `maxStalenessPeriod > 0`, certificates must satisfy:
   ```
   block.timestamp <= referenceTimestamp + maxStalenessPeriod
   ```

2. **Root Validity Check**: The global root at the certificate's reference timestamp must not be disabled by the `OperatorTableUpdater`

3. **Timestamp Existence**: The reference timestamp must have valid operator table data

Setting `maxStalenessPeriod = 0` disables staleness checks, allowing certificates to be verified indefinitely against a specific operator table.
