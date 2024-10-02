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
    ) 
        AllocationManagerStorage(
            _delegation, 
            _avsDirectory, 
            _DEALLOCATION_DELAY, 
            _ALLOCATION_CONFIGURATION_DELAY
        ) 
    {
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
        _DOMAIN_SEPARATOR = _calculateDomainSeparator();
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
    function completePendingDeallocations(
        address operator,
        IStrategy[] calldata strategies,
        uint16[] calldata numToComplete
    ) external onlyWhenNotPaused(PAUSED_STAKE_ALLOCATIONS_AND_DEALLOCATIONS) {
        require(strategies.length == numToComplete.length, InputArrayLengthMismatch());
        require(delegation.isOperator(operator), OperatorNotRegistered());
        for (uint256 i = 0; i < strategies.length; ++i) {
            _completePendingDeallocations({operator: operator, strategy: strategies[i], numToComplete: numToComplete[i]});
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
    ) external onlyWhenNotPaused(PAUSED_STAKE_ALLOCATIONS_AND_DEALLOCATIONS) {
        if (msg.sender != operator) {
            _verifyOperatorSignature(operator, allocations, operatorSignature);
        }
        require(delegation.isOperator(operator), OperatorNotRegistered());

        (bool isSet, uint32 operatorAllocationDelay) = allocationDelay(operator);
        require(isSet, UninitializedAllocationDelay());

        for (uint256 i = 0; i < allocations.length; ++i) {
            MagnitudeAllocation calldata allocation = allocations[i];
            require(allocation.operatorSets.length == allocation.magnitudes.length, InputArrayLengthMismatch());
            require(avsDirectory.isOperatorSetBatch(allocation.operatorSets), InvalidOperatorSet());
            
            // 1. For the given (operator,strategy) complete any pending deallocations to update free magnitude
            _completePendingDeallocations({
                operator: operator,
                strategy: allocation.strategy,
                numToComplete: type(uint16).max
            });

            {
                // 2. Check current totalMagnitude matches expected value. This is to check for slashing race conditions
                // where an operator gets slashed from an operatorSet and as a result all the configured allocations have larger
                // proprtional magnitudes relative to each other.

                // Load the operator's total magnitude for the strategy.
                (bool exists,, uint224 currentTotalMagnitude) = _totalMagnitudeUpdate[operator][allocation.strategy].latestSnapshot();

                // If the operator has no total magnitude snapshot, set it to WAD, which denotes an unslashed operator.
                if (!exists) {
                    currentTotalMagnitude = WAD;
                    _totalMagnitudeUpdate[operator][allocation.strategy].push({key: uint32(block.timestamp), value: currentTotalMagnitude});
                    operatorFreeMagnitudeInfo[operator][allocation.strategy].freeMagnitude = uint64(currentTotalMagnitude);
                }

                // 3. set allocations for the strategy after updating freeMagnitude
                require(uint64(currentTotalMagnitude) == allocation.expectedTotalMagnitude, InvalidExpectedTotalMagnitude());
            }

            for (uint256 j = 0; j < allocation.operatorSets.length; ++j) {
                // Check that there are no pending allocations & deallocations for the operator, operatorSet, strategy
                MagnitudeInfo memory mInfo = _getCurrentEffectiveMagnitude(operator, allocation.strategy, _encodeOperatorSet(allocation.operatorSets[j]));
                require(block.timestamp >= mInfo.effectTimestamp, ModificationAlreadyPending());

                // Calculate the new pending diff with this modification
                mInfo.pendingMagnitudeDiff = int128(uint128(allocation.magnitudes[j])) - int128(uint128(mInfo.currentMagnitude));
                require(mInfo.pendingMagnitudeDiff != 0, SameMagnitude());
                
                // Handle deallocation/allocation and modification effect timestamp
                if (mInfo.pendingMagnitudeDiff < 0) {
                    // This is a deallocation

                    // 1. push PendingFreeMagnitude and respective array index into (op,opSet,Strategy) queued deallocations
                    deallocationQueue[operator][allocation.strategy].push(_encodeOperatorSet(allocation.operatorSets[j]));

                    // 2. Update the effect timestamp for the deallocation
                    mInfo.effectTimestamp = uint32(block.timestamp) + DEALLOCATION_DELAY;
                } else if (mInfo.pendingMagnitudeDiff > 0) {
                    // This is an allocation

                    // 1. decrement free magnitude by incremented amount
                    uint64 magnitudeToAllocate = uint64(uint128(mInfo.pendingMagnitudeDiff));
                    FreeMagnitudeInfo memory freeInfo = operatorFreeMagnitudeInfo[operator][allocation.strategy];
                    require(freeInfo.freeMagnitude >= magnitudeToAllocate, InsufficientAllocatableMagnitude());
                    freeInfo.freeMagnitude -= magnitudeToAllocate;

                    // 2. Update the effectTimestamp for the allocation
                    mInfo.effectTimestamp = uint32(block.timestamp) + operatorAllocationDelay; 

                    operatorFreeMagnitudeInfo[operator][allocation.strategy] = freeInfo;
                }

                // Allocate magnitude which will take effect at the `effectTimestamp`
                _operatorMagnitudeInfo[operator][allocation.strategy][_encodeOperatorSet(allocation.operatorSets[j])] = mInfo;
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
     * @param bipsToSlash the number of bips to slash, this will be proportional to the
     * operator's slashable stake allocation for the operatorSet
     */
    function slashOperator(
        address operator,
        uint32 operatorSetId,
        IStrategy[] calldata strategies,
        uint16 bipsToSlash
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_SLASHING) {
        require(0 < bipsToSlash && bipsToSlash <= BIPS_FACTOR, InvalidBipsToSlash());
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        require(avsDirectory.isOperatorSlashable(operator, operatorSet), InvalidOperator());
        bytes32 operatorSetKey = _encodeOperatorSet(operatorSet);

        for (uint256 i = 0; i < strategies.length; ++i) {
            // 1. Slash from pending deallocations and allocations
            MagnitudeInfo memory mInfo = _getCurrentEffectiveMagnitude(operator, strategies[i], operatorSetKey);
            
            uint64 slashedMagnitude = uint64(mInfo.currentMagnitude * bipsToSlash / BIPS_FACTOR);
            mInfo.currentMagnitude -= slashedMagnitude;
            // if there is a pending deallocation, slash pending deallocation proportionally
            if (mInfo.pendingMagnitudeDiff < 0) {
                uint128 slashedPending = uint128(uint128(-mInfo.pendingMagnitudeDiff )* bipsToSlash / BIPS_FACTOR);
                mInfo.pendingMagnitudeDiff += int128(slashedPending);
            }
            // update operatorMagnitudeInfo
            _operatorMagnitudeInfo[operator][strategies[i]][operatorSetKey] = mInfo;

            // 2. update totalMagnitude, get total magnitude and subtract slashedMagnitude
            // this will be reflected in the conversion of delegatedShares to shares in the DM
            Snapshots.History storage totalMagnitudes = _totalMagnitudeUpdate[operator][strategies[i]];
            totalMagnitudes.push({key: uint32(block.timestamp), value: totalMagnitudes.latest() - slashedMagnitude});
        }
    }

    /**
     * @notice Called by an operator to cancel a salt that has been used to register with an AVS.
     *
     * @param salt A unique and single use value associated with the approver signature.
     */
    function cancelSalt(
        bytes32 salt
    ) external override {
        // Mutate `operatorSaltIsSpent` to `true` to prevent future spending.
        operatorSaltIsSpent[msg.sender][salt] = true;
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
     * @notice For a single strategy, complete pending deallocations and free up their magnitude
     * @param operator address to update freeMagnitude for
     * @param strategy the strategy to update freeMagnitude for
     * @param numToComplete the number of pending free magnitudes deallocations to complete
     * @dev read through pending free magnitudes and add to freeMagnitude if completableTimestamp is >= block timestamp
     * In addition to updating freeMagnitude, updates next starting index to read from for pending free magnitudes after completing
     */
    function _completePendingDeallocations(address operator, IStrategy strategy, uint16 numToComplete) internal {
        FreeMagnitudeInfo memory freeInfo = operatorFreeMagnitudeInfo[operator][strategy];
        
        uint256 numDeallocations = deallocationQueue[operator][strategy].length;
        uint256 completed;

        while (freeInfo.nextPendingIndex < numDeallocations && completed < numToComplete) {
            bytes32 opsetKey = deallocationQueue[operator][strategy][freeInfo.nextPendingIndex];
            MagnitudeInfo memory mInfo = _operatorMagnitudeInfo[operator][strategy][opsetKey];

            // deallocationQueue is ordered by `effectTimestamp`. If we reach a pending deallocation
            // that cannot be completed, we're done.
            if (block.timestamp < mInfo.effectTimestamp) {
                break;
            }

            // We know that this is a deallocation because this `opsetKey` was in the pending deallocation set
            // Therefore, `pendingMagnitudeDiff` MUST be negative.
            uint64 freeMagnitudeToAdd = uint64(uint128(-mInfo.pendingMagnitudeDiff));
            mInfo.pendingMagnitudeDiff = 0;
            mInfo.currentMagnitude -= freeMagnitudeToAdd;

            // Add newly-freed magnitude to FreeMagnitudeInfo
            freeInfo.freeMagnitude += freeMagnitudeToAdd;
            freeInfo.nextPendingIndex++;
            ++completed;

            // Update MagnitudeInfo in storage
            _operatorMagnitudeInfo[operator][strategy][opsetKey] = mInfo;
        }

        operatorFreeMagnitudeInfo[operator][strategy] = freeInfo;
    }

    /// @dev Fetch the operator's current magnitude, applying a pending diff if the effect timestamp is passed
    /// @notice This may return something that is not recorded in state. Remember to store this updated value if needed!
    function _getCurrentEffectiveMagnitude(address operator, IStrategy strategy, bytes32 operatorSetKey) internal view returns (MagnitudeInfo memory) {
        MagnitudeInfo memory mInfo = _operatorMagnitudeInfo[operator][strategy][operatorSetKey];

        // If the magnitude change is not yet in effect, return unaltered
        if (block.timestamp < mInfo.effectTimestamp) {
            return mInfo;
        }

        // Otherwise, calculate the new current magnitude and return the modified struct
        if (mInfo.pendingMagnitudeDiff >= 0) {
            mInfo.currentMagnitude += uint64(uint128(mInfo.pendingMagnitudeDiff));
        } else {
            mInfo.currentMagnitude -= uint64(uint128(-mInfo.pendingMagnitudeDiff));
        }

        mInfo.pendingMagnitudeDiff = 0;
        return mInfo;
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
                MagnitudeInfo memory mInfo = _getCurrentEffectiveMagnitude(operator, strategies[i], _encodeOperatorSet(operatorSets[j]));
                slashableMagnitudes[i][j] = mInfo.currentMagnitude;
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
        FreeMagnitudeInfo memory info = operatorFreeMagnitudeInfo[operator][strategy];
        uint256 numDeallocations = deallocationQueue[operator][strategy].length;
        uint64 freeMagnitudeToAdd = 0;
        for (uint192 i = info.nextPendingIndex; i < numDeallocations; ++i) {
            bytes32 opsetKey = deallocationQueue[operator][strategy][i];
            MagnitudeInfo memory opsetMagnitudeInfo = _operatorMagnitudeInfo[operator][strategy][opsetKey];
            if (block.timestamp < opsetMagnitudeInfo.effectTimestamp) {
                break;
            }
            freeMagnitudeToAdd += uint64(uint128(-opsetMagnitudeInfo.pendingMagnitudeDiff));
        }
        return info.freeMagnitude + freeMagnitudeToAdd;
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
     * @notice Returns the current total magnitude of an operator for a given strategy
     * @param operator the operator to get the total magnitude for
     * @param strategy the strategy to get the total magnitude for
     * @return totalMagnitude the total magnitude for the strategy
     */
    function getTotalMagnitude(address operator, IStrategy strategy) external view returns (uint64) {
        uint64 totalMagnitude;
        (bool exists,, uint224 value) = _totalMagnitudeUpdate[operator][strategy].latestSnapshot();
        if (!exists) {
            totalMagnitude = WAD;
        } else {
            totalMagnitude = uint64(value);
        }
        return totalMagnitude;
    }

    /**
     * @notice Returns the latest pending allocation of an operator for a given strategy and operatorSets.
     * One of the assumptions here is we don't allow more than one pending allocation for an operatorSet at a time.
     * If that changes, we would need to change this function to return all pending allocations for an operatorSet.
     * @param operator the operator to get the pending allocations for
     * @param strategy the strategy to get the pending allocations for
     * @param operatorSets the operatorSets to get the pending allocations for
     * @return pendingMagnitudes the pending allocations for each operatorSet
     * @return timestamps the timestamps for each pending allocation
     */
    function getPendingAllocations(
        address operator,
        IStrategy strategy,
        OperatorSet[] calldata operatorSets
    ) external view returns (uint64[] memory pendingMagnitudes, uint32[] memory timestamps) {
        pendingMagnitudes = new uint64[](operatorSets.length);
        timestamps = new uint32[](operatorSets.length);
        for (uint256 i = 0; i < operatorSets.length; ++i) {
            MagnitudeInfo memory opsetMagnitudeInfo = _operatorMagnitudeInfo[operator][strategy][_encodeOperatorSet(operatorSets[i])];

            if (opsetMagnitudeInfo.effectTimestamp < block.timestamp && opsetMagnitudeInfo.pendingMagnitudeDiff > 0) {
                pendingMagnitudes[i] = opsetMagnitudeInfo.currentMagnitude + uint64(uint128(opsetMagnitudeInfo.pendingMagnitudeDiff));
                timestamps[i] = opsetMagnitudeInfo.effectTimestamp;
            } else {
                pendingMagnitudes[i] = 0;
                timestamps[i] = 0;
            }
        }
    }

    /**
     * @notice Returns the pending deallocations of an operator for a given strategy and operatorSets.
     * One of the assumptions here is we don't allow more than one pending deallocation for an operatorSet at a time.
     * If that changes, we would need to change this function to return all pending deallocations for an operatorSet.
     * @param operator the operator to get the pending deallocations for
     * @param strategy the strategy to get the pending deallocations for
     * @param operatorSets the operatorSets to get the pending deallocations for
     * @return pendingMagnitudeDiffs the pending deallocation diffs for each operatorSet
     * @return timestamps the timestamps for each pending dealloction
     */
    function getPendingDeallocations(
        address operator,
        IStrategy strategy,
        OperatorSet[] calldata operatorSets
    ) external view returns (uint64[] memory pendingMagnitudeDiffs, uint32[] memory timestamps) {
        pendingMagnitudeDiffs = new uint64[](operatorSets.length);
        timestamps = new uint32[](operatorSets.length);

        for (uint256 i = 0; i < operatorSets.length; ++i) {
            MagnitudeInfo memory opsetMagnitudeInfo = _operatorMagnitudeInfo[operator][strategy][_encodeOperatorSet(operatorSets[i])];

            if (opsetMagnitudeInfo.effectTimestamp < block.timestamp && opsetMagnitudeInfo.pendingMagnitudeDiff < 0) {
                pendingMagnitudeDiffs[i] = uint64(uint128(-opsetMagnitudeInfo.pendingMagnitudeDiff));
                timestamps[i] = opsetMagnitudeInfo.effectTimestamp;
                pendingMagnitudeDiffs[i] = 0;
            }
        }
    }

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
    function _calculateDigestHash(
        bytes32 structHash
    ) internal view returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", _calculateDomainSeparator(), structHash));
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
