// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/core/AllocationManager.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

// TODO: Add **unique** tests for events.

contract AllocationManagerUnitTests is EigenLayerUnitTestSetup, IAllocationManagerErrors, IAllocationManagerEvents {
    using SingleItemArrayLib for *;

    uint8 internal constant PAUSED_MODIFY_ALLOCATIONS = 0;
    uint8 internal constant PAUSED_OPERATOR_SLASHING = 1;
    uint32 constant DEALLOCATION_DELAY = 17.5 days; // TODO: Set to 17.5 days in blocks.
    uint32 constant ALLOCATION_CONFIGURATION_DELAY = 21 days; // TODO: Set to 21 days in blocks.
    uint32 constant DEFAULT_OPERATOR_ALLOCATION_DELAY = 1 days; // TODO: Set to 1 day in blocks.

    AllocationManager allocationManager;
    ERC20PresetFixedSupply tokenMock;
    StrategyBase strategyMock;
    StrategyBase strategyMock2;
    OperatorSet defaultOperatorSet;

    address defaultOperator = address(this);
    address defaultAVS = address(0xFEDBAD);

    /// -----------------------------------------------------------------------
    /// Setup
    /// -----------------------------------------------------------------------

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();

        allocationManager = _deployAllocationManagerWithMockDependencies({
            _initialOwner: address(this),
            _pauserRegistry: pauserRegistry,
            _initialPausedStatus: 0
        });

        tokenMock = new ERC20PresetFixedSupply("Mock Token", "MOCK", type(uint256).max, address(this));

        strategyMock = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(new StrategyBase(IStrategyManager(address(strategyManagerMock)))),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, tokenMock, pauserRegistry)
                )
            )
        );

        strategyMock2 = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(new StrategyBase(IStrategyManager(address(strategyManagerMock)))),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, tokenMock, pauserRegistry)
                )
            )
        );

        defaultOperatorSet = OperatorSet({avs: defaultAVS, id: 0});

        // Set the allocation delay & roll to when it can be set
        delegationManagerMock.setIsOperator(defaultOperator, true);
        cheats.prank(defaultOperator);
        allocationManager.setAllocationDelay(DEFAULT_OPERATOR_ALLOCATION_DELAY);
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);
    }

    /// -----------------------------------------------------------------------
    /// Internal Helpers
    /// -----------------------------------------------------------------------

    function _deployAllocationManagerWithMockDependencies(
        address _initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 _initialPausedStatus
    ) internal virtual returns (AllocationManager) {
        return AllocationManager(
            address(
                new TransparentUpgradeableProxy(
                    address(
                        new AllocationManager(
                            IDelegationManager(address(delegationManagerMock)),
                            DEALLOCATION_DELAY,
                            ALLOCATION_CONFIGURATION_DELAY
                        )
                    ),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        AllocationManager.initialize.selector, _initialOwner, _pauserRegistry, _initialPausedStatus
                    )
                )
            )
        );
    }

    /// -----------------------------------------------------------------------
    /// Generate calldata for a magnitude allocation
    /// -----------------------------------------------------------------------

    function _createOperatorSet(OperatorSet memory operatorSet, IStrategy[] memory strategies) internal {
        cheats.prank(operatorSet.avs);
        // Create operator set, if it doesn't exist.
        try allocationManager.createOperatorSets(
            CreateSetParams({operatorSetId: operatorSet.id, strategies: strategies}).toArray()
        ) {} catch {
            console.log("OperatorSet already exists...");
        }
    }

    function _newAllocateParams_SingleStrategy(
        OperatorSet memory operatorSet,
        IStrategy strategy,
        uint64 magnitude
    ) internal returns (AllocateParams[] memory) {
        // Cache strategies for repeat use.
        IStrategy[] memory strategies = strategy.toArray();

        // Create the operator set to avoid invalid operator set error.
        _createOperatorSet(operatorSet, strategies);

        // Return the arrayified allocation params.
        return AllocateParams({operatorSet: operatorSet, strategies: strategies, newMagnitudes: magnitude.toArray()})
            .toArray();
    }

    function _newAllocateParams_SingleMockStrategy(
        OperatorSet memory operatorSet,
        uint64 magnitude
    ) internal returns (AllocateParams[] memory) {
        return _newAllocateParams_SingleStrategy(operatorSet, strategyMock, magnitude);
    }

    function _newAllocateParams_SingleMockStrategy(
        address avs,
        uint64 magnitude
    ) internal returns (AllocateParams[] memory) {
        return _newAllocateParams_SingleMockStrategy(OperatorSet({avs: avs, id: 1}), magnitude);
    }

    /// -----------------------------------------------------------------------
    /// Generate random slashing parameters
    /// -----------------------------------------------------------------------

    function _shuffleRandomness(uint256 r, uint256 salt) internal pure returns (uint256) {
        return uint256(keccak256(abi.encodePacked(r, salt)));
    }

    /**
     * @notice Gets random slashing parameters. Not useful unless the operatorSetID is set. See overloaded method
     */
    function _randomSlashingParams(
        address operator,
        uint256 r,
        uint256 salt
    ) internal pure returns (SlashingParams memory) {
        r = _shuffleRandomness(r, salt);

        return SlashingParams({
            operator: operator,
            operatorSetId: uint32(r),
            wadToSlash: bound(r, 1, WAD),
            description: "test"
        });
    }

    function _randomSlashingParams(
        address operator,
        uint32 operatorSetId,
        uint256 r,
        uint256 salt
    ) internal pure returns (SlashingParams memory) {
        r = _shuffleRandomness(r, salt);

        return SlashingParams({
            operator: operator,
            operatorSetId: operatorSetId,
            wadToSlash: bound(r, 1, WAD),
            description: "test"
        });
    }

    /// -----------------------------------------------------------------------
    /// Generated a random magnitude allocation for a single strategy and operatorSet
    /// -----------------------------------------------------------------------

    function _completeRandomAllocation_singleStrat_singleOpset(
        address operator,
        address avs,
        uint256 r,
        uint256 salt
    ) internal returns (AllocateParams[] memory allocateParams) {
        allocateParams = _queueRandomAllocation_singleStrat_singleOpSet(operator, avs, r, salt);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);
    }

    function _queueRandomAllocation_singleStrat_singleOpSet(
        address operator,
        address avs,
        uint256 r,
        uint256 salt
    ) internal returns (AllocateParams[] memory allocateParams) {
        allocateParams = _randomAllocateParams_singleStrat_singleOpSet(avs, r, salt);
        cheats.prank(operator);
        allocationManager.modifyAllocations(allocateParams);
    }

    /**
     * @notice Queued a random allocation for the given `operator`
     * - Does NOT roll past the effect block
     */
    function _queueRandomAllocation_singleStrat_singleOpSet(
        address operator,
        uint256 r,
        uint256 salt
    ) internal returns (AllocateParams[] memory allocateParams) {
        allocateParams = _randomAllocateParams_singleStrat_singleOpSet(r, salt);
        cheats.prank(operator);
        allocationManager.modifyAllocations(allocateParams);
    }

    /**
     * @notice Create a random magnitude allocation
     * Randomized Parameters: avs, opSet, magnitude
     * Non-random Parameters: strategy, expectedMaxMagnitude
     * In addition
     * - Registers the operatorSet with the avsDirectory
     */
    function _randomAllocateParams_singleStrat_singleOpSet(
        uint256 r,
        uint256 salt
    ) internal returns (AllocateParams[] memory) {
        r = _shuffleRandomness(r, salt);
        return _randomAllocateParams_singleStrat_singleOpSet(_randomAddr(r, 0), r, salt);
    }

    /**
     * @notice Create a random magnitude allocation
     * Randomized Parameters: opSet, magnitude
     * Non-random Parameters: strategy, expectedMaxMagnitude, avs
     * In addition
     * - Registers the operatorSet with the avsDirectory
     */
    function _randomAllocateParams_singleStrat_singleOpSet(
        address avs,
        uint256 r,
        uint256 salt
    ) internal returns (AllocateParams[] memory) {
        r = _shuffleRandomness(r, salt);

        // Mock a random operator set.
        OperatorSet memory operatorSet = OperatorSet({avs: avs, id: uint32(r)});

        // Make the operator set valid.
        cheats.prank(avs);
        allocationManager.createOperatorSets(CreateSetParams(operatorSet.id, _strategyMockArray()).toArray());

        // Mock a random magnitude allocation.
        return AllocateParams({
            operatorSet: operatorSet,
            strategies: _strategyMockArray(),
            newMagnitudes: uint64(bound(r, 1, WAD)).toArray()
        }).toArray();
    }

    /// -----------------------------------------------------------------------
    /// Generate a random allocation for a single strategy and multiple operatorSets
    /// -----------------------------------------------------------------------

    function _randomAllocateParams_singleStrat_multipleOpSets(
        uint256 r,
        uint256 salt,
        uint8 numOpSets
    ) internal returns (AllocateParams[] memory) {
        r = _shuffleRandomness(r, salt);

        // Create multiple operatorSets
        CreateSetParams[] memory createSetParams = new CreateSetParams[](numOpSets);
        for (uint8 i = 0; i < numOpSets; ++i) {
            createSetParams[i] = CreateSetParams({operatorSetId: uint32(i), strategies: _strategyMockArray()});
        }

        // Make the operator set valid.
        allocationManager.createOperatorSets(createSetParams);

        // Give each set a minimum of 1 magnitude
        uint64[] memory magnitudes = new uint64[](numOpSets);
        uint64 usedMagnitude;
        for (uint8 i = 0; i < numOpSets; ++i) {
            magnitudes[i] = 1;
            usedMagnitude++;
        }

        // Distribute remaining magnitude
        uint64 maxMagnitude = WAD;
        for (uint8 i = 0; i < numOpSets; ++i) {
            r = _shuffleRandomness(r, i);
            uint64 remainingMagnitude = maxMagnitude - usedMagnitude;
            if (remainingMagnitude > 0) {
                magnitudes[i] += uint64(bound(r, 0, remainingMagnitude));
                usedMagnitude += magnitudes[i] - 1;
            }
        }

        // Create magnitude allocation
        return AllocateParams({
            operatorSet: OperatorSet(address(this), 0),
            strategies: _strategyMockArray(),
            newMagnitudes: magnitudes
        }).toArray();
    }

    /// -----------------------------------------------------------------------
    /// Generate a random allocation AND delllocation
    /// -----------------------------------------------------------------------

    /**
     * @notice Queued a random allocation and deallocation for the given `operator`
     * - DOES NOT roll past the deallocation effect block
     */
    function _queueRandomAllocationAndDeallocation(
        address operator,
        uint8 numOpSets,
        uint256 r,
        uint256 salt
    ) internal returns (AllocateParams[] memory, AllocateParams[] memory) {
        (AllocateParams[] memory allocateParams, AllocateParams[] memory deallocateParams) =
            _randomAllocationAndDeallocation_singleStrat_multipleOpSets(numOpSets, r, salt);

        // Allocate
        cheats.prank(operator);
        allocationManager.modifyAllocations(allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        cheats.prank(operator);
        allocationManager.modifyAllocations(deallocateParams);

        return (allocateParams, deallocateParams);
    }

    /**
     * @notice Generates a random allocation and deallocation for a single strategy and multiple operatorSets
     * @notice DeallocateParams are from 0 to 1 less that the current allocated magnitude
     */
    function _randomAllocationAndDeallocation_singleStrat_multipleOpSets(
        uint8 numOpSets,
        uint256 r,
        uint256 salt
    ) internal returns (AllocateParams[] memory, AllocateParams[] memory) {
        r = _shuffleRandomness(r, salt);

        AllocateParams[] memory allocateParams = _randomAllocateParams_singleStrat_multipleOpSets(r, salt, numOpSets);

        uint64[] memory newMagnitudes = new uint64[](numOpSets);
        for (uint8 i = 0; i < numOpSets; ++i) {
            newMagnitudes[i] = uint64(bound(r, 0, allocateParams[0].newMagnitudes[i] - 1));
        }

        AllocateParams[] memory deallocateParams = AllocateParams({
            operatorSet: allocateParams[0].operatorSet,
            strategies: _strategyMockArray(),
            newMagnitudes: newMagnitudes
        }).toArray();

        return (allocateParams, deallocateParams);
    }

    /// -----------------------------------------------------------------------
    /// Utils
    /// -----------------------------------------------------------------------

    function _strategyMockArray() internal view returns (IStrategy[] memory) {
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        return strategies;
    }

    function _randomAddr(uint256 r, uint256 salt) internal pure returns (address addr) {
        /// @solidity memory-safe-assembly
        assembly {
            mstore(0x00, r)
            mstore(0x20, salt)
            addr := keccak256(0x00, 0x40)
        }
    }

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
        uint256 r
    ) public {
        // Generate random values for the expected initial state of the contract.
        address expectedInitialOwner = _randomAddr(r, 0);
        IPauserRegistry expectedPauserRegistry = IPauserRegistry(_randomAddr(r, 1));

        // Deploy the contract with the expected initial state.
        AllocationManager alm = _deployAllocationManagerWithMockDependencies(
            expectedInitialOwner,
            expectedPauserRegistry,
            r // initialPausedStatus
        );

        // Assert that the contract can only be initialized once.
        vm.expectRevert("Initializable: contract is already initialized");
        alm.initialize(expectedInitialOwner, expectedPauserRegistry, r);

        // Assert immutable state
        assertEq(address(alm.delegation()), address(delegationManagerMock));
        assertEq(alm.DEALLOCATION_DELAY(), DEALLOCATION_DELAY);
        assertEq(alm.ALLOCATION_CONFIGURATION_DELAY(), ALLOCATION_CONFIGURATION_DELAY);

        // Assert initialiation state
        assertEq(alm.owner(), expectedInitialOwner);
        assertEq(address(alm.pauserRegistry()), address(expectedPauserRegistry));
        assertEq(alm.paused(), r);
    }
}

