// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/core/AllocationManager.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/test/mocks/MockAVSRegistrar.sol";

contract AllocationManagerUnitTests is EigenLayerUnitTestSetup, IAllocationManagerErrors, IAllocationManagerEvents {
    using StdStyle for *;
    using SingleItemArrayLib for *;

    /// -----------------------------------------------------------------------
    /// Constants
    /// -----------------------------------------------------------------------

    /// NOTE: Raising these values directly increases cpu time for tests.
    uint256 internal constant FUZZ_MAX_ALLOCATIONS = 8;
    uint256 internal constant FUZZ_MAX_STRATS = 8;
    uint256 internal constant FUZZ_MAX_OP_SETS = 8;

    uint8 internal constant PAUSED_MODIFY_ALLOCATIONS = 0;
    uint8 internal constant PAUSED_OPERATOR_SLASHING = 1;
    uint8 internal constant PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION = 2;

    uint32 constant ASSUMED_BLOCK_TIME = 12 seconds;
    uint32 constant DEALLOCATION_DELAY = 14 days / ASSUMED_BLOCK_TIME;
    uint32 constant ALLOCATION_CONFIGURATION_DELAY = 21 days / ASSUMED_BLOCK_TIME;
    uint32 constant DEFAULT_OPERATOR_ALLOCATION_DELAY = 1 days / ASSUMED_BLOCK_TIME;

    uint256 constant DEFAULT_OPERATOR_SHARES = 1e18;

    /// -----------------------------------------------------------------------
    /// Mocks
    /// -----------------------------------------------------------------------

    AllocationManager allocationManager;
    ERC20PresetFixedSupply tokenMock;
    StrategyBase strategyMock;

    /// -----------------------------------------------------------------------
    /// Defaults
    /// -----------------------------------------------------------------------

    OperatorSet defaultOperatorSet;
    IStrategy[] defaultStrategies;
    address defaultOperator = address(this);
    address defaultAVS = address(new MockAVSRegistrar());

    /// -----------------------------------------------------------------------
    /// Setup
    /// -----------------------------------------------------------------------

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();
        _initializeAllocationManager(address(this), pauserRegistry, 0);
        tokenMock = new ERC20PresetFixedSupply("Mock Token", "MOCK", type(uint256).max, address(this));
        strategyMock = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(new StrategyBase(IStrategyManager(address(strategyManagerMock)), pauserRegistry)),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, tokenMock)
                )
            )
        );

        defaultStrategies = strategyMock.toArray();
        defaultOperatorSet = OperatorSet(defaultAVS, 0);

        _createOperatorSet(defaultOperatorSet, defaultStrategies);
        _registerOperator(defaultOperator);
        _setAllocationDelay(defaultOperator, DEFAULT_OPERATOR_ALLOCATION_DELAY);
        _registerForOperatorSet(defaultOperator, defaultOperatorSet);
        _grantDelegatedStake(defaultOperator, defaultOperatorSet, DEFAULT_OPERATOR_SHARES);
    }

    /// -----------------------------------------------------------------------
    /// Internal Helpers
    /// -----------------------------------------------------------------------

    function _initializeAllocationManager(
        address _initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 _initialPausedStatus
    ) internal returns (AllocationManager) {
        return allocationManager = AllocationManager(
            address(
                new TransparentUpgradeableProxy(
                    address(
                        new AllocationManager(
                            IDelegationManager(address(delegationManagerMock)),
                            _pauserRegistry,
                            IPermissionController(address(permissionController)),
                            DEALLOCATION_DELAY,
                            ALLOCATION_CONFIGURATION_DELAY
                        )
                    ),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(AllocationManager.initialize.selector, _initialOwner, _initialPausedStatus)
                )
            )
        );
    }

    function _registerOperator(
        address operator
    ) internal {
        delegationManagerMock.setIsOperator(operator, true);
    }

    function _setAllocationDelay(address operator, uint32 delay) internal {
        cheats.prank(operator);
        allocationManager.setAllocationDelay(operator, delay);
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);
    }

    function _createOperatorSet(
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies
    ) internal returns (OperatorSet memory) {
        cheats.prank(operatorSet.avs);
        allocationManager.createOperatorSets(
            operatorSet.avs,
            CreateSetParams({operatorSetId: operatorSet.id, strategies: strategies}).toArray()
        );
        return operatorSet;
    }

    function _createOperatorSets(OperatorSet[] memory operatorSets, IStrategy[] memory strategies) internal {
        CreateSetParams[] memory createSetParams = new CreateSetParams[](operatorSets.length);

        for (uint256 i; i < operatorSets.length; ++i) {
            createSetParams[i] = CreateSetParams({operatorSetId: operatorSets[i].id, strategies: strategies});
        }

        cheats.prank(operatorSets[0].avs);
        allocationManager.createOperatorSets(operatorSets[0].avs, createSetParams);
    }

    function _registerForOperatorSet(address operator, OperatorSet memory operatorSet) internal {
        cheats.prank(operator);
        allocationManager.registerForOperatorSets(
            operator,
            RegisterParams({avs: operatorSet.avs, operatorSetIds: operatorSet.id.toArrayU32(), data: ""})
        );
    }

    function _grantDelegatedStake(address operator, OperatorSet memory operatorSet, uint256 stake) internal {
        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);
        delegationManagerMock.setOperatorsShares(operator, strategies, stake);
    }

    function _registerForOperatorSets(address operator, OperatorSet[] memory operatorSets) internal {
        cheats.startPrank(operator);
        for (uint256 i; i < operatorSets.length; ++i) {
            allocationManager.registerForOperatorSets(
                operator,
                RegisterParams({avs: operatorSets[i].avs, operatorSetIds: operatorSets[i].id.toArrayU32(), data: ""})
            );
        }
        cheats.stopPrank();
    }

    function _checkAllocationStorage(
        Allocation memory allocation,
        uint256 expectedCurrentMagnitude,
        int256 expectedPendingDiff,
        uint256 expectedEffectBlock
    ) internal view {
        console.log("Check Allocation Storage:".yellow());
        console.log("   currentMagnitude = %d", allocation.currentMagnitude);
        console.log("   pendingDiff = %d", allocation.pendingDiff);
        console.log("   effectBlock = %d", allocation.effectBlock);
        console.log("   currentBlock = %d", block.number);
        console.log("\n");

        assertApproxEqAbs(expectedCurrentMagnitude, allocation.currentMagnitude, 1, "currentMagnitude != expected");
        assertEq(expectedPendingDiff, allocation.pendingDiff, "pendingDiff != expected");
        assertEq(expectedEffectBlock, allocation.effectBlock, "effectBlock != expected");
    }

    function _checkAllocationStorage(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy strategy,
        uint256 expectedCurrentMagnitude,
        int256 expectedPendingDiff,
        uint256 expectedEffectBlock
    ) internal view {
        Allocation memory getAllocation = allocationManager.getAllocation(operator, operatorSet, strategy);
        Allocation memory getAllocations =
            allocationManager.getAllocations(operator.toArray(), operatorSet, strategy)[0];
        _checkAllocationStorage(getAllocation, expectedCurrentMagnitude, expectedPendingDiff, expectedEffectBlock);
        _checkAllocationStorage(getAllocations, expectedCurrentMagnitude, expectedPendingDiff, expectedEffectBlock);
    }

    function _checkSlashableStake(
        OperatorSet memory operatorSet,
        address operator,
        IStrategy[] memory strategies,
        uint256 expectedStake
    ) internal view {
        _checkSlashableStake(operatorSet, operator, strategies, expectedStake, block.number);
    }

    function _checkSlashableStake(
        OperatorSet memory operatorSet,
        address operator,
        IStrategy[] memory strategies,
        uint256 expectedStake,
        uint256 futureBlock
    ) internal view {
        uint256[] memory slashableStake = allocationManager.getMinimumSlashableStake({
            operatorSet: operatorSet,
            operators: operator.toArray(),
            strategies: strategies,
            futureBlock: uint32(futureBlock)
        })[0];

        for (uint256 i = 0; i < strategies.length; i++) {
            console.log(StdStyle.yellow("Check Slashable Stake:"));
            console.log("   slashableStake[%d] = %d", i, slashableStake[i]);
            console.log("\n");
            assertApproxEqAbs(slashableStake[i], expectedStake, 1, "slashableStake != expected");
        }
    }

    function _checkAllocationEvents(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy strategy,
        uint64 currentMagnitude,
        uint64 encumberedMagnitude,
        uint32 effectBlock
    ) internal {
        cheats.expectEmit(true, false, false, false, address(allocationManager));
        emit EncumberedMagnitudeUpdated(operator, strategy, encumberedMagnitude);
        cheats.expectEmit(true, false, false, false, address(allocationManager));
        emit AllocationUpdated(operator, operatorSet, strategy, currentMagnitude, effectBlock);
    }

    function _checkSlashEvents(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        uint256[] memory wadToSlash,
        string memory description
    ) internal {
        cheats.expectEmit(true, false, false, false, address(allocationManager));
        emit OperatorSlashed(operator, operatorSet, strategies, wadToSlash, description);
    }

    /// -----------------------------------------------------------------------
    /// Allocate/deallocate params
    /// -----------------------------------------------------------------------

    /// @dev Create allocate params, allocating `magnitude` to each strategy in the set
    function _newAllocateParams(
        OperatorSet memory operatorSet,
        uint64 magnitude
    ) internal view returns (AllocateParams[] memory) {
        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);
        uint64[] memory newMagnitudes = new uint64[](strategies.length);

        for (uint256 i; i < strategies.length; ++i) {
            newMagnitudes[i] = magnitude;
        }

        return
            AllocateParams({operatorSet: operatorSet, strategies: strategies, newMagnitudes: newMagnitudes}).toArray();
    }

    /// @dev Create allocate params for multiple operator sets
    function _newAllocateParams(
        OperatorSet[] memory operatorSets,
        uint64 magnitude
    ) internal view returns (AllocateParams[] memory) {
        AllocateParams[] memory allocateParams = new AllocateParams[](operatorSets.length);

        for (uint256 i; i < operatorSets.length; ++i) {
            allocateParams[i] = _newAllocateParams(operatorSets[i], magnitude)[0];
        }

        return allocateParams;
    }

    /// @dev Create random allocation params to the default operator set and strategy
    function _randAllocateParams_DefaultOpSet() internal returns (AllocateParams[] memory) {
        return _randAllocateParams_SingleMockStrategy(defaultOperatorSet.toArray());
    }

    /// @dev Create allocate params for random magnitudes to the same default strategy across multiple operator sets
    function _randAllocateParams_SingleMockStrategy(
        OperatorSet[] memory operatorSets
    ) internal returns (AllocateParams[] memory) {
        // Give each set a minimum of 1 magnitude
        uint64[] memory magnitudes = new uint64[](operatorSets.length);
        uint64 usedMagnitude;
        for (uint8 i = 0; i < magnitudes.length; ++i) {
            magnitudes[i] = 1;
            usedMagnitude++;
        }

        // Distribute remaining magnitude
        uint64 maxMagnitude = WAD;
        for (uint8 i = 0; i < magnitudes.length; ++i) {
            uint64 remainingMagnitude = maxMagnitude - usedMagnitude;
            if (remainingMagnitude > 0) {
                magnitudes[i] += uint64(random().Uint256(0, remainingMagnitude));
                usedMagnitude += magnitudes[i] - 1;
            }
        }

        AllocateParams[] memory params = new AllocateParams[](magnitudes.length);
        for (uint256 i; i < params.length; ++i) {
            params[i] = AllocateParams({
                operatorSet: operatorSets[i],
                strategies: defaultStrategies,
                newMagnitudes: magnitudes[i].toArrayU64()
            });
        }

        return params;
    }

    /// @dev Create allocate params for random magnitudes to the same default strategy across multiple operator sets
    /// NOTE: this variant allocates ALL magnitude (1 WAD)
    function _randAllocateParams_SingleMockStrategy_AllocAll(
        OperatorSet[] memory operatorSets
    ) internal returns (AllocateParams[] memory) {
        // Give each set a minimum of 1 magnitude
        uint64[] memory magnitudes = new uint64[](operatorSets.length);
        uint64 usedMagnitude;
        for (uint8 i = 0; i < magnitudes.length; ++i) {
            magnitudes[i] = 1;
            usedMagnitude++;
        }

        // Distribute remaining magnitude
        uint64 maxMagnitude = WAD;
        for (uint8 i = 0; i < magnitudes.length; ++i) {
            uint64 remainingMagnitude = maxMagnitude - usedMagnitude;
            if (remainingMagnitude > 0) {
                magnitudes[i] += uint64(random().Uint64(0, remainingMagnitude));
                usedMagnitude += magnitudes[i] - 1;
            }
        }

        // If there's any left, dump it on a random set
        uint64 magnitudeLeft = maxMagnitude - usedMagnitude;
        if (magnitudeLeft > 0) {
            uint256 randIdx = random().Uint256(0, magnitudes.length - 1);
            magnitudes[randIdx] += magnitudeLeft;
            usedMagnitude += magnitudeLeft;
        }

        AllocateParams[] memory params = new AllocateParams[](magnitudes.length);
        for (uint256 i; i < params.length; ++i) {
            params[i] = AllocateParams({
                operatorSet: operatorSets[i],
                strategies: defaultStrategies,
                newMagnitudes: magnitudes[i].toArrayU64()
            });
        }

        return params;
    }

    /// @dev Create allocate/deallocate params to the same default strategy across multiple sets
    function _randAllocAndDeallocParams_SingleMockStrategy(
        OperatorSet[] memory operatorSets
    ) internal returns (AllocateParams[] memory, AllocateParams[] memory) {
        AllocateParams[] memory allocateParams = _randAllocateParams_SingleMockStrategy(operatorSets);
        AllocateParams[] memory deallocateParams = new AllocateParams[](allocateParams.length);

        // Generate a random deallocation for each operator set
        for (uint256 i; i < deallocateParams.length; ++i) {
            deallocateParams[i] = AllocateParams({
                operatorSet: allocateParams[i].operatorSet,
                strategies: allocateParams[i].strategies,
                newMagnitudes: uint64(random().Uint256({min: 0, max: allocateParams[i].newMagnitudes[0] - 1})).toArrayU64()
            });
        }

        return (allocateParams, deallocateParams);
    }

    /// -----------------------------------------------------------------------
    /// Utils
    /// -----------------------------------------------------------------------

    function _maxNumToClear() internal pure returns (uint16[] memory) {
        uint16[] memory numToClear = new uint16[](1);
        numToClear[0] = type(uint16).max;
        return numToClear;
    }
}

