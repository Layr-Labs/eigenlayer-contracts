// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/core/RewardsCoordinator.sol";
import "src/contracts/strategies/StrategyBase.sol";

import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/test/mocks/Reenterer.sol";
import "src/test/mocks/ERC20Mock.sol";

/**
 * @notice Unit testing of the RewardsCoordinator contract
 * Contracts tested: RewardsCoordinator
 * Contracts not mocked: StrategyBase, PauserRegistry
 */
contract RewardsCoordinatorUnitTests is EigenLayerUnitTestSetup, IRewardsCoordinatorEvents, IRewardsCoordinatorErrors {
    // used for stack too deep
    struct FuzzAVSRewardsSubmission {
        address avs;
        uint startTimestamp;
        uint duration;
        uint amount;
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
    uint mockTokenInitialSupply = 1e38 - 1;
    StrategyAndMultiplier[] defaultStrategyAndMultipliers;

    // Config Variables
    /// @notice intervals(epochs) are 1 day: https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-001.md#updated-calculation-interval-seconds
    uint32 CALCULATION_INTERVAL_SECONDS = 1 days;

    /// @notice Max duration is 5 epochs (2 weeks * 5 = 10 weeks in seconds)
    uint32 MAX_REWARDS_DURATION = 70 days;

    /// @notice Lower bound start range is ~3 months into the past, multiple of CALCULATION_INTERVAL_SECONDS
    uint32 MAX_RETROACTIVE_LENGTH = 84 days;
    /// @notice Upper bound start range is ~1 month into the future, multiple of CALCULATION_INTERVAL_SECONDS
    uint32 MAX_FUTURE_LENGTH = 28 days;
    /// @notice absolute min timestamp that a rewards can start at
    uint32 GENESIS_REWARDS_TIMESTAMP = 1_712_188_800;
    /// @notice Equivalent to 100%, but in basis points.
    uint16 internal constant ONE_HUNDRED_IN_BIPS = 10_000;

    /// @notice Delay in timestamp before a posted root can be claimed against
    uint32 activationDelay = 7 days;
    /// @notice the split for all operators across all avss
    uint16 defaultSplitBips = 1000;

    IERC20[] rewardTokens;

    // RewardsCoordinator Constants

    /// @dev Index for flag that pauses calling createAVSRewardsSubmission
    uint8 internal constant PAUSED_AVS_REWARDS_SUBMISSION = 0;

    /// @dev Index for flag that pauses calling createRewardsForAllSubmission
    uint8 internal constant PAUSED_REWARDS_FOR_ALL_SUBMISSION = 1;

    /// @dev Index for flag that pauses claiming
    uint8 internal constant PAUSED_PROCESS_CLAIM = 2;

    /// @dev Index for flag that pauses submitRoots
    uint8 internal constant PAUSED_SUBMIT_ROOTS = 3;

    /// @dev Index for flag that pauses rewardAllStakersAndOperators
    uint8 internal constant PAUSED_REWARD_ALL_STAKERS_AND_OPERATORS = 4;

    /// @dev Index for flag that pauses calling createOperatorDirectedAVSRewardsSubmission
    uint8 internal constant PAUSED_OPERATOR_DIRECTED_AVS_REWARDS_SUBMISSION = 5;

    /// @dev Index for flag that pauses calling setOperatorAVSSplit
    uint8 internal constant PAUSED_OPERATOR_AVS_SPLIT = 6;

    /// @dev Index for flag that pauses calling setOperatorPISplit
    uint8 internal constant PAUSED_OPERATOR_PI_SPLIT = 7;

    /// @dev Index for flag that pauses calling setOperatorSetSplit
    uint8 internal constant PAUSED_OPERATOR_SET_SPLIT = 8;

    /// @dev Index for flag that pauses calling setOperatorSetPerformanceRewardsSubmission
    uint8 internal constant PAUSED_OPERATOR_DIRECTED_OPERATOR_SET_REWARDS_SUBMISSION = 9;

    // RewardsCoordinator entities
    address rewardsUpdater = address(1000);
    address defaultAVS = address(1001);
    address defaultClaimer = address(1002);
    address rewardsForAllSubmitter = address(1003);
    address defaultAppointee = address(1004);

    function setUp() public virtual override {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        // Deploy RewardsCoordinator proxy and implementation
        rewardsCoordinatorImplementation = new RewardsCoordinator(
            IRewardsCoordinatorTypes.RewardsCoordinatorConstructorParams({
                delegationManager: IDelegationManager(address(delegationManagerMock)),
                strategyManager: IStrategyManager(address(strategyManagerMock)),
                allocationManager: IAllocationManager(address(allocationManagerMock)),
                pauserRegistry: pauserRegistry,
                permissionController: IPermissionController(address(permissionController)),
                CALCULATION_INTERVAL_SECONDS: CALCULATION_INTERVAL_SECONDS,
                MAX_REWARDS_DURATION: MAX_REWARDS_DURATION,
                MAX_RETROACTIVE_LENGTH: MAX_RETROACTIVE_LENGTH,
                MAX_FUTURE_LENGTH: MAX_FUTURE_LENGTH,
                GENESIS_REWARDS_TIMESTAMP: GENESIS_REWARDS_TIMESTAMP,
                version: "v9.9.9"
            })
        );

        rewardsCoordinator = RewardsCoordinator(
            address(
                new TransparentUpgradeableProxy(
                    address(rewardsCoordinatorImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        RewardsCoordinator.initialize.selector,
                        address(this), // initOwner
                        0, // 0 is initialPausedStatus
                        rewardsUpdater,
                        activationDelay,
                        defaultSplitBips
                    )
                )
            )
        );

        // Deploy mock token and strategy
        token1 = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, address(this));
        token2 = new ERC20PresetFixedSupply("jeo boden", "MOCK2", mockTokenInitialSupply, address(this));
        token3 = new ERC20PresetFixedSupply("pepe wif avs", "MOCK3", mockTokenInitialSupply, address(this));

        strategyImplementation = new StrategyBase(IStrategyManager(address(strategyManagerMock)), pauserRegistry, "v9.9.9");
        strategyMock1 = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, token1)
                )
            )
        );
        strategyMock2 = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, token2)
                )
            )
        );
        strategyMock3 = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, token3)
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

        defaultStrategyAndMultipliers.push(StrategyAndMultiplier(IStrategy(address(strategies[0])), 1e18));
        defaultStrategyAndMultipliers.push(StrategyAndMultiplier(IStrategy(address(strategies[1])), 2e18));
        defaultStrategyAndMultipliers.push(StrategyAndMultiplier(IStrategy(address(strategies[2])), 3e18));

        rewardsCoordinator.setRewardsForAllSubmitter(rewardsForAllSubmitter, true);
        rewardsCoordinator.setRewardsUpdater(rewardsUpdater);

        // Exclude from fuzzed tests
        isExcludedFuzzAddress[address(rewardsCoordinator)] = true;
        isExcludedFuzzAddress[address(rewardsUpdater)] = true;

        // Set the timestamp to some time after the genesis rewards timestamp
        cheats.warp(GENESIS_REWARDS_TIMESTAMP + 5 days);
    }

    /// @notice deploy token to owner and approve rewardsCoordinator. Used for deploying reward tokens
    function _deployMockRewardTokens(address owner, uint numTokens) internal virtual {
        cheats.startPrank(owner);
        for (uint i = 0; i < numTokens; ++i) {
            IERC20 token = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, owner);
            rewardTokens.push(token);
            token.approve(address(rewardsCoordinator), mockTokenInitialSupply);
        }
        cheats.stopPrank();
    }

    function _getBalanceForTokens(IERC20[] memory tokens, address holder) internal view returns (uint[] memory) {
        uint[] memory balances = new uint[](tokens.length);
        for (uint i = 0; i < tokens.length; ++i) {
            balances[i] = tokens[i].balanceOf(holder);
        }
        return balances;
    }

    function _maxTimestamp(uint32 timestamp1, uint32 timestamp2) internal pure returns (uint32) {
        return timestamp1 > timestamp2 ? timestamp1 : timestamp2;
    }

    function _assertRewardsClaimedEvents(bytes32 root, RewardsMerkleClaim memory claim, address recipient) internal {
        address earner = claim.earnerLeaf.earner;
        address claimer = rewardsCoordinator.claimerFor(earner);
        if (claimer == address(0)) claimer = earner;
        IERC20 token;
        uint claimedAmount;
        for (uint i = 0; i < claim.tokenLeaves.length; ++i) {
            token = claim.tokenLeaves[i].token;
            claimedAmount = rewardsCoordinator.cumulativeClaimed(earner, token);

            cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
            emit RewardsClaimed(root, earner, claimer, recipient, token, claim.tokenLeaves[i].cumulativeEarnings - claimedAmount);
        }
    }

    /// @notice given address and array of reward tokens, return array of cumulativeClaimed amonts
    function _getCumulativeClaimed(address earner, RewardsMerkleClaim memory claim) internal view returns (uint[] memory) {
        uint[] memory totalClaimed = new uint[](claim.tokenLeaves.length);

        for (uint i = 0; i < claim.tokenLeaves.length; ++i) {
            totalClaimed[i] = rewardsCoordinator.cumulativeClaimed(earner, claim.tokenLeaves[i].token);
        }

        return totalClaimed;
    }

    /// @notice given a claim, return the new cumulativeEarnings for each token
    function _getCumulativeEarnings(RewardsMerkleClaim memory claim) internal pure returns (uint[] memory) {
        uint[] memory earnings = new uint[](claim.tokenLeaves.length);

        for (uint i = 0; i < claim.tokenLeaves.length; ++i) {
            earnings[i] = claim.tokenLeaves[i].cumulativeEarnings;
        }

        return earnings;
    }

    function _getClaimTokenBalances(address earner, RewardsMerkleClaim memory claim) internal view returns (uint[] memory) {
        uint[] memory balances = new uint[](claim.tokenLeaves.length);

        for (uint i = 0; i < claim.tokenLeaves.length; ++i) {
            balances[i] = claim.tokenLeaves[i].token.balanceOf(earner);
        }

        return balances;
    }

    /// @dev Sort to ensure that the array is in ascending order for strategies
    function _sortArrayAsc(IStrategy[] memory arr) internal pure returns (IStrategy[] memory) {
        uint l = arr.length;
        for (uint i = 0; i < l; i++) {
            for (uint j = i + 1; j < l; j++) {
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

    function test_setClaimerFor_UAM_revert_staker(address earner, address claimer) public filterFuzzedAddressInputs(earner) {
        cheats.prank(earner);
        cheats.expectRevert(InvalidEarner.selector);
        rewardsCoordinator.setClaimerFor(earner, claimer);
    }

    function test_setClaimerFor_UAM_AVS() public {
        address avs = address(1000);
        address claimer = address(1001);

        // Set AVS
        allocationManagerMock.setAVSSetCount(avs, 1);

        // Initialize UAM
        cheats.prank(avs);
        permissionController.setAppointee(
            avs, defaultAppointee, address(rewardsCoordinator), bytes4(keccak256("setClaimerFor(address,address)"))
        );

        // Set claimer for AVS
        cheats.startPrank(defaultAppointee);
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit ClaimerForSet(avs, rewardsCoordinator.claimerFor(avs), claimer);
        rewardsCoordinator.setClaimerFor(avs, claimer);
        cheats.stopPrank();

        assertEq(claimer, rewardsCoordinator.claimerFor(avs), "claimerFor not set");
    }

    function test_setClaimerFor_UAM_Operator() public {
        address operator = address(1000);
        address claimer = address(1001);

        // Set operator
        delegationManagerMock.setIsOperator(operator, true);

        // Initialize UAM
        cheats.prank(operator);
        permissionController.setAppointee(
            operator, defaultAppointee, address(rewardsCoordinator), bytes4(keccak256("setClaimerFor(address,address)"))
        );

        // Set claimer for operator
        cheats.startPrank(defaultAppointee);
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit ClaimerForSet(operator, rewardsCoordinator.claimerFor(operator), claimer);
        rewardsCoordinator.setClaimerFor(operator, claimer);
        cheats.stopPrank();

        assertEq(claimer, rewardsCoordinator.claimerFor(operator), "claimerFor not set");
    }

    function testFuzz_setActivationDelay(uint32 activationDelay) public {
        cheats.startPrank(rewardsCoordinator.owner());
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit ActivationDelaySet(rewardsCoordinator.activationDelay(), activationDelay);
        rewardsCoordinator.setActivationDelay(activationDelay);
        assertEq(activationDelay, rewardsCoordinator.activationDelay(), "activationDelay not set");
        cheats.stopPrank();
    }

    function testFuzz_setActivationDelay_Revert_WhenNotOwner(address caller, uint32 activationDelay)
        public
        filterFuzzedAddressInputs(caller)
    {
        cheats.assume(caller != rewardsCoordinator.owner());
        cheats.prank(caller);
        cheats.expectRevert("Ownable: caller is not the owner");
        rewardsCoordinator.setActivationDelay(activationDelay);
    }

    function testFuzz_setDefaultOperatorSplit(uint16 defaultSplitBips) public {
        cheats.startPrank(rewardsCoordinator.owner());
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit DefaultOperatorSplitBipsSet(rewardsCoordinator.defaultOperatorSplitBips(), defaultSplitBips);
        rewardsCoordinator.setDefaultOperatorSplit(defaultSplitBips);
        assertEq(defaultSplitBips, rewardsCoordinator.defaultOperatorSplitBips(), "defaultOperatorSplitBips not set");
        cheats.stopPrank();
    }

    function testFuzz_setDefaultOperatorSplit_Revert_WhenNotOwner(address caller, uint16 defaultSplitBips)
        public
        filterFuzzedAddressInputs(caller)
    {
        cheats.assume(caller != rewardsCoordinator.owner());
        cheats.prank(caller);
        cheats.expectRevert("Ownable: caller is not the owner");
        rewardsCoordinator.setDefaultOperatorSplit(defaultSplitBips);
    }

    function testFuzz_setRewardsUpdater(address newRewardsUpdater) public {
        cheats.startPrank(rewardsCoordinator.owner());
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit RewardsUpdaterSet(rewardsCoordinator.rewardsUpdater(), newRewardsUpdater);
        rewardsCoordinator.setRewardsUpdater(newRewardsUpdater);
        assertEq(newRewardsUpdater, rewardsCoordinator.rewardsUpdater(), "rewardsUpdater not set");
        cheats.stopPrank();
    }

    function testFuzz_setRewardsUpdater_Revert_WhenNotOwner(address caller, address newRewardsUpdater)
        public
        filterFuzzedAddressInputs(caller)
    {
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
        assertEq(newValue, rewardsCoordinator.isRewardsForAllSubmitter(submitter), "isRewardsForAllSubmitter not set");
        cheats.stopPrank();
    }

    function testFuzz_setRewardsForAllSubmitter_Revert_WhenNotOwner(address caller, address submitter, bool newValue)
        public
        filterFuzzedAddressInputs(caller)
    {
        cheats.assume(caller != rewardsCoordinator.owner());
        cheats.prank(caller);
        cheats.expectRevert("Ownable: caller is not the owner");
        rewardsCoordinator.setRewardsForAllSubmitter(submitter, newValue);
    }
}

contract RewardsCoordinatorUnitTests_setOperatorAVSSplit is RewardsCoordinatorUnitTests {
    // Revert when paused
    function testFuzz_Revert_WhenPaused(address operator, address avs, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        split = uint16(bound(split, 0, ONE_HUNDRED_IN_BIPS));
        cheats.prank(pauser);
        rewardsCoordinator.pause(2 ** PAUSED_OPERATOR_AVS_SPLIT);

        cheats.prank(operator);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        rewardsCoordinator.setOperatorAVSSplit(operator, avs, split);
    }

    // Revert when split is greater than 100%
    function testFuzz_Revert_WhenSplitGreaterThan100(address operator, address avs, uint16 split)
        public
        filterFuzzedAddressInputs(operator)
    {
        cheats.assume(operator != address(0));
        split = uint16(bound(split, ONE_HUNDRED_IN_BIPS + 1, type(uint16).max));

        cheats.prank(operator);
        cheats.expectRevert(SplitExceedsMax.selector);
        rewardsCoordinator.setOperatorAVSSplit(operator, avs, split);
    }

    function testFuzz_setOperatorAVSSplit(address operator, address avs, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        split = uint16(bound(split, 0, ONE_HUNDRED_IN_BIPS));
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;
        uint16 oldSplit = rewardsCoordinator.getOperatorAVSSplit(operator, avs);

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorAVSSplitBipsSet(operator, operator, avs, activatedAt, oldSplit, split);
        cheats.prank(operator);
        rewardsCoordinator.setOperatorAVSSplit(operator, avs, split);

        assertEq(oldSplit, rewardsCoordinator.getOperatorAVSSplit(operator, avs), "Incorrect Operator split");
        cheats.warp(activatedAt);
        assertEq(split, rewardsCoordinator.getOperatorAVSSplit(operator, avs), "Incorrect Operator split");
    }

    function testFuzz_setOperatorAVSSplit_UAM(address operator, address avs, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));

        // Set UAM
        cheats.prank(operator);
        permissionController.setAppointee(
            operator, defaultAppointee, address(rewardsCoordinator), IRewardsCoordinator.setOperatorAVSSplit.selector
        );

        split = uint16(bound(split, 0, ONE_HUNDRED_IN_BIPS));
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;
        uint16 oldSplit = rewardsCoordinator.getOperatorAVSSplit(operator, avs);

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorAVSSplitBipsSet(defaultAppointee, operator, avs, activatedAt, oldSplit, split);
        cheats.prank(defaultAppointee);
        rewardsCoordinator.setOperatorAVSSplit(operator, avs, split);

        assertEq(oldSplit, rewardsCoordinator.getOperatorAVSSplit(operator, avs), "Incorrect Operator split");
        cheats.warp(activatedAt);
        assertEq(split, rewardsCoordinator.getOperatorAVSSplit(operator, avs), "Incorrect Operator split");
    }

    // Testing that the split has been initialized for the first time.
    function testFuzz_setOperatorAVSSplitFirstTime(address operator, address avs, uint16 split)
        public
        filterFuzzedAddressInputs(operator)
    {
        cheats.assume(operator != address(0));
        split = uint16(bound(split, 0, ONE_HUNDRED_IN_BIPS));
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;
        uint16 oldSplit = rewardsCoordinator.getOperatorAVSSplit(operator, avs);
        // Check that the split returns the default split before initialization for the first time.
        assertEq(oldSplit, rewardsCoordinator.defaultOperatorSplitBips(), "Operator split is not Default split before Initialization");

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorAVSSplitBipsSet(operator, operator, avs, activatedAt, oldSplit, split);
        cheats.prank(operator);
        rewardsCoordinator.setOperatorAVSSplit(operator, avs, split);

        cheats.prank(address(this)); // Owner of RewardsCoordinator
        // Change default split to check if it is returned before activation
        rewardsCoordinator.setDefaultOperatorSplit(5000);
        // Check that the split returns the default split before activation for the first time.
        assertEq(
            rewardsCoordinator.defaultOperatorSplitBips(),
            rewardsCoordinator.getOperatorAVSSplit(operator, avs),
            "Operator split is not Default split before Activation for first time"
        );

        cheats.warp(activatedAt);
        assertEq(split, rewardsCoordinator.getOperatorAVSSplit(operator, avs), "Incorrect Operator split");
    }

    // Testing the split setting for a second time prior to the earlier activation.
    function testFuzz_Revert_setOperatorAVSSplitSecondTimeBeforePriorActivation(
        address operator,
        address avs,
        uint16 firstSplit,
        uint16 secondSplit,
        uint32 warpTime
    ) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        firstSplit = uint16(bound(firstSplit, 0, ONE_HUNDRED_IN_BIPS));
        secondSplit = uint16(bound(secondSplit, 0, ONE_HUNDRED_IN_BIPS));
        warpTime = uint32(bound(warpTime, uint32(block.timestamp), uint32(block.timestamp) + activationDelay));

        // Setting First Split
        cheats.prank(operator);
        rewardsCoordinator.setOperatorAVSSplit(operator, avs, firstSplit);
        // Warping to time before activation of First split
        cheats.warp(warpTime);

        // Trying to set Second Split
        cheats.prank(operator);
        cheats.expectRevert(PreviousSplitPending.selector);
        rewardsCoordinator.setOperatorAVSSplit(operator, avs, secondSplit);
    }

    // Testing the split setting for a second time after earlier activation.
    function testFuzz_setOperatorAVSSplitSecondTimeAfterPriorActivation(
        address operator,
        address avs,
        uint16 firstSplit,
        uint16 secondSplit,
        uint32 warpTime
    ) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        firstSplit = uint16(bound(firstSplit, 0, ONE_HUNDRED_IN_BIPS));
        secondSplit = uint16(bound(secondSplit, 0, ONE_HUNDRED_IN_BIPS));
        warpTime = uint32(bound(warpTime, uint32(block.timestamp) + activationDelay + 1, type(uint32).max - activationDelay));

        // Setting First Split
        cheats.prank(operator);
        rewardsCoordinator.setOperatorAVSSplit(operator, avs, firstSplit);
        // Warping to time after activation of First split
        cheats.warp(warpTime);
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;

        // Setting Second Split
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorAVSSplitBipsSet(operator, operator, avs, activatedAt, firstSplit, secondSplit);
        cheats.prank(operator);
        rewardsCoordinator.setOperatorAVSSplit(operator, avs, secondSplit);

        assertEq(firstSplit, rewardsCoordinator.getOperatorAVSSplit(operator, avs), "Incorrect Operator split");
        cheats.warp(activatedAt);
        assertEq(secondSplit, rewardsCoordinator.getOperatorAVSSplit(operator, avs), "Incorrect Operator split");
    }
}

