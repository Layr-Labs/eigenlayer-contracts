// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/utils/math/Math.sol";
import "../interfaces/IServiceManager.sol";
import "../interfaces/IStakeRegistry.sol";
import "../interfaces/IRegistryCoordinator.sol";
import "./VoteWeigherBase.sol";

/**
 * @title Interface for a `Registry` that keeps track of stakes of operators for up to 256 quroums.
 * @author Layr Labs, Inc.
 */
contract StakeRegistry is VoteWeigherBase, IStakeRegistry {

    IRegistryCoordinator public registryCoordinator;

    // TODO: set these on initialization
    /// @notice In order to register, an operator must have at least `minimumStakeFirstQuorum` or `minimumStakeSecondQuorum`, as
    /// evaluated by this contract's 'VoteWeigher' logic.
    uint96[256] public minimumStakeForQuorum;

    /// @notice array of the history of the total stakes for each quorum -- marked as internal since getTotalStakeFromIndex is a getter for this
    OperatorStakeUpdate[][256] internal totalStakeHistory;

    /// @notice mapping from operator's operatorId to the history of their stake updates
    mapping(bytes32 => mapping(uint8 => OperatorStakeUpdate[])) public operatorIdToStakeHistory;

    // EVENTS
    /// @notice emitted whenever the stake of `operator` is updated
    event StakeUpdate(
        bytes32 operatorId,
        uint8 quorumNumber,
        uint96 stake
    );

    /**
     * @notice Irrevocably sets the (immutable) `delegation` & `strategyManager` addresses, and `NUMBER_OF_QUORUMS` variable.
     */
    constructor(
        IStrategyManager _strategyManager,
        IServiceManager _serviceManager
    ) VoteWeigherBase(_strategyManager, _serviceManager)
    // solhint-disable-next-line no-empty-blocks
    {
    }

    /**
     * @notice Adds empty first entries to the dynamic arrays `totalStakeHistory` and `totalOperatorsHistory`,
     * to record an initial condition of zero operators with zero total stake.
     * Adds `_quorumStrategiesConsideredAndMultipliers` for each quorum the Registry is being initialized with
     */
    function _initialize(
        uint96[] memory _minimumStakeForQuorum,
        StrategyAndWeightingMultiplier[][] memory _quorumStrategiesConsideredAndMultipliers
    ) internal virtual onlyInitializing {
        // sanity check lengths
        require(_minimumStakeForQuorum.length == _quorumStrategiesConsideredAndMultipliers.length, "Registry._initialize: minimumStakeForQuorum length mismatch");
        // push an empty OperatorStakeUpdate struct to the total stake history to record starting with zero stake
        // TODO: Address this @ gpsanant
        OperatorStakeUpdate memory _totalStakeUpdate;
        for (uint quorumNumber = 0; quorumNumber < 256;) {
            totalStakeHistory[quorumNumber].push(_totalStakeUpdate);
            unchecked {
                ++quorumNumber;
            }
        }

        // add the strategies considered and multipliers for each quorum
        for (uint8 quorumNumber = 0; quorumNumber < _quorumStrategiesConsideredAndMultipliers.length;) {
            minimumStakeForQuorum[quorumNumber] = _minimumStakeForQuorum[quorumNumber];
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
     * @notice Returns the `index`-th entry in the dynamic array of total stake, `totalStakeHistory` for quorum `quorumNumber`.
     * @param quorumNumber The quorum number to get the stake for.
     * @param index Array index for lookup, within the dynamic array `totalStakeHistory[quorumNumber]`.
     */
    function getTotalStakeUpdateForQuorumFromIndex(uint8 quorumNumber, uint256 index) external view returns (OperatorStakeUpdate memory) {
        return totalStakeHistory[quorumNumber][index];
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
     * `totalStakeHistory[quorumNumber]` array if it was the stake at `blockNumber`. Reverts otherwise.
     * @param quorumNumber The quorum number to get the stake for.
     * @param index Array index for lookup, within the dynamic array `totalStakeHistory[quorumNumber]`.
     * @param blockNumber Block number to make sure the stake is from.
     * @dev Function will revert if `index` is out-of-bounds.
     */
    function getTotalStakeAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) external view returns (uint96) {
        OperatorStakeUpdate memory totalStakeUpdate = totalStakeHistory[quorumNumber][index];
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

    /// @notice Returns the stake weight from the latest entry in `totalStakeHistory` for quorum `quorumNumber`.
    function getCurrentTotalStakeForQuorum(uint8 quorumNumber) external view returns (uint96) {
        // no chance of underflow / error in next line, since an empty entry is pushed in the constructor
        return totalStakeHistory[quorumNumber][totalStakeHistory[quorumNumber].length - 1].stake;
    }

    function getLengthOfOperatorIdStakeHistoryForQuorum(bytes32 operatorId, uint8 quorumNumber) external view returns (uint256) {
        return operatorIdToStakeHistory[operatorId][quorumNumber].length;
    }

    function getLengthOfTotalStakeHistoryForQuorum(uint8 quorumNumber) external view returns (uint256) {
        return totalStakeHistory[quorumNumber].length;
    }

    /**
     * @notice Checks that the `operator` was active at the `blockNumber`, using the specified `stakeHistoryIndex` as proof.
     * @param operatorId is the id of the operator of interest
     * @param blockNumber is the block number of interest
     * @param quorumNumber is the quorum number which the operator had stake in
     * @param stakeHistoryIndex specifies an index in `operatorIdToStakeHistory[operatorId]`
     * @return 'true' if it is succesfully proven that  the `operator` was active at the `blockNumber`, and 'false' otherwise
     * @dev In order for this function to return 'true', the inputs must satisfy all of the following list:
     * 1) `operatorIdToStakeHistory[operatorId][quorumNumber][index].updateBlockNumber <= blockNumber`
     * 2) `operatorIdToStakeHistory[operatorId][quorumNumber][index].nextUpdateBlockNumber` must be either `0` (signifying no next update) or
     * is must be strictly greater than `blockNumber`
     * 3) `operatorIdToStakeHistory[operatorId][quorumNumber][index].stake > 0` i.e. the operator had nonzero stake
     * @dev Note that a return value of 'false' does not guarantee that the `operator` was inactive at `blockNumber`, since a
     * bad `stakeHistoryIndex` can be supplied in order to obtain a response of 'false'.
     */
    function checkOperatorActiveAtBlockNumber(
        bytes32 operatorId,
        uint256 blockNumber,
        uint8 quorumNumber,
        uint256 stakeHistoryIndex
        ) external view returns (bool)
    {
        // pull the stake history entry specified by `stakeHistoryIndex`
        OperatorStakeUpdate memory operatorStakeUpdate = operatorIdToStakeHistory[operatorId][quorumNumber][stakeHistoryIndex];
        return (
            // check that the update specified by `stakeHistoryIndex` occurred at or prior to `blockNumber`
            (operatorStakeUpdate.updateBlockNumber <= blockNumber)
            &&
            // if there is a next update, then check that the next update occurred strictly after `blockNumber`
            (operatorStakeUpdate.nextUpdateBlockNumber == 0 || operatorStakeUpdate.nextUpdateBlockNumber > blockNumber)
            &&
            /// verify that the stake was non-zero at the time (note: here was use the assumption that the operator was 'inactive'
            /// once their stake fell to zero)
            operatorStakeUpdate.stake != 0 // this implicitly checks that the operator was a part of the quorum of interest
        );
    }

    /**
     * @notice Checks that the `operator` was inactive at the `blockNumber`, using the specified `stakeHistoryIndex` for `quorumNumber` as proof.
     * @param operatorId is the id of the operator of interest
     * @param blockNumber is the block number of interest
     * @param quorumNumber is the quorum number which the operator had no stake in
     * @param stakeHistoryIndex specifies an index in `operatorIdToStakeHistory[operatorId]`
     * @return 'true' if it is succesfully proven that  the `operator` was inactive at the `blockNumber`, and 'false' otherwise
     * @dev In order for this function to return 'true', the inputs must satisfy all of the following list:
     * 1) `operatorIdToStakeHistory[operatorId][quorumNumber][index].updateBlockNumber <= blockNumber`
     * 2) `operatorIdToStakeHistory[operatorId][quorumNumber][index].nextUpdateBlockNumber` must be either `0` (signifying no next update) or
     * is must be strictly greater than `blockNumber`
     * 3) `operatorIdToStakeHistory[operatorId][quorumNumber][index].stake == 0` i.e. the operator had zero stake
     * @dev Note that a return value of 'false' does not guarantee that the `operator` was active at `blockNumber`, since a
     * bad `stakeHistoryIndex` can be supplied in order to obtain a response of 'false'.
     * @dev One precondition that must be checked is that the operator is a part of the given `quorumNumber`
     */
    function checkOperatorInactiveAtBlockNumber(
        bytes32 operatorId,
        uint256 blockNumber,
        uint8 quorumNumber,
        uint256 stakeHistoryIndex
        ) external view returns (bool)
    {
        // special case for `operatorIdToStakeHistory[operatorId]` having lenght zero -- in which case we know the operator was never registered
        if (operatorIdToStakeHistory[operatorId][quorumNumber].length == 0) {
            return true;
        }
        // pull the stake history entry specified by `stakeHistoryIndex`
        OperatorStakeUpdate memory operatorStakeUpdate = operatorIdToStakeHistory[operatorId][quorumNumber][stakeHistoryIndex];
        return (
            // check that the update specified by `stakeHistoryIndex` occurred at or prior to `blockNumber`
            (operatorStakeUpdate.updateBlockNumber <= blockNumber)
            &&
            // if there is a next update, then check that the next update occurred strictly after `blockNumber`
            (operatorStakeUpdate.nextUpdateBlockNumber == 0 || operatorStakeUpdate.nextUpdateBlockNumber > blockNumber)
            &&
            /// verify that the stake was zero at the time (note: here was use the assumption that the operator was 'inactive'
            /// once their stake fell to zero)
            operatorStakeUpdate.stake == 0
        );
    }

    // MUTATING FUNCTIONS

    /// @notice Adjusts the `minimumStakeFirstQuorum` -- i.e. the node stake (weight) requirement for inclusion in the 1st quorum.
    function setMinimumStakeForQuorum(uint8 quorumNumber, uint96 minimumStake) external onlyServiceManagerOwner {
        minimumStakeForQuorum[quorumNumber] = minimumStake;
    }

    /**
     * @notice Registers the `operator` with `operatorId` for the specified `quorumNumbers`.
     * @param operator The address of the operator to register.
     * @param operatorId The id of the operator to register.
     * @param quorumNumbers The quorum numbers the operator is registering for, where each byte is an 8 bit integer quorumNumber.
     * @dev access restricted to the RegistryCoordinator
     * @dev Preconditions:
     *         1) `quorumNumbers` has no duplicates
     *         2) `quorumNumbers.length` != 0
     *         3) `quorumNumbers` is ordered in ascending order
     *         4) the operator is not already registered
     */
    function registerOperator(address operator, bytes32 operatorId, bytes memory quorumNumbers) external virtual {
        _registerOperator(operator, operatorId, quorumNumbers);
    }

    /**
     * @notice Deregisters the operator with `operatorId` for the specified `quorumNumbers`.
     * @param operator The address of the operator to deregister.
     * @param operatorId The id of the operator to deregister.
     * @param quorumNumbers The quourm numbers the operator is deregistering from, where each byte is an 8 bit integer quorumNumber.
     * @dev access restricted to the RegistryCoordinator
     * @dev Preconditions:
     *         1) `quorumNumbers` has no duplicates
     *         2) `quorumNumbers.length` != 0
     *         3) `quorumNumbers` is ordered in ascending order
     *         4) the operator is not already deregistered
     *         5) `quorumNumbers` is the same as the parameter use when registering
     */
    function deregisterOperator(address operator, bytes32 operatorId, bytes memory quorumNumbers) external virtual {
        _deregisterOperator(operator, operatorId, quorumNumbers);
    }

    /**
     * @notice Used for updating information on deposits of nodes.
     * @param operators are the addresses of the operators whose stake information is getting updated
     * @param operatorIds are the ids of the operators whose stake information is getting updated
     * @param quorumBitmaps are the bitmap of the quorums that each operator in `operators` is part of
     * @param prevElements are the elements before this middleware in the operator's linked list within the slasher
     * @dev Precondition:
     *          1) `quorumBitmaps[i]` should be the bitmap that represents the quorums that `operators[i]` registered for
     */
    function updateStakes(address[] memory operators, bytes32[] memory operatorIds, uint256[] memory quorumBitmaps, uint256[] memory prevElements) external {
        // for each quorum, loop through operators and see if they are apart of the quorum
        // if they are, get their new weight and update their individual stake history and thes
        // quorum's total stake history accordingly
        for (uint8 quorumNumber = 0; quorumNumber < quorumCount;) {
            OperatorStakeUpdate memory totalStakeUpdate;
            // for each operator
            for(uint i = 0; i < operatorIds.length;) {
                // if the operator is apart of the quorum
                if (quorumBitmaps[i] >> quorumNumber & 1 == 1) {
                    // if the total stake has not been loaded yet, load it
                    if (totalStakeUpdate.updateBlockNumber == 0) {
                        totalStakeUpdate = totalStakeHistory[quorumNumber][totalStakeHistory[quorumNumber].length - 1];
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

    function _registerOperator(address operator, bytes32 operatorId, bytes memory quorumNumbers) internal {
        require(
            slasher.contractCanSlashOperatorUntilBlock(operator, address(serviceManager)) == type(uint32).max,
            "RegistryBase._addRegistrant: operator must be opted into slashing by the serviceManager"
        );
        
        // calculate stakes for each quorum the operator is trying to join
        _registerStake(operator, operatorId, quorumNumbers);

        // record a stake update not bonding the operator at all (unbonded at 0), because they haven't served anything yet
        serviceManager.recordFirstStakeUpdate(operator, 0);
    }

    /**
     * TODO: critique: "Currently only `_registrationStakeEvaluation` uses the `uint256 registrantType` input -- we should **EITHER** store this
     * and keep using it in other places as well, **OR** stop using it altogether"
     */
    /**
     * @notice Used inside of inheriting contracts to validate the registration of `operator` and find their `OperatorStake`.
     * @dev This function does **not** update the stored state of the operator's stakes -- storage updates are performed elsewhere.
     */
    function _registerStake(address operator, bytes32 operatorId, bytes memory quorumNumbers)
        internal
    {
        uint8 quorumNumbersLength = uint8(quorumNumbers.length);
        // check the operator is registering for only valid quorums
        require(uint8(quorumNumbers[quorumNumbersLength - 1]) < quorumCount, "StakeRegistry._registerStake: greatest quorumNumber must be less than quorumCount");
        OperatorStakeUpdate memory _newTotalStakeUpdate;
        // add the `updateBlockNumber` info
        _newTotalStakeUpdate.updateBlockNumber = uint32(block.number);
        // for each quorum, evaluate stake and add to total stake
        for (uint8 quorumNumbersIndex = 0; quorumNumbersIndex < quorumNumbersLength;) {
            // get the next quourumNumber
            uint8 quorumNumber = uint8(quorumNumbers[quorumNumbersIndex]);
            // evaluate the stake for the operator
            // since we don't use the first output, this will use 1 extra sload when deregistered operator's register again
            (, uint96 stake) = _updateOperatorStake(operator, operatorId, quorumNumber);
            // @JEFF: This reverts pretty late, but i think that's fine. wdyt?
            // check if minimum requirement has been met, will be 0 if not
            require(stake != 0, "StakeRegistry._registerStake: Operator does not meet minimum stake requirement for quorum");
            // add operator stakes to total stake before update (in memory)
            _newTotalStakeUpdate.stake = totalStakeHistory[quorumNumber][totalStakeHistory[quorumNumber].length - 1].stake + stake;
            // update storage of total stake
            _recordTotalStakeUpdate(quorumNumber, _newTotalStakeUpdate);
            unchecked {
                ++quorumNumbersIndex;
            }
        }
    }

    function _deregisterOperator(address operator, bytes32 operatorId, bytes memory quorumNumbers) internal {
        // remove the operator's stake
        _removeOperatorStake(operatorId, quorumNumbers);

        // @notice Registrant must continue to serve until the latest block at which an active task expires. this info is used in challenges
        uint32 latestServeUntilBlock = serviceManager.latestServeUntilBlock();

        // record a stake update unbonding the operator after `latestServeUntilBlock`
        serviceManager.recordLastStakeUpdateAndRevokeSlashingAbility(operator, latestServeUntilBlock);
    }

    /**
     * @notice Removes the stakes of the operator
     */
    function _removeOperatorStake(bytes32 operatorId, bytes memory quorumNumbers) internal {
        uint8 quorumNumbersLength = uint8(quorumNumbers.length);
        // check the operator is deregistering from only valid quorums
        require(uint8(quorumNumbers[quorumNumbersLength - 1]) < quorumCount, "StakeRegistry._registerStake: greatest quorumNumber must be less than quorumCount");
        OperatorStakeUpdate memory _operatorStakeUpdate;
        // add the `updateBlockNumber` info
        _operatorStakeUpdate.updateBlockNumber = uint32(block.number);
        OperatorStakeUpdate memory _newTotalStakeUpdate;
        // add the `updateBlockNumber` info
        _newTotalStakeUpdate.updateBlockNumber = uint32(block.number);
        // loop through the operator's quorums and remove the operator's stake for each quorum
        for (uint8 quorumNumbersIndex = 0; quorumNumbersIndex < quorumNumbersLength;) {
            uint8 quorumNumber = uint8(quorumNumbers[quorumNumbersIndex]);
            // update the operator's stake
            uint96 stakeBeforeUpdate = _recordOperatorStakeUpdate(operatorId, quorumNumber, _operatorStakeUpdate);
            // subtract the amounts staked by the operator that is getting deregistered from the total stake before deregistration
            // copy latest totalStakes to memory
            _newTotalStakeUpdate.stake = totalStakeHistory[quorumNumber][totalStakeHistory.length - 1].stake - stakeBeforeUpdate;
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
            operatorStakeUpdate.stake = uint96(0);
        }
        // initialize stakeBeforeUpdate to 0
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
        _totalStake.updateBlockNumber = uint32(block.number);
        totalStakeHistory[quorumNumber][totalStakeHistory.length - 1].nextUpdateBlockNumber = uint32(block.number);
        totalStakeHistory[quorumNumber].push(_totalStake);
    }

    /// @notice Validates that the `operatorStake` was accurate at the given `blockNumber`
    function _validateOperatorStakeUpdateAtBlockNumber(OperatorStakeUpdate memory operatorStakeUpdate, uint32 blockNumber) internal pure {
        require(
            operatorStakeUpdate.updateBlockNumber <= blockNumber,
            "RegistryBase._validateOperatorStakeAtBlockNumber: operatorStakeUpdate is from after blockNumber"
        );
        require(
            operatorStakeUpdate.nextUpdateBlockNumber == 0 || operatorStakeUpdate.nextUpdateBlockNumber > blockNumber,
            "RegistryBase._validateOperatorStakeAtBlockNumber: there is a newer operatorStakeUpdate available before blockNumber"
        );
    }

    // storage gap
    uint256[50] private __GAP;
}