contract AllocationManagerUnitTests_Initialization_Setters is AllocationManagerUnitTests {
    /// -----------------------------------------------------------------------
    /// initialize()
    /// -----------------------------------------------------------------------

    /// @dev Asserts the following:
    /// 1. The fn can only be called once, during deployment.
    /// 2. The fn initializes the contract state correctly (owner, pauserRegistry, and initialPausedStatus).
    function testFuzz_Initialize(
        Randomness r
    ) public rand(r) {
        // Generate random values for the expected initial state of the contract.
        address expectedInitialOwner = r.Address();
        IPauserRegistry expectedPauserRegistry = IPauserRegistry(r.Address());

        // Deploy the contract with the expected initial state.
        uint256 initialPausedStatus = r.Uint256();
        AllocationManager alm =
            _initializeAllocationManager(expectedInitialOwner, expectedPauserRegistry, initialPausedStatus);

        // Assert that the contract can only be initialized once.
        vm.expectRevert("Initializable: contract is already initialized");
        alm.initialize(expectedInitialOwner, initialPausedStatus);

        // Assert immutable state
        assertEq(address(alm.delegation()), address(delegationManagerMock));
        assertEq(alm.DEALLOCATION_DELAY(), DEALLOCATION_DELAY);
        assertEq(alm.ALLOCATION_CONFIGURATION_DELAY(), ALLOCATION_CONFIGURATION_DELAY);

        // Assert initialiation state
        assertEq(alm.owner(), expectedInitialOwner);
        assertEq(alm.paused(), initialPausedStatus);
    }
}

