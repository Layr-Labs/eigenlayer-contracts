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
}
