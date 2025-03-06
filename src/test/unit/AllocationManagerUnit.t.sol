// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/harnesses/AllocationManagerHarness.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/test/mocks/MockAVSRegistrar.sol";

contract AllocationManagerUnitTests is EigenLayerUnitTestSetup, IAllocationManagerErrors, IAllocationManagerEvents {
    using StdStyle for *;
    using ArrayLib for *;

    /// -----------------------------------------------------------------------
    /// Constants
    /// -----------------------------------------------------------------------

    /// NOTE: Raising these values directly increases cpu time for tests.
    uint internal constant FUZZ_MAX_ALLOCATIONS = 8;
    uint internal constant FUZZ_MAX_STRATS = 8;
    uint internal constant FUZZ_MAX_OP_SETS = 8;

    uint8 internal constant PAUSED_MODIFY_ALLOCATIONS = 0;
    uint8 internal constant PAUSED_OPERATOR_SLASHING = 1;
    uint8 internal constant PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION = 2;

    uint32 constant ASSUMED_BLOCK_TIME = 12 seconds;
    uint32 constant DEALLOCATION_DELAY = 14 days / ASSUMED_BLOCK_TIME;
    uint32 constant ALLOCATION_CONFIGURATION_DELAY = 17.5 days / ASSUMED_BLOCK_TIME;
    uint32 constant DEFAULT_OPERATOR_ALLOCATION_DELAY = 1 days / ASSUMED_BLOCK_TIME;

    uint constant DEFAULT_OPERATOR_SHARES = 1e18;

    /// -----------------------------------------------------------------------
    /// Mocks
    /// -----------------------------------------------------------------------

    AllocationManagerHarness allocationManager;
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
    /// Internal Storage Helpers
    /// -----------------------------------------------------------------------
    mapping(IStrategy => uint64) _encumberedMagnitudes;

    /// -----------------------------------------------------------------------
    /// Setup
    /// -----------------------------------------------------------------------

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();
        _initializeAllocationManager(address(this), pauserRegistry, 0);
        tokenMock = new ERC20PresetFixedSupply("Mock Token", "MOCK", type(uint).max, address(this));
        strategyMock = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(new StrategyBase(IStrategyManager(address(strategyManagerMock)), pauserRegistry, "v9.9.9")),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, tokenMock)
                )
            )
        );
        defaultStrategies = strategyMock.toArray();
        defaultOperatorSet = OperatorSet(defaultAVS, 0);

        cheats.prank(defaultAVS);
        allocationManager.updateAVSMetadataURI(defaultAVS, "https://example.com");

        _createOperatorSet(defaultOperatorSet, defaultStrategies);
        _registerOperator(defaultOperator);
        _setAllocationDelay(defaultOperator, DEFAULT_OPERATOR_ALLOCATION_DELAY);
        _registerForOperatorSet(defaultOperator, defaultOperatorSet);
        _grantDelegatedStake(defaultOperator, defaultOperatorSet, DEFAULT_OPERATOR_SHARES);
    }

    /// -----------------------------------------------------------------------
    /// Internal Helpers
    /// -----------------------------------------------------------------------

    function _initializeAllocationManager(address _initialOwner, IPauserRegistry _pauserRegistry, uint _initialPausedStatus)
        internal
        returns (AllocationManagerHarness)
    {
        return allocationManager = AllocationManagerHarness(
            address(
                new TransparentUpgradeableProxy(
                    address(
                        new AllocationManagerHarness(
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

    function _registerOperator(address operator) internal {
        delegationManagerMock.setIsOperator(operator, true);
    }

    function _setAllocationDelay(address operator, uint32 delay) internal {
        cheats.prank(operator);
        allocationManager.setAllocationDelay(operator, delay);
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY + 1);
    }

    function _createOperatorSet(OperatorSet memory operatorSet, IStrategy[] memory strategies) internal returns (OperatorSet memory) {
        cheats.prank(operatorSet.avs);
        allocationManager.createOperatorSets(
            operatorSet.avs, CreateSetParams({operatorSetId: operatorSet.id, strategies: strategies}).toArray()
        );
        return operatorSet;
    }

    function _createOperatorSets(OperatorSet[] memory operatorSets, IStrategy[] memory strategies) internal {
        CreateSetParams[] memory createSetParams = new CreateSetParams[](operatorSets.length);

        for (uint i; i < operatorSets.length; ++i) {
            createSetParams[i] = CreateSetParams({operatorSetId: operatorSets[i].id, strategies: strategies});
        }

        cheats.prank(operatorSets[0].avs);
        allocationManager.createOperatorSets(operatorSets[0].avs, createSetParams);
    }

    function _registerForOperatorSet(address operator, OperatorSet memory operatorSet) internal {
        cheats.prank(operator);
        allocationManager.registerForOperatorSets(
            operator, RegisterParams({avs: operatorSet.avs, operatorSetIds: operatorSet.id.toArrayU32(), data: ""})
        );
    }

    function _grantDelegatedStake(address operator, OperatorSet memory operatorSet, uint stake) internal {
        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);
        delegationManagerMock.setOperatorsShares(operator, strategies, stake);
    }

    function _registerForOperatorSets(address operator, OperatorSet[] memory operatorSets) internal {
        cheats.startPrank(operator);
        for (uint i; i < operatorSets.length; ++i) {
            allocationManager.registerForOperatorSets(
                operator, RegisterParams({avs: operatorSets[i].avs, operatorSetIds: operatorSets[i].id.toArrayU32(), data: ""})
            );
        }
        cheats.stopPrank();
    }

    struct Magnitudes {
        uint encumbered;
        uint max;
        uint allocatable;
    }

    /**
     * Get expected post slash storage values
     * Assumes that:
     * 1. WAD is max before slash
     * 2. encumbered is equal to magnitude before slash
     */
    function _getExpectedSlashVals(uint wadToSlash, uint64 magBeforeSlash)
        internal
        pure
        returns (uint wadSlashed, uint64 newCurrentMag, uint64 newMaxMag, uint64 newEncumberedMag)
    {
        return _getExpectedSlashVals(wadToSlash, magBeforeSlash, magBeforeSlash);
    }
    /**
     * Get expected post slash storage values
     * Assumes that:
     * 1. WAD is max before slash
     */

    function _getExpectedSlashVals(uint wadToSlash, uint64 magBeforeSlash, uint64 encumberedMagBeforeSlash)
        internal
        pure
        returns (uint wadSlashed, uint64 newCurrentMag, uint64 newMaxMag, uint64 newEncumberedMag)
    {
        // Get slippage to apply to returned values - we basically recreate mulWadRoundUp here
        uint64 slippage = _calculateSlippage(magBeforeSlash, wadToSlash);
        // Get the magnitude to slash - this value is rounded UP in the implementation
        uint64 slashedMag = uint64((uint(magBeforeSlash) * wadToSlash / WAD + slippage));
        wadSlashed = slashedMag;
        newCurrentMag = magBeforeSlash - slashedMag;
        newMaxMag = WAD - slashedMag;
        newEncumberedMag = encumberedMagBeforeSlash - slashedMag;
    }

    /// @dev Returns 0 or 1, depending on the remainder of the division
    function _calculateSlippage(uint64 magnitude, uint wadToSlash) internal pure returns (uint64) {
        return mulmod(magnitude, wadToSlash, WAD) > 0 ? 1 : 0;
    }

    function _checkAllocationStorage(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy strategy,
        Allocation memory expectedAllocation,
        Magnitudes memory expectedMagnitudes
    ) internal view {
        Allocation memory allocation = allocationManager.getAllocation(operator, operatorSet, strategy);

        console.log("\nChecking Allocation Storage:".yellow());
        console.log("   currentMagnitude: %d", allocation.currentMagnitude);
        console.log("   pendingDiff: %d", allocation.pendingDiff);
        console.log("   effectBlock: %d", allocation.effectBlock);

        assertEq(expectedAllocation.currentMagnitude, allocation.currentMagnitude, "currentMagnitude != expected");
        assertEq(expectedAllocation.pendingDiff, allocation.pendingDiff, "pendingDiff != expected");
        assertEq(expectedAllocation.effectBlock, allocation.effectBlock, "effectBlock != expected");

        uint encumberedMagnitude = allocationManager.getEncumberedMagnitude(operator, strategy);
        uint maxMagnitude = allocationManager.getMaxMagnitudes(operator, strategy.toArray())[0];
        uint allocatableMagnitude = allocationManager.getAllocatableMagnitude(operator, strategy);

        console.log("   encumberedMagnitude: %d", encumberedMagnitude);
        console.log("   maxMagnitude: %d", maxMagnitude);
        console.log("   allocatableMagnitude: %d", allocatableMagnitude);

        assertEq(expectedMagnitudes.encumbered, encumberedMagnitude, "encumberedMagnitude != expected");
        assertEq(expectedMagnitudes.max, maxMagnitude, "maxMagnitude != expected");
        assertEq(expectedMagnitudes.allocatable, allocatableMagnitude, "allocatableMagnitude != expected");

        // Check `getMaxMagnitudes` alias for coverage.
        assertEq(expectedMagnitudes.max, allocationManager.getMaxMagnitudes(operator.toArray(), strategy)[0], "maxMagnitude != expected");

        // Check `getAllocations` alias for coverage.
        Allocation memory getAllocations = allocationManager.getAllocations(operator.toArray(), operatorSet, strategy)[0];
        assertEq(expectedAllocation.currentMagnitude, getAllocations.currentMagnitude, "currentMagnitude != expected");
        assertEq(expectedAllocation.pendingDiff, getAllocations.pendingDiff, "pendingDiff != expected");
        assertEq(expectedAllocation.effectBlock, getAllocations.effectBlock, "effectBlock != expected");

        console.log("Success!".green().bold());
    }

    /// @dev Check that the deallocation queue is in ascending order of effectBlocks
    function _checkDeallocationQueueOrder(address operator, IStrategy strategy, uint numDeallocations) internal view {
        uint32 lastEffectBlock = 0;

        for (uint i = 0; i < numDeallocations; ++i) {
            bytes32 operatorSetKey = allocationManager.deallocationQueueAtIndex(operator, strategy, i);
            Allocation memory allocation = allocationManager.getAllocation(operator, OperatorSetLib.decode(operatorSetKey), strategy);

            assertTrue(lastEffectBlock <= allocation.effectBlock, "Deallocation queue is not in ascending order of effectBlocks");

            lastEffectBlock = allocation.effectBlock;
        }
    }

    function _checkSlashableStake(
        OperatorSet memory operatorSet,
        address operator,
        IStrategy[] memory strategies,
        uint expectedSlashableStake
    ) internal view {
        _checkSlashableStake(operatorSet, operator, strategies, expectedSlashableStake, block.number);
    }

    function _checkSlashableStake(
        OperatorSet memory operatorSet,
        address operator,
        IStrategy[] memory strategies,
        uint expectedSlashableStake,
        uint futureBlock
    ) internal view {
        uint[][] memory slashableStake = allocationManager.getMinimumSlashableStake({
            operatorSet: operatorSet,
            operators: operator.toArray(),
            strategies: strategies,
            futureBlock: uint32(futureBlock)
        });

        for (uint i = 0; i < strategies.length; i++) {
            console.log("\nChecking Slashable Stake:".yellow());
            console.log("   slashableStake[%d] = %d", i, slashableStake[0][i]);
            assertEq(slashableStake[0][i], expectedSlashableStake, "slashableStake != expected");
        }

        console.log("Success!".green().bold());
    }

    function _checkAllocationEvents(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy strategy,
        uint64 magnitude,
        uint64 encumberedMagnitude,
        uint32 effectBlock
    ) internal {
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(operator, strategy, encumberedMagnitude);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AllocationUpdated(operator, operatorSet, strategy, magnitude, effectBlock);
    }

    function _checkDeallocationEvent(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy strategy,
        uint64 magnitude,
        uint32 effectBlock
    ) internal {
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AllocationUpdated(operator, operatorSet, strategy, magnitude, effectBlock);
    }

    function _checkClearDeallocationQueueEvents(address operator, IStrategy strategy, uint64 encumberedMagnitude) internal {
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(operator, strategy, encumberedMagnitude);
    }

    function _checkSlashEvents(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy strategy,
        uint wadToSlash,
        string memory description,
        uint64 currentMag,
        uint64 maxMag,
        uint64 encumberedMag
    ) internal {
        return _checkSlashEvents(
            operator,
            operatorSet,
            strategy.toArray(),
            wadToSlash.toArrayU256(),
            description,
            currentMag.toArrayU64(),
            maxMag.toArrayU64(),
            encumberedMag.toArrayU64()
        );
    }

    function _checkSlashEvents(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        uint[] memory wadToSlash,
        string memory description,
        uint64[] memory currentMags,
        uint64[] memory maxMags,
        uint64[] memory encumberedMags
    ) internal {
        for (uint i = 0; i < strategies.length; i++) {
            // If there is nothing slashed, we don't emit events for encumbered magnitude
            if (wadToSlash[i] == 0) continue;
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit EncumberedMagnitudeUpdated(operator, strategies[i], encumberedMags[i]);
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit AllocationUpdated(operator, operatorSet, strategies[i], currentMags[i], uint32(block.number));
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit MaxMagnitudeUpdated(operator, strategies[i], maxMags[i]);
        }
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSlashed(operator, operatorSet, strategies, wadToSlash, description);
    }

    /// -----------------------------------------------------------------------
    /// Allocate/deallocate params
    /// -----------------------------------------------------------------------

    /// @dev Create allocate params, allocating `magnitude` to each strategy in the set
    function _newAllocateParams(OperatorSet memory operatorSet, uint64 magnitude) internal view returns (AllocateParams[] memory) {
        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);
        uint64[] memory newMagnitudes = new uint64[](strategies.length);

        for (uint i; i < strategies.length; ++i) {
            newMagnitudes[i] = magnitude;
        }

        return AllocateParams({operatorSet: operatorSet, strategies: strategies, newMagnitudes: newMagnitudes}).toArray();
    }

    /// @dev Create allocate params for multiple operator sets
    function _newAllocateParams(OperatorSet[] memory operatorSets, uint64 magnitude) internal view returns (AllocateParams[] memory) {
        AllocateParams[] memory allocateParams = new AllocateParams[](operatorSets.length);

        for (uint i; i < operatorSets.length; ++i) {
            allocateParams[i] = _newAllocateParams(operatorSets[i], magnitude)[0];
        }

        return allocateParams;
    }

    /// @dev Create random allocation params to the default operator set and strategy
    function _randAllocateParams_DefaultOpSet() internal returns (AllocateParams[] memory) {
        return _randAllocateParams_SingleMockStrategy(defaultOperatorSet.toArray());
    }

    /// @dev Create allocate params for random magnitudes to the same default strategy across multiple operator sets
    function _randAllocateParams_SingleMockStrategy(OperatorSet[] memory operatorSets) internal returns (AllocateParams[] memory) {
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
        for (uint i; i < params.length; ++i) {
            params[i] =
                AllocateParams({operatorSet: operatorSets[i], strategies: defaultStrategies, newMagnitudes: magnitudes[i].toArrayU64()});
        }

        return params;
    }

    /// @dev Create allocate params for random magnitudes to the same default strategy across multiple operator sets
    /// NOTE: this variant allocates ALL magnitude (1 WAD)
    function _randAllocateParams_SingleMockStrategy_AllocAll(OperatorSet[] memory operatorSets)
        internal
        returns (AllocateParams[] memory)
    {
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
            uint randIdx = random().Uint256(0, magnitudes.length - 1);
            magnitudes[randIdx] += magnitudeLeft;
            usedMagnitude += magnitudeLeft;
        }

        AllocateParams[] memory params = new AllocateParams[](magnitudes.length);
        for (uint i; i < params.length; ++i) {
            params[i] =
                AllocateParams({operatorSet: operatorSets[i], strategies: defaultStrategies, newMagnitudes: magnitudes[i].toArrayU64()});
        }

        return params;
    }

    /// @dev Create allocate/deallocate params to the same default strategy across multiple sets
    function _randAllocAndDeallocParams_SingleMockStrategy(OperatorSet[] memory operatorSets)
        internal
        returns (AllocateParams[] memory, AllocateParams[] memory)
    {
        AllocateParams[] memory allocateParams = _randAllocateParams_SingleMockStrategy(operatorSets);
        AllocateParams[] memory deallocateParams = new AllocateParams[](allocateParams.length);

        // Generate a random deallocation for each operator set
        for (uint i; i < deallocateParams.length; ++i) {
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

    function _defaultAllocEffectBlock() internal view returns (uint32) {
        return uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);
    }
}

contract AllocationManagerUnitTests_Initialization_Setters is AllocationManagerUnitTests {
    /// -----------------------------------------------------------------------
    /// initialize()
    /// -----------------------------------------------------------------------

    /// @dev Asserts the following:
    /// 1. The fn can only be called once, during deployment.
    /// 2. The fn initializes the contract state correctly (owner, pauserRegistry, and initialPausedStatus).
    function testFuzz_Initialize(Randomness r) public rand(r) {
        // Generate random values for the expected initial state of the contract.
        address expectedInitialOwner = r.Address();
        IPauserRegistry expectedPauserRegistry = IPauserRegistry(r.Address());

        // Deploy the contract with the expected initial state.
        uint initialPausedStatus = r.Uint256();
        AllocationManager alm = _initializeAllocationManager(expectedInitialOwner, expectedPauserRegistry, initialPausedStatus);

        // Assert that the contract can only be initialized once.
        vm.expectRevert("Initializable: contract is already initialized");
        alm.initialize(expectedInitialOwner, initialPausedStatus);

        // Assert immutable state
        assertEq(address(alm.delegation()), address(delegationManagerMock));
        assertEq(alm.DEALLOCATION_DELAY(), DEALLOCATION_DELAY);
        assertEq(alm.ALLOCATION_CONFIGURATION_DELAY(), ALLOCATION_CONFIGURATION_DELAY);

        // Assert initialization state
        assertEq(alm.owner(), expectedInitialOwner);
        assertEq(alm.paused(), initialPausedStatus);
    }
}

contract AllocationManagerUnitTests_SlashOperator is AllocationManagerUnitTests {
    using ArrayLib for *;
    using SlashingLib for *;

    /// -----------------------------------------------------------------------
    /// slashOperator()
    /// -----------------------------------------------------------------------

    function _randSlashingParams(address operator, uint32 operatorSetId) internal returns (SlashingParams memory) {
        return SlashingParams({
            operator: operator,
            operatorSetId: operatorSetId,
            strategies: defaultStrategies,
            wadsToSlash: random().Uint256(1, WAD).toArrayU256(),
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
        slashingParams.wadsToSlash[0] = 0;

        cheats.prank(defaultAVS);
        cheats.expectRevert(InvalidWadToSlash.selector);
        allocationManager.slashOperator(defaultAVS, slashingParams);
    }

    function test_revert_slashGreaterThanWAD() public {
        SlashingParams memory slashingParams = _randSlashingParams(defaultOperator, 0);
        slashingParams.wadsToSlash[0] = WAD + 1;

        cheats.prank(defaultAVS);
        cheats.expectRevert(InvalidWadToSlash.selector);
        allocationManager.slashOperator(defaultAVS, slashingParams);
    }

    function test_revert_NotRegisteredToSet() public {
        cheats.prank(defaultAVS);
        cheats.expectRevert(OperatorNotSlashable.selector);
        allocationManager.slashOperator(defaultAVS, _randSlashingParams(random().Address(), 0));
    }

    function test_revert_NotMemberOfSet() public {
        cheats.prank(defaultAVS);
        cheats.expectRevert(OperatorNotSlashable.selector);
        allocationManager.slashOperator(defaultAVS, _randSlashingParams(random().Address(), 0));
    }

    function test_revert_InputArrayLengthMismatch() public {
        SlashingParams memory slashingParams = _randSlashingParams(defaultOperator, 0);
        slashingParams.strategies = slashingParams.strategies.setLength(2);

        cheats.prank(defaultAVS);
        cheats.expectRevert(InputArrayLengthMismatch.selector);
        allocationManager.slashOperator(defaultAVS, slashingParams);
    }

    function test_revert_StrategiesMustBeInAscendingOrder() public {
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = IStrategy(address(3));
        strategies[1] = IStrategy(address(2));
        strategies[2] = IStrategy(address(1));

        _createOperatorSet(OperatorSet(defaultAVS, 1), strategies);
        _registerForOperatorSet(defaultOperator, OperatorSet(defaultAVS, 1));

        cheats.prank(defaultAVS);
        cheats.expectRevert(StrategiesMustBeInAscendingOrder.selector);
        allocationManager.slashOperator(
            defaultAVS,
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: 1,
                strategies: strategies,
                wadsToSlash: uint(0.5 ether).toArrayU256(3),
                description: "test"
            })
        );
    }

    function test_revert_StrategyNotInOperatorSet() public {
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = IStrategy(address(1));
        strategies[1] = IStrategy(address(2));
        strategies[2] = IStrategy(address(3));

        _createOperatorSet(OperatorSet(defaultAVS, 1), strategies);
        _registerForOperatorSet(defaultOperator, OperatorSet(defaultAVS, 1));

        strategies = strategies.setLength(4);
        strategies[3] = IStrategy(address(4));

        cheats.prank(defaultAVS);
        cheats.expectRevert(StrategyNotInOperatorSet.selector);
        allocationManager.slashOperator(
            defaultAVS,
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: 1,
                strategies: defaultStrategies,
                wadsToSlash: uint(0.5 ether).toArrayU256(),
                description: "test"
            })
        );
    }

    /**
     * Attempts to slash an operator before the allocation is active
     * Validates:
     * 1. The events of the slash indicate nothing was slashed
     * 2. Storage is not mutated post slash
     * 3. The operator's allocation takes effect as normal post slash
     */
    function test_operatorAllocated_notActive() public {
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, WAD);

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        uint64 encumberedMagnitudeBefore = allocationManager.getEncumberedMagnitude(defaultOperator, strategyMock);
        uint64 maxMagnitudeBefore = allocationManager.getMaxMagnitudes(defaultOperator, strategyMock.toArray())[0];

        // The only slash event we expect is the OperatorSlashed. Validate the number
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSlashed(defaultOperator, defaultOperatorSet, defaultStrategies, uint(0).toArrayU256(), "test");

        uint numLogsBefore = cheats.getRecordedLogs().length;
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(
            defaultAVS,
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: allocateParams[0].operatorSet.id,
                strategies: defaultStrategies,
                wadsToSlash: WAD.toArrayU256(),
                description: "test"
            })
        );
        uint numLogsAfter = cheats.getRecordedLogs().length;

        // Assert only 1 log was emitted
        assertEq(numLogsAfter, numLogsBefore + 1, "Incorrect number of logs emitted");

        // Assert encumberedMagnitude and maxMagnitude are unchanged
        assertEq(
            encumberedMagnitudeBefore,
            allocationManager.getEncumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude mutated"
        );

        assertEq(maxMagnitudeBefore, allocationManager.getMaxMagnitudes(defaultOperator, strategyMock.toArray())[0], "maxMagnitude mutated");

        // Roll to effect block and validate allocation
        uint32 effectBlock = uint32(block.number) + DEFAULT_OPERATOR_ALLOCATION_DELAY;
        uint64 pendingIncrease = allocateParams[0].newMagnitudes[0];
        cheats.roll(effectBlock);

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: pendingIncrease, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({encumbered: pendingIncrease, max: WAD, allocatable: WAD - pendingIncrease})
        });
    }

    /**
     * Allocates all magnitude to for a single strategy to an operatorSet. Slashes 25%
     * Validates:
     * 1. Events are emitted
     * 2. Allocation & info introspection
     * 3. Slashable stake introspection
     */
    function test_slashPostAllocation() public {
        // Generate allocation for this operator set, we allocate max
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, WAD);

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            wadToSlash: uint(25e16),
            description: "test",
            currentMag: uint64(75e16),
            maxMag: uint64(75e16),
            encumberedMag: uint64(75e16)
        });

        // Slash operator for 25%
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(
            defaultAVS,
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: defaultOperatorSet.id,
                strategies: defaultStrategies,
                wadsToSlash: 25e16.toArrayU256(),
                description: "test"
            })
        );

        // Check storage
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: 75e16, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({encumbered: 75e16, max: 75e16, allocatable: 0})
        });

        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: DEFAULT_OPERATOR_SHARES.mulWad(75e16)
        });
    }

    /// @notice Same test as above, but fuzzes the allocation
    function testFuzz_slashPostAllocation(Randomness r) public rand(r) {
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();

        // Allocate magnitude and roll forward to completable block
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        SlashingParams memory slashingParams = _randSlashingParams(defaultOperator, defaultOperatorSet.id);

        (uint expectedWadSlashed, uint64 expectedCurrentMag, uint64 expectedMaxMag, uint64 expectedEncumberedMag) =
            _getExpectedSlashVals({magBeforeSlash: allocateParams[0].newMagnitudes[0], wadToSlash: slashingParams.wadsToSlash[0]});

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            wadToSlash: expectedWadSlashed,
            description: "test",
            currentMag: expectedCurrentMag,
            maxMag: expectedMaxMag,
            encumberedMag: expectedEncumberedMag
        });

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: expectedCurrentMag, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({
                encumbered: expectedEncumberedMag,
                max: expectedMaxMag,
                allocatable: expectedMaxMag - expectedEncumberedMag
            })
        });

        uint slashedStake = DEFAULT_OPERATOR_SHARES.mulWad(expectedWadSlashed); // Wad is same as slashed mag since we start with max mag
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: (DEFAULT_OPERATOR_SHARES - slashedStake).mulWad(expectedCurrentMag.divWad(expectedMaxMag))
        });
    }

    /**
     * Allocates half of magnitude for a single strategy to an operatorSet. Then allocates again. Slashes 50%
     * Asserts that:
     * 1. Events are emitted
     * 2. Encumbered mag is updated
     * 3. Max mag is updated
     * 4. Calculations for `getAllocatableMagnitude` and `getAllocation` are correct
     * 5. The second allocation is not slashed from
     */
    function testFuzz_slash_oneCompletedAlloc_onePendingAlloc(Randomness r) public rand(r) {
        uint wadToSlash = r.Uint256(0.01 ether, WAD);

        // Generate allocation for `strategyMock`, we allocate half
        {
            AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, 5e17);
            cheats.prank(defaultOperator);
            allocationManager.modifyAllocations(defaultOperator, allocateParams);
            cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);
        }

        // Allocate the other half
        AllocateParams[] memory allocateParams2 = _newAllocateParams(defaultOperatorSet, WAD);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams2);
        uint32 secondAllocEffectBlock = _defaultAllocEffectBlock();

        // Slash operator for 50%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: defaultOperatorSet.id,
            strategies: defaultStrategies,
            wadsToSlash: wadToSlash.toArrayU256(),
            description: "test"
        });

        (uint expectedWadSlashed, uint64 expectedCurrentMag, uint64 expectedMaxMag, uint64 expectedEncumberedMag) =
            _getExpectedSlashVals({magBeforeSlash: 5e17, encumberedMagBeforeSlash: WAD, wadToSlash: wadToSlash});

        uint slashedStake = DEFAULT_OPERATOR_SHARES.mulWad(expectedWadSlashed); // Wad is same as slashed mag since we start with max mag
        uint newTotalStake = DEFAULT_OPERATOR_SHARES - slashedStake;

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            wadToSlash: expectedWadSlashed,
            description: "test",
            currentMag: expectedCurrentMag,
            maxMag: expectedMaxMag,
            encumberedMag: expectedEncumberedMag
        });

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: expectedCurrentMag, pendingDiff: 5e17, effectBlock: secondAllocEffectBlock}),
            expectedMagnitudes: Magnitudes({encumbered: expectedEncumberedMag, max: expectedMaxMag, allocatable: 0})
        });

        // Slashable stake should include first allocation and slashed magnitude
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: newTotalStake.mulWad(expectedCurrentMag.divWad(expectedMaxMag))
        });

        cheats.roll(secondAllocEffectBlock);

        uint64 newSlashableMagnitude = expectedCurrentMag + 5e17;
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: newSlashableMagnitude, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({encumbered: expectedEncumberedMag, max: expectedMaxMag, allocatable: 0})
        });

        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: newTotalStake.mulWad(newSlashableMagnitude.divWad(expectedMaxMag))
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
     * 2. Storage properly updated after each slash
     * 3. Slashed amounts are rounded up to ensure magnitude is always slashed
     */
    function test_repeatUntilFullSlash() public {
        // Generate allocation for `strategyMock`, we allocate 100% to opSet 0
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, WAD);

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // 1. Slash operator for 99% in opSet 0 bringing their magnitude to 1e16
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: defaultOperatorSet.id,
            strategies: defaultStrategies,
            wadsToSlash: 99e16.toArrayU256(),
            description: "test"
        });

        uint64 expectedEncumberedMagnitude = 1e16; // After slashing 99%, only 1% expected encumberedMagnitude
        uint64 magnitudeAfterSlash = 1e16;
        uint64 maxMagnitudeAfterSlash = 1e16; // 1e15 is maxMagnitude

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            wadToSlash: uint(99e16),
            description: "test",
            currentMag: magnitudeAfterSlash,
            maxMag: maxMagnitudeAfterSlash,
            encumberedMag: expectedEncumberedMagnitude
        });

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Check storage
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: magnitudeAfterSlash, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({
                encumbered: expectedEncumberedMagnitude,
                max: maxMagnitudeAfterSlash,
                allocatable: maxMagnitudeAfterSlash - expectedEncumberedMagnitude
            })
        });

        // 2. Slash operator again for 99.99% in opSet 0 bringing their magnitude to 1e14
        slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: defaultOperatorSet.id,
            strategies: defaultStrategies,
            wadsToSlash: 9999e14.toArrayU256(),
            description: "test"
        });

        expectedEncumberedMagnitude = 1e12; // After slashing 99.99%, only 0.01% expected encumberedMagnitude
        magnitudeAfterSlash = 1e12;
        maxMagnitudeAfterSlash = 1e12;

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            wadToSlash: uint(9999e14),
            description: "test",
            currentMag: magnitudeAfterSlash,
            maxMag: maxMagnitudeAfterSlash,
            encumberedMag: expectedEncumberedMagnitude
        });

        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Check storage
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: magnitudeAfterSlash, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({
                encumbered: expectedEncumberedMagnitude,
                max: maxMagnitudeAfterSlash,
                allocatable: maxMagnitudeAfterSlash - expectedEncumberedMagnitude
            })
        });

        // 3. Slash operator again for 99.9999999999999% in opSet 0
        slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: defaultOperatorSet.id,
            strategies: defaultStrategies,
            wadsToSlash: uint(WAD - 1e3).toArrayU256(),
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
            strategy: strategyMock,
            wadToSlash: WAD,
            description: "test",
            currentMag: magnitudeAfterSlash,
            maxMag: maxMagnitudeAfterSlash,
            encumberedMag: expectedEncumberedMagnitude
        });

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Check storage
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: 0, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({encumbered: 0, max: 0, allocatable: 0})
        });

        // Check slashable amount after final slash
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: 0
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
    function testFuzz_SlashWhileDeallocationPending(Randomness r) public rand(r) {
        // Initialize state
        AllocateParams[] memory allocateParams = r.AllocateParams(defaultAVS, 1, 1);
        AllocateParams[] memory deallocateParams = r.DeallocateParams(allocateParams);
        CreateSetParams[] memory createSetParams = r.CreateSetParams(allocateParams);
        RegisterParams memory registerParams = r.RegisterParams(allocateParams);
        SlashingParams memory slashingParams = r.SlashingParams(defaultOperator, allocateParams[0]);

        delegationManagerMock.setOperatorShares(defaultOperator, allocateParams[0].strategies[0], DEFAULT_OPERATOR_SHARES);

        cheats.prank(defaultAVS);
        allocationManager.createOperatorSets(defaultAVS, createSetParams);
        cheats.startPrank(defaultOperator);
        allocationManager.registerForOperatorSets(defaultOperator, registerParams);

        // Allocate
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY + 1);
        cheats.stopPrank();

        uint magnitudeAllocated = allocateParams[0].newMagnitudes[0];
        uint magnitudeDeallocated = magnitudeAllocated - deallocateParams[0].newMagnitudes[0];
        uint magnitudeSlashed = (magnitudeAllocated * slashingParams.wadsToSlash[0] / WAD)
            + _calculateSlippage(uint64(magnitudeAllocated), slashingParams.wadsToSlash[0]);
        uint expectedCurrentMagnitude = magnitudeAllocated - magnitudeSlashed;
        int128 expectedPendingDiff =
            -int128(uint128(magnitudeDeallocated - magnitudeDeallocated.mulWadRoundUp(slashingParams.wadsToSlash[0])));

        // Manually check slash events since we have a deallocation pending
        // Deallocation update is emitted first
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AllocationUpdated(
            defaultOperator,
            allocateParams[0].operatorSet,
            allocateParams[0].strategies[0],
            uint64(uint128(int128(uint128(expectedCurrentMagnitude)) + expectedPendingDiff)),
            deallocationEffectBlock
        );
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, allocateParams[0].strategies[0], uint64(magnitudeAllocated - magnitudeSlashed));
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AllocationUpdated(
            defaultOperator,
            allocateParams[0].operatorSet,
            allocateParams[0].strategies[0],
            uint64(expectedCurrentMagnitude),
            uint32(block.number)
        );
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit MaxMagnitudeUpdated(defaultOperator, allocateParams[0].strategies[0], uint64(WAD - magnitudeSlashed));
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSlashed(
            defaultOperator, allocateParams[0].operatorSet, allocateParams[0].strategies, magnitudeSlashed.toArrayU256(), ""
        );

        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: allocateParams[0].operatorSet,
            strategy: allocateParams[0].strategies[0],
            expectedAllocation: Allocation({
                currentMagnitude: uint64(expectedCurrentMagnitude),
                pendingDiff: expectedPendingDiff,
                effectBlock: deallocationEffectBlock
            }),
            expectedMagnitudes: Magnitudes({encumbered: expectedCurrentMagnitude, max: uint64(WAD - magnitudeSlashed), allocatable: 0})
        });

        cheats.roll(deallocationEffectBlock);
        allocationManager.clearDeallocationQueue(defaultOperator, allocateParams[0].strategies, _maxNumToClear());

        uint64 newMag = uint64(uint128(int128(uint128(expectedCurrentMagnitude)) + expectedPendingDiff));

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: allocateParams[0].operatorSet,
            strategy: allocateParams[0].strategies[0],
            expectedAllocation: Allocation({currentMagnitude: newMag, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({
                encumbered: newMag,
                max: uint64(WAD - magnitudeSlashed),
                allocatable: uint128(-expectedPendingDiff) // This works because we allocated all in the randomization allocation helper
            })
        });
    }

    /**
     * Allocates all magnitude to a single opSet. Then slashes the entire magnitude
     * Validates:
     * 1. Storage post slash
     * 2. The operator cannot allocate again
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
            strategy: strategyMock,
            wadToSlash: WAD,
            description: "test",
            currentMag: 0,
            maxMag: 0,
            encumberedMag: 0
        });

        // Slash operator for 100%
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(
            defaultAVS,
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: allocateParams[0].operatorSet.id,
                strategies: defaultStrategies,
                wadsToSlash: WAD.toArrayU256(),
                description: "test"
            })
        );

        // Validate storage post slash
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: 0, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({encumbered: 0, max: 0, allocatable: 0})
        });

        OperatorSet memory operatorSet = _createOperatorSet(OperatorSet(defaultAVS, random().Uint32()), defaultStrategies);
        AllocateParams[] memory allocateParams2 = _newAllocateParams(operatorSet, 1);

        // Attempt to allocate
        cheats.prank(defaultOperator);
        cheats.expectRevert(InsufficientMagnitude.selector);
        allocationManager.modifyAllocations(defaultOperator, allocateParams2);
    }

    /**
     * Allocates all magnitude to a single opSet. Deallocates magnitude. Slashes all
     * Asserts that:
     * 1. The Allocation is 0 after slash
     * 2. Them storage post slash for encumbered and maxMags is zero
     */
    function test_slash_allocateAll_deallocateAll() public {
        // Allocate all magnitude
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(defaultOperatorSet, WAD));
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate all
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(defaultOperatorSet, 0));

        // Validate event for the deallocation
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AllocationUpdated(defaultOperator, defaultOperatorSet, strategyMock, 0, uint32(block.number + DEALLOCATION_DELAY + 1));

        // Slash operator for 100%
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(
            defaultAVS,
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: defaultOperatorSet.id,
                strategies: defaultStrategies,
                wadsToSlash: WAD.toArrayU256(),
                description: "test"
            })
        );

        // forgefmt: disable-next-item
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: 0, pendingDiff: 0, effectBlock: uint32(block.number) + DEALLOCATION_DELAY + 1}),
            expectedMagnitudes: Magnitudes({encumbered: 0, max: 0, allocatable: 0})
        });

        // Complete deallocation
        cheats.roll(uint32(block.number) + DEALLOCATION_DELAY + 1);
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());

        // Validate allocatable amount is 0
        assertEq(0, allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock), "Allocatable magnitude should be 0");
    }

    /**
     * Slashes the operator after deallocation, even if the deallocation has not been cleared.
     * Validates that:
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
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY + 1);

        // Warp to deallocation effect block
        cheats.roll(deallocationEffectBlock);

        // Slash operator for 25%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: defaultOperatorSet.id,
            strategies: defaultStrategies,
            wadsToSlash: 25e16.toArrayU256(),
            description: "test"
        });

        uint64 expectedEncumberedMagnitude = 375e15; // 25e16 is slashed. 5e17 was previously
        uint64 magnitudeAfterSlash = 375e15;
        uint64 maxMagnitudeAfterSlash = 875e15; // Operator can only allocate up to 75e16 magnitude since 25% is slashed
        uint expectedSlashedMagnitude = SlashingLib.mulWadRoundUp(5e17, 25e16);

        // Slash Operator, only emit events assuming that there is no deallocation
        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            wadToSlash: expectedSlashedMagnitude,
            description: "test",
            currentMag: magnitudeAfterSlash,
            maxMag: maxMagnitudeAfterSlash,
            encumberedMag: expectedEncumberedMagnitude
        });

        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());

        uint64 allocatableMagnitudeAfterSlash = allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock);

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: magnitudeAfterSlash, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({
                encumbered: expectedEncumberedMagnitude,
                max: maxMagnitudeAfterSlash,
                allocatable: allocatableMagnitudeAfterSlash
            })
        });
    }

    /**
     * Allocates to multiple operatorSets for a strategy. Only slashes from one operatorSet.
     * Validates:
     * 1. The first operatorSet has less slashable shares post slash
     * 2. The second operatorSet has the same number slashable shares post slash (within slippage)
     * 3. The PROPORTION that is slashable for opSet 2 has increased
     */
    function testFuzz_allocateMultipleOpsets_slashSingleOpset(Randomness r) public rand(r) {
        // Get magnitude to allocate
        uint64 magnitudeToAllocate = r.Uint64(1, 5e17);
        uint wadToSlash = r.Uint256(1, 1e18);

        OperatorSet memory operatorSet = OperatorSet(defaultAVS, 1);
        OperatorSet memory operatorSet2 = OperatorSet(defaultAVS, 2);

        // Allocate 40% to firstOperatorSet, 40% to secondOperatorSet
        AllocateParams[] memory allocateParams = new AllocateParams[](2);
        allocateParams[0] = _newAllocateParams(_createOperatorSet(OperatorSet(defaultAVS, 1), defaultStrategies), magnitudeToAllocate)[0];
        allocateParams[1] = _newAllocateParams(_createOperatorSet(OperatorSet(defaultAVS, 2), defaultStrategies), magnitudeToAllocate)[0];

        // Register operator for both operatorSets
        _registerForOperatorSet(defaultOperator, operatorSet);
        _registerForOperatorSet(defaultOperator, operatorSet2);

        // Modify allocations
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Get slashable shares for each operatorSet
        uint[][] memory opset2SlashableSharesBefore =
            allocationManager.getMinimumSlashableStake(operatorSet2, defaultOperator.toArray(), defaultStrategies, uint32(block.number));
        // Slash operator on operatorSet1 for 50%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocateParams[0].operatorSet.id,
            strategies: defaultStrategies,
            wadsToSlash: wadToSlash.toArrayU256(),
            description: "test"
        });

        (, uint64 expectedCurrentMag, uint64 expectedMaxMag, uint64 expectedEncumberedMag) = _getExpectedSlashVals({
            magBeforeSlash: allocateParams[0].newMagnitudes[0],
            wadToSlash: slashingParams.wadsToSlash[0],
            encumberedMagBeforeSlash: allocateParams[0].newMagnitudes[0] * 2
        });

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Validate storage operatorSet1
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: operatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: expectedCurrentMag, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({
                encumbered: expectedEncumberedMag,
                max: expectedMaxMag,
                allocatable: expectedMaxMag - expectedEncumberedMag
            })
        });

        // Validate storage for operatorSet2
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: operatorSet2,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: magnitudeToAllocate, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({
                encumbered: expectedEncumberedMag,
                max: expectedMaxMag,
                allocatable: expectedMaxMag - expectedEncumberedMag
            })
        });

        // Check proportion after slash
        uint opSet2PortionOfMaxMagnitudeAfterSlash = uint(magnitudeToAllocate) * WAD / expectedMaxMag;
        assertGt(
            opSet2PortionOfMaxMagnitudeAfterSlash,
            magnitudeToAllocate, // This is the same as proportion before slash
            "opSet2 should have a greater proportion to slash from previous"
        );

        // Assert that slashable stake is the same - we add slippage here due to rounding error from the slash itself
        uint[][] memory minSlashableStake =
            allocationManager.getMinimumSlashableStake(operatorSet2, defaultOperator.toArray(), defaultStrategies, uint32(block.number));
        assertEq(opset2SlashableSharesBefore[0][0], minSlashableStake[0][0] + 1, "opSet2 slashable shares should be the same");
    }

    /**
     * Allocates to multiple strategies for the given operatorSetKey. Slashes from both strategies Validates a slash propagates to both strategies.
     * Validates that
     * 1. Proper events are emitted for each strategy slashed
     * 2. Each strategy is slashed proportional to its allocation
     * 3. Storage is updated for each strategy, opSet
     */
    function testFuzz_allocateMultipleStrategies_slashMultiple(Randomness r) public rand(r) {
        // Initialize random params
        uint64 strategy1Magnitude = r.Uint64(1, 1e18);
        uint64 strategy2Magnitude = r.Uint64(1, 1e18);
        uint wadToSlash = r.Uint256(1, 1e18);

        // Crate and allocate to operatorSets
        OperatorSet memory operatorSet = OperatorSet(defaultAVS, random().Uint32());
        _createOperatorSet(operatorSet, random().StrategyArray(2));
        _registerForOperatorSet(defaultOperator, operatorSet);

        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);
        uint[] memory wadsToSlash = new uint[](strategies.length);
        {
            if (strategies[1] < strategies[0]) {
                IStrategy temp = strategies[0];
                strategies[0] = strategies[1];
                strategies[1] = temp;
            }
            AllocateParams memory allocateParams =
                AllocateParams({operatorSet: operatorSet, strategies: strategies, newMagnitudes: new uint64[](2)});
            allocateParams.newMagnitudes[0] = strategy1Magnitude;
            allocateParams.newMagnitudes[1] = strategy2Magnitude;

            cheats.prank(defaultOperator);
            allocationManager.modifyAllocations(defaultOperator, allocateParams.toArray());
            cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

            for (uint i = 0; i < strategies.length; i++) {
                wadsToSlash[i] = wadToSlash;
            }
        }

        // Store post-slash vars to check against
        uint64[] memory expectedEncumberedMags = new uint64[](2);
        uint[] memory expectedSlashedMagnitude = new uint[](2);
        uint64[] memory expectedMagnitudeAfterSlash = new uint64[](2);
        uint64[] memory expectedMaxMagnitudeAfterSlash = new uint64[](2);

        {
            (
                uint strat1ExpectedWadSlashed,
                uint64 strat1ExpectedCurrentMag,
                uint64 strat1ExpectedMaxMag,
                uint64 strat1ExpectedEncumberedMag
            ) = _getExpectedSlashVals({magBeforeSlash: strategy1Magnitude, wadToSlash: wadToSlash});
            expectedEncumberedMags[0] = strat1ExpectedEncumberedMag;
            expectedSlashedMagnitude[0] = strat1ExpectedWadSlashed;
            expectedMagnitudeAfterSlash[0] = strat1ExpectedCurrentMag;
            expectedMaxMagnitudeAfterSlash[0] = strat1ExpectedMaxMag;
        }
        {
            (
                uint strat2ExpectedWadSlashed,
                uint64 strat2ExpectedCurrentMag,
                uint64 strat2ExpectedMaxMag,
                uint64 strat2ExpectedEncumberedMag
            ) = _getExpectedSlashVals({magBeforeSlash: strategy2Magnitude, wadToSlash: wadToSlash});
            expectedEncumberedMags[1] = strat2ExpectedEncumberedMag;
            expectedSlashedMagnitude[1] = strat2ExpectedWadSlashed;
            expectedMagnitudeAfterSlash[1] = strat2ExpectedCurrentMag;
            expectedMaxMagnitudeAfterSlash[1] = strat2ExpectedMaxMag;
        }

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: operatorSet,
            strategies: strategies,
            wadToSlash: expectedSlashedMagnitude,
            description: "test",
            currentMags: expectedMagnitudeAfterSlash,
            maxMags: expectedMaxMagnitudeAfterSlash,
            encumberedMags: expectedEncumberedMags
        });

        // Slash Operator
        {
            SlashingParams memory slashingParams = SlashingParams({
                operator: defaultOperator,
                operatorSetId: operatorSet.id,
                strategies: strategies,
                wadsToSlash: wadsToSlash,
                description: "test"
            });
            cheats.prank(defaultAVS);
            allocationManager.slashOperator(defaultAVS, slashingParams);
        }

        // Check storage
        for (uint i; i < strategies.length; ++i) {
            _checkAllocationStorage({
                operator: defaultOperator,
                operatorSet: operatorSet,
                strategy: strategies[i],
                expectedAllocation: Allocation({currentMagnitude: expectedMagnitudeAfterSlash[i], pendingDiff: 0, effectBlock: 0}),
                expectedMagnitudes: Magnitudes({
                    encumbered: expectedEncumberedMags[i],
                    max: expectedMaxMagnitudeAfterSlash[i],
                    allocatable: expectedMaxMagnitudeAfterSlash[i] - expectedEncumberedMags[i]
                })
            });
        }
    }

    /// @dev Allocates magnitude. Deallocates some. Slashes a portion, and then allocates up to the max available magnitude
    function testFuzz_allocate_deallocate_allocateMax(Randomness r) public rand(r) {
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
        cheats.roll(block.number + DEALLOCATION_DELAY + 1);
        cheats.stopPrank();

        // Slash operator for some random amount (1% -> 99%).
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: operatorSet.id,
            strategies: strategy.toArray(),
            wadsToSlash: (r.Uint256(0.01 ether, 0.99 ether)).toArrayU256(),
            description: "test"
        });

        (uint expectedWadSlashed, uint64 expectedCurrentMag, uint64 expectedMaxMag, uint64 expectedEncumberedMag) =
            _getExpectedSlashVals({magBeforeSlash: deallocateParams[0].newMagnitudes[0], wadToSlash: slashingParams.wadsToSlash[0]});

        _checkSlashEvents({
            operator: defaultOperator,
            operatorSet: operatorSet,
            strategy: allocateParams[0].strategies[0],
            wadToSlash: expectedWadSlashed,
            description: "test",
            currentMag: expectedCurrentMag,
            maxMag: expectedMaxMag,
            encumberedMag: expectedEncumberedMag
        });

        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Clear deallocation queue.
        allocationManager.clearDeallocationQueue(defaultOperator, strategy.toArray(), _maxNumToClear());

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: operatorSet,
            strategy: strategy,
            expectedAllocation: Allocation({currentMagnitude: expectedCurrentMag, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({
                encumbered: expectedEncumberedMag,
                max: expectedMaxMag,
                allocatable: expectedMaxMag - expectedEncumberedMag
            })
        });

        // Allocate up to max magnitude
        AllocateParams[] memory allocateParams2 = _newAllocateParams(operatorSet, expectedMaxMag);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams2);

        int128 pendingDiff = int128(uint128(expectedMaxMag - expectedCurrentMag));

        // Check storage
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: operatorSet,
            strategy: strategy,
            expectedAllocation: Allocation({
                currentMagnitude: expectedCurrentMag,
                pendingDiff: pendingDiff,
                effectBlock: _defaultAllocEffectBlock()
            }),
            expectedMagnitudes: Magnitudes({encumbered: expectedMaxMag, max: expectedMaxMag, allocatable: 0})
        });
    }

    /**
     * Allocates magnitude to an operator, deallocates all, warps to deallocation effect block and attempts to slash
     * Asserts that the operator is not slashed
     */
    function test_noFundsSlashedAfterDeallocationDelay() public {
        // Allocate all magnitude
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(defaultOperatorSet, WAD));
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate all
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(defaultOperatorSet, 0));

        // Warp to deallocation effect block
        cheats.roll(block.number + DEALLOCATION_DELAY + 1);

        // Slash operator for all wad
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(
            defaultAVS,
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: defaultOperatorSet.id,
                strategies: defaultStrategies,
                wadsToSlash: WAD.toArrayU256(),
                description: "test"
            })
        );

        // Assert that the operator's max magnitude and allocatable magnitude are still WAD
        assertEq(WAD, allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock), "Allocatable magnitude should be WAD");
        assertEq(WAD, allocationManager.getMaxMagnitude(defaultOperator, strategyMock), "Max magnitude should be WAD");
    }

    function testRevert_noFundsSlashedAfterDeregistration() public {
        // Allocate all magnitude
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(defaultOperatorSet, WAD));
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deregister operator
        DeregisterParams memory deregisterParams =
            DeregisterParams({operator: defaultOperator, avs: defaultAVS, operatorSetIds: defaultOperatorSet.id.toArrayU32()});
        cheats.prank(defaultOperator);
        allocationManager.deregisterFromOperatorSets(deregisterParams);

        // Warp to deallocation effect block
        cheats.roll(block.number + DEALLOCATION_DELAY + 1);

        // Slash operator for all wad
        cheats.expectRevert(OperatorNotSlashable.selector);
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(
            defaultAVS,
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: defaultOperatorSet.id,
                strategies: defaultStrategies,
                wadsToSlash: WAD.toArrayU256(),
                description: "test"
            })
        );
    }

    function test_deallocationSlashedJustBeforeEffectBlock() public {
        // Allocate all magnitude
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(defaultOperatorSet, WAD));
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate all
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(defaultOperatorSet, 0));

        // Warp to just before deallocation effect block
        cheats.roll(block.number + DEALLOCATION_DELAY);

        // Slash operator for all wad
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(
            defaultAVS,
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: defaultOperatorSet.id,
                strategies: defaultStrategies,
                wadsToSlash: WAD.toArrayU256(),
                description: "test"
            })
        );

        // Assert that the operator has no max magnitude or allocatable magnitude
        assertEq(0, allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock), "Allocatable magnitude should be 0");
        assertEq(0, allocationManager.getMaxMagnitude(defaultOperator, strategyMock), "Max magnitude should be 0");
    }

    function test_deregisteredOperatorSlashableBeforeDelay() public {
        // Allocate all magnitude
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(defaultOperatorSet, WAD));
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deregister operator
        DeregisterParams memory deregisterParams =
            DeregisterParams({operator: defaultOperator, avs: defaultAVS, operatorSetIds: defaultOperatorSet.id.toArrayU32()});
        cheats.prank(defaultOperator);
        allocationManager.deregisterFromOperatorSets(deregisterParams);

        // Roll to the slashableUntil at block
        cheats.roll(block.number + DEALLOCATION_DELAY);

        // Slash operator for all wad
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(
            defaultAVS,
            SlashingParams({
                operator: defaultOperator,
                operatorSetId: defaultOperatorSet.id,
                strategies: defaultStrategies,
                wadsToSlash: WAD.toArrayU256(),
                description: "test"
            })
        );

        // Assert that the operator has no max magnitude or allocatable magnitude
        assertEq(0, allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock), "Allocatable magnitude should be 0");
        assertEq(0, allocationManager.getMaxMagnitude(defaultOperator, strategyMock), "Max magnitude should be 0");
    }
}

