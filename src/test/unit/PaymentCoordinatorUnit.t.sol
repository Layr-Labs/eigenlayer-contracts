// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/core/PaymentCoordinator.sol";
import "src/contracts/strategies/StrategyBase.sol";

import "src/test/events/IPaymentCoordinatorEvents.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/test/mocks/Reenterer.sol";
import "src/test/mocks/ERC20Mock.sol";

/**
 * @notice Unit testing of the PaymentCoordinator contract
 * Contracts tested: PaymentCoordinator
 * Contracts not mocked: StrategyBase, PauserRegistry
 */
contract PaymentCoordinatorUnitTests is EigenLayerUnitTestSetup, IPaymentCoordinatorEvents {
    // used for stack too deep
    struct FuzzPayForRange {
        address avs;
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
    uint256 mockTokenInitialSupply = 1e38 - 1;
    IPaymentCoordinator.StrategyAndMultiplier[] defaultStrategyAndMultipliers;

    // Config Variables
    /// @notice intervals(epochs) are 1 weeks
    uint32 CALCULATION_INTERVAL_SECONDS = 7 days;

    /// @notice Max duration is 5 epochs (2 weeks * 5 = 10 weeks in seconds)
    uint32 MAX_PAYMENT_DURATION = 70 days;

    /// @notice Lower bound start range is ~3 months into the past, multiple of CALCULATION_INTERVAL_SECONDS
    uint32 MAX_RETROACTIVE_LENGTH = 84 days;
    /// @notice Upper bound start range is ~1 month into the future, multiple of CALCULATION_INTERVAL_SECONDS
    uint32 MAX_FUTURE_LENGTH = 28 days;
    /// @notice absolute min timestamp that a payment can start at
    uint32 GENESIS_PAYMENT_TIMESTAMP = 1712188800;

    /// @notice Delay in timestamp before a posted root can be claimed against
    uint32 activationDelay = 7 days;
    /// @notice the commission for all operators across all avss
    uint16 globalCommissionBips = 1000;

    IERC20[] paymentTokens;

    // PaymentCoordinator Constants

    /// @dev Index for flag that pauses payForRange payments
    uint8 internal constant PAUSED_PAY_FOR_RANGE = 0;

    /// @dev Index for flag that pauses payAllForRange payments
    uint8 internal constant PAUSED_PAY_ALL_FOR_RANGE = 1;

    /// @dev Index for flag that pauses claiming
    uint8 internal constant PAUSED_CLAIM_PAYMENTS = 2;

    /// @dev Index for flag that pauses submitRoots
    uint8 internal constant PAUSED_SUBMIT_ROOTS = 3;

    // PaymentCoordinator entities
    address paymentUpdater = address(1000);
    address defaultAVS = address(1001);
    address defaultClaimer = address(1002);
    address payAllSubmitter = address(1003);

    function setUp() public virtual override {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        // Deploy PaymentCoordinator proxy and implementation
        paymentCoordinatorImplementation = new PaymentCoordinator(
            delegationManagerMock,
            strategyManagerMock,
            CALCULATION_INTERVAL_SECONDS,
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
                        globalCommissionBips
                    )
                )
            )
        );

        // Deploy mock token and strategy
        token1 = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, address(this));
        token2 = new ERC20PresetFixedSupply("jeo boden", "MOCK2", mockTokenInitialSupply, address(this));
        token3 = new ERC20PresetFixedSupply("pepe wif avs", "MOCK3", mockTokenInitialSupply, address(this));

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
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = strategyMock1;
        strategies[1] = strategyMock2;
        strategies[2] = strategyMock3;
        strategies = _sortArrayAsc(strategies);

        strategyManagerMock.setStrategyWhitelist(strategies[0], true);
        strategyManagerMock.setStrategyWhitelist(strategies[1], true);
        strategyManagerMock.setStrategyWhitelist(strategies[2], true);

        defaultStrategyAndMultipliers.push(
            IPaymentCoordinator.StrategyAndMultiplier(IStrategy(address(strategies[0])), 1e18)
        );
        defaultStrategyAndMultipliers.push(
            IPaymentCoordinator.StrategyAndMultiplier(IStrategy(address(strategies[1])), 2e18)
        );
        defaultStrategyAndMultipliers.push(
            IPaymentCoordinator.StrategyAndMultiplier(IStrategy(address(strategies[2])), 3e18)
        );

        paymentCoordinator.setPayAllForRangeSubmitter(payAllSubmitter, true);
        paymentCoordinator.setPaymentUpdater(paymentUpdater);

        // Exclude from fuzzed tests
        addressIsExcludedFromFuzzedInputs[address(paymentCoordinator)] = true;
        addressIsExcludedFromFuzzedInputs[address(paymentUpdater)] = true;

        // Set the timestamp to some time after the genesis payment timestamp
        cheats.warp(GENESIS_PAYMENT_TIMESTAMP + 5 days);
    }

    /// @notice deploy token to owner and approve paymentCoordinator. Used for deploying payment tokens
    function _deployMockPaymentTokens(address owner, uint256 numTokens) internal virtual {
        cheats.startPrank(owner);
        for (uint256 i = 0; i < numTokens; ++i) {
            IERC20 token = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, owner);
            paymentTokens.push(token);
            token.approve(address(paymentCoordinator), mockTokenInitialSupply);
        }
        cheats.stopPrank();
    }

    function _getBalanceForTokens(IERC20[] memory tokens, address holder) internal view returns (uint256[] memory) {
        uint256[] memory balances = new uint256[](tokens.length);
        for (uint256 i = 0; i < tokens.length; ++i) {
            balances[i] = tokens[i].balanceOf(holder);
        }
        return balances;
    }

    function _maxTimestamp(uint32 timestamp1, uint32 timestamp2) internal pure returns (uint32) {
        return timestamp1 > timestamp2 ? timestamp1 : timestamp2;
    }

    function _assertPaymentClaimedEvents(bytes32 root, IPaymentCoordinator.PaymentMerkleClaim memory claim, address recipient) internal {
        address earner = claim.earnerLeaf.earner;
        address claimer = paymentCoordinator.claimerFor(earner);
        if (claimer == address(0)) {
            claimer = earner;
        }
        IERC20 token;
        uint256 claimedAmount;
        for (uint256 i = 0; i < claim.tokenLeaves.length; ++i) {
            token = claim.tokenLeaves[i].token;
            claimedAmount = paymentCoordinator.cumulativeClaimed(earner, token);

            cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
            emit PaymentClaimed(
                root,
                earner,
                claimer,
                recipient,
                token,
                claim.tokenLeaves[i].cumulativeEarnings - claimedAmount
            );
        }
    }

    /// @notice given address and array of payment tokens, return array of cumulativeClaimed amonts
    function _getCumulativeClaimed(
        address earner,
        IPaymentCoordinator.PaymentMerkleClaim memory claim
    ) internal view returns (uint256[] memory) {
        uint256[] memory totalClaimed = new uint256[](claim.tokenLeaves.length);

        for (uint256 i = 0; i < claim.tokenLeaves.length; ++i) {
            totalClaimed[i] = paymentCoordinator.cumulativeClaimed(earner, claim.tokenLeaves[i].token);
        }

        return totalClaimed;
    }

    /// @notice given a claim, return the new cumulativeEarnings for each token
    function _getCumulativeEarnings(
        IPaymentCoordinator.PaymentMerkleClaim memory claim
    ) internal pure returns (uint256[] memory) {
        uint256[] memory earnings = new uint256[](claim.tokenLeaves.length);

        for (uint256 i = 0; i < claim.tokenLeaves.length; ++i) {
            earnings[i] = claim.tokenLeaves[i].cumulativeEarnings;
        }

        return earnings;
    }

    function _getClaimTokenBalances(
        address earner,
        IPaymentCoordinator.PaymentMerkleClaim memory claim
    ) internal view returns (uint256[] memory) {
        uint256[] memory balances = new uint256[](claim.tokenLeaves.length);

        for (uint256 i = 0; i < claim.tokenLeaves.length; ++i) {
            balances[i] = claim.tokenLeaves[i].token.balanceOf(earner);
        }

        return balances;
    }

    /// @dev Sort to ensure that the array is in ascending order for strategies
    function _sortArrayAsc(IStrategy[] memory arr) internal pure returns (IStrategy[] memory) {
        uint256 l = arr.length;
        for (uint256 i = 0; i < l; i++) {
            for (uint256 j = i + 1; j < l; j++) {
                if (address(arr[i]) > address(arr[j])) {
                    IStrategy temp = arr[i];
                    arr[i] = arr[j];
                    arr[j] = temp;
                }
            }
        }
        return arr;
    }
}

