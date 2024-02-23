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

    // function testCallOracleForkTest() public {
    //    emit log_named_uint("timestamp", block.timestamp);
    //     bytes32 blockRoot = oracle.timestampToBlockRoot(block.timestamp);

    // }

    // function testCallOracleForkTestFromEOA() public {
    //     address eoa = address(0x123);
    //     cheats.startPrank(eoa);
    //     emit log_named_uint("timestamp", block.timestamp);
    //     (bool success, bytes memory data) = address(0x000F3df6D732807Ef1319fB7B8bB8522d0Beac02).call(abi.encodePacked(block.timestamp));
    //     require(success, "Precompile call failed");
    // }
}