contract AllocationManagerUnitTests_ModifyAllocations is AllocationManagerUnitTests {
    using ArrayLib for *;
    using OperatorSetLib for *;
    using SlashingLib for *;

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

    /**
     * @notice Regression tests for the bugfix where pending modifications were checked by
     * require(allocation.pendingDiff == 0, ModificationAlreadyPending());
     * which would overwrite the effectBlock, pendingDiff if a pendingDiff
     * of a deallocation was slashed to become 0.
     *
     * This test checks that the effectBlock, pendingDiff are not overwritten even if the pendingDiff is 0
     * when attempting to modify allocations again
     */
    function test_modifyAllocations_PendingDiffZero() public {
        // Step 1: Allocate to the operator set
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, 501);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Step 2: Roll blocks forward until the allocation effectBlock
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Step 3: Deallocate from the operator set
        AllocateParams[] memory deallocateParams = _newAllocateParams(defaultOperatorSet, 500);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);

        Allocation memory allocation = allocationManager.getAllocation(defaultOperator, defaultOperatorSet, strategyMock);
        uint32 originalEffectBlock = allocation.effectBlock;

        // Step 4: Slash the operator to adjust pendingDiff to 0, slashing rounds up the amount of magnitude to slash
        // so with an existing deallocation/pendingDiff of 1, it should result in a pendingDiff of 0
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: defaultOperatorSet.id,
            strategies: defaultStrategies,
            wadsToSlash: 5e17.toArrayU256(),
            description: "Test slashing"
        });
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);
        allocation = allocationManager.getAllocation(defaultOperator, defaultOperatorSet, strategyMock);
        assertEq(allocation.pendingDiff, 0, "Pending diff should be 0");
        assertEq(allocation.effectBlock, originalEffectBlock, "Effect block should not have changed");

        // Step 5: Modify allocations again (Should not be called)
        AllocateParams[] memory newAllocateParams = _newAllocateParams(defaultOperatorSet, 1000);
        cheats.prank(defaultOperator);
        cheats.expectRevert(ModificationAlreadyPending.selector);
        allocationManager.modifyAllocations(defaultOperator, newAllocateParams);

        // Assert that the allocation was modified without reverting
        allocation = allocationManager.getAllocation(defaultOperator, defaultOperatorSet, strategyMock);
        assertEq(allocation.currentMagnitude, 250, "Allocation should be updated to 250 after slashing 50%");

        // Note: These 2 assertions fail prior to the bugfix and if we kept the same
        // require(allocation.pendingDiff == 0, ModificationAlreadyPending());
        // in the code. The effectBlock, pendingDiff would get overwritten with the new modification
        // but the deallocationQueue would now be unordered(in terms of effectBlocks) with this overwrite.
        assertEq(allocation.effectBlock, originalEffectBlock, "Effect block should not have changed");
        assertEq(allocation.pendingDiff, 0, "Pending diff should still be 0");
    }

    /**
     * @notice Regression tests for the bugfix where pending modifications were checked by
     * require(allocation.pendingDiff == 0, ModificationAlreadyPending());
     * which would overwrite the effectBlock, pendingDiff if a pendingDiff
     * of a deallocation was slashed to become 0.
     *
     * This test checks that the deallocationQueue is ascending ordered by effectBlocks
     */
    function test_modifyAllocations_PendingDiffZero_CheckOrderedDeallocationQueue() public {
        // Step 1: Register the operator to multiple operator sets
        OperatorSet memory operatorSet1 = OperatorSet(defaultAVS, 1);
        OperatorSet memory operatorSet2 = OperatorSet(defaultAVS, 2);
        _createOperatorSet(operatorSet1, defaultStrategies);
        _createOperatorSet(operatorSet2, defaultStrategies);
        _registerForOperatorSet(defaultOperator, operatorSet1);
        _registerForOperatorSet(defaultOperator, operatorSet2);

        // Step 2: Allocate to both operator sets
        AllocateParams[] memory allocateParams1 = _newAllocateParams(operatorSet1, 1001);
        AllocateParams[] memory allocateParams2 = _newAllocateParams(operatorSet2, 1000);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams1);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams2);

        // Step 3: Roll blocks forward until the allocation effectBlock
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Step 4: Deallocate from both operator sets
        AllocateParams[] memory deallocateParams1 = _newAllocateParams(operatorSet1, 1000);
        AllocateParams[] memory deallocateParams2 = _newAllocateParams(operatorSet2, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams1);
        // roll blocks forward so that deallocations have different effectBlocks
        cheats.roll(block.number + 100);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams2);

        // Step 5: Slash the first deallocation to adjust pendingDiff to 0
        SlashingParams memory slashingParams1 = SlashingParams({
            operator: defaultOperator,
            operatorSetId: operatorSet1.id,
            strategies: defaultStrategies,
            wadsToSlash: 5e17.toArrayU256(),
            description: "Test slashing"
        });
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams1);

        // Step 6: Modify allocations again for operatorSet1 making another deallocation and
        // overwriting/increasing the effectBlock
        // roll blocks forward so that deallocations have different effectBlocks
        cheats.roll(block.number + 100);
        // Note: this should revert but previously it would not prior to the bugfix
        AllocateParams[] memory newAllocateParams1 = _newAllocateParams(operatorSet1, 400);
        cheats.prank(defaultOperator);
        cheats.expectRevert(ModificationAlreadyPending.selector);
        allocationManager.modifyAllocations(defaultOperator, newAllocateParams1);

        // Assert that the deallocationQueue is unordered for the 2 deallocations in queue
        _checkDeallocationQueueOrder(defaultOperator, defaultStrategies[0], 2);
    }

    /**
     * @notice Regression tests for the bugfix where pending modifications were checked by
     * require(allocation.pendingDiff == 0, ModificationAlreadyPending());
     * which would overwrite the effectBlock, pendingDiff if a pendingDiff
     * of a deallocation was slashed to become 0.
     *
     * This test checks that the deallocationQueue is ascending ordered by effectBlocks
     * In this case the new modifyAllocations call is an allocation
     * where the effectBlock is increased and the deallocationQueue is unordered as well because the operator
     * allocationDelay configured to be long enough.
     */
    function test_modifyAllocations_PendingDiffZero_Allocation() public {
        // Step 1: Register the operator to multiple operator sets
        OperatorSet memory operatorSet1 = OperatorSet(defaultAVS, 1);
        OperatorSet memory operatorSet2 = OperatorSet(defaultAVS, 2);
        _createOperatorSet(operatorSet1, defaultStrategies);
        _createOperatorSet(operatorSet2, defaultStrategies);
        _registerForOperatorSet(defaultOperator, operatorSet1);
        _registerForOperatorSet(defaultOperator, operatorSet2);

        // Step 2: Allocate to both operator sets
        AllocateParams[] memory allocateParams1 = _newAllocateParams(operatorSet1, 1001);
        AllocateParams[] memory allocateParams2 = _newAllocateParams(operatorSet2, 1000);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams1);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams2);

        // Step 3: Update operator allocation delay
        cheats.prank(defaultOperator);
        allocationManager.setAllocationDelay(defaultOperator, DEALLOCATION_DELAY + 10 days);
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);

        // Step 4: Deallocate from both operator sets
        AllocateParams[] memory deallocateParams1 = _newAllocateParams(operatorSet1, 1000);
        AllocateParams[] memory deallocateParams2 = _newAllocateParams(operatorSet2, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams1);
        // roll blocks forward so that deallocations have different effectBlocks
        cheats.roll(block.number + 100);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams2);

        // Step 5: Slash the first deallocation to adjust pendingDiff to 0
        SlashingParams memory slashingParams1 = SlashingParams({
            operator: defaultOperator,
            operatorSetId: operatorSet1.id,
            strategies: defaultStrategies,
            wadsToSlash: 5e17.toArrayU256(),
            description: "Test slashing"
        });
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams1);

        // Step 6: Modify allocations again for operatorSet1 making an allocation and
        // overwriting/increasing the effectBlock
        // Note: this should revert but previously it would not prior to the bugfix
        AllocateParams[] memory newAllocateParams1 = _newAllocateParams(operatorSet1, 5000);
        cheats.prank(defaultOperator);
        cheats.expectRevert(ModificationAlreadyPending.selector);
        allocationManager.modifyAllocations(defaultOperator, newAllocateParams1);

        // Assert that the deallocationQueue is unordered for the 2 deallocations in queue
        _checkDeallocationQueueOrder(defaultOperator, defaultStrategies[0], 2);
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

    function testFuzz_revert_insufficientAllocatableMagnitude(Randomness r) public rand(r) {
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

    function test_revert_safeCastOverflow() public {
        // setup additional operatorSets for tests
        OperatorSet memory opSet1 = OperatorSet(defaultAVS, 1);
        _createOperatorSet(opSet1, defaultStrategies);
        _registerOperator(defaultOperator);
        _setAllocationDelay(defaultOperator, DEFAULT_OPERATOR_ALLOCATION_DELAY);
        _registerForOperatorSet(defaultOperator, opSet1);

        OperatorSet memory opSet2 = OperatorSet(defaultAVS, 2);
        _createOperatorSet(opSet2, defaultStrategies);
        _registerOperator(defaultOperator);
        _setAllocationDelay(defaultOperator, DEFAULT_OPERATOR_ALLOCATION_DELAY);
        _registerForOperatorSet(defaultOperator, opSet2);

        // 1. Allocate all available magnitude for the strategy (WAD)
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();
        allocateParams[0].newMagnitudes[0] = WAD;
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        assertEq(allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock), 0, "Allocatable magnitude should be 0");
        assertEq(allocationManager.getEncumberedMagnitude(defaultOperator, strategyMock), WAD, "Encumbered magnitude should be WAD");

        // 2. allocate to another operatorSet for the same strategy to reset encumberedMagnitude back to 0
        allocateParams[0].operatorSet = opSet1;
        allocateParams[0].newMagnitudes[0] = type(uint64).max - WAD + 1;
        cheats.prank(defaultOperator);
        cheats.expectRevert("SafeCast: value doesn't fit in 64 bits");
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // 3. after resetting encumberedMagnitude, attempt to allocate to opSet2 with WAD
        allocateParams[0].operatorSet = opSet2;
        allocateParams[0].newMagnitudes[0] = WAD;
        cheats.prank(defaultOperator);
        cheats.expectRevert(InsufficientMagnitude.selector);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // 4. after resetting encumberedMagnitude, attempt to allocate to opSet2 with 1
        allocateParams[0].operatorSet = opSet2;
        allocateParams[0].newMagnitudes[0] = 1;
        cheats.prank(defaultOperator);
        cheats.expectRevert(InsufficientMagnitude.selector);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
    }

    /**
     * @notice Tests edge cases around allocation delay:
     * 1. Set allocation delay to a value greater than ALLOCATION_CONFIGURATION_DELAY
     * 2. Allocate magnitude before the configured delay is hit
     * 3. Set allocation delay to a value less than ALLOCATION_CONFIGURATION_DELAY
     * 4. Allocate magnitude after allocation in step 2 takes effect, but before the new delay is hit
     * Validates that you should be able to allocate in step 4 since there is no other pending modifications
     */
    function testFuzz_ShouldBeAbleToAllocateSoonerThanLastDelay(Randomness r) public rand(r) {
        uint32 firstDelay = r.Uint32(ALLOCATION_CONFIGURATION_DELAY, type(uint24).max);
        uint32 secondDelay = r.Uint32(1, ALLOCATION_CONFIGURATION_DELAY - 1);
        uint64 half = 0.5 ether;

        cheats.prank(defaultAVS);
        allocationManager.createOperatorSets(defaultAVS, CreateSetParams(1, defaultStrategies).toArray());

        cheats.startPrank(defaultOperator);

        allocationManager.setAllocationDelay(defaultOperator, firstDelay);

        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(defaultOperatorSet, half));

        // Validate storage - the `firstDelay` should not be applied yet
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: 0, pendingDiff: int64(half), effectBlock: _defaultAllocEffectBlock()}),
            expectedMagnitudes: Magnitudes({encumbered: half, max: WAD, allocatable: WAD - half})
        });

        allocationManager.setAllocationDelay(defaultOperator, secondDelay);

        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY + 1);
        allocationManager.modifyAllocations(defaultOperator, _newAllocateParams(defaultOperatorSet, half + 1));

        cheats.stopPrank();

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: half, pendingDiff: int64(1), effectBlock: uint32(block.number + secondDelay)}),
            expectedMagnitudes: Magnitudes({encumbered: half + 1, max: WAD, allocatable: WAD - (half + 1)})
        });
    }

    /**
     * @notice Allocates a random magnitude to the default operatorSet.
     * Validates:
     * 1. Storage is clear prior to allocation
     * 2. Events are emitted
     * 3. Allocation storage/introspection after allocation
     * 4. Allocation storage/introspection after roll to allocation effect block
     */
    function testFuzz_allocate_singleStrat_singleOperatorSet(Randomness r) public rand(r) {
        // Create allocation
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();

        // Save vars to check against
        uint64 magnitude = allocateParams[0].newMagnitudes[0];
        uint32 effectBlock = _defaultAllocEffectBlock();

        // Check that the operator has no allocated sets/strats before allocation
        OperatorSet[] memory allocatedSets = allocationManager.getAllocatedSets(defaultOperator);
        IStrategy[] memory allocatedStrats = allocationManager.getAllocatedStrategies(defaultOperator, defaultOperatorSet);
        assertEq(allocatedSets.length, 0, "should not have any allocated sets before allocation");
        assertEq(allocatedStrats.length, 0, "should not have any allocated strats before allocation");

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            magnitude: magnitude,
            encumberedMagnitude: magnitude,
            effectBlock: effectBlock
        });

        // Allocate magnitude
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Check storage Prior to Completion

        // 1. Validate allocated sets and strategies
        allocatedSets = allocationManager.getAllocatedSets(defaultOperator);
        allocatedStrats = allocationManager.getAllocatedStrategies(defaultOperator, defaultOperatorSet);
        assertEq(allocatedSets.length, 1, "should have a single allocated set");
        assertEq(allocatedSets[0].key(), defaultOperatorSet.key(), "should be allocated to default set");
        assertEq(allocatedStrats.length, 1, "should have a single allocated strategy to default set");
        assertEq(address(allocatedStrats[0]), address(strategyMock), "should have allocated default strat");

        // 2. Validate allocation + info
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: 0, pendingDiff: int128(uint128(magnitude)), effectBlock: effectBlock}),
            expectedMagnitudes: Magnitudes({encumbered: magnitude, max: WAD, allocatable: WAD - magnitude})
        });

        // 3. Check allocation and info after roll to completion
        cheats.roll(effectBlock);
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: magnitude, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({encumbered: magnitude, max: WAD, allocatable: WAD - magnitude})
        });
    }

    /**
     * @notice Allocates magnitude for a single strategy to multiple operatorSets
     * Validates:
     * 1. Events
     * 2. Allocation storage/introspection after allocation
     * 3. Allocation storage/introspection after roll to allocation effect block
     */
    function testFuzz_allocate_singleStrat_multipleSets(Randomness r) public rand(r) {
        uint8 numOpSets = uint8(r.Uint256(1, FUZZ_MAX_OP_SETS));

        // Create and register for operator sets, each with a single default strategy
        OperatorSet[] memory operatorSets = r.OperatorSetArray(defaultAVS, numOpSets);
        AllocateParams[] memory allocateParams = _randAllocateParams_SingleMockStrategy(operatorSets);

        _createOperatorSets(operatorSets, defaultStrategies);
        _registerForOperatorSets(defaultOperator, operatorSets);

        // Save vars to check against
        uint32 effectBlock = _defaultAllocEffectBlock();
        uint64 usedMagnitude;
        for (uint i; i < allocateParams.length; ++i) {
            usedMagnitude += allocateParams[i].newMagnitudes[0];
        }

        // Validate events
        for (uint i; i < allocateParams.length; ++i) {
            // There is only one strategy in each allocateParams, so we don't need a nested for loop
            _checkAllocationEvents({
                operator: defaultOperator,
                operatorSet: operatorSets[i],
                strategy: strategyMock,
                magnitude: allocateParams[i].newMagnitudes[0],
                encumberedMagnitude: _encumberedMagnitudes[strategyMock] + allocateParams[i].newMagnitudes[0],
                effectBlock: effectBlock
            });
            _encumberedMagnitudes[strategyMock] += allocateParams[i].newMagnitudes[0];
        }

        // Allocate magnitude
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Check storage

        // 1. Sanity check number of allocated sets
        OperatorSet[] memory allocatedSets = allocationManager.getAllocatedSets(defaultOperator);
        assertEq(allocatedSets.length, numOpSets, "should have multiple allocated sets");

        // 2. Check storage after allocation
        for (uint i; i < allocateParams.length; ++i) {
            _checkAllocationStorage({
                operator: defaultOperator,
                operatorSet: operatorSets[i],
                strategy: strategyMock,
                expectedAllocation: Allocation({
                    currentMagnitude: 0,
                    pendingDiff: int128(uint128(allocateParams[i].newMagnitudes[0])),
                    effectBlock: effectBlock
                }),
                expectedMagnitudes: Magnitudes({
                    encumbered: _encumberedMagnitudes[strategyMock],
                    max: WAD,
                    allocatable: WAD - _encumberedMagnitudes[strategyMock]
                })
            });

            IStrategy[] memory allocatedStrats = allocationManager.getAllocatedStrategies(defaultOperator, operatorSets[i]);
            assertEq(allocatedStrats.length, 1, "should have a single allocated strategy to each set");
            assertEq(address(allocatedStrats[0]), address(strategyMock), "should have allocated default strat");
            assertEq(allocatedSets[i].key(), operatorSets[i].key(), "should be allocated to expected set");
        }

        // 3. Check storage after roll to completion
        cheats.roll(effectBlock);
        for (uint i; i < allocateParams.length; ++i) {
            _checkAllocationStorage({
                operator: defaultOperator,
                operatorSet: operatorSets[i],
                strategy: strategyMock,
                expectedAllocation: Allocation({currentMagnitude: allocateParams[i].newMagnitudes[0], pendingDiff: 0, effectBlock: 0}),
                expectedMagnitudes: Magnitudes({
                    encumbered: _encumberedMagnitudes[strategyMock],
                    max: WAD,
                    allocatable: WAD - _encumberedMagnitudes[strategyMock]
                })
            });
        }
    }

    /**
     * @notice Allocates once, warps to allocation effect block, then allocates again
     * Validates:
     * 1. Events for each allocation
     * 2. Allocation storage/introspection immediately after each allocation
     */
    function testFuzz_allocateMultipleTimes(Randomness r) public rand(r) {
        uint64 firstAlloc = r.Uint64(1, WAD - 1);
        uint64 secondAlloc = r.Uint64(firstAlloc + 1, WAD);

        // Validate events
        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            magnitude: firstAlloc,
            encumberedMagnitude: firstAlloc,
            effectBlock: _defaultAllocEffectBlock()
        });

        // Allocate magnitude
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, firstAlloc);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: 0, pendingDiff: int64(firstAlloc), effectBlock: _defaultAllocEffectBlock()}),
            expectedMagnitudes: Magnitudes({encumbered: firstAlloc, max: WAD, allocatable: WAD - firstAlloc})
        });

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Allocate magnitude again
        allocateParams = _newAllocateParams(defaultOperatorSet, secondAlloc);

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            magnitude: secondAlloc,
            encumberedMagnitude: secondAlloc,
            effectBlock: _defaultAllocEffectBlock()
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Check storage
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({
                currentMagnitude: firstAlloc,
                pendingDiff: int64(secondAlloc - firstAlloc),
                effectBlock: _defaultAllocEffectBlock()
            }),
            expectedMagnitudes: Magnitudes({encumbered: secondAlloc, max: WAD, allocatable: WAD - secondAlloc})
        });
    }

    /**
     * Allocates maximum magnitude to multiple strategies for the same operatorSet
     * Validates that encumbered magnitude is max for each strategy
     */
    function testFuzz_allocateMaxToMultipleStrategies(Randomness r) public rand(r) {
        uint numStrats = r.Uint256(2, FUZZ_MAX_STRATS);

        OperatorSet memory operatorSet = OperatorSet(defaultAVS, r.Uint32());
        IStrategy[] memory strategies = r.StrategyArray(numStrats);

        _createOperatorSet(operatorSet, strategies);
        _registerForOperatorSet(defaultOperator, operatorSet);

        for (uint i; i < numStrats; ++i) {
            _checkAllocationEvents({
                operator: defaultOperator,
                operatorSet: operatorSet,
                strategy: strategies[i],
                magnitude: WAD,
                encumberedMagnitude: WAD,
                effectBlock: _defaultAllocEffectBlock()
            });
        }

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(
            defaultOperator,
            AllocateParams({operatorSet: operatorSet, strategies: strategies, newMagnitudes: WAD.toArrayU64(numStrats)}).toArray()
        );

        for (uint i; i < numStrats; ++i) {
            assertEq(WAD, allocationManager.getEncumberedMagnitude(defaultOperator, strategies[i]), "encumberedMagnitude not max");
        }
    }

    /**
     * Allocates to `firstMod` magnitude and then deallocate to `secondMod` magnitude
     * Validates:
     * 1. Events are valid for the allocation and deallocation
     * 2. Storage after the allocation is made
     * 3. Storage after the deallocation is made
     * 4. Storage after the deallocation effect block is hit
     * 5. Storage after the deallocation queue is cleared (specifically encumbered mag is decreased)
     */
    function testFuzz_allocate_deallocate_whenRegistered(Randomness r) public rand(r) {
        // Bound allocation and deallocation
        uint64 firstMod = r.Uint64(1, WAD);
        uint64 secondMod = r.Uint64(0, firstMod - 1);

        // Allocate magnitude to default registered set
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, firstMod);

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            magnitude: firstMod,
            encumberedMagnitude: firstMod,
            effectBlock: _defaultAllocEffectBlock()
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        allocateParams = _newAllocateParams(defaultOperatorSet, secondMod);

        _checkDeallocationEvent({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            magnitude: secondMod,
            effectBlock: uint32(block.number + DEALLOCATION_DELAY + 1)
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Check storage after dealloc
        uint32 effectBlock = uint32(block.number + DEALLOCATION_DELAY + 1);

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({
                currentMagnitude: firstMod,
                pendingDiff: -int128(uint128(firstMod - secondMod)),
                effectBlock: effectBlock
            }),
            expectedMagnitudes: Magnitudes({encumbered: firstMod, max: WAD, allocatable: WAD - firstMod})
        });

        // Check storage after roll to completion
        cheats.roll(effectBlock);
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: secondMod, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({encumbered: secondMod, max: WAD, allocatable: WAD - secondMod})
        });

        // Check storage after clearing deallocation queue
        allocationManager.clearDeallocationQueue(defaultOperator, strategyMock.toArray(), uint16(1).toArrayU16());
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: secondMod, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({encumbered: secondMod, max: WAD, allocatable: WAD - secondMod})
        });
    }

    /**
     * Allocates to an operatorSet, then fully deallocates after the strategy is removed from the set.
     * Validates that the deallocation takes effect immediately after the strategy is removed
     */
    function test_allocate_removeStrategyFromSet_fullyDeallocate() public {
        // Allocate
        AllocateParams[] memory allocateParams = _randAllocateParams_SingleMockStrategy(defaultOperatorSet.toArray());
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Remove strategy from operatorSet
        cheats.prank(defaultAVS);
        allocationManager.removeStrategiesFromOperatorSet(defaultAVS, defaultOperatorSet.id, defaultStrategies);

        // Deallocate All Instantly
        AllocateParams[] memory deallocateParams = allocateParams;
        deallocateParams[0].newMagnitudes[0] = 0;

        // We check the allocation event and not the deallocation event since the encumbered mag is updated too!
        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            magnitude: 0,
            encumberedMagnitude: 0,
            effectBlock: uint32(block.number)
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: 0, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({encumbered: 0, max: WAD, allocatable: WAD})
        });
    }

    /**
     * Allocates to an operatorSet, deallocates, then removes a strategy from the operatorSet
     * Validates that:
     * 1. The deallocation still completes at its expected time
     */
    function testFuzz_allocate_deallocate_removeStrategyFromSet(Randomness r) public {
        // Allocate
        AllocateParams[] memory allocateParams = _randAllocateParams_SingleMockStrategy(defaultOperatorSet.toArray());
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        AllocateParams[] memory deallocateParams = r.DeallocateParams(allocateParams);
        uint32 deallocEffectBlock = uint32(block.number + DEALLOCATION_DELAY + 1);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);

        // Remove strategy from operatorSet
        cheats.prank(defaultAVS);
        allocationManager.removeStrategiesFromOperatorSet(defaultAVS, defaultOperatorSet.id, defaultStrategies);

        // Roll to just before deallocation complete block & clear deallocation queue for sanity
        cheats.roll(deallocEffectBlock - 1);
        allocationManager.clearDeallocationQueue(defaultOperator, strategyMock.toArray(), _maxNumToClear());

        int128 expectedDiff = -int128(uint128(allocateParams[0].newMagnitudes[0] - deallocateParams[0].newMagnitudes[0]));
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({
                currentMagnitude: allocateParams[0].newMagnitudes[0],
                pendingDiff: expectedDiff,
                effectBlock: deallocEffectBlock
            }),
            expectedMagnitudes: Magnitudes({
                encumbered: allocateParams[0].newMagnitudes[0],
                max: WAD,
                allocatable: WAD - allocateParams[0].newMagnitudes[0]
            })
        });

        // Roll to deallocation complete block
        cheats.roll(deallocEffectBlock);

        // Note that the encumbered mag hasn't been updated since we haven't cleared the deallocaction queue!
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: deallocateParams[0].newMagnitudes[0], pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({
                encumbered: deallocateParams[0].newMagnitudes[0],
                max: WAD,
                allocatable: WAD - deallocateParams[0].newMagnitudes[0]
            })
        });
    }

    /**
     * Allocates to an operator set, then fully deallocates when not registered to the set.
     * Validates that:
     * 1. Events are properly emitted post instantaneous deallocation
     * 2. The deallocation is instant & can be reallocated immediately
     * 3. Storage/introspection post combined deallocation/allocation
     */
    function testFuzz_allocate_fullyDeallocate_reallocate_WhenNotRegistered(Randomness r) public rand(r) {
        // Bound allocation and deallocation
        uint64 firstMod = r.Uint64(1, WAD);

        // Create new operator sets that the operator is not registered for
        OperatorSet memory operatorSetA = _createOperatorSet(OperatorSet(defaultAVS, r.Uint32()), defaultStrategies);
        OperatorSet memory operatorSetB = _createOperatorSet(OperatorSet(defaultAVS, r.Uint32()), defaultStrategies);

        // Allocate magnitude to operator set
        AllocateParams[] memory allocateParams = _newAllocateParams(operatorSetA, firstMod);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate instantly and reallocate all magnitude to second operator set
        allocateParams = new AllocateParams[](2);
        allocateParams[0] = _newAllocateParams(operatorSetA, 0)[0];
        allocateParams[1] = _newAllocateParams(operatorSetB, firstMod)[0];

        // We check the allocation event and not the deallocation event since
        // encumbered magnitude is also updated here
        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: operatorSetA,
            strategy: strategyMock,
            magnitude: 0,
            encumberedMagnitude: 0,
            effectBlock: uint32(block.number) // Instant deallocation
        });
        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: operatorSetB,
            strategy: strategyMock,
            magnitude: firstMod,
            encumberedMagnitude: firstMod,
            effectBlock: _defaultAllocEffectBlock()
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Check storage after deallocation
        // Check operator set A
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: operatorSetA,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: 0, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({encumbered: firstMod, max: WAD, allocatable: WAD - firstMod}) // This is from opsetB
        });

        // Check operator set B
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: operatorSetB,
            strategy: strategyMock,
            expectedAllocation: Allocation({
                currentMagnitude: 0,
                pendingDiff: int128(uint128(firstMod)),
                effectBlock: _defaultAllocEffectBlock()
            }),
            expectedMagnitudes: Magnitudes({encumbered: firstMod, max: WAD, allocatable: WAD - firstMod})
        });
    }

    /**
     * Allocates all magnitude to a single strategy across multiple operatorSets. Deallocates fully, and then reallocates
     * Validates:
     * 1. Events are emitted for the allocation, deallocation, and reallocation (including the deallocation queue clear)
     * 2. Stake is fully allocated & encumbered mag used up
     * 3. Stake can be reallocated after the deallocation delay
     */
    function testFuzz_allocate_fromClearedDeallocQueue(Randomness r) public rand(r) {
        uint numOpSets = r.Uint256(1, FUZZ_MAX_OP_SETS);

        // Create multiple operator sets, register, and allocate to each. Ensure all magnitude is fully allocated.
        OperatorSet[] memory deallocSets = r.OperatorSetArray(defaultAVS, numOpSets);
        _createOperatorSets(deallocSets, defaultStrategies);
        _registerForOperatorSets(defaultOperator, deallocSets);
        AllocateParams[] memory allocateParams = _randAllocateParams_SingleMockStrategy_AllocAll(deallocSets);

        for (uint i; i < allocateParams.length; ++i) {
            // There is only one strategy each allocateParams, so we don't need a nested for loop
            _checkAllocationEvents({
                operator: defaultOperator,
                operatorSet: allocateParams[i].operatorSet,
                strategy: strategyMock,
                magnitude: allocateParams[i].newMagnitudes[0],
                encumberedMagnitude: _encumberedMagnitudes[strategyMock] + allocateParams[i].newMagnitudes[0],
                effectBlock: _defaultAllocEffectBlock()
            });
            _encumberedMagnitudes[strategyMock] += allocateParams[i].newMagnitudes[0];
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

        for (uint i; i < numOpSets; ++i) {
            _checkDeallocationEvent({
                operator: defaultOperator,
                operatorSet: deallocSets[i],
                strategy: strategyMock,
                magnitude: 0,
                effectBlock: uint32(block.number + DEALLOCATION_DELAY + 1)
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
        cheats.roll(block.number + DEALLOCATION_DELAY + 1);

        // Check that we now have sufficient allocatable magnitude
        assertEq(
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock), WAD, "operator should have all magnitude allocatable"
        );

        // Create and register for a new operator set with the same default strategy.
        // If we try to allocate to this new set, it should clear the deallocation queue,
        // allowing all magnitude to be allocated
        OperatorSet memory finalOpSet = _createOperatorSet(OperatorSet(defaultAVS, r.Uint32()), defaultStrategies);
        _registerForOperatorSet(defaultOperator, finalOpSet);
        AllocateParams[] memory finalAllocParams = _newAllocateParams(finalOpSet, WAD);

        _checkClearDeallocationQueueEvents({operator: defaultOperator, strategy: strategyMock, encumberedMagnitude: 0});

        _checkAllocationEvents({
            operator: defaultOperator,
            operatorSet: finalOpSet,
            strategy: strategyMock,
            magnitude: WAD,
            encumberedMagnitude: WAD,
            effectBlock: _defaultAllocEffectBlock()
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, finalAllocParams);

        // Check that all magnitude will be allocated to the new set, and each prior set
        // has a zeroed-out allocation
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: finalOpSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: 0, pendingDiff: int128(uint128(WAD)), effectBlock: _defaultAllocEffectBlock()}),
            expectedMagnitudes: Magnitudes({encumbered: WAD, max: WAD, allocatable: 0})
        });
    }

    /**
     * Allocates all mag and then deallocates all mag
     * Validates
     * 1. Events for the deallocation
     * 2. Storage after deallocation
     * 3. Storage after clearing the deallocation queue
     */
    function test_deallocate_all() public {
        // Allocate all
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, WAD);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate all
        allocateParams[0].newMagnitudes[0] = 0;

        _checkDeallocationEvent({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            magnitude: 0,
            effectBlock: uint32(block.number + DEALLOCATION_DELAY + 1)
        });

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Warp to completion and clear deallocation queue
        cheats.roll(block.number + DEALLOCATION_DELAY + 1);
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, uint16(1).toArrayU16());

        // Check storage
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: 0, pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({encumbered: 0, max: WAD, allocatable: WAD})
        });
    }

    /**
     * Allocates, deallocates, and then clears the deallocation queue. Multiple strategies & sets in a single operatorSet
     * Validates:
     * 1. Events for allocation, deallocation, and deallocation queue clear
     * 2. Storage after allocation & after allocation effect block
     * 3. Storage after deallocation & after deallocation effect block
     */
    function testFuzz_lifecycle_allocate_deallocate_MultipleSetsAndStrats(Randomness r) public rand(r) {
        uint numAllocations = r.Uint256(2, FUZZ_MAX_ALLOCATIONS);
        uint numStrats = r.Uint256(2, FUZZ_MAX_STRATS);

        AllocateParams[] memory allocateParams = r.AllocateParams(defaultAVS, numAllocations, numStrats);
        AllocateParams[] memory deallocateParams = r.DeallocateParams(allocateParams);
        CreateSetParams[] memory createSetParams = r.CreateSetParams(allocateParams);

        cheats.prank(defaultAVS);
        allocationManager.createOperatorSets(defaultAVS, createSetParams);

        for (uint i = 0; i < allocateParams.length; i++) {
            _registerForOperatorSet(defaultOperator, allocateParams[i].operatorSet);
        }

        // Allocate
        for (uint i; i < allocateParams.length; ++i) {
            for (uint j; j < allocateParams[i].strategies.length; ++j) {
                _checkAllocationEvents({
                    operator: defaultOperator,
                    operatorSet: allocateParams[i].operatorSet,
                    strategy: allocateParams[i].strategies[j],
                    magnitude: allocateParams[i].newMagnitudes[j],
                    encumberedMagnitude: allocateParams[i].newMagnitudes[j],
                    effectBlock: _defaultAllocEffectBlock()
                });
            }
        }

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Check storage after allocation
        for (uint i; i < allocateParams.length; ++i) {
            for (uint j = 0; j < allocateParams[i].strategies.length; j++) {
                _checkAllocationStorage({
                    operator: defaultOperator,
                    operatorSet: allocateParams[i].operatorSet,
                    strategy: allocateParams[i].strategies[j],
                    expectedAllocation: Allocation({
                        currentMagnitude: 0,
                        pendingDiff: int128(uint128(allocateParams[i].newMagnitudes[j])),
                        effectBlock: _defaultAllocEffectBlock()
                    }),
                    expectedMagnitudes: Magnitudes({
                        encumbered: allocateParams[i].newMagnitudes[j],
                        max: WAD,
                        allocatable: WAD - allocateParams[i].newMagnitudes[j]
                    })
                });
            }
        }

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Check storage after roll to completion
        for (uint i; i < allocateParams.length; ++i) {
            for (uint j = 0; j < allocateParams[i].strategies.length; j++) {
                _checkAllocationStorage({
                    operator: defaultOperator,
                    operatorSet: allocateParams[i].operatorSet,
                    strategy: allocateParams[i].strategies[j],
                    expectedAllocation: Allocation({currentMagnitude: allocateParams[i].newMagnitudes[j], pendingDiff: 0, effectBlock: 0}),
                    expectedMagnitudes: Magnitudes({
                        encumbered: allocateParams[i].newMagnitudes[j],
                        max: WAD,
                        allocatable: WAD - allocateParams[i].newMagnitudes[j]
                    })
                });
            }
        }

        // Deallocate

        for (uint i; i < deallocateParams.length; ++i) {
            for (uint j = 0; j < deallocateParams[i].strategies.length; j++) {
                _checkDeallocationEvent({
                    operator: defaultOperator,
                    operatorSet: deallocateParams[i].operatorSet,
                    strategy: deallocateParams[i].strategies[j],
                    magnitude: deallocateParams[i].newMagnitudes[j],
                    effectBlock: uint32(block.number + DEALLOCATION_DELAY + 1)
                });
            }
        }

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);

        for (uint i; i < allocateParams.length; ++i) {
            for (uint j = 0; j < allocateParams[i].strategies.length; j++) {
                int128 expectedDiff = -int128(uint128(allocateParams[i].newMagnitudes[j] - deallocateParams[i].newMagnitudes[j]));
                _checkAllocationStorage({
                    operator: defaultOperator,
                    operatorSet: deallocateParams[i].operatorSet,
                    strategy: deallocateParams[i].strategies[j],
                    expectedAllocation: Allocation({
                        currentMagnitude: allocateParams[i].newMagnitudes[j],
                        pendingDiff: expectedDiff,
                        effectBlock: uint32(block.number + DEALLOCATION_DELAY + 1)
                    }),
                    expectedMagnitudes: Magnitudes({
                        encumbered: allocateParams[i].newMagnitudes[j],
                        max: WAD,
                        allocatable: WAD - allocateParams[i].newMagnitudes[j]
                    })
                });
            }
        }

        // Warp to deallocation complete block
        cheats.roll(block.number + DEALLOCATION_DELAY + 1);
        for (uint i; i < allocateParams.length; ++i) {
            for (uint j = 0; j < allocateParams[i].strategies.length; j++) {
                _checkAllocationStorage({
                    operator: defaultOperator,
                    operatorSet: deallocateParams[i].operatorSet,
                    strategy: allocateParams[i].strategies[j],
                    expectedAllocation: Allocation({currentMagnitude: deallocateParams[i].newMagnitudes[j], pendingDiff: 0, effectBlock: 0}),
                    expectedMagnitudes: Magnitudes({
                        encumbered: deallocateParams[i].newMagnitudes[j],
                        max: WAD,
                        allocatable: WAD - deallocateParams[i].newMagnitudes[j]
                    })
                });
            }
        }
    }
}

