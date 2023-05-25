// SPDX-License-Identifier: BUSL-1.1

pragma solidity =0.8.12;

/**
 * @title Library for converting between an array of bytes and a bitmap.
 * @author Layr Labs, Inc.
 */
library BytesArrayBitmaps {
    /**
     * @notice Converts an array of bytes into a bitmap.
     * @param bytesArray The array of bytes to convert/compress into a bitmap.
     * @return The resulting bitmap.
     * @dev Each byte in the input is processed as indicating a single bit to flip in the bitmap
     */
    function bytesArrayToBitmap(bytes calldata bytesArray) internal pure returns (uint256) {
        // sanity-check on input. a too-long input would fail later on due to having duplicate entry(s)
        require(bytesArray.length <= 256, "BytesArrayBitmaps.bytesArrayToBitmap: bytesArray is too long");
        // initialize the empty bitmap, to be built inside the loop
        uint256 bitmap;
        // initialize an empty byte object, to be re-used inside the loop
        bytes1 singleByte;
        // initialize an empty uint256 to be used as a bitmask inside the loop
        uint256 bitMask;
        // loop through each byte in the array to construct the bitmap
        for (uint256 i = 0; i < bytesArray.length; ++i) {
            // pull the next byte out of the array
            singleByte = bytesArray[i];
            // construct a single-bit mask from the numerical value of the byte
            bitMask = uint256(1 << uint8(singleByte));
            // check that the entry is not a repeat
            require(bitmap & bitMask == 0, "BytesArrayBitmaps.bytesArrayToBitmap: repeat entry in bytesArray");
            // add the entry to the bitmap
            bitmap = (bitmap | bitMask);
        }
        return bitmap;
    }

    /**
     * @notice Converts a bitmap into an array of bytes.
     * @param bitmap The bitmap to decompress/convert to an array of bytes.
     * @return bytesArray The resulting bitmap array of bytes.
     * @dev Each byte in the input is processed as indicating a single bit to flip in the bitmap
     */
    function bitmapToBytesArray(uint256 bitmap) internal pure returns (bytes memory bytesArray) {
        // initialize an empty uint256 to be used as a bitmask inside the loop
        uint256 bitMask;
        // loop through each index in the bitmap to construct the array
        for (uint256 i = 0; i < 256; ++i) {
            // construct a single-bit mask for the i-th bit
            bitMask = uint256(1 << i);
            // check if the i-th bit is flipped in the bitmap
            if (bitmap & bitMask != 0) {
                // if the i-th bit is flipped, then add a byte encoding the value 'i' to the `bytesArray`
                bytesArray = bytes.concat(bytesArray, bytes1(uint8(i)));
            }
        }
        return bytesArray;
    }
}