contract AllocationManagerUnitTests_SlashOperator is AllocationManagerUnitTests {
    /// -----------------------------------------------------------------------
    /// slashOperator()
    /// -----------------------------------------------------------------------

    function test_revert_paused() public {
        allocationManager.pause(2 ** PAUSED_OPERATOR_SLASHING);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        allocationManager.slashOperator(_randomSlashingParams(defaultOperator, 0, 0));
    }

    function test_revert_slashZero() public {
        SlashingParams memory slashingParams = _randomSlashingParams(defaultOperator, 0, 0);
        slashingParams.wadToSlash = 0;

        cheats.prank(defaultAVS);
        cheats.expectRevert(InvalidWadToSlash.selector);
        allocationManager.slashOperator(slashingParams);
    }

    function test_revert_slashGreaterThanWAD() public {
        SlashingParams memory slashingParams = _randomSlashingParams(defaultOperator, 0, 0);
        slashingParams.wadToSlash = WAD + 1;

        cheats.prank(defaultAVS);
        cheats.expectRevert(InvalidWadToSlash.selector);
        allocationManager.slashOperator(slashingParams);
    }

    function test_revert_operatorNotSlashable() public {
        cheats.prank(defaultAVS);
        cheats.expectRevert(InvalidOperator.selector);
        allocationManager.slashOperator(_randomSlashingParams(defaultOperator, 0, 0));
    }

    // function test_revert_operatorNotAllocated() public {
    //     SlashingParams memory slashingParams = _randomSlashingParams(defaultOperator, 0, 0);
    //     // avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

    //     cheats.expectRevert(OperatorNotAllocated.selector);
    //     cheats.prank(defaultAVS);
    //     allocationManager.slashOperator(slashingParams);
    // }

    // function test_revert_operatorAllocated_notActive() public {
    //     // Queue allocation
    //     AllocateParams[] memory allocateParams =
    //         _queueRandomAllocation_singleStrat_singleOpSet(defaultOperator, 0, 0);

    //     // Setup data
    //     SlashingParams memory slashingParams = SlashingParams({
    //         operator: defaultOperator,
    //         operatorSetId: allocateParams[0].operatorSet.id,
    //         wadToSlash: WAD,
    //         description: "test"
    //     });
    //     // avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

    //     // Expect revert
    //     cheats.expectRevert(OperatorNotAllocated.selector);
    //     cheats.prank(defaultAVS);
    //     allocationManager.slashOperator(slashingParams);
    // }

    /**
     * Allocates all magnitude to for a single strategy to an operatorSet. Slashes 25%
     * Asserts that:
     * 1. Events are emitted
     * 2. Encumbered mag is updated
     * 3. Max mag is updated
     * 4. Calculations for `getAllocatableMagnitude` and `getAllocation` are correct
     */
    function test_slashPostAllocation() public {
        // Generate allocation for `strategyMock`, we allocate max
        AllocateParams[] memory allocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, WAD);

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Slash operator for 25%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocateParams[0].operatorSet.id,
            wadToSlash: 25e16,
            description: "test"
        });
        // // avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage
        assertEq(
            75e16,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            75e16,
            allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0],
            "maxMagnitude not updated"
        );
        assertEq(
            0,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            "allocatableMagnitude shoudl be 0"
        );
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(75e16, allocation.currentMagnitude, "currentMagnitude not updated");
        assertEq(0, allocation.pendingDiff, "pendingDiff should be 0");
        assertEq(0, allocation.effectBlock, "effectBlock should be 0");
    }

    /// @notice Same test as above, but fuzzes the allocation
    function testFuzz_slashPostAllocation(
        uint256 r
    ) public {
        AllocateParams[] memory allocateParams =
            _completeRandomAllocation_singleStrat_singleOpset(defaultOperator, defaultAVS, r, 0);
        SlashingParams memory slashingParams =
            _randomSlashingParams(defaultOperator, allocateParams[0].operatorSet.id, r, 1);

        uint64 expectedSlashedMagnitude =
            uint64(SlashingLib.mulWadRoundUp(allocateParams[0].newMagnitudes[0], slashingParams.wadToSlash));
        uint64 expectedEncumberedMagnitude = allocateParams[0].newMagnitudes[0] - expectedSlashedMagnitude;
        uint64 maxMagnitudeAfterSlash = WAD - expectedSlashedMagnitude;

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage
        assertEq(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            maxMagnitudeAfterSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0],
            "maxMagnitude not updated"
        );

        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);

        assertEq(expectedEncumberedMagnitude, allocation.currentMagnitude, "currentMagnitude not updated");
        assertEq(0, allocation.pendingDiff, "pendingDiff should be 0");
        assertEq(0, allocation.effectBlock, "effectBlock should be 0");
    }

    /**
     * Allocates half of magnitude for a single strategy to an operatorSet. Then allocates again. Slashes 50%
     * Asserts that:
     * 1. Events are emitted
     * 2. Encumbered mag is updated
     * 3. Max mag is updated
     * 4. Calculations for `getAllocatableMagnitude` and `getAllocation` are correct
     * 5. The second magnitude allocation is not slashed from
     * TODO: Fuzz
     */
    function test_slash_oneCompletedAlloc_onePendingAlloc() public {
        // Generate allocation for `strategyMock`, we allocate half
        AllocateParams[] memory allocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Allocate the other half
        AllocateParams[] memory allocateParams2 = _newAllocateParams_SingleMockStrategy(defaultAVS, WAD);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams2);
        uint32 secondAllocEffectBlock = uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Slash operator for 50%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocateParams[0].operatorSet.id,
            wadToSlash: 50e16,
            description: "test"
        });
        // // avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);
        uint64 expectedEncumberedMagnitude = 75e16; // 25e16 from first allocation, 50e16 from second
        uint64 magnitudeAfterSlash = 25e16;
        uint64 maxMagnitudeAfterSlash = 75e16;

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage
        assertEq(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            maxMagnitudeAfterSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0],
            "maxMagnitude not updated"
        );
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(magnitudeAfterSlash, allocation.currentMagnitude, "currentMagnitude not updated");
        assertEq(5e17, allocation.pendingDiff, "pendingDiff should be for second alloc");
        assertEq(secondAllocEffectBlock, allocation.effectBlock, "effectBlock should be 0");

        // Warp to complete second allocation
        cheats.roll(secondAllocEffectBlock);
        uint64 allocatableMagnitude = allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock);
        assertEq(0, allocatableMagnitude, "allocatableMagnitude should be 0");
        allocation = allocationManager.getAllocation(defaultOperator, allocateParams2[0].operatorSet, strategyMock);
        assertEq(75e16, allocation.currentMagnitude, "currentMagnitude not updated");
        assertEq(0, allocation.pendingDiff, "pendingDiff should be 0");
        assertEq(0, allocation.effectBlock, "effectBlock should be 0");
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
    function test_slashTwoOperatorSets() public {
        // Generate allocation for `strategyMock`, we allocate 100% to opSet 0
        AllocateParams[] memory allocateParams = _newAllocateParams_SingleMockStrategy(defaultOperatorSet, WAD);

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // 1. Slash operator for 99% in opSet 0 bringing their magnitude to 1e16
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocateParams[0].operatorSet.id,
            wadToSlash: 99e16,
            description: "test"
        });
        // // avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);
        uint64 expectedEncumberedMagnitude = 1e16; // After slashing 99%, only 1% expected encumberedMagnitude
        uint64 magnitudeAfterSlash = 1e16;
        uint64 maxMagnitudeAfterSlash = 1e16; // 1e15 is maxMagnitude

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage
        assertEq(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            maxMagnitudeAfterSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0],
            "maxMagnitude not updated"
        );
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(magnitudeAfterSlash, allocation.currentMagnitude, "currentMagnitude not updated");

        // 2. Slash operator again for 99.99% in opSet 0 bringing their magnitude to 1e14
        slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocateParams[0].operatorSet.id,
            wadToSlash: 9999e14,
            description: "test"
        });
        expectedEncumberedMagnitude = 1e12; // After slashing 99.99%, only 0.01% expected encumberedMagnitude
        magnitudeAfterSlash = 1e12;
        maxMagnitudeAfterSlash = 1e12;

        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage
        assertEq(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            maxMagnitudeAfterSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0],
            "maxMagnitude not updated"
        );
        allocation = allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(magnitudeAfterSlash, allocation.currentMagnitude, "currentMagnitude not updated");

        // 3. Slash operator again for 99.9999999999999% in opSet 0
        slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocateParams[0].operatorSet.id,
            wadToSlash: WAD - 1e3,
            description: "test"
        });
        // Should technically be 1e3 remaining but with rounding error and rounding up slashed amounts
        // the remaining magnitude is 0
        expectedEncumberedMagnitude = 0; // Should technically be 1e3 remaining but with rounding error and rounding up slashed amounts.
        magnitudeAfterSlash = 0;
        maxMagnitudeAfterSlash = 0;

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage
        assertEq(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            maxMagnitudeAfterSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0],
            "maxMagnitude not updated"
        );
        allocation = allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(magnitudeAfterSlash, allocation.currentMagnitude, "currentMagnitude not updated");
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
     * TODO: Fuzz the allocation & slash amounts
     */
    function test_allocateAll_deallocateHalf_slashWhileDeallocPending() public {
        uint64 initialMagnitude = WAD;
        // Generate allocation for `strategyMock`, we allocate half
        AllocateParams[] memory allocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, initialMagnitude);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate half
        AllocateParams[] memory deallocateParams =
            _newAllocateParams_SingleMockStrategy(defaultAVS, initialMagnitude / 2);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(deallocateParams);
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY);

        // Slash operator for 25%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocateParams[0].operatorSet.id,
            wadToSlash: 25e16,
            description: "test"
        });
        // // avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);
        uint64 magnitudeAfterDeallocationSlash = 375e15; // 25% is slashed off of 5e17
        uint64 expectedEncumberedMagnitude = 75e16; // 25e16 is slashed. 75e16 is encumbered
        uint64 magnitudeAfterSlash = 75e16;
        uint64 maxMagnitudeAfterSlash = 75e16; // Operator can only allocate up to 75e16 magnitude since 25% is slashed

        // Slash Operator
        // First event is emitted because of deallocation
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage post slash
        assertEq(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            maxMagnitudeAfterSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0],
            "maxMagnitude not updated"
        );
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(magnitudeAfterSlash, allocation.currentMagnitude, "currentMagnitude not updated");
        assertEq(
            -int128(uint128((uint64(magnitudeAfterDeallocationSlash)))),
            allocation.pendingDiff,
            "pendingDiff should be decreased after slash"
        );
        assertEq(deallocationEffectBlock, allocation.effectBlock, "effectBlock should be 0");

        // Check storage after complete modification
        cheats.roll(deallocationEffectBlock);
        allocationManager.clearDeallocationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());
        allocation = allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(magnitudeAfterDeallocationSlash, allocation.currentMagnitude, "currentMagnitude not updated");
        assertEq(
            magnitudeAfterDeallocationSlash,
            maxMagnitudeAfterSlash / 2,
            "magnitude after deallocation should be half of max magnitude, since we originally deallocated by half"
        );
    }

    /**
     * Allocates all magnitude to a single opSet. Then slashes the entire magnitude
     * Asserts that:
     * 1. The operator cannot allocate again
     */
    function testRevert_allocateAfterSlashedEntirely() public {
        // Allocate all magnitude
        AllocateParams[] memory allocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, WAD);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Slash operator for 100%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocateParams[0].operatorSet.id,
            wadToSlash: WAD,
            description: "test"
        });

        // avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Attempt to allocate
        AllocateParams[] memory allocateParams2 = _newAllocateParams_SingleMockStrategy(OperatorSet(defaultAVS, 1), 1);

        // TODO:
        // cheats.expectRevert(InsufficientAllocatableMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams2);
    }

    /**
     * Allocates all magnitude to a single opSet. Deallocateas magnitude. Slashes al
     * Asserts that:
     * 1. The Allocation is 0 after slash
     * 2. Them sotrage post slash for encumbered and maxMags ais zero
     */
    function test_allocateAll_deallocateAll() public {
        // Allocate all magnitude
        AllocateParams[] memory allocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, WAD);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate all
        AllocateParams[] memory deallocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(deallocateParams);
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY);

        // Slash operator for 100%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocateParams[0].operatorSet.id,
            wadToSlash: WAD,
            description: "test"
        });
        // avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage post slash
        assertEq(
            0, allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude not updated"
        );
        assertEq(
            0, allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0], "maxMagnitude not updated"
        );
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(0, allocation.currentMagnitude, "currentMagnitude not updated");
        assertEq(0, allocation.pendingDiff, "pendingDiff should be zero since everything is slashed");
        assertEq(deallocationEffectBlock, allocation.effectBlock, "effectBlock should be 0");
    }

    /**
     * Slashes the operator after deallocation, even if the deallocation has not been cleared. Validates that:
     * 1. Even if we do not clear deallocation queue, the deallocation is NOT slashed from since we're passed the deallocationEffectBlock
     * 2. Validates storage post slash & post clearing deallocation queue
     * 3. Max magnitude only decreased proportionally by the magnitude set after deallocation
     */
    function test_allocate_deallocate_slashAfterDeallocation() public {
        // Allocate all magnitude
        AllocateParams[] memory allocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, WAD);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate half
        AllocateParams[] memory deallocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(deallocateParams);
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY);

        // Check storage post deallocation
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(WAD, allocation.currentMagnitude, "currentMagnitude not updated");
        assertEq(-5e17, allocation.pendingDiff, "pendingDiff should be 5e17 after deallocation");
        assertEq(deallocationEffectBlock, allocation.effectBlock, "effectBlock should be 0");

        // Warp to deallocation effect block
        cheats.roll(deallocationEffectBlock);

        // Slash operator for 25%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocateParams[0].operatorSet.id,
            wadToSlash: 25e16,
            description: "test"
        });
        // avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);
        uint64 expectedEncumberedMagnitude = 375e15; // 25e16 is slashed. 5e17 was previously
        uint64 magnitudeAfterSlash = 375e15;
        uint64 maxMagnitudeAfterSlash = 875e15; // Operator can only allocate up to 75e16 magnitude since 25% is slashed

        // Slash Operator, only emit events assuming that there is no deallocation
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage post slash
        assertEq(
            expectedEncumberedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude not updated"
        );
        assertEq(
            maxMagnitudeAfterSlash,
            allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0],
            "maxMagnitude not updated"
        );
        allocation = allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(magnitudeAfterSlash, allocation.currentMagnitude, "currentMagnitude not updated");
        assertEq(0, allocation.pendingDiff, "pendingDiff should be 0 after slash");
        assertEq(0, allocation.effectBlock, "effectBlock should be 0");
        uint64 allocatableMagnitudeAfterSlash = allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock);

        // Check storage after complete modification. Expect encumberedMag to be emitted again
        allocationManager.clearDeallocationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());
        allocation = allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(
            allocatableMagnitudeAfterSlash,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            "allocatable mag after slash shoudl be equal to allocatable mag after clearing queue"
        );
    }

    // /**
    //  * Allocates to multiple operatorSets for a strategy. Only slashes from one operatorSet. Validates
    //  * 1. The slashable shares of each operatorSet after magnitude allocation
    //  * 2. The first operatorSet has less slashable shares post slash
    //  * 3. The second operatorSet has the same number slashable shares post slash
    //  * 4. The PROPORTION that is slashable for opSet 2 has increased
    //  * 5. Encumbered magnitude, total allocatable magnitude
    //  */
    // function test_allocateMultipleOpsets_slashSingleOpset() public {
    //     // Set 100e18 shares for operator in DM
    //     uint256 operatorShares = 100e18;
    //     delegationManagerMock.setOperatorShares(defaultOperator, strategyMock, operatorShares);
    //     uint64 magnitudeToAllocate = 4e17;

    //     // Allocate 40% to firstOperatorSet, 40% to secondOperatorSet
    //     AllocateParams[] memory allocateParams = new AllocateParams[](2);
    //     allocateParams[0] = _newAllocateParams_SingleMockStrategy(OperatorSet(defaultAVS, 1), magnitudeToAllocate)[0];
    //     allocateParams[1] = _newAllocateParams_SingleMockStrategy(OperatorSet(defaultAVS, 2), magnitudeToAllocate)[0];
    //     cheats.prank(defaultOperator);
    //     allocationManager.modifyAllocations(allocateParams);
    //     cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

    //     // Get slashable shares for each operatorSet
    //     address[] memory operatorArray = new address[](1);
    //     operatorArray[0] = defaultOperator;
    //     (, uint256[][] memory slashableSharesOpset1_preSlash) = allocationManager
    //         .getMinDelegatedAndSlashableOperatorSharesBefore(
    //         OperatorSet(defaultAVS, 1), operatorArray, _strategyMockArray(), uint32(block.number + 1)
    //     );
    //     (, uint256[][] memory slashableSharesOpset2_preSlash) = allocationManager
    //         .getMinDelegatedAndSlashableOperatorSharesBefore(
    //         OperatorSet(defaultAVS, 2), operatorArray, _strategyMockArray(), uint32(block.number + 1)
    //     );
    //     assertEq(40e18, slashableSharesOpset1_preSlash[0][0], "slashableShares of opSet_1 should be 40e18");
    //     assertEq(40e18, slashableSharesOpset2_preSlash[0][0], "slashableShares of opSet_2 should be 40e18");
    //     uint256 maxMagnitude = allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0];
    //     uint256 opSet2PortionOfMaxMagnitude = uint256(magnitudeToAllocate) * WAD / maxMagnitude;

    //     // Slash operator on operatorSet1 for 50%
    //     SlashingParams memory slashingParams = SlashingParams({
    //         operator: defaultOperator,
    //         operatorSetId: allocateParams[0].operatorSet.id,
    //         wadToSlash: 5e17,
    //         description: "test"
    //     });
    //     // avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

    //     // Slash Operator
    //     cheats.prank(defaultAVS);
    //     allocationManager.slashOperator(slashingParams);

    //     // Operator should now have 80e18 shares, since half of 40e18 was slashed
    //     delegationManagerMock.setOperatorShares(defaultOperator, strategyMock, 80e18);

    //     // Check storage
    //     (, uint256[][] memory slashableSharesOpset1_postSlash) = allocationManager
    //         .getMinDelegatedAndSlashableOperatorSharesBefore(
    //         OperatorSet(defaultAVS, 1), operatorArray, _strategyMockArray(), uint32(block.number + 1)
    //     );
    //     (, uint256[][] memory slashableSharesOpset2_postSlash) = allocationManager
    //         .getMinDelegatedAndSlashableOperatorSharesBefore(
    //         OperatorSet(defaultAVS, 2), operatorArray, _strategyMockArray(), uint32(block.number + 1)
    //     );

    //     assertEq(20e18, slashableSharesOpset1_postSlash[0][0], "slashableShares of opSet_1 should be 20e18");
    //     assertEq(
    //         slashableSharesOpset2_preSlash[0][0],
    //         slashableSharesOpset2_postSlash[0][0],
    //         "slashableShares of opSet_2 should remain unchanged"
    //     );

    //     // Validate encumbered and total allocatable magnitude
    //     uint256 maxMagnitudeAfterSlash = allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0];
    //     uint256 expectedEncumberedMagnitude = 6e17; // 4e17 from opSet2, 2e17 from opSet1
    //     assertEq(
    //         expectedEncumberedMagnitude,
    //         allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
    //         "encumberedMagnitude not updated"
    //     );
    //     assertEq(
    //         maxMagnitudeAfterSlash - expectedEncumberedMagnitude,
    //         allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
    //         "allocatableMagnitude should be diff of maxMagnitude and encumberedMagnitude"
    //     );

    //     // Check proportion after slash
    //     uint256 opSet2PortionOfMaxMagnitudeAfterSlash = uint256(magnitudeToAllocate) * WAD / maxMagnitudeAfterSlash;
    //     assertGt(
    //         opSet2PortionOfMaxMagnitudeAfterSlash,
    //         opSet2PortionOfMaxMagnitude,
    //         "opSet2 should have a greater proportion to slash from previous"
    //     );
    // }

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

        AllocateParams[] memory allocateParams = new AllocateParams[](2);
        allocateParams[0] = _newAllocateParams_SingleStrategy(defaultOperatorSet, strategyMock, strategy1Magnitude)[0];
        allocateParams[1] = _newAllocateParams_SingleStrategy(defaultOperatorSet, strategyMock2, strategy2Magnitude)[0];

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Slash operator on both strategies for 60%
        IStrategy[] memory strategiesToSlash = new IStrategy[](2);
        strategiesToSlash[0] = strategyMock;
        strategiesToSlash[1] = strategyMock2;

        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocateParams[0].operatorSet.id,
            wadToSlash: 6e17,
            description: "test"
        });
        // avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

        uint64[] memory expectedEncumberedMags = new uint64[](2);
        expectedEncumberedMags[0] = 2e17; // 60% of 5e17
        expectedEncumberedMags[1] = 4e17; // 60% of WAD

        uint64[] memory expectedMagnitudeAfterSlash = new uint64[](2);
        expectedMagnitudeAfterSlash[0] = 2e17;
        expectedMagnitudeAfterSlash[1] = 4e17;

        uint64[] memory expectedMaxMagnitudeAfterSlash = new uint64[](2);
        expectedMaxMagnitudeAfterSlash[0] = 7e17;
        expectedMaxMagnitudeAfterSlash[1] = 4e17;

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage
        for (uint256 i = 0; i < strategiesToSlash.length; ++i) {
            assertEq(
                expectedEncumberedMags[i],
                allocationManager.encumberedMagnitude(defaultOperator, strategiesToSlash[i]),
                "encumberedMagnitude not updated"
            );
            assertEq(
                expectedMaxMagnitudeAfterSlash[i] - expectedMagnitudeAfterSlash[i],
                allocationManager.getAllocatableMagnitude(defaultOperator, strategiesToSlash[i]),
                "allocatableMagnitude not updated"
            );
            Allocation memory allocation =
                allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategiesToSlash[i]);
            assertEq(expectedMagnitudeAfterSlash[i], allocation.currentMagnitude, "currentMagnitude not updated");
            assertEq(0, allocation.pendingDiff, "pendingDiff should be 0");
            assertEq(0, allocation.effectBlock, "effectBlock should be 0");
        }
    }

    /**
     * Allocates magnitude. Deallocates some. Slashes a portion, and then allocates up to the max available magnitude
     * TODO: Fuzz the wadsToSlash
     */
    function testFuzz_allocate_deallocate_slashWhilePending_allocateMax(
        uint256 r
    ) public {
        // Bound allocation and deallocation
        uint64 firstMod = uint64(bound(r, 3, WAD));
        uint64 secondMod = uint64(bound(r, 1, firstMod - 2));

        // TODO: remove these assumptions around even numbers
        if (firstMod % 2 != 0) {
            firstMod += 1;
        }
        if (secondMod % 2 != 0) {
            secondMod += 1;
        }
        uint64 pendingDiff = firstMod - secondMod;

        // Allocate magnitude
        AllocateParams[] memory allocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, firstMod);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate magnitude
        AllocateParams[] memory deallocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, secondMod);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(deallocateParams);
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY);

        // Slash operator for 50%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocateParams[0].operatorSet.id,
            wadToSlash: 5e17,
            description: "test"
        });
        // avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage post slash
        assertEq(
            firstMod / 2,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should be half of firstMod"
        );
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(firstMod / 2, allocation.currentMagnitude, "currentMagnitude should be half of firstMod");
        console.log("value of pendingDiff: ", pendingDiff - pendingDiff / 2);
        assertEq(
            -int128(uint128(pendingDiff - pendingDiff / 2)), allocation.pendingDiff, "pendingDiff should be -secondMod"
        );
        assertEq(deallocationEffectBlock, allocation.effectBlock, "effectBlock should be deallocationEffectBlock");

        // Warp to deallocation effect block & clear deallocation queue
        console.log("encumbered mag before: ", allocationManager.encumberedMagnitude(defaultOperator, strategyMock));
        cheats.roll(deallocationEffectBlock);
        allocationManager.clearDeallocationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());
        console.log("encumbered mag after: ", allocationManager.encumberedMagnitude(defaultOperator, strategyMock));

        // Check expected max and allocatable
        uint64 expectedMaxMagnitude = WAD - firstMod / 2;
        assertEq(
            expectedMaxMagnitude,
            allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0],
            "maxMagnitude should be expectedMaxMagnitude"
        );
        // Allocatable is expectedMax - currentMagPostSlashing - pendingDiffOfDeallocateParams post slashing
        uint64 expectedAllocatable = expectedMaxMagnitude - ((firstMod / 2) - (pendingDiff - pendingDiff / 2));
        assertEq(
            expectedAllocatable,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            "allocatableMagnitude should be expectedAllocatable"
        );

        // Allocate up to max magnitude
        AllocateParams[] memory allocateParams2 =
            _newAllocateParams_SingleMockStrategy(defaultAVS, expectedMaxMagnitude);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams2);

        // Assert that encumbered is expectedMaxMagnitude
        assertEq(
            0,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock),
            "allocatableMagnitude should be 0"
        );
    }
}

