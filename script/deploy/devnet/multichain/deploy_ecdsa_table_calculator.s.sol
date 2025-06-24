// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Core Contracts
import "src/contracts/core/AllocationManager.sol";
import "src/contracts/permissions/KeyRegistrar.sol";

// Multichain Contracts
import "src/contracts/multichain/ECDSATableCalculator.sol";

// Forge
import "forge-std/Script.sol";
import "forge-std/Test.sol";

// Run with:
// forge script script/deploy/devnet/multichain/deploy_ecdsa_table_calculator.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run()" --verify $ETHERSCAN_API_KEY
contract DeployECDSATableCalculator is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    // Owner address (same as used in other multichain deployments)
    address owner = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;

    // Existing contract addresses on devnet
    AllocationManager public allocationManager = AllocationManager(0xFdD5749e11977D60850E06bF5B13221Ad95eb6B4);
    KeyRegistrar public keyRegistrar = KeyRegistrar(0x1C84Bb62fE7791e173014A879C706445fa893BbE);

    // Contract to be deployed
    ECDSATableCalculator public ecdsaTableCalculator;

    // Configuration
    uint256 public constant LOOKAHEAD_BLOCKS = 100; // Same as BN254TableCalculator

    function run() public {
        uint256 chainId = block.chainid;
        emit log_named_uint("Deploying ECDSATableCalculator on ChainID", chainId);

        /**
         *
         *                     CONTRACT DEPLOYMENT
         *
         */
        vm.startBroadcast();

        console.log("Deploying ECDSATableCalculator...");
        console.log("  KeyRegistrar:", address(keyRegistrar));
        console.log("  AllocationManager:", address(allocationManager));
        console.log("  Lookahead Blocks:", LOOKAHEAD_BLOCKS);

        // Deploy ECDSATableCalculator as a non-upgradeable contract
        ecdsaTableCalculator = new ECDSATableCalculator(
            keyRegistrar,
            allocationManager,
            LOOKAHEAD_BLOCKS
        );

        console.log("ECDSATableCalculator deployed at:", address(ecdsaTableCalculator));

        vm.stopBroadcast();

        /**
         *
         *                     VERIFICATION
         *
         */
        
        console.log("\n=== Deployment Verification ===");
        console.log("ECDSATableCalculator:");
        console.log("  Address:", address(ecdsaTableCalculator));
        console.log("  KeyRegistrar:", address(ecdsaTableCalculator.keyRegistrar()));
        console.log("  AllocationManager:", address(ecdsaTableCalculator.allocationManager()));
        console.log("  Lookahead Blocks:", ecdsaTableCalculator.LOOKAHEAD_BLOCKS());

        /**
         *
         *                     OUTPUT
         *
         */
        string memory parent_object = "parent object";

        // Serialize deployment information
        string memory deployment_info = "deployment_info";
        vm.serializeAddress(deployment_info, "ecdsaTableCalculator", address(ecdsaTableCalculator));
        vm.serializeAddress(deployment_info, "deployer", msg.sender);
        vm.serializeUint(deployment_info, "lookaheadBlocks", LOOKAHEAD_BLOCKS);
        string memory deployment_info_output = vm.serializeUint(deployment_info, "timestamp", block.timestamp);

        // Serialize dependencies
        string memory dependencies = "dependencies";
        vm.serializeAddress(dependencies, "keyRegistrar", address(keyRegistrar));
        string memory dependencies_output = vm.serializeAddress(dependencies, "allocationManager", address(allocationManager));

        // Serialize configuration
        string memory config = "configuration";
        vm.serializeAddress(config, "owner", owner);
        string memory config_output = vm.serializeUint(config, "chainId", chainId);

        // Combine outputs
        vm.serializeString(parent_object, "deployment_info", deployment_info_output);
        vm.serializeString(parent_object, "dependencies", dependencies_output);
        vm.serializeString(parent_object, "configuration", config_output);
        string memory finalJson = vm.serializeUint(parent_object, "chainId", chainId);

        // Write to file
        vm.writeJson(
            finalJson,
            string.concat(
                "script/output/devnet/multichain/deploy_ecdsa_table_calculator_chainid_",
                vm.toString(block.chainid),
                "_",
                vm.toString(block.timestamp),
                ".json"
            )
        );

        console.log("\nDeployment complete! Output written to file.");
        console.log("\n========================================");
        console.log("IMPORTANT: Save this address!");
        console.log("ECDSATableCalculator:", address(ecdsaTableCalculator));
        console.log("========================================\n");
    }
} 