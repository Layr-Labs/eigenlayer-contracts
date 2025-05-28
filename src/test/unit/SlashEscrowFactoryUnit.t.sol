// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {MockERC20} from "src/test/mocks/MockERC20.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/contracts/core/SlashEscrowFactory.sol";
import "src/contracts/core/SlashEscrow.sol";

contract SlashEscrowFactoryUnitTests is EigenLayerUnitTestSetup, ISlashEscrowFactoryEvents {
    /// @notice default address for burning slashed shares and transferring underlying tokens
    address public constant DEFAULT_BURN_ADDRESS = 0x00000000000000000000000000000000000E16E4;

    /// @dev Index for flag that pauses calling `releaseSlashEscrow`
    uint8 constant PAUSED_RELEASE_ESCROW = 0;

    SlashEscrowFactory factory;

    OperatorSet defaultOperatorSet;
    IStrategy defaultStrategy;
    MockERC20 defaultToken;
    uint defaultSlashId;
    address defaultRedistributionRecipient;
    address defaultOwner;
    SlashEscrow slashEscrowImplementation;
    uint32 defaultGlobalDelayBlocks = uint32(4 days / 12 seconds);

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();

        defaultOperatorSet = OperatorSet(cheats.randomAddress(), 0);
        defaultStrategy = IStrategy(cheats.randomAddress());
        defaultToken = new MockERC20();
        defaultSlashId = 1;
        defaultRedistributionRecipient = address(cheats.randomAddress());
        defaultOwner = address(cheats.randomAddress());
        slashEscrowImplementation = new SlashEscrow();
        allocationManagerMock.setRedistributionRecipient(defaultOperatorSet, defaultRedistributionRecipient);

        factory = SlashEscrowFactory(
            address(
                new TransparentUpgradeableProxy(
                    address(
                        new SlashEscrowFactory(
                            IAllocationManager(address(allocationManagerMock)),
                            IStrategyManager(address(strategyManagerMock)),
                            IPauserRegistry(address(pauserRegistry)),
                            ISlashEscrow(address(slashEscrowImplementation)),
                            "1.0.0"
                        )
                    ),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(SlashEscrowFactory.initialize.selector, defaultOwner, 0, uint32(defaultGlobalDelayBlocks))
                )
            )
        );
    }

    function _rollForwardDefaultEscrowDelay() internal {
        cheats.roll(block.number + defaultGlobalDelayBlocks + 1);
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
        return factory.getPendingUnderlyingAmountForStrategy(operatorSet, slashId, strategy);
    }

    /// @dev Starts a escrow for a given strategy and token.
    /// - Calls as the `StrategyManager`.
    /// - Asserts that the `StartEscrow` event is emitted.
    /// - Mocks the strategy sending the underlying token to the `SlashEscrow`.
    function _initiateSlashEscrow(OperatorSet memory operatorSet, uint slashId, IStrategy strategy, MockERC20 token, uint underlyingAmount)
        internal
    {
        cheats.prank(address(strategyManagerMock));
        cheats.expectEmit(true, true, true, true);
        emit StartEscrow(operatorSet, slashId, strategy, uint32(block.number));
        factory.initiateSlashEscrow(operatorSet, slashId, strategy);
        deal(address(token), address(factory.getSlashEscrow(operatorSet, slashId)), underlyingAmount);
    }

    /// @dev Calls the `releaseSlashEscrow` function as the redistribution recipient.
    /// - Asserts that the `Escrow` event is emitted
    function _releaseSlashEscrow(OperatorSet memory operatorSet, uint slashId) internal {
        (IStrategy[] memory strategies) = factory.getPendingStrategiesForSlashId(operatorSet, slashId);

        address redistributionRecipient = allocationManagerMock.getRedistributionRecipient(operatorSet);

        for (uint i = 0; i < strategies.length; i++) {
            cheats.expectEmit(true, true, true, true);
            emit EscrowComplete(operatorSet, slashId, strategies[i], redistributionRecipient);
        }

        // If the redistribution recipient is any address
        if (redistributionRecipient != DEFAULT_BURN_ADDRESS) cheats.prank(defaultRedistributionRecipient);
        else cheats.prank(cheats.randomAddress());
        factory.releaseSlashEscrow(operatorSet, slashId);
    }

    /// @dev Calls the `releaseSlashEscrow` function as the redistribution recipient.
    /// - Asserts that the `Escrow` event is emitted
    function _releaseSlashEscrowByStrategy(OperatorSet memory operatorSet, uint slashId, IStrategy strategy) internal {
        address redistributionRecipient = allocationManagerMock.getRedistributionRecipient(operatorSet);
        // If the redistribution recipient is any address
        if (redistributionRecipient != DEFAULT_BURN_ADDRESS) cheats.prank(redistributionRecipient);
        else cheats.prank(cheats.randomAddress());

        cheats.expectEmit(true, true, true, true);
        emit EscrowComplete(operatorSet, slashId, strategy, redistributionRecipient);

        factory.releaseSlashEscrowByStrategy(operatorSet, slashId, strategy);

        assertEq(factory.computeSlashEscrowSalt(operatorSet, slashId), keccak256(abi.encodePacked(operatorSet.key(), slashId)));
        assertTrue(factory.isDeployedSlashEscrow(operatorSet, slashId));
        ISlashEscrow slashEscrow = factory.getSlashEscrow(operatorSet, slashId);
        assertTrue(factory.isDeployedSlashEscrow(slashEscrow));
        assertTrue(slashEscrow.verifyDeploymentParameters(factory, slashEscrowImplementation, operatorSet, slashId));
    }

    /// @dev Asserts that the operator set and slash ID are pending, and that the strategy and underlying amount are in the pending escrows.
    function _checkStartEscrows(
        OperatorSet memory operatorSet,
        uint slashId,
        IStrategy strategy,
        MockERC20 token,
        uint expectedUnderlyingAmount,
        uint expectedCount
    ) internal {
        // Assert that the operator set and slash ID are pending.
        assertTrue(factory.isPendingOperatorSet(operatorSet));
        assertTrue(factory.isPendingSlashId(operatorSet, slashId));

        OperatorSet[] memory pendingOperatorSets = factory.getPendingOperatorSets();
        uint[] memory pendingSlashIds = factory.getPendingSlashIds(operatorSet);
        assertEq(pendingOperatorSets[pendingOperatorSets.length - 1].key(), operatorSet.key());
        assertEq(pendingSlashIds[pendingSlashIds.length - 1], slashId);

        // Assert that the underlying amount in escrow for the (operator set, slash ID, strategy) is correct.
        assertEq(_getPendingUnderlyingAmountForStrategy(operatorSet, slashId, strategy, token), expectedUnderlyingAmount);

        // Assert that the number of pending escrows is correct.
        (IStrategy[] memory strategies) = factory.getPendingStrategiesForSlashId(operatorSet, slashId);

        assertEq(strategies.length, expectedCount);
        assertEq(factory.getTotalPendingStrategiesForSlashId(operatorSet, slashId), expectedCount);

        // Assert that the start block for the (operator set, slash ID) is correct.
        assertEq(factory.getEscrowStartBlock(operatorSet, slashId), uint32(block.number));

        // Assert that the escrow is deployed
        assertEq(factory.computeSlashEscrowSalt(operatorSet, slashId), keccak256(abi.encodePacked(operatorSet.key(), slashId)));
        assertTrue(factory.isDeployedSlashEscrow(operatorSet, slashId));
        ISlashEscrow slashEscrow = factory.getSlashEscrow(operatorSet, slashId);
        assertTrue(factory.isDeployedSlashEscrow(slashEscrow));
        assertTrue(slashEscrow.verifyDeploymentParameters(factory, slashEscrowImplementation, operatorSet, slashId));
    }

    /// @dev Sets the redistribution recipient to the default burn address if `shouldBurn` is true.
    function _setRedistributionRecipient(bool shouldBurn) internal {
        if (shouldBurn) allocationManagerMock.setRedistributionRecipient(defaultOperatorSet, DEFAULT_BURN_ADDRESS);
    }
}