contract PaymentCoordinatorUnitTests_initializeAndSetters is PaymentCoordinatorUnitTests {
    function testFuzz_setClaimerFor(address earner, address claimer) public filterFuzzedAddressInputs(earner) {
        cheats.startPrank(earner);
        cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
        emit ClaimerForSet(earner, paymentCoordinator.claimerFor(earner), claimer);
        paymentCoordinator.setClaimerFor(claimer);
        assertEq(claimer, paymentCoordinator.claimerFor(earner), "claimerFor not set");
        cheats.stopPrank();
    }

    function testFuzz_setActivationDelay(uint32 activationDelay) public {
        cheats.startPrank(paymentCoordinator.owner());
        cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
        emit ActivationDelaySet(paymentCoordinator.activationDelay(), activationDelay);
        paymentCoordinator.setActivationDelay(activationDelay);
        assertEq(activationDelay, paymentCoordinator.activationDelay(), "activationDelay not set");
        cheats.stopPrank();
    }

    function testFuzz_setActivationDelay_Revert_WhenNotOwner(
        address caller,
        uint32 activationDelay
    ) public filterFuzzedAddressInputs(caller) {
        cheats.assume(caller != paymentCoordinator.owner());
        cheats.prank(caller);
        cheats.expectRevert("Ownable: caller is not the owner");
        paymentCoordinator.setActivationDelay(activationDelay);
    }

    function testFuzz_setGlobalOperatorCommission(uint16 globalCommissionBips) public {
        cheats.startPrank(paymentCoordinator.owner());
        cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
        emit GlobalCommissionBipsSet(paymentCoordinator.globalOperatorCommissionBips(), globalCommissionBips);
        paymentCoordinator.setGlobalOperatorCommission(globalCommissionBips);
        assertEq(
            globalCommissionBips,
            paymentCoordinator.globalOperatorCommissionBips(),
            "globalOperatorCommissionBips not set"
        );
        cheats.stopPrank();
    }

    function testFuzz_setGlobalOperatorCommission_Revert_WhenNotOwner(
        address caller,
        uint16 globalCommissionBips
    ) public filterFuzzedAddressInputs(caller) {
        cheats.assume(caller != paymentCoordinator.owner());
        cheats.prank(caller);
        cheats.expectRevert("Ownable: caller is not the owner");
        paymentCoordinator.setGlobalOperatorCommission(globalCommissionBips);
    }

    function testFuzz_setPaymentUpdater(address newPaymentUpdater) public {
        cheats.startPrank(paymentCoordinator.owner());
        cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
        emit PaymentUpdaterSet(paymentCoordinator.paymentUpdater(), newPaymentUpdater);
        paymentCoordinator.setPaymentUpdater(newPaymentUpdater);
        assertEq(newPaymentUpdater, paymentCoordinator.paymentUpdater(), "paymentUpdater not set");
        cheats.stopPrank();
    }

    function testFuzz_setPaymentUpdater_Revert_WhenNotOwner(
        address caller,
        address newPaymentUpdater
    ) public filterFuzzedAddressInputs(caller) {
        cheats.assume(caller != paymentCoordinator.owner());
        cheats.prank(caller);
        cheats.expectRevert("Ownable: caller is not the owner");
        paymentCoordinator.setPaymentUpdater(newPaymentUpdater);
    }

    function testFuzz_setPayAllForRangeSubmitter(address submitter, bool newValue) public {
        cheats.startPrank(paymentCoordinator.owner());
        cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
        emit PayAllForRangeSubmitterSet(submitter, paymentCoordinator.isPayAllForRangeSubmitter(submitter), newValue);
        paymentCoordinator.setPayAllForRangeSubmitter(submitter, newValue);
        assertEq(
            newValue,
            paymentCoordinator.isPayAllForRangeSubmitter(submitter),
            "isPayAllForRangeSubmitter not set"
        );
        cheats.stopPrank();
    }

    function testFuzz_setPayAllForRangeSubmitter_Revert_WhenNotOwner(
        address caller,
        address submitter,
        bool newValue
    ) public filterFuzzedAddressInputs(caller) {
        cheats.assume(caller != paymentCoordinator.owner());
        cheats.prank(caller);
        cheats.expectRevert("Ownable: caller is not the owner");
        paymentCoordinator.setPayAllForRangeSubmitter(submitter, newValue);
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
        amount = bound(amount, 1, 1e38-1);
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
            startTimestamp: uint32(block.timestamp),
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
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(paymentCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 paymentToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_PAYMENT_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        IPaymentCoordinator.StrategyAndMultiplier[] memory emptyStratsAndMultipliers;
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: emptyStratsAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call payForRange() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("PaymentCoordinator._payForRange: no strategies set");
        paymentCoordinator.payForRange(rangePayments);
    }

    // Revert when amount > 1e38-1
    function testFuzz_Revert_AmountTooLarge(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        // 1. Bound fuzz inputs
        amount = bound(amount, 1e38, type(uint256).max);
        IERC20 paymentToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", amount, avs);
        duration = bound(duration, 0, MAX_PAYMENT_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. Call payForRange() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("PaymentCoordinator._payForRange: amount too large");
        paymentCoordinator.payForRange(rangePayments);
    }

    function testFuzz_Revert_WhenDuplicateStrategies(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 paymentToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_PAYMENT_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        IPaymentCoordinator.StrategyAndMultiplier[]
            memory dupStratsAndMultipliers = new IPaymentCoordinator.StrategyAndMultiplier[](2);
        dupStratsAndMultipliers[0] = defaultStrategyAndMultipliers[0];
        dupStratsAndMultipliers[1] = defaultStrategyAndMultipliers[0];
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: dupStratsAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call payForRange() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(
            "PaymentCoordinator._payForRange: strategies must be in ascending order to handle duplicates"
        );
        paymentCoordinator.payForRange(rangePayments);
    }

    // Revert with exceeding max duration
    function testFuzz_Revert_WhenExceedingMaxDuration(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(paymentCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 paymentToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, MAX_PAYMENT_DURATION + 1, type(uint32).max);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call payForRange() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("PaymentCoordinator._payForRange: duration exceeds MAX_PAYMENT_DURATION");
        paymentCoordinator.payForRange(rangePayments);
    }

    // Revert with invalid interval seconds
    function testFuzz_Revert_WhenInvalidIntervalSeconds(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(paymentCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 paymentToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_PAYMENT_DURATION);
        cheats.assume(duration % CALCULATION_INTERVAL_SECONDS != 0);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call payForRange() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(
            "PaymentCoordinator._payForRange: duration must be a multiple of CALCULATION_INTERVAL_SECONDS"
        );
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

        // 1. Bound fuzz inputs to valid ranges and amounts
        fuzzBlockTimestamp = bound(fuzzBlockTimestamp, uint256(MAX_RETROACTIVE_LENGTH), block.timestamp);
        cheats.warp(fuzzBlockTimestamp);

        IERC20 paymentToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_PAYMENT_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            0,
            uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
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

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 paymentToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_PAYMENT_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            block.timestamp + uint256(MAX_FUTURE_LENGTH) + 1 + CALCULATION_INTERVAL_SECONDS,
            type(uint32).max
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call payForRange() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("PaymentCoordinator._payForRange: startTimestamp too far in the future");
        paymentCoordinator.payForRange(rangePayments);
    }

    // Revert with non whitelisted strategy
    function testFuzz_Revert_WhenInvalidStrategy(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(paymentCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 paymentToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_PAYMENT_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        defaultStrategyAndMultipliers[0].strategy = IStrategy(address(999));
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
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
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(paymentCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 paymentToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_PAYMENT_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call payForRange() with expected event emitted
        uint256 avsBalanceBefore = paymentToken.balanceOf(avs);
        uint256 paymentCoordinatorBalanceBefore = paymentToken.balanceOf(address(paymentCoordinator));

        cheats.startPrank(avs);
        paymentToken.approve(address(paymentCoordinator), amount);
        uint256 currPaymentNonce = paymentCoordinator.paymentNonce(avs);
        bytes32 rangePaymentHash = keccak256(abi.encode(avs, currPaymentNonce, rangePayments[0]));

        cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
        emit RangePaymentCreated(avs, currPaymentNonce, rangePaymentHash, rangePayments[0]);
        paymentCoordinator.payForRange(rangePayments);
        cheats.stopPrank();

        assertTrue(paymentCoordinator.isRangePaymentHash(avs, rangePaymentHash), "Range payment hash not submitted");
        assertEq(currPaymentNonce + 1, paymentCoordinator.paymentNonce(avs), "Payment nonce not incremented");
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

        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](numPayments);
        bytes32[] memory rangePaymentHashes = new bytes32[](numPayments);
        uint256 startPaymentNonce = paymentCoordinator.paymentNonce(param.avs);
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
            param.duration = param.duration - (param.duration % CALCULATION_INTERVAL_SECONDS);
            param.startTimestamp = bound(
                param.startTimestamp + i,
                uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                    CALCULATION_INTERVAL_SECONDS -
                    1,
                block.timestamp + uint256(MAX_FUTURE_LENGTH)
            );
            param.startTimestamp = param.startTimestamp - (param.startTimestamp % CALCULATION_INTERVAL_SECONDS);

            // 2. Create range payment input param
            IPaymentCoordinator.RangePayment memory rangePayment = IPaymentCoordinator.RangePayment({
                strategiesAndMultipliers: defaultStrategyAndMultipliers,
                token: paymentTokens[i],
                amount: amounts[i],
                startTimestamp: uint32(param.startTimestamp),
                duration: uint32(param.duration)
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
            paymentCoordinator.paymentNonce(param.avs),
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
            startTimestamp: uint32(block.timestamp),
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
    function testFuzz_payAllForRange_SinglePayment(uint256 startTimestamp, uint256 duration, uint256 amount) public {
        cheats.prank(paymentCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 paymentToken = new ERC20PresetFixedSupply(
            "dog wif hat",
            "MOCK1",
            mockTokenInitialSupply,
            payAllSubmitter
        );
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_PAYMENT_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create range payment input param
        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](1);
        rangePayments[0] = IPaymentCoordinator.RangePayment({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call payForRange() with expected event emitted
        uint256 submitterBalanceBefore = paymentToken.balanceOf(payAllSubmitter);
        uint256 paymentCoordinatorBalanceBefore = paymentToken.balanceOf(address(paymentCoordinator));

        cheats.startPrank(payAllSubmitter);
        paymentToken.approve(address(paymentCoordinator), amount);
        uint256 currPaymentNonce = paymentCoordinator.paymentNonce(payAllSubmitter);
        bytes32 rangePaymentHash = keccak256(abi.encode(payAllSubmitter, currPaymentNonce, rangePayments[0]));

        cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
        emit RangePaymentForAllCreated(payAllSubmitter, currPaymentNonce, rangePaymentHash, rangePayments[0]);
        paymentCoordinator.payAllForRange(rangePayments);
        cheats.stopPrank();

        assertTrue(
            paymentCoordinator.isRangePaymentForAllHash(payAllSubmitter, rangePaymentHash),
            "Range payment hash not submitted"
        );
        assertEq(
            currPaymentNonce + 1,
            paymentCoordinator.paymentNonce(payAllSubmitter),
            "Payment nonce not incremented"
        );
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

        IPaymentCoordinator.RangePayment[] memory rangePayments = new IPaymentCoordinator.RangePayment[](numPayments);
        bytes32[] memory rangePaymentHashes = new bytes32[](numPayments);
        uint256 startPaymentNonce = paymentCoordinator.paymentNonce(payAllSubmitter);
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
            param.duration = param.duration - (param.duration % CALCULATION_INTERVAL_SECONDS);
            param.startTimestamp = bound(
                param.startTimestamp + i,
                uint256(_maxTimestamp(GENESIS_PAYMENT_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                    CALCULATION_INTERVAL_SECONDS -
                    1,
                block.timestamp + uint256(MAX_FUTURE_LENGTH)
            );
            param.startTimestamp = param.startTimestamp - (param.startTimestamp % CALCULATION_INTERVAL_SECONDS);

            // 2. Create range payment input param
            IPaymentCoordinator.RangePayment memory rangePayment = IPaymentCoordinator.RangePayment({
                strategiesAndMultipliers: defaultStrategyAndMultipliers,
                token: paymentTokens[i],
                amount: amounts[i],
                startTimestamp: uint32(param.startTimestamp),
                duration: uint32(param.duration)
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
            paymentCoordinator.paymentNonce(payAllSubmitter),
            "Payment nonce not incremented properly"
        );

        for (uint256 i = 0; i < numPayments; ++i) {
            assertTrue(
                paymentCoordinator.isRangePaymentForAllHash(payAllSubmitter, rangePaymentHashes[i]),
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
        paymentCoordinator.submitRoot(bytes32(0), 0);
    }

    function test_Revert_WhenSubmitRootPaused() public {
        cheats.prank(pauser);
        paymentCoordinator.pause(2 ** PAUSED_SUBMIT_ROOTS);

        cheats.expectRevert("Pausable: index is paused");
        paymentCoordinator.submitRoot(bytes32(0), 0);
    }

    /// @notice submits root with correct values and adds to root storage array
    /// - checks activatedAt has added activationDelay
    function testFuzz_submitRoot(bytes32 root, uint32 paymentCalculationEndTimestamp) public {
        // fuzz avoiding overflows and valid activatedAt values
        cheats.assume(paymentCoordinator.currPaymentCalculationEndTimestamp() < paymentCalculationEndTimestamp);
        cheats.assume(paymentCalculationEndTimestamp < block.timestamp);

        uint32 expectedRootIndex = uint32(paymentCoordinator.getDistributionRootsLength());
        uint32 activatedAt = uint32(block.timestamp) + paymentCoordinator.activationDelay();

        cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
        emit DistributionRootSubmitted(expectedRootIndex, root, paymentCalculationEndTimestamp, activatedAt);
        cheats.prank(paymentUpdater);
        paymentCoordinator.submitRoot(root, paymentCalculationEndTimestamp);

        (
            bytes32 submittedRoot,
            uint32 submittedPaymentCalculationEndTimestamp,
            uint32 submittedActivatedAt
        ) = paymentCoordinator.distributionRoots(expectedRootIndex);

        assertEq(
            expectedRootIndex,
            paymentCoordinator.getDistributionRootsLength() - 1,
            "root not added to roots array"
        );
        assertEq(activatedAt, submittedActivatedAt, "activatedAt not correct");
        assertEq(root, submittedRoot, "root not set");
        assertEq(
            paymentCalculationEndTimestamp,
            submittedPaymentCalculationEndTimestamp,
            "paymentCalculationEndTimestamp not set"
        );
        assertEq(
            paymentCoordinator.currPaymentCalculationEndTimestamp(),
            paymentCalculationEndTimestamp,
            "currPaymentCalculationEndTimestamp not set"
        );
    }

    /// @notice Submits multiple roots and checks root index from hash is correct
    function testFuzz_getRootIndexFromHash(bytes32 root, uint16 numRoots, uint256 index) public {
        numRoots = uint16(bound(numRoots, 1, 100));
        index = bound(index, 0, uint256(numRoots - 1));

        bytes32[] memory roots = new bytes32[](numRoots);
        cheats.startPrank(paymentUpdater);
        for (uint16 i = 0; i < numRoots; ++i) {
            roots[i] = keccak256(abi.encodePacked(root, i));

            uint32 activationDelay = uint32(block.timestamp) + paymentCoordinator.activationDelay();
            paymentCoordinator.submitRoot(roots[i], uint32(block.timestamp - 1));
            cheats.warp(activationDelay);
        }
        cheats.stopPrank();

        assertEq(index, paymentCoordinator.getRootIndexFromHash(roots[index]), "root index not found");
    }
}

/// @notice Tests for sets of JSON data with different distribution roots
contract PaymentCoordinatorUnitTests_processClaim is PaymentCoordinatorUnitTests {
    using stdStorage for StdStorage;

    /// @notice earner address used for proofs
    address earner = 0xF2288D736d27C1584Ebf7be5f52f9E4d47251AeE;

    /// @notice mock token bytecode
    bytes mockTokenBytecode;

    uint32 prevRootCalculationEndTimestamp;

    // Temp storage for managing stack in _parseProofData
    bytes32 merkleRoot;
    uint32 earnerIndex;
    bytes earnerTreeProof;
    address proofEarner;
    bytes32 earnerTokenRoot;

    function setUp() public virtual override {
        PaymentCoordinatorUnitTests.setUp();

        // Create mock token to use bytecode later to etch
        IERC20 mockToken = new ERC20Mock();
        mockTokenBytecode = address(mockToken).code;
    }

    /// @notice Claim against latest submitted root, rootIndex 3
    /// Limit fuzz runs to speed up tests since these require reading from JSON
    /// forge-config: default.fuzz.runs = 10
    function testFuzz_processClaim_LatestRoot(
        bool setClaimerFor,
        address claimerFor
    ) public filterFuzzedAddressInputs(claimerFor) {
        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            paymentCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = _parseAllProofs();
        IPaymentCoordinator.PaymentMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        (bytes32 root, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
        cheats.warp(activatedAt);

        // Claim against root and check balances before/after, and check it matches the difference between
        // cumulative claimed and earned.
        cheats.startPrank(claimer);
        assertTrue(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

        uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
        uint256[] memory earnings = _getCumulativeEarnings(claim);
        uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

        _assertPaymentClaimedEvents(root, claim, claimer);
        paymentCoordinator.processClaim(claim, claimer);

        uint256[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

        for (uint256 i = 0; i < totalClaimedBefore.length; ++i) {
            assertEq(
                earnings[i] - totalClaimedBefore[i],
                tokenBalancesAfter[i] - tokenBalancesBefore[i],
                "Token balance not incremented by earnings amount"
            );
        }

        cheats.stopPrank();
    }

    /// @notice Claim against an old root that isn't the latest
    /// forge-config: default.fuzz.runs = 10
    function testFuzz_processClaim_OldRoot(
        bool setClaimerFor,
        address claimerFor
    ) public filterFuzzedAddressInputs(claimerFor) {
        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            paymentCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = _parseAllProofs();
        IPaymentCoordinator.PaymentMerkleClaim memory claim = claims[0];

        uint32 rootIndex = claim.rootIndex;
        (bytes32 root, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
        cheats.warp(activatedAt);

        // Claim against root and check balances before/after, and check it matches the difference between
        // cumulative claimed and earned.
        cheats.startPrank(claimer);
        assertTrue(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

        uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
        uint256[] memory earnings = _getCumulativeEarnings(claim);
        uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

        _assertPaymentClaimedEvents(root, claim, claimer);
        paymentCoordinator.processClaim(claim, claimer);

        uint256[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

        for (uint256 i = 0; i < totalClaimedBefore.length; ++i) {
            assertEq(
                earnings[i] - totalClaimedBefore[i],
                tokenBalancesAfter[i] - tokenBalancesBefore[i],
                "Token balance not incremented by earnings amount"
            );
        }

        cheats.stopPrank();
    }

    /// @notice Claim against all roots in order, rootIndex 0, 1, 2
    /// forge-config: default.fuzz.runs = 10
    function testFuzz_processClaim_Sequential(
        bool setClaimerFor,
        address claimerFor
    ) public filterFuzzedAddressInputs(claimerFor) {
        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            paymentCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = _parseAllProofs();
        IPaymentCoordinator.PaymentMerkleClaim memory claim = claims[0];

        // 1. Claim against first root
        {
            uint32 rootIndex = claim.rootIndex;
            (bytes32 root, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
            cheats.warp(activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

            uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint256[] memory earnings = _getCumulativeEarnings(claim);
            uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            _assertPaymentClaimedEvents(root, claim, claimer);
            paymentCoordinator.processClaim(claim, claimer);

            uint256[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint256 i = 0; i < totalClaimedBefore.length; ++i) {
                assertEq(
                    earnings[i] - totalClaimedBefore[i],
                    tokenBalancesAfter[i] - tokenBalancesBefore[i],
                    "Token balance not incremented by earnings amount"
                );
            }

            cheats.stopPrank();
        }

        // 2. Claim against second root
        claim = claims[1];
        {
            uint32 rootIndex = claim.rootIndex;
            (bytes32 root, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
            cheats.warp(activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

            uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint256[] memory earnings = _getCumulativeEarnings(claim);
            uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            _assertPaymentClaimedEvents(root, claim, claimer);
            paymentCoordinator.processClaim(claim, claimer);

            uint256[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint256 i = 0; i < totalClaimedBefore.length; ++i) {
                assertEq(
                    earnings[i] - totalClaimedBefore[i],
                    tokenBalancesAfter[i] - tokenBalancesBefore[i],
                    "Token balance not incremented by earnings amount"
                );
            }

            cheats.stopPrank();
        }

        // 3. Claim against third and latest root
        claim = claims[2];
        {
            uint32 rootIndex = claim.rootIndex;
            (bytes32 root, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
            cheats.warp(activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

            uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint256[] memory earnings = _getCumulativeEarnings(claim);
            uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            _assertPaymentClaimedEvents(root, claim, claimer);
            paymentCoordinator.processClaim(claim, claimer);

            uint256[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint256 i = 0; i < totalClaimedBefore.length; ++i) {
                assertEq(
                    earnings[i] - totalClaimedBefore[i],
                    tokenBalancesAfter[i] - tokenBalancesBefore[i],
                    "Token balance not incremented by earnings amount"
                );
            }

            cheats.stopPrank();
        }
    }

    /// @notice Claim against rootIndex 0 and claim again. Balances should not increment.
    /// forge-config: default.fuzz.runs = 10
    function testFuzz_processClaim_Revert_WhenReuseSameClaimAgain(
        bool setClaimerFor,
        address claimerFor
    ) public filterFuzzedAddressInputs(claimerFor) {
        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            paymentCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = _parseAllProofs();
        IPaymentCoordinator.PaymentMerkleClaim memory claim = claims[0];

        // 1. Claim against first root
        {
            uint32 rootIndex = claim.rootIndex;
            (bytes32 root, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
            cheats.warp(activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

            uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint256[] memory earnings = _getCumulativeEarnings(claim);
            uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            _assertPaymentClaimedEvents(root, claim, claimer);
            paymentCoordinator.processClaim(claim, claimer);

            uint256[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint256 i = 0; i < totalClaimedBefore.length; ++i) {
                assertEq(
                    earnings[i] - totalClaimedBefore[i],
                    tokenBalancesAfter[i] - tokenBalancesBefore[i],
                    "Token balance not incremented by earnings amount"
                );
            }

            cheats.stopPrank();
        }

        // 2. Claim against first root again, expect a revert
        {
            uint32 rootIndex = claim.rootIndex;
            (bytes32 root, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
            cheats.warp(activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

            cheats.expectRevert(
                "PaymentCoordinator.processClaim: cumulativeEarnings must be gt than cumulativeClaimed"
            );
            paymentCoordinator.processClaim(claim, claimer);

            cheats.stopPrank();
        }
    }

    /// @notice Claim against latest submitted root, rootIndex 3 but modify some of the leaf values
    /// forge-config: default.fuzz.runs = 10
    function testFuzz_processClaim_Revert_WhenInvalidTokenClaim(
        bool setClaimerFor,
        address claimerFor
    ) public filterFuzzedAddressInputs(claimerFor) {
        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            paymentCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = _parseAllProofs();
        IPaymentCoordinator.PaymentMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        (, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
        cheats.warp(activatedAt);

        // Modify Earnings
        claim.tokenLeaves[0].cumulativeEarnings = 1e20;
        claim.tokenLeaves[1].cumulativeEarnings = 1e20;

        // Check claim is not valid from both checkClaim() and processClaim() throwing a revert
        cheats.startPrank(claimer);

        cheats.expectRevert("PaymentCoordinator._verifyTokenClaim: invalid token claim proof");
        assertFalse(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

        cheats.expectRevert("PaymentCoordinator._verifyTokenClaim: invalid token claim proof");
        paymentCoordinator.processClaim(claim, claimer);

        cheats.stopPrank();
    }

    /// @notice Claim against latest submitted root, rootIndex 3 but modify some of the leaf values
    /// forge-config: default.fuzz.runs = 10
    function testFuzz_processClaim_Revert_WhenInvalidEarnerClaim(
        bool setClaimerFor,
        address claimerFor,
        address invalidEarner
    ) public filterFuzzedAddressInputs(claimerFor) {
        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            paymentCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = _parseAllProofs();
        IPaymentCoordinator.PaymentMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        (, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
        cheats.warp(activatedAt);

        // Modify Earner
        claim.earnerLeaf.earner = invalidEarner;

        // Check claim is not valid from both checkClaim() and processClaim() throwing a revert
        cheats.startPrank(claimer);

        cheats.expectRevert("PaymentCoordinator._verifyEarnerClaimProof: invalid earner claim proof");
        assertFalse(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

        cheats.expectRevert("PaymentCoordinator._verifyEarnerClaimProof: invalid earner claim proof");
        paymentCoordinator.processClaim(claim, claimer);

        cheats.stopPrank();
    }

    /// @notice Claim against latest submitted root, rootIndex 3 but write to cumulativeClaimed storage.
    /// Expects a revert when calling processClaim()
    function testFuzz_processClaim_Revert_WhenCumulativeClaimedUnderflow(
        bool setClaimerFor,
        address claimerFor
    ) public filterFuzzedAddressInputs(claimerFor) {
        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            paymentCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = _parseAllProofs();
        IPaymentCoordinator.PaymentMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        (bytes32 root, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
        cheats.warp(activatedAt);

        assertTrue(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

        // Set cumulativeClaimed to be max uint256, should revert when attempting to claim
        stdstore
            .target(address(paymentCoordinator))
            .sig("cumulativeClaimed(address,address)")
            .with_key(claim.earnerLeaf.earner)
            .with_key(address(claim.tokenLeaves[0].token))
            .checked_write(type(uint256).max);
        cheats.startPrank(claimer);
        cheats.expectRevert("PaymentCoordinator.processClaim: cumulativeEarnings must be gt than cumulativeClaimed");
        paymentCoordinator.processClaim(claim, claimer);
        cheats.stopPrank();
    }

    /// @notice Claim against latest submitted root, rootIndex 3 but with larger tokenIndex used that could pass the proofs.
    /// Expects a revert as we check for this in processClaim()
    function testFuzz_processClaim_Revert_WhenTokenIndexBitwiseAddedTo(
        bool setClaimerFor,
        address claimerFor,
        uint8 numShift
    ) public filterFuzzedAddressInputs(claimerFor) {
        cheats.assume(0 < numShift && numShift < 16);
        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            paymentCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = _parseAllProofs();
        IPaymentCoordinator.PaymentMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        (bytes32 root, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
        cheats.warp(activatedAt);

        assertTrue(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

        // Take the tokenIndex and add a significant bit so that the actual index number is increased
        // but still with the same least significant bits for valid proofs
        uint8 proofLength = uint8(claim.tokenTreeProofs[0].length);
        claim.tokenIndices[0] = claim.tokenIndices[0] | uint32(1 << (numShift + proofLength / 32));
        cheats.startPrank(claimer);
        cheats.expectRevert("PaymentCoordinator._verifyTokenClaim: invalid tokenLeafIndex");
        paymentCoordinator.processClaim(claim, claimer);
        cheats.stopPrank();
    }

    /// @notice Claim against latest submitted root, rootIndex 3 but with larger earnerIndex used that could pass the proofs.
    /// Expects a revert as we check for this in processClaim()
    function testFuzz_processClaim_Revert_WhenEarnerIndexBitwiseAddedTo(
        bool setClaimerFor,
        address claimerFor,
        uint8 numShift
    ) public filterFuzzedAddressInputs(claimerFor) {
        cheats.assume(0 < numShift && numShift < 16);
        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            paymentCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = _parseAllProofs();
        IPaymentCoordinator.PaymentMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        (bytes32 root, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
        cheats.warp(activatedAt);

        assertTrue(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

        // Take the tokenIndex and add a significant bit so that the actual index number is increased
        // but still with the same least significant bits for valid proofs
        uint8 proofLength = uint8(claim.earnerTreeProof.length);
        claim.earnerIndex = claim.earnerIndex | uint32(1 << (numShift + proofLength / 32));
        cheats.startPrank(claimer);
        cheats.expectRevert("PaymentCoordinator._verifyEarnerClaimProof: invalid earnerLeafIndex");
        paymentCoordinator.processClaim(claim, claimer);
        cheats.stopPrank();
    }

    /// @notice tests with earnerIndex and tokenIndex set to max value and using alternate claim proofs
    function testFuzz_processClaim_WhenMaxEarnerIndexAndTokenIndex(
        bool setClaimerFor,
        address claimerFor,
        uint8 numShift
    ) public filterFuzzedAddressInputs(claimerFor) {
        // Hardcode earner address to earner in alternate claim proofs
        earner = 0x25A1B7322f9796B26a4Bec125913b34C292B28D6;

        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            paymentCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = _parseAllProofsMaxEarnerAndLeafIndices();
        IPaymentCoordinator.PaymentMerkleClaim memory claim = claims[0];

        // 1. Claim against first root where earner tree is full tree and earner and token index is last index of that tree height
        {
            uint32 rootIndex = claim.rootIndex;
            (bytes32 root, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
            cheats.warp(activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

            uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint256[] memory earnings = _getCumulativeEarnings(claim);
            uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            // +1 since earnerIndex is 0-indexed
            // Here the earnerIndex is 7 in a full binary tree and the number of bytes32 hash proofs is 3
            assertEq(
                claim.earnerIndex + 1,
                (1 << ((claim.earnerTreeProof.length / 32))),
                "EarnerIndex not set to max value"
            );
            // +1 since tokenIndex is 0-indexed
            // Here the tokenIndex is also 7 in a full binary tree and the number of bytes32 hash proofs is 3
            assertEq(
                claim.tokenIndices[0] + 1,
                (1 << ((claim.tokenTreeProofs[0].length / 32))),
                "TokenIndex not set to max value"
            );
            _assertPaymentClaimedEvents(root, claim, claimer);
            paymentCoordinator.processClaim(claim, claimer);

            uint256[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint256 i = 0; i < totalClaimedBefore.length; ++i) {
                assertEq(
                    earnings[i] - totalClaimedBefore[i],
                    tokenBalancesAfter[i] - tokenBalancesBefore[i],
                    "Token balance not incremented by earnings amount"
                );
            }

            cheats.stopPrank();
        }
    }

    /// @notice tests with single token leaf for the earner's subtree. tokenTreeProof for the token in the claim should be empty
    function testFuzz_processClaim_WhenSingleTokenLeaf(
        bool setClaimerFor,
        address claimerFor,
        uint8 numShift
    ) public filterFuzzedAddressInputs(claimerFor) {
        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            paymentCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = _parseAllProofsSingleTokenLeaf();
        IPaymentCoordinator.PaymentMerkleClaim memory claim = claims[0];

        // 1. Claim against first root where earner tree is full tree and earner and token index is last index of that tree height
        {
            uint32 rootIndex = claim.rootIndex;
            (bytes32 root, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
            cheats.warp(activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

            uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint256[] memory earnings = _getCumulativeEarnings(claim);
            uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            // Single tokenLeaf in earner's subtree, should be 0 index
            assertEq(
                claim.tokenIndices[0],
                0,
                "TokenIndex should be 0"
            );
            assertEq(
                claim.tokenTreeProofs[0].length,
                0,
                "TokenTreeProof should be empty"
            );
            _assertPaymentClaimedEvents(root, claim, claimer);
            paymentCoordinator.processClaim(claim, claimer);

            uint256[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint256 i = 0; i < totalClaimedBefore.length; ++i) {
                assertEq(
                    earnings[i] - totalClaimedBefore[i],
                    tokenBalancesAfter[i] - tokenBalancesBefore[i],
                    "Token balance not incremented by earnings amount"
                );
            }

            cheats.stopPrank();
        }
    }

    /// @notice tests with single earner leaf in the merkle tree. earnerTreeProof in claim should be empty
    function testFuzz_processClaim_WhenSingleEarnerLeaf(
        bool setClaimerFor,
        address claimerFor,
        uint8 numShift
    ) public filterFuzzedAddressInputs(claimerFor) {
        // Hardcode earner address to earner in alternate claim proofs
        earner = 0x0D6bA28b9919CfCDb6b233469Cc5Ce30b979e08E;

        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            paymentCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = _parseAllProofsSingleEarnerLeaf();
        IPaymentCoordinator.PaymentMerkleClaim memory claim = claims[0];

        // 1. Claim against first root where earner tree is full tree and earner and token index is last index of that tree height
        {
            uint32 rootIndex = claim.rootIndex;
            (bytes32 root, , uint32 activatedAt) = paymentCoordinator.distributionRoots(rootIndex);
            cheats.warp(activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(paymentCoordinator.checkClaim(claim), "PaymentCoordinator.checkClaim: claim not valid");

            uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint256[] memory earnings = _getCumulativeEarnings(claim);
            uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            // Earner Leaf in merkle tree, should be 0 index
            assertEq(
                claim.earnerIndex,
                0,
                "EarnerIndex should be 0"
            );
            assertEq(
                claim.earnerTreeProof.length,
                0,
                "EarnerTreeProof should be empty"
            );
            _assertPaymentClaimedEvents(root, claim, claimer);
            paymentCoordinator.processClaim(claim, claimer);

            uint256[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint256 i = 0; i < totalClaimedBefore.length; ++i) {
                assertEq(
                    earnings[i] - totalClaimedBefore[i],
                    tokenBalancesAfter[i] - tokenBalancesBefore[i],
                    "Token balance not incremented by earnings amount"
                );
            }

            cheats.stopPrank();
        }
    }

    /// @notice Set address with ERC20Mock bytecode and mint amount to paymentCoordinator for
    /// balance for testing processClaim()
    function _setAddressAsERC20(address randAddress, uint256 mintAmount) internal {
        cheats.etch(randAddress, mockTokenBytecode);
        ERC20Mock(randAddress).mint(address(paymentCoordinator), mintAmount);
    }

    /// @notice parse proofs from json file and submitRoot()
    function _parseProofData(string memory filePath) internal returns (IPaymentCoordinator.PaymentMerkleClaim memory) {
        cheats.readFile(filePath);

        string memory claimProofData = cheats.readFile(filePath);

        // Parse PaymentMerkleClaim
        merkleRoot = abi.decode(stdJson.parseRaw(claimProofData, ".Root"), (bytes32));
        earnerIndex = abi.decode(stdJson.parseRaw(claimProofData, ".EarnerIndex"), (uint32));
        earnerTreeProof = abi.decode(stdJson.parseRaw(claimProofData, ".EarnerTreeProof"), (bytes));
        proofEarner = stdJson.readAddress(claimProofData, ".EarnerLeaf.Earner");
        require(earner == proofEarner, "earner in test and json file do not match");
        earnerTokenRoot = abi.decode(stdJson.parseRaw(claimProofData, ".EarnerLeaf.EarnerTokenRoot"), (bytes32));
        uint256 numTokenLeaves = stdJson.readUint(claimProofData, ".TokenLeavesNum");
        uint256 numTokenTreeProofs = stdJson.readUint(claimProofData, ".TokenTreeProofsNum");

        IPaymentCoordinator.TokenTreeMerkleLeaf[] memory tokenLeaves = new IPaymentCoordinator.TokenTreeMerkleLeaf[](
            numTokenLeaves
        );
        uint32[] memory tokenIndices = new uint32[](numTokenLeaves);
        for (uint256 i = 0; i < numTokenLeaves; ++i) {
            string memory tokenKey = string.concat(".TokenLeaves[", cheats.toString(i), "].Token");
            string memory amountKey = string.concat(".TokenLeaves[", cheats.toString(i), "].CumulativeEarnings");
            string memory leafIndicesKey = string.concat(".LeafIndices[", cheats.toString(i), "]");

            IERC20 token = IERC20(stdJson.readAddress(claimProofData, tokenKey));
            uint256 cumulativeEarnings = stdJson.readUint(claimProofData, amountKey);
            tokenLeaves[i] = IPaymentCoordinator.TokenTreeMerkleLeaf({
                token: token,
                cumulativeEarnings: cumulativeEarnings
            });
            tokenIndices[i] = uint32(stdJson.readUint(claimProofData, leafIndicesKey));

            /// DeployCode ERC20 to Token Address
            // deployCodeTo("ERC20PresetFixedSupply.sol", address(tokenLeaves[i].token));
            _setAddressAsERC20(address(token), cumulativeEarnings);
        }
        bytes[] memory tokenTreeProofs = new bytes[](numTokenTreeProofs);
        for (uint256 i = 0; i < numTokenTreeProofs; ++i) {
            string memory tokenTreeProofKey = string.concat(".TokenTreeProofs[", cheats.toString(i), "]");
            tokenTreeProofs[i] = abi.decode(stdJson.parseRaw(claimProofData, tokenTreeProofKey), (bytes));
        }

        uint32 rootCalculationEndTimestamp = uint32(block.timestamp);
        uint32 activatedAt = uint32(block.timestamp) + paymentCoordinator.activationDelay();
        prevRootCalculationEndTimestamp = rootCalculationEndTimestamp;
        cheats.warp(activatedAt);

        uint32 rootIndex = uint32(paymentCoordinator.getDistributionRootsLength());

        cheats.prank(paymentUpdater);
        paymentCoordinator.submitRoot(merkleRoot, prevRootCalculationEndTimestamp);

        IPaymentCoordinator.PaymentMerkleClaim memory newClaim = IPaymentCoordinator.PaymentMerkleClaim({
            rootIndex: rootIndex,
            earnerIndex: earnerIndex,
            earnerTreeProof: earnerTreeProof,
            earnerLeaf: IPaymentCoordinator.EarnerTreeMerkleLeaf({earner: earner, earnerTokenRoot: earnerTokenRoot}),
            tokenIndices: tokenIndices,
            tokenTreeProofs: tokenTreeProofs,
            tokenLeaves: tokenLeaves
        });

        return newClaim;
    }

    function _parseAllProofs() internal virtual returns (IPaymentCoordinator.PaymentMerkleClaim[] memory) {
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = new IPaymentCoordinator.PaymentMerkleClaim[](3);

        claims[0] = _parseProofData("src/test/test-data/paymentCoordinator/processClaimProofs_Root1.json");
        claims[1] = _parseProofData("src/test/test-data/paymentCoordinator/processClaimProofs_Root2.json");
        claims[2] = _parseProofData("src/test/test-data/paymentCoordinator/processClaimProofs_Root3.json");

        return claims;
    }

    function _parseAllProofsMaxEarnerAndLeafIndices() internal virtual returns (IPaymentCoordinator.PaymentMerkleClaim[] memory) {
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = new IPaymentCoordinator.PaymentMerkleClaim[](1);

        claims[0] = _parseProofData("src/test/test-data/paymentCoordinator/processClaimProofs_MaxEarnerAndLeafIndices.json");

        return claims;
    }

    function _parseAllProofsSingleTokenLeaf() internal virtual returns (IPaymentCoordinator.PaymentMerkleClaim[] memory) {
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = new IPaymentCoordinator.PaymentMerkleClaim[](1);

        claims[0] = _parseProofData("src/test/test-data/paymentCoordinator/processClaimProofs_SingleTokenLeaf.json");

        return claims;
    }

    function _parseAllProofsSingleEarnerLeaf() internal virtual returns (IPaymentCoordinator.PaymentMerkleClaim[] memory) {
        IPaymentCoordinator.PaymentMerkleClaim[] memory claims = new IPaymentCoordinator.PaymentMerkleClaim[](1);

        claims[0] = _parseProofData("src/test/test-data/paymentCoordinator/processClaimProofs_SingleEarnerLeaf.json");

        return claims;
    }
}
