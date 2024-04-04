// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/core/PaymentCoordinator.sol";
import "src/contracts/strategies/StrategyBase.sol";

import "src/test/events/IPaymentCoordinatorEvents.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

/**
 * @notice Unit testing of the PaymentCoordinator contract
 * Contracts tested: PaymentCoordinator
 * Contracts not mocked: StrategyBase, PauserRegistry
 */
contract PaymentCoordinatorUnitTests is EigenLayerUnitTestSetup, IPaymentCoordinatorEvents {
    // Contract under test
    PaymentCoordinator public paymentCoordinator;
    PaymentCoordinator public paymentCoordinatorImplementation;

    // Mocks
    IERC20 token1;
    IERC20 token2;
    IERC20 token3;
    IStrategy strategyMock1;
    IStrategy strategyMock2;
    IStrategy strategyMock3;
    StrategyBase strategyImplementation;
    uint256 mockTokenInitialSupply = 10e50;
    IPaymentCoordinator.StrategyAndMultiplier[] defaultStrategyAndMultipliers;

    // Config Variables
    /// @notice Max duration is 5 epochs (2 weeks * 5 = 10 weeks in seconds)
    uint64 MAX_PAYMENT_DURATION = 86400 * 70;

    /// @notice Lower bound start range is 3 months into the past
    uint64 MAX_RETROACTIVE_LENGTH = 86400 * 90;
    /// @notice Upper bound start range is 1 month into the future
    uint64 MAX_FUTURE_LENGTH = 86400 * 30;
    /// @notice absolute min timestamp that a payment can start at
    uint64 GENESIS_PAYMENT_TIMESTAMP = 1712092632;

    /// @notice Delay in timestamp before a posted root can be claimed against
    uint64 activationDelay = 86400 * 7;
    /// @notice intervals(epochs) are 2 weeks
    uint64 calculationIntervalSeconds = 86400 * 14;
    /// @notice the commission for all operators across all avss
    uint16 globalCommissionBips = 1000;

    // PaymentCoordinator Constants

    /// @dev Index for flag that pauses payForRange payments
    uint8 internal constant PAUSED_PAY_FOR_RANGE = 0;

    /// @dev Index for flag that pauses payAllForRange payments
    uint8 internal constant PAUSED_PAY_ALL_FOR_RANGE = 1;

    /// @dev Index for flag that pauses
    uint8 internal constant PAUSED_CLAIM_PAYMENTS = 2;

    // PaymentCoordinator entities
    address paymentUpdater = address(1000);
    address defaultAVS = address(1001);
    address defaultClaimer = address(1002);
    address payAllSubmitter = address(1003);

    function setUp() public override {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        // Deploy PaymentCoordinator proxy and implementation
        paymentCoordinatorImplementation = new PaymentCoordinator(
            delegationManagerMock,
            strategyManagerMock,
            MAX_PAYMENT_DURATION,
            MAX_RETROACTIVE_LENGTH,
            MAX_FUTURE_LENGTH,
            GENESIS_PAYMENT_TIMESTAMP
        );
        paymentCoordinator = PaymentCoordinator(
            address(
                new TransparentUpgradeableProxy(
                    address(paymentCoordinatorImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        PaymentCoordinator.initialize.selector,
                        address(this), // initOwner
                        pauserRegistry,
                        0, // 0 is initialPausedStatus
                        paymentUpdater,
                        activationDelay,
                        calculationIntervalSeconds,
                        globalCommissionBips
                    )
                )
            )
        );

        // Deploy mock token and strategy
        token1 = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, address(this));
        token1 = new ERC20PresetFixedSupply("jeo boden", "MOCK2", mockTokenInitialSupply, address(this));
        token1 = new ERC20PresetFixedSupply("pepe wif avs", "MOCK3", mockTokenInitialSupply, address(this));

        strategyImplementation = new StrategyBase(strategyManagerMock);
        strategyMock1 = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, token1, pauserRegistry)
                )
            )
        );
        strategyMock2 = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, token2, pauserRegistry)
                )
            )
        );
        strategyMock3 = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, token3, pauserRegistry)
                )
            )
        );

        strategyManagerMock.setStrategyWhitelist(strategyMock1, true);
        strategyManagerMock.setStrategyWhitelist(strategyMock2, true);
        strategyManagerMock.setStrategyWhitelist(strategyMock3, true);

        defaultStrategyAndMultipliers.push(
            IPaymentCoordinator.StrategyAndMultiplier(IStrategy(address(strategyMock1)), 1e18)
        );
        defaultStrategyAndMultipliers.push(
            IPaymentCoordinator.StrategyAndMultiplier(IStrategy(address(strategyMock2)), 2e18)
        );
        defaultStrategyAndMultipliers.push(
            IPaymentCoordinator.StrategyAndMultiplier(IStrategy(address(strategyMock3)), 3e18)
        );

        // Exclude from fuzzed tests
        addressIsExcludedFromFuzzedInputs[address(paymentCoordinator)] = true;
        addressIsExcludedFromFuzzedInputs[address(paymentUpdater)] = true;

        // Set the timestamp to some time after the genesis payment timestamp
        cheats.warp(GENESIS_PAYMENT_TIMESTAMP + 5 days);
    }

    /// @notice deploy token to owner and approve paymentCoordinator. Used for deploying payment tokens
    function _deployMockPaymentTokens(address owner, uint256 numTokens) internal returns (IERC20[] memory) {
        cheats.startPrank(owner);
        IERC20[] memory tokens = new IERC20[](numTokens);
        for (uint256 i = 0; i < numTokens; ++i) {
            IERC20 token = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, owner);
            tokens[i] = token;
            token.approve(address(paymentCoordinator), mockTokenInitialSupply);
        }
        cheats.stopPrank();
        return tokens;
    }

    function _maxTimestamp(uint64 timestamp1, uint64 timestamp2) internal returns (uint64) {
        return timestamp1 > timestamp2 ? timestamp1 : timestamp2;
    }
}

