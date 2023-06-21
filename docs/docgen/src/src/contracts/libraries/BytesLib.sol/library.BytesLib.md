# BytesLib
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/libraries/BytesLib.sol)


## Functions
### concat


```solidity
function concat(bytes memory _preBytes, bytes memory _postBytes) internal pure returns (bytes memory);
```

### concatStorage


```solidity
function concatStorage(bytes storage _preBytes, bytes memory _postBytes) internal;
```

### slice


```solidity
function slice(bytes memory _bytes, uint256 _start, uint256 _length) internal pure returns (bytes memory);
```

### toAddress


```solidity
function toAddress(bytes memory _bytes, uint256 _start) internal pure returns (address);
```

### toUint8


```solidity
function toUint8(bytes memory _bytes, uint256 _start) internal pure returns (uint8);
```

### toUint16


```solidity
function toUint16(bytes memory _bytes, uint256 _start) internal pure returns (uint16);
```

### toUint32


```solidity
function toUint32(bytes memory _bytes, uint256 _start) internal pure returns (uint32);
```

### toUint64


```solidity
function toUint64(bytes memory _bytes, uint256 _start) internal pure returns (uint64);
```

### toUint96


```solidity
function toUint96(bytes memory _bytes, uint256 _start) internal pure returns (uint96);
```

### toUint128


```solidity
function toUint128(bytes memory _bytes, uint256 _start) internal pure returns (uint128);
```

### toUint256


```solidity
function toUint256(bytes memory _bytes, uint256 _start) internal pure returns (uint256);
```

### toBytes32


```solidity
function toBytes32(bytes memory _bytes, uint256 _start) internal pure returns (bytes32);
```

### equal


```solidity
function equal(bytes memory _preBytes, bytes memory _postBytes) internal pure returns (bool);
```

### equalStorage


```solidity
function equalStorage(bytes storage _preBytes, bytes memory _postBytes) internal view returns (bool);
```

