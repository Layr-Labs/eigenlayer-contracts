// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "script/releases/CrosschainDeployLib.sol";
import "src/test/integration/IntegrationDeployer.t.sol";

/// @notice Sanity check for the crosschain deploy lib
/// @dev We use the integration testing suite as it has RPC urls in our CI
contract Integration_CrosschainDeployLibTest is IntegrationDeployer {
    address public constant DEPLOYER = 0xC10E5F3AF465Fe85A7077390797dc5ae89DAB9F1;
    address public constant EXPECTED_EMPTY_CONTRACT = 0x2e2F802C33b5725Cb2Aa2d805cb874902B695796;
    address public constant EXPECTED_PROXY = 0x4368AC3537895e2704dfCba86f6185166BDf4aE4;

    function setUp() public override {
        super.setUp();
        if (isForktest()) return;
    }

    function test_mainnet() public {
        // Deploy empty contract
        address mainnetEmptyContract = CrosschainDeployLib.deployEmptyContract(DEPLOYER);
        assertEq(mainnetEmptyContract, EXPECTED_EMPTY_CONTRACT, "mainnetEmptyContract != EXPECTED_EMPTY_CONTRACT");

        // Deploy proxy
        address mainnetProxy = address(CrosschainDeployLib.deployCrosschainProxy(DEPLOYER, mainnetEmptyContract, "ExampleContract"));
        assertEq(mainnetProxy, EXPECTED_PROXY, "mainnetProxy != EXPECTED_PROXY");
    }

    function test_base() public {
        cheats.createSelectFork(cheats.rpcUrl("base"), 37_783_655);

        // Deploy empty contract
        address baseEmptyContract = CrosschainDeployLib.deployEmptyContract(DEPLOYER);
        assertEq(baseEmptyContract, EXPECTED_EMPTY_CONTRACT, "baseEmptyContract != EXPECTED_EMPTY_CONTRACT");

        // Deploy proxy
        address baseProxy = address(CrosschainDeployLib.deployCrosschainProxy(DEPLOYER, baseEmptyContract, "ExampleContract"));
        assertEq(baseProxy, EXPECTED_PROXY, "baseProxy != EXPECTED_PROXY");
    }
}
