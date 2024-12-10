// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Deploy} from "./1-deployContracts.s.sol";
import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
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
    using Env for *;
    using Encode for *;

    function _getMultisigTransactionCalldata() internal returns (bytes memory) {
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls()
            /// core/
            .append({
                to: Env.proxyAdmin(),
                data: Encode.proxyAdmin.upgrade({
                    proxy: address(Env.proxy.avsDirectory()),
                    impl: address(Env.impl.avsDirectory())
                })
            })
            .append({
                to: Env.proxyAdmin(),
                data: Encode.proxyAdmin.upgrade({
                    proxy: address(Env.proxy.delegationManager()),
                    impl: address(Env.impl.delegationManager())
                })
            })
            .append({
                to: Env.proxyAdmin(),
                data: Encode.proxyAdmin.upgrade({
                    proxy: address(Env.proxy.rewardsCoordinator()),
                    impl: address(Env.impl.rewardsCoordinator())
                })
            })
            .append({
                to: Env.proxyAdmin(),
                data: Encode.proxyAdmin.upgrade({
                    proxy: address(Env.proxy.strategyManager()),
                    impl: address(Env.impl.strategyManager())
                })
            })
            /// pods/
            .append({
                to: address(Env.proxy.eigenPod()),
                data: Encode.upgradeableBeacon.upgradeTo({
                    newImpl: address(Env.impl.eigenPod())
                })
            })
            .append({
                to: Env.proxyAdmin(),
                data: Encode.proxyAdmin.upgrade({
                    proxy: address(Env.proxy.eigenPodManager()),
                    impl: address(Env.impl.eigenPodManager())
                })
            })
            /// strategies/
            .append({
                to: Env.proxyAdmin(),
                data: Encode.proxyAdmin.upgrade({
                    proxy: address(Env.proxy.eigenStrategy()),
                    impl: address(Env.impl.eigenStrategy())
                })
            })
            .append({
                to: address(Env.proxy.strategyBase()),
                data: Encode.upgradeableBeacon.upgradeTo({
                    newImpl: address(Env.impl.strategyBase())
                })
            })
            .append({
                to: Env.proxyAdmin(),
                data: Encode.proxyAdmin.upgrade({
                    proxy: address(Env.proxy.strategyFactory()),
                    impl: address(Env.impl.strategyFactory())
                })
            });

        /// Add call to upgrade each pre-longtail strategy instance
        uint count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint i = 0; i < count; i++) {
            address proxyInstance = address(Env.instance.strategyBaseTVLLimits(i));

            executorCalls.append({
                to: Env.proxyAdmin(),
                data: Encode.proxyAdmin.upgrade({
                    proxy: proxyInstance,
                    impl: address(Env.impl.strategyFactory())
                })
            });
        }

        return Encode.gnosisSafe.execTransaction({
            from: address(Env.timelockController()),
            to: Env.multiSendCallOnly(),
            op: Encode.Operation.DelegateCall,
            data: Encode.multiSend(executorCalls)
        });
    }

    function options() internal virtual override returns (MultisigOptions memory) {
        return MultisigOptions(
            Env.opsMultisig(),
            Operation.DelegateCall
        );
    }

    function runAsMultisig() internal virtual override {
        bytes memory call = _getMultisigTransactionCalldata();

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

        TimelockController timelock = Env.timelockController();
        bytes memory call = _getMultisigTransactionCalldata();
        bytes32 txHash = timelock.hashOperation(Env.executorMultisig(), 0, call, 0, 0);
        assertEq(timelock.isOperationPending(txHash), true, "Transaction should be queued.");
    }
}
