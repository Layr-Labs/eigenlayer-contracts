// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {MockERC20} from "src/test/mocks/MockERC20.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/contracts/core/SlashEscrowFactory.sol";
import "src/contracts/core/SlashEscrow.sol";

contract SlashEscrowFactoryUnitTests is EigenLayerUnitTestSetup, ISlashEscrowFactoryEvents {
    /// @notice default address for burning slashed shares and transferring underlying tokens
    address public constant DEFAULT_BURN_ADDRESS = 0x00000000000000000000000000000000000E16E4;

    /// @notice The pause status for the `releaseSlashEscrow` function.
    /// @dev Allows all burn or redistribution outflows to be temporarily halted.
    uint8 public constant PAUSED_BURN_OR_REDISTRIBUTE_SHARES = 0;

    SlashEscrowFactory router;

    OperatorSet defaultOperatorSet;
    IStrategy defaultStrategy;
    MockERC20 defaultToken;
    uint defaultSlashId;
    address defaultRedistributionRecipient;
    address defaultOwner;

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();

        defaultOperatorSet = OperatorSet(cheats.randomAddress(), 0);
        defaultStrategy = IStrategy(cheats.randomAddress());
        defaultToken = new MockERC20();
        defaultSlashId = 1;
        defaultRedistributionRecipient = address(cheats.randomAddress());
        defaultOwner = address(cheats.randomAddress());

        allocationManagerMock.setRedistributionRecipient(defaultOperatorSet, defaultRedistributionRecipient);

        router = SlashEscrowFactory(
            address(
                new TransparentUpgradeableProxy(
                    address(
                        new SlashEscrowFactory(
                            IAllocationManager(address(allocationManagerMock)),
                            IStrategyManager(address(strategyManagerMock)),
                            IPauserRegistry(address(pauserRegistry)),
                            ISlashEscrow(address(new SlashEscrow())),
                            "1.0.0"
                        )
                    ),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(SlashEscrowFactory.initialize.selector, defaultOwner, 0, uint32(4 days / 12 seconds))
                )
            )
        );
    }

    function _rollForwardDefaultBurnOrRedistributionDelay() internal {
        cheats.roll(block.number + 4 days / 12 seconds + 1);
    }

    /// @dev Sets the return value for the next call to `strategy.underlyingToken()`.
    function _mockStrategyUnderlyingTokenCall(IStrategy strategy, address underlyingToken) internal {
        cheats.mockCall(address(strategy), abi.encodeWithSelector(IStrategy.underlyingToken.selector), abi.encode(underlyingToken));
    }

    /// @dev Returns the pending underlying amount for a given strategy and token.
    function _getPendingUnderlyingAmountForStrategy(OperatorSet memory operatorSet, uint slashId, IStrategy strategy, MockERC20 token)
        internal
        returns (uint)
    {
        _mockStrategyUnderlyingTokenCall(strategy, address(token));
        return router.getPendingUnderlyingAmountForStrategy(operatorSet, slashId, strategy);
    }

    /// @dev Starts a burn or redistribution for a given strategy and token.
    /// - Calls as the `StrategyManager`.
    /// - Asserts that the `StartBurnOrRedistribution` event is emitted.
    /// - Mocks the strategy sending the underlying token to the `SlashEscrow`.
    function _initiateSlashEscrow(OperatorSet memory operatorSet, uint slashId, IStrategy strategy, MockERC20 token, uint underlyingAmount)
        internal
    {
        cheats.prank(address(strategyManagerMock));
        cheats.expectEmit(true, true, true, true);
        emit StartBurnOrRedistribution(operatorSet, slashId, strategy, uint32(block.number));
        router.initiateSlashEscrow(operatorSet, slashId, strategy);
        deal(address(token), address(router.getSlashEscrow(operatorSet, slashId)), underlyingAmount);
    }

    /// @dev Calls the `releaseSlashEscrow` function as the redistribution recipient.
    /// - Asserts that the `BurnOrRedistribution` event is emitted only for strategies whose delay has elapsed.
    function _releaseSlashEscrow(OperatorSet memory operatorSet, uint slashId) internal {
        (IStrategy[] memory strategies) = router.getPendingStrategiesForSlashId(operatorSet, slashId);

        address redistributionRecipient = allocationManagerMock.getRedistributionRecipient(operatorSet);
        uint startBlock = router.getBurnOrRedistributionStartBlock(operatorSet, slashId);

        for (uint i = strategies.length; i > 0; i--) {
            uint strategyDelay = router.getStrategyBurnOrRedistributionDelay(strategies[i - 1]);
            if (block.number > startBlock + strategyDelay) {
                cheats.expectEmit(true, true, true, true);
                emit BurnOrRedistribution(operatorSet, slashId, strategies[i - 1], redistributionRecipient);
            }
        }

        cheats.prank(defaultRedistributionRecipient);
        router.releaseSlashEscrow(operatorSet, slashId);
    }

    /// @dev Asserts that the operator set and slash ID are pending, and that the strategy and underlying amount are in the pending burn or redistributions.
    function _checkStartBurnOrRedistributions(
        OperatorSet memory operatorSet,
        uint slashId,
        IStrategy strategy,
        MockERC20 token,
        uint expectedUnderlyingAmount,
        uint expectedCount
    ) internal {
        // Assert that the operator set and slash ID are pending.
        assertTrue(router.isPendingOperatorSet(operatorSet));
        assertTrue(router.isPendingSlashId(operatorSet, slashId));

        // Assert that the underlying amount in escrow for the (operator set, slash ID, strategy) is correct.
        assertEq(_getPendingUnderlyingAmountForStrategy(operatorSet, slashId, strategy, token), expectedUnderlyingAmount);

        // Assert that the number of pending burn or redistributions is correct.
        (IStrategy[] memory strategies) = router.getPendingStrategiesForSlashId(operatorSet, slashId);

        assertEq(strategies.length, expectedCount);
        assertEq(router.getPendingStrategiesForSlashIdCount(operatorSet, slashId), expectedCount);

        // Assert that the start block for the (operator set, slash ID) is correct.
        assertEq(router.getBurnOrRedistributionStartBlock(operatorSet, slashId), uint32(block.number));
    }
}