contract AllocationManagerUnitTests_SlashOperator is AllocationManagerUnitTests {
    using SingleItemArrayLib for *;
    using SlashingLib for *;

    /// -----------------------------------------------------------------------
    /// slashOperator()
    /// -----------------------------------------------------------------------

    function _randSlashingParams(address operator, uint32 operatorSetId) internal returns (SlashingParams memory) {
        return SlashingParams({
            operator: operator,
            operatorSetId: operatorSetId,
            wadToSlash: random().Uint256(1, WAD),
            description: "test"
        });
    }

    function test_revert_paused() public {
        allocationManager.pause(2 ** PAUSED_OPERATOR_SLASHING);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        allocationManager.slashOperator(defaultAVS, _randSlashingParams(defaultOperator, 0));
    }

    function test_revert_slashZero() public {
        SlashingParams memory slashingParams = _randSlashingParams(defaultOperator, 0);
        slashingParams.wadToSlash = 0;

        cheats.prank(defaultAVS);
        cheats.expectRevert(InvalidWadToSlash.selector);
        allocationManager.slashOperator(defaultAVS, slashingParams);
    }

    function test_revert_slashGreaterThanWAD() public {
        SlashingParams memory slashingParams = _randSlashingParams(defaultOperator, 0);
        slashingParams.wadToSlash = WAD + 1;

        cheats.prank(defaultAVS);
        cheats.expectRevert(InvalidWadToSlash.selector);
        allocationManager.slashOperator(defaultAVS, slashingParams);
    }

    function test_revert_NotMemberOfSet() public {
        cheats.prank(defaultAVS);
        cheats.expectRevert(NotMemberOfSet.selector);
        allocationManager.slashOperator(defaultAVS, _randSlashingParams(random().Address(), 0));
    }

    function test_revert_operatorAllocated_notActive() public {
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, WAD);

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        cheats.prank(defaultAVS);
        allocationManager.slashOperator(
            defaultAVS, 
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: allocateParams[0].operatorSet.id,
                wadToSlash: WAD,
                description: "test"
            })
        );

        uint256 effectBlock = block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY;

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: 0,
            expectedPendingDiff: int64(allocateParams[0].newMagnitudes[0]),
            expectedEffectBlock: effectBlock
        });

        cheats.roll(effectBlock);

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: allocateParams[0].newMagnitudes[0],
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });
    }

    /**
     * Allocates all magnitude to for a single strategy to an operatorSet. Slashes 25%
     * Asserts that:
     * 1. Events are emitted
     * 2. Encumbered mag is updated
     * 3. Max mag is updated
     * 4. Calculations for `getAllocatableMagnitude` and `getAllocation` are correct
     */
    function test_slashPostAllocation() public {
        // Generate allocation for this operator set, we allocate max
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, WAD);

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        _checkSlashEvents(defaultOperator, defaultOperatorSet, defaultStrategies, uint256(25e16).toArrayU256(), "test");

        // Slash operator for 25%
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(
            defaultAVS,
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: defaultOperatorSet.id,
                wadToSlash: 25e16,
                description: "test"
            })
        );

        // Check storage
        assertEq(
            75e16,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            75e16, allocationManager.getMaxMagnitudes(defaultOperator, defaultStrategies)[0], "maxMagnitude not updated"
        );
        assertEq(
            0,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            "allocatableMagnitude shoudl be 0"
        );
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: 75e16,
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedStake: DEFAULT_OPERATOR_SHARES.mulWad(75e16)
        });
    }

    /// @notice Same test as above, but fuzzes the allocation
    function testFuzz_slashPostAllocation(
        Randomness r
    ) public rand(r) {
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();

        // Allocate magnitude and roll forward to completable block
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        SlashingParams memory slashingParams = _randSlashingParams(defaultOperator, defaultOperatorSet.id);

        uint64 allocatedMagnitude = allocateParams[0].newMagnitudes[0];
        uint64 expectedSlashedMagnitude =
            uint64(SlashingLib.mulWadRoundUp(allocatedMagnitude, slashingParams.wadToSlash));
        uint64 expectedEncumberedMagnitude = allocatedMagnitude - expectedSlashedMagnitude;
        uint64 maxMagnitudeAfterSlash = WAD - expectedSlashedMagnitude;
        uint256 slashedStake = DEFAULT_OPERATOR_SHARES.mulWad(expectedSlashedMagnitude);
        uint256 newSlashableMagnitude = uint256(expectedEncumberedMagnitude).divWad(maxMagnitudeAfterSlash);

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategies: defaultStrategies,
            wadToSlash: uint256(expectedSlashedMagnitude).toArrayU256(),
            description: "test"
        });

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Check storage
        assertEq(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            maxMagnitudeAfterSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, defaultStrategies)[0],
            "maxMagnitude not updated"
        );
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: expectedEncumberedMagnitude,
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedStake: (DEFAULT_OPERATOR_SHARES - slashedStake).mulWad(newSlashableMagnitude)
        });
    }

    /**
     * Allocates half of magnitude for a single strategy to an operatorSet. Then allocates again. Slashes 50%
     * Asserts that:
     * 1. Events are emitted
     * 2. Encumbered mag is updated
     * 3. Max mag is updated
     * 4. Calculations for `getAllocatableMagnitude` and `getAllocation` are correct
     * 5. The second magnitude allocation is not slashed from
     */
    function testFuzz_slash_oneCompletedAlloc_onePendingAlloc(
        Randomness r
    ) public rand(r) {
        uint64 wadToSlash = r.Uint64(0.01 ether, WAD);

        // Generate allocation for `strategyMock`, we allocate half
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Check slashable stake after the first allocation
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedStake: DEFAULT_OPERATOR_SHARES.mulWad(5e17)
        });

        // Allocate the other half
        AllocateParams[] memory allocateParams2 = _newAllocateParams(defaultOperatorSet, WAD);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams2);
        uint32 secondAllocEffectBlock = uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Check slashable stake hasn't changed after the second allocation
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedStake: DEFAULT_OPERATOR_SHARES.mulWad(5e17)
        });

        // Check minimum slashable stake would not change even after the second allocation becomes effective
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedStake: DEFAULT_OPERATOR_SHARES.mulWad(5e17),
            futureBlock: secondAllocEffectBlock
        });

        // Slash operator for 50%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: defaultOperatorSet.id,
            wadToSlash: wadToSlash,
            description: "test"
        });

        uint64 totalAllocated = 0.5 ether;
        uint64 expectedEncumberedMagnitude = (WAD - uint64(uint256(totalAllocated) * uint256(wadToSlash) / WAD));
        uint64 magnitudeAfterSlash = totalAllocated - uint64(uint256(totalAllocated) * uint256(wadToSlash) / WAD);
        uint64 maxMagnitudeAfterSlash = expectedEncumberedMagnitude;

        uint64 expectedSlashedMagnitude = uint64(totalAllocated.mulWadRoundUp(slashingParams.wadToSlash));
        uint256 newSlashableMagnitude = uint256(magnitudeAfterSlash).divWad(maxMagnitudeAfterSlash);
        uint256 slashedStake = DEFAULT_OPERATOR_SHARES.mulWad(expectedSlashedMagnitude);
        uint256 newTotalStake = DEFAULT_OPERATOR_SHARES - slashedStake;

        // // STACK TOO DEEP
        // // _checkSlashEvents({
        // //     operator: defaultOperator,
        // //     operatorSet: defaultOperatorSet,
        // //     strategies: defaultStrategies,
        // //     wadToSlash: uint256(wadToSlash).toArrayU256(),
        // //     description: "test"
        // // });

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Check storage
        assertApproxEqAbs(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            1,
            "encumberedMagnitude not updated"
        );
        assertApproxEqAbs(
            maxMagnitudeAfterSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, defaultStrategies)[0],
            1,
            "maxMagnitude not updated"
        );

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: magnitudeAfterSlash,
            expectedPendingDiff: 5e17,
            expectedEffectBlock: secondAllocEffectBlock
        });

        // Slashable stake should include first allocation and slashed magnitude
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedStake: newTotalStake.mulWad(newSlashableMagnitude)
        });

        cheats.roll(secondAllocEffectBlock);

        assertEq(
            0,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            "allocatableMagnitude should be 0"
        );

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: magnitudeAfterSlash + 0.5 ether,
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });

        newSlashableMagnitude = allocateParams2[0].newMagnitudes[0];
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedStake: newTotalStake.mulWad(newSlashableMagnitude)
        });
    }

    /**
     * Allocates 100% magnitude for a single strategy to an operatorSet.
     * First slashes 99% from the operatorSet, slashes 99.99% a second time, and on the third slash, slashes
     * 99.9999999999999% which should get rounded up to 100% or WAD wadSlashed leaving the operator with no magnitude
     * in the operatorSet, 0 encumbered magnitude, and 0 max magnitude.
     *
     * Asserts that:
     * 1. Events are emitted
     * 2. Encumbered mag is updated
     * 3. Max mag is updated
     * 4. Calculations for `getAllocatableMagnitude` and `getAllocation` are correct
     * 5. Slashed amounts are rounded up to ensure magnitude is always slashed
     */
    function test_repeatUntilFullSlash() public {
        // Generate allocation for `strategyMock`, we allocate 100% to opSet 0
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, WAD);

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Check slashable amount after initial allocation
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedStake: DEFAULT_OPERATOR_SHARES
        });

        // 1. Slash operator for 99% in opSet 0 bringing their magnitude to 1e16
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: defaultOperatorSet.id,
            wadToSlash: 99e16,
            description: "test"
        });

        uint64 expectedEncumberedMagnitude = 1e16; // After slashing 99%, only 1% expected encumberedMagnitude
        uint64 magnitudeAfterSlash = 1e16;
        uint64 maxMagnitudeAfterSlash = 1e16; // 1e15 is maxMagnitude

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategies: defaultStrategies,
            wadToSlash: uint256(99e16).toArrayU256(),
            description: "test"
        });

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Check storage
        assertEq(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            maxMagnitudeAfterSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, defaultStrategies)[0],
            "maxMagnitude not updated"
        );
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, defaultOperatorSet, strategyMock);
        assertEq(magnitudeAfterSlash, allocation.currentMagnitude, "currentMagnitude not updated");

        // Check slashable amount after first slash
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedStake: DEFAULT_OPERATOR_SHARES.mulWad(1e16)
        });

        // 2. Slash operator again for 99.99% in opSet 0 bringing their magnitude to 1e14
        slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: defaultOperatorSet.id,
            wadToSlash: 9999e14,
            description: "test"
        });
        expectedEncumberedMagnitude = 1e12; // After slashing 99.99%, only 0.01% expected encumberedMagnitude
        magnitudeAfterSlash = 1e12;
        maxMagnitudeAfterSlash = 1e12;

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategies: defaultStrategies,
            wadToSlash: uint256(9999e14).toArrayU256(),
            description: "test"
        });

        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Check storage
        assertEq(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            maxMagnitudeAfterSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, defaultStrategies)[0],
            "maxMagnitude not updated"
        );
        allocation = allocationManager.getAllocation(defaultOperator, defaultOperatorSet, strategyMock);
        assertEq(magnitudeAfterSlash, allocation.currentMagnitude, "currentMagnitude not updated");

        // Check slashable amount after second slash
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedStake: DEFAULT_OPERATOR_SHARES.mulWad(1e12)
        });

        // 3. Slash operator again for 99.9999999999999% in opSet 0
        slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: defaultOperatorSet.id,
            wadToSlash: WAD - 1e3,
            description: "test"
        });
        // Should technically be 1e3 remaining but with rounding error and rounding up slashed amounts
        // the remaining magnitude is 0
        expectedEncumberedMagnitude = 0; // Should technically be 1e3 remaining but with rounding error and rounding up slashed amounts.
        magnitudeAfterSlash = 0;
        maxMagnitudeAfterSlash = 0;

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategies: defaultStrategies,
            wadToSlash: uint256(WAD - 1e3).toArrayU256(),
            description: "test"
        });

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Check storage
        assertEq(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            maxMagnitudeAfterSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, defaultStrategies)[0],
            "maxMagnitude not updated"
        );
        allocation = allocationManager.getAllocation(defaultOperator, defaultOperatorSet, strategyMock);
        assertEq(magnitudeAfterSlash, allocation.currentMagnitude, "currentMagnitude not updated");

        // Check slashable amount after final slash
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedStake: 0
        });
    }

    /**
     * Allocates all of magnitude to a single strategy to an operatorSet. Deallocate half. Finally, slash while deallocation is pending
     * Asserts that:
     * 1. Events are emitted, including for deallocation
     * 2. Encumbered mag is updated
     * 3. Max mag is updated
     * 4. Calculations for `getAllocatableMagnitude` and `getAllocation` are correct
     * 5. The deallocation is slashed from
     * 6. Pending magnitude updates post deallocation are valid
     */
    function testFuzz_SlashWhileDeallocationPending(
        Randomness r
    ) public rand(r) {
        AllocateParams[] memory allocateParams = r.AllocateParams(defaultAVS, 1, 1);
        AllocateParams[] memory deallocateParams = r.DeallocateParams(allocateParams);
        CreateSetParams[] memory createSetParams = r.CreateSetParams(allocateParams);
        RegisterParams memory registerParams = r.RegisterParams(allocateParams);
        SlashingParams memory slashingParams = r.SlashingParams(defaultOperator, allocateParams[0]);

        console.log("wadToSlash: %d", slashingParams.wadToSlash);

        delegationManagerMock.setOperatorShares(
            defaultOperator, allocateParams[0].strategies[0], DEFAULT_OPERATOR_SHARES
        );

        cheats.prank(defaultAVS);
        allocationManager.createOperatorSets(defaultAVS, createSetParams);
        cheats.startPrank(defaultOperator);
        allocationManager.registerForOperatorSets(defaultOperator, registerParams);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY);
        cheats.stopPrank();

        // Check slashable stake after deallocation (still pending; no change)
        _checkSlashableStake({
            operatorSet: allocateParams[0].operatorSet,
            operator: defaultOperator,
            strategies: allocateParams[0].strategies,
            expectedStake: allocateParams[0].newMagnitudes[0]
        });

        // Check slashable stake after deallocation takes effect, before slashing
        _checkSlashableStake({
            operatorSet: allocateParams[0].operatorSet,
            operator: defaultOperator,
            strategies: allocateParams[0].strategies,
            expectedStake: deallocateParams[0].newMagnitudes[0],
            futureBlock: deallocationEffectBlock
        });

        uint256 magnitudeAllocated = allocateParams[0].newMagnitudes[0];
        uint256 magnitudeDeallocated = magnitudeAllocated - deallocateParams[0].newMagnitudes[0];
        uint256 magnitudeSlashed = magnitudeAllocated.mulWad(slashingParams.wadToSlash);
        uint256 expectedCurrentMagnitude = magnitudeAllocated - magnitudeSlashed;
        int128 expectedPendingDiff =
            -int128(uint128(magnitudeDeallocated - magnitudeDeallocated.mulWadRoundUp(slashingParams.wadToSlash)));

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: allocateParams[0].operatorSet,
            strategies: allocateParams[0].strategies,
            wadToSlash: slashingParams.wadToSlash.toArrayU256(),
            description: "test"
        });

        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: allocateParams[0].operatorSet,
            strategy: allocateParams[0].strategies[0],
            expectedCurrentMagnitude: expectedCurrentMagnitude,
            expectedPendingDiff: expectedPendingDiff,
            expectedEffectBlock: deallocationEffectBlock
        });

        // Check slashable stake after slash
        _checkSlashableStake({
            operatorSet: allocateParams[0].operatorSet,
            operator: defaultOperator,
            strategies: allocateParams[0].strategies,
            expectedStake: expectedCurrentMagnitude
        });

        // Check slashable stake after deallocation takes effect
        _checkSlashableStake({
            operatorSet: allocateParams[0].operatorSet,
            operator: defaultOperator,
            strategies: allocateParams[0].strategies,
            expectedStake: expectedCurrentMagnitude - uint128(-expectedPendingDiff) - 1,
            futureBlock: deallocationEffectBlock
        });

        assertEq(
            expectedCurrentMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, allocateParams[0].strategies[0]),
            "encumberedMagnitude not updated"
        );
        assertEq(
            WAD - slashingParams.wadToSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, allocateParams[0].strategies)[0],
            "maxMagnitude not updated"
        );

        cheats.roll(deallocationEffectBlock);
        allocationManager.clearDeallocationQueue(defaultOperator, allocateParams[0].strategies, _maxNumToClear());

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: allocateParams[0].operatorSet,
            strategy: allocateParams[0].strategies[0],
            expectedCurrentMagnitude: deallocateParams[0].newMagnitudes[0]
                - deallocateParams[0].newMagnitudes[0] * slashingParams.wadToSlash / WAD,
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });

        // Check slashable stake after slash and deallocation
        _checkSlashableStake({
            operatorSet: allocateParams[0].operatorSet,
            operator: defaultOperator,
            strategies: allocateParams[0].strategies,
            expectedStake: expectedCurrentMagnitude - uint128(-expectedPendingDiff) - 1
        });
    }

    /**
     * Allocates all magnitude to a single opSet. Then slashes the entire magnitude
     * Asserts that:
     * 1. The operator cannot allocate again
     */
    function testRevert_allocateAfterSlashedEntirely() public {
        // Allocate all magnitude
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, WAD);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategies: defaultStrategies,
            wadToSlash: WAD.toArrayU256(),
            description: "test"
        });

        // Slash operator for 100%
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(
            defaultAVS,
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: allocateParams[0].operatorSet.id,
                wadToSlash: WAD,
                description: "test"
            })
        );

        OperatorSet memory operatorSet =
            _createOperatorSet(OperatorSet(defaultAVS, random().Uint32()), defaultStrategies);
        AllocateParams[] memory allocateParams2 = _newAllocateParams(operatorSet, 1);

        // Attempt to allocate
        cheats.prank(defaultOperator);
        cheats.expectRevert(InsufficientMagnitude.selector);
        allocationManager.modifyAllocations(defaultOperator, allocateParams2);
    }

    /**
     * Allocates all magnitude to a single opSet. Deallocateas magnitude. Slashes al
     * Asserts that:
     * 1. The Allocation is 0 after slash
     * 2. Them sotrage post slash for encumbered and maxMags ais zero
     */
    function test_allocateAll_deallocateAll() public {
        // Allocate all magnitude
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(defaultOperatorSet, WAD));
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate all
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(defaultOperatorSet, 0));

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategies: defaultStrategies,
            wadToSlash: WAD.toArrayU256(),
            description: "test"
        });

        // Slash operator for 100%
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(
            defaultAVS,
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: defaultOperatorSet.id,
                wadToSlash: WAD,
                description: "test"
            })
        );

        assertEq(
            0, allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude not updated"
        );
        assertEq(
            0, allocationManager.getMaxMagnitudes(defaultOperator, defaultStrategies)[0], "maxMagnitude not updated"
        );
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: 0,
            expectedPendingDiff: 0,
            expectedEffectBlock: block.number + DEALLOCATION_DELAY
        });
    }

    /**
     * Slashes the operator after deallocation, even if the deallocation has not been cleared. Validates that:
     * 1. Even if we do not clear deallocation queue, the deallocation is NOT slashed from since we're passed the deallocationEffectBlock
     * 2. Validates storage post slash & post clearing deallocation queue
     * 3. Max magnitude only decreased proportionally by the magnitude set after deallocation
     */
    function test_allocate_deallocate_slashAfterDeallocation() public {
        // Allocate all magnitude
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, WAD);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate half
        AllocateParams[] memory deallocateParams = _newAllocateParams(defaultOperatorSet, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY);

        // Check storage post deallocation
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: WAD,
            expectedPendingDiff: -5e17,
            expectedEffectBlock: deallocationEffectBlock
        });

        // Warp to deallocation effect block
        cheats.roll(deallocationEffectBlock);

        // Slash operator for 25%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: defaultOperatorSet.id,
            wadToSlash: 25e16,
            description: "test"
        });

        uint64 expectedEncumberedMagnitude = 375e15; // 25e16 is slashed. 5e17 was previously
        uint64 magnitudeAfterSlash = 375e15;
        uint64 maxMagnitudeAfterSlash = 875e15; // Operator can only allocate up to 75e16 magnitude since 25% is slashed

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategies: defaultStrategies,
            wadToSlash: uint256(25e16).toArrayU256(),
            description: "test"
        });

        // Slash Operator, only emit events assuming that there is no deallocation
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Check storage post slash
        assertEq(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            maxMagnitudeAfterSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, defaultStrategies)[0],
            "maxMagnitude not updated"
        );
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: magnitudeAfterSlash,
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });

        uint64 allocatableMagnitudeAfterSlash = allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock);

        // Check storage after complete modification.
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());
        assertEq(
            allocatableMagnitudeAfterSlash,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            "allocatable mag after slash shoudl be equal to allocatable mag after clearing queue"
        );
    }

    /**
     * Allocates to multiple operatorSets for a strategy. Only slashes from one operatorSet. Validates
     * 1. The slashable shares of each operatorSet after magnitude allocation
     * 2. The first operatorSet has less slashable shares post slash
     * 3. The second operatorSet has the same number slashable shares post slash
     * 4. The PROPORTION that is slashable for opSet 2 has increased
     * 5. Encumbered magnitude, total allocatable magnitude
     */
    function test_allocateMultipleOpsets_slashSingleOpset() public {
        // Set 100e18 shares for operator in DM
        uint256 operatorShares = 100e18;
        delegationManagerMock.setOperatorShares(defaultOperator, strategyMock, operatorShares);
        uint64 magnitudeToAllocate = 4e17;

        OperatorSet memory operatorSet = OperatorSet(defaultAVS, 1);
        OperatorSet memory operatorSet2 = OperatorSet(defaultAVS, 2);

        // Allocate 40% to firstOperatorSet, 40% to secondOperatorSet
        AllocateParams[] memory allocateParams = new AllocateParams[](2);
        allocateParams[0] = _newAllocateParams(
            _createOperatorSet(OperatorSet(defaultAVS, 1), defaultStrategies), magnitudeToAllocate
        )[0];
        allocateParams[1] = _newAllocateParams(
            _createOperatorSet(OperatorSet(defaultAVS, 2), defaultStrategies), magnitudeToAllocate
        )[0];

        _registerForOperatorSet(defaultOperator, operatorSet);
        _registerForOperatorSet(defaultOperator, operatorSet2);

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Get slashable shares for each operatorSet
        address[] memory operatorArray = new address[](1);
        operatorArray[0] = defaultOperator;

        uint256 maxMagnitude = allocationManager.getMaxMagnitudes(defaultOperator, defaultStrategies)[0];
        uint256 opSet2PortionOfMaxMagnitude = uint256(magnitudeToAllocate) * WAD / maxMagnitude;

        // Slash operator on operatorSet1 for 50%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocateParams[0].operatorSet.id,
            wadToSlash: 5e17,
            description: "test"
        });

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: operatorSet,
            strategies: defaultStrategies,
            wadToSlash: uint256(5e17).toArrayU256(),
            description: "test"
        });

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Operator should now have 80e18 shares, since half of 40e18 was slashed
        delegationManagerMock.setOperatorShares(defaultOperator, strategyMock, 80e18);

        // Validate encumbered and total allocatable magnitude
        uint256 maxMagnitudeAfterSlash = allocationManager.getMaxMagnitudes(defaultOperator, defaultStrategies)[0];
        uint256 expectedEncumberedMagnitude = 6e17; // 4e17 from opSet2, 2e17 from opSet1
        assertEq(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            maxMagnitudeAfterSlash - expectedEncumberedMagnitude,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            "allocatableMagnitude should be diff of maxMagnitude and encumberedMagnitude"
        );

        // Check proportion after slash
        uint256 opSet2PortionOfMaxMagnitudeAfterSlash = uint256(magnitudeToAllocate) * WAD / maxMagnitudeAfterSlash;
        assertGt(
            opSet2PortionOfMaxMagnitudeAfterSlash,
            opSet2PortionOfMaxMagnitude,
            "opSet2 should have a greater proportion to slash from previous"
        );
    }

    /**
     * Allocates to multiple strategies for the given operatorSetKey. Slashes from both strategies Validates a slash propogates to both strategies.
     * Validates that
     * 1. Proper events are emitted for each strategy slashed
     * 2. Each strategy is slashed proportional to its allocation
     * 3. Storage is updated for each strategy, opSet
     */
    function test_allocateMultipleStrategies_slashMultiple() public {
        // Allocate to each strategy
        uint64 strategy1Magnitude = 5e17;
        uint64 strategy2Magnitude = WAD;

        OperatorSet memory operatorSet = OperatorSet(defaultAVS, random().Uint32());
        _createOperatorSet(operatorSet, random().StrategyArray(2));
        _registerForOperatorSet(defaultOperator, operatorSet);

        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);
        AllocateParams memory allocateParams =
            AllocateParams({operatorSet: operatorSet, strategies: strategies, newMagnitudes: new uint64[](2)});
        allocateParams.newMagnitudes[0] = strategy1Magnitude;
        allocateParams.newMagnitudes[1] = strategy2Magnitude;

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams.toArray());
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Slash operator on both strategies for 60%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: operatorSet.id,
            wadToSlash: 6e17,
            description: "test"
        });

        uint64[] memory expectedEncumberedMags = new uint64[](2);
        expectedEncumberedMags[0] = 2e17; // 60% of 5e17
        expectedEncumberedMags[1] = 4e17; // 60% of WAD

        uint64[] memory expectedMagnitudeAfterSlash = new uint64[](2);
        expectedMagnitudeAfterSlash[0] = 2e17;
        expectedMagnitudeAfterSlash[1] = 4e17;

        uint64[] memory expectedMaxMagnitudeAfterSlash = new uint64[](2);
        expectedMaxMagnitudeAfterSlash[0] = 7e17;
        expectedMaxMagnitudeAfterSlash[1] = 4e17;

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: operatorSet,
            strategies: strategies,
            wadToSlash: uint256(6e17).toArrayU256(),
            description: "test"
        });

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Check storage
        for (uint256 i; i < strategies.length; ++i) {
            assertEq(
                expectedEncumberedMags[i],
                allocationManager.encumberedMagnitude(defaultOperator, strategies[i]),
                "encumberedMagnitude not updated"
            );
            assertEq(
                expectedMaxMagnitudeAfterSlash[i] - expectedMagnitudeAfterSlash[i],
                allocationManager.getAllocatableMagnitude(defaultOperator, strategies[i]),
                "allocatableMagnitude not updated"
            );
            _checkAllocationStorage({
                operator: defaultOperator,
                operatorSet: operatorSet,
                strategy: strategies[i],
                expectedCurrentMagnitude: expectedMagnitudeAfterSlash[i],
                expectedPendingDiff: 0,
                expectedEffectBlock: 0
            });
        }
    }

    /// @dev Allocates magnitude. Deallocates some. Slashes a portion, and then allocates up to the max available magnitude
    function testFuzz_allocate_deallocate_slashWhilePending_allocateMax(
        Randomness r
    ) public rand(r) {
        AllocateParams[] memory allocateParams = r.AllocateParams({avs: defaultAVS, numAllocations: 1, numStrats: 1});
        AllocateParams[] memory deallocateParams = r.DeallocateParams(allocateParams);
        CreateSetParams[] memory createSetParams = r.CreateSetParams(allocateParams);
        OperatorSet memory operatorSet = allocateParams[0].operatorSet;
        IStrategy strategy = allocateParams[0].strategies[0];

        cheats.prank(defaultAVS);
        allocationManager.createOperatorSets(defaultAVS, createSetParams);

        _registerForOperatorSet(defaultOperator, operatorSet);

        // Allocate some magnitude, then deallocate some.
        cheats.startPrank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);
        cheats.roll(block.number + DEALLOCATION_DELAY);
        cheats.stopPrank();

        // Slash operator for some random amount (1% -> 99%).
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: operatorSet.id,
            wadToSlash: r.Uint64(0.01 ether, 0.99 ether),
            description: "test"
        });

        uint256 magnitudeBeforeSlash = deallocateParams[0].newMagnitudes[0];
        uint256 slashedMagnitude = magnitudeBeforeSlash * slashingParams.wadToSlash / WAD;
        uint256 currentMagnitude = magnitudeBeforeSlash - slashedMagnitude - 1;
        uint256 maxMagnitude = WAD - slashedMagnitude - 1;

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: operatorSet,
            strategies: strategy.toArray(),
            wadToSlash: uint256(slashingParams.wadToSlash).toArrayU256(),
            description: "test"
        });

        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        assertEq(
            currentMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategy),
            "encumberedMagnitude should be half of firstMod"
        );

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: operatorSet,
            strategy: strategy,
            expectedCurrentMagnitude: uint64(currentMagnitude),
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });

        // Clear deallocation queue.
        allocationManager.clearDeallocationQueue(defaultOperator, strategy.toArray(), _maxNumToClear());

        assertEq(
            maxMagnitude,
            allocationManager.getMaxMagnitudes(defaultOperator, strategy.toArray())[0],
            "maxMagnitude should be expectedMaxMagnitude"
        );

        assertEq(
            maxMagnitude - currentMagnitude,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategy),
            "allocatableMagnitude should be expectedAllocatable"
        );

        // Allocate up to max magnitude
        AllocateParams[] memory allocateParams2 = _newAllocateParams(operatorSet, uint64(maxMagnitude));
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams2);

        // Assert that encumbered is expectedMaxMagnitude
        assertEq(
            0, allocationManager.getAllocatableMagnitude(defaultOperator, strategy), "allocatableMagnitude should be 0"
        );
    }
}

