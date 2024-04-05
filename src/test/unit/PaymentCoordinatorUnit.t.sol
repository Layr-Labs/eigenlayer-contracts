// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/core/PaymentCoordinator.sol";
import "src/contracts/strategies/StrategyBase.sol";

import "src/test/events/IPaymentCoordinatorEvents.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/test/mocks/Reenterer.sol";

/**
 * @notice Unit testing of the PaymentCoordinator contract
 * Contracts tested: PaymentCoordinator
 * Contracts not mocked: StrategyBase, PauserRegistry
 */
contract PaymentCoordinatorUnitTests is EigenLayerUnitTestSetup, IPaymentCoordinatorEvents {
    // used for stack too deep
    struct FuzzPayForRange {
        address avs;
        bool retroactivePaymentsEnabled;
        uint256 startTimestamp;
        uint256 duration;
        uint256 amount;
    }

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

    IERC20[] paymentTokens;

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

        paymentCoordinator.setPayAllForRangeSubmitter(payAllSubmitter, true);

        // Exclude from fuzzed tests
        addressIsExcludedFromFuzzedInputs[address(paymentCoordinator)] = true;
        addressIsExcludedFromFuzzedInputs[address(paymentUpdater)] = true;

        // Set the timestamp to some time after the genesis payment timestamp
        cheats.warp(GENESIS_PAYMENT_TIMESTAMP + 5 days);
    }

    /// @notice deploy token to owner and approve paymentCoordinator. Used for deploying payment tokens
    function _deployMockPaymentTokens(address owner, uint256 numTokens) internal {
        cheats.startPrank(owner);
        for (uint256 i = 0; i < numTokens; ++i) {
            IERC20 token = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, owner);
            paymentTokens.push(token);
            token.approve(address(paymentCoordinator), mockTokenInitialSupply);
        }
        cheats.stopPrank();
    }

    function _getBalanceForTokens(IERC20[] memory tokens, address holder) internal returns (uint256[] memory) {
        uint256[] memory balances = new uint256[](tokens.length);
        for (uint256 i = 0; i < tokens.length; ++i) {
            balances[i] = tokens[i].balanceOf(holder);
        }
        return balances;
    }

    function _maxTimestamp(uint64 timestamp1, uint64 timestamp2) internal returns (uint64) {
        return timestamp1 > timestamp2 ? timestamp1 : timestamp2;
    }
}

