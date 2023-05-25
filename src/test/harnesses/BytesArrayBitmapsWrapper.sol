// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/libraries/BytesArrayBitmaps.sol";

// wrapper around the BytesArrayBitmaps library that exposes the internal functions
contract BytesArrayBitmapsWrapper {
    function bytesArrayToBitmap(bytes calldata bytesArray) external pure returns (uint256) {
        return BytesArrayBitmaps.bytesArrayToBitmap(bytesArray);
    }

    function bitmapToBytesArray(uint256 bitmap) external pure returns (bytes memory bytesArray) {
        return BytesArrayBitmaps.bitmapToBytesArray(bitmap);
    }
}