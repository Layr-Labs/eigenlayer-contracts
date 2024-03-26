// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../contracts/interfaces/IBeaconChainOracle.sol";



contract BeaconChainOracleMock is IBeaconChainOracle {

    bytes32 public mockBeaconChainStateRoot;

    function getOracleBlockRootAtTimestamp() external view returns(bytes32) {
        return mockBeaconChainStateRoot;
    }

    function setOracleBlockRootAtTimestamp(bytes32 beaconChainStateRoot) external {
        mockBeaconChainStateRoot = beaconChainStateRoot;
    }

    function timestampToBlockRoot(uint256 /*blockNumber*/) external view returns(bytes32) {
        return mockBeaconChainStateRoot;
    }

    function isOracleSigner(address /*_oracleSigner*/) external pure returns(bool) {
        return true;
    }

    function hasVoted(uint64 /*blockNumber*/, address /*oracleSigner*/) external pure returns(bool) {
        return true;
    }

    function stateRootVotes(uint64 /*blockNumber*/, bytes32 /*stateRoot*/) external pure returns(uint256) {
        return 0;
    }

    function totalOracleSigners() external pure returns(uint256) {
        return 0;
    }

    function threshold() external pure returns(uint256) {
        return 0;
    }

    function setThreshold(uint256 /*_threshold*/) external pure {}

    function addOracleSigners(address[] memory /*_oracleSigners*/) external pure {}

    function removeOracleSigners(address[] memory /*_oracleSigners*/) external pure {}

    function voteForBeaconChainStateRoot(uint64 /*blockNumber*/, bytes32 /*stateRoot*/) external pure {}

    function latestConfirmedOracleBlockNumber() external pure returns(uint64) {}
}