contract RewardsCoordinatorUnitTests_setOperatorPISplit is RewardsCoordinatorUnitTests {
    // Revert when paused
    function testFuzz_Revert_WhenPaused(address operator, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        split = uint16(bound(split, 0, ONE_HUNDRED_IN_BIPS));
        cheats.prank(pauser);
        rewardsCoordinator.pause(2 ** PAUSED_OPERATOR_PI_SPLIT);

        cheats.prank(operator);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        rewardsCoordinator.setOperatorPISplit(operator, split);
    }

    // Revert when split is greater than 100%
    function testFuzz_Revert_WhenSplitGreaterThan100(address operator, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        split = uint16(bound(split, ONE_HUNDRED_IN_BIPS + 1, type(uint16).max));

        cheats.prank(operator);
        cheats.expectRevert(SplitExceedsMax.selector);
        rewardsCoordinator.setOperatorPISplit(operator, split);
    }

    function testFuzz_setOperatorPISplit(address operator, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));

        split = uint16(bound(split, 0, ONE_HUNDRED_IN_BIPS));
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;
        uint16 oldSplit = rewardsCoordinator.getOperatorPISplit(operator);

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorPISplitBipsSet(operator, operator, activatedAt, oldSplit, split);
        cheats.prank(operator);
        rewardsCoordinator.setOperatorPISplit(operator, split);

        assertEq(oldSplit, rewardsCoordinator.getOperatorPISplit(operator), "Incorrect Operator split");
        cheats.warp(activatedAt);
        assertEq(split, rewardsCoordinator.getOperatorPISplit(operator), "Incorrect Operator split");
    }

    function testFuzz_setOperatorPISplit_UAM(address operator, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));

        // Set UAM
        cheats.prank(operator);
        permissionController.setAppointee(
            operator, defaultAppointee, address(rewardsCoordinator), IRewardsCoordinator.setOperatorPISplit.selector
        );

        split = uint16(bound(split, 0, ONE_HUNDRED_IN_BIPS));
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;
        uint16 oldSplit = rewardsCoordinator.getOperatorPISplit(operator);

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorPISplitBipsSet(defaultAppointee, operator, activatedAt, oldSplit, split);
        cheats.prank(defaultAppointee);
        rewardsCoordinator.setOperatorPISplit(operator, split);

        assertEq(oldSplit, rewardsCoordinator.getOperatorPISplit(operator), "Incorrect Operator split");
        cheats.warp(activatedAt);
        assertEq(split, rewardsCoordinator.getOperatorPISplit(operator), "Incorrect Operator split");
    }

    // Testing that the split has been initialized for the first time.
    function testFuzz_setOperatorPISplitFirstTime(address operator, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        split = uint16(bound(split, 0, ONE_HUNDRED_IN_BIPS));
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;
        uint16 oldSplit = rewardsCoordinator.getOperatorPISplit(operator);
        // Check that the split returns the default split before initialization for the first time.
        assertEq(oldSplit, rewardsCoordinator.defaultOperatorSplitBips(), "Operator split is not Default split before Initialization");

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorPISplitBipsSet(operator, operator, activatedAt, oldSplit, split);
        cheats.prank(operator);
        rewardsCoordinator.setOperatorPISplit(operator, split);

        cheats.prank(address(this)); // Owner of RewardsCoordinator
        // Change default split to check if it is returned before activation
        rewardsCoordinator.setDefaultOperatorSplit(5000);
        // Check that the split returns the default split before activation for the first time.
        assertEq(
            rewardsCoordinator.defaultOperatorSplitBips(),
            rewardsCoordinator.getOperatorPISplit(operator),
            "Operator split is not Default split before Activation for first time"
        );

        cheats.warp(activatedAt);
        assertEq(split, rewardsCoordinator.getOperatorPISplit(operator), "Incorrect Operator split");
    }

    // Testing the split setting for a second time prior to the earlier activation.
    function testFuzz_Revert_setOperatorPISplitSecondTimeBeforePriorActivation(
        address operator,
        uint16 firstSplit,
        uint16 secondSplit,
        uint32 warpTime
    ) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        firstSplit = uint16(bound(firstSplit, 0, ONE_HUNDRED_IN_BIPS));
        secondSplit = uint16(bound(secondSplit, 0, ONE_HUNDRED_IN_BIPS));
        warpTime = uint32(bound(warpTime, uint32(block.timestamp), uint32(block.timestamp) + activationDelay));

        // Setting First Split
        cheats.prank(operator);
        rewardsCoordinator.setOperatorPISplit(operator, firstSplit);
        // Warping to time before activation of First split
        cheats.warp(warpTime);

        // Trying to set Second Split
        cheats.prank(operator);
        cheats.expectRevert(PreviousSplitPending.selector);
        rewardsCoordinator.setOperatorPISplit(operator, secondSplit);
    }

    // Testing the split setting for a second time after earlier activation.
    function testFuzz_setOperatorPISplitSecondTimeAfterPriorActivation(
        address operator,
        uint16 firstSplit,
        uint16 secondSplit,
        uint32 warpTime
    ) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        firstSplit = uint16(bound(firstSplit, 0, ONE_HUNDRED_IN_BIPS));
        secondSplit = uint16(bound(secondSplit, 0, ONE_HUNDRED_IN_BIPS));
        warpTime = uint32(bound(warpTime, uint32(block.timestamp) + activationDelay + 1, type(uint32).max - activationDelay));

        // Setting First Split
        cheats.prank(operator);
        rewardsCoordinator.setOperatorPISplit(operator, firstSplit);
        // Warping to time after activation of First split
        cheats.warp(warpTime);
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;

        // Setting Second Split
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorPISplitBipsSet(operator, operator, activatedAt, firstSplit, secondSplit);
        cheats.prank(operator);
        rewardsCoordinator.setOperatorPISplit(operator, secondSplit);

        assertEq(firstSplit, rewardsCoordinator.getOperatorPISplit(operator), "Incorrect Operator split");
        cheats.warp(activatedAt);
        assertEq(secondSplit, rewardsCoordinator.getOperatorPISplit(operator), "Incorrect Operator split");
    }
}

