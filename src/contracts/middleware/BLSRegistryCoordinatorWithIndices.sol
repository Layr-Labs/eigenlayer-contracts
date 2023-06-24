// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

import "../interfaces/IBLSRegistryCoordinatorWithIndices.sol";
import "../interfaces/IServiceManager.sol";
import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IVoteWeigher.sol";
import "../interfaces/IStakeRegistry.sol";
import "../interfaces/IIndexRegistry.sol";

import "../libraries/BytesArrayBitmaps.sol";

/**
 * @title A `RegistryCoordinator` that has three registries:
 *      1) a `StakeRegistry` that keeps track of operators' stakes (this is actually the contract itself, via inheritance)
 *      2) a `BLSPubkeyRegistry` that keeps track of operators' BLS public keys and aggregate BLS public keys for each quorum
 *      3) an `IndexRegistry` that keeps track of an ordered list of operators for each quorum
 * 
 * @author Layr Labs, Inc.
 */
contract BLSRegistryCoordinatorWithIndices is Initializable, IBLSRegistryCoordinatorWithIndices {
    using BN254 for BN254.G1Point;

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
    mapping(uint8 => OperatorSetParam) public quorumOperatorSetParams;
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
            quorumOperatorSetParams[i] = _operatorSetParams[i];
        }
    }

    /// @notice Returns task number from when `operator` has been registered.
    function getFromTaskNumberForOperator(address operator) external view returns (uint32) {
        return _operators[operator].fromTaskNumber;
    }

    /// @notice Returns the operator struct for the given `operator`
    function getOperator(address operator) external view returns (Operator memory) {
        return _operators[operator];
    }

    /// @notice Returns the operatorId for the given `operator`
    function getOperatorId(address operator) external view returns (bytes32) {
        return _operators[operator].operatorId;
    }

    /// @notice Returns the quorum bitmap for the given `operatorId` at the given `blockNumber` via the `index`
    function getQuorumBitmapByOperatorIdAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) external view returns (uint192) {
        QuorumBitmapUpdate memory quorumBitmapUpdate = _operatorIdToQuorumBitmapHistory[operatorId][index];
        require(
            quorumBitmapUpdate.updateBlockNumber <= blockNumber, 
            "BLSRegistryCoordinator.getQuorumBitmapByOperatorIdAtBlockNumberByIndex: quorumBitmapUpdate is from after blockNumber"
        );
        require(
            quorumBitmapUpdate.nextUpdateBlockNumber > blockNumber, 
            "BLSRegistryCoordinator.getQuorumBitmapByOperatorIdAtBlockNumberByIndex: quorumBitmapUpdate is from before blockNumber"
        );
        return quorumBitmapUpdate.quorumBitmap;
    }

    /// @notice Returns the current quorum bitmap for the given `operatorId`
    function getCurrentQuorumBitmapByOperatorId(bytes32 operatorId) external view returns (uint192) {
        return _operatorIdToQuorumBitmapHistory[operatorId][_operatorIdToQuorumBitmapHistory[operatorId].length - 1].quorumBitmap;
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
        quorumOperatorSetParams[quorumNumber] = operatorSetParam;
    }

    /**
     * @notice Registers msg.sender as an operator with the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registering for
     * @param registrationData is the data that is decoded to get the operator's registration information
     * @dev `registrationData` should be a G1 point representing the operator's BLS public key
     */
    function registerOperatorWithCoordinator(bytes calldata quorumNumbers, bytes calldata registrationData) external {
        // get the operator's BLS public key
        BN254.G1Point memory pubkey = abi.decode(registrationData, (BN254.G1Point));
        // call internal function to register the operator
        _registerOperatorWithCoordinator(msg.sender, quorumNumbers, pubkey);
    }

    /**
     * @notice Registers msg.sender as an operator with the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registering for
     * @param pubkey is the BLS public key of the operator
     */
    function registerOperatorWithCoordinator(bytes calldata quorumNumbers, BN254.G1Point memory pubkey) external {
        _registerOperatorWithCoordinator(msg.sender, quorumNumbers, pubkey);
    }

    /**
     * @notice Registers msg.sender as an operator with the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registering for
     * @param pubkey is the BLS public key of the operator
     */
    function registerOperatorWithCoordinator(
        bytes calldata quorumNumbers, 
        BN254.G1Point memory pubkey,
        OperatorKickParam[] calldata operatorKickParams
    ) external {
        require(quorumNumbers.length == operatorKickParams.length, "BLSIndexRegistryCoordinator.registerOperatorWithCoordinator: quorumNumbers and operatorKickParams must be the same length");
        // register the operator
        _registerOperatorWithCoordinator(msg.sender, quorumNumbers, pubkey);

        // get the registering operator's operatorId
        bytes32 registeringOperatorId = _operators[msg.sender].operatorId;

        // kick the operators
        for (uint256 i = 0; i < quorumNumbers.length; i++) {
            // check that the quorum has reached the max operator count
            uint8 quorumNumber = uint8(quorumNumbers[i]);
            OperatorSetParam memory operatorSetParam = quorumOperatorSetParams[quorumNumber];
            {
                uint32 numOperatorsForQuorum = indexRegistry.totalOperatorsForQuorum(quorumNumber);
                require(
                    numOperatorsForQuorum == operatorSetParam.maxOperatorCount + 1,
                    "BLSIndexRegistryCoordinator.registerOperatorWithCoordinator: quorum has not reached max operator count"
                );

                // get the total stake for the quorum
                uint96 totalStakeForQuorum = stakeRegistry.getCurrentTotalStakeForQuorum(quorumNumber);
                bytes32 operatorToKickId = _operators[operatorKickParams[i].operator].operatorId;
                uint96 operatorToKickStake = stakeRegistry.getCurrentOperatorStakeForQuorum(operatorToKickId, quorumNumber);
                uint96 registeringOperatorStake = stakeRegistry.getCurrentOperatorStakeForQuorum(registeringOperatorId, quorumNumber);

                // check the registering operator has more than the kick percentage of the operator to kick's stake
                require(
                    registeringOperatorStake > operatorToKickStake * operatorSetParam.kickPercentageOfOperatorStake / 100,
                    "BLSIndexRegistryCoordinator.registerOperatorWithCoordinator: registering operator has less than kickPercentageOfOperatorStake"
                );
                
                // check that the operator to kick has less than the kick percentage of the average stake
                require(
                    operatorToKickStake < totalStakeForQuorum * operatorSetParam.kickPercentageOfAverageStake / 100 / numOperatorsForQuorum,
                    "BLSIndexRegistryCoordinator.registerOperatorWithCoordinator: operator to kick has more than kickPercentageOfAverageStake"
                );
                // check the that the operator to kick has lss than the kick percentage of the total stake
                require(
                    operatorToKickStake < totalStakeForQuorum * operatorSetParam.kickPercentageOfTotalStake / 100,
                    "BLSIndexRegistryCoordinator.registerOperatorWithCoordinator: operator to kick has more than kickPercentageOfTotalStake"
                );
            }
            
            // kick the operator
            _deregisterOperatorWithCoordinator(
                operatorKickParams[i].operator, 
                quorumNumbers[i:i+1], 
                operatorKickParams[i].pubkey, 
                operatorKickParams[i].operatorIdsToSwap, 
                operatorKickParams[i].globalOperatorListIndex
            );
        }
    }

    /**
     * @notice Deregisters the msg.sender as an operator from the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registered for
     * @param deregistrationData is the the data that is decoded to get the operator's deregisteration information
     * @dev `deregistrationData` should be a tuple of the operator's BLS public key, the list of operator ids to swap, 
     * and the operator's index in the global operator list
     */
    function deregisterOperatorWithCoordinator(bytes calldata quorumNumbers, bytes calldata deregistrationData) external {
        // get the operator's deregisteration information
        (BN254.G1Point memory pubkey, bytes32[] memory operatorIdsToSwap, uint32 globalOperatorListIndex) 
            = abi.decode(deregistrationData, (BN254.G1Point, bytes32[], uint32));
        // call internal function to deregister the operator
        _deregisterOperatorWithCoordinator(msg.sender, quorumNumbers, pubkey, operatorIdsToSwap, globalOperatorListIndex);
    }

    /**
     * @notice Deregisters the msg.sender as an operator from the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registered for
     * @param pubkey is the BLS public key of the operator
     * @param operatorIdsToSwap is the list of the operator ids that the should swap for the deregistering operator's index
     * @param globalOperatorListIndex is the operator's index in the global operator list in the IndexRegistry
     */
    function deregisterOperatorWithCoordinator(bytes calldata quorumNumbers, BN254.G1Point memory pubkey, bytes32[] memory operatorIdsToSwap, uint32 globalOperatorListIndex) external {
        _deregisterOperatorWithCoordinator(msg.sender, quorumNumbers, pubkey, operatorIdsToSwap, globalOperatorListIndex);
    }

    function _registerOperatorWithCoordinator(address operator, bytes calldata quorumNumbers, BN254.G1Point memory pubkey) internal {
        require(
            slasher.contractCanSlashOperatorUntilBlock(operator, address(serviceManager)) == type(uint32).max,
            "StakeRegistry._registerOperator: operator must be opted into slashing by the serviceManager"
        );
        
        // check that the sender is not already registered
        require(_operators[operator].status != OperatorStatus.REGISTERED, "BLSIndexRegistryCoordinator._registerOperatorWithCoordinator: operator already registered");

        // get the quorum bitmap from the quorum numbers
        uint192 quorumBitmap = uint192(BytesArrayBitmaps.orderedBytesArrayToBitmap_Yul(quorumNumbers));
        require(quorumBitmap != 0, "BLSIndexRegistryCoordinator._registerOperatorWithCoordinator: quorumBitmap cannot be 0");

        // register the operator with the BLSPubkeyRegistry and get the operatorId (in this case, the pubkeyHash) back
        bytes32 operatorId = blsPubkeyRegistry.registerOperator(operator, quorumNumbers, pubkey);

        // register the operator with the StakeRegistry
        stakeRegistry.registerOperator(operator, operatorId, quorumNumbers);

        // register the operator with the IndexRegistry
        indexRegistry.registerOperator(operatorId, quorumNumbers);

        // set the operatorId to quorum bitmap history
        _operatorIdToQuorumBitmapHistory[operatorId].push(QuorumBitmapUpdate({
            updateBlockNumber: uint32(block.number),
            nextUpdateBlockNumber: 0,
            quorumBitmap: quorumBitmap
        }));

        // set the operator struct
        _operators[operator] = Operator({
            operatorId: operatorId,
            fromTaskNumber: serviceManager.taskNumber(),
            status: OperatorStatus.REGISTERED
        });

        // record a stake update not bonding the operator at all (unbonded at 0), because they haven't served anything yet
        serviceManager.recordFirstStakeUpdate(operator, 0);
    }

    function _deregisterOperatorWithCoordinator(address operator, bytes calldata quorumNumbers, BN254.G1Point memory pubkey, bytes32[] memory operatorIdsToSwap, uint32 globalOperatorListIndex) internal {
        require(_operators[operator].status == OperatorStatus.REGISTERED, "BLSIndexRegistryCoordinator._deregisterOperatorWithCoordinator: operator is not registered");

        // get the operatorId of the operator
        bytes32 operatorId = _operators[operator].operatorId;
        require(operatorId == pubkey.hashG1Point(), "BLSIndexRegistryCoordinator._deregisterOperatorWithCoordinator: operatorId does not match pubkey hash");

        // get the quorumNumbers of the operator
        uint192 quorumsToRemoveBitmap = uint192(BytesArrayBitmaps.orderedBytesArrayToBitmap_Yul(quorumNumbers));
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
        blsPubkeyRegistry.deregisterOperator(operator, completeDeregistration, quorumNumbers, pubkey);

        // deregister the operator from the StakeRegistry
        stakeRegistry.deregisterOperator(operatorId, quorumNumbers);

        // deregister the operator from the IndexRegistry
        indexRegistry.deregisterOperator(operatorId, completeDeregistration, quorumNumbers, operatorIdsToSwap, globalOperatorListIndex);

        // set the toBlockNumber of the operator's quorum bitmap update
        _operatorIdToQuorumBitmapHistory[operatorId][operatorQuorumBitmapHistoryLengthMinusOne].nextUpdateBlockNumber = uint32(block.number);
        
        // if it is not a complete deregistration, add a new quorum bitmap update
        if (!completeDeregistration) {
            _operatorIdToQuorumBitmapHistory[operatorId].push(QuorumBitmapUpdate({
                updateBlockNumber: uint32(block.number),
                nextUpdateBlockNumber: 0,
                quorumBitmap: quorumBitmapBeforeUpdate & ~quorumsToRemoveBitmap // this removes the quorumsToRemoveBitmap from the quorumBitmapBeforeUpdate
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