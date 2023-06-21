# BytesArrayBitmapsUnitTests
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/unit/BytesArrayBitmapsUnit.t.sol)

**Inherits:**
Test


## State Variables
### cheats

```solidity
Vm cheats = Vm(HEVM_ADDRESS);
```


### bytesArrayBitmapsWrapper

```solidity
BytesArrayBitmapsWrapper public bytesArrayBitmapsWrapper;
```


## Functions
### setUp


```solidity
function setUp() public;
```

### testEmptyArrayEncoding


```solidity
function testEmptyArrayEncoding() public view;
```

### testSingleByteEncoding


```solidity
function testSingleByteEncoding(uint8 fuzzedNumber) public view;
```

### testTwoByteEncoding


```solidity
function testTwoByteEncoding(uint8 firstFuzzedNumber, uint8 secondFuzzedNumber) public;
```

### testBytesArrayToBitmapToBytesArray


```solidity
function testBytesArrayToBitmapToBytesArray(bytes memory originalBytesArray) public view;
```

### testBytesArrayToBitmapToBytesArray_Yul


```solidity
function testBytesArrayToBitmapToBytesArray_Yul(bytes memory originalBytesArray) public view;
```

### testBytesArrayToBitmapToBytesArray_OrderedVersion


```solidity
function testBytesArrayToBitmapToBytesArray_OrderedVersion(bytes memory originalBytesArray) public view;
```

### testBytesArrayToBitmapToBytesArray_OrderedVersion_Yul


```solidity
function testBytesArrayToBitmapToBytesArray_OrderedVersion_Yul(bytes memory originalBytesArray) public view;
```

### testBitMapToBytesArrayToBitmap


```solidity
function testBitMapToBytesArrayToBitmap(uint256 originalBitmap) public view;
```

### testBytesArrayToBitmap_OrderedVersion_Yul_SpecificInput


```solidity
function testBytesArrayToBitmap_OrderedVersion_Yul_SpecificInput() public;
```

### testBytesArrayToBitmap_OrderedVersion_SpecificInput


```solidity
function testBytesArrayToBitmap_OrderedVersion_SpecificInput() public;
```

### testBytesArrayToBitmap_SpecificInput


```solidity
function testBytesArrayToBitmap_SpecificInput() public;
```

### testBytesArrayToBitmap_Yul_SpecificInput


```solidity
function testBytesArrayToBitmap_Yul_SpecificInput() public;
```

