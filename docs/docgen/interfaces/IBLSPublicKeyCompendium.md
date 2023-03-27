# Solidity API

## IBLSPublicKeyCompendium

### operatorToPubkeyHash

```solidity
function operatorToPubkeyHash(address operator) external view returns (bytes32)
```

mapping from operator address to pubkey hash.
Returns *zero* if the `operator` has never registered, and otherwise returns the hash of the public key of the operator.

### pubkeyHashToOperator

```solidity
function pubkeyHashToOperator(bytes32 pubkeyHash) external view returns (address)
```

mapping from pubkey hash to operator address.
Returns *zero* if no operator has ever registered the public key corresponding to `pubkeyHash`,
and otherwise returns the (unique) registered operator who owns the BLS public key that is the preimage of `pubkeyHash`.

### registerBLSPublicKey

```solidity
function registerBLSPublicKey(uint256 s, struct BN254.G1Point rPoint, struct BN254.G1Point pubkeyG1, struct BN254.G2Point pubkeyG2) external
```

Called by an operator to register themselves as the owner of a BLS public key and reveal their G1 and G2 public key.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| s | uint256 | is the field element of the operator's Schnorr signature |
| rPoint | struct BN254.G1Point | is the group element of the operator's Schnorr signature |
| pubkeyG1 | struct BN254.G1Point | is the the G1 pubkey of the operator |
| pubkeyG2 | struct BN254.G2Point | is the G2 with the same private key as the pubkeyG1 |

