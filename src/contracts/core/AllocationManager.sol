// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../permissions/Pausable.sol";
import "../libraries/EIP1271SignatureUtils.sol";
import "../libraries/SlashingLib.sol";
import "./AllocationManagerStorage.sol";

contract AllocationManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AllocationManagerStorage,
    ReentrancyGuardUpgradeable
{
    using Snapshots for Snapshots.History;
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;
    using SlashingLib for uint256;

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the immutable addresses of the strategy mananger, delegationManage,
     * and eigenpodManager contracts
     */
    constructor(
        IDelegationManager _delegation,
        IAVSDirectory _avsDirectory,
        uint32 _DEALLOCATION_DELAY,
        uint32 _ALLOCATION_CONFIGURATION_DELAY
    ) AllocationManagerStorage(_delegation, _avsDirectory, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY) {
        _disableInitializers();
    }

    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, and paused status.
     * minWithdrawalDelayBlocks is set only once here
     */
    function initialize(
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 initialPausedStatus
    ) external initializer {
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _transferOwnership(initialOwner);
    }

    /**
     * @notice Called by the delagation manager to set delay when operators register.
     * @param operator The operator to set the delay on behalf of.
     * @param delay The allocation delay in seconds.
     * @dev msg.sender is assumed to be the delegation manager.
     */
    function setAllocationDelay(address operator, uint32 delay) external {
        require(msg.sender == address(delegation), OnlyDelegationManager());
        _setAllocationDelay(operator, delay);
    }

    /**
     * @notice Called by operators to set their allocation delay.
     * @param delay the allocation delay in seconds
     * @dev msg.sender is assumed to be the operator
     */
    function setAllocationDelay(
        uint32 delay
    ) external {
        require(delegation.isOperator(msg.sender), OperatorNotRegistered());
        _setAllocationDelay(msg.sender, delay);
    }

    /**
     * @notice This function takes a list of strategies and adds all completable modifications for each strategy, 
     * updating the freeMagnitudes of the operator as needed.
     *
     * @param operator address to complete modifications for
     * @param strategies a list of strategies to complete modifications for
     * @param numToComplete a list of number of pending modifications to complete for each strategy
     *
     * @dev can be called permissionlessly by anyone
     */
    function completePendingModifications(
        address operator,
        IStrategy[] calldata strategies,
        uint16[] calldata numToComplete
    ) external onlyWhenNotPaused(PAUSED_MODIFY_ALLOCATIONS) {
        require(strategies.length == numToComplete.length, InputArrayLengthMismatch());
        require(delegation.isOperator(operator), OperatorNotRegistered());
        for (uint256 i = 0; i < strategies.length; ++i) {
            _completePendingModifications({operator: operator, strategy: strategies[i], numToComplete: numToComplete[i]});
        }
    }

    /**
     * @notice Modifies the propotions of slashable stake allocated to a list of operatorSets for a set of strategies
     * @param allocations array of magnitude adjustments for multiple strategies and corresponding operator sets
     * @dev Updates freeMagnitude for the updated strategies
     * @dev msg.sender is used as operator
     * @dev For each allocation, allocation.operatorSets MUST be ordered in ascending order according to the
     * encoding of the operatorSet. This is to prevent duplicate operatorSets being passed in. The easiest way to ensure
     * ordering is to sort allocated operatorSets by address first, and then sort for each avs by ascending operatorSetIds.
     */
    function modifyAllocations(
        MagnitudeAllocation[] calldata allocations
    ) external onlyWhenNotPaused(PAUSED_MODIFY_ALLOCATIONS) {
        (bool isSet, uint32 operatorAllocationDelay) = allocationDelay(msg.sender);
        require(isSet, UninitializedAllocationDelay());

        for (uint256 i = 0; i < allocations.length; ++i) {
            MagnitudeAllocation calldata allocation = allocations[i];
            require(allocation.operatorSets.length == allocation.magnitudes.length, InputArrayLengthMismatch());
            require(avsDirectory.isOperatorSetBatch(allocation.operatorSets), InvalidOperatorSet());

            // 1. For the given (operator,strategy) complete any pending deallocations to update free magnitude
            _completePendingModifications({
                operator: msg.sender,
                strategy: allocation.strategy,
                numToComplete: type(uint16).max
            });

            {
                // 2. Check current totalMagnitude matches expected value. This is to check for slashing race conditions
                // where an operator gets slashed from an operatorSet and as a result all the configured allocations have larger
                // proprtional magnitudes relative to each other.

                // Load the operator's total magnitude for the strategy.
                (bool exists,, uint224 currentTotalMagnitude) =
                    _totalMagnitudeUpdate[msg.sender][allocation.strategy].latestSnapshot();

                // If the operator has no total magnitude snapshot, set it to WAD, which denotes an unslashed operator.
                if (!exists) {
                    currentTotalMagnitude = WAD;
                    _totalMagnitudeUpdate[msg.sender][allocation.strategy].push({
                        key: uint32(block.timestamp),
                        value: currentTotalMagnitude
                    });
                    allocatableMagnitude[msg.sender][allocation.strategy] = uint64(currentTotalMagnitude);
                }

                require(
                    uint64(currentTotalMagnitude) == allocation.expectedTotalMagnitude, InvalidExpectedTotalMagnitude()
                );
            }

            for (uint256 j = 0; j < allocation.operatorSets.length; ++j) {
                bytes32 operatorSetKey = _encodeOperatorSet(allocation.operatorSets[j]);
                // Check that there are no pending allocations & deallocations for the operator, operatorSet, strategy
                (bool noPendingModification, MagnitudeInfo memory mInfo) = 
                    _completePendingModification(msg.sender, allocation.strategy, operatorSetKey);
                require(noPendingModification, ModificationAlreadyPending());

                // Calculate the new pending diff with this modification
                mInfo.pendingMagnitudeDiff =
                    int128(uint128(allocation.magnitudes[j])) - int128(uint128(mInfo.currentMagnitude));
                require(mInfo.pendingMagnitudeDiff != 0, SameMagnitude());

                // Handle deallocation/allocation and modification effect timestamp
                if (mInfo.pendingMagnitudeDiff < 0) {
                    // This is a deallocation

                    // 1. Update the effect timestamp for the deallocation
                    mInfo.effectTimestamp = uint32(block.timestamp) + DEALLOCATION_DELAY;
                } else if (mInfo.pendingMagnitudeDiff > 0) {
                    // This is an allocation

                    // 1. decrement free magnitude by incremented amount
                    uint64 magnitudeToAllocate = uint64(uint128(mInfo.pendingMagnitudeDiff));
                    require(allocatableMagnitude[msg.sender][allocation.strategy] >= magnitudeToAllocate, InsufficientAllocatableMagnitude());
                    allocatableMagnitude[msg.sender][allocation.strategy] -= magnitudeToAllocate;

                    // 2. Update the effectTimestamp for the allocation
                    mInfo.effectTimestamp = uint32(block.timestamp) + operatorAllocationDelay;
                }

                // Add the operatorSet to the modification queue
                modificationQueue[msg.sender][allocation.strategy].pushBack(operatorSetKey);
                // Allocate magnitude which will take effect at the `effectTimestamp`
                _operatorMagnitudeInfo[msg.sender][allocation.strategy][operatorSetKey] = mInfo;
            }
        }
    }

    /**
     * @notice Called by an AVS to slash an operator for given operatorSetId, list of strategies, and bipsToSlash.
     * For each given (operator, operatorSetId, strategy) tuple, bipsToSlash
     * bips of the operatorSet's slashable stake allocation will be slashed
     *
     * @param operator the address to slash
     * @param operatorSetId the ID of the operatorSet the operator is being slashed on behalf of
     * @param strategies the set of strategies to slash
     * @param wadToSlash the parts in 1e18 to slash, this will be proportional to the
     * operator's slashable stake allocation for the operatorSet
     */
    function slashOperator(
        address operator,
        uint32 operatorSetId,
        IStrategy[] calldata strategies,
        uint256 wadToSlash
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_SLASHING) {
        require(wadToSlash > 0, InvalidWadToSlash());
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        require(avsDirectory.isOperatorSlashable(operator, operatorSet), InvalidOperator());
        bytes32 operatorSetKey = _encodeOperatorSet(operatorSet);

        for (uint256 i = 0; i < strategies.length; ++i) {
            // 1. Complete pending modifications for the operator, strategy, and operatorSet
            (,MagnitudeInfo memory mInfo) = _completePendingModification(operator, strategies[i], operatorSetKey);

            // 2. Calculate the slashed magnitude and update operatorMagnitudeInfo
            uint64 slashedMagnitude = uint64(uint256(mInfo.currentMagnitude).mulWad(wadToSlash));
            mInfo.currentMagnitude -= slashedMagnitude;
            // if there is a pending deallocation, slash pending deallocation proportionally
            if (mInfo.pendingMagnitudeDiff < 0) {
                int128 slashedPending = int128(uint128(uint256(uint128(-mInfo.pendingMagnitudeDiff)).mulWad(wadToSlash)));
                mInfo.pendingMagnitudeDiff += int128(slashedPending);
            }
            // update operatorMagnitudeInfo
            _operatorMagnitudeInfo[operator][strategies[i]][operatorSetKey] = mInfo;

            // 3. update totalMagnitude, get total magnitude and subtract slashedMagnitude
            Snapshots.History storage totalMagnitudes = _totalMagnitudeUpdate[operator][strategies[i]];
            uint64 totalMagnitudeBeforeSlashing = uint64(totalMagnitudes.latest());
            totalMagnitudes.push({key: uint32(block.timestamp), value: totalMagnitudeBeforeSlashing - slashedMagnitude});

            // 4. Decrease operators shares in the DM
            delegation.decreaseOperatorShares(operator, strategies[i], totalMagnitudeBeforeSlashing, uint64(totalMagnitudes.latest()));
        }
    }

    /**
     * @notice Helper for setting an operators allocation delay.
     * @param operator The operator to set the delay on behalf of.
     * @param delay The allocation delay in seconds.
     */
    function _setAllocationDelay(address operator, uint32 delay) internal {
        require(delay != 0, InvalidDelay());

        AllocationDelayInfo memory info = _allocationDelayInfo[operator];

        if (info.pendingDelay != 0 && block.timestamp >= info.pendingDelayEffectTimestamp) {
            info.delay = info.pendingDelay;
        }

        info.pendingDelay = delay;
        info.pendingDelayEffectTimestamp = uint32(block.timestamp + ALLOCATION_CONFIGURATION_DELAY);

        _allocationDelayInfo[operator] = info;
        emit AllocationDelaySet(operator, delay);
    }

    /**
     * @notice For a single strategy, complete pending modifications for an operator, strategy
     * @param operator address to update allocatableMagnitude for
     * @param strategy the strategy to update allocatableMagnitude for
     * @param numToComplete the number of pending free magnitudes deallocations to complete
     * @dev read through pending free magnitudes and add to allocatableMagnitude if completableTimestamp is >= block timestamp
     * In addition to updating allocatableMagnitude, updates next starting index to read from for pending free magnitudes after completing
     */
    function _completePendingModifications(address operator, IStrategy strategy, uint16 numToComplete) internal {
        uint256 numCompleted;
        while (modificationQueue[operator][strategy].length() > 0 && numCompleted < numToComplete) {
            bytes32 opsetKey = modificationQueue[operator][strategy].front();
            (bool completed,) = _completePendingModification(operator, strategy, opsetKey);
            if (completed) {
                // if the modification was completed, remove it from the queue
                modificationQueue[operator][strategy].popFront();
                ++numCompleted;
            } else {
                // else, the modification is still pending, and so are all subsequent modifications, so break
                break;
            }
        }
    }
    
    // returns whether the modification was completed
    function _completePendingModification(address operator, IStrategy strategy, bytes32 operatorSetKey) internal returns (bool, MagnitudeInfo memory) {
        MagnitudeInfo memory mInfo = _operatorMagnitudeInfo[operator][strategy][operatorSetKey];
        // if the effect timestamp is not reached, return false
        if (block.timestamp < mInfo.effectTimestamp) {
            return (false, mInfo);
        } else if (mInfo.effectTimestamp == 0) {
            // if the effect timestamp is 0, there is no pending modification, so we can avoid the below logic
            return (true, mInfo);
        }

        // update stored magnitudes
        if (mInfo.pendingMagnitudeDiff > 0) {
            mInfo.currentMagnitude += uint64(uint128(mInfo.pendingMagnitudeDiff));
        } else {
            mInfo.currentMagnitude -= uint64(uint128(-mInfo.pendingMagnitudeDiff));
            // if this is a deallocation, add the magnitude back to freeMagnitude
            allocatableMagnitude[operator][strategy] += uint64(uint128(-mInfo.pendingMagnitudeDiff));
        }
        mInfo.pendingMagnitudeDiff = 0;

        // update the magnitude info in storage
        _operatorMagnitudeInfo[operator][strategy][operatorSetKey] = mInfo;

        return (true, mInfo);
    }

    /**
     * @notice Returns the allocation delay of an operator
     * @param operator The operator to get the allocation delay for
     */
    function allocationDelay(
        address operator
    ) public view returns (bool isSet, uint32 delay) {
        AllocationDelayInfo memory info = _allocationDelayInfo[operator];
        if (info.pendingDelay != 0 && block.timestamp >= info.pendingDelayEffectTimestamp) {
            return (true, info.pendingDelay);
        } else {
            return (info.delay != 0, info.delay);
        }
    }

    /**
     * @param operator the operator to get the slashable magnitude for
     * @param strategies the strategies to get the slashable magnitude for
     *
     * @return operatorSets the operator sets the operator is a member of and the current slashable magnitudes for each strategy
     */
    function getSlashableMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (OperatorSet[] memory, uint64[][] memory) {
        OperatorSet[] memory operatorSets = avsDirectory.getOperatorSetsOfOperator(operator, 0, type(uint256).max);
        uint64[][] memory slashableMagnitudes = new uint64[][](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            slashableMagnitudes[i] = new uint64[](operatorSets.length);
            for (uint256 j = 0; j < operatorSets.length; ++j) {
                bytes32 operatorSetKey = _encodeOperatorSet(operatorSets[j]);
                MagnitudeInfo memory mInfo = _operatorMagnitudeInfo[operator][strategies[i]][operatorSetKey];
                slashableMagnitudes[i][j] = mInfo.currentMagnitude;
                if (block.timestamp >= mInfo.effectTimestamp && mInfo.effectTimestamp != 0) {
                    if (mInfo.pendingMagnitudeDiff < 0) {
                        slashableMagnitudes[i][j] -= uint64(uint128(-mInfo.pendingMagnitudeDiff));
                    } else if (mInfo.pendingMagnitudeDiff > 0) {
                        slashableMagnitudes[i][j] += uint64(uint128(mInfo.pendingMagnitudeDiff));
                    }
                }
            }
        }
        return (operatorSets, slashableMagnitudes);
    }

    /**
     * @notice Get the allocatable magnitude for an operator and strategy based on number of pending deallocations
     * that could be completed at the same time. This is the sum of freeMagnitude and the sum of all pending completable deallocations.
     * @param operator the operator to get the allocatable magnitude for
     * @param strategy the strategy to get the allocatable magnitude for
     */
    function getAllocatableMagnitude(address operator, IStrategy strategy) external view returns (uint64) {
        uint64 freeableMagnitude = 0;
        for (uint256 i = 0; i < modificationQueue[operator][strategy].length(); ++i) {
            bytes32 opsetKey = modificationQueue[operator][strategy].at(i);
            MagnitudeInfo memory mInfo = _operatorMagnitudeInfo[operator][strategy][opsetKey];
            if (block.timestamp >= mInfo.effectTimestamp && mInfo.effectTimestamp != 0 && mInfo.pendingMagnitudeDiff < 0) {
                freeableMagnitude += uint64(uint128(-mInfo.pendingMagnitudeDiff));
            } else {
                break;
            }
        }
        return allocatableMagnitude[operator][strategy] + freeableMagnitude;
    }

    /**
     * @notice Returns the current total magnitudes of an operator for a given set of strategies
     * @param operator the operator to get the total magnitude for
     * @param strategies the strategies to get the total magnitudes for
     * @return totalMagnitudes the total magnitudes for each strategy
     */
    function getTotalMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (uint64[] memory) {
        uint64[] memory totalMagnitudes = new uint64[](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            (bool exists,, uint224 value) = _totalMagnitudeUpdate[operator][strategies[i]].latestSnapshot();
            if (!exists) {
                totalMagnitudes[i] = WAD;
            } else {
                totalMagnitudes[i] = uint64(value);
            }
        }
        return totalMagnitudes;
    }

    /**
     * @notice Returns the total magnitudes of an operator for a given set of strategies at a given timestamp
     * @param operator the operator to get the total magnitude for
     * @param strategies the strategies to get the total magnitudes for
     * @param timestamp the timestamp to get the total magnitudes at
     * @return totalMagnitudes the total magnitudes for each strategy
     */
    function getTotalMagnitudesAtTimestamp(
        address operator,
        IStrategy[] calldata strategies,
        uint32 timestamp
    ) external view returns (uint64[] memory) {
        uint64[] memory totalMagnitudes = new uint64[](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            (uint224 value, uint256 pos) = _totalMagnitudeUpdate[operator][strategies[i]].upperLookupWithPos(timestamp);
            // if there is no existing total magnitude snapshot
            if (value == 0 && pos == 0) {
                totalMagnitudes[i] = WAD;
            } else {
                totalMagnitudes[i] = uint64(value);
            }
        }
        return totalMagnitudes;
    }

    /**
     * @notice Returns the pending modifications of an operator for a given strategy and operatorSets.
     * @param operator the operator to get the pending modification for
     * @param strategy the strategy to get the pending modification for
     * @param operatorSets the operatorSets to get the pending modification for
     * @return timestamps the timestamps for each pending dealloction
     * @return pendingMagnitudeDiffs the pending modification diffs for each operatorSet
     */
    function getPendingModifications(
        address operator,
        IStrategy strategy,
        OperatorSet[] calldata operatorSets
    ) external view returns (uint32[] memory timestamps, int128[] memory pendingMagnitudeDiffs) {
        timestamps = new uint32[](operatorSets.length);
        pendingMagnitudeDiffs = new int128[](operatorSets.length);

        for (uint256 i = 0; i < operatorSets.length; ++i) {
            MagnitudeInfo memory opsetMagnitudeInfo =
                _operatorMagnitudeInfo[operator][strategy][_encodeOperatorSet(operatorSets[i])];

            if (opsetMagnitudeInfo.effectTimestamp < block.timestamp && opsetMagnitudeInfo.effectTimestamp != 0) {
                pendingMagnitudeDiffs[i] = opsetMagnitudeInfo.pendingMagnitudeDiff;
                timestamps[i] = opsetMagnitudeInfo.effectTimestamp;
            }
        }
    }

    /// @dev Returns an `OperatorSet` encoded into a 32-byte value.
    /// @param operatorSet The `OperatorSet` to encode.
    function _encodeOperatorSet(
        OperatorSet memory operatorSet
    ) internal pure returns (bytes32) {
        return bytes32(abi.encodePacked(operatorSet.avs, uint96(operatorSet.operatorSetId)));
    }

    /// @dev Returns an `OperatorSet` decoded from an encoded 32-byte value.
    /// @param encoded The encoded `OperatorSet` to decode.
    /// @dev Assumes `encoded` is encoded via `_encodeOperatorSet(operatorSet)`.
    function _decodeOperatorSet(
        bytes32 encoded
    ) internal pure returns (OperatorSet memory) {
        return OperatorSet({
            avs: address(uint160(uint256(encoded) >> 96)),
            operatorSetId: uint32(uint256(encoded) & type(uint96).max)
        });
    }
}