contract SlashEscrowFactoryUnitTests_initialize is SlashEscrowFactoryUnitTests {
    function test_initialize() public {
        assertEq(address(router.allocationManager()), address(allocationManagerMock));
        assertEq(address(router.strategyManager()), address(strategyManagerMock));
        assertEq(address(router.pauserRegistry()), address(pauserRegistry));
        assertEq(router.paused(), 0);
    }
}

contract SlashEscrowFactoryUnitTests_initiateSlashEscrow is SlashEscrowFactoryUnitTests {
    /// @dev Asserts only the `StrategyManager` can call `initiateSlashEscrow`.
    function test_initiateSlashEscrow_onlyStrategyManager() public {
        cheats.prank(pauser);
        router.pause(PAUSED_BURN_OR_REDISTRIBUTE_SHARES);
        cheats.expectRevert(ISlashEscrowFactoryErrors.OnlyStrategyManager.selector);
        router.initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy);
    }

    function test_initiateSlashEscrow_correctness(uint underlyingAmount) public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        _checkStartBurnOrRedistributions(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount, 1);
    }
}

contract SlashEscrowFactoryUnitTests_releaseSlashEscrow is SlashEscrowFactoryUnitTests {
    /// @dev Asserts that the function reverts if the caller is not the redistribution recipient.
    function testFuzz_releaseSlashEscrow_onlyRedistributionRecipient(uint underlyingAmount) public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        cheats.expectRevert(ISlashEscrowFactoryErrors.OnlyRedistributionRecipient.selector);
        router.releaseSlashEscrow(defaultOperatorSet, defaultSlashId);
    }

    /// @dev Asserts that the function reverts if the redistribution is paused.
    function testFuzz_releaseSlashEscrow_paused(uint underlyingAmount) public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        cheats.prank(pauser);
        router.pauseRedistribution(defaultOperatorSet, defaultSlashId);
        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        router.releaseSlashEscrow(defaultOperatorSet, defaultSlashId);
    }

    /// @dev Asserts that the function reverts if the operator set and slash ID do not exist.
    /// NOTE: `releaseSlashEscrow` does not revert when a slash ID does not exist for an operator set.
    function testFuzz_releaseSlashEscrow_nonexistentSlashIdForOperatorSet(uint underlyingAmount) public {
        cheats.prank(defaultRedistributionRecipient);
        router.releaseSlashEscrow(defaultOperatorSet, defaultSlashId);
        assertFalse(router.isPendingOperatorSet(defaultOperatorSet));
        assertFalse(router.isPendingSlashId(defaultOperatorSet, defaultSlashId));
    }

    /// @dev Tests that multiple strategies can be burned or redistributed in a single call
    function testFuzz_releaseSlashEscrow_multipleStrategies_sameDelay(uint r) public {
        // Initialize arrays to store test data for multiple strategies
        uint numStrategies = bound(r, 2, 10);
        IStrategy[] memory strategies = new IStrategy[](numStrategies);
        MockERC20[] memory tokens = new MockERC20[](numStrategies);
        uint[] memory underlyingAmounts = new uint[](numStrategies);

        // Set up each strategy with random data and start burn/redistribution
        for (uint i = 0; i < numStrategies; i++) {
            // Generate random strategy address and token
            strategies[i] = IStrategy(cheats.randomAddress());
            tokens[i] = new MockERC20();
            underlyingAmounts[i] = cheats.randomUint();

            // Start burn/redistribution for this strategy
            _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i]);
            // Verify the burn/redistribution was started correctly
            _checkStartBurnOrRedistributions(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i], i + 1);
        }

        // Advance time to allow burn/redistribution to occur
        _rollForwardDefaultBurnOrRedistributionDelay();

        // Set up mock calls for each strategy's underlying token
        for (uint i = numStrategies; i > 0; i--) {
            _mockStrategyUnderlyingTokenCall(strategies[i - 1], address(tokens[i - 1]));
        }

        // Execute the burn/redistribution
        _releaseSlashEscrow(defaultOperatorSet, defaultSlashId);

        // Checks

        // Assert that the operator set and slash ID are no longer pending.
        assertFalse(router.isPendingOperatorSet(defaultOperatorSet));
        assertFalse(router.isPendingSlashId(defaultOperatorSet, defaultSlashId));
        assertEq(router.getTotalPendingOperatorSets(), 0);
        assertEq(router.getTotalPendingSlashIds(defaultOperatorSet), 0);

        // Assert that the strategies and underlying amounts are no longer in the pending burn or redistributions.
        assertEq(router.getPendingStrategiesForSlashIdCount(defaultOperatorSet, defaultSlashId), 0);

        // Assert that the underlying amounts are no longer set.
        for (uint i = numStrategies; i > 0; i--) {
            assertEq(_getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategies[i - 1], tokens[i - 1]), 0);
        }

        // Assert that the start block for the (operator set, slash ID) is no longer set.
        assertEq(router.getBurnOrRedistributionStartBlock(defaultOperatorSet, defaultSlashId), 0);
    }

    /// @dev Tests that multiple strategies with different delays are processed correctly
    function testFuzz_releaseSlashEscrow_multipleStrategies_differentDelays(uint r) public {
        // Initialize arrays to store test data for multiple strategies
        uint numStrategies = bound(r, 2, 10);
        IStrategy[] memory strategies = new IStrategy[](numStrategies);
        MockERC20[] memory tokens = new MockERC20[](numStrategies);
        uint[] memory underlyingAmounts = new uint[](numStrategies);
        uint[] memory delays = new uint[](numStrategies);

        // Set up each strategy with random data and different delays
        for (uint i = 0; i < numStrategies; i++) {
            // Generate random strategy address and token
            strategies[i] = IStrategy(cheats.randomAddress());
            tokens[i] = new MockERC20();
            underlyingAmounts[i] = cheats.randomUint();

            // Set different delays for each strategy (increasing delays)
            delays[i] = (i + 1) * 1 days / 12 seconds; // 1 day, 2 days, 3 days, etc.

            // Set the strategy-specific delay
            cheats.prank(defaultOwner);
            router.setStrategyBurnOrRedistributionDelay(strategies[i], delays[i]);

            // Start burn/redistribution for this strategy
            _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i]);
            // Verify the burn/redistribution was started correctly
            _checkStartBurnOrRedistributions(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i], i + 1);
        }

        // Advance time to allow some strategies to be processed (2 days worth of blocks)
        cheats.roll(block.number + 2 days / 12 seconds);

        // Set up mock calls for each strategy's underlying token
        for (uint i = 0; i < numStrategies; i++) {
            _mockStrategyUnderlyingTokenCall(strategies[i], address(tokens[i]));
        }

        // Execute the burn/redistribution
        _releaseSlashEscrow(defaultOperatorSet, defaultSlashId);

        // Verify that only strategies with delays < 2 days were processed
        for (uint i = 0; i < numStrategies; i++) {
            if (delays[i] < 2 days / 12 seconds) {
                // Strategy should have been processed
                assertEq(_getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i]), 0);
            } else {
                // Strategy should still be pending
                assertEq(
                    _getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i]),
                    underlyingAmounts[i]
                );
            }
        }

        // Advance time further to process remaining strategies (4 days worth of blocks)
        cheats.roll(block.number + 2 days / 12 seconds);

        // Execute the burn/redistribution again
        _releaseSlashEscrow(defaultOperatorSet, defaultSlashId);

        // Verify that all strategies have been processed
        assertFalse(router.isPendingOperatorSet(defaultOperatorSet));
        assertFalse(router.isPendingSlashId(defaultOperatorSet, defaultSlashId));
        assertEq(router.getTotalPendingOperatorSets(), 0);
        assertEq(router.getTotalPendingSlashIds(defaultOperatorSet), 0);
        assertEq(router.getPendingStrategiesForSlashIdCount(defaultOperatorSet, defaultSlashId), 0);

        // Verify that all underlying amounts are cleared
        for (uint i = 0; i < numStrategies; i++) {
            assertEq(_getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i]), 0);
        }

        // Verify that the start block is cleared
        assertEq(router.getBurnOrRedistributionStartBlock(defaultOperatorSet, defaultSlashId), 0);
    }
}

