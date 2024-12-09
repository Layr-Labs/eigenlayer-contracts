// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Deploy} from "./1-deployContracts.s.sol";
import "../Env.sol";

import {MultisigCall, MultisigCallUtils, MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "zeus-templates/utils/EncGnosisSafe.sol";
import {IMultiSend} from "zeus-templates/interfaces/IMultiSend.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

/// core/
import "src/contracts/core/AllocationManager.sol";
import "src/contracts/core/AVSDirectory.sol";
import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/RewardsCoordinator.sol";
import "src/contracts/core/StrategyManager.sol";

/// permissions/
import "src/contracts/permissions/PauserRegistry.sol";
import "src/contracts/permissions/PermissionController.sol";

/// pods/
import "src/contracts/pods/EigenPod.sol";
import "src/contracts/pods/EigenPodManager.sol";

/// strategies/
import "src/contracts/strategies/EigenStrategy.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/contracts/strategies/StrategyBaseTVLLimits.sol";
import "src/contracts/strategies/StrategyFactory.sol";

/**
 * Purpose: 
 *      * enqueue a multisig transaction which;
 *             - upgrades all the relevant contracts, and
 *             - unpauses the system.
 *  This should be run via the protocol council multisig.
 */
contract QueueAndUnpause is MultisigBuilder, Deploy {
    using MultisigCallUtils for *;
    using EncGnosisSafe for *;
    using MultisigCallUtils for *;
    using Strings for *;
    using Env for *;

    MultisigCall[] private _executorCalls;
    MultisigCall[] private _opsCalls;

    function _getMultisigTransactionCalldata() internal returns (bytes memory) {

        // MultisigCall[] storage executorCalls = MultisigCallUtils.newCalls();

        _executorCalls
            .append({
                to: Env.proxyAdmin(),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(address(Env.proxy.avsDirectory())),
                        address(Env.impl.avsDirectory())
                    )
                )
            })
            .append({
                to: Env.proxyAdmin(),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(address(Env.proxy.delegationManager())),
                        address(Env.impl.delegationManager())
                    )
                )
            })
            .append({
                to: Env.proxyAdmin(),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(address(Env.proxy.rewardsCoordinator())),
                        address(Env.impl.rewardsCoordinator())
                    )
                )
            })
            .append({
                to: Env.proxyAdmin(),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(address(Env.proxy.strategyManager())),
                        address(Env.impl.strategyManager())
                    )
                )
            })
            .append({
                to: Env.proxyAdmin(),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(address(Env.proxy.eigenPodManager())),
                        address(Env.impl.eigenPodManager())
                    )
                )
            })
            .append({
                to: address(Env.proxy.eigenPod()),
                data: abi.encodeCall(
                    UpgradeableBeacon.upgradeTo,
                    (
                        address(Env.impl.eigenPod())
                    )
                )
            })
            .append({
                to: Env.proxyAdmin(),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(address(Env.proxy.eigenStrategy())),
                        address(Env.impl.eigenStrategy())
                    )
                )
            })  
            .append({
                to: Env.proxyAdmin(),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(address(Env.proxy.strategyFactory())),
                        address(Env.impl.strategyFactory())
                    )
                )
            })  
            .append({
                to: address(Env.proxy.strategyBase()),
                data: abi.encodeCall(
                    UpgradeableBeacon.upgradeTo,
                    (
                        address(Env.impl.strategyBase())
                    )
                )
            });

        uint count = Env.instance.strategyBaseTVLLimits_Count();
        
        // zDeployedInstanceCount("PreLongtailStrats");
        for (uint i = 0; i < count; i++) {
            address instance = address(Env.instance.strategyBaseTVLLimits(i));

            _executorCalls.append({
                to: Env.proxyAdmin(),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(instance),
                        address(Env.impl.strategyBaseTVLLimits())
                    )
                )
            });
        }

        return EncGnosisSafe.calldataToExecTransaction({
            from: address(Env.timelockController()),
            to: Env.multiSendCallOnly(),
            data: MultisigCallUtils.encodeMultisendTxs(IMultiSend(address(0)), _executorCalls),
            op: EncGnosisSafe.Operation.DelegateCall
        });
    }

    function options() internal virtual override returns (MultisigOptions memory) {
        return MultisigOptions(
            Env.opsMultisig(),
            Operation.DelegateCall
        );
    }

    function runAsMultisig() internal virtual override {
        (bytes memory call) = _getMultisigTransactionCalldata();

        TimelockController timelock = Env.timelockController();
        timelock.schedule(
            Env.executorMultisig(), 
            0, 
            call,
            0,
            bytes32(0),  
            timelock.getMinDelay()
        );
    }

    function testDeploy() virtual public override {
        runAsEOA();
        execute();

        // TODO hack
        delete _executorCalls;

        TimelockController timelock = Env.timelockController();
        bytes memory call = _getMultisigTransactionCalldata();
        bytes32 txHash = timelock.hashOperation(Env.executorMultisig(), 0, call, 0, 0);
        assertEq(timelock.isOperationPending(txHash), true, "Transaction should be queued.");
    }
}
