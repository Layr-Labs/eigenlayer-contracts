// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/interfaces/IPermissionController.sol";
import "src/contracts/interfaces/IStrategyFactory.sol";
import "src/contracts/interfaces/ISlashEscrowFactory.sol";

import "src/test/mocks/ERC20Mock.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/TimeMachine.t.sol";
import "src/test/utils/Logger.t.sol";

import "src/test/utils/ArrayLib.sol";
import "src/contracts/interfaces/IAVSRegistrar.sol";

interface IAVSDeployer {
    function delegationManager() external view returns (IDelegationManager);
    function allocationManager() external view returns (IAllocationManager);
    function strategyManager() external view returns (IStrategyManager);
    function strategyFactory() external view returns (IStrategyFactory);
    function permissionController() external view returns (IPermissionController);
    function slashEscrowFactory() external view returns (ISlashEscrowFactory);
    function timeMachine() external view returns (TimeMachine);
}

contract AVS is Logger, IAllocationManagerTypes, IAVSRegistrar {
    using print for *;
    using ArrayLib for *;

    IStrategy constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    // TODO: fix later for same reason as User.t.sol
    IAllocationManager immutable allocationManager;
    IPermissionController immutable permissionController;
    IDelegationManager immutable delegationManager;
    IStrategyManager immutable strategyManager;
    IStrategyFactory immutable strategyFactory;
    ISlashEscrowFactory immutable slashEscrowFactory;
    TimeMachine immutable timeMachine;
    string _NAME;

    uint32 totalOperatorSets;

    constructor(string memory name) {
        IAVSDeployer deployer = IAVSDeployer(msg.sender);
        allocationManager = deployer.allocationManager();
        permissionController = deployer.permissionController();
        delegationManager = deployer.delegationManager();
        strategyManager = deployer.strategyManager();
        strategyFactory = deployer.strategyFactory();
        slashEscrowFactory = deployer.slashEscrowFactory();
        timeMachine = deployer.timeMachine();
        _NAME = name;
        cheats.label(address(this), NAME_COLORED());
    }

    modifier createSnapshot() virtual {
        timeMachine.createSnapshot();
        _;
    }

    receive() external payable {}

    function NAME() public view override returns (string memory) {
        return _NAME;
    }

    function supportsAVS(address) external pure override returns (bool) {
        return true;
    }

    /// -----------------------------------------------------------------------
    /// AllocationManager
    /// -----------------------------------------------------------------------

    function updateAVSMetadataURI(string memory uri) public createSnapshot {
        print.method("updateAVSMetadataURI");

        console.log("Setting AVS metadata URI to: %s", uri);
        _tryPrankAppointee_AllocationManager(IAllocationManager.updateAVSMetadataURI.selector);
        allocationManager.updateAVSMetadataURI(address(this), uri);

        print.gasUsed();
    }

    function createOperatorSets(IStrategy[][] memory strategies) public createSnapshot returns (OperatorSet[] memory operatorSets) {
        print.method("createOperatorSets");

        uint len = strategies.length;
        CreateSetParams[] memory p = new CreateSetParams[](len);
        operatorSets = new OperatorSet[](len);

        for (uint i; i < len; ++i) {
            p[i] = CreateSetParams({operatorSetId: totalOperatorSets++, strategies: strategies[i]});

            operatorSets[i] = OperatorSet(address(this), p[i].operatorSetId);
        }

        print.createOperatorSets(p);

        allocationManager.createOperatorSets(address(this), p);

        print.gasUsed();
    }

    function createRedistributingOperatorSets(IStrategy[][] memory strategies, address[] memory redistributionRecipients)
        public
        createSnapshot
        returns (OperatorSet[] memory operatorSets)
    {
        print.method("createRedistributingOperatorSets");

        uint len = strategies.length;
        CreateSetParams[] memory p = new CreateSetParams[](len);
        operatorSets = new OperatorSet[](len);

        for (uint i; i < len; ++i) {
            p[i] = CreateSetParams({operatorSetId: totalOperatorSets++, strategies: strategies[i]});

            operatorSets[i] = OperatorSet(address(this), p[i].operatorSetId);
        }

        print.createOperatorSets(p);

        allocationManager.createRedistributingOperatorSets(address(this), p, redistributionRecipients);

        print.gasUsed();
    }

    function createOperatorSet(IStrategy[] memory strategies) public createSnapshot returns (OperatorSet memory operatorSet) {
        print.method("createOperatorSets");

        operatorSet = OperatorSet(address(this), totalOperatorSets++);

        CreateSetParams[] memory p = CreateSetParams({operatorSetId: operatorSet.id, strategies: strategies}).toArray();

        print.createOperatorSets(p);
        allocationManager.createOperatorSets(address(this), p);
        print.gasUsed();
    }

    function createRedistributingOperatorSet(IStrategy[] memory strategies, address redistributionRecipient)
        public
        createSnapshot
        returns (OperatorSet memory operatorSet)
    {
        print.method("createRedistributingOperatorSet");

        operatorSet = OperatorSet(address(this), totalOperatorSets++);

        CreateSetParams[] memory p = CreateSetParams({operatorSetId: operatorSet.id, strategies: strategies}).toArray();

        print.createOperatorSets(p);
        allocationManager.createRedistributingOperatorSets(address(this), p, redistributionRecipient.toArray());
        print.gasUsed();
    }

    function slashOperator(SlashingParams memory params) public createSnapshot returns (uint slashId, uint[] memory shares) {
        for (uint i; i < params.strategies.length; ++i) {
            string memory strategyName = params.strategies[i] == beaconChainETHStrategy
                ? "Native ETH"
                : IERC20Metadata(address(params.strategies[i].underlyingToken())).name();

            print.method(
                "slashOperator",
                string.concat(
                    "{operator: ",
                    User(payable(params.operator)).NAME_COLORED(),
                    ", operatorSetId: ",
                    cheats.toString(params.operatorSetId),
                    ", strategy: ",
                    strategyName,
                    ", wadToSlash: ",
                    params.wadsToSlash[i].asWad(),
                    "}"
                )
            );
        }
        cheats.label(address(slashEscrowFactory.getSlashEscrow(OperatorSet(address(this), params.operatorSetId), slashId)), "slashEscrow");
        _tryPrankAppointee_AllocationManager(IAllocationManager.slashOperator.selector);
        (slashId, shares) = allocationManager.slashOperator(address(this), params);
        print.gasUsed();
    }

    function slashOperator(User operator, uint32 operatorSetId, IStrategy[] memory strategies, uint[] memory wadsToSlash)
        public
        createSnapshot
        returns (SlashingParams memory p, uint slashId, uint[] memory shares)
    {
        p = SlashingParams({
            operator: address(operator),
            operatorSetId: operatorSetId,
            strategies: strategies,
            wadsToSlash: wadsToSlash,
            description: "bad operator"
        });

        for (uint i; i < strategies.length; ++i) {
            string memory strategyName =
                strategies[i] == beaconChainETHStrategy ? "Native ETH" : IERC20Metadata(address(strategies[i].underlyingToken())).name();

            print.method(
                "slashOperator",
                string.concat(
                    "{operator: ",
                    operator.NAME_COLORED(),
                    ", operatorSetId: ",
                    cheats.toString(operatorSetId),
                    ", strategy: ",
                    strategyName,
                    ", wadToSlash: ",
                    wadsToSlash[i].asWad(),
                    "}"
                )
            );
        }

        _tryPrankAppointee_AllocationManager(IAllocationManager.slashOperator.selector);
        (slashId, shares) = allocationManager.slashOperator(address(this), p);
        print.gasUsed();
    }

    function clearBurnOrRedistributableShares(OperatorSet memory operatorSet, uint slashId) public createSnapshot {
        print.method("clearBurnOrRedistributableShares");
        strategyManager.clearBurnOrRedistributableShares(operatorSet, slashId);
        print.gasUsed();
    }

    function clearBurnOrRedistributableSharesByStrategy(OperatorSet memory operatorSet, uint slashId, IStrategy strategy)
        public
        createSnapshot
    {
        print.method("clearBurnOrRedistributableSharesByStrategy");
        strategyManager.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId, strategy);
        print.gasUsed();
    }

    function deregisterFromOperatorSets(User operator, uint32[] memory operatorSetIds) public createSnapshot {
        print.method("deregisterFromOperatorSets");

        DeregisterParams memory p = DeregisterParams({operator: address(operator), avs: address(this), operatorSetIds: operatorSetIds});

        print.deregisterFromOperatorSets(p);
        _tryPrankAppointee_AllocationManager(IAllocationManager.deregisterFromOperatorSets.selector);
        allocationManager.deregisterFromOperatorSets(p);
        print.gasUsed();
    }

    function setAVSRegistrar(IAVSRegistrar registrar) public createSnapshot {
        print.method("setAVSRegistrar");
        console.log("Setting AVS registrar to: %s", address(registrar));
        _tryPrankAppointee_AllocationManager(IAllocationManager.setAVSRegistrar.selector);
        allocationManager.setAVSRegistrar(address(this), registrar);
        print.gasUsed();
    }

    function addStrategiesToOperatorSet(uint32 operatorSetId, IStrategy[] memory strategies) public createSnapshot {
        print.method("addStrategiesToOperatorSet");

        console.log("Adding strategies to operator set: %d", operatorSetId);

        for (uint i; i < strategies.length; ++i) {
            console.log("   strategy: %s", address(strategies[i]));
        }
        _tryPrankAppointee_AllocationManager(IAllocationManager.addStrategiesToOperatorSet.selector);
        allocationManager.addStrategiesToOperatorSet(address(this), operatorSetId, strategies);
        print.gasUsed();
    }

    function removeStrategiesFromOperatorSet(uint32 operatorSetId, IStrategy[] memory strategies) public createSnapshot {
        print.method("removeStrategiesFromOperatorSet");

        console.log("Removing strategies from operator set: %d", operatorSetId);

        for (uint i; i < strategies.length; ++i) {
            console.log("   strategy: %s", address(strategies[i]));
        }
        _tryPrankAppointee_AllocationManager(IAllocationManager.removeStrategiesFromOperatorSet.selector);
        allocationManager.removeStrategiesFromOperatorSet(address(this), operatorSetId, strategies);
        print.gasUsed();
    }

    /// -----------------------------------------------------------------------
    /// IAVSRegistrar
    /// -----------------------------------------------------------------------

    function registerOperator(address operator, address avsIdentifier, uint32[] calldata operatorSetIds, bytes calldata data)
        external
        override
    {}

    function deregisterOperator(address operator, address avsIdentifier, uint32[] calldata operatorSetIds) external override {}

    /// -----------------------------------------------------------------------
    /// Internal Helpers
    /// -----------------------------------------------------------------------

    // function allocationManager public view returns (AllocationManager) {
    //     return AllocationManager(address(delegationManager.allocationManager));
    // }

    // function permissionController public view returns (PermissionController) {
    //     return PermissionController(address(delegationManager.permissionController));
    // }

    function _tryPrankAppointee(address target, bytes4 selector) internal {
        address[] memory appointees = permissionController.getAppointees(address(this), target, selector);
        if (appointees.length != 0) cheats.prank(appointees[0]);
    }

    function _tryPrankAppointee_AllocationManager(bytes4 selector) internal {
        return _tryPrankAppointee(address(allocationManager), selector);
    }
}
