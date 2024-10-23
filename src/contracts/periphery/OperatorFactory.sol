// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity >=0.8.0;

import "./Operator.sol";
import "@openzeppelin/contracts/utils/Create2.sol";

contract OperatorFactory {
    function createOperator() external returns (Operator) {
        return Operator(Create2.deploy(0, bytes32(uint256(uint160(msg.sender))), type(Operator).creationCode));
    }

    function getOperatorContract(
        address owner
    ) external view returns (Operator) {
        return
            Operator(Create2.computeAddress(bytes32(uint256(uint160(owner))), keccak256(type(Operator).creationCode)));
    }
}
