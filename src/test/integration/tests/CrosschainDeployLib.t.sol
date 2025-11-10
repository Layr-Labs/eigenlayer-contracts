// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "script/releases/CrosschainDeployLib.sol";
import "src/test/integration/IntegrationDeployer.t.sol";

/// @notice Sanity check for the crosschain deploy lib
/// @dev We use the integration testing suite as it has RPC urls in our CI
contract Integration_CrosschainDeployLibTest is IntegrationDeployer {
    function test_SameAddressEveryChain() public {
        // Skip if we're not on a fork test profile
        if (!isForktest()) return;

        address deployer = 0xC10E5F3AF465Fe85A7077390797dc5ae89DAB9F1;

        vm.startPrank(deployer);

        // Set env for tests to pass
        vm.setEnv("ZEUS_ENV", "mainnet");

        // Test empty contract deployment
        uint hoodiFork = vm.createSelectFork(vm.envString("RPC_HOODI"), 1_549_268);
        address hoodiExpected = CrosschainDeployLib.deployEmptyContract(deployer);
        uint mainnetFork = vm.createSelectFork(vm.envString("RPC_MAINNET"), 22_819_288);
        address mainnetExpected = CrosschainDeployLib.deployEmptyContract(deployer);
        assertEq(hoodiExpected, mainnetExpected, "hoodiExpected != mainnetExpected");

        // Test proxy deployment
        vm.selectFork(hoodiFork);
        address hoodiProxy = address(CrosschainDeployLib.deployCrosschainProxy(deployer, hoodiExpected, "ExampleContract"));
        vm.selectFork(mainnetFork);
        address mainnetProxy = address(CrosschainDeployLib.deployCrosschainProxy(deployer, mainnetExpected, "ExampleContract"));
        assertEq(hoodiProxy, mainnetProxy, "hoodiProxy != mainnetProxy");

        // Test address prediction
        assertEq(
            CrosschainDeployLib.computeCrosschainAddress(deployer, keccak256(type(EmptyContract).creationCode), "EmptyContract"),
            hoodiExpected
        );
        assertEq(CrosschainDeployLib.computeCrosschainUpgradeableProxyAddress(deployer, hoodiExpected, "ExampleContract"), hoodiProxy);
    }
}
