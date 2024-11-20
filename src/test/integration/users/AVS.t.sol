// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/core/AllocationManager.sol";
import "src/contracts/strategies/StrategyFactory.sol";

import "src/test/mocks/ERC20Mock.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/TimeMachine.t.sol";
import "src/test/utils/Logger.t.sol";

import "src/test/utils/SingleItemArrayLib.sol";

interface IAVSDeployer {
    function allocationManager() external view returns (AllocationManager);
    function strategyFactory() external view returns (StrategyFactory);
    function timeMachine() external view returns (TimeMachine);
}

contract AVS is Logger, IAllocationManagerTypes {
    using SingleItemArrayLib for *;

    AllocationManager immutable allocationManager;
    StrategyFactory immutable strategyFactory;
    TimeMachine immutable timeMachine;
    string _NAME;

    uint256 operatorSetId;

    constructor(
        string memory name
    ) {
        IAVSDeployer deployer = IAVSDeployer(msg.sender);
        allocationManager = deployer.allocationManager();
        strategyFactory = deployer.strategyFactory();
        timeMachine = deployer.timeMachine();
        _NAME = name;
    }

    modifier createSnapshot() virtual {
        timeMachine.createSnapshot();
        _;
    }

    receive() external payable {}

    function NAME() public view override returns (string memory) {
        return _NAME;
    }

    /// -----------------------------------------------------------------------
    /// AllocationManager Interactions
    /// -----------------------------------------------------------------------
    
    function createOperatorSets(
        uint256 numOpSets,
        uint256 numStrategies
    ) public createSnapshot returns (uint32[] memory operatorSetIds, IStrategy[][] memory strategies) {
        print.method("createOperatorSets");

        CreateSetParams[] memory p = new CreateSetParams[](numOpSets);
        operatorSetIds = new uint32[](numOpSets);
        strategies = new IStrategy[][](numOpSets);

        for (uint256 i; i < numOpSets; ++i) {
            p[i].operatorSetId = operatorSetId++;
            for (uint256 j; j < numStrategies; ++j) {
                IERC20 token = IERC20(new ERC20Mock());
                p[i].strategies[j] = strategyFactory.deployNewStrategy(token);
            }

            operatorSetIds[i] = p[i].operatorSetId;
            strategies[i] = p[i].strategies;
        }

        print.createOperatorSets(p);
        allocationManager.createOperatorSets(p);
    }

    function createOperatorSet(IStrategy[] memory strategies) public createSnapshot returns (uint32 operatorSetId) {
        print.method("createOperatorSets");

        CreateSetParams[] memory p = CreateSetParams({
            operatorSetId: operatorSetId,
            strategies: strategies
        }).toArray();

        print.createOperatorSets(p);
        allocationManager.createOperatorSets(p);
    }

    function slashOperator(User operator, uint32 operatorSetId, uint256 wadToSlash) public createSnapshot {
        print.method("slashOperator");

        SlashingParams memory p = SlashingParams({
            operator: address(operator),
            operatorSetId: operatorSetId,
            wadToSlash: wadToSlash,
            description: "bad operator"
        });

        print.slashOperator(p);
        allocationManager.slashOperator(p);
    }

    function deregisterFromOperatorSets(User operator, uint32[] memory operatorSetIds) public createSnapshot {
        print.method("deregisterFromOperatorSets");

        DeregisterParams memory p = DeregisterParams({
            operator: address(operator),
            avs: address(this),
            operatorSetIds: operatorSetIds
        });

        print.deregisterFromOperatorSets(p);
        allocationManager.deregisterFromOperatorSets(p);
    }

    function setAVSRegistrar(
        IAVSRegistrar registrar
    ) public createSnapshot {
        print.method("setAVSRegistrar");
        console.log("Setting AVS registrar to: %s", address(registrar));
        allocationManager.setAVSRegistrar(registrar);
    }

    function addStrategiesToOperatorSet(uint32 operatorSetId, IStrategy[] memory strategies) public createSnapshot {
        print.method("addStrategiesToOperatorSet");

        console.log("Adding strategies to operator set: %d", operatorSetId);

        for (uint256 i; i < strategies.length; ++i) {
            console.log("   strategy: %s", address(strategies[i]));
        }

        allocationManager.addStrategiesToOperatorSet(operatorSetId, strategies);
    }

    function removeStrategiesFromOperatorSet(
        uint32 operatorSetId,
        IStrategy[] memory strategies
    ) public createSnapshot {
        print.method("removeStrategiesFromOperatorSet");

        console.log("Removing strategies from operator set: %d", operatorSetId);

        for (uint256 i; i < strategies.length; ++i) {
            console.log("   strategy: %s", address(strategies[i]));
        }

        allocationManager.removeStrategiesFromOperatorSet(operatorSetId, strategies);
    }
}
