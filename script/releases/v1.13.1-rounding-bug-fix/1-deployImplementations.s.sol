// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import "../TestUtils.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";

/// Purpose: use an EOA to deploy DelegationManager for the queued slash accounting fix.
contract DeployImplementations is CoreContractsDeployer {
    using Env for *;

    function _runAsEOA() internal virtual override {
        vm.startBroadcast();

        // v1.13.1 changes
        deployDelegationManager();

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        runAsEOA();

        TestUtils.validateDelegationManagerImmutables(Env.impl.delegationManager());
        TestUtils.validateDelegationManagerInitialized(Env.impl.delegationManager());
        TestUtils.validateDelegationManagerVersion();
    }
}