contract PaymentCoordinatorUnitTests_initializeAndSetters is PaymentCoordinatorUnitTests {
    function testFuzz_setClaimerFor(address earner, address claimer) public filterFuzzedAddressInputs(earner) {
        cheats.prank(earner);
        paymentCoordinator.setClaimerFor(earner, claimer);
        assertEq(claimer, paymentCoordinator.claimerFor(earner), "claimerFor not set");
    }

    function testFuzz_setCalculationIntervalSeconds(uint64 intervalSeconds) public {
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setCalculationIntervalSeconds(intervalSeconds);
        assertEq(
            intervalSeconds,
            paymentCoordinator.calculationIntervalSeconds(),
            "calculationIntervalSeconds not set"
        );
    }

    function testFuzz_setRetroactivePaymentsEnabled(bool retroactivePaymentsEnabled) public {
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setRetroactivePaymentsEnabled(retroactivePaymentsEnabled);
        assertEq(
            retroactivePaymentsEnabled,
            paymentCoordinator.retroactivePaymentsEnabled(),
            "retroactivePaymentsEnabled not set"
        );
    }

    function testFuzz_setActivationDelay(uint64 activationDelay) public {
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setActivationDelay(activationDelay);
        assertEq(activationDelay, paymentCoordinator.activationDelay(), "activationDelay not set");
    }

    function testFuzz_setGlobalOperatorCommission(uint16 globalCommissionBips) public {
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setGlobalOperatorCommission(globalCommissionBips);
        assertEq(
            globalCommissionBips,
            paymentCoordinator.globalOperatorCommissionBips(),
            "globalOperatorCommissionBips not set"
        );
    }

    function testFuzz_setPaymentUpdater(address newPaymentUpdater) public {
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setPaymentUpdater(newPaymentUpdater);
        assertEq(newPaymentUpdater, paymentCoordinator.paymentUpdater(), "paymentUpdater not set");
    }

    function testFuzz_setPayAllForRangeSubmitter(address submitter, bool newValue) public {
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setPayAllForRangeSubmitter(defaultClaimer, newValue);
        assertEq(
            newValue,
            paymentCoordinator.isPayAllForRangeSubmitter(defaultClaimer),
            "isPayAllForRangeSubmitter not set"
        );
    }
}

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
    function test_Revert_WhenReentrancy(uint256 amount) public {
        Reenterer reenterer = new Reenterer();

        reenterer.prepareReturnData(abi.encode(amount));

        address targetToUse = address(paymentCoordinator);
        uint256 msgValueToUse = 0;

        _deployMockPaymentTokens(address(this), 1);

        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(reenterer)),
            amount: amount,
            startTimestamp: uint64(block.timestamp),
            duration: 0
        });

        bytes memory calldataToUse = abi.encodeWithSelector(PaymentCoordinator.payForRange.selector, rangePayments);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        cheats.expectRevert();
        paymentCoordinator.payForRange(rangePayments);
    }

    // Revert with 0 length strats and multipliers
    function testFuzz_Revert_WhenEmptyStratsAndMultipliers(
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
        IPaymentCoordinator.StrategyAndMultiplier[] memory emptyStratsAndMultipliers;
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: emptyStratsAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint64(startTimestamp),
            duration: uint64(duration)
        });

        // 3. call payForRange() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("PaymentCoordinator._payForRange: no strategies set");
        paymentCoordinator.payForRange(rangePayments);
    }

    // Revert with exceeding max duration
    function testFuzz_Revert_WhenExceedingMaxDuration(
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
        duration = bound(duration, MAX_PAYMENT_DURATION + 1, type(uint64).max);
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

        // 3. call payForRange() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("PaymentCoordinator._payForRange: duration exceeds MAX_PAYMENT_DURATION");
        paymentCoordinator.payForRange(rangePayments);
    }

    // Revert with invalid interval seconds
    function testFuzz_Revert_WhenInvalidIntervalSeconds(
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
        cheats.assume(duration % calculationIntervalSeconds != 0);
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

        // 3. call payForRange() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(
            "PaymentCoordinator._payForRange: duration must be a multiple of calculationIntervalSeconds"
        );
        paymentCoordinator.payForRange(rangePayments);
    }

    // Revert with retroactive payments disable and set to past
    function testFuzz_Revert_WhenRetroactivePaymentsDisabled(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setRetroactivePaymentsEnabled(false);

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 paymentToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_PAYMENT_DURATION);
        duration = duration - (duration % calculationIntervalSeconds);
        startTimestamp = bound(startTimestamp, block.timestamp - uint256(MAX_RETROACTIVE_LENGTH), block.timestamp - 1);

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint64(startTimestamp),
            duration: uint64(duration)
        });

        // 3. call payForRange() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("PaymentCoordinator._payForRange: startTimestamp cannot be in the past");
        paymentCoordinator.payForRange(rangePayments);
    }

    // Revert with retroactive payments enabled and set too far in past
    // - either before genesis payment timestamp
    // - before max retroactive length
    function testFuzz_Revert_WhenPaymentTooStale(
        uint256 fuzzBlockTimestamp,
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setRetroactivePaymentsEnabled(true);

        // 1. Bound fuzz inputs to valid ranges and amounts
        fuzzBlockTimestamp = bound(fuzzBlockTimestamp, uint256(MAX_RETROACTIVE_LENGTH), block.timestamp);
        cheats.warp(fuzzBlockTimestamp);

        IERC20 paymentToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_PAYMENT_DURATION);
        duration = duration - (duration % calculationIntervalSeconds);
        startTimestamp = bound(
            startTimestamp,
            0,
            uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint64(block.timestamp) - MAX_RETROACTIVE_LENGTH)) - 1
        );

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint64(startTimestamp),
            duration: uint64(duration)
        });

        // 3. call payForRange() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("PaymentCoordinator._payForRange: startTimestamp too far in the past");
        paymentCoordinator.payForRange(rangePayments);
    }

    // Revert with start timestamp past max future length
    function testFuzz_Revert_WhenPaymentTooFarInFuture(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setRetroactivePaymentsEnabled(true);

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 paymentToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_PAYMENT_DURATION);
        duration = duration - (duration % calculationIntervalSeconds);
        startTimestamp = bound(startTimestamp, block.timestamp + uint256(MAX_FUTURE_LENGTH) + 1, type(uint64).max);

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint64(startTimestamp),
            duration: uint64(duration)
        });

        // 3. call payForRange() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("PaymentCoordinator._payForRange: startTimestamp too far in the future");
        paymentCoordinator.payForRange(rangePayments);
    }

    // Revert with non whitelisted strategy
    function testFuzz_Revert_WhenInvalidStrategy(
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
        defaultStrategyAndMultipliers[0].strategy = IStrategy(address(999));
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint64(startTimestamp),
            duration: uint64(duration)
        });

        // 3. call payForRange() with expected event emitted
        cheats.prank(avs);
        cheats.expectRevert("PaymentCoordinator._payForRange: invalid strategy considered");
        paymentCoordinator.payForRange(rangePayments);
    }

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
        uint256 avsBalanceBefore = paymentToken.balanceOf(avs);
        uint256 paymentCoordinatorBalanceBefore = paymentToken.balanceOf(address(paymentCoordinator));

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
        assertEq(
            avsBalanceBefore - amount,
            paymentToken.balanceOf(avs),
            "AVS balance not decremented by amount of range payment"
        );
        assertEq(
            paymentCoordinatorBalanceBefore + amount,
            paymentToken.balanceOf(address(paymentCoordinator)),
            "PaymentCoordinator balance not incremented by amount of range payment"
        );
    }

    /**
     * @notice test multiple range payments asserting for the following
     * - correct event emitted
     * - payment nonce incrementation by numPayments, and range payment hashes being set in storage.
     * - range payment hash being set in storage
     * - token balances before and after of avs and paymentCoordinator
     */
    function testFuzz_payForRange_MultiplePayments(
        FuzzPayForRange memory param,
        uint256 numPayments
    ) public filterFuzzedAddressInputs(param.avs) {
        cheats.assume(2 <= numPayments && numPayments <= 10);
        cheats.assume(param.avs != address(0));
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setRetroactivePaymentsEnabled(param.retroactivePaymentsEnabled);

        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](numPayments);
        bytes32[] memory rangePaymentHashes = new bytes32[](numPayments);
        uint256 startPaymentNonce = paymentCoordinator.paymentNonce();
        _deployMockPaymentTokens(param.avs, numPayments);

        uint256[] memory avsBalancesBefore = _getBalanceForTokens(paymentTokens, param.avs);
        uint256[] memory paymentCoordinatorBalancesBefore = _getBalanceForTokens(
            paymentTokens,
            address(paymentCoordinator)
        );
        uint256[] memory amounts = new uint256[](numPayments);

        // Create multiple range payments and their expected event
        for (uint256 i = 0; i < numPayments; ++i) {
            // 1. Bound fuzz inputs to valid ranges and amounts using randSeed for each
            param.amount = bound(param.amount + i, 1, mockTokenInitialSupply);
            amounts[i] = param.amount;
            param.duration = bound(param.duration + i, 0, MAX_PAYMENT_DURATION);
            param.duration = param.duration - (param.duration % calculationIntervalSeconds);
            if (param.retroactivePaymentsEnabled) {
                param.startTimestamp = bound(
                    param.startTimestamp + i,
                    uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint64(block.timestamp) - MAX_RETROACTIVE_LENGTH)),
                    block.timestamp + uint256(MAX_FUTURE_LENGTH)
                );
            } else {
                param.startTimestamp = bound(
                    param.startTimestamp + i,
                    block.timestamp,
                    block.timestamp + uint256(MAX_FUTURE_LENGTH)
                );
            }

            // 2. Create range payment input param
            IPaymentCoordinator.RangePayment memory rangePayment = IPaymentCoordinator.RangePayment({
                strategiesAndMultipliers: defaultStrategyAndMultipliers,
                token: paymentTokens[i],
                amount: param.amount,
                startTimestamp: uint64(param.startTimestamp),
                duration: uint64(param.duration)
            });
            rangePayments[i] = rangePayment;

            // 3. expected event emitted for this rangePayment
            rangePaymentHashes[i] = keccak256(abi.encode(param.avs, startPaymentNonce + i, rangePayments[i]));
            cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
            emit RangePaymentCreated(param.avs, startPaymentNonce + i, rangePaymentHashes[i], rangePayments[i]);
        }

        // 4. call payForRange()
        cheats.prank(param.avs);
        paymentCoordinator.payForRange(rangePayments);

        // 5. Check for paymentNonce() and rangePaymentHashes being set
        assertEq(
            startPaymentNonce + numPayments,
            paymentCoordinator.paymentNonce(),
            "Payment nonce not incremented properly"
        );

        for (uint256 i = 0; i < numPayments; ++i) {
            assertTrue(
                paymentCoordinator.isRangePaymentHash(param.avs, rangePaymentHashes[i]),
                "Range payment hash not submitted"
            );
            assertEq(
                avsBalancesBefore[i] - amounts[i],
                paymentTokens[i].balanceOf(param.avs),
                "AVS balance not decremented by amount of range payment"
            );
            assertEq(
                paymentCoordinatorBalancesBefore[i] + amounts[i],
                paymentTokens[i].balanceOf(address(paymentCoordinator)),
                "PaymentCoordinator balance not incremented by amount of range payment"
            );
        }
    }
}

