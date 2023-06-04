# BytesArrayBitmaps
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/libraries/BytesArrayBitmaps.sol)

**Author:**
Layr Labs, Inc.


## State Variables
### MAX_BYTE_ARRAY_LENGTH
Byte arrays are meant to contain unique bytes.
If the array length exceeds 256, then it's impossible for all entries to be unique.
This constant captures the max allowed array length (inclusive, i.e. 256 is allowed).


```solidity
uint256 constant MAX_BYTE_ARRAY_LENGTH = 256;
```


## Functions
### bytesArrayToBitmap

Converts an array of bytes into a bitmap.

*Each byte in the input is processed as indicating a single bit to flip in the bitmap*

*This function will also revert if the `bytesArray` input contains any duplicate entries (i.e. duplicate bytes).*


```solidity
function bytesArrayToBitmap(bytes calldata bytesArray) internal pure returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`bytesArray`|`bytes`|The array of bytes to convert/compress into a bitmap.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|The resulting bitmap.|


### orderedBytesArrayToBitmap

Converts an ordered array of bytes into a bitmap.

*Each byte in the input is processed as indicating a single bit to flip in the bitmap.*

*This function will eventually revert in the event that the `orderedBytesArray` is not properly ordered (in ascending order).*

*This function will also revert if the `orderedBytesArray` input contains any duplicate entries (i.e. duplicate bytes).*


```solidity
function orderedBytesArrayToBitmap(bytes calldata orderedBytesArray) internal pure returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`orderedBytesArray`|`bytes`|The array of bytes to convert/compress into a bitmap. Must be in strictly ascending order.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|The resulting bitmap.|


### orderedBytesArrayToBitmap_Yul

Converts an ordered array of bytes into a bitmap. Optimized, Yul-heavy version of `orderedBytesArrayToBitmap`.

*Each byte in the input is processed as indicating a single bit to flip in the bitmap.*

*This function will eventually revert in the event that the `orderedBytesArray` is not properly ordered (in ascending order).*

*This function will also revert if the `orderedBytesArray` input contains any duplicate entries (i.e. duplicate bytes).*


```solidity
function orderedBytesArrayToBitmap_Yul(bytes calldata orderedBytesArray) internal pure returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`orderedBytesArray`|`bytes`|The array of bytes to convert/compress into a bitmap. Must be in strictly ascending order.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|The resulting bitmap.|


### bytesArrayToBitmap_Yul

Converts an array of bytes into a bitmap. Optimized, Yul-heavy version of `bytesArrayToBitmap`.

*Each byte in the input is processed as indicating a single bit to flip in the bitmap.*

*This function will eventually revert in the event that the `bytesArray` is not properly ordered (in ascending order).*

*This function will also revert if the `bytesArray` input contains any duplicate entries (i.e. duplicate bytes).*


```solidity
function bytesArrayToBitmap_Yul(bytes calldata bytesArray) internal pure returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`bytesArray`|`bytes`|The array of bytes to convert/compress into a bitmap.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|The resulting bitmap.|


### isArrayStrictlyAscendingOrdered

Utility function for checking if a bytes array is strictly ordered, in ascending order.

*This function returns 'true' for the edge case of the `bytesArray` having zero length.
It also returns 'false' early for arrays with length in excess of MAX_BYTE_ARRAY_LENGTH (i.e. so long that they cannot be strictly ordered)*


```solidity
function isArrayStrictlyAscendingOrdered(bytes calldata bytesArray) internal pure returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`bytesArray`|`bytes`|the bytes array of interest|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|Returns 'true' if the array is ordered in strictly ascending order, and 'false' otherwise.|


### bitmapToBytesArray

Converts a bitmap into an array of bytes.

*Each byte in the input is processed as indicating a single bit to flip in the bitmap*


```solidity
function bitmapToBytesArray(uint256 bitmap) internal pure returns (bytes memory bytesArray);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`bitmap`|`uint256`|The bitmap to decompress/convert to an array of bytes.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`bytesArray`|`bytes`|The resulting bitmap array of bytes.|


