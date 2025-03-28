// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console2} from "forge-std/Test.sol";
import {DeployFromScratch} from "script/deploy/local/deploy_from_scratch.slashing.s.sol";

// NOTE: Run the following command to deploy from scratch in an anvil instance:
// RUST_LOG=forge,foundry=trace forge script script/deploy/local/Deploy_From_Scratch.s.sol --slow \
//   --rpc-url http://127.0.0.1:8545 \
//   --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
//   --broadcast \
//   --sig "run(string memory configFile)" \
//   -- local/deploy_from_scratch.slashing.anvil.config.json
contract DeployTest is Test {
    DeployFromScratch public deployer;

    function setUp() public {
        deployer = new DeployFromScratch();
    }

    function test_DeployFromScratch() public {
        // Deploy, expecting no revert.
        deployer.run("local/deploy_from_scratch.slashing.anvil.config.json");
    }
}