contract AllocationManagerUnitTests_ClearDeallocationQueue is AllocationManagerUnitTests {
    using ArrayLib for *;

    /// -----------------------------------------------------------------------
    /// clearDeallocationQueue()
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
    function testFuzz_allocate(Randomness r) public rand(r) {
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
            expectedAllocation: Allocation({currentMagnitude: allocateParams[0].newMagnitudes[0], pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({
                encumbered: allocateParams[0].newMagnitudes[0],
                max: WAD,
                allocatable: WAD - allocateParams[0].newMagnitudes[0]
            })
        });
    }

    /**
     * @notice Allocates magnitude to an operator registered for some operator sets, and then
     * - Clears deallocation queue when nothing can be completed
     * - After the first clear, asserts the allocation info takes into account the deallocation
     * - Clears deallocation queue when the dealloc can be completed
     * - Assert events & validates storage after the deallocateParams are completed
     */
    function testFuzz_allocate_deallocate(Randomness r) public rand(r) {
        // Generate a random allocation and subsequent deallocation from the default operator set
        (AllocateParams[] memory allocateParams, AllocateParams[] memory deallocateParams) =
            _randAllocAndDeallocParams_SingleMockStrategy(defaultOperatorSet.toArray());

        // Allocate
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // Roll to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        _checkDeallocationEvent({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            magnitude: deallocateParams[0].newMagnitudes[0],
            effectBlock: uint32(block.number + DEALLOCATION_DELAY + 1)
        });
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);

        // Clear queue - since we have not rolled forward, this should be a no-op
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());

        // Validate storage - encumbered magnitude should just be allocateParams (we only have 1 allocation)
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({
                currentMagnitude: allocateParams[0].newMagnitudes[0],
                pendingDiff: -int128(uint128(allocateParams[0].newMagnitudes[0] - deallocateParams[0].newMagnitudes[0])),
                effectBlock: uint32(block.number + DEALLOCATION_DELAY + 1)
            }),
            expectedMagnitudes: Magnitudes({
                encumbered: allocateParams[0].newMagnitudes[0],
                max: WAD,
                allocatable: WAD - allocateParams[0].newMagnitudes[0]
            })
        });

        // Warp to deallocation complete block
        cheats.roll(block.number + DEALLOCATION_DELAY + 1);

        // Clear queue
        _checkClearDeallocationQueueEvents({
            operator: defaultOperator,
            strategy: strategyMock,
            encumberedMagnitude: deallocateParams[0].newMagnitudes[0]
        });
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());

        // Validate storage - encumbered magnitude should just be deallocateParams (we only have 1 deallocation)
        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: defaultOperatorSet,
            strategy: strategyMock,
            expectedAllocation: Allocation({currentMagnitude: deallocateParams[0].newMagnitudes[0], pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({
                encumbered: deallocateParams[0].newMagnitudes[0],
                max: WAD,
                allocatable: WAD - deallocateParams[0].newMagnitudes[0]
            })
        });
    }

    /**
     * Allocates, deallocates, and then allocates again. Asserts that
     * - The deallocation does not block state updates from the second allocation, even though the allocation has an earlier
     *   effect block
     */
    function test_allocate_deallocate_allocate() public {
        // Allocate half of mag to default operator set
        AllocateParams[] memory firstAllocation = _newAllocateParams(defaultOperatorSet, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, firstAllocation);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate half from default operator set
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY + 1);
        AllocateParams[] memory firstDeallocation = _newAllocateParams(defaultOperatorSet, 25e16);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, firstDeallocation);
        Allocation memory allocation = allocationManager.getAllocation(defaultOperator, defaultOperatorSet, strategyMock);
        assertEq(deallocationEffectBlock, allocation.effectBlock, "effect block not correct");

        // Create and register for a new operator set
        OperatorSet memory newOperatorSet = _createOperatorSet(OperatorSet(defaultAVS, random().Uint32()), defaultStrategies);
        _registerForOperatorSet(defaultOperator, newOperatorSet);

        // Allocate 33e16 mag to new operator set
        uint32 allocationEffectBlock = _defaultAllocEffectBlock();
        AllocateParams[] memory secondAllocation = _newAllocateParams(newOperatorSet, 33e16);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, secondAllocation);
        allocation = allocationManager.getAllocation(defaultOperator, newOperatorSet, strategyMock);
        assertEq(allocationEffectBlock, allocation.effectBlock, "effect block not correct");

        // Warp to allocation effect block & clear the queue - clearing is a noop here
        cheats.roll(allocationEffectBlock);
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());

        // Validate `getAllocatableMagnitude`. Allocatable magnitude should be the difference between the max magnitude and the encumbered magnitude
        uint64 allocatableMagnitude = allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock);
        assertEq(WAD - 33e16 - 5e17, allocatableMagnitude, "allocatableMagnitude not correct");

        // Validate that we can allocate again for opset2. This should not revert
        AllocateParams[] memory thirdAllocation = _newAllocateParams(newOperatorSet, 10e16);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, thirdAllocation);

        // Warp & validate deallocation
        cheats.roll(deallocationEffectBlock);
        _checkClearDeallocationQueueEvents({operator: defaultOperator, strategy: strategyMock, encumberedMagnitude: 58e16});
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());
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
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY + 1);
        (, uint32 storedDelay) = allocationManager.getAllocationDelay(defaultOperator);
        assertEq(allocationDelay, storedDelay, "allocation delay not valid");

        // Allocate half of mag to default operator set
        AllocateParams[] memory firstAllocation = _newAllocateParams(defaultOperatorSet, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, firstAllocation);
        cheats.roll(block.number + allocationDelay);

        // Create and register for a second operator set
        OperatorSet memory newOperatorSet = _createOperatorSet(OperatorSet(defaultAVS, random().Uint32()), defaultStrategies);
        _registerForOperatorSet(defaultOperator, newOperatorSet);

        // Allocate half of mag to opset2
        AllocateParams[] memory secondAllocation = _newAllocateParams(newOperatorSet, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, secondAllocation);

        uint32 allocationEffectBlock = uint32(block.number + allocationDelay);
        Allocation memory allocation = allocationManager.getAllocation(defaultOperator, newOperatorSet, strategyMock);
        assertEq(allocationEffectBlock, allocation.effectBlock, "effect block not correct");

        // Deallocate all from opSet1
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY + 1);
        AllocateParams[] memory firstDeallocation = _newAllocateParams(defaultOperatorSet, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, firstDeallocation);
        allocation = allocationManager.getAllocation(defaultOperator, defaultOperatorSet, strategyMock);
        assertEq(deallocationEffectBlock, allocation.effectBlock, "effect block not correct");

        // Warp to deallocation effect block & clear the queue
        cheats.roll(deallocationEffectBlock);
        allocationManager.clearDeallocationQueue(defaultOperator, defaultStrategies, _maxNumToClear());

        // At this point, we should be able to allocate again to opSet1 AND have only 5e17 encumbered magnitude
        assertEq(5e17, allocationManager.getEncumberedMagnitude(defaultOperator, strategyMock), "encumbered magnitude not correct");
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
        cheats.expectRevert(InvalidOperator.selector);
        allocationManager.setAllocationDelay(operatorToSet, 1);
    }

    function test_revert_callerNotAuthorized() public {
        cheats.expectRevert(InvalidCaller.selector);
        allocationManager.setAllocationDelay(operatorToSet, 1);
    }

    function test_regression_boundary() public {
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(operatorToSet, 0);

        // Warp to the allocation config delay - we should not be set yet
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);
        (bool isSet, uint32 delay) = allocationManager.getAllocationDelay(operatorToSet);
        assertFalse(isSet, "isSet should not be set");

        // Warp to the next block - we should be set now
        cheats.roll(block.number + 1);
        (isSet, delay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
    }

    function testFuzz_setDelay(Randomness r) public rand(r) {
        uint32 delay = r.Uint32(0, type(uint32).max);

        // Set delay
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AllocationDelaySet(operatorToSet, delay, uint32(block.number + ALLOCATION_CONFIGURATION_DELAY + 1));
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(operatorToSet, delay);

        // Check values after set
        (bool isSet, uint32 returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertFalse(isSet, "isSet should not be set");
        assertEq(0, returnedDelay, "returned delay should be 0");

        // Warp to effect block
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY + 1);

        // Check values after config delay
        (isSet, returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(delay, returnedDelay, "delay not set");
    }

    function test_fuzz_setDelay_multipleTimesWithinConfigurationDelay(Randomness r) public rand(r) {
        uint32 firstDelay = r.Uint32(1, type(uint32).max);
        uint32 secondDelay = r.Uint32(1, type(uint32).max);

        // Set delay
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(operatorToSet, firstDelay);

        // Warp just before effect block
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);

        // Set delay again
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AllocationDelaySet(operatorToSet, secondDelay, uint32(block.number + ALLOCATION_CONFIGURATION_DELAY + 1));
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(operatorToSet, secondDelay);

        // Warp to effect block of first delay
        cheats.roll(block.number + 1);

        // Assert that the delay is still not set
        (bool isSet, uint32 returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertFalse(isSet, "isSet should not be set");
        assertEq(0, returnedDelay, "returned delay should be 0");

        // Warp to effect block of second delay
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY + 1);
        (isSet, returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(secondDelay, returnedDelay, "delay not set");
    }

    function testFuzz_multipleDelays(Randomness r) public rand(r) {
        uint32 firstDelay = r.Uint32(1, type(uint32).max);
        uint32 secondDelay = r.Uint32(1, type(uint32).max);

        // Set delay
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(operatorToSet, firstDelay);

        // Warp to effect block of first delay
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY + 1);

        // Set delay again
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(operatorToSet, secondDelay);

        // Assert that first delay is set
        (bool isSet, uint32 returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(firstDelay, returnedDelay, "delay not set");

        // Warp to effect block of second delay
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY + 1);

        // Check values after second delay
        (isSet, returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(secondDelay, returnedDelay, "delay not set");
    }

    function testFuzz_setDelay_DMCaller(Randomness r) public rand(r) {
        uint32 delay = r.Uint32(1, type(uint32).max);

        cheats.prank(address(delegationManagerMock));
        allocationManager.setAllocationDelay(operatorToSet, delay);

        // Warp to effect block
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY + 1);
        (bool isSet, uint32 returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(delay, returnedDelay, "delay not set");
    }
}

