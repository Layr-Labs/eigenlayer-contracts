// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../harnesses/BitmapUtilsWrapper.sol";
// import "../../contracts/libraries/BitmapUtils.sol";

import "forge-std/Test.sol";

contract BitmapUtilsUnitTests is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    BitmapUtilsWrapper public bitmapUtilsWrapper;

    function setUp() public {
        bitmapUtilsWrapper = new BitmapUtilsWrapper();
    }

    // ensure that the bitmap encoding of an emtpy bytes array is an emtpy bitmap (function doesn't revert and approriately returns uint256(0))
    function testEmptyArrayEncoding() public view {
        bytes memory emptyBytesArray;
        uint256 returnedBitMap = bitmapUtilsWrapper.bytesArrayToBitmap(emptyBytesArray);
        require(returnedBitMap == 0, "BitmapUtilsUnitTests.testEmptyArrayEncoding: empty array not encoded to empty bitmap");
    }

    // ensure that the bitmap encoding of a single uint8 (i.e. a single byte) matches the expected output
    function testSingleByteEncoding(uint8 fuzzedNumber) public view {
        bytes1 singleByte = bytes1(fuzzedNumber);
        bytes memory bytesArray = abi.encodePacked(singleByte);
        uint256 returnedBitMap = bitmapUtilsWrapper.bytesArrayToBitmap(bytesArray);
        uint256 bitMask = uint256(1 << fuzzedNumber);
        require(returnedBitMap == bitMask, "BitmapUtilsUnitTests.testSingleByteEncoding: non-equivalence");
    }

    // ensure that the bitmap encoding of a two uint8's (i.e. a two byte array) matches the expected output
    function testTwoByteEncoding(uint8 firstFuzzedNumber, uint8 secondFuzzedNumber) public {
        bytes1 firstSingleByte = bytes1(firstFuzzedNumber);
        bytes1 secondSingleByte = bytes1(secondFuzzedNumber);
        bytes memory bytesArray = abi.encodePacked(firstSingleByte, secondSingleByte);
        if (firstFuzzedNumber == secondFuzzedNumber) {
            cheats.expectRevert(bytes("BitmapUtils.bytesArrayToBitmap: repeat entry in bytesArray"));
            bitmapUtilsWrapper.bytesArrayToBitmap(bytesArray);
        } else {
            uint256 returnedBitMap = bitmapUtilsWrapper.bytesArrayToBitmap(bytesArray);
            uint256 firstBitMask = uint256(1 << firstFuzzedNumber);
            uint256 secondBitMask = uint256(1 << secondFuzzedNumber);
            uint256 combinedBitMask = firstBitMask | secondBitMask;
            require(returnedBitMap == combinedBitMask, "BitmapUtilsUnitTests.testTwoByteEncoding: non-equivalence");
        }
    }

    // ensure that converting bytes array => bitmap => bytes array is returns the original bytes array (i.e. is lossless and artifactless)
    // note that this only works on ordered arrays, because unordered arrays will be returned ordered
    function testBytesArrayToBitmapToBytesArray(bytes memory originalBytesArray) public view {
        // filter down to only ordered inputs
        cheats.assume(bitmapUtilsWrapper.isArrayStrictlyAscendingOrdered(originalBytesArray));
        uint256 bitmap = bitmapUtilsWrapper.bytesArrayToBitmap(originalBytesArray);
        bytes memory returnedBytesArray = bitmapUtilsWrapper.bitmapToBytesArray(bitmap);
        // emit log_named_bytes("originalBytesArray", originalBytesArray);
        // emit log_named_uint("bitmap", bitmap);
        // emit log_named_bytes("returnedBytesArray", returnedBytesArray);
        require(keccak256(abi.encodePacked(originalBytesArray)) == keccak256(abi.encodePacked(returnedBytesArray)),
            "BitmapUtilsUnitTests.testBytesArrayToBitmapToBytesArray: output doesn't match input");
    }

    // ensure that converting bytes array => bitmap => bytes array is returns the original bytes array (i.e. is lossless and artifactless)
    // note that this only works on ordered arrays, because unordered arrays will be returned ordered
    function testBytesArrayToBitmapToBytesArray_Yul(bytes memory originalBytesArray) public view {
        // filter down to only ordered inputs
        cheats.assume(bitmapUtilsWrapper.isArrayStrictlyAscendingOrdered(originalBytesArray));
        uint256 bitmap = bitmapUtilsWrapper.bytesArrayToBitmap_Yul(originalBytesArray);
        bytes memory returnedBytesArray = bitmapUtilsWrapper.bitmapToBytesArray(bitmap);
        // emit log_named_bytes("originalBytesArray", originalBytesArray);
        // emit log_named_uint("bitmap", bitmap);
        // emit log_named_bytes("returnedBytesArray", returnedBytesArray);
        require(keccak256(abi.encodePacked(originalBytesArray)) == keccak256(abi.encodePacked(returnedBytesArray)),
            "BitmapUtilsUnitTests.testBytesArrayToBitmapToBytesArray: output doesn't match input");
    }

    // ensure that converting bytes array => bitmap => bytes array is returns the original bytes array (i.e. is lossless and artifactless)
    // note that this only works on ordered arrays
    function testBytesArrayToBitmapToBytesArray_OrderedVersion(bytes memory originalBytesArray) public view {
        // filter down to only ordered inputs
        cheats.assume(bitmapUtilsWrapper.isArrayStrictlyAscendingOrdered(originalBytesArray));
        uint256 bitmap = bitmapUtilsWrapper.orderedBytesArrayToBitmap(originalBytesArray);
        bytes memory returnedBytesArray = bitmapUtilsWrapper.bitmapToBytesArray(bitmap);
        require(keccak256(abi.encodePacked(originalBytesArray)) == keccak256(abi.encodePacked(returnedBytesArray)),
            "BitmapUtilsUnitTests.testBytesArrayToBitmapToBytesArray: output doesn't match input");
    }

    // ensure that converting bytes array => bitmap => bytes array is returns the original bytes array (i.e. is lossless and artifactless)
    // note that this only works on ordered arrays
    function testBytesArrayToBitmapToBytesArray_OrderedVersion_Yul(bytes memory originalBytesArray) public view {
        // filter down to only ordered inputs
        cheats.assume(bitmapUtilsWrapper.isArrayStrictlyAscendingOrdered(originalBytesArray));
        uint256 bitmap = bitmapUtilsWrapper.orderedBytesArrayToBitmap(originalBytesArray);
        bytes memory returnedBytesArray = bitmapUtilsWrapper.bitmapToBytesArray(bitmap);
        require(keccak256(abi.encodePacked(originalBytesArray)) == keccak256(abi.encodePacked(returnedBytesArray)),
            "BitmapUtilsUnitTests.testBytesArrayToBitmapToBytesArray: output doesn't match input");
    }

    // ensure that converting bitmap => bytes array => bitmap is returns the original bitmap (i.e. is lossless and artifactless)
    function testBitMapToBytesArrayToBitmap(uint256 originalBitmap) public view {
        bytes memory bytesArray = bitmapUtilsWrapper.bitmapToBytesArray(originalBitmap);
        uint256 returnedBitMap = bitmapUtilsWrapper.bytesArrayToBitmap(bytesArray);
        require(returnedBitMap == originalBitmap, "BitmapUtilsUnitTests.testBitMapToArrayToBitmap: output doesn't match input");
    }

    // testing one function for a specific input. used for comparing gas costs
    function testBytesArrayToBitmap_OrderedVersion_Yul_SpecificInput(/*bytes memory originalBytesArray*/) public {
        bytes memory originalBytesArray =
            abi.encodePacked(bytes1(uint8(5)), bytes1(uint8(6)), bytes1(uint8(7)), bytes1(uint8(8)), bytes1(uint8(9)), bytes1(uint8(10)), bytes1(uint8(11)), bytes1(uint8(12)));
        // bytes memory originalBytesArray = abi.encodePacked(bytes1(uint8(5)));
        uint256 gasLeftBefore = gasleft();
        bitmapUtilsWrapper.orderedBytesArrayToBitmap_Yul(originalBytesArray);
        uint256 gasLeftAfter = gasleft();
        uint256 gasSpent = gasLeftBefore - gasLeftAfter;
        emit log_named_uint("gasSpent", gasSpent);
    }

    // testing one function for a specific input. used for comparing gas costs
    function testBytesArrayToBitmap_OrderedVersion_SpecificInput(/*bytes memory originalBytesArray*/) public {
        bytes memory originalBytesArray =
            abi.encodePacked(bytes1(uint8(5)), bytes1(uint8(6)), bytes1(uint8(7)), bytes1(uint8(8)), bytes1(uint8(9)), bytes1(uint8(10)), bytes1(uint8(11)), bytes1(uint8(12)));
        // bytes memory originalBytesArray = abi.encodePacked(bytes1(uint8(5)));
        uint256 gasLeftBefore = gasleft();
        bitmapUtilsWrapper.orderedBytesArrayToBitmap(originalBytesArray);
        uint256 gasLeftAfter = gasleft();
        uint256 gasSpent = gasLeftBefore - gasLeftAfter;
        emit log_named_uint("gasSpent", gasSpent);
    }

    // testing one function for a specific input. used for comparing gas costs
    function testBytesArrayToBitmap_SpecificInput(/*bytes memory originalBytesArray*/) public {
        bytes memory originalBytesArray =
            abi.encodePacked(bytes1(uint8(5)), bytes1(uint8(6)), bytes1(uint8(7)), bytes1(uint8(8)), bytes1(uint8(9)), bytes1(uint8(10)), bytes1(uint8(11)), bytes1(uint8(12)));
        // bytes memory originalBytesArray = abi.encodePacked(bytes1(uint8(5)));
        uint256 gasLeftBefore = gasleft();
        bitmapUtilsWrapper.bytesArrayToBitmap(originalBytesArray);
        uint256 gasLeftAfter = gasleft();
        uint256 gasSpent = gasLeftBefore - gasLeftAfter;
        emit log_named_uint("gasSpent", gasSpent);
    }

    // testing one function for a specific input. used for comparing gas costs
    function testBytesArrayToBitmap_Yul_SpecificInput(/*bytes memory originalBytesArray*/) public {
        bytes memory originalBytesArray =
            abi.encodePacked(bytes1(uint8(5)), bytes1(uint8(6)), bytes1(uint8(7)), bytes1(uint8(8)), bytes1(uint8(9)), bytes1(uint8(10)), bytes1(uint8(11)), bytes1(uint8(12)));
        // bytes memory originalBytesArray = abi.encodePacked(bytes1(uint8(5)));
        uint256 gasLeftBefore = gasleft();
        bitmapUtilsWrapper.bytesArrayToBitmap_Yul(originalBytesArray);
        uint256 gasLeftAfter = gasleft();
        uint256 gasSpent = gasLeftBefore - gasLeftAfter;
        emit log_named_uint("gasSpent", gasSpent);
    }
}