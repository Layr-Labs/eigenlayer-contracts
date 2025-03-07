// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {DeployFresh} from "./1-deploy-proxies-and-implementations.s.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "src/contracts/libraries/BeaconChainProofs.sol";
import "forge-std/console.sol";

/**
 * Purpose: update the implementation contracts on the previously deployed blank proxies. 
 */
contract UpgradeContracts is MultisigBuilder, DeployFresh {
    using Env for *;

    function _runAsMultisig() prank(Env.protocolCouncilMultisig()) internal virtual override {
        ProxyAdmin pa = ProxyAdmin(Env.proxyAdmin());

        pa.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(Env.proxy.delegationManager()))),
            address(Env.impl.delegationManager()),
            abi.encodeCall(
                DelegationManager.initialize,
                (
                    Env.executorMultisig(),
                    Env.DELEGATION_INIT_PAUSED_STATUS()
                )
            )
        );
        
        pa.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(Env.proxy.strategyManager()))),
            address(Env.impl.strategyManager()),
            abi.encodeCall(
                StrategyManager.initialize,
                (
                    Env.executorMultisig(),
                    Env.opsMultisig(),
                    Env.STRATEGY_MANAGER_INIT_PAUSED_STATUS()
                )
            )
        );

        pa.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(Env.proxy.avsDirectory()))),
            address(Env.impl.avsDirectory()),
            abi.encodeCall(
                AVSDirectory.initialize,
                ( 
                    Env.executorMultisig(),
                    0
                )
            )
        );

        pa.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(Env.proxy.eigenPodManager()))),
            address(Env.impl.eigenPodManager()),
            abi.encodeCall(
                EigenPodManager.initialize,
                (
                    Env.executorMultisig(),
                    Env.EIGENPOD_MANAGER_INIT_PAUSED_STATUS()
                )
            )
        );

        pa.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(Env.proxy.rewardsCoordinator()))),
            address(Env.impl.rewardsCoordinator()),
            abi.encodeCall(
                RewardsCoordinator.initialize,
                (
                    Env.executorMultisig(),
                    Env.REWARDS_PAUSE_STATUS(),
                    Env.REWARDS_UPDATER(),
                    Env.ACTIVATION_DELAY(),
                    Env.DEFAULT_SPLIT_BIPS()
                )
            )
        );

        pa.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(Env.proxy.allocationManager()))),
            address(Env.impl.allocationManager()),
            abi.encodeCall(
                AllocationManager.initialize,
                (
                    Env.executorMultisig(),
                    Env.ALLOCATION_MANAGER_INIT_PAUSED_STATUS()
                )
            )
        );

        pa.upgrade(
            ITransparentUpgradeableProxy(payable(address(Env.proxy.permissionController()))),
            address(Env.impl.permissionController())
        );
    }


    function testDeploy() public override virtual {} // prevent duplicate test.

    function testScript() public virtual {
        _runAsEOA();
        execute();
    }
}