contract AllocationManagerUnitTests_registerForOperatorSets is AllocationManagerUnitTests {
    using ArrayLib for *;

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

    function testFuzz_registerForOperatorSets_InvalidOperator_x(Randomness r) public rand(r) {
        address operator = r.Address();
        cheats.prank(operator);
        cheats.expectRevert(InvalidOperator.selector);
        allocationManager.registerForOperatorSets(operator, defaultRegisterParams);
    }

    function testFuzz_registerForOperatorSets_InvalidOperatorSet(Randomness r) public rand(r) {
        cheats.prank(defaultOperator);
        cheats.expectRevert(InvalidOperatorSet.selector);
        defaultRegisterParams.operatorSetIds[0] = 1; // invalid id
        allocationManager.registerForOperatorSets(defaultOperator, defaultRegisterParams); // invalid id
    }

    function testFuzz_registerForOperatorSets_AlreadyMemberOfSet(Randomness r) public rand(r) {
        cheats.prank(defaultOperator);
        cheats.expectRevert(AlreadyMemberOfSet.selector);
        allocationManager.registerForOperatorSets(defaultOperator, defaultRegisterParams);
    }

    function testFuzz_registerForOperatorSets_Correctness(Randomness r) public rand(r) {
        address operator = r.Address();
        uint numOpSets = r.Uint256(1, FUZZ_MAX_OP_SETS);
        uint32[] memory operatorSetIds = new uint32[](numOpSets);
        CreateSetParams[] memory createSetParams = new CreateSetParams[](numOpSets);

        _registerOperator(operator);

        for (uint i; i < numOpSets; ++i) {
            operatorSetIds[i] = r.Uint32(1, type(uint32).max);
            createSetParams[i].operatorSetId = operatorSetIds[i];
            createSetParams[i].strategies = defaultStrategies;
        }

        cheats.prank(defaultAVS);
        allocationManager.createOperatorSets(defaultAVS, createSetParams);

        for (uint j; j < numOpSets; ++j) {
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit OperatorAddedToOperatorSet(operator, OperatorSet(defaultAVS, operatorSetIds[j]));
        }

        cheats.expectCall(
            defaultAVS, abi.encodeWithSelector(IAVSRegistrar.registerOperator.selector, operator, defaultAVS, operatorSetIds, "")
        );

        cheats.prank(operator);
        allocationManager.registerForOperatorSets(operator, RegisterParams(defaultAVS, operatorSetIds, ""));

        assertEq(allocationManager.getRegisteredSets(operator).length, numOpSets, "should be registered for all sets");

        for (uint k; k < numOpSets; ++k) {
            OperatorSet memory operatorSet = OperatorSet(defaultAVS, operatorSetIds[k]);
            assertTrue(allocationManager.isMemberOfOperatorSet(operator, operatorSet), "should be member of set");
            assertEq(allocationManager.getMembers(OperatorSet(defaultAVS, operatorSetIds[k]))[0], operator, "should be member of set");
        }
    }
}

