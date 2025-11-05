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

        // Test empty contract deployment
        uint baseFork = vm.createSelectFork(cheats.rpcUrl("base"), 37_783_655);
        address baseExpected = CrosschainDeployLib.deployEmptyContract(deployer);
        uint mainnetFork = vm.createSelectFork(cheats.rpcUrl("mainnet"), 22_819_288);
        address mainnetExpected = CrosschainDeployLib.deployEmptyContract(deployer);
        assertEq(baseExpected, mainnetExpected, "baseExpected != mainnetExpected");

        // Test proxy deployment
        vm.selectFork(baseFork);
        address baseProxy = address(CrosschainDeployLib.deployCrosschainProxy(deployer, baseExpected, "ExampleContract"));
        vm.selectFork(mainnetFork);
        address mainnetProxy = address(CrosschainDeployLib.deployCrosschainProxy(deployer, mainnetExpected, "ExampleContract"));
        assertEq(baseProxy, mainnetProxy, "baseProxy != mainnetProxy");

        // Test address prediction
        assertEq(
            CrosschainDeployLib.computeCrosschainAddress(deployer, keccak256(type(EmptyContract).creationCode), "EmptyContract"),
            baseExpected
        );
        assertEq(CrosschainDeployLib.computeCrosschainUpgradeableProxyAddress(deployer, baseExpected, "ExampleContract"), baseProxy);
    }
}
