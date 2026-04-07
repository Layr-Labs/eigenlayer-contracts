// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import "../TestUtils.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";

/// Purpose: deploy the patched DurationVaultStrategy implementation for Sepolia.
contract DeployImplementations is CoreContractsDeployer {
    using Env for *;

    function _runAsEOA() internal virtual override {
        _requireSepoliaPatchEnv();

        vm.startBroadcast();
        deployDurationVaultStrategy();
        vm.stopBroadcast();
    }

    function testScript() public virtual {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        if (!_isSepoliaPatchEnv()) {
            return;
        }

        runAsEOA();

        TestUtils.validateDurationVaultStrategyImplConstructors();
    }

    function _isSepoliaPatchEnv() internal view returns (bool) {
        return Env._strEq(Env.env(), "testnet-sepolia") && Env._strEq(Env.envVersion(), "1.12.0")
            && Env._strEq(Env.deployVersion(), "1.12.1");
    }

    function _requireSepoliaPatchEnv() internal view {
        require(Env._strEq(Env.env(), "testnet-sepolia"), "only testnet-sepolia");
        require(Env._strEq(Env.envVersion(), "1.12.0"), "expected env version 1.12.0");
        require(Env._strEq(Env.deployVersion(), "1.12.1"), "expected deploy version 1.12.1");
    }
}
