// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "src/contracts/interfaces/IAVSDirectory.sol";

contract AVSDirectoryMock is Test {
    receive() external payable {}
    fallback() external payable {}

    mapping(address => mapping(bytes32 => bool)) public _isOperatorSlashable;
    mapping(bytes32 => bool) public _isOperatorSetBatch;

    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) public virtual view returns (bool) {
        return _isOperatorSlashable[operator][bytes32(abi.encode(operatorSet))];
    }

    function isOperatorSetBatch(OperatorSet[] memory operatorSets) public virtual view returns (bool) {
        return _isOperatorSetBatch[keccak256(abi.encode(operatorSets))];
    }

    function setIsOperatorSlashable(
        address operator, 
        OperatorSet memory operatorSet, 
        bool value
    ) public virtual {
        _isOperatorSlashable[operator][bytes32(abi.encode(operatorSet))] = value;
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
}