// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

import "../interfaces/IBLSRegistryCoordinatorWithIndices.sol";
import "../interfaces/IServiceManager.sol";
import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IVoteWeigher.sol";
import "../interfaces/IStakeRegistry.sol";
import "../interfaces/IIndexRegistry.sol";

import "../libraries/BitmapUtils.sol";

import "forge-std/Test.sol";

/**
 * @title A `RegistryCoordinator` that has three registries:
 *      1) a `StakeRegistry` that keeps track of operators' stakes (this is actually the contract itself, via inheritance)
 *      2) a `BLSPubkeyRegistry` that keeps track of operators' BLS public keys and aggregate BLS public keys for each quorum
 *      3) an `IndexRegistry` that keeps track of an ordered list of operators for each quorum
 * 
 * @author Layr Labs, Inc.
 */
contract BLSRegistryCoordinatorWithIndices is Initializable, IBLSRegistryCoordinatorWithIndices, Test {
    using BN254 for BN254.G1Point;

    uint16 internal constant BIPS_DENOMINATOR = 10000;

    /// @notice the EigenLayer Slasher
    ISlasher public immutable slasher;
    /// @notice the Service Manager for the service that this contract is coordinating
    IServiceManager public immutable serviceManager;
    /// @notice the BLS Pubkey Registry contract that will keep track of operators' BLS public keys
    IBLSPubkeyRegistry public immutable blsPubkeyRegistry;
    /// @notice the Stake Registry contract that will keep track of operators' stakes
    IStakeRegistry public immutable stakeRegistry;
    /// @notice the Index Registry contract that will keep track of operators' indexes
    IIndexRegistry public immutable indexRegistry;
    /// @notice the mapping from quorum number to the maximum number of operators that can be registered for that quorum
    mapping(uint8 => OperatorSetParam) internal _quorumOperatorSetParams;
    /// @notice the mapping from operator's operatorId to the updates of the bitmap of quorums they are registered for
    mapping(bytes32 => QuorumBitmapUpdate[]) internal _operatorIdToQuorumBitmapHistory;
    /// @notice the mapping from operator's address to the operator struct
    mapping(address => Operator) internal _operators;
    /// @notice the dynamic-length array of the registries this coordinator is coordinating
    address[] public registries;

    modifier onlyServiceManagerOwner {
        require(msg.sender == serviceManager.owner(), "BLSIndexRegistryCoordinator.onlyServiceManagerOwner: caller is not the service manager owner");
        _;
    }

    constructor(
        ISlasher _slasher,
        IServiceManager _serviceManager,
        IStakeRegistry _stakeRegistry,
        IBLSPubkeyRegistry _blsPubkeyRegistry,
        IIndexRegistry _indexRegistry
    ) {
        slasher = _slasher;
        serviceManager = _serviceManager;
        stakeRegistry = _stakeRegistry;
        blsPubkeyRegistry = _blsPubkeyRegistry;
        indexRegistry = _indexRegistry;
    }

    function initialize(OperatorSetParam[] memory _operatorSetParams) external initializer {
        // the stake registry is this contract itself
        registries.push(address(stakeRegistry));
        registries.push(address(blsPubkeyRegistry));
        registries.push(address(indexRegistry));

        // set the operator set params
        require(IVoteWeigher(address(stakeRegistry)).quorumCount() == _operatorSetParams.length, "BLSIndexRegistryCoordinator: operator set params length mismatch");
        for (uint8 i = 0; i < _operatorSetParams.length; i++) {
            _setOperatorSetParams(i, _operatorSetParams[i]);
        }
    }

    /// @notice Returns the operator set params for the given `quorumNumber`
    function getOperatorSetParams(uint8 quorumNumber) external view returns (OperatorSetParam memory) {
        return _quorumOperatorSetParams[quorumNumber];
    }

    /// @notice Returns the operator struct for the given `operator`
    function getOperator(address operator) external view returns (Operator memory) {
        return _operators[operator];
    }

    /// @notice Returns the operatorId for the given `operator`
    function getOperatorId(address operator) external view returns (bytes32) {
        return _operators[operator].operatorId;
    }

    /// @notice Returns the indices of the quorumBitmaps for the provided `operatorIds` at the given `blockNumber`
    function getQuorumBitmapIndicesByOperatorIdsAtBlockNumber(uint32 blockNumber, bytes32[] memory operatorIds) external view returns (uint32[] memory) {
        uint32[] memory indices = new uint32[](operatorIds.length);
        for (uint256 i = 0; i < operatorIds.length; i++) {
            uint32 length = uint32(_operatorIdToQuorumBitmapHistory[operatorIds[i]].length);
            for (uint32 j = 0; j < length; j++) {
                if (_operatorIdToQuorumBitmapHistory[operatorIds[i]][length - j - 1].updateBlockNumber <= blockNumber) {
                    require(
                        _operatorIdToQuorumBitmapHistory[operatorIds[i]][length - j - 1].nextUpdateBlockNumber == 0 ||
                        _operatorIdToQuorumBitmapHistory[operatorIds[i]][length - j - 1].nextUpdateBlockNumber > blockNumber,
                        "BLSRegistryCoordinator.getQuorumBitmapIndicesByOperatorIdsAtBlockNumber: operatorId has no quorumBitmaps at blockNumber"
                    );
                    indices[i] = length - j - 1;
                    break;
                }
            }
        }
        return indices;
    }

    /**
     * @notice Returns the quorum bitmap for the given `operatorId` at the given `blockNumber` via the `index`
     * @dev reverts if `index` is incorrect 
     */ 
    function getQuorumBitmapByOperatorIdAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) external view returns (uint192) {
        QuorumBitmapUpdate memory quorumBitmapUpdate = _operatorIdToQuorumBitmapHistory[operatorId][index];
        require(
            quorumBitmapUpdate.updateBlockNumber <= blockNumber, 
            "BLSRegistryCoordinator.getQuorumBitmapByOperatorIdAtBlockNumberByIndex: quorumBitmapUpdate is from after blockNumber"
        );
        // if the next update is at or before the block number, then the quorum provided index is too early
        // if the nex update  block number is 0, then this is the latest update
        require(
            quorumBitmapUpdate.nextUpdateBlockNumber > blockNumber || quorumBitmapUpdate.nextUpdateBlockNumber == 0, 
            "BLSRegistryCoordinator.getQuorumBitmapByOperatorIdAtBlockNumberByIndex: quorumBitmapUpdate is from before blockNumber"
        );
        return quorumBitmapUpdate.quorumBitmap;
    }

    /// @notice Returns the `index`th entry in the operator with `operatorId`'s bitmap history
    function getQuorumBitmapUpdateByOperatorIdByIndex(bytes32 operatorId, uint256 index) external view returns (QuorumBitmapUpdate memory) {
        return _operatorIdToQuorumBitmapHistory[operatorId][index];
    }

    /// @notice Returns the current quorum bitmap for the given `operatorId`
    function getCurrentQuorumBitmapByOperatorId(bytes32 operatorId) external view returns (uint192) {
        uint256 quorumBitmapHistoryLength = _operatorIdToQuorumBitmapHistory[operatorId].length;
        if (quorumBitmapHistoryLength == 0) {
            revert("BLSRegistryCoordinator.getCurrentQuorumBitmapByOperatorId: no quorum bitmap history for operatorId");
        }
        require(_operatorIdToQuorumBitmapHistory[operatorId][quorumBitmapHistoryLength - 1].nextUpdateBlockNumber == 0, "BLSRegistryCoordinator.getCurrentQuorumBitmapByOperatorId: operator is not registered");
        return _operatorIdToQuorumBitmapHistory[operatorId][quorumBitmapHistoryLength - 1].quorumBitmap;
    }

    /// @notice Returns the length of the quorum bitmap history for the given `operatorId`
    function getQuorumBitmapUpdateByOperatorIdLength(bytes32 operatorId) external view returns (uint256) {
        return _operatorIdToQuorumBitmapHistory[operatorId].length;
    }

    /// @notice Returns the number of registries
    function numRegistries() external view returns (uint256) {
        return registries.length;
    }

    /**
     * @notice Sets parameters of the operator set for the given `quorumNumber`
     * @param quorumNumber is the quorum number to set the maximum number of operators for
     * @param operatorSetParam is the parameters of the operator set for the `quorumNumber`
     */
    function setOperatorSetParams(uint8 quorumNumber, OperatorSetParam memory operatorSetParam) external onlyServiceManagerOwner {
        _setOperatorSetParams(quorumNumber, operatorSetParam);
    }

    /**
     * @notice Registers msg.sender as an operator with the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registering for
     * @param registrationData is the data that is decoded to get the operator's registration information
     * @dev `registrationData` should be a G1 point representing the operator's BLS public key and their socket
     */
    function registerOperatorWithCoordinator(bytes calldata quorumNumbers, bytes calldata registrationData) external {
        // get the operator's BLS public key
        (BN254.G1Point memory pubkey, string memory socket) = abi.decode(registrationData, (BN254.G1Point, string));
        // call internal function to register the operator
        _registerOperatorWithCoordinatorAndNoOverfilledQuorums(msg.sender, quorumNumbers, pubkey, socket);
    }

    /**
     * @notice Registers msg.sender as an operator with the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registering for
     * @param pubkey is the BLS public key of the operator
     * @param socket is the socket of the operator
     */
    function registerOperatorWithCoordinator(bytes calldata quorumNumbers, BN254.G1Point memory pubkey, string calldata socket) external {
        _registerOperatorWithCoordinatorAndNoOverfilledQuorums(msg.sender, quorumNumbers, pubkey, socket);
    }

    /**
     * @notice Registers msg.sender as an operator with the middleware when the quorum operator limit is full. To register 
     * while maintaining the limit, the operator chooses another registered opeerator with lower stake to kick.
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registering for
     * @param pubkey is the BLS public key of the operator
     * @param operatorKickParams are the parameters for the deregistration of the operator that is being kicked
     */
    function registerOperatorWithCoordinator(
        bytes calldata quorumNumbers, 
        BN254.G1Point memory pubkey,
        string calldata socket,
        OperatorKickParam[] calldata operatorKickParams
    ) external {
        // register the operator
        uint32[] memory numOperatorsPerQuorum = _registerOperatorWithCoordinator(msg.sender, quorumNumbers, pubkey, socket);

        // get the registering operator's operatorId
        bytes32 registeringOperatorId = _operators[msg.sender].operatorId;

        // kick the operators
        for (uint256 i = 0; i < quorumNumbers.length; i++) {
            // check that the quorum has reached the max operator count
            uint8 quorumNumber = uint8(quorumNumbers[i]);
            OperatorSetParam memory operatorSetParam = _quorumOperatorSetParams[quorumNumber];
            {
                // if the number of operators for the quorum is less than or equal to the max operator count, 
                // then the quorum has not reached the max operator count
                if(numOperatorsPerQuorum[i] <= operatorSetParam.maxOperatorCount) {
                    continue;
                }

                // get the total stake for the quorum
                uint96 totalStakeForQuorum = stakeRegistry.getCurrentTotalStakeForQuorum(quorumNumber);
                bytes32 operatorToKickId = _operators[operatorKickParams[i].operator].operatorId;
                uint96 operatorToKickStake = stakeRegistry.getCurrentOperatorStakeForQuorum(operatorToKickId, quorumNumber);
                uint96 registeringOperatorStake = stakeRegistry.getCurrentOperatorStakeForQuorum(registeringOperatorId, quorumNumber);

                // check the registering operator has more than the kick BIPs of the operator to kick's stake
                require(
                    registeringOperatorStake > operatorToKickStake * operatorSetParam.kickBIPsOfOperatorStake / BIPS_DENOMINATOR,
                    "BLSIndexRegistryCoordinator.registerOperatorWithCoordinator: registering operator has less than kickBIPsOfOperatorStake"
                );
                
                // check the that the operator to kick has less than the kick BIPs of the total stake
                require(
                    operatorToKickStake < totalStakeForQuorum * operatorSetParam.kickBIPsOfTotalStake / BIPS_DENOMINATOR,
                    "BLSIndexRegistryCoordinator.registerOperatorWithCoordinator: operator to kick has more than kickBIPSOfTotalStake"
                );
            }
            
            // kick the operator
            _deregisterOperatorWithCoordinator(
                operatorKickParams[i].operator, 
                quorumNumbers[i:i+1], 
                operatorKickParams[i].pubkey, 
                operatorKickParams[i].operatorIdsToSwap
            );
        }
    }

    /**
     * @notice Deregisters the msg.sender as an operator from the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registered for
     * @param deregistrationData is the the data that is decoded to get the operator's deregistration information
     * @dev `deregistrationData` should be a tuple of the operator's BLS public key, the list of operator ids to swap
     */
    function deregisterOperatorWithCoordinator(bytes calldata quorumNumbers, bytes calldata deregistrationData) external {
        // get the operator's deregistration information
        (BN254.G1Point memory pubkey, bytes32[] memory operatorIdsToSwap) 
            = abi.decode(deregistrationData, (BN254.G1Point, bytes32[]));
        // call internal function to deregister the operator
        _deregisterOperatorWithCoordinator(msg.sender, quorumNumbers, pubkey, operatorIdsToSwap);
    }

    /**
     * @notice Deregisters the msg.sender as an operator from the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registered for
     * @param pubkey is the BLS public key of the operator
     * @param operatorIdsToSwap is the list of the operator ids tho swap the index of the operator with in each 
     * quorum when removing the operator from the quorum's ordered list
     */
    function deregisterOperatorWithCoordinator(bytes calldata quorumNumbers, BN254.G1Point memory pubkey, bytes32[] memory operatorIdsToSwap) external {
        _deregisterOperatorWithCoordinator(msg.sender, quorumNumbers, pubkey, operatorIdsToSwap);
    }

    // INTERNAL FUNCTIONS

    function _setOperatorSetParams(uint8 quorumNumber, OperatorSetParam memory operatorSetParam) internal {
        _quorumOperatorSetParams[quorumNumber] = operatorSetParam;
        emit OperatorSetParamsUpdated(quorumNumber, operatorSetParam);
    }

    /// @return numOperatorsPerQuorum is the list of number of operators per quorum in quorumNumberss
    function _registerOperatorWithCoordinator(address operator, bytes calldata quorumNumbers, BN254.G1Point memory pubkey, string memory socket) internal returns(uint32[] memory) {
        // require(
        //     slasher.contractCanSlashOperatorUntilBlock(operator, address(serviceManager)) == type(uint32).max,
        //     "StakeRegistry._registerOperator: operator must be opted into slashing by the serviceManager"
        // );
        
        // check that the sender is not already registered
        require(_operators[operator].status != OperatorStatus.REGISTERED, "BLSIndexRegistryCoordinator._registerOperatorWithCoordinator: operator already registered");

        // get the quorum bitmap from the quorum numbers
        uint256 quorumBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);

        require(quorumBitmap != 0, "BLSIndexRegistryCoordinator._registerOperatorWithCoordinator: quorumBitmap cannot be 0");
        require(quorumBitmap <= type(uint192).max, "BLSIndexRegistryCoordinator._registerOperatorWithCoordinator: quorumBitmap cant have more than 192 set bits");

        // register the operator with the BLSPubkeyRegistry and get the operatorId (in this case, the pubkeyHash) back
        bytes32 operatorId = blsPubkeyRegistry.registerOperator(operator, quorumNumbers, pubkey);

        // register the operator with the StakeRegistry
        stakeRegistry.registerOperator(operator, operatorId, quorumNumbers);

        // register the operator with the IndexRegistry
        uint32[] memory numOperatorsPerQuorum = indexRegistry.registerOperator(operatorId, quorumNumbers);

        // set the operatorId to quorum bitmap history
        _operatorIdToQuorumBitmapHistory[operatorId].push(QuorumBitmapUpdate({
            updateBlockNumber: uint32(block.number),
            nextUpdateBlockNumber: 0,
            quorumBitmap: uint192(quorumBitmap)
        }));

        // set the operator struct
        _operators[operator] = Operator({
            operatorId: operatorId,
            status: OperatorStatus.REGISTERED
        });

        // record a stake update not bonding the operator at all (unbonded at 0), because they haven't served anything yet
        // serviceManager.recordFirstStakeUpdate(operator, 0);

        emit OperatorSocketUpdate(operatorId, socket);

        return numOperatorsPerQuorum;
    }

    function _registerOperatorWithCoordinatorAndNoOverfilledQuorums(address operator, bytes calldata quorumNumbers, BN254.G1Point memory pubkey, string memory socket) internal {
        uint32[] memory numOperatorsPerQuorum = _registerOperatorWithCoordinator(operator, quorumNumbers, pubkey, socket);
        for (uint i = 0; i < numOperatorsPerQuorum.length; i++) {
            require(
                numOperatorsPerQuorum[i] <= _quorumOperatorSetParams[uint8(quorumNumbers[i])].maxOperatorCount,
                "BLSIndexRegistryCoordinator._registerOperatorWithCoordinatorAndNoOverfilledQuorums: quorum is overfilled"
            );
        }
    }

    function _deregisterOperatorWithCoordinator(address operator, bytes calldata quorumNumbers, BN254.G1Point memory pubkey, bytes32[] memory operatorIdsToSwap) internal {
        require(_operators[operator].status == OperatorStatus.REGISTERED, "BLSIndexRegistryCoordinator._deregisterOperatorWithCoordinator: operator is not registered");

        // get the operatorId of the operator
        bytes32 operatorId = _operators[operator].operatorId;
        require(operatorId == pubkey.hashG1Point(), "BLSIndexRegistryCoordinator._deregisterOperatorWithCoordinator: operatorId does not match pubkey hash");

        // get the quorumNumbers of the operator
        uint256 quorumsToRemoveBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);
        require(quorumsToRemoveBitmap <= type(uint192).max, "BLSIndexRegistryCoordinator._deregisterOperatorWithCoordinator: quorumsToRemoveBitmap cant have more than 192 set bits");
        uint256 operatorQuorumBitmapHistoryLengthMinusOne = _operatorIdToQuorumBitmapHistory[operatorId].length - 1;
        uint192 quorumBitmapBeforeUpdate = _operatorIdToQuorumBitmapHistory[operatorId][operatorQuorumBitmapHistoryLengthMinusOne].quorumBitmap;
        // check that the quorumNumbers of the operator matches the quorumNumbers passed in
        require(
            quorumBitmapBeforeUpdate & quorumsToRemoveBitmap == quorumsToRemoveBitmap,
            "BLSIndexRegistryCoordinator._deregisterOperatorWithCoordinator: cannot deregister operator for quorums that it is not a part of"
        );
        // check if the operator is completely deregistering
        bool completeDeregistration = quorumBitmapBeforeUpdate == quorumsToRemoveBitmap;
        
        // deregister the operator from the BLSPubkeyRegistry
        blsPubkeyRegistry.deregisterOperator(operator, quorumNumbers, pubkey);

        // deregister the operator from the StakeRegistry
        stakeRegistry.deregisterOperator(operatorId, quorumNumbers);

        // deregister the operator from the IndexRegistry
        indexRegistry.deregisterOperator(operatorId, quorumNumbers, operatorIdsToSwap);

        // set the toBlockNumber of the operator's quorum bitmap update
        _operatorIdToQuorumBitmapHistory[operatorId][operatorQuorumBitmapHistoryLengthMinusOne].nextUpdateBlockNumber = uint32(block.number);
        
        // if it is not a complete deregistration, add a new quorum bitmap update
        if (!completeDeregistration) {
            _operatorIdToQuorumBitmapHistory[operatorId].push(QuorumBitmapUpdate({
                updateBlockNumber: uint32(block.number),
                nextUpdateBlockNumber: 0,
                quorumBitmap: quorumBitmapBeforeUpdate & ~uint192(quorumsToRemoveBitmap) // this removes the quorumsToRemoveBitmap from the quorumBitmapBeforeUpdate
            }));
        } else {
            // @notice Registrant must continue to serve until the latest block at which an active task expires. this info is used in challenges
            uint32 latestServeUntilBlock = serviceManager.latestServeUntilBlock();

            // record a stake update unbonding the operator after `latestServeUntilBlock`
            serviceManager.recordLastStakeUpdateAndRevokeSlashingAbility(operator, latestServeUntilBlock);
            // set the status of the operator to DEREGISTERED
            _operators[operator].status = OperatorStatus.DEREGISTERED;
        }
    }
}