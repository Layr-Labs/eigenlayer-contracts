// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/libraries/BitmapUtils.sol";

import "forge-std/Test.sol";

// wrapper around the BitmapUtils library that exposes the internal functions
contract BitmapUtilsWrapper is Test {
    function bytesArrayToBitmap(bytes calldata bytesArray) external pure returns (uint256) {
        return BitmapUtils.bytesArrayToBitmap(bytesArray);
    }

    function orderedBytesArrayToBitmap(bytes calldata orderedBytesArray) external pure returns (uint256) {
        return BitmapUtils.orderedBytesArrayToBitmap(orderedBytesArray);
    }

    function isArrayStrictlyAscendingOrdered(bytes calldata bytesArray) external pure returns (bool) {
        return BitmapUtils.isArrayStrictlyAscendingOrdered(bytesArray);
    }

    function bitmapToBytesArray(uint256 bitmap) external pure returns (bytes memory bytesArray) {
        return BitmapUtils.bitmapToBytesArray(bitmap);
    }

    function orderedBytesArrayToBitmap_Yul(bytes calldata orderedBytesArray) external pure returns (uint256) {
        return BitmapUtils.orderedBytesArrayToBitmap_Yul(orderedBytesArray);
    }

    function bytesArrayToBitmap_Yul(bytes calldata bytesArray) external pure returns (uint256) {
        return BitmapUtils.bytesArrayToBitmap_Yul(bytesArray);
    }
}