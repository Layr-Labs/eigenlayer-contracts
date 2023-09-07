// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./StakeRegistryStorage.sol";

abstract contract StakeRegistryBase is StakeRegistryStorage {

    function _beforeRegisterOperator(address operator, bytes32 operatorId, bytes memory quorumNumbers) internal virtual;

    function _afterRegisterOperator(address operator, bytes32 operatorId, bytes memory quorumNumbers) internal virtual;
    
    function _beforeDeregisterOperator(bytes32 operatorId, bytes memory quorumNumbers) internal virtual;

    function _afterDeregisterOperator(bytes32 operatorId, bytes memory quorumNumbers) internal virtual;

    function _registerOperator(address operator, bytes32 operatorId, bytes memory quorumNumbers) internal virtual;

    function _deregisterOperator(bytes32 operatorId, bytes memory quorumNumbers) internal virtual;

}
