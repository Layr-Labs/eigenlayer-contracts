// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {DeployProtocolRegistryProxy} from "./1-deployProtocolRegistryProxy.s.sol";
import {DeployProtocolRegistryImpl} from "./2-deployProtocolRegistryImpl.s.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";

import "../Env.sol";
import "../TestUtils.sol";

/// Purpose: Upgrade Protocol Registry Proxy to point to the implementation. Also transfer control to the ProxyAdmin.
contract UpgradeProtocolRegistry is DeployProtocolRegistryImpl {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsMultisig() internal virtual override prank(Env.multichainDeployerMultisig()) {
        // Upgrade the proxies to point to the actual implementations
        ITransparentUpgradeableProxy protocolRegistryProxy =
            ITransparentUpgradeableProxy(payable(address(Env.proxy.protocolRegistry())));
        protocolRegistryProxy.upgradeToAndCall(address(Env.impl.protocolRegistry()), abi.encodeWithSelector(ProtocolRegistry.initialize.selector, Env.executorMultisig(), Env.pauserMultisig()));

        // Transfer proxy admin ownership
        protocolRegistryProxy.changeAdmin(Env.proxyAdmin());
    }

    function testScript() public virtual override {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        // 1. Deploy the Protocol Registry Proxy
        // If proxies are not deployed, deploy them
        if (!_areProxiesDeployed()) {
            DeployProtocolRegistryProxy._runAsMultisig();
            _unsafeResetHasPranked(); // reset hasPranked so we can use it in the execute()
        } else {
            // Since the proxies are already deployed, we need to update the env with the proper addresses
            _addContractsToEnv();
        }

        // If the proxy has been upgraded already, return
        if (_areProxiesDeployed() && Env.getProxyAdminBySlot(address(Env.proxy.protocolRegistry())) == Env.proxyAdmin())
        {
            return;
        }

        // 2. Deploy the Protocol Registry Implementation
        _mode = OperationalMode.EOA; // Set to EOA mode so we can deploy the impls in the EOA script
        DeployProtocolRegistryImpl._runAsEOA();

        // 3. Upgrade the Protocol Registry Proxy
        execute();

        // 4. Validate the Protocol Registry
        TestUtils.validateDestinationProxyAdmins();
        TestUtils.validateProtocolRegistryInitialized(Env.proxy.protocolRegistry());
        // ProtocolRegistry has no constructor, so we don't need to validate the constructors
        // ProtocolRegistry has no initial storage, so we don't need to validate the storage
        TestUtils.validateDestinationImplAddressesMatchProxy();
    }
}
