// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/AVSDirectory.sol";
import "../../src/contracts/interfaces/IAVSDirectory.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// use forge:
// RUST_LOG=forge,foundry=trace forge script script/tasks/register_operator_to_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile)" -- <DEPLOYMENT_OUTPUT_JSON>
// RUST_LOG=forge,foundry=trace forge script script/tasks/register_operator_to_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile)" -- local/slashing_output.json
contract registerOperatorToOperatorSets is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    function run(string memory configFile) public {
        // Load config
        string memory deployConfigPath = string(bytes(string.concat("script/output/", configFile)));
        string memory config_data = vm.readFile(deployConfigPath);

        // Pull avs directory address
        address avsDir = stdJson.readAddress(config_data, ".addresses.avsDirectory");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        // Attach the AVSDirectory
        AVSDirectory avsDirectory = AVSDirectory(avsDir);

        // Use privateKey to register as an operator
        address operator = cheats.addr(vm.envUint("PRIVATE_KEY"));
        uint256 expiry = type(uint256).max;
        uint32[] memory oids = new uint32[](1);
        oids[0] = 1;

        // Sign as Operator
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(
            operator, avsDirectory.calculateOperatorAVSRegistrationDigestHash(operator, operator, bytes32(uint256(0) + 1), expiry)
        );

        // // Upgrade avsDirectory
        // if (!avsDirectory.isOperatorSetAVS(operator)) {
        //     avsDirectory.becomeOperatorSetAVS();
        // }

        // Create OperatorSet(s)
        // avsDirectory.createOperatorSets(oids);

        // Register the Operator to the AVS
        avsDirectory.registerOperatorToAVS(
            operator, ISignatureUtils.SignatureWithSaltAndExpiry(abi.encodePacked(r, s, v), bytes32(uint256(0) + 1), expiry)
        );

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();
    }
}
