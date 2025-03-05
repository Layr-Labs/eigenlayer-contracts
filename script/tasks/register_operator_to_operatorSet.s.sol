// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/AVSDirectory.sol";
import "../../src/contracts/interfaces/IAVSDirectory.sol";
import "../../src/contracts/core/AllocationManager.sol";
import "../../src/contracts/interfaces/IAllocationManager.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// Define dummy AVSRegistrar contract to prevent revert
contract AVSRegistrar is IAVSRegistrar {
    function registerOperator(
        address operator,
        address avsIdentifier,
        uint32[] calldata operatorSetIds,
        bytes calldata data
    ) external {}
    function deregisterOperator(address operator, address avsIdentifier, uint32[] calldata operatorSetIds) external {}

    function supportsAVS(
        address /*avs*/
    ) external pure returns (bool) {
        return true;
    }

    fallback() external {}
}

// use forge:
// RUST_LOG=forge,foundry=trace forge script script/tasks/register_operator_to_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile)" -- <DEPLOYMENT_OUTPUT_JSON>
// RUST_LOG=forge,foundry=trace forge script script/tasks/register_operator_to_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile)" -- local/slashing_output.json
contract RegisterOperatorToOperatorSets is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    function run(
        string memory configFile
    ) public {
        // Load config
        string memory deployConfigPath = string(bytes(string.concat("script/output/", configFile)));
        string memory config_data = vm.readFile(deployConfigPath);

        // Pull avs directory address
        address avsDir = stdJson.readAddress(config_data, ".addresses.avsDirectory");
        address allocManager = stdJson.readAddress(config_data, ".addresses.allocationManager");
        address strategy = stdJson.readAddress(config_data, ".addresses.strategy");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        // Attach to the deployed contracts
        AVSDirectory avsDirectory = AVSDirectory(avsDir);
        AllocationManager allocationManager = AllocationManager(allocManager);

        // Use privateKey to register as an operator
        address operator = cheats.addr(vm.envUint("PRIVATE_KEY"));
        uint256 expiry = type(uint256).max;
        uint32[] memory oids = new uint32[](1);
        oids[0] = 1;

        // Sign as Operator
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(
            operator,
            avsDirectory.calculateOperatorAVSRegistrationDigestHash(operator, operator, bytes32(uint256(0) + 1), expiry)
        );

        // Add strategies to array
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = IStrategy(strategy);

        // Create OperatorSet(s)
        IAllocationManagerTypes.CreateSetParams[] memory sets = new IAllocationManagerTypes.CreateSetParams[](1);
        sets[0] = IAllocationManagerTypes.CreateSetParams({operatorSetId: 1, strategies: strategies});
        allocationManager.createOperatorSets(msg.sender, sets);

        // Register the Operator to the AVS
        avsDirectory.registerOperatorToAVS(
            operator,
            ISignatureUtilsMixinTypes.SignatureWithSaltAndExpiry(
                abi.encodePacked(r, s, v), bytes32(uint256(0) + 1), expiry
            )
        );

        // Deploy and set registrar.
        allocationManager.setAVSRegistrar(msg.sender, new AVSRegistrar());

        // Register OperatorSet(s)
        IAllocationManagerTypes.RegisterParams memory register =
            IAllocationManagerTypes.RegisterParams({avs: operator, operatorSetIds: oids, data: ""});
        allocationManager.registerForOperatorSets(operator, register);

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();
    }
}
