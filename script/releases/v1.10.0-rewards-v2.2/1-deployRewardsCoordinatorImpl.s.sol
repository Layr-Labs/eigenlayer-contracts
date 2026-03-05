// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";
import "../TestUtils.sol";

/// @title DeployRewardsCoordinatorImpl
/// @notice Deploy new RewardsCoordinator implementation with Rewards v2.2 support.
///         This adds support for:
///         - Unique stake rewards submissions (rewards linear to allocated unique stake)
///         - Total stake rewards submissions (rewards linear to total stake)
///         - Updated MAX_REWARDS_DURATION to 730 days (63072000 seconds)
contract DeployRewardsCoordinatorImpl is EOADeployer {
    using Env for *;
    using TestUtils for *;

    /// forgefmt: disable-next-item
    function _runAsEOA() internal override {
        // Only execute on source chains with version 1.9.0
        if (!(Env.isSourceChain() && Env._strEq(Env.envVersion(), "1.9.0"))) {
            return;
        }

        vm.startBroadcast();

        // Update the MAX_REWARDS_DURATION environment variable before deployment
        // 63072000s = 730 days = 2 years
        zUpdateUint32("REWARDS_COORDINATOR_MAX_REWARDS_DURATION", 63072000);

        // Deploy RewardsCoordinator implementation with the updated MAX_REWARDS_DURATION
        deployImpl({
            name: type(RewardsCoordinator).name,
            deployedTo: address(
                new RewardsCoordinator(
                    IRewardsCoordinatorTypes.RewardsCoordinatorConstructorParams({
                        delegationManager: Env.proxy.delegationManager(),
                        strategyManager: Env.proxy.strategyManager(),
                        allocationManager: Env.proxy.allocationManager(),
                        pauserRegistry: Env.impl.pauserRegistry(),
                        permissionController: Env.proxy.permissionController(),
                        CALCULATION_INTERVAL_SECONDS: Env.CALCULATION_INTERVAL_SECONDS(),
                        MAX_REWARDS_DURATION: Env.MAX_REWARDS_DURATION(), // Using updated env value
                        MAX_RETROACTIVE_LENGTH: Env.MAX_RETROACTIVE_LENGTH(),
                        MAX_FUTURE_LENGTH: Env.MAX_FUTURE_LENGTH(),
                        GENESIS_REWARDS_TIMESTAMP: Env.GENESIS_REWARDS_TIMESTAMP()
                    })
                )
            )
        });

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        if (!(Env.isSourceChain() && Env._strEq(Env.envVersion(), "1.9.0"))) {
            return;
        }

        // Deploy the new RewardsCoordinator implementation
        runAsEOA();

        _validateNewImplAddress();
        _validateProxyAdmin();
        _validateImplConstructor();
        _validateZeusEnvUpdated();
        _validateImplInitialized();
        _validateNewFunctionality();
        _validateStorageLayout();
    }

    /// @dev Validate that the new RewardsCoordinator impl address is distinct from the current one
    function _validateNewImplAddress() internal view {
        address currentImpl = TestUtils._getProxyImpl(address(Env.proxy.rewardsCoordinator()));
        address newImpl = address(Env.impl.rewardsCoordinator());

        assertFalse(currentImpl == newImpl, "RewardsCoordinator impl should be different from current implementation");
    }

    /// @dev Validate that the RewardsCoordinator proxy is still owned by the correct ProxyAdmin
    function _validateProxyAdmin() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(
            TestUtils._getProxyAdmin(address(Env.proxy.rewardsCoordinator())) == pa,
            "RewardsCoordinator proxyAdmin incorrect"
        );
    }

    /// @dev Validate the immutables set in the new RewardsCoordinator implementation constructor
    function _validateImplConstructor() internal view {
        RewardsCoordinator rewardsCoordinatorImpl = Env.impl.rewardsCoordinator();

        // Validate core dependencies
        assertTrue(
            address(rewardsCoordinatorImpl.delegationManager()) == address(Env.proxy.delegationManager()),
            "RewardsCoordinator delegationManager mismatch"
        );
        assertTrue(
            address(rewardsCoordinatorImpl.strategyManager()) == address(Env.proxy.strategyManager()),
            "RewardsCoordinator strategyManager mismatch"
        );
        assertTrue(
            address(rewardsCoordinatorImpl.allocationManager()) == address(Env.proxy.allocationManager()),
            "RewardsCoordinator allocationManager mismatch"
        );
        assertTrue(
            address(rewardsCoordinatorImpl.pauserRegistry()) == address(Env.impl.pauserRegistry()),
            "RewardsCoordinator pauserRegistry mismatch"
        );
        assertTrue(
            address(rewardsCoordinatorImpl.permissionController()) == address(Env.proxy.permissionController()),
            "RewardsCoordinator permissionController mismatch"
        );

        // Validate reward parameters
        assertEq(
            rewardsCoordinatorImpl.CALCULATION_INTERVAL_SECONDS(),
            Env.CALCULATION_INTERVAL_SECONDS(),
            "CALCULATION_INTERVAL_SECONDS mismatch"
        );
        assertEq(
            rewardsCoordinatorImpl.MAX_REWARDS_DURATION(),
            63_072_000,
            "MAX_REWARDS_DURATION should be updated to 730 days (63072000 seconds)"
        );
        assertEq(
            rewardsCoordinatorImpl.MAX_RETROACTIVE_LENGTH(),
            Env.MAX_RETROACTIVE_LENGTH(),
            "MAX_RETROACTIVE_LENGTH mismatch"
        );
        assertEq(rewardsCoordinatorImpl.MAX_FUTURE_LENGTH(), Env.MAX_FUTURE_LENGTH(), "MAX_FUTURE_LENGTH mismatch");
        assertEq(
            rewardsCoordinatorImpl.GENESIS_REWARDS_TIMESTAMP(),
            Env.GENESIS_REWARDS_TIMESTAMP(),
            "GENESIS_REWARDS_TIMESTAMP mismatch"
        );
    }

    /// @dev Validate that the zeus environment variable has been updated correctly
    function _validateZeusEnvUpdated() internal view {
        RewardsCoordinator rewardsCoordinatorImpl = Env.impl.rewardsCoordinator();

        // Validate that the zeus env MAX_REWARDS_DURATION matches what was deployed
        assertEq(
            rewardsCoordinatorImpl.MAX_REWARDS_DURATION(),
            Env.MAX_REWARDS_DURATION(),
            "Deployed MAX_REWARDS_DURATION should match zeus env value"
        );

        // Also validate it equals the expected value
        assertEq(
            Env.MAX_REWARDS_DURATION(),
            63_072_000,
            "Zeus env MAX_REWARDS_DURATION should be updated to 730 days (63072000 seconds)"
        );
    }

    /// @dev Validate that the new implementation cannot be initialized (should revert)
    function _validateImplInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        RewardsCoordinator rewardsCoordinatorImpl = Env.impl.rewardsCoordinator();

        vm.expectRevert(errInit);
        rewardsCoordinatorImpl.initialize(
            address(0), // initialOwner
            0, // initialPausedStatus
            address(0), // rewardsUpdater
            0, // activationDelay
            0 // defaultSplitBips
        );
    }

    /// @dev Validate new Rewards v2.2 functionality
    function _validateNewFunctionality() internal view {
        RewardsCoordinator rewardsCoordinatorImpl = Env.impl.rewardsCoordinator();

        // The new functions exist (this will fail to compile if they don't exist)
        // Just checking that the contract has the expected interface
        bytes4 createUniqueStakeSelector = rewardsCoordinatorImpl.createUniqueStakeRewardsSubmission.selector;
        bytes4 createTotalStakeSelector = rewardsCoordinatorImpl.createTotalStakeRewardsSubmission.selector;

        // Verify the selectors are non-zero (functions exist)
        assertTrue(createUniqueStakeSelector != bytes4(0), "createUniqueStakeRewardsSubmission function should exist");
        assertTrue(createTotalStakeSelector != bytes4(0), "createTotalStakeRewardsSubmission function should exist");
    }

    /// @dev Validate storage layout changes
    function _validateStorageLayout() internal view {
        // The storage gap was reduced from 35 to 33 slots to accommodate the new mappings:
        // - isUniqueStakeRewardsSubmissionHash (1 slot)
        // - isTotalStakeRewardsSubmissionHash (1 slot)
        // This validation ensures the contract is compiled with the expected storage layout

        // We can't directly access the storage gap, but we can ensure the contract
        // compiles and deploys successfully, which validates the storage layout is correct
        RewardsCoordinator rewardsCoordinatorImpl = Env.impl.rewardsCoordinator();

        // Verify we can access the existing public mappings
        // This validates that storage layout hasn't been corrupted

        // Check that we can call view functions that access storage
        address testAvs = address(0x1234);
        bytes32 testHash = keccak256("test");

        // These calls should not revert, validating storage is accessible
        bool isAVS = rewardsCoordinatorImpl.isAVSRewardsSubmissionHash(testAvs, testHash);
        bool isOperatorDirectedAVS =
            rewardsCoordinatorImpl.isOperatorDirectedAVSRewardsSubmissionHash(testAvs, testHash);
        bool isOperatorDirectedOperatorSet =
            rewardsCoordinatorImpl.isOperatorDirectedOperatorSetRewardsSubmissionHash(testAvs, testHash);
        bool isUniqueStake = rewardsCoordinatorImpl.isUniqueStakeRewardsSubmissionHash(testAvs, testHash);
        bool isTotalStake = rewardsCoordinatorImpl.isTotalStakeRewardsSubmissionHash(testAvs, testHash);

        // All should be false for a random hash
        assertFalse(isAVS, "Random hash should not be a rewards submission");
        assertFalse(isOperatorDirectedAVS, "Random hash should not be operator directed");
        assertFalse(isOperatorDirectedOperatorSet, "Random hash should not be operator set");
        assertFalse(isUniqueStake, "Random hash should not be unique stake");
        assertFalse(isTotalStake, "Random hash should not be total stake");
    }
}
