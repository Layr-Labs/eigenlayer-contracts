# Solidity API

## BLSPublicKeyCompendium

### ZERO_PK_HASH

```solidity
bytes32 ZERO_PK_HASH
```

### operatorToPubkeyHash

```solidity
mapping(address => bytes32) operatorToPubkeyHash
```

mapping from operator address to pubkey hash

### pubkeyHashToOperator

```solidity
mapping(bytes32 => address) pubkeyHashToOperator
```

mapping from pubkey hash to operator address

### NewPubkeyRegistration

```solidity
event NewPubkeyRegistration(address operator, struct BN254.G1Point pubkeyG1, struct BN254.G2Point pubkeyG2)
```

Emitted when `operator` registers with the public key `pk`.

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

