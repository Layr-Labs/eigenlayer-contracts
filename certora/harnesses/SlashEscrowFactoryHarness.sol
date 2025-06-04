// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../src/contracts/core/SlashEscrowFactory.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @title SlashEscrowFactoryHarness
 * @dev Minimal harness contract to expose only the internal state needed for verification
 *      without inventing methods that don't exist in the actual contracts
 */
contract SlashEscrowFactoryHarness is SlashEscrowFactory {
    using OperatorSetLib for *;
    using EnumerableSet for *;
    constructor(
        IAllocationManager _allocationManager,
        IStrategyManager _strategyManager,
        IPauserRegistry _pauserRegistry,
        ISlashEscrow _slashEscrowImplementation,
        string memory _version
    ) SlashEscrowFactory(_allocationManager, _strategyManager, _pauserRegistry, _slashEscrowImplementation, _version ) {}
    
 
    /**
     * @dev Get pending slash IDs count for an operator set
     */
    function getPendingSlashIdsCount(OperatorSet calldata operatorSet) external view returns (uint256) {
        return _pendingSlashIds[operatorSet.key()].length();
    }
    
    /**
     * @dev Check if a slash ID is pending for an operator set
     */
    function isSlashIdPending(OperatorSet calldata operatorSet, uint256 slashId) external view returns (bool) {
        return _pendingSlashIds[operatorSet.key()].contains(slashId);
    }
    
    /**
     * @dev Get pending strategies count for a slash ID
     */
    function getPendingStrategiesCount(OperatorSet calldata operatorSet, uint256 slashId) external view returns (uint256) {
        return _pendingStrategiesForSlashId[operatorSet.key()][slashId].length();
    }
    
    /**
     * @dev Check if a strategy is pending for a slash ID
     */
    function isStrategyPending(OperatorSet calldata operatorSet, uint256 slashId, IStrategy strategy) external view returns (bool) {
        return _pendingStrategiesForSlashId[operatorSet.key()][slashId].contains(address(strategy));
    }
    
    /**
     * @dev Get WAD constant (1e18)
     */
    function WAD() external pure returns (uint256) {
        return 1e18;
    }
    
     /**
     * @dev Create an operator set key for testing
     */
    function createOperatorSetKey(address avs, uint32 operatorSetId) external pure returns (bytes32) {
        return OperatorSetLib.key(OperatorSet(avs, operatorSetId));
    }

     /**
     * @dev Create an operator set key for testing
     */
    function getOperatorSetKey(OperatorSet calldata op) external pure returns (bytes32) {
        return OperatorSetLib.key(op);
    }
    
    /**
     * @dev Decode an operator set key for testing
     */
    function decodeOperatorSetKey(bytes32 operatorSetKey) external pure returns (OperatorSet memory) {
        return OperatorSetLib.decode(operatorSetKey);
    }
}