contract SlashEscrowFactoryUnitTests_pauseRedistribution is SlashEscrowFactoryUnitTests {
    function test_pauseRedistribution_onlyPauser() public {
        cheats.prank(cheats.randomAddress());
        cheats.expectRevert(IPausable.OnlyPauser.selector);
        router.pauseRedistribution(defaultOperatorSet, defaultSlashId);
    }

    function test_pauseRedistribution_correctness() public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, 100);

        cheats.prank(pauser);
        router.pauseRedistribution(defaultOperatorSet, defaultSlashId);
        assertTrue(router.isBurnOrRedistributionPaused(defaultOperatorSet, defaultSlashId));

        _rollForwardDefaultBurnOrRedistributionDelay();

        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        router.releaseSlashEscrow(defaultOperatorSet, defaultSlashId);
    }
}

contract SlashEscrowFactoryUnitTests_unpauseRedistribution is SlashEscrowFactoryUnitTests {
    function test_unpauseRedistribution_onlyUnpauser() public {
        cheats.prank(cheats.randomAddress());
        cheats.expectRevert(IPausable.OnlyUnpauser.selector);
        router.unpauseRedistribution(defaultOperatorSet, defaultSlashId);
    }

    function test_unpauseRedistribution_correctness() public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, 100);

        cheats.prank(pauser);
        router.pauseRedistribution(defaultOperatorSet, defaultSlashId);
        assertTrue(router.isBurnOrRedistributionPaused(defaultOperatorSet, defaultSlashId));

        _rollForwardDefaultBurnOrRedistributionDelay();

        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        router.releaseSlashEscrow(defaultOperatorSet, defaultSlashId);

        cheats.prank(unpauser);
        router.unpauseRedistribution(defaultOperatorSet, defaultSlashId);
        assertFalse(router.isBurnOrRedistributionPaused(defaultOperatorSet, defaultSlashId));

        // Should not revert...
        _mockStrategyUnderlyingTokenCall(defaultStrategy, address(defaultToken));
        _releaseSlashEscrow(defaultOperatorSet, defaultSlashId);
    }
}

