// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {Queue} from "./2-multisig.s.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract Execute is Queue {
    using Env for *;

    function _runAsMultisig() prank(Env.protocolCouncilMultisig()) internal override(Queue) {
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

    function testScript() public virtual override(Queue){
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
        assertFalse(timelock.isOperationPending(txHash), "Transaction should NOT be queued.");

        // 1. Queue Upgrade
        Queue._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        // 2. Warp past delay
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA
        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");

        // 3- execute
        execute();

        assertTrue(timelock.isOperationDone(txHash), "Transaction should be complete.");

        // 4. Validate
        _validateNewImplAddresses(true);
        _validateProxyConstructors();
        _validateProxyDomainSeparators();
    }

    function _validateProxyDomainSeparators() internal view {
        bytes32 zeroDomainSeparator = bytes32(0);

        assertFalse(Env.proxy.avsDirectory().domainSeparator() == zeroDomainSeparator, "avsD.domainSeparator is zero");
        assertFalse(Env.proxy.delegationManager().domainSeparator() == zeroDomainSeparator, "dm.domainSeparator is zero");
        assertFalse(Env.proxy.strategyManager().domainSeparator() == zeroDomainSeparator, "rc.domainSeparator is zero");
    }

    function _validateProxyConstructors() internal view {
        AVSDirectory avsDirectory = Env.proxy.avsDirectory();
        assertTrue(avsDirectory.delegation() == Env.proxy.delegationManager(), "avsD.dm invalid");
        assertTrue(avsDirectory.pauserRegistry() == Env.impl.pauserRegistry(), "avsD.pR invalid");

        DelegationManager delegation = Env.proxy.delegationManager();
        assertTrue(delegation.strategyManager() == Env.proxy.strategyManager(), "dm.sm invalid");
        assertTrue(delegation.eigenPodManager() == Env.proxy.eigenPodManager(), "dm.epm invalid");
        assertTrue(delegation.allocationManager() == Env.proxy.allocationManager(), "dm.alm invalid");
        assertTrue(delegation.pauserRegistry() == Env.impl.pauserRegistry(), "dm.pR invalid");
        assertTrue(delegation.permissionController() == Env.proxy.permissionController(), "dm.pc invalid");
        assertTrue(delegation.minWithdrawalDelayBlocks() == Env.MIN_WITHDRAWAL_DELAY(), "dm.withdrawalDelay invalid");

        RewardsCoordinator rewards = Env.proxy.rewardsCoordinator();
        assertTrue(rewards.delegationManager() == Env.proxy.delegationManager(), "rc.dm invalid");
        assertTrue(rewards.strategyManager() == Env.proxy.strategyManager(), "rc.sm invalid");
        assertTrue(rewards.allocationManager() == Env.proxy.allocationManager(), "rc.alm invalid");
        assertTrue(rewards.pauserRegistry() == Env.impl.pauserRegistry(), "rc.pR invalid");
        assertTrue(rewards.permissionController() == Env.proxy.permissionController(), "rc.pc invalid");
        assertTrue(rewards.CALCULATION_INTERVAL_SECONDS() == Env.CALCULATION_INTERVAL_SECONDS(), "rc.calcInterval invalid");
        assertTrue(rewards.MAX_REWARDS_DURATION() == Env.MAX_REWARDS_DURATION(), "rc.rewardsDuration invalid");
        assertTrue(rewards.MAX_RETROACTIVE_LENGTH() == Env.MAX_RETROACTIVE_LENGTH(), "rc.retroLength invalid");
        assertTrue(rewards.MAX_FUTURE_LENGTH() == Env.MAX_FUTURE_LENGTH(), "rc.futureLength invalid");
        assertTrue(rewards.GENESIS_REWARDS_TIMESTAMP() == Env.GENESIS_REWARDS_TIMESTAMP(), "rc.genesis invalid");
        
        StrategyManager strategyManager = Env.proxy.strategyManager();
        assertTrue(strategyManager.delegation() == Env.proxy.delegationManager(), "sm.dm invalid");
        assertTrue(strategyManager.pauserRegistry() == Env.impl.pauserRegistry(), "sm.pR invalid");
    }

    function _validateRC() internal view {
        RewardsCoordinator rewards = Env.proxy.rewardsCoordinator();
        assertEq(
            rewards.defaultOperatorSplitBips(),
            Env.DEFAULT_SPLIT_BIPS(),
            "expected defaultOperatorSplitBips"
        );

        uint256 pausedStatusAfter = rewards.paused();

        assertEq(
            Env.REWARDS_PAUSE_STATUS(),
            pausedStatusAfter,
            "expected paused status to be the same before and after initialization"
        );

        assertEq(
            rewards.CALCULATION_INTERVAL_SECONDS(),
            Env.CALCULATION_INTERVAL_SECONDS(),
            "expected REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );
        assertEq(
            rewards.CALCULATION_INTERVAL_SECONDS(),
            1 days,
            "expected REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );
        assertGt(
            rewards.CALCULATION_INTERVAL_SECONDS(),
            0,
            "expected non-zero REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );

        assertEq(rewards.MAX_REWARDS_DURATION(), Env.MAX_REWARDS_DURATION());
        assertGt(rewards.MAX_REWARDS_DURATION(), 0);

        assertEq(
            rewards.MAX_RETROACTIVE_LENGTH(),
            Env.MAX_RETROACTIVE_LENGTH()
        );
        assertGt(rewards.MAX_RETROACTIVE_LENGTH(), 0);

        assertEq(rewards.MAX_FUTURE_LENGTH(), Env.MAX_FUTURE_LENGTH());
        assertGt(rewards.MAX_FUTURE_LENGTH(), 0);

        assertEq(
            rewards.GENESIS_REWARDS_TIMESTAMP(),
            Env.GENESIS_REWARDS_TIMESTAMP()
        );
        assertGt(rewards.GENESIS_REWARDS_TIMESTAMP(), 0);
    }
}