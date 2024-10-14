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

    function _randomMagnitudeAllocation(
        uint256 r,
        uint256 salt
    ) internal view returns (IAllocationManagerTypes.MagnitudeAllocation memory) {
        r = uint256(keccak256(abi.encodePacked(r, salt)));

        // Mock a random operator set.
        OperatorSet[] memory operatorSets = new OperatorSet[](1);
        operatorSets[0] = OperatorSet({avs: _randomAddr(r, 0), operatorSetId: uint32(r)});

        uint64[] memory magnitudes = new uint64[](1);
        magnitudes[0] = uint64(bound(r, 1, 1e18));

        // Mock a random magnitude allocation.
        return IAllocationManagerTypes.MagnitudeAllocation({
            strategy: strategyMock,
            expectedMaxMagnitude: 1e18, // magnitude starts at 100%
            operatorSets: operatorSets,
            magnitudes: magnitudes
        });
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

    /// @dev Asserts the following:
    /// 1. Fn can only be called by operator with actively set allocation delay.
    /// 2. Fn can only be called with valid operator sets.
    /// 3. Fn can only be called when unpaused.
    function testFuzz_modifyAllocations_Checks(
        uint256 r
    ) public {
        // Mock random operator and allocation delay..
        address operator = _randomAddr(r, 0);
        uint32 delay = uint32(bound(r, 1, type(uint32).max));

        // Mock random magnitude allocation.
        IAllocationManagerTypes.MagnitudeAllocation[] memory a = new IAllocationManagerTypes.MagnitudeAllocation[](1);
        a[0] = _randomMagnitudeAllocation(r, 0);

        // Assert that allocation cannot be made unless allocation delay has been set (aka if the operator has registered).
        cheats.prank(operator);
        cheats.expectRevert(IAllocationManagerErrors.UninitializedAllocationDelay.selector);
        allocationManager.modifyAllocations(a);

        // Mock allocation delay being set and active.
        cheats.prank(address(delegationManagerMock));
        allocationManager.setAllocationDelay(operator, delay);
        cheats.warp(block.timestamp + ALLOCATION_CONFIGURATION_DELAY );

        // Assert that fn can only be called with valid operator sets.
        cheats.prank(operator);
        cheats.expectRevert(IAllocationManagerErrors.InvalidOperatorSet.selector);
        allocationManager.modifyAllocations(a);

        // Mock valid operator sets.
        avsDirectoryMock.setIsOperatorSetBatch(a[0].operatorSets, true);

        // Pause allocation modifications.
        allocationManager.pause(2 ** PAUSED_MODIFY_ALLOCATIONS);

        // Assert that fn can only be called when unpaused.
        cheats.prank(operator);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        allocationManager.modifyAllocations(a);

        // Unpause allocation modifications.
        cheats.prank(unpauser);
        allocationManager.unpause(2 ** PAUSED_MODIFY_ALLOCATIONS);
    }
}

contract AllocationManagerUnitTests_ClearModificationQueue is AllocationManagerUnitTests {
    /// -----------------------------------------------------------------------
    /// clearModificationQueue()
    /// -----------------------------------------------------------------------

    /// @dev Asserts the following:
    /// 1. Fn array input lengths must match.
    /// 2. Fn can only be called by registered operator.
    /// 3. Fn can only be called when unpaused.
    function testFuzz_clearModificationQueue_Checks(uint256 r) public {
        // Mock random operator address, stratgies array, and numToClear array.
        address operator = _randomAddr(r, 0);
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        uint16[] memory numToClear = new uint16[](2);

        // Assert that fn inputs arrays match.
        cheats.expectRevert(IAllocationManagerErrors.InputArrayLengthMismatch.selector);
        allocationManager.clearModificationQueue(operator, strategies, numToClear);

        // Adjust numToClear array to match strategies array length.
        numToClear = new uint16[](1);

        // Assert that operator cannot clear the modification queue unless registered.
        cheats.expectRevert(IAllocationManagerErrors.OperatorNotRegistered.selector);
        allocationManager.clearModificationQueue(operator, strategies, numToClear);

        // Mock operator being registered.
        delegationManagerMock.setIsOperator(operator, true);

        // Pause allocation modifications.
        allocationManager.pause(2 ** PAUSED_MODIFY_ALLOCATIONS);

        // Assert that fn can only be called when unpaused.
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        allocationManager.clearModificationQueue(operator, strategies, numToClear);

        // Unpause allocation modifications.
        cheats.prank(unpauser);
        allocationManager.unpause(2 ** PAUSED_MODIFY_ALLOCATIONS);

        // TODO
    }
}

contract AllocationManagerUnitTests_SetAllocationDelay is AllocationManagerUnitTests {

    /// -----------------------------------------------------------------------
    /// setAllocationDelay() + getAllocationDelay()
    /// -----------------------------------------------------------------------

    /// @dev Asserts the following:
    ///     Calling as the operator:
    ///         1. Fn can only be called by registed operators..
    ///         2. Fn emits `AllocationDelaySet` event.
    ///         3. Fn properly mutates storage such that getAllocationDelay returns the correct value.
    ///     Calling as the DM:
    ///         2. Only the DM can set allocation delay.
    ///         3. Fn emits `AllocationDelaySet` event.
    ///         4. Fn properly mutates storage such that getAllocationDelay returns the correct value.
    function testFuzz_allocationDelay_Checks(uint256 r) public {
        // Mock random operator address and delay.
        address operator = _randomAddr(r, 0);
        uint32 expectedDelay = uint32(bound(r, 1, type(uint32).max));
        uint32 effectTimestamp = uint32(block.timestamp + ALLOCATION_CONFIGURATION_DELAY);

        // If r is even, pretend to be the operator, else pretend to be the DM.
        if (r % 2 == 0) {
            // Assert unregister operator cannot set allocation delay.
            cheats.prank(operator);
            cheats.expectRevert(IAllocationManagerErrors.OperatorNotRegistered.selector);
            allocationManager.setAllocationDelay(expectedDelay);

            // Mock operator being registered.
            delegationManagerMock.setIsOperator(operator, true);

            // Assert `AllocationDelaySet` event is emitted.
            cheats.expectEmit(true, false, false, false);
            emit AllocationDelaySet(operator, expectedDelay, effectTimestamp);

            // Set the allocation delay as the operator.
            cheats.prank(operator);
            allocationManager.setAllocationDelay(expectedDelay);            
        } else {
            // Assert only the DM can set allocation delay.
            cheats.expectRevert(IAllocationManagerErrors.OnlyDelegationManager.selector);
            allocationManager.setAllocationDelay(operator, expectedDelay);

            // Assert `AllocationDelaySet` event is emitted.
            cheats.expectEmit(true, false, false, false);
            emit AllocationDelaySet(operator, expectedDelay, effectTimestamp);

            // Set the allocation delay as the DM.
            cheats.prank(address(delegationManagerMock));
            allocationManager.setAllocationDelay(operator, expectedDelay);
        }

        // Assert that the delay does not change until the allocation config delay has elapsed.
        (bool isSet, uint32 delay) = allocationManager.getAllocationDelay(operator);
        assertEq(delay, 0);
        assertTrue(!isSet);

        // Assert that the delay does not change until the allocation config delay has elapsed.
        cheats.warp(effectTimestamp);
        (isSet, delay) = allocationManager.getAllocationDelay(operator);
        assertEq(delay, expectedDelay);
        assertTrue(isSet);
    }
}