// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Deploy} from "./1-deployContracts.s.sol";
import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

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

    function _runAsMultisig() prank(Env.opsMultisig()) internal virtual override {
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

    /// @dev Get the calldata to be sent from the timelock to the executor
    function _getCalldataToExecutor() internal returns (bytes memory) {
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
                to: address(Env.beacon.eigenPod()),
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
                to: address(Env.beacon.strategyBase()),
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
                    impl: address(Env.impl.strategyBaseTVLLimits())
                })
            });
        }

        // /// Finally, add a call unpausing the EigenPodManager
        // /// We will end up pausing it in step 3, so the unpause will
        // /// go through as part of execution (step 5)
        executorCalls.append({
            to: address(Env.proxy.eigenPodManager()),
            data: abi.encodeCall(Pausable.unpause, 0)
        });

        return Encode.gnosisSafe.execTransaction({
            from: address(Env.timelockController()),
            to: Env.multiSendCallOnly(),
            op: Encode.Operation.DelegateCall,
            data: Encode.multiSend(executorCalls)
        });
    }

    function testScript() public virtual override {
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
  
        execute();

        // Check that the upgrade has been added to the timelock
        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
    }
}
