// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {Queue} from "./2-multisig.s.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract Execute is Queue {
    using Env for *;

    function _runAsMultisig()
        internal
        override(Queue)
        prank(Env.protocolCouncilMultisig())
    {
        bytes memory calldata_to_executor = _getCalldataToExecutor();

        TimelockController timelock = Env.timelockController();
        timelock.execute({
            target: Env.executorMultisig(),
            value: 0,
            payload: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });
    }

    function testScript() public virtual override(Queue) {
        // 0. Deploy Impls
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
        assertFalse(
            timelock.isOperationPending(txHash),
            "Transaction should NOT be queued."
        );

        // 1. Queue Upgrade
        Queue._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        // 2. Warp past delay
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA
        assertEq(
            timelock.isOperationReady(txHash),
            true,
            "Transaction should be executable."
        );

        // 3- execute
        execute();

        assertTrue(
            timelock.isOperationDone(txHash),
            "Transaction should be complete."
        );

        // 4. Validate
        _validateNewImplAddresses(true);
        _validateProxyConstructors();
        _validateProxyInitialized();
    }

    function _validateProxyConstructors() internal view {
        RewardsCoordinator rewards = Env.proxy.rewardsCoordinator();
        assertTrue(
            rewards.delegationManager() == Env.proxy.delegationManager(),
            "rc.dm invalid"
        );
        assertTrue(
            rewards.strategyManager() == Env.proxy.strategyManager(),
            "rc.sm invalid"
        );
        assertTrue(
            rewards.allocationManager() == Env.proxy.allocationManager(),
            "rc.alm invalid"
        );
        assertTrue(
            rewards.pauserRegistry() == Env.impl.pauserRegistry(),
            "rc.pR invalid"
        );
        assertTrue(
            rewards.permissionController() == Env.proxy.permissionController(),
            "rc.pc invalid"
        );
        assertTrue(
            rewards.CALCULATION_INTERVAL_SECONDS() ==
                Env.CALCULATION_INTERVAL_SECONDS(),
            "rc.calcInterval invalid"
        );
        assertTrue(
            rewards.MAX_REWARDS_DURATION() == Env.MAX_REWARDS_DURATION(),
            "rc.rewardsDuration invalid"
        );
        assertTrue(
            rewards.MAX_RETROACTIVE_LENGTH() == Env.MAX_RETROACTIVE_LENGTH(),
            "rc.retroLength invalid"
        );
        assertTrue(
            rewards.MAX_FUTURE_LENGTH() == Env.MAX_FUTURE_LENGTH(),
            "rc.futureLength invalid"
        );
        assertTrue(
            rewards.GENESIS_REWARDS_TIMESTAMP() ==
                Env.GENESIS_REWARDS_TIMESTAMP(),
            "rc.genesis invalid"
        );
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateProxyInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        RewardsCoordinator rewards = Env.proxy.rewardsCoordinator();
        vm.expectRevert(errInit);
        rewards.initialize(address(0), 0, address(0), 0, 0);
    }
}
