// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/middleware/RegistryBase.sol";

import "forge-std/Test.sol";

contract MiddlewareVoteWeigherMock is RegistryBase {
    uint8 _NUMBER_OF_QUORUMS = 2;

    constructor(
        IDelegationManager _delegation,
        IStrategyManager _strategyManager,
        IServiceManager _serviceManager
    )
    RegistryBase(_strategyManager, _serviceManager, _NUMBER_OF_QUORUMS)
    {}

    function initialize(
        uint256[] memory _quorumBips,
        StrategyAndWeightingMultiplier[] memory _firstQuorumStrategiesConsideredAndMultipliers,
        StrategyAndWeightingMultiplier[] memory _secondQuorumStrategiesConsideredAndMultipliers
    ) external initializer {
        VoteWeigherBase._initialize(_quorumBips);

        // push an empty OperatorStake struct to the total stake history to record starting with zero stake
        OperatorStake memory _totalStake;
        totalStakeHistory.push(_totalStake);

        // push an empty OperatorIndex struct to the total operators history to record starting with zero operators
        OperatorIndex memory _totalOperators;
        totalOperatorsHistory.push(_totalOperators);

        _addStrategiesConsideredAndMultipliers(0, _firstQuorumStrategiesConsideredAndMultipliers);
        _addStrategiesConsideredAndMultipliers(1, _secondQuorumStrategiesConsideredAndMultipliers);
    }

    function registerOperator(address operator, uint32 serveUntil) public {        
        require(slasher.canSlash(operator, address(serviceManager)), "Not opted into slashing");
        serviceManager.recordFirstStakeUpdate(operator, serveUntil);

    }

    function deregisterOperator(address operator) public {
        uint32 latestServeUntilBlock = serviceManager.latestServeUntilBlock();
        serviceManager.recordLastStakeUpdateAndRevokeSlashingAbility(operator, latestServeUntilBlock);
    }

    function propagateStakeUpdate(address operator, uint32 blockNumber, uint256 prevElement) external {
        uint32 serveUntilBlock = serviceManager.latestServeUntilBlock();
        serviceManager.recordStakeUpdate(operator, blockNumber, serveUntilBlock, prevElement);
    }
}