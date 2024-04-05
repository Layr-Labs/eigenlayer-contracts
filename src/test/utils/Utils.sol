pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "src/contracts/strategies/StrategyBase.sol";

contract Utils {
    address constant dummyAdmin = address(uint160(uint256(keccak256("DummyAdmin"))));

    function deployNewStrategy(IERC20 token, IStrategyManager strategyManager, IPauserRegistry pauserRegistry, address admin) public returns (StrategyBase) {
        StrategyBase newStrategy = new StrategyBase(strategyManager);
        newStrategy = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(newStrategy),
                    address(admin),
                    ""
                )
            )
        );
        newStrategy.initialize(token, pauserRegistry);
        return newStrategy;
    }
}
