[middleware-repo]: https://github.com/Layr-Labs/eigenlayer-middleware/

## AVSDirectory

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`AVSDirectory.sol`](../../src/contracts/core/AVSDirectory.sol) | Singleton | Transparent proxy |

*Note: The AVSDirectory is deprecated as of the slashing release. This documentation is kept for historical purposes. AVSs are recommended to use the [AllocationManager](./AllocationManager.md) for the new operatorSet and slashing model.*

The `AVSDirectory` once handled interactions between AVSs and the EigenLayer core contracts. Once registered as an Operator in EigenLayer core (via the `DelegationManager`), Operators can register with one or more AVSs (via the AVS's contracts) to begin providing services to them offchain. As a part of registering with an AVS, the AVS will record this registration in the core contracts by calling into the `AVSDirectory`.

For more information on AVS contracts, see the [middleware repo][middleware-repo].

Currently, the only interactions between AVSs and the core contracts is to track whether Operators are currently registered for the AVS. This is handled by two methods:
<!-- no toc -->
* [`AVSDirectory.registerOperatorToAVS`](#registeroperatortoavs)
* [`AVSDirectory.deregisterOperatorFromAVS`](#deregisteroperatorfromavs)

---

#### `registerOperatorToAVS`

```solidity
function registerOperatorToAVS(
    address operator,
    ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
) 
    external 
    onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS)
```

Allows the caller (an AVS) to register an `operator` with itself, given the provided signature is valid.

*Effects*:
* Sets the `operator`'s status to `REGISTERED` for the AVS

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS`
* `operator` MUST already be a registered Operator (via the `DelegationManager`)
* `operator` MUST NOT already be registered with the AVS
* `operatorSignature` must be a valid, unused, unexpired signature from the `operator`. The signature is an ECDSA signature by the operator over the [`OPERATOR_AVS_REGISTRATION_TYPEHASH`](../../src/contracts/core/DelegationManagerStorage.sol). Expiry is a utc timestamp in seconds. Salt is used only once per signature to prevent replay attacks.

*As of M2*:
* Operator registration/deregistration in the AVSDirectory does not have any sort of consequences for the Operator or its shares. The AllocationManager handles with operator sets, tying in rewards for services and slashing for misbehavior.

#### `deregisterOperatorFromAVS`

```solidity
function deregisterOperatorFromAVS(
    address operator
) 
    external 
    onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS)
```

Allows the caller (an AVS) to deregister an `operator` with itself

*Effects*:
* Sets the `operator's` status to `UNREGISTERED` for the AVS

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS`
* `operator` MUST already be registered with the AVS

*As of M2*:
* Operator registration/deregistration in the AVSDirectory does not have any sort of consequences for the Operator or its shares. The AllocationManager handles this with operator sets, tying in rewards for services and slashing for misbehavior.

#### `cancelSalt`

```solidity
function cancelSalt(bytes32 salt) external
```

Allows the caller (an Operator) to cancel a signature salt before it is used to register for an AVS.

*Effects*:
* Sets `operatorSaltIsSpent[msg.sender][salt]` to `true`

*Requirements*:
* Salt MUST NOT already be cancelled