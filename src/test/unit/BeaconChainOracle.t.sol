// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;


import "forge-std/Test.sol";
import "../../contracts/pods/BeaconChainOracle.sol";



contract BeaconChainOracleTest is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    event OracleBlockRootAdded(uint256 timestamp, bytes32 blockRoot);


    BeaconChainOracle oracle;
    function setUp() public {
        oracle = new BeaconChainOracle();
    }

    function testCallOracleForkTest() public {
       emit log_named_uint("timestamp", block.timestamp);
        bytes32 blockRoot = oracle.timestampToBlockRoot(block.timestamp);

    }
}