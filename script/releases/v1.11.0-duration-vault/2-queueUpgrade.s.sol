// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {DeployContracts} from "./1-deployContracts.s.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import {CoreUpgradeQueueBuilder} from "../CoreUpgradeQueueBuilder.sol";
import {IProtocolRegistry, IProtocolRegistryTypes} from "src/contracts/interfaces/IProtocolRegistry.sol";
import "../Env.sol";
import "../TestUtils.sol";

/// Purpose: Queue the upgrade for Duration Vault feature.
/// This script queues upgrades to:
/// - StrategyManager proxy
/// - EigenStrategy proxy
/// - StrategyBase beacon
/// - StrategyBaseTVLLimits proxies
/// - StrategyFactory proxy
/// - Sets DurationVaultStrategy beacon on StrategyFactory
/// - Registers DurationVaultStrategy beacon in ProtocolRegistry
contract QueueUpgrade is DeployContracts, MultisigBuilder {
    using Env for *;
    using Encode for *;
    using CoreUpgradeQueueBuilder for MultisigCall[];

    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        // Only execute on version 1.10.0
        if (!Env._strEq(Env.envVersion(), "1.10.0")) {
            return;
        }

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

        /// core/
        executorCalls.upgradeStrategyManager();

        /// strategies/
        executorCalls.upgradeEigenStrategy();
        executorCalls.upgradeStrategyBase();
        executorCalls.upgradeStrategyBaseTVLLimits();

        // Upgrade StrategyFactory (beacons are now immutable in the implementation)
        executorCalls.upgradeStrategyFactory();

        // Register the DurationVaultStrategy beacon in the protocol registry and tick version
        _appendProtocolRegistryUpgrade(executorCalls);

        return Encode.gnosisSafe
            .execTransaction({
                from: address(Env.timelockController()),
                to: Env.multiSendCallOnly(),
                op: Encode.Operation.DelegateCall,
                data: Encode.multiSend(executorCalls)
            });
    }

    /// @notice Appends the protocol registry upgrade to register the DurationVaultStrategy beacon
    function _appendProtocolRegistryUpgrade(
        MultisigCall[] storage calls
    ) internal {
        address[] memory addresses = new address[](1);
        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](1);
        string[] memory names = new string[](1);

        IProtocolRegistryTypes.DeploymentConfig memory unpausableConfig =
            IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});

        // Register DurationVaultStrategy beacon
        addresses[0] = address(Env.beacon.durationVaultStrategy());
        configs[0] = unpausableConfig;
        names[0] = type(DurationVaultStrategy).name;

        calls.append({
            to: address(Env.proxy.protocolRegistry()),
            data: abi.encodeWithSelector(
                IProtocolRegistry.ship.selector, addresses, configs, names, Env.deployVersion()
            )
        });
    }

    function testScript() public virtual override {
        if (!Env.isCoreProtocolDeployed() || !Env.isSource() || !Env._strEq(Env.envVersion(), "1.10.0")) {
            return;
        }

        // Deploy the contracts first
        super.runAsEOA();

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

        execute();

        // Check that the upgrade has been added to the timelock
        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
    }
}
