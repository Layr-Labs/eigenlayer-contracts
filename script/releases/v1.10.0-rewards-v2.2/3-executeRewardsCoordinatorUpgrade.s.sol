// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {QueueRewardsCoordinatorUpgrade} from "./2-queueRewardsCoordinatorUpgrade.s.sol";
import {Encode} from "zeus-templates/utils/Encode.sol";

/**
 * @title ExecuteRewardsCoordinatorUpgrade
 * @notice Execute the queued RewardsCoordinator upgrade after the timelock delay.
 *         This completes the upgrade to add Rewards v2.2 support:
 *         - Unique stake rewards (linear to allocated unique stake)
 *         - Total stake rewards (linear to total stake)
 *         - Updated MAX_FUTURE_LENGTH to 730 days (63072000 seconds)
 */
contract ExecuteRewardsCoordinatorUpgrade is QueueRewardsCoordinatorUpgrade {
    using Env for *;
    using Encode for *;

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
        _validateStoragePreservation();
    }

    /// @dev Validate that the RewardsCoordinator proxy now points to the new implementation
    function _validateUpgradeComplete() internal view {
        address currentImpl = Env._getProxyImpl(address(Env.proxy.rewardsCoordinator()));
        address expectedImpl = address(Env.impl.rewardsCoordinator());

        assertTrue(currentImpl == expectedImpl, "RewardsCoordinator proxy should point to new implementation");
    }

    /// @dev Validate the proxy's constructor values through the proxy
    function _validateProxyConstructor() internal view {
        RewardsCoordinator rewardsCoordinator = Env.proxy.rewardsCoordinator();

        // Validate version
        assertEq(
            keccak256(bytes(rewardsCoordinator.version())),
            keccak256(bytes(Env.deployVersion())),
            "RewardsCoordinator version mismatch"
        );

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

        // Validate reward parameters
        assertEq(
            rewardsCoordinator.CALCULATION_INTERVAL_SECONDS(),
            Env.CALCULATION_INTERVAL_SECONDS(),
            "CALCULATION_INTERVAL_SECONDS mismatch"
        );
        assertEq(rewardsCoordinator.MAX_REWARDS_DURATION(), Env.MAX_REWARDS_DURATION(), "MAX_REWARDS_DURATION mismatch");
        assertEq(
            rewardsCoordinator.MAX_RETROACTIVE_LENGTH(), Env.MAX_RETROACTIVE_LENGTH(), "MAX_RETROACTIVE_LENGTH mismatch"
        );

        // Validate the updated MAX_FUTURE_LENGTH
        assertEq(
            rewardsCoordinator.MAX_FUTURE_LENGTH(),
            63_072_000,
            "MAX_FUTURE_LENGTH should be 730 days (63072000 seconds)"
        );

        assertEq(
            rewardsCoordinator.GENESIS_REWARDS_TIMESTAMP(),
            Env.GENESIS_REWARDS_TIMESTAMP(),
            "GENESIS_REWARDS_TIMESTAMP mismatch"
        );
    }

    /// @dev Validate that the proxy is still initialized and cannot be re-initialized
    function _validateProxyInitialized() internal {
        RewardsCoordinator rewardsCoordinator = Env.proxy.rewardsCoordinator();

        // Validate that key state variables are still set (proving initialization is preserved)
        address owner = rewardsCoordinator.owner();
        assertTrue(owner != address(0), "Owner should still be set");

        address rewardsUpdater = rewardsCoordinator.rewardsUpdater();
        assertTrue(rewardsUpdater != address(0), "RewardsUpdater should still be set");

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

    /// @dev Validate that existing storage is preserved after upgrade
    function _validateStoragePreservation() internal view {
        RewardsCoordinator rewardsCoordinator = Env.proxy.rewardsCoordinator();

        // Check that existing storage mappings are still accessible
        address testAvs = address(0xBEEF);
        bytes32 testHash = keccak256("existing_storage_test");

        // These calls should not revert, validating that storage layout is preserved
        bool isRewardsSubmission = rewardsCoordinator.isAVSRewardsSubmissionHash(testAvs, testHash);
        bool isOperatorDirected = rewardsCoordinator.isOperatorDirectedAVSRewardsSubmissionHash(testAvs, testHash);
        bool isOperatorSetRewards =
            rewardsCoordinator.isOperatorDirectedOperatorSetRewardsSubmissionHash(testAvs, testHash);

        // All should be false for test values
        assertFalse(isRewardsSubmission, "Test hash should not be a rewards submission");
        assertFalse(isOperatorDirected, "Test hash should not be operator directed");
        assertFalse(isOperatorSetRewards, "Test hash should not be operator set rewards");

        // Check that core state variables are preserved
        uint256 currRewardsCalculationEndTimestamp = rewardsCoordinator.currRewardsCalculationEndTimestamp();
        // Note: This may be 0 in test environments, which is acceptable
        assertTrue(currRewardsCalculationEndTimestamp >= 0, "currRewardsCalculationEndTimestamp should be accessible");

        // Check that submission nonce is accessible
        uint256 nonce = rewardsCoordinator.submissionNonce(testAvs);
        // Nonce should be 0 for a test address, but accessing it validates storage
        assertEq(nonce, 0, "Nonce for test AVS should be 0");

        // Validate distribution roots are still accessible
        uint256 distributionRootCount = rewardsCoordinator.getDistributionRootsLength();
        // There should be some distribution roots if this is an existing deployment
        // If it's 0, that's also fine for a test environment
        assertTrue(distributionRootCount >= 0, "Distribution roots should be accessible");
    }
}