contract SlashEscrowFactoryUnitTests_initialize is SlashEscrowFactoryUnitTests {
    function test_initialize() public {
        assertEq(address(factory.allocationManager()), address(allocationManagerMock));
        assertEq(address(factory.strategyManager()), address(strategyManagerMock));
        assertEq(address(factory.pauserRegistry()), address(pauserRegistry));
        assertEq(address(factory.slashEscrowImplementation()), address(slashEscrowImplementation));
        assertEq(factory.getGlobalEscrowDelay(), defaultGlobalDelayBlocks);
        assertEq(factory.paused(), 0);
    }
}

contract SlashEscrowFactoryUnitTests_initiateSlashEscrow is SlashEscrowFactoryUnitTests {
    /// @dev Asserts only the `StrategyManager` can call `initiateSlashEscrow`.
    function test_initiateSlashEscrow_onlyStrategyManager() public {
        cheats.expectRevert(ISlashEscrowFactoryErrors.OnlyStrategyManager.selector);
        factory.initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy);
    }

    function testFuzz_initiateSlashEscrow_multipleStrategies(uint r) public {
        // Initialize arrays to store test data for multiple strategies
        uint numStrategies = bound(r, 2, 10);
        // Set up each strategy with random data and start escrow
        for (uint i = 0; i < numStrategies; i++) {
            // Generate random strategy address and token
            IStrategy strategy = IStrategy(cheats.randomAddress());
            MockERC20 token = new MockERC20();
            uint underlyingAmount = cheats.randomUint();

            // Start escrow for this strategy
            _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategy, token, underlyingAmount);
            // Verify the escrow was started correctly
            _checkStartEscrows(defaultOperatorSet, defaultSlashId, strategy, token, underlyingAmount, i + 1);
        }
    }

    function test_initiateSlashEscrow_correctness(uint underlyingAmount) public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        _checkStartEscrows(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount, 1);
    }
}

