// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import "../TestUtils.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";

/// Purpose: deploy the StrategyManager implementation with a configurable slash resolution delay.
contract DeployImplementations is CoreContractsDeployer {
    using Env for *;

    function _runAsEOA() internal virtual override {
        vm.startBroadcast();

        // v1.13.2: StrategyManager with a configurable slash resolution delay.
        deployStrategyManager();

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        if (!Env.isCoreProtocolDeployed() || Env._versionGte(Env.envVersion(), Env.deployVersion())) {
            return;
        }

        runAsEOA();

        StrategyManager strategyManager = Env.impl.strategyManager();
        assertTrue(address(strategyManager).code.length != 0, "strategyManager implementation not deployed");
        TestUtils.validateStrategyManagerImmutables(strategyManager);
        TestUtils.validateStrategyManagerInitialized(strategyManager);
        TestUtils.validateStrategyManagerVersion();

        // Fresh implementation storage: no delay and no setter (both are configured through the proxy).
        assertEq(strategyManager.SLASH_RESOLUTION_DELAY_BLOCKS(), 0, "implementation delay should be zero");
        assertEq(strategyManager.slashResolutionDelaySetter(), address(0), "implementation setter should be unset");
    }
}