contract AllocationManagerUnitTests_deregisterFromOperatorSets is AllocationManagerUnitTests {
    using ArrayLib for *;

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

    function testFuzz_deregisterFromOperatorSets_InvalidOperatorSet(Randomness r) public rand(r) {
        defaultDeregisterParams.operatorSetIds = uint32(1).toArrayU32(); // invalid id
        cheats.prank(defaultOperator);
        cheats.expectRevert(InvalidOperatorSet.selector);
        allocationManager.deregisterFromOperatorSets(defaultDeregisterParams);
    }

    function testFuzz_deregisterFromOperatorSets_NotMemberOfSet(Randomness r) public rand(r) {
        defaultDeregisterParams.operator = r.Address();
        cheats.prank(defaultDeregisterParams.operator);
        cheats.expectRevert(NotMemberOfSet.selector);
        allocationManager.deregisterFromOperatorSets(defaultDeregisterParams);
    }

    function testFuzz_deregisterFromOperatorSets_Correctness(Randomness r) public rand(r) {
        uint numOpSets = r.Uint256(1, FUZZ_MAX_OP_SETS);
        uint32[] memory operatorSetIds = new uint32[](numOpSets);
        CreateSetParams[] memory createSetParams = new CreateSetParams[](numOpSets);

        for (uint i; i < numOpSets; ++i) {
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

        for (uint j; j < numOpSets; ++j) {
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit OperatorRemovedFromOperatorSet(operator, OperatorSet(defaultAVS, operatorSetIds[j]));
        }

        cheats.expectCall(
            defaultAVS, abi.encodeWithSelector(IAVSRegistrar.deregisterOperator.selector, operator, defaultAVS, operatorSetIds)
        );

        bool callFromAVS = r.Boolean();
        if (callFromAVS) cheats.prank(defaultAVS);
        else cheats.prank(operator);
        allocationManager.deregisterFromOperatorSets(DeregisterParams(operator, defaultAVS, operatorSetIds));

        assertEq(allocationManager.getRegisteredSets(operator).length, 0, "should not be registered for any sets");

        for (uint k; k < numOpSets; ++k) {
            assertFalse(
                allocationManager.isMemberOfOperatorSet(operator, OperatorSet(defaultAVS, operatorSetIds[k])), "should not be member of set"
            );
            assertEq(allocationManager.getMemberCount(OperatorSet(defaultAVS, operatorSetIds[k])), 0, "should not be member of set");
        }
    }
}

contract AllocationManagerUnitTests_addStrategiesToOperatorSet is AllocationManagerUnitTests {
    using ArrayLib for *;

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

    function testFuzz_addStrategiesToOperatorSet_Correctness(Randomness r) public rand(r) {
        uint numStrategies = r.Uint256(1, FUZZ_MAX_STRATS);

        IStrategy[] memory strategies = new IStrategy[](numStrategies);

        for (uint i; i < numStrategies; ++i) {
            strategies[i] = IStrategy(r.Address());
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit StrategyAddedToOperatorSet(defaultOperatorSet, strategies[i]);
        }

        cheats.prank(defaultAVS);
        allocationManager.addStrategiesToOperatorSet(defaultAVS, defaultOperatorSet.id, strategies);

        IStrategy[] memory strategiesInSet = allocationManager.getStrategiesInOperatorSet(defaultOperatorSet);

        for (uint j; j < numStrategies; ++j) {
            assertTrue(strategiesInSet[j + 1] == strategies[j], "should be strat of set");
        }
    }
}

contract AllocationManagerUnitTests_removeStrategiesFromOperatorSet is AllocationManagerUnitTests {
    using ArrayLib for *;

    function test_removeStrategiesFromOperatorSet_InvalidOperatorSet() public {
        cheats.prank(defaultAVS);
        cheats.expectRevert(InvalidOperatorSet.selector);
        allocationManager.removeStrategiesFromOperatorSet(defaultAVS, 1, defaultStrategies);
    }

    function test_removeStrategiesFromOperatorSet_StrategyNotInOperatorSet() public {
        cheats.prank(defaultAVS);
        cheats.expectRevert(StrategyNotInOperatorSet.selector);
        allocationManager.removeStrategiesFromOperatorSet(defaultAVS, defaultOperatorSet.id, IStrategy(random().Address()).toArray());
    }

    function testFuzz_removeStrategiesFromOperatorSet_Correctness(Randomness r) public rand(r) {
        uint numStrategies = r.Uint256(1, FUZZ_MAX_STRATS);
        IStrategy[] memory strategies = r.StrategyArray(numStrategies);

        cheats.prank(defaultAVS);
        allocationManager.addStrategiesToOperatorSet(defaultAVS, defaultOperatorSet.id, strategies);

        for (uint i; i < numStrategies; ++i) {
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit StrategyRemovedFromOperatorSet(defaultOperatorSet, strategies[i]);
        }

        assertEq(allocationManager.getStrategiesInOperatorSet(defaultOperatorSet).length, numStrategies + 1, "sanity check");

        cheats.prank(defaultAVS);
        allocationManager.removeStrategiesFromOperatorSet(defaultAVS, defaultOperatorSet.id, strategies);

        // The original strategy should still be in the operator set.
        assertEq(allocationManager.getStrategiesInOperatorSet(defaultOperatorSet).length, 1, "should not be strat of set");
    }
}

contract AllocationManagerUnitTests_createOperatorSets is AllocationManagerUnitTests {
    using ArrayLib for *;

    function testRevert_createOperatorSets_InvalidOperatorSet() public {
        cheats.prank(defaultAVS);
        cheats.expectRevert(InvalidOperatorSet.selector);
        allocationManager.createOperatorSets(defaultAVS, CreateSetParams(defaultOperatorSet.id, defaultStrategies).toArray());
    }

    function testRevert_createOperatorSets_NonexistentAVSMetadata(Randomness r) public rand(r) {
        address avs = r.Address();
        cheats.expectRevert(NonexistentAVSMetadata.selector);
        cheats.prank(avs);
        allocationManager.createOperatorSets(avs, CreateSetParams(defaultOperatorSet.id, defaultStrategies).toArray());
    }

    function testFuzz_createOperatorSets_Correctness(Randomness r) public rand(r) {
        address avs = r.Address();
        uint numOpSets = r.Uint256(1, FUZZ_MAX_OP_SETS);
        uint numStrategies = r.Uint256(1, FUZZ_MAX_STRATS);

        cheats.prank(avs);
        allocationManager.updateAVSMetadataURI(avs, "https://example.com");

        CreateSetParams[] memory createSetParams = new CreateSetParams[](numOpSets);

        for (uint i; i < numOpSets; ++i) {
            createSetParams[i].operatorSetId = r.Uint32(1, type(uint32).max);
            createSetParams[i].strategies = r.StrategyArray(numStrategies);
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit OperatorSetCreated(OperatorSet(avs, createSetParams[i].operatorSetId));
            for (uint j; j < numStrategies; ++j) {
                cheats.expectEmit(true, true, true, true, address(allocationManager));
                emit StrategyAddedToOperatorSet(OperatorSet(avs, createSetParams[i].operatorSetId), createSetParams[i].strategies[j]);
            }
        }

        cheats.prank(avs);
        allocationManager.createOperatorSets(avs, createSetParams);

        for (uint k; k < numOpSets; ++k) {
            OperatorSet memory opSet = OperatorSet(avs, createSetParams[k].operatorSetId);
            assertTrue(allocationManager.isOperatorSet(opSet), "should be operator set");
            IStrategy[] memory strategiesInSet = allocationManager.getStrategiesInOperatorSet(opSet);
            assertEq(strategiesInSet.length, numStrategies, "strategiesInSet length should be numStrategies");
            for (uint l; l < numStrategies; ++l) {
                assertTrue(
                    allocationManager.getStrategiesInOperatorSet(opSet)[l] == createSetParams[k].strategies[l], "should be strat of set"
                );
            }
        }

        assertEq(createSetParams.length, allocationManager.getOperatorSetCount(avs), "should be correct number of sets");
    }
}

contract AllocationManagerUnitTests_setAVSRegistrar is AllocationManagerUnitTests {
    function test_getAVSRegistrar() public {
        address randomAVS = random().Address();
        IAVSRegistrar avsRegistrar = allocationManager.getAVSRegistrar(randomAVS);
        assertEq(address(avsRegistrar), address(randomAVS), "AVS registrar should return default");
    }

    function testFuzz_setAVSRegistrar_Correctness(Randomness r) public rand(r) {
        address avs = r.Address();
        IAVSRegistrar avsRegistrar = IAVSRegistrar(defaultAVS);

        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AVSRegistrarSet(avs, avsRegistrar);

        cheats.prank(avs);
        allocationManager.setAVSRegistrar(avs, avsRegistrar);
        assertTrue(avsRegistrar == allocationManager.getAVSRegistrar(avs), "should be set");
    }
}

contract AllocationManagerUnitTests_updateAVSMetadataURI is AllocationManagerUnitTests {
    function test_updateAVSMetadataURI_Correctness() public {
        string memory newURI = "test";
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AVSMetadataURIUpdated(defaultAVS, newURI);
        cheats.prank(defaultAVS);
        allocationManager.updateAVSMetadataURI(defaultAVS, newURI);
    }
}

contract AllocationManagerUnitTests_getStrategyAllocations is AllocationManagerUnitTests {
    using ArrayLib for *;

    function testFuzz_getStrategyAllocations_Correctness(Randomness r) public rand(r) {
        AllocateParams[] memory allocateParams = r.AllocateParams(defaultAVS, 1, 1);
        CreateSetParams[] memory createSetParams = r.CreateSetParams(allocateParams);

        cheats.prank(defaultAVS);
        allocationManager.createOperatorSets(defaultAVS, createSetParams);

        cheats.startPrank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.stopPrank();

        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        (OperatorSet[] memory operatorSets,) = allocationManager.getStrategyAllocations(defaultOperator, allocateParams[0].strategies[0]);

        assertEq(operatorSets[0].avs, allocateParams[0].operatorSet.avs, "should be defaultAVS");
        assertEq(operatorSets[0].id, allocateParams[0].operatorSet.id, "should be defaultOperatorSet");

        _checkAllocationStorage({
            operator: defaultOperator,
            operatorSet: operatorSets[0],
            strategy: createSetParams[0].strategies[0],
            expectedAllocation: Allocation({currentMagnitude: allocateParams[0].newMagnitudes[0], pendingDiff: 0, effectBlock: 0}),
            expectedMagnitudes: Magnitudes({
                encumbered: allocateParams[0].newMagnitudes[0],
                max: WAD,
                allocatable: WAD - allocateParams[0].newMagnitudes[0]
            })
        });
    }
}

contract AllocationManagerUnitTests_getSlashableStake is AllocationManagerUnitTests {
    using SlashingLib for *;
    using ArrayLib for *;

    /**
     * Allocates half of magnitude for a single strategy to an operatorSet. Then allocates again. Slashes 50%
     * of the first allocation. Validates slashable stake at each step.
     */
    function test_allocate_onePendingAllocation(Randomness r) public rand(r) {
        // Generate allocation for `strategyMock`, we allocate half
        {
            AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, 5e17);
            cheats.prank(defaultOperator);
            allocationManager.modifyAllocations(defaultOperator, allocateParams);
            cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);
        }

        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: DEFAULT_OPERATOR_SHARES.mulWad(5e17)
        });

        // Allocate the other half
        AllocateParams[] memory allocateParams2 = _newAllocateParams(defaultOperatorSet, WAD);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams2);
        uint32 secondAllocEffectBlock = _defaultAllocEffectBlock();

        // Check minimum slashable stake remains the same
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: DEFAULT_OPERATOR_SHARES.mulWad(5e17)
        });

        // Check minimum slashable stake would not change even after the second allocation becomes effective
        // This is because the allocation is not effective yet & we're getting a MINIMUM
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: DEFAULT_OPERATOR_SHARES.mulWad(5e17),
            futureBlock: secondAllocEffectBlock + 1
        });

        // Check minimum slashable stake after the second allocation becomes effective
        cheats.roll(secondAllocEffectBlock);
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: DEFAULT_OPERATOR_SHARES
        });
    }

    /**
     * Allocates to `firstMod` magnitude and then deallocate to `secondMod` magnitude
     * Validates slashable stake at each step after allocation and deallocation
     */
    function testFuzz_allocate_deallocate_validateSlashableStake(Randomness r) public rand(r) {
        // Bound allocation and deallocation
        uint64 firstMod = r.Uint64(1, WAD);
        uint64 secondMod = r.Uint64(0, firstMod - 1);

        // Allocate magnitude to default registered set
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, firstMod);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // 1. Validate slashable stake.
        // This value should be 0 even at the effectBlock since its minimal slashable stake
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: 0,
            futureBlock: _defaultAllocEffectBlock()
        });

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // 2. Check slashable stake after allocation effect block
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: DEFAULT_OPERATOR_SHARES.mulWad(firstMod)
        });

        // Deallocate
        allocateParams = _newAllocateParams(defaultOperatorSet, secondMod);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        // 3. Check slashable stake after deallocation - should be same at current block
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: DEFAULT_OPERATOR_SHARES.mulWad(firstMod),
            futureBlock: uint32(block.number)
        });

        // 4. Check slashable stake at the deallocation effect block
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: DEFAULT_OPERATOR_SHARES.mulWad(secondMod),
            futureBlock: uint32(block.number + DEALLOCATION_DELAY + 1)
        });

        // Warp to deallocation effect block
        cheats.roll(block.number + DEALLOCATION_DELAY + 1);

        // 5. Check slashable stake at the deallocation effect block
        _checkSlashableStake({
            operatorSet: defaultOperatorSet,
            operator: defaultOperator,
            strategies: defaultStrategies,
            expectedSlashableStake: DEFAULT_OPERATOR_SHARES.mulWad(secondMod)
        });
    }

    /**
     * Allocates all of magnitude to a single strategy to an operatorSet.
     * Deallocate some portion. Finally, slash while deallocation is pending
     */
    function testFuzz_SlashWhileDeallocationPending(Randomness r) public rand(r) {
        // Initialize state
        AllocateParams[] memory allocateParams = r.AllocateParams(defaultAVS, 1, 1);
        AllocateParams[] memory deallocateParams = r.DeallocateParams(allocateParams);
        CreateSetParams[] memory createSetParams = r.CreateSetParams(allocateParams);
        RegisterParams memory registerParams = r.RegisterParams(allocateParams);
        SlashingParams memory slashingParams = r.SlashingParams(defaultOperator, allocateParams[0]);

        delegationManagerMock.setOperatorShares(defaultOperator, allocateParams[0].strategies[0], DEFAULT_OPERATOR_SHARES);

        cheats.prank(defaultAVS);
        allocationManager.createOperatorSets(defaultAVS, createSetParams);
        cheats.startPrank(defaultOperator);
        allocationManager.registerForOperatorSets(defaultOperator, registerParams);

        // Allocate
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        allocationManager.modifyAllocations(defaultOperator, deallocateParams);
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY + 1);
        cheats.stopPrank();

        // Check slashable stake after deallocation (still pending; no change)
        _checkSlashableStake({
            operatorSet: allocateParams[0].operatorSet,
            operator: defaultOperator,
            strategies: allocateParams[0].strategies,
            expectedSlashableStake: allocateParams[0].newMagnitudes[0]
        });

        // Check slashable stake after deallocation takes effect, before slashing
        _checkSlashableStake({
            operatorSet: allocateParams[0].operatorSet,
            operator: defaultOperator,
            strategies: allocateParams[0].strategies,
            expectedSlashableStake: deallocateParams[0].newMagnitudes[0],
            futureBlock: deallocationEffectBlock
        });

        uint magnitudeAllocated = allocateParams[0].newMagnitudes[0];
        uint magnitudeDeallocated = magnitudeAllocated - deallocateParams[0].newMagnitudes[0];
        uint magnitudeSlashed = magnitudeAllocated.mulWad(slashingParams.wadsToSlash[0]);
        uint expectedCurrentMagnitude = magnitudeAllocated - magnitudeSlashed;
        int128 expectedPendingDiff =
            -int128(uint128(magnitudeDeallocated - magnitudeDeallocated.mulWadRoundUp(slashingParams.wadsToSlash[0])));

        // Slash
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashingParams);

        // Check slashable stake after slash
        _checkSlashableStake({
            operatorSet: allocateParams[0].operatorSet,
            operator: defaultOperator,
            strategies: allocateParams[0].strategies,
            expectedSlashableStake: expectedCurrentMagnitude
        });

        // Check slashable stake after deallocation takes effect
        // Add 1 slippage for rounding down slashable stake
        _checkSlashableStake({
            operatorSet: allocateParams[0].operatorSet,
            operator: defaultOperator,
            strategies: allocateParams[0].strategies,
            expectedSlashableStake: expectedCurrentMagnitude - uint128(-expectedPendingDiff) - 1,
            futureBlock: deallocationEffectBlock
        });

        cheats.roll(deallocationEffectBlock);
        allocationManager.clearDeallocationQueue(defaultOperator, allocateParams[0].strategies, _maxNumToClear());

        // Check slashable stake after slash and deallocation
        // Add 1 slippage for rounding down slashable stake
        _checkSlashableStake({
            operatorSet: allocateParams[0].operatorSet,
            operator: defaultOperator,
            strategies: allocateParams[0].strategies,
            expectedSlashableStake: expectedCurrentMagnitude - uint128(-expectedPendingDiff) - 1
        });
    }

    function testFuzz_allocate_deregister(Randomness r) public rand(r) {
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        cheats.prank(defaultOperator);
        allocationManager.deregisterFromOperatorSets(
            DeregisterParams(defaultOperator, defaultOperatorSet.avs, defaultOperatorSet.id.toArrayU32())
        );

        // Check slashable stake right after deregistration
        _checkSlashableStake({
            operatorSet: allocateParams[0].operatorSet,
            operator: defaultOperator,
            strategies: allocateParams[0].strategies,
            expectedSlashableStake: DEFAULT_OPERATOR_SHARES.mulWad(5e17)
        });

        // Assert slashable stake after deregistration delay is 0
        cheats.roll(block.number + DEALLOCATION_DELAY + 1);
        _checkSlashableStake({
            operatorSet: allocateParams[0].operatorSet,
            operator: defaultOperator,
            strategies: allocateParams[0].strategies,
            expectedSlashableStake: 0
        });
    }
}

