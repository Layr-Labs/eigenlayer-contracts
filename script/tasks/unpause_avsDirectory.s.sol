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
// RUST_LOG=forge,foundry=trace forge script script/tasks/unpause_avsDirectory.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address avsDir)" -- <AVS_DIRECTORY_ADDRESS>
// RUST_LOG=forge,foundry=trace forge script script/tasks/unpause_avsDirectory.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address avsDir)" -- 0x5FC8d32690cc91D4c39d9d3abcBD16989F875707
contract unpauseAVSDirectory is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    function run(address avsDir) public {
        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        AVSDirectory avsDirectory = AVSDirectory(avsDir);

        avsDirectory.unpause(0);

        vm.stopBroadcast();
    }
}