contract SlashEscrowFactoryUnitTests_setStrategyBurnOrRedistributionDelay is SlashEscrowFactoryUnitTests {
    function test_setStrategyBurnOrRedistributionDelay_onlyOwner() public {
        cheats.prank(cheats.randomAddress());
        cheats.expectRevert("Ownable: caller is not the owner");
        router.setStrategyBurnOrRedistributionDelay(defaultStrategy, 100);
    }

    function test_setStrategyBurnOrRedistributionDelay_correctness() public {
        cheats.prank(defaultOwner);
        router.setStrategyBurnOrRedistributionDelay(defaultStrategy, 10 days / 12 seconds);
        // Returns global delay since strategy delay is larger than global delay.
        assertEq(router.getStrategyBurnOrRedistributionDelay(defaultStrategy), 4 days / 12 seconds);

        cheats.prank(defaultOwner);
        router.setStrategyBurnOrRedistributionDelay(defaultStrategy, 1 days / 12 seconds);
        // Returns strategy delay since strategy delay is smaller than global delay.
        assertEq(router.getStrategyBurnOrRedistributionDelay(defaultStrategy), 1 days / 12 seconds);
    }
}

contract SlashEscrowFactoryUnitTests_setGlobalBurnOrRedistributionDelay is SlashEscrowFactoryUnitTests {
    function test_setGlobalBurnOrRedistributionDelay_onlyOwner() public {
        cheats.prank(cheats.randomAddress());
        cheats.expectRevert("Ownable: caller is not the owner");
        router.setGlobalBurnOrRedistributionDelay(100);
    }
}

