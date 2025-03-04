// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/DelegationManager.sol";
import "forge-std/Script.sol";
import "forge-std/Test.sol";

contract DelegateToOperator is Script, Test, ISignatureUtilsMixinTypes {
    Vm cheats = Vm(VM_ADDRESS);

    function run(string memory configFile, address operator, uint256 operatorPrivateKey, uint256 saltInt) public {
        // Load config
        string memory deployConfigPath = string(bytes(string.concat("script/output/", configFile)));
        string memory config_data = vm.readFile(deployConfigPath);

        // Pull strategy manager address
        address delegationManager = stdJson.readAddress(config_data, ".addresses.delegationManager");

        // Bind the dm contract
        DelegationManager dm = DelegationManager(delegationManager);

        // Set salt and expiry
        bytes32 salt = bytes32(uint256(0) + saltInt);
        uint256 expiry = type(uint256).max;
        address sender = msg.sender;

        // Operator signs off-chain
        bytes32 digest = dm.calculateDelegationApprovalDigestHash(sender, operator, operator, salt, expiry);

        // Operator signs the message
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(operatorPrivateKey, digest);

        // Sender executes on-chain
        vm.startBroadcast();

        dm.delegateTo(operator, SignatureWithExpiry({signature: abi.encodePacked(r, s, v), expiry: expiry}), salt);

        vm.stopBroadcast();
    }
}
