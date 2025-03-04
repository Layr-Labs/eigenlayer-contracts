// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/DelegationManager.sol";
import "forge-std/Script.sol";
import "forge-std/Test.sol";

contract DelegateToOperator is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    function run(
        string memory configFile
    ) public {
        // Load config
        string memory deployConfigPath = string(bytes(string.concat("script/output/", configFile)));
        string memory config_data = vm.readFile(deployConfigPath);
        address sender = msg.sender;

        // Pull strategy manager address
        address delegationManager = stdJson.readAddress(config_data, ".addresses.delegationManager");

        // Bind the dm contract
        DelegationManager dm = DelegationManager(delegationManager);

        vm.startBroadcast();

        // undelegate the msg sender
        dm.undelegate(sender);

        vm.stopBroadcast();
    }
}
