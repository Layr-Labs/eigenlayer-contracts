# Endian
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/libraries/Endian.sol)


## Functions
### fromLittleEndianUint64

Converts a little endian-formatted uint64 to a big endian-formatted uint64

*Note that the input is formatted as a 'bytes32' type (i.e. 256 bits), but it is immediately truncated to a uint64 (i.e. 64 bits)
through a right-shift/shr operation.*


```solidity
function fromLittleEndianUint64(bytes32 lenum) internal pure returns (uint64 n);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`lenum`|`bytes32`|little endian-formatted uint64 input, provided as 'bytes32' type|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`n`|`uint64`|The big endian-formatted uint64|


