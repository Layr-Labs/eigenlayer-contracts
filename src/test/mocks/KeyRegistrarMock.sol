// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

contract KeyRegistrarMock {
    using OperatorSetLib for OperatorSet;

    mapping(bytes32 => IKeyRegistrarTypes.CurveType) public operatorSetCurveTypes;

    function getOperatorSetCurveType(OperatorSet memory operatorSet) external view returns (IKeyRegistrarTypes.CurveType) {
        bytes32 key = keccak256(abi.encode(operatorSet.avs, operatorSet.id));
        return operatorSetCurveTypes[key];
    }

    // Helper function for testing
    function setOperatorSetCurveType(OperatorSet memory operatorSet, IKeyRegistrarTypes.CurveType curveType) external {
        bytes32 key = keccak256(abi.encode(operatorSet.avs, operatorSet.id));
        operatorSetCurveTypes[key] = curveType;
    }
}
