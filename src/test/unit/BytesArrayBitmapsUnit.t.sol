// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../harnesses/BytesArrayBitmapsWrapper.sol";
// import "../../contracts/libraries/BytesArrayBitmaps.sol";

import "forge-std/Test.sol";

contract BytesArrayBitmapsUnitTests is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    BytesArrayBitmapsWrapper public bytesArrayBitmapsWrapper;

    function setUp() public {
        bytesArrayBitmapsWrapper = new BytesArrayBitmapsWrapper();
    }

    // ensure that the bitmap encoding of a single uint8 (i.e. a single byte) matches the expected output
    function testSingleByteEncoding(uint8 fuzzedNumber) public view {
        bytes1 singleByte = bytes1(fuzzedNumber);
        bytes memory bytesArray = abi.encodePacked(singleByte);
        uint256 returnedBitMap = bytesArrayBitmapsWrapper.bytesArrayToBitmap(bytesArray);
        uint256 bitMask = uint256(1 << fuzzedNumber);
        require(returnedBitMap == bitMask, "BytesArrayBitmapsUnitTests.testSingleByteEncoding: non-equivalence");
    }

    // ensure that the bitmap encoding of a two uint8's (i.e. a two byte array) matches the expected output
    function testTwoByteEncoding(uint8 firstFuzzedNumber, uint8 secondFuzzedNumber) public {
        bytes1 firstSingleByte = bytes1(firstFuzzedNumber);
        bytes1 secondSingleByte = bytes1(secondFuzzedNumber);
        bytes memory bytesArray = abi.encodePacked(firstSingleByte, secondSingleByte);
        if (firstFuzzedNumber == secondFuzzedNumber) {
            cheats.expectRevert(bytes("BytesArrayBitmaps.bytesArrayToBitmap: repeat entry in bytesArray"));
            bytesArrayBitmapsWrapper.bytesArrayToBitmap(bytesArray);
        } else {
            uint256 returnedBitMap = bytesArrayBitmapsWrapper.bytesArrayToBitmap(bytesArray);
            uint256 firstBitMask = uint256(1 << firstFuzzedNumber);
            uint256 secondBitMask = uint256(1 << secondFuzzedNumber);
            uint256 combinedBitMask = firstBitMask | secondBitMask;
            require(returnedBitMap == combinedBitMask, "BytesArrayBitmapsUnitTests.testTwoByteEncoding: non-equivalence");
        }
    }

    // ensure that converting bytes array => bitmap => bytes array is returns the original bytes array (i.e. is lossless and artifactless)
    // note that this only works on ordered arrays
    function testBytesArrayToBitmapToBytesArray(bytes memory originalBytesArray) public view {
        // filter down to only ordered inputs
        cheats.assume(bytesArrayBitmapsWrapper.isArrayStrictlyAscendingOrdered(originalBytesArray));
        uint256 bitmap = bytesArrayBitmapsWrapper.bytesArrayToBitmap(originalBytesArray);
        bytes memory returnedBytesArray = bytesArrayBitmapsWrapper.bitmapToBytesArray(bitmap);
        // emit log_named_bytes("originalBytesArray", originalBytesArray);
        // emit log_named_uint("bitmap", bitmap);
        // emit log_named_bytes("returnedBytesArray", returnedBytesArray);
        require(keccak256(abi.encodePacked(originalBytesArray)) == keccak256(abi.encodePacked(returnedBytesArray)),
            "BytesArrayBitmapsUnitTests.testBytesArrayToBitmapToBytesArray: output doesn't match input");
    }

    // ensure that converting bytes array => bitmap => bytes array is returns the original bytes array (i.e. is lossless and artifactless)
    // note that this only works on ordered arrays
    function testBytesArrayToBitmapToBytesArray_OrderedVersion(bytes memory originalBytesArray) public view {
        // filter down to only ordered inputs
        cheats.assume(bytesArrayBitmapsWrapper.isArrayStrictlyAscendingOrdered(originalBytesArray));
        uint256 bitmap = bytesArrayBitmapsWrapper.orderedBytesArrayToBitmap(originalBytesArray);
        bytes memory returnedBytesArray = bytesArrayBitmapsWrapper.bitmapToBytesArray(bitmap);
        require(keccak256(abi.encodePacked(originalBytesArray)) == keccak256(abi.encodePacked(returnedBytesArray)),
            "BytesArrayBitmapsUnitTests.testBytesArrayToBitmapToBytesArray: output doesn't match input");
    }

    // ensure that converting bytes array => bitmap => bytes array is returns the original bytes array (i.e. is lossless and artifactless)
    // note that this only works on ordered arrays
    function testBytesArrayToBitmapToBytesArray_OrderedVersion_Yul(bytes memory originalBytesArray) public view {
        // filter down to only ordered inputs
        cheats.assume(bytesArrayBitmapsWrapper.isArrayStrictlyAscendingOrdered(originalBytesArray));
        uint256 bitmap = bytesArrayBitmapsWrapper.orderedBytesArrayToBitmap(originalBytesArray);
        bytes memory returnedBytesArray = bytesArrayBitmapsWrapper.bitmapToBytesArray(bitmap);
        require(keccak256(abi.encodePacked(originalBytesArray)) == keccak256(abi.encodePacked(returnedBytesArray)),
            "BytesArrayBitmapsUnitTests.testBytesArrayToBitmapToBytesArray: output doesn't match input");
    }

    // ensure that converting bitmap => bytes array => bitmap is returns the original bitmap (i.e. is lossless and artifactless)
    function testBitMapToBytesArrayToBitmap(uint256 originalBitmap) public view {
        bytes memory bytesArray = bytesArrayBitmapsWrapper.bitmapToBytesArray(originalBitmap);
        uint256 returnedBitMap = bytesArrayBitmapsWrapper.bytesArrayToBitmap(bytesArray);
        require(returnedBitMap == originalBitmap, "BytesArrayBitmapsUnitTests.testBitMapToArrayToBitmap: output doesn't match input");
    }

    // ensure that converting bytes array => bitmap => bytes array is returns the original bytes array (i.e. is lossless and artifactless)
    // note that this only works on ordered arrays
    function testBytesArrayToBitmapToBytesArray_OrderedVersion_Yul_SpecificInput(/*bytes memory originalBytesArray*/) public view {
        bytes memory originalBytesArray =
            abi.encodePacked(bytes1(uint8(5)), bytes1(uint8(6)), bytes1(uint8(7)), bytes1(uint8(8)), bytes1(uint8(9)), bytes1(uint8(10)), bytes1(uint8(11)), bytes1(uint8(12)));
        // bytes memory originalBytesArray = abi.encodePacked(bytes1(uint8(5)));
        bytesArrayBitmapsWrapper.orderedBytesArrayToBitmap_Yul(originalBytesArray);
        // uint256 bitmap = bytesArrayBitmapsWrapper.orderedBytesArrayToBitmap_Yul(originalBytesArray);
        // bytes memory returnedBytesArray = bytesArrayBitmapsWrapper.bitmapToBytesArray(bitmap);
        // require(keccak256(abi.encodePacked(originalBytesArray)) == keccak256(abi.encodePacked(returnedBytesArray)),
        //     "BytesArrayBitmapsUnitTests.testBytesArrayToBitmapToBytesArray: output doesn't match input");
    }

    // ensure that converting bytes array => bitmap => bytes array is returns the original bytes array (i.e. is lossless and artifactless)
    // note that this only works on ordered arrays
    function testBytesArrayToBitmapToBytesArray_OrderedVersion_SpecificInput(/*bytes memory originalBytesArray*/) public view {
        bytes memory originalBytesArray =
            abi.encodePacked(bytes1(uint8(5)), bytes1(uint8(6)), bytes1(uint8(7)), bytes1(uint8(8)), bytes1(uint8(9)), bytes1(uint8(10)), bytes1(uint8(11)), bytes1(uint8(12)));
        // bytes memory originalBytesArray = abi.encodePacked(bytes1(uint8(5)));
        bytesArrayBitmapsWrapper.orderedBytesArrayToBitmap(originalBytesArray);
        // uint256 bitmap = bytesArrayBitmapsWrapper.orderedBytesArrayToBitmap(originalBytesArray);
        // bytes memory returnedBytesArray = bytesArrayBitmapsWrapper.bitmapToBytesArray(bitmap);
        // require(keccak256(abi.encodePacked(originalBytesArray)) == keccak256(abi.encodePacked(returnedBytesArray)),
        //     "BytesArrayBitmapsUnitTests.testBytesArrayToBitmapToBytesArray: output doesn't match input");
    }

    // ensure that converting bytes array => bitmap => bytes array is returns the original bytes array (i.e. is lossless and artifactless)
    // note that this only works on ordered arrays
    function testBytesArrayToBitmapToBytesArray_SpecificInput(/*bytes memory originalBytesArray*/) public view {
        bytes memory originalBytesArray =
            abi.encodePacked(bytes1(uint8(5)), bytes1(uint8(6)), bytes1(uint8(7)), bytes1(uint8(8)), bytes1(uint8(9)), bytes1(uint8(10)), bytes1(uint8(11)), bytes1(uint8(12)));
        // bytes memory originalBytesArray = abi.encodePacked(bytes1(uint8(5)));
        bytesArrayBitmapsWrapper.bytesArrayToBitmap(originalBytesArray);
        // uint256 bitmap = bytesArrayBitmapsWrapper.bytesArrayToBitmap(originalBytesArray);
        // bytes memory returnedBytesArray = bytesArrayBitmapsWrapper.bitmapToBytesArray(bitmap);
        // require(keccak256(abi.encodePacked(originalBytesArray)) == keccak256(abi.encodePacked(returnedBytesArray)),
        //     "BytesArrayBitmapsUnitTests.testBytesArrayToBitmapToBytesArray: output doesn't match input");
    }
}