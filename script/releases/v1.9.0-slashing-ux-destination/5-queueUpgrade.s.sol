// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {UpgradeProtocolRegistry} from "./3-upgradeProtocolRegistry.s.sol";
import {DeployCoreContracts} from "./4-deployCoreContracts.s.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {CoreUpgradeQueueBuilder} from "../CoreUpgradeQueueBuilder.sol";
import {IProtocolRegistryTypes} from "src/contracts/interfaces/IProtocolRegistry.sol";
import "../Env.sol";
import "../TestUtils.sol";

contract QueueUpgrade is DeployCoreContracts {
    using Env for *;
    using Encode for *;
    using CoreUpgradeQueueBuilder for MultisigCall[];

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

        // Add the protocol registry as a pauser to the pauser registry
        Env.impl.pauserRegistry().setIsPauser(address(Env.proxy.protocolRegistry()), true);
    }

    function _getCalldataToExecutor() internal returns (bytes memory) {
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls();

        /// multichain
        executorCalls.upgradeBN254CertificateVerifier();
        executorCalls.upgradeECDSACertificateVerifier();
        executorCalls.upgradeOperatorTableUpdater();

        /// avs
        executorCalls.upgradeTaskMailbox();

        // Add the protocol registry upgrade to the executor calls
        _appendProtocolRegistryUpgrade(executorCalls);

        return Encode.gnosisSafe
            .execTransaction({
                from: address(Env.timelockController()),
                to: Env.multiSendCallOnly(),
                op: Encode.Operation.DelegateCall,
                data: Encode.multiSend(executorCalls)
            });
    }

    function _appendProtocolRegistryUpgrade(
        MultisigCall[] storage calls
    ) internal {
        // We want to add all addresses that are deployed to the protocol registry
        address[] memory addresses = new address[](6);
        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](6);
        string[] memory names = new string[](6);

        IProtocolRegistryTypes.DeploymentConfig memory pausableConfig =
            IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});
        IProtocolRegistryTypes.DeploymentConfig memory unpausableConfig =
            IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});

        /// permissions/
        addresses[0] = address(Env.impl.pauserRegistry());
        configs[0] = unpausableConfig;
        names[0] = type(PauserRegistry).name;

        /// core/
        addresses[1] = address(Env.proxy.protocolRegistry());
        configs[1] = unpausableConfig;
        names[1] = type(ProtocolRegistry).name;

        /// multichain/
        addresses[2] = address(Env.proxy.bn254CertificateVerifier());
        configs[2] = unpausableConfig;
        names[2] = type(BN254CertificateVerifier).name;

        addresses[3] = address(Env.proxy.ecdsaCertificateVerifier());
        configs[3] = unpausableConfig;
        names[3] = type(ECDSACertificateVerifier).name;

        addresses[4] = address(Env.proxy.operatorTableUpdater());
        configs[4] = pausableConfig;
        names[4] = type(OperatorTableUpdater).name;

        /// avs/
        addresses[5] = address(Env.proxy.taskMailbox());
        configs[5] = unpausableConfig;
        names[5] = type(TaskMailbox).name;

        // Lastly, append to the multisig calls
        calls.append({
            to: address(Env.proxy.protocolRegistry()),
            data: abi.encodeWithSelector(
                IProtocolRegistry.ship.selector, addresses, configs, names, Env.deployVersion()
            )
        });
    }

    function testScript() public virtual override {
        if (Env.isCoreProtocolDeployed()) {
            return;
        }

        // Complete previous steps
        _completeProtocolRegistryUpgrade();

        // Deploy the core contracts
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
