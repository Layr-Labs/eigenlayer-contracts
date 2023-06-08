// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IRegistryCoordinator.sol";
import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IIndexRegistry.sol";

import "../libraries/BytesArrayBitmaps.sol";

import "./StakeRegistry.sol";

contract BLSIndexRegistryCoordinator is StakeRegistry, IRegistryCoordinator {
    using BN254 for BN254.G1Point;

    /// @notice the BLS Pubkey Registry contract that will keep track of operators' BLS public keys
    IBLSPubkeyRegistry public immutable blsPubkeyRegistry;
    /// @notice the Index Registry contract that will keep track of operators' indexes
    IIndexRegistry public immutable indexRegistry;
    /// @notice the mapping from operator's operatorId to the bitmap of quorums they are registered for
    mapping(bytes32 => uint256) public operatorIdToQuorumBitmap;
    /// @notice the mapping from operator's address to the operator struct
    mapping(address => Operator) public operators;
    /// @notice the dynamic length array of the registries this coordinator is coordinating
    address[] public registries;

    constructor(
        IBLSPubkeyRegistry _blsPubkeyRegistry,
        IIndexRegistry _indexRegistry,
        IStrategyManager _strategyManager,
        IServiceManager _serviceManager
    ) StakeRegistry(IRegistryCoordinator(address(this)), _strategyManager, _serviceManager) {
        blsPubkeyRegistry = _blsPubkeyRegistry;
        indexRegistry = _indexRegistry;
    }

    function initialize(
        uint96[] memory _minimumStakeForQuorum,
        StrategyAndWeightingMultiplier[][] memory _quorumStrategiesConsideredAndMultipliers
    ) external override initializer {
        // the stake registry is this contract itself
        registries.push(address(this));
        registries.push(address(blsPubkeyRegistry));
        registries.push(address(indexRegistry));

        StakeRegistry._initialize(_minimumStakeForQuorum, _quorumStrategiesConsideredAndMultipliers);
    }

    /// @notice Returns task number from when `operator` has been registered.
    function getFromTaskNumberForOperator(address operator) external view returns (uint32) {
        return operators[operator].fromTaskNumber;
    }

    /// @notice Returns the operatorId for the given `operator`
    function getOperatorId(address operator) external view returns (bytes32) {
        return operators[operator].operatorId;
    }

    /// @notice Returns the number of registries
    function numRegistries() external view returns (uint256) {
        return registries.length;
    }

    function registerOperator(address, bytes32, bytes calldata) external override pure {
        revert("BLSIndexRegistryCoordinator.registerOperator: cannot use overrided StakeRegistry.registerOperator on BLSIndexRegistryCoordinator");
    }

    /**
     * @notice Registers the operator with the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registering for
     * @param registrationData is the data that is decoded to get the operator's registration information
     */
    function registerOperator(bytes calldata quorumNumbers, bytes calldata registrationData) external {
        // get the operator's BLS public key
        BN254.G1Point memory pubkey = abi.decode(registrationData, (BN254.G1Point));
        // call internal function to register the operator
        _registerOperator(msg.sender, quorumNumbers, pubkey);
    }

    /**
     * @notice Registers the operator with the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registering for
     * @param pubkey is the BLS public key of the operator
     */
    function registerOperator(bytes calldata quorumNumbers, BN254.G1Point memory pubkey) external {
        _registerOperator(msg.sender, quorumNumbers, pubkey);
    }

    function deregisterOperator(address, bytes32, bytes calldata) external override pure {
        revert("BLSIndexRegistryCoordinator.deregisterOperator: cannot use overrided StakeRegistry.deregisterOperator on BLSIndexRegistryCoordinator");
    }
    
    /**
     * @notice Deregisters the operator from the middleware
     * @param deregistrationData is the the data that is decoded to get the operator's deregisteration information
     */
    function deregisterOperator(bytes calldata deregistrationData) external {
        // get the operator's deregisteration information
        (BN254.G1Point memory pubkey, uint32[] memory quorumToOperatorListIndexes, uint32 globalOperatorListIndex) 
            = abi.decode(deregistrationData, (BN254.G1Point, uint32[], uint32));
        // call internal function to deregister the operator
        _deregisterOperator(msg.sender, pubkey, quorumToOperatorListIndexes, globalOperatorListIndex);
    }

    /**
     * @notice Deregisters the operator from the middleware
     * @param pubkey is the BLS public key of the operator
     * @param quorumToOperatorListIndexes is the list of the operator's indexes in the quorums they are registered for in the IndexRegistry
     * @param globalOperatorListIndex is the operator's index in the global operator list in the IndexRegistry
     */
    function deregisterOperator(BN254.G1Point memory pubkey, uint32[] memory quorumToOperatorListIndexes, uint32 globalOperatorListIndex) external {
        _deregisterOperator(msg.sender, pubkey, quorumToOperatorListIndexes, globalOperatorListIndex);
    }

    function _registerOperator(address operator, bytes calldata quorumNumbers, BN254.G1Point memory pubkey) internal {
        // TODO: check that the sender is not already registered

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
        operators[operator] = Operator({
            operatorId: operatorId,
            fromTaskNumber: serviceManager.taskNumber(),
            status: OperatorStatus.REGISTERED
        });

        // emit event
        emit OperatorRegistered(operator, operatorId);
    }

    function _deregisterOperator(address operator, BN254.G1Point memory pubkey, uint32[] memory quorumToOperatorListIndexes, uint32 globalOperatorListIndex) internal {
        require(operators[operator].status == OperatorStatus.REGISTERED, "BLSIndexRegistryCoordinator._deregisterOperator: operator is not registered");

        // get the operatorId of the operator
        bytes32 operatorId = operators[operator].operatorId;
        require(operatorId == pubkey.hashG1Point(), "BLSIndexRegistryCoordinator._deregisterOperator: operatorId does not match pubkey hash");

        // get the quorumNumbers of the operator
        bytes memory quorumNumbers = BytesArrayBitmaps.bitmapToBytesArray(operatorIdToQuorumBitmap[operatorId]);
        
        // deregister the operator from the BLSPubkeyRegistry
        blsPubkeyRegistry.deregisterOperator(operator, quorumNumbers, pubkey);

        // deregister the operator from the IndexRegistry
        indexRegistry.deregisterOperator(operatorId, quorumNumbers, quorumToOperatorListIndexes, globalOperatorListIndex);

        // deregister the operator from the StakeRegistry
        _deregisterOperator(operator, operatorId, quorumNumbers);

        // set the status of the operator to DEREGISTERED
        operators[operator].status = OperatorStatus.DEREGISTERED;

        // emit event
        emit OperatorDeregistered(operator, operatorId);
    }
}