contract PaymentCoordinatorUnitTests_initializeAndSetters is PaymentCoordinatorUnitTests {}

contract PaymentCoordinatorUnitTests_payForRange is PaymentCoordinatorUnitTests {
    // Revert when paused
    function test_Revert_WhenPaused() public {
        cheats.prank(pauser);
        paymentCoordinator.pause(2 ** PAUSED_PAY_FOR_RANGE);

        cheats.expectRevert("Pausable: index is paused");
        IPaymentCoordinator.RangePayment[] memory rangePayments;
        paymentCoordinator.payForRange(rangePayments);
    }

    // Revert from reentrancy
    function test_Revert_WhenReentrancy() public {}

    // Revert from duplicate payment hash submitted
    function testFuzz_Revert_WhenDuplicateHash() public {}

    // Revert with 0 length strats and multipliers
    function testFuzz_Revert_WhenEmptyStratsAndMultipliers() public {}

    // Revert with exceeding max duration
    function testFuzz_Revert_WhenExceedingMaxDuration() public {}

    // Revert with invalid interval seconds
    function testFuzz_Revert_WhenInvalidIntervalSeconds() public {}

    // Revert with retroactive payments disable and set to past
    function testFuzz_Revert_WhenRetroactivePaymentsDisabled() public {}

    // Revert with retroactive payments enabled and set too far in past
    // - either before genesis payment timestamp
    // - before max retroactive length
    function testFuzz_Revert_WhenPaymentTooStale() public {}

    // Revert with start timestamp past max future length
    function testFuzz_Revert_WhenPaymentTooFarInFuture() public {}

    // Revert with non whitelisted strategy
    function testFuzz_Revert_WhenInvalidStrategy() public {}

    /**
     * @notice test a single range payment asserting for the following
     * - correct event emitted
     * - payment nonce incrementation by 1, and range payment hash being set in storage.
     * - range payment hash being set in storage
     * - token balance before and after of avs and paymentCoordinator
     */
    function testFuzz_payForRange_SinglePayment(
        address avs,
        bool retroactivePaymentsEnabled,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setRetroactivePaymentsEnabled(retroactivePaymentsEnabled);

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 paymentToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_PAYMENT_DURATION);
        duration = duration - (duration % calculationIntervalSeconds);
        startTimestamp;
        if (retroactivePaymentsEnabled) {
            startTimestamp = bound(
                startTimestamp,
                uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint64(block.timestamp) - MAX_RETROACTIVE_LENGTH)),
                block.timestamp + uint256(MAX_FUTURE_LENGTH)
            );
        } else {
            startTimestamp = bound(startTimestamp, block.timestamp, block.timestamp + uint256(MAX_FUTURE_LENGTH));
        }

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint64(startTimestamp),
            duration: uint64(duration)
        });

        // 3. call payForRange() with expected event emitted
        cheats.startPrank(avs);
        paymentToken.approve(address(paymentCoordinator), amount);
        uint256 currPaymentNonce = paymentCoordinator.paymentNonce();
        bytes32 rangePaymentHash = keccak256(abi.encode(avs, currPaymentNonce, rangePayments[0]));

        cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
        emit RangePaymentCreated(avs, currPaymentNonce, rangePaymentHash, rangePayments[0]);
        paymentCoordinator.payForRange(rangePayments);
        cheats.stopPrank();

        assertTrue(paymentCoordinator.isRangePaymentHash(avs, rangePaymentHash), "Range payment hash not submitted");
        assertEq(currPaymentNonce + 1, paymentCoordinator.paymentNonce(), "Payment nonce not incremented");
    }

    /**
     * @notice test multiple range payments asserting for the following
     * - correct event emitted
     * - payment nonce incrementation by numPayments, and range payment hashes being set in storage.
     * - range payment hash being set in storage
     * - token balances before and after of avs and paymentCoordinator
     */
    function testFuzz_payForRange_MultiplePayments(
        address avs,
        bool retroactivePaymentsEnabled,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount,
        uint256 numPayments
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(2 <= numPayments && numPayments <= 10);
        cheats.assume(avs != address(0));
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setRetroactivePaymentsEnabled(retroactivePaymentsEnabled);

        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](numPayments);
        bytes32[] memory rangePaymentHashes = new bytes32[](numPayments);
        uint256 startPaymentNonce = paymentCoordinator.paymentNonce();
        IERC20[] memory paymentTokens = _deployMockPaymentTokens(avs, numPayments);

        // Create multiple range payments and their expected event
        for (uint256 i = 0; i < numPayments; ++i) {
            // 1. Bound fuzz inputs to valid ranges and amounts using randSeed for each
            amount = bound(amount + i, 1, mockTokenInitialSupply);
            duration = bound(duration + i, 0, MAX_PAYMENT_DURATION);
            duration = duration - (duration % calculationIntervalSeconds);
            startTimestamp;
            if (retroactivePaymentsEnabled) {
                startTimestamp = bound(
                    startTimestamp + i,
                    uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint64(block.timestamp) - MAX_RETROACTIVE_LENGTH)),
                    block.timestamp + uint256(MAX_FUTURE_LENGTH)
                );
            } else {
                startTimestamp = bound(
                    startTimestamp + i,
                    block.timestamp,
                    block.timestamp + uint256(MAX_FUTURE_LENGTH)
                );
            }

            // 2. Create range payment input param
            IPaymentCoordinator.RangePayment memory rangePayment = IPaymentCoordinator.RangePayment({
                strategiesAndMultipliers: defaultStrategyAndMultipliers,
                token: paymentTokens[i],
                amount: amount,
                startTimestamp: uint64(startTimestamp),
                duration: uint64(duration)
            });
            rangePayments[i] = rangePayment;

            // 3. expected event emitted for this rangePayment
            rangePaymentHashes[i] = keccak256(abi.encode(avs, startPaymentNonce + i, rangePayments[i]));
            cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
            emit RangePaymentCreated(avs, startPaymentNonce + i, rangePaymentHashes[i], rangePayments[i]);
        }

        // 4. call payForRange()
        cheats.prank(avs);
        paymentCoordinator.payForRange(rangePayments);

        // 5. Check for paymentNonce() and rangePaymentHashes being set
        assertEq(
            startPaymentNonce + numPayments,
            paymentCoordinator.paymentNonce(),
            "Payment nonce not incremented properly"
        );

        for (uint256 i = 0; i < numPayments; ++i) {
            assertTrue(
                paymentCoordinator.isRangePaymentHash(avs, rangePaymentHashes[i]),
                "Range payment hash not submitted"
            );
        }
    }
}

contract PaymentCoordinatorUnitTests_payAllForRange is PaymentCoordinatorUnitTests {}
// Revert if paused

// Revert if not payAllforrangeSubmitter

// Test single range payment
// - check for proper event emitted
// - check token balance before after of msg.sender and paymentCoordinator

// Test multiple range payments
// - check for proper event emitted
// - check token balance before after of msg.sender and paymentCoordinator
// - have 3 payments, 2 with same token, 1 with different token

contract PaymentCoordinatorUnitTests_submitRoot is PaymentCoordinatorUnitTests {
    // only callable by paymentUpdater
    function testFuzz_Revert_WhenNotPaymentUpdater(address invalidPaymentUpdater) public {
        cheats.prank(invalidPaymentUpdater);

        cheats.expectRevert("PaymentCoordinator: caller is not the paymentUpdater");
        paymentCoordinator.submitRoot(bytes32(0), 0, 0);
    }
}
