// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_Redistribution_Base is UpgradeTest {
    AVS avs;
    OperatorSet operatorSet;

    User operator;
    AllocateParams allocateParams;
    SlashingParams slashParams;

    User staker;
    IStrategy[] strategies;
    IERC20[] tokens; // underlying token for each strategy
    uint[] initTokenBalances;
    uint[] initDepositShares;

    /// Shared setup:
    ///
    /// 1. Generate staker, operator, and AVS
    /// 2. Staker deposits and delegates to operator
    /// 3. AVS creates an operator set containing the strategies held by the staker
    /// 4. Operator allocates to operator set
    /// 5. Operator registers for operator set
    /// 6. Operator is randomly slashed by the operatorSet
    function _init() internal override {
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator_NoAssets();
        (avs,) = _newRandomAVS();
        tokens = _getUnderlyingTokens(strategies);

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);

        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // 3. Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);

        // 4. Operator allocates to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);

        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // 5. Operator registers for operator set
        operator.registerForOperatorSet(operatorSet);

        // 6. Operator is randomly slashed by the operatorSet
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator_PreRedistribution(slashParams);
    }
}

contract Integration_Upgrade_Redistribution is Integration_Upgrade_Redistribution_Base {
    function testFuzz_upgrade_burn(uint24 r) public rand(r) {
        // 1. Upgrade contracts
        _upgradeEigenLayerContracts();

        // 2. Burn shares
        (address[] memory strategiesWithBurnableShares,) = strategyManager.getStrategiesWithBurnableShares();
        for (uint i = 0; i < strategiesWithBurnableShares.length; i++) {
            strategyManager.burnShares(IStrategy(strategiesWithBurnableShares[i]));
        }

        // Assert that there are no strategies with burned shares
        (address[] memory strategiesWithBurnedShares,) = strategyManager.getStrategiesWithBurnableShares();
        assertEq(strategiesWithBurnedShares.length, 0);
    }

    function testFuzz_burn_update_operatorSet(uint24 r) public rand(r) {
        // 1. Burn shares
        (address[] memory strategiesWithBurnableShares,) = strategyManager.getStrategiesWithBurnableShares();
        for (uint i = 0; i < strategiesWithBurnableShares.length; i++) {
            strategyManager.burnShares(IStrategy(strategiesWithBurnableShares[i]));
        }

        // 2. Upgrade contracts
        _upgradeEigenLayerContracts();

        // Assert that there are no strategies with burned shares
        (address[] memory strategiesWithBurnedShares,) = strategyManager.getStrategiesWithBurnableShares();
        assertEq(strategiesWithBurnedShares.length, 0);
    }
}
