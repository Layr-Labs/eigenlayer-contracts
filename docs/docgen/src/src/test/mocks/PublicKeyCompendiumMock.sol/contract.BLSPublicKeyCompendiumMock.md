# BLSPublicKeyCompendiumMock
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/PublicKeyCompendiumMock.sol)

**Inherits:**
[IBLSPublicKeyCompendium](/docs/docgen/src/src/contracts/interfaces/IBLSPublicKeyCompendium.sol/interface.IBLSPublicKeyCompendium.md), DSTest

**Author:**
Layr Labs, Inc.


## State Variables
### operatorToPubkeyHash
mapping from operator address to pubkey hash


```solidity
mapping(address => bytes32) public operatorToPubkeyHash;
```


### pubkeyHashToOperator
mapping from pubkey hash to operator address


```solidity
mapping(bytes32 => address) public pubkeyHashToOperator;
```


## Functions
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


### registerPublicKey


```solidity
function registerPublicKey(BN254.G1Point memory pk) external;
```