contract RewardsCoordinatorUnitsTests_setOperatorSetSplit is RewardsCoordinatorUnitTests {
    OperatorSet operatorSet;

    function setUp() public virtual override {
        RewardsCoordinatorUnitTests.setUp();
        operatorSet = OperatorSet(address(this), 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);
    }

    // Revert when paused
    function testFuzz_Revert_WhenPaused(address operator, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        split = uint16(bound(split, 0, ONE_HUNDRED_IN_BIPS));
        cheats.prank(pauser);
        rewardsCoordinator.pause(2 ** PAUSED_OPERATOR_SET_SPLIT);

        cheats.prank(operator);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        rewardsCoordinator.setOperatorSetSplit(operator, operatorSet, split);
    }

    // Revert when split is greater than 100%
    function testFuzz_Revert_WhenSplitGreaterThan100(address operator, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        split = uint16(bound(split, ONE_HUNDRED_IN_BIPS + 1, type(uint16).max));

        cheats.prank(operator);
        cheats.expectRevert(SplitExceedsMax.selector);
        rewardsCoordinator.setOperatorSetSplit(operator, operatorSet, split);
    }

    function testFuzz_Revert_InvalidOperatorSet(address operator, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        operatorSet.id = 2; // Invalid operator set id
        cheats.prank(operator);
        cheats.expectRevert(InvalidOperatorSet.selector);
        rewardsCoordinator.setOperatorSetSplit(operator, operatorSet, split);
    }

    function testFuzz_setOperatorSetSplit(address operator, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));

        split = uint16(bound(split, 0, ONE_HUNDRED_IN_BIPS));
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;
        uint16 oldSplit = rewardsCoordinator.getOperatorSetSplit(operator, operatorSet);

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorSetSplitBipsSet(operator, operator, operatorSet, activatedAt, oldSplit, split);
        cheats.prank(operator);
        rewardsCoordinator.setOperatorSetSplit(operator, operatorSet, split);

        assertEq(oldSplit, rewardsCoordinator.getOperatorSetSplit(operator, operatorSet), "Incorrect Operator split");
        cheats.warp(activatedAt);
        assertEq(split, rewardsCoordinator.getOperatorSetSplit(operator, operatorSet), "Incorrect Operator split");
    }

    function testFuzz_setOperatorSetSplit_UAM(address operator, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));

        // Set UAM
        cheats.prank(operator);
        permissionController.setAppointee(
            operator, defaultAppointee, address(rewardsCoordinator), IRewardsCoordinator.setOperatorSetSplit.selector
        );

        split = uint16(bound(split, 0, ONE_HUNDRED_IN_BIPS));
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;
        uint16 oldSplit = rewardsCoordinator.getOperatorSetSplit(operator, operatorSet);

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorSetSplitBipsSet(defaultAppointee, operator, operatorSet, activatedAt, oldSplit, split);
        cheats.prank(defaultAppointee);
        rewardsCoordinator.setOperatorSetSplit(operator, operatorSet, split);

        assertEq(oldSplit, rewardsCoordinator.getOperatorSetSplit(operator, operatorSet), "Incorrect Operator split");
        cheats.warp(activatedAt);
        assertEq(split, rewardsCoordinator.getOperatorSetSplit(operator, operatorSet), "Incorrect Operator split");
    }

    // Testing that the split has been initialized for the first time.
    function testFuzz_setOperatorSetSplitFirstTime(address operator, uint16 split) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        split = uint16(bound(split, 0, ONE_HUNDRED_IN_BIPS));
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;
        uint16 oldSplit = rewardsCoordinator.getOperatorSetSplit(operator, operatorSet);
        assertEq(oldSplit, defaultSplitBips, "Operator split is not Default split before Initialization");

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorSetSplitBipsSet(operator, operator, operatorSet, activatedAt, oldSplit, split);
        cheats.prank(operator);
        rewardsCoordinator.setOperatorSetSplit(operator, operatorSet, split);

        assertEq(oldSplit, rewardsCoordinator.getOperatorSetSplit(operator, operatorSet), "Incorrect Operator split");
        cheats.warp(activatedAt);
        assertEq(split, rewardsCoordinator.getOperatorSetSplit(operator, operatorSet), "Incorrect Operator split");
    }

    // Testing the split setting for a second time prior to the earlier activation.
    function testFuzz_Revert_setOperatorSetSplitSecondTimeBeforePriorActivation(
        address operator,
        uint16 firstSplit,
        uint16 secondSplit,
        uint32 warpTime
    ) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        firstSplit = uint16(bound(firstSplit, 0, ONE_HUNDRED_IN_BIPS));
        secondSplit = uint16(bound(secondSplit, 0, ONE_HUNDRED_IN_BIPS));
        warpTime = uint32(bound(warpTime, uint32(block.timestamp), uint32(block.timestamp) + activationDelay));

        // Setting First Split
        cheats.prank(operator);
        rewardsCoordinator.setOperatorSetSplit(operator, operatorSet, firstSplit);
        // Warping to time before activation of First split
        cheats.warp(warpTime);

        // Trying to set Second Split
        cheats.prank(operator);
        cheats.expectRevert(PreviousSplitPending.selector);
        rewardsCoordinator.setOperatorSetSplit(operator, operatorSet, secondSplit);
    }

    // Testing the split setting for a second time after earlier activation.
    function testFuzz_setOperatorSetSplitSecondTimeAfterPriorActivation(
        address operator,
        uint16 firstSplit,
        uint16 secondSplit,
        uint32 warpTime
    ) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(0));
        firstSplit = uint16(bound(firstSplit, 0, ONE_HUNDRED_IN_BIPS));
        secondSplit = uint16(bound(secondSplit, 0, ONE_HUNDRED_IN_BIPS));
        warpTime = uint32(bound(warpTime, uint32(block.timestamp) + activationDelay + 1, type(uint32).max - activationDelay));

        // Setting First Split
        cheats.prank(operator);
        rewardsCoordinator.setOperatorSetSplit(operator, operatorSet, firstSplit);
        // Warping to time after activation of First split
        cheats.warp(warpTime);
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;

        // Setting Second Split
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorSetSplitBipsSet(operator, operator, operatorSet, activatedAt, firstSplit, secondSplit);
        cheats.prank(operator);
        rewardsCoordinator.setOperatorSetSplit(operator, operatorSet, secondSplit);

        assertEq(firstSplit, rewardsCoordinator.getOperatorSetSplit(operator, operatorSet), "Incorrect Operator split");
        cheats.warp(activatedAt);
        assertEq(secondSplit, rewardsCoordinator.getOperatorSetSplit(operator, operatorSet), "Incorrect Operator split");
    }
}

