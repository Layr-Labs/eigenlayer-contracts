// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/core/AllocationManager.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

contract AllocationManagerUnitTests is EigenLayerUnitTestSetup, IAllocationManagerErrors, IAllocationManagerEvents {
    uint8 internal constant PAUSED_MODIFY_ALLOCATIONS = 0;
    uint8 internal constant PAUSED_OPERATOR_SLASHING = 1;
    uint32 constant DEALLOCATION_DELAY = 17.5 days;
    uint32 constant ALLOCATION_CONFIGURATION_DELAY = 21 days;

    AllocationManager allocationManager;
    ERC20PresetFixedSupply tokenMock;
    StrategyBase strategyMock;
    StrategyBase strategyMock2;

    address defaultOperator = address(this);
    address defaultAVS = address(0xFEDBAD);
    uint32 constant DEFAULT_OPERATOR_ALLOCATION_DELAY = 1 days;
    

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

        // Set the allocation delay & warp to when it can be set
        delegationManagerMock.setIsOperator(defaultOperator, true);
        cheats.prank(defaultOperator);
        allocationManager.setAllocationDelay(DEFAULT_OPERATOR_ALLOCATION_DELAY);
        cheats.warp(block.timestamp + ALLOCATION_CONFIGURATION_DELAY);
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
                            IAVSDirectory(address(avsDirectoryMock)),
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

    /**
     * @notice Generated magnitue allocation calldata for a given `avsToSet`, `strategy`, and `operatorSetId`
     */
    function _generateMagnitudeAllocationCalldata_opSetAndStrategy(
        address avsToSet,
        IStrategy strategy,
        uint32 operatorSetId,
        uint64 magnitudeToSet,
        uint64 expectedMaxMagnitude
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        OperatorSet[] memory operatorSets = new OperatorSet[](1);
        operatorSets[0] = OperatorSet({avs: avsToSet, operatorSetId: operatorSetId});

        // Set operatorSet to being valid
        avsDirectoryMock.setIsOperatorSetBatch(operatorSets, true);

        uint64[] memory magnitudes = new uint64[](1);
        magnitudes[0] = magnitudeToSet;

        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            new IAllocationManagerTypes.MagnitudeAllocation[](1);
        allocations[0] = IAllocationManagerTypes.MagnitudeAllocation({
            strategy: strategy,
            expectedMaxMagnitude: expectedMaxMagnitude,
            operatorSets: operatorSets,
            magnitudes: magnitudes
        });

        return allocations;
    }

    /**
     * @notice Generates magnitudeAllocation calldata for a given operatorSet and avs for `strategyMock`
     */
    function _generateMagnitudeAllocationCalldataForOpSet(
        address avsToSet,
        uint32 operatorSetId,
        uint64 magnitudeToSet,
        uint64 expectedMaxMagnitude
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        return _generateMagnitudeAllocationCalldata_opSetAndStrategy(
            avsToSet,
            strategyMock,
            operatorSetId,
            magnitudeToSet,
            expectedMaxMagnitude
        );
    }

    /**
     * @notice Generates magnitudeAllocation calldata for the `strategyMock` on operatorSet 1 with a provided magnitude.
     */
    function _generateMagnitudeAllocationCalldata(
        address avsToSet,
        uint64 magnitudeToSet,
        uint64 expectedMaxMagnitude
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        return _generateMagnitudeAllocationCalldataForOpSet(avsToSet, 1, magnitudeToSet, expectedMaxMagnitude);
    }

    /// -----------------------------------------------------------------------
    /// Generate random slashing parameters
    /// -----------------------------------------------------------------------

    /**
     * @notice Gets random slashing parameters. Not useful unless the operatorSetID is set. See overloaded method
     */
    function _randomSlashingParams(
        address operator,
        uint256 r,
        uint256 salt
    ) internal view returns (IAllocationManagerTypes.SlashingParams memory) {
        r = uint256(keccak256(abi.encodePacked(r, salt)));

        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;

        return IAllocationManagerTypes.SlashingParams({
            operator: operator,
            operatorSetId: uint32(r),
            strategies: strategies,
            wadToSlash: bound(r, 1, 1e18),
            description: "test"
        });
    }

    function _randomSlashingParams(
        address operator,
        uint32 operatorSetId,
        uint256 r,
        uint256 salt
    ) internal view returns (IAllocationManagerTypes.SlashingParams memory) {
        r = uint256(keccak256(abi.encodePacked(r, salt)));

        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;

        return IAllocationManagerTypes.SlashingParams({
            operator: operator,
            operatorSetId: operatorSetId,
            strategies: strategies,
            wadToSlash: bound(r, 1, 1e18),
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
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _queueRandomAllocation_singleStrat_singleOpSet(operator, avs, r, salt);

        // Warp to allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        return allocations;
    }

    function _queueRandomAllocation_singleStrat_singleOpSet(
        address operator,
        address avs,
        uint256 r,
        uint256 salt
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_singleOpSet(avs, r, salt);
        cheats.prank(operator);
        allocationManager.modifyAllocations(allocations);

        return allocations;
    }

    /**
     * @notice Queued a random allocation for the given `operator`
     * - Does NOT warp past the effect timestamp
     */
    function _queueRandomAllocation_singleStrat_singleOpSet(
        address operator,
        uint256 r,
        uint256 salt
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_singleOpSet(r, salt);
        cheats.prank(operator);
        allocationManager.modifyAllocations(allocations);

        return allocations;
    }

    /**
     * @notice Create a random magnitude allocation
     * Randomized Parameters: avs, opSet, magnitude
     * Non-random Parameters: strategy, expectedMaxMagnitude
     * In addition
     * - Registers the operatorSet with the avsDirectory
     */
    function _randomMagnitudeAllocation_singleStrat_singleOpSet(
        uint256 r,
        uint256 salt
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        r = uint256(keccak256(abi.encodePacked(r, salt)));
        address avs = _randomAddr(r, 0);
        return _randomMagnitudeAllocation_singleStrat_singleOpSet(avs, r, salt);
    }

    /**
     * @notice Create a random magnitude allocation
     * Randomized Parameters: opSet, magnitude
     * Non-random Parameters: strategy, expectedMaxMagnitude, avs
     * In addition
     * - Registers the operatorSet with the avsDirectory
     */
    function _randomMagnitudeAllocation_singleStrat_singleOpSet(
        address avs,
        uint256 r,
        uint256 salt
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        r = uint256(keccak256(abi.encodePacked(r, salt)));

        // Mock a random operator set.
        OperatorSet[] memory operatorSets = new OperatorSet[](1);
        operatorSets[0] = OperatorSet({avs: avs, operatorSetId: uint32(r)});

        // Set operatorSet to being valid
        avsDirectoryMock.setIsOperatorSetBatch(operatorSets, true);

        uint64[] memory magnitudes = new uint64[](1);
        magnitudes[0] = uint64(bound(r, 1, 1e18));

        // Mock a random magnitude allocation.
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            new IAllocationManagerTypes.MagnitudeAllocation[](1);
        allocations[0] = IAllocationManagerTypes.MagnitudeAllocation({
            strategy: strategyMock,
            expectedMaxMagnitude: 1e18, // magnitude starts at 100%
            operatorSets: operatorSets,
            magnitudes: magnitudes
        });
        return allocations;
    }

    /// -----------------------------------------------------------------------
    /// Generate a random allocation for a single strategy and multiple operatorSets
    /// -----------------------------------------------------------------------

    function _randomMagnitudeAllocation_singleStrat_multipleOpSets(
        uint256 r,
        uint256 salt,
        uint8 numOpSets
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        r = uint256(keccak256(abi.encodePacked(r, salt)));

        // Create multiple operatorSets
        OperatorSet[] memory operatorSets = new OperatorSet[](numOpSets);
        for (uint8 i = 0; i < numOpSets; i++) {
            operatorSets[i] = OperatorSet({avs: _randomAddr(r, i), operatorSetId: uint32(r + i)});
        }
        avsDirectoryMock.setIsOperatorSetBatch(operatorSets, true);

        // Give each set a minimum of 1 magnitude
        uint64[] memory magnitudes = new uint64[](numOpSets);
        uint64 usedMagnitude;
        for (uint8 i = 0; i < numOpSets; i++) {
            magnitudes[i] = 1;
            usedMagnitude++;
        }

        // Distribute remaining magnitude
        uint64 maxMagnitude = 1e18;
        for (uint8 i = 0; i < numOpSets; i++) {
            r = uint256(keccak256(abi.encodePacked(r, i)));
            uint64 remainingMagnitude = maxMagnitude - usedMagnitude;
            if (remainingMagnitude > 0) {
                magnitudes[i] += uint64(bound(r, 0, remainingMagnitude));
                usedMagnitude += magnitudes[i] - 1;
            }
        }

        // Create magnitude allocation
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            new IAllocationManagerTypes.MagnitudeAllocation[](1);
        allocations[0] = IAllocationManagerTypes.MagnitudeAllocation({
            strategy: strategyMock,
            expectedMaxMagnitude: 1e18, // magnitude starts at 100%
            operatorSets: operatorSets,
            magnitudes: magnitudes
        });
        return allocations;
    }

    /// -----------------------------------------------------------------------
    /// Generate a random allocation AND delllocation
    /// -----------------------------------------------------------------------

    /**
     * @notice Queued a random allocation and deallocation for the given `operator`
     * - DOES NOT warp past the deallocation effect timestamp
     */
    function _queueRandomAllocationAndDeallocation(
        address operator,
        uint8 numOpSets,
        uint256 r,
        uint256 salt
    )
        internal
        returns (
            IAllocationManagerTypes.MagnitudeAllocation[] memory,
            IAllocationManagerTypes.MagnitudeAllocation[] memory
        )
    {
        (MagnitudeAllocation[] memory allocations, MagnitudeAllocation[] memory deallocations) =
            _randomAllocationAndDeallocation_singleStrat_multipleOpSets(numOpSets, r, salt);

        // Allocate
        cheats.prank(operator);
        allocationManager.modifyAllocations(allocations);

        // Warp to allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        cheats.prank(operator);
        allocationManager.modifyAllocations(deallocations);

        return (allocations, deallocations);
    }

    /**
     * @notice Generates a random allocation and deallocation for a single strategy and multiple operatorSets
     * @notice Deallocations are from 0 to 1 less that the current allocated magnitude
     */
    function _randomAllocationAndDeallocation_singleStrat_multipleOpSets(
        uint8 numOpSets,
        uint256 r,
        uint256 salt
    )
        internal
        returns (
            IAllocationManagerTypes.MagnitudeAllocation[] memory,
            IAllocationManagerTypes.MagnitudeAllocation[] memory
        )
    {
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            new IAllocationManagerTypes.MagnitudeAllocation[](1);
        allocations = _randomMagnitudeAllocation_singleStrat_multipleOpSets(r, salt, numOpSets);

        // Deallocate random magnitude from each of thsoe operatorSets
        r = uint256(keccak256(abi.encodePacked(r, salt)));
        uint64[] memory newMags = new uint64[](numOpSets);
        for (uint8 i = 0; i < numOpSets; i++) {
            newMags[i] = uint64(bound(r, 0, allocations[0].magnitudes[i] - 1));
        }

        // Create deallocations
        IAllocationManagerTypes.MagnitudeAllocation[] memory deallocations =
            new IAllocationManagerTypes.MagnitudeAllocation[](1);
        deallocations[0] = IAllocationManagerTypes.MagnitudeAllocation({
            strategy: strategyMock,
            expectedMaxMagnitude: 1e18, // magnitude starts at 100%
            operatorSets: allocations[0].operatorSets,
            magnitudes: newMags
        });

        return (allocations, deallocations);
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

    function _operatorSet(address avs, uint32 operatorSetId) internal pure returns (OperatorSet memory) {
        return OperatorSet({avs: avs, operatorSetId: operatorSetId});
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
        assertEq(address(alm.avsDirectory()), address(avsDirectoryMock));
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

        cheats.expectRevert(IAllocationManagerErrors.InvalidWadToSlash.selector);
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);
    }

    function test_revert_slashGreaterThanWAD() public {
        SlashingParams memory slashingParams = _randomSlashingParams(defaultOperator, 0, 0);
        slashingParams.wadToSlash = 1e18 + 1;

        cheats.expectRevert(IAllocationManagerErrors.InvalidWadToSlash.selector);
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);
    }

    function test_revert_operatorNotSlashable() public {
        SlashingParams memory slashingParams = _randomSlashingParams(defaultOperator, 0, 0);
        avsDirectoryMock.setIsOperatorSlashable(
            slashingParams.operator, defaultAVS, slashingParams.operatorSetId, false
        );

        cheats.expectRevert(IAllocationManagerErrors.InvalidOperator.selector);
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);
    }

    function test_revert_operatorNotAllocated() public {
        SlashingParams memory slashingParams = _randomSlashingParams(defaultOperator, 0, 0);
        avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

        cheats.expectRevert(IAllocationManagerErrors.OperatorNotAllocated.selector);
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);
    }

    function test_revert_operatorAllocated_notActive() public {
        // Queue allocation
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _queueRandomAllocation_singleStrat_singleOpSet(defaultOperator, 0, 0);

        // Setup data
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocations[0].operatorSets[0].operatorSetId,
            strategies: _strategyMockArray(),
            wadToSlash: 1e18,
            description: "test"
        });
        avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

        // Expect revert
        cheats.expectRevert(IAllocationManagerErrors.OperatorNotAllocated.selector);
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);
    }

    /**
     * Allocates all magnitude to for a single strategy to an operatorSet. Slashes 25%
     * Asserts that:
     * 1. Events are emitted
     * 2. Encumbered mag is updated
     * 3. Max mag is updated
     * 4. Calculations for `getAllocatableMagnitude` and `getAllocationInfo` are correct
     */
    function test_slashPostAllocation() public {
        // Generate allocation for `strategyMock`, we allocate max
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _generateMagnitudeAllocationCalldata(defaultAVS, 1e18, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Slash operator for 25%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocations[0].operatorSets[0].operatorSetId,
            strategies: _strategyMockArray(),
            wadToSlash: 25e16,
            description: "test"
        });
        avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

        // Slash Operator
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, strategyMock, 75e16);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSetMagnitudeUpdated(
            defaultOperator, allocations[0].operatorSets[0], strategyMock, 75e16, uint32(block.timestamp)
        );
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit MaxMagnitudeUpdated(defaultOperator, strategyMock, 75e16);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        uint256[] memory wadSlashed = new uint256[](1);
        wadSlashed[0] = 25e16;
        emit OperatorSlashed(
            slashingParams.operator,
            _operatorSet(defaultAVS, slashingParams.operatorSetId),
            slashingParams.strategies,
            wadSlashed,
            slashingParams.description
        );
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
        MagnitudeInfo[] memory mInfos =
            allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(75e16, mInfos[0].currentMagnitude, "currentMagnitude not updated");
        assertEq(0, mInfos[0].pendingDiff, "pendingDiff should be 0");
        assertEq(0, mInfos[0].effectTimestamp, "effectTimestamp should be 0");
    }

    /// @notice Same test as above, but fuzzes the allocation
    function testFuzz_slashPostAllocation(uint256 r, uint256 salt) public {
        // Complete Allocation for `strategyMock`
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _completeRandomAllocation_singleStrat_singleOpset(defaultOperator, defaultAVS, r, 0);

        // Setup data
        SlashingParams memory slashingParams =
            _randomSlashingParams(defaultOperator, allocations[0].operatorSets[0].operatorSetId, r, 1);
        avsDirectoryMock.setIsOperatorSlashable(
            slashingParams.operator, defaultAVS, allocations[0].operatorSets[0].operatorSetId, true
        );
        uint64 expectedSlashedMagnitude =
            uint64(SlashingLib.mulWad(allocations[0].magnitudes[0], slashingParams.wadToSlash));
        uint64 expectedEncumberedMagnitude = allocations[0].magnitudes[0] - expectedSlashedMagnitude;
        uint64 maxMagnitudeAfterSlash = WAD - expectedSlashedMagnitude;
        uint256[] memory wadSlashed = new uint256[](1);
        wadSlashed[0] = expectedSlashedMagnitude;

        // Slash Operator
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, strategyMock, expectedEncumberedMagnitude);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSetMagnitudeUpdated(
            defaultOperator,
            allocations[0].operatorSets[0],
            strategyMock,
            expectedEncumberedMagnitude,
            uint32(block.timestamp)
        );
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit MaxMagnitudeUpdated(defaultOperator, strategyMock, maxMagnitudeAfterSlash);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSlashed(
            slashingParams.operator,
            _operatorSet(defaultAVS, slashingParams.operatorSetId),
            slashingParams.strategies,
            wadSlashed,
            slashingParams.description
        );
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
        MagnitudeInfo[] memory mInfos =
            allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(expectedEncumberedMagnitude, mInfos[0].currentMagnitude, "currentMagnitude not updated");
        assertEq(0, mInfos[0].pendingDiff, "pendingDiff should be 0");
        assertEq(0, mInfos[0].effectTimestamp, "effectTimestamp should be 0");
    }

    /**
     * Allocates half of magnitude for a single strategy to an operatorSet. Then allocates again. Slashes 50%
     * Asserts that:
     * 1. Events are emitted
     * 2. Encumbered mag is updated
     * 3. Max mag is updated
     * 4. Calculations for `getAllocatableMagnitude` and `getAllocationInfo` are correct
     * 5. The second magnitude allocation is not slashed from
     * TODO: Fuzz
     */
    function test_slash_oneCompletedAlloc_onePendingAlloc() public {
        // Generate allocation for `strategyMock`, we allocate half
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _generateMagnitudeAllocationCalldata(defaultAVS, 5e17, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Allocate the other half
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations2 = _generateMagnitudeAllocationCalldata(defaultAVS, 1e18, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations2);
        uint32 secondAllocEffectTimestamp = uint32(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Slash operator for 50%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocations[0].operatorSets[0].operatorSetId,
            strategies: _strategyMockArray(),
            wadToSlash: 50e16,
            description: "test"
        });
        avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);
        uint64 expectedEncumberedMagnitude = 75e16; // 25e16 from first allocation, 50e16 from second
        uint64 magnitudeAfterSlash = 25e16;
        uint64 maxMagnitudeAfterSlash = 75e16;
        uint256[] memory wadSlashed = new uint256[](1);
        wadSlashed[0] = 25e16;

        // Slash Operator
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, strategyMock, expectedEncumberedMagnitude);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSetMagnitudeUpdated(defaultOperator, allocations[0].operatorSets[0], strategyMock, magnitudeAfterSlash, uint32(block.timestamp));
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit MaxMagnitudeUpdated(defaultOperator, strategyMock, maxMagnitudeAfterSlash);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSlashed(slashingParams.operator, _operatorSet(defaultAVS, slashingParams.operatorSetId), slashingParams.strategies, wadSlashed, slashingParams.description);
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage
        assertEq(expectedEncumberedMagnitude, allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude not updated");
        assertEq(maxMagnitudeAfterSlash, allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0], "maxMagnitude not updated");
        MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(magnitudeAfterSlash, mInfos[0].currentMagnitude, "currentMagnitude not updated");
        assertEq(5e17, mInfos[0].pendingDiff, "pendingDiff should be for second alloc");
        assertEq(secondAllocEffectTimestamp, mInfos[0].effectTimestamp, "effectTimestamp should be 0");

        // Warp to complete second allocation
        cheats.warp(secondAllocEffectTimestamp);
        uint64 allocatableMagnitude = allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock);
        assertEq(0, allocatableMagnitude, "allocatableMagnitude should be 0");
        mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations2[0].operatorSets);
        assertEq(75e16, mInfos[0].currentMagnitude, "currentMagnitude not updated");
        assertEq(0, mInfos[0].pendingDiff, "pendingDiff should be 0");
        assertEq(0, mInfos[0].effectTimestamp, "effectTimestamp should be 0");
    }

    /**
     * Allocates all of magnitude to a single strategy to an operatorSet. Deallocate half. Finally, slash while deallocation is pending
     * Asserts that:
     * 1. Events are emitted, including for deallocation
     * 2. Encumbered mag is updated
     * 3. Max mag is updated
     * 4. Calculations for `getAllocatableMagnitude` and `getAllocationInfo` are correct
     * 5. The deallocation is slashed from
     * 6. Pending magnitude updates post deallocation are valid
     * TODO: Fuzz the allocation & slash amounts
     */
    function test_allocateAll_deallocateHalf_slashWhileDeallocPending() public {
        uint64 initialMagnitude = 1e18;
        // Generate allocation for `strategyMock`, we allocate half
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _generateMagnitudeAllocationCalldata(defaultAVS, initialMagnitude, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate half
        IAllocationManagerTypes.MagnitudeAllocation[] memory deallocations = _generateMagnitudeAllocationCalldata(defaultAVS, initialMagnitude / 2, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(deallocations);
        uint32 deallocationEffectTimestamp = uint32(block.timestamp + DEALLOCATION_DELAY);

        // Slash operator for 25%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocations[0].operatorSets[0].operatorSetId,
            strategies: _strategyMockArray(),
            wadToSlash: 25e16,
            description: "test"
        });
        avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);
        uint64 magnitudeAfterDeallocationSlash = 375e15; // 25% is slashed off of 5e17
        uint64 expectedEncumberedMagnitude = 75e16; // 25e16 is slashed. 75e16 is encumbered
        uint64 magnitudeAfterSlash = 75e16;
        uint64 maxMagnitudeAfterSlash = 75e16; // Operator can only allocate up to 75e16 magnitude since 25% is slashed

        // Slash Operator
        // First event is emitted because of deallocation
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSetMagnitudeUpdated(defaultOperator, allocations[0].operatorSets[0], strategyMock, magnitudeAfterDeallocationSlash, deallocationEffectTimestamp);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, strategyMock, expectedEncumberedMagnitude);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSetMagnitudeUpdated(defaultOperator, allocations[0].operatorSets[0], strategyMock, magnitudeAfterSlash, uint32(block.timestamp));
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit MaxMagnitudeUpdated(defaultOperator, strategyMock, maxMagnitudeAfterSlash);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        uint256[] memory wadSlashed = new uint256[](1);
        wadSlashed[0] = 25e16;
        emit OperatorSlashed(slashingParams.operator, _operatorSet(defaultAVS, slashingParams.operatorSetId), slashingParams.strategies, wadSlashed, slashingParams.description);
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage post slash
        assertEq(expectedEncumberedMagnitude, allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude not updated");
        assertEq(maxMagnitudeAfterSlash, allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0], "maxMagnitude not updated");
        MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(magnitudeAfterSlash, mInfos[0].currentMagnitude, "currentMagnitude not updated");
        assertEq(-int128(uint128((uint64(magnitudeAfterDeallocationSlash)))), mInfos[0].pendingDiff, "pendingDiff should be decreased after slash");
        assertEq(deallocationEffectTimestamp, mInfos[0].effectTimestamp, "effectTimestamp should be 0");

        // Check storage after complete modification
        cheats.warp(deallocationEffectTimestamp);
        allocationManager.clearModificationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());
        mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(magnitudeAfterDeallocationSlash, mInfos[0].currentMagnitude, "currentMagnitude not updated");
        assertEq(magnitudeAfterDeallocationSlash, maxMagnitudeAfterSlash / 2, "magnitude after deallocation should be half of max magnitude, since we originally deallocated by half");
    }

    /**
     * Allocates all magnitude to a single opSet. Then slashes the entire magnitude
     * Asserts that:
     * 1. The operator cannot allocate again
     */
    function testRevert_allocateAfterSlashedEntirely() public {
        // Allocate all magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _generateMagnitudeAllocationCalldata(defaultAVS, 1e18, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Slash operator for 100%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocations[0].operatorSets[0].operatorSetId,
            strategies: _strategyMockArray(),
            wadToSlash: 1e18,
            description: "test"
        });
        avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Attempt to allocate
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations2 = _generateMagnitudeAllocationCalldata(defaultAVS, 1, 0);
        cheats.expectRevert(IAllocationManagerErrors.InsufficientAllocatableMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations2);
    }

    /**
     * Allocates all magnitude to a single opSet. Deallocateas magnitude. Slashes al
     * Asserts that:
     * 1. The MagnitudeInfo is 0 after slash
     * 2. Them sotrage post slash for encumbered and maxMags ais zero
     */
    function test_allocateAll_deallocateAll() public {
        // Allocate all magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _generateMagnitudeAllocationCalldata(defaultAVS, 1e18, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate all
        IAllocationManagerTypes.MagnitudeAllocation[] memory deallocations = _generateMagnitudeAllocationCalldata(defaultAVS, 0, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(deallocations);
        uint32 deallocationEffectTimestamp = uint32(block.timestamp + DEALLOCATION_DELAY);

        // Slash operator for 100%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocations[0].operatorSets[0].operatorSetId,
            strategies: _strategyMockArray(),
            wadToSlash: 1e18,
            description: "test"
        });
        avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

        // Slash Operator
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSetMagnitudeUpdated(defaultOperator, allocations[0].operatorSets[0], strategyMock, 0, deallocationEffectTimestamp);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, strategyMock, 0);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSetMagnitudeUpdated(defaultOperator, allocations[0].operatorSets[0], strategyMock, 0, uint32(block.timestamp));
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit MaxMagnitudeUpdated(defaultOperator, strategyMock, 0);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        uint256[] memory wadSlashed = new uint256[](1);
        wadSlashed[0] = 1e18;
        emit OperatorSlashed(slashingParams.operator, _operatorSet(defaultAVS, slashingParams.operatorSetId), slashingParams.strategies, wadSlashed, slashingParams.description);

        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage post slash
        assertEq(0, allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude not updated");
        assertEq(0, allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0], "maxMagnitude not updated");
        MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(0, mInfos[0].currentMagnitude, "currentMagnitude not updated");
        assertEq(0, mInfos[0].pendingDiff, "pendingDiff should be zero since everything is slashed");
        assertEq(deallocationEffectTimestamp, mInfos[0].effectTimestamp, "effectTimestamp should be 0");
    }

    /**
     * Slashes the operator after deallocation, even if the deallocation has not been cleared. Validates that:
     * 1. Even if we do not clear modification queue, the deallocation is NOT slashed from since we're passed the deallocationEffectTimestamp
     * 2. Validates storage post slash & post clearing modification queue
     * 3. Total magnitude only decreased proportionally by the magnitude set after deallocation
     */
    function test_allocate_deallocate_slashAfterDeallocation() public {
        // Allocate all magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _generateMagnitudeAllocationCalldata(defaultAVS, 1e18, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate half
        IAllocationManagerTypes.MagnitudeAllocation[] memory deallocations = _generateMagnitudeAllocationCalldata(defaultAVS, 5e17, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(deallocations);
        uint32 deallocationEffectTimestamp = uint32(block.timestamp + DEALLOCATION_DELAY);

        // Check storage post deallocation
        MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(1e18, mInfos[0].currentMagnitude, "currentMagnitude not updated");
        assertEq(-5e17, mInfos[0].pendingDiff, "pendingDiff should be 5e17 after deallocation");
        assertEq(deallocationEffectTimestamp, mInfos[0].effectTimestamp, "effectTimestamp should be 0");

        // Warp to deallocation effect timestamp
        cheats.warp(deallocationEffectTimestamp);


        // Slash operator for 25%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocations[0].operatorSets[0].operatorSetId,
            strategies: _strategyMockArray(),
            wadToSlash: 25e16,
            description: "test"
        });
        avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);
        uint64 expectedEncumberedMagnitude = 375e15; // 25e16 is slashed. 5e17 was previously
        uint64 magnitudeAfterSlash = 375e15;
        uint64 maxMagnitudeAfterSlash = 875e15; // Operator can only allocate up to 75e16 magnitude since 25% is slashed

        // Slash Operator, only emit events assuming that there is no deallocation
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, strategyMock, expectedEncumberedMagnitude);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSetMagnitudeUpdated(defaultOperator, allocations[0].operatorSets[0], strategyMock, magnitudeAfterSlash, uint32(block.timestamp));
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit MaxMagnitudeUpdated(defaultOperator, strategyMock, maxMagnitudeAfterSlash);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        uint256[] memory wadSlashed = new uint256[](1);
        wadSlashed[0] = 125e15;
        emit OperatorSlashed(slashingParams.operator, _operatorSet(defaultAVS, slashingParams.operatorSetId), slashingParams.strategies, wadSlashed, slashingParams.description);
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage post slash
        assertEq(expectedEncumberedMagnitude, allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude not updated");
        assertEq(maxMagnitudeAfterSlash, allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0], "maxMagnitude not updated");
        mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(magnitudeAfterSlash, mInfos[0].currentMagnitude, "currentMagnitude not updated");
        assertEq(0, mInfos[0].pendingDiff, "pendingDiff should be 0 after slash");
        assertEq(0, mInfos[0].effectTimestamp, "effectTimestamp should be 0");
        uint64 allocatableMagnitudeAfterSlash = allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock);

        // Check storage after complete modification. Expect encumberedMag to be emitted again
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, strategyMock, expectedEncumberedMagnitude);
        allocationManager.clearModificationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());
        mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(allocatableMagnitudeAfterSlash, allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock), "allocatable mag after slash shoudl be equal to allocatable mag after clearing queue");
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

        // Allocate 40% to firstOperatorSet, 40% to secondOperatorSet
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = new IAllocationManagerTypes.MagnitudeAllocation[](2);
        allocations[0] = _generateMagnitudeAllocationCalldataForOpSet(defaultAVS, 1, magnitudeToAllocate, 1e18)[0];
        allocations[1] = _generateMagnitudeAllocationCalldataForOpSet(defaultAVS, 2, magnitudeToAllocate, 1e18)[0];
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Get slashable shares for each operatorSet
        address[] memory operatorArray = new address[](1);
        operatorArray[0] = defaultOperator;
        (, uint256[][] memory slashableSharesOpset1_preSlash) = allocationManager.getMinDelegatedAndSlashableOperatorShares(
            _operatorSet(defaultAVS, 1),
            operatorArray,
            _strategyMockArray(),
            uint32(block.timestamp + 1)
        );
        (, uint256[][] memory slashableSharesOpset2_preSlash) = allocationManager.getMinDelegatedAndSlashableOperatorShares(
            _operatorSet(defaultAVS, 2),
            operatorArray,
            _strategyMockArray(),
            uint32(block.timestamp + 1)
        );
        assertEq(40e18, slashableSharesOpset1_preSlash[0][0], "slashableShares of opSet_1 should be 40e18");
        assertEq(40e18, slashableSharesOpset2_preSlash[0][0], "slashableShares of opSet_2 should be 40e18");
        uint256 maxMagnitude = allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0];
        uint256 opSet2PortionOfMaxMagnitude = uint256(magnitudeToAllocate) * 1e18 / maxMagnitude;

        // Slash operator on operatorSet1 for 50%
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocations[0].operatorSets[0].operatorSetId,
            strategies: _strategyMockArray(),
            wadToSlash: 5e17,
            description: "test"
        });
        avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Operator should now have 80e18 shares, since half of 40e18 was slashed
        delegationManagerMock.setOperatorShares(defaultOperator, strategyMock, 80e18);

        // Check storage
        (, uint256[][] memory slashableSharesOpset1_postSlash) = allocationManager.getMinDelegatedAndSlashableOperatorShares(
            _operatorSet(defaultAVS, 1),
            operatorArray,
            _strategyMockArray(),
            uint32(block.timestamp + 1)
        );
        (, uint256[][] memory slashableSharesOpset2_postSlash) = allocationManager.getMinDelegatedAndSlashableOperatorShares(
            _operatorSet(defaultAVS, 2),
            operatorArray,
            _strategyMockArray(),
            uint32(block.timestamp + 1)
        );
        
        assertEq(20e18, slashableSharesOpset1_postSlash[0][0], "slashableShares of opSet_1 should be 20e18");
        assertEq(slashableSharesOpset2_preSlash[0][0], slashableSharesOpset2_postSlash[0][0], "slashableShares of opSet_2 should remain unchanged");

        // Validate encumbered and total allocatable magnitude
        uint256 maxMagnitudeAfterSlash = allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0];
        uint256 expectedEncumberedMagnitude = 6e17; // 4e17 from opSet2, 2e17 from opSet1
        assertEq(expectedEncumberedMagnitude, allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude not updated");
        assertEq(maxMagnitudeAfterSlash - expectedEncumberedMagnitude, allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock), "allocatableMagnitude should be diff of maxMagnitude and encumberedMagnitude");

        // Check proportion after slash
        uint256 opSet2PortionOfMaxMagnitudeAfterSlash = uint256(magnitudeToAllocate) * 1e18 / maxMagnitudeAfterSlash;
        assertGt(opSet2PortionOfMaxMagnitudeAfterSlash, opSet2PortionOfMaxMagnitude, "opSet2 should have a greater proportion to slash from previous");
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
        uint64 strategy2Magnitude = 1e18;
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = new IAllocationManagerTypes.MagnitudeAllocation[](2);
        allocations[0] = _generateMagnitudeAllocationCalldata_opSetAndStrategy(defaultAVS, strategyMock, 1, strategy1Magnitude, 1e18)[0];
        allocations[1] = _generateMagnitudeAllocationCalldata_opSetAndStrategy(defaultAVS, strategyMock2, 1, strategy2Magnitude, 1e18)[0];
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);


        // Slash operator on both strategies for 60%
        IStrategy[] memory strategiesToSlash = new IStrategy[](2);
        strategiesToSlash[0] = strategyMock;
        strategiesToSlash[1] = strategyMock2;
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocations[0].operatorSets[0].operatorSetId,
            strategies: strategiesToSlash,
            wadToSlash: 6e17,
            description: "test"
        });
        avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

        uint64[] memory expectedEncumberedMags = new uint64[](2);
        expectedEncumberedMags[0] = 2e17; // 60% of 5e17
        expectedEncumberedMags[1] = 4e17; // 60% of 1e18

        uint64[] memory expectedMagnitudeAfterSlash = new uint64[](2);
        expectedMagnitudeAfterSlash[0] = 2e17;
        expectedMagnitudeAfterSlash[1] = 4e17; 

        uint64[] memory expectedMaxMagnitudeAfterSlash = new uint64[](2);
        expectedMaxMagnitudeAfterSlash[0] = 7e17;
        expectedMaxMagnitudeAfterSlash[1] = 4e17;

        // Expect emits
        for(uint256 i = 0; i < strategiesToSlash.length; i++) {
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit EncumberedMagnitudeUpdated(defaultOperator, strategiesToSlash[i], expectedEncumberedMags[i]);
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit OperatorSetMagnitudeUpdated(defaultOperator, _operatorSet(defaultAVS, slashingParams.operatorSetId), strategiesToSlash[i], expectedMagnitudeAfterSlash[i], uint32(block.timestamp));
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit MaxMagnitudeUpdated(defaultOperator, strategiesToSlash[i], expectedMaxMagnitudeAfterSlash[i]);
        }
        uint256[] memory wadSlashed = new uint256[](2);
        wadSlashed[0] = 3e17;
        wadSlashed[1] = 6e17;
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSlashed(slashingParams.operator, _operatorSet(defaultAVS, slashingParams.operatorSetId), slashingParams.strategies, wadSlashed, slashingParams.description);
        
        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage
        for(uint256 i = 0; i < strategiesToSlash.length; i++) {
            assertEq(expectedEncumberedMags[i], allocationManager.encumberedMagnitude(defaultOperator, strategiesToSlash[i]), "encumberedMagnitude not updated");
            assertEq(expectedMaxMagnitudeAfterSlash[i] - expectedMagnitudeAfterSlash[i], allocationManager.getAllocatableMagnitude(defaultOperator, strategiesToSlash[i]), "allocatableMagnitude not updated");
            MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategiesToSlash[i], allocations[0].operatorSets);
            assertEq(expectedMagnitudeAfterSlash[i], mInfos[0].currentMagnitude, "currentMagnitude not updated");
            assertEq(0, mInfos[0].pendingDiff, "pendingDiff should be 0");
            assertEq(0, mInfos[0].effectTimestamp, "effectTimestamp should be 0");
        }
    }

    /**
     * Allocates magnitude. Deallocates some. Slashes a portion, and then allocates up to the max available magnitude
     * TODO: Fuzz the wadsToSlash
     */
    function testFuzz_allocate_deallocate_slashWhilePending_allocateMax(uint256 r) public {
        // Bound allocation and deallocation
        uint64 firstMod = uint64(bound(r, 3, 1e18));
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
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _generateMagnitudeAllocationCalldata(defaultAVS, firstMod, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory deallocations = _generateMagnitudeAllocationCalldata(defaultAVS, secondMod, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(deallocations);
        uint32 deallocationEffectTimestamp = uint32(block.timestamp + DEALLOCATION_DELAY);

        // Slash operator for 50%  
        SlashingParams memory slashingParams = SlashingParams({
            operator: defaultOperator,
            operatorSetId: allocations[0].operatorSets[0].operatorSetId,
            strategies: _strategyMockArray(),
            wadToSlash: 5e17,
            description: "test"
        });
        avsDirectoryMock.setIsOperatorSlashable(slashingParams.operator, defaultAVS, slashingParams.operatorSetId, true);

        // Slash Operator
        cheats.prank(defaultAVS);
        allocationManager.slashOperator(slashingParams);

        // Check storage post slash
        assertEq(firstMod / 2, allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude should be half of firstMod");
        MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(firstMod / 2, mInfos[0].currentMagnitude, "currentMagnitude should be half of firstMod");
        console.log("value of pendingDiff: ", pendingDiff - pendingDiff/2);
        assertEq(-int128(uint128(pendingDiff - pendingDiff/2)), mInfos[0].pendingDiff, "pendingDiff should be -secondMod");
        assertEq(deallocationEffectTimestamp, mInfos[0].effectTimestamp, "effectTimestamp should be deallocationEffectTimestamp");

        // Warp to deallocation effect timestamp & clear modification queue
        console.log("encumbered mag before: ", allocationManager.encumberedMagnitude(defaultOperator, strategyMock));
        cheats.warp(deallocationEffectTimestamp);
        allocationManager.clearModificationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());
        console.log("encumbered mag after: ", allocationManager.encumberedMagnitude(defaultOperator, strategyMock));

        // Check expected max and allocatable
        uint64 expectedMaxMagnitude = 1e18 - firstMod / 2;
        assertEq(expectedMaxMagnitude, allocationManager.getMaxMagnitudes(defaultOperator, _strategyMockArray())[0], "maxMagnitude should be expectedMaxMagnitude");
        // Allocatable is expectedMax - currentMagPostSlashing - pendingDiffOfDeallocations post slashing
        uint64 expectedAllocatable = expectedMaxMagnitude - ((firstMod / 2) - (pendingDiff - pendingDiff / 2));
        assertEq(expectedAllocatable, allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock), "allocatableMagnitude should be expectedAllocatable");

        // Allocate up to max magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations2 = _generateMagnitudeAllocationCalldata(defaultAVS, expectedMaxMagnitude, expectedMaxMagnitude);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations2);

        // Assert that encumbered is expectedMaxMagnitude
        assertEq(0, allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock), "allocatableMagnitude should be 0");
    }
}

