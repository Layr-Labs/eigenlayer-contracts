// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {DeployProtocolRegistryProxy} from "./1-deployProtocolRegistryProxy.s.sol";
import "../Env.sol";
import "../TestUtils.sol";

/**
 * Purpose: Deploy Protocol Registry implementation
 */
contract DeployProtocolRegistryImpl is EOADeployer, DeployProtocolRegistryProxy {
    using Env for *;

    function _runAsEOA() internal virtual override {
        vm.startBroadcast();

        // Deploy Protocol Registry implementation
        deployImpl({name: type(ProtocolRegistry).name, deployedTo: address(new ProtocolRegistry())});

        vm.stopBroadcast();
    }

    function testScript() public virtual override {
        // 1. Deploy Protocol Registry Proxy
        // Only deploy the proxies if they haven't been deployed yet
        /// @dev This is needed in the production environment tests since this step would fail if the proxies are already deployed
        if (!_areProxiesDeployed()) {
            DeployProtocolRegistryProxy._runAsMultisig();
            _unsafeResetHasPranked(); // reset hasPranked so we can use it in the execute()
        } else {
            // Since the proxies are already deployed, we need to update the env with the proper addresses
            _addContractsToEnv();
        }

        // 2. Deploy Implementation
        runAsEOA();

        // Validate the implementation
        // Protocol Registry has no constructor, so we don't need to validate the immutables
        TestUtils.validateProtocolRegistryInitialized(Env.impl.protocolRegistry());
    }
}