contract SlashEscrowFactoryUnitTests_releaseSlashEscrow is SlashEscrowFactoryUnitTests {
    /// @dev Asserts that the function reverts if the caller is not the redistribution recipient.
    function testFuzz_releaseSlashEscrow_OnlyRedistributionRecipient(uint underlyingAmount) public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        cheats.expectRevert(ISlashEscrowFactoryErrors.OnlyRedistributionRecipient.selector);
        factory.releaseSlashEscrow(defaultOperatorSet, defaultSlashId);
    }

    function testFuzz_releaseSlashEscrow_globalPause(uint underlyingAmount) public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        cheats.prank(pauser);
        factory.pause(2 ** PAUSED_RELEASE_ESCROW);
        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        factory.releaseSlashEscrow(defaultOperatorSet, defaultSlashId);
    }

    /// @dev Asserts that the function reverts if the redistribution is paused.
    function testFuzz_releaseSlashEscrow_paused(uint underlyingAmount) public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        cheats.prank(pauser);
        factory.pauseEscrow(defaultOperatorSet, defaultSlashId);
        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        factory.releaseSlashEscrow(defaultOperatorSet, defaultSlashId);
    }

    /// @dev Asserts that the function reverts if the operator set and slash ID do not exist.
    /// NOTE: `releaseSlashEscrow` does not revert when a slash ID does not exist for an operator set.
    function testFuzz_releaseSlashEscrow_nonexistentSlashIdForOperatorSet(uint underlyingAmount) public {
        cheats.prank(defaultRedistributionRecipient);
        factory.releaseSlashEscrow(defaultOperatorSet, defaultSlashId);
        assertFalse(factory.isPendingOperatorSet(defaultOperatorSet));
        assertFalse(factory.isPendingSlashId(defaultOperatorSet, defaultSlashId));
    }

    function testFuzz_releaseSlashEscrow_delayNotElapsed(uint underlyingAmount) public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        cheats.roll(block.number + defaultGlobalDelayBlocks);
        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(ISlashEscrowFactoryErrors.EscrowDelayNotElapsed.selector);
        factory.releaseSlashEscrow(defaultOperatorSet, defaultSlashId);
    }

    /// @dev Tests that multiple strategies can be burned or redistributed in a single call
    function testFuzz_releaseSlashEscrow_multipleStrategies_sameDelay(uint r) public {
        // Initialize arrays to store test data for multiple strategies
        uint numStrategies = bound(r, 2, 10);
        IStrategy[] memory strategies = new IStrategy[](numStrategies);
        MockERC20[] memory tokens = new MockERC20[](numStrategies);
        uint[] memory underlyingAmounts = new uint[](numStrategies);

        // // Randomly update the redistribution to be the default burn address
        // _setRedistributionRecipient(r % 2 == 0);

        // Set up each strategy with random data and start escrow
        for (uint i = 0; i < numStrategies; i++) {
            // Generate random strategy address and token
            strategies[i] = IStrategy(cheats.randomAddress());
            tokens[i] = new MockERC20();
            underlyingAmounts[i] = cheats.randomUint();

            // Start escrow for this strategy
            _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i]);
            // Verify the escrow was started correctly
            _checkStartEscrows(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i], i + 1);
        }

        // Advance time to allow escrow to occur
        _rollForwardDefaultEscrowDelay();

        // Set up mock calls for each strategy's underlying token
        for (uint i = 0; i < numStrategies; i++) {
            _mockStrategyUnderlyingTokenCall(strategies[i], address(tokens[i]));
        }

        // Execute the escrow
        _releaseSlashEscrow(defaultOperatorSet, defaultSlashId);

        // Checks

        // Assert that the operator set and slash ID are no longer pending.
        assertFalse(factory.isPendingOperatorSet(defaultOperatorSet));
        assertFalse(factory.isPendingSlashId(defaultOperatorSet, defaultSlashId));
        assertEq(factory.getTotalPendingOperatorSets(), 0);
        assertEq(factory.getTotalPendingSlashIds(defaultOperatorSet), 0);

        // Assert that the strategies and underlying amounts are no longer in the pending escrows.
        assertEq(factory.getTotalPendingStrategiesForSlashId(defaultOperatorSet, defaultSlashId), 0);

        // Assert that the underlying amounts are no longer set.
        for (uint i = 0; i < numStrategies; i++) {
            assertEq(_getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i]), 0);
        }

        // Assert that the start block for the (operator set, slash ID) is no longer set.
        assertEq(factory.getEscrowStartBlock(defaultOperatorSet, defaultSlashId), 0);
    }

    /// @dev Tests that multiple strategies with different delays are processed correctly
    function testFuzz_releaseSlashEscrow_multipleStrategies_differentDelays(uint r) public {
        // Initialize arrays to store test data for multiple strategies
        uint numStrategies = bound(r, 2, 10);
        IStrategy[] memory strategies = new IStrategy[](numStrategies);
        MockERC20[] memory tokens = new MockERC20[](numStrategies);
        uint[] memory underlyingAmounts = new uint[](numStrategies);
        uint32[] memory delays = new uint32[](numStrategies);

        // Randomly update the redistribution to be the default burn address
        _setRedistributionRecipient(r % 2 == 0);

        // Set global delay less than all the strategy delays.
        cheats.prank(defaultOwner);
        factory.setGlobalEscrowDelay(0.5 days / 12 seconds);

        // Set up each strategy with random data and different delays
        for (uint i = 0; i < numStrategies; i++) {
            // Generate random strategy address and token
            strategies[i] = IStrategy(cheats.randomAddress());
            tokens[i] = new MockERC20();
            underlyingAmounts[i] = cheats.randomUint();

            // Set different delays for each strategy (increasing delays)
            delays[i] = uint32((i + 1) * 1 days / 12 seconds); // 1 day, 2 days, 3 days, etc.

            // Set the strategy-specific delay
            cheats.prank(defaultOwner);
            factory.setStrategyEscrowDelay(strategies[i], delays[i]);

            // Start escrow for this strategy
            _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i]);
            // Verify the escrow was started correctly
            _checkStartEscrows(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i], i + 1);
        }

        // Assert that we cannot release the slash escrow until the delay has elapsed for ALL strategies
        cheats.roll(block.number + 1 days / 12 seconds * numStrategies);
        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(ISlashEscrowFactoryErrors.EscrowDelayNotElapsed.selector);
        factory.releaseSlashEscrow(defaultOperatorSet, defaultSlashId);

        // Roll forward to the next block
        cheats.roll(block.number + 1);

        // Set up mock calls for each strategy's underlying token
        for (uint i = 0; i < numStrategies; i++) {
            _mockStrategyUnderlyingTokenCall(strategies[i], address(tokens[i]));
        }

        // Execute the escrow
        _releaseSlashEscrow(defaultOperatorSet, defaultSlashId);

        // Verify that all strategies have been processed
        assertFalse(factory.isPendingOperatorSet(defaultOperatorSet));
        assertFalse(factory.isPendingSlashId(defaultOperatorSet, defaultSlashId));
        assertEq(factory.getTotalPendingOperatorSets(), 0);
        assertEq(factory.getTotalPendingSlashIds(defaultOperatorSet), 0);
        assertEq(factory.getTotalPendingStrategiesForSlashId(defaultOperatorSet, defaultSlashId), 0);

        // Verify that all underlying amounts are cleared
        for (uint i = 0; i < numStrategies; i++) {
            assertEq(_getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i]), 0);
        }

        // Verify that the start block is cleared
        assertEq(factory.getEscrowStartBlock(defaultOperatorSet, defaultSlashId), 0);
    }

    /// @dev Tests that operatorSets are only cleared once all slash IDs are released
    function testFuzz_releaseSlashEscrow_multipleReleases(uint r) public {
        uint numEscrows = bound(r, 2, 5);

        IStrategy[] memory strategies = new IStrategy[](1);
        MockERC20[] memory tokens = new MockERC20[](1);
        uint[] memory underlyingAmounts = new uint[](1);

        underlyingAmounts[0] = cheats.randomUint() / numEscrows;

        strategies[0] = IStrategy(cheats.randomAddress());
        tokens[0] = new MockERC20();

        // Set up numEscrows slash escrows for the same operator set
        for (uint i = 0; i < numEscrows; i++) {
            // Start escrow for this slash
            _initiateSlashEscrow(defaultOperatorSet, defaultSlashId + i, strategies[0], tokens[0], underlyingAmounts[0]);
            // Verify the escrow was started correctly
            _checkStartEscrows(defaultOperatorSet, defaultSlashId + i, strategies[0], tokens[0], underlyingAmounts[0], 1);
        }

        _rollForwardDefaultEscrowDelay();

        // Release the first n-1 slash escrows
        for (uint i = 0; i < numEscrows - 1; i++) {
            _releaseSlashEscrow(defaultOperatorSet, defaultSlashId + i);
        }

        // Assert that the operator set is still pending
        assertTrue(factory.isPendingOperatorSet(defaultOperatorSet));
        assertTrue(factory.isPendingSlashId(defaultOperatorSet, defaultSlashId + numEscrows - 1));
        assertEq(factory.getTotalPendingOperatorSets(), 1);
        assertEq(factory.getTotalPendingSlashIds(defaultOperatorSet), 1);

        // Release the last escrow
        _releaseSlashEscrow(defaultOperatorSet, defaultSlashId + numEscrows - 1);

        // Assert that the operator set is no longer pending
        assertFalse(factory.isPendingOperatorSet(defaultOperatorSet));
        assertFalse(factory.isPendingSlashId(defaultOperatorSet, defaultSlashId + numEscrows - 1));
        assertEq(factory.getTotalPendingOperatorSets(), 0);
    }
}

