// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import "../TestUtils.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";
import {EmissionsController} from "src/contracts/core/EmissionsController.sol";

/// Purpose: Deploy implementations for the v1.14.0 EigenPod disable
/// and emissions burn release.
contract DeployImplementations is CoreContractsDeployer {
    using Env for *;

    /// @notice Catch up environments that are still missing the v1.12.0 incentive council deployment.
    function _needsIncentiveCouncilUpgrade() internal view returns (bool) {
        return !Env._versionGte(Env.envVersion(), "1.12.0");
    }

    function _runAsEOA() internal virtual override {
        vm.startBroadcast();

        if (_needsIncentiveCouncilUpgrade()) {
            _deployEmissionsControllerProxy();
            zUpdateUint32("REWARDS_COORDINATOR_MAX_REWARDS_DURATION", 63_072_000);
            deployRewardsCoordinator();
        }

        deployDelegationManager();
        deployEigenPodManager();
        deployEigenPod();
        deployEmissionsController();

        vm.stopBroadcast();
    }

    function _deployEmissionsControllerProxy() internal onlyEOA {
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

    function testScript() public virtual onlyIfUpgradeRequired("1.14.0") {
        runAsEOA();

        TestUtils.validateDelegationManagerImmutables(Env.impl.delegationManager());
        TestUtils.validateDelegationManagerInitialized(Env.impl.delegationManager());
        TestUtils.validateDelegationManagerVersion();

        TestUtils.validateEigenPodManagerImmutables(Env.impl.eigenPodManager());
        TestUtils.validateEigenPodManagerInitialized(Env.impl.eigenPodManager());
        TestUtils.validateEigenPodImmutables(Env.impl.eigenPod());

        TestUtils.validateEmissionsControllerImmutables(Env.impl.emissionsController());
        TestUtils.validateEmissionsControllerInitialized(Env.impl.emissionsController());

        if (_needsIncentiveCouncilUpgrade()) {
            TestUtils.validateRewardsCoordinatorImmutables(Env.impl.rewardsCoordinator());
            TestUtils.validateRewardsCoordinatorInitialized(Env.impl.rewardsCoordinator());
        }
    }
}
