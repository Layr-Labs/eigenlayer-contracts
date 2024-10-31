// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/AVSDirectory.sol";
import "../../src/contracts/interfaces/IAVSDirectory.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// use forge:
// RUST_LOG=forge,foundry=trace forge script script/tasks/register_operator_to_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address avsDir)" -- <AVS_DIRECTORY_ADDRESS>
// RUST_LOG=forge,foundry=trace forge script script/tasks/register_operator_to_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address avsDir)" -- 0x5FC8d32690cc91D4c39d9d3abcBD16989F875707
contract registerOperatorToOperatorSets is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    uint256 delegationSignerPrivateKey = uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);

    function run(address avsDir) public {
        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        AVSDirectory avsDirectory = AVSDirectory(avsDir);

        address operator = cheats.addr(delegationSignerPrivateKey);
        uint256 expiry = type(uint256).max;
        uint32[] memory oids = new uint32[](1);
        oids[0] = 1;

        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(
            operator, avsDirectory.calculateOperatorSetRegistrationDigestHash(operator, oids, bytes32(uint256(0) + 1), expiry)
        );

        if (!avsDirectory.isOperatorSetAVS(operator)) {
            avsDirectory.becomeOperatorSetAVS();
        }
        avsDirectory.createOperatorSets(oids);
        avsDirectory.registerOperatorToOperatorSets(
            operator, oids, ISignatureUtils.SignatureWithSaltAndExpiry(abi.encodePacked(r, s, v), bytes32(uint256(0) + 1), expiry)
        );

        vm.stopBroadcast();
    }
}
