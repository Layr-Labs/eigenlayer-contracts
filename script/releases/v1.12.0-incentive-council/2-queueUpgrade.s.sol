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

import {EmissionsController} from "src/contracts/core/EmissionsController.sol";
import {RewardsCoordinator} from "src/contracts/core/RewardsCoordinator.sol";
import {IProtocolRegistry, IProtocolRegistryTypes} from "src/contracts/interfaces/IProtocolRegistry.sol";
import {IBackingEigen} from "src/contracts/interfaces/IBackingEigen.sol";

/// Purpose: Queue the upgrade for Duration Vault and Incentive Council features.
/// This script queues upgrades to:
/// - EmissionsController (new proxy + initialize, or upgrade for preprod)
/// - RewardsCoordinator (rewards v2.2 + protocol fees, with conditional reinitialization)
/// - StrategyManager, EigenStrategy, StrategyBase beacon, StrategyBaseTVLLimits, StrategyFactory
/// - Register DurationVaultStrategy beacon + EmissionsController in ProtocolRegistry
/// - Transfer minting rights from legacy hopper to EmissionsController
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

        // 1. Upgrade EmissionsController proxy to the new implementation and initialize
        // Preprod needs to be upgraded, not redeployed.
        if (Env._strEq(Env.envVersion(), "1.12.0")) {
            executorCalls.append({
                to: Env.proxyAdmin(),
                data: abi.encodeCall(
                    IProxyAdmin.upgrade,
                    (
                        ITransparentProxy(payable(address(Env.proxy.emissionsController()))),
                        address(Env.impl.emissionsController())
                    )
                )
            });
        } else {
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
        }
        // 2. Upgrade strategy contracts for Duration Vault feature.
        executorCalls.upgradeStrategyManager();
        executorCalls.upgradeEigenStrategy();
        executorCalls.upgradeStrategyBase();
        executorCalls.upgradeStrategyBaseTVLLimits();
        executorCalls.upgradeStrategyFactory();

        // 3. Upgrade RewardsCoordinator to the new implementation.
        // Check if RewardsCoordinator has already been reinitialized by reading _initialized from storage.
        // Slot 0 contains: _initialized (uint8 at offset 0) and _initializing (bool at offset 1)
        // If _initialized >= 2, reinitializer(2) has already been called.
        RewardsCoordinator rc = Env.proxy.rewardsCoordinator();
        bytes32 slot0 = vm.load(address(rc), bytes32(uint256(0)));
        uint8 initializedVersion = uint8(uint256(slot0) & 0xFF); // Mask to get only the first byte

        if (initializedVersion < 2) {
            // Not yet reinitialized - perform upgrade with reinitialization
            executorCalls.upgradeAndReinitializeRewardsCoordinator({
                initialOwner: Env.opsMultisig(),
                initialPausedStatus: 2,
                rewardsUpdater: Env.REWARDS_UPDATER(),
                activationDelay: Env.ACTIVATION_DELAY(),
                defaultSplitBips: Env.DEFAULT_SPLIT_BIPS(),
                feeRecipient: Env.incentiveCouncilMultisig()
            });
        } else {
            // Already reinitialized - just upgrade without calling initialize
            executorCalls.upgradeRewardsCoordinator();
        }
        // 4. Register DurationVaultStrategy beacon and EmissionsController in protocol registry.
        address[] memory addresses = new address[](2);
        addresses[0] = address(Env.beacon.durationVaultStrategy());
        addresses[1] = address(Env.proxy.emissionsController());

        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](2);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});
        configs[1] = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});

        string[] memory names = new string[](2);
        names[0] = type(DurationVaultStrategy).name;
        names[1] = type(EmissionsController).name;

        executorCalls.append({
            to: address(Env.proxy.protocolRegistry()),
            data: abi.encodeCall(IProtocolRegistry.ship, (addresses, configs, names, Env.deployVersion()))
        });
        // 5. Remove minting rights from the old hopper.
        executorCalls.append({
            to: address(Env.proxy.beigen()),
            data: abi.encodeCall(IBackingEigen.setIsMinter, (Env.legacyTokenHopper(), false))
        });
        // 6. Grant minting rights to the EmissionsController.
        executorCalls.append({
            to: address(Env.proxy.beigen()),
            data: abi.encodeCall(IBackingEigen.setIsMinter, (address(Env.proxy.emissionsController()), true))
        });

        // Execute the calls.
        return Encode.gnosisSafe
            .execTransaction({
                from: address(Env.timelockController()),
                to: Env.multiSendCallOnly(),
                op: Encode.Operation.DelegateCall,
                data: Encode.multiSend(executorCalls)
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

        // Check that the upgrade does not exist in the timelock
        assertFalse(timelock.isOperationPending(txHash), "Transaction should NOT be queued.");
        _validatePreUpgradeState();

        execute();

        // Check that the upgrade has been added to the timelock
        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
    }

    /// @notice Validates state before queuing the upgrade
    function _validatePreUpgradeState() internal view {
        // Ensure new implementations are deployed
        require(
            address(Env.impl.emissionsController()).code.length > 0, "EmissionsController implementation not deployed"
        );
        require(
            address(Env.impl.rewardsCoordinator()).code.length > 0, "RewardsCoordinator implementation not deployed"
        );

        // Ensure EmissionsController proxy is deployed
        require(address(Env.proxy.emissionsController()).code.length > 0, "EmissionsController proxy not deployed");
        // Ensure old hopper currently has minting rights
        IBackingEigen beigen = Env.proxy.beigen();
        if (Env.legacyTokenHopper() != address(0)) {
            require(beigen.isMinter(Env.legacyTokenHopper()), "Old hopper should have minting rights before upgrade");
        }
        // Ensure EmissionsController does NOT have minting rights yet
        // Skip check if we are on preprod, as the EmissionsController is already deployed.
        if (!Env._strEq(Env.envVersion(), "1.12.0")) {
            require(
                !beigen.isMinter(address(Env.proxy.emissionsController())),
                "EmissionsController should NOT have minting rights before upgrade"
            );
        }
    }
}

