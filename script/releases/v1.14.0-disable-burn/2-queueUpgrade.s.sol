// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import {DeployImplementations} from "./1-deployImplementations.s.sol";
import {CoreUpgradeQueueBuilder} from "../CoreUpgradeQueueBuilder.sol";
import "../Env.sol";
import "../TestUtils.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import {IProtocolRegistry, IProtocolRegistryTypes} from "src/contracts/interfaces/IProtocolRegistry.sol";
import {IBackingEigen} from "src/contracts/interfaces/IBackingEigen.sol";
import {EmissionsController} from "src/contracts/core/EmissionsController.sol";

/// Purpose: Queue the v1.14.0 core upgrade.
/// This script queues upgrades to:
/// - DelegationManager (disabled-pod queued withdrawal cleanup)
/// - EigenPodManager and EigenPod (permanent pod disable)
/// - EmissionsController (burn distributions)
/// It also grants EmissionsController permission to burn bEIGEN and ships v1.14.0.
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

        executorCalls.upgradeDelegationManager();
        executorCalls.upgradeEigenPodManager();
        executorCalls.upgradeEigenPod();

        if (needsIncentiveCouncil) {
            _queueIncentiveCouncilCatchUp(executorCalls);
        } else {
            executorCalls.upgradeEmissionsController();
        }

        executorCalls.append({
            to: address(Env.proxy.beigen()),
            data: abi.encodeCall(IBackingEigen.setAllowedFrom, (address(Env.proxy.emissionsController()), true))
        });

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

    function _queueIncentiveCouncilCatchUp(
        MultisigCall[] storage executorCalls
    ) internal {
        executorCalls.upgradeAndInitializeEmissionsController({
            initialOwner: Env.opsMultisig(),
            initialIncentiveCouncil: Env.incentiveCouncilMultisig(),
            initialPausedStatus: 0
        });

        executorCalls.upgradeAndReinitializeRewardsCoordinator({
            initialOwner: Env.opsMultisig(),
            initialPausedStatus: 2,
            rewardsUpdater: Env.REWARDS_UPDATER(),
            activationDelay: Env.ACTIVATION_DELAY(),
            defaultSplitBips: Env.DEFAULT_SPLIT_BIPS(),
            feeRecipient: Env.incentiveCouncilMultisig()
        });

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

    function testScript() public virtual override onlyIfUpgradeRequired("1.14.0") {
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

        TestUtils.validateDelegationManagerImmutables(Env.impl.delegationManager());
        TestUtils.validateDelegationManagerInitialized(Env.impl.delegationManager());
        TestUtils.validateDelegationManagerVersion();
        TestUtils.validateEigenPodManagerImmutables(Env.impl.eigenPodManager());
        TestUtils.validateEigenPodManagerInitialized(Env.impl.eigenPodManager());
        TestUtils.validateEigenPodImmutables(Env.impl.eigenPod());
        TestUtils.validateEmissionsControllerImmutables(Env.impl.emissionsController());
        TestUtils.validateEmissionsControllerInitialized(Env.impl.emissionsController());

        execute();

        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
    }
}
