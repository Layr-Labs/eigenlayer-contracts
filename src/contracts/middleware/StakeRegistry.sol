// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/utils/math/Math.sol";
import "../interfaces/IServiceManager.sol";
import "../interfaces/IQuorumRegistry.sol";
import "./VoteWeigherBase.sol";

/**
 * @title An abstract Registry-type contract that is signature scheme agnostic.
 * @author Layr Labs, Inc.
 * @notice This contract is used for
 * - registering new operators
 * - committing to and finalizing de-registration as an operator
 * - updating the stakes of the operator
 * @dev This contract is missing key functions. See `BLSRegistry` or `ECDSARegistry` for examples that inherit from this contract.
 */
abstract contract RegistryBase is VoteWeigherBase, IQuorumRegistry {

    // TODO: set these on initialization
    /// @notice In order to register, an operator must have at least `minimumStakeFirstQuorum` or `minimumStakeSecondQuorum`, as
    /// evaluated by this contract's 'VoteWeigher' logic.
    uint96[256] public minimumStakeForQuorum;

    /// @notice array of the history of the total stakes for each quorum -- marked as internal since getTotalStakeFromIndex is a getter for this
    OperatorStakeUpdate[][256] internal totalStakeHistory;

    /// @notice mapping from operator's operatorId to the history of their stake updates
    mapping(bytes32 => mapping(uint8 => OperatorStakeUpdate[])) public operatorIdToStakeHistory;

    /// @notice mapping from operator's operatorId to the history of their index in the array of all operators
    mapping(bytes32 => OperatorIndex[]) public operatorIdToIndexHistory;

    // EVENTS
    /// @notice emitted whenever the stake of `operator` is updated
    event StakeUpdate(
        address operator,
        uint8 quorumNumber,
        uint96 stake,
        uint32 updateBlockNumber,
        uint32 prevUpdateBlockNumber
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
    function getTotalStakeUpdateForQuorumFromIndex(uint256 quorumNumber, uint256 index) external view returns (OperatorStakeUpdate memory) {
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
        OperatorStakeUpdate memory operatorStakeUpdate = getMostRecentStakeUpdateByOperator(operatorId, quorumNumber);
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

    function getLengthOfOperatorIdIndexHistory(bytes32 operatorId) external view returns (uint256) {
        return operatorIdToIndexHistory[operatorId].length;
    }

    function getLengthOfTotalStakeHistoryForQuorum(uint8 quorumNumber) external view returns (uint256) {
        return totalStakeHistory[quorumNumber].length;
    }

    function getLengthOfTotalOperatorsHistory() external view returns (uint256) {
        return totalOperatorsHistory.length;
    }

    /// @notice Returns task number from when `operator` has been registered.
    function getFromTaskNumberForOperator(address operator) external view returns (uint32) {
        return registry[operator].fromTaskNumber;
    }

    /// @notice Returns the current number of operators of this service.
    function numOperators() public view returns (uint32) {
        return uint32(operatorList.length);
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
            operatorStakeUpdate.stake != 0 // this implicitly checks that the quorum bitmap was 1 for the quorum of interest
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

    function updateSocket(string calldata newSocket) external {
        require(
            registry[msg.sender].status == IQuorumRegistry.Status.ACTIVE,
            "RegistryBase.updateSocket: Can only update socket if active on the service"
        );
        emit SocketUpdate(msg.sender, newSocket);
    }


    // INTERNAL FUNCTIONS
    /**
     * @notice Called when the total number of operators has changed.
     * Sets the `toBlockNumber` field on the last entry *so far* in thedynamic array `totalOperatorsHistory` to the current `block.number`,
     * recording that the previous entry is *no longer the latest* and the block number at which the next was added.
     * Pushes a new entry to `totalOperatorsHistory`, with `index` field set equal to the new amount of operators, recording the new number
     * of total operators (and leaving the `toBlockNumber` field at zero, signaling that this is the latest entry in the array)
     */
    function _updateTotalOperatorsHistory() internal {
        // set the 'toBlockNumber' field on the last entry *so far* in 'totalOperatorsHistory' to the current block number
        totalOperatorsHistory[totalOperatorsHistory.length - 1].toBlockNumber = uint32(block.number);
        // push a new entry to 'totalOperatorsHistory', with 'index' field set equal to the new amount of operators
        OperatorIndex memory _totalOperators;
        _totalOperators.index = uint32(operatorList.length);
        totalOperatorsHistory.push(_totalOperators);
    }

    /**
     * @notice Remove the operator from active status. Removes the operator with the given `operatorId` from the `index` in `operatorList`,
     * updates operatorList and index histories, and performs other necessary updates for removing operator
     */
    function _removeOperator(address operator, bytes32 operatorId, uint32 index) internal virtual {
        // remove the operator's stake
        _removeOperatorStake(operatorId);

        // store blockNumber at which operator index changed (stopped being applicable)
        operatorIdToIndexHistory[operatorId][operatorIdToIndexHistory[operatorId].length - 1].toBlockNumber =
            uint32(block.number);

        // remove the operator at `index` from the `operatorList`
        address swappedOperator = _removeOperatorFromOperatorListAndIndex(index);

        // @notice Registrant must continue to serve until the latest block at which an active task expires. this info is used in challenges
        uint32 latestServeUntilBlock = serviceManager.latestServeUntilBlock();
        // committing to not signing off on any more middleware tasks
        registry[operator].status = IQuorumRegistry.Status.INACTIVE;

        // record a stake update unbonding the operator after `latestServeUntilBlock`
        serviceManager.recordLastStakeUpdateAndRevokeSlashingAbility(operator, latestServeUntilBlock);

        // Emit `Deregistration` event
        emit Deregistration(operator, swappedOperator, index);
    }

    /**
     * @notice Removes the stakes of the operator
     */
    function _removeOperatorStake(bytes32 operatorId) internal {
        // loop through the operator's quorum bitmap and remove the operator's stake for each quorum
        uint256 quorumBitmap = operatorIdToQuorumBitmap[operatorId];
        for (uint quorumNumber = 0; quorumNumber < quorumCount;) {
            if (quorumBitmap >> quorumNumber & 1 == 1) {
                _removeOperatorStakeForQuorum(operatorId, uint8(quorumNumber));
            }
            unchecked {
                quorumNumber++;
            }
        }

    }

    /**
     * @notice Removes the stakes of the operator with operatorId `operatorId` for the quorum with number `quorumNumber`
     */
    function _removeOperatorStakeForQuorum(bytes32 operatorId, uint8 quorumNumber) internal {
        // gas saving by caching length here
        uint256 operatorIdToStakeHistoryLengthMinusOne = operatorIdToStakeHistory[operatorId][quorumNumber].length - 1;

        // determine current stakes
        OperatorStakeUpdate memory currentStakeUpdate =
            operatorIdToStakeHistory[operatorId][quorumNumber][operatorIdToStakeHistoryLengthMinusOne];
        //set nextUpdateBlockNumber in current stakes
        operatorIdToStakeHistory[operatorId][quorumNumber][operatorIdToStakeHistoryLengthMinusOne].nextUpdateBlockNumber =
            uint32(block.number);

        /**
         * @notice recording the information pertaining to change in stake for this operator in the history. operator stakes are set to 0 here.
         */
        operatorIdToStakeHistory[operatorId][quorumNumber].push(
            OperatorStakeUpdate({
                // recording the current block number where the operator stake got updated
                updateBlockNumber: uint32(block.number),
                // mark as 0 since the next update has not yet occurred
                nextUpdateBlockNumber: 0,
                // setting the operator's stake to 0
                stake: 0
            })
        );

        // subtract the amounts staked by the operator that is getting deregistered from the total stake
        // copy latest totalStakes to memory
        OperatorStakeUpdate memory currentTotalStakeUpdate = totalStakeHistory[quorumNumber][totalStakeHistory.length - 1];
        currentTotalStakeUpdate.stake -= currentStakeUpdate.stake;
        // update storage of total stake
        _recordTotalStakeUpdate(quorumNumber, currentTotalStakeUpdate);

        emit StakeUpdate(
            msg.sender,
            // new stakes are zero
            quorumNumber,
            0,
            uint32(block.number),
            currentStakeUpdate.updateBlockNumber
            );
    }

    /// @notice Adds the Operator `operator` with the given `pubkeyHash` to the `operatorList` and performs necessary related updates.
    function _addRegistrant(
        address operator,
        bytes32 pubkeyHash,
        uint256 quorumBitmap
    )
        internal virtual
    {

        require(
            slasher.contractCanSlashOperatorUntilBlock(operator, address(serviceManager)) == type(uint32).max,
            "RegistryBase._addRegistrant: operator must be opted into slashing by the serviceManager"
        );
        // store the Operator's info in mapping
        registry[operator] = Operator({
            pubkeyHash: pubkeyHash,
            status: IQuorumRegistry.Status.ACTIVE,
            fromTaskNumber: serviceManager.taskNumber()
        });

        // store the operator's quorum bitmap
        pubkeyHashToQuorumBitmap[pubkeyHash] = quorumBitmap;

        // add the operator to the list of operators
        operatorList.push(operator);
        
        // calculate stakes for each quorum the operator is trying to join
        _registerStake(operator, pubkeyHash, quorumBitmap);

        // record `operator`'s index in list of operators
        OperatorIndex memory operatorIndex;
        operatorIndex.index = uint32(operatorList.length - 1);
        pubkeyHashToIndexHistory[pubkeyHash].push(operatorIndex);

        // Update totalOperatorsHistory array
        _updateTotalOperatorsHistory();

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
    function _registerStake(address operator, bytes32 pubkeyHash, uint256 quorumBitmap)
        internal
    {
        // verify that the `operator` is not already registered
        require(
            registry[operator].status == IQuorumRegistry.Status.INACTIVE,
            "RegistryBase._registerStake: Operator is already registered"
        );

        OperatorStakeUpdate memory _operatorStakeUpdate;
        // add the `updateBlockNumber` info
        _operatorStakeUpdate.updateBlockNumber = uint32(block.number);
        OperatorStakeUpdate memory _newTotalStakeUpdate;
        // add the `updateBlockNumber` info
        _newTotalStakeUpdate.updateBlockNumber = uint32(block.number);
        // for each quorum, evaluate stake and add to total stake
        for (uint8 quorumNumber = 0; quorumNumber < quorumCount;) {
            // evaluate the stake for the operator
            if(quorumBitmap >> quorumNumber & 1 == 1) {
                _operatorStakeUpdate.stake = uint96(weightOfOperator(operator, quorumNumber));
                // check if minimum requirement has been met
                require(_operatorStakeUpdate.stake >= minimumStakeForQuorum[quorumNumber], "RegistryBase._registerStake: Operator does not meet minimum stake requirement for quorum");
                // check special case that operator is re-registering (and thus already has some history)
                if (pubkeyHashToStakeHistory[pubkeyHash][quorumNumber].length != 0) {
                    // correctly set the 'nextUpdateBlockNumber' field for the re-registering operator's oldest history entry
                    pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][pubkeyHashToStakeHistory[pubkeyHash][quorumNumber].length - 1].nextUpdateBlockNumber
                        = uint32(block.number);
                }
                // push the new stake for the operator to storage
                pubkeyHashToStakeHistory[pubkeyHash][quorumNumber].push(_operatorStakeUpdate);

                // get the total stake for the quorum
                _newTotalStakeUpdate = totalStakeHistory[quorumNumber][totalStakeHistory[quorumNumber].length - 1];
                // add operator stakes to total stake (in memory)
                _newTotalStakeUpdate.stake = uint96(_newTotalStakeUpdate.stake + _operatorStakeUpdate.stake);
                // update storage of total stake
                _recordTotalStakeUpdate(quorumNumber, _newTotalStakeUpdate);

                emit StakeUpdate(
                    operator,
                    quorumNumber,
                    _operatorStakeUpdate.stake,
                    uint32(block.number),
                    // no previous update block number -- use 0 instead
                    0 // TODO: Decide whether this needs to be set in re-registration edge case
                );
            }
            unchecked {
                ++quorumNumber;
            }
        }
    }

    /**
     * @notice Finds the updated stake for `operator`, stores it and records the update.
     * @dev **DOES NOT UPDATE `totalStake` IN ANY WAY** -- `totalStake` updates must be done elsewhere.
     */
    function _updateOperatorStake(address operator, bytes32 pubkeyHash, uint256 quorumBitmap, uint8 quorumNumber, uint32 prevUpdateBlockNumber)
        internal
        returns (OperatorStakeUpdate memory operatorStakeUpdate)
    {   
        // if the operator is part of the quorum
        if (quorumBitmap >> quorumNumber & 1 == 1) {
            // determine new stakes
            operatorStakeUpdate.updateBlockNumber = uint32(block.number);
            operatorStakeUpdate.stake = weightOfOperator(operator, quorumNumber);

            // check if minimum requirements have been met
            if (operatorStakeUpdate.stake < minimumStakeForQuorum[quorumNumber]) {
                operatorStakeUpdate.stake = uint96(0);
            }

            // set nextUpdateBlockNumber in prev stakes
            pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][pubkeyHashToStakeHistory[pubkeyHash][quorumNumber].length - 1].nextUpdateBlockNumber =
                uint32(block.number);
            // push new stake to storage
            pubkeyHashToStakeHistory[pubkeyHash][quorumNumber].push(operatorStakeUpdate);
            
            emit StakeUpdate(
                operator,
                quorumNumber,
                operatorStakeUpdate.stake,
                uint32(block.number),
                prevUpdateBlockNumber
            );
        }
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