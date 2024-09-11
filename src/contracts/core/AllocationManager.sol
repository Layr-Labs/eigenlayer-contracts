// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../permissions/Pausable.sol";
import "../libraries/EIP1271SignatureUtils.sol";
import "../libraries/SlashingConstants.sol";
import "./AllocationManagerStorage.sol";

contract AllocationManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AllocationManagerStorage,
    ReentrancyGuardUpgradeable
{
    using Snapshots for Snapshots.History;

    /// @dev BIPS factor for slashable bips
    uint256 internal constant BIPS_FACTOR = 10_000;

    /// @dev Maximum number of pending updates that can be queued for allocations/deallocations
    uint256 public constant MAX_PENDING_UPDATES = 1;

    /// @dev Returns the chain ID from the time the contract was deployed.
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the immutable addresses of the strategy mananger, delegationManager, slasher,
     * and eigenpodManager contracts
     */
    constructor(
        IDelegationManager _delegation,
        IAVSDirectory _avsDirectory
    ) AllocationManagerStorage(_delegation, _avsDirectory) {
        _disableInitializers();
        ORIGINAL_CHAIN_ID = block.chainid;
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
        _DOMAIN_SEPARATOR = _calculateDomainSeparator();
        _transferOwnership(initialOwner);
    }

    /**
     * @notice For all pending deallocations that have become completable, their pending free magnitude can be
     * added back to the free magnitude of the (operator, strategy) amount. This function takes a list of strategies
     * and adds all completable deallocations for each strategy, updating the freeMagnitudes of the operator
     *
     * @param operator address to complete deallocations for
     * @param strategies a list of strategies to complete deallocations for
     * @param numToComplete a list of number of pending free magnitude deallocations to complete for each strategy
     *
     * @dev can be called permissionlessly by anyone
     */
    function updateFreeMagnitude(
        address operator,
        IStrategy[] calldata strategies,
        uint16[] calldata numToComplete
    ) external {
        require(strategies.length == numToComplete.length, InputArrayLengthMismatch());
        require(delegation.isOperator(operator), OperatorNotRegistered());
        for (uint256 i = 0; i < strategies.length; ++i) {
            _updateFreeMagnitude({operator: operator, strategy: strategies[i], numToComplete: numToComplete[i]});
        }
    }

    /**
     * @notice Modifies the propotions of slashable stake allocated to a list of operatorSets for a set of strategies
     * @param operator address to modify allocations for
     * @param allocations array of magnitude adjustments for multiple strategies and corresponding operator sets
     * @param operatorSignature signature of the operator if msg.sender is not the operator
     * @dev Updates freeMagnitude for the updated strategies
     * @dev Must be called by the operator or with a valid operator signature
     * @dev For each allocation, allocation.operatorSets MUST be ordered in ascending order according to the
     * encoding of the operatorSet. This is to prevent duplicate operatorSets being passed in. The easiest way to ensure
     * ordering is to sort allocated operatorSets by address first, and then sort for each avs by ascending operatorSetIds.
     */
    function modifyAllocations(
        address operator,
        MagnitudeAllocation[] calldata allocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external {
        if (msg.sender != operator) {
            _verifyOperatorSignature(operator, allocations, operatorSignature);
        }
        require(delegation.isOperator(operator), OperatorNotRegistered());
        IDelegationManager.AllocationDelayDetails memory details = delegation.operatorAllocationDelay(operator);
        require(details.isSet, UninitializedAllocationDelay());
        // effect timestamp for allocations to take effect. This is configurable by operators
        uint32 effectTimestamp = uint32(block.timestamp) + details.allocationDelay;
        // completable timestamp for deallocations
        uint32 completableTimestamp = uint32(block.timestamp) + SlashingConstants.DEALLOCATION_DELAY;

        for (uint256 i = 0; i < allocations.length; ++i) {
            // 1. For the given (operator,strategy) clear all pending free magnitude for the strategy and update freeMagnitude
            _updateFreeMagnitude({
                operator: operator,
                strategy: allocations[i].strategy,
                numToComplete: type(uint16).max
            });

            // 2. Check current totalMagnitude matches expected value. This is to check for slashing race conditions
            // where an operator gets slashed from an operatorSet and as a result all the configured allocations have larger
            // proprtional magnitudes relative to eachother. This check prevents any surprising behavior.
            uint64 currentTotalMagnitude = _getLatestTotalMagnitude(operator, allocations[i].strategy);
            require(currentTotalMagnitude == allocations[i].expectedTotalMagnitude, InvalidExpectedTotalMagnitude());

            // 3. set allocations for the strategy after updating freeMagnitude
            _modifyAllocations({
                operator: operator,
                allocation: allocations[i],
                allocationEffectTimestamp: effectTimestamp,
                deallocationCompletableTimestamp: completableTimestamp
            });
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
     * @param bipsToSlash the number of bips to slash, this will be proportional to the
     * operator's slashable stake allocation for the operatorSet
     */
    function slashOperator(
        address operator,
        uint32 operatorSetId,
        IStrategy[] calldata strategies,
        uint16 bipsToSlash
    ) external {
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        bytes32 operatorSetKey = _encodeOperatorSet(operatorSet);
        require(isOperatorSlashable(operator, operatorSet), InvalidOperator());

        for (uint256 i = 0; i < strategies.length; ++i) {
            // 1. calculate slashed magnitude from current allocation
            // update current and all following queued magnitude updates for (operator, strategy, operatorSetId) tuple
            uint64 slashedMagnitude;
            {
                uint64 currentMagnitude = uint64(
                    _magnitudeUpdate[operator][strategies[i]][operatorSetKey].upperLookupRecent(uint32(block.timestamp))
                );
                // slash nonzero current magnitude
                if (currentMagnitude != 0) {
                    /// TODO: add wrapping library for rounding up for slashing accounting
                    slashedMagnitude = uint64(uint256(bipsToSlash) * uint256(currentMagnitude) / BIPS_FACTOR);

                    _magnitudeUpdate[operator][strategies[i]][operatorSetKey].decrementAtAndFutureSnapshots({
                        key: uint32(block.timestamp),
                        decrementValue: slashedMagnitude
                    });
                }
            }

            // 2. if there are any pending deallocations then need to update and decrement if they fall within slashable window
            // loop backwards through _queuedDeallocationIndices, each element contains an array index to
            // corresponding deallocation to access in pendingFreeMagnitude
            // if completable, then break
            //      (since ordered by completableTimestamps, older deallocations will also be completable and outside slashable window)
            // if NOT completable, then add to slashed magnitude
            {
                uint256 queuedDeallocationIndicesLen =
                    _queuedDeallocationIndices[operator][strategies[i]][operatorSetKey].length;
                for (uint256 j = queuedDeallocationIndicesLen; j > 0; --j) {
                    // index of pendingFreeMagnitude/deallocation to check for slashing
                    uint256 index = _queuedDeallocationIndices[operator][strategies[i]][operatorSetKey][j - 1];
                    PendingFreeMagnitude storage pendingFreeMagnitude =
                        _pendingFreeMagnitude[operator][strategies[i]][index];

                    // Reached pendingFreeMagnitude/deallocation that is completable and not within slashability window,
                    // therefore older deallocations will also be completable. Since this is ordered by completableTimestamps break loop now
                    if (pendingFreeMagnitude.completableTimestamp >= uint32(block.timestamp)) {
                        break;
                    }

                    // pending deallocation is still within slashable window, slash magnitudeDiff and add to slashedMagnitude
                    uint64 slashedAmount =
                        uint64(uint256(bipsToSlash) * uint256(pendingFreeMagnitude.magnitudeDiff) / BIPS_FACTOR);
                    pendingFreeMagnitude.magnitudeDiff -= slashedAmount;
                    slashedMagnitude += slashedAmount;
                }
            }

            // 3. update totalMagnitude, get total magnitude and subtract slashedMagnitude
            _totalMagnitudeUpdate[operator][strategies[i]].push({
                key: uint32(block.timestamp),
                value: _getLatestTotalMagnitude(operator, strategies[i]) - slashedMagnitude
            });
        }
    }

    /**
     * @notice Called by an operator to cancel a salt that has been used to register with an AVS.
     *
     * @param salt A unique and single use value associated with the approver signature.
     */
    function cancelSalt(bytes32 salt) external override {
        // Mutate `operatorSaltIsSpent` to `true` to prevent future spending.
        operatorSaltIsSpent[msg.sender][salt] = true;
    }

    /**
     * @notice For a single strategy, update freeMagnitude by adding completable pending free magnitudes
     * @param operator address to update freeMagnitude for
     * @param strategy the strategy to update freeMagnitude for
     * @param numToComplete the number of pending free magnitudes deallocations to complete
     * @dev read through pending free magnitudes and add to freeMagnitude if completableTimestamp is >= block timestamp
     * In addition to updating freeMagnitude, updates next starting index to read from for pending free magnitudes after completing
     */
    function _updateFreeMagnitude(address operator, IStrategy strategy, uint16 numToComplete) internal {
        OperatorMagnitudeInfo storage info = operatorMagnitudeInfo[operator][strategy];
        (uint64 freeMagnitudeToAdd, uint192 completed) =
            _getPendingFreeMagnitude(operator, strategy, numToComplete, info.nextPendingFreeMagnitudeIndex);
        info.freeMagnitude += freeMagnitudeToAdd;
        info.nextPendingFreeMagnitudeIndex += completed;
    }

    /**
     * @notice For a single strategy, modify magnitude allocations for each of the specified operatorSets
     * @param operator address to modify allocations for
     * @param allocation the magnitude allocations to modify for a single strategy
     * @param allocationEffectTimestamp the timestamp when the allocations will take effect
     * @param deallocationCompletableTimestamp the timestamp when the deallocations will be completable
     * @dev For each allocation, allocation.operatorSets MUST be ordered in ascending order according to the
     * encoding of the operatorSet. This is to prevent duplicate operatorSets being passed in. The easiest way to ensure
     * ordering is to sort allocated operatorSets by address first, and then sort for each avs by ascending operatorSetIds.
     */
    function _modifyAllocations(
        address operator,
        MagnitudeAllocation calldata allocation,
        uint32 allocationEffectTimestamp,
        uint32 deallocationCompletableTimestamp
    ) internal {
        require(allocation.operatorSets.length == allocation.magnitudes.length, InputArrayLengthMismatch());

        // OperatorSet[] calldata opSets = allocation.operatorSets;

        bytes32 prevOperatorSet = bytes32(0);

        for (uint256 i = 0; i < allocation.operatorSets.length; ++i) {
            require(
                avsDirectory.isOperatorSet(allocation.operatorSets[i].avs, allocation.operatorSets[i].operatorSetId),
                InvalidOperatorSet()
            );
            // use encoding of operatorSet to ensure ordering and also used to use OperatorSet struct as key in mappings
            bytes32 operatorSetKey = _encodeOperatorSet(allocation.operatorSets[i]);
            require(prevOperatorSet < operatorSetKey, OperatorSetsNotInAscendingOrder());
            prevOperatorSet = operatorSetKey;

            // Read current magnitude allocation including its respective array index and length.
            // We'll use these values later to check the number of pending allocations/deallocations.
            (uint224 currentMagnitude, uint256 pos, uint256 length) = _magnitudeUpdate[operator][allocation.strategy][operatorSetKey]
                .upperLookupRecentWithPos(uint32(block.timestamp));

            // Check that there is at MOST `MAX_PENDING_UPDATES` combined allocations & deallocations for the operator, operatorSet, strategy
            {
                uint256 numPendingAllocations;
                // if no lookup found (currentMagnitude == 0 && pos == 0), then we are at the beginning of the array
                // the number of pending allocations is simply length
                if (currentMagnitude == 0 && pos == 0) {
                    numPendingAllocations = length;
                    // if lookup found, then we take the difference between length-1 and pos
                } else {
                    numPendingAllocations = length - pos - 1;
                }
                uint256 numPendingDeallocations =
                    _getNumQueuedDeallocations(operator, allocation.strategy, operatorSetKey);

                require(
                    numPendingAllocations + numPendingDeallocations < MAX_PENDING_UPDATES,
                    PendingAllocationOrDeallocation()
                );
            }

            if (allocation.magnitudes[i] < uint64(currentMagnitude)) {
                // Newly configured magnitude is less than current value.
                // Therefore we handle this as a deallocation

                // Note: MAX_PENDING_UPDATES == 1, so we do not have to decrement any allocations

                // 1. push PendingFreeMagnitude and respective array index into (op,opSet,Strategy) queued deallocations
                uint256 index = _pendingFreeMagnitude[operator][allocation.strategy].length;
                _pendingFreeMagnitude[operator][allocation.strategy].push(
                    PendingFreeMagnitude({
                        magnitudeDiff: uint64(currentMagnitude) - allocation.magnitudes[i],
                        completableTimestamp: deallocationCompletableTimestamp
                    })
                );
                _queuedDeallocationIndices[operator][allocation.strategy][operatorSetKey].push(index);
            } else if (allocation.magnitudes[i] > uint64(currentMagnitude)) {
                // Newly configured magnitude is greater than current value.
                // Therefore we handle this as an allocation

                // 1. allocate magnitude which will take effect in the future 21 days from now
                _magnitudeUpdate[operator][allocation.strategy][operatorSetKey].push({
                    key: allocationEffectTimestamp,
                    value: allocation.magnitudes[i]
                });

                // 2. decrement free magnitude by incremented amount
                OperatorMagnitudeInfo storage info = operatorMagnitudeInfo[operator][allocation.strategy];
                require(
                    info.freeMagnitude >= allocation.magnitudes[i] - uint64(currentMagnitude),
                    InsufficientAllocatableMagnitude()
                );
                info.freeMagnitude -= allocation.magnitudes[i] - uint64(currentMagnitude);
            }
        }
    }

    /**
     * @notice Get the number of queued dealloations for the given (operator, strategy, operatorSetKey) tuple
     * @param operator address to get queued deallocations for
     * @param strategy the strategy to get queued deallocations for
     * @param operatorSetKey the encoded operatorSet to get queued deallocations for
     */
    function _getNumQueuedDeallocations(
        address operator,
        IStrategy strategy,
        bytes32 operatorSetKey
    ) internal view returns (uint256) {
        uint256 numQueuedDeallocations;

        uint256 length = _queuedDeallocationIndices[operator][strategy][operatorSetKey].length;

        for (uint256 i = length; i > 0; --i) {
            // index of pendingFreeMagnitude/deallocation to check for slashing
            uint256 index = _queuedDeallocationIndices[operator][strategy][operatorSetKey][i - 1];
            PendingFreeMagnitude memory pendingFreeMagnitude = _pendingFreeMagnitude[operator][strategy][index];

            // If completableTimestamp is greater than completeUntilTimestamp, break
            if (pendingFreeMagnitude.completableTimestamp < uint32(block.timestamp)) {
                ++numQueuedDeallocations;
            } else {
                break;
            }
        }

        return numQueuedDeallocations;
    }

    /// @dev gets the latest total magnitude or overwrites it if it is not set
    function _getLatestTotalMagnitude(address operator, IStrategy strategy) internal returns (uint64) {
        (bool exists,, uint224 totalMagnitude) = _totalMagnitudeUpdate[operator][strategy].latestSnapshot();
        if (!exists) {
            totalMagnitude = SlashingConstants.INITIAL_TOTAL_MAGNITUDE;
            _totalMagnitudeUpdate[operator][strategy].push({key: uint32(block.timestamp), value: totalMagnitude});
            operatorMagnitudeInfo[operator][strategy].freeMagnitude = SlashingConstants.INITIAL_TOTAL_MAGNITUDE;
        }

        return uint64(totalMagnitude);
    }

    /// @dev gets the pending free magnitude available to add by completing numToComplete pending deallocations
    /// and returns the next index to start from if completed.
    function _getPendingFreeMagnitude(
        address operator,
        IStrategy strategy,
        uint16 numToComplete,
        uint192 nextIndex
    ) internal view returns (uint64 freeMagnitudeToAdd, uint192 completed) {
        uint256 pendingFreeMagnitudeLength = _pendingFreeMagnitude[operator][strategy].length;
        uint16 completed = 0;
        freeMagnitudeToAdd = 0;
        while (nextIndex < pendingFreeMagnitudeLength && completed < numToComplete) {
            PendingFreeMagnitude memory pendingFreeMagnitude = _pendingFreeMagnitude[operator][strategy][nextIndex];
            // pendingFreeMagnitude is ordered by completableTimestamp. If we reach one that is not completable yet, then break
            // loop until completableTimestamp is < block.timestamp
            if (pendingFreeMagnitude.completableTimestamp < uint32(block.timestamp)) {
                break;
            }

            // pending free magnitude can be added to freeMagnitude
            freeMagnitudeToAdd += pendingFreeMagnitude.magnitudeDiff;
            ++nextIndex;
            ++completed;
        }
        return (freeMagnitudeToAdd, nextIndex);
    }

    /**
     * @param operator the operator to get the slashable magnitude for
     * @param strategies the strategies to get the slashable magnitude for
     *
     * @return operatorSets the operator sets the operator is a member of and the current slashable magnitudes for each strategy
     */
    function getCurrentSlashableMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (OperatorSet[] memory, uint64[][] memory) {
        OperatorSet[] memory operatorSets = avsDirectory.getOperatorSetsOfOperator(operator, 0, type(uint256).max);
        uint64[][] memory slashableMagnitudes = new uint64[][](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            slashableMagnitudes[i] = new uint64[](operatorSets.length);
            for (uint256 j = 0; j < operatorSets.length; ++j) {
                slashableMagnitudes[i][j] = uint64(
                    _magnitudeUpdate[operator][strategies[i]][_encodeOperatorSet(operatorSets[j])].upperLookupLinear(
                        uint32(block.timestamp)
                    )
                );
            }
        }
        return (operatorSets, slashableMagnitudes);
    }

    /**
     * @param operator the operator to get the slashable magnitude for
     * @param strategies the strategies to get the slashable magnitude for
     * @param timestamp the timestamp to get the slashable magnitude for
     *
     * @return operatorSets the operator sets the operator is a member of and the slashable magnitudes for each strategy
     */
    function getSlashableMagnitudes(
        address operator,
        IStrategy[] calldata strategies,
        uint32 timestamp
    ) external view returns (OperatorSet[] memory, uint64[][] memory) {
        OperatorSet[] memory operatorSets = avsDirectory.getOperatorSetsOfOperator(operator, 0, type(uint256).max);
        uint64[][] memory slashableMagnitudes = new uint64[][](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            slashableMagnitudes[i] = new uint64[](operatorSets.length);
            for (uint256 j = 0; j < operatorSets.length; ++j) {
                slashableMagnitudes[i][j] = uint64(
                    _magnitudeUpdate[operator][strategies[i]][_encodeOperatorSet(operatorSets[j])].upperLookupRecent(
                        timestamp
                    )
                );
            }
        }
        return (operatorSets, slashableMagnitudes);
    }

    /**
     * @param operator the operator to get the slashable ppm for
     * @param operatorSet the operatorSet to get the slashable ppm for
     * @param strategies the strategies to get the slashable ppm for
     * @param timestamp the timestamp to get the slashable ppm for for
     * @param linear whether the search should be linear (from the most recent) or binary
     *
     * @return slashablePPM the slashable ppm of the given list of strategies allocated to
     * the given OperatorSet for the given operator and timestamp
     */
    function getSlashablePPM(
        address operator,
        OperatorSet calldata operatorSet,
        IStrategy[] calldata strategies,
        uint32 timestamp,
        bool linear
    ) public view returns (uint24[] memory) {
        uint24[] memory slashablePPM = new uint24[](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            slashablePPM[i] = _getSlashablePPM(operator, operatorSet, strategies[i], timestamp, linear);
        }
        return slashablePPM;
    }

    function _getSlashablePPM(
        address operator,
        OperatorSet calldata operatorSet,
        IStrategy strategy,
        uint32 timestamp,
        bool linear
    ) public view returns (uint24) {
        uint64 totalMagnitude;
        if (linear) {
            totalMagnitude = uint64(_totalMagnitudeUpdate[operator][strategy].upperLookupLinear(timestamp));
        } else {
            totalMagnitude = uint64(_totalMagnitudeUpdate[operator][strategy].upperLookup(timestamp));
        }
        // return early if totalMagnitude is 0
        if (totalMagnitude == 0) {
            return 0;
        }

        uint64 currentMagnitude;
        bytes32 operatorSetKey = _encodeOperatorSet(operatorSet);
        if (linear) {
            currentMagnitude = uint64(_magnitudeUpdate[operator][strategy][operatorSetKey].upperLookupLinear(timestamp));
        } else {
            currentMagnitude = uint64(_magnitudeUpdate[operator][strategy][operatorSetKey].upperLookup(timestamp));
        }

        return uint16(currentMagnitude * 1e6 / totalMagnitude);
    }

    /**
     * @notice Get the allocatable magnitude for an operator and strategy based on number of pending deallocations
     * that could be completed at the same time. This is the sum of freeMagnitude and the sum of all pending completable deallocations.
     * @param operator the operator to get the allocatable magnitude for
     * @param strategy the strategy to get the allocatable magnitude for
     * @param numToComplete the number of pending free magnitudes deallocations to complete
     */
    function getAllocatableMagnitude(
        address operator,
        IStrategy strategy,
        uint16 numToComplete
    ) external view returns (uint64) {
        OperatorMagnitudeInfo storage info = operatorMagnitudeInfo[operator][strategy];
        (uint64 freeMagnitudeToAdd,) =
            _getPendingFreeMagnitude(operator, strategy, numToComplete, info.nextPendingFreeMagnitudeIndex);
        return info.freeMagnitude + freeMagnitudeToAdd;
    }

    /// @notice operator is slashable by operatorSet if currently registered OR last deregistered within 21 days
    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        (, uint32 lastDeregisteredTimestamp) =
            avsDirectory.operatorSetStatus(operatorSet.avs, operator, operatorSet.operatorSetId);
        return avsDirectory.isMember(operator, operatorSet)
            || lastDeregisteredTimestamp + SlashingConstants.DEALLOCATION_DELAY >= block.timestamp;
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
                totalMagnitudes[i] = SlashingConstants.INITIAL_TOTAL_MAGNITUDE;
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
            (uint224 value, uint256 pos,) =
                _totalMagnitudeUpdate[operator][strategies[i]].upperLookupRecentWithPos(timestamp);
            // if there is no existing total magnitude snapshot
            if (value == 0 && pos == 0) {
                totalMagnitudes[i] = SlashingConstants.INITIAL_TOTAL_MAGNITUDE;
            } else {
                totalMagnitudes[i] = uint64(value);
            }
        }
        return totalMagnitudes;
    }

    /**
     * @notice Returns the current total magnitude of an operator for a given strategy
     * @param operator the operator to get the total magnitude for
     * @param strategy the strategy to get the total magnitude for
     * @return totalMagnitude the total magnitude for the strategy
     */
    function getTotalMagnitude(address operator, IStrategy strategy) external view returns (uint64) {
        uint64 totalMagnitude;
        (bool exists,, uint224 value) = _totalMagnitudeUpdate[operator][strategy].latestSnapshot();
        if (!exists) {
            totalMagnitude = SlashingConstants.INITIAL_TOTAL_MAGNITUDE;
        } else {
            totalMagnitude = uint64(value);
        }

        return totalMagnitude;
    }

    /**
     * @notice Returns the total magnitude of an operator for a given strategy at a given timestamp
     * @param operator the operator to get the total magnitude for
     * @param strategy the strategy to get the total magnitude for
     * @param timestamp the timestamp to get the total magnitude at
     * @return totalMagnitude the total magnitude for the strategy
     */
    function getTotalMagnitudeAtTimestamp(
        address operator,
        IStrategy strategy,
        uint32 timestamp
    ) external view returns (uint64) {
        uint64 totalMagnitude = SlashingConstants.INITIAL_TOTAL_MAGNITUDE;
        (uint224 value, uint256 pos,) = _totalMagnitudeUpdate[operator][strategy].upperLookupRecentWithPos(timestamp);

        // if there is no existing total magnitude snapshot
        if (value == 0 && pos == 0) {
            totalMagnitude = SlashingConstants.INITIAL_TOTAL_MAGNITUDE;
        } else {
            totalMagnitude = uint64(value);
        }

        return totalMagnitude;
    }

    // /**
    //  * @notice fetches the minimum slashable shares for a certain operator and operatorSet for a list of strategies
    //  * from the current timestamp until the given timestamp
    //  *
    //  * @param operator the operator to get the minimum slashable shares for
    //  * @param operatorSet the operatorSet to get the minimum slashable shares for
    //  * @param strategies the strategies to get the minimum slashable shares for
    //  * @param timestamp the timestamp to the minimum slashable shares before
    //  *
    //  * @dev used to get the slashable stakes of operators to weigh over a given slashability window
    //  *
    //  * @return the list of share amounts for each strategy
    //  */
    // function getMinimumSlashableSharesBefore(
    //     address operator,
    //     OperatorSet calldata operatorSet,
    //     IStrategy[] calldata strategies,
    //     uint32 timestamp
    // ) external view returns (uint256[] calldata) {}

    /// @dev Verify operator's signature and spend salt
    function _verifyOperatorSignature(
        address operator,
        MagnitudeAllocation[] calldata allocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) internal {
        // check the signature expiry
        require(operatorSignature.expiry >= block.timestamp, SignatureExpired());
        // Assert operator's signature cannot be replayed.
        require(!avsDirectory.operatorSaltIsSpent(operator, operatorSignature.salt), SaltSpent());

        bytes32 digestHash = calculateMagnitudeAllocationDigestHash(
            operator, allocations, operatorSignature.salt, operatorSignature.expiry
        );

        // Assert operator's signature is valid.
        EIP1271SignatureUtils.checkSignature_EIP1271(operator, digestHash, operatorSignature.signature);
        // Spend salt.
        operatorSaltIsSpent[operator][operatorSignature.salt] = true;
    }

    /**
     * @notice Calculates the digest hash to be signed by an operator to modify magnitude allocations
     * @param operator The operator to allocate or deallocate magnitude for.
     * @param allocations The magnitude allocations/deallocations to be made.
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid.
     */
    function calculateMagnitudeAllocationDigestHash(
        address operator,
        MagnitudeAllocation[] calldata allocations,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        return _calculateDigestHash(
            keccak256(abi.encode(MAGNITUDE_ADJUSTMENT_TYPEHASH, operator, allocations, salt, expiry))
        );
    }

    /// @notice Getter function for the current EIP-712 domain separator for this contract.
    /// @dev The domain separator will change in the event of a fork that changes the ChainID.
    function domainSeparator() public view override returns (bytes32) {
        return _calculateDomainSeparator();
    }

    /// @notice Internal function for calculating the current domain separator of this contract
    function _calculateDomainSeparator() internal view returns (bytes32) {
        if (block.chainid == ORIGINAL_CHAIN_ID) {
            return _DOMAIN_SEPARATOR;
        } else {
            return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("EigenLayer")), block.chainid, address(this)));
        }
    }

    /// @notice Returns an EIP-712 encoded hash struct.
    function _calculateDigestHash(bytes32 structHash) internal view returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", _calculateDomainSeparator(), structHash));
    }

    /// @dev Returns an `OperatorSet` encoded into a 32-byte value.
    /// @param operatorSet The `OperatorSet` to encode.
    function _encodeOperatorSet(OperatorSet memory operatorSet) internal pure returns (bytes32) {
        return bytes32(abi.encodePacked(operatorSet.avs, uint96(operatorSet.operatorSetId)));
    }

    /// @dev Returns an `OperatorSet` decoded from an encoded 32-byte value.
    /// @param encoded The encoded `OperatorSet` to decode.
    /// @dev Assumes `encoded` is encoded via `_encodeOperatorSet(operatorSet)`.
    function _decodeOperatorSet(bytes32 encoded) internal pure returns (OperatorSet memory) {
        return OperatorSet({
            avs: address(uint160(uint256(encoded) >> 96)),
            operatorSetId: uint32(uint256(encoded) & type(uint96).max)
        });
    }
}
