# SignatureCompaction
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/utils/SignatureCompaction.sol)


## State Variables
### HALF_CURVE_ORDER

```solidity
bytes32 internal constant HALF_CURVE_ORDER = 0x7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0;
```


## Functions
### ecrecoverPacked


```solidity
function ecrecoverPacked(bytes32 hash, bytes32 r, bytes32 vs) internal pure returns (address);
```

### packSignature


```solidity
function packSignature(bytes32 r, bytes32 s, uint8 v) internal pure returns (bytes32, bytes32);
```

### packVS


```solidity
function packVS(bytes32 s, uint8 v) internal pure returns (bytes32);
```