contract SlashEscrowFactoryUnitTests_releaseSlashEscrowByStrategy is SlashEscrowFactoryUnitTests {
    /// @dev Asserts that the function reverts if the caller is not the redistribution recipient.
    function testFuzz_releaseSlashEscrowByStrategy_OnlyRedistributionRecipient(uint underlyingAmount) public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        cheats.expectRevert(ISlashEscrowFactoryErrors.OnlyRedistributionRecipient.selector);
        factory.releaseSlashEscrowByStrategy(defaultOperatorSet, defaultSlashId, defaultStrategy);
    }

    function testFuzz_releaseSlashEscrowByStrategy_globalPause(uint underlyingAmount) public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        cheats.prank(pauser);
        factory.pause(2 ** PAUSED_RELEASE_ESCROW);
        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        factory.releaseSlashEscrowByStrategy(defaultOperatorSet, defaultSlashId, defaultStrategy);
    }

    /// @dev Asserts that the function reverts if the redistribution is paused.
    function testFuzz_releaseSlashEscrowByStrategy_paused(uint underlyingAmount) public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        cheats.prank(pauser);
        factory.pauseEscrow(defaultOperatorSet, defaultSlashId);
        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        factory.releaseSlashEscrowByStrategy(defaultOperatorSet, defaultSlashId, defaultStrategy);
    }

    /// @dev Asserts that the function reverts if the operator set and slash ID do not exist.
    /// NOTE: `releaseSlashEscrowByStrategy` DOES revert when a slash ID does not exist for an operator set.
    function testFuzz_releaseSlashEscrowByStrategy_nonexistentSlashIdForOperatorSet(uint underlyingAmount) public {
        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert();
        factory.releaseSlashEscrowByStrategy(defaultOperatorSet, defaultSlashId, defaultStrategy);
        assertFalse(factory.isPendingOperatorSet(defaultOperatorSet));
        assertFalse(factory.isPendingSlashId(defaultOperatorSet, defaultSlashId));
    }

    function testFuzz_releaseSlashEscrowByStrategy_delayNotElapsed(uint underlyingAmount) public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        cheats.roll(block.number + defaultGlobalDelayBlocks);
        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(ISlashEscrowFactoryErrors.EscrowDelayNotElapsed.selector);
        factory.releaseSlashEscrowByStrategy(defaultOperatorSet, defaultSlashId, defaultStrategy);
    }

    function testFuzz_releaseSlashEscrowByStrategy_multipleStrategies(uint underlyingAmount) public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, underlyingAmount);
        cheats.roll(block.number + defaultGlobalDelayBlocks);
        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(ISlashEscrowFactoryErrors.EscrowDelayNotElapsed.selector);
        factory.releaseSlashEscrowByStrategy(defaultOperatorSet, defaultSlashId, defaultStrategy);
    }

    /// @dev Tests that multiple strategies can be burned or redistributed across multiple calls
    function testFuzz_releaseSlashEscrowByStrategy__multipleStrategies_sameDelay(uint r) public {
        // Initialize arrays to store test data for multiple strategies
        uint numStrategies = bound(r, 2, 10);
        IStrategy[] memory strategies = new IStrategy[](numStrategies);
        MockERC20[] memory tokens = new MockERC20[](numStrategies);
        uint[] memory underlyingAmounts = new uint[](numStrategies);

        // Randomly update the redistribution to be the default burn address
        _setRedistributionRecipient(r % 2 == 0);

        // Set up each strategy with random data and start escrow
        for (uint i = 0; i < numStrategies; i++) {
            // Generate random strategy address and token
            strategies[i] = IStrategy(cheats.randomAddress());
            tokens[i] = new MockERC20();
            underlyingAmounts[i] = cheats.randomUint();

            // Start escrow for this strategy
            _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i]);
            // Verify the escrow was started correctly
            _checkStartEscrows(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i], i + 1);
        }

        // Advance time to allow escrow to occur
        _rollForwardDefaultEscrowDelay();

        // Set up mock calls for each strategy's underlying token
        for (uint i = 0; i < numStrategies; i++) {
            _mockStrategyUnderlyingTokenCall(strategies[i], address(tokens[i]));
        }

        // Release the first n-1 slash escrows
        for (uint i = 0; i < numStrategies - 1; i++) {
            _releaseSlashEscrowByStrategy(defaultOperatorSet, defaultSlashId, strategies[i]);
        }

        // Assert that the slashId and operatorSet are still pending
        assertTrue(factory.isPendingOperatorSet(defaultOperatorSet));
        assertTrue(factory.isPendingSlashId(defaultOperatorSet, defaultSlashId));
        assertEq(factory.getTotalPendingOperatorSets(), 1);
        assertEq(factory.getTotalPendingSlashIds(defaultOperatorSet), 1);
        assertEq(factory.getTotalPendingStrategiesForSlashId(defaultOperatorSet, defaultSlashId), 1);

        // Release the last slash escrow
        _releaseSlashEscrowByStrategy(defaultOperatorSet, defaultSlashId, strategies[numStrategies - 1]);

        // Checks

        // Assert that the operator set and slash ID are no longer pending.
        assertFalse(factory.isPendingOperatorSet(defaultOperatorSet));
        assertFalse(factory.isPendingSlashId(defaultOperatorSet, defaultSlashId));
        assertEq(factory.getTotalPendingOperatorSets(), 0);
        assertEq(factory.getTotalPendingSlashIds(defaultOperatorSet), 0);

        // Assert that the strategies and underlying amounts are no longer in the pending escrows.
        assertEq(factory.getTotalPendingStrategiesForSlashId(defaultOperatorSet, defaultSlashId), 0);

        // Assert that the underlying amounts are no longer set.
        for (uint i = 0; i < numStrategies; i++) {
            assertEq(_getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i]), 0);
        }

        // Assert that the start block for the (operator set, slash ID) is no longer set.
        assertEq(factory.getEscrowStartBlock(defaultOperatorSet, defaultSlashId), 0);
    }

    /// @dev Tests that multiple strategies can be burned or redistributed across multiple calls
    function testFuzz_releaseSlashEscrowByStrategy__multipleStrategies_byIndexThenAll(uint r) public {
        // Initialize arrays to store test data for multiple strategies
        uint numStrategies = bound(r, 2, 10);
        IStrategy[] memory strategies = new IStrategy[](numStrategies);
        MockERC20[] memory tokens = new MockERC20[](numStrategies);
        uint[] memory underlyingAmounts = new uint[](numStrategies);

        // Randomly update the redistribution to be the default burn address
        _setRedistributionRecipient(r % 2 == 0);

        // Set up each strategy with random data and start escrow
        for (uint i = 0; i < numStrategies; i++) {
            // Generate random strategy address and token
            strategies[i] = IStrategy(cheats.randomAddress());
            tokens[i] = new MockERC20();
            underlyingAmounts[i] = cheats.randomUint();

            // Start escrow for this strategy
            _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i]);
            // Verify the escrow was started correctly
            _checkStartEscrows(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i], i + 1);
        }

        // Advance time to allow escrow to occur
        _rollForwardDefaultEscrowDelay();

        // Release the first index
        _releaseSlashEscrowByStrategy(defaultOperatorSet, defaultSlashId, strategies[0]);

        // Assert that the slashId and operatorSet are still pending
        assertTrue(factory.isPendingOperatorSet(defaultOperatorSet));
        assertTrue(factory.isPendingSlashId(defaultOperatorSet, defaultSlashId));
        assertEq(factory.getTotalPendingOperatorSets(), 1);
        assertEq(factory.getTotalPendingSlashIds(defaultOperatorSet), 1);
        assertEq(factory.getTotalPendingStrategiesForSlashId(defaultOperatorSet, defaultSlashId), numStrategies - 1);

        // Release remaining
        _releaseSlashEscrow(defaultOperatorSet, defaultSlashId);

        // Checks

        // Assert that the operator set and slash ID are no longer pending.
        assertFalse(factory.isPendingOperatorSet(defaultOperatorSet));
        assertFalse(factory.isPendingSlashId(defaultOperatorSet, defaultSlashId));
        assertEq(factory.getTotalPendingOperatorSets(), 0);
        assertEq(factory.getTotalPendingSlashIds(defaultOperatorSet), 0);

        // Assert that the strategies and underlying amounts are no longer in the pending escrows.
        assertEq(factory.getTotalPendingStrategiesForSlashId(defaultOperatorSet, defaultSlashId), 0);

        // Assert that the underlying amounts are no longer set.
        for (uint i = 0; i < numStrategies; i++) {
            assertEq(_getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i]), 0);
        }

        // Assert that the start block for the (operator set, slash ID) is no longer set.
        assertEq(factory.getEscrowStartBlock(defaultOperatorSet, defaultSlashId), 0);
    }
}

