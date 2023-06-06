// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/middleware/StakeRegistry.sol";

// wrapper around the StakeRegistry contract that exposes the internal functions for unit testing.
contract StakeRegistryHarness is StakeRegistry {
    mapping(uint8 => mapping(address => uint96)) public weightOfOperatorForQuorum;

    constructor(
        IRegistryCoordinator _registryCoordinator,
        IStrategyManager _strategyManager,
        IServiceManager _serviceManager
    ) StakeRegistry(_registryCoordinator, _strategyManager, _serviceManager) {}

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

    // mocked function so we can set this arbitrarily without having to mock other elements
    function weightOfOperator(uint8 quorumNumber, address operator) public override view returns(uint96) {
        return weightOfOperatorForQuorum[quorumNumber][operator];
    }

    // mocked function so we can set this arbitrarily without having to mock other elements
    function setOperatorWeight(uint8 quorumNumber, address operator, uint96 weight) external {
        weightOfOperatorForQuorum[quorumNumber][operator] = weight;
    }
}