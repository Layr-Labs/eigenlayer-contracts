// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {DeployToken} from "./3-deployToken.s.sol";
import {DeployPauser} from "./2-deployPauser.s.sol";
import {DeployGovernance} from "./1-deployGovernance.s.sol";
import {DeployCore} from "./4-deployCore.s.sol";
import "../Env.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";

/// @notice Queues an upgrade for the previously deployed empty proxies.
/// @dev The Beacons have already been set to the correct implementation contracts in step 4 of the upgrade script.
/// Upgrades the following contracts:
/// - DelegationManager
/// - StrategyManager
/// - AllocationManager
/// - RewardsCoordinator
/// - AVSDirectory
/// - EigenPodManager
/// - StrategyFactory
/// - ReleaseManager
/// - PermissionController
/// - KeyRegistrar
/// - EigenStrategy
contract Queue is MultisigBuilder, DeployCore {
    using Env for *;
    using Encode for *;

    ProxyAdmin public proxyAdmin;

    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        proxyAdmin = ProxyAdmin(address(Env.proxyAdmin()));
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
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls().append({
            to: Env.proxyAdmin(),
            data: abi.encodeCall(
                proxyAdmin.upgradeAndCall,
                (
                    ITransparentUpgradeableProxy(payable(address(Env.proxy.delegationManager()))),
                    address(Env.impl.delegationManager()),
                    abi.encodeCall(DelegationManager.initialize, (0))
                )
            )
        }).append({
            to: Env.proxyAdmin(),
            data: abi.encodeCall(
                proxyAdmin.upgradeAndCall,
                (
                    ITransparentUpgradeableProxy(payable(address(Env.proxy.strategyManager()))),
                    address(Env.impl.strategyManager()),
                    abi.encodeCall(
                        StrategyManager.initialize,
                        (address(Env.executorMultisig()), address(Env.proxy.strategyFactory()), 0)
                    )
                )
            )
        }).append({
            to: Env.proxyAdmin(),
            data: abi.encodeCall(
                proxyAdmin.upgradeAndCall,
                (
                    ITransparentUpgradeableProxy(payable(address(Env.proxy.allocationManager()))),
                    address(Env.impl.allocationManager()),
                    abi.encodeCall(AllocationManager.initialize, (0))
                )
            )
        }).append({
            to: Env.proxyAdmin(),
            data: abi.encodeCall(
                proxyAdmin.upgradeAndCall,
                (
                    ITransparentUpgradeableProxy(payable(address(Env.proxy.rewardsCoordinator()))),
                    address(Env.impl.rewardsCoordinator()),
                    abi.encodeCall(
                        RewardsCoordinator.initialize,
                        (
                            address(Env.opsMultisig()),
                            Env.REWARDS_PAUSE_STATUS(),
                            Env.REWARDS_UPDATER(),
                            Env.ACTIVATION_DELAY(),
                            Env.DEFAULT_SPLIT_BIPS()
                        )
                    )
                )
            )
        }).append({
            to: Env.proxyAdmin(),
            data: abi.encodeCall(
                proxyAdmin.upgradeAndCall,
                (
                    ITransparentUpgradeableProxy(payable(address(Env.proxy.avsDirectory()))),
                    address(Env.impl.avsDirectory()),
                    abi.encodeCall(AVSDirectory.initialize, (address(Env.executorMultisig()), 0))
                )
            )
        }).append({
            to: Env.proxyAdmin(),
            data: abi.encodeCall(
                proxyAdmin.upgradeAndCall,
                (
                    ITransparentUpgradeableProxy(payable(address(Env.proxy.eigenPodManager()))),
                    address(Env.impl.eigenPodManager()),
                    abi.encodeCall(EigenPodManager.initialize, (address(Env.executorMultisig()), 0))
                )
            )
        }).append({
            to: Env.proxyAdmin(),
            data: abi.encodeCall(
                proxyAdmin.upgradeAndCall,
                (
                    ITransparentUpgradeableProxy(payable(address(Env.proxy.strategyFactory()))),
                    address(Env.impl.strategyFactory()),
                    abi.encodeCall(StrategyFactory.initialize, (address(Env.opsMultisig()), 0, Env.beacon.strategyBase()))
                )
            )
        }).append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.releaseManager()),
                impl: address(Env.impl.releaseManager())
            })
        }).append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.permissionController()),
                impl: address(Env.impl.permissionController())
            })
        }).append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.keyRegistrar()),
                impl: address(Env.impl.keyRegistrar())
            })
        }).append({
            to: Env.proxyAdmin(),
            data: abi.encodeCall(
                proxyAdmin.upgradeAndCall,
                (
                    ITransparentUpgradeableProxy(payable(address(Env.proxy.eigenStrategy()))),
                    address(Env.impl.eigenStrategy()),
                    abi.encodeCall(
                        EigenStrategy.initialize, (IEigen(address(Env.proxy.eigen())), IERC20(address(Env.proxy.beigen())))
                    )
                )
            )
        });

        // Add a call to set the proof timestamp setter timestamp to the operations multisig. This value will be set in step 10 of the upgrade.
        executorCalls.append({
            to: address(Env.proxy.eigenPodManager()),
            data: abi.encodeCall(EigenPodManager.setProofTimestampSetter, (address(Env.opsMultisig())))
        });

        return Encode.gnosisSafe.execTransaction({
            from: address(Env.timelockController()),
            to: address(Env.multiSendCallOnly()),
            op: Encode.Operation.DelegateCall,
            data: Encode.multiSend(executorCalls)
        });
    }

    function testScript() public virtual override {
        // Run all the previous steps of the upgrade script
        _mode = OperationalMode.EOA;
        DeployGovernance._runAsEOA();
        DeployPauser._runAsEOA();
        DeployToken._runAsEOA();
        DeployCore._runAsEOA();

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