contract SlashEscrowFactoryUnitTests_pauseEscrow is SlashEscrowFactoryUnitTests {
    function test_pauseEscrow_onlyPauser() public {
        cheats.prank(cheats.randomAddress());
        cheats.expectRevert(IPausable.OnlyPauser.selector);
        factory.pauseEscrow(defaultOperatorSet, defaultSlashId);
    }

    function test_pauseEscrow_correctness() public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, 100);

        cheats.prank(pauser);
        factory.pauseEscrow(defaultOperatorSet, defaultSlashId);
        assertTrue(factory.isEscrowPaused(defaultOperatorSet, defaultSlashId));

        _rollForwardDefaultEscrowDelay();

        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        factory.releaseSlashEscrow(defaultOperatorSet, defaultSlashId);
    }
}

contract SlashEscrowFactoryUnitTests_unpauseEscrow is SlashEscrowFactoryUnitTests {
    function test_unpauseEscrow_onlyUnpauser() public {
        cheats.prank(cheats.randomAddress());
        cheats.expectRevert(IPausable.OnlyUnpauser.selector);
        factory.unpauseEscrow(defaultOperatorSet, defaultSlashId);
    }

    function test_unpauseEscrow_correctness() public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, 100);

        cheats.prank(pauser);
        factory.pauseEscrow(defaultOperatorSet, defaultSlashId);
        assertTrue(factory.isEscrowPaused(defaultOperatorSet, defaultSlashId));

        _rollForwardDefaultEscrowDelay();

        cheats.prank(defaultRedistributionRecipient);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        factory.releaseSlashEscrow(defaultOperatorSet, defaultSlashId);

        cheats.prank(unpauser);
        factory.unpauseEscrow(defaultOperatorSet, defaultSlashId);
        assertFalse(factory.isEscrowPaused(defaultOperatorSet, defaultSlashId));

        // Should not revert...
        _mockStrategyUnderlyingTokenCall(defaultStrategy, address(defaultToken));
        _releaseSlashEscrow(defaultOperatorSet, defaultSlashId);
    }
}