contract AllocationManagerUnitTests_ModifyAllocations is AllocationManagerUnitTests {
    using SingleItemArrayLib for *;
    using OperatorSetLib for *;

    function test_revert_paused() public {
        allocationManager.pause(2 ** PAUSED_MODIFY_ALLOCATIONS);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        allocationManager.modifyAllocations(address(this), new AllocateParams[](0));
    }

    function test_revert_invalidCaller() public {
        address invalidOperator = address(0x2);
        cheats.expectRevert(InvalidCaller.selector);
        allocationManager.modifyAllocations(invalidOperator, new AllocateParams[](0));
    }

    function test_revert_allocationDelayNotSet() public {
        address invalidOperator = address(0x2);
        cheats.prank(invalidOperator);
        cheats.expectRevert(UninitializedAllocationDelay.selector);
        allocationManager.modifyAllocations(invalidOperator, new AllocateParams[](0));
    }

    function test_revert_allocationDelayNotInEffect() public {
        address operator = address(0x2);
        _registerOperator(operator);

        cheats.startPrank(operator);
        allocationManager.setAllocationDelay(operator, 5);
        // even though the operator has an allocation delay set, it is not in effect
        // and modifyAllocations should still be blocked
        cheats.expectRevert(UninitializedAllocationDelay.selector);
        allocationManager.modifyAllocations(operator, new AllocateParams[](0));
    }

    function test_revert_lengthMismatch() public {
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();
        allocateParams[0].newMagnitudes = new uint64[](0);

        cheats.expectRevert(InputArrayLengthMismatch.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
    }

    function test_revert_invalidOperatorSet() public {
        AllocateParams[] memory allocateParams = AllocateParams({
            operatorSet: OperatorSet(random().Address(), 0),
            strategies: defaultStrategies,
            newMagnitudes: uint64(0.5 ether).toArrayU64()
        }).toArray();

        cheats.expectRevert(InvalidOperatorSet.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
    }

    function test_revert_multiAlloc_modificationAlreadyPending_diffTx() public {
        // Allocate magnitude
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();
        cheats.startPrank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Warp to just before allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY - 1);

        // Attempt to allocate magnitude again
        cheats.expectRevert(ModificationAlreadyPending.selector);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.stopPrank();
    }

    function test_revert_multiAlloc_modificationAlreadyPending_sameTx() public {
        // Allocate magnitude
        AllocateParams[] memory allocateParams = new AllocateParams[](2);
        allocateParams[0] = _randAllocateParams_DefaultOpSet()[0];
        allocateParams[1] = allocateParams[0];

        cheats.expectRevert(ModificationAlreadyPending.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
    }

    function test_revert_allocateZeroMagnitude() public {
        // Allocate exact same magnitude as initial allocation (0)
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();
        allocateParams[0].newMagnitudes[0] = 0;

        cheats.expectRevert(SameMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
    }

    function test_revert_allocateSameMagnitude() public {
        // Allocate nonzero magnitude
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Attempt to allocate no magnitude (ie. same magnitude)
        cheats.expectRevert(SameMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
    }

    function testFuzz_revert_insufficientAllocatableMagnitude(
        Randomness r
    ) public rand(r) {
        // Allocate some magnitude
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Attempt to allocate more magnitude than the operator has
        // uint64 allocatedMag = allocateParams[0].newMagnitudes[0];
        allocateParams[0].newMagnitudes[0] = WAD + 1;
        cheats.expectRevert(InsufficientMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
    }

    function testFuzz_revert_overAllocate(
        Randomness r
    ) public rand(r) {
        uint8 numOpSets = uint8(r.Uint256(2, FUZZ_MAX_OP_SETS));

        // Create and register for operator sets
        OperatorSet[] memory operatorSets = r.OperatorSetArray(defaultAVS, numOpSets);
        _createOperatorSets(operatorSets, defaultStrategies);

        AllocateParams[] memory allocateParams = _randAllocateParams_SingleMockStrategy(operatorSets);
        uint256 randIdx = r.Uint256(0, allocateParams.length - 1);

        allocateParams[randIdx].newMagnitudes[0] = WAD + 1;

        // Overallocate
        cheats.expectRevert(InsufficientMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
    }

    // TODO: yash
    // function test_allocateMaxToMultipleStrategies(
    //     Randomness r
    // ) public rand(r) {
    //     // Create a handful of operator sets under the same AVS, each with a unique strategy
    //     OperatorSet[] memory operatorSets = _newOperatorSets_SingleUniqueStrategy(defaultAVS, r.Uint256(2, 10));

    //     // Register for each operator set
    //     _registerForOperatorSets(defaultOperator, operatorSets);

    //     // Allocate max to each operator set
    //     AllocateParams[] memory allocateParams = new AllocateParams[](operatorSets.length);
    //     for (uint256 i = 0; i < operatorSets.length; i++) {
    //         allocateParams[i] = AllocateParams({
    //             operatorSet: operatorSets[i],
    //             strategies: allocationManager.getStrategiesInOperatorSet(operatorSets[i]),
    //             newMagnitudes: WAD.toArrayU64()
    //         });
    //     }

    //     cheats.prank(defaultOperator);
    //     allocationManager.modifyAllocations(defaultOperator, allocateParams);

    //     // Ensure encumbered magnitude is updated for each strategy
    //     for (uint256 i = 0; i < allocateParams.length; i++) {
    //         assertEq(
    //             WAD,
    //             allocationManager.encumberedMagnitude(defaultOperator, allocateParams[i].strategies[0]),
    //             "encumberedMagnitude not max"
    //         );
    //     }
    // }

    function test_revert_allocateDeallocate_modificationPending() public {
        // Allocate
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Deallocate
        allocateParams[0].newMagnitudes[0] -= 1;
        cheats.expectRevert(ModificationAlreadyPending.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
    }

    function test_revert_deallocateTwice_modificationPending() public {
        // Allocate
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Warp past allocation complete timestsamp
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        allocateParams[0].newMagnitudes[0] -= 1;
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Deallocate again -> expect revert
        cheats.expectRevert(ModificationAlreadyPending.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
    }

    /// @dev Set allocation delay > ALLOCATION_CONFIGURATION_DELAY, allocate,
    /// set allocation delay to < ALLOCATION_CONFIGURATION_DELAY, allocate again
    /// once new delay is sect.
    ///
    /// NOTE: Should be able to allocate faster than `ALLOCATION_CONFIGURATION_DELAY`.
    function testFuzz_ShouldBeAbleToAllocateSoonerThanLastDelay(
        Randomness r
    ) public rand(r) {
        uint32 firstDelay = r.Uint32(ALLOCATION_CONFIGURATION_DELAY, type(uint24).max);
        uint32 secondDelay = r.Uint32(1, ALLOCATION_CONFIGURATION_DELAY - 1);
        uint64 half = 0.5 ether;

        cheats.prank(defaultAVS);
        allocationManager.createOperatorSets(defaultAVS, CreateSetParams(1, defaultStrategies).toArray());

        cheats.startPrank(defaultOperator);

        allocationManager.setAllocationDelay(defaultOperator, firstDelay);

        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(defaultOperatorSet, half));

        allocationManager.setAllocationDelay(defaultOperator, secondDelay);
        cheats.roll(block.number + secondDelay);
        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(OperatorSet(defaultAVS, 1), half));

        cheats.stopPrank();
    }

    function testFuzz_allocate_singleStrat_singleOperatorSet(
        Randomness r
    ) public rand(r) {
        // Create allocation
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();

        // Save vars to check against
        uint64 magnitude = allocateParams[0].newMagnitudes[0];
        uint32 effectBlock = uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Check that the operator has no allocated sets/strats before allocation
        OperatorSet[] memory allocatedSets = allocationManager.getAllocatedSets(defaultOperator);
        IStrategy[] memory allocatedStrats =
            allocationManager.getAllocatedStrategies(defaultOperator, defaultOperatorSet);
        assertEq(allocatedSets.length, 0, "should not have any allocated sets before allocation");
        assertEq(allocatedStrats.length, 0, "should not have any allocated strats before allocation");

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            currentMagnitude: magnitude,
            encumberedMagnitude: magnitude,
            effectBlock: effectBlock
        });

        // Allocate magnitude
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Check storage

        allocatedSets = allocationManager.getAllocatedSets(defaultOperator);
        allocatedStrats = allocationManager.getAllocatedStrategies(defaultOperator, defaultOperatorSet);
        assertEq(allocatedSets.length, 1, "should have a single allocated set");
        assertEq(allocatedSets[0].key(), defaultOperatorSet.key(), "should be allocated to default set");
        assertEq(allocatedStrats.length, 1, "should have a single allocated strategy to default set");
        assertEq(address(allocatedStrats[0]), address(strategyMock), "should have allocated default strat");

        assertEq(
            magnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            WAD - magnitude,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            "allocatableMagnitude not calcualted correctly"
        );
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: 0,
            expectedPendingDiff: int128(uint128(magnitude)),
            expectedEffectBlock: effectBlock
        });

        // Check storage after roll to completion
        cheats.roll(effectBlock);
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: magnitude,
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });
    }

    function testFuzz_allocate_singleStrat_multipleSets(
        Randomness r
    ) public rand(r) {
        uint8 numOpSets = uint8(r.Uint256(1, FUZZ_MAX_OP_SETS));

        // Create and register for operator sets, each with a single default strategy
        OperatorSet[] memory operatorSets = r.OperatorSetArray(defaultAVS, numOpSets);
        AllocateParams[] memory allocateParams = _randAllocateParams_SingleMockStrategy(operatorSets);

        _createOperatorSets(operatorSets, defaultStrategies);
        _registerForOperatorSets(defaultOperator, operatorSets);

        // Save vars to check against
        uint32 effectBlock = uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);
        uint64 usedMagnitude;
        for (uint256 i; i < allocateParams.length; ++i) {
            usedMagnitude += allocateParams[i].newMagnitudes[0];
        }

        // Check that the operator has no allocated sets/strats before allocation
        OperatorSet[] memory allocatedSets = allocationManager.getAllocatedSets(defaultOperator);
        IStrategy[] memory allocatedStrats =
            allocationManager.getAllocatedStrategies(defaultOperator, defaultOperatorSet);
        assertEq(allocatedSets.length, 0, "should not have any allocated sets before allocation");
        assertEq(allocatedStrats.length, 0, "should not have any allocated strats before allocation");

        for (uint256 i; i < allocateParams.length; ++i) {
            _checkAllocationEvents({
                operator: defaultOperator,
                operatorSet: operatorSets[i],
                strategy: strategyMock,
                currentMagnitude: allocateParams[i].newMagnitudes[0],
                encumberedMagnitude: allocateParams[i].newMagnitudes[0],
                effectBlock: effectBlock
            });
        }

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Check storage
        assertEq(
            usedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            WAD - usedMagnitude,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            "allocatableMagnitude not calcualted correctly"
        );

        allocatedSets = allocationManager.getAllocatedSets(defaultOperator);
        assertEq(allocatedSets.length, numOpSets, "should have multiple allocated sets");

        for (uint256 i; i < allocateParams.length; ++i) {
            _checkAllocationStorage({
                operator: defaultOperator,
                operatorSet: operatorSets[i],
                strategy: strategyMock,
                expectedCurrentMagnitude: 0,
                expectedPendingDiff: int128(uint128(allocateParams[i].newMagnitudes[0])),
                expectedEffectBlock: effectBlock
            });

            allocatedStrats = allocationManager.getAllocatedStrategies(defaultOperator, operatorSets[i]);
            assertEq(allocatedStrats.length, 1, "should have a single allocated strategy to each set");
            assertEq(address(allocatedStrats[0]), address(strategyMock), "should have allocated default strat");
            assertEq(allocatedSets[i].key(), operatorSets[i].key(), "should be allocated to expected set");
        }

        // Check storage after roll to completion
        cheats.roll(effectBlock);
        for (uint256 i; i < allocateParams.length; ++i) {
            _checkAllocationStorage({
                operator: defaultOperator,
                operatorSet: operatorSets[i],
                strategy: strategyMock,
                expectedCurrentMagnitude: allocateParams[i].newMagnitudes[0],
                expectedPendingDiff: 0,
                expectedEffectBlock: 0
            });
        }
    }

    function testFuzz_allocateMultipleTimes(
        Randomness r
    ) public rand(r) {
        uint64 firstAlloc = r.Uint64(1, WAD - 1);
        uint64 secondAlloc = r.Uint64(firstAlloc + 1, WAD);

        // Check that the operator has no allocated sets/strats before allocation
        OperatorSet[] memory allocatedSets = allocationManager.getAllocatedSets(defaultOperator);
        IStrategy[] memory allocatedStrats =
            allocationManager.getAllocatedStrategies(defaultOperator, defaultOperatorSet);
        assertEq(allocatedSets.length, 0, "should not have any allocated sets before allocation");
        assertEq(allocatedStrats.length, 0, "should not have any allocated strats before allocation");

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            currentMagnitude: firstAlloc,
            encumberedMagnitude: firstAlloc,
            effectBlock: uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY)
        });

        // Allocate magnitude
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, firstAlloc);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Allocate magnitude again
        allocateParams = _newAllocateParams(defaultOperatorSet, secondAlloc);

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            currentMagnitude: firstAlloc + secondAlloc,
            encumberedMagnitude: firstAlloc + secondAlloc,
            effectBlock: uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY)
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Check storage
        assertEq(
            secondAlloc,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );

        allocatedSets = allocationManager.getAllocatedSets(defaultOperator);
        allocatedStrats = allocationManager.getAllocatedStrategies(defaultOperator, defaultOperatorSet);
        assertEq(allocatedSets.length, 1, "should have a single allocated set");
        assertEq(allocatedSets[0].key(), defaultOperatorSet.key(), "should be allocated to default set");
        assertEq(allocatedStrats.length, 1, "should have a single allocated strategy to default set");
        assertEq(address(allocatedStrats[0]), address(strategyMock), "should have allocated default strat");
    }

    function testFuzz_allocateMaxToMultipleStrategies(
        Randomness r
    ) public rand(r) {
        uint256 numStrats = r.Uint256(2, FUZZ_MAX_STRATS);

        OperatorSet memory operatorSet = OperatorSet(defaultAVS, r.Uint32());
        IStrategy[] memory strategies = r.StrategyArray(numStrats);

        _createOperatorSet(operatorSet, strategies);
        _registerForOperatorSet(defaultOperator, operatorSet);

        for (uint256 i; i < numStrats; ++i) {
            _checkAllocationEvents({
                operator: defaultOperator,
                operatorSet: operatorSet,
                strategy: strategies[i],
                currentMagnitude: WAD,
                encumberedMagnitude: WAD,
                effectBlock: uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY)
            });
        }

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(
            defaultOperator,
            AllocateParams({operatorSet: operatorSet, strategies: strategies, newMagnitudes: WAD.toArrayU64(numStrats)})
                .toArray()
        );

        for (uint256 i; i < numStrats; ++i) {
            assertEq(
                WAD,
                allocationManager.encumberedMagnitude(defaultOperator, strategies[i]),
                "encumberedMagnitude not max"
            );
        }
    }

    /**
     * Allocates to `firstMod` magnitude and then deallocate to `secondMod` magnitude
     * Validates the storage
     * - 1. After deallocation is called
     * - 2. After the deallocationd delay is hit
     * - 3. After the deallocation queue is cleared
     */
    function testFuzz_allocate_deallocate_whenRegistered(
        Randomness r
    ) public rand(r) {
        // Bound allocation and deallocation
        uint64 firstMod = r.Uint64(1, WAD);
        uint64 secondMod = r.Uint64(0, firstMod - 1);

        // Allocate magnitude to default registered set
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, firstMod);

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            currentMagnitude: firstMod,
            encumberedMagnitude: firstMod,
            effectBlock: uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY)
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        allocateParams = _newAllocateParams(defaultOperatorSet, secondMod);

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            currentMagnitude: firstMod - secondMod,
            encumberedMagnitude: firstMod - secondMod,
            effectBlock: uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY)
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Check storage after dealloc
        assertEq(
            firstMod,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should not be updated"
        );
        assertEq(
            WAD - firstMod,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            "allocatableMagnitude not calcualted correctly"
        );

        uint32 effectBlock = uint32(block.number + DEALLOCATION_DELAY);

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: firstMod,
            expectedPendingDiff: -int128(uint128(firstMod - secondMod)),
            expectedEffectBlock: effectBlock
        });

        // Check storage after roll to completion
        cheats.roll(effectBlock);
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: secondMod,
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });
        assertEq(
            firstMod,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should not be updated"
        );

        // Check storage after clearing deallocation queue
        allocationManager.clearDeallocationQueue(defaultOperator, strategyMock.toArray(), uint16(1).toArrayU16());
        assertEq(
            secondMod,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should be updated"
        );
    }

    /**
     * Allocates to an operator set, then fully deallocates when not registered to the set.
     * Checks that deallocation is instant and can be reallocated instantly.
     */
    function testFuzz_allocate_fullyDeallocate_reallocate_WhenNotRegistered(
        Randomness r
    ) public rand(r) {
        // Bound allocation and deallocation
        uint64 firstMod = r.Uint64(1, WAD);

        // Create a new operator sets that the operator is not registered for
        OperatorSet memory operatorSetA = _createOperatorSet(OperatorSet(defaultAVS, r.Uint32()), defaultStrategies);
        OperatorSet memory operatorSetB = _createOperatorSet(OperatorSet(defaultAVS, r.Uint32()), defaultStrategies);

        // Allocate magnitude to operator set
        AllocateParams[] memory allocateParams = _newAllocateParams(operatorSetA, firstMod);

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: operatorSetA,
            strategy: strategyMock,
            currentMagnitude: firstMod,
            encumberedMagnitude: firstMod,
            effectBlock: uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY)
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        assertEq(
            firstMod,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should equal firstMod"
        );

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate instantly and reallocate all magnitude to second operator set
        allocateParams = new AllocateParams[](2);
        allocateParams[0] = _newAllocateParams(operatorSetA, 0)[0];
        allocateParams[1] = _newAllocateParams(operatorSetB, firstMod)[0];

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: operatorSetA,
            strategy: strategyMock,
            currentMagnitude: 0,
            encumberedMagnitude: 0,
            effectBlock: 0
        });
        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: operatorSetB,
            strategy: strategyMock,
            currentMagnitude: firstMod,
            encumberedMagnitude: firstMod,
            effectBlock: uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY)
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Check storage after dealloc
        assertEq(
            firstMod,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should not be changed"
        );
        assertEq(
            WAD - firstMod,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            "allocatableMagnitude not calculated correctly"
        );

        // Check operator set A
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: operatorSetA,
            strategy: strategyMock,
            expectedCurrentMagnitude: 0,
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });

        // Check operator set B
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: operatorSetB,
            strategy: strategyMock,
            expectedCurrentMagnitude: 0,
            expectedPendingDiff: int64(firstMod),
            expectedEffectBlock: block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY
        });
    }

    /**
     * Allocate to an operator set using magnitude that is only available if the deallocation
     * queue is cleared
     */
    function testFuzz_allocate_fromClearedDeallocQueue(
        Randomness r
    ) public rand(r) {
        uint256 numOpSets = r.Uint256(1, FUZZ_MAX_OP_SETS);

        // Create multiple operator sets, register, and allocate to each. Ensure all magnitude is fully allocated.
        OperatorSet[] memory deallocSets = r.OperatorSetArray(defaultAVS, numOpSets);
        _createOperatorSets(deallocSets, defaultStrategies);
        _registerForOperatorSets(defaultOperator, deallocSets);
        AllocateParams[] memory allocateParams = _randAllocateParams_SingleMockStrategy_AllocAll(deallocSets);

        for (uint256 i; i < numOpSets; ++i) {
            _checkAllocationEvents({
                operator: defaultOperator,
                operatorSet: deallocSets[i],
                strategy: strategyMock,
                currentMagnitude: WAD,
                encumberedMagnitude: WAD,
                effectBlock: uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY)
            });
        }

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        assertEq(
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            0,
            "operator should not have any remaining allocatable magnitude"
        );

        // Move forward to allocation completion
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate fully from each operator set
        AllocateParams[] memory deallocateParams = _newAllocateParams(deallocSets, 0);

        for (uint256 i; i < numOpSets; ++i) {
            _checkAllocationEvents({
                operator: defaultOperator,
                operatorSet: deallocSets[i],
                strategy: strategyMock,
                currentMagnitude: 0,
                encumberedMagnitude: 0,
                effectBlock: 0
            });
        }

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);

        assertEq(
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            0,
            "operator should still not have any allocatable magnitude"
        );

        // Move forward to deallocation completion
        cheats.roll(block.number + DEALLOCATION_DELAY);

        // Check that we now have sufficient allocatable magnitude
        assertEq(
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            WAD,
            "operator should have all magnitude allocatable"
        );

        // Create and register for a new operator set with the same default strategy.
        // If we try to allocate to this new set, it should clear the deallocation queue,
        // allowing all magnitude to be allocated
        OperatorSet memory finalOpSet = _createOperatorSet(OperatorSet(defaultAVS, r.Uint32()), defaultStrategies);
        _registerForOperatorSet(defaultOperator, finalOpSet);
        AllocateParams[] memory finalAllocParams = _newAllocateParams(finalOpSet, WAD);

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: finalOpSet,
            strategy: strategyMock,
            currentMagnitude: WAD,
            encumberedMagnitude: WAD,
            effectBlock: uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY)
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, finalAllocParams);

        // Check that all magnitude will be allocated to the new set, and each prior set
        // has a zeroed-out allocation
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: finalOpSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: 0,
            expectedPendingDiff: int64(WAD),
            expectedEffectBlock: block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY
        });
        assertEq(
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            0,
            "operator should not have any remaining allocatable magnitude"
        );
        assertEq(
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            WAD,
            "all magnitude should be allocated"
        );

        for (uint256 i; i < deallocSets.length; ++i) {
            _checkAllocationStorage({
                operator: defaultOperator,
                operatorSet: deallocSets[i],
                strategy: strategyMock,
                expectedCurrentMagnitude: 0,
                expectedPendingDiff: 0,
                expectedEffectBlock: 0
            });
        }
    }

    function test_deallocate_all() public {
        // Allocate
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, WAD);

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            currentMagnitude: WAD,
            encumberedMagnitude: WAD,
            effectBlock: uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY)
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        allocateParams[0].newMagnitudes[0] = 0;

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            currentMagnitude: 0,
            encumberedMagnitude: 0,
            effectBlock: 0
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Warp to completion and clear deallocation queue
        cheats.roll(block.number + DEALLOCATION_DELAY);
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, uint16(1).toArrayU16());

        // Check storage
        assertEq(
            0,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should be updated"
        );
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: 0,
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });
    }

    function testFuzz_allocate_deallocate_singleStrat_multipleOperatorSets(
        Randomness r
    ) public rand(r) {
        uint8 numOpSets = uint8(r.Uint256(1, FUZZ_MAX_OP_SETS));

        // Create and register for operator sets, each with a single default strategy
        OperatorSet[] memory operatorSets = r.OperatorSetArray(defaultAVS, numOpSets);
        _createOperatorSets(operatorSets, defaultStrategies);
        _registerForOperatorSets(defaultOperator, operatorSets);

        (AllocateParams[] memory allocateParams, AllocateParams[] memory deallocateParams) =
            _randAllocAndDeallocParams_SingleMockStrategy(operatorSets);

        // Allocate
        for (uint256 i; i < allocateParams.length; ++i) {
            _checkAllocationEvents({
                operator: defaultOperator,
                operatorSet: operatorSets[i],
                strategy: strategyMock,
                currentMagnitude: allocateParams[i].newMagnitudes[0],
                encumberedMagnitude: allocateParams[i].newMagnitudes[0],
                effectBlock: uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY)
            });
        }

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        uint64 encumberedMagnitudeAfterAllocation = allocationManager.encumberedMagnitude(defaultOperator, strategyMock);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Calculate post-deallocation magnitude
        // We can add each entry to this value because each operator set is using the same strategy
        uint64 postDeallocMag;
        for (uint256 i; i < deallocateParams.length; ++i) {
            postDeallocMag += deallocateParams[i].newMagnitudes[0];
            _checkAllocationEvents({
                operator: defaultOperator,
                operatorSet: operatorSets[i],
                strategy: strategyMock,
                currentMagnitude: deallocateParams[i].newMagnitudes[0],
                encumberedMagnitude: deallocateParams[i].newMagnitudes[0],
                effectBlock: uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY)
            });
        }

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);

        // Check storage after dealloc
        assertEq(
            encumberedMagnitudeAfterAllocation,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should not be updated"
        );

        for (uint256 i; i < allocateParams.length; ++i) {
            _checkAllocationStorage({
                operator: defaultOperator,
                operatorSet: operatorSets[i],
                strategy: strategyMock,
                expectedCurrentMagnitude: allocateParams[i].newMagnitudes[0],
                expectedPendingDiff: -int64(allocateParams[i].newMagnitudes[0] - deallocateParams[i].newMagnitudes[0]),
                expectedEffectBlock: block.number + DEALLOCATION_DELAY
            });
        }

        // Check storage after roll to completion
        cheats.roll(block.number + DEALLOCATION_DELAY);

        for (uint256 i; i < allocateParams.length; ++i) {
            _checkAllocationStorage({
                operator: defaultOperator,
                operatorSet: operatorSets[i],
                strategy: strategyMock,
                expectedCurrentMagnitude: deallocateParams[i].newMagnitudes[0],
                expectedPendingDiff: 0,
                expectedEffectBlock: 0
            });
        }

        // Clear deallocation queue
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, type(uint16).max.toArrayU16());
        // Check storage after clearing deallocation queue
        assertEq(
            postDeallocMag,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should be updated"
        );
    }

    function testFuzz_MultipleSetsAndStrats(
        Randomness r
    ) public rand(r) {
        uint256 numAllocations = r.Uint256(2, FUZZ_MAX_ALLOCATIONS);
        uint256 numStrats = r.Uint256(2, FUZZ_MAX_STRATS);

        AllocateParams[] memory allocateParams = r.AllocateParams(defaultAVS, numAllocations, numStrats);
        AllocateParams[] memory deallocateParams = r.DeallocateParams(allocateParams);
        CreateSetParams[] memory createSetParams = r.CreateSetParams(allocateParams);

        cheats.prank(defaultAVS);
        allocationManager.createOperatorSets(defaultAVS, createSetParams);

        for (uint256 i; i < allocateParams.length; ++i) {
            for (uint256 j; j < allocateParams[i].strategies.length; ++j) {
                _checkAllocationEvents({
                    operator: defaultOperator,
                    operatorSet: allocateParams[i].operatorSet,
                    strategy: allocateParams[i].strategies[j],
                    currentMagnitude: allocateParams[i].newMagnitudes[j],
                    encumberedMagnitude: allocateParams[i].newMagnitudes[j],
                    effectBlock: uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY)
                });
            }
        }

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        for (uint256 i; i < allocateParams.length; ++i) {
            for (uint256 j = 0; j < allocateParams[i].strategies.length; j++) {
                _checkAllocationStorage({
                    operator: defaultOperator,
                    operatorSet: allocateParams[i].operatorSet,
                    strategy: allocateParams[i].strategies[j],
                    expectedCurrentMagnitude: 0,
                    expectedPendingDiff: int64(allocateParams[i].newMagnitudes[j]),
                    expectedEffectBlock: block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY
                });
            }
        }

        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        for (uint256 i; i < allocateParams.length; ++i) {
            for (uint256 j = 0; j < allocateParams[i].strategies.length; j++) {
                _checkAllocationStorage({
                    operator: defaultOperator,
                    operatorSet: allocateParams[i].operatorSet,
                    strategy: allocateParams[i].strategies[j],
                    expectedCurrentMagnitude: allocateParams[i].newMagnitudes[j],
                    expectedPendingDiff: 0,
                    expectedEffectBlock: 0
                });
            }
        }

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);

        // Deallocations are immediate if the operator's allocation is not slashable.
        for (uint256 i; i < allocateParams.length; ++i) {
            for (uint256 j = 0; j < allocateParams[i].strategies.length; j++) {
                _checkAllocationStorage({
                    operator: defaultOperator,
                    operatorSet: allocateParams[i].operatorSet,
                    strategy: allocateParams[i].strategies[j],
                    expectedCurrentMagnitude: deallocateParams[i].newMagnitudes[j],
                    expectedPendingDiff: 0,
                    expectedEffectBlock: 0
                });
            }
        }
    }
}

contract AllocationManagerUnitTests_ClearDeallocationQueue is AllocationManagerUnitTests {
    using SingleItemArrayLib for *;

    /// -----------------------------------------------------------------------
    /// clearModificationQueue()
    /// -----------------------------------------------------------------------

    function test_revert_paused() public {
        allocationManager.pause(2 ** PAUSED_MODIFY_ALLOCATIONS);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        allocationManager.clearDeallocationQueue(defaultOperator, new IStrategy[](0), new uint16[](0));
    }

    function test_revert_arrayMismatch() public {
        IStrategy[] memory strategies = new IStrategy[](1);
        uint16[] memory numToClear = new uint16[](2);

        cheats.expectRevert(InputArrayLengthMismatch.selector);
        allocationManager.clearDeallocationQueue(defaultOperator, strategies, numToClear);
    }

    /**
     * @notice Allocates magnitude to an operator and then
     * - Clears deallocation queue when only an allocation exists
     * - Clears deallocation queue when the alloc can be completed - asserts emit has been emitted
     * - Validates storage after the second clear
     */
    function testFuzz_allocate(
        Randomness r
    ) public rand(r) {
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();

        // Allocate magnitude
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Attempt to clear queue, assert no events emitted
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());
        Vm.Log[] memory entries = vm.getRecordedLogs();
        assertEq(0, entries.length, "should not have emitted any events");

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Clear queue - this is a noop
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());
        entries = vm.getRecordedLogs();
        assertEq(0, entries.length, "should not have emitted any events 2");

        // Validate allocation is no longer pending
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: allocateParams[0].newMagnitudes[0],
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });
    }

    /**
     * @notice Allocates magnitude to an operator registered for some operator sets, and then
     * - Clears deallocation queue when nothing can be completed
     * - After the first clear, asserts the allocation info takes into account the deallocation
     * - Clears deallocation queue when the dealloc can be completed
     * - Assert events & validates storage after the deallocateParams are completed
     */
    function testFuzz_allocate_deallocate_whenRegistered(
        Randomness r
    ) public rand(r) {
        // Generate a random allocation and subsequent deallocation from the default operator set
        (AllocateParams[] memory allocateParams, AllocateParams[] memory deallocateParams) =
            _randAllocAndDeallocParams_SingleMockStrategy(defaultOperatorSet.toArray());

        // Allocate
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Roll to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);

        // Clear queue - since we have not rolled forward, this should be a no-op
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());
        assertEq(
            allocateParams[0].newMagnitudes[0],
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should not be updated"
        );

        // Validate storage - encumbered magnitude should just be allocateParams (we only have 1 allocation)
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: allocateParams[0].newMagnitudes[0],
            expectedPendingDiff: -int128(uint128(allocateParams[0].newMagnitudes[0] - deallocateParams[0].newMagnitudes[0])),
            expectedEffectBlock: block.number + DEALLOCATION_DELAY
        });

        // Warp to deallocation complete block
        cheats.roll(block.number + DEALLOCATION_DELAY);

        // Clear queue
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());

        // Validate storage - encumbered magnitude should just be deallocateParams (we only have 1 deallocation)
        assertEq(
            deallocateParams[0].newMagnitudes[0],
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should be updated"
        );
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedCurrentMagnitude: deallocateParams[0].newMagnitudes[0],
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });
    }

    /**
     * Allocates, deallocates, and then allocates again. Asserts that
     * - The deallocation does not block state updates from the second allocation, even though the allocation has an earlier
     *   effect block
     */
    function test_allocate_deallocate_allocate_whenRegistered() public {
        // Allocate half of mag to default operator set
        AllocateParams[] memory firstAllocation = _newAllocateParams(defaultOperatorSet, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, firstAllocation);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate half from default operator set
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY);
        AllocateParams[] memory firstDeallocation = _newAllocateParams(defaultOperatorSet, 25e16);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, firstDeallocation);
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, defaultOperatorSet, strategyMock);
        assertEq(deallocationEffectBlock, allocation.effectBlock, "effect block not correct");

        // Create and register for a new operator set
        OperatorSet memory newOperatorSet =
            _createOperatorSet(OperatorSet(defaultAVS, random().Uint32()), defaultStrategies);
        _registerForOperatorSet(defaultOperator, newOperatorSet);

        // Allocate 33e16 mag to new operator set
        uint32 allocationEffectBlock = uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);
        AllocateParams[] memory secondAllocation = _newAllocateParams(newOperatorSet, 33e16);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, secondAllocation);
        allocation = allocationManager.getAllocation(defaultOperator, newOperatorSet, strategyMock);
        assertEq(allocationEffectBlock, allocation.effectBlock, "effect block not correct");

        // Warp to allocation effect block & clear the queue
        cheats.roll(allocationEffectBlock);
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());

        // Validate `getAllocatableMagnitude`. Allocatable magnitude should be the difference between the max magnitude and the encumbered magnitude
        uint64 allocatableMagnitude = allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock);
        assertEq(WAD - 33e16 - 5e17, allocatableMagnitude, "allocatableMagnitude not correct");

        // Validate that we can allocate again for opset2. This should not revert
        AllocateParams[] memory thirdAllocation = _newAllocateParams(newOperatorSet, 10e16);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, thirdAllocation);
    }

    /**
     * Allocates to opset1, allocates to opset2, deallocates from opset1. Asserts that the allocation, which has a higher
     * effect block is not blocking the deallocation.
     * The allocs/deallocs looks like
     * 1. (allocation, opSet2, mag: 5e17, effectBlock: 50th day)
     * 2. (deallocation, opSet1, mag: 0, effectBlock: 42.5 day)
     *
     * The deallocation queue looks like
     * 1. (deallocation, opSet1, mag: 0, effectBlock: 42.5 day)
     */
    function test_regression_deallocationNotBlocked() public {
        // Set allocation delay to be longer than the deallocation delay
        uint32 allocationDelay = DEALLOCATION_DELAY * 2;
        cheats.prank(defaultOperator);
        allocationManager.setAllocationDelay(defaultOperator, allocationDelay);
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);
        (, uint32 storedDelay) = allocationManager.getAllocationDelay(defaultOperator);
        assertEq(allocationDelay, storedDelay, "allocation delay not valid");

        // Allocate half of mag to default operator set
        AllocateParams[] memory firstAllocation = _newAllocateParams(defaultOperatorSet, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, firstAllocation);
        cheats.roll(block.number + allocationDelay);

        // Create and register for a second operator set
        OperatorSet memory newOperatorSet =
            _createOperatorSet(OperatorSet(defaultAVS, random().Uint32()), defaultStrategies);
        _registerForOperatorSet(defaultOperator, newOperatorSet);

        // Allocate half of mag to opset2
        AllocateParams[] memory secondAllocation = _newAllocateParams(newOperatorSet, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, secondAllocation);

        uint32 allocationEffectBlock = uint32(block.number + allocationDelay);
        Allocation memory allocation = allocationManager.getAllocation(defaultOperator, newOperatorSet, strategyMock);
        assertEq(allocationEffectBlock, allocation.effectBlock, "effect block not correct");

        // Deallocate all from opSet1
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY);
        AllocateParams[] memory firstDeallocation = _newAllocateParams(defaultOperatorSet, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, firstDeallocation);
        allocation = allocationManager.getAllocation(defaultOperator, defaultOperatorSet, strategyMock);
        assertEq(deallocationEffectBlock, allocation.effectBlock, "effect block not correct");

        // Warp to deallocation effect block & clear the queue
        cheats.roll(deallocationEffectBlock);
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());

        // At this point, we should be able to allocate again to opSet1 AND have only 5e17 encumbered magnitude
        assertEq(
            5e17,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumbered magnitude not correct"
        );
        AllocateParams[] memory thirdAllocation = _newAllocateParams(defaultOperatorSet, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, thirdAllocation);
    }
}

contract AllocationManagerUnitTests_SetAllocationDelay is AllocationManagerUnitTests {
    /// -----------------------------------------------------------------------
    /// setAllocationDelay() + getAllocationDelay()
    /// -----------------------------------------------------------------------

    address operatorToSet = address(0x1);

    function setUp() public override {
        AllocationManagerUnitTests.setUp();
        _registerOperator(operatorToSet);
    }

    function test_revert_callerNotOperator() public {
        delegationManagerMock.setIsOperator(operatorToSet, false);
        cheats.prank(operatorToSet);
        cheats.expectRevert(OperatorNotRegistered.selector);
        allocationManager.setAllocationDelay(operatorToSet, 1);
    }

    function test_revert_callerNotAuthorized() public {
        cheats.expectRevert(InvalidCaller.selector);
        allocationManager.setAllocationDelay(operatorToSet, 1);
    }

    function testFuzz_setDelay(
        Randomness r
    ) public rand(r) {
        uint32 delay = r.Uint32(0, type(uint32).max);

        // Set delay
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AllocationDelaySet(operatorToSet, delay, uint32(block.number + ALLOCATION_CONFIGURATION_DELAY));
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(operatorToSet, delay);

        // Check values after set
        (bool isSet, uint32 returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertFalse(isSet, "isSet should not be set");
        assertEq(0, returnedDelay, "returned delay should be 0");

        // Warp to effect block
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);

        // Check values after config delay
        (isSet, returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(delay, returnedDelay, "delay not set");
    }

    function test_fuzz_setDelay_multipleTimesWithinConfigurationDelay(
        Randomness r
    ) public rand(r) {
        uint32 firstDelay = r.Uint32(1, type(uint32).max);
        uint32 secondDelay = r.Uint32(1, type(uint32).max);

        // Set delay
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(operatorToSet, firstDelay);

        // Warp just before effect block
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY - 1);

        // Set delay again
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AllocationDelaySet(operatorToSet, secondDelay, uint32(block.number + ALLOCATION_CONFIGURATION_DELAY));
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(operatorToSet, secondDelay);

        // Warp to effect block of first delay
        cheats.roll(block.number + 1);

        // Assert that the delay is still not set
        (bool isSet, uint32 returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertFalse(isSet, "isSet should not be set");
        assertEq(0, returnedDelay, "returned delay should be 0");

        // Warp to effect block of second delay
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);
        (isSet, returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(secondDelay, returnedDelay, "delay not set");
    }

    function testFuzz_multipleDelays(
        Randomness r
    ) public rand(r) {
        uint32 firstDelay = r.Uint32(1, type(uint32).max);
        uint32 secondDelay = r.Uint32(1, type(uint32).max);

        // Set delay
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(operatorToSet, firstDelay);

        // Warp to effect block of first delay
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);

        // Set delay again
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(operatorToSet, secondDelay);

        // Assert that first delay is set
        (bool isSet, uint32 returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(firstDelay, returnedDelay, "delay not set");

        // Warp to effect block of second delay
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);

        // Check values after second delay
        (isSet, returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(secondDelay, returnedDelay, "delay not set");
    }

    function testFuzz_setDelay_DMCaller(
        Randomness r
    ) public rand(r) {
        uint32 delay = r.Uint32(1, type(uint32).max);

        cheats.prank(address(delegationManagerMock));
        allocationManager.setAllocationDelay(operatorToSet, delay);

        // Warp to effect block
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);
        (bool isSet, uint32 returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(delay, returnedDelay, "delay not set");
    }
}

contract AllocationManagerUnitTests_registerForOperatorSets is AllocationManagerUnitTests {
    using SingleItemArrayLib for *;

    RegisterParams defaultRegisterParams;

    function setUp() public override {
        AllocationManagerUnitTests.setUp();
        defaultRegisterParams = RegisterParams(defaultAVS, defaultOperatorSet.id.toArrayU32(), "");
    }

    function test_registerForOperatorSets_Paused() public {
        allocationManager.pause(2 ** PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        allocationManager.registerForOperatorSets(defaultOperator, defaultRegisterParams);
    }

    function testFuzz_registerForOperatorSets_InvalidOperator_x(
        Randomness r
    ) public rand(r) {
        address operator = r.Address();
        cheats.prank(operator);
        cheats.expectRevert(InvalidOperator.selector);
        allocationManager.registerForOperatorSets(operator, defaultRegisterParams);
    }

    function testFuzz_registerForOperatorSets_InvalidOperatorSet(
        Randomness r
    ) public rand(r) {
        cheats.prank(defaultOperator);
        cheats.expectRevert(InvalidOperatorSet.selector);
        defaultRegisterParams.operatorSetIds[0] = 1; // invalid id
        allocationManager.registerForOperatorSets(defaultOperator, defaultRegisterParams); // invalid id
    }

    function testFuzz_registerForOperatorSets_AlreadyMemberOfSet(
        Randomness r
    ) public rand(r) {
        cheats.prank(defaultOperator);
        cheats.expectRevert(AlreadyMemberOfSet.selector);
        allocationManager.registerForOperatorSets(defaultOperator, defaultRegisterParams);
    }

    function testFuzz_registerForOperatorSets_Correctness(
        Randomness r
    ) public rand(r) {
        address operator = r.Address();
        uint256 numOpSets = r.Uint256(1, FUZZ_MAX_OP_SETS);
        uint32[] memory operatorSetIds = new uint32[](numOpSets);
        CreateSetParams[] memory createSetParams = new CreateSetParams[](numOpSets);

        _registerOperator(operator);

        for (uint256 i; i < numOpSets; ++i) {
            operatorSetIds[i] = r.Uint32(1, type(uint32).max);
            createSetParams[i].operatorSetId = operatorSetIds[i];
            createSetParams[i].strategies = defaultStrategies;
        }

        cheats.prank(defaultAVS);
        allocationManager.createOperatorSets(defaultAVS, createSetParams);

        for (uint256 j; j < numOpSets; ++j) {
            cheats.expectEmit(true, true, false, false, address(allocationManager));
            emit OperatorAddedToOperatorSet(operator, OperatorSet(defaultAVS, operatorSetIds[j]));
        }

        cheats.expectCall(
            defaultAVS, abi.encodeWithSelector(IAVSRegistrar.registerOperator.selector, operator, operatorSetIds, "")
        );

        cheats.prank(operator);
        allocationManager.registerForOperatorSets(operator, RegisterParams(defaultAVS, operatorSetIds, ""));

        assertEq(allocationManager.getRegisteredSets(operator).length, numOpSets, "should be registered for all sets");

        for (uint256 k; k < numOpSets; ++k) {
            assertEq(
                allocationManager.getMembers(OperatorSet(defaultAVS, operatorSetIds[k]))[0],
                operator,
                "should be member of set"
            );
        }
    }
}

contract AllocationManagerUnitTests_deregisterFromOperatorSets is AllocationManagerUnitTests {
    using SingleItemArrayLib for *;

    DeregisterParams defaultDeregisterParams;

    function setUp() public override {
        AllocationManagerUnitTests.setUp();
        defaultDeregisterParams = DeregisterParams(defaultOperator, defaultAVS, defaultOperatorSet.id.toArrayU32());
    }

    function test_deregisterFromOperatorSets_Paused() public {
        allocationManager.pause(2 ** PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        allocationManager.deregisterFromOperatorSets(defaultDeregisterParams);
    }

    function test_deregisterFromOperatorSets_revert_invalidCaller_notOperator() public {
        address randomOperator = address(0x1);
        defaultDeregisterParams.operator = randomOperator;

        cheats.expectRevert(InvalidCaller.selector);
        allocationManager.deregisterFromOperatorSets(defaultDeregisterParams);
    }

    function test_deregisterFromOperatorSets_revert_invalidCaller_notAVS() public {
        address randomAVS = address(0x1);

        cheats.prank(randomAVS);
        cheats.expectRevert(InvalidCaller.selector);
        allocationManager.deregisterFromOperatorSets(defaultDeregisterParams);
    }

    function testFuzz_deregisterFromOperatorSets_InvalidOperatorSet(
        Randomness r
    ) public rand(r) {
        defaultDeregisterParams.operatorSetIds = uint32(1).toArrayU32(); // invalid id
        cheats.prank(defaultOperator);
        cheats.expectRevert(InvalidOperatorSet.selector);
        allocationManager.deregisterFromOperatorSets(defaultDeregisterParams);
    }

    function testFuzz_deregisterFromOperatorSets_NotMemberOfSet(
        Randomness r
    ) public rand(r) {
        defaultDeregisterParams.operator = r.Address();
        cheats.prank(defaultDeregisterParams.operator);
        cheats.expectRevert(NotMemberOfSet.selector);
        allocationManager.deregisterFromOperatorSets(defaultDeregisterParams);
    }

    function testFuzz_deregisterFromOperatorSets_Correctness(
        Randomness r
    ) public rand(r) {
        uint256 numOpSets = r.Uint256(1, FUZZ_MAX_OP_SETS);
        uint32[] memory operatorSetIds = new uint32[](numOpSets);
        CreateSetParams[] memory createSetParams = new CreateSetParams[](numOpSets);

        for (uint256 i; i < numOpSets; ++i) {
            operatorSetIds[i] = r.Uint32(1, type(uint32).max);
            createSetParams[i].operatorSetId = operatorSetIds[i];
            createSetParams[i].strategies = defaultStrategies;
        }

        cheats.prank(defaultAVS);
        allocationManager.createOperatorSets(defaultAVS, createSetParams);

        address operator = r.Address();
        _registerOperator(operator);

        cheats.prank(operator);
        allocationManager.registerForOperatorSets(operator, RegisterParams(defaultAVS, operatorSetIds, ""));

        for (uint256 j; j < numOpSets; ++j) {
            cheats.expectEmit(true, true, false, false, address(allocationManager));
            emit OperatorRemovedFromOperatorSet(operator, OperatorSet(defaultAVS, operatorSetIds[j]));
        }

        cheats.expectCall(
            defaultAVS, abi.encodeWithSelector(IAVSRegistrar.deregisterOperator.selector, operator, operatorSetIds)
        );

        cheats.prank(operator);
        allocationManager.deregisterFromOperatorSets(DeregisterParams(operator, defaultAVS, operatorSetIds));

        assertEq(allocationManager.getRegisteredSets(operator).length, 0, "should not be registered for any sets");

        for (uint256 k; k < numOpSets; ++k) {
            assertEq(
                allocationManager.getMemberCount(OperatorSet(defaultAVS, operatorSetIds[k])),
                0,
                "should not be member of set"
            );
        }
    }
}

contract AllocationManagerUnitTests_addStrategiesToOperatorSet is AllocationManagerUnitTests {
    function test_addStrategiesToOperatorSet_InvalidOperatorSet() public {
        cheats.prank(defaultAVS);
        cheats.expectRevert(InvalidOperatorSet.selector);
        allocationManager.addStrategiesToOperatorSet(defaultAVS, 1, defaultStrategies);
    }

    function test_addStrategiesToOperatorSet_StrategyAlreadyInOperatorSet() public {
        cheats.prank(defaultAVS);
        cheats.expectRevert(StrategyAlreadyInOperatorSet.selector);
        allocationManager.addStrategiesToOperatorSet(defaultAVS, defaultOperatorSet.id, defaultStrategies);
    }

    function testFuzz_addStrategiesToOperatorSet_Correctness(
        Randomness r
    ) public rand(r) {
        uint256 numStrategies = r.Uint256(1, FUZZ_MAX_STRATS);

        IStrategy[] memory strategies = new IStrategy[](numStrategies);

        for (uint256 i; i < numStrategies; ++i) {
            strategies[i] = IStrategy(r.Address());
            cheats.expectEmit(true, false, false, false, address(allocationManager));
            emit StrategyAddedToOperatorSet(defaultOperatorSet, strategies[i]);
        }

        cheats.prank(defaultAVS);
        allocationManager.addStrategiesToOperatorSet(defaultAVS, defaultOperatorSet.id, strategies);

        IStrategy[] memory strategiesInSet = allocationManager.getStrategiesInOperatorSet(defaultOperatorSet);

        for (uint256 j; j < numStrategies; ++j) {
            assertTrue(strategiesInSet[j + 1] == strategies[j], "should be strat of set");
        }
    }
}

contract AllocationManagerUnitTests_removeStrategiesFromOperatorSet is AllocationManagerUnitTests {
    using SingleItemArrayLib for *;

    function test_removeStrategiesFromOperatorSet_InvalidOperatorSet() public {
        cheats.prank(defaultAVS);
        cheats.expectRevert(InvalidOperatorSet.selector);
        allocationManager.removeStrategiesFromOperatorSet(defaultAVS, 1, defaultStrategies);
    }

    function test_removeStrategiesFromOperatorSet_StrategyNotInOperatorSet() public {
        cheats.prank(defaultAVS);
        cheats.expectRevert(StrategyNotInOperatorSet.selector);
        allocationManager.removeStrategiesFromOperatorSet(
            defaultAVS, defaultOperatorSet.id, IStrategy(random().Address()).toArray()
        );
    }

    function testFuzz_removeStrategiesFromOperatorSet_Correctness(
        Randomness r
    ) public rand(r) {
        uint256 numStrategies = r.Uint256(1, FUZZ_MAX_STRATS);
        IStrategy[] memory strategies = r.StrategyArray(numStrategies);

        cheats.prank(defaultAVS);
        allocationManager.addStrategiesToOperatorSet(defaultAVS, defaultOperatorSet.id, strategies);

        for (uint256 i; i < numStrategies; ++i) {
            cheats.expectEmit(true, false, false, false, address(allocationManager));
            emit StrategyRemovedFromOperatorSet(defaultOperatorSet, strategies[i]);
        }

        assertEq(
            allocationManager.getStrategiesInOperatorSet(defaultOperatorSet).length, numStrategies + 1, "sanity check"
        );

        cheats.prank(defaultAVS);
        allocationManager.removeStrategiesFromOperatorSet(defaultAVS, defaultOperatorSet.id, strategies);

        // The orginal strategy should still be in the operator set.
        assertEq(
            allocationManager.getStrategiesInOperatorSet(defaultOperatorSet).length, 1, "should not be strat of set"
        );
    }
}

contract AllocationManagerUnitTests_createOperatorSets is AllocationManagerUnitTests {
    using SingleItemArrayLib for *;

    function test_createOperatorSets_InvalidOperatorSet() public {
        cheats.prank(defaultAVS);
        cheats.expectRevert(InvalidOperatorSet.selector);
        allocationManager.createOperatorSets(defaultAVS, CreateSetParams(defaultOperatorSet.id, defaultStrategies).toArray());
    }

    function testFuzz_createOperatorSets_Correctness(
        Randomness r
    ) public rand(r) {
        address avs = r.Address();
        uint256 numOpSets = r.Uint256(1, FUZZ_MAX_OP_SETS);
        uint256 numStrategies = r.Uint256(1, FUZZ_MAX_STRATS);

        CreateSetParams[] memory createSetParams = new CreateSetParams[](numOpSets);

        for (uint256 i; i < numOpSets; ++i) {
            createSetParams[i].operatorSetId = r.Uint32(1, type(uint32).max);
            createSetParams[i].strategies = r.StrategyArray(numStrategies);
            cheats.expectEmit(true, false, false, false, address(allocationManager));
            emit OperatorSetCreated(OperatorSet(avs, createSetParams[i].operatorSetId));
            for (uint256 j; j < numStrategies; ++j) {
                cheats.expectEmit(true, false, false, false, address(allocationManager));
                emit StrategyAddedToOperatorSet(
                    OperatorSet(avs, createSetParams[i].operatorSetId), createSetParams[i].strategies[j]
                );
            }
        }

        cheats.prank(avs);
        allocationManager.createOperatorSets(avs, createSetParams);

        for (uint256 k; k < numOpSets; ++k) {
            OperatorSet memory opSet = OperatorSet(avs, createSetParams[k].operatorSetId);
            assertTrue(allocationManager.isOperatorSet(opSet), "should be operator set");
            IStrategy[] memory strategiesInSet = allocationManager.getStrategiesInOperatorSet(opSet);
            assertEq(strategiesInSet.length, numStrategies, "strategiesInSet length should be numStrategies");
            for (uint256 l; l < numStrategies; ++l) {
                assertTrue(
                    allocationManager.getStrategiesInOperatorSet(opSet)[l] == createSetParams[k].strategies[l],
                    "should be strat of set"
                );
            }
        }

        assertEq(createSetParams.length, allocationManager.getOperatorSetCount(avs), "should be correct number of sets");
    }
}

contract AllocationManagerUnitTests_setAVSRegistrar is AllocationManagerUnitTests {
    function testFuzz_setAVSRegistrar_Correctness(
        Randomness r
    ) public rand(r) {
        address avs = r.Address();
        IAVSRegistrar avsRegistrar = IAVSRegistrar(r.Address());

        cheats.expectEmit(true, false, false, false, address(allocationManager));
        emit AVSRegistrarSet(avs, avsRegistrar);

        cheats.prank(avs);
        allocationManager.setAVSRegistrar(avs, avsRegistrar);
        assertTrue(avsRegistrar == allocationManager.getAVSRegistrar(avs), "should be set");
    }
}

contract AllocationManagerUnitTests_updateAVSMetadataURI is AllocationManagerUnitTests {
    function test_updateAVSMetadataURI_Correctness() public {
        string memory newURI = "test";
        cheats.expectEmit(true, false, false, false, address(allocationManager));
        emit AVSMetadataURIUpdated(defaultAVS, newURI);
        cheats.prank(defaultAVS);
        allocationManager.updateAVSMetadataURI(defaultAVS, newURI);
    }
}

contract AllocationManagerUnitTests_getStrategyAllocations is AllocationManagerUnitTests {
    function testFuzz_getStrategyAllocations_Correctness(
        Randomness r
    ) public rand(r) {
        AllocateParams[] memory allocateParams = r.AllocateParams(defaultAVS, 1, 1);
        CreateSetParams[] memory createSetParams = r.CreateSetParams(allocateParams);

        cheats.prank(defaultAVS);
        allocationManager.createOperatorSets(defaultAVS, createSetParams);

        cheats.startPrank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.stopPrank();

        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        (OperatorSet[] memory operatorSets, Allocation[] memory allocations) =
            allocationManager.getStrategyAllocations(defaultOperator, allocateParams[0].strategies[0]);

        assertEq(operatorSets[0].avs, allocateParams[0].operatorSet.avs, "should be defaultAVS");
        assertEq(operatorSets[0].id, allocateParams[0].operatorSet.id, "should be defaultOperatorSet");

        _checkAllocationStorage({
            allocation: allocations[0],
            expectedCurrentMagnitude: allocateParams[0].newMagnitudes[0],
            expectedPendingDiff: 0,
            expectedEffectBlock: 0
        });
    }
}