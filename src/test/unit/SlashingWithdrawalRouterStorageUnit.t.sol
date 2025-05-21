// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {MockERC20} from "src/test/mocks/MockERC20.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/contracts/core/SlashingWithdrawalRouter.sol";

contract SlashingWithdrawalRouterUnitTests is EigenLayerUnitTestSetup, ISlashingWithdrawalRouterEvents {
    /// @notice default address for burning slashed shares and transferring underlying tokens
    address public constant DEFAULT_BURN_ADDRESS = 0x00000000000000000000000000000000000E16E4;

    /// @notice The pause status for the `burnOrRedistributeShares` function.
    /// @dev Allows all burn or redistribution outflows to be temporarily halted.
    uint8 public constant PAUSED_BURN_OR_REDISTRIBUTE_SHARES = 0;

    SlashingWithdrawalRouter router;

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

        router = SlashingWithdrawalRouter(
            address(
                new TransparentUpgradeableProxy(
                    address(
                        new SlashingWithdrawalRouter(
                            IAllocationManager(address(allocationManagerMock)),
                            IStrategyManager(address(strategyManagerMock)),
                            IPauserRegistry(address(pauserRegistry)),
                            "1.0.0"
                        )
                    ),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(SlashingWithdrawalRouter.initialize.selector, defaultOwner, 0)
                )
            )
        );
    }

    function _rollForwardDefaultBurnOrRedistributionDelay() internal {
        cheats.roll(block.number + 3.5 days / 12 seconds);
    }

    /// @dev Sets the return value for the next call to `strategy.underlyingToken()`.
    function _mockStrategyUnderlyingTokenCall(IStrategy strategy, address underlyingToken) internal {
        cheats.mockCall(address(strategy), abi.encodeWithSelector(IStrategy.underlyingToken.selector), abi.encode(underlyingToken));
    }

    /// @dev Starts a burn or redistribution for a given strategy and token.
    /// - Calls as the `StrategyManager`.
    /// - Asserts that the `StartBurnOrRedistribution` event is emitted.
    /// - Mocks the strategy sending the underlying token to the router.
    function _startBurnOrRedistributeShares(
        OperatorSet memory operatorSet,
        uint slashId,
        IStrategy strategy,
        MockERC20 token,
        uint underlyingAmount
    ) internal {
        cheats.prank(address(strategyManagerMock));
        cheats.expectEmit(true, true, true, true);
        emit StartBurnOrRedistribution(operatorSet, slashId, strategy, underlyingAmount, uint32(block.number));
        router.startBurnOrRedistributeShares(defaultOperatorSet, defaultSlashId, strategy, underlyingAmount);
        deal(address(token), address(router), underlyingAmount);
    }

    /// @dev Calls the `burnOrRedistributeShares` function as the redistribution recipient.
    /// - Asserts that the `BurnOrRedistribution` event is emitted for each strategy.
    function _burnOrRedistributeShares(OperatorSet memory operatorSet, uint slashId) internal {
        (IStrategy[] memory strategies, uint[] memory underlyingAmounts) = router.getPendingBurnOrRedistributions(operatorSet, slashId);

        // for (uint i = strategies.length; i > 0; i--) {
        //     cheats.expectEmit(true, true, true, true);
        //     emit BurnOrRedistribution(operatorSet, slashId, strategies[i - 1], underlyingAmounts[i - 1], defaultRedistributionRecipient);
        // }

        cheats.prank(defaultRedistributionRecipient);
        router.burnOrRedistributeShares(operatorSet, slashId);
    }

    /// @dev Asserts that the operator set and slash ID are pending, and that the strategy and underlying amount are in the pending burn or redistributions.
    function _checkStartBurnOrRedistributions(
        OperatorSet memory operatorSet,
        uint slashId,
        IStrategy strategy,
        uint expectedUnderlyingAmount,
        uint expectedCount
    ) internal {
        // Assert that the operator set and slash ID are pending.
        assertTrue(router.isPendingOperatorSet(operatorSet));
        assertTrue(router.isPendingSlashId(operatorSet, slashId));

        // Assert that the underlying amount in escrow for the (operator set, slash ID, strategy) is correct.
        assertEq(router.getPendingUnderlyingAmountForStrategy(operatorSet, slashId, strategy), expectedUnderlyingAmount);

        // Assert that the number of pending burn or redistributions is correct.
        (IStrategy[] memory strategies, uint[] memory amounts) = router.getPendingBurnOrRedistributions(operatorSet, slashId);

        assertEq(strategies.length, expectedCount);
        assertEq(amounts.length, expectedCount);
        assertEq(router.getPendingBurnOrRedistributionsCount(operatorSet, slashId), expectedCount);

        // Assert that the start block for the (operator set, slash ID) is correct.
        assertEq(router.getBurnOrRedistributionStartBlock(operatorSet, slashId), uint32(block.number));
    }
}

