// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IRegistryCoordinator.sol";
import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IIndexRegistry.sol";

import "../libraries/BytesArrayBitmaps.sol";

import "./StakeRegistry.sol";

/**
 * @title A `RegistryCoordinator` that has three registries:
 *      1) a `StakeRegistry` that keeps track of operators' stakes (this is actually the contract itself, via inheritance)
 *      2) a `BLSPubkeyRegistry` that keeps track of operators' BLS public keys and aggregate BLS public keys for each quorum
 *      3) an `IndexRegistry` that keeps track of an ordered list of operators for each quorum
 * 
 * @author Layr Labs, Inc.
 */
contract BLSIndexRegistryCoordinator is StakeRegistry, IRegistryCoordinator {
    using BN254 for BN254.G1Point;

    /// @notice the BLS Pubkey Registry contract that will keep track of operators' BLS public keys
    IBLSPubkeyRegistry public immutable blsPubkeyRegistry;
    /// @notice the Index Registry contract that will keep track of operators' indexes
    IIndexRegistry public immutable indexRegistry;
    /// @notice the mapping from operator's operatorId to the bitmap of quorums they are registered for
    mapping(bytes32 => uint256) public operatorIdToQuorumBitmap;
    /// @notice the mapping from operator's address to the operator struct
    mapping(address => Operator) internal _operators;
    /// @notice the dynamic-length array of the registries this coordinator is coordinating
    address[] public registries;

    constructor(
        IStrategyManager _strategyManager,
        IServiceManager _serviceManager,
        IBLSPubkeyRegistry _blsPubkeyRegistry,
        IIndexRegistry _indexRegistry
    ) StakeRegistry(_strategyManager, _serviceManager) {
        blsPubkeyRegistry = _blsPubkeyRegistry;
        indexRegistry = _indexRegistry;
    }

    function initialize(
        uint96[] memory _minimumStakeForQuorum,
        StrategyAndWeightingMultiplier[][] memory _quorumStrategiesConsideredAndMultipliers
    ) external initializer {
        // the stake registry is this contract itself
        registries.push(address(this));
        registries.push(address(blsPubkeyRegistry));
        registries.push(address(indexRegistry));

        // this contract is the registry coordinator for the stake registry
        StakeRegistry._initialize(IRegistryCoordinator(address(this)), _minimumStakeForQuorum, _quorumStrategiesConsideredAndMultipliers);
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

    /// @notice Returns the number of registries
    function numRegistries() external view returns (uint256) {
        return registries.length;
    }

    /// @notice disabled function on the StakeRegistry because this is the registry coordinator
    function registerOperator(address, bytes32, bytes calldata) external override pure {
        revert("BLSIndexRegistryCoordinator.registerOperator: cannot use overrided StakeRegistry.registerOperator on BLSIndexRegistryCoordinator");
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
    function registerOperator(bytes calldata quorumNumbers, BN254.G1Point memory pubkey) external {
        _registerOperatorWithCoordinator(msg.sender, quorumNumbers, pubkey);
    }

    /// @notice disabled function on the StakeRegistry because this is the registry coordinator
    function deregisterOperator(address, bytes32, bytes calldata) external override pure {
        revert("BLSIndexRegistryCoordinator.deregisterOperator: cannot use overrided StakeRegistry.deregisterOperator on BLSIndexRegistryCoordinator");
    }
    
    /**
     * @notice Deregisters the msg.sender as an operator from the middleware
     * @param deregistrationData is the the data that is decoded to get the operator's deregisteration information
     * @dev `deregistrationData` should be a tuple of the operator's BLS public key, the list of operator ids to swap,
     * and the operator's index in the global operator list
     */
    function deregisterOperatorWithCoordinator(bytes calldata deregistrationData) external {
        // get the operator's deregisteration information
        (BN254.G1Point memory pubkey, bytes32[] memory operatorIdsToSwap, uint32 globalOperatorListIndex) 
            = abi.decode(deregistrationData, (BN254.G1Point, bytes32[], uint32));
        // call internal function to deregister the operator
        _deregisterOperatorWithCoordinator(msg.sender, pubkey, operatorIdsToSwap, globalOperatorListIndex);
    }

    /**
     * @notice Deregisters the msg.sender as an operator from the middleware
     * @param pubkey is the BLS public key of the operator
     * @param operatorIdsToSwap is the list of the operator ids that the should swap for the deregistering operator's index
     * @param globalOperatorListIndex is the operator's index in the global operator list in the IndexRegistry
     */
    function deregisterOperatorWithCoordinator(BN254.G1Point memory pubkey, bytes32[] memory operatorIdsToSwap, uint32 globalOperatorListIndex) external {
        _deregisterOperatorWithCoordinator(msg.sender, pubkey, operatorIdsToSwap, globalOperatorListIndex);
    }

    function _registerOperatorWithCoordinator(address operator, bytes calldata quorumNumbers, BN254.G1Point memory pubkey) internal {
        // check that the sender is not already registered
        require(_operators[operator].status != OperatorStatus.REGISTERED, "BLSIndexRegistryCoordinator: operator already registered");

        // get the quorum bitmap from the quorum numbers
        uint256 quorumBitmap = BytesArrayBitmaps.orderedBytesArrayToBitmap_Yul(quorumNumbers);
        require(quorumBitmap != 0, "BLSIndexRegistryCoordinator: quorumBitmap cannot be 0");

        // register the operator with the BLSPubkeyRegistry and get the operatorId (in this case, the pubkeyHash) back
        bytes32 operatorId = blsPubkeyRegistry.registerOperator(operator, quorumNumbers, pubkey);

        // register the operator with the IndexRegistry
        indexRegistry.registerOperator(operatorId, quorumNumbers);

        // register the operator with the StakeRegistry
        _registerOperator(operator, operatorId, quorumNumbers);

        // set the operatorId to quorum bitmap mapping
        operatorIdToQuorumBitmap[operatorId] = quorumBitmap;

        // set the operator struct
        _operators[operator] = Operator({
            operatorId: operatorId,
            fromTaskNumber: serviceManager.taskNumber(),
            status: OperatorStatus.REGISTERED
        });
    }

    function _deregisterOperatorWithCoordinator(address operator, BN254.G1Point memory pubkey, bytes32[] memory operatorIdsToSwap, uint32 globalOperatorListIndex) internal {
        require(_operators[operator].status == OperatorStatus.REGISTERED, "BLSIndexRegistryCoordinator._deregisterOperator: operator is not registered");

        // get the operatorId of the operator
        bytes32 operatorId = _operators[operator].operatorId;
        require(operatorId == pubkey.hashG1Point(), "BLSIndexRegistryCoordinator._deregisterOperator: operatorId does not match pubkey hash");

        // get the quorumNumbers of the operator
        bytes memory quorumNumbers = BytesArrayBitmaps.bitmapToBytesArray(operatorIdToQuorumBitmap[operatorId]);
        
        // deregister the operator from the BLSPubkeyRegistry
        blsPubkeyRegistry.deregisterOperator(operator, quorumNumbers, pubkey);

        // deregister the operator from the IndexRegistry
        indexRegistry.deregisterOperator(operatorId, quorumNumbers, operatorIdsToSwap, globalOperatorListIndex);

        // deregister the operator from the StakeRegistry
        _deregisterOperator(operator, operatorId, quorumNumbers);

        // set the status of the operator to DEREGISTERED
        _operators[operator].status = OperatorStatus.DEREGISTERED;
    }
}