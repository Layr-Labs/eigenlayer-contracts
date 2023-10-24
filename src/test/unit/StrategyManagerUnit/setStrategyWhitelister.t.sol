// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/unit/StrategyManagerUnit/StrategyManagerUnit.t.sol";

contract StrategyManagerUnitTests_setStrategyWhitelister is StrategyManagerUnitTests {
    function testSetStrategyWhitelister(address newWhitelister) external {
        address previousStrategyWhitelister = strategyManager.strategyWhitelister();
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyWhitelisterChanged(previousStrategyWhitelister, newWhitelister);
        strategyManager.setStrategyWhitelister(newWhitelister);
        require(
            strategyManager.strategyWhitelister() == newWhitelister,
            "strategyManager.strategyWhitelister() != newWhitelister"
        );
    }

    function testSetStrategyWhitelisterRevertsWhenCalledByNotOwner(
        address notOwner
    ) external filterFuzzedAddressInputs(notOwner) {
        cheats.assume(notOwner != strategyManager.owner());
        address newWhitelister = address(this);
        cheats.startPrank(notOwner);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        strategyManager.setStrategyWhitelister(newWhitelister);
        cheats.stopPrank();
    }
}
