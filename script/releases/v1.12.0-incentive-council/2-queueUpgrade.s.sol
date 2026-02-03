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

/// Steps:
/// 1. Queue EmissionsController upgrade and initialization.
/// 2. Queue RewardsCoordinator upgrade (with conditional reinitialization based on current state).
/// 3. Queue add EmissionsController to protocol registry.
/// 4. Queue remove minting rights from old hopper.
/// 5. Queue grant minting rights to EmissionsController.
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
        // 2. Upgrade RewardsCoordinator to the new implementation.
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
                initialPausedStatus: 0,
                rewardsUpdater: Env.REWARDS_UPDATER(),
                activationDelay: Env.ACTIVATION_DELAY(),
                defaultSplitBips: Env.DEFAULT_SPLIT_BIPS(),
                feeRecipient: Env.incentiveCouncilMultisig()
            });
        } else {
            // Already reinitialized - just upgrade without calling initialize
            executorCalls.upgradeRewardsCoordinator();
        }
        // 3. Add EmissionsController to the protocol registry.
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
        // 4. Remove minting rights from the old hopper.
        executorCalls.append({
            to: address(Env.proxy.beigen()),
            data: abi.encodeCall(IBackingEigen.setIsMinter, (Env.legacyTokenHopper(), false))
        });
        // 5. Grant minting rights to the EmissionsController.
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
        require(
            !beigen.isMinter(address(Env.proxy.emissionsController())),
            "EmissionsController should NOT have minting rights before upgrade"
        );
    }
}

