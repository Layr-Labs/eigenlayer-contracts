// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/middleware/StakeRegistry.sol";

// wrapper around the StakeRegistry contract that exposes the internal functions for unit testing.
contract StakeRegistryHarness is StakeRegistry {
    mapping(uint8 => mapping(address => uint96)) private _weightOfOperatorForQuorum;

    constructor(
        IRegistryCoordinator _registryCoordinator,
        IStrategyManager _strategyManager,
        IServiceManager _serviceManager
    ) StakeRegistry(_registryCoordinator, _strategyManager, _serviceManager) {
    }

    function recordOperatorStakeUpdate(bytes32 operatorId, uint8 quorumNumber, OperatorStakeUpdate memory operatorStakeUpdate) external returns(uint96) {
        return _recordOperatorStakeUpdate(operatorId, quorumNumber, operatorStakeUpdate);
    }

    function updateOperatorStake(address operator, bytes32 operatorId, uint8 quorumNumber) external returns (uint96, uint96) {
        return _updateOperatorStake(operator, operatorId, quorumNumber);
    }

    function recordTotalStakeUpdate(uint8 quorumNumber, OperatorStakeUpdate memory totalStakeUpdate) external {
        _recordTotalStakeUpdate(quorumNumber, totalStakeUpdate);
    }

    // mocked function so we can set this arbitrarily without having to mock other elements
    function weightOfOperatorForQuorum(uint8 quorumNumber, address operator) public override view returns(uint96) {
        return _weightOfOperatorForQuorum[quorumNumber][operator];
    }

    // mocked function so we can set this arbitrarily without having to mock other elements
    function setOperatorWeight(uint8 quorumNumber, address operator, uint96 weight) external {
        _weightOfOperatorForQuorum[quorumNumber][operator] = weight;
    }

    // mocked function to register an operator without having to mock other elements
    function registerOperatorNonCoordinator(address operator, bytes32 operatorId, bytes calldata quorumNumbers) external {
        _registerOperator(operator, operatorId, quorumNumbers);
    }
}