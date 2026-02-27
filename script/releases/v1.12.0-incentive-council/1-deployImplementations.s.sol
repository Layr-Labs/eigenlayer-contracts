// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import "../TestUtils.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";
import {EmissionsController} from "src/contracts/core/EmissionsController.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

/// Purpose: use an EOA to deploy all new/updated contracts for Duration Vault and Incentive Council features.
/// Contracts deployed:
/// /// Core
/// - RewardsCoordinator (rewards v2.2 + protocol fees)
/// - StrategyManager (updated with beforeAddShares/beforeRemoveShares hooks)
/// - EmissionsController implementation + proxy
/// /// Strategies
/// - EigenStrategy, StrategyBase, StrategyBaseTVLLimits (updated for Duration Vault hooks)
/// - StrategyFactory (updated with duration vault beacon support)
/// - DurationVaultStrategy (new beacon implementation + beacon)
contract DeployImplementations is CoreContractsDeployer {
    using Env for *;

    function _runAsEOA() internal virtual override {
        vm.startBroadcast();

        /// core/
        // Deploy EmissionsController proxy first (RewardsCoordinator needs its address as immutable)
        // Preprod already has EmissionsController proxy deployed
        if (!(Env._strEq(Env.envVersion(), "1.12.0"))) {
            deployEmissionsControllerProxy();
        }
        deployEmissionsController();

        // Update the MAX_REWARDS_DURATION environment variable before deploying RewardsCoordinator
        // 63072000s = 730 days = 2 years
        zUpdateUint32("REWARDS_COORDINATOR_MAX_REWARDS_DURATION", 63_072_000);
        deployRewardsCoordinator();
        deployStrategyManager();

        /// strategies/
        deployEigenStrategy();
        deployStrategyBase();
        deployStrategyBaseTVLLimits();

        // Deploy DurationVaultStrategy implementation and beacon
        DurationVaultStrategy durationVaultImpl = deployDurationVaultStrategy();
        _deployDurationVaultBeacon(address(durationVaultImpl));

        // Deploy StrategyFactory (requires beacons to exist for immutable constructor args)
        deployStrategyFactory();

        vm.stopBroadcast();
    }

    function _deployDurationVaultBeacon(
        address implementation
    ) internal onlyEOA {
        UpgradeableBeacon beacon = new UpgradeableBeacon(implementation);
        beacon.transferOwnership(Env.executorMultisig());
        deployBeacon({name: type(DurationVaultStrategy).name, deployedTo: address(beacon)});
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

        runAsEOA();

        // Validate incentive council implementations
        TestUtils.validateEmissionsControllerInitialized(Env.impl.emissionsController());
        TestUtils.validateRewardsCoordinatorConstructor(Env.impl.rewardsCoordinator());
        TestUtils.validateEmissionsControllerImmutables(Env.impl.emissionsController());
        TestUtils.validateRewardsCoordinatorImmutables(Env.impl.rewardsCoordinator());

        // Validate duration vault implementations
        TestUtils.validateDurationVaultStrategyImplConstructors();

        TestUtils.validateProxyAdmins();
        TestUtils.validateImplConstructors();
        TestUtils.validateImplsNotInitializable();
    }
}

