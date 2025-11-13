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
    }

    function _getCalldataToExecutor() internal returns (bytes memory) {
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls();

        /// permissions
        executorCalls.upgradePermissionController();
        executorCalls.upgradeKeyRegistrar();

        /// core
        executorCalls.upgradeAllocationManager();
        executorCalls.upgradeAVSDirectory();
        executorCalls.upgradeDelegationManager();
        // protocol registry was deployed in the previous step
        executorCalls.upgradeReleaseManager();
        executorCalls.upgradeRewardsCoordinator();
        executorCalls.upgradeStrategyManager();

        /// pods
        executorCalls.upgradeEigenPodManager();
        executorCalls.upgradeEigenPod();

        /// strategies
        executorCalls.upgradeEigenStrategy();
        executorCalls.upgradeStrategyBase();
        executorCalls.upgradeStrategyBaseTVLLimits();
        executorCalls.upgradeStrategyFactory();

        /// multichain
        executorCalls.upgradeBN254CertificateVerifier();
        executorCalls.upgradeCrossChainRegistry();
        executorCalls.upgradeECDSACertificateVerifier();
        executorCalls.upgradeOperatorTableUpdater();

        /// avs
        executorCalls.upgradeTaskMailbox();

        // Add the protocol registry upgrade to the executor calls
        _appendProtocolRegistryUpgrade(executorCalls);

        // Lastly, add the protocol registry as a pauser to the pauser registry
        executorCalls.append({
            to: address(Env.impl.pauserRegistry()),
            data: abi.encodeWithSelector(PauserRegistry.setIsPauser.selector, address(Env.proxy.protocolRegistry()), true)
        });

        return Encode.gnosisSafe.execTransaction({
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
        address[] memory addresses = new address[](22);
        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](22);
        string[] memory names = new string[](22);

        IProtocolRegistryTypes.DeploymentConfig memory pausableConfig =
            IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});
        IProtocolRegistryTypes.DeploymentConfig memory unpausableConfig =
            IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});

        /**
         * permissions/
         */
        addresses[0] = address(Env.impl.pauserRegistry());
        configs[0] = unpausableConfig;
        names[0] = type(PauserRegistry).name;

        addresses[1] = address(Env.proxy.permissionController());
        configs[1] = unpausableConfig;
        names[1] = type(PermissionController).name;

        addresses[2] = address(Env.proxy.keyRegistrar());
        configs[2] = unpausableConfig;
        names[2] = type(KeyRegistrar).name;

        /**
         * core/
         */
        addresses[3] = address(Env.proxy.allocationManager());
        configs[3] = pausableConfig;
        names[3] = type(AllocationManager).name;

        addresses[4] = address(Env.proxy.avsDirectory());
        configs[4] = pausableConfig;
        names[4] = type(AVSDirectory).name;

        addresses[5] = address(Env.proxy.delegationManager());
        configs[5] = pausableConfig;
        names[5] = type(DelegationManager).name;

        addresses[6] = address(Env.proxy.protocolRegistry());
        configs[6] = unpausableConfig;
        names[6] = type(ProtocolRegistry).name;

        addresses[7] = address(Env.proxy.releaseManager());
        configs[7] = unpausableConfig;
        names[7] = type(ReleaseManager).name;

        addresses[8] = address(Env.proxy.rewardsCoordinator());
        configs[8] = pausableConfig;
        names[8] = type(RewardsCoordinator).name;

        addresses[9] = address(Env.proxy.strategyManager());
        configs[9] = pausableConfig;
        names[9] = type(StrategyManager).name;

        /**
         * pods/
         */
        addresses[10] = address(Env.proxy.eigenPodManager());
        configs[10] = pausableConfig;
        names[10] = type(EigenPodManager).name;

        addresses[11] = address(Env.beacon.eigenPod());
        configs[11] = unpausableConfig;
        names[11] = type(EigenPod).name;

        /**
         * strategies/
         */
        addresses[12] = address(Env.proxy.eigenStrategy());
        configs[12] = pausableConfig;
        names[12] = type(EigenStrategy).name;

        addresses[13] = address(Env.beacon.strategyBase());
        configs[13] = unpausableConfig;
        names[13] = type(StrategyBase).name;

        addresses[14] = address(Env.proxy.strategyFactory());
        configs[14] = pausableConfig;
        names[14] = type(StrategyFactory).name;

        /**
         * multichain/
         */
        addresses[15] = address(Env.proxy.bn254CertificateVerifier());
        configs[15] = unpausableConfig;
        names[15] = type(BN254CertificateVerifier).name;

        addresses[16] = address(Env.proxy.crossChainRegistry());
        configs[16] = pausableConfig;
        names[16] = type(CrossChainRegistry).name;

        addresses[17] = address(Env.proxy.ecdsaCertificateVerifier());
        configs[17] = unpausableConfig;
        names[17] = type(ECDSACertificateVerifier).name;

        addresses[18] = address(Env.proxy.operatorTableUpdater());
        configs[18] = pausableConfig;
        names[18] = type(OperatorTableUpdater).name;

        /**
         * avs/
         */
        addresses[19] = address(Env.proxy.taskMailbox());
        configs[19] = unpausableConfig;
        names[19] = type(TaskMailbox).name;

        /**
         * token
         */
        addresses[20] = address(Env.proxy.beigen());
        configs[20] = unpausableConfig;
        names[20] = type(BackingEigen).name;

        addresses[21] = address(Env.proxy.eigen());
        configs[21] = unpausableConfig;
        names[21] = type(Eigen).name;

        // Append to the multisig calls
        calls.append({
            to: address(Env.proxy.protocolRegistry()),
            data: abi.encodeWithSelector(IProtocolRegistry.ship.selector, addresses, configs, names, Env.deployVersion())
        });

        // Now, if we have any strategy base TVLLimits, we need to add them to the protocol registry
        uint256 count = Env.instance.strategyBaseTVLLimits_Count();
        if (count > 0) {
            address[] memory strategyAddresses = new address[](count);
            IProtocolRegistryTypes.DeploymentConfig[] memory strategyConfigs =
                new IProtocolRegistryTypes.DeploymentConfig[](count);
            string[] memory strategyNames = new string[](count);
            for (uint256 i = 0; i < count; i++) {
                strategyAddresses[i] = address(Env.instance.strategyBaseTVLLimits(i));
                strategyConfigs[i] = pausableConfig;
                strategyNames[i] = string.concat(type(StrategyBaseTVLLimits).name, "_", Strings.toString(i));
            }

            calls.append({
                to: address(Env.proxy.protocolRegistry()),
                data: abi.encodeWithSelector(
                    IProtocolRegistry.ship.selector, strategyAddresses, strategyConfigs, strategyNames, Env.deployVersion()
                )
            });
        }
    }

    function testScript() public virtual override {
        if (!Env.isCoreProtocolDeployed()) {
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
