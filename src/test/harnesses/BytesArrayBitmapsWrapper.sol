// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/libraries/BytesArrayBitmaps.sol";

import "forge-std/Test.sol";

// wrapper around the BytesArrayBitmaps library that exposes the internal functions
contract BytesArrayBitmapsWrapper is Test {
    function bytesArrayToBitmap(bytes calldata bytesArray) external pure returns (uint256) {
        return BytesArrayBitmaps.bytesArrayToBitmap(bytesArray);
    }

    function orderedBytesArrayToBitmap(bytes calldata orderedBytesArray) external pure returns (uint256) {
        return BytesArrayBitmaps.orderedBytesArrayToBitmap(orderedBytesArray);
    }

    function isArrayStrictlyAscendingOrdered(bytes calldata bytesArray) external pure returns (bool) {
        return BytesArrayBitmaps.isArrayStrictlyAscendingOrdered(bytesArray);
    }

    function bitmapToBytesArray(uint256 bitmap) external pure returns (bytes memory bytesArray) {
        return BytesArrayBitmaps.bitmapToBytesArray(bitmap);
    }

    function orderedBytesArrayToBitmap_Yul(bytes calldata orderedBytesArray) external pure returns (uint256) {
        return BytesArrayBitmaps.orderedBytesArrayToBitmap_Yul(orderedBytesArray);
    }

    function bytesArrayToBitmap_Yul(bytes calldata bytesArray) external pure returns (uint256) {
        return BytesArrayBitmaps.bytesArrayToBitmap_Yul(bytesArray);
    }
}