// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/utils/math/Math.sol";
import "../interfaces/IServiceManager.sol";
import "../interfaces/IQuorumRegistry.sol";
import "./VoteWeigherBase.sol";

/**
 * @title An abstract Registry-type contract that is signature scheme agnostic.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
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

    /// @notice used for storing Operator info on each operator while registration
    mapping(address => Operator) public registry;

    /// @notice used for storing the list of current and past registered operators
    address[] public operatorList;

    /// @notice used for storing the quorums which the operator is participating in
    mapping(bytes32 => uint256) public pubkeyHashToQuorumBitmap;

    /// @notice array of the history of the total stakes for each quorum -- marked as internal since getTotalStakeFromIndex is a getter for this
    OperatorStakeUpdate[][256] internal totalStakeHistory;

    /// @notice array of the history of the number of operators, and the taskNumbers at which the number of operators changed
    OperatorIndex[] public totalOperatorsHistory;

    /// @notice mapping from operator's pubkeyhash to the history of their stake updates
    mapping(bytes32 => mapping(uint8 => OperatorStakeUpdate[])) public pubkeyHashToStakeHistory;

    /// @notice mapping from operator's pubkeyhash to the history of their index in the array of all operators
    mapping(bytes32 => OperatorIndex[]) public pubkeyHashToIndexHistory;

    // EVENTS
    /// @notice emitted when `operator` updates their socket address to `socket`
    event SocketUpdate(address operator, string socket);

    /// @notice emitted whenever the stake of `operator` is updated
    event StakeUpdate(
        address operator,
        uint8 quorumNumber,
        uint96 stake,
        uint32 updateBlockNumber,
        uint32 prevUpdateBlockNumber
    );

    /**
     * @notice Emitted whenever an operator deregisters.
     * The `swapped` address is the address returned by an internal call to the `_popRegistrant` function.
     */
    event Deregistration(address operator, address swapped, uint32 deregisteredIndex);

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

        // push an empty OperatorIndex struct to the total operators history to record starting with zero operators
        OperatorIndex memory _totalOperators;
        totalOperatorsHistory.push(_totalOperators);

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
     * @notice Looks up the `operator`'s index in the dynamic array `operatorList` at the specified `blockNumber`.
     * @param index Used to specify the entry within the dynamic array `pubkeyHashToIndexHistory[pubkeyHash]` to 
     * read data from, where `pubkeyHash` is looked up from `operator`'s registration info
     * @param blockNumber Is the desired block number at which we wish to query the operator's position in the `operatorList` array
     * @dev Function will revert in the event that the specified `index` input does not identify the appropriate entry in the
     * array `pubkeyHashToIndexHistory[pubkeyHash]` to pull the info from.
    */
    function getOperatorIndex(address operator, uint32 blockNumber, uint32 index) external view returns (uint32) {
        // look up the operator's stored pubkeyHash
        bytes32 pubkeyHash = getOperatorPubkeyHash(operator);

        /**
         * Since the 'to' field represents the blockNumber at which a new index started, it is OK if the 
         * previous array entry has 'to' == blockNumber, so we check not strict inequality here
         */
        require(
            index == 0 || pubkeyHashToIndexHistory[pubkeyHash][index - 1].toBlockNumber <= blockNumber,
            "RegistryBase.getOperatorIndex: Operator indexHistory index is too high"
        );
        OperatorIndex memory operatorIndex = pubkeyHashToIndexHistory[pubkeyHash][index];
        /**
         * When deregistering, the operator does *not* serve the current block number -- 'to' gets set (from zero) to the current block number.
         * Since the 'to' field represents the blocknumber at which a new index started, we want to check strict inequality here.
        */
        require(
            operatorIndex.toBlockNumber == 0 || blockNumber < operatorIndex.toBlockNumber,
            "RegistryBase.getOperatorIndex: indexHistory index is too low"
        );
        return operatorIndex.index;
    }

    /**
     * @notice Looks up the number of total operators at the specified `blockNumber`.
     * @param index Input used to specify the entry within the dynamic array `totalOperatorsHistory` to read data from.
     * @dev This function will revert if the provided `index` is out of bounds.
    */
    function getTotalOperators(uint32 blockNumber, uint32 index) external view returns (uint32) {
        /**
         * Since the 'to' field represents the blockNumber at which a new index started, it is OK if the 
         * previous array entry has 'to' == blockNumber, so we check not strict inequality here
         */
        require(
            index == 0 || totalOperatorsHistory[index - 1].toBlockNumber <= blockNumber,
            "RegistryBase.getTotalOperators: TotalOperatorsHistory index is too high"
        );

    
        OperatorIndex memory operatorIndex = totalOperatorsHistory[index];

        // since the 'to' field represents the blockNumber at which a new index started, we want to check strict inequality here

        require(
            operatorIndex.toBlockNumber == 0 || blockNumber < operatorIndex.toBlockNumber,
            "RegistryBase.getTotalOperators: TotalOperatorsHistory index is too low"
        );
        return operatorIndex.index;
    }

    /// @notice Returns whether or not the `operator` is currently an active operator, i.e. is "registered".
    function isActiveOperator(address operator) external view virtual returns (bool) {
        return (registry[operator].status == IQuorumRegistry.Status.ACTIVE);
    }

    /// @notice Returns the stored pubkeyHash for the specified `operator`.
    function getOperatorPubkeyHash(address operator) public view returns (bytes32) {
        return registry[operator].pubkeyHash;
    }

    /**
     * @notice Returns the `index`-th entry in the `pubkeyHashToStakeHistory[pubkeyHash][quorumNumber]` array.
     * @param quorumNumber The quorum number to get the stake for.
     * @param pubkeyHash Hash of the public key of the operator of interest.
     * @param index Array index for lookup, within the dynamic array `pubkeyHashToStakeHistory[pubkeyHash][quorumNumber]`.
     * @dev Function will revert if `index` is out-of-bounds.
     */
    function getStakeUpdateForQuorumFromPubkeyHashAndIndex(uint8 quorumNumber, bytes32 pubkeyHash, uint256 index)
        external
        view
        returns (OperatorStakeUpdate memory)
    {
        return pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][index];
    }

    /**
     * @notice Returns the `index`-th entry in the dynamic array of total stake, `totalStakeHistory` for quorum `quorumNumber`.
     * @dev Function will revert in the event that `index` is out-of-bounds.
     */
    function getTotalStakeUpdateForQuorumFromIndex(uint256 quorumNumber, uint256 index) external view returns (OperatorStakeUpdate memory) {
        return totalStakeHistory[quorumNumber][index];
    }

    /**
     * @notice Returns the stake weight corresponding to `pubkeyHash` for quorum `quorumNumber` at `blockNumber`. Reverts otherwise.
     * @param quorumNumber The quorum number to get the stake for.
     * @param pubkeyHash Hash of the public key of the operator of interest.
     * @param index Array index for lookup, within the dynamic array `pubkeyHashToStakeHistory[pubkeyHash][quorumNumber]`.
     * @param blockNumber Block number to make sure the stake is from.
     * @dev Function will revert if `index` is out-of-bounds.
     */
    function getStakeForQuorumAtBlockNumberFromPubkeyHashAndIndex(uint8 quorumNumber, uint32 blockNumber, bytes32 pubkeyHash, uint256 index)
        external
        view
        returns (uint96)
    {
        OperatorStakeUpdate memory operatorStakeUpdate = pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][index];
        _validateOperatorStakeUpdateAtBlockNumber(operatorStakeUpdate, blockNumber);
        return operatorStakeUpdate.stake;
    }

    /**
     * @notice Returns the total stake weight for quorum `quorumNumber` at `blockNumber`. Reverts otherwise.
     * @param quorumNumber The quorum number to get the stake for.
     * @param index Array index for lookup, within the dynamic array `totalStakeHistory[quorumNumber]`.
     * @param blockNumber Block number to make sure the stake is from.
     * @dev Function will revert if `index` is out-of-bounds.
     */
    function getTotalStakeAtBlockNumberFromIndex(uint256 quorumNumber, uint32 blockNumber, uint256 index) external view returns (uint96) {
        OperatorStakeUpdate memory totalStakeUpdate = totalStakeHistory[quorumNumber][index];
        _validateOperatorStakeUpdateAtBlockNumber(totalStakeUpdate, blockNumber);
        return totalStakeUpdate.stake;
    }

    /**
     * @notice Returns the most recent stake weight for the `operator` for a certain quorum
     * @dev Function returns an OperatorStakeUpdate struct with **every entry equal to 0** in the event that the operator has no stake history
     */
    function getMostRecentStakeUpdateByOperator(address operator, uint8 quorumNumber) public view returns (OperatorStakeUpdate memory) {
        bytes32 pubkeyHash = getOperatorPubkeyHash(operator);
        uint256 historyLength = pubkeyHashToStakeHistory[pubkeyHash][quorumNumber].length;
        OperatorStakeUpdate memory operatorStakeUpdate;
        if (historyLength == 0) {
            return operatorStakeUpdate;
        } else {
            operatorStakeUpdate = pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][historyLength - 1];
            return operatorStakeUpdate;
        }
    }

    function getStakeHistoryLengthForQuorumNumber(bytes32 pubkeyHash, uint8 quorumNumber) external view returns (uint256) {
        return pubkeyHashToStakeHistory[pubkeyHash][quorumNumber].length;
    }

    /**
     * @notice Returns the most recent stake weight for the `operator` for quorum `quorumNumber`
     * @dev Function returns weight of **0** in the event that the operator has no stake history
     */
    function getCurrentOperatorStakeForQuorum(address operator, uint8 quorumNumber) external view returns (uint96) {
        OperatorStakeUpdate memory operatorStakeUpdate = getMostRecentStakeUpdateByOperator(operator, quorumNumber);
        return operatorStakeUpdate.stake;
    }

    /// @notice Returns the stake weight from the latest entry in `totalStakeHistory` for quorum `quorumNumber`.
    function getCurrentTotalStakeForQuorum(uint8 quorumNumber) external view returns (uint96) {
        // no chance of underflow / error in next line, since an empty entry is pushed in the constructor
        return totalStakeHistory[quorumNumber][totalStakeHistory[quorumNumber].length - 1].stake;
    }

    function getLengthOfPubkeyHashStakeHistoryForQuorum(bytes32 pubkeyHash, uint8 quorumNumber) external view returns (uint256) {
        return pubkeyHashToStakeHistory[pubkeyHash][quorumNumber].length;
    }

    function getLengthOfPubkeyHashIndexHistory(bytes32 pubkeyHash) external view returns (uint256) {
        return pubkeyHashToIndexHistory[pubkeyHash].length;
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
     * @param operator is the operator of interest
     * @param blockNumber is the block number of interest
     * @param quorumNumber is the quorum number which the operator had stake in
     * @param stakeHistoryIndex specifies an index in `pubkeyHashToStakeHistory[pubkeyHash]`, where `pubkeyHash` is looked up
     * in `registry[operator].pubkeyHash`
     * @return 'true' if it is succesfully proven that  the `operator` was active at the `blockNumber`, and 'false' otherwise
     * @dev In order for this function to return 'true', the inputs must satisfy all of the following list:
     * 1) `pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][index].updateBlockNumber <= blockNumber`
     * 2) `pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][index].nextUpdateBlockNumber` must be either `0` (signifying no next update) or
     * is must be strictly greater than `blockNumber`
     * 3) `pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][index].stake > 0` i.e. the operator had nonzero stake
     * @dev Note that a return value of 'false' does not guarantee that the `operator` was inactive at `blockNumber`, since a
     * bad `stakeHistoryIndex` can be supplied in order to obtain a response of 'false'.
     */
    function checkOperatorActiveAtBlockNumber(
        address operator,
        uint256 blockNumber,
        uint8 quorumNumber,
        uint256 stakeHistoryIndex
        ) external view returns (bool)
    {
        // fetch the `operator`'s pubkey hash
        bytes32 pubkeyHash = registry[operator].pubkeyHash;
        // pull the stake history entry specified by `stakeHistoryIndex`
        OperatorStakeUpdate memory operatorStakeUpdate = pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][stakeHistoryIndex];
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
     * @param operator is the operator of interest
     * @param blockNumber is the block number of interest
     * @param quorumNumber is the quorum number which the operator had no stake in
     * @param stakeHistoryIndex specifies an index in `pubkeyHashToStakeHistory[pubkeyHash]`, where `pubkeyHash` is looked up
     * in `registry[operator].pubkeyHash`
     * @return 'true' if it is succesfully proven that  the `operator` was inactive at the `blockNumber`, and 'false' otherwise
     * @dev In order for this function to return 'true', the inputs must satisfy all of the following list:
     * 1) `pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][index].updateBlockNumber <= blockNumber`
     * 2) `pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][index].nextUpdateBlockNumber` must be either `0` (signifying no next update) or
     * is must be strictly greater than `blockNumber`
     * 3) `pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][index].stake == 0` i.e. the operator had zero stake
     * @dev Note that a return value of 'false' does not guarantee that the `operator` was active at `blockNumber`, since a
     * bad `stakeHistoryIndex` can be supplied in order to obtain a response of 'false'.
     */
    function checkOperatorInactiveAtBlockNumber(
        address operator,
        uint256 blockNumber,
        uint8 quorumNumber,
        uint256 stakeHistoryIndex
        ) external view returns (bool)
    {
        bytes32 pubkeyHash = registry[operator].pubkeyHash;
        
        require(pubkeyHashToQuorumBitmap[pubkeyHash] >> quorumNumber & 1 == 1, "RegistryBase._checkOperatorInactiveAtBlockNumber: operator was not part of quorum");
        // special case for `pubkeyHashToStakeHistory[pubkeyHash]` having lenght zero -- in which case we know the operator was never registered
        if (pubkeyHashToStakeHistory[pubkeyHash][quorumNumber].length == 0) {
            return true;
        }
        // pull the stake history entry specified by `stakeHistoryIndex`
        OperatorStakeUpdate memory operatorStakeUpdate = pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][stakeHistoryIndex];
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
     * @notice Remove the operator from active status. Removes the operator with the given `pubkeyHash` from the `index` in `operatorList`,
     * updates operatorList and index histories, and performs other necessary updates for removing operator
     */
    function _removeOperator(address operator, bytes32 pubkeyHash, uint32 index) internal virtual {
        // remove the operator's stake
        _removeOperatorStake(pubkeyHash);

        // store blockNumber at which operator index changed (stopped being applicable)
        pubkeyHashToIndexHistory[pubkeyHash][pubkeyHashToIndexHistory[pubkeyHash].length - 1].toBlockNumber =
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
    function _removeOperatorStake(bytes32 pubkeyHash) internal {
        // loop through the operator's quorum bitmap and remove the operator's stake for each quorum
        uint256 quorumBitmap = pubkeyHashToQuorumBitmap[pubkeyHash];
        for (uint quorumNumber = 0; quorumNumber < quorumCount;) {
            if (quorumBitmap >> quorumNumber & 1 == 1) {
                _removeOperatorStakeForQuorum(pubkeyHash, uint8(quorumNumber));
            }
            unchecked {
                quorumNumber++;
            }
        }

    }

    /**
     * @notice Removes the stakes of the operator with pubkeyHash `pubkeyHash` for the quorum with number `quorumNumber`
     */
    function _removeOperatorStakeForQuorum(bytes32 pubkeyHash, uint8 quorumNumber) internal {
        // gas saving by caching length here
        uint256 pubkeyHashToStakeHistoryLengthMinusOne = pubkeyHashToStakeHistory[pubkeyHash][quorumNumber].length - 1;

        // determine current stakes
        OperatorStakeUpdate memory currentStakeUpdate =
            pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][pubkeyHashToStakeHistoryLengthMinusOne];
        //set nextUpdateBlockNumber in current stakes
        pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][pubkeyHashToStakeHistoryLengthMinusOne].nextUpdateBlockNumber =
            uint32(block.number);

        /**
         * @notice recording the information pertaining to change in stake for this operator in the history. operator stakes are set to 0 here.
         */
        pubkeyHashToStakeHistory[pubkeyHash][quorumNumber].push(
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

    /**
     * @notice Removes the registrant at the given `index` from the `operatorList`
     * @return swappedOperator is the operator who was swapped with the removed operator in the operatorList,
     * or the *zero address* in the case that the removed operator was already the list operator in the operatorList.
     */
    function _removeOperatorFromOperatorListAndIndex(uint32 index) internal returns (address swappedOperator) {
        // gas saving by caching length here
        uint256 operatorListLengthMinusOne = operatorList.length - 1;
        // Update index info for operator at end of list, if they are not the same as the removed operator
        if (index < operatorListLengthMinusOne) {
            // get existing operator at end of list, and retrieve their pubkeyHash
            swappedOperator = operatorList[operatorListLengthMinusOne];
            Operator memory registrant = registry[swappedOperator];
            bytes32 pubkeyHash = registrant.pubkeyHash;
            // store blockNumber at which operator index changed
            // same operation as above except pubkeyHash is now different (since different operator)
            pubkeyHashToIndexHistory[pubkeyHash][pubkeyHashToIndexHistory[pubkeyHash].length - 1].toBlockNumber =
                uint32(block.number);
            // push new 'OperatorIndex' struct to operator's array of historical indices, with 'index' set equal to 'index' input
            OperatorIndex memory operatorIndex;
            operatorIndex.index = index;
            pubkeyHashToIndexHistory[pubkeyHash].push(operatorIndex);

            // move 'swappedOperator' into 'index' slot in operatorList (swapping them with removed operator)
            operatorList[index] = swappedOperator;
        }

        // slither-disable-next-line costly-loop
        operatorList.pop();

        // Update totalOperatorsHistory
        _updateTotalOperatorsHistory();

        return swappedOperator;
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

    /// @notice Verify that the `operator` is an active operator and that they've provided the correct `index`
    function _deregistrationCheck(address operator, uint32 index) internal view {
        require(
            registry[operator].status == IQuorumRegistry.Status.ACTIVE,
            "RegistryBase._deregistrationCheck: Operator is not registered"
        );

        require(operator == operatorList[index], "RegistryBase._deregistrationCheck: Incorrect index supplied");
    }

    // storage gap
    uint256[50] private __GAP;
}