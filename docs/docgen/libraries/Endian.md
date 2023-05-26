# Solidity API

## Endian

### fromLittleEndianUint64

```solidity
function fromLittleEndianUint64(bytes32 lenum) internal pure returns (uint64 n)
```

Converts a little endian-formatted uint64 to a big endian-formatted uint64

_Note that the input is formatted as a 'bytes32' type (i.e. 256 bits), but it is immediately truncated to a uint64 (i.e. 64 bits)
through a right-shift/shr operation._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| lenum | bytes32 | little endian-formatted uint64 input, provided as 'bytes32' type |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| n | uint64 | The big endian-formatted uint64 |

