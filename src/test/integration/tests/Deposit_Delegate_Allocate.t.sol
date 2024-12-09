// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

contract Integration_Deposit_Delegate_Allocate is IntegrationCheckUtils {
    function testFuzz_deposit_delegate_allocate(uint24 _random) public {
        // Configure the random parameters for the test
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
            _userTypes: DEFAULT | ALT_METHODS
        });
        
        // Create a staker and an operator with a nonzero balance and corresponding strategies
        (AVS avs, OperatorSet[] memory operatorSets) = _newRandomAVS();
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator, ,) = _newRandomOperator();

        // Upgrade contracts if forkType is not local
        _upgradeEigenLayerContracts();

        // 1. Delegate to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, new uint256[](strategies.length)); // Initial shares are zero

        // 2. Deposit into strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        // Check that the deposit increased operator shares the staker is delegated to
        check_Deposit_State(staker, strategies, shares);
        assert_Snap_Added_OperatorShares(operator, strategies, shares, "operator should have received shares");

        operator.registerForOperatorSets(operatorSets);
        
        for (uint i; i < operatorSets.length; ++i) {
            uint256 len = allocationManager.getStrategiesInOperatorSet(operatorSets[i]).length;
            operator.modifyAllocations(operatorSets[i], _randMagnitudes({ sum: 1 ether / uint64(operatorSets.length), len: len }));
            (IStrategy[] memory strategiesToSlash, uint[] memory wadsToSlash) = _randStrategiesAndWadsToSlash(operatorSets[i]);
            avs.slashOperator(operator, operatorSets[i].id, strategiesToSlash, wadsToSlash);
        }

        // TODO: write checks for slashing...
    }
}