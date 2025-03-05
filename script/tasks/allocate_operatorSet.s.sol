// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/AVSDirectory.sol";
import "../../src/contracts/core/AllocationManager.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// use forge:
// RUST_LOG=forge,foundry=trace forge script script/tasks/allocate_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile,address strategy,address avs,uint32 operatorSetId,uint64 magnitude)" -- <DEPLOYMENT_OUTPUT_JSON> <STRATEGY_ADDRESS> <AVS_ADDRESS> <OPERATOR_SET_ID> <MAGNITUDE>
// RUST_LOG=forge,foundry=trace forge script script/tasks/allocate_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile,address strategy,address avs,uint32 operatorSetId,uint64 magnitude)" -- local/slashing_output.json 0x8aCd85898458400f7Db866d53FCFF6f0D49741FF 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 00000001 0500000000000000000
contract AllocateOperatorSet is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    function run(
        string memory configFile,
        address strategy,
        address avs,
        uint32 operatorSetId,
        uint64 magnitude
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

        // Correct array initialization
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = IStrategy(strategy);

        // Set OperatorSets
        OperatorSet[] memory sets = new OperatorSet[](1);
        sets[0] = OperatorSet({avs: avs, id: operatorSetId});

        // Set new mangitudes
        uint64[] memory magnitudes = new uint64[](1);
        magnitudes[0] = magnitude;

        // Define a single MagnitudeAllocation and wrap it in an array
        IAllocationManagerTypes.AllocateParams[] memory allocations = new IAllocationManagerTypes.AllocateParams[](1);
        allocations[0] = IAllocationManagerTypes.AllocateParams({
            operatorSet: sets[0],
            strategies: strategies,
            newMagnitudes: magnitudes
        });

        // Perform allocation
        am.modifyAllocations(msg.sender, allocations);

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();
    }
}