contract AllocationManagerUnitTests_ModifyAllocations is AllocationManagerUnitTests {
    /// -----------------------------------------------------------------------
    /// modifyAllocations()
    /// -----------------------------------------------------------------------
    function test_revert_paused() public {
        allocationManager.pause(2 ** PAUSED_MODIFY_ALLOCATIONS);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        allocationManager.modifyAllocations(new IAllocationManagerTypes.MagnitudeAllocation[](0));
    }

    function test_revert_allocationDelayNotSet() public {
        address invalidOperator = address(0x2);
        cheats.prank(invalidOperator);
        cheats.expectRevert(IAllocationManagerErrors.UninitializedAllocationDelay.selector);
        allocationManager.modifyAllocations(new IAllocationManagerTypes.MagnitudeAllocation[](0));
    }

    function test_revert_lengthMismatch() public {
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);
        allocations[0].operatorSets = new OperatorSet[](0);

        cheats.expectRevert(IAllocationManagerErrors.InputArrayLengthMismatch.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function test_revert_invalidOperatorSet() public {
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);

        // Set operatorSet to being invalid
        avsDirectoryMock.setIsOperatorSetBatch(allocations[0].operatorSets, false);

        cheats.expectRevert(IAllocationManagerErrors.InvalidOperatorSet.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function test_revert_invalidExpectedTotalMagnitude() public {
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);
        allocations[0].expectedMaxMagnitude = 1e18 + 1;

        cheats.expectRevert(IAllocationManagerErrors.InvalidExpectedTotalMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function test_revert_multiAlloc_modificationAlreadyPending_diffTx() public {
        // Allocate magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);
        cheats.startPrank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Warp to just before allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY - 1);

        // Attempt to allocate magnitude again
        cheats.expectRevert(IAllocationManagerErrors.ModificationAlreadyPending.selector);
        allocationManager.modifyAllocations(allocations);
        cheats.stopPrank();
    }

    function test_revert_multiAlloc_modificationAlreadyPending_sameTx() public {
        // Allocate magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            new IAllocationManagerTypes.MagnitudeAllocation[](2);
        allocations[0] = _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0)[0];
        allocations[1] = allocations[0];

        cheats.expectRevert(IAllocationManagerErrors.ModificationAlreadyPending.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function test_revert_allocateZeroMagnitude() public {
        // Allocate exact same magnitude as initial allocation (0)
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);
        allocations[0].magnitudes[0] = 0;

        cheats.expectRevert(IAllocationManagerErrors.SameMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function test_revert_allocateSameMagnitude() public {
        // Allocate nonzero magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Warp to allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Attempt to allocate no magnitude (ie. same magnitude)
        cheats.expectRevert(IAllocationManagerErrors.SameMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function testFuzz_revert_insufficientAllocatableMagnitude(uint256 r) public {
        // Allocate some magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_singleOpSet(r, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Warp to allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Attempt to allocate more magnitude than the operator has
        uint64 allocatedMag = allocations[0].magnitudes[0];
        allocations[0].magnitudes[0] = 1e18 + 1;
        cheats.expectRevert(IAllocationManagerErrors.InsufficientAllocatableMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function testFuzz_allocate_singleStrat_singleOperatorSet(uint256 r) public {
        // Create allocation
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_singleOpSet(r, 0);

        // Save vars to check against
        IStrategy strategy = allocations[0].strategy;
        uint64 magnitude = allocations[0].magnitudes[0];
        uint32 effectTimestamp = uint32(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Expect emits
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, strategy, magnitude);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSetMagnitudeUpdated(
            defaultOperator, allocations[0].operatorSets[0], strategy, magnitude, effectTimestamp
        );

        // Allocate magnitude
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

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
        MagnitudeInfo[] memory mInfos =
            allocationManager.getAllocationInfo(defaultOperator, strategy, allocations[0].operatorSets);
        assertEq(0, mInfos[0].currentMagnitude, "currentMagnitude should not be updated");
        assertEq(int128(uint128(magnitude)), mInfos[0].pendingDiff, "pendingMagnitude not updated");
        assertEq(effectTimestamp, mInfos[0].effectTimestamp, "effectTimestamp not updated");

        // Check storage after warp to completion
        cheats.warp(effectTimestamp);
        mInfos = allocationManager.getAllocationInfo(defaultOperator, strategy, allocations[0].operatorSets);
        assertEq(magnitude, mInfos[0].currentMagnitude, "currentMagnitude not updated");
        assertEq(0, mInfos[0].pendingDiff, "pendingMagnitude not updated");
        assertEq(0, mInfos[0].effectTimestamp, "effectTimestamp not updated");
    }

    function testFuzz_allocate_singleStrat_multipleSets(uint256 r) public {
        uint8 numOpSets = uint8(bound(r, 1, type(uint8).max));

        MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_multipleOpSets(r, 0, numOpSets);

        // Save vars to check against
        IStrategy strategy = allocations[0].strategy;
        uint64[] memory magnitudes = allocations[0].magnitudes;
        uint32 effectTimestamp = uint32(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Expect emits
        uint64 usedMagnitude;
        for (uint256 i = 0; i < numOpSets; i++) {
            usedMagnitude += magnitudes[i];
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit EncumberedMagnitudeUpdated(defaultOperator, strategy, usedMagnitude);
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit OperatorSetMagnitudeUpdated(
                defaultOperator, allocations[0].operatorSets[i], strategy, magnitudes[i], effectTimestamp
            );
        }
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

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
        MagnitudeInfo[] memory mInfos =
            allocationManager.getAllocationInfo(defaultOperator, strategy, allocations[0].operatorSets);
        for (uint256 i = 0; i < numOpSets; i++) {
            assertEq(0, mInfos[i].currentMagnitude, "currentMagnitude should not be updated");
            assertEq(int128(uint128(magnitudes[i])), mInfos[i].pendingDiff, "pendingMagnitude not updated");
            assertEq(effectTimestamp, mInfos[i].effectTimestamp, "effectTimestamp not updated");
        }

        // Check storage after warp to completion
        cheats.warp(effectTimestamp);
        mInfos = allocationManager.getAllocationInfo(defaultOperator, strategy, allocations[0].operatorSets);
        for (uint256 i = 0; i < numOpSets; i++) {
            assertEq(magnitudes[i], mInfos[i].currentMagnitude, "currentMagnitude not updated");
            assertEq(0, mInfos[i].pendingDiff, "pendingMagnitude not updated");
            assertEq(0, mInfos[i].effectTimestamp, "effectTimestamp not updated");
        }
    }

    function testFuzz_allocateMultipleTimes(
        uint256 r
    ) public {
        // Assumptions
        uint64 firstAlloc = uint64(bound(r, 1, type(uint64).max));
        uint64 secondAlloc = uint64(bound(r, 0, 1e18));
        cheats.assume(firstAlloc < secondAlloc);

        // Allocate magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _generateMagnitudeAllocationCalldata(defaultAVS, firstAlloc, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Warp to allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Allocate magnitude again
        allocations = _generateMagnitudeAllocationCalldata(defaultAVS, secondAlloc, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Check storage
        assertEq(
            secondAlloc,
            allocationManager.encumberedMagnitude(defaultOperator, allocations[0].strategy),
            "encumberedMagnitude not updated"
        );
    }

    function testFuzz_revert_overAllocate(
        uint256 r
    ) public {
        uint8 numOpSets = uint8(bound(r, 2, type(uint8).max));
        MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_multipleOpSets(r, 0, numOpSets);

        allocations[0].magnitudes[numOpSets - 1] = 1e18 + 1;

        // Overallocate
        cheats.expectRevert(IAllocationManagerErrors.InsufficientAllocatableMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function test_allocateMaxToMultipleStrategies() public {
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            new IAllocationManagerTypes.MagnitudeAllocation[](2);
        allocations[0] = _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0)[0];
        allocations[0].magnitudes[0] = 1e18;

        allocations[1] = _randomMagnitudeAllocation_singleStrat_singleOpSet(1, 1)[0];
        allocations[1].magnitudes[0] = 1e18;
        allocations[1].strategy = IStrategy(address(uint160(2))); // Set a different strategy

        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Assert maxMagnitude is encumbered
        assertEq(
            1e18,
            allocationManager.encumberedMagnitude(defaultOperator, allocations[0].strategy),
            "encumberedMagnitude not max"
        );
        assertEq(
            1e18,
            allocationManager.encumberedMagnitude(defaultOperator, allocations[1].strategy),
            "encumberedMagnitude not max"
        );
    }

    function test_revert_allocateDeallocate_modificationPending() public {
        // Allocate
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Deallocate
        allocations[0].magnitudes[0] -= 1;
        cheats.expectRevert(IAllocationManagerErrors.ModificationAlreadyPending.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function test_revert_deallocateTwice_modificationPending() public {
        // Allocate
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Warp past allocation complete timestsamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        allocations[0].magnitudes[0] -= 1;
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Deallocate again -> expect revert
        cheats.expectRevert(IAllocationManagerErrors.ModificationAlreadyPending.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    /**
     * Allocates to `firstMod` magnitude and then deallocate to `secondMod` magnitude
     * Validates the storage
     * - 1. After deallocation is alled
     * - 2. After the deallocationd elay is hit
     * - 3. After the modification queue is cleared
     */
    function testFuzz_allocate_deallocate(uint256 r) public {
        // Bound allocation and deallocation
        uint64 firstMod = uint64(bound(r, 1, 1e18));
        uint64 secondMod = uint64(bound(r, 0, firstMod - 1));

        // Allocate
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _generateMagnitudeAllocationCalldata(defaultAVS, firstMod, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Warp to allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        allocations = _generateMagnitudeAllocationCalldata(defaultAVS, secondMod, 1e18);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, strategyMock, firstMod);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSetMagnitudeUpdated(
            defaultOperator,
            allocations[0].operatorSets[0],
            strategyMock,
            secondMod,
            uint32(block.timestamp + DEALLOCATION_DELAY)
        );
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

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
        MagnitudeInfo[] memory mInfos =
            allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(firstMod, mInfos[0].currentMagnitude, "currentMagnitude should not be updated");
        int128 expectedDiff = -int128(uint128(firstMod - secondMod));
        assertEq(expectedDiff, mInfos[0].pendingDiff, "pendingMagnitude not updated");
        uint32 effectTimestamp = uint32(block.timestamp + DEALLOCATION_DELAY);
        assertEq(effectTimestamp, mInfos[0].effectTimestamp, "effectTimestamp not updated");

        // Check storage after warp to completion
        cheats.warp(effectTimestamp);
        mInfos =
            allocationManager.getAllocationInfo(defaultOperator, allocations[0].strategy, allocations[0].operatorSets);
        assertEq(secondMod, mInfos[0].currentMagnitude, "currentMagnitude not updated");
        assertEq(0, mInfos[0].pendingDiff, "pendingMagnitude not updated");
        assertEq(0, mInfos[0].effectTimestamp, "effectTimestamp not updated");
        assertEq(
            firstMod,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should not be updated"
        );

        // Check storage after clearing modification queue
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        uint16[] memory numToClear = new uint16[](1);
        numToClear[0] = 1;
        allocationManager.clearModificationQueue(defaultOperator, strategies, numToClear);
        assertEq(
            secondMod,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should be updated"
        );
    }

    function test_deallocate_all() public {
        // Allocate
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations =
            _generateMagnitudeAllocationCalldata(defaultAVS, 1e18, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Warp to allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        allocations[0].magnitudes[0] = 0;
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Warp to completion and clear modification queue
        cheats.warp(block.timestamp + DEALLOCATION_DELAY);
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        uint16[] memory numToClear = new uint16[](1);
        numToClear[0] = 1;
        allocationManager.clearModificationQueue(defaultOperator, strategies, numToClear);

        // Check storage
        assertEq(
            0,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should be updated"
        );
        MagnitudeInfo[] memory mInfos =
            allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(0, mInfos[0].currentMagnitude, "currentMagnitude should be 0");
        assertEq(0, mInfos[0].pendingDiff, "pendingMagnitude should be 0");
        assertEq(0, mInfos[0].effectTimestamp, "effectTimestamp should be 0");
    }

    function testFuzz_allocate_deallocate_singleStrat_multipleOperatorSets(
        uint256 r
    ) public {
        uint8 numOpSets = uint8(bound(r, 0, type(uint8).max));
        (MagnitudeAllocation[] memory allocations, MagnitudeAllocation[] memory deallocations) =
            _randomAllocationAndDeallocation_singleStrat_multipleOpSets(numOpSets, r, 0);


        // Allocate
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
        uint64 encumberedMagnitudeAfterAllocation =
            allocationManager.encumberedMagnitude(defaultOperator, allocations[0].strategy);

        // Warp to allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        uint64 postDeallocMag;
        for (uint256 i = 0; i < numOpSets; i++) {
            postDeallocMag += deallocations[0].magnitudes[i];
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit EncumberedMagnitudeUpdated(
                defaultOperator, deallocations[0].strategy, encumberedMagnitudeAfterAllocation
            );
            // pendingNewMags[i] = allocations[0].magnitudes[i] - deallocations[0].magnitudes[i];
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit OperatorSetMagnitudeUpdated(
                defaultOperator,
                deallocations[0].operatorSets[i],
                deallocations[0].strategy,
                deallocations[0].magnitudes[i],
                uint32(block.timestamp + DEALLOCATION_DELAY)
            );
        }
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(deallocations);

        // Check storage after dealloc
        assertEq(
            encumberedMagnitudeAfterAllocation,
            allocationManager.encumberedMagnitude(defaultOperator, allocations[0].strategy),
            "encumberedMagnitude should not be updated"
        );
        MagnitudeInfo[] memory mInfos =
            allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        for (uint256 i = 0; i < mInfos.length; i++) {
            assertEq(allocations[0].magnitudes[i], mInfos[i].currentMagnitude, "currentMagnitude should not be updated");
            int128 expectedDiff = -int128(uint128(allocations[0].magnitudes[i] - deallocations[0].magnitudes[i]));
            assertEq(expectedDiff, mInfos[i].pendingDiff, "pendingMagnitude not updated");
            uint32 effectTimestamp = uint32(block.timestamp + DEALLOCATION_DELAY);
            assertEq(effectTimestamp, mInfos[i].effectTimestamp, "effectTimestamp not updated");
        }

        // Check storage after warp to completion
        cheats.warp(block.timestamp + DEALLOCATION_DELAY);
        mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        for (uint256 i = 0; i < mInfos.length; i++) {
            assertEq(deallocations[0].magnitudes[i], mInfos[i].currentMagnitude, "currentMagnitude not updated");
            assertEq(0, mInfos[i].pendingDiff, "pendingMagnitude not updated");
            assertEq(0, mInfos[i].effectTimestamp, "effectTimestamp not updated");
        }

        // Clear modification queue
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        uint16[] memory numToClear = new uint16[](1);
        numToClear[0] = numOpSets;
        allocationManager.clearModificationQueue(defaultOperator, strategies, numToClear);

        // Check storage after clearing modification queue
        assertEq(
            postDeallocMag,
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should be updated"
        );
    }

    /**
     * Allocates, deallocates, and then allocates again
     */
    function test_regression_allocate_deallocate_allocate() public {
        uint32 allocationDelay = 15 days;
        // Set allocation delay to be 15 days
        cheats.prank(defaultOperator);
        allocationManager.setAllocationDelay(allocationDelay);
        cheats.warp(block.timestamp + ALLOCATION_CONFIGURATION_DELAY);
        (,uint32 storedDelay) = allocationManager.getAllocationDelay(defaultOperator);
        assertEq(allocationDelay, storedDelay, "allocation delay not valid");

        // Allocate half of mag to opset1
        IAllocationManagerTypes.MagnitudeAllocation[] memory firstAllocation =
            _generateMagnitudeAllocationCalldataForOpSet(defaultAVS, 1, 5e17, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(firstAllocation);
        cheats.warp(block.timestamp + 15 days);

        // Deallocate half from opset1.
        uint32 t = uint32(block.timestamp);
        uint32 deallocationEffectTimestamp = uint32(block.timestamp + DEALLOCATION_DELAY);
        IAllocationManagerTypes.MagnitudeAllocation[] memory firstDeallocation =
            _generateMagnitudeAllocationCalldataForOpSet(defaultAVS, 1, 25e16, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(firstDeallocation);
        MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, firstDeallocation[0].operatorSets);
        assertEq(deallocationEffectTimestamp, mInfos[0].effectTimestamp, "effect timestamp not correct");
        
        // Allocate 33e16 mag to opset2
        uint32 allocationEffectTimestamp = uint32(block.timestamp + allocationDelay);
        IAllocationManagerTypes.MagnitudeAllocation[] memory secondAllocation =
            _generateMagnitudeAllocationCalldataForOpSet(defaultAVS, 2, 33e16, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(secondAllocation);
        mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, secondAllocation[0].operatorSets);
        console.log("deallocation effect timestamp: ", deallocationEffectTimestamp);
        console.log("allocation effect timestamp: ", allocationEffectTimestamp);
        assertEq(allocationEffectTimestamp, mInfos[0].effectTimestamp, "effect timestamp not correct");
        assertLt(allocationEffectTimestamp, deallocationEffectTimestamp, "invalid test setup");

        // Warp to allocation effect timestamp & clear the queue
        cheats.warp(allocationEffectTimestamp);
        allocationManager.clearModificationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());

        // Validate `getAllocatableMagnitude`. Allocatable magnitude should be the difference between the total magnitude and the encumbered magnitude
        uint64 allocatableMagnitude = allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock);
        assertEq(WAD - 33e16 - 5e17, allocatableMagnitude, "allocatableMagnitude not correct");

        // Validate that we can allocate again for opset2. This should revert?
        IAllocationManagerTypes.MagnitudeAllocation[] memory thirdAllocation =
            _generateMagnitudeAllocationCalldataForOpSet(defaultAVS, 2, 10e16, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(thirdAllocation);
    }
}

contract AllocationManagerUnitTests_ClearModificationQueue is AllocationManagerUnitTests {
    /// -----------------------------------------------------------------------
    /// clearModificationQueue()
    /// -----------------------------------------------------------------------

    function test_revert_paused() public {
        allocationManager.pause(2 ** PAUSED_MODIFY_ALLOCATIONS);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        allocationManager.clearModificationQueue(defaultOperator, new IStrategy[](0), new uint16[](0));
    }

    function test_revert_arrayMismatch() public {
        IStrategy[] memory strategies = new IStrategy[](1);
        uint16[] memory numToClear = new uint16[](2);

        cheats.expectRevert(IAllocationManagerErrors.InputArrayLengthMismatch.selector);
        allocationManager.clearModificationQueue(defaultOperator, strategies, numToClear);
    }

    function test_revert_operatorNotRegistered() public {
        // Deregister operator
        delegationManagerMock.setIsOperator(defaultOperator, false);

        cheats.expectRevert(IAllocationManagerErrors.OperatorNotRegistered.selector);
        allocationManager.clearModificationQueue(defaultOperator, new IStrategy[](0), new uint16[](0));
    }

    /**
     * @notice Allocates magnitude to an operator and then
     * - Clears modification queue when nothing can be completed
     * - Clears modification queue when the alloc can be completed - asserts emit has been emitted
     * - Validates storage after the second clear
     */
    function testFuzz_allocate(
        uint256 r
    ) public {
        // Allocate magnitude
        IAllocationManager.MagnitudeAllocation[] memory allocations =
            _queueRandomAllocation_singleStrat_singleOpSet(defaultOperator, r, 0);

        // Attempt to clear queue, assert no events emitted
        allocationManager.clearModificationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());
        Vm.Log[] memory entries = vm.getRecordedLogs();
        assertEq(0, entries.length, "should not have emitted any events");

        // Warp to allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Clear queue
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, strategyMock, allocations[0].magnitudes[0]);
        allocationManager.clearModificationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());

        // Validate storage (although this is technically tested in allocation tests, adding for sanity)
        // TODO: maybe add a harness here to actually introspect storage
        IAllocationManager.MagnitudeInfo[] memory mInfos =
            allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(allocations[0].magnitudes[0], mInfos[0].currentMagnitude, "currentMagnitude should be 0");
        assertEq(0, mInfos[0].pendingDiff, "pendingMagnitude should be 0");
        assertEq(0, mInfos[0].effectTimestamp, "effectTimestamp should be 0");
    }

    /**
     * @notice Allocates magnitude to an operator and then
     * - Clears modification queue when nothing can be completed
     * - After the first clear, asserts the allocation info takes into account the deallocation
     * - Clears modification queue when the dealloc can be completed
     * - Assert events & validates storage after the deallocations are completed
     */
    function testFuzz_allocate_deallocate(uint256 r) public {
        // Complete allocations & add a deallocation
        (MagnitudeAllocation[] memory allocations, MagnitudeAllocation[] memory deallocations) =
        _queueRandomAllocationAndDeallocation(
            defaultOperator,
            1, // numOpSets
            r,
            0 // salt
        );

        // Clear queue & check storage
        allocationManager.clearModificationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());
        assertEq(
            allocations[0].magnitudes[0],
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should not be updated"
        );

        // Validate storage - encumbered magnitude should just be allocations (we only have 1 allocation)
        IAllocationManager.MagnitudeInfo[] memory mInfos =
            allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        int128 pendingDiff = -int128(uint128(allocations[0].magnitudes[0] - deallocations[0].magnitudes[0]));
        assertEq(allocations[0].magnitudes[0], mInfos[0].currentMagnitude, "currentMagnitude should be 0");
        assertEq(pendingDiff, mInfos[0].pendingDiff, "pendingMagnitude should be 0");
        assertEq(block.timestamp + DEALLOCATION_DELAY, mInfos[0].effectTimestamp, "effectTimestamp should be 0");

        // Warp to deallocation complete timestamp
        cheats.warp(block.timestamp + DEALLOCATION_DELAY);

        // Clear queue
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, strategyMock, deallocations[0].magnitudes[0]);
        allocationManager.clearModificationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());

        // Validate storage - encumbered magnitude should just be deallocations (we only have 1 deallocation)
        assertEq(
            deallocations[0].magnitudes[0],
            allocationManager.encumberedMagnitude(defaultOperator, strategyMock),
            "encumberedMagnitude should be updated"
        );
        mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, deallocations[0].operatorSets);
        assertEq(deallocations[0].magnitudes[0], mInfos[0].currentMagnitude, "currentMagnitude should be 0");
        assertEq(0, mInfos[0].pendingDiff, "pendingMagnitude should be 0");
        assertEq(0, mInfos[0].effectTimestamp, "effectTimestamp should be 0");
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
        cheats.expectRevert(IAllocationManagerErrors.OperatorNotRegistered.selector);
        allocationManager.setAllocationDelay(1);
    }

    function test_revert_zeroAllocationDelay() public {
        cheats.expectRevert(IAllocationManagerErrors.InvalidAllocationDelay.selector);
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(0);
    }

    function testFuzz_setDelay(
        uint256 r
    ) public {
        uint32 delay = uint32(bound(r, 1, type(uint32).max));

        // Set delay
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AllocationDelaySet(operatorToSet, delay, uint32(block.timestamp + ALLOCATION_CONFIGURATION_DELAY));
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(delay);

        // Check values after set
        (bool isSet, uint32 returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertFalse(isSet, "isSet should not be set");
        assertEq(0, returnedDelay, "returned delay should be 0");

        // Warp to effect timestamp
        cheats.warp(block.timestamp + ALLOCATION_CONFIGURATION_DELAY);

        // Check values after config delay
        (isSet, returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(delay, returnedDelay, "delay not set");
    }

    function test_fuzz_setDelay_multipleTimesWithinConfigurationDelay(
        uint32 firstDelay, uint32 secondDelay
    ) public {
        firstDelay = uint32(bound(firstDelay, 1, type(uint32).max));
        secondDelay = uint32(bound(secondDelay, 1, type(uint32).max));
        cheats.assume(firstDelay != secondDelay);

        // Set delay
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(firstDelay);

        // Warp just before effect timestamp
        cheats.warp(block.timestamp + ALLOCATION_CONFIGURATION_DELAY - 1);

        // Set delay again
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit AllocationDelaySet(operatorToSet, secondDelay, uint32(block.timestamp + ALLOCATION_CONFIGURATION_DELAY));
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(secondDelay);

        // Warp to effect timestamp of first delay
        cheats.warp(block.timestamp + 1);

        // Assert that the delay is still not set
        (bool isSet, uint32 returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertFalse(isSet, "isSet should not be set");
        assertEq(0, returnedDelay, "returned delay should be 0");

        // Warp to effect timestamp of second delay
        cheats.warp(block.timestamp + ALLOCATION_CONFIGURATION_DELAY);
        (isSet, returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(secondDelay, returnedDelay, "delay not set");
    }

    function testFuzz_multipleDelays(
        uint32 firstDelay, uint32 secondDelay
    ) public {
        firstDelay = uint32(bound(firstDelay, 1, type(uint32).max));
        secondDelay = uint32(bound(secondDelay, 1, type(uint32).max));
        cheats.assume(firstDelay != secondDelay);

        // Set delay
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(firstDelay);

        // Warp to effect timestamp of first delay
        cheats.warp(block.timestamp + ALLOCATION_CONFIGURATION_DELAY);

        // Set delay again
        cheats.prank(operatorToSet);
        allocationManager.setAllocationDelay(secondDelay);

        // Assert that first delay is set
        (bool isSet, uint32 returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(firstDelay, returnedDelay, "delay not set");

        // Warp to effect timestamp of second delay
        cheats.warp(block.timestamp + ALLOCATION_CONFIGURATION_DELAY);

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

        // Warp to effect timestamp
        cheats.warp(block.timestamp + ALLOCATION_CONFIGURATION_DELAY);
        (bool isSet, uint32 returnedDelay) = allocationManager.getAllocationDelay(operatorToSet);
        assertTrue(isSet, "isSet should be set");
        assertEq(delay, returnedDelay, "delay not set");
    }
}

/**
 * @notice TODO Lifecycle tests - These tests combine multiple functionalities of the AllocationManager
 * 1. Set allocation delay > 21 days (configuration), Allocate, modify allocation delay to < 21 days, try to allocate again once new delay is set (should be able to allocate faster than 21 deays)
 * 2. Allocate across multiple strategies and multiple operatorSets
 * 3. lifecycle fuzz test allocating/deallocating across multiple opSets/strategies
 * 4. HIGH PRIO - add uint16.max allocations/deallocations and then clear them
 * 5. Correctness of slashable magnitudes
 * 6. HIGH PRIO - get gas costs of `getSlashableMagnitudes`
 */
