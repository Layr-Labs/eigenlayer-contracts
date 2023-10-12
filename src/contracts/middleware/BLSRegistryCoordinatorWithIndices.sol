// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/utils/cryptography/draft-EIP712.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

import "../interfaces/IBLSRegistryCoordinatorWithIndices.sol";
import "../interfaces/ISocketUpdater.sol";
import "../interfaces/IServiceManager.sol";
import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IVoteWeigher.sol";
import "../interfaces/IStakeRegistry.sol";
import "../interfaces/IIndexRegistry.sol";
import "../interfaces/IPauserRegistry.sol";

import "../libraries/EIP1271SignatureUtils.sol";
import "../libraries/BitmapUtils.sol";

import "../permissions/Pausable.sol";

/**
 * @title A `RegistryCoordinator` that has three registries:
 *      1) a `StakeRegistry` that keeps track of operators' stakes
 *      2) a `BLSPubkeyRegistry` that keeps track of operators' BLS public keys and aggregate BLS public keys for each quorum
 *      3) an `IndexRegistry` that keeps track of an ordered list of operators for each quorum
 * 
 * @author Layr Labs, Inc.
 */
contract BLSRegistryCoordinatorWithIndices is EIP712, Initializable, IBLSRegistryCoordinatorWithIndices, ISocketUpdater, Pausable {
    using BN254 for BN254.G1Point;

    /// @notice The EIP-712 typehash for the `DelegationApproval` struct used by the contract
    bytes32 public constant OPERATOR_CHURN_APPROVAL_TYPEHASH =
        keccak256("OperatorChurnApproval(bytes32 registeringOperatorId,OperatorKickParam[] operatorKickParams)OperatorKickParam(address operator,BN254.G1Point pubkey,bytes32[] operatorIdsToSwap)BN254.G1Point(uint256 x,uint256 y)");
    /// @notice The maximum value of a quorum bitmap
    uint256 internal constant MAX_QUORUM_BITMAP = type(uint192).max;
    /// @notice The basis point denominator
    uint16 internal constant BIPS_DENOMINATOR = 10000;
    /// @notice Index for flag that pauses operator registration
    uint8 internal constant PAUSED_REGISTER_OPERATOR = 0;
    /// @notice Index for flag that pauses operator deregistration
    uint8 internal constant PAUSED_DEREGISTER_OPERATOR = 1;

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

    /// @notice the mapping from quorum number to a quorums operator cap and kick parameters
    mapping(uint8 => OperatorSetParam) internal _quorumOperatorSetParams;
    /// @notice the mapping from operator's operatorId to the updates of the bitmap of quorums they are registered for
    mapping(bytes32 => QuorumBitmapUpdate[]) internal _operatorIdToQuorumBitmapHistory;
    /// @notice the mapping from operator's address to the operator struct
    mapping(address => Operator) internal _operators;
    /// @notice whether the salt has been used for an operator churn approval
    mapping(bytes32 => bool) public isChurnApproverSaltUsed;

    /// @notice the dynamic-length array of the registries this coordinator is coordinating
    address[] public registries;
    /// @notice the address of the entity allowed to sign off on operators getting kicked out of the AVS during registration
    address public churnApprover;
    /// @notice the address of the entity allowed to eject operators from the AVS
    address public ejector;

    modifier onlyServiceManagerOwner {
        require(msg.sender == serviceManager.owner(), "BLSRegistryCoordinatorWithIndices.onlyServiceManagerOwner: caller is not the service manager owner");
        _;
    }

    modifier onlyEjector {
        require(msg.sender == ejector, "BLSRegistryCoordinatorWithIndices.onlyEjector: caller is not the ejector");
        _;
    }

    constructor(
        ISlasher _slasher,
        IServiceManager _serviceManager,
        IStakeRegistry _stakeRegistry,
        IBLSPubkeyRegistry _blsPubkeyRegistry,
        IIndexRegistry _indexRegistry
    ) EIP712("AVSRegistryCoordinator", "v0.0.1") {
        slasher = _slasher;
        serviceManager = _serviceManager;
        stakeRegistry = _stakeRegistry;
        blsPubkeyRegistry = _blsPubkeyRegistry;
        indexRegistry = _indexRegistry;
    }

    function initialize(
        address _churnApprover,
        address _ejector,
        OperatorSetParam[] memory _operatorSetParams,
        IPauserRegistry _pauserRegistry,
        uint256 _initialPausedStatus
    ) external initializer {
        // set initial paused status
        _initializePauser(_pauserRegistry, _initialPausedStatus);
        // set the churnApprover
        _setChurnApprover(_churnApprover);
        // set the ejector
        _setEjector(_ejector);
        // add registry contracts to the registries array
        registries.push(address(stakeRegistry));
        registries.push(address(blsPubkeyRegistry));
        registries.push(address(indexRegistry));

        // set the operator set params
        require(IVoteWeigher(address(stakeRegistry)).quorumCount() == _operatorSetParams.length, "BLSRegistryCoordinatorWithIndices: operator set params length mismatch");
        for (uint8 i = 0; i < _operatorSetParams.length; i++) {
            _setOperatorSetParams(i, _operatorSetParams[i]);
        }
    }
    
    // VIEW FUNCTIONS

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

    /// @notice Returns the operator address for the given `operatorId`
    function getOperatorFromId(bytes32 operatorId) external view returns (address) {
        return blsPubkeyRegistry.getOperatorFromPubkeyHash(operatorId);
    }

    /// @notice Returns the status for the given `operator`
    function getOperatorStatus(address operator) external view returns (IRegistryCoordinator.OperatorStatus) {
        return _operators[operator].status;
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
                        "BLSRegistryCoordinatorWithIndices.getQuorumBitmapIndicesByOperatorIdsAtBlockNumber: operatorId has no quorumBitmaps at blockNumber"
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
            "BLSRegistryCoordinatorWithIndices.getQuorumBitmapByOperatorIdAtBlockNumberByIndex: quorumBitmapUpdate is from after blockNumber"
        );
        // if the next update is at or before the block number, then the quorum provided index is too early
        // if the nex update  block number is 0, then this is the latest update
        require(
            quorumBitmapUpdate.nextUpdateBlockNumber > blockNumber || quorumBitmapUpdate.nextUpdateBlockNumber == 0, 
            "BLSRegistryCoordinatorWithIndices.getQuorumBitmapByOperatorIdAtBlockNumberByIndex: quorumBitmapUpdate is from before blockNumber"
        );
        return quorumBitmapUpdate.quorumBitmap;
    }

    /// @notice Returns the `index`th entry in the operator with `operatorId`'s bitmap history
    function getQuorumBitmapUpdateByOperatorIdByIndex(bytes32 operatorId, uint256 index) external view returns (QuorumBitmapUpdate memory) {
        return _operatorIdToQuorumBitmapHistory[operatorId][index];
    }

    /// @notice Returns the current quorum bitmap for the given `operatorId` or 0 if the operator is not registered for any quorum
    function getCurrentQuorumBitmapByOperatorId(bytes32 operatorId) external view returns (uint192) {
        uint256 quorumBitmapHistoryLength = _operatorIdToQuorumBitmapHistory[operatorId].length;
        if (quorumBitmapHistoryLength == 0 || _operatorIdToQuorumBitmapHistory[operatorId][quorumBitmapHistoryLength - 1].nextUpdateBlockNumber != 0) {
            return 0;
        }
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
     * @notice Public function for the the churnApprover signature hash calculation when operators are being kicked from quorums
     * @param registeringOperatorId The is of the registering operator 
     * @param operatorKickParams The parameters needed to kick the operator from the quorums that have reached their caps
     * @param salt The salt to use for the churnApprover's signature
     * @param expiry The desired expiry time of the churnApprover's signature
     */
    function calculateOperatorChurnApprovalDigestHash(
        bytes32 registeringOperatorId,
        OperatorKickParam[] memory operatorKickParams,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        // calculate the digest hash
        return _hashTypedDataV4(keccak256(abi.encode(OPERATOR_CHURN_APPROVAL_TYPEHASH, registeringOperatorId, operatorKickParams, salt, expiry)));
    }

    // STATE CHANGING FUNCTIONS

    /**
     * @notice Sets parameters of the operator set for the given `quorumNumber`
     * @param quorumNumber is the quorum number to set the maximum number of operators for
     * @param operatorSetParam is the parameters of the operator set for the `quorumNumber`
     * @dev only callable by the service manager owner
     */
    function setOperatorSetParams(uint8 quorumNumber, OperatorSetParam memory operatorSetParam) external onlyServiceManagerOwner {
        _setOperatorSetParams(quorumNumber, operatorSetParam);
    }

    /**
     * @notice Sets the churnApprover
     * @param _churnApprover is the address of the churnApprover
     * @dev only callable by the service manager owner
     */
    function setChurnApprover(address _churnApprover) external onlyServiceManagerOwner {
        _setChurnApprover(_churnApprover);
    }

    /**
     * @notice Sets the ejector
     * @param _ejector is the address of the ejector
     * @dev only callable by the service manager owner
     */
    function setEjector(address _ejector) external onlyServiceManagerOwner {
        _setEjector(_ejector);
    }

    /**
     * @notice Registers msg.sender as an operator with the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registering for
     * @param registrationData is the data that is decoded to get the operator's registration information
     * @dev `registrationData` should be a G1 point representing the operator's BLS public key and their socket
     */
    function registerOperatorWithCoordinator(
        bytes calldata quorumNumbers,
        bytes calldata registrationData
    ) external onlyWhenNotPaused(PAUSED_REGISTER_OPERATOR) {
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
    function registerOperatorWithCoordinator(
        bytes calldata quorumNumbers,
        BN254.G1Point memory pubkey,
        string calldata socket
    ) external onlyWhenNotPaused(PAUSED_REGISTER_OPERATOR) {
        _registerOperatorWithCoordinatorAndNoOverfilledQuorums(msg.sender, quorumNumbers, pubkey, socket);
    }

    /**
     * @notice Registers msg.sender as an operator with the middleware when the quorum operator limit is full. To register 
     * while maintaining the limit, the operator chooses another registered operator with lower stake to kick.
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registering for
     * @param pubkey is the BLS public key of the operator
     * @param operatorKickParams are the parameters for the deregistration of the operator that is being kicked from each 
     * quorum that will be filled after the operator registers. These parameters should include an operator, their pubkey, 
     * and ids of the operators to swap with the kicked operator. 
     * @param signatureWithSaltAndExpiry is the signature of the churnApprover on the operator kick params with a salt and expiry
     */
    function registerOperatorWithCoordinator(
        bytes calldata quorumNumbers, 
        BN254.G1Point memory pubkey,
        string calldata socket,
        OperatorKickParam[] calldata operatorKickParams,
        SignatureWithSaltAndExpiry memory signatureWithSaltAndExpiry
    ) external onlyWhenNotPaused(PAUSED_REGISTER_OPERATOR) {
        // register the operator
        uint32[] memory numOperatorsPerQuorum = _registerOperatorWithCoordinator(msg.sender, quorumNumbers, pubkey, socket);

        // get the registering operator's operatorId and set the operatorIdsToSwap to it because the registering operator is the one with the greatest index
        bytes32[] memory operatorIdsToSwap = new bytes32[](1);
        operatorIdsToSwap[0] = pubkey.hashG1Point();

        // verify the churnApprover's signature
        _verifyChurnApproverSignatureOnOperatorChurnApproval(operatorIdsToSwap[0], operatorKickParams, signatureWithSaltAndExpiry);

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

                require(
                    operatorKickParams[i].quorumNumber == quorumNumber, 
                    "BLSRegistryCoordinatorWithIndices.registerOperatorWithCoordinator: quorumNumber not the same as signed"
                );

                // get the total stake for the quorum
                uint96 totalStakeForQuorum = stakeRegistry.getCurrentTotalStakeForQuorum(quorumNumber);
                bytes32 operatorToKickId = _operators[operatorKickParams[i].operator].operatorId;
                uint96 operatorToKickStake = stakeRegistry.getCurrentOperatorStakeForQuorum(operatorToKickId, quorumNumber);
                uint96 registeringOperatorStake = stakeRegistry.getCurrentOperatorStakeForQuorum(operatorIdsToSwap[0], quorumNumber);

                // check the registering operator has more than the kick BIPs of the operator to kick's stake
                require(
                    registeringOperatorStake > operatorToKickStake * operatorSetParam.kickBIPsOfOperatorStake / BIPS_DENOMINATOR,
                    "BLSRegistryCoordinatorWithIndices.registerOperatorWithCoordinator: registering operator has less than kickBIPsOfOperatorStake"
                );
                
                // check the that the operator to kick has less than the kick BIPs of the total stake
                require(
                    operatorToKickStake < totalStakeForQuorum * operatorSetParam.kickBIPsOfTotalStake / BIPS_DENOMINATOR,
                    "BLSRegistryCoordinatorWithIndices.registerOperatorWithCoordinator: operator to kick has more than kickBIPSOfTotalStake"
                );
            }
            
            // kick the operator
            _deregisterOperatorWithCoordinator(
                operatorKickParams[i].operator, 
                quorumNumbers[i:i+1], 
                operatorKickParams[i].pubkey, 
                operatorIdsToSwap
            );
        }
    }

    /**
     * @notice Deregisters the msg.sender as an operator from the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registered for
     * @param deregistrationData is the the data that is decoded to get the operator's deregistration information
     * @dev `deregistrationData` should be a tuple of the operator's BLS public key, the list of operator ids to swap
     */
    function deregisterOperatorWithCoordinator(
        bytes calldata quorumNumbers,
        bytes calldata deregistrationData
    ) external onlyWhenNotPaused(PAUSED_DEREGISTER_OPERATOR) {
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
     * @param operatorIdsToSwap is the list of the operator ids to swap the index of the operator with in each 
     * quorum when removing the operator from the quorum's ordered list. The provided operator ids should be the 
     * those of the operator's with the largest index in each quorum that the operator is deregistering from, in
     * ascending order of quorum number.
     */
    function deregisterOperatorWithCoordinator(
        bytes calldata quorumNumbers,
        BN254.G1Point memory pubkey,
        bytes32[] memory operatorIdsToSwap
    ) external onlyWhenNotPaused(PAUSED_DEREGISTER_OPERATOR) {
        _deregisterOperatorWithCoordinator(msg.sender, quorumNumbers, pubkey, operatorIdsToSwap);
    }

    /**
     * @notice Ejects the provided operator from the provided quorums from the AVS
     * @param operator is the operator to eject
     * @param quorumNumbers are the quorum numbers to eject the operator from
     * @param pubkey is the BLS public key of the operator
     * @param operatorIdsToSwap is the list of the operator ids to swap the index of the operator with in each 
     * quorum when removing the operator from the quorum's ordered list. The provided operator ids should be the 
     * those of the operator's with the largest index in each quorum that the operator is being ejected from, in
     * ascending order of quorum number.
     */
    function ejectOperatorFromCoordinator(
        address operator, 
        bytes calldata quorumNumbers, 
        BN254.G1Point memory pubkey, 
        bytes32[] memory operatorIdsToSwap
    ) external onlyEjector {
        _deregisterOperatorWithCoordinator(operator, quorumNumbers, pubkey, operatorIdsToSwap);
    }

    /**
     * @notice Updates the socket of the msg.sender given they are a registered operator
     * @param socket is the new socket of the operator
     */
    function updateSocket(string memory socket) external {
        require(_operators[msg.sender].status == OperatorStatus.REGISTERED, "BLSRegistryCoordinatorWithIndicies.updateSocket: operator is not registered");
        emit OperatorSocketUpdate(_operators[msg.sender].operatorId, socket);
    }

    // INTERNAL FUNCTIONS

    function _setOperatorSetParams(uint8 quorumNumber, OperatorSetParam memory operatorSetParam) internal {
        _quorumOperatorSetParams[quorumNumber] = operatorSetParam;
        emit OperatorSetParamsUpdated(quorumNumber, operatorSetParam);
    }
    
    function _setChurnApprover(address newChurnApprover) internal {
        emit ChurnApproverUpdated(churnApprover, newChurnApprover);
        churnApprover = newChurnApprover;
    }

    function _setEjector(address newEjector) internal {
        emit EjectorUpdated(ejector, newEjector);
        ejector = newEjector;
    }

    /// @return numOperatorsPerQuorum is the list of number of operators per quorum in quorumNumberss
    function _registerOperatorWithCoordinator(address operator, bytes calldata quorumNumbers, BN254.G1Point memory pubkey, string memory socket) internal returns(uint32[] memory) {
        // require(
        //     slasher.contractCanSlashOperatorUntilBlock(operator, address(serviceManager)) == type(uint32).max,
        //     "StakeRegistry._registerOperator: operator must be opted into slashing by the serviceManager"
        // );
        
        _beforeRegisterOperator(operator, quorumNumbers);
        // get the quorum bitmap from the quorum numbers
        uint256 quorumBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);
        require(quorumBitmap <= MAX_QUORUM_BITMAP, "BLSRegistryCoordinatorWithIndices._registerOperatorWithCoordinator: quorumBitmap exceeds of max bitmap size");
        require(quorumBitmap != 0, "BLSRegistryCoordinatorWithIndices._registerOperatorWithCoordinator: quorumBitmap cannot be 0");
        // register the operator with the BLSPubkeyRegistry and get the operatorId (in this case, the pubkeyHash) back
        bytes32 operatorId = blsPubkeyRegistry.registerOperator(operator, quorumNumbers, pubkey);

        uint256 operatorQuorumBitmapHistoryLength = _operatorIdToQuorumBitmapHistory[operatorId].length;
        if(operatorQuorumBitmapHistoryLength > 0) {
            uint256 prevQuorumBitmap = _operatorIdToQuorumBitmapHistory[operatorId][operatorQuorumBitmapHistoryLength - 1].quorumBitmap;
            require(prevQuorumBitmap & quorumBitmap == 0, "BLSRegistryCoordinatorWithIndices._registerOperatorWithCoordinator: operator already registered for some quorums being registered for");
            // new stored quorumBitmap is the previous quorumBitmap or'd with the new quorumBitmap to register for
            quorumBitmap |= prevQuorumBitmap;
        }

        // register the operator with the StakeRegistry
        stakeRegistry.registerOperator(operator, operatorId, quorumNumbers);

        // register the operator with the IndexRegistry
        uint32[] memory numOperatorsPerQuorum = indexRegistry.registerOperator(operatorId, quorumNumbers);

        uint256 quorumBitmapHistoryLength = _operatorIdToQuorumBitmapHistory[operatorId].length;
        if(quorumBitmapHistoryLength != 0) {
            // set the toBlockNumber of the previous quorum bitmap update
            _operatorIdToQuorumBitmapHistory[operatorId][quorumBitmapHistoryLength - 1].nextUpdateBlockNumber = uint32(block.number);
        }

        // set the operatorId to quorum bitmap history
        _operatorIdToQuorumBitmapHistory[operatorId].push(QuorumBitmapUpdate({
            updateBlockNumber: uint32(block.number),
            nextUpdateBlockNumber: 0,
            quorumBitmap: uint192(quorumBitmap)
        }));

        // set the operator struct
        if (_operators[operator].status != OperatorStatus.REGISTERED) {
            // if the operator is not already registered, then they are registering for the first time
            _operators[operator] = Operator({
                operatorId: operatorId,
                status: OperatorStatus.REGISTERED
            });
        }

        _afterRegisterOperator(operator, quorumNumbers);

        // record a stake update not bonding the operator at all (unbonded at 0), because they haven't served anything yet
        // serviceManager.recordFirstStakeUpdate(operator, 0);

        emit OperatorRegistered(operator, operatorId);

        emit OperatorSocketUpdate(operatorId, socket);

        return numOperatorsPerQuorum;
    }

    function _registerOperatorWithCoordinatorAndNoOverfilledQuorums(address operator, bytes calldata quorumNumbers, BN254.G1Point memory pubkey, string memory socket) internal {
        uint32[] memory numOperatorsPerQuorum = _registerOperatorWithCoordinator(operator, quorumNumbers, pubkey, socket);
        for (uint i = 0; i < numOperatorsPerQuorum.length; i++) {
            require(
                numOperatorsPerQuorum[i] <= _quorumOperatorSetParams[uint8(quorumNumbers[i])].maxOperatorCount,
                "BLSRegistryCoordinatorWithIndices._registerOperatorWithCoordinatorAndNoOverfilledQuorums: quorum is overfilled"
            );
        }
    }

    function _deregisterOperatorWithCoordinator(address operator, bytes calldata quorumNumbers, BN254.G1Point memory pubkey, bytes32[] memory operatorIdsToSwap) internal {
        require(_operators[operator].status == OperatorStatus.REGISTERED, "BLSRegistryCoordinatorWithIndices._deregisterOperatorWithCoordinator: operator is not registered");

        // get the operatorId of the operator
        bytes32 operatorId = _operators[operator].operatorId;
        require(operatorId == pubkey.hashG1Point(), "BLSRegistryCoordinatorWithIndices._deregisterOperatorWithCoordinator: operatorId does not match pubkey hash");

        // get the quorumNumbers of the operator
        uint256 quorumsToRemoveBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);
        require(quorumsToRemoveBitmap <= MAX_QUORUM_BITMAP, "BLSRegistryCoordinatorWithIndices._deregisterOperatorWithCoordinator: quorumsToRemoveBitmap exceeds of max bitmap size");
        uint256 operatorQuorumBitmapHistoryLengthMinusOne = _operatorIdToQuorumBitmapHistory[operatorId].length - 1;
        uint192 quorumBitmapBeforeUpdate = _operatorIdToQuorumBitmapHistory[operatorId][operatorQuorumBitmapHistoryLengthMinusOne].quorumBitmap;
        // check that the quorumNumbers of the operator matches the quorumNumbers passed in
        require(
            quorumBitmapBeforeUpdate & quorumsToRemoveBitmap == quorumsToRemoveBitmap,
            "BLSRegistryCoordinatorWithIndices._deregisterOperatorWithCoordinator: cannot deregister operator for quorums that it is not a part of"
        );
        // check if the operator is completely deregistering
        bool completeDeregistration = quorumBitmapBeforeUpdate == quorumsToRemoveBitmap;
        
        _beforeDeregisterOperator(operator, quorumNumbers);

        // deregister the operator from the BLSPubkeyRegistry
        blsPubkeyRegistry.deregisterOperator(operator, quorumNumbers, pubkey);

        // deregister the operator from the StakeRegistry
        stakeRegistry.deregisterOperator(operatorId, quorumNumbers);

        // deregister the operator from the IndexRegistry
        indexRegistry.deregisterOperator(operatorId, quorumNumbers, operatorIdsToSwap);

        _afterDeregisterOperator(operator, quorumNumbers);

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
            // uint32 latestServeUntilBlock = serviceManager.latestServeUntilBlock();

            // record a stake update unbonding the operator after `latestServeUntilBlock`
            // serviceManager.recordLastStakeUpdateAndRevokeSlashingAbility(operator, latestServeUntilBlock);
            // set the status of the operator to DEREGISTERED
            _operators[operator].status = OperatorStatus.DEREGISTERED;

            emit OperatorDeregistered(operator, operatorId);
        }
    }

    /**
     * @dev Hook that is called before any operator registration to insert additional logic.
     * @param operator The address of the operator to register.
     * @param quorumNumbers The quorum numbers the operator is registering for, where each byte is an 8 bit integer quorumNumber.
     */
    function _beforeRegisterOperator(address operator, bytes memory quorumNumbers) internal virtual{} 

    /**
     * @dev Hook that is called after any operator registration to insert additional logic.
     * @param operator The address of the operator to register.
     * @param quorumNumbers The quorum numbers the operator is registering for, where each byte is an 8 bit integer quorumNumber.
     */
    function _afterRegisterOperator(address operator, bytes memory quorumNumbers) internal virtual {}
    
    /**
     * @dev Hook that is called before any operator deregistration to insert additional logic.
     * @param operator The address of the operator to deregister.
     * @param quorumNumbers The quorum numbers the operator is registering for, where each byte is an 8 bit integer quorumNumber.
     */
    function _beforeDeregisterOperator(address operator, bytes memory quorumNumbers) internal virtual {}

    /**
     * @dev Hook that is called after any operator deregistration to insert additional logic.
     * @param operator The address of the operator to deregister.
     * @param quorumNumbers The quorum numbers the operator is registering for, where each byte is an 8 bit integer quorumNumber.
     */
    function _afterDeregisterOperator(address operator, bytes memory quorumNumbers) internal virtual {}

    /// @notice verifies churnApprover's signature on operator churn approval and increments the churnApprover nonce
    function _verifyChurnApproverSignatureOnOperatorChurnApproval(bytes32 registeringOperatorId, OperatorKickParam[] memory operatorKickParams, SignatureWithSaltAndExpiry memory signatureWithSaltAndExpiry) internal {
        // make sure the salt hasn't been used already
        require(!isChurnApproverSaltUsed[signatureWithSaltAndExpiry.salt], "BLSRegistryCoordinatorWithIndices._verifyChurnApproverSignatureOnOperatorChurnApproval: churnApprover salt already used");
        require(signatureWithSaltAndExpiry.expiry >= block.timestamp, "BLSRegistryCoordinatorWithIndices._verifyChurnApproverSignatureOnOperatorChurnApproval: churnApprover signature expired");   

        // set salt used to true
        isChurnApproverSaltUsed[signatureWithSaltAndExpiry.salt] = true;    

        // check the churnApprover's signature 
        EIP1271SignatureUtils.checkSignature_EIP1271(churnApprover, calculateOperatorChurnApprovalDigestHash(registeringOperatorId, operatorKickParams, signatureWithSaltAndExpiry.salt, signatureWithSaltAndExpiry.expiry), signatureWithSaltAndExpiry.signature);
    }
}
