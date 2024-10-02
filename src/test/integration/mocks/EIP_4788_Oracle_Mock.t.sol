// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

contract EIP_4788_Oracle_Mock {
    mapping(uint256 => bytes32) blockRoots;

    uint256 constant HISTORY_BUFFER_LENGTH = 8191;

    fallback() external {
        require(msg.data.length == 32, "4788OracleMock.fallback: malformed msg.data");

        uint256 timestamp = abi.decode(msg.data, (uint256));
        require(timestamp != 0, "4788OracleMock.fallback: timestamp is 0");

        bytes32 blockRoot = blockRoots[timestamp];
        require(blockRoot != 0, "4788OracleMock.fallback: no block root found. DID YOU USE CHEATS.WARP?");

        assembly {
            mstore(0, blockRoot)
            return(0, 32)
        }
    }

    function timestampToBlockRoot(
        uint256 timestamp
    ) public view returns (bytes32) {
        return blockRoots[uint64(timestamp)];
    }

    function setBlockRoot(uint64 timestamp, bytes32 blockRoot) public {
        blockRoots[timestamp] = blockRoot;
    }
}