contract PaymentCoordinatorUnitTests_payAllForRange is PaymentCoordinatorUnitTests {
    // Revert when paused
    function test_Revert_WhenPaused() public {
        cheats.prank(pauser);
        paymentCoordinator.pause(2 ** PAUSED_PAY_ALL_FOR_RANGE);

        cheats.expectRevert("Pausable: index is paused");
        IPaymentCoordinator.RangePayment[] memory rangePayments;
        paymentCoordinator.payAllForRange(rangePayments);
    }

    // Revert from reentrancy
    function test_Revert_WhenReentrancy(uint256 amount) public {
        Reenterer reenterer = new Reenterer();

        reenterer.prepareReturnData(abi.encode(amount));

        address targetToUse = address(paymentCoordinator);
        uint256 msgValueToUse = 0;

        _deployMockPaymentTokens(address(this), 1);

        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(reenterer)),
            amount: amount,
            startTimestamp: uint64(block.timestamp),
            duration: 0
        });

        bytes memory calldataToUse = abi.encodeWithSelector(PaymentCoordinator.payForRange.selector, rangePayments);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        cheats.prank(payAllSubmitter);
        cheats.expectRevert();
        paymentCoordinator.payAllForRange(rangePayments);
    }

    function testFuzz_Revert_WhenNotPayAllForRangeSubmitter(
        address invalidSubmitter
    ) public filterFuzzedAddressInputs(invalidSubmitter) {
        cheats.assume(invalidSubmitter != payAllSubmitter);

        cheats.expectRevert("PaymentCoordinator: caller is not a valid payAllForRange submitter");
        IPaymentCoordinator.RangePayment[] memory rangePayments;
        paymentCoordinator.payAllForRange(rangePayments);
    }

    /**
     * @notice test a single range payment asserting for the following
     * - correct event emitted
     * - payment nonce incrementation by 1, and range payment hash being set in storage.
     * - range payment hash being set in storage
     * - token balance before and after of payAllForRangeSubmitter and paymentCoordinator
     */
    function testFuzz_payAllForRange_SinglePayment(
        bool retroactivePaymentsEnabled,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public {
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setRetroactivePaymentsEnabled(retroactivePaymentsEnabled);

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 paymentToken = new ERC20PresetFixedSupply(
            "dog wif hat",
            "MOCK1",
            mockTokenInitialSupply,
            payAllSubmitter
        );
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
        uint256 submitterBalanceBefore = paymentToken.balanceOf(payAllSubmitter);
        uint256 paymentCoordinatorBalanceBefore = paymentToken.balanceOf(address(paymentCoordinator));

        cheats.startPrank(payAllSubmitter);
        paymentToken.approve(address(paymentCoordinator), amount);
        uint256 currPaymentNonce = paymentCoordinator.paymentNonce();
        bytes32 rangePaymentHash = keccak256(abi.encode(payAllSubmitter, currPaymentNonce, rangePayments[0]));

        cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
        emit RangePaymentForAllCreated(payAllSubmitter, currPaymentNonce, rangePaymentHash, rangePayments[0]);
        paymentCoordinator.payAllForRange(rangePayments);
        cheats.stopPrank();

        assertTrue(
            paymentCoordinator.isRangePaymentHash(payAllSubmitter, rangePaymentHash),
            "Range payment hash not submitted"
        );
        assertEq(currPaymentNonce + 1, paymentCoordinator.paymentNonce(), "Payment nonce not incremented");
        assertEq(
            submitterBalanceBefore - amount,
            paymentToken.balanceOf(payAllSubmitter),
            "PayAllForRange Submitter balance not decremented by amount of range payment"
        );
        assertEq(
            paymentCoordinatorBalanceBefore + amount,
            paymentToken.balanceOf(address(paymentCoordinator)),
            "PaymentCoordinator balance not incremented by amount of range payment"
        );
    }

    /**
     * @notice test multiple range payments asserting for the following
     * - correct event emitted
     * - payment nonce incrementation by numPayments, and range payment hashes being set in storage.
     * - range payment hash being set in storage
     * - token balances before and after of payAllForRange submitter and paymentCoordinator
     */
    function testFuzz_payAllForRange_MultiplePayments(FuzzPayForRange memory param, uint256 numPayments) public {
        cheats.assume(2 <= numPayments && numPayments <= 10);
        cheats.prank(paymentCoordinator.owner());
        paymentCoordinator.setRetroactivePaymentsEnabled(param.retroactivePaymentsEnabled);

        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](numPayments);
        bytes32[] memory rangePaymentHashes = new bytes32[](numPayments);
        uint256 startPaymentNonce = paymentCoordinator.paymentNonce();
        _deployMockPaymentTokens(payAllSubmitter, numPayments);

        uint256[] memory submitterBalancesBefore = _getBalanceForTokens(paymentTokens, payAllSubmitter);
        uint256[] memory paymentCoordinatorBalancesBefore = _getBalanceForTokens(
            paymentTokens,
            address(paymentCoordinator)
        );
        uint256[] memory amounts = new uint256[](numPayments);

        // Create multiple range payments and their expected event
        for (uint256 i = 0; i < numPayments; ++i) {
            // 1. Bound fuzz inputs to valid ranges and amounts using randSeed for each
            param.amount = bound(param.amount + i, 1, mockTokenInitialSupply);
            amounts[i] = param.amount;
            param.duration = bound(param.duration + i, 0, MAX_PAYMENT_DURATION);
            param.duration = param.duration - (param.duration % calculationIntervalSeconds);
            if (param.retroactivePaymentsEnabled) {
                param.startTimestamp = bound(
                    param.startTimestamp + i,
                    uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint64(block.timestamp) - MAX_RETROACTIVE_LENGTH)),
                    block.timestamp + uint256(MAX_FUTURE_LENGTH)
                );
            } else {
                param.startTimestamp = bound(
                    param.startTimestamp + i,
                    block.timestamp,
                    block.timestamp + uint256(MAX_FUTURE_LENGTH)
                );
            }

            // 2. Create range payment input param
            IPaymentCoordinator.RangePayment memory rangePayment = IPaymentCoordinator.RangePayment({
                strategiesAndMultipliers: defaultStrategyAndMultipliers,
                token: paymentTokens[i],
                amount: amounts[i],
                startTimestamp: uint64(param.startTimestamp),
                duration: uint64(param.duration)
            });
            rangePayments[i] = rangePayment;

            // 3. expected event emitted for this rangePayment
            rangePaymentHashes[i] = keccak256(abi.encode(payAllSubmitter, startPaymentNonce + i, rangePayments[i]));
            cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
            emit RangePaymentForAllCreated(
                payAllSubmitter,
                startPaymentNonce + i,
                rangePaymentHashes[i],
                rangePayments[i]
            );
        }

        // 4. call payForRange()
        cheats.prank(payAllSubmitter);
        paymentCoordinator.payAllForRange(rangePayments);

        // 5. Check for paymentNonce() and rangePaymentHashes being set
        assertEq(
            startPaymentNonce + numPayments,
            paymentCoordinator.paymentNonce(),
            "Payment nonce not incremented properly"
        );

        for (uint256 i = 0; i < numPayments; ++i) {
            assertTrue(
                paymentCoordinator.isRangePaymentHash(payAllSubmitter, rangePaymentHashes[i]),
                "Range payment hash not submitted"
            );
            assertEq(
                submitterBalancesBefore[i] - amounts[i],
                paymentTokens[i].balanceOf(payAllSubmitter),
                "PayAllForRange Submitter balance not decremented by amount of range payment"
            );
            assertEq(
                paymentCoordinatorBalancesBefore[i] + amounts[i],
                paymentTokens[i].balanceOf(address(paymentCoordinator)),
                "PaymentCoordinator balance not incremented by amount of range payment"
            );
        }
    }
}

contract PaymentCoordinatorUnitTests_submitRoot is PaymentCoordinatorUnitTests {
    // only callable by paymentUpdater
    function testFuzz_Revert_WhenNotPaymentUpdater(
        address invalidPaymentUpdater
    ) public filterFuzzedAddressInputs(invalidPaymentUpdater) {
        cheats.prank(invalidPaymentUpdater);

        cheats.expectRevert("PaymentCoordinator: caller is not the paymentUpdater");
        paymentCoordinator.submitRoot(bytes32(0), 0, 0);
    }

    function testFuzz_Revert_WhenActivatedAtInPast(
        bytes32 root,
        uint64 paymentCalculationEndTimestamp,
        uint64 activatedAt
    ) public {
        cheats.prank(paymentUpdater);
        cheats.assume(activatedAt < block.timestamp);

        cheats.expectRevert("PaymentCoordinator.submitRoot: activatedAt can't be in the past");
        paymentCoordinator.submitRoot(root, paymentCalculationEndTimestamp, activatedAt);
    }

    /// @notice submits root with correct values and adds to root storage array
    function testFuzz_submitRoot() public {}
}

contract PaymentCoordinatorUnitTests_processClaim is PaymentCoordinatorUnitTests {}
