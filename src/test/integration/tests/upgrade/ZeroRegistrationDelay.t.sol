// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_ZeroRegistrationDelay is UpgradeTest {
    User operator;
    User operator2;

    AVS avs;
    IStrategy[] strategies;
    OperatorSet operatorSet;

    function _init() internal override {
        (, strategies,) = _newRandomStaker();
        (avs,) = _newRandomAVS();
        operatorSet = avs.createOperatorSet(strategies);
    }

    function testFuzz_register_upgrade_registerNew(uint24 r) public rand(r) {
        // 1. Register operator
        operator = _newRandomOperator_NoAssets();

        (bool isSetOperator1,) = allocationManager.getAllocationDelay(address(operator));
        // The delay should not be set since we adjusted `_newRandomOperator` to roll forward only 1 block after registration
        assertFalse(isSetOperator1, "isSet should be false");

        // 2. Upgrade contracts
        _upgradeEigenLayerContracts();

        // 3. Register new operator
        operator2 = _newRandomOperator_NoAssets();
        (bool isSetOperator2,) = allocationManager.getAllocationDelay(address(operator2));

        // The delay should be set since we are now with a 1 block set allocation delay
        assertTrue(isSetOperator2, "isSet should be true");

        // The delay for operator1 should still not be set
        (isSetOperator1,) = allocationManager.getAllocationDelay(address(operator));
        assertFalse(isSetOperator1, "isSet should be false");

        // 4. Assert that operator1 is set immediately
        rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY + 1});
        (isSetOperator1,) = allocationManager.getAllocationDelay(address(operator));
        assertTrue(isSetOperator1, "isSet should be true");
    }
}
