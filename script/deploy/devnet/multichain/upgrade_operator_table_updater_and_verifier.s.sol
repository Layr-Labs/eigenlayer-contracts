// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// OpenZeppelin Contracts
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

// Core Contracts
import "src/contracts/multichain/OperatorTableUpdater.sol";
import "src/contracts/multichain/BN254CertificateVerifier.sol";
import "src/contracts/interfaces/IECDSACertificateVerifier.sol";

// Forge
import "forge-std/Script.sol";
import "forge-std/Test.sol";

// Run with:
// forge script script/deploy/devnet/multichain/upgrade_operator_table_updater_and_verifier.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run()"
contract UpgradeOperatorTableUpdaterAndVerifier is Script, Test {
    // ========== STATE VARIABLES ==========
    
    // Proxy Admin (controls upgrades)
    ProxyAdmin public proxyAdmin = ProxyAdmin(0xb5Aa3De8Fb7AD2e418da9CD70D2824FE93Ce68a5);
    
    // Existing Proxy Addresses
    OperatorTableUpdater public operatorTableUpdaterProxy = OperatorTableUpdater(0xd7230B89E5E2ed1FD068F0FF9198D7960243f12a);
    BN254CertificateVerifier public bn254CertificateVerifierProxy = BN254CertificateVerifier(0xf462d03A82C1F3496B0DFe27E978318eD1720E1f);
    
    // Dependencies for OperatorTableUpdater constructor
    // Note: For ECDSA verifier, you can use address(0) or an empty contract if not needed
    IECDSACertificateVerifier public ecdsaCertificateVerifier = IECDSACertificateVerifier(address(0));
    string public constant VERSION = "0.0.2"; // Update version for the upgrade
    
    // New Implementation contracts (will be deployed)
    OperatorTableUpdater public operatorTableUpdaterImplementation;
    BN254CertificateVerifier public bn254CertificateVerifierImplementation;
    
    function run() public {
        uint256 chainId = block.chainid;
        emit log_named_uint("Upgrading contracts on ChainID", chainId);
        
        /**
         *
         *                     DEPLOY NEW IMPLEMENTATIONS
         *
         */
        vm.startBroadcast();
        
        // Due to circular dependency, we need to deploy in a specific order:
        // 1. Deploy OperatorTableUpdater implementation using existing BN254CertificateVerifier proxy
        console.log("Deploying new OperatorTableUpdater implementation...");
        operatorTableUpdaterImplementation = new OperatorTableUpdater(
            bn254CertificateVerifierProxy,  // Use existing proxy
            ecdsaCertificateVerifier,
            VERSION
        );
        console.log("OperatorTableUpdater implementation deployed at:", address(operatorTableUpdaterImplementation));
        
        // 2. Upgrade OperatorTableUpdater proxy to new implementation
        console.log("\nUpgrading OperatorTableUpdater proxy...");
        proxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(operatorTableUpdaterProxy))),
            address(operatorTableUpdaterImplementation)
        );
        console.log("OperatorTableUpdater proxy upgraded successfully!");
        
        // 3. Now deploy BN254CertificateVerifier implementation using the upgraded OperatorTableUpdater proxy
        console.log("\nDeploying new BN254CertificateVerifier implementation...");
        bn254CertificateVerifierImplementation = new BN254CertificateVerifier(
            operatorTableUpdaterProxy  // Use the upgraded proxy
        );
        console.log("BN254CertificateVerifier implementation deployed at:", address(bn254CertificateVerifierImplementation));
        
        // 4. Upgrade BN254CertificateVerifier proxy
        console.log("\nUpgrading BN254CertificateVerifier proxy...");
        proxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(bn254CertificateVerifierProxy))),
            address(bn254CertificateVerifierImplementation)
        );
        console.log("BN254CertificateVerifier proxy upgraded successfully!");

        // 5. Set the max staleness period for the operator set
        console.log("\nSetting max staleness period for the operator set...");
        OperatorSet memory opSet = OperatorSet({
            avs: 0xD638d3779456898dff17EBFe5D62F5B7a92D61d7,
            id: 0
        });
        bn254CertificateVerifierProxy.setMaxStalenessPeriod(opSet, 0 days);
        console.log("Max staleness period set successfully!");
        
        vm.stopBroadcast();
        
        /**
         *
         *                     VERIFICATION
         *
         */
        
        // Verify the upgrades
        console.log("\n=== Upgrade Verification ===");
        
        // Check OperatorTableUpdater
        console.log("\nOperatorTableUpdater:");
        console.log("  Proxy address:", address(operatorTableUpdaterProxy));
        console.log("  New implementation:", address(operatorTableUpdaterImplementation));
        console.log("  Version:", operatorTableUpdaterProxy.version());
        
        // Verify dependencies
        console.log("  BN254CertificateVerifier:", address(operatorTableUpdaterProxy.bn254CertificateVerifier()));
        console.log("  ECDSACertificateVerifier:", address(operatorTableUpdaterProxy.ecdsaCertificateVerifier()));
        
        // Check current state
        console.log("  Latest Reference Timestamp:", operatorTableUpdaterProxy.getLatestReferenceTimestamp());
        console.log("  Global Root Confirmation Threshold:", operatorTableUpdaterProxy.globalRootConfirmationThreshold());
        
        // Check BN254CertificateVerifier
        console.log("\nBN254CertificateVerifier:");
        console.log("  Proxy address:", address(bn254CertificateVerifierProxy));
        console.log("  New implementation:", address(bn254CertificateVerifierImplementation));
        
        /**
         *
         *                     OUTPUT
         *
         */
        string memory parent_object = "parent object";
        
        // Serialize upgrade information
        string memory upgrade_info = "upgrade_info";
        vm.serializeAddress(upgrade_info, "proxyAdmin", address(proxyAdmin));
        vm.serializeAddress(upgrade_info, "operatorTableUpdaterProxy", address(operatorTableUpdaterProxy));
        vm.serializeAddress(upgrade_info, "operatorTableUpdaterNewImpl", address(operatorTableUpdaterImplementation));
        vm.serializeAddress(upgrade_info, "bn254CertificateVerifierProxy", address(bn254CertificateVerifierProxy));
        vm.serializeAddress(upgrade_info, "bn254CertificateVerifierNewImpl", address(bn254CertificateVerifierImplementation));
        vm.serializeString(upgrade_info, "version", VERSION);
        string memory upgrade_info_output = vm.serializeUint(upgrade_info, "timestamp", block.timestamp);
        
        // Serialize dependencies
        string memory dependencies = "dependencies";
        string memory dependencies_output = vm.serializeAddress(dependencies, "ecdsaCertificateVerifier", address(ecdsaCertificateVerifier));
        
        // Serialize state after upgrade
        string memory state_info = "state_after_upgrade";
        vm.serializeUint(state_info, "latestReferenceTimestamp", operatorTableUpdaterProxy.getLatestReferenceTimestamp());
        string memory state_info_output = vm.serializeUint(state_info, "globalRootConfirmationThreshold", operatorTableUpdaterProxy.globalRootConfirmationThreshold());
        
        // Combine outputs
        vm.serializeString(parent_object, "upgrade_info", upgrade_info_output);
        vm.serializeString(parent_object, "dependencies", dependencies_output);
        vm.serializeString(parent_object, "state_after_upgrade", state_info_output);
        string memory finalJson = vm.serializeUint(parent_object, "chainId", chainId);
        
        // Write to file
        vm.writeJson(
            finalJson,
            string.concat(
                "script/output/devnet/multichain/upgrade_operator_table_updater_and_verifier_chainid_",
                vm.toString(block.chainid),
                "_",
                vm.toString(block.timestamp),
                ".json"
            )
        );
        
        console.log("\nUpgrade complete! Output written to file.");
    }
} 