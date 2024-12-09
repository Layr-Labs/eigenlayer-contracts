// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Deploy} from "./1-deployContracts.s.sol";

import {MultisigCall, MultisigCallUtils, MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import {EigenLabsUpgrade} from "../EigenLabsUpgrade.s.sol";
import "zeus-templates/utils/EncGnosisSafe.sol";
import {MultisigCallUtils, MultisigCall} from "zeus-templates/utils/MultisigCallUtils.sol";
import {IMultiSend} from "zeus-templates/interfaces/IMultiSend.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

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
    using EigenLabsUpgrade for *;
    using EncGnosisSafe for *;
    using MultisigCallUtils for *;
    using Strings for *;

    MultisigCall[] private _executorCalls;
    MultisigCall[] private _opsCalls;

    function _getMultisigTransactionCalldata() internal returns (bytes memory) {

        // MultisigCall[] storage executorCalls = MultisigCallUtils.newCalls();

        _executorCalls
            .append({
                to: zDeployedContract(type(ProxyAdmin).name),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(payable(zDeployedProxy(type(AVSDirectory).name))),
                        zDeployedImpl(type(AVSDirectory).name)
                    )
                )
            })
            .append({
                to: zDeployedContract(type(ProxyAdmin).name),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(payable(zDeployedProxy(type(DelegationManager).name))),
                        zDeployedImpl(type(DelegationManager).name)
                    )
                )
            })
            .append({
                to: zDeployedContract(type(ProxyAdmin).name),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(payable(zDeployedProxy(type(RewardsCoordinator).name))),
                        zDeployedImpl(type(RewardsCoordinator).name)
                    )
                )
            })
            .append({
                to: zDeployedContract(type(ProxyAdmin).name),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(payable(zDeployedProxy(type(StrategyManager).name))),
                        zDeployedImpl(type(StrategyManager).name)
                    )
                )
            })
            .append({
                to: zDeployedContract(type(ProxyAdmin).name),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(payable(zDeployedProxy(type(EigenPodManager).name))),
                        zDeployedImpl(type(EigenPodManager).name)
                    )
                )
            })
            .append({
                to: zDeployedProxy(type(EigenPod).name),
                data: abi.encodeCall(
                    UpgradeableBeacon.upgradeTo,
                    (
                        zDeployedImpl(type(EigenPod).name)
                    )
                )
            })
            .append({
                to: zDeployedContract(type(ProxyAdmin).name),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(payable(zDeployedProxy(type(EigenStrategy).name))),
                        zDeployedImpl(type(EigenStrategy).name)
                    )
                )
            })  
            .append({
                to: zDeployedContract(type(ProxyAdmin).name),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(payable(zDeployedProxy(type(StrategyFactory).name))),
                        zDeployedImpl(type(StrategyFactory).name)
                    )
                )
            })  
            .append({
                to: zDeployedProxy("StrategyBeacon"), // TODO - this needs to be `StrategyBase.name`
                data: abi.encodeCall(
                    UpgradeableBeacon.upgradeTo,
                    (
                        zDeployedImpl("StrategyBeacon")
                    )
                )
            });

        uint count = zDeployedInstanceCount("PreLongtailStrats");
        for (uint i = 0; i < count; i++) {
            _executorCalls.append({
                to: zDeployedContract(type(ProxyAdmin).name),
                data: abi.encodeCall(
                    ProxyAdmin.upgrade,
                    (
                        ITransparentUpgradeableProxy(zDeployedInstance(type(StrategyBaseTVLLimits).name, i)),
                        zDeployedContract("PreLongtailStrats")
                    )
                )
            });
        }

        return EncGnosisSafe.calldataToExecTransaction({
            from: zAddress("timelockController"),
            to: zAddress("MultiSendCallOnly"),
            data: MultisigCallUtils.encodeMultisendTxs(IMultiSend(address(0)), _executorCalls),
            op: EncGnosisSafe.Operation.DelegateCall
        });
    }

    function options() internal virtual override returns (MultisigOptions memory) {
        return MultisigOptions(
            this._operationsMultisig(),
            Operation.DelegateCall
        );
    }

    function runAsMultisig() internal virtual override {
        (bytes memory call) = _getMultisigTransactionCalldata();

        TimelockController timelock = TimelockController(payable(this._timelock()));
        timelock.schedule(
            this._executorMultisig(), 
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

        TimelockController timelock = TimelockController(payable(this._timelock()));
        bytes memory call = _getMultisigTransactionCalldata();
        bytes32 txHash = timelock.hashOperation(this._executorMultisig(), 0, call, 0, 0);
        assertEq(timelock.isOperationPending(txHash), true, "Transaction should be queued.");
    }
}
