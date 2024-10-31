// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/AVSDirectory.sol";
import "../../src/contracts/core/AllocationManager.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// use forge:
// RUST_LOG=forge,foundry=trace forge script script/tasks/allocate_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address allocationManager,address strategy,address operator,uint32 operatorSetId,uint64 magnitude)" -- <ALLOCATION_MANAGER_ADDRESS> <STRATEGY_ADDRESS> <OPERATOR_ADDRESS> <OPERATOR_SET_ID> <MAGNITUDE>
// RUST_LOG=forge,foundry=trace forge script script/tasks/allocate_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address allocationManager,address strategy,address operator,uint32 operatorSetId,uint64 magnitude)" -- 0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6 0x8aCd85898458400f7Db866d53FCFF6f0D49741FF 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 00000001 0500000000000000000
contract allocateOperatorSet is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    function run(address allocationManager, address strategy, address operator, uint32 operatorSetId, uint64 magnitude) public {
        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        AllocationManager am = AllocationManager(allocationManager);

        // Correct array initialization
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = IStrategy(strategy);

        OperatorSet[] memory sets = new OperatorSet[](1);
        sets[0] = OperatorSet({
            avs: operator,
            operatorSetId: operatorSetId
        });

        uint64[] memory magnitudes = new uint64[](1);
        magnitudes[0] = magnitude;

        // Define a single MagnitudeAllocation and wrap it in an array
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = new IAllocationManagerTypes.MagnitudeAllocation[](1);
        allocations[0] = IAllocationManagerTypes.MagnitudeAllocation({
            strategy: IStrategy(strategy),
            expectedMaxMagnitude: 1000000000000000000,
            operatorSets: sets,
            magnitudes: magnitudes
        });

        // Perform allocation
        am.modifyAllocations(allocations);
        
        vm.stopBroadcast();
    }
}


