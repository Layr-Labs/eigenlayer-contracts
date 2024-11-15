// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/core/AllocationManager.sol";
import "src/contracts/strategies/StrategyFactory.sol";

import "src/test/mocks/ERC20Mock.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/TimeMachine.t.sol";
import "src/test/integration/utils/PrintUtils.t.sol";

import "src/test/utils/SingleItemArrayLib.sol";

interface IAVSDeployer {
    function allocationManager() external view returns (AllocationManager);
    function strategyFactory() external view returns (StrategyFactory);
    function timeMachine() external view returns (TimeMachine);
}

contract AVS is PrintUtils, IAllocationManagerTypes {
    using SingleItemArrayLib for *;

    Vm cheats = Vm(VM_ADDRESS);
    AllocationManager immutable allocationManager;
    StrategyFactory immutable strategyFactory;
    TimeMachine immutable timeMachine;
    string _NAME;

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
        _logM("createOperatorSets");

        CreateSetParams[] memory p = new CreateSetParams[](numOpSets);
        operatorSetIds = new uint32[](numOpSets);
        strategies = new IStrategy[][](numOpSets);

        for (uint256 i = 0; i < numOpSets; ++i) {
            p[i].operatorSetId = uint32(cheats.randomUint({bits: 32}));
            for (uint256 j = 0; j < numStrategies; ++j) {
                IERC20 token = IERC20(new ERC20Mock());
                p[i].strategies[j] = strategyFactory.deployNewStrategy(token);
            }

            operatorSetIds[i] = p[i].operatorSetId;
            strategies[i] = p[i].strategies;
        }

        _logCreateOperatorSets(p);

        allocationManager.createOperatorSets(p);
    }

    function createOperatorSet(IStrategy[] memory strategies) public createSnapshot returns (uint32 operatorSetId) {
        _logM("createOperatorSets");

        CreateSetParams[] memory p = CreateSetParams({
            operatorSetId: operatorSetId,
            strategies: strategies
        }).toArray();

        _logCreateOperatorSets(p);

        allocationManager.createOperatorSets(p);
    }

    function slashOperator(User operator, uint32 operatorSetId, uint256 wadToSlash) public createSnapshot {
        _logM("slashOperator");

        SlashingParams memory p = SlashingParams({
            operator: address(operator),
            operatorSetId: operatorSetId,
            wadToSlash: wadToSlash,
            description: "bad operator"
        });

        _logSlashOperator(p);

        allocationManager.slashOperator(p);
    }

    function slashOperator(User operator, uint32 operatorSetId) public {
        slashOperator(operator, operatorSetId, cheats.randomUint(0.1 ether, 1 ether));
    }

    function deregisterFromOperatorSets(User operator, uint32[] memory operatorSetIds) public createSnapshot {
        _logM("deregisterFromOperatorSets");

        DeregisterParams memory p = DeregisterParams({
            operator: address(operator),
            avs: address(this),
            operatorSetIds: operatorSetIds
        });

        _logDeregisterFromOperatorSets(p);

        allocationManager.deregisterFromOperatorSets(p);
    }

    function setAVSRegistrar(
        IAVSRegistrar registrar
    ) public createSnapshot {
        _logM("setAVSRegistrar");

        console.log("Setting AVS registrar to: %s", address(registrar));

        allocationManager.setAVSRegistrar(registrar);
    }

    function updateAVSMetadataURI(
        string memory uri
    ) public createSnapshot {
        _logM("updateAVSMetadataURI");

        console.log("Updating AVS metadata URI to: %s", uri);

        allocationManager.updateAVSMetadataURI(uri);
    }

    function addStrategiesToOperatorSet(uint32 operatorSetId, IStrategy[] memory strategies) public createSnapshot {
        _logM("addStrategiesToOperatorSet");

        console.log("Adding strategies to operator set: %d", operatorSetId);

        for (uint256 i = 0; i < strategies.length; ++i) {
            console.log("   strategy: %s", address(strategies[i]));
        }

        allocationManager.addStrategiesToOperatorSet(operatorSetId, strategies);
    }

    function removeStrategiesFromOperatorSet(
        uint32 operatorSetId,
        IStrategy[] memory strategies
    ) public createSnapshot {
        _logM("removeStrategiesFromOperatorSet");

        console.log("Removing strategies from operator set: %d", operatorSetId);

        for (uint256 i = 0; i < strategies.length; ++i) {
            console.log("   strategy: %s", address(strategies[i]));
        }

        allocationManager.removeStrategiesFromOperatorSet(operatorSetId, strategies);
    }

    /// -----------------------------------------------------------------------
    /// Logging
    /// -----------------------------------------------------------------------

    function _logSlashOperator(
        SlashingParams memory p
    ) private pure {
        console.log("Slashing operator: %s", address(p.operator));
        console.log("   operatorSetId: %d", p.operatorSetId);
        console.log("   wadToSlash: %d", p.wadToSlash);
        console.log("   description: %s", p.description);
    }

    function _logCreateOperatorSets(
        CreateSetParams[] memory p
    ) private pure {
        console.log("Creating operator sets:");
        for (uint256 i = 0; i < p.length; ++i) {
            console.log("   Creating operator set: %d", p[i].operatorSetId);
            for (uint256 j = 0; j < p[i].strategies.length; j++) {
                console.log("       strategy: %s", address(p[i].strategies[j]));
            }
        }
    }

    function _logDeregisterFromOperatorSets(
        DeregisterParams memory p
    ) private pure {
        console.log("Deregistering operator: %s", address(p.operator));
        console.log("   from operator sets:");
        for (uint256 i = 0; i < p.operatorSetIds.length; ++i) {
            console.log("       operatorSetId: %d", p.operatorSetIds[i]);
        }
    }
}
