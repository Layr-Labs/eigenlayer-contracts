// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {MockERC20} from "forge-std/mocks/MockERC20.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/contracts/core/SlashingWithdrawalRouter.sol";

contract SlashingWithdrawalRouterUnitTests is EigenLayerUnitTestSetup, ISlashingWithdrawalRouterEvents {
    /// @notice The pause status for the `burnOrRedistributeShares` function.
    /// @dev Allows all burn or redistribution outflows to be temporarily halted.
    uint8 public constant PAUSED_BURN_OR_REDISTRIBUTE_SHARES = 0;

    SlashingWithdrawalRouter slashingWithdrawalRouter;

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
        slashingWithdrawalRouter = SlashingWithdrawalRouter(
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
        cheats.warp(100);
    }

    function _mockStrategyUnderlyingTokenCall(IStrategy strategy, address underlyingToken) internal {
        cheats.mockCall(address(strategy), abi.encodeWithSelector(IStrategy.underlyingToken.selector), abi.encode(underlyingToken));
    }
}

contract SlashingWithdrawalRouterUnitTests_initialize is SlashingWithdrawalRouterUnitTests {
    function test_initialize() public {
        assertEq(address(slashingWithdrawalRouter.allocationManager()), address(allocationManagerMock));
        assertEq(address(slashingWithdrawalRouter.strategyManager()), address(strategyManagerMock));
        assertEq(address(slashingWithdrawalRouter.pauserRegistry()), address(pauserRegistry));
        assertEq(slashingWithdrawalRouter.paused(), 0);
    }
}

contract SlashingWithdrawalRouterUnitTests_startBurnOrRedistributeShares is SlashingWithdrawalRouterUnitTests {
    function test_startBurnOrRedistributeShares_onlyStrategyManager() public {
        cheats.prank(pauser);
        slashingWithdrawalRouter.pause(PAUSED_BURN_OR_REDISTRIBUTE_SHARES);

        cheats.expectRevert(ISlashingWithdrawalRouterErrors.OnlyStrategyManager.selector);
        slashingWithdrawalRouter.startBurnOrRedistributeShares(defaultOperatorSet, defaultSlashId, defaultStrategy, 0);
    }

    function test_startBurnOrRedistributeShares_onlyRedistributionRecipient(uint underlyingAmount) public {
        allocationManagerMock.setRedistributionRecipient(defaultOperatorSet, defaultRedistributionRecipient);

        cheats.prank(address(strategyManagerMock));
        cheats.expectEmit(true, true, true, true);
        emit StartBurnOrRedistribution(defaultOperatorSet, defaultSlashId, defaultStrategy, underlyingAmount, uint32(block.number));
        slashingWithdrawalRouter.startBurnOrRedistributeShares(defaultOperatorSet, defaultSlashId, defaultStrategy, underlyingAmount);
        deal(address(defaultToken), address(slashingWithdrawalRouter), underlyingAmount);

        cheats.expectRevert(ISlashingWithdrawalRouterErrors.OnlyRedistributionRecipient.selector);
        slashingWithdrawalRouter.burnOrRedistributeShares(defaultOperatorSet, defaultSlashId);
    }

    function test_startBurnOrRedistributeShares_correctness(uint underlyingAmount) public {
        allocationManagerMock.setRedistributionRecipient(defaultOperatorSet, defaultRedistributionRecipient);

        cheats.prank(address(strategyManagerMock));
        cheats.expectEmit(true, true, true, true);
        emit StartBurnOrRedistribution(defaultOperatorSet, defaultSlashId, defaultStrategy, underlyingAmount, uint32(block.number));
        slashingWithdrawalRouter.startBurnOrRedistributeShares(defaultOperatorSet, defaultSlashId, defaultStrategy, underlyingAmount);
        deal(address(defaultToken), address(slashingWithdrawalRouter), underlyingAmount);

        (IStrategy[] memory strategies, uint[] memory amounts) =
            slashingWithdrawalRouter.getPendingBurnOrRedistributions(defaultOperatorSet, defaultSlashId);
        assertEq(strategies.length, 1);
        assertEq(amounts.length, 1);
        assertEq(address(strategies[0]), address(defaultStrategy));
        assertEq(amounts[0], underlyingAmount);

        uint[] memory slashIds = slashingWithdrawalRouter.getPendingSlashIds(defaultOperatorSet);
        assertEq(slashIds.length, 1);
        assertEq(slashIds[0], defaultSlashId);
    }
}

