// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import "../TestUtils.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";

contract DeployImplementations is CoreContractsDeployer {
    using Env for *;

    function _runAsEOA() internal virtual override {
        vm.startBroadcast();

        deployStrategyManager();

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        runAsEOA();

        TestUtils.validateStrategyManagerImmutables(Env.impl.strategyManager());
        TestUtils.validateStrategyManagerInitialized(Env.impl.strategyManager());
        require(
            keccak256(bytes(Env.impl.strategyManager().version())) == keccak256(bytes("1.13.0")),
            "strategyManager version must be 1.13.0"
        );
        TestUtils.validateStrategyManagerSlashResolutionDelay(Env.impl.strategyManager());
    }
}
