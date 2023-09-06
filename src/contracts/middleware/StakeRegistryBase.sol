// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

abstract contract StakeRegistryBase {

    function _beforeRegisterOperator(bytes32 operatorId, bytes memory quorumNumbers) internal virtual;

    function _afterRegisterOperator(bytes32 operatorId, bytes memory quorumNumbers) internal virtual;
    
    function _beforeDeregisterOperator(bytes32 operatorId, bytes memory quorumNumbers) internal virtual;

    function _afterDeregisterOperator(bytes32 operatorId, bytes memory quorumNumbers) internal virtual;


    function _registerOperator(address operator, bytes32 operatorId, bytes memory quorumNumbers) internal virtual {
        _beforeRegisterOperator(operatorId, quorumNumbers);

        _registerOperator(operator, operatorId, quorumNumbers);

        _afterRegisterOperator(operatorId, quorumNumbers);
    }

    function _deregisterOperator(bytes32 operatorId, bytes memory quorumNumbers) internal virtual {
        _beforeDeregisterOperator(operatorId, quorumNumbers);

        _deregisterOperator(operatorId, quorumNumbers);

        _afterDeregisterOperator(operatorId, quorumNumbers);
    }
}