contract RewardsCoordinatorUnitTests_createAVSRewardsSubmission is RewardsCoordinatorUnitTests {
    // Revert when paused
    function test_Revert_WhenPaused() public {
        cheats.prank(pauser);
        rewardsCoordinator.pause(2 ** PAUSED_AVS_REWARDS_SUBMISSION);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        RewardsSubmission[] memory rewardsSubmissions;
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert from reentrancy
    function test_Revert_WhenReentrancy(uint amount) public {
        amount = bound(amount, 1, 1e38 - 1);
        Reenterer reenterer = new Reenterer();

        reenterer.prepareReturnData(abi.encode(amount));

        address targetToUse = address(rewardsCoordinator);
        uint msgValueToUse = 0;

        _deployMockRewardTokens(address(this), 1);

        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(reenterer)),
            amount: amount,
            startTimestamp: uint32(block.timestamp),
            duration: CALCULATION_INTERVAL_SECONDS
        });

        bytes memory calldataToUse = abi.encodeWithSelector(RewardsCoordinator.createAVSRewardsSubmission.selector, rewardsSubmissions);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        cheats.expectRevert();
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert with 0 length strats and multipliers
    function testFuzz_Revert_WhenEmptyStratsAndMultipliers(address avs, uint startTimestamp, uint duration, uint amount)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp + uint(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        StrategyAndMultiplier[] memory emptyStratsAndMultipliers;
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: emptyStratsAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(InputArrayLengthZero.selector);
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert when amount > 1e38-1
    function testFuzz_Revert_AmountTooLarge(address avs, uint startTimestamp, uint duration, uint amount)
        public
        filterFuzzedAddressInputs(avs)
    {
        // 1. Bound fuzz inputs
        amount = bound(amount, 1e38, type(uint).max);
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", amount, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp + uint(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. Call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(AmountExceedsMax.selector);
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    function testFuzz_Revert_WhenDuplicateStrategies(address avs, uint startTimestamp, uint duration, uint amount)
        public
        filterFuzzedAddressInputs(avs)
    {
        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp + uint(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        StrategyAndMultiplier[] memory dupStratsAndMultipliers = new StrategyAndMultiplier[](2);
        dupStratsAndMultipliers[0] = defaultStrategyAndMultipliers[0];
        dupStratsAndMultipliers[1] = defaultStrategyAndMultipliers[0];
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: dupStratsAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(StrategiesNotInAscendingOrder.selector);
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert with exceeding max duration
    function testFuzz_Revert_WhenExceedingMaxDuration(address avs, uint startTimestamp, uint duration, uint amount)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, MAX_REWARDS_DURATION + 1, type(uint32).max);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp + uint(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(DurationExceedsMax.selector);
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert with invalid interval seconds
    function testFuzz_Revert_WhenInvalidIntervalSeconds(address avs, uint startTimestamp, uint duration, uint amount)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        cheats.assume(duration % CALCULATION_INTERVAL_SECONDS != 0);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp + uint(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(InvalidDurationRemainder.selector);
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert when duration is 0
    function testFuzz_Revert_WhenDurationIsZero(address avs, uint startTimestamp, uint amount) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp + uint(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: 0
        });

        // 3. call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(DurationIsZero.selector);
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert with retroactive rewards enabled and set too far in past
    // - either before genesis rewards timestamp
    // - before max retroactive length
    function testFuzz_Revert_WhenRewardsSubmissionTooStale(
        uint fuzzBlockTimestamp,
        address avs,
        uint startTimestamp,
        uint duration,
        uint amount
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        fuzzBlockTimestamp = bound(fuzzBlockTimestamp, uint(MAX_RETROACTIVE_LENGTH), block.timestamp);
        cheats.warp(fuzzBlockTimestamp);

        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp =
            bound(startTimestamp, 0, uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) - 1);
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(StartTimestampTooFarInPast.selector);
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert with start timestamp past max future length
    function testFuzz_Revert_WhenRewardsSubmissionTooFarInFuture(address avs, uint startTimestamp, uint duration, uint amount)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp =
            bound(startTimestamp, block.timestamp + uint(MAX_FUTURE_LENGTH) + 1 + CALCULATION_INTERVAL_SECONDS, type(uint32).max);
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(StartTimestampTooFarInFuture.selector);
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    // Revert with non whitelisted strategy
    function testFuzz_Revert_WhenInvalidStrategy(address avs, uint startTimestamp, uint duration, uint amount)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp + uint(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        defaultStrategyAndMultipliers[0].strategy = IStrategy(address(999));
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected event emitted
        cheats.prank(avs);
        cheats.expectRevert(StrategyNotWhitelisted.selector);
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    /**
     * @notice test a single rewards submission asserting for the following
     * - correct event emitted
     * - submission nonce incrementation by 1, and rewards submission hash being set in storage.
     * - rewards submission hash being set in storage
     * - token balance before and after of avs and rewardsCoordinator
     */
    function testFuzz_createAVSRewardsSubmission_SingleSubmission(address avs, uint startTimestamp, uint duration, uint amount)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp + uint(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected event emitted
        uint avsBalanceBefore = rewardToken.balanceOf(avs);
        uint rewardsCoordinatorBalanceBefore = rewardToken.balanceOf(address(rewardsCoordinator));

        cheats.startPrank(avs);
        rewardToken.approve(address(rewardsCoordinator), amount);
        uint currSubmissionNonce = rewardsCoordinator.submissionNonce(avs);
        bytes32 rewardsSubmissionHash = keccak256(abi.encode(avs, currSubmissionNonce, rewardsSubmissions[0]));

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit AVSRewardsSubmissionCreated(avs, currSubmissionNonce, rewardsSubmissionHash, rewardsSubmissions[0]);
        rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
        cheats.stopPrank();

        assertTrue(rewardsCoordinator.isAVSRewardsSubmissionHash(avs, rewardsSubmissionHash), "rewards submission hash not submitted");
        assertEq(currSubmissionNonce + 1, rewardsCoordinator.submissionNonce(avs), "submission nonce not incremented");
        assertEq(avsBalanceBefore - amount, rewardToken.balanceOf(avs), "AVS balance not decremented by amount of rewards submission");
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
    function testFuzz_createAVSRewardsSubmission_MultipleSubmissions(FuzzAVSRewardsSubmission memory param, uint numSubmissions)
        public
        filterFuzzedAddressInputs(param.avs)
    {
        numSubmissions = bound(numSubmissions, 2, 10);
        cheats.assume(param.avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](numSubmissions);
        bytes32[] memory rewardsSubmissionHashes = new bytes32[](numSubmissions);
        uint startSubmissionNonce = rewardsCoordinator.submissionNonce(param.avs);
        _deployMockRewardTokens(param.avs, numSubmissions);

        uint[] memory avsBalancesBefore = _getBalanceForTokens(rewardTokens, param.avs);
        uint[] memory rewardsCoordinatorBalancesBefore = _getBalanceForTokens(rewardTokens, address(rewardsCoordinator));
        uint[] memory amounts = new uint[](numSubmissions);

        // Create multiple rewards submissions and their expected event
        for (uint i = 0; i < numSubmissions; ++i) {
            // 1. Bound fuzz inputs to valid ranges and amounts using randSeed for each
            param.amount = bound(param.amount + i, 1, mockTokenInitialSupply);
            amounts[i] = param.amount;
            param.duration = bound(param.duration + i, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
            param.duration = param.duration - (param.duration % CALCULATION_INTERVAL_SECONDS);
            param.startTimestamp = bound(
                param.startTimestamp + i,
                uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH))
                    + CALCULATION_INTERVAL_SECONDS - 1,
                block.timestamp + uint(MAX_FUTURE_LENGTH)
            );
            param.startTimestamp = param.startTimestamp - (param.startTimestamp % CALCULATION_INTERVAL_SECONDS);

            // 2. Create rewards submission input param
            RewardsSubmission memory rewardsSubmission = RewardsSubmission({
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

        for (uint i = 0; i < numSubmissions; ++i) {
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

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        RewardsSubmission[] memory rewardsSubmissions;
        rewardsCoordinator.createRewardsForAllSubmission(rewardsSubmissions);
    }

    // Revert from reentrancy
    function test_Revert_WhenReentrancy(uint amount) public {
        Reenterer reenterer = new Reenterer();

        reenterer.prepareReturnData(abi.encode(amount));

        address targetToUse = address(rewardsCoordinator);
        uint msgValueToUse = 0;

        _deployMockRewardTokens(address(this), 1);

        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(reenterer)),
            amount: amount,
            startTimestamp: uint32(block.timestamp),
            duration: CALCULATION_INTERVAL_SECONDS
        });

        bytes memory calldataToUse = abi.encodeWithSelector(RewardsCoordinator.createAVSRewardsSubmission.selector, rewardsSubmissions);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        cheats.prank(rewardsForAllSubmitter);
        cheats.expectRevert();
        rewardsCoordinator.createRewardsForAllSubmission(rewardsSubmissions);
    }

    function testFuzz_Revert_WhenNotRewardsForAllSubmitter(address invalidSubmitter) public filterFuzzedAddressInputs(invalidSubmitter) {
        cheats.assume(invalidSubmitter != rewardsForAllSubmitter);

        cheats.expectRevert(UnauthorizedCaller.selector);
        RewardsSubmission[] memory rewardsSubmissions;
        rewardsCoordinator.createRewardsForAllSubmission(rewardsSubmissions);
    }

    // Revert when duration is 0
    function testFuzz_Revert_WhenDurationIsZero(uint startTimestamp, uint amount) public {
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, rewardsForAllSubmitter);
        amount = bound(amount, 1, mockTokenInitialSupply);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp + uint(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: 0
        });

        // 3. call createRewardsForAllSubmission() with expected revert
        cheats.prank(rewardsForAllSubmitter);
        cheats.expectRevert(DurationIsZero.selector);
        rewardsCoordinator.createRewardsForAllSubmission(rewardsSubmissions);
    }

    /**
     * @notice test a single rewards submission asserting for the following
     * - correct event emitted
     * - submission nonce incrementation by 1, and rewards submission hash being set in storage.
     * - rewards submission hash being set in storage
     * - token balance before and after of RewardsForAllSubmitter and rewardsCoordinator
     */
    function testFuzz_createRewardsForAllSubmission_SingleSubmission(uint startTimestamp, uint duration, uint amount) public {
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, rewardsForAllSubmitter);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp + uint(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected event emitted
        uint submitterBalanceBefore = rewardToken.balanceOf(rewardsForAllSubmitter);
        uint rewardsCoordinatorBalanceBefore = rewardToken.balanceOf(address(rewardsCoordinator));

        cheats.startPrank(rewardsForAllSubmitter);
        rewardToken.approve(address(rewardsCoordinator), amount);
        uint currSubmissionNonce = rewardsCoordinator.submissionNonce(rewardsForAllSubmitter);
        bytes32 rewardsSubmissionHash = keccak256(abi.encode(rewardsForAllSubmitter, currSubmissionNonce, rewardsSubmissions[0]));

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit RewardsSubmissionForAllCreated(rewardsForAllSubmitter, currSubmissionNonce, rewardsSubmissionHash, rewardsSubmissions[0]);
        rewardsCoordinator.createRewardsForAllSubmission(rewardsSubmissions);
        cheats.stopPrank();

        assertTrue(
            rewardsCoordinator.isRewardsSubmissionForAllHash(rewardsForAllSubmitter, rewardsSubmissionHash),
            "rewards submission hash not submitted"
        );
        assertEq(currSubmissionNonce + 1, rewardsCoordinator.submissionNonce(rewardsForAllSubmitter), "submission nonce not incremented");
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
    function testFuzz_createRewardsForAllSubmission_MultipleSubmissions(FuzzAVSRewardsSubmission memory param, uint numSubmissions)
        public
    {
        numSubmissions = bound(numSubmissions, 2, 10);
        cheats.prank(rewardsCoordinator.owner());

        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](numSubmissions);
        bytes32[] memory rewardsSubmissionHashes = new bytes32[](numSubmissions);
        uint startSubmissionNonce = rewardsCoordinator.submissionNonce(rewardsForAllSubmitter);
        _deployMockRewardTokens(rewardsForAllSubmitter, numSubmissions);

        uint[] memory submitterBalancesBefore = _getBalanceForTokens(rewardTokens, rewardsForAllSubmitter);
        uint[] memory rewardsCoordinatorBalancesBefore = _getBalanceForTokens(rewardTokens, address(rewardsCoordinator));
        uint[] memory amounts = new uint[](numSubmissions);

        // Create multiple rewards submissions and their expected event
        for (uint i = 0; i < numSubmissions; ++i) {
            // 1. Bound fuzz inputs to valid ranges and amounts using randSeed for each
            param.amount = bound(param.amount + i, 1, mockTokenInitialSupply);
            amounts[i] = param.amount;
            param.duration = bound(param.duration + i, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
            param.duration = param.duration - (param.duration % CALCULATION_INTERVAL_SECONDS);
            param.startTimestamp = bound(
                param.startTimestamp + i,
                uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH))
                    + CALCULATION_INTERVAL_SECONDS - 1,
                block.timestamp + uint(MAX_FUTURE_LENGTH)
            );
            param.startTimestamp = param.startTimestamp - (param.startTimestamp % CALCULATION_INTERVAL_SECONDS);

            // 2. Create rewards submission input param
            RewardsSubmission memory rewardsSubmission = RewardsSubmission({
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
                rewardsForAllSubmitter, startSubmissionNonce + i, rewardsSubmissionHashes[i], rewardsSubmissions[i]
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

        for (uint i = 0; i < numSubmissions; ++i) {
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

contract RewardsCoordinatorUnitTests_createRewardsForAllEarners is RewardsCoordinatorUnitTests {
    // Revert when paused
    function test_Revert_WhenPaused() public {
        cheats.prank(pauser);
        rewardsCoordinator.pause(2 ** PAUSED_REWARD_ALL_STAKERS_AND_OPERATORS);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        RewardsSubmission[] memory rewardsSubmissions;
        rewardsCoordinator.createRewardsForAllEarners(rewardsSubmissions);
    }

    // Revert from reentrancy
    function test_Revert_WhenReentrancy(uint amount) public {
        Reenterer reenterer = new Reenterer();

        reenterer.prepareReturnData(abi.encode(amount));

        address targetToUse = address(rewardsCoordinator);
        uint msgValueToUse = 0;

        _deployMockRewardTokens(address(this), 1);

        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(reenterer)),
            amount: amount,
            startTimestamp: uint32(block.timestamp),
            duration: CALCULATION_INTERVAL_SECONDS
        });

        bytes memory calldataToUse = abi.encodeWithSelector(RewardsCoordinator.createAVSRewardsSubmission.selector, rewardsSubmissions);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        cheats.prank(rewardsForAllSubmitter);
        cheats.expectRevert();
        rewardsCoordinator.createRewardsForAllEarners(rewardsSubmissions);
    }

    function testFuzz_Revert_WhenNotRewardsForAllSubmitter(address invalidSubmitter) public filterFuzzedAddressInputs(invalidSubmitter) {
        cheats.assume(invalidSubmitter != rewardsForAllSubmitter);

        cheats.expectRevert(UnauthorizedCaller.selector);
        RewardsSubmission[] memory rewardsSubmissions;
        rewardsCoordinator.createRewardsForAllEarners(rewardsSubmissions);
    }

    // Revert when duration is 0
    function testFuzz_Revert_WhenDurationIsZero(uint startTimestamp, uint amount) public {
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, rewardsForAllSubmitter);
        amount = bound(amount, 1, mockTokenInitialSupply);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp + uint(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: 0
        });

        // 3. call createRewardsForAllEarners() with expected revert
        cheats.prank(rewardsForAllSubmitter);
        cheats.expectRevert(DurationIsZero.selector);
        rewardsCoordinator.createRewardsForAllEarners(rewardsSubmissions);
    }

    /**
     * @notice test a single rewards submission asserting for the following
     * - correct event emitted
     * - submission nonce incrementation by 1, and rewards submission hash being set in storage.
     * - rewards submission hash being set in storage
     * - token balance before and after of RewardsForAllSubmitter and rewardsCoordinator
     */
    function testFuzz_createRewardsForAllSubmission_SingleSubmission(uint startTimestamp, uint duration, uint amount) public {
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, rewardsForAllSubmitter);
        amount = bound(amount, 1, mockTokenInitialSupply);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp + uint(MAX_FUTURE_LENGTH)
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create rewards submission input param
        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](1);
        rewardsSubmissions[0] = RewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            amount: amount,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration)
        });

        // 3. call createAVSRewardsSubmission() with expected event emitted
        uint submitterBalanceBefore = rewardToken.balanceOf(rewardsForAllSubmitter);
        uint rewardsCoordinatorBalanceBefore = rewardToken.balanceOf(address(rewardsCoordinator));

        cheats.startPrank(rewardsForAllSubmitter);
        rewardToken.approve(address(rewardsCoordinator), amount);
        uint currSubmissionNonce = rewardsCoordinator.submissionNonce(rewardsForAllSubmitter);
        bytes32 rewardsSubmissionHash = keccak256(abi.encode(rewardsForAllSubmitter, currSubmissionNonce, rewardsSubmissions[0]));

        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit RewardsSubmissionForAllEarnersCreated(
            rewardsForAllSubmitter, currSubmissionNonce, rewardsSubmissionHash, rewardsSubmissions[0]
        );
        rewardsCoordinator.createRewardsForAllEarners(rewardsSubmissions);
        cheats.stopPrank();

        assertTrue(
            rewardsCoordinator.isRewardsSubmissionForAllEarnersHash(rewardsForAllSubmitter, rewardsSubmissionHash),
            "rewards submission hash not submitted"
        );
        assertEq(currSubmissionNonce + 1, rewardsCoordinator.submissionNonce(rewardsForAllSubmitter), "submission nonce not incremented");
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
    function testFuzz_createRewardsForAllSubmission_MultipleSubmissions(FuzzAVSRewardsSubmission memory param, uint numSubmissions)
        public
    {
        numSubmissions = bound(numSubmissions, 2, 10);
        cheats.prank(rewardsCoordinator.owner());

        RewardsSubmission[] memory rewardsSubmissions = new RewardsSubmission[](numSubmissions);
        bytes32[] memory rewardsSubmissionHashes = new bytes32[](numSubmissions);
        uint startSubmissionNonce = rewardsCoordinator.submissionNonce(rewardsForAllSubmitter);
        _deployMockRewardTokens(rewardsForAllSubmitter, numSubmissions);

        uint[] memory submitterBalancesBefore = _getBalanceForTokens(rewardTokens, rewardsForAllSubmitter);
        uint[] memory rewardsCoordinatorBalancesBefore = _getBalanceForTokens(rewardTokens, address(rewardsCoordinator));
        uint[] memory amounts = new uint[](numSubmissions);

        // Create multiple rewards submissions and their expected event
        for (uint i = 0; i < numSubmissions; ++i) {
            // 1. Bound fuzz inputs to valid ranges and amounts using randSeed for each
            param.amount = bound(param.amount + i, 1, mockTokenInitialSupply);
            amounts[i] = param.amount;
            param.duration = bound(param.duration + i, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
            param.duration = param.duration - (param.duration % CALCULATION_INTERVAL_SECONDS);
            param.startTimestamp = bound(
                param.startTimestamp + i,
                uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH))
                    + CALCULATION_INTERVAL_SECONDS - 1,
                block.timestamp + uint(MAX_FUTURE_LENGTH)
            );
            param.startTimestamp = param.startTimestamp - (param.startTimestamp % CALCULATION_INTERVAL_SECONDS);

            // 2. Create rewards submission input param
            RewardsSubmission memory rewardsSubmission = RewardsSubmission({
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
            emit RewardsSubmissionForAllEarnersCreated(
                rewardsForAllSubmitter, startSubmissionNonce + i, rewardsSubmissionHashes[i], rewardsSubmissions[i]
            );
        }

        // 4. call createAVSRewardsSubmission()
        cheats.prank(rewardsForAllSubmitter);
        rewardsCoordinator.createRewardsForAllEarners(rewardsSubmissions);

        // 5. Check for submissionNonce() and rewardsSubmissionHashes being set
        assertEq(
            startSubmissionNonce + numSubmissions,
            rewardsCoordinator.submissionNonce(rewardsForAllSubmitter),
            "submission nonce not incremented properly"
        );

        for (uint i = 0; i < numSubmissions; ++i) {
            assertTrue(
                rewardsCoordinator.isRewardsSubmissionForAllEarnersHash(rewardsForAllSubmitter, rewardsSubmissionHashes[i]),
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

contract RewardsCoordinatorUnitTests_createOperatorDirectedAVSRewardsSubmission is RewardsCoordinatorUnitTests {
    // used for stack too deep
    struct FuzzOperatorDirectedAVSRewardsSubmission {
        address avs;
        uint startTimestamp;
        uint duration;
    }

    OperatorReward[] defaultOperatorRewards;

    function setUp() public virtual override {
        RewardsCoordinatorUnitTests.setUp();

        address[] memory operators = new address[](3);
        operators[0] = makeAddr("operator1");
        operators[1] = makeAddr("operator2");
        operators[2] = makeAddr("operator3");
        operators = _sortAddressArrayAsc(operators);

        defaultOperatorRewards.push(OperatorReward(operators[0], 1e18));
        defaultOperatorRewards.push(OperatorReward(operators[1], 2e18));
        defaultOperatorRewards.push(OperatorReward(operators[2], 3e18));

        // Set the timestamp to when Rewards v2 will realisticly go out (i.e 6 months)
        cheats.warp(GENESIS_REWARDS_TIMESTAMP + 168 days);
    }

    /// @dev Sort to ensure that the array is in ascending order for addresses
    function _sortAddressArrayAsc(address[] memory arr) internal pure returns (address[] memory) {
        uint l = arr.length;
        for (uint i = 0; i < l; i++) {
            for (uint j = i + 1; j < l; j++) {
                if (arr[i] > arr[j]) {
                    address temp = arr[i];
                    arr[i] = arr[j];
                    arr[j] = temp;
                }
            }
        }
        return arr;
    }

    function _getTotalRewardsAmount(OperatorReward[] memory operatorRewards) internal pure returns (uint) {
        uint totalAmount = 0;
        for (uint i = 0; i < operatorRewards.length; ++i) {
            totalAmount += operatorRewards[i].amount;
        }
        return totalAmount;
    }

    // Revert when paused
    function test_Revert_WhenPaused() public {
        cheats.prank(pauser);
        rewardsCoordinator.pause(2 ** PAUSED_OPERATOR_DIRECTED_AVS_REWARDS_SUBMISSION);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions;
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(address(this), operatorDirectedRewardsSubmissions);
    }

    // Revert from reentrancy
    function testFuzz_Revert_WhenReentrancy(uint startTimestamp, uint duration) public {
        // 1. Bound fuzz inputs to valid ranges and amounts
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Deploy Reenterer
        Reenterer reenterer = new Reenterer();

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(reenterer)),
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        address targetToUse = address(rewardsCoordinator);
        uint msgValueToUse = 0;
        bytes memory calldataToUse = abi.encodeWithSelector(
            RewardsCoordinator.createOperatorDirectedAVSRewardsSubmission.selector, address(reenterer), operatorDirectedRewardsSubmissions
        );
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse);

        cheats.prank(address(reenterer));
        cheats.expectRevert("ReentrancyGuard: reentrant call");
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(address(reenterer), operatorDirectedRewardsSubmissions);
    }

    // Revert with 0 length strats and multipliers
    function testFuzz_Revert_WhenEmptyStratsAndMultipliers(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        StrategyAndMultiplier[] memory emptyStratsAndMultipliers;
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: emptyStratsAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(InputArrayLengthZero.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    // Revert with 0 length operator rewards
    function testFuzz_Revert_WhenEmptyOperatorRewards(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        OperatorReward[] memory emptyOperatorRewards;
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: emptyOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(InputArrayLengthZero.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    // Revert when operator is zero address
    function testFuzz_Revert_WhenOperatorIsZeroAddress(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        defaultOperatorRewards[0].operator = address(0);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(InvalidAddressZero.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    // Revert when duplicate operators
    function testFuzz_Revert_WhenDuplicateOperators(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        OperatorReward[] memory dupOperatorRewards = new OperatorReward[](2);
        dupOperatorRewards[0] = defaultOperatorRewards[0];
        dupOperatorRewards[1] = defaultOperatorRewards[0];
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: dupOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(OperatorsNotInAscendingOrder.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    // Revert when operator amount is zero
    function testFuzz_Revert_WhenOperatorAmountIsZero(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        defaultOperatorRewards[0].amount = 0;
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(AmountIsZero.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    // Revert when operator amount is zero
    function testFuzz_Revert_TotalAmountTooLarge(address avs, uint startTimestamp, uint duration, uint amount)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        amount = bound(amount, 1e38, type(uint).max - 5e18);
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        defaultOperatorRewards[0].amount = amount;
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(AmountExceedsMax.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    // Revert with exceeding max duration
    function testFuzz_Revert_WhenExceedingMaxDuration(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, MAX_REWARDS_DURATION + 1, type(uint32).max);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(DurationExceedsMax.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    // Revert with invalid interval seconds
    function testFuzz_Revert_WhenInvalidIntervalSeconds(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        cheats.assume(duration % CALCULATION_INTERVAL_SECONDS != 0);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(InvalidDurationRemainder.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    // Revert when duration is 0
    function testFuzz_Revert_WhenDurationIsZero(address avs, uint startTimestamp) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: 0,
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(DurationIsZero.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    // Revert with invalid interval start timestamp
    function testFuzz_Revert_WhenInvalidIntervalStartTimestamp(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        cheats.assume(startTimestamp % CALCULATION_INTERVAL_SECONDS != 0);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(InvalidStartTimestampRemainder.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    // Revert with retroactive rewards enabled and set too far in past
    // - either before genesis rewards timestamp
    // - before max retroactive length
    function testFuzz_Revert_WhenRewardsSubmissionTooStale(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(startTimestamp, 0, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH - 1);
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(StartTimestampTooFarInPast.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    // Revert when not retroactive
    function testFuzz_Revert_WhenRewardsSubmissionNotRetroactive(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(startTimestamp, block.timestamp - duration + CALCULATION_INTERVAL_SECONDS, type(uint32).max - duration);
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(SubmissionNotRetroactive.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    // Revert with non whitelisted strategy
    function testFuzz_Revert_WhenInvalidStrategy(address avs, uint startTimestamp, uint duration) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        defaultStrategyAndMultipliers[0].strategy = IStrategy(address(999));
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(StrategyNotWhitelisted.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    // Revert when duplicate strategies
    function testFuzz_Revert_WhenDuplicateStrategies(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        StrategyAndMultiplier[] memory dupStratsAndMultipliers = new StrategyAndMultiplier[](2);
        dupStratsAndMultipliers[0] = defaultStrategyAndMultipliers[0];
        dupStratsAndMultipliers[1] = defaultStrategyAndMultipliers[0];
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: dupStratsAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(StrategiesNotInAscendingOrder.selector);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
    }

    /**
     * @notice test a single rewards submission asserting for the following
     * - correct event emitted
     * - submission nonce incrementation by 1, and rewards submission hash being set in storage.
     * - rewards submission hash being set in storage
     * - token balance before and after of avs and rewardsCoordinator
     */
    function testFuzz_createOperatorDirectedAVSRewardsSubmission_SingleSubmission(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected event emitted
        uint avsBalanceBefore = rewardToken.balanceOf(avs);
        uint rewardsCoordinatorBalanceBefore = rewardToken.balanceOf(address(rewardsCoordinator));

        cheats.startPrank(avs);
        uint amount = _getTotalRewardsAmount(defaultOperatorRewards);
        rewardToken.approve(address(rewardsCoordinator), amount);
        uint currSubmissionNonce = rewardsCoordinator.submissionNonce(avs);
        bytes32 rewardsSubmissionHash = keccak256(abi.encode(avs, currSubmissionNonce, operatorDirectedRewardsSubmissions[0]));
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorDirectedAVSRewardsSubmissionCreated(
            avs, avs, rewardsSubmissionHash, currSubmissionNonce, operatorDirectedRewardsSubmissions[0]
        );
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
        cheats.stopPrank();

        assertTrue(
            rewardsCoordinator.isOperatorDirectedAVSRewardsSubmissionHash(avs, rewardsSubmissionHash),
            "rewards submission hash not submitted"
        );
        assertEq(currSubmissionNonce + 1, rewardsCoordinator.submissionNonce(avs), "submission nonce not incremented");
        assertEq(avsBalanceBefore - amount, rewardToken.balanceOf(avs), "AVS balance not decremented by amount of rewards submission");
        assertEq(
            rewardsCoordinatorBalanceBefore + amount,
            rewardToken.balanceOf(address(rewardsCoordinator)),
            "RewardsCoordinator balance not incremented by amount of rewards submission"
        );
    }

    /**
     * @notice Same test as above, uses UAM
     * - correct event emitted
     * - submission nonce incrementation by 1, and rewards submission hash being set in storage.
     * - rewards submission hash being set in storage
     * - token balance before and after of avs and rewardsCoordinator
     */
    function testFuzz_createOperatorDirectedAVSRewardsSubmission_SingleSubmission_UAM(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        // Set UAM
        cheats.prank(avs);
        permissionController.setAppointee(
            avs, defaultAppointee, address(rewardsCoordinator), IRewardsCoordinator.createOperatorDirectedAVSRewardsSubmission.selector
        );

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, defaultAppointee);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions =
            new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedAVSRewardsSubmission() with expected event emitted
        uint submitterBalanceBefore = rewardToken.balanceOf(defaultAppointee);
        uint rewardsCoordinatorBalanceBefore = rewardToken.balanceOf(address(rewardsCoordinator));

        cheats.startPrank(defaultAppointee);
        uint amount = _getTotalRewardsAmount(defaultOperatorRewards);
        rewardToken.approve(address(rewardsCoordinator), amount);
        uint currSubmissionNonce = rewardsCoordinator.submissionNonce(avs);
        bytes32 rewardsSubmissionHash = keccak256(abi.encode(avs, currSubmissionNonce, operatorDirectedRewardsSubmissions[0]));
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorDirectedAVSRewardsSubmissionCreated(
            defaultAppointee, avs, rewardsSubmissionHash, currSubmissionNonce, operatorDirectedRewardsSubmissions[0]
        );
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avs, operatorDirectedRewardsSubmissions);
        cheats.stopPrank();

        assertTrue(
            rewardsCoordinator.isOperatorDirectedAVSRewardsSubmissionHash(avs, rewardsSubmissionHash),
            "rewards submission hash not submitted"
        );
        assertEq(currSubmissionNonce + 1, rewardsCoordinator.submissionNonce(avs), "submission nonce not incremented");
        assertEq(
            submitterBalanceBefore - amount,
            rewardToken.balanceOf(defaultAppointee),
            "Submitter balance not decremented by amount of rewards submission"
        );
        assertEq(
            rewardsCoordinatorBalanceBefore + amount,
            rewardToken.balanceOf(address(rewardsCoordinator)),
            "RewardsCoordinator balance not incremented by amount of rewards submission"
        );
    }

    /**
     * @notice test a multiple rewards submission asserting for the following
     * - correct event emitted
     * - submission nonce incrementation by 1, and rewards submission hash being set in storage.
     * - rewards submission hash being set in storage
     * - token balance before and after of avs and rewardsCoordinator
     */
    function testFuzz_createOperatorDirectedAVSRewardsSubmission_MultipleSubmissions(
        FuzzOperatorDirectedAVSRewardsSubmission memory param,
        uint numSubmissions
    ) public filterFuzzedAddressInputs(param.avs) {
        numSubmissions = bound(numSubmissions, 2, 10);
        cheats.assume(param.avs != address(0));
        cheats.prank(rewardsCoordinator.owner());

        OperatorDirectedRewardsSubmission[] memory rewardsSubmissions = new OperatorDirectedRewardsSubmission[](numSubmissions);
        bytes32[] memory rewardsSubmissionHashes = new bytes32[](numSubmissions);
        uint startSubmissionNonce = rewardsCoordinator.submissionNonce(param.avs);
        _deployMockRewardTokens(param.avs, numSubmissions);

        uint[] memory avsBalancesBefore = _getBalanceForTokens(rewardTokens, param.avs);
        uint[] memory rewardsCoordinatorBalancesBefore = _getBalanceForTokens(rewardTokens, address(rewardsCoordinator));
        uint[] memory amounts = new uint[](numSubmissions);

        // Create multiple rewards submissions and their expected event
        for (uint i = 0; i < numSubmissions; ++i) {
            // 1. Bound fuzz inputs to valid ranges and amounts using randSeed for each
            amounts[i] = _getTotalRewardsAmount(defaultOperatorRewards);
            param.duration = bound(param.duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
            param.duration = param.duration - (param.duration % CALCULATION_INTERVAL_SECONDS);
            param.startTimestamp = bound(
                param.startTimestamp + i,
                uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH))
                    + CALCULATION_INTERVAL_SECONDS - 1,
                block.timestamp + uint(MAX_FUTURE_LENGTH)
            );
            param.startTimestamp = param.startTimestamp - (param.startTimestamp % CALCULATION_INTERVAL_SECONDS);

            param.duration = bound(param.duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
            param.duration = param.duration - (param.duration % CALCULATION_INTERVAL_SECONDS);
            param.startTimestamp = bound(
                param.startTimestamp,
                uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH))
                    + CALCULATION_INTERVAL_SECONDS - 1,
                block.timestamp - param.duration - 1
            );
            param.startTimestamp = param.startTimestamp - (param.startTimestamp % CALCULATION_INTERVAL_SECONDS);

            // 2. Create rewards submission input param
            OperatorDirectedRewardsSubmission memory rewardsSubmission = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
                strategiesAndMultipliers: defaultStrategyAndMultipliers,
                token: rewardTokens[i],
                operatorRewards: defaultOperatorRewards,
                startTimestamp: uint32(param.startTimestamp),
                duration: uint32(param.duration),
                description: ""
            });
            rewardsSubmissions[i] = rewardsSubmission;

            // 3. expected event emitted for this rewardsSubmission
            rewardsSubmissionHashes[i] = keccak256(abi.encode(param.avs, startSubmissionNonce + i, rewardsSubmissions[i]));
            cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
            emit OperatorDirectedAVSRewardsSubmissionCreated(
                param.avs, param.avs, rewardsSubmissionHashes[i], startSubmissionNonce + i, rewardsSubmissions[i]
            );
        }

        // 4. call createAVSRewardsSubmission()
        cheats.prank(param.avs);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(param.avs, rewardsSubmissions);

        // 5. Check for submissionNonce() and rewardsSubmissionHashes being set
        assertEq(
            startSubmissionNonce + numSubmissions,
            rewardsCoordinator.submissionNonce(param.avs),
            "submission nonce not incremented properly"
        );

        for (uint i = 0; i < numSubmissions; ++i) {
            assertTrue(
                rewardsCoordinator.isOperatorDirectedAVSRewardsSubmissionHash(param.avs, rewardsSubmissionHashes[i]),
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

contract RewardsCoordinatorUnitTests_createOperatorDirectedOperatorSetRewardsSubmission is RewardsCoordinatorUnitTests {
    OperatorSet operatorSet;
    // used for stack too deep

    struct FuzzOperatorDirectedAVSRewardsSubmission {
        address avs;
        uint startTimestamp;
        uint duration;
    }

    OperatorReward[] defaultOperatorRewards;

    function setUp() public virtual override {
        RewardsCoordinatorUnitTests.setUp();

        address[] memory operators = new address[](3);
        operators[0] = makeAddr("operator1");
        operators[1] = makeAddr("operator2");
        operators[2] = makeAddr("operator3");
        operators = _sortAddressArrayAsc(operators);

        defaultOperatorRewards.push(OperatorReward(operators[0], 1e18));
        defaultOperatorRewards.push(OperatorReward(operators[1], 2e18));
        defaultOperatorRewards.push(OperatorReward(operators[2], 3e18));

        // Set the timestamp to when Rewards v2 will realisticly go out (i.e 6 months)
        cheats.warp(GENESIS_REWARDS_TIMESTAMP + 168 days);
        operatorSet = OperatorSet(address(this), 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);
    }

    /// @dev Sort to ensure that the array is in ascending order for addresses
    function _sortAddressArrayAsc(address[] memory arr) internal pure returns (address[] memory) {
        uint l = arr.length;
        for (uint i = 0; i < l; i++) {
            for (uint j = i + 1; j < l; j++) {
                if (arr[i] > arr[j]) {
                    address temp = arr[i];
                    arr[i] = arr[j];
                    arr[j] = temp;
                }
            }
        }
        return arr;
    }

    function _getTotalRewardsAmount(OperatorReward[] memory operatorRewards) internal pure returns (uint) {
        uint totalAmount = 0;
        for (uint i = 0; i < operatorRewards.length; ++i) {
            totalAmount += operatorRewards[i].amount;
        }
        return totalAmount;
    }

    // Revert when paused
    function test_Revert_WhenPaused() public {
        cheats.prank(pauser);
        rewardsCoordinator.pause(2 ** PAUSED_OPERATOR_DIRECTED_OPERATOR_SET_REWARDS_SUBMISSION);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions;
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    function testFuzz_Revert_InvalidOperatorSet(uint32 id) public {
        cheats.assume(id != 1);
        operatorSet.id = id;
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions;
        cheats.prank(operatorSet.avs);
        cheats.expectRevert(InvalidOperatorSet.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert from reentrancy
    function testFuzz_Revert_WhenReentrancy(uint startTimestamp, uint duration) public {
        // 1. Bound fuzz inputs to valid ranges and amounts
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Deploy Reenterer
        Reenterer reenterer = new Reenterer();

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: IERC20(address(reenterer)),
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        operatorSet = OperatorSet(address(reenterer), 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);
        address targetToUse = address(rewardsCoordinator);
        uint msgValueToUse = 0;
        bytes memory calldataToUse = abi.encodeWithSelector(
            RewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission.selector, operatorSet, operatorDirectedRewardsSubmissions
        );
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse);

        cheats.prank(address(reenterer));
        cheats.expectRevert("ReentrancyGuard: reentrant call");
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert with 0 length strats and multipliers
    function testFuzz_Revert_WhenEmptyStratsAndMultipliers(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        StrategyAndMultiplier[] memory emptyStratsAndMultipliers;
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: emptyStratsAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(InputArrayLengthZero.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert with 0 length operator rewards
    function testFuzz_Revert_WhenEmptyOperatorRewards(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        OperatorReward[] memory emptyOperatorRewards;
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: emptyOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(InputArrayLengthZero.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert when operator is zero address
    function testFuzz_Revert_WhenOperatorIsZeroAddress(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        defaultOperatorRewards[0].operator = address(0);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(InvalidAddressZero.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert when duplicate operators
    function testFuzz_Revert_WhenDuplicateOperators(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        OperatorReward[] memory dupOperatorRewards = new OperatorReward[](2);
        dupOperatorRewards[0] = defaultOperatorRewards[0];
        dupOperatorRewards[1] = defaultOperatorRewards[0];
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: dupOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(OperatorsNotInAscendingOrder.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert when operator amount is zero
    function testFuzz_Revert_WhenOperatorAmountIsZero(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        defaultOperatorRewards[0].amount = 0;
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(AmountIsZero.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert when operator amount is zero
    function testFuzz_Revert_TotalAmountTooLarge(address avs, uint startTimestamp, uint duration, uint amount)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        amount = bound(amount, 1e38, type(uint).max - 5e18);
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        defaultOperatorRewards[0].amount = amount;
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(AmountExceedsMax.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert with exceeding max duration
    function testFuzz_Revert_WhenExceedingMaxDuration(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, MAX_REWARDS_DURATION + 1, type(uint32).max);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(DurationExceedsMax.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert with invalid interval seconds
    function testFuzz_Revert_WhenInvalidIntervalSeconds(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        cheats.assume(duration % CALCULATION_INTERVAL_SECONDS != 0);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(InvalidDurationRemainder.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert when duration is 0
    function testFuzz_Revert_WhenDurationIsZero(address avs, uint startTimestamp) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: 0,
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(DurationIsZero.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert with invalid interval start timestamp
    function testFuzz_Revert_WhenInvalidIntervalStartTimestamp(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        cheats.assume(startTimestamp % CALCULATION_INTERVAL_SECONDS != 0);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(InvalidStartTimestampRemainder.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert with retroactive rewards enabled and set too far in past
    // - either before genesis rewards timestamp
    // - before max retroactive length
    function testFuzz_Revert_WhenRewardsSubmissionTooStale(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(startTimestamp, 0, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH - 1);
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(StartTimestampTooFarInPast.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert when not retroactive
    function testFuzz_Revert_WhenRewardsSubmissionNotRetroactive(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(startTimestamp, block.timestamp - duration + CALCULATION_INTERVAL_SECONDS, type(uint32).max - duration);
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(SubmissionNotRetroactive.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert with non whitelisted strategy
    function testFuzz_Revert_WhenInvalidStrategy(address avs, uint startTimestamp, uint duration) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        defaultStrategyAndMultipliers[0].strategy = IStrategy(address(999));
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(StrategyNotWhitelisted.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    // Revert when duplicate strategies
    function testFuzz_Revert_WhenDuplicateStrategies(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        StrategyAndMultiplier[] memory dupStratsAndMultipliers = new StrategyAndMultiplier[](2);
        dupStratsAndMultipliers[0] = defaultStrategyAndMultipliers[0];
        dupStratsAndMultipliers[1] = defaultStrategyAndMultipliers[0];
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: dupStratsAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected revert
        cheats.prank(avs);
        cheats.expectRevert(StrategiesNotInAscendingOrder.selector);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
    }

    /**
     * @notice test a single rewards submission asserting for the following
     * - correct event emitted
     * - submission nonce incrementation by 1, and rewards submission hash being set in storage.
     * - rewards submission hash being set in storage
     * - token balance before and after of avs and rewardsCoordinator
     */
    function testFuzz_createOperatorDirectedOperatorSetRewardsSubmission_SingleSubmission(address avs, uint startTimestamp, uint duration)
        public
        filterFuzzedAddressInputs(avs)
    {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, avs);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions = new OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected event emitted
        uint avsBalanceBefore = rewardToken.balanceOf(avs);
        uint rewardsCoordinatorBalanceBefore = rewardToken.balanceOf(address(rewardsCoordinator));

        cheats.startPrank(avs);
        uint amount = _getTotalRewardsAmount(defaultOperatorRewards);
        rewardToken.approve(address(rewardsCoordinator), amount);
        uint currSubmissionNonce = rewardsCoordinator.submissionNonce(avs);
        bytes32 rewardsSubmissionHash = keccak256(abi.encode(avs, currSubmissionNonce, operatorDirectedRewardsSubmissions[0]));
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorDirectedOperatorSetRewardsSubmissionCreated(
            avs, rewardsSubmissionHash, operatorSet, currSubmissionNonce, operatorDirectedRewardsSubmissions[0]
        );
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
        cheats.stopPrank();

        assertTrue(
            rewardsCoordinator.isOperatorDirectedOperatorSetRewardsSubmissionHash(avs, rewardsSubmissionHash),
            "rewards submission hash not submitted"
        );
        assertEq(currSubmissionNonce + 1, rewardsCoordinator.submissionNonce(avs), "submission nonce not incremented");
        assertEq(avsBalanceBefore - amount, rewardToken.balanceOf(avs), "AVS balance not decremented by amount of rewards submission");
        assertEq(
            rewardsCoordinatorBalanceBefore + amount,
            rewardToken.balanceOf(address(rewardsCoordinator)),
            "RewardsCoordinator balance not incremented by amount of rewards submission"
        );
    }

    /**
     * @notice Same test as above, uses UAM
     * - correct event emitted
     * - submission nonce incrementation by 1, and rewards submission hash being set in storage.
     * - rewards submission hash being set in storage
     * - token balance before and after of avs and rewardsCoordinator
     */
    function testFuzz_createOperatorDirectedOperatorSetRewardsSubmission_SingleSubmission_UAM(
        address avs,
        uint startTimestamp,
        uint duration
    ) public filterFuzzedAddressInputs(avs) {
        cheats.assume(avs != address(0));

        operatorSet = OperatorSet(avs, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        // Set UAM
        cheats.prank(avs);
        permissionController.setAppointee(
            avs,
            defaultAppointee,
            address(rewardsCoordinator),
            IRewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission.selector
        );

        // 1. Bound fuzz inputs to valid ranges and amounts
        IERC20 rewardToken = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, defaultAppointee);
        duration = bound(duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
        duration = duration - (duration % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = bound(
            startTimestamp,
            uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH)) + CALCULATION_INTERVAL_SECONDS
                - 1,
            block.timestamp - duration - 1
        );
        startTimestamp = startTimestamp - (startTimestamp % CALCULATION_INTERVAL_SECONDS);

        // 2. Create operator directed rewards submission input param
        IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory operatorDirectedRewardsSubmissions =
            new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
        operatorDirectedRewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: rewardToken,
            operatorRewards: defaultOperatorRewards,
            startTimestamp: uint32(startTimestamp),
            duration: uint32(duration),
            description: ""
        });

        // 3. call createOperatorDirectedOperatorSetRewardsSubmission() with expected event emitted
        uint submitterBalanceBefore = rewardToken.balanceOf(defaultAppointee);
        uint rewardsCoordinatorBalanceBefore = rewardToken.balanceOf(address(rewardsCoordinator));

        cheats.startPrank(defaultAppointee);
        uint amount = _getTotalRewardsAmount(defaultOperatorRewards);
        rewardToken.approve(address(rewardsCoordinator), amount);
        uint currSubmissionNonce = rewardsCoordinator.submissionNonce(avs);
        bytes32 rewardsSubmissionHash = keccak256(abi.encode(avs, currSubmissionNonce, operatorDirectedRewardsSubmissions[0]));
        cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
        emit OperatorDirectedOperatorSetRewardsSubmissionCreated(
            defaultAppointee, rewardsSubmissionHash, operatorSet, currSubmissionNonce, operatorDirectedRewardsSubmissions[0]
        );
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, operatorDirectedRewardsSubmissions);
        cheats.stopPrank();

        assertTrue(
            rewardsCoordinator.isOperatorDirectedOperatorSetRewardsSubmissionHash(avs, rewardsSubmissionHash),
            "rewards submission hash not submitted"
        );
        assertEq(currSubmissionNonce + 1, rewardsCoordinator.submissionNonce(avs), "submission nonce not incremented");
        assertEq(
            submitterBalanceBefore - amount,
            rewardToken.balanceOf(defaultAppointee),
            "Submitter balance not decremented by amount of rewards submission"
        );
        assertEq(
            rewardsCoordinatorBalanceBefore + amount,
            rewardToken.balanceOf(address(rewardsCoordinator)),
            "RewardsCoordinator balance not incremented by amount of rewards submission"
        );
    }

    /**
     * @notice test a multiple rewards submission asserting for the following
     * - correct event emitted
     * - submission nonce incrementation by 1, and rewards submission hash being set in storage.
     * - rewards submission hash being set in storage
     * - token balance before and after of avs and rewardsCoordinator
     */
    function testFuzz_createOperatorDirectedOperatorSetRewardsSubmission_MultipleSubmissions(
        FuzzOperatorDirectedAVSRewardsSubmission memory param,
        uint numSubmissions
    ) public filterFuzzedAddressInputs(param.avs) {
        numSubmissions = bound(numSubmissions, 2, 10);
        cheats.assume(param.avs != address(0));

        operatorSet = OperatorSet(param.avs, 2);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);

        cheats.prank(rewardsCoordinator.owner());

        OperatorDirectedRewardsSubmission[] memory rewardsSubmissions = new OperatorDirectedRewardsSubmission[](numSubmissions);
        bytes32[] memory rewardsSubmissionHashes = new bytes32[](numSubmissions);
        uint startSubmissionNonce = rewardsCoordinator.submissionNonce(param.avs);
        _deployMockRewardTokens(param.avs, numSubmissions);

        uint[] memory avsBalancesBefore = _getBalanceForTokens(rewardTokens, param.avs);
        uint[] memory rewardsCoordinatorBalancesBefore = _getBalanceForTokens(rewardTokens, address(rewardsCoordinator));
        uint[] memory amounts = new uint[](numSubmissions);

        // Create multiple rewards submissions and their expected event
        for (uint i = 0; i < numSubmissions; ++i) {
            // 1. Bound fuzz inputs to valid ranges and amounts using randSeed for each
            amounts[i] = _getTotalRewardsAmount(defaultOperatorRewards);
            param.duration = bound(param.duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
            param.duration = param.duration - (param.duration % CALCULATION_INTERVAL_SECONDS);
            param.startTimestamp = bound(
                param.startTimestamp + i,
                uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH))
                    + CALCULATION_INTERVAL_SECONDS - 1,
                block.timestamp + uint(MAX_FUTURE_LENGTH)
            );
            param.startTimestamp = param.startTimestamp - (param.startTimestamp % CALCULATION_INTERVAL_SECONDS);

            param.duration = bound(param.duration, CALCULATION_INTERVAL_SECONDS, MAX_REWARDS_DURATION);
            param.duration = param.duration - (param.duration % CALCULATION_INTERVAL_SECONDS);
            param.startTimestamp = bound(
                param.startTimestamp,
                uint(_maxTimestamp(GENESIS_REWARDS_TIMESTAMP, uint32(block.timestamp) - MAX_RETROACTIVE_LENGTH))
                    + CALCULATION_INTERVAL_SECONDS - 1,
                block.timestamp - param.duration - 1
            );
            param.startTimestamp = param.startTimestamp - (param.startTimestamp % CALCULATION_INTERVAL_SECONDS);

            // 2. Create rewards submission input param
            OperatorDirectedRewardsSubmission memory rewardsSubmission = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
                strategiesAndMultipliers: defaultStrategyAndMultipliers,
                token: rewardTokens[i],
                operatorRewards: defaultOperatorRewards,
                startTimestamp: uint32(param.startTimestamp),
                duration: uint32(param.duration),
                description: ""
            });
            rewardsSubmissions[i] = rewardsSubmission;

            // 3. expected event emitted for this rewardsSubmission
            rewardsSubmissionHashes[i] = keccak256(abi.encode(param.avs, startSubmissionNonce + i, rewardsSubmissions[i]));
            cheats.expectEmit(true, true, true, true, address(rewardsCoordinator));
            emit OperatorDirectedOperatorSetRewardsSubmissionCreated(
                param.avs, rewardsSubmissionHashes[i], operatorSet, startSubmissionNonce + i, rewardsSubmissions[i]
            );
        }

        // 4. call createAVSRewardsSubmission()
        cheats.prank(param.avs);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, rewardsSubmissions);

        // 5. Check for submissionNonce() and rewardsSubmissionHashes being set
        assertEq(
            startSubmissionNonce + numSubmissions,
            rewardsCoordinator.submissionNonce(param.avs),
            "submission nonce not incremented properly"
        );

        for (uint i = 0; i < numSubmissions; ++i) {
            assertTrue(
                rewardsCoordinator.isOperatorDirectedOperatorSetRewardsSubmissionHash(param.avs, rewardsSubmissionHashes[i]),
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

contract RewardsCoordinatorUnitTests_submitRoot is RewardsCoordinatorUnitTests {
    // only callable by rewardsUpdater
    function testFuzz_Revert_WhenNotRewardsUpdater(address invalidRewardsUpdater) public filterFuzzedAddressInputs(invalidRewardsUpdater) {
        cheats.prank(invalidRewardsUpdater);

        cheats.expectRevert(UnauthorizedCaller.selector);
        rewardsCoordinator.submitRoot(bytes32(0), 0);
    }

    function test_Revert_WhenSubmitRootPaused() public {
        cheats.prank(pauser);
        rewardsCoordinator.pause(2 ** PAUSED_SUBMIT_ROOTS);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
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

        assertEq(expectedRootIndex, rewardsCoordinator.getDistributionRootsLength() - 1, "root not added to roots array");
        assertEq(root, rewardsCoordinator.getCurrentDistributionRoot().root, "getCurrentDistributionRoot view function failed");
        assertEq(
            root, rewardsCoordinator.getDistributionRootAtIndex(expectedRootIndex).root, "getDistributionRootAtIndex view function failed"
        );
        assertEq(activatedAt, distributionRoot.activatedAt, "activatedAt not correct");
        assertEq(root, distributionRoot.root, "root not set");
        assertEq(rewardsCalculationEndTimestamp, distributionRoot.rewardsCalculationEndTimestamp, "rewardsCalculationEndTimestamp not set");
        assertEq(
            rewardsCoordinator.currRewardsCalculationEndTimestamp(),
            rewardsCalculationEndTimestamp,
            "currRewardsCalculationEndTimestamp not set"
        );
    }

    /// @notice Submits multiple roots and checks root index from hash is correct
    function testFuzz_getRootIndexFromHash(bytes32 root, uint16 numRoots, uint index) public {
        numRoots = uint16(bound(numRoots, 1, 100));
        index = bound(index, 0, uint(numRoots - 1));

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

    /// @notice Claim against latest submitted root, rootIndex 3 using batch claim.
    /// Limit fuzz runs to speed up tests since these require reading from JSON
    /// forge-config: default.fuzz.runs = 10
    function testFuzz_processClaims_LatestRoot(bool setClaimerFor, address claimerFor) public filterFuzzedAddressInputs(claimerFor) {
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
        RewardsMerkleClaim[] memory claims = _parseAllProofs();
        RewardsMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        // Claim against root and check balances before/after, and check it matches the difference between
        // cumulative claimed and earned.
        cheats.startPrank(claimer);
        assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        uint[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
        uint[] memory earnings = _getCumulativeEarnings(claim);
        uint[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

        _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
        IRewardsCoordinator.RewardsMerkleClaim[] memory batchClaim = new IRewardsCoordinator.RewardsMerkleClaim[](1);
        batchClaim[0] = claim;
        rewardsCoordinator.processClaims(batchClaim, claimer);

        uint[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

        for (uint i = 0; i < totalClaimedBefore.length; ++i) {
            assertEq(
                earnings[i] - totalClaimedBefore[i],
                tokenBalancesAfter[i] - tokenBalancesBefore[i],
                "Token balance not incremented by earnings amount"
            );
        }

        cheats.stopPrank();
    }

    /// @notice Claim against latest submitted root, rootIndex 3
    /// Limit fuzz runs to speed up tests since these require reading from JSON
    /// forge-config: default.fuzz.runs = 10
    function testFuzz_processClaim_LatestRoot(bool setClaimerFor, address claimerFor) public filterFuzzedAddressInputs(claimerFor) {
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
        RewardsMerkleClaim[] memory claims = _parseAllProofs();
        RewardsMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        // Claim against root and check balances before/after, and check it matches the difference between
        // cumulative claimed and earned.
        cheats.startPrank(claimer);
        assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        uint[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
        uint[] memory earnings = _getCumulativeEarnings(claim);
        uint[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

        _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
        rewardsCoordinator.processClaim(claim, claimer);

        uint[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

        for (uint i = 0; i < totalClaimedBefore.length; ++i) {
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
    function testFuzz_processClaim_OldRoot(bool setClaimerFor, address claimerFor) public filterFuzzedAddressInputs(claimerFor) {
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
        RewardsMerkleClaim[] memory claims = _parseAllProofs();
        RewardsMerkleClaim memory claim = claims[0];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        // Claim against root and check balances before/after, and check it matches the difference between
        // cumulative claimed and earned.
        cheats.startPrank(claimer);
        assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        uint[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
        uint[] memory earnings = _getCumulativeEarnings(claim);
        uint[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

        _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
        rewardsCoordinator.processClaim(claim, claimer);

        uint[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

        for (uint i = 0; i < totalClaimedBefore.length; ++i) {
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
    function testFuzz_processClaim_Sequential(bool setClaimerFor, address claimerFor) public filterFuzzedAddressInputs(claimerFor) {
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
        RewardsMerkleClaim[] memory claims = _parseAllProofs();
        RewardsMerkleClaim memory claim = claims[0];

        // 1. Claim against first root
        {
            uint32 rootIndex = claim.rootIndex;
            IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
            cheats.warp(distributionRoot.activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

            uint[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint[] memory earnings = _getCumulativeEarnings(claim);
            uint[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

            uint[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint i = 0; i < totalClaimedBefore.length; ++i) {
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

            uint[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint[] memory earnings = _getCumulativeEarnings(claim);
            uint[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

            uint[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint i = 0; i < totalClaimedBefore.length; ++i) {
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

            uint[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint[] memory earnings = _getCumulativeEarnings(claim);
            uint[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

            uint[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint i = 0; i < totalClaimedBefore.length; ++i) {
                assertEq(
                    earnings[i] - totalClaimedBefore[i],
                    tokenBalancesAfter[i] - tokenBalancesBefore[i],
                    "Token balance not incremented by earnings amount"
                );
            }

            cheats.stopPrank();
        }
    }

    function testFuzz_processClaim_Revert_WhenRootDisabled(bool setClaimerFor, address claimerFor, bytes32 root)
        public
        filterFuzzedAddressInputs(claimerFor)
    {
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
        rewardsCoordinator.submitRoot(root, 1);
        uint32 rootIndex = 0;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        rewardsCoordinator.disableRoot(rootIndex);
        cheats.stopPrank();

        cheats.warp(distributionRoot.activatedAt);

        cheats.startPrank(claimer);
        // rootIndex in claim is 0, which is disabled
        RewardsMerkleClaim memory claim;
        cheats.expectRevert(RootDisabled.selector);
        rewardsCoordinator.processClaim(claim, claimer);
        cheats.stopPrank();
    }

    /// @notice Claim against rootIndex 0 and claim again. Balances should not increment.
    /// forge-config: default.fuzz.runs = 10
    function testFuzz_processClaim_Revert_WhenReuseSameClaimAgain(bool setClaimerFor, address claimerFor)
        public
        filterFuzzedAddressInputs(claimerFor)
    {
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
        RewardsMerkleClaim[] memory claims = _parseAllProofs();
        RewardsMerkleClaim memory claim = claims[0];

        // 1. Claim against first root
        {
            uint32 rootIndex = claim.rootIndex;
            IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
            cheats.warp(distributionRoot.activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

            uint[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint[] memory earnings = _getCumulativeEarnings(claim);
            uint[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

            uint[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint i = 0; i < totalClaimedBefore.length; ++i) {
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

            cheats.expectRevert(EarningsNotGreaterThanClaimed.selector);
            rewardsCoordinator.processClaim(claim, claimer);

            cheats.stopPrank();
        }
    }

    /// @notice Claim against latest submitted root, rootIndex 3 but modify some of the leaf values
    /// forge-config: default.fuzz.runs = 10
    function testFuzz_processClaim_Revert_WhenInvalidTokenClaim(bool setClaimerFor, address claimerFor)
        public
        filterFuzzedAddressInputs(claimerFor)
    {
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
        RewardsMerkleClaim[] memory claims = _parseAllProofs();
        RewardsMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        // Modify Earnings
        claim.tokenLeaves[0].cumulativeEarnings = 1e20;
        claim.tokenLeaves[1].cumulativeEarnings = 1e20;

        // Check claim is not valid from both checkClaim() and processClaim() throwing a revert
        cheats.startPrank(claimer);

        cheats.expectRevert(InvalidClaimProof.selector);
        assertFalse(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        cheats.expectRevert(InvalidClaimProof.selector);
        rewardsCoordinator.processClaim(claim, claimer);

        cheats.stopPrank();
    }

    /// @notice Claim against latest submitted root, rootIndex 3 but modify some of the leaf values
    /// forge-config: default.fuzz.runs = 10
    function testFuzz_processClaim_Revert_WhenInvalidEarnerClaim(bool setClaimerFor, address claimerFor, address invalidEarner)
        public
        filterFuzzedAddressInputs(claimerFor)
    {
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
        RewardsMerkleClaim[] memory claims = _parseAllProofs();
        RewardsMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        // Modify Earner
        claim.earnerLeaf.earner = invalidEarner;

        // Check claim is not valid from both checkClaim() and processClaim() throwing a revert
        cheats.startPrank(claimer);

        cheats.expectRevert(InvalidClaimProof.selector);
        assertFalse(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        cheats.expectRevert(InvalidClaimProof.selector);
        rewardsCoordinator.processClaim(claim, claimer);

        cheats.stopPrank();
    }

    /// @notice Claim against latest submitted root, rootIndex 3 but write to cumulativeClaimed storage.
    /// Expects a revert when calling processClaim()
    function testFuzz_processClaim_Revert_WhenCumulativeClaimedUnderflow(bool setClaimerFor, address claimerFor)
        public
        filterFuzzedAddressInputs(claimerFor)
    {
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
        RewardsMerkleClaim[] memory claims = _parseAllProofs();
        RewardsMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        // Set cumulativeClaimed to be max uint256, should revert when attempting to claim
        stdstore.target(address(rewardsCoordinator)).sig("cumulativeClaimed(address,address)").with_key(claim.earnerLeaf.earner).with_key(
            address(claim.tokenLeaves[0].token)
        ).checked_write(type(uint).max);
        cheats.startPrank(claimer);
        cheats.expectRevert(EarningsNotGreaterThanClaimed.selector);
        rewardsCoordinator.processClaim(claim, claimer);
        cheats.stopPrank();
    }

    /// @notice Claim against latest submitted root, rootIndex 3 but with larger tokenIndex used that could pass the proofs.
    /// Expects a revert as we check for this in processClaim()
    function testFuzz_processClaim_Revert_WhenTokenIndexBitwiseAddedTo(bool setClaimerFor, address claimerFor, uint8 numShift)
        public
        filterFuzzedAddressInputs(claimerFor)
    {
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
        RewardsMerkleClaim[] memory claims = _parseAllProofs();
        RewardsMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        // Take the tokenIndex and add a significant bit so that the actual index number is increased
        // but still with the same least significant bits for valid proofs
        uint8 proofLength = uint8(claim.tokenTreeProofs[0].length);
        claim.tokenIndices[0] = claim.tokenIndices[0] | uint32(1 << (numShift + proofLength / 32));
        cheats.startPrank(claimer);
        cheats.expectRevert(InvalidTokenLeafIndex.selector);
        rewardsCoordinator.processClaim(claim, claimer);
        cheats.stopPrank();
    }

    /// @notice Claim against latest submitted root, rootIndex 3 but with larger earnerIndex used that could pass the proofs.
    /// Expects a revert as we check for this in processClaim()
    function testFuzz_processClaim_Revert_WhenEarnerIndexBitwiseAddedTo(bool setClaimerFor, address claimerFor, uint8 numShift)
        public
        filterFuzzedAddressInputs(claimerFor)
    {
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
        RewardsMerkleClaim[] memory claims = _parseAllProofs();
        RewardsMerkleClaim memory claim = claims[2];

        uint32 rootIndex = claim.rootIndex;
        IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
        cheats.warp(distributionRoot.activatedAt);

        assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

        // Take the tokenIndex and add a significant bit so that the actual index number is increased
        // but still with the same least significant bits for valid proofs
        uint8 proofLength = uint8(claim.earnerTreeProof.length);
        claim.earnerIndex = claim.earnerIndex | uint32(1 << (numShift + proofLength / 32));
        cheats.startPrank(claimer);
        cheats.expectRevert(InvalidEarnerLeafIndex.selector);
        rewardsCoordinator.processClaim(claim, claimer);
        cheats.stopPrank();
    }

    /// @notice tests with earnerIndex and tokenIndex set to max value and using alternate claim proofs
    function testFuzz_processClaim_WhenMaxEarnerIndexAndTokenIndex(bool setClaimerFor, address claimerFor)
        public
        filterFuzzedAddressInputs(claimerFor)
    {
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
        RewardsMerkleClaim[] memory claims = _parseAllProofsMaxEarnerAndLeafIndices();
        RewardsMerkleClaim memory claim = claims[0];

        // 1. Claim against first root where earner tree is full tree and earner and token index is last index of that tree height
        {
            uint32 rootIndex = claim.rootIndex;
            IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
            cheats.warp(distributionRoot.activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

            uint[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint[] memory earnings = _getCumulativeEarnings(claim);
            uint[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            // +1 since earnerIndex is 0-indexed
            // Here the earnerIndex is 7 in a full binary tree and the number of bytes32 hash proofs is 3
            assertEq(claim.earnerIndex + 1, (1 << ((claim.earnerTreeProof.length / 32))), "EarnerIndex not set to max value");
            // +1 since tokenIndex is 0-indexed
            // Here the tokenIndex is also 7 in a full binary tree and the number of bytes32 hash proofs is 3
            assertEq(claim.tokenIndices[0] + 1, (1 << ((claim.tokenTreeProofs[0].length / 32))), "TokenIndex not set to max value");
            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

            uint[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint i = 0; i < totalClaimedBefore.length; ++i) {
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
    function testFuzz_processClaim_WhenSingleTokenLeaf(bool setClaimerFor, address claimerFor)
        public
        filterFuzzedAddressInputs(claimerFor)
    {
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
        RewardsMerkleClaim[] memory claims = _parseAllProofsSingleTokenLeaf();
        RewardsMerkleClaim memory claim = claims[0];

        // 1. Claim against first root where earner tree is full tree and earner and token index is last index of that tree height
        {
            uint32 rootIndex = claim.rootIndex;
            IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
            cheats.warp(distributionRoot.activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

            uint[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint[] memory earnings = _getCumulativeEarnings(claim);
            uint[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            // Single tokenLeaf in earner's subtree, should be 0 index
            assertEq(claim.tokenIndices[0], 0, "TokenIndex should be 0");
            assertEq(claim.tokenTreeProofs[0].length, 0, "TokenTreeProof should be empty");
            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

            uint[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint i = 0; i < totalClaimedBefore.length; ++i) {
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
    function testFuzz_processClaim_WhenSingleEarnerLeaf(bool setClaimerFor, address claimerFor)
        public
        filterFuzzedAddressInputs(claimerFor)
    {
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
        RewardsMerkleClaim[] memory claims = _parseAllProofsSingleEarnerLeaf();
        RewardsMerkleClaim memory claim = claims[0];

        // 1. Claim against first root where earner tree is full tree and earner and token index is last index of that tree height
        {
            uint32 rootIndex = claim.rootIndex;
            IRewardsCoordinator.DistributionRoot memory distributionRoot = rewardsCoordinator.getDistributionRootAtIndex(rootIndex);
            cheats.warp(distributionRoot.activatedAt);

            // Claim against root and check balances before/after, and check it matches the difference between
            // cumulative claimed and earned.
            cheats.startPrank(claimer);
            assertTrue(rewardsCoordinator.checkClaim(claim), "RewardsCoordinator.checkClaim: claim not valid");

            uint[] memory totalClaimedBefore = _getCumulativeClaimed(earner, claim);
            uint[] memory earnings = _getCumulativeEarnings(claim);
            uint[] memory tokenBalancesBefore = _getClaimTokenBalances(claimer, claim);

            // Earner Leaf in merkle tree, should be 0 index
            assertEq(claim.earnerIndex, 0, "EarnerIndex should be 0");
            assertEq(claim.earnerTreeProof.length, 0, "EarnerTreeProof should be empty");
            _assertRewardsClaimedEvents(distributionRoot.root, claim, claimer);
            rewardsCoordinator.processClaim(claim, claimer);

            uint[] memory tokenBalancesAfter = _getClaimTokenBalances(claimer, claim);

            for (uint i = 0; i < totalClaimedBefore.length; ++i) {
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
    function _setAddressAsERC20(address randAddress, uint mintAmount) internal {
        cheats.etch(randAddress, mockTokenBytecode);
        ERC20Mock(randAddress).mint(address(rewardsCoordinator), mintAmount);
    }

    /// @notice parse proofs from json file and submitRoot()
    function _parseProofData(string memory filePath) internal returns (RewardsMerkleClaim memory) {
        cheats.readFile(filePath);

        string memory claimProofData = cheats.readFile(filePath);

        // Parse RewardsMerkleClaim
        merkleRoot = abi.decode(stdJson.parseRaw(claimProofData, ".Root"), (bytes32));
        earnerIndex = abi.decode(stdJson.parseRaw(claimProofData, ".EarnerIndex"), (uint32));
        earnerTreeProof = abi.decode(stdJson.parseRaw(claimProofData, ".EarnerTreeProof"), (bytes));
        proofEarner = stdJson.readAddress(claimProofData, ".EarnerLeaf.Earner");
        require(earner == proofEarner, "earner in test and json file do not match");
        earnerTokenRoot = abi.decode(stdJson.parseRaw(claimProofData, ".EarnerLeaf.EarnerTokenRoot"), (bytes32));
        uint numTokenLeaves = stdJson.readUint(claimProofData, ".TokenLeavesNum");
        uint numTokenTreeProofs = stdJson.readUint(claimProofData, ".TokenTreeProofsNum");

        TokenTreeMerkleLeaf[] memory tokenLeaves = new TokenTreeMerkleLeaf[](numTokenLeaves);
        uint32[] memory tokenIndices = new uint32[](numTokenLeaves);
        for (uint i = 0; i < numTokenLeaves; ++i) {
            string memory tokenKey = string.concat(".TokenLeaves[", cheats.toString(i), "].Token");
            string memory amountKey = string.concat(".TokenLeaves[", cheats.toString(i), "].CumulativeEarnings");
            string memory leafIndicesKey = string.concat(".LeafIndices[", cheats.toString(i), "]");

            IERC20 token = IERC20(stdJson.readAddress(claimProofData, tokenKey));
            uint cumulativeEarnings = stdJson.readUint(claimProofData, amountKey);
            tokenLeaves[i] = TokenTreeMerkleLeaf({token: token, cumulativeEarnings: cumulativeEarnings});
            tokenIndices[i] = uint32(stdJson.readUint(claimProofData, leafIndicesKey));

            /// DeployCode ERC20 to Token Address
            // deployCodeTo("ERC20PresetFixedSupply.sol", address(tokenLeaves[i].token));
            _setAddressAsERC20(address(token), cumulativeEarnings);
        }
        bytes[] memory tokenTreeProofs = new bytes[](numTokenTreeProofs);
        for (uint i = 0; i < numTokenTreeProofs; ++i) {
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

        RewardsMerkleClaim memory newClaim = RewardsMerkleClaim({
            rootIndex: rootIndex,
            earnerIndex: earnerIndex,
            earnerTreeProof: earnerTreeProof,
            earnerLeaf: EarnerTreeMerkleLeaf({earner: earner, earnerTokenRoot: earnerTokenRoot}),
            tokenIndices: tokenIndices,
            tokenTreeProofs: tokenTreeProofs,
            tokenLeaves: tokenLeaves
        });

        return newClaim;
    }

    function _parseAllProofs() internal virtual returns (RewardsMerkleClaim[] memory) {
        RewardsMerkleClaim[] memory claims = new RewardsMerkleClaim[](3);

        claims[0] = _parseProofData("src/test/test-data/rewardsCoordinator/processClaimProofs_Root1.json");
        claims[1] = _parseProofData("src/test/test-data/rewardsCoordinator/processClaimProofs_Root2.json");
        claims[2] = _parseProofData("src/test/test-data/rewardsCoordinator/processClaimProofs_Root3.json");

        return claims;
    }

    function _parseAllProofsMaxEarnerAndLeafIndices() internal virtual returns (RewardsMerkleClaim[] memory) {
        RewardsMerkleClaim[] memory claims = new RewardsMerkleClaim[](1);

        claims[0] = _parseProofData("src/test/test-data/rewardsCoordinator/processClaimProofs_MaxEarnerAndLeafIndices.json");

        return claims;
    }

    function _parseAllProofsSingleTokenLeaf() internal virtual returns (RewardsMerkleClaim[] memory) {
        RewardsMerkleClaim[] memory claims = new RewardsMerkleClaim[](1);

        claims[0] = _parseProofData("src/test/test-data/rewardsCoordinator/processClaimProofs_SingleTokenLeaf.json");

        return claims;
    }

    function _parseAllProofsSingleEarnerLeaf() internal virtual returns (RewardsMerkleClaim[] memory) {
        RewardsMerkleClaim[] memory claims = new RewardsMerkleClaim[](1);

        claims[0] = _parseProofData("src/test/test-data/rewardsCoordinator/processClaimProofs_SingleEarnerLeaf.json");

        return claims;
    }
}
