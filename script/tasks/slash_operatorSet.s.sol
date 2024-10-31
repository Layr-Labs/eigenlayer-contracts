// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/AllocationManager.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// use forge:
// RUST_LOG=forge,foundry=trace forge script script/tasks/slash_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address allocationManager,address strategy,address operator,uint32 operatorSetId,uint256 wadToSlash)" -- <ALLOCATION_MANAGER_ADDRESS> <STRATEGY_ADDRESS> <OPERATOR_ADDRESS> <OPERATOR_SET_ID> <WADS_TO_SLASH>
// RUST_LOG=forge,foundry=trace forge script script/tasks/slash_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address allocationManager,address strategy,address operator,uint32 operatorSetId,uint256 wadToSlash)" -- 0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6 0x8aCd85898458400f7Db866d53FCFF6f0D49741FF 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 00000001 05000000
contract slashOperatorSet is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    function run(address allocationManager, address strategy, address operator, uint32 operatorSetId, uint256 wadToSlash) public {
        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        AllocationManager am = AllocationManager(allocationManager);

        // Correct array initialization
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = IStrategy(strategy);

        // Define SlashingParams struct instance with correct array initialization
        IAllocationManagerTypes.SlashingParams memory slashing = IAllocationManagerTypes.SlashingParams({
            operator: operator,
            operatorSetId: operatorSetId,
            strategies: strategies,
            wadToSlash: wadToSlash,
            description: "slashed"
        });

        // Perform slashing
        am.slashOperator(slashing);
        
        vm.stopBroadcast();
    }
}