contract SlashingWithdrawalRouterUnitTests_burnOrRedistributeShares is SlashingWithdrawalRouterUnitTests {
    function _startBurnOrRedistributeShares(IStrategy strategy, MockERC20 token, uint underlyingAmount) internal {
        cheats.prank(address(strategyManagerMock));
        cheats.expectEmit(true, true, true, true);
        emit StartBurnOrRedistribution(defaultOperatorSet, defaultSlashId, strategy, underlyingAmount, uint32(block.number));
        slashingWithdrawalRouter.startBurnOrRedistributeShares(defaultOperatorSet, defaultSlashId, strategy, underlyingAmount);
        deal(address(token), address(slashingWithdrawalRouter), underlyingAmount);
    }

    function testFuzz_burnOrRedistributeShares_onlyRedistributionRecipient(uint underlyingAmount) public {
        _startBurnOrRedistributeShares(defaultStrategy, defaultToken, underlyingAmount);
        cheats.expectRevert(ISlashingWithdrawalRouterErrors.OnlyRedistributionRecipient.selector);
        slashingWithdrawalRouter.burnOrRedistributeShares(defaultOperatorSet, defaultSlashId);
    }

    function testFuzz_burnOrRedistributeShares_paused(uint underlyingAmount) public {
        _startBurnOrRedistributeShares(defaultStrategy, defaultToken, underlyingAmount);

        cheats.prank(pauser);
        cheats.expectEmit(true, true, true, true);
        emit RedistributionPaused(defaultOperatorSet, defaultSlashId);
        slashingWithdrawalRouter.pauseRedistribution(defaultOperatorSet, defaultSlashId);

        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        slashingWithdrawalRouter.burnOrRedistributeShares(defaultOperatorSet, defaultSlashId);

        cheats.prank(unpauser);
        cheats.expectEmit(true, true, true, true);
        emit RedistributionUnpaused(defaultOperatorSet, defaultSlashId);
        slashingWithdrawalRouter.unpauseRedistribution(defaultOperatorSet, defaultSlashId);
    }

    function testFuzz_burnOrRedistributeShares_correctness(uint underlyingAmount, uint underlyingAmount2) public {
        // Create a second token and strategy for testing multiple redistributions.
        MockERC20 token2 = new MockERC20();
        IStrategy strategy2 = IStrategy(cheats.randomAddress());

        // Start redistribution process for both strategies.
        _startBurnOrRedistributeShares(defaultStrategy, defaultToken, underlyingAmount);
        _startBurnOrRedistributeShares(strategy2, token2, underlyingAmount2);

        // Check all view functions BEFORE burn/redistribution
        // Check pending slash IDs
        uint[] memory slashIds = slashingWithdrawalRouter.getPendingSlashIds(defaultOperatorSet);
        assertEq(slashIds.length, 1);
        assertEq(slashIds[0], defaultSlashId);

        // Check pending burn/redistributions
        (IStrategy[] memory strategies, uint[] memory amounts) =
            slashingWithdrawalRouter.getPendingBurnOrRedistributions(defaultOperatorSet, defaultSlashId);
        assertEq(strategies.length, 2);
        assertEq(amounts.length, 2);
        assertEq(address(strategies[0]), address(defaultStrategy));
        assertEq(address(strategies[1]), address(strategy2));
        assertEq(amounts[0], underlyingAmount);
        assertEq(amounts[1], underlyingAmount2);

        // Check pending burn/redistributions count
        uint count = slashingWithdrawalRouter.getPendingBurnOrRedistributionsCount(defaultOperatorSet, defaultSlashId);
        assertEq(count, 2);

        // Check pending underlying amounts for each strategy
        uint amount1 = slashingWithdrawalRouter.getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, defaultStrategy);
        uint amount2 = slashingWithdrawalRouter.getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategy2);
        assertEq(amount1, underlyingAmount);
        assertEq(amount2, underlyingAmount2);

        // Check redistribution pause status
        bool isPaused = slashingWithdrawalRouter.isBurnOrRedistributionPaused(defaultOperatorSet, defaultSlashId);
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
        slashingWithdrawalRouter.burnOrRedistributeShares(defaultOperatorSet, defaultSlashId);

        // Verify tokens were correctly transferred to the redistribution recipient.
        assertEq(defaultToken.balanceOf(defaultRedistributionRecipient), underlyingAmount);
        assertEq(token2.balanceOf(defaultRedistributionRecipient), underlyingAmount2);

        // Check all view functions AFTER burn/redistribution
        // Check pending slash IDs are cleared
        slashIds = slashingWithdrawalRouter.getPendingSlashIds(defaultOperatorSet);
        assertEq(slashIds.length, 0);

        // Check pending burn/redistributions are cleared
        (strategies, amounts) = slashingWithdrawalRouter.getPendingBurnOrRedistributions(defaultOperatorSet, defaultSlashId);
        assertEq(strategies.length, 0);
        assertEq(amounts.length, 0);

        // Check pending burn/redistributions count is zero
        count = slashingWithdrawalRouter.getPendingBurnOrRedistributionsCount(defaultOperatorSet, defaultSlashId);
        assertEq(count, 0);

        // Check pending underlying amounts are zero for both strategies
        amount1 = slashingWithdrawalRouter.getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, defaultStrategy);
        amount2 = slashingWithdrawalRouter.getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategy2);
        assertEq(amount1, 0);
        assertEq(amount2, 0);

        // Check redistribution pause status remains unchanged
        isPaused = slashingWithdrawalRouter.isBurnOrRedistributionPaused(defaultOperatorSet, defaultSlashId);
        assertFalse(isPaused);
    }
}