contract SlashEscrowFactoryUnitTests_setStrategyEscrowDelay is SlashEscrowFactoryUnitTests {
    function test_setStrategyEscrowDelay_onlyOwner() public {
        cheats.prank(cheats.randomAddress());
        cheats.expectRevert("Ownable: caller is not the owner");
        factory.setStrategyEscrowDelay(defaultStrategy, uint32(25));
    }

    function testFuzz_setStrategyEscrowDelay_correctness(uint32 delay) public {
        delay = uint32(bound(delay, defaultGlobalDelayBlocks + 1, type(uint).max));
        cheats.prank(defaultOwner);
        cheats.expectEmit(true, true, true, true);
        emit StrategyEscrowDelaySet(defaultStrategy, delay);
        factory.setStrategyEscrowDelay(defaultStrategy, delay);
        // Returns strategy delay since strategy delay is larger than global delay.
        assertEq(factory.getStrategyEscrowDelay(defaultStrategy), delay);
    }
}

contract SlashEscrowFactoryUnitTests_getBurnOrRedistributionDelay is SlashEscrowFactoryUnitTests {
    function testFuzz_getBurnOrRedistributionDelay_correctness(uint r) public {
        // Initialize arrays to store test data for multiple strategies
        uint numStrategies = bound(r, 2, 10);
        IStrategy[] memory strategies = new IStrategy[](numStrategies);
        MockERC20[] memory tokens = new MockERC20[](numStrategies);
        uint[] memory underlyingAmounts = new uint[](numStrategies);
        uint32[] memory delays = new uint32[](numStrategies);

        // Set global delay less than all the strategy delays.
        cheats.prank(defaultOwner);
        factory.setGlobalEscrowDelay(0.5 days / 12 seconds);

        // Set up each strategy with random data and different delays
        for (uint i = 0; i < numStrategies; i++) {
            // Generate random strategy address and token
            strategies[i] = IStrategy(cheats.randomAddress());
            tokens[i] = new MockERC20();
            underlyingAmounts[i] = cheats.randomUint();

            // Set different delays for each strategy (increasing delays)
            delays[i] = uint32((i + 1) * 1 days / 12 seconds); // 1 day, 2 days, 3 days, etc.

            // Set the strategy-specific delay
            cheats.prank(defaultOwner);
            factory.setStrategyEscrowDelay(strategies[i], delays[i]);

            // Start escrow for this strategy
            _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i]);
            // Verify the escrow was started correctly
            _checkStartEscrows(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i], i + 1);
        }

        // The complete block should be the maximum delay across all strategies
        assertEq(factory.getEscrowCompleteBlock(defaultOperatorSet, defaultSlashId), block.number + delays[numStrategies - 1] + 1);
    }
}

contract SlashEscrowFactoryUnitTests_getEscrowDelay is SlashEscrowFactoryUnitTests {
    function testFuzz_getEscrowDelay_correctness(uint r) public {
        // Initialize arrays to store test data for multiple strategies
        uint numStrategies = bound(r, 2, 10);
        IStrategy[] memory strategies = new IStrategy[](numStrategies);
        MockERC20[] memory tokens = new MockERC20[](numStrategies);
        uint[] memory underlyingAmounts = new uint[](numStrategies);
        uint32[] memory delays = new uint32[](numStrategies);

        // Set global delay less than all the strategy delays.
        cheats.prank(defaultOwner);
        factory.setGlobalEscrowDelay(0.5 days / 12 seconds);

        // Set up each strategy with random data and different delays
        for (uint i = 0; i < numStrategies; i++) {
            // Generate random strategy address and token
            strategies[i] = IStrategy(cheats.randomAddress());
            tokens[i] = new MockERC20();
            underlyingAmounts[i] = cheats.randomUint();

            // Set different delays for each strategy (increasing delays)
            delays[i] = uint32((i + 1) * 1 days / 12 seconds); // 1 day, 2 days, 3 days, etc.

            // Set the strategy-specific delay
            cheats.prank(defaultOwner);
            factory.setStrategyEscrowDelay(strategies[i], delays[i]);

            // Start escrow for this strategy
            _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i]);
            // Verify the escrow was started correctly
            _checkStartEscrows(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i], i + 1);
        }

        // The complete block should be the maximum delay across all strategies
        assertEq(factory.getEscrowCompleteBlock(defaultOperatorSet, defaultSlashId), block.number + delays[numStrategies - 1] + 1);
    }
}

