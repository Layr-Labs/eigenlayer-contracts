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
    RegistryBase(_strategyManager, _serviceManager)
    {}

    function initialize(
        uint256[] memory _quorumBips,
        uint96[] memory _minimumStakeForQuorums,
        StrategyAndWeightingMultiplier[][] memory _quorumStrategiesConsideredAndMultipliers
    ) public virtual initializer {
        RegistryBase._initialize(
            _quorumBips,
            _minimumStakeForQuorums,
            _quorumStrategiesConsideredAndMultipliers
        );
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

    //  TODO: Fix this
    function getTotalStakeForQuorumFromIndex(uint256 quorumNumber, uint256 index) external view returns (OperatorStakeUpdate memory) {
        return totalStakeHistory[quorumNumber][index];
    }
}