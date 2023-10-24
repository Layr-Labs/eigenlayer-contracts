// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/unit/StrategyManagerUnit/addStrategiesToDepositWhitelist.t.sol";

contract StrategyManagerUnitTests_removeStrategiesFromDepositWhitelist is
    StrategyManagerUnitTests_addStrategiesToDepositWhitelist
{
    function testRemoveStrategiesFromDepositWhitelist(
        uint8 numberOfStrategiesToAdd,
        uint8 numberOfStrategiesToRemove
    ) external {
        // sanity filtering on fuzzed input
        cheats.assume(numberOfStrategiesToAdd <= 16);
        cheats.assume(numberOfStrategiesToRemove <= 16);
        cheats.assume(numberOfStrategiesToRemove <= numberOfStrategiesToAdd);

        IStrategy[] memory strategiesAdded = testAddStrategiesToDepositWhitelist(numberOfStrategiesToAdd);

        IStrategy[] memory strategiesToRemove = new IStrategy[](numberOfStrategiesToRemove);
        // loop that selectively copies from array to other array
        for (uint256 i = 0; i < numberOfStrategiesToRemove; ++i) {
            strategiesToRemove[i] = strategiesAdded[i];
        }

        cheats.startPrank(strategyManager.strategyWhitelister());
        for (uint256 i = 0; i < strategiesToRemove.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyRemovedFromDepositWhitelist(strategiesToRemove[i]);
        }
        strategyManager.removeStrategiesFromDepositWhitelist(strategiesToRemove);
        cheats.stopPrank();

        for (uint256 i = 0; i < numberOfStrategiesToAdd; ++i) {
            if (i < numberOfStrategiesToRemove) {
                require(
                    !strategyManager.strategyIsWhitelistedForDeposit(strategiesToRemove[i]),
                    "strategy not properly removed from whitelist"
                );
            } else {
                require(
                    strategyManager.strategyIsWhitelistedForDeposit(strategiesAdded[i]),
                    "strategy improperly removed from whitelist?"
                );
            }
        }
    }

    function testRemoveStrategiesFromDepositWhitelistRevertsWhenCalledByNotStrategyWhitelister(
        address notStrategyWhitelister
    ) external filterFuzzedAddressInputs(notStrategyWhitelister) {
        cheats.assume(notStrategyWhitelister != strategyManager.strategyWhitelister());
        IStrategy[] memory strategyArray = testAddStrategiesToDepositWhitelist(1);

        cheats.startPrank(notStrategyWhitelister);
        cheats.expectRevert(bytes("StrategyManager.onlyStrategyWhitelister: not the strategyWhitelister"));
        strategyManager.removeStrategiesFromDepositWhitelist(strategyArray);
        cheats.stopPrank();
    }
}
