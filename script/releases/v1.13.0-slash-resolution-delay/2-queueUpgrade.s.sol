// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import {IProxyAdmin} from "zeus-templates/interfaces/IProxyAdmin.sol";
import {
    ITransparentUpgradeableProxy as ITransparentProxy
} from "zeus-templates/interfaces/ITransparentUpgradeableProxy.sol";
import {DeployImplementations} from "./1-deployImplementations.s.sol";
import {CoreUpgradeQueueBuilder} from "../CoreUpgradeQueueBuilder.sol";
import "../Env.sol";
import "../TestUtils.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import {IProtocolRegistry, IProtocolRegistryTypes} from "src/contracts/interfaces/IProtocolRegistry.sol";
import {IBackingEigen} from "src/contracts/interfaces/IBackingEigen.sol";
import {EmissionsController} from "src/contracts/core/EmissionsController.sol";

/// Purpose: Queue the upgrade for the slash resolution delay + duration vault blacklist fix.
/// This script queues an upgrade to:
/// - StrategyManager (adds SLASH_RESOLUTION_DELAY_BLOCKS constant and burn/redistribution pause flag)
/// - StrategyFactory (adds EIGEN/bEIGEN immutables, replaces blacklist check for duration vaults)
/// - DurationVaultStrategy beacon (removes deposit-time blacklist check)
/// - RewardsCoordinator upgrade + reinitialize (if env missed v1.12.0)
/// - EmissionsController deploy + initialize (if env missed v1.12.0)
/// - ProtocolRegistry semantic version bump to 1.13.0
contract QueueUpgrade is DeployImplementations, MultisigBuilder {
    using Env for *;
    using Encode for *;
    using CoreUpgradeQueueBuilder for *;

    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        bytes memory calldata_to_executor = _getCalldataToExecutor();

        TimelockController timelock = Env.timelockController();
        timelock.schedule({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor,
            predecessor: 0,
            salt: 0,
            delay: timelock.getMinDelay()
        });
    }

    function _getCalldataToExecutor() internal returns (bytes memory) {
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls();

        bool needsIncentiveCouncil = _needsIncentiveCouncilUpgrade();

        // 1. Upgrade StrategyManager proxy to new implementation
        executorCalls.upgradeStrategyManager();

        // 2. Upgrade StrategyFactory proxy to new implementation (adds EIGEN/bEIGEN immutables)
        executorCalls.upgradeStrategyFactory();

        // 3. Upgrade DurationVaultStrategy beacon to new implementation (removes blacklist check)
        executorCalls.upgradeDurationVaultStrategy();

        // 4. Conditionally catch up with v1.12.0 incentive council changes
        if (needsIncentiveCouncil) {
            _queueIncentiveCouncilCatchUp(executorCalls);
        }

        // 5. Bump protocol registry semantic version to 1.13.0
        //    If catching up, also register EmissionsController in the protocol registry.
        if (needsIncentiveCouncil) {
            address[] memory addresses = new address[](1);
            addresses[0] = address(Env.proxy.emissionsController());

            IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](1);
            configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});

            string[] memory names = new string[](1);
            names[0] = type(EmissionsController).name;

            executorCalls.append({
                to: address(Env.proxy.protocolRegistry()),
                data: abi.encodeCall(IProtocolRegistry.ship, (addresses, configs, names, Env.deployVersion()))
            });
        } else {
            executorCalls.append({
                to: address(Env.proxy.protocolRegistry()),
                data: abi.encodeCall(
                    IProtocolRegistry.ship,
                    (
                        new address[](0),
                        new IProtocolRegistryTypes.DeploymentConfig[](0),
                        new string[](0),
                        Env.deployVersion()
                    )
                )
            });
        }

        return Encode.gnosisSafe
            .execTransaction({
                from: address(Env.timelockController()),
                to: Env.multiSendCallOnly(),
                op: Encode.Operation.DelegateCall,
                data: Encode.multiSend(executorCalls)
            });
    }

    /// @notice Queues v1.12.0 incentive council changes for envs that missed that release.
    /// Upgrades RewardsCoordinator (with reinitialize), initializes EmissionsController,
    /// and transfers bEIGEN minting rights.
    function _queueIncentiveCouncilCatchUp(
        MultisigCall[] storage executorCalls
    ) internal {
        // Initialize EmissionsController proxy
        executorCalls.append({
            to: Env.proxyAdmin(),
            data: abi.encodeCall(
                IProxyAdmin.upgradeAndCall,
                (
                    ITransparentProxy(payable(address(Env.proxy.emissionsController()))),
                    address(Env.impl.emissionsController()),
                    abi.encodeCall(
                        EmissionsController.initialize,
                        (
                            Env.opsMultisig(), // initialOwner
                            Env.incentiveCouncilMultisig(), // initialIncentiveCouncil
                            0 // initialPausedStatus
                        )
                    )
                )
            )
        });

        // Upgrade + reinitialize RewardsCoordinator
        executorCalls.upgradeAndReinitializeRewardsCoordinator({
            initialOwner: Env.opsMultisig(),
            initialPausedStatus: 2,
            rewardsUpdater: Env.REWARDS_UPDATER(),
            activationDelay: Env.ACTIVATION_DELAY(),
            defaultSplitBips: Env.DEFAULT_SPLIT_BIPS(),
            feeRecipient: Env.incentiveCouncilMultisig()
        });

        // Transfer bEIGEN minting rights to EmissionsController
        if (Env.legacyTokenHopper() != address(0)) {
            executorCalls.append({
                to: address(Env.proxy.beigen()),
                data: abi.encodeCall(IBackingEigen.setIsMinter, (Env.legacyTokenHopper(), false))
            });
        }
        executorCalls.append({
            to: address(Env.proxy.beigen()),
            data: abi.encodeCall(IBackingEigen.setIsMinter, (address(Env.proxy.emissionsController()), true))
        });
    }

    function testScript() public virtual override {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        runAsEOA();

        TimelockController timelock = Env.timelockController();
        bytes memory calldata_to_executor = _getCalldataToExecutor();
        bytes32 txHash = timelock.hashOperation({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });

        assertFalse(timelock.isOperationPending(txHash), "Transaction should NOT be queued.");

        // Validate new implementations
        TestUtils.validateStrategyManagerImmutables(Env.impl.strategyManager());
        TestUtils.validateStrategyManagerVersion();
        TestUtils.validateStrategyManagerSlashResolutionDelay(Env.impl.strategyManager());
        TestUtils.validateStrategyFactoryImmutables(Env.impl.strategyFactory());
        TestUtils.validateDurationVaultStrategyImmutables(Env.impl.durationVaultStrategy());

        execute();

        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
    }
}
