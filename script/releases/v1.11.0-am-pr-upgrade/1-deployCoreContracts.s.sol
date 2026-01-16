// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";
import "../Env.sol";
import "../TestUtils.sol";

/// Purpose: use an EOA to deploy new implementations for this upgrade.
/// Contracts deployed:
/// - AllocationManagerView (dependency of AllocationManager)
/// - AllocationManager
/// - ProtocolRegistry
contract DeployCoreContracts is CoreContractsDeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        /// core/
        deployAllocationManagerView();
        deployAllocationManager();
        deployProtocolRegistry();

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        // Deploy the core contracts
        runAsEOA();

        // Run tests
        TestUtils.validateProxyAdmins();
        TestUtils.validateImplConstructors();
        // TestUtils.validateImplsNotInitializable();
    }
}
