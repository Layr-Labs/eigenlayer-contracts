# IBLSPublicKeyCompendium
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/interfaces/IBLSPublicKeyCompendium.sol)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service


## Functions
### operatorToPubkeyHash

mapping from operator address to pubkey hash.
Returns *zero* if the `operator` has never registered, and otherwise returns the hash of the public key of the operator.


```solidity
function operatorToPubkeyHash(address operator) external view returns (bytes32);
```

### pubkeyHashToOperator

mapping from pubkey hash to operator address.
Returns *zero* if no operator has ever registered the public key corresponding to `pubkeyHash`,
and otherwise returns the (unique) registered operator who owns the BLS public key that is the preimage of `pubkeyHash`.


```solidity
function pubkeyHashToOperator(bytes32 pubkeyHash) external view returns (address);
```

### registerBLSPublicKey

Called by an operator to register themselves as the owner of a BLS public key and reveal their G1 and G2 public key.


```solidity
function registerBLSPublicKey(
    uint256 s,
    BN254.G1Point memory rPoint,
    BN254.G1Point memory pubkeyG1,
    BN254.G2Point memory pubkeyG2
) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`s`|`uint256`|is the field element of the operator's Schnorr signature|
|`rPoint`|`BN254.G1Point`|is the group element of the operator's Schnorr signature|
|`pubkeyG1`|`BN254.G1Point`|is the the G1 pubkey of the operator|
|`pubkeyG2`|`BN254.G2Point`|is the G2 with the same private key as the pubkeyG1|