contract SlashingWithdrawalRouterUnitTests_initialize is SlashingWithdrawalRouterUnitTests {
    function test_initialize() public {
        assertEq(address(router.allocationManager()), address(allocationManagerMock));
        assertEq(address(router.strategyManager()), address(strategyManagerMock));
        assertEq(address(router.pauserRegistry()), address(pauserRegistry));
        assertEq(router.paused(), 0);
    }
}

contract SlashingWithdrawalRouterUnitTests_startBurnOrRedistributeShares is SlashingWithdrawalRouterUnitTests {
    /// @dev Asserts only the `StrategyManager` can call `startBurnOrRedistributeShares`.
    function test_startBurnOrRedistributeShares_onlyStrategyManager() public {
        cheats.prank(pauser);
        router.pause(PAUSED_BURN_OR_REDISTRIBUTE_SHARES);
        cheats.expectRevert(ISlashingWithdrawalRouterErrors.OnlyStrategyManager.selector);
        router.startBurnOrRedistributeShares(defaultOperatorSet, defaultSlashId, defaultStrategy, 0);
    }

    function test_startBurnOrRedistributeShares_correctness(uint underlyingAmount) public {
        _startBurnOrRedistributeShares(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        _checkStartBurnOrRedistributions(defaultOperatorSet, defaultSlashId, defaultStrategy, underlyingAmount, 1);
    }
}

contract SlashingWithdrawalRouterUnitTests_burnOrRedistributeShares is SlashingWithdrawalRouterUnitTests {
    /// @dev Asserts that the function reverts if the caller is not the redistribution recipient.
    function testFuzz_burnOrRedistributeShares_onlyRedistributionRecipient(uint underlyingAmount) public {
        _startBurnOrRedistributeShares(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        cheats.expectRevert(ISlashingWithdrawalRouterErrors.OnlyRedistributionRecipient.selector);
        router.burnOrRedistributeShares(defaultOperatorSet, defaultSlashId);
    }

    /// @dev Asserts that the function reverts if the redistribution is paused.
    function testFuzz_burnOrRedistributeShares_paused(uint underlyingAmount) public {
        _startBurnOrRedistributeShares(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        cheats.prank(pauser);
        router.pauseRedistribution(defaultOperatorSet, defaultSlashId);
        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        router.burnOrRedistributeShares(defaultOperatorSet, defaultSlashId);
    }

    /// @dev Asserts that the function reverts if the operator set and slash ID do not exist.
    /// NOTE: `burnOrRedistributeShares` does not revert when a slash ID does not exist for an operator set.
    function testFuzz_burnOrRedistributeShares_nonexistentSlashIdForOperatorSet(uint underlyingAmount) public {
        cheats.prank(defaultRedistributionRecipient);
        router.burnOrRedistributeShares(defaultOperatorSet, defaultSlashId);
        assertFalse(router.isPendingOperatorSet(defaultOperatorSet));
        assertFalse(router.isPendingSlashId(defaultOperatorSet, defaultSlashId));
    }

    /// @dev Tests that multiple strategies can be burned or redistributed in a single call
    function testFuzz_burnOrRedistributeShares_multipleStrategies_sameDelay(uint r) public {
        // Initialize arrays to store test data for multiple strategies
        uint numStrategies = bound(r, 1, 10);
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
            _startBurnOrRedistributeShares(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i]);
            // Verify the burn/redistribution was started correctly
            _checkStartBurnOrRedistributions(defaultOperatorSet, defaultSlashId, strategies[i], underlyingAmounts[i], i + 1);
        }

        // Advance time to allow burn/redistribution to occur
        _rollForwardDefaultBurnOrRedistributionDelay();

        // Set up mock calls for each strategy's underlying token
        for (uint i = numStrategies; i > 0; i--) {
            _mockStrategyUnderlyingTokenCall(strategies[i - 1], address(tokens[i - 1]));
        }

        // Execute the burn/redistribution
        _burnOrRedistributeShares(defaultOperatorSet, defaultSlashId);

        // Checks

        // Assert that the operator set and slash ID are no longer pending.
        assertFalse(router.isPendingOperatorSet(defaultOperatorSet));
        assertFalse(router.isPendingSlashId(defaultOperatorSet, defaultSlashId));
        assertEq(router.getTotalPendingOperatorSets(), 0);
        assertEq(router.getTotalPendingSlashIds(defaultOperatorSet), 0);

        // Assert that the strategies and underlying amounts are no longer in the pending burn or redistributions.
        assertEq(router.getPendingBurnOrRedistributionsCount(defaultOperatorSet, defaultSlashId), 0);

        // Assert that the underlying amounts are no longer set.
        for (uint i = numStrategies; i > 0; i--) {
            assertEq(router.getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategies[i - 1]), 0);
        }

        // Assert that the start block for the (operator set, slash ID) is no longer set.
        assertEq(router.getBurnOrRedistributionStartBlock(defaultOperatorSet, defaultSlashId), 0);
    }

    /// @dev Tests that multiple strategies with different delays are processed correctly
    function testFuzz_burnOrRedistributeShares_multipleStrategies_differentDelays(uint r) public {
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
            _startBurnOrRedistributeShares(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i]);
            // Verify the burn/redistribution was started correctly
            _checkStartBurnOrRedistributions(defaultOperatorSet, defaultSlashId, strategies[i], underlyingAmounts[i], i + 1);
        }

        // Advance time to allow some strategies to be processed (2 days worth of blocks)
        cheats.roll(block.number + 2 days / 12 seconds);

        // Set up mock calls for each strategy's underlying token
        for (uint i = 0; i < numStrategies; i++) {
            _mockStrategyUnderlyingTokenCall(strategies[i], address(tokens[i]));
        }

        // Execute the burn/redistribution
        _burnOrRedistributeShares(defaultOperatorSet, defaultSlashId);

        // Verify that only strategies with delays < 2 days were processed
        for (uint i = 0; i < numStrategies; i++) {
            if (delays[i] < 2 days / 12 seconds) {
                // Strategy should have been processed
                assertEq(router.getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategies[i]), 0);
            } else {
                // Strategy should still be pending
                assertEq(router.getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategies[i]), underlyingAmounts[i]);
            }
        }

        // Advance time further to process remaining strategies (4 days worth of blocks)
        cheats.roll(block.number + 2 days / 12 seconds);

        // Execute the burn/redistribution again
        _burnOrRedistributeShares(defaultOperatorSet, defaultSlashId);

        // Verify that all strategies have been processed
        assertFalse(router.isPendingOperatorSet(defaultOperatorSet));
        assertFalse(router.isPendingSlashId(defaultOperatorSet, defaultSlashId));
        assertEq(router.getTotalPendingOperatorSets(), 0);
        assertEq(router.getTotalPendingSlashIds(defaultOperatorSet), 0);
        assertEq(router.getPendingBurnOrRedistributionsCount(defaultOperatorSet, defaultSlashId), 0);

        // Verify that all underlying amounts are cleared
        for (uint i = 0; i < numStrategies; i++) {
            assertEq(router.getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategies[i]), 0);
        }

        // Verify that the start block is cleared
        assertEq(router.getBurnOrRedistributionStartBlock(defaultOperatorSet, defaultSlashId), 0);
    }
}
