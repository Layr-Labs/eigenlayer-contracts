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
import "src/contracts/interfaces/IAVSRegistrar.sol";

interface IAVSDeployer {
    function allocationManager() external view returns (AllocationManager);
    function strategyFactory() external view returns (StrategyFactory);
    function timeMachine() external view returns (TimeMachine);
}

contract AVS is Logger, IAllocationManagerTypes, IAVSRegistrar {
    using print for *;
    using SingleItemArrayLib for *;

    AllocationManager immutable allocationManager;
    StrategyFactory immutable strategyFactory;
    TimeMachine immutable timeMachine;
    string _NAME;

    uint32 totalOperatorSets;

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
    /// AllocationManager
    /// -----------------------------------------------------------------------

    function createOperatorSets(
        IStrategy[][] memory strategies
    ) public createSnapshot returns (OperatorSet[] memory operatorSets) {
        print.method("createOperatorSets");

        uint256 len = strategies.length;
        CreateSetParams[] memory p = new CreateSetParams[](len);
        operatorSets = new OperatorSet[](len);

        for (uint256 i; i < len; ++i) {
            p[i] = CreateSetParams({operatorSetId: totalOperatorSets++, strategies: strategies[i]});

            operatorSets[i] = OperatorSet(address(this), p[i].operatorSetId);
        }

        print.createOperatorSets(p);
        allocationManager.createOperatorSets(address(this), p);
        print.gasUsed();
    }

    function createOperatorSet(
        IStrategy[] memory strategies
    ) public createSnapshot returns (OperatorSet memory operatorSet) {
        print.method("createOperatorSets");

        operatorSet = OperatorSet(address(this), totalOperatorSets++);

        CreateSetParams[] memory p = CreateSetParams({operatorSetId: operatorSet.id, strategies: strategies}).toArray();

        print.createOperatorSets(p);
        allocationManager.createOperatorSets(address(this), p);
        print.gasUsed();
    }

    function slashOperator(User operator, uint32 operatorSetId, uint256 wadToSlash) public createSnapshot {
        print.method(
            "slashOperator",
            string.concat(
                "{operator: ",
                operator.NAME_COLORED(),
                ", operatorSetId: ",
                cheats.toString(operatorSetId),
                ", wadToSlash: ",
                wadToSlash.asWad(),
                "}"
            )
        );

        SlashingParams memory p = SlashingParams({
            operator: address(operator),
            operatorSetId: operatorSetId,
            wadToSlash: wadToSlash,
            description: "bad operator"
        });

        allocationManager.slashOperator(address(this), p);
        print.gasUsed();
    }

    function deregisterFromOperatorSets(User operator, uint32[] memory operatorSetIds) public createSnapshot {
        print.method("deregisterFromOperatorSets");

        DeregisterParams memory p =
            DeregisterParams({operator: address(operator), avs: address(this), operatorSetIds: operatorSetIds});

        print.deregisterFromOperatorSets(p);
        allocationManager.deregisterFromOperatorSets(p);
        print.gasUsed();
    }

    function setAVSRegistrar(
        IAVSRegistrar registrar
    ) public createSnapshot {
        print.method("setAVSRegistrar");
        console.log("Setting AVS registrar to: %s", address(registrar));
        allocationManager.setAVSRegistrar(address(this), registrar);
        print.gasUsed();
    }

    function addStrategiesToOperatorSet(uint32 operatorSetId, IStrategy[] memory strategies) public createSnapshot {
        print.method("addStrategiesToOperatorSet");

        console.log("Adding strategies to operator set: %d", operatorSetId);

        for (uint256 i; i < strategies.length; ++i) {
            console.log("   strategy: %s", address(strategies[i]));
        }

        allocationManager.addStrategiesToOperatorSet(address(this), operatorSetId, strategies);
        print.gasUsed();
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

        allocationManager.removeStrategiesFromOperatorSet(address(this), operatorSetId, strategies);
        print.gasUsed();
    }

    /// -----------------------------------------------------------------------
    /// IAVSRegistrar
    /// -----------------------------------------------------------------------

    function registerOperator(
        address operator,
        uint32[] calldata operatorSetIds,
        bytes calldata data
    ) external override {}

    function deregisterOperator(address operator, uint32[] calldata operatorSetIds) external override {}
}