contract SlashEscrowFactoryUnitTests_getPendingBurnOrRedistributions is SlashEscrowFactoryUnitTests {
    function test_getPendingBurnOrRedistributions_singleSlashId() public {
        // Start burn/redistribution for a single strategy
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, 100);

        // Get pending burn/redistributions for the specific slash ID
        (IStrategy[] memory strategies) = router.getPendingStrategiesForSlashId(defaultOperatorSet, defaultSlashId);

        // Verify results
        assertEq(strategies.length, 1);
        assertEq(address(strategies[0]), address(defaultStrategy));
    }

    function test_getPendingBurnOrRedistributions_multipleStrategies() public {
        // Create multiple strategies and start burn/redistributions
        IStrategy strategy1 = IStrategy(cheats.randomAddress());
        IStrategy strategy2 = IStrategy(cheats.randomAddress());
        MockERC20 token1 = new MockERC20();
        MockERC20 token2 = new MockERC20();

        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategy1, token1, 100);
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategy2, token2, 200);

        // Get pending burn/redistributions for the specific slash ID
        (IStrategy[] memory strategies) = router.getPendingStrategiesForSlashId(defaultOperatorSet, defaultSlashId);

        // Verify results
        assertEq(strategies.length, 2);
        assertEq(address(strategies[0]), address(strategy1));
        assertEq(address(strategies[1]), address(strategy2));
        assertEq(_getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategy1, token1), 100);
        assertEq(_getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategy2, token2), 200);
    }

    function test_getPendingBurnOrRedistributions_multipleSlashIds() public {
        // Create multiple slash IDs and strategies
        uint slashId1 = 1;
        uint slashId2 = 2;
        IStrategy strategy1 = IStrategy(cheats.randomAddress());
        IStrategy strategy2 = IStrategy(cheats.randomAddress());
        MockERC20 token1 = new MockERC20();
        MockERC20 token2 = new MockERC20();

        // Start burn/redistributions for different slash IDs
        _initiateSlashEscrow(defaultOperatorSet, slashId1, strategy1, token1, 100);
        _initiateSlashEscrow(defaultOperatorSet, slashId2, strategy2, token2, 200);

        // Get pending burn/redistributions for all slash IDs of the operator set
        (IStrategy[][] memory strategies) = router.getPendingStrategiesForSlashIds(defaultOperatorSet);

        // Verify results
        assertEq(strategies.length, 2);

        // Check first slash ID
        assertEq(strategies[0].length, 1);
        assertEq(address(strategies[0][0]), address(strategy1));
        assertEq(_getPendingUnderlyingAmountForStrategy(defaultOperatorSet, slashId1, strategy1, token1), 100);

        // Check second slash ID
        assertEq(strategies[1].length, 1);
        assertEq(address(strategies[1][0]), address(strategy2));
        assertEq(_getPendingUnderlyingAmountForStrategy(defaultOperatorSet, slashId2, strategy2, token2), 200);
    }

    function test_getPendingBurnOrRedistributions_multipleOperatorSets() public {
        // Create multiple operator sets and strategies
        OperatorSet memory operatorSet1 = OperatorSet(cheats.randomAddress(), 0);
        OperatorSet memory operatorSet2 = OperatorSet(cheats.randomAddress(), 0);
        IStrategy strategy1 = IStrategy(cheats.randomAddress());
        IStrategy strategy2 = IStrategy(cheats.randomAddress());
        MockERC20 token1 = new MockERC20();
        MockERC20 token2 = new MockERC20();

        // Set redistribution recipients for the operator sets
        allocationManagerMock.setRedistributionRecipient(operatorSet1, defaultRedistributionRecipient);
        allocationManagerMock.setRedistributionRecipient(operatorSet2, defaultRedistributionRecipient);

        // Start burn/redistributions for different operator sets
        _initiateSlashEscrow(operatorSet1, defaultSlashId, strategy1, token1, 100);
        _initiateSlashEscrow(operatorSet2, defaultSlashId, strategy2, token2, 200);

        // Get pending burn/redistributions for all operator sets
        (IStrategy[][][] memory strategies) = router.getPendingStrategiesForSlashIds();

        // Verify results
        assertEq(strategies.length, 2);

        // Check first operator set
        assertEq(strategies[0].length, 1);
        assertEq(strategies[0][0].length, 1);
        assertEq(address(strategies[0][0][0]), address(strategy1));
        assertEq(_getPendingUnderlyingAmountForStrategy(operatorSet1, defaultSlashId, strategy1, token1), 100);

        // Check second operator set
        assertEq(strategies[1].length, 1);
        assertEq(strategies[1][0].length, 1);
        assertEq(address(strategies[1][0][0]), address(strategy2));
        assertEq(_getPendingUnderlyingAmountForStrategy(operatorSet2, defaultSlashId, strategy2, token2), 200);
    }

    function test_getPendingBurnOrRedistributions_empty() public {
        // Test with no pending burn/redistributions
        (IStrategy[] memory strategies) = router.getPendingStrategiesForSlashId(defaultOperatorSet, defaultSlashId);
        assertEq(strategies.length, 0);

        (IStrategy[][] memory strategies2) = router.getPendingStrategiesForSlashIds(defaultOperatorSet);
        assertEq(strategies2.length, 0);

        (IStrategy[][][] memory strategies3) = router.getPendingStrategiesForSlashIds();
        assertEq(strategies3.length, 0);
    }
}