contract AllocationManagerUnitTests_isOperatorSlashable is AllocationManagerUnitTests {
    using SlashingLib for *;
    using ArrayLib for *;

    function test_registeredOperator() public view {
        assertEq(
            allocationManager.isOperatorSlashable(defaultOperator, defaultOperatorSet), true, "registered operator should be slashable"
        );
        assertEq(allocationManager.isMemberOfOperatorSet(defaultOperator, defaultOperatorSet), true, "operator should be member of set");
    }

    function test_deregisteredOperatorAndSlashable() public {
        // 1. deregister defaultOperator from defaultOperator set
        uint32 deregisterBlock = uint32(block.number);
        cheats.prank(defaultOperator);
        allocationManager.deregisterFromOperatorSets(
            DeregisterParams(defaultOperator, defaultOperatorSet.avs, defaultOperatorSet.id.toArrayU32())
        );
        assertEq(
            allocationManager.isMemberOfOperatorSet(defaultOperator, defaultOperatorSet), false, "operator should not be member of set"
        );

        // 2. assert operator is still slashable
        assertEq(
            allocationManager.isOperatorSlashable(defaultOperator, defaultOperatorSet), true, "deregistered operator should be slashable"
        );

        // 3. roll blocks forward to slashableUntil block assert still slashable
        cheats.roll(deregisterBlock + DEALLOCATION_DELAY);
        assertEq(
            allocationManager.isOperatorSlashable(defaultOperator, defaultOperatorSet), true, "deregistered operator should be slashable"
        );

        // 4. roll 1 block forward and assert not slashable
        cheats.roll(deregisterBlock + DEALLOCATION_DELAY + 1);
        assertEq(
            allocationManager.isOperatorSlashable(defaultOperator, defaultOperatorSet),
            false,
            "deregistered operator should not be slashable"
        );
    }
}

