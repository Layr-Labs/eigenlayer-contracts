// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

// Core Contracts
import "../../../src/contracts/core/AllocationManager.sol";
import "../../../src/contracts/permissions/PermissionController.sol";

// Multichain Contracts
import "../../../src/contracts/interfaces/ICrossChainRegistry.sol";
import "../../../src/contracts/multichain/BN254TableCalculator.sol";
import "../../../src/contracts/permissions/KeyRegistrar.sol";
import "../../../src/test/mocks/EmptyContract.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// forge script script/deploy/devnet/deploy_multichain_l1.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast --sig "run()"
contract DeployMultichain_L1 is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    // EigenLayer Contracts on PREPROD
    EmptyContract public emptyContract;
    AllocationManager public allocationManager = AllocationManager(0xFdD5749e11977D60850E06bF5B13221Ad95eb6B4);
    PermissionController public permissionController = PermissionController(0xa2348c77802238Db39f0CefAa500B62D3FDD682b);

    // Multichain Contracts
    ProxyAdmin public proxyAdmin;
    KeyRegistrar public keyRegistrar;
    KeyRegistrar public keyRegistrarImplementation;
    ICrossChainRegistry public crossChainRegistry;
    ICrossChainRegistry public crossChainRegistryImplementation;
    BN254TableCalculator public bn254TableCalculator;

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

        // Key Registrar
        keyRegistrar = KeyRegistrar(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));

        // Cross Chain Registry
        // crossChainRegistry = ICrossChainRegistry(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));

        // BN254 Table Calculator

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        keyRegistrarImplementation = new KeyRegistrar(permissionController, allocationManager, "9.9.9");
        // crossChainRegistryImplementation = new CrossChainRegistry(IKeyRegistrar(address(keyRegistrar)), IBN254TableCalculator(address(bn254TableCalculator)));

        // Third, upgrade the proxies to point to the new implementations
        proxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(keyRegistrar))), address(keyRegistrarImplementation));
        // proxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(crossChainRegistry))), address(crossChainRegistryImplementation));

        // Fourth, deploy the non-upgradeable contracts
        bn254TableCalculator = new BN254TableCalculator(keyRegistrar, allocationManager, 100);

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
        vm.serializeAddress(deployed_addresses, "keyRegistrar", address(keyRegistrar));
        vm.serializeAddress(deployed_addresses, "keyRegistrarImplementation", address(keyRegistrarImplementation));
        string memory deployed_addresses_output = vm.serializeAddress(deployed_addresses, "bn254TableCalculator", address(bn254TableCalculator));

        string memory finalJson = vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);

        vm.writeJson(finalJson, "script/output/devnet/deploy_multichain_l1.json");
    }
}