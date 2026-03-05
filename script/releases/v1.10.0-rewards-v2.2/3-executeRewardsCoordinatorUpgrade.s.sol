// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import "../TestUtils.sol";
import {QueueRewardsCoordinatorUpgrade} from "./2-queueRewardsCoordinatorUpgrade.s.sol";
import {Encode} from "zeus-templates/utils/Encode.sol";

/// @title ExecuteRewardsCoordinatorUpgrade
/// @notice Execute the queued RewardsCoordinator upgrade after the timelock delay.
///         This completes the upgrade to add Rewards v2.2 support:
///         - Unique stake rewards (linear to allocated unique stake)
///         - Total stake rewards (linear to total stake)
///         - Updated MAX_REWARDS_DURATION to 730 days (63072000 seconds)
contract ExecuteRewardsCoordinatorUpgrade is QueueRewardsCoordinatorUpgrade {
    using Env for *;
    using Encode for *;
    using TestUtils for *;

    function _runAsMultisig() internal override prank(Env.protocolCouncilMultisig()) {
        if (!(Env.isSourceChain() && Env._strEq(Env.envVersion(), "1.9.0"))) {
            return;
        }

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

    function testScript() public virtual override {
        if (!(Env.isSourceChain() && Env._strEq(Env.envVersion(), "1.9.0"))) {
            return;
        }

        // 1 - Deploy. The new RewardsCoordinator implementation has been deployed
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

        // 2 - Queue. Check that the operation IS ready
        QueueRewardsCoordinatorUpgrade._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued");
        assertFalse(timelock.isOperationReady(txHash), "Transaction should NOT be ready immediately");
        assertFalse(timelock.isOperationDone(txHash), "Transaction should NOT be complete");

        // 3 - Warp past the timelock delay
        vm.warp(block.timestamp + timelock.getMinDelay());
        assertTrue(timelock.isOperationReady(txHash), "Transaction should be ready for execution");

        // 4 - Execute the upgrade
        execute();
        assertTrue(timelock.isOperationDone(txHash), "v1.10.0 RewardsCoordinator upgrade should be complete");

        // 5 - Validate the upgrade was successful
        _validateUpgradeComplete();
        _validateProxyAdmin();
        _validateProxyConstructor();
        _validateProxyInitialized();
        _validateNewFunctionalityThroughProxy();
    }

    /// @dev Validate that the RewardsCoordinator proxy now points to the new implementation
    function _validateUpgradeComplete() internal view {
        address currentImpl = TestUtils._getProxyImpl(address(Env.proxy.rewardsCoordinator()));
        address expectedImpl = address(Env.impl.rewardsCoordinator());

        assertTrue(currentImpl == expectedImpl, "RewardsCoordinator proxy should point to new implementation");
    }

    /// @dev Validate the proxy's constructor values through the proxy
    function _validateProxyConstructor() internal view {
        RewardsCoordinator rewardsCoordinator = Env.proxy.rewardsCoordinator();

        // Validate core dependencies
        assertTrue(
            address(rewardsCoordinator.delegationManager()) == address(Env.proxy.delegationManager()),
            "RewardsCoordinator delegationManager mismatch"
        );
        assertTrue(
            address(rewardsCoordinator.strategyManager()) == address(Env.proxy.strategyManager()),
            "RewardsCoordinator strategyManager mismatch"
        );
        assertTrue(
            address(rewardsCoordinator.allocationManager()) == address(Env.proxy.allocationManager()),
            "RewardsCoordinator allocationManager mismatch"
        );
        assertTrue(
            address(rewardsCoordinator.pauserRegistry()) == address(Env.impl.pauserRegistry()),
            "RewardsCoordinator pauserRegistry mismatch"
        );
        assertTrue(
            address(rewardsCoordinator.permissionController()) == address(Env.proxy.permissionController()),
            "RewardsCoordinator permissionController mismatch"
        );

        // Validate reward parameters
        assertEq(
            rewardsCoordinator.CALCULATION_INTERVAL_SECONDS(),
            Env.CALCULATION_INTERVAL_SECONDS(),
            "CALCULATION_INTERVAL_SECONDS mismatch"
        );

        // Validate the updated MAX_REWARDS_DURATION
        assertEq(
            rewardsCoordinator.MAX_REWARDS_DURATION(),
            63_072_000,
            "MAX_REWARDS_DURATION should be updated to 730 days (63072000 seconds)"
        );

        assertEq(
            rewardsCoordinator.MAX_RETROACTIVE_LENGTH(), Env.MAX_RETROACTIVE_LENGTH(), "MAX_RETROACTIVE_LENGTH mismatch"
        );
        assertEq(rewardsCoordinator.MAX_FUTURE_LENGTH(), Env.MAX_FUTURE_LENGTH(), "MAX_FUTURE_LENGTH mismatch");

        assertEq(
            rewardsCoordinator.GENESIS_REWARDS_TIMESTAMP(),
            Env.GENESIS_REWARDS_TIMESTAMP(),
            "GENESIS_REWARDS_TIMESTAMP mismatch"
        );
    }

    /// @dev Validate that the proxy is still initialized and cannot be re-initialized
    function _validateProxyInitialized() internal {
        RewardsCoordinator rewardsCoordinator = Env.proxy.rewardsCoordinator();

        // Validate the existing initializable state variables are still set
        assertTrue(rewardsCoordinator.paused() == Env.REWARDS_PAUSE_STATUS(), "Paused status should still be set");
        assertTrue(rewardsCoordinator.owner() == Env.opsMultisig(), "Owner should still be set");
        assertTrue(rewardsCoordinator.rewardsUpdater() == Env.REWARDS_UPDATER(), "RewardsUpdater should still be set");
        assertTrue(
            rewardsCoordinator.activationDelay() == Env.ACTIVATION_DELAY(), "Activation delay should still be set"
        );
        assertTrue(
            rewardsCoordinator.defaultOperatorSplitBips() == Env.DEFAULT_SPLIT_BIPS(),
            "Default split bips should still be set"
        );

        // Attempt to re-initialize should fail
        bytes memory errInit = "Initializable: contract is already initialized";
        vm.expectRevert(errInit);
        rewardsCoordinator.initialize(
            address(0x1234), // initialOwner
            0, // initialPausedStatus
            address(0x9ABC), // rewardsUpdater
            0, // activationDelay
            0 // defaultSplitBips
        );
    }

    /// @dev Validate new Rewards v2.2 functionality through the proxy
    function _validateNewFunctionalityThroughProxy() internal view {
        RewardsCoordinator rewardsCoordinator = Env.proxy.rewardsCoordinator();

        // The new functions should be accessible through the proxy
        bytes4 createUniqueStakeSelector = rewardsCoordinator.createUniqueStakeRewardsSubmission.selector;
        bytes4 createTotalStakeSelector = rewardsCoordinator.createTotalStakeRewardsSubmission.selector;

        // Verify the selectors are non-zero (functions exist)
        assertTrue(
            createUniqueStakeSelector != bytes4(0), "createUniqueStakeRewardsSubmission function should exist on proxy"
        );
        assertTrue(
            createTotalStakeSelector != bytes4(0), "createTotalStakeRewardsSubmission function should exist on proxy"
        );

        // Test that we can access the new storage mappings
        address testAvs = address(0xDEAD);
        bytes32 testHash = keccak256("test_rewards_v2.2");

        // These should all return false for test values, but accessing them validates the storage layout
        bool isUniqueStake = rewardsCoordinator.isUniqueStakeRewardsSubmissionHash(testAvs, testHash);
        bool isTotalStake = rewardsCoordinator.isTotalStakeRewardsSubmissionHash(testAvs, testHash);

        assertFalse(isUniqueStake, "Test hash should not be a unique stake submission");
        assertFalse(isTotalStake, "Test hash should not be a total stake submission");
    }
}
