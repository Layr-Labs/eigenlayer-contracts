// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/AllocationManager.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// use forge:
// RUST_LOG=forge,foundry=trace forge script script/tasks/slash_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile,address operator,uint32 operatorSetId,uint256 wadToSlash)" -- <DEPLOYMENT_OUTPUT_JSON> <OPERATOR_ADDRESS> <OPERATOR_SET_ID> <WADS_TO_SLASH>
// RUST_LOG=forge,foundry=trace forge script script/tasks/slash_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile,address operator,uint32 operatorSetId,uint256 wadToSlash)" -- local/slashing_output.json 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 00000001 05000000
contract SlashOperatorSet is Script, Test, IAllocationManagerTypes {
    Vm cheats = Vm(VM_ADDRESS);

    function run(
        string memory configFile,
        address operator,
        uint32 operatorSetId,
        IStrategy[] memory strategies,
        uint256[] memory wadsToSlash
    ) public {
        // Load config
        string memory deployConfigPath = string(bytes(string.concat("script/output/", configFile)));
        string memory config_data = vm.readFile(deployConfigPath);

        // Pull allocationManager address
        address allocationManager = stdJson.readAddress(config_data, ".addresses.allocationManager");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        // Attach to the AllocationManager
        AllocationManager am = AllocationManager(allocationManager);

        // Define SlashingParams struct instance with correct array initialization
        SlashingParams memory slashing = SlashingParams({
            operator: operator,
            operatorSetId: operatorSetId,
            strategies: strategies,
            wadsToSlash: wadsToSlash,
            description: "slashed"
        });

        // Perform slashing
        am.slashOperator(operator, slashing);

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();
    }
}