contract AllocationManagerUnitTests_getMaxMagnitudesAtBlock is AllocationManagerUnitTests {
    using ArrayLib for *;

    function testFuzz_correctness(Randomness r) public rand(r) {
        // Randomly allocate
        AllocateParams[] memory allocateParams = _randAllocateParams_DefaultOpSet();
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Slash first time
        SlashingParams memory slashParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: defaultOperatorSet.id,
            strategies: defaultStrategies,
            wadsToSlash: r.Uint256(0.1 ether, 0.99 ether).toArrayU256(),
            description: "test"
        });
        uint32 firstSlashBlock = uint32(block.number);
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashParams);
        uint64 maxMagnitudeAfterFirstSlash = allocationManager.getMaxMagnitude(defaultOperator, strategyMock);

        // Warp to random block
        uint32 secondSlashBlock = uint32(block.number + r.Uint32());
        cheats.roll(secondSlashBlock);

        // Slash second time
        slashParams.wadsToSlash = (r.Uint64(0.1 ether, 0.99 ether)).toArrayU256();
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(defaultAVS, slashParams);
        uint64 maxMagnitudeAfterSecondSlash = allocationManager.getMaxMagnitude(defaultOperator, strategyMock);

        // Warp to a block after the second slash
        cheats.roll(block.number + r.Uint32());

        // Validate get max magnitudes at block
        assertEq(
            allocationManager.getMaxMagnitudesAtBlock(defaultOperator, defaultStrategies, firstSlashBlock)[0],
            maxMagnitudeAfterFirstSlash,
            "max magnitude after first slash not correct"
        );

        assertEq(
            allocationManager.getMaxMagnitudesAtBlock(defaultOperator, defaultStrategies, secondSlashBlock)[0],
            maxMagnitudeAfterSecondSlash,
            "max magnitude after second slash not correct"
        );
    }
}

