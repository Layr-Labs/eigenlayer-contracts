# BN254
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/libraries/BN254.sol)

**Author:**
Layr Labs, Inc.

Contains BN254 parameters, common operations (addition, scalar mul, pairing), and BLS signature functionality.


## State Variables
### FP_MODULUS

```solidity
uint256 internal constant FP_MODULUS = 21888242871839275222246405745257275088696311157297823662689037894645226208583;
```


### FR_MODULUS

```solidity
uint256 internal constant FR_MODULUS = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
```


### G2x1
*Generator point in F_q2 is of the form: (x0 + ix1, y0 + iy1).*


```solidity
uint256 internal constant G2x1 = 11559732032986387107991004021392285783925812861821192530917403151452391805634;
```


### G2x0

```solidity
uint256 internal constant G2x0 = 10857046999023057135944570762232829481370756359578518086990519993285655852781;
```


### G2y1

```solidity
uint256 internal constant G2y1 = 4082367875863433681332203403145435568316851327593401208105741076214120093531;
```


### G2y0

```solidity
uint256 internal constant G2y0 = 8495653923123431417604973247489272438418190587263600148770280649306958101930;
```


### nG2x1
*Generator point in F_q2 is of the form: (x0 + ix1, y0 + iy1).*


```solidity
uint256 internal constant nG2x1 = 11559732032986387107991004021392285783925812861821192530917403151452391805634;
```


### nG2x0

```solidity
uint256 internal constant nG2x0 = 10857046999023057135944570762232829481370756359578518086990519993285655852781;
```


### nG2y1

```solidity
uint256 internal constant nG2y1 = 17805874995975841540914202342111839520379459829704422454583296818431106115052;
```


### nG2y0

```solidity
uint256 internal constant nG2y0 = 13392588948715843804641432497768002650278120570034223513918757245338268106653;
```


### powersOfTauMerkleRoot

```solidity
bytes32 internal constant powersOfTauMerkleRoot = 0x22c998e49752bbb1918ba87d6d59dd0e83620a311ba91dd4b2cc84990b31b56f;
```


## Functions
### generatorG2

returns the G2 generator

*mind the ordering of the 1s and 0s!
this is because of the (unknown to us) convention used in the bn254 pairing precompile contract
"Elements a * i + b of F_p^2 are encoded as two elements of F_p, (a, b)."
https://github.com/ethereum/EIPs/blob/master/EIPS/eip-197.md#encoding*


```solidity
function generatorG2() internal pure returns (G2Point memory);
```

### negGeneratorG2


```solidity
function negGeneratorG2() internal pure returns (G2Point memory);
```

### negate


```solidity
function negate(G1Point memory p) internal pure returns (G1Point memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`p`|`G1Point`|Some point in G1.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`G1Point`|The negation of `p`, i.e. p.plus(p.negate()) should be zero.|


### plus


```solidity
function plus(G1Point memory p1, G1Point memory p2) internal view returns (G1Point memory r);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`r`|`G1Point`|the sum of two points of G1|


### scalar_mul


```solidity
function scalar_mul(G1Point memory p, uint256 s) internal view returns (G1Point memory r);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`r`|`G1Point`|the product of a point on G1 and a scalar, i.e. p == p.scalar_mul(1) and p.plus(p) == p.scalar_mul(2) for all points p.|


### pairing


```solidity
function pairing(G1Point memory a1, G2Point memory a2, G1Point memory b1, G2Point memory b2)
    internal
    view
    returns (bool);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|The result of computing the pairing check e(p1[0], p2[0]) *  .... * e(p1[n], p2[n]) == 1 For example, pairing([P1(), P1().negate()], [P2(), P2()]) should return true.|


### safePairing

This function is functionally the same as pairing(), however it specifies a gas limit
the user can set, as a precompile may use the entire gas budget if it reverts.


```solidity
function safePairing(G1Point memory a1, G2Point memory a2, G1Point memory b1, G2Point memory b2, uint256 pairingGas)
    internal
    view
    returns (bool, bool);
```

### hashG1Point

*used for BLS signatures*


```solidity
function hashG1Point(BN254.G1Point memory pk) internal pure returns (bytes32);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes32`|the keccak256 hash of the G1 Point|


### hashToG1

adapted from https://github.com/HarryR/solcrypto/blob/master/contracts/altbn128.sol


```solidity
function hashToG1(bytes32 _x) internal view returns (uint256, uint256);
```

### findYFromX

Given X, find Y
where y = sqrt(x^3 + b)
Returns: (x^3 + b), y


```solidity
function findYFromX(uint256 x) internal view returns (uint256, uint256);
```

### expMod


```solidity
function expMod(uint256 _base, uint256 _exponent, uint256 _modulus) internal view returns (uint256 retval);
```

## Structs
### G1Point

```solidity
struct G1Point {
    uint256 X;
    uint256 Y;
}
```

### G2Point

```solidity
struct G2Point {
    uint256[2] X;
    uint256[2] Y;
}
```