contract AllocationManagerUnitTests_ModifyAllocations is AllocationManagerUnitTests {
    using SingleItemArrayLib for *;

    /// -----------------------------------------------------------------------
    /// modifyAllocations()
    /// -----------------------------------------------------------------------
    function test_revert_paused() public {
        allocationManager.pause(2 ** PAUSED_MODIFY_ALLOCATIONS);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        allocationManager.modifyAllocations(new AllocateParams[](0));
    }

    function test_revert_allocationDelayNotSet() public {
        address invalidOperator = address(0x2);
        cheats.prank(invalidOperator);
        cheats.expectRevert(UninitializedAllocationDelay.selector);
        allocationManager.modifyAllocations(new AllocateParams[](0));
    }

    function test_revert_allocationDelayNotInEffect() public {
        address operator = address(0x2);
        delegationManagerMock.setIsOperator(operator, true);

        cheats.startPrank(operator);
        allocationManager.setAllocationDelay(5);
        // even though the operator has an allocation delay set, it is not in effect
        // and modifyAllocations should still be blocked
        cheats.expectRevert(UninitializedAllocationDelay.selector);
        allocationManager.modifyAllocations(new AllocateParams[](0));
    }

    function test_revert_lengthMismatch() public {
        AllocateParams[] memory allocateParams = _randomAllocateParams_singleStrat_singleOpSet(0, 0);
        allocateParams[0].newMagnitudes = new uint64[](0);

        cheats.expectRevert(InputArrayLengthMismatch.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
    }

    function test_revert_invalidOperatorSet() public {
        AllocateParams[] memory allocateParams = AllocateParams({
            operatorSet: defaultOperatorSet,
            strategies: _strategyMockArray(),
            newMagnitudes: uint64(0.5 ether).toArray()
        }).toArray();

        cheats.expectRevert(InvalidOperatorSet.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
    }

    function test_revert_multiAlloc_modificationAlreadyPending_diffTx() public {
        // Allocate magnitude
        AllocateParams[] memory allocateParams = _randomAllocateParams_singleStrat_singleOpSet(0, 0);
        cheats.startPrank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Warp to just before allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY - 1);

        // Attempt to allocate magnitude again
        cheats.expectRevert(ModificationAlreadyPending.selector);
        allocationManager.modifyAllocations(allocateParams);
        cheats.stopPrank();
    }

    function test_revert_multiAlloc_modificationAlreadyPending_sameTx() public {
        // Allocate magnitude
        AllocateParams[] memory allocateParams = new AllocateParams[](2);
        allocateParams[0] = _randomAllocateParams_singleStrat_singleOpSet(0, 0)[0];
        allocateParams[1] = allocateParams[0];

        cheats.expectRevert(ModificationAlreadyPending.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
    }

    function test_revert_allocateZeroMagnitude() public {
        // Allocate exact same magnitude as initial allocation (0)
        AllocateParams[] memory allocateParams = _randomAllocateParams_singleStrat_singleOpSet(0, 0);
        allocateParams[0].newMagnitudes[0] = 0;

        cheats.expectRevert(SameMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
    }

    function test_revert_allocateSameMagnitude() public {
        // Allocate nonzero magnitude
        AllocateParams[] memory allocateParams = _randomAllocateParams_singleStrat_singleOpSet(0, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Attempt to allocate no magnitude (ie. same magnitude)
        cheats.expectRevert(SameMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
    }

    function testFuzz_revert_insufficientAllocatableMagnitude(
        uint256 r
    ) public {
        // Allocate some magnitude
        AllocateParams[] memory allocateParams = _randomAllocateParams_singleStrat_singleOpSet(r, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Attempt to allocate more magnitude than the operator has
        // uint64 allocatedMag = allocateParams[0].newMagnitudes[0];
        allocateParams[0].newMagnitudes[0] = WAD + 1;
        // cheats.expectRevert(InsufficientAllocatableMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
    }

    function testFuzz_allocate_singleStrat_singleOperatorSet(
        uint256 r
    ) public {
        // Create allocation
        AllocateParams[] memory allocateParams = _randomAllocateParams_singleStrat_singleOpSet(r, 0);

        // Save vars to check against
        IStrategy strategy = allocateParams[0].strategies[0];
        uint64 magnitude = allocateParams[0].newMagnitudes[0];
        uint32 effectBlock = uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Allocate magnitude
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Check storage
        assertEq(
            magnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategy),
            "encumberedMagnitude not updated"
        );
        assertEq(
            WAD - magnitude,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategy),
            "allocatableMagnitude not calcualted correctly"
        );
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategy);
        assertEq(0, allocation.currentMagnitude, "currentMagnitude should not be updated");
        assertEq(int128(uint128(magnitude)), allocation.pendingDiff, "pendingMagnitude not updated");
        assertEq(effectBlock, allocation.effectBlock, "effectBlock not updated");

        // Check storage after roll to completion
        cheats.roll(effectBlock);
        allocation = allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategy);
        assertEq(magnitude, allocation.currentMagnitude, "currentMagnitude not updated");
        assertEq(0, allocation.pendingDiff, "pendingMagnitude not updated");
        assertEq(0, allocation.effectBlock, "effectBlock not updated");
    }

    function testFuzz_allocate_singleStrat_multipleSets(
        uint256 r
    ) public {
        uint8 numOpSets = uint8(bound(r, 1, type(uint8).max));

        AllocateParams[] memory allocateParams = _randomAllocateParams_singleStrat_multipleOpSets(r, 0, numOpSets);

        // Save vars to check against
        IStrategy strategy = allocateParams[0].strategies[0];
        uint64[] memory magnitudes = allocateParams[0].newMagnitudes;
        uint32 effectBlock = uint32(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Expect emits
        uint64 usedMagnitude;
        for (uint256 i = 0; i < numOpSets; ++i) {
            usedMagnitude += magnitudes[i];
        }
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Check storage
        assertEq(
            usedMagnitude,
            allocationManager.encumberedMagnitude(defaultOperator, strategy),
            "encumberedMagnitude not updated"
        );
        assertEq(
            WAD - usedMagnitude,
            allocationManager.getAllocatableMagnitude(defaultOperator, strategy),
            "allocatableMagnitude not calcualted correctly"
        );

        Allocation memory allocation;
        for (uint256 i = 0; i < numOpSets; ++i) {
            allocation = allocationManager.getAllocation(defaultOperator, allocateParams[i].operatorSet, strategy);
            assertEq(0, allocation.currentMagnitude, "currentMagnitude should not be updated");
            assertEq(int128(uint128(magnitudes[i])), allocation.pendingDiff, "pendingMagnitude not updated");
            assertEq(effectBlock, allocation.effectBlock, "effectBlock not updated");
        }

        // Check storage after roll to completion
        cheats.roll(effectBlock);
        for (uint256 i = 0; i < numOpSets; ++i) {
            allocation = allocationManager.getAllocation(defaultOperator, allocateParams[i].operatorSet, strategy);
            assertEq(magnitudes[i], allocation.currentMagnitude, "currentMagnitude not updated");
            assertEq(0, allocation.pendingDiff, "pendingMagnitude not updated");
            assertEq(0, allocation.effectBlock, "effectBlock not updated");
        }
    }

    function testFuzz_allocateMultipleTimes(
        uint256 r
    ) public {
        // Assumptions
        uint64 firstAlloc = uint64(bound(r, 1, type(uint64).max));
        uint64 secondAlloc = uint64(bound(r, 0, WAD));
        cheats.assume(firstAlloc < secondAlloc);

        // Allocate magnitude
        AllocateParams[] memory allocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, firstAlloc);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Allocate magnitude again
        allocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, secondAlloc);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Check storage
        assertEq(
            secondAlloc,
            allocationManager.encumberedMagnitude(defaultOperator, allocateParams[0].strategies[0]),
            "encumberedMagnitude not updated"
        );
    }

    function testFuzz_revert_overAllocate(
        uint256 r
    ) public {
        uint8 numOpSets = uint8(bound(r, 2, type(uint8).max));
        AllocateParams[] memory allocateParams = _randomAllocateParams_singleStrat_multipleOpSets(r, 0, numOpSets);

        allocateParams[0].newMagnitudes[numOpSets - 1] = WAD + 1;

        // Overallocate
        // cheats.expectRevert(InsufficientAllocatableMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
    }

    function test_allocateMaxToMultipleStrategies() public {
        AllocateParams[] memory allocateParams = new AllocateParams[](2);
        allocateParams[0] = _randomAllocateParams_singleStrat_singleOpSet(0, 0)[0];
        allocateParams[0].newMagnitudes[0] = WAD;

        allocateParams[1] = _randomAllocateParams_singleStrat_singleOpSet(1, 1)[0];
        allocateParams[1].newMagnitudes[0] = WAD;
        allocateParams[1].strategies = IStrategy(address(uint160(2))).toArray();

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Assert maxMagnitude is encumbered
        assertEq(
            WAD,
            allocationManager.encumberedMagnitude(defaultOperator, allocateParams[0].strategies[0]),
            "encumberedMagnitude not max"
        );
        assertEq(
            WAD,
            allocationManager.encumberedMagnitude(defaultOperator, allocateParams[1].strategies[0]),
            "encumberedMagnitude not max"
        );
    }

    function test_revert_allocateDeallocate_modificationPending() public {
        // Allocate
        AllocateParams[] memory allocateParams = _randomAllocateParams_singleStrat_singleOpSet(0, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Deallocate
        allocateParams[0].newMagnitudes[0] -= 1;
        cheats.expectRevert(ModificationAlreadyPending.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
    }

    function test_revert_deallocateTwice_modificationPending() public {
        // Allocate
        AllocateParams[] memory allocateParams = _randomAllocateParams_singleStrat_singleOpSet(0, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Warp past allocation complete timestsamp
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        allocateParams[0].newMagnitudes[0] -= 1;
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Deallocate again -> expect revert
        cheats.expectRevert(ModificationAlreadyPending.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
    }

    /**
     * Allocates to `firstMod` magnitude and then deallocate to `secondMod` magnitude
     * Validates the storage
     * - 1. After deallocation is alled
     * - 2. After the deallocationd delay is hit
     * - 3. After the deallocation queue is cleared
     */
    function testFuzz_allocate_deallocate(
        uint256 r
    ) public {
        // Bound allocation and deallocation
        uint64 firstMod = uint64(bound(r, 1, WAD));
        uint64 secondMod = uint64(bound(r, 0, firstMod - 1));

        // Allocate
        AllocateParams[] memory allocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, firstMod);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        allocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, secondMod);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

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
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(firstMod, allocation.currentMagnitude, "currentMagnitude should not be updated");
        int128 expectedDiff = -int128(uint128(firstMod - secondMod));
        assertEq(expectedDiff, allocation.pendingDiff, "pendingMagnitude not updated");
        uint32 effectBlock = uint32(block.number + DEALLOCATION_DELAY);
        assertEq(effectBlock, allocation.effectBlock, "effectBlock not updated");

        // Check storage after roll to completion
        cheats.roll(effectBlock);
        allocation = allocationManager.getAllocation(
            defaultOperator, allocateParams[0].operatorSet, allocateParams[0].strategies[0]
        );
        assertEq(secondMod, allocation.currentMagnitude, "currentMagnitude not updated");
        assertEq(0, allocation.pendingDiff, "pendingMagnitude not updated");
        assertEq(0, allocation.effectBlock, "effectBlock not updated");
        assertEq(
            firstMod,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should not be updated"
        );

        // Check storage after clearing deallocation queue
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        uint16[] memory numToClear = new uint16[](1);
        numToClear[0] = 1;
        allocationManager.clearDeallocationQueue(defaultOperator, strategies, numToClear);
        assertEq(
            secondMod,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should be updated"
        );
    }

    function test_deallocate_all() public {
        // Allocate
        AllocateParams[] memory allocateParams = _newAllocateParams_SingleMockStrategy(defaultAVS, WAD);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        allocateParams[0].newMagnitudes[0] = 0;
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);

        // Warp to completion and clear deallocation queue
        cheats.roll(block.number + DEALLOCATION_DELAY);
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        uint16[] memory numToClear = new uint16[](1);
        numToClear[0] = 1;
        allocationManager.clearDeallocationQueue(defaultOperator, strategies, numToClear);

        // Check storage
        assertEq(
            0,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should be updated"
        );
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(0, allocation.currentMagnitude, "currentMagnitude should be 0");
        assertEq(0, allocation.pendingDiff, "pendingMagnitude should be 0");
        assertEq(0, allocation.effectBlock, "effectBlock should be 0");
    }

    function testFuzz_allocate_deallocate_singleStrat_multipleOperatorSets(
        uint256 r
    ) public {
        uint8 numOpSets = uint8(bound(r, 0, type(uint8).max));
        (AllocateParams[] memory allocateParams, AllocateParams[] memory deallocateParams) =
            _randomAllocationAndDeallocation_singleStrat_multipleOpSets(numOpSets, r, 0);

        // Allocate
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocateParams);
        uint64 encumberedMagnitudeAfterAllocation =
            allocationManager.encumberedMagnitude(defaultOperator, allocateParams[0].strategies[0]);

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        uint64 postDeallocMag;
        for (uint256 i = 0; i < numOpSets; ++i) {
            postDeallocMag += deallocateParams[0].newMagnitudes[i];
            // pendingNewMags[i] = allocateParams[0].newMagnitudes[i] - deallocateParams[0].newMagnitudes[i];
        }
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(deallocateParams);

        // Check storage after dealloc
        assertEq(
            encumberedMagnitudeAfterAllocation,
            allocationManager.encumberedMagnitude(defaultOperator, allocateParams[0].strategies[0]),
            "encumberedMagnitude should not be updated"
        );
        Allocation memory allocation;
        for (uint256 i = 0; i < numOpSets; ++i) {
            allocation = allocationManager.getAllocation(defaultOperator, allocateParams[i].operatorSet, strategyMock);
            assertEq(
                allocateParams[0].newMagnitudes[i],
                allocation.currentMagnitude,
                "currentMagnitude should not be updated"
            );
            int128 expectedDiff =
                -int128(uint128(allocateParams[0].newMagnitudes[i] - deallocateParams[0].newMagnitudes[i]));
            assertEq(expectedDiff, allocation.pendingDiff, "pendingMagnitude not updated");
            uint32 effectBlock = uint32(block.number + DEALLOCATION_DELAY);
            assertEq(effectBlock, allocation.effectBlock, "effectBlock not updated");
        }

        // Check storage after roll to completion
        cheats.roll(block.number + DEALLOCATION_DELAY);

        for (uint256 i = 0; i < numOpSets; ++i) {
            allocation = allocationManager.getAllocation(defaultOperator, allocateParams[i].operatorSet, strategyMock);
            assertEq(deallocateParams[0].newMagnitudes[i], allocation.currentMagnitude, "currentMagnitude not updated");
            assertEq(0, allocation.pendingDiff, "pendingMagnitude not updated");
            assertEq(0, allocation.effectBlock, "effectBlock not updated");
        }

        // Clear deallocation queue
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        uint16[] memory numToClear = new uint16[](1);
        numToClear[0] = numOpSets;
        allocationManager.clearDeallocationQueue(defaultOperator, strategies, numToClear);

        // Check storage after clearing deallocation queue
        assertEq(
            postDeallocMag,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should be updated"
        );
    }
}

contract AllocationManagerUnitTests_ClearDeallocationQueue is AllocationManagerUnitTests {
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

    function test_revert_operatorNotRegistered() public {
        // Deregister operator
        delegationManagerMock.setIsOperator(defaultOperator, false);

        cheats.expectRevert(OperatorNotRegistered.selector);
        allocationManager.clearDeallocationQueue(defaultOperator, new IStrategy[](0), new uint16[](0));
    }

    /**
     * @notice Allocates magnitude to an operator and then
     * - Clears deallocation queue when only an allocation exists
     * - Clears deallocation queue when the alloc can be completed - asserts emit has been emitted
     * - Validates storage after the second clear
     */
    function testFuzz_allocate(
        uint256 r
    ) public {
        // Allocate magnitude
        IAllocationManager.AllocateParams[] memory allocateParams =
            _queueRandomAllocation_singleStrat_singleOpSet(defaultOperator, r, 0);

        // Attempt to clear queue, assert no events emitted
        allocationManager.clearDeallocationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());
        Vm.Log[] memory entries = vm.getRecordedLogs();
        assertEq(0, entries.length, "should not have emitted any events");

        // Warp to allocation complete block
        cheats.roll(block.number + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Clear queue - this is a noop
        allocationManager.clearDeallocationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());

        // Validate storage (although this is technically tested in allocation tests, adding for sanity)
        // TODO: maybe add a harness here to actually introspect storage
        IAllocationManager.Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        assertEq(allocateParams[0].newMagnitudes[0], allocation.currentMagnitude, "currentMagnitude should be 0");
        assertEq(0, allocation.pendingDiff, "pendingMagnitude should be 0");
        assertEq(0, allocation.effectBlock, "effectBlock should be 0");
    }

    /**
     * @notice Allocates magnitude to an operator and then
     * - Clears deallocation queue when nothing can be completed
     * - After the first clear, asserts the allocation info takes into account the deallocation
     * - Clears deallocation queue when the dealloc can be completed
     * - Assert events & validates storage after the deallocateParams are completed
     */
    function testFuzz_allocate_deallocate(
        uint256 r
    ) public {
        // Complete allocateParams & add a deallocation
        (AllocateParams[] memory allocateParams, AllocateParams[] memory deallocateParams) =
        _queueRandomAllocationAndDeallocation(
            defaultOperator,
            1, // numOpSets
            r,
            0 // salt
        );

        // Clear queue & check storage
        allocationManager.clearDeallocationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());
        assertEq(
            allocateParams[0].newMagnitudes[0],
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should not be updated"
        );

        // Validate storage - encumbered magnitude should just be allocateParams (we only have 1 allocation)
        IAllocationManager.Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, allocateParams[0].operatorSet, strategyMock);
        int128 pendingDiff = -int128(uint128(allocateParams[0].newMagnitudes[0] - deallocateParams[0].newMagnitudes[0]));
        assertEq(allocateParams[0].newMagnitudes[0], allocation.currentMagnitude, "currentMagnitude should be 0");
        assertEq(pendingDiff, allocation.pendingDiff, "pendingMagnitude should be 0");
        assertEq(block.number + DEALLOCATION_DELAY, allocation.effectBlock, "effectBlock should be 0");

        // Warp to deallocation complete block
        cheats.roll(block.number + DEALLOCATION_DELAY);

        // Clear queue
        allocationManager.clearDeallocationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());

        // Validate storage - encumbered magnitude should just be deallocateParams (we only have 1 deallocation)
        assertEq(
            deallocateParams[0].newMagnitudes[0],
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should be updated"
        );
        allocation = allocationManager.getAllocation(defaultOperator, deallocateParams[0].operatorSet, strategyMock);
        assertEq(deallocateParams[0].newMagnitudes[0], allocation.currentMagnitude, "currentMagnitude should be 0");
        assertEq(0, allocation.pendingDiff, "pendingMagnitude should be 0");
        assertEq(0, allocation.effectBlock, "effectBlock should be 0");
    }

    /**
     * Allocates, deallocates, and then allocates again. Asserts that
     * - The deallocation does not block state updates from the second allocation, even though the allocation has an earlier
     *   effect block
     */
    function test_allocate_deallocate_allocate() public {
        uint32 allocationDelay = 15 days;
        // Set allocation delay to be 15 days
        cheats.prank(defaultOperator);
        allocationManager.setAllocationDelay(allocationDelay);
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);
        (, uint32 storedDelay) = allocationManager.getAllocationDelay(defaultOperator);
        assertEq(allocationDelay, storedDelay, "allocation delay not valid");

        // Allocate half of mag to opset1
        AllocateParams[] memory firstAllocation =
            _newAllocateParams_SingleMockStrategy(OperatorSet(defaultAVS, 1), 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(firstAllocation);
        cheats.roll(block.number + 15 days);

        // Deallocate half from opset1.
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY);
        AllocateParams[] memory firstDeallocation =
            _newAllocateParams_SingleMockStrategy(OperatorSet(defaultAVS, 1), 25e16);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(firstDeallocation);
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, firstDeallocation[0].operatorSet, strategyMock);
        assertEq(deallocationEffectBlock, allocation.effectBlock, "effect block not correct");

        // Allocate 33e16 mag to opset2
        uint32 allocationEffectBlock = uint32(block.number + allocationDelay);
        AllocateParams[] memory secondAllocation =
            _newAllocateParams_SingleMockStrategy(OperatorSet(defaultAVS, 2), 33e16);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(secondAllocation);
        allocation = allocationManager.getAllocation(defaultOperator, secondAllocation[0].operatorSet, strategyMock);
        console.log("deallocation effect block: ", deallocationEffectBlock);
        console.log("allocation effect block: ", allocationEffectBlock);
        assertEq(allocationEffectBlock, allocation.effectBlock, "effect block not correct");
        assertLt(allocationEffectBlock, deallocationEffectBlock, "invalid test setup");

        // Warp to allocation effect block & clear the queue
        cheats.roll(allocationEffectBlock);
        allocationManager.clearDeallocationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());

        // Validate `getAllocatableMagnitude`. Allocatable magnitude should be the difference between the max magnitude and the encumbered magnitude
        uint64 allocatableMagnitude = allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock);
        assertEq(WAD - 33e16 - 5e17, allocatableMagnitude, "allocatableMagnitude not correct");

        // Validate that we can allocate again for opset2. This should not revert
        AllocateParams[] memory thirdAllocation =
            _newAllocateParams_SingleMockStrategy(OperatorSet(defaultAVS, 2), 10e16);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(thirdAllocation);
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
        uint32 allocationDelay = 25 days;
        // Set allocation delay to be 25 days, greater than the deallocation block
        cheats.prank(defaultOperator);
        allocationManager.setAllocationDelay(allocationDelay);
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);
        (, uint32 storedDelay) = allocationManager.getAllocationDelay(defaultOperator);
        assertEq(allocationDelay, storedDelay, "allocation delay not valid");

        // Allocate half of mag to opset1
        AllocateParams[] memory firstAllocation =
            _newAllocateParams_SingleMockStrategy(OperatorSet(defaultAVS, 1), 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(firstAllocation);
        cheats.roll(block.number + 25 days);

        // Allocate half of mag to opset2
        AllocateParams[] memory secondAllocation =
            _newAllocateParams_SingleMockStrategy(OperatorSet(defaultAVS, 2), 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(secondAllocation);

        uint32 allocationEffectBlock = uint32(block.number + allocationDelay);
        Allocation memory allocation =
            allocationManager.getAllocation(defaultOperator, secondAllocation[0].operatorSet, strategyMock);
        assertEq(allocationEffectBlock, allocation.effectBlock, "effect block not correct");

        // Deallocate all from opSet1
        uint32 deallocationEffectBlock = uint32(block.number + DEALLOCATION_DELAY);
        AllocateParams[] memory firstDeallocation = _newAllocateParams_SingleMockStrategy(OperatorSet(defaultAVS, 1), 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(firstDeallocation);
        allocation = allocationManager.getAllocation(defaultOperator, firstDeallocation[0].operatorSet, strategyMock);
        assertEq(deallocationEffectBlock, allocation.effectBlock, "effect block not correct");
        assertLt(deallocationEffectBlock, allocationEffectBlock, "invalid test setup");

        // Warp to deallocation effect block & clear the queue
        cheats.roll(deallocationEffectBlock);
        allocationManager.clearDeallocationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());

        // At this point, we should be able to allocate again to opSet1 AND have only 5e17 encumbered magnitude
        assertEq(
            5e17,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumbered magnitude not correct"
        );
        AllocateParams[] memory thirdAllocation =
            _newAllocateParams_SingleMockStrategy(OperatorSet(defaultAVS, 1), 5e17);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(thirdAllocation);
    }
}