contract AllocationManagerUnitTests_getAllocatedStake is AllocationManagerUnitTests {
    using ArrayLib for *;
    using SlashingLib for *;

    /**
     * Allocates to `firstMod` magnitude and validates allocated stake is correct
     */
    function testFuzz_allocate(Randomness r) public rand(r) {
        // Generate allocation for `strategyMock`, we allocate half
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        uint[][] memory allocatedStake =
            allocationManager.getAllocatedStake(defaultOperatorSet, defaultOperator.toArray(), defaultStrategies);

        assertEq(allocatedStake[0][0], DEFAULT_OPERATOR_SHARES.mulWad(5e17), "allocated stake not correct");
    }

    /**
     * Allocates to `firstMod` magnitude and then deallocate to `secondMod` magnitude
     * Validates allocated stake is updated even after deallocation is cleared in storage
     */
    function testFuzz_allocate_deallocate(Randomness r) public rand(r) {
        // Bound allocation and deallocation
        uint64 firstMod = r.Uint64(1, WAD);
        uint64 secondMod = r.Uint64(0, firstMod - 1);

        // 1. Allocate magnitude to default registered set & warp to allocation complete block
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, firstMod);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // 2. Deallocate
        allocateParams = _newAllocateParams(defaultOperatorSet, secondMod);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);

        // 3. Check allocated stake right after deallocation - shouldn't be updated
        uint[][] memory allocatedStake =
            allocationManager.getAllocatedStake(defaultOperatorSet, defaultOperator.toArray(), defaultStrategies);
        assertEq(allocatedStake[0][0], DEFAULT_OPERATOR_SHARES.mulWad(firstMod), "allocated stake not correct");

        // 4. Check slashable stake at the deallocation effect block
        // Warp to deallocation effect block
        cheats.roll(block.number + DEALLOCATION_DELAY + 1);

        allocatedStake = allocationManager.getAllocatedStake(defaultOperatorSet, defaultOperator.toArray(), defaultStrategies);
        assertEq(allocatedStake[0][0], DEFAULT_OPERATOR_SHARES.mulWad(secondMod), "allocated stake not correct");
    }

    /**
     * Allocates to `firstMod` magnitude and then deregisters the operator.
     * Validates allocated stake is nonzero even after deregistration delay
     */
    function testFuzz_allocate_deregister(Randomness r) public rand(r) {
        // 1. Generate allocation for `strategyMock`, we allocate half
        AllocateParams[] memory allocateParams = _newAllocateParams(defaultOperatorSet, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(defaultOperator, allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // 2. Deregister from operator set & warp to deregistration effect block
        cheats.prank(defaultOperator);
        allocationManager.deregisterFromOperatorSets(
            DeregisterParams(defaultOperator, defaultOperatorSet.avs, defaultOperatorSet.id.toArrayU32())
        );
        cheats.roll(block.number + DEALLOCATION_DELAY + 1);

        // 3. Check allocated stake
        uint[][] memory allocatedStake =
            allocationManager.getAllocatedStake(defaultOperatorSet, defaultOperator.toArray(), defaultStrategies);
        assertEq(allocatedStake[0][0], DEFAULT_OPERATOR_SHARES.mulWad(5e17), "allocated stake should remain same after deregistration");
    }
}
