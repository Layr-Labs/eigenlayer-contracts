// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/contracts/interfaces/IBeaconChainOracle.sol";

contract BeaconChainOracleMock is IBeaconChainOracle {

    mapping(uint64 => bytes32) blockRoots;

    function timestampToBlockRoot(uint timestamp) public view returns (bytes32) {
        return blockRoots[uint64(timestamp)];
    }

    function setBlockRoot(uint64 timestamp, bytes32 blockRoot) public {
        blockRoots[timestamp] = blockRoot;
    }
}