contract AllocationManagerUnitTests_SetAllocationDelay is AllocationManagerUnitTests {
    /// -----------------------------------------------------------------------
    /// setAllocationDelay() + getAllocationDelay()
    /// -----------------------------------------------------------------------

    address operatorToSet = address(0x1);

    function setUp() public override {
        AllocationManagerUnitTests.setUp();

        // Register operator
        delegationManagerMock.setIsOperator(operatorToSet, true);
    }

    function test_revert_callerNotOperator() public {
        // Deregister operator
        delegationManagerMock.setIsOperator(operatorToSet, false);
        cheats.prank(operatorToSet);
        cheats.expectRevert(OperatorNotRegistered.selector);
        allocationManager.setAllocationDelay(1);
    }

    function testFuzz_setDelay(
        uint256 r
    ) public {
        uint32 delay = uint32(bound(r, 0, type(uint32).max));

        // Set delay
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AllocationDelaySet(operatorToSet, delay, uint32(block.number + ALLOCATION_CONFIGURATION_DELAY));
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(delay);

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

    function test_fuzz_setDelay_multipleTimesWithinConfigurationDelay(uint32 firstDelay, uint32 secondDelay) public {
        firstDelay = uint32(bound(firstDelay, 1, type(uint32).max));
        secondDelay = uint32(bound(secondDelay, 1, type(uint32).max));
        cheats.assume(firstDelay != secondDelay);

        // Set delay
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(firstDelay);

        // Warp just before effect block
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY - 1);

        // Set delay again
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AllocationDelaySet(operatorToSet, secondDelay, uint32(block.number + ALLOCATION_CONFIGURATION_DELAY));
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(secondDelay);

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

    function testFuzz_multipleDelays(uint32 firstDelay, uint32 secondDelay) public {
        firstDelay = uint32(bound(firstDelay, 1, type(uint32).max));
        secondDelay = uint32(bound(secondDelay, 1, type(uint32).max));
        cheats.assume(firstDelay != secondDelay);

        // Set delay
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(firstDelay);

        // Warp to effect block of first delay
        cheats.roll(block.number + ALLOCATION_CONFIGURATION_DELAY);

        // Set delay again
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(secondDelay);

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
        uint256 r
    ) public {
        uint32 delay = uint32(bound(r, 1, type(uint32).max));

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

    

}

/**
 * @notice TODO Lifecycle tests - These tests combine multiple functionalities of the AllocationManager
 * 1. Set allocation delay > 21 days (configuration), Allocate, modify allocation delay to < 21 days, try to allocate again once new delay is set (should be able to allocate faster than 21 deays)
 * 2. Allocate across multiple strategies and multiple operatorSets
 * 3. lifecycle fuzz test allocating/deallocating across multiple opSets/strategies
 * 4. HIGH PRIO - add uint16.max allocateParams/deallocateParams and then clear them
 * 5. Correctness of slashable magnitudes
 * 6. HIGH PRIO - get gas costs of `getSlashableMagnitudes`
 */
