// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// OpenZeppelin Contracts
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

// Core Contracts
import "src/contracts/permissions/PauserRegistry.sol";
import "src/contracts/core/AllocationManager.sol";
import "src/contracts/permissions/PermissionController.sol";

// Multichain Contracts
import "src/contracts/multichain/CrossChainRegistry.sol";
import "src/contracts/multichain/BN254TableCalculator.sol";
import "src/contracts/permissions/KeyRegistrar.sol";
import "src/test/mocks/EmptyContract.sol";

// Forge
import "forge-std/Script.sol";
import "forge-std/Test.sol";

// forge script script/deploy/devnet/deploy_multichain_l1.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast --sig "run()" --verify $ETHERSCAN_API_KEY
contract DeployMultichain_L1 is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    address owner = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;

    // EigenLayer Contracts on PREPROD
    EmptyContract public emptyContract;
    AllocationManager public allocationManager = AllocationManager(0xFdD5749e11977D60850E06bF5B13221Ad95eb6B4);
    PermissionController public permissionController = PermissionController(0xa2348c77802238Db39f0CefAa500B62D3FDD682b);
    PauserRegistry public pauserRegistry = PauserRegistry(0x50712285cE831a6B9a11214A430f28999A5b4DAe);

    // Multichain Contracts
    ProxyAdmin public proxyAdmin;
    KeyRegistrar public keyRegistrar;
    KeyRegistrar public keyRegistrarImplementation;
    CrossChainRegistry public crossChainRegistry;
    CrossChainRegistry public crossChainRegistryImplementation;
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

        // First, deploy the *proxy* contracts, using the *empty contract* as inputs
        // Key Registrar
        keyRegistrar =
            KeyRegistrar(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));

        // Cross Chain Registry
        crossChainRegistry = CrossChainRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );

        // BN254 Table Calculator

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        keyRegistrarImplementation = new KeyRegistrar(permissionController, allocationManager, "9.9.9");
        crossChainRegistryImplementation =
            new CrossChainRegistry(allocationManager, keyRegistrar, permissionController, pauserRegistry, "9.9.9");

        // Third, upgrade the proxies to point to the new implementations
        proxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(keyRegistrar))), address(keyRegistrarImplementation)
        );
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(crossChainRegistry))),
            address(crossChainRegistryImplementation),
            abi.encodeWithSelector(CrossChainRegistry.initialize.selector, owner, 0)
        );

        // Fourth, deploy the non-upgradeable contracts
        bn254TableCalculator = new BN254TableCalculator(keyRegistrar, allocationManager, 100);

        // Transfer ownership to the 0xDA address
        proxyAdmin.transferOwnership(owner);

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
        vm.serializeAddress(deployed_addresses, "crossChainRegistry", address(crossChainRegistry));
        vm.serializeAddress(
            deployed_addresses, "crossChainRegistryImplementation", address(crossChainRegistryImplementation)
        );
        string memory deployed_addresses_output =
            vm.serializeAddress(deployed_addresses, "bn254TableCalculator", address(bn254TableCalculator));

        string memory finalJson = vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);

        vm.writeJson(finalJson, "script/output/devnet/multichain/deploy_multichain_l1.json");
    }
}
