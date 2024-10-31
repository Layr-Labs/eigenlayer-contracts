// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/DelegationManager.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// use cast:
// cast send <DELEGATION_MANAGER_ADDRESS> "registerAsOperator((address,address,uint32),uint256,string)" \
// "(address(0), <OPERATOR_ADDRESS>, 0)" \
// 0 \
// "<METADATA_URI>" \
// --private-key <YOUR_PRIVATE_KEY>

// use forge:
// RUST_LOG=forge,foundry=trace forge script script/tasks/register_as_operator.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address delegationManager,address operator,string memory metadataURI)" -- <DELEGATION_MANAGER_ADDRESS> <OPERATOR_ADDRESS> <METADATA_URI>
// RUST_LOG=forge,foundry=trace forge script script/tasks/register_as_operator.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address delegationManager,address operator,string metadataURI)" -- 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 "test"
contract registerAsOperator is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    function run(address delegationManager, address operator, string memory metadataURI) public {
        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        DelegationManager delegation = DelegationManager(delegationManager);

        // Define OperatorDetails struct instance
        IDelegationManagerTypes.OperatorDetails memory operatorDetails = IDelegationManagerTypes.OperatorDetails({
            __deprecated_earningsReceiver: address(0),
            delegationApprover: operator,
            __deprecated_stakerOptOutWindowBlocks: 0
        });

        delegation.registerAsOperator(operatorDetails, 0, metadataURI);
        
        vm.stopBroadcast();
    }
}
