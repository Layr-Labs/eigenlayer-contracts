// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/core/AllocationManager.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

contract AllocationManagerUnitTests is EigenLayerUnitTestSetup, IAllocationManagerErrors, IAllocationManagerEvents{
    uint8 internal constant PAUSED_MODIFY_ALLOCATIONS = 0;
    uint8 internal constant PAUSED_OPERATOR_SLASHING = 1;
    uint32 constant DEALLOCATION_DELAY = 17.5 days;
    uint32 constant ALLOCATION_CONFIGURATION_DELAY = 21 days;

    AllocationManager allocationManager;
    ERC20PresetFixedSupply tokenMock;
    StrategyBase strategyMock;

    address defaultOperator = address(this);
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

    function _randomAddr(uint256 r, uint256 salt) internal pure returns (address addr) {
        /// @solidity memory-safe-assembly
        assembly {
            mstore(0x00, r)
            mstore(0x20, salt)
            addr := keccak256(0x00, 0x40)
        }
    }

    function _randomSlashingParams(
        uint256 r,
        uint256 salt
    ) internal view returns (IAllocationManagerTypes.SlashingParams memory) {
        r = uint256(keccak256(abi.encodePacked(r, salt)));

        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;

        return IAllocationManagerTypes.SlashingParams({
            operator: _randomAddr(r, 0),
            operatorSetId: uint32(r),
            strategies: strategies,
            wadToSlash: bound(r, 1, 1e18),
            description: "test"
        });
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
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _randomMagnitudeAllocation_singleStrat_singleOpSet(r, salt);
        cheats.prank(operator);
        allocationManager.modifyAllocations(allocations);

        return allocations;
    }

    /**
     * @notice Queued a random allocation and deallocation for the given `operator`
     * - DOES NOT warp past the deallocation effect timestamp
     */
    function _queueRandomAllocationAndDeallocation(
        address operator,
        uint256 r,
        uint256 salt,
        uint8 numOpSets
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory, IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        (MagnitudeAllocation[] memory allocations, MagnitudeAllocation[] memory deallocations) = 
            _randomAllocationAndDeallocation_singleStrat_multipleOpSets(r, salt, numOpSets);

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
     * @notice Generated magnitudeAllocation calldata for the `strategyMock` with a provided magnitude.
     */
    function _generateMagnitudeAllocationCalldata(
        uint64 magnitudeToSet, 
        uint64 expectedMaxMagnitude
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        OperatorSet[] memory operatorSets = new OperatorSet[](1);
        operatorSets[0] = OperatorSet({avs: _randomAddr(0, 0), operatorSetId: 1});

        // Set operatorSet to being valid
        avsDirectoryMock.setIsOperatorSetBatch(operatorSets, true);

        uint64[] memory magnitudes = new uint64[](1);
        magnitudes[0] = magnitudeToSet;

        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = new IAllocationManagerTypes.MagnitudeAllocation[](1);
        allocations[0] = IAllocationManagerTypes.MagnitudeAllocation({
            strategy: strategyMock,
            expectedMaxMagnitude: expectedMaxMagnitude,
            operatorSets: operatorSets,
            magnitudes: magnitudes
        });

        return allocations;
    }

    /**
     * @notice Create a random magnitude allocation.
     * In addition
     * - Registers the operatorSet with the avsDirectory
     */
    function _randomMagnitudeAllocation_singleStrat_singleOpSet(
        uint256 r,
        uint256 salt
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        r = uint256(keccak256(abi.encodePacked(r, salt)));

        // Mock a random operator set.
        OperatorSet[] memory operatorSets = new OperatorSet[](1);
        operatorSets[0] = OperatorSet({avs: _randomAddr(r, 0), operatorSetId: uint32(r)});

        // Set operatorSet to being valid
        avsDirectoryMock.setIsOperatorSetBatch(operatorSets, true);

        uint64[] memory magnitudes = new uint64[](1);
        magnitudes[0] = uint64(bound(r, 1, 1e18));

        // Mock a random magnitude allocation.
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = new IAllocationManagerTypes.MagnitudeAllocation[](1);
        allocations[0] = IAllocationManagerTypes.MagnitudeAllocation({
            strategy: strategyMock,
            expectedMaxMagnitude: 1e18, // magnitude starts at 100%
            operatorSets: operatorSets,
            magnitudes: magnitudes
        });
        return allocations;
    }

    function _randomAllocationAndDeallocation_singleStrat_multipleOpSets(
        uint256 r,
        uint256 salt,
        uint8 numOpSets
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory, IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = new IAllocationManagerTypes.MagnitudeAllocation[](1);
        allocations = _randomMagnitudeAllocation_singleStrat_multipleOpSets(r, salt, numOpSets);

        // Deallocate random magnitude from each of thsoe operatorSets
        r = uint256(keccak256(abi.encodePacked(r, salt)));
        uint64[] memory newMags = new uint64[](numOpSets);
        for(uint8 i = 0; i < numOpSets; i++) {
            newMags[i] = uint64(bound(r, 0, allocations[0].magnitudes[i] - 1));
        }

        // Create deallocations
        IAllocationManagerTypes.MagnitudeAllocation[] memory deallocations = new IAllocationManagerTypes.MagnitudeAllocation[](1);
        deallocations[0] = IAllocationManagerTypes.MagnitudeAllocation({
            strategy: strategyMock,
            expectedMaxMagnitude: 1e18, // magnitude starts at 100%
            operatorSets: allocations[0].operatorSets,
            magnitudes: newMags
        });

        return (allocations, deallocations);
    }

    function _randomMagnitudeAllocation_singleStrat_multipleOpSets(
        uint256 r,
        uint256 salt,
        uint8 numOpSets
    ) internal returns (IAllocationManagerTypes.MagnitudeAllocation[] memory) {
        r = uint256(keccak256(abi.encodePacked(r, salt)));

        // Create multiple operatorSets
        OperatorSet[] memory operatorSets = new OperatorSet[](numOpSets);
        for(uint8 i = 0; i < numOpSets; i++) {
            operatorSets[i] = OperatorSet({avs: _randomAddr(r, i), operatorSetId: uint32(r + i)});
        }
        avsDirectoryMock.setIsOperatorSetBatch(operatorSets, true);

        // Give each set a minimum of 1 magnitude
        uint64[] memory magnitudes = new uint64[](numOpSets);
        uint64 usedMagnitude;
        for(uint8 i = 0; i < numOpSets; i++) {
            magnitudes[i] = 1;
            usedMagnitude++;
        }

        // Distribute remaining magnitude
        uint64 maxMagnitude = 1e18;
        for(uint8 i = 0; i < numOpSets; i++) {
            r = uint256(keccak256(abi.encodePacked(r, i)));
            uint64 remainingMagnitude = maxMagnitude - usedMagnitude;
            if (remainingMagnitude > 0) {
                magnitudes[i] += uint64(bound(r, 0, remainingMagnitude));
                usedMagnitude += magnitudes[i] - 1;
            }
        }

        // Create magnitude allocation
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = new IAllocationManagerTypes.MagnitudeAllocation[](1);
        allocations[0] = IAllocationManagerTypes.MagnitudeAllocation({
            strategy: strategyMock,
            expectedMaxMagnitude: 1e18, // magnitude starts at 100%
            operatorSets: operatorSets,
            magnitudes: magnitudes
        });
        return allocations;
    }

    // function _getMockStrategyMagnitudeAllocation(uint256 magnitude) internal view returns (IAllocationManagerTypes.MagnitudeAllocation[] memory) {

    // }
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

    /// @dev Asserts the following:
    /// 1. Fn can only be called if `IAVSDirectory.isOperatorSlashable(operator)` returns true.
    /// 2. Fn can only be called if `wadToSlash` is no greater than 100% (1e18) and no less than 0.000000000000000001% (1).
    /// 3. Fn can only be called when unpaused.
    function testFuzz_SlashOperator_Checks(
        uint256 r
    ) public {
        // Mock random slashing params, and operator set.
        IAllocationManagerTypes.SlashingParams memory p = _randomSlashingParams(r, 0);
        OperatorSet memory operatorSet = OperatorSet({avs: _randomAddr(r, 0), operatorSetId: p.operatorSetId});

        // Assert that operator's cannot be  slashed unless `IAVSDirectory.isOperatorSlashable` returns true.
        cheats.prank(operatorSet.avs);
        cheats.expectRevert(IAllocationManagerErrors.InvalidOperator.selector);
        allocationManager.slashOperator(p);

        // Mock  `IAVSDirectory.isOperatorSlashable` returning true.
        avsDirectoryMock.setIsOperatorSlashable(p.operator, operatorSet, true);

        // Assert `wadToSlash` cannot be greater than 100% (1e18).
        p.wadToSlash = 1e18 + 1;
        cheats.prank(operatorSet.avs);
        cheats.expectRevert(IAllocationManagerErrors.InvalidWadToSlash.selector);
        allocationManager.slashOperator(p);

        // Assert `wadToSlash` cannot be 0.
        p.wadToSlash = 0;
        cheats.prank(operatorSet.avs);
        cheats.expectRevert(IAllocationManagerErrors.InvalidWadToSlash.selector);
        allocationManager.slashOperator(p);

        // Pause operator slashing.
        allocationManager.pause(2 ** PAUSED_OPERATOR_SLASHING);

        // Assert that fn can only be called when unpaused.
        cheats.prank(operatorSet.avs);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        allocationManager.slashOperator(p);

        // Unpause operator slashing.
        cheats.prank(unpauser);
        allocationManager.unpause(2 ** PAUSED_OPERATOR_SLASHING);

        // TODO
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
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);
        allocations[0].operatorSets = new OperatorSet[](0);

        cheats.expectRevert(IAllocationManagerErrors.InputArrayLengthMismatch.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function test_revert_invalidOperatorSet() public {
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0); 
        
        // Set operatorSet to being invalid
        avsDirectoryMock.setIsOperatorSetBatch(allocations[0].operatorSets, false);

        cheats.expectRevert(IAllocationManagerErrors.InvalidOperatorSet.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function test_revert_invalidExpectedTotalMagnitude() public {
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0); 
        allocations[0].expectedMaxMagnitude = 1e18 + 1;

        cheats.expectRevert(IAllocationManagerErrors.InvalidExpectedTotalMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function test_revert_multiAlloc_modificationAlreadyPending_diffTx() public {
        // Allocate magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);
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
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = new IAllocationManagerTypes.MagnitudeAllocation[](2);
        allocations[0] = _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0)[0];
        allocations[1] = allocations[0];

        cheats.expectRevert(IAllocationManagerErrors.ModificationAlreadyPending.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function test_revert_allocateZeroMagnitude() public {
        // Allocate exact same magnitude as initial allocation (0)
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);
        allocations[0].magnitudes[0] = 0;

        cheats.expectRevert(IAllocationManagerErrors.SameMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function test_revert_allocateSameMagnitude() public {
        // Allocate nonzero magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Warp to allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Attempt to allocate no magnitude (ie. same magnitude)
        cheats.expectRevert(IAllocationManagerErrors.SameMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function testFuzz_revert_insufficientAllocatableMagnitude(uint256 r, uint256 salt) public {
        // Allocate some magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _randomMagnitudeAllocation_singleStrat_singleOpSet(r, salt);
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

    function testFuzz_allocate_singleStrat_singleOperatorSet(uint256 r, uint256 salt) public {
        // Create allocation
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _randomMagnitudeAllocation_singleStrat_singleOpSet(r, salt); 

        // Save vars to check against
        IStrategy strategy = allocations[0].strategy;
        uint64 magnitude = allocations[0].magnitudes[0];
        uint32 effectTimestamp = uint32(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Expect emits
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, strategy, magnitude);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSetMagnitudeUpdated(defaultOperator, allocations[0].operatorSets[0], strategy, magnitude, effectTimestamp);

        // Allocate magnitude
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Check storage
        assertEq(magnitude, allocationManager.encumberedMagnitude(defaultOperator, strategy), "encumberedMagnitude not updated");
        assertEq(WAD - magnitude, allocationManager.getAllocatableMagnitude(defaultOperator, strategy), "allocatableMagnitude not calcualted correctly");
        MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategy, allocations[0].operatorSets);
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

    function testFuzz_allocate_singleStrat_multipleSets(uint256 r, uint256 salt, uint8 numOpSets) public {
        cheats.assume(numOpSets > 1);

        MagnitudeAllocation[] memory allocations = _randomMagnitudeAllocation_singleStrat_multipleOpSets(r, salt, numOpSets);  

        // Save vars to check against
        IStrategy strategy = allocations[0].strategy;
        uint64[] memory magnitudes = allocations[0].magnitudes;
        uint32 effectTimestamp = uint32(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Expect emits
        uint64 usedMagnitude;
        for(uint256 i = 0; i < numOpSets; i++) {
            usedMagnitude += magnitudes[i];
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit EncumberedMagnitudeUpdated(defaultOperator, strategy, usedMagnitude);
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit OperatorSetMagnitudeUpdated(defaultOperator, allocations[0].operatorSets[i], strategy, magnitudes[i], effectTimestamp);
        }
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Check storage
        assertEq(usedMagnitude, allocationManager.encumberedMagnitude(defaultOperator, strategy), "encumberedMagnitude not updated");
        assertEq(WAD - usedMagnitude, allocationManager.getAllocatableMagnitude(defaultOperator, strategy), "allocatableMagnitude not calcualted correctly");
        MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategy, allocations[0].operatorSets);
        for(uint256 i = 0; i < numOpSets; i++) {
            assertEq(0, mInfos[i].currentMagnitude, "currentMagnitude should not be updated");
            assertEq(int128(uint128(magnitudes[i])), mInfos[i].pendingDiff, "pendingMagnitude not updated");
            assertEq(effectTimestamp, mInfos[i].effectTimestamp, "effectTimestamp not updated");
        }

        // Check storage after warp to completion
        cheats.warp(effectTimestamp);
        mInfos = allocationManager.getAllocationInfo(defaultOperator, strategy, allocations[0].operatorSets);
        for(uint256 i = 0; i < numOpSets; i++) {
            assertEq(magnitudes[i], mInfos[i].currentMagnitude, "currentMagnitude not updated");
            assertEq(0, mInfos[i].pendingDiff, "pendingMagnitude not updated");
            assertEq(0, mInfos[i].effectTimestamp, "effectTimestamp not updated");
        }
    }

    function testFuzz_allocateMultipleTimes(uint64 firstAlloc, uint64 secondAlloc) public {
        // Assumptions
        cheats.assume(firstAlloc > 0);
        cheats.assume(firstAlloc < secondAlloc);
        cheats.assume(secondAlloc <= 1e18);

        // Allocate magnitude
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _generateMagnitudeAllocationCalldata(firstAlloc, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Warp to allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Allocate magnitude again
        allocations = _generateMagnitudeAllocationCalldata(secondAlloc, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Check storage
        assertEq(secondAlloc, allocationManager.encumberedMagnitude(defaultOperator, allocations[0].strategy), "encumberedMagnitude not updated");
    }

    function testFuzz_revert_overAllocate(uint256 r, uint256 salt, uint8 numOpSets) public {
        cheats.assume(numOpSets > 1);
        MagnitudeAllocation[] memory allocations = _randomMagnitudeAllocation_singleStrat_multipleOpSets(r, salt, numOpSets);  

        allocations[0].magnitudes[numOpSets - 1] = 1e18 + 1;

        // Overallocate
        cheats.expectRevert(IAllocationManagerErrors.InsufficientAllocatableMagnitude.selector);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
    }

    function test_allocateMaxToMultipleStrategies() public {
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = new IAllocationManagerTypes.MagnitudeAllocation[](2);
        allocations[0] = _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0)[0];
        allocations[0].magnitudes[0] = 1e18;

        allocations[1] = _randomMagnitudeAllocation_singleStrat_singleOpSet(1, 1)[0];
        allocations[1].magnitudes[0] = 1e18;
        allocations[1].strategy = IStrategy(address(uint160(2))); // Set a different strategy
        
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Assert maxMagnitude is encumbered
        assertEq(1e18, allocationManager.encumberedMagnitude(defaultOperator, allocations[0].strategy), "encumberedMagnitude not max");
        assertEq(1e18, allocationManager.encumberedMagnitude(defaultOperator, allocations[1].strategy), "encumberedMagnitude not max");
    }

    function test_revert_allocateDeallocate_modificationPending() public {
        // Allocate
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);
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
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _randomMagnitudeAllocation_singleStrat_singleOpSet(0, 0);
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
    function testFuzz_allocate_deallocate(uint64 firstMod, uint64 secondMod) public {
        // Bound allocation and deallocation
        firstMod = uint64(bound(firstMod, 1, 1e18));
        secondMod = uint64(bound(secondMod, 0, firstMod - 1));

        // Allocate
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _generateMagnitudeAllocationCalldata(firstMod, 1e18);
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Warp to allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        allocations = _generateMagnitudeAllocationCalldata(secondMod, 1e18);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit EncumberedMagnitudeUpdated(defaultOperator, strategyMock, firstMod);
        cheats.expectEmit(true, true, true, true, address(allocationManager));
        emit OperatorSetMagnitudeUpdated(defaultOperator, allocations[0].operatorSets[0], strategyMock, secondMod, uint32(block.timestamp + DEALLOCATION_DELAY));
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);

        // Check storage after dealloc
        assertEq(firstMod, allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude should not be updated");
        assertEq(WAD - firstMod, allocationManager.getAllocatableMagnitude(defaultOperator, strategyMock), "allocatableMagnitude not calcualted correctly");
        MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(firstMod, mInfos[0].currentMagnitude, "currentMagnitude should not be updated");
        int128 expectedDiff = -int128(uint128(firstMod - secondMod));
        assertEq(expectedDiff, mInfos[0].pendingDiff, "pendingMagnitude not updated");
        uint32 effectTimestamp = uint32(block.timestamp + DEALLOCATION_DELAY);
        assertEq(effectTimestamp, mInfos[0].effectTimestamp, "effectTimestamp not updated");

        // Check storage after warp to completion
        cheats.warp(effectTimestamp);
        mInfos = allocationManager.getAllocationInfo(defaultOperator, allocations[0].strategy, allocations[0].operatorSets);
        assertEq(secondMod, mInfos[0].currentMagnitude, "currentMagnitude not updated");
        assertEq(0, mInfos[0].pendingDiff, "pendingMagnitude not updated");
        assertEq(0, mInfos[0].effectTimestamp, "effectTimestamp not updated");
        assertEq(firstMod, allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude should not be updated");

        // Check storage after clearing modification queue
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        uint16[] memory numToClear = new uint16[](1);
        numToClear[0] = 1;
        allocationManager.clearModificationQueue(defaultOperator, strategies, numToClear);
        assertEq(secondMod, allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude should be updated");
    }

    function test_deallocate_all() public {
        // Allocate
        IAllocationManagerTypes.MagnitudeAllocation[] memory allocations = _generateMagnitudeAllocationCalldata(1e18, 1e18);
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
        assertEq(0, allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude should be updated");
        MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        assertEq(0, mInfos[0].currentMagnitude, "currentMagnitude should be 0");
        assertEq(0, mInfos[0].pendingDiff, "pendingMagnitude should be 0");
        assertEq(0, mInfos[0].effectTimestamp, "effectTimestamp should be 0");
    }

    function testFuzz_allocate_deallocate_singleStrat_multipleOperatorSets(uint256 r, uint256 salt, uint8 numOpSets) public {
        (MagnitudeAllocation[] memory allocations, MagnitudeAllocation[] memory deallocations) = 
            _randomAllocationAndDeallocation_singleStrat_multipleOpSets(r, salt, numOpSets);

        // Allocate
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(allocations);
        uint64 encumberedMagnitudeAfterAllocation = allocationManager.encumberedMagnitude(defaultOperator, allocations[0].strategy);

        // Warp to allocation complete timestamp
        cheats.warp(block.timestamp + DEFAULT_OPERATOR_ALLOCATION_DELAY);

        // Deallocate
        uint64 postDeallocMag;
        for(uint256 i = 0; i < numOpSets; i++) {
            postDeallocMag += deallocations[0].magnitudes[i];
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit EncumberedMagnitudeUpdated(defaultOperator, deallocations[0].strategy, encumberedMagnitudeAfterAllocation);
            // pendingNewMags[i] = allocations[0].magnitudes[i] - deallocations[0].magnitudes[i];
            cheats.expectEmit(true, true, true, true, address(allocationManager));
            emit OperatorSetMagnitudeUpdated(defaultOperator, deallocations[0].operatorSets[i], deallocations[0].strategy, deallocations[0].magnitudes[i], uint32(block.timestamp + DEALLOCATION_DELAY));
        }
        cheats.prank(defaultOperator);
        allocationManager.modifyAllocations(deallocations);

        // Check storage after dealloc
        assertEq(encumberedMagnitudeAfterAllocation, allocationManager.encumberedMagnitude(defaultOperator, allocations[0].strategy), "encumberedMagnitude should not be updated");
        MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        for(uint256 i = 0; i < mInfos.length; i++) {
            assertEq(allocations[0].magnitudes[i], mInfos[i].currentMagnitude, "currentMagnitude should not be updated");
            int128 expectedDiff = -int128(uint128(allocations[0].magnitudes[i] - deallocations[0].magnitudes[i]));
            assertEq(expectedDiff, mInfos[i].pendingDiff, "pendingMagnitude not updated");
            uint32 effectTimestamp = uint32(block.timestamp + DEALLOCATION_DELAY);
            assertEq(effectTimestamp, mInfos[i].effectTimestamp, "effectTimestamp not updated");
        }

        // Check storage after warp to completion
        cheats.warp(block.timestamp + DEALLOCATION_DELAY);
        mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
        for(uint256 i = 0; i < mInfos.length; i++) {
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
        assertEq(postDeallocMag, allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude should be updated");
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
    function testFuzz_allocate(uint256 r, uint256 salt) public {
        // Allocate magnitude
        IAllocationManager.MagnitudeAllocation[] memory allocations = _queueRandomAllocation_singleStrat_singleOpSet(defaultOperator, r, salt);

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
        IAllocationManager.MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
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
    function testFuzz_allocate_deallocate(uint256 r, uint256 salt) public {
        // Complete allocations & add a deallocation
        (MagnitudeAllocation[] memory allocations, MagnitudeAllocation[] memory deallocations) = 
            _queueRandomAllocationAndDeallocation(defaultOperator, r, salt, 1 /** numOpSets */);

        // Clear queue & check storage
        allocationManager.clearModificationQueue(defaultOperator, _strategyMockArray(), _maxNumToClear());
        assertEq(allocations[0].magnitudes[0], allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude should not be updated");

        // Validate storage - encumbered magnitude should just be allocations (we only have 1 allocation)
        IAllocationManager.MagnitudeInfo[] memory mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, allocations[0].operatorSets);
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
        assertEq(deallocations[0].magnitudes[0], allocationManager.encumberedMagnitude(defaultOperator, strategyMock), "encumberedMagnitude should be updated");
        mInfos = allocationManager.getAllocationInfo(defaultOperator, strategyMock, deallocations[0].operatorSets);
        assertEq(deallocations[0].magnitudes[0], mInfos[0].currentMagnitude, "currentMagnitude should be 0");
        assertEq(0, mInfos[0].pendingDiff, "pendingMagnitude should be 0");
        assertEq(0, mInfos[0].effectTimestamp, "effectTimestamp should be 0");
    }

    function _strategyMockArray() internal view returns (IStrategy[] memory) {
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        return strategies;
    }

    function _maxNumToClear() internal view returns (uint16[] memory) {
        uint16[] memory numToClear = new uint16[](1);
        numToClear[0] = type(uint16).max;
        return numToClear;
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

    function testFuzz_setDelay(uint32 delay) public {
        delay = uint32(bound(delay, 1, type(uint32).max));

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

    function test_fuzz_setDelay_multipleTimesWithinConfigurationDelay(uint32 firstDelay, uint32 secondDelay) public {
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

    function testFuzz_multipleDelays(uint32 firstDelay, uint32 secondDelay) public {
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

    function testFuzz_setDelay_DMCaller(uint32 delay) public {
        cheats.assume(delay > 0);

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
 */