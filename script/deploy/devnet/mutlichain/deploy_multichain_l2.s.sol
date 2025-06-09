// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// OpenZeppelin Contracts
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

// Core Contracats
import "src/contracts/multichain/OperatorTableUpdater.sol";
import "src/contracts/multichain/BN254CertificateVerifier.sol";
import "src/contracts/interfaces/IECDSATableCalculator.sol";
import "src/test/mocks/EmptyContract.sol";

// Import the necessary types
import "src/contracts/interfaces/IBN254TableCalculator.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/libraries/BN254.sol";

// Forge
import "forge-std/Script.sol";
import "forge-std/Test.sol";

// forge script script/deploy/devnet/mutlichain/deploy_multichain_l2.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast --sig "run()" --verify $ETHERSCAN_API_KEY
contract DeployMultichain_L2 is Script, Test {
    address globalOwner = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;

    // Multichain Contracts
    EmptyContract public emptyContract;
    ProxyAdmin public proxyAdmin;
    OperatorTableUpdater public operatorTableUpdater;
    OperatorTableUpdater public operatorTableUpdaterImplementation;
    BN254CertificateVerifier public bn254CertificateVerifier;
    BN254CertificateVerifier public bn254CertificateVerifierImplementation;

    function run() public {
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        /**
         *
         *                     CONTRACT DEPLOYMENT
         *
         */
        vm.startBroadcast();

        emptyContract = new EmptyContract();
        proxyAdmin = new ProxyAdmin();

        // First, deploy the *proxy* contracts, using the *empty contract* as inputs
        // Operator Table Updater
        operatorTableUpdater = OperatorTableUpdater(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );

        // BN254 Certificate Verifier
        bn254CertificateVerifier = BN254CertificateVerifier(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        operatorTableUpdaterImplementation = new OperatorTableUpdater(
            bn254CertificateVerifier, IECDSACertificateVerifier(address(emptyContract)), "0.0.1"
        );
        bn254CertificateVerifierImplementation = new BN254CertificateVerifier(operatorTableUpdater);

        // Third, upgrade the proxies to point to the new implementations
        proxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(operatorTableUpdater))),
            address(operatorTableUpdaterImplementation)
        );

        proxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(bn254CertificateVerifier))),
            address(bn254CertificateVerifierImplementation)
        );

        proxyAdmin.transferOwnership(globalOwner);
        vm.stopBroadcast();

        /**
         *
         *                     INITIALIZATION
         *
         */

        // Read the JSON data
        string memory globalRootOperatorInfoPath = "script/output/devnet/multichain/globalRootOperatorInfo.json";
        string memory jsonData = vm.readFile(globalRootOperatorInfoPath);

        // Parse operatorSetInfo
        uint256 numOperators = vm.parseJsonUint(jsonData, ".operatorSetInfo.numOperators");
        bytes32 operatorInfoTreeRoot = vm.parseJsonBytes32(jsonData, ".operatorSetInfo.operatorInfoTreeRoot");

        // Parse APK (aggregatePubkey)
        uint256 apkX = vm.parseJsonUint(jsonData, ".operatorSetInfo.apk.x");
        uint256 apkY = vm.parseJsonUint(jsonData, ".operatorSetInfo.apk.y");
        BN254.G1Point memory aggregatePubkey = BN254.G1Point({X: apkX, Y: apkY});

        // Parse totalWeights array
        uint256[] memory totalWeights = vm.parseJsonUintArray(jsonData, ".operatorSetInfo.totalWeights");

        // Create BN254OperatorSetInfo struct
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory globalRootConfirmerSetInfo = IBN254TableCalculatorTypes
            .BN254OperatorSetInfo({
            operatorInfoTreeRoot: operatorInfoTreeRoot,
            numOperators: numOperators,
            aggregatePubkey: aggregatePubkey,
            totalWeights: totalWeights
        });

        // Parse operatorSetConfig
        address owner = vm.parseJsonAddress(jsonData, ".operatorSetConfig.owner");
        uint32 maxStalenessPeriod = uint32(vm.parseJsonUint(jsonData, ".operatorSetConfig.maxStalenessPeriod"));

        ICrossChainRegistryTypes.OperatorSetConfig memory globalRootConfirmerSetConfig =
            ICrossChainRegistryTypes.OperatorSetConfig({owner: owner, maxStalenessPeriod: maxStalenessPeriod});

        // Create the global root confirmer OperatorSet
        // For devnet, we'll use address(this) as the AVS and id 0
        OperatorSet memory globalRootConfirmerSet = OperatorSet({avs: owner, id: 0});

        // Set confirmation threshold (e.g., 100% = 10_000 bps)
        uint16 globalRootConfirmationThreshold = 10_000;

        // Use current block timestamp as reference timestamp
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Initialize the operatorTableUpdater
        vm.startBroadcast();
        operatorTableUpdater.initialize(
            globalOwner,
            globalRootConfirmerSet,
            globalRootConfirmationThreshold,
            referenceTimestamp,
            globalRootConfirmerSetInfo,
            globalRootConfirmerSetConfig
        );
        vm.stopBroadcast();

        /**
         *
         *                     OUTPUT
         *
         */
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(deployed_addresses, "emptyContract", address(emptyContract));
        vm.serializeAddress(deployed_addresses, "proxyAdmin", address(proxyAdmin));
        vm.serializeAddress(deployed_addresses, "operatorTableUpdater", address(operatorTableUpdater));
        vm.serializeAddress(
            deployed_addresses, "operatorTableUpdaterImplementation", address(operatorTableUpdaterImplementation)
        );
        vm.serializeAddress(deployed_addresses, "bn254CertificateVerifier", address(bn254CertificateVerifier));
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "bn254CertificateVerifierImplementation",
            address(bn254CertificateVerifierImplementation)
        );

        string memory finalJson = vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
        vm.writeJson(
            finalJson,
            string.concat(
                "script/output/devnet/multichain/deploy_multichain_l2_chainid_", vm.toString(block.chainid), ".json"
            )
        );
    }
}
