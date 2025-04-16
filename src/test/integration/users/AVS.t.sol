// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/test/integration/users/User.t.sol";

contract AVS is Logger, IAllocationManagerTypes {
    using print for *;
    using ArrayLib for *;

    // TODO: fix later for same reason as User.t.sol
    string _NAME;

    uint32 totalOperatorSets;

    ConfigGetters public config;

    constructor(string memory name) {
        config = ConfigGetters(address(msg.sender));
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

    /// -----------------------------------------------------------------------
    /// AllocationManager
    /// -----------------------------------------------------------------------

    function updateAVSMetadataURI(string memory uri) public createSnapshot {
        print.method("updateAVSMetadataURI");

        console.log("Setting AVS metadata URI to: %s", uri);
        _tryPrankAppointee_AllocationManager(IAllocationManager.updateAVSMetadataURI.selector);
        config.allocationManager().updateAVSMetadataURI(address(this), uri);

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

        config.allocationManager().createOperatorSets(address(this), p);

        print.gasUsed();
    }

    function createOperatorSet(IStrategy[] memory strategies) public createSnapshot returns (OperatorSet memory operatorSet) {
        print.method("createOperatorSets");

        operatorSet = OperatorSet(address(this), totalOperatorSets++);

        CreateSetParams[] memory p = CreateSetParams({operatorSetId: operatorSet.id, strategies: strategies}).toArray();

        print.createOperatorSets(p);
        config.allocationManager().createOperatorSets(address(this), p);
        print.gasUsed();
    }

    function slashOperator(SlashingParams memory params) public createSnapshot {
        for (uint i; i < params.strategies.length; ++i) {
            string memory strategyName = params.strategies[i] == BEACONCHAIN_ETH_STRAT
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
                    cheats.toString(params.wadsToSlash[i]),
                    "}"
                )
            );
        }

        _tryPrankAppointee_AllocationManager(IAllocationManager.slashOperator.selector);
        config.allocationManager().slashOperator(address(this), params);
        print.gasUsed();
    }

    function slashOperator(User operator, uint32 operatorSetId, IStrategy[] memory strategies, uint[] memory wadsToSlash)
        public
        createSnapshot
        returns (SlashingParams memory p)
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
                strategies[i] == BEACONCHAIN_ETH_STRAT ? "Native ETH" : IERC20Metadata(address(strategies[i].underlyingToken())).name();

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
                    cheats.toString(wadsToSlash[i]),
                    "}"
                )
            );
        }

        _tryPrankAppointee_AllocationManager(IAllocationManager.slashOperator.selector);
        config.allocationManager().slashOperator(address(this), p);
        print.gasUsed();
    }

    function deregisterFromOperatorSets(User operator, uint32[] memory operatorSetIds) public createSnapshot {
        print.method("deregisterFromOperatorSets");

        DeregisterParams memory p = DeregisterParams({operator: address(operator), avs: address(this), operatorSetIds: operatorSetIds});

        print.deregisterFromOperatorSets(p);
        _tryPrankAppointee_AllocationManager(IAllocationManager.deregisterFromOperatorSets.selector);
        config.allocationManager().deregisterFromOperatorSets(p);
        print.gasUsed();
    }

    function setAVSRegistrar(IAVSRegistrar registrar) public createSnapshot {
        print.method("setAVSRegistrar");
        console.log("Setting AVS registrar to: %s", address(registrar));
        _tryPrankAppointee_AllocationManager(IAllocationManager.setAVSRegistrar.selector);
        config.allocationManager().setAVSRegistrar(address(this), registrar);
        print.gasUsed();
    }

    function addStrategiesToOperatorSet(uint32 operatorSetId, IStrategy[] memory strategies) public createSnapshot {
        print.method("addStrategiesToOperatorSet");

        console.log("Adding strategies to operator set: %d", operatorSetId);

        for (uint i; i < strategies.length; ++i) {
            console.log("   strategy: %s", address(strategies[i]));
        }
        _tryPrankAppointee_AllocationManager(IAllocationManager.addStrategiesToOperatorSet.selector);
        config.allocationManager().addStrategiesToOperatorSet(address(this), operatorSetId, strategies);
        print.gasUsed();
    }

    function removeStrategiesFromOperatorSet(uint32 operatorSetId, IStrategy[] memory strategies) public createSnapshot {
        print.method("removeStrategiesFromOperatorSet");

        console.log("Removing strategies from operator set: %d", operatorSetId);

        for (uint i; i < strategies.length; ++i) {
            console.log("   strategy: %s", address(strategies[i]));
        }
        _tryPrankAppointee_AllocationManager(IAllocationManager.removeStrategiesFromOperatorSet.selector);
        config.allocationManager().removeStrategiesFromOperatorSet(address(this), operatorSetId, strategies);
        print.gasUsed();
    }

    /// -----------------------------------------------------------------------
    /// IAVSRegistrar
    /// -----------------------------------------------------------------------

    function registerOperator(address operator, address avsIdentifier, uint32[] calldata operatorSetIds, bytes calldata data) external {}

    function deregisterOperator(address operator, address avsIdentifier, uint32[] calldata operatorSetIds) external {}

    function supportsAVS(address) external pure returns (bool) {
        return true;
    }

    /// -----------------------------------------------------------------------
    /// Internal Helpers
    /// -----------------------------------------------------------------------

    function _tryPrankAppointee(address target, bytes4 selector) internal {
        address[] memory appointees = config.permissionController().getAppointees(address(this), target, selector);
        if (appointees.length != 0) cheats.prank(appointees[0]);
    }

    function _tryPrankAppointee_AllocationManager(bytes4 selector) internal {
        return _tryPrankAppointee(address(config.allocationManager()), selector);
    }
}
