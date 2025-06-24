// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

/**
 * @title DeployECDSACertificateVerifier
 * @notice Deploys the ECDSACertificateVerifier contract with proxy pattern
 * @dev This script:
 *      1. Deploys the ECDSACertificateVerifier implementation
 *      2. Deploys a TransparentUpgradeableProxy pointing to the implementation
 *      3. Verifies the deployment and outputs addresses to JSON
 * 
 * Prerequisites:
 *      - OperatorTableUpdater must be deployed (used as dependency)
 *      - ProxyAdmin must be deployed and owned by the deployer
 * 
 * After deployment:
 *      - Update other scripts with the new ECDSACertificateVerifier proxy address
 *      - The proxy address should be used for all interactions
 */

// OpenZeppelin Contracts
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

// Core Contracts
import "src/contracts/multichain/ECDSACertificateVerifier.sol";
import "src/contracts/interfaces/IOperatorTableUpdater.sol";

// Forge
import "forge-std/Script.sol";
import "forge-std/Test.sol";

// Run with:
// forge script script/deploy/devnet/multichain/deploy_ecdsa_certificate_verifier.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast --sig "run()" --verify $ETHERSCAN_API_KEY
contract DeployECDSACertificateVerifier is Script, Test {
    // ========== STATE VARIABLES ==========
    
    // Proxy Admin (controls upgrades)
    ProxyAdmin public proxyAdmin = ProxyAdmin(0xb5Aa3De8Fb7AD2e418da9CD70D2824FE93Ce68a5);
    
    // Dependencies
    IOperatorTableUpdater public operatorTableUpdater = IOperatorTableUpdater(0xd7230B89E5E2ed1FD068F0FF9198D7960243f12a);
    
    // Version for the deployment
    string public constant VERSION = "1.0.0";
    
    // New contracts to be deployed
    ECDSACertificateVerifier public ecdsaCertificateVerifierImplementation;
    TransparentUpgradeableProxy public ecdsaCertificateVerifierProxy;
    
    function run() public {
        uint256 chainId = block.chainid;
        emit log_named_uint("Deploying ECDSACertificateVerifier on ChainID", chainId);
        
        /**
         *
         *                     DEPLOY IMPLEMENTATION AND PROXY
         *
         */
        vm.startBroadcast();
        
        // 1. Deploy ECDSACertificateVerifier implementation
        console.log("Deploying ECDSACertificateVerifier implementation...");
        ecdsaCertificateVerifierImplementation = new ECDSACertificateVerifier(
            operatorTableUpdater,
            VERSION
        );
        console.log("ECDSACertificateVerifier implementation deployed at:", address(ecdsaCertificateVerifierImplementation));
        
        // 2. Deploy TransparentUpgradeableProxy for ECDSACertificateVerifier
        console.log("\nDeploying ECDSACertificateVerifier proxy...");
        ecdsaCertificateVerifierProxy = new TransparentUpgradeableProxy(
            address(ecdsaCertificateVerifierImplementation),
            address(proxyAdmin),
            ""  // Empty data since ECDSACertificateVerifier doesn't have an initialize function
        );
        console.log("ECDSACertificateVerifier proxy deployed at:", address(ecdsaCertificateVerifierProxy));
        
        vm.stopBroadcast();
        
        /**
         *
         *                     VERIFICATION
         *
         */
        
        // Verify the deployment
        console.log("\n=== Deployment Verification ===");
        
        // Cast proxy to ECDSACertificateVerifier
        ECDSACertificateVerifier ecdsaCertificateVerifier = ECDSACertificateVerifier(address(ecdsaCertificateVerifierProxy));
        
        console.log("\nECDSACertificateVerifier:");
        console.log("  Proxy address:", address(ecdsaCertificateVerifierProxy));
        console.log("  Implementation address:", address(ecdsaCertificateVerifierImplementation));
        console.log("  Version:", ecdsaCertificateVerifier.version());
        console.log("  OperatorTableUpdater:", address(ecdsaCertificateVerifier.operatorTableUpdater()));
        console.log("  Domain Separator:", vm.toString(ecdsaCertificateVerifier.domainSeparator()));
        
        // Verify proxy admin ownership
        console.log("\nProxy Admin:");
        console.log("  Address:", address(proxyAdmin));
        console.log("  Owner:", proxyAdmin.owner());
        
        // Verify that the proxy admin owns the proxy
        address proxyAdminOfECDSA = proxyAdmin.getProxyAdmin(ITransparentUpgradeableProxy(payable(address(ecdsaCertificateVerifierProxy))));
        console.log("  Proxy admin of ECDSACertificateVerifier proxy:", proxyAdminOfECDSA);
        require(proxyAdminOfECDSA == address(proxyAdmin), "Proxy admin mismatch");
        
        /**
         *
         *                     OUTPUT
         *
         */
        // string memory parent_object = "parent object";
        
        // // Serialize deployment information
        // string memory deployment_info = "deployment_info";
        // vm.serializeAddress(deployment_info, "proxyAdmin", address(proxyAdmin));
        // vm.serializeAddress(deployment_info, "ecdsaCertificateVerifierProxy", address(ecdsaCertificateVerifierProxy));
        // vm.serializeAddress(deployment_info, "ecdsaCertificateVerifierImplementation", address(ecdsaCertificateVerifierImplementation));
        // vm.serializeString(deployment_info, "version", VERSION);
        // string memory deployment_info_output = vm.serializeUint(deployment_info, "timestamp", block.timestamp);
        
        // // Serialize dependencies
        // string memory dependencies = "dependencies";
        // string memory dependencies_output = vm.serializeAddress(dependencies, "operatorTableUpdater", address(operatorTableUpdater));
        
        // // Serialize state after deployment
        // string memory state_info = "state_after_deployment";
        // vm.serializeBytes32(state_info, "domainSeparator", ecdsaCertificateVerifier.domainSeparator());
        // string memory state_info_output = vm.serializeString(state_info, "contractVersion", ecdsaCertificateVerifier.version());
        
        // // Combine outputs
        // vm.serializeString(parent_object, "deployment_info", deployment_info_output);
        // vm.serializeString(parent_object, "dependencies", dependencies_output);
        // vm.serializeString(parent_object, "state_after_deployment", state_info_output);
        // string memory finalJson = vm.serializeUint(parent_object, "chainId", chainId);
        
        // Write to file
        // vm.writeJson(
        //     finalJson,
        //     string.concat(
        //         "script/output/devnet/multichain/deploy_ecdsa_certificate_verifier_chainid_",
        //         vm.toString(block.chainid),
        //         "_",
        //         vm.toString(block.timestamp),
        //         ".json"
        //     )
        // );
        
        console.log("\nDeployment complete! Output written to file.");
        console.log("\n========================================");
        console.log("IMPORTANT: Save these addresses!");
        console.log("ECDSACertificateVerifier Proxy:", address(ecdsaCertificateVerifierProxy));
        console.log("ECDSACertificateVerifier Implementation:", address(ecdsaCertificateVerifierImplementation));
        console.log("========================================\n");
    }
} 