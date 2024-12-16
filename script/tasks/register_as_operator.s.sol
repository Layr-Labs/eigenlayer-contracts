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
// RUST_LOG=forge,foundry=trace forge script script/tasks/register_as_operator.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile,address operator,string memory metadataURI)" -- <DEPLOYMENT_OUTPUT_JSON> <OPERATOR_ADDRESS> <METADATA_URI>
// RUST_LOG=forge,foundry=trace forge script script/tasks/register_as_operator.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile,address operator,string metadataURI)" -- local/slashing_output.json 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 "test"
contract RegisterAsOperator is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    function run(string memory configFile, address operator, string memory metadataURI) public {
        // Load config
        string memory deployConfigPath = string(bytes(string.concat("script/output/", configFile)));
        string memory config_data = vm.readFile(deployConfigPath);

        // Pull delegation manager address
        address delegationManager = stdJson.readAddress(config_data, ".addresses.delegationManager");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        // Attach the delegationManager
        DelegationManager delegation = DelegationManager(delegationManager);

        // Register the sender as an Operator
        delegation.registerAsOperator(operator, 0, metadataURI);

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();
    }
}
