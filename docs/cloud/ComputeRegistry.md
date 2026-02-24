# ComputeRegistry

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`ComputeRegistry.sol`](../../src/contracts/cloud/ComputeRegistry.sol) | Singleton | Transparent proxy |
| [`ComputeRegistryStorage.sol`](../../src/contracts/cloud/ComputeRegistryStorage.sol) | Storage | |
| [`IComputeRegistry.sol`](../../src/contracts/interfaces/IComputeRegistry.sol) | Interface | |

Libraries and Mixins:

| File | Notes |
| -------- | -------- |
| [`PermissionControllerMixin.sol`](../../src/contracts/mixins/PermissionControllerMixin.sol) | Account delegation |
| [`SignatureUtilsMixin.sol`](../../src/contracts/mixins/SignatureUtilsMixin.sol) | EIP-712 signature verification |
| [`OperatorSetLib.sol`](../../src/contracts/libraries/OperatorSetLib.sol) | Encode/decode operator sets |

## Overview

The `ComputeRegistry` manages the registration of operator sets for compute services. It serves as the entry point for AVSs to register their operator sets for compute capabilities by signing Terms of Service (TOS) agreements. The contract ensures that operator sets meet specific requirements before registration, including proper key configuration, active generation reservations, and at least one release.

The `ComputeRegistry` acts as a gatekeeper for compute services, coordinating with multiple protocol components:
* **AllocationManager**: Validates operator set existence
* **KeyRegistrar**: Ensures proper cryptographic key configuration
* **CrossChainRegistry**: Verifies multichain compatibility through generation reservations
* **ReleaseManager**: Confirms operator sets have available releases
* **PermissionController**: Manages authorization for registration/deregistration actions

## Key Concepts

### Terms of Service (TOS)

The TOS is a legal agreement that operator sets must sign to register for compute services. The TOS hash is immutable and set at contract deployment. All registrations require a valid EIP-712 signature of the TOS agreement with the following structure:

```solidity
struct TOSSignature {
    address signer;      // Address that signed the TOS
    bytes32 tosHash;     // Hash of the Terms of Service
    bytes signature;     // EIP-712 signature
}
```

The signature must be created using the EIP-712 standard with:
* Domain name: "EigenLayer"
* Domain version: "1" (major version only)
* Type hash: `TOSAgreement(bytes32 tosHash,address avs,uint32 operatorSetId,address signer,uint256 expiry)`
* Expiry: MAX_EXPIRY (type(uint256).max) - signatures never expire

### Registration Requirements

Before an operator set can register for compute services, it must satisfy the following requirements:

1. **Valid Operator Set**: Must exist in the AllocationManager
2. **Curve Type Configuration**: Must have a configured curve type (ECDSA or BN254) in the KeyRegistrar
3. **Active Generation Reservation**: Must have an active generation reservation in the CrossChainRegistry for multichain support
4. **Available Releases**: Must have at least one release in the ReleaseManager

These checks ensure operator sets are fully configured and ready for compute operations before registration.

---

## Registration Flow

### `registerForCompute`

```solidity
/**
 * @notice Registers an operator set for compute services
 * @param operatorSet The operator set to register
 * @param signature The EIP-712 signature of the Terms of Service agreement
 * @dev The caller must have permission to call on behalf of the operatorSet.avs through the PermissionController
 * @dev The signature must be a valid EIP-712 signature of the Terms of Service with expiry set to MAX_EXPIRY
 * @dev Reverts for:
 *      - InvalidPermissions: Caller does not have permission to call on behalf of operatorSet.avs
 *      - InvalidOperatorSet: The operator set does not exist in the AllocationManager
 *      - OperatorSetAlreadyRegistered: The operator set is already registered for compute
 *      - CurveTypeNotSet: The operator set has not configured a curve type in the KeyRegistrar
 *      - NoActiveGenerationReservation: The operator set does not have an active generation reservation
 *      - NoReleases: The operator set does not have any releases in the ReleaseManager
 *      - InvalidSignature: The provided signature is invalid or does not match the expected signer
 * @dev Emits the following events:
 *      - OperatorSetRegistered: When the operator set is successfully registered with the TOS signature
 */
function registerForCompute(OperatorSet memory operatorSet, bytes memory signature) external;
```

Registers an operator set for compute services by validating all requirements and storing the TOS signature.

*Effects*:
* Marks the operator set as registered for compute
* Stores the TOS signature information (signer, TOS hash, and signature bytes)
* Emits an `OperatorSetRegistered` event with full signature details

*Requirements*:
* Caller MUST have permission to call on behalf of the operator set's AVS (via PermissionController)
* Operator set MUST exist in the AllocationManager
* Operator set MUST NOT already be registered
* Operator set MUST have a configured curve type (not NONE) in the KeyRegistrar
* Operator set MUST have an active generation reservation in the CrossChainRegistry
* Operator set MUST have at least one release in the ReleaseManager
* Signature MUST be valid and match the expected TOS agreement format

---

## Deregistration Flow

### `deregisterFromCompute`

```solidity
/**
 * @notice Deregisters an operator set from compute services
 * @param operatorSet The operator set to deregister
 * @dev The caller must have permission to call on behalf of the operatorSet.avs through the PermissionController
 * @dev Reverts for:
 *      - InvalidPermissions: Caller does not have permission to call on behalf of operatorSet.avs
 *      - InvalidOperatorSet: The operator set does not exist in the AllocationManager
 *      - OperatorSetNotRegistered: The operator set is not registered for compute
 * @dev Emits the following events:
 *      - OperatorSetDeregistered: When the operator set is successfully deregistered
 */
function deregisterFromCompute(OperatorSet memory operatorSet) external;
```

