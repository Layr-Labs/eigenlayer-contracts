# KeyRegistrar

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`KeyRegistrar.sol`](../../src/contracts/permissions/KeyRegistrar.sol) | Singleton | Transparent proxy |

The `KeyRegistrar` manages cryptographic keys for operators across different operator sets. It supports both ECDSA and BN254 key types and ensures global uniqueness of keys across all operator sets.

Key features:
* **Per-OperatorSet Configuration**: Each operator set must be configured with a specific curve type before keys can be registered
* **Global Key Registry**: Keys are globally unique - once registered, a key cannot be reused across operatorSets or operators

Keys are stored in a 2-way mapping:
1. (operator, operatorSet) to key
2. keyHash to operator address

---

## Operator Set Configuration

An AVS must configure the operator set with a specific curve type.

### `configureOperatorSet`

```solidity
/**
 * @notice Configures an operator set with curve type
 * @param operatorSet The operator set to configure
 * @param curveType Type of curve (ECDSA, BN254)
 * @dev Only authorized callers for the AVS can configure operator sets
 */
function configureOperatorSet(OperatorSet memory operatorSet, CurveType curveType) external;
```

Configures an operator set to use a specific cryptographic curve type. This must be called before any keys can be registered for the operator set. 
*Note: Registering for an operatorSet in the core protocol does not require a key to be registered. However, the AVS may have logic that gates registration based on a key being registered in the `KeyRegistrar`.*

*Effects*:
* Sets the curve type for the specified operator set
* Emits an `OperatorSetConfigured` event

*Requirements*:
* Caller MUST be authorized for the AVS (via PermissionController)
* The operator set MUST NOT already be configured
* The curve type MUST be either ECDSA or BN254

---

## Key Registration

Key registration is segmented by curve type: ECDSA and BN254. 

### ECDSA Key Registration

#### `registerKey` (ECDSA)

```solidity
/**
 * @notice Registers a cryptographic key for an operator with a specific operator set
 * @param operator Address of the operator to register key for
 * @param operatorSet The operator set to register the key for
 * @param pubkey Public key bytes
 * @param signature Signature proving ownership (only needed for BN254 keys)
 * @dev Can be called by operator directly or by addresses they've authorized via PermissionController
 * @dev Reverts if key is already registered
 */
function registerKey(
    address operator,
    OperatorSet memory operatorSet,
    bytes calldata pubkey,
    bytes calldata signature
) external;
```

For ECDSA keys:
- `pubkey`: 20 bytes representing the Ethereum address
- `signature`: EIP-712 signature from the key's private key

*Effects*:
* Registers the key for the operator in the specified operator set
* Adds the key to the global registry
* Associates the key hash with the operator address
* Emits a `KeyRegistered` event with curve type ECDSA

*Requirements*:
* Caller MUST be the operator or authorized via PermissionController
* The operator MUST NOT be already registered for the operatorSet in the `KeyRegistrar`
* The key MUST be exactly 20 bytes
* The key MUST NOT be the zero address
* The key MUST NOT already be registered globally (by hash)
* The signature MUST be valid

#### `getECDSAKeyRegistrationMessageHash`

```solidity
/**
 * @notice Returns the message hash for ECDSA key registration
 * @param operator The operator address
 * @param operatorSet The operator set
 * @param keyAddress The address of the key
 * @return The message hash for signing
 */
function getECDSAKeyRegistrationMessageHash(
    address operator,
    OperatorSet memory operatorSet,
    address keyAddress
) external view returns (bytes32);
```

Returns the message hash that must be signed over for ECDSA key registration. 


### BN254 Key Registration

BN254 keys registration requires passing in G1 and G2 points.

#### `registerKey` (BN254)

```solidity
/**
 * @notice Registers a cryptographic key for an operator with a specific operator set
 * @param operator Address of the operator to register key for
 * @param operatorSet The operator set to register the key for
 * @param pubkey Public key bytes
 * @param signature Signature proving ownership (only needed for BN254 keys)
 * @dev Can be called by operator directly or by addresses they've authorized via PermissionController
 * @dev Reverts if key is already registered
 */
function registerKey(
    address operator,
    OperatorSet memory operatorSet,
    bytes calldata pubkey,
    bytes calldata signature
) external;
```

For BN254 keys:
- `pubkey`: [Encoded](#encodebn254keydata) BN254 key data containing G1 and G2 points
- `signature`: BN254 signature proving ownership over a set [digest](#getbn254keyregistrationmessagehash)

*Effects*:
* Registers the BN254 key for the operator in the specified operator set
* Adds the key hash to the global registry
* Emits a `KeyRegistered` event with curve type BN254

*Requirements*:
* Caller MUST be the operator or authorized via PermissionController
* The operator MUST NOT be already registered for the operatorSet in the `KeyRegistrar`
* The key MUST contain valid G1 and G2 points
* The G1 point MUST NOT be the zero point
* The key MUST NOT already be registered globally (by hash)
* The signature MUST be valid

#### `encodeBN254KeyData`

```solidity
/**
 * @notice Encodes the BN254 key data into a bytes array
 * @param g1Point The BN254 G1 public key
 * @param g2Point The BN254 G2 public key
 * @return The encoded key data
 */
function encodeBN254KeyData(
    BN254.G1Point memory g1Point,
    BN254.G2Point memory g2Point
) external pure returns (bytes memory);
```

Utility function to properly encode BN254 key data for registration.

#### `getBN254KeyRegistrationMessageHash`

```solidity
/**
 * @notice Returns the message hash for BN254 key registration
 * @param operator The operator address
 * @param operatorSet The operator set
 * @param keyData The BN254 key data
 * @return The message hash for signing
 */
function getBN254KeyRegistrationMessageHash(
    address operator,
    OperatorSet memory operatorSet,
    bytes calldata keyData
) external view returns (bytes32);
```

Returns the message hash that must be signed over for BN254 key registration. 

---

## Key Deregistration

### `deregisterKey`

```solidity
/**
 * @notice Deregisters a cryptographic key for an operator with a specific operator set
 * @param operator Address of the operator to deregister key for
 * @param operatorSet The operator set to deregister the key from
 * @dev Can be called by avs directly or by addresses they've authorized via PermissionController
 * @dev Reverts if key was not registered
 * @dev Keys remain in global key registry to prevent reuse
 */
function deregisterKey(address operator, OperatorSet memory operatorSet) external;
```

Removes an operator's key from the specified operator set. Note that the key remains in the global registry to prevent reuse.

*Effects*:
* Removes the key from the operator's record for the operator set
* Emits a `KeyDeregistered` event
* The key remains in the global registry

*Requirements*:
* Caller MUST be authorized for the AVS (via PermissionController)
* The operator set MUST be configured
* The operator MUST have a registered key for this operator set

---
