// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import "../TestUtils.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";
import {EmissionsController} from "src/contracts/core/EmissionsController.sol";

contract DeployImplementations is CoreContractsDeployer {
    using Env for *;

    /// @notice `testnet-hoodi` is the only env still on a pre-v1.12.0 deployment.
    ///      All other envs (mainnet, base, sepolia, preprod-hoodi, base-sepolia) already
    ///      have these contracts. so this to avoid ci failing on testnet-hoodi
    function _needsIncentiveCouncilUpgrade() internal view returns (bool) {
        return Env._strEq(Env.env(), "testnet-hoodi");
    }

    function _runAsEOA() internal virtual override {
        vm.startBroadcast();

        // v1.13.0 changes
        deployStrategyManager();
        deployStrategyFactory();
        deployDurationVaultStrategy();

        // Catch up with v1.12.0 incentive council changes if env hasn't been upgraded
        if (_needsIncentiveCouncilUpgrade()) {
            deployEmissionsControllerProxy();
            deployEmissionsController();
            // Update MAX_REWARDS_DURATION before deploying RewardsCoordinator (63072000s = 730 days)
            zUpdateUint32("REWARDS_COORDINATOR_MAX_REWARDS_DURATION", 63_072_000);
            deployRewardsCoordinator();
        }

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

        runAsEOA();

        // StrategyManager validations
        TestUtils.validateStrategyManagerImmutables(Env.impl.strategyManager());
        TestUtils.validateStrategyManagerInitialized(Env.impl.strategyManager());
        require(
            keccak256(bytes(Env.impl.strategyManager().version())) == keccak256(bytes("1.13.0")),
            "strategyManager version must be 1.13.0"
        );
        TestUtils.validateStrategyManagerSlashResolutionDelay(Env.impl.strategyManager());

        // StrategyFactory validations
        TestUtils.validateStrategyFactoryImmutables(Env.impl.strategyFactory());
        TestUtils.validateStrategyFactoryInitialized(Env.impl.strategyFactory());

        // DurationVaultStrategy validations
        TestUtils.validateDurationVaultStrategyImmutables(Env.impl.durationVaultStrategy());

        // Incentive council catch-up validations (only when env was pre-v1.12.0)
        if (address(Env.proxy.emissionsController()).code.length > 0) {
            TestUtils.validateEmissionsControllerImmutables(Env.impl.emissionsController());
            TestUtils.validateRewardsCoordinatorImmutables(Env.impl.rewardsCoordinator());
            TestUtils.validateRewardsCoordinatorInitialized(Env.impl.rewardsCoordinator());
        }
    }
}