Removes an operator set from compute services registration.

*Effects*:
* Removes the operator set's registration status
* Clears the stored TOS signature
* Emits an `OperatorSetDeregistered` event

*Requirements*:
* Caller MUST have permission to call on behalf of the operator set's AVS (via PermissionController)
* Operator set MUST exist in the AllocationManager
* Operator set MUST be currently registered for compute

*Note*: After deregistration, an operator set can re-register by signing a new TOS agreement and meeting all registration requirements again.

---

## View Functions

### Registration Status

#### `isOperatorSetRegistered`

```solidity
function isOperatorSetRegistered(bytes32 operatorSetKey) external view returns (bool);
```

Checks if an operator set is registered for compute services.

*Parameters*:
* `operatorSetKey`: The key of the operator set (computed using `OperatorSetLib.key()`)

*Returns*:
* `true` if the operator set is registered, `false` otherwise

#### `getOperatorSetTosSignature`

```solidity
function getOperatorSetTosSignature(OperatorSet memory operatorSet) external view returns (TOSSignature memory);
```

Retrieves the stored TOS signature for a registered operator set.

*Parameters*:
* `operatorSet`: The operator set to query

*Returns*:
* The TOS signature struct containing signer, TOS hash, and signature bytes
* Returns an empty struct if the operator set is not registered

### TOS Agreement Helpers

#### `calculateTOSAgreementDigest`

```solidity
function calculateTOSAgreementDigest(
    OperatorSet memory operatorSet,
    address signer
) external view returns (bytes32);
```

Calculates the EIP-712 digest hash that should be signed for TOS agreement.

*Parameters*:
* `operatorSet`: The operator set that is agreeing to the TOS
* `signer`: The address that will sign the agreement

*Returns*:
* The EIP-712 digest hash ready for signing

*Usage*: AVSs can use this function to generate the correct digest for signing the TOS agreement off-chain.

### Contract References

The following immutable contract references are available:

* `RELEASE_MANAGER()`: Returns the ReleaseManager contract address
* `ALLOCATION_MANAGER()`: Returns the AllocationManager contract address
* `KEY_REGISTRAR()`: Returns the KeyRegistrar contract address
* `CROSS_CHAIN_REGISTRY()`: Returns the CrossChainRegistry contract address
* `TOS_HASH()`: Returns the immutable TOS hash
* `MAX_EXPIRY()`: Returns the maximum expiry value (type(uint256).max)
* `TOS_AGREEMENT_TYPEHASH()`: Returns the EIP-712 type hash for TOS agreements

---

## Error Codes

| Error | Code | Description |
| ----- | ---- | ----------- |
| `InvalidTOSSignature` | `0x04bf729c` | Invalid or mismatched TOS signature |
| `OperatorSetAlreadyRegistered` | `0x1503562a` | Operator set is already registered |
| `OperatorSetNotRegistered` | `0x3a2e3ac6` | Operator set is not registered |
| `InvalidOperatorSet` | `0x7ec5c154` | Operator set doesn't exist in AllocationManager |
| `CurveTypeNotSet` | `0x3104b8e7` | Curve type not configured in KeyRegistrar |
| `NoActiveGenerationReservation` | `0xd0147d2d` | No active generation reservation in CrossChainRegistry |

---

## Events

### `OperatorSetRegistered`

```solidity
event OperatorSetRegistered(
    OperatorSet indexed operatorSet,
    address indexed signer,
    bytes32 indexed tosHash,
    bytes signature
);
```

Emitted when an operator set successfully registers for compute services.

### `OperatorSetDeregistered`

```solidity
event OperatorSetDeregistered(OperatorSet indexed operatorSet);
```

Emitted when an operator set is deregistered from compute services.

---

## Usage Patterns

### Registering for Compute Services

1. **Prepare the Operator Set**:
   - Ensure the operator set exists in the AllocationManager
   - Configure the curve type in the KeyRegistrar
   - Establish an active generation reservation in the CrossChainRegistry
   - Deploy at least one release in the ReleaseManager

2. **Generate TOS Signature**:
   ```solidity
   // Calculate the digest to sign
   bytes32 digest = computeRegistry.calculateTOSAgreementDigest(operatorSet, signer);
   
   // Sign the digest off-chain (using EIP-712)
   bytes memory signature = signEIP712(digest, signerPrivateKey);
   ```

3. **Register**:
   ```solidity
   // Caller must have permission via PermissionController
   computeRegistry.registerForCompute(operatorSet, signature);
   ```

### Deregistering

```solidity
// Caller must have permission via PermissionController
computeRegistry.deregisterFromCompute(operatorSet);
```

---

## Security Considerations

1. **Permission Control**: All registration and deregistration actions require proper authorization through the PermissionController. AVSs should carefully manage who has permission to register/deregister their operator sets.

2. **Signature Validation**: The contract uses EIP-712 for secure signature verification. Signatures are bound to the specific operator set and cannot be reused for different operator sets.

3. **Immutable TOS**: The Terms of Service hash is immutable once deployed, ensuring consistent agreement terms for all registrations.

4. **Dependency Validation**: The contract validates dependencies (KeyRegistrar, CrossChainRegistry, ReleaseManager) to ensure operator sets are fully configured before allowing compute registration.

5. **Re-registration**: After deregistration, operator sets can re-register, but must provide a new signature and meet all requirements again.

---