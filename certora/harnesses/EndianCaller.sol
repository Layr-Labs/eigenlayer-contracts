// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../src/contracts/libraries/Endian.sol";

contract EndianCaller
{
    function getByteAt(bytes32 data, uint index) external pure returns (bytes1)
    {
        return data[index];
    }

    function fromLittleEndianUint64(bytes32 lenum) external pure returns (uint64 n) {
        return Endian.fromLittleEndianUint64(lenum);
    }

    // copied from src/test/integration/mocks/BeaconChainMock.t.sol
    function toLittleEndianUint64(uint64 num) external pure returns (bytes32) {
        uint256 lenum;
    
        // Rearrange the bytes from big-endian to little-endian format
        lenum |= uint256((num & 0xFF) << 56);
        lenum |= uint256((num & 0xFF00) << 40);
        lenum |= uint256((num & 0xFF0000) << 24);
        lenum |= uint256((num & 0xFF000000) << 8);
        lenum |= uint256((num & 0xFF00000000) >> 8);
        lenum |= uint256((num & 0xFF0000000000) >> 24);
        lenum |= uint256((num & 0xFF000000000000) >> 40);
        lenum |= uint256((num & 0xFF00000000000000) >> 56);

        // Shift the little-endian bytes to the end of the bytes32 value
        return bytes32(lenum << 192);
        //return bytes32(lenum);
    }
}
