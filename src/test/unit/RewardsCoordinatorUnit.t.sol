// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/core/RewardsCoordinator.sol";
import "src/contracts/strategies/StrategyBase.sol";

import "src/test/events/IRewardsCoordinatorEvents.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/test/mocks/Reenterer.sol";
import "src/test/mocks/ERC20Mock.sol";

/**
 * @notice Unit testing of the RewardsCoordinator contract
 * Contracts tested: RewardsCoordinator
 * Contracts not mocked: StrategyBase, PauserRegistry
 */
contract RewardsCoordinatorUnitTests is EigenLayerUnitTestSetup, IRewardsCoordinatorEvents {
    // used for stack too deep
    struct FuzzAVSRewardsSubmission {
        address avs;
        uint256 startTimestamp;
        uint256 duration;
        uint256 amount;
    }

    // Contract under test
    RewardsCoordinator public rewardsCoordinator;
    RewardsCoordinator public rewardsCoordinatorImplementation;

    // Mocks
    IERC20 token1;
    IERC20 token2;
    IERC20 token3;
    IStrategy strategyMock1;
    IStrategy strategyMock2;
    IStrategy strategyMock3;
    StrategyBase strategyImplementation;
    uint256 mockTokenInitialSupply = 1e38 - 1;
    IRewardsCoordinator.StrategyAndMultiplier[] defaultStrategyAndMultipliers;

    // Config Variables
    /// @notice intervals(epochs) are 1 weeks
    uint32 CALCULATION_INTERVAL_SECONDS = 7 days;

    /// @notice Max duration is 5 epochs (2 weeks * 5 = 10 weeks in seconds)
    uint32 MAX_REWARDS_DURATION = 70 days;

    /// @notice Lower bound start range is ~3 months into the past, multiple of CALCULATION_INTERVAL_SECONDS
    uint32 MAX_RETROACTIVE_LENGTH = 84 days;
    /// @notice Upper bound start range is ~1 month into the future, multiple of CALCULATION_INTERVAL_SECONDS
    uint32 MAX_FUTURE_LENGTH = 28 days;
    /// @notice absolute min timestamp that a rewards can start at
    uint32 GENESIS_REWARDS_TIMESTAMP = 1712188800;

    /// @notice Lower bound start range is ~1 month into the past, multiple of CALCULATION_INTERVAL_SECONDS
    uint32 OPERATOR_SET_MAX_RETROACTIVE_LENGTH = 28 days;
    /// @notice absolute min timestamp (seconds) that an operatorSet rewards submission can start at
    uint32 OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP = 1720656000;


    /// @notice Delay in timestamp before a posted root can be claimed against
    uint32 activationDelay = 7 days;
    /// @notice The activation delay until an updated operator's commission bips takes effect
    uint32 OPERATOR_COMMISSION_ACTIVATION_DELAY = 17.5 days;
    /// @notice The maximum commission bips that can be set for an operator
    uint16 MAX_COMMISSION_BIPS = 10000;
    /// @notice the commission for all operators across all avss
    uint16 globalCommissionBips = 1000;

    IERC20[] rewardTokens;

    // RewardsCoordinator Constants

    /// @dev Index for flag that pauses calling createAVSRewardsSubmission
    uint8 internal constant PAUSED_AVS_REWARDS_SUBMISSION = 0;

    /// @dev Index for flag that pauses calling createRewardsForAllSubmission
    uint8 internal constant PAUSED_REWARDS_FOR_ALL_SUBMISSION = 1;

    /// @dev Index for flag that pauses claiming
    uint8 internal constant PAUSED_PROCESS_CLAIM = 2;

    /// @dev Index for flag that pauses submitRoots
    uint8 internal constant PAUSED_SUBMIT_DISABLE_ROOTS = 3;

    /// @dev Index for flag that pauses calling rewardOperatorSetForRange
    uint8 internal constant PAUSED_REWARD_OPERATOR_SET = 4;

    // RewardsCoordinator entities
    address rewardsUpdater = address(1000);
    address defaultAVS = address(1001);
    address defaultClaimer = address(1002);
    address rewardsForAllSubmitter = address(1003);

    function setUp() public virtual override {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        // Deploy RewardsCoordinator proxy and implementation
        rewardsCoordinatorImplementation = new RewardsCoordinator(
            delegationManagerMock,
            strategyManagerMock,
            avsDirectoryMock,
            CALCULATION_INTERVAL_SECONDS,
            MAX_REWARDS_DURATION,
            MAX_RETROACTIVE_LENGTH,
            MAX_FUTURE_LENGTH,
            GENESIS_REWARDS_TIMESTAMP,
            OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP,
            OPERATOR_SET_MAX_RETROACTIVE_LENGTH
        );
        rewardsCoordinator = RewardsCoordinator(
            address(
                new TransparentUpgradeableProxy(
                    address(rewardsCoordinatorImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        RewardsCoordinator.initialize.selector,
                        address(this), // initOwner
                        pauserRegistry,
                        0, // 0 is initialPausedStatus
                        rewardsUpdater,
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
            IRewardsCoordinator.StrategyAndMultiplier(IStrategy(address(strategies[0])), 1e18)
        );
        defaultStrategyAndMultipliers.push(
            IRewardsCoordinator.StrategyAndMultiplier(IStrategy(address(strategies[1])), 2e18)
        );
        defaultStrategyAndMultipliers.push(
            IRewardsCoordinator.StrategyAndMultiplier(IStrategy(address(strategies[2])), 3e18)
        );

        rewardsCoordinator.setRewardsForAllSubmitter(rewardsForAllSubmitter, true);
        rewardsCoordinator.setRewardsUpdater(rewardsUpdater);

        // Exclude from fuzzed tests
        addressIsExcludedFromFuzzedInputs[address(rewardsCoordinator)] = true;
        addressIsExcludedFromFuzzedInputs[address(rewardsUpdater)] = true;

        // Set the timestamp to some time after the genesis rewards timestamp
        cheats.warp(GENESIS_REWARDS_TIMESTAMP + 5 days);
    }

    /// @notice deploy token to owner and approve rewardsCoordinator. Used for deploying reward tokens
    function _deployMockRewardTokens(address owner, uint256 numTokens) internal virtual {
        cheats.startPrank(owner);
        for (uint256 i = 0; i < numTokens; ++i) {
            IERC20 token = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, owner);
            rewardTokens.push(token);
            token.approve(address(rewardsCoordinator), mockTokenInitialSupply);
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

    function _assertRewardsClaimedEvents(bytes32 root, IRewardsCoordinator.RewardsMerkleClaim memory claim, address recipient) internal {
        address earner = claim.earnerLeaf.earner;
        address claimer = rewardsCoordinator.claimerFor(earner);
        if (claimer == address(0)) {
            claimer = earner;
        }
        IERC20 token;
        uint256 claimedAmount;
        for (uint256 i = 0; i < claim.tokenLeaves.length; ++i) {
            token = claim.tokenLeaves[i].token;
            claimedAmount = rewardsCoordinator.cumulativeClaimed(earner, token);

            cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
            emit RewardsClaimed(
                root,
                earner,
                claimer,
                recipient,
                token,
                claim.tokenLeaves[i].cumulativeEarnings - claimedAmount
            );
        }
    }

    /// @notice given address and array of reward tokens, return array of cumulativeClaimed amonts
    function _getCumulativeClaimed(
        address earner,
        IRewardsCoordinator.RewardsMerkleClaim memory claim
    ) internal view returns (uint256[] memory) {
        uint256[] memory totalClaimed = new uint256[](claim.tokenLeaves.length);

        for (uint256 i = 0; i < claim.tokenLeaves.length; ++i) {
            totalClaimed[i] = rewardsCoordinator.cumulativeClaimed(earner, claim.tokenLeaves[i].token);
        }

        return totalClaimed;
    }

    /// @notice given a claim, return the new cumulativeEarnings for each token
    function _getCumulativeEarnings(
        IRewardsCoordinator.RewardsMerkleClaim memory claim
    ) internal pure returns (uint256[] memory) {
        uint256[] memory earnings = new uint256[](claim.tokenLeaves.length);

        for (uint256 i = 0; i < claim.tokenLeaves.length; ++i) {
            earnings[i] = claim.tokenLeaves[i].cumulativeEarnings;
        }

        return earnings;
    }

    function _getClaimTokenBalances(
        address earner,
        IRewardsCoordinator.RewardsMerkleClaim memory claim
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

contract RewardsCoordinatorUnitTests_initializeAndSetters is RewardsCoordinatorUnitTests {
    function testFuzz_setClaimerFor(address earner, address claimer) public filterFuzzedAddressInputs(earner) {
        cheats.startPrank(earner);
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit ClaimerForSet(earner, rewardsCoordinator.claimerFor(earner), claimer);
        rewardsCoordinator.setClaimerFor(claimer);
        assertEq(claimer, rewardsCoordinator.claimerFor(earner), "claimerFor not set");
        cheats.stopPrank();
    }

    function testFuzz_setActivationDelay(uint32 activationDelay) public {
        cheats.startPrank(rewardsCoordinator.owner());
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit ActivationDelaySet(rewardsCoordinator.activationDelay(), activationDelay);
        rewardsCoordinator.setActivationDelay(activationDelay);
        assertEq(activationDelay, rewardsCoordinator.activationDelay(), "activationDelay not set");
        cheats.stopPrank();
    }

    function testFuzz_setActivationDelay_Revert_WhenNotOwner(
        address caller,
        uint32 activationDelay
    ) public filterFuzzedAddressInputs(caller) {
        cheats.assume(caller != rewardsCoordinator.owner());
        cheats.prank(caller);
        cheats.expectRevert("Ownable: caller is not the owner");
        rewardsCoordinator.setActivationDelay(activationDelay);
    }

    function testFuzz_setGlobalOperatorCommission(uint16 globalCommissionBips) public {
        cheats.startPrank(rewardsCoordinator.owner());
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit GlobalCommissionBipsSet(rewardsCoordinator.globalOperatorCommissionBips(), globalCommissionBips);
        rewardsCoordinator.setGlobalOperatorCommission(globalCommissionBips);
        assertEq(
            globalCommissionBips,
            rewardsCoordinator.globalOperatorCommissionBips(),
            "globalOperatorCommissionBips not set"
        );
        cheats.stopPrank();
    }

    function testFuzz_setGlobalOperatorCommission_Revert_WhenNotOwner(
        address caller,
        uint16 globalCommissionBips
    ) public filterFuzzedAddressInputs(caller) {
        cheats.assume(caller != rewardsCoordinator.owner());
        cheats.prank(caller);
        cheats.expectRevert("Ownable: caller is not the owner");
        rewardsCoordinator.setGlobalOperatorCommission(globalCommissionBips);
    }

    function testFuzz_setRewardsUpdater(address newRewardsUpdater) public {
        cheats.startPrank(rewardsCoordinator.owner());
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit RewardsUpdaterSet(rewardsCoordinator.rewardsUpdater(), newRewardsUpdater);
        rewardsCoordinator.setRewardsUpdater(newRewardsUpdater);
        assertEq(newRewardsUpdater, rewardsCoordinator.rewardsUpdater(), "rewardsUpdater not set");
        cheats.stopPrank();
    }

    function testFuzz_setRewardsUpdater_Revert_WhenNotOwner(
        address caller,
        address newRewardsUpdater
    ) public filterFuzzedAddressInputs(caller) {
        cheats.assume(caller != rewardsCoordinator.owner());
        cheats.prank(caller);
        cheats.expectRevert("Ownable: caller is not the owner");
        rewardsCoordinator.setRewardsUpdater(newRewardsUpdater);
    }

    function testFuzz_setRewardsForAllSubmitter(address submitter, bool newValue) public {
        cheats.startPrank(rewardsCoordinator.owner());
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit RewardsForAllSubmitterSet(submitter, rewardsCoordinator.isRewardsForAllSubmitter(submitter), newValue);
        rewardsCoordinator.setRewardsForAllSubmitter(submitter, newValue);
        assertEq(
            newValue,
            rewardsCoordinator.isRewardsForAllSubmitter(submitter),
            "isRewardsForAllSubmitter not set"
        );
        cheats.stopPrank();
    }

    function testFuzz_setRewardsForAllSubmitter_Revert_WhenNotOwner(
        address caller,
        address submitter,
        bool newValue
    ) public filterFuzzedAddressInputs(caller) {
        cheats.assume(caller != rewardsCoordinator.owner());
        cheats.prank(caller);
        cheats.expectRevert("Ownable: caller is not the owner");
        rewardsCoordinator.setRewardsForAllSubmitter(submitter, newValue);
    }
}

contract RewardsCoordinatorUnitTests_createAVSRewardsSubmission is RewardsCoordinatorUnitTests {
    // Revert when paused
    function test_Revert_WhenPaused() public {
        cheats.prank(pauser);
        rewardsCoordinator.pause(2 ** PAUSED_AVS_REWARDS_SUBMISSION);

        cheats.expectRevert("Pausable: index is paused");
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions;
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert from reentrancy
    function test_Revert_WhenReentrancy(uint256 amount) public {
        amount = bound(amount, 1, 1e38-1);
        Reenterer reenterer = new Reenterer();

        reenterer.prepareReturnData(abi.encode(amount));

        address targetToUse = address(rewardsCoordinator);
        uint256 msgValueToUse = 0;

        _deployMockRewardTokens(address(this), 1);

        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(reenterer)),
            amount: amount,
            startTimestamp: uint32(block.timestamp),
            duration: 0
        });

        bytes memory calldataToUse = abi.encodeWithSelector(RewardsCoordinator.createAVSRewardsSubmission.selector, rewardsSubmissions);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        cheats.expectRevert();
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert with 0 length strats and multipliers
    function testFuzz_Revert_WhenEmptyStratsAndMultipliers(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        IRewardsCoordinator.StrategyAndMultiplier[] memory emptyStratsAndMultipliers;
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: emptyStratsAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: no strategies set");
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
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
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", amount, avs);
        duration = bound(duration, 0, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. Call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: amount too large");
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    function testFuzz_Revert_WhenDuplicateStrategies(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        IRewardsCoordinator.StrategyAndMultiplier[]
            memory dupStratsAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](2);
        dupStratsAndMultipliers[0] = defaultStrategyAndMultipliers[0];
        dupStratsAndMultipliers[1] = defaultStrategyAndMultipliers[0];
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: dupStratsAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(
            "RewardsCoordinator._validateRewardsSubmission: strategies must be in ascending order to handle duplicates"
        );
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert with exceeding max duration
    function testFuzz_Revert_WhenExceedingMaxDuration(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, MAX_REWARDS_DURATION + 1, type(uint32).max);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: duration exceeds MAX_REWARDS_DURATION");
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert with invalid interval seconds
    function testFuzz_Revert_WhenInvalidIntervalSeconds(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_REWARDS_DURATION);
        cheats.assume(duration % CALCULATION_INTERVAL_SECONDS != 0);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(
            "RewardsCoordinator._validateRewardsSubmission: duration must be a multiple of CALCULATION_INTERVAL_SECONDS"
        );
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert with retroactive rewards enabled and set too far in past
    // - either before genesis rewards timestamp
    // - before max retroactive length
    function testFuzz_Revert_WhenRewardsSubmissionTooStale(
        uint256 fuzzBlockTimestamp,
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        fuzzBlockTimestamp = bound(fuzzBlockTimestamp, uint256(MAX_RETROACTIVE_LENGTH), block.timestamp);
        cheats.warp(fuzzBlockTimestamp);

        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            0,
            uint256(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: startTimestamp too far in the past");
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert with start timestamp past max future length
    function testFuzz_Revert_WhenRewardsSubmissionTooFarInFuture(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            block.timestamp + uint256(MAX_FUTURE_LENGTH) + 1 + CALCULATION_INTERVAL_SECONDS,
            type(uint32).max
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: startTimestamp too far in the future");
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert with non whitelisted strategy
    function testFuzz_Revert_WhenInvalidStrategy(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        defaultStrategyAndMultipliers[0].strategy = IStrategy(address(999));
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected event emitted
        cheats.prank(avs);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: invalid strategy considered");
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    /**
     * @notice test a single rewards submission asserting for the following
     * - correct event emitted
     * - submission nonce incrementation by 1, and rewards submission hash being set in storage.
     * - rewards submission hash being set in storage
     * - token balance before and after of avs and rewardsCoordinator
     */
    function testFuzz_createAVSRewardsSubmission_SingleSubmission(
        address avs,
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected event emitted
        uint256 avsBalanceBefore = rewardToken.balanceOf(avs);
        uint256 rewardsCoordinatorBalanceBefore = rewardToken.balanceOf(address(rewardsCoordinator));

        cheats.startPrank(avs);
        rewardToken.approve(address(rewardsCoordinator), amount);
        uint256 currSubmissionNonce = rewardsCoordinator.submissionNonce(avs);
        bytes32 rewardsSubmissionHash = keccak256(abi.encode(avs, currSubmissionNonce, rewardsSubmissions[0]));

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit AVSRewardsSubmissionCreated(avs, currSubmissionNonce, rewardsSubmissionHash, rewardsSubmissions[0]);
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
        cheats.stopPrank();

        assertTrue(rewardsCoordinator.isAVSRewardsSubmissionHash(avs, rewardsSubmissionHash), "rewards submission hash not submitted");
        assertEq(currSubmissionNonce + 1, rewardsCoordinator.submissionNonce(avs), "submission nonce not incremented");
        assertEq(
            avsBalanceBefore - amount,
            rewardToken.balanceOf(avs),
            "AVS balance not decremented by amount of rewards submission"
        );
        assertEq(
            rewardsCoordinatorBalanceBefore + amount,
            rewardToken.balanceOf(address(rewardsCoordinator)),
            "RewardsCoordinator balance not incremented by amount of rewards submission"
        );
    }

    /**
     * @notice test multiple rewards submissions asserting for the following
     * - correct event emitted
     * - submission nonce incrementation by numSubmissions, and rewards submission hashes being set in storage.
     * - rewards submission hash being set in storage
     * - token balances before and after of avs and rewardsCoordinator
     */
    function testFuzz_createAVSRewardsSubmission_MultipleSubmissions(
        FuzzAVSRewardsSubmission memory param,
        uint256 numSubmissions
    ) public filterFuzzedAddressInputs(param.avs) {
        cheats.assume(2 <= numSubmissions && numSubmissions <= 10);
        cheats.assume(param.avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](numSubmissions);
        bytes32[] memory rewardsSubmissionHashes = new bytes32[](numSubmissions);
        uint256 startSubmissionNonce = rewardsCoordinator.submissionNonce(param.avs);
        _deployMockRewardTokens(param.avs, numSubmissions);

        uint256[] memory avsBalancesBefore = _getBalanceForTokens(rewardTokens, param.avs);
        uint256[] memory rewardsCoordinatorBalancesBefore = _getBalanceForTokens(
            rewardTokens,
            address(rewardsCoordinator)
        );
        uint256[] memory amounts = new uint256[](numSubmissions);

        // Create multiple rewards submissions and their expected event
        for (uint256 i = 0; i < numSubmissions; ++i) {
            // 1. Bound fuzz inputs to valid ranges and amounts using randSeed for each
            param.amount = bound(param.amount + i, 1, mockTokenInitialSupply);
            amounts[i] = param.amount;
            param.duration = bound(param.duration + i, 0, MAX_REWARDS_DURATION);
            param.duration = param.duration - (param.duration % CALCULATION_INTERVAL_SECONDS);
            param.startTimestamp = bound(
                param.startTimestamp + i,
                uint256(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                    CALCULATION_INTERVAL_SECONDS -
                    1,
                block.timestamp + uint256(MAX_FUTURE_LENGTH)
            );
            param.startTimestamp = param.startTimestamp - (param.startTimestamp % CALCULATION_INTERVAL_SECONDS);

            // 2. Create rewards submission input param
            IRewardsCoordinator.RewardsSubmission memory rewardsSubmission = IRewardsCoordinator.RewardsSubmission({
                strategiesAndMultipliers: defaultStrategyAndMultipliers,
                token: rewardTokens[i],
                amount: amounts[i],
                startTimestamp: uint32(param.startTimestamp),
                duration: uint32(param.duration)
            });
            rewardsSubmissions[i] = rewardsSubmission;

            // 3. expected event emitted for this rewardsSubmission
            rewardsSubmissionHashes[i] = keccak256(abi.encode(param.avs, startSubmissionNonce + i, rewardsSubmissions[i]));
            cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
            emit AVSRewardsSubmissionCreated(param.avs, startSubmissionNonce + i, rewardsSubmissionHashes[i], rewardsSubmissions[i]);
        }

        // 4. call createAVSRewardsSubmission()
        cheats.prank(param.avs);
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);

        // 5. Check for submissionNonce() and rewardsSubmissionHashes being set
        assertEq(
            startSubmissionNonce + numSubmissions,
            rewardsCoordinator.submissionNonce(param.avs),
            "submission nonce not incremented properly"
        );

        for (uint256 i = 0; i < numSubmissions; ++i) {
            assertTrue(
                rewardsCoordinator.isAVSRewardsSubmissionHash(param.avs, rewardsSubmissionHashes[i]),
                "rewards submission hash not submitted"
            );
            assertEq(
                avsBalancesBefore[i] - amounts[i],
                rewardTokens[i].balanceOf(param.avs),
                "AVS balance not decremented by amount of rewards submission"
            );
            assertEq(
                rewardsCoordinatorBalancesBefore[i] + amounts[i],
                rewardTokens[i].balanceOf(address(rewardsCoordinator)),
                "RewardsCoordinator balance not incremented by amount of rewards submission"
            );
        }
    }
}

contract RewardsCoordinatorUnitTests_createRewardsForAllSubmission is RewardsCoordinatorUnitTests {
    // Revert when paused
    function test_Revert_WhenPaused() public {
        cheats.prank(pauser);
        rewardsCoordinator.pause(2 ** PAUSED_REWARDS_FOR_ALL_SUBMISSION);

        cheats.expectRevert("Pausable: index is paused");
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions;
        rewardsCoordinator.createRewardsForAllSubmission(rewardsSubmissions);
    }

    // Revert from reentrancy
    function test_Revert_WhenReentrancy(uint256 amount) public {
        Reenterer reenterer = new Reenterer();

        reenterer.prepareReturnData(abi.encode(amount));

        address targetToUse = address(rewardsCoordinator);
        uint256 msgValueToUse = 0;

        _deployMockRewardTokens(address(this), 1);

        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(reenterer)),
            amount: amount,
            startTimestamp: uint32(block.timestamp),
            duration: 0
        });

        bytes memory calldataToUse = abi.encodeWithSelector(RewardsCoordinator.createAVSRewardsSubmission.selector, rewardsSubmissions);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        cheats.prank(rewardsForAllSubmitter);
        cheats.expectRevert();
        rewardsCoordinator.createRewardsForAllSubmission(rewardsSubmissions);
    }

    function testFuzz_Revert_WhenNotRewardsForAllSubmitter(
        address invalidSubmitter
    ) public filterFuzzedAddressInputs(invalidSubmitter) {
        cheats.assume(invalidSubmitter != rewardsForAllSubmitter);

        cheats.expectRevert("RewardsCoordinator: caller is not a valid createRewardsForAllSubmission submitter");
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions;
        rewardsCoordinator.createRewardsForAllSubmission(rewardsSubmissions);
    }

    /**
     * @notice test a single rewards submission asserting for the following
     * - correct event emitted
     * - submission nonce incrementation by 1, and rewards submission hash being set in storage.
     * - rewards submission hash being set in storage
     * - token balance before and after of RewardsForAllSubmitter and rewardsCoordinator
     */
    function testFuzz_createRewardsForAllSubmission_SingleSubmission(uint256 startTimestamp, uint256 duration, uint256 amount) public {
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply(
            "dog wif hat",
            "MOCK1",
            mockTokenInitialSupply,
            rewardsForAllSubmitter
        );
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected event emitted
        uint256 submitterBalanceBefore = rewardToken.balanceOf(rewardsForAllSubmitter);
        uint256 rewardsCoordinatorBalanceBefore = rewardToken.balanceOf(address(rewardsCoordinator));

        cheats.startPrank(rewardsForAllSubmitter);
        rewardToken.approve(address(rewardsCoordinator), amount);
        uint256 currSubmissionNonce = rewardsCoordinator.submissionNonce(rewardsForAllSubmitter);
        bytes32 rewardsSubmissionHash = keccak256(abi.encode(rewardsForAllSubmitter, currSubmissionNonce, rewardsSubmissions[0]));

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit RewardsSubmissionForAllCreated(rewardsForAllSubmitter, currSubmissionNonce, rewardsSubmissionHash, rewardsSubmissions[0]);
        rewardsCoordinator.createRewardsForAllSubmission(rewardsSubmissions);
        cheats.stopPrank();

        assertTrue(
            rewardsCoordinator.isRewardsSubmissionForAllHash(rewardsForAllSubmitter, rewardsSubmissionHash),
            "rewards submission hash not submitted"
        );
        assertEq(
            currSubmissionNonce + 1,
            rewardsCoordinator.submissionNonce(rewardsForAllSubmitter),
            "submission nonce not incremented"
        );
        assertEq(
            submitterBalanceBefore - amount,
            rewardToken.balanceOf(rewardsForAllSubmitter),
            "createRewardsForAllSubmission Submitter balance not decremented by amount of rewards submission"
        );
        assertEq(
            rewardsCoordinatorBalanceBefore + amount,
            rewardToken.balanceOf(address(rewardsCoordinator)),
            "RewardsCoordinator balance not incremented by amount of rewards submission"
        );
    }

    /**
     * @notice test multiple rewards submissions asserting for the following
     * - correct event emitted
     * - submission nonce incrementation by numSubmissions, and rewards submission hashes being set in storage.
     * - rewards submission hash being set in storage
     * - token balances before and after of createRewardsForAllSubmission submitter and rewardsCoordinator
     */
    function testFuzz_createRewardsForAllSubmission_MultipleSubmissions(FuzzAVSRewardsSubmission memory param, uint256 numSubmissions) public {
        cheats.assume(2 <= numSubmissions && numSubmissions <= 10);
        cheats.prank(rewardsCoordinator.owner());

        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](numSubmissions);
        bytes32[] memory rewardsSubmissionHashes = new bytes32[](numSubmissions);
        uint256 startSubmissionNonce = rewardsCoordinator.submissionNonce(rewardsForAllSubmitter);
        _deployMockRewardTokens(rewardsForAllSubmitter, numSubmissions);

        uint256[] memory submitterBalancesBefore = _getBalanceForTokens(rewardTokens, rewardsForAllSubmitter);
        uint256[] memory rewardsCoordinatorBalancesBefore = _getBalanceForTokens(
            rewardTokens,
            address(rewardsCoordinator)
        );
        uint256[] memory amounts = new uint256[](numSubmissions);

        // Create multiple rewards submissions and their expected event
        for (uint256 i = 0; i < numSubmissions; ++i) {
            // 1. Bound fuzz inputs to valid ranges and amounts using randSeed for each
            param.amount = bound(param.amount + i, 1, mockTokenInitialSupply);
            amounts[i] = param.amount;
            param.duration = bound(param.duration + i, 0, MAX_REWARDS_DURATION);
            param.duration = param.duration - (param.duration % CALCULATION_INTERVAL_SECONDS);
            param.startTimestamp = bound(
                param.startTimestamp + i,
                uint256(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) +
                    CALCULATION_INTERVAL_SECONDS -
                    1,
                block.timestamp + uint256(MAX_FUTURE_LENGTH)
            );
            param.startTimestamp = param.startTimestamp - (param.startTimestamp % CALCULATION_INTERVAL_SECONDS);

            // 2. Create rewards submission input param
            IRewardsCoordinator.RewardsSubmission memory rewardsSubmission = IRewardsCoordinator.RewardsSubmission({
                strategiesAndMultipliers: defaultStrategyAndMultipliers,
                token: rewardTokens[i],
                amount: amounts[i],
                startTimestamp: uint32(param.startTimestamp),
                duration: uint32(param.duration)
            });
            rewardsSubmissions[i] = rewardsSubmission;

            // 3. expected event emitted for this rewardsSubmission
            rewardsSubmissionHashes[i] = keccak256(abi.encode(rewardsForAllSubmitter, startSubmissionNonce + i, rewardsSubmissions[i]));
            cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
            emit RewardsSubmissionForAllCreated(
                rewardsForAllSubmitter,
                startSubmissionNonce + i,
                rewardsSubmissionHashes[i],
                rewardsSubmissions[i]
            );
        }

        // 4. call createAVSRewardsSubmission()
        cheats.prank(rewardsForAllSubmitter);
        rewardsCoordinator.createRewardsForAllSubmission(rewardsSubmissions);

        // 5. Check for submissionNonce() and rewardsSubmissionHashes being set
        assertEq(
            startSubmissionNonce + numSubmissions,
            rewardsCoordinator.submissionNonce(rewardsForAllSubmitter),
            "submission nonce not incremented properly"
        );

        for (uint256 i = 0; i < numSubmissions; ++i) {
            assertTrue(
                rewardsCoordinator.isRewardsSubmissionForAllHash(rewardsForAllSubmitter, rewardsSubmissionHashes[i]),
                "rewards submission hash not submitted"
            );
            assertEq(
                submitterBalancesBefore[i] - amounts[i],
                rewardTokens[i].balanceOf(rewardsForAllSubmitter),
                "RewardsForAllSubmitter Submitter balance not decremented by amount of rewards submission"
            );
            assertEq(
                rewardsCoordinatorBalancesBefore[i] + amounts[i],
                rewardTokens[i].balanceOf(address(rewardsCoordinator)),
                "RewardsCoordinator balance not incremented by amount of rewards submission"
            );
        }
    }
}

contract RewardsCoordinatorUnitTests_submitRoot is RewardsCoordinatorUnitTests {
    // only callable by rewardsUpdater
    function testFuzz_Revert_WhenNotRewardsUpdater(
        address invalidRewardsUpdater
    ) public filterFuzzedAddressInputs(invalidRewardsUpdater) {
        cheats.prank(invalidRewardsUpdater);

        cheats.expectRevert("RewardsCoordinator: caller is not the rewardsUpdater");
        rewardsCoordinator.submitRoot(bytes32(0), 0);
    }

    function test_Revert_WhenSubmitRootPaused() public {
        cheats.prank(pauser);
        rewardsCoordinator.pause(2 ** PAUSED_SUBMIT_DISABLE_ROOTS);

        cheats.expectRevert("Pausable: index is paused");
        rewardsCoordinator.submitRoot(bytes32(0), 0);
    }

    /// @notice submits root with correct values and adds to root storage array
    /// - checks activatedAt has added activationDelay
    function testFuzz_submitRoot(bytes32 root, uint32 rewardsCalculationEndTimestamp) public {
        // fuzz avoiding overflows and valid activatedAt values
        cheats.assume(rewardsCoordinator.currRewardsCalculationEndTimestamp() < rewardsCalculationEndTimestamp);
        cheats.assume(rewardsCalculationEndTimestamp < block.timestamp);

        uint32 expectedRootIndex = uint32(rewardsCoordinator.getDistributionRootsLength());
        uint32 activatedAt = uint32(block.timestamp) + rewardsCoordinator.activationDelay();

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit DistributionRootSubmitted(expectedRootIndex, root, rewardsCalculationEndTimestamp, activatedAt);
        cheats.prank(rewardsUpdater);
        rewardsCoordinator.submitRoot(root, rewardsCalculationEndTimestamp);

        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(expectedRootIndex);

        assertEq(
            expectedRootIndex,
            rewardsCoordinator.getDistributionRootsLength() - 1,
            "root not added to roots array"
        );
        assertEq(
            root,
            rewardsCoordinator.getCurrentDistributionRoot().root,
            "getCurrentDistributionRoot view function failed"
        );
        assertEq(
            root,
            rewardsCoordinator.getDistributionRootAtIndex(expectedRootIndex).root,
            "getDistributionRootAtIndex view function failed"
        );
        assertEq(activatedAt, distributionRoot.activatedAt, "activatedAt not correct");
        assertEq(root, distributionRoot.root, "root not set");
        assertEq(
            rewardsCalculationEndTimestamp,
            distributionRoot.rewardsCalculationEndTimestamp,
            "rewardsCalculationEndTimestamp not set"
        );
        assertEq(
            rewardsCoordinator.currRewardsCalculationEndTimestamp(),
            rewardsCalculationEndTimestamp,
            "currRewardsCalculationEndTimestamp not set"
        );
    }

    /// @notice Submits multiple roots and checks root index from hash is correct
    function testFuzz_getRootIndexFromHash(bytes32 root, uint16 numRoots, uint256 index) public {
        numRoots = uint16(bound(numRoots, 1, 100));
        index = bound(index, 0, uint256(numRoots - 1));

        bytes32[] memory roots = new bytes32[](numRoots);
        cheats.startPrank(rewardsUpdater);
        for (uint16 i = 0; i < numRoots; ++i) {
            roots[i] = keccak256(abi.encodePacked(root, i));

            uint32 activationDelay = uint32(block.timestamp) + rewardsCoordinator.activationDelay();
            rewardsCoordinator.submitRoot(roots[i], uint32(block.timestamp - 1));
            cheats.warp(activationDelay);
        }
        cheats.stopPrank();

        assertEq(index, rewardsCoordinator.getRootIndexFromHash(roots[index]), "root index not found");
    }
}

/// @notice Tests for sets of JSON data with different distribution roots
contract RewardsCoordinatorUnitTests_processClaim is RewardsCoordinatorUnitTests {
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
        RewardsCoordinatorUnitTests.setUp();

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
            rewardsCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = _parseAllProofs();
        IRewardsCoordinator.RewardsMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        // Claim against root and check balances before/after, and check it matches the difference between
        // cumulative claimed and earned.
        cheats.startPrank(claimer);
        assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
        uint256[] memory earnings = _getCumulativeEarnings(claim);
        uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

        _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
        rewardsCoordinator.processClaim(claim, claimer);

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
            rewardsCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = _parseAllProofs();
        IRewardsCoordinator.RewardsMerkleClaim memory claim = claims[0];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        // Claim against root and check balances before/after, and check it matches the difference between
        // cumulative claimed and earned.
        cheats.startPrank(claimer);
        assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
        uint256[] memory earnings = _getCumulativeEarnings(claim);
        uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

        _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
        rewardsCoordinator.processClaim(claim, claimer);

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
            rewardsCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = _parseAllProofs();
        IRewardsCoordinator.RewardsMerkleClaim memory claim = claims[0];

        // 1. Claim against first root
        {
            uint32 rootIndex = claim.rootIndex;
            IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
            cheats.warp(distributionRoot.activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

            uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint256[] memory earnings = _getCumulativeEarnings(claim);
            uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

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
            IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
            cheats.warp(distributionRoot.activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

            uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint256[] memory earnings = _getCumulativeEarnings(claim);
            uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

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
            IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
            cheats.warp(distributionRoot.activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

            uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint256[] memory earnings = _getCumulativeEarnings(claim);
            uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

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

    function testFuzz_processClaim_Revert_WhenRootDisabled(
        bool setClaimerFor,
        address claimerFor,
        bytes32 merkleRoot
    ) public filterFuzzedAddressInputs(claimerFor) {
        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            rewardsCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Submit a root and disable it
        cheats.startPrank(rewardsUpdater);
        rewardsCoordinator.submitRoot(merkleRoot, 1);
        uint32 rootIndex = 0;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        rewardsCoordinator.disableRoot(rootIndex);
        cheats.stopPrank();
        
        cheats.warp(distributionRoot.activatedAt);

        cheats.startPrank(claimer);
        // rootIndex in claim is 0, which is disabled
        IRewardsCoordinator.RewardsMerkleClaim memory claim;
        cheats.expectRevert("RewardsCoordinator._checkClaim: root is disabled");
        rewardsCoordinator.processClaim(claim, claimer);
        cheats.stopPrank();
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
            rewardsCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = _parseAllProofs();
        IRewardsCoordinator.RewardsMerkleClaim memory claim = claims[0];

        // 1. Claim against first root
        {
            uint32 rootIndex = claim.rootIndex;
            IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
            cheats.warp(distributionRoot.activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

            uint256[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint256[] memory earnings = _getCumulativeEarnings(claim);
            uint256[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

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
            IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
            cheats.warp(distributionRoot.activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

            cheats.expectRevert(
                "RewardsCoordinator.processClaim: cumulativeEarnings must be gt than cumulativeClaimed"
            );
            rewardsCoordinator.processClaim(claim, claimer);

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
            rewardsCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = _parseAllProofs();
        IRewardsCoordinator.RewardsMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        // Modify Earnings
        claim.tokenLeaves[0].cumulativeEarnings = 1e20;
        claim.tokenLeaves[1].cumulativeEarnings = 1e20;

        // Check claim is not valid from both checkClaim() and processClaim() throwing a revert
        cheats.startPrank(claimer);

        cheats.expectRevert("RewardsCoordinator._verifyTokenClaim: invalid token claim proof");
        assertFalse(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        cheats.expectRevert("RewardsCoordinator._verifyTokenClaim: invalid token claim proof");
        rewardsCoordinator.processClaim(claim, claimer);

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
            rewardsCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = _parseAllProofs();
        IRewardsCoordinator.RewardsMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        // Modify Earner
        claim.earnerLeaf.earner = invalidEarner;

        // Check claim is not valid from both checkClaim() and processClaim() throwing a revert
        cheats.startPrank(claimer);

        cheats.expectRevert("RewardsCoordinator._verifyEarnerClaimProof: invalid earner claim proof");
        assertFalse(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        cheats.expectRevert("RewardsCoordinator._verifyEarnerClaimProof: invalid earner claim proof");
        rewardsCoordinator.processClaim(claim, claimer);

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
            rewardsCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = _parseAllProofs();
        IRewardsCoordinator.RewardsMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        // Set cumulativeClaimed to be max uint256, should revert when attempting to claim
        stdstore
            .target(address(rewardsCoordinator))
            .sig("cumulativeClaimed(address,address)")
            .with_key(claim.earnerLeaf.earner)
            .with_key(address(claim.tokenLeaves[0].token))
            .checked_write(type(uint256).max);
        cheats.startPrank(claimer);
        cheats.expectRevert("RewardsCoordinator.processClaim: cumulativeEarnings must be gt than cumulativeClaimed");
        rewardsCoordinator.processClaim(claim, claimer);
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
            rewardsCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = _parseAllProofs();
        IRewardsCoordinator.RewardsMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        // Take the tokenIndex and add a significant bit so that the actual index number is increased
        // but still with the same least significant bits for valid proofs
        uint8 proofLength = uint8(claim.tokenTreeProofs[0].length);
        claim.tokenIndices[0] = claim.tokenIndices[0] | uint32(1 << (numShift + proofLength / 32));
        cheats.startPrank(claimer);
        cheats.expectRevert("RewardsCoordinator._verifyTokenClaim: invalid tokenLeafIndex");
        rewardsCoordinator.processClaim(claim, claimer);
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
            rewardsCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = _parseAllProofs();
        IRewardsCoordinator.RewardsMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        // Take the tokenIndex and add a significant bit so that the actual index number is increased
        // but still with the same least significant bits for valid proofs
        uint8 proofLength = uint8(claim.earnerTreeProof.length);
        claim.earnerIndex = claim.earnerIndex | uint32(1 << (numShift + proofLength / 32));
        cheats.startPrank(claimer);
        cheats.expectRevert("RewardsCoordinator._verifyEarnerClaimProof: invalid earnerLeafIndex");
        rewardsCoordinator.processClaim(claim, claimer);
        cheats.stopPrank();
    }

    /// @notice tests with earnerIndex and tokenIndex set to max value and using alternate claim proofs
    function testFuzz_processClaim_WhenMaxEarnerIndexAndTokenIndex(
        bool setClaimerFor,
        address claimerFor
    ) public filterFuzzedAddressInputs(claimerFor) {
        // Hardcode earner address to earner in alternate claim proofs
        earner = 0x25A1B7322f9796B26a4Bec125913b34C292B28D6;

        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            rewardsCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = _parseAllProofsMaxEarnerAndLeafIndices();
        IRewardsCoordinator.RewardsMerkleClaim memory claim = claims[0];

        // 1. Claim against first root where earner tree is full tree and earner and token index is last index of that tree height
        {
            uint32 rootIndex = claim.rootIndex;
            IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
            cheats.warp(distributionRoot.activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

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
            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

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
        address claimerFor
    ) public filterFuzzedAddressInputs(claimerFor) {
        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            rewardsCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = _parseAllProofsSingleTokenLeaf();
        IRewardsCoordinator.RewardsMerkleClaim memory claim = claims[0];

        // 1. Claim against first root where earner tree is full tree and earner and token index is last index of that tree height
        {
            uint32 rootIndex = claim.rootIndex;
            IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
            cheats.warp(distributionRoot.activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

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
            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

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
        address claimerFor
    ) public filterFuzzedAddressInputs(claimerFor) {
        // Hardcode earner address to earner in alternate claim proofs
        earner = 0x0D6bA28b9919CfCDb6b233469Cc5Ce30b979e08E;

        // if setClaimerFor is true, set the earners claimer to the fuzzed address
        address claimer;
        if (setClaimerFor) {
            cheats.prank(earner);
            rewardsCoordinator.setClaimerFor(claimerFor);
            claimer = claimerFor;
        } else {
            claimer = earner;
        }

        // Parse all 3 claim proofs for distributionRoots 0,1,2 respectively
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = _parseAllProofsSingleEarnerLeaf();
        IRewardsCoordinator.RewardsMerkleClaim memory claim = claims[0];

        // 1. Claim against first root where earner tree is full tree and earner and token index is last index of that tree height
        {
            uint32 rootIndex = claim.rootIndex;
            IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
            cheats.warp(distributionRoot.activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

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
            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

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

    /// @notice Set address with ERC20Mock bytecode and mint amount to rewardsCoordinator for
    /// balance for testing processClaim()
    function _setAddressAsERC20(address randAddress, uint256 mintAmount) internal {
        cheats.etch(randAddress, mockTokenBytecode);
        ERC20Mock(randAddress).mint(address(rewardsCoordinator), mintAmount);
    }

    /// @notice parse proofs from json file and submitRoot()
    function _parseProofData(string memory filePath) internal returns (IRewardsCoordinator.RewardsMerkleClaim memory) {
        cheats.readFile(filePath);

        string memory claimProofData = cheats.readFile(filePath);

        // Parse RewardsMerkleClaim
        merkleRoot = abi.decode(stdJson.parseRaw(claimProofData, ".Root"), (bytes32));
        earnerIndex = abi.decode(stdJson.parseRaw(claimProofData, ".EarnerIndex"), (uint32));
        earnerTreeProof = abi.decode(stdJson.parseRaw(claimProofData, ".EarnerTreeProof"), (bytes));
        proofEarner = stdJson.readAddress(claimProofData, ".EarnerLeaf.Earner");
        require(earner == proofEarner, "earner in test and json file do not match");
        earnerTokenRoot = abi.decode(stdJson.parseRaw(claimProofData, ".EarnerLeaf.EarnerTokenRoot"), (bytes32));
        uint256 numTokenLeaves = stdJson.readUint(claimProofData, ".TokenLeavesNum");
        uint256 numTokenTreeProofs = stdJson.readUint(claimProofData, ".TokenTreeProofsNum");

        IRewardsCoordinator.TokenTreeMerkleLeaf[] memory tokenLeaves = new IRewardsCoordinator.TokenTreeMerkleLeaf[](
            numTokenLeaves
        );
        uint32[] memory tokenIndices = new uint32[](numTokenLeaves);
        for (uint256 i = 0; i < numTokenLeaves; ++i) {
            string memory tokenKey = string.concat(".TokenLeaves[", cheats.toString(i), "].Token");
            string memory amountKey = string.concat(".TokenLeaves[", cheats.toString(i), "].CumulativeEarnings");
            string memory leafIndicesKey = string.concat(".LeafIndices[", cheats.toString(i), "]");

            IERC20 token = IERC20(stdJson.readAddress(claimProofData, tokenKey));
            uint256 cumulativeEarnings = stdJson.readUint(claimProofData, amountKey);
            tokenLeaves[i] = IRewardsCoordinator.TokenTreeMerkleLeaf({
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
        uint32 activatedAt = uint32(block.timestamp) + rewardsCoordinator.activationDelay();
        prevRootCalculationEndTimestamp = rootCalculationEndTimestamp;
        cheats.warp(activatedAt);

        uint32 rootIndex = uint32(rewardsCoordinator.getDistributionRootsLength());

        cheats.prank(rewardsUpdater);
        rewardsCoordinator.submitRoot(merkleRoot, prevRootCalculationEndTimestamp);

        IRewardsCoordinator.RewardsMerkleClaim memory newClaim = IRewardsCoordinator.RewardsMerkleClaim({
            rootIndex: rootIndex,
            earnerIndex: earnerIndex,
            earnerTreeProof: earnerTreeProof,
            earnerLeaf: IRewardsCoordinator.EarnerTreeMerkleLeaf({earner: earner, earnerTokenRoot: earnerTokenRoot}),
            tokenIndices: tokenIndices,
            tokenTreeProofs: tokenTreeProofs,
            tokenLeaves: tokenLeaves
        });

        return newClaim;
    }

    function _parseAllProofs() internal virtual returns (IRewardsCoordinator.RewardsMerkleClaim[] memory) {
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = new IRewardsCoordinator.RewardsMerkleClaim[](3);

        claims[0] = _parseProofData("src/test/test-data/rewardsCoordinator/processClaimProofs_Root1.json");
        claims[1] = _parseProofData("src/test/test-data/rewardsCoordinator/processClaimProofs_Root2.json");
        claims[2] = _parseProofData("src/test/test-data/rewardsCoordinator/processClaimProofs_Root3.json");

        return claims;
    }

    function _parseAllProofsMaxEarnerAndLeafIndices() internal virtual returns (IRewardsCoordinator.RewardsMerkleClaim[] memory) {
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = new IRewardsCoordinator.RewardsMerkleClaim[](1);

        claims[0] = _parseProofData("src/test/test-data/rewardsCoordinator/processClaimProofs_MaxEarnerAndLeafIndices.json");

        return claims;
    }

    function _parseAllProofsSingleTokenLeaf() internal virtual returns (IRewardsCoordinator.RewardsMerkleClaim[] memory) {
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = new IRewardsCoordinator.RewardsMerkleClaim[](1);

        claims[0] = _parseProofData("src/test/test-data/rewardsCoordinator/processClaimProofs_SingleTokenLeaf.json");

        return claims;
    }

    function _parseAllProofsSingleEarnerLeaf() internal virtual returns (IRewardsCoordinator.RewardsMerkleClaim[] memory) {
        IRewardsCoordinator.RewardsMerkleClaim[] memory claims = new IRewardsCoordinator.RewardsMerkleClaim[](1);

        claims[0] = _parseProofData("src/test/test-data/rewardsCoordinator/processClaimProofs_SingleEarnerLeaf.json");

        return claims;
    }
}

contract RewardsCoordinatorUnitTests_disableRoot is RewardsCoordinatorUnitTests {
    function testFuzz_disableRoot_Revert_WhenNotRewardsUpdater(
        address invalidRewardsUpdater
    ) public filterFuzzedAddressInputs(invalidRewardsUpdater) {
        cheats.prank(invalidRewardsUpdater);

        cheats.expectRevert("RewardsCoordinator: caller is not the rewardsUpdater");
        rewardsCoordinator.disableRoot(0);
    }

    function test_disableRoot_Revert_WhenPaused() public {
        cheats.prank(pauser);
        rewardsCoordinator.pause(2 ** PAUSED_SUBMIT_DISABLE_ROOTS);

        cheats.expectRevert("Pausable: index is paused");
        rewardsCoordinator.disableRoot(0);
    }

    function testFuzz_disableRoot_Revert_WhenRootIndexOutOfBounds(uint32 rootIndex) public {
        // submitRoots with indexes 0,1,2
        _submitRoots();
        cheats.assume(rootIndex > 2);

        cheats.expectRevert("RewardsCoordinator.disableRoot: invalid rootIndex");
        cheats.prank(rewardsUpdater);
        rewardsCoordinator.disableRoot(rootIndex);
    }

    function testFuzz_disableRoot_Revert_WhenRootAlreadyDisabled(uint32 rootIndex) public {
        // submitRoots with indexes 0,1,2
        _submitRoots();
        rootIndex = uint32(bound(uint256(rootIndex), 0, 2));
        
        cheats.startPrank(rewardsUpdater);
        rewardsCoordinator.disableRoot(rootIndex);

        cheats.expectRevert("RewardsCoordinator.disableRoot: root already disabled");
        rewardsCoordinator.disableRoot(rootIndex);
        cheats.stopPrank();
    }

    function testeFuzz_disableRoot_Revert_WhenRootAlreadyActivated(uint32 rootIndex) public {
        // submitRoots with indexes 0,1,2
        _submitRoots();
        rootIndex = uint32(bound(uint256(rootIndex), 0, 2));
        cheats.warp(block.timestamp + activationDelay);

        cheats.expectRevert("RewardsCoordinator.disableRoot: root already activated");
        cheats.prank(rewardsUpdater);
        rewardsCoordinator.disableRoot(rootIndex);
    }

    function testFuzz_disableRoot(uint32 rootIndex, uint256 randSalt) public {
        // Set timestamp to some value in the future
        cheats.warp(bound(randSalt, 1e5, 1e6));
        uint256 currTimestamp = block.timestamp;
        // submitRoots with indexes 0,1,2
        _submitRoots();
        rootIndex = uint32(bound(uint256(rootIndex), 0, 2));

        cheats.warp(currTimestamp);

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit DistributionRootDisabled(rootIndex);
        cheats.prank(rewardsUpdater);
        rewardsCoordinator.disableRoot(rootIndex);

        assertEq(
            rewardsCoordinator.getDistributionRootAtIndex(rootIndex).disabled,
            true,
            "root should be disabled"
        );
    }

    function _submitRoots() internal {
        cheats.warp(1e6);
        
        // submitRoots with indexes 0,1,2
        cheats.startPrank(rewardsUpdater);
        rewardsCoordinator.submitRoot(bytes32("0"), uint32(block.timestamp - 1));

        cheats.warp(block.timestamp + 2);
        rewardsCoordinator.submitRoot(bytes32("1"), uint32(block.timestamp - 1));

        cheats.warp(block.timestamp + 2);
        rewardsCoordinator.submitRoot(bytes32("2"), uint32(block.timestamp - 1));

        cheats.stopPrank();
    }
}

contract RewardsCoordinatorUnitTests_operatorCommission is RewardsCoordinatorUnitTests {
    uint8 numberOfRewardTypes = 1;

    function testFuzz_operatorCommissionBips_EmptyHistory(
        address operator,
        IAVSDirectory.OperatorSet calldata operatorSet,
        uint8 rewardTypeEnum
    ) public {
        rewardTypeEnum = uint8(bound(rewardTypeEnum, 0, numberOfRewardTypes - 1));
        IRewardsCoordinator.RewardType rewardType = IRewardsCoordinator.RewardType(rewardTypeEnum);
        rewardType = IRewardsCoordinator.RewardType(uint8(bound(uint256(rewardType), 0, 1)));
        // Check operator commission
        uint32 operatorCommissionBipsStored = rewardsCoordinator.getOperatorCommissionBips(operator, operatorSet, rewardType);
        assertEq(
            globalCommissionBips,
            operatorCommissionBipsStored,
            "Incorrect default operator commission bips"
        );
    }

    function testFuzz_setOperatorCommissionBips_Reverts_WhenBipsGreaterThanMax(
        address operator,
        IAVSDirectory.OperatorSet calldata operatorSet,
        uint8 rewardTypeEnum,
        uint16 newOperatorCommissionBips
    ) public filterFuzzedAddressInputs(operator) {
        rewardTypeEnum = uint8(bound(rewardTypeEnum, 0, numberOfRewardTypes - 1));
        IRewardsCoordinator.RewardType rewardType = IRewardsCoordinator.RewardType(rewardTypeEnum);
        cheats.assume(newOperatorCommissionBips > MAX_COMMISSION_BIPS);
        cheats.expectRevert("RewardsCoordinator.setOperatorCommissionBips: commissionBips too high");
        cheats.prank(operator);
        rewardsCoordinator.setOperatorCommissionBips(operatorSet, rewardType, newOperatorCommissionBips);
    }

    /// @notice test setting operator commission bips to a new value with empty history
    function testFuzz_setOperatorCommissionBips_EmptyHistory(
        address operator,
        IAVSDirectory.OperatorSet calldata operatorSet,
        uint8 rewardTypeEnum,
        uint16 newOperatorCommissionBips,
        uint256 randSalt
    ) public filterFuzzedAddressInputs(operator) {
        rewardTypeEnum = uint8(bound(rewardTypeEnum, 0, numberOfRewardTypes - 1));
        IRewardsCoordinator.RewardType rewardType = IRewardsCoordinator.RewardType(rewardTypeEnum);
        cheats.assume(newOperatorCommissionBips != globalCommissionBips);
        cheats.assume(newOperatorCommissionBips <= MAX_COMMISSION_BIPS);
        // 1. Set operator commission and check updated value hasn't taken effect yet
        uint16 prevOperatorCommissionBips = rewardsCoordinator.getOperatorCommissionBips(operator, operatorSet, rewardType);
        cheats.prank(operator);
        uint32 effectTimestamp = rewardsCoordinator.setOperatorCommissionBips(
            operatorSet,
            rewardType,
            newOperatorCommissionBips
        );
        assertTrue(
            rewardsCoordinator.getOperatorCommissionBips(operator, operatorSet, rewardType) == prevOperatorCommissionBips,
            "Operator commission bips should not be updated"
        );
        assertTrue(
            rewardsCoordinator.getOperatorCommissionBips(operator, operatorSet, rewardType) != newOperatorCommissionBips,
            "Operator commission bips should not be updated"
        );

        // 2. warp timestamp forwards to random time that is not the effectTimestamp
        uint256 warpTimestamp = bound(randSalt, block.timestamp, effectTimestamp - 1);
        cheats.warp(warpTimestamp);
        assertTrue(
            rewardsCoordinator.getOperatorCommissionBips(operator, operatorSet, rewardType) != newOperatorCommissionBips,
            "Operator commission bips should still be not updated"
        );

        // 3. warp timestamp forwards to effectTimestamp, operator commission should be updated now
        cheats.warp(effectTimestamp);
        assertTrue(
            rewardsCoordinator.getOperatorCommissionBips(operator, operatorSet, rewardType) == newOperatorCommissionBips,
            "Operator commission bips should be updated"
        );
    }

    /// @notice test setting operator commission bips to a new value with non empty history
    function testFuzz_setOperatorCommissionBips_WithHistory(
        address operator,
        IAVSDirectory.OperatorSet calldata operatorSet,
        uint8 rewardTypeEnum,
        uint16 newOperatorCommissionBips,
        uint256 randSalt
    ) public filterFuzzedAddressInputs(operator) {
        rewardTypeEnum = uint8(bound(rewardTypeEnum, 0, numberOfRewardTypes - 1));
        IRewardsCoordinator.RewardType rewardType = IRewardsCoordinator.RewardType(rewardTypeEnum);
        cheats.assume(newOperatorCommissionBips <= MAX_COMMISSION_BIPS);

        // 1. Set operator commission to initial value with existing history
        uint16 prevOperatorCommissionBips = uint16(bound(randSalt, 0, MAX_COMMISSION_BIPS));
        cheats.assume(prevOperatorCommissionBips != newOperatorCommissionBips);
        cheats.prank(operator);
        uint32 effectTimestamp = rewardsCoordinator.setOperatorCommissionBips(
            operatorSet,
            rewardType,
            prevOperatorCommissionBips
        );
        cheats.warp(effectTimestamp);
        assertTrue(
            rewardsCoordinator.getOperatorCommissionBips(operator, operatorSet, rewardType) == prevOperatorCommissionBips,
            "1. Operator commission bips should be initialized properly"
        );
        assertEq(
            rewardsCoordinator.getOperatorCommissionUpdateHistoryLength(operator, operatorSet, rewardType),
            1,
            "2. Invalid commission history length"
        );

        // 2. Set operator commission and check updated value hasn't taken effect yet
        cheats.prank(operator);
        effectTimestamp = rewardsCoordinator.setOperatorCommissionBips(
            operatorSet,
            rewardType,
            newOperatorCommissionBips
        );
        assertEq(
            rewardsCoordinator.getOperatorCommissionUpdateHistoryLength(operator, operatorSet, rewardType),
            2,
            "3. Invalid commission history length"
        );
        assertTrue(
            rewardsCoordinator.getOperatorCommissionBips(operator, operatorSet, rewardType) == prevOperatorCommissionBips,
            "4. Operator commission bips should not be updated"
        );
        assertTrue(
            rewardsCoordinator.getOperatorCommissionBips(operator, operatorSet, rewardType) != newOperatorCommissionBips,
            "5. Operator commission bips should not be updated"
        );

        // 3. warp timestamp forwards to random time that is not the effectTimestamp
        uint256 warpTimestamp = bound(randSalt, block.timestamp, effectTimestamp - 1);

        cheats.warp(warpTimestamp);
        assertTrue(
            rewardsCoordinator.getOperatorCommissionBips(operator, operatorSet, rewardType) == prevOperatorCommissionBips,
            "6. Operator commission bips should still be not updated"
        );

        // 4. warp timestamp forwards to effectTimestamp, operator commission should be updated now
        cheats.warp(effectTimestamp);
        assertTrue(
            rewardsCoordinator.getOperatorCommissionBips(operator, operatorSet, rewardType) == newOperatorCommissionBips,
            "7.Operator commission bips should be updated"
        );
    }

    /// @notice test setting operator commission bips to a new value with pending update
    /// has same effectTimestamp. Should overrwrite previous value
    function testFuzz_setOperatorCommissionBips_SameBlockUpdate(
        address operator,
        IAVSDirectory.OperatorSet calldata operatorSet,
        uint8 rewardTypeEnum,
        uint16 pendingCommissionBips,
        uint16 overwriteCommissionBips
    ) public filterFuzzedAddressInputs(operator) {
        rewardTypeEnum = uint8(bound(rewardTypeEnum, 0, numberOfRewardTypes - 1));
        IRewardsCoordinator.RewardType rewardType = IRewardsCoordinator.RewardType(rewardTypeEnum);
        cheats.assume(pendingCommissionBips <= MAX_COMMISSION_BIPS);
        cheats.assume(overwriteCommissionBips <= MAX_COMMISSION_BIPS);
        cheats.prank(operator);
        uint32 effectTimestamp = rewardsCoordinator.setOperatorCommissionBips(
            operatorSet,
            rewardType,
            pendingCommissionBips
        );
        assertEq(
            rewardsCoordinator.getOperatorCommissionUpdateHistoryLength(operator, operatorSet, rewardType),
            1,
            "1. Invalid commission history length"
        );
        cheats.prank(operator);
        uint32 effectTimestamp2 = rewardsCoordinator.setOperatorCommissionBips(
            operatorSet,
            rewardType,
            overwriteCommissionBips
        );
        assertEq(
            rewardsCoordinator.getOperatorCommissionUpdateHistoryLength(operator, operatorSet, rewardType),
            1,
            "2. Invalid commission history length"
        );
        assertEq(
            effectTimestamp,
            effectTimestamp2,
            "3. Effect timestamp should be the same"
        );
        cheats.warp(effectTimestamp);
        assertEq(
            rewardsCoordinator.getOperatorCommissionBips(operator, operatorSet, rewardType),
            overwriteCommissionBips,
            "4. Operator commission bips should be updated"
        );
    }
}

contract RewardsCoordinatorUnitTests_rewardOperatorSetForRange is RewardsCoordinatorUnitTests {
    uint32 defaultOperatorSetId = 0;
    uint32 validDuration = 2 weeks;
    uint32 validStartTimestamp = OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP;

    function test_Revert_WhenPaused() public {
        cheats.prank(pauser);
        rewardsCoordinator.pause(2 ** PAUSED_REWARD_OPERATOR_SET);

        cheats.expectRevert("Pausable: index is paused");
        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions;
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
    }

    function test_Revert_WhenReentrancy() public {
        uint256 amount = 1e38-1;
        Reenterer reenterer = new Reenterer();

        reenterer.prepareReturnData(abi.encode(amount));

        address targetToUse = address(rewardsCoordinator);
        uint256 msgValueToUse = 0;

        _deployMockRewardTokens(address(this), 1);

        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.OperatorSetRewardsSubmission({
            rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
            operatorSetId: defaultOperatorSetId,
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(reenterer)),
            amount: amount,
            startTimestamp: uint32(block.timestamp),
            duration: 0
        });

        bytes memory calldataToUse = abi.encodeWithSelector(RewardsCoordinator.createAVSRewardsSubmission.selector, rewardsSubmissions);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        cheats.expectRevert();
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
    }

    function test_Revert_EmptyStratsAndMultipliers() public {
        IRewardsCoordinator.StrategyAndMultiplier[] memory emptyStratsAndMultipliers;       
        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.OperatorSetRewardsSubmission({
            rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
            operatorSetId: defaultOperatorSetId,
            strategiesAndMultipliers: emptyStratsAndMultipliers,
            token: IERC20(address(this)),
            amount: 1,
            startTimestamp: uint32(block.timestamp),
            duration: 0
        });

        cheats.prank(defaultAVS);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: no strategies set");
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
    }

    function test_Revert_ZeroAmount() public {
        uint256 amount = 0;

        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.OperatorSetRewardsSubmission({
            rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
            operatorSetId: defaultOperatorSetId,
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(this)),
            amount: amount,
            startTimestamp: uint32(block.timestamp),
            duration: 0
        });

        cheats.prank(defaultAVS);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: amount cannot be 0");
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
    }

    function test_Revert_AmountTooLarge() public {
        uint256 amount = 1e38;

        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.OperatorSetRewardsSubmission({
            rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
            operatorSetId: defaultOperatorSetId,
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(this)),
            amount: amount,
            startTimestamp: uint32(block.timestamp),
            duration: 0
        });

        cheats.prank(defaultAVS);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: amount too large");
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
    }

    function test_Revert_DurationTooLarge() public {
        uint32 duration = MAX_REWARDS_DURATION + 1;

        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.OperatorSetRewardsSubmission({
            rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
            operatorSetId: defaultOperatorSetId,
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(this)),
            amount: 1,
            startTimestamp: uint32(block.timestamp),
            duration: duration
        });

        cheats.prank(defaultAVS);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: duration exceeds MAX_REWARDS_DURATION");
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
    }

    function test_Revert_DurationInvalidMultiple(uint256 duration) public {
        uint256 duration = CALCULATION_INTERVAL_SECONDS - 1;

        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.OperatorSetRewardsSubmission({
            rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
            operatorSetId: defaultOperatorSetId,
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(this)),
            amount: 1,
            startTimestamp: uint32(block.timestamp),
            duration: uint32(duration)
        });

        cheats.prank(defaultAVS);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: duration must be a multiple of CALCULATION_INTERVAL_SECONDS");
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
    }

    function test_Revert_InvalidStartTimestamp() public {
        uint32 startTimestamp = OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP + 1;

        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.OperatorSetRewardsSubmission({
            rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
            operatorSetId: defaultOperatorSetId,
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(this)),
            amount: 1,
            startTimestamp: startTimestamp,
            duration: validDuration
        });

        cheats.prank(defaultAVS);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: startTimestamp must be a multiple of CALCULATION_INTERVAL_SECONDS");
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
    }

    function test_Revert_StartTimestampPriorToGenesis() public {
        uint256 startTimestamp = OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP - CALCULATION_INTERVAL_SECONDS;

        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.OperatorSetRewardsSubmission({
            rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
            operatorSetId: defaultOperatorSetId,
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(this)),
            amount: 1,
            startTimestamp: uint32(startTimestamp),
            duration: validDuration
        });

        cheats.prank(defaultAVS);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: startTimestamp too far in the past");
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
    }

    /// @notice Warps block timestamp & attemps to go make a reward to genesis
    function test_Revert_StartTimestampGoesFarBack() public {
        /// Warp to a timestamp past max retroactive length
        cheats.warp(OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP + OPERATOR_SET_MAX_RETROACTIVE_LENGTH + CALCULATION_INTERVAL_SECONDS);
        uint256 startTimestamp = OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP;

        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.OperatorSetRewardsSubmission({
            rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
            operatorSetId: defaultOperatorSetId,
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(this)),
            amount: 1,
            startTimestamp: uint32(startTimestamp),
            duration: validDuration
        });

        cheats.prank(defaultAVS);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: startTimestamp too far in the past");
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
    }

    function test_Revert_StartTimestampTooFarInFuture() public {
        cheats.warp(OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP);

        uint256 startTimestamp = OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP + MAX_FUTURE_LENGTH + CALCULATION_INTERVAL_SECONDS;

        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.OperatorSetRewardsSubmission({
            rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
            operatorSetId: defaultOperatorSetId,
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(this)),
            amount: 1,
            startTimestamp: uint32(startTimestamp),
            duration: validDuration
        });

        cheats.prank(defaultAVS);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: startTimestamp too far in the future");
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
    }

    function test_Revert_InvalidStrategyConsidered() public {
        cheats.warp(OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP);
        defaultStrategyAndMultipliers[0].strategy = IStrategy(address(999));

        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.OperatorSetRewardsSubmission({
            rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
            operatorSetId: defaultOperatorSetId,
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(this)),
            amount: 1,
            startTimestamp: validStartTimestamp,
            duration: validDuration
        });

        cheats.prank(defaultAVS);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: invalid strategy considered");
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
    }

    function test_Revert_StrategiesNotOrdered() public {
        cheats.warp(OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP);

        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](1);
        IRewardsCoordinator.StrategyAndMultiplier[]
            memory dupStratsAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](2);
        dupStratsAndMultipliers[0] = defaultStrategyAndMultipliers[0];
        dupStratsAndMultipliers[1] = defaultStrategyAndMultipliers[0];
        rewardsSubmissions[0] = IRewardsCoordinator.OperatorSetRewardsSubmission({
            rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
            operatorSetId: defaultOperatorSetId,
            strategiesAndMultipliers: dupStratsAndMultipliers,
            token: IERC20(address(this)),
            amount: 1,
            startTimestamp: validStartTimestamp,
            duration: validDuration
        });

        cheats.prank(defaultAVS);
        cheats.expectRevert("RewardsCoordinator._validateRewardsSubmission: strategies must be in ascending order to handle duplicates");
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
    }

    function testFuzz_RewardOperatorSetForRange_SingleSubmission(
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount
    ) public {
        cheats.warp(OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP);

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, defaultAVS);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, 0, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint256(_maxTimestamp(OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - OPERATOR_SET_MAX_RETROACTIVE_LENGTH)) +
                CALCULATION_INTERVAL_SECONDS -
                1,
            block.timestamp + uint256(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.OperatorSetRewardsSubmission({
            rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
            operatorSetId: defaultOperatorSetId,
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected event emitted
        uint256 avsBalanceBefore = rewardToken.balanceOf(defaultAVS);
        uint256 rewardsCoordinatorBalanceBefore = rewardToken.balanceOf(address(rewardsCoordinator));

        cheats.startPrank(defaultAVS);
        rewardToken.approve(address(rewardsCoordinator), amount);
        uint256 currSubmissionNonce = rewardsCoordinator.submissionNonce(defaultAVS);
        bytes32 rewardsSubmissionHash = keccak256(abi.encode(defaultAVS, currSubmissionNonce, rewardsSubmissions[0]));

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorSetRewardCreated(defaultAVS, currSubmissionNonce, rewardsSubmissionHash, rewardsSubmissions[0]);
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);
        cheats.stopPrank();

        assertTrue(rewardsCoordinator.isAVSRewardsSubmissionHash(defaultAVS, rewardsSubmissionHash), "rewards submission hash not submitted");
        assertEq(currSubmissionNonce + 1, rewardsCoordinator.submissionNonce(defaultAVS), "submission nonce not incremented");
        assertEq(
            avsBalanceBefore - amount,
            rewardToken.balanceOf(defaultAVS),
            "AVS balance not decremented by amount of rewards submission"
        );
        assertEq(
            rewardsCoordinatorBalanceBefore + amount,
            rewardToken.balanceOf(address(rewardsCoordinator)),
            "RewardsCoordinator balance not incremented by amount of rewards submission"
        );
    }

    /**
     * @notice test multiple rewards submissions asserting for the following
     * - correct event emitted
     * - submission nonce incrementation by numSubmissions, and rewards submission hashes being set in storage.
     * - rewards submission hash being set in storage
     * - token balances before and after of avs and rewardsCoordinator
     */
    function testFuzz_RewardOperatorSetForRange_MultipleSubmissions(
        uint256 startTimestamp,
        uint256 duration,
        uint256 amount,
        uint256 numSubmissions
    ) public {
        cheats.assume(2 <= numSubmissions && numSubmissions <= 10);
        cheats.warp(OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP);

        IRewardsCoordinator.OperatorSetRewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.OperatorSetRewardsSubmission[](numSubmissions);
        bytes32[] memory rewardsSubmissionHashes = new bytes32[](numSubmissions);
        uint256 startSubmissionNonce = rewardsCoordinator.submissionNonce(defaultAVS);
        _deployMockRewardTokens(defaultAVS, numSubmissions);

        uint256[] memory avsBalancesBefore = _getBalanceForTokens(rewardTokens, defaultAVS);
        uint256[] memory rewardsCoordinatorBalancesBefore = _getBalanceForTokens(
            rewardTokens,
            address(rewardsCoordinator)
        );
        uint256[] memory amounts = new uint256[](numSubmissions);

        // Create multiple rewards submissions and their expected event
        for (uint256 i = 0; i < numSubmissions; ++i) {
            // 1. Bound fuzz inputs to valid ranges and amounts using randSeed for each
            amount = bound(amount + i, 1, mockTokenInitialSupply);
            amounts[i] = amount;
            duration = bound(duration + i, 0, MAX_REWARDS_DURATION);
            duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
            startTimestamp = bound(
                startTimestamp + i,
                uint256(_maxTimestamp(OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - OPERATOR_SET_MAX_RETROACTIVE_LENGTH)) +
                    CALCULATION_INTERVAL_SECONDS -
                    1,
                block.timestamp + uint256(MAX_FUTURE_LENGTH)
            );
            startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

            // 2. Create rewards submission input param
            IRewardsCoordinator.OperatorSetRewardsSubmission memory rewardsSubmission = IRewardsCoordinator.OperatorSetRewardsSubmission({
                rewardType: IRewardsCoordinator.RewardType.DELEGATED_STAKE,
                operatorSetId: defaultOperatorSetId,
                strategiesAndMultipliers: defaultStrategyAndMultipliers,
                token: rewardTokens[i],
                amount: amounts[i],
                startTimestamp: uint32(startTimestamp),
                duration: uint32(duration)
            });
            rewardsSubmissions[i] = rewardsSubmission;

            // 3. expected event emitted for this rewardsSubmission
            rewardsSubmissionHashes[i] = keccak256(abi.encode(defaultAVS, startSubmissionNonce + i, rewardsSubmissions[i]));
            cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
            emit OperatorSetRewardCreated(defaultAVS, startSubmissionNonce + i, rewardsSubmissionHashes[i], rewardsSubmissions[i]);
        }

        // 4. call createAVSRewardsSubmission()
        cheats.prank(defaultAVS);
        rewardsCoordinator.rewardOperatorSetForRange(rewardsSubmissions);

        // 5. Check for submissionNonce() and rewardsSubmissionHashes being set
        assertEq(
            startSubmissionNonce + numSubmissions,
            rewardsCoordinator.submissionNonce(defaultAVS),
            "submission nonce not incremented properly"
        );

        for (uint256 i = 0; i < numSubmissions; ++i) {
            assertTrue(
                rewardsCoordinator.isAVSRewardsSubmissionHash(defaultAVS, rewardsSubmissionHashes[i]),
                "rewards submission hash not submitted"
            );
            assertEq(
                avsBalancesBefore[i] - amounts[i],
                rewardTokens[i].balanceOf(defaultAVS),
                "AVS balance not decremented by amount of rewards submission"
            );
            assertEq(
                rewardsCoordinatorBalancesBefore[i] + amounts[i],
                rewardTokens[i].balanceOf(address(rewardsCoordinator)),
                "RewardsCoordinator balance not incremented by amount of rewards submission"
            );
        }
    }
}
