// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../contracts/core/AllocationManager.sol";
import "forge-std/Test.sol";
import "../TestConstants.sol";

contract AllocationManagerHarness is AllocationManager {
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;

    constructor(
        IAllocationManagerView _allocationManagerView,
        IDelegationManager _delegation,
        IStrategy _eigenStrategy,
        IPauserRegistry _pauserRegistry,
        IPermissionController _permissionController,
        uint32 _DEALLOCATION_DELAY,
        uint32 _ALLOCATION_CONFIGURATION_DELAY
    )
        AllocationManager(
            _allocationManagerView,
            _delegation,
            _eigenStrategy,
            _pauserRegistry,
            _permissionController,
            _DEALLOCATION_DELAY,
            _ALLOCATION_CONFIGURATION_DELAY
        )
    {}

    function deallocationQueueAtIndex(address operator, IStrategy strategy, uint index) external view returns (bytes32) {
        return deallocationQueue[operator][strategy].at(index);
    }

    function setSlasherToZero(OperatorSet memory operatorSet) external {
        _slashers[operatorSet.key()] = SlasherParams(address(0), address(0), 0);
    }

    /// @notice Returns the raw SlasherParams struct from storage for testing purposes.
    /// @dev This bypasses the in-memory application of pending slasher that getSlasher() does.
    function getSlasherParams(OperatorSet memory operatorSet) external view returns (SlasherParams memory) {
        return _slashers[operatorSet.key()];
    }

    /// @notice Returns the raw AllocationDelayInfo struct from storage for testing purposes.
    /// @dev This bypasses the in-memory application of pending delay that getAllocationDelay() does.
    function getAllocationDelayInfoRaw(address operator) external view returns (AllocationDelayInfo memory) {
        return _allocationDelayInfo[operator];
    }
}
