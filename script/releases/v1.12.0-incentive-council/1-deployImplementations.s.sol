// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import "../TestUtils.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";
import {EmissionsController} from "src/contracts/core/EmissionsController.sol";

/// Steps:
/// 1. Deploy new RewardsCoordinator implementation (fees added).
/// 2. Deploy new EmissionsController implementation.
/// 3. Deploy EmissionsController proxy (empty implementation).
/// zeus run --env <env> --command "forge script script/releases/v9.9.9-incentive-council/1-deployImplementations.s.sol --rpc-url <rpc_url> --sig 'testScript()'"
contract DeployImplementations is CoreContractsDeployer {
    using Env for *;

    function _runAsEOA() internal virtual override {
        vm.startBroadcast();

        deployRewardsCoordinator();
        deployEmissionsController();
        deployEmissionsControllerProxy();

        vm.stopBroadcast();
    }

    function deployEmissionsControllerProxy() internal onlyEOA {
        deployProxy({
            name: type(EmissionsController).name,
            deployedTo: address(
                ITransparentUpgradeableProxy(
                    payable(new TransparentUpgradeableProxy({
                            _logic: address(Env.impl.emptyContract()),
                            admin_: Env.proxyAdmin(),
                            _data: ""
                        }))
                )
            )
        });
    }

    function testScript() public virtual {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        // Deploy implementations as EOA
        runAsEOA();

        // Check that the implementations have initialization disabled.
        TestUtils.validateEmissionsControllerInitialized(Env.impl.emissionsController());
        TestUtils.validateRewardsCoordinatorConstructor(Env.impl.rewardsCoordinator());

        // Check that the implementations have correct constructor/immutables.
        TestUtils.validateEmissionsControllerImmutables(Env.impl.emissionsController());
        TestUtils.validateRewardsCoordinatorImmutables(Env.impl.rewardsCoordinator());

        // Only validate proxy admins - skip full impl validation since preprod may not have v1.9.0+
        TestUtils.validateProxyAdmins();
        TestUtils.validateImplConstructors();
        TestUtils.validateImplsNotInitializable();
    }
}

