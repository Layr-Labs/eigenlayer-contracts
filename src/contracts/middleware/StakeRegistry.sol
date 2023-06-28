// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/utils/math/Math.sol";
import "../interfaces/IServiceManager.sol";
import "../interfaces/IStakeRegistry.sol";
import "../interfaces/IRegistryCoordinator.sol";
import "./StakeRegistryStorage.sol";

/**
 * @title A `Registry` that keeps track of stakes of operators for up to 256 quorums.
 * Specifically, it keeps track of
 *      1) The stake of each operator in all the quorums they are a part of for block ranges
 *      2) The total stake of all operators in each quorum for block ranges
 *      3) The minimum stake required to register for each quorum
 * It allows an additional functionality (in addition to registering and deregistering) to update the stake of an operator.
 * @author Layr Labs, Inc.
 */
contract StakeRegistry is StakeRegistryStorage {
    
    modifier onlyRegistryCoordinator() {
        require(msg.sender == address(registryCoordinator), "StakeRegistry.onlyRegistryCoordinator: caller is not the RegistryCoordinator");
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator,
        IStrategyManager _strategyManager,
        IServiceManager _serviceManager
    ) StakeRegistryStorage(_registryCoordinator, _strategyManager, _serviceManager)
    // solhint-disable-next-line no-empty-blocks
    {
    }

    /**
     * @notice Sets the minimum stake for each quorum and adds `_quorumStrategiesConsideredAndMultipliers` for each 
     * quorum the Registry is being initialized with
     */
    function initialize(
        uint96[] memory _minimumStakeForQuorum,
        StrategyAndWeightingMultiplier[][] memory _quorumStrategiesConsideredAndMultipliers
    ) external virtual initializer {
        _initialize(_minimumStakeForQuorum, _quorumStrategiesConsideredAndMultipliers);
    }

    function _initialize(
        uint96[] memory _minimumStakeForQuorum,
        StrategyAndWeightingMultiplier[][] memory _quorumStrategiesConsideredAndMultipliers
    ) internal virtual onlyInitializing {
        // sanity check lengths
        require(_minimumStakeForQuorum.length == _quorumStrategiesConsideredAndMultipliers.length, "Registry._initialize: minimumStakeForQuorum length mismatch");

        // add the strategies considered and multipliers for each quorum
        for (uint8 quorumNumber = 0; quorumNumber < _quorumStrategiesConsideredAndMultipliers.length;) {
            _setMinimumStakeForQuorum(quorumNumber, _minimumStakeForQuorum[quorumNumber]);
            _createQuorum(_quorumStrategiesConsideredAndMultipliers[quorumNumber]);
            unchecked {
                ++quorumNumber;
            }
        }
    }

    /**
     * @notice Returns the `index`-th entry in the `operatorIdToStakeHistory[operatorId][quorumNumber]` array.
     * @param quorumNumber The quorum number to get the stake for.
     * @param operatorId The id of the operator of interest.
     * @param index Array index for lookup, within the dynamic array `operatorIdToStakeHistory[operatorId][quorumNumber]`.
     * @dev Function will revert if `index` is out-of-bounds.
     */
    function getStakeUpdateForQuorumFromOperatorIdAndIndex(uint8 quorumNumber, bytes32 operatorId, uint256 index)
        external
        view
        returns (OperatorStakeUpdate memory)
    {
        return operatorIdToStakeHistory[operatorId][quorumNumber][index];
    }

    /**
     * @notice Returns the `index`-th entry in the dynamic array of total stake, `_totalStakeHistory` for quorum `quorumNumber`.
     * @param quorumNumber The quorum number to get the stake for.
     * @param index Array index for lookup, within the dynamic array `_totalStakeHistory[quorumNumber]`.
     */
    function getTotalStakeUpdateForQuorumFromIndex(uint8 quorumNumber, uint256 index) external view returns (OperatorStakeUpdate memory) {
        return _totalStakeHistory[quorumNumber][index];
    }

    /// @notice Returns the indices of the operator stakes for the provided `quorumNumber` at the given `blockNumber`
    function getStakeUpdateIndexForOperatorIdForQuorumAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber)
        external
        view
        returns (uint32)
    {
        uint32 length = uint32(operatorIdToStakeHistory[operatorId][quorumNumber].length);
        for (uint32 i = 0; i < length; i++) {
            if (operatorIdToStakeHistory[operatorId][quorumNumber][length - i - 1].updateBlockNumber <= blockNumber) {
                return length - i - 1;
            }
        }
        revert("StakeRegistry.getStakeUpdateIndexForOperatorIdForQuorumAtBlockNumber: no stake update found for operatorId and quorumNumber");
    }


    /// @notice Returns the indices of the total stakes for the provided `quorumNumbers` at the given `blockNumber`
    function getTotalStakeIndicesByQuorumNumbersAtBlockNumber(uint32 blockNumber, bytes calldata quourmNumbers) external view returns(uint32[] memory) {
        uint32[] memory indices = new uint32[](quourmNumbers.length);
        for (uint256 i = 0; i < quourmNumbers.length; i++) {
            uint8 quorumNumber = uint8(quourmNumbers[i]);
            uint32 length = uint32(_totalStakeHistory[quorumNumber].length);
            for (uint32 j = 0; j < length; j++) {
                if (_totalStakeHistory[quorumNumber][length - j - 1].updateBlockNumber <= blockNumber) {
                    indices[i] = length - j - 1;
                    break;
                }
            }
        }
        return indices;
    }

    /**
     * @notice Returns the stake weight corresponding to `operatorId` for quorum `quorumNumber`, at the
     * `index`-th entry in the `operatorIdToStakeHistory[operatorId][quorumNumber]` array if it was the operator's
     * stake at `blockNumber`. Reverts otherwise.
     * @param quorumNumber The quorum number to get the stake for.
     * @param operatorId The id of the operator of interest.
     * @param index Array index for lookup, within the dynamic array `operatorIdToStakeHistory[operatorId][quorumNumber]`.
     * @param blockNumber Block number to make sure the stake is from.
     * @dev Function will revert if `index` is out-of-bounds.
     */
    function getStakeForQuorumAtBlockNumberFromOperatorIdAndIndex(uint8 quorumNumber, uint32 blockNumber, bytes32 operatorId, uint256 index)
        external
        view
        returns (uint96)
    {
        OperatorStakeUpdate memory operatorStakeUpdate = operatorIdToStakeHistory[operatorId][quorumNumber][index];
        _validateOperatorStakeUpdateAtBlockNumber(operatorStakeUpdate, blockNumber);
        return operatorStakeUpdate.stake;
    }

    /**
     * @notice Returns the total stake weight for quorum `quorumNumber`, at the `index`-th entry in the 
     * `_totalStakeHistory[quorumNumber]` array if it was the stake at `blockNumber`. Reverts otherwise.
     * @param quorumNumber The quorum number to get the stake for.
     * @param index Array index for lookup, within the dynamic array `_totalStakeHistory[quorumNumber]`.
     * @param blockNumber Block number to make sure the stake is from.
     * @dev Function will revert if `index` is out-of-bounds.
     */
    function getTotalStakeAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) external view returns (uint96) {
        OperatorStakeUpdate memory totalStakeUpdate = _totalStakeHistory[quorumNumber][index];
        _validateOperatorStakeUpdateAtBlockNumber(totalStakeUpdate, blockNumber);
        return totalStakeUpdate.stake;
    }

    /**
     * @notice Returns the most recent stake weight for the `operatorId` for a certain quorum
     * @dev Function returns an OperatorStakeUpdate struct with **every entry equal to 0** in the event that the operator has no stake history
     */
    function getMostRecentStakeUpdateByOperatorId(bytes32 operatorId, uint8 quorumNumber) public view returns (OperatorStakeUpdate memory) {
        uint256 historyLength = operatorIdToStakeHistory[operatorId][quorumNumber].length;
        OperatorStakeUpdate memory operatorStakeUpdate;
        if (historyLength == 0) {
            return operatorStakeUpdate;
        } else {
            operatorStakeUpdate = operatorIdToStakeHistory[operatorId][quorumNumber][historyLength - 1];
            return operatorStakeUpdate;
        }
    }

    function getStakeHistoryLengthForQuorumNumber(bytes32 operatorId, uint8 quorumNumber) external view returns (uint256) {
        return operatorIdToStakeHistory[operatorId][quorumNumber].length;
    }

    /**
     * @notice Returns the most recent stake weight for the `operatorId` for quorum `quorumNumber`
     * @dev Function returns weight of **0** in the event that the operator has no stake history
     */
    function getCurrentOperatorStakeForQuorum(bytes32 operatorId, uint8 quorumNumber) external view returns (uint96) {
        OperatorStakeUpdate memory operatorStakeUpdate = getMostRecentStakeUpdateByOperatorId(operatorId, quorumNumber);
        return operatorStakeUpdate.stake;
    }

    /**
     * @notice Returns the stake weight from the latest entry in `_totalStakeHistory` for quorum `quorumNumber`.
     * @dev Will revert if `_totalStakeHistory[quorumNumber]` is empty.
     */
    function getCurrentTotalStakeForQuorum(uint8 quorumNumber) external view returns (uint96) {
        // no chance of underflow / error in next line, since an empty entry is pushed in the constructor
        return _totalStakeHistory[quorumNumber][_totalStakeHistory[quorumNumber].length - 1].stake;
    }

    function getLengthOfOperatorIdStakeHistoryForQuorum(bytes32 operatorId, uint8 quorumNumber) external view returns (uint256) {
        return operatorIdToStakeHistory[operatorId][quorumNumber].length;
    }

    function getLengthOfTotalStakeHistoryForQuorum(uint8 quorumNumber) external view returns (uint256) {
        return _totalStakeHistory[quorumNumber].length;
    }

    // MUTATING FUNCTIONS

    /// @notice Adjusts the `minimumStakeFirstQuorum` -- i.e. the node stake (weight) requirement for inclusion in the 1st quorum.
    function setMinimumStakeForQuorum(uint8 quorumNumber, uint96 minimumStake) external onlyServiceManagerOwner {
        _setMinimumStakeForQuorum(quorumNumber, minimumStake);
    }

    /**
     * @notice Registers the `operator` with `operatorId` for the specified `quorumNumbers`.
     * @param operator The address of the operator to register.
     * @param operatorId The id of the operator to register.
     * @param quorumNumbers The quorum numbers the operator is registering for, where each byte is an 8 bit integer quorumNumber.
     * @dev access restricted to the RegistryCoordinator
     * @dev Preconditions (these are assumed, not validated in this contract):
     *         1) `quorumNumbers` has no duplicates
     *         2) `quorumNumbers.length` != 0
     *         3) `quorumNumbers` is ordered in ascending order
     *         4) the operator is not already registered
     */
    function registerOperator(address operator, bytes32 operatorId, bytes calldata quorumNumbers) external virtual onlyRegistryCoordinator {
        _registerOperator(operator, operatorId, quorumNumbers);
    }

    /**
     * @notice Deregisters the operator with `operatorId` for the specified `quorumNumbers`.
     * @param operatorId The id of the operator to deregister.
     * @param quorumNumbers The quorum numbers the operator is deregistering from, where each byte is an 8 bit integer quorumNumber.
     * @dev access restricted to the RegistryCoordinator
     * @dev Preconditions (these are assumed, not validated in this contract):
     *         1) `quorumNumbers` has no duplicates
     *         2) `quorumNumbers.length` != 0
     *         3) `quorumNumbers` is ordered in ascending order
     *         4) the operator is not already deregistered
     *         5) `quorumNumbers` is a subset of the quorumNumbers that the operator is registered for
     */
    function deregisterOperator(bytes32 operatorId, bytes calldata quorumNumbers) external virtual onlyRegistryCoordinator {
        _deregisterOperator(operatorId, quorumNumbers);
    }

    /**
     * @notice Used for updating information on deposits of nodes.
     * @param operators are the addresses of the operators whose stake information is getting updated
     * @param operatorIds are the ids of the operators whose stake information is getting updated
     * @param prevElements are the elements before this middleware in the operator's linked list within the slasher
     * @dev Precondition:
     *          1) `quorumBitmaps[i]` should be the bitmap that represents the quorums that `operators[i]` registered for
     * @dev reverts if there are no operators registered with index out of bounds
     */
    function updateStakes(address[] calldata operators, bytes32[] calldata operatorIds, uint256[] calldata prevElements) external {
        // for each quorum, loop through operators and see if they are a part of the quorum
        // if they are, get their new weight and update their individual stake history and the
        // quorum's total stake history accordingly
        for (uint8 quorumNumber = 0; quorumNumber < quorumCount;) {
            OperatorStakeUpdate memory totalStakeUpdate;
            // for each operator
            for(uint i = 0; i < operatorIds.length;) {
                uint192 quorumBitmap = registryCoordinator.getCurrentQuorumBitmapByOperatorId(operatorIds[i]);
                // if the operator is a part of the quorum
                if (quorumBitmap >> quorumNumber & 1 == 1) {
                    // if the total stake has not been loaded yet, load it
                    if (totalStakeUpdate.updateBlockNumber == 0) {
                        totalStakeUpdate = _totalStakeHistory[quorumNumber][_totalStakeHistory[quorumNumber].length - 1];
                    }
                    bytes32 operatorId = operatorIds[i];
                    // update the operator's stake based on current state
                    (uint96 stakeBeforeUpdate, uint96 stakeAfterUpdate) = _updateOperatorStake(operators[i], operatorId, quorumNumber);
                    // calculate the new total stake for the quorum
                    totalStakeUpdate.stake = totalStakeUpdate.stake - stakeBeforeUpdate + stakeAfterUpdate;
                }
                unchecked {
                    ++i;
                }
            }
            // if the total stake for this quorum was updated, record it in storage
            if (totalStakeUpdate.updateBlockNumber != 0) {
                // update the total stake history for the quorum
                _recordTotalStakeUpdate(quorumNumber, totalStakeUpdate);
            }
            unchecked {
                ++quorumNumber;
            }
        }

        // record stake updates in the EigenLayer Slasher
        for (uint i = 0; i < operators.length;) {
            serviceManager.recordStakeUpdate(operators[i], uint32(block.number), serviceManager.latestServeUntilBlock(), prevElements[i]);
            unchecked {
                ++i;
            }
        }     
    }

    // INTERNAL FUNCTIONS

    function _setMinimumStakeForQuorum(uint8 quorumNumber, uint96 minimumStake) internal {
        minimumStakeForQuorum[quorumNumber] = minimumStake;
        emit MinimumStakeForQuorumUpdated(quorumNumber, minimumStake);
    }

    /** 
     * @notice Updates the stake for the operator with `operatorId` for the specified `quorumNumbers`. The total stake
     * for each quorum is updated accordingly in addition to the operator's individual stake history.
     */ 
    function _registerOperator(address operator, bytes32 operatorId, bytes memory quorumNumbers) internal {
        // check the operator is registering for only valid quorums
        require(uint8(quorumNumbers[quorumNumbers.length - 1]) < quorumCount, "StakeRegistry._registerOperator: greatest quorumNumber must be less than quorumCount");
        OperatorStakeUpdate memory _newTotalStakeUpdate;
        // add the `updateBlockNumber` info
        _newTotalStakeUpdate.updateBlockNumber = uint32(block.number);
        // for each quorum, evaluate stake and add to total stake
        for (uint8 quorumNumbersIndex = 0; quorumNumbersIndex < quorumNumbers.length;) {
            // get the next quorumNumber
            uint8 quorumNumber = uint8(quorumNumbers[quorumNumbersIndex]);
            // evaluate the stake for the operator
            // since we don't use the first output, this will use 1 extra sload when deregistered operator's register again
            (, uint96 stake) = _updateOperatorStake(operator, operatorId, quorumNumber);
            // @JEFF: This reverts pretty late, but i think that's fine. wdyt?
            // check if minimum requirement has been met, will be 0 if not
            require(stake != 0, "StakeRegistry._registerOperator: Operator does not meet minimum stake requirement for quorum");
            // add operator stakes to total stake before update (in memory)
            uint256 _totalStakeHistoryLength = _totalStakeHistory[quorumNumber].length;
            // add calculate the total stake for the quorum
            uint96 totalStakeAfterUpdate = stake;
            if (_totalStakeHistoryLength != 0) {
                // only add the stake if there is a previous total stake
                // overwrite `stake` variable
                totalStakeAfterUpdate += _totalStakeHistory[quorumNumber][_totalStakeHistoryLength - 1].stake;
            }
            _newTotalStakeUpdate.stake = totalStakeAfterUpdate;
            // update storage of total stake
            _recordTotalStakeUpdate(quorumNumber, _newTotalStakeUpdate);
            unchecked {
                ++quorumNumbersIndex;
            }
        }
    }

    /**
     * @notice Removes the stakes of the operator with `operatorId` from the quorums specified in `quorumNumbers`
     * the total stake of the quorums specified in `quorumNumbers` will be updated and so will the operator's individual
     * stake updates. These operator's individual stake updates will have a 0 stake value for the latest update.
     */
    function _deregisterOperator(bytes32 operatorId, bytes memory quorumNumbers) internal {
        // check the operator is deregistering from only valid quorums
        OperatorStakeUpdate memory _operatorStakeUpdate;
        // add the `updateBlockNumber` info
        _operatorStakeUpdate.updateBlockNumber = uint32(block.number);
        OperatorStakeUpdate memory _newTotalStakeUpdate;
        // add the `updateBlockNumber` info
        _newTotalStakeUpdate.updateBlockNumber = uint32(block.number);
        // loop through the operator's quorums and remove the operator's stake for each quorum
        for (uint8 quorumNumbersIndex = 0; quorumNumbersIndex < quorumNumbers.length;) {
            uint8 quorumNumber = uint8(quorumNumbers[quorumNumbersIndex]);
            // update the operator's stake
            uint96 stakeBeforeUpdate = _recordOperatorStakeUpdate(operatorId, quorumNumber, _operatorStakeUpdate);
            // subtract the amounts staked by the operator that is getting deregistered from the total stake before deregistration
            // copy latest totalStakes to memory
            _newTotalStakeUpdate.stake = _totalStakeHistory[quorumNumber][_totalStakeHistory[quorumNumber].length - 1].stake - stakeBeforeUpdate;
            // update storage of total stake
            _recordTotalStakeUpdate(quorumNumber, _newTotalStakeUpdate);

            emit StakeUpdate(
                operatorId,
                quorumNumber,
                // new stakes are zero
                0
            );
            unchecked {
                ++quorumNumbersIndex;
            }
        }
    }

    /**
     * @notice Finds the updated stake for `operator`, stores it and records the update.
     * @dev **DOES NOT UPDATE `totalStake` IN ANY WAY** -- `totalStake` updates must be done elsewhere.
     */
    function _updateOperatorStake(address operator, bytes32 operatorId, uint8 quorumNumber)
        internal
        returns (uint96, uint96)
    {   
        // determine new stakes
        OperatorStakeUpdate memory operatorStakeUpdate;
        operatorStakeUpdate.updateBlockNumber = uint32(block.number);
        operatorStakeUpdate.stake = weightOfOperator(quorumNumber, operator);

        // check if minimum requirements have been met
        if (operatorStakeUpdate.stake < minimumStakeForQuorum[quorumNumber]) {
            // set staker to 0
            operatorStakeUpdate.stake = uint96(0);
        }
        // get stakeBeforeUpdate and update with new stake
        uint96 stakeBeforeUpdate = _recordOperatorStakeUpdate(operatorId, quorumNumber, operatorStakeUpdate);
    
        emit StakeUpdate(
            operatorId,
            quorumNumber,
            operatorStakeUpdate.stake
        );

        return (stakeBeforeUpdate, operatorStakeUpdate.stake);
    }

    /// @notice Records that `operatorId`'s current stake is now param @operatorStakeUpdate
    function _recordOperatorStakeUpdate(bytes32 operatorId, uint8 quorumNumber, OperatorStakeUpdate memory operatorStakeUpdate) internal returns(uint96) {
        // initialize stakeBeforeUpdate to 0
        uint96 stakeBeforeUpdate;
        uint256 operatorStakeHistoryLength = operatorIdToStakeHistory[operatorId][quorumNumber].length; 
        if (operatorStakeHistoryLength != 0) {
            // set nextUpdateBlockNumber in prev stakes
            operatorIdToStakeHistory[operatorId][quorumNumber][operatorStakeHistoryLength - 1].nextUpdateBlockNumber =
                uint32(block.number);
            // load stake before update into memory if it exists
            stakeBeforeUpdate = operatorIdToStakeHistory[operatorId][quorumNumber][operatorStakeHistoryLength - 1].stake;
        }
        // push new stake to storage
        operatorIdToStakeHistory[operatorId][quorumNumber].push(operatorStakeUpdate);
        return stakeBeforeUpdate;
    }

    /// @notice Records that the `totalStake` is now equal to the input param @_totalStake
    function _recordTotalStakeUpdate(uint8 quorumNumber, OperatorStakeUpdate memory _totalStake) internal {
        uint256 _totalStakeHistoryLength = _totalStakeHistory[quorumNumber].length;
        if (_totalStakeHistoryLength != 0) {
            _totalStakeHistory[quorumNumber][_totalStakeHistoryLength - 1].nextUpdateBlockNumber = uint32(block.number);
        }
        _totalStake.updateBlockNumber = uint32(block.number);
        _totalStakeHistory[quorumNumber].push(_totalStake);
    }

    /// @notice Validates that the `operatorStake` was accurate at the given `blockNumber`
    function _validateOperatorStakeUpdateAtBlockNumber(OperatorStakeUpdate memory operatorStakeUpdate, uint32 blockNumber) internal pure {
        require(
            operatorStakeUpdate.updateBlockNumber <= blockNumber,
            "StakeRegistry._validateOperatorStakeAtBlockNumber: operatorStakeUpdate is from after blockNumber"
        );
        require(
            operatorStakeUpdate.nextUpdateBlockNumber == 0 || operatorStakeUpdate.nextUpdateBlockNumber > blockNumber,
            "StakeRegistry._validateOperatorStakeAtBlockNumber: there is a newer operatorStakeUpdate available before blockNumber"
        );
    }

    // storage gap
    uint256[50] private __GAP;
}