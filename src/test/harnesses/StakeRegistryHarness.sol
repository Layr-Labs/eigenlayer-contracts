// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/middleware/StakeRegistry.sol";

// wrapper around the StakeRegistry contract that exposes the internal functions for unit testing.
contract StakeRegistryHarness is StakeRegistry {
    constructor(
        IStrategyManager _strategyManager,
        IServiceManager _serviceManager
    ) StakeRegistry(_strategyManager, _serviceManager) {}

    function recordOperatorStakeUpdate(bytes32 operatorId, uint8 quorumNumber, OperatorStakeUpdate memory operatorStakeUpdate) external returns(uint96) {
        return _recordOperatorStakeUpdate(operatorId, quorumNumber, operatorStakeUpdate);
    }

    function updateOperatorStake(address operator, bytes32 operatorId, uint8 quorumNumber) external returns (uint96, uint96) {
        return _updateOperatorStake(operator, operatorId, quorumNumber);
    }

    function registerStake(address operator, bytes32 operatorId, bytes memory quorumNumbers) external {
        _registerStake(operator, operatorId, quorumNumbers);
    }

    function removeOperatorStake(bytes32 operatorId, bytes memory quorumNumbers) external {
        _removeOperatorStake(operatorId, quorumNumbers);
    }
}