contract SlashEscrowFactoryUnitTests_setGlobalEscrowDelay is SlashEscrowFactoryUnitTests {
    function test_setGlobalEscrowDelay_onlyOwner() public {
        cheats.prank(cheats.randomAddress());
        cheats.expectRevert("Ownable: caller is not the owner");
        factory.setGlobalEscrowDelay(100);
    }

    function testFuzz_setGlobalEscrowDelay_correctness(uint32 delay) public {
        cheats.prank(defaultOwner);
        cheats.expectEmit(true, true, true, true);
        emit GlobalEscrowDelaySet(delay);
        factory.setGlobalEscrowDelay(delay);
        assertEq(factory.getGlobalEscrowDelay(), delay);
    }
}

contract SlashEscrowFactoryUnitTests_getPendingEscrows is SlashEscrowFactoryUnitTests {
    function test_getPendingEscrows_singleSlashId() public {
        // Start escrow for a single strategy
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, 100);

        // Get pending escrows for the specific slash ID
        (IStrategy[] memory strategies) = factory.getPendingStrategiesForSlashId(defaultOperatorSet, defaultSlashId);

        // Verify results
        assertEq(strategies.length, 1);
        assertEq(address(strategies[0]), address(defaultStrategy));
    }

    function test_getPendingEscrows_multipleStrategies() public {
        // Create multiple strategies and start escrows
        IStrategy strategy1 = IStrategy(cheats.randomAddress());
        IStrategy strategy2 = IStrategy(cheats.randomAddress());
        MockERC20 token1 = new MockERC20();
        MockERC20 token2 = new MockERC20();

        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategy1, token1, 100);
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategy2, token2, 200);

        // Get pending escrows for the specific slash ID
        (IStrategy[] memory strategies) = factory.getPendingStrategiesForSlashId(defaultOperatorSet, defaultSlashId);

        // Verify results
        assertEq(strategies.length, 2);
        assertEq(address(strategies[0]), address(strategy1));
        assertEq(address(strategies[1]), address(strategy2));
        assertEq(_getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategy1, token1), 100);
        assertEq(_getPendingUnderlyingAmountForStrategy(defaultOperatorSet, defaultSlashId, strategy2, token2), 200);
    }

    function test_getPendingEscrows_multipleSlashIds() public {
        // Create multiple slash IDs and strategies
        uint slashId1 = 1;
        uint slashId2 = 2;
        IStrategy strategy1 = IStrategy(cheats.randomAddress());
        IStrategy strategy2 = IStrategy(cheats.randomAddress());
        MockERC20 token1 = new MockERC20();
        MockERC20 token2 = new MockERC20();

        // Start escrows for different slash IDs
        _initiateSlashEscrow(defaultOperatorSet, slashId1, strategy1, token1, 100);
        _initiateSlashEscrow(defaultOperatorSet, slashId2, strategy2, token2, 200);

        // Get pending escrows for all slash IDs of the operator set
        (IStrategy[][] memory strategies) = factory.getPendingStrategiesForSlashIds(defaultOperatorSet);

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

    function test_getPendingEscrows_empty() public {
        // Test with no pending escrows
        (IStrategy[] memory strategies) = factory.getPendingStrategiesForSlashId(defaultOperatorSet, defaultSlashId);
        assertEq(strategies.length, 0);

        (IStrategy[][] memory strategies2) = factory.getPendingStrategiesForSlashIds(defaultOperatorSet);
        assertEq(strategies2.length, 0);
        assertEq(factory.getTotalPendingStrategiesForSlashId(defaultOperatorSet, defaultSlashId), 0);
    }
}

contract SlashEscrowFactoryUnitTests_getEscrowCompleteBlock is SlashEscrowFactoryUnitTests {
    function test_getEscrowCompleteBlock() public {
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, 100);
        assertEq(factory.getEscrowCompleteBlock(defaultOperatorSet, defaultSlashId), block.number + defaultGlobalDelayBlocks + 1);
    }

    function testFuzz_getEscrowCompleteBlock_multipleStrategies(uint r) public {
        // Initialize arrays to store test data for multiple strategies
        uint numStrategies = bound(r, 2, 10);
        IStrategy[] memory strategies = new IStrategy[](numStrategies);
        MockERC20[] memory tokens = new MockERC20[](numStrategies);
        uint[] memory underlyingAmounts = new uint[](numStrategies);
        uint32[] memory delays = new uint32[](numStrategies);

        // Set global delay less than all the strategy delays.
        cheats.prank(defaultOwner);
        factory.setGlobalEscrowDelay(0.5 days / 12 seconds);

        // Set up each strategy with random data and different delays
        for (uint i = 0; i < numStrategies; i++) {
            // Generate random strategy address and token
            strategies[i] = IStrategy(cheats.randomAddress());
            tokens[i] = new MockERC20();
            underlyingAmounts[i] = cheats.randomUint();

            // Set different delays for each strategy (increasing delays)
            delays[i] = uint32((i + 1) * 1 days / 12 seconds); // 1 day, 2 days, 3 days, etc.

            // Set the strategy-specific delay
            cheats.prank(defaultOwner);
            factory.setStrategyEscrowDelay(strategies[i], delays[i]);

            // Start escrow for this strategy
            _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i]);
            // Verify the escrow was started correctly
            _checkStartEscrows(defaultOperatorSet, defaultSlashId, strategies[i], tokens[i], underlyingAmounts[i], i + 1);
        }

        // The complete block should be the maximum delay across all strategies
        assertEq(factory.getEscrowCompleteBlock(defaultOperatorSet, defaultSlashId), block.number + delays[numStrategies - 1] + 1);
    }
}

