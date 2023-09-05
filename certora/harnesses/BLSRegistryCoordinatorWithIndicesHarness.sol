// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../munged/middleware/BLSRegistryCoordinatorWithIndices.sol";

contract BLSRegistryCoordinatorWithIndicesHarness is BLSRegistryCoordinatorWithIndices {
    constructor(
        ISlasher _slasher,
        IServiceManager _serviceManager,
        IStakeRegistry _stakeRegistry,
        IBLSPubkeyRegistry _blsPubkeyRegistry,
        IIndexRegistry _indexRegistry
    )
        BLSRegistryCoordinatorWithIndices(_slasher, _serviceManager, _stakeRegistry, _blsPubkeyRegistry, _indexRegistry)
    {}

    // @notice function based upon `BitmapUtils.bytesArrayToBitmap`, used to determine if an array contains any duplicates
    function bytesArrayContainsDuplicates(bytes calldata bytesArray) public pure returns (bool) {
        uint256 bitmap;
        // sanity-check on input. a too-long input would fail later on due to having duplicate entry(s)
        if (bytesArray.length > 256) {
            return false;
        }

        // initialize the empty bitmap, to be built inside the loop
        uint256 bitmap;
        // initialize an empty uint256 to be used as a bitmask inside the loop
        uint256 bitMask;

        // loop through each byte in the array to construct the bitmap
        for (uint256 i = 0; i < bytesArray.length; ++i) {
            // construct a single-bit mask from the numerical value of the next byte out of the array
            bitMask = uint256(1 << uint8(bytesArray[i]));
            // check that the entry is not a repeat
            if (bitmap & bitMask != 0) {
                return false;
            }
            // add the entry to the bitmap
            bitmap = (bitmap | bitMask);
        }

        // if the loop is completed without returning early, then the array contains no duplicates
        return true;
    }
}