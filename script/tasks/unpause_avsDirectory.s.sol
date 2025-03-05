// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/AVSDirectory.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// use cast:
//
// cast send <TOKEN_ADDRESS> "approve(address,uint256)" \
// <STRATEGY_MANAGER_ADDRESS> \
// <AMOUNT> \
// --private-key <YOUR_PRIVATE_KEY>
//
// cast send <STRATEGY_MANAGER_ADDRESS> "depositIntoStrategy(address,address,uint256)" \
// <STRATEGY_ADDRESS> \
// <TOKEN_ADDRESS> \
// <AMOUNT> \
// --private-key <YOUR_PRIVATE_KEY>

// use forge:
// RUST_LOG=forge,foundry=trace forge script script/tasks/unpause_avsDirectory.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile)" -- <DEPLOYMENT_OUTPUT_JSON>
// RUST_LOG=forge,foundry=trace forge script script/tasks/unpause_avsDirectory.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile)" -- local/slashing_output.json
contract UnpauseAVSDirectory is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    function run(
        string memory configFile
    ) public {
        // Load config
        string memory deployConfigPath = string(bytes(string.concat("script/output/", configFile)));
        string memory config_data = vm.readFile(deployConfigPath);

        // Pull avs directory address
        address avsDir = stdJson.readAddress(config_data, ".addresses.avsDirectory");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        // Attach to the AVSDirectory
        AVSDirectory avsDirectory = AVSDirectory(avsDir);

        // Unpause the AVSDirectory
        avsDirectory.unpause(0);

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();
    }
}