contract SlashEscrowFactoryUnitTests_getPendingEscrowsFull is SlashEscrowFactoryUnitTests {
    function test_getPendingEscrows() public {
        uint32 dayInBlocks = 1 days / 12 seconds;
        uint32 initialBlock = uint32(block.number);
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, 100);

        cheats.roll(block.number + dayInBlocks);
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId + 1, defaultStrategy, defaultToken, 100);

        cheats.roll(block.number + dayInBlocks);
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId + 2, defaultStrategy, defaultToken, 100);

        (OperatorSet[] memory operatorSets, bool[] memory isRedistributing, uint[][] memory slashIds, uint32[][] memory completeBlocks) =
            factory.getPendingEscrows();
        assertEq(operatorSets.length, 1);
        assertEq(operatorSets[0].key(), defaultOperatorSet.key());
        assertEq(isRedistributing.length, 1);
        assertEq(isRedistributing[0], true);
        assertEq(slashIds.length, 1);
        assertEq(slashIds[0].length, 3);
        assertEq(completeBlocks.length, 1);
        assertEq(completeBlocks[0].length, 3);

        // Assert that the complete blocks are correct
        assertEq(completeBlocks[0][0], initialBlock + defaultGlobalDelayBlocks + 1);
        assertEq(completeBlocks[0][1], initialBlock + defaultGlobalDelayBlocks + dayInBlocks + 1);
        assertEq(completeBlocks[0][2], initialBlock + defaultGlobalDelayBlocks + dayInBlocks * 2 + 1);
    }
}

contract SlashEscrowFactoryUnitTests_SlashEscrowProxy is SlashEscrowFactoryUnitTests {
    address maliciousCaller;

    function setUp() public override {
        super.setUp();
        _initiateSlashEscrow(defaultOperatorSet, defaultSlashId, defaultStrategy, defaultToken, 100);
        maliciousCaller = cheats.randomAddress();
    }

    function test_SlashEscrowProxy_InvalidDeploymentParameters_BadFactory() public {
        ISlashEscrow slashEscrow = factory.getSlashEscrow(defaultOperatorSet, defaultSlashId);

        cheats.prank(maliciousCaller);
        cheats.expectRevert(ISlashEscrow.InvalidDeploymentParameters.selector);
        slashEscrow.releaseTokens(
            ISlashEscrowFactory(maliciousCaller),
            slashEscrowImplementation,
            defaultOperatorSet,
            defaultSlashId,
            defaultRedistributionRecipient,
            defaultStrategy
        );
    }

    function test_SlashEscrowProxy_InvalidDeploymentParameters_BadSlashEscrowImplementation() public {
        ISlashEscrow slashEscrow = factory.getSlashEscrow(defaultOperatorSet, defaultSlashId);

        cheats.prank(maliciousCaller);
        cheats.expectRevert(ISlashEscrow.InvalidDeploymentParameters.selector);
        slashEscrow.releaseTokens(
            factory, ISlashEscrow(maliciousCaller), defaultOperatorSet, defaultSlashId, defaultRedistributionRecipient, defaultStrategy
        );
    }

    function test_SlashEscrowProxy_InvalidDeploymentParameters_BadOperatorSet() public {
        ISlashEscrow slashEscrow = factory.getSlashEscrow(defaultOperatorSet, defaultSlashId);

        cheats.prank(maliciousCaller);
        cheats.expectRevert(ISlashEscrow.InvalidDeploymentParameters.selector);
        slashEscrow.releaseTokens(
            factory,
            slashEscrowImplementation,
            OperatorSet(maliciousCaller, 1),
            defaultSlashId,
            defaultRedistributionRecipient,
            defaultStrategy
        );
    }

    function test_SlashEscrowProxy_InvalidDeploymentParameters_BadSlashId(uint slashId) public {
        ISlashEscrow slashEscrow = factory.getSlashEscrow(defaultOperatorSet, defaultSlashId);

        cheats.assume(slashId != defaultSlashId);

        cheats.prank(maliciousCaller);
        cheats.expectRevert(ISlashEscrow.InvalidDeploymentParameters.selector);
        slashEscrow.releaseTokens(
            factory, slashEscrowImplementation, defaultOperatorSet, slashId, defaultRedistributionRecipient, defaultStrategy
        );
    }

    function test_SlashEscrowProxy_OnlySlashEscrowFactory() public {
        ISlashEscrow slashEscrow = factory.getSlashEscrow(defaultOperatorSet, defaultSlashId);

        cheats.expectRevert(ISlashEscrow.OnlySlashEscrowFactory.selector);
        slashEscrow.releaseTokens(
            factory, slashEscrowImplementation, defaultOperatorSet, defaultSlashId, defaultRedistributionRecipient, defaultStrategy
        );
    }

    function test_SlashEscrowProxy_OnlySlashEscrowFactory_BadRecipient() public {
        ISlashEscrow slashEscrow = factory.getSlashEscrow(defaultOperatorSet, defaultSlashId);

        cheats.prank(maliciousCaller);
        cheats.expectRevert(ISlashEscrow.OnlySlashEscrowFactory.selector);
        slashEscrow.releaseTokens(factory, slashEscrowImplementation, defaultOperatorSet, defaultSlashId, maliciousCaller, defaultStrategy);
    }
}
