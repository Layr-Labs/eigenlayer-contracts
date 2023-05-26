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

    // function orderedBytesArrayToBitmap_Yul(bytes calldata orderedBytesArray) external pure returns (uint256) {
    //     return BytesArrayBitmaps.orderedBytesArrayToBitmap_Yul(orderedBytesArray);
    // }

    /**
     * @notice Converts an ordered array of bytes into a bitmap.
     * @param orderedBytesArray The array of bytes to convert/compress into a bitmap. Must be in strictly ascending order.
     * @return The resulting bitmap.
     * @dev Each byte in the input is processed as indicating a single bit to flip in the bitmap.
     * @dev This function will eventually revert in the event that the `orderedBytesArray` is not properly ordered (in ascending order).
     */
    function orderedBytesArrayToBitmap_Yul(bytes calldata orderedBytesArray) external pure returns (uint256) {
        // sanity-check on input. a too-long input would fail later on due to having duplicate entry(s)
        require(orderedBytesArray.length <= 256, "BytesArrayBitmaps.orderedBytesArrayToBitmap: orderedBytesArray is too long");
        // return empty bitmap early if length of array is 0
        if (orderedBytesArray.length == 0) {
            return uint256(0);
        }

        // NOTE: routine 6 has the best performance

        // // BEGIN WORKING ROUTINE 1
        // uint256 bitmap;
        // uint256 startReadingAt;
        // uint256 previousByte;
        // uint256 currentByte;

        // assembly {
        //     startReadingAt := sub(orderedBytesArray.offset, 31)
        //     previousByte := 
        //         and(
        //             calldataload(
        //                 startReadingAt
        //             ),
        //             0x00000000000000000000000000000000000000000000000000000000000000FF
        //         )
        //     bitmap :=
        //         or(bitmap,
        //             shl(previousByte,
        //                 1
        //             )
        //         )
        // }

        // // TODO: add custom revert message in Yul, rather than using `revert(0, 0)`
        // for (uint256 i = 1; i < orderedBytesArray.length;) {
        //     assembly {
        //         currentByte := 
        //             and(
        //                 calldataload(
        //                     add(
        //                         startReadingAt,
        //                         i
        //                     )
        //                 ),
        //                 0x00000000000000000000000000000000000000000000000000000000000000FF
        //             )
        //         bitmap :=
        //             or(bitmap,
        //                 shl(currentByte,
        //                     1
        //                 )
        //             )
        //         if iszero(gt(currentByte, previousByte)) {revert(0, 0)}
        //         previousByte := currentByte
        //         i := add(i, 1)              
        //     }
        // }
        // // END WORKING ROUTINE 1

        // // BEGIN WORKING ROUTINE 2
        // uint256 bitmap;

        // assembly {                
        //     bitmap :=
        //         shl(
        //             and(
        //                 calldataload(
        //                     sub(orderedBytesArray.offset, 31)
        //                 ),
        //                 0x00000000000000000000000000000000000000000000000000000000000000FF
        //             ),
        //             1
        //         )
        // }

        // // TODO: add custom revert message in Yul, rather than using `revert(0, 0)`
        // for (uint256 i = 1; i < orderedBytesArray.length;) {
        //     assembly {                    
        //         bitmap :=
        //             or(
        //                 bitmap,
        //                 shl(
        //                     and(
        //                         calldataload(
        //                             add(
        //                                 sub(orderedBytesArray.offset, 31),
        //                                 i
        //                             )
        //                         ),
        //                         0x00000000000000000000000000000000000000000000000000000000000000FF
        //                     ),
        //                     1
        //                 )
        //             )
        //         if iszero(gt(
        //                 and(
        //                     calldataload(
        //                         add(
        //                             sub(orderedBytesArray.offset, 31),
        //                             i
        //                         )
        //                     ),
        //                     0x00000000000000000000000000000000000000000000000000000000000000FF
        //                 ),
        //                 and(
        //                     calldataload(
        //                         sub(orderedBytesArray.offset, 31)
        //                     ),
        //                     0x00000000000000000000000000000000000000000000000000000000000000FF
        //                 )
        //             )) {revert(0, 0)}
        //         i := add(i, 1)              
        //     }
        // }
        // // END WORKING ROUTINE 2

        // // BEGIN WORKING ROUTINE 3
        // uint256 bitmap;
        // uint256 startReadingAt;
        // uint256 previousByte;
        // uint256 currentByte;

        // assembly {
        //     startReadingAt := sub(orderedBytesArray.offset, 31)

        //     previousByte := 
        //         and(
        //             calldataload(
        //                 startReadingAt
        //             ),
        //             0x00000000000000000000000000000000000000000000000000000000000000FF
        //         )

        //     bitmap :=
        //         shl(
        //             previousByte,
        //             1
        //         )
        // }

        // // TODO: add custom revert message in Yul, rather than using `revert(0, 0)`
        // for (uint256 i = 1; i < orderedBytesArray.length;) {
        //     assembly {
        //         currentByte := 
        //             and(
        //                 calldataload(
        //                     add(
        //                         startReadingAt,
        //                         i
        //                     )
        //                 ),
        //                 0x00000000000000000000000000000000000000000000000000000000000000FF
        //             )

        //         bitmap :=
        //             or(bitmap,
        //                 shl(currentByte,
        //                     1
        //                 )
        //             )

        //         if iszero(gt(currentByte, previousByte)) {revert(0, 0)}

        //         previousByte := currentByte

        //         i := add(i, 1)              
        //     }
        // }
        // // END WORKING ROUTINE 3

        // // BEGIN WORKING ROUTINE 4
        // uint256 bitmap;
        // uint256 startReadingAt;
        // uint256 previousByte;
        // uint256 currentByte;

        // // TODO: add custom revert message in Yul, rather than using `revert(0, 0)`
        // assembly {
        //     startReadingAt := sub(orderedBytesArray.offset, 31)

        //     previousByte := 
        //         and(
        //             calldataload(
        //                 startReadingAt
        //             ),
        //             0x00000000000000000000000000000000000000000000000000000000000000FF
        //         )

        //     bitmap :=
        //         shl(
        //             previousByte,
        //             1
        //         )

        //     for { let i := 1 } lt(i, orderedBytesArray.length) { i := add(i, 1) } {
        //         currentByte := 
        //             and(
        //                 calldataload(
        //                     add(
        //                         startReadingAt,
        //                         i
        //                     )
        //                 ),
        //                 0x00000000000000000000000000000000000000000000000000000000000000FF
        //             )

        //         bitmap :=
        //             or(bitmap,
        //                 shl(currentByte,
        //                     1
        //                 )
        //             )

        //         if iszero(gt(currentByte, previousByte)) {revert(0, 0)}

        //         previousByte := currentByte
        //     }
        // }
        // // END WORKING ROUTINE 4

        // // BEGIN WORKING ROUTINE 5
        // uint256 bitmap;
        // uint256 previousByte;
        // uint256 currentByte;

        // assembly {
        //     previousByte := 
        //         shr(
        //             248,
        //             calldataload(
        //                 orderedBytesArray.offset
        //             )
        //         )
        //     bitmap :=
        //         shl(previousByte,
        //             1
        //         )
        // }

        // // emit log_named_bytes("orderedBytesArray", orderedBytesArray);
        // // emit log_named_bytes32("bytes32(previousByte)", bytes32(previousByte));
        // // emit log_named_uint("bitmap", bitmap);

        // // TODO: add custom revert message in Yul, rather than using `revert(0, 0)`
        // for (uint256 i = 1; i < orderedBytesArray.length;) {
        //     assembly {
        //         currentByte := 
        //             shr(
        //                 248,
        //                 calldataload(
        //                     add(
        //                         orderedBytesArray.offset,
        //                         i
        //                     )
        //                 )
        //             )
        //         bitmap :=
        //             or(bitmap,
        //                 shl(currentByte,
        //                     1
        //                 )
        //             )
        //         if iszero(gt(currentByte, previousByte)) {revert(0, 0)}
        //         previousByte := currentByte
        //         i := add(i, 1)              
        //     }
        // }
        // // END WORKING ROUTINE 5


        // // BEGIN WORKING ROUTINE 6
        // uint256 bitmap;
        // uint256 mask;

        // assembly {
        //     bitmap :=
        //         shl(
        //             shr(
        //                 248,
        //                 calldataload(
        //                     orderedBytesArray.offset
        //                 )
        //             ),
        //             1
        //         )
        // }

        // // TODO: add custom revert message in Yul, rather than using `revert(0, 0)`
        // for (uint256 i = 1; i < orderedBytesArray.length;) {
        //     assembly {
        //         mask := 
        //             shl(
        //                 shr(
        //                     248,
        //                     calldataload(
        //                         add(
        //                             orderedBytesArray.offset,
        //                             i
        //                         )
        //                     )
        //                 ),
        //                 1
        //             )
        //         if iszero(gt(mask, bitmap)) {revert(0, 0)}
        //         bitmap := or(bitmap, mask)
        //         i := add(i, 1)              
        //     }
        // }
        // // END WORKING ROUTINE 6

        // BEGIN WORKING ROUTINE 7
        // TODO: add custom revert message in Yul, rather than using `revert(0, 0)`
        // uint256 bitmap;

        assembly {
            // get first entry in bitmap (single byte => single-bit mask)
            let bitmap :=
                shl(
                    // pull out single byte to get the correct value for the left shift
                    shr(
                        248,
                        calldataload(
                            orderedBytesArray.offset
                        )
                    ),
                    1
                )
            // loop through other entries (byte by byte)
            for { let i := 1 } lt(i, orderedBytesArray.length) { i := add(i, 1) } {
                // first construct the single-bit mask by left-shifting a '1'
                let mask := 
                    shl(
                        // pull out single byte to get the correct value for the left shift
                        shr(
                            248,
                            calldataload(
                                add(
                                    orderedBytesArray.offset,
                                    i
                                )
                            )
                        ),
                        1
                    )
                // check strictly ascending ordering by comparing the mask to the bitmap so far (revert if mask isn't greater than bitmap)
                if iszero(gt(mask, bitmap)) {revert(0, 0)}
                // update the bitmap by adding the single bit in the mask
                bitmap := or(bitmap, mask)
            }
            // after the loop is complete, store the bitmap at the value encoded at the free memory pointer, then return it
            mstore(mload(0x40), bitmap)
            return(mload(0x40), 32)
        }
        // END WORKING ROUTINE 7

        // return bitmap;
    }

}