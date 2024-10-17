// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "src/contracts/interfaces/IStrategy.sol";
import "src/contracts/interfaces/IAVSDirectory.sol";

contract AVSDirectoryMock is Test {
    receive() external payable {}
    fallback() external payable {}

    mapping(bytes32 => bool) _isOperatorSlashable;
    mapping(bytes32 => bool) _isOperatorSetBatch;
    mapping(bytes32 => bool) _isOperatorSetStrategy;
    mapping(bytes32 => bool) _isOperatorSetStrategyBatch;
    
    // Getters

    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) public virtual view returns (bool) {
        return _isOperatorSlashable[keccak256(abi.encode(operator, operatorSet))];
    }

    function isOperatorSetBatch(OperatorSet[] memory operatorSets) public virtual view returns (bool) {
        return _isOperatorSetBatch[keccak256(abi.encode(operatorSets))];
    }

    function isOperatorSetStrategy(OperatorSet calldata operatorSet, IStrategy strategy) external view returns (bool) {
        return _isOperatorSetStrategy[keccak256(abi.encode(operatorSet, strategy))];
    }

    function isOperatorSetStrategyBatch(
        OperatorSet calldata operatorSet,
        IStrategy[] calldata strategies
    ) external view returns (bool) {
        return _isOperatorSetStrategyBatch[keccak256(abi.encode(operatorSet, strategies))];
    }

    // Setters

    function setIsOperatorSlashable(
        address operator, 
        OperatorSet memory operatorSet, 
        bool value
    ) public virtual {
        _isOperatorSlashable[keccak256(abi.encode(operator, operatorSet))] = value;
    }

    function setIsOperatorSlashable(
        address operator,
        address avs,
        uint32 operatorSetId,
        bool value
    ) public virtual {
        OperatorSet memory operatorSet = OperatorSet({
            avs: avs,
            operatorSetId: operatorSetId
        });
        setIsOperatorSlashable(operator, operatorSet, value);
    }

    function setIsOperatorSetBatch(OperatorSet[] memory operatorSets, bool value) public virtual {
        _isOperatorSetBatch[keccak256(abi.encode(operatorSets))] = value;
    }

    function setIsOperatorSetStrategy(OperatorSet memory operatorSet, IStrategy strategy, bool value) public virtual {
        _isOperatorSetStrategy[keccak256(abi.encode(operatorSet, strategy))] = value;
    }

    function setIsOperatorSetStrategyBatch(OperatorSet memory operatorSet, IStrategy[] memory strategies, bool value) public virtual {
        _isOperatorSetStrategyBatch[keccak256(abi.encode(operatorSet, strategies))] = value;
    }
}