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

    function testFuzz_burnOrRedistributeShares_correctness(uint underlyingAmount, uint underlyingAmount2) public {
        // Create a second token and strategy for testing multiple redistributions.
        MockERC20 token2 = new MockERC20();
        IStrategy strategy2 = IStrategy(cheats.randomAddress());

        // Start redistribution process for both strategies.
        _startBurnOrRedistributeShares(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        _checkStartBurnOrRedistributions(defaultOperatorSet, defaultSlashId, defaultStrategy, underlyingAmount, 1);

        _startBurnOrRedistributeShares(defaultOperatorSet, defaultSlashId, strategy2, token2, underlyingAmount2);
        _checkStartBurnOrRedistributions(defaultOperatorSet, defaultSlashId, strategy2, underlyingAmount2, 2);

        // Check redistribution pause status
        bool isPaused = router.isBurnOrRedistributionPaused(defaultOperatorSet, defaultSlashId);
        assertFalse(isPaused);

        // Execute the redistribution as the authorized recipient.
        cheats.prank(defaultRedistributionRecipient);
        _mockStrategyUnderlyingTokenCall(defaultStrategy, address(defaultToken));
        _mockStrategyUnderlyingTokenCall(strategy2, address(token2));

        cheats.roll(block.number + 3.5 days / 12 seconds);

        // Verify correct events are emitted for both strategies.
        cheats.expectEmit(true, true, true, true);
        emit BurnOrRedistribution(defaultOperatorSet, defaultSlashId, strategy2, underlyingAmount2, defaultRedistributionRecipient);
        cheats.expectEmit(true, true, true, true);
        emit BurnOrRedistribution(defaultOperatorSet, defaultSlashId, defaultStrategy, underlyingAmount, defaultRedistributionRecipient);
        router.burnOrRedistributeShares(defaultOperatorSet, defaultSlashId);

        // // Verify tokens were correctly transferred to the redistribution recipient.
        // assertEq(defaultToken.balanceOf(defaultRedistributionRecipient), underlyingAmount);
        // assertEq(token2.balanceOf(defaultRedistributionRecipient), underlyingAmount2);

        // // Check pending burn/redistributions are cleared
        // (strategies, amounts) = router.getPendingBurnOrRedistributions(defaultOperatorSet, defaultSlashId);
        // assertEq(strategies.length, 0);
        // assertEq(amounts.length, 0);

        // // Check pending burn/redistributions count is zero
        // count = router.getPendingBurnOrRedistributionsCount(defaultOperatorSet, defaultSlashId);
        // assertEq(count, 0);

        // // Check pending underlying amounts are zero for both strategies
        // amount1 = router.getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, defaultStrategy);
        // amount2 = router.getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategy2);
        // assertEq(amount1, 0);
        // assertEq(amount2, 0);

        // // Check redistribution pause status remains unchanged
        // isPaused = router.isBurnOrRedistributionPaused(defaultOperatorSet, defaultSlashId);
        // assertFalse(isPaused);
    }
}
