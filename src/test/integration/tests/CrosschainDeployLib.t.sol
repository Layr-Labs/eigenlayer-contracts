// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "script/releases/CrosschainDeployLib.sol";
import "src/test/integration/IntegrationDeployer.t.sol";

/// @notice Sanity check for the crosschain deploy lib
/// @dev We use the integration testing suite as it has RPC urls in our CI
contract Integration_CrosschainDeployLibTest_Testnet is IntegrationDeployer {
    // Addresses, taken from zeus
    address hoodiDeployer = 0xA591635DE4C254BD3fa9C9Db9000eA6488344C28;
    address hoodiExpectedEmptyContract = 0xa6192470D7D4c39f8F392167ADde283F60b34E15;
    address hoodiExpectedProxyContract = 0xB99CC53e8db7018f557606C2a5B066527bF96b26; // Using type(TaskMailbox).name as name
    address mainnetDeployer = 0xa3053EF25F1F7d9D55a7655372B8a31D0f40eCA9;
    address mainnetExpectedEmptyContract = 0x70323BC7261136A002ab392F921D67ea52096dCf;
    address mainnetExpectedProxyContract = 0x132b466d9d5723531F68797519DfED701aC2C749;

    // Forks
    uint mainnetFork;
    uint hoodiFork;

    function setUp() public virtual override {
        // Skip if we're not on a fork test profile
        if (!isForktest()) vm.skip(true);

        mainnetFork = vm.createSelectFork(vm.envString("RPC_MAINNET"), 22_819_288);
        hoodiFork = vm.createSelectFork(vm.envString("RPC_HOODI"), 1_291_826);
    }

    /// @dev We test hoodi and mainnet in same function due to flakiness of `vm.setEnv`
    function test_contracts_hoodi() public {
        // Hoodi Fork
        vm.selectFork(hoodiFork);
        vm.setEnv("ZEUS_ENV", "preprod-hoodi");

        // Deploy empty contract and proxy contract
        cheats.startPrank(hoodiDeployer);
        address emptyContract = CrosschainDeployLib.deployEmptyContract(hoodiDeployer);
        address proxyContract = address(CrosschainDeployLib.deployCrosschainProxy(hoodiDeployer, emptyContract, type(TaskMailbox).name));
        cheats.stopPrank();

        assertEq(emptyContract, hoodiExpectedEmptyContract, "emptyContract != hoodiExpectedEmptyContract");
        assertEq(proxyContract, hoodiExpectedProxyContract, "proxyContract != hoodiExpectedProxyContract");

        // Mainnet fork
        vm.selectFork(mainnetFork);
        vm.setEnv("ZEUS_ENV", "mainnet");

        // Deploy empty contract and proxy contract
        cheats.startPrank(mainnetDeployer);
        emptyContract = CrosschainDeployLib.deployEmptyContract(mainnetDeployer);
        proxyContract = address(CrosschainDeployLib.deployCrosschainProxy(mainnetDeployer, emptyContract, type(TaskMailbox).name));
        cheats.stopPrank();

        assertEq(emptyContract, mainnetExpectedEmptyContract, "emptyContract != mainnetExpectedEmptyContract");
        assertEq(proxyContract, mainnetExpectedProxyContract, "proxyContract != mainnetExpectedProxyContract");
    }

    function test_address_prediction() public {
        // Get on mainnet fork
        vm.selectFork(mainnetFork);

        // Empty Contract
        bytes32 initCodeHash = keccak256(CrosschainDeployLib.EMPTY_CONTRACT_CREATION_CODE_PRODUCTION);
        address expectedEmptyContract =
            CrosschainDeployLib.computeCrosschainAddress(mainnetDeployer, initCodeHash, type(EmptyContract).name);
        assertEq(expectedEmptyContract, mainnetExpectedEmptyContract, "expectedEmptyContract != mainnetExpectedEmptyContract");

        // Proxy Contract
        // Recreate the creationCode, because using vm.setEnv() is flaky in multiple functions in the same test contract
        initCodeHash = keccak256(
            abi.encodePacked(
                CrosschainDeployLib.PROXY_CONTRACT_CREATION_CODE_PRODUCTION, abi.encode(mainnetExpectedEmptyContract, mainnetDeployer, "")
            )
        );
        address expectedProxyContract = CrosschainDeployLib.computeCrosschainAddress(mainnetDeployer, initCodeHash, type(TaskMailbox).name);
        assertEq(expectedProxyContract, mainnetExpectedProxyContract, "expectedEmptyContract != mainnetExpectedProxyContract");
    }
}
