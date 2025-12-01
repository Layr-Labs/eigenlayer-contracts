// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";
import {DeployProtocolRegistryProxy} from "./1-deployProtocolRegistryProxy.s.sol";
import {DeployProtocolRegistryImpl} from "./2-deployProtocolRegistryImpl.s.sol";
import {UpgradeProtocolRegistry} from "./3-upgradeProtocolRegistry.s.sol";
import "../Env.sol";
import "../TestUtils.sol";

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * Purpose: use an EOA to deploy all of the new contracts for this upgrade.
 * Contracts deployed:
 * /// Multichain
 * - BN254CertificateVerifier
 * - CrossChainRegistry
 * - ECDSACertificateVerifier
 * - OperatorTableUpdater
 * /// AVS
 * - TaskMailbox
 */
contract DeployCoreContracts is UpgradeProtocolRegistry {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        /**
         * multichain/
         */
        deployBN254CertificateVerifier();
        // CrossChainRegistry only deployed on destination chain
        deployECDSACertificateVerifier();
        deployOperatorTableUpdater();

        /**
         * avs/
         */
        deployTaskMailbox();

        vm.stopBroadcast();
    }

    function testScript() public virtual override {
        if (Env.isCoreProtocolDeployed()) {
            return;
        }

        // Deploy protocol registry and initialize it
        _completeProtocolRegistryUpgrade();

        // Deploy the core contracts
        runAsEOA();

        // Run tests
        TestUtils.validateDestinationProxyAdmins();
        TestUtils.validateDestinationImplConstructors();
        TestUtils.validateDestinationImplsNotInitializable();

        // Check individual version addresses
        TestUtils.validateECDSACertificateVerifierVersion();
    }

    function _completeProtocolRegistryUpgrade() internal {
        // If proxies are not deployed, deploy them
        if (!_areProxiesDeployed()) {
            DeployProtocolRegistryProxy._runAsMultisig();
            _unsafeResetHasPranked(); // reset hasPranked so we can use it in the execute()
        } else {
            // Since the proxies are already deployed, we need to update the env with the proper addresses
            _addContractsToEnv();
        }

        // 2. Deploy the Protocol Registry Implementation
        _mode = OperationalMode.EOA; // Set to EOA mode so we can deploy the impls in the EOA script
        DeployProtocolRegistryImpl._runAsEOA();

        // 3. Upgrade the Protocol Registry Proxy, only if the admin is not already the proxyAdmin
        if (Env.getProxyAdminBySlot(address(Env.proxy.protocolRegistry())) != Env.proxyAdmin()) {
            UpgradeProtocolRegistry._runAsMultisig();
            _unsafeResetHasPranked(); // reset hasPranked so we can use in the future
        }
    }
}
