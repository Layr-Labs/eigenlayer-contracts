// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/unit/StrategyManagerUnit/StrategyManagerUnit.t.sol";

contract StrategyManagerUnitTests_initialize is StrategyManagerUnitTests {
    function testCannotReinitialize() external {
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        strategyManager.initialize(initialOwner, initialOwner, pauserRegistry, 0);
    }
}
