// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_ALMSlashBase is IntegrationCheckUtils {
    AVS avs;
    OperatorSet operatorSet;

    User operator;
    AllocateParams allocateParams;

    User staker;
    IStrategy[] strategies;
    uint[] initTokenBalances;
    uint[] initDepositShares;

    /// Shared setup:
    /// 
    /// 1. Generate staker, operator, and AVS
    /// 2. Staker deposits and delegates to operator
    /// 3. AVS creates an operator set containing the strategies held by the staker
    /// 4. Operator allocates to operator set
    /// 5. Operator registers for operator set
    /// NOTE: Steps 4 and 5 are done in random order, as these should not have an outcome on the test
    function _init() internal virtual override {
        _configAssetTypes(HOLDS_LST);
        // (staker, strategies, initTokenBalances) = _newRandomStaker();
        // operator = _newRandomOperator_NoAssets();
        // (avs,) = _newRandomAVS();

        (staker, strategies, initTokenBalances) = _newBasicStaker();
        operator = _newRandomOperator_NoAssets();
        (avs,) = _newRandomAVS();

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, initDepositShares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 3. Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);

        // randomly choose between:
        // register -> allocate / allocate -> register
        if (_randBool()) {
            // register -> allocate
            operator.registerForOperatorSet(operatorSet);
            check_Registration_State_NoAllocation(operator, operatorSet, allStrats);

            allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
            operator.modifyAllocations(allocateParams);
            check_IncrAlloc_State_Slashable(operator, allocateParams);
        } else {
            // allocate -> register
            allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
            operator.modifyAllocations(allocateParams);
            check_IncrAlloc_State_NotSlashable(operator, allocateParams);

            operator.registerForOperatorSet(operatorSet);
            check_Registration_State_PendingAllocation(operator, allocateParams);
        }
   
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
    }
}

contract Integration_InitSlash is Integration_ALMSlashBase {

    SlashingParams slashParams;

    function testFuzz_slashSingle(uint24 _r) public rand(_r) {
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);
    } 

    function testFuzz_slashMulti_WithdrawTokens(uint24 _r) public rand(_r) {
        for (uint i = 0; i < 25; i++) {
            slashParams = _genSlashing_Rand(operator, operatorSet);
            avs.slashOperator(slashParams);
            check_Base_Slashing_State(operator, allocateParams, slashParams);
        }

        // undelegate
        uint[] memory shares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory roots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, roots, strategies, initDepositShares, shares);

        _rollBlocksForCompleteWithdrawals(withdrawals);

        // try withdrawing as tokens:
        IERC20[] memory tokens = _getUnderlyingTokens(strategies);
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);

        staker.completeWithdrawalsAsTokens(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, shares, tokens, expectedTokens);
        }
    }

    function testFuzz_slashMulti_WithdrawShares(uint24 _r) public rand(_r) {
        for (uint i = 0; i < 25; i++) {
            slashParams = _genSlashing_Rand(operator, operatorSet);
            avs.slashOperator(slashParams);
            check_Base_Slashing_State(operator, allocateParams, slashParams);
        }

        // undelegate
        uint[] memory shares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory roots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, roots, strategies, initDepositShares, shares);

        _rollBlocksForCompleteWithdrawals(withdrawals);

        // try withdrawing as shares
        staker.completeWithdrawalsAsShares(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], strategies, shares);
        }
    }
}

contract Integration_SlashingWithdrawals is Integration_ALMSlashBase {
<<<<<<< HEAD

    function testFuzz_slash_undelegate_completeAsTokens(
        uint24 _random
    ) public rand(_random) {
        // 4. Slash operator
        SlashingParams memory slashingParams;

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 3. Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);

        // randomly choose between:
        // register -> allocate / allocate -> register
        if (_randBool()) {
            // register -> allocate
            operator.registerForOperatorSet(operatorSet);
            check_Registration_State_NoAllocation(operator, operatorSet, allStrats);

            allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
            operator.modifyAllocations(allocateParams);
            check_IncrAlloc_State_Slashable(operator, allocateParams);
        } else {
            // allocate -> register
            allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
            operator.modifyAllocations(allocateParams);
            check_IncrAlloc_State_NotSlashable(operator, allocateParams);

            operator.registerForOperatorSet(operatorSet);
            check_Registration_State_PendingAllocation(operator, allocateParams);
        }
   
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
    }

<<<<<<<< HEAD:src/test/integration/tests/Deposit_Delegate_Allocate.t.sol
    function testFuzz_deposit_delegate_upgrade_allocate(
        uint24 _random
    ) public {
        // Configure the random parameters for the test
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
            _userTypes: DEFAULT | ALT_METHODS
        });

        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);

        // 2. Delegate to an operator
        staker.delegateTo(operator);

        _upgradeEigenLayerContracts(); // Upgrade contracts if forkType is not local
        (AVS avs,) = _newRandomAVS();

        // 3. Set allocation delay for operator
        operator.setAllocationDelay(1);
        rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY});

        // 4. Create an operator set and register an operator.
        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);

        // 5. Allocate to operator set.
        IAllocationManagerTypes.AllocateParams memory allocateParams =
            operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );
    }

    function testFuzz_deposit_delegate_allocate_slash_undelegate_completeAsTokens(
        uint24 _random
    ) public {
        _configRand({_randomSeed: _random, _assetTypes: HOLDS_ALL, _userTypes: DEFAULT});
        _upgradeEigenLayerContracts(); // Upgrade contracts if forkType is not local

        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);

        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // Create an operator set and register an operator.
        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);

        // 3. Allocate to operator set.
        IAllocationManagerTypes.AllocateParams memory allocateParams =
            operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );
=======
>>>>>>> b6c49a95 (test: slash invariants (#1083))

    function testFuzz_slash_undelegate_completeAsTokens(
        uint24 _random
    ) public rand(_random) {
        // 4. Slash operator
        SlashingParams memory slashingParams;
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker deposit shares should be slashed");
        }

        // 5. Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 6. Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens =
                _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State_AfterSlash(staker, operator, withdrawals[i], allocateParams, slashingParams, expectedTokens);
        }

        // Check Final State
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");
        assert_HasUnderlyingTokenBalances_AfterSlash(
            staker,
            allocateParams,
            slashingParams,
            initTokenBalances,
            "staker should once again have original token initTokenBalances minus slashed"
        );
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_slash_undelegate_completeAsShares(
        uint24 _random
<<<<<<< HEAD
    ) public rand(_random) {
        // 4. Slash operator
        SlashingParams memory slashingParams;
<<<<<<< HEAD
=======
<<<<<<<< HEAD:src/test/integration/tests/Deposit_Delegate_Allocate.t.sol
    ) public {
        _configRand({_randomSeed: _random, _assetTypes: HOLDS_ALL, _userTypes: DEFAULT});
        _upgradeEigenLayerContracts(); // Upgrade contracts if forkType is not local

        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();


        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint256[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
        // TODO - post-deposit and post-delegate checks?

        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // Create an operator set and register an operator.
        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);

        // 3. Allocate to operator set.
        IAllocationManagerTypes.AllocateParams memory allocateParams =
            operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );

========
    ) public rand(_random) {
>>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036)):src/test/integration/tests/SlashingWithdrawals.t.sol
        // 4. Slash operator
        IAllocationManagerTypes.SlashingParams memory slashingParams;
>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036))
=======
>>>>>>> d6510cc7 (test(integration): implement registration and allocation invariants (#1042))
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker deposit shares should be slashed");
        }

        // 5. Undelegate from an operator
<<<<<<< HEAD
<<<<<<< HEAD
        Withdrawal[] memory withdrawals = staker.undelegate();
=======
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.undelegate();
>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036))
=======
        Withdrawal[] memory withdrawals = staker.undelegate();
>>>>>>> d6510cc7 (test(integration): implement registration and allocation invariants (#1042))
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State_AfterSlash(staker, operator, withdrawals[i], allocateParams, slashingParams);
        }

        // Check final state:
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker not have any underlying tokens");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_queue_slash_completeAsTokens(
        uint24 _random
<<<<<<< HEAD
    ) public rand(_random) {
        // 4. Queue withdrawal
        Withdrawal[] memory withdrawals =
<<<<<<< HEAD
=======
<<<<<<<< HEAD:src/test/integration/tests/Deposit_Delegate_Allocate.t.sol
    ) public {
        _configRand({_randomSeed: _random, _assetTypes: HOLDS_ALL, _userTypes: DEFAULT});
        _upgradeEigenLayerContracts(); // Upgrade contracts if forkType is not local

        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // Create an operator set and register an operator.
        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);

        // 3. Allocate to operator set.
        IAllocationManagerTypes.AllocateParams memory allocateParams =
            operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );

========
    ) public rand(_random) {
>>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036)):src/test/integration/tests/SlashingWithdrawals.t.sol
        // 4. Queue withdrawal
        IDelegationManagerTypes.Withdrawal[] memory withdrawals =
>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036))
=======
>>>>>>> d6510cc7 (test(integration): implement registration and allocation invariants (#1042))
            staker.queueWithdrawals(strategies, _calculateExpectedShares(strategies, initTokenBalances));
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 5. Slash operator
<<<<<<< HEAD
<<<<<<< HEAD
        SlashingParams memory slashingParams;
=======
        IAllocationManagerTypes.SlashingParams memory slashingParams;
>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036))
=======
        SlashingParams memory slashingParams;
>>>>>>> d6510cc7 (test(integration): implement registration and allocation invariants (#1042))
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker deposit shares should be slashed");
        }

        // 6. Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens =
                _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State_AfterSlash(
                staker, operator, withdrawals[i], allocateParams, slashingParams, expectedTokens
            );
        }

        // Check Final State
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");
        assert_HasUnderlyingTokenBalances_AfterSlash(
            staker,
            allocateParams,
            slashingParams,
            initTokenBalances,
            "staker should once again have original token initTokenBalances minus slashed"
        );
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_queue_slash_completeAsShares(
        uint24 _random
<<<<<<< HEAD
    ) public rand(_random) {
        // 4. Queue withdrawal
        Withdrawal[] memory withdrawals =
<<<<<<< HEAD
=======
<<<<<<<< HEAD:src/test/integration/tests/Deposit_Delegate_Allocate.t.sol
    ) public {
        _configRand({_randomSeed: _random, _assetTypes: HOLDS_ALL, _userTypes: DEFAULT});
        _upgradeEigenLayerContracts(); // Upgrade contracts if forkType is not local

        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // Create an operator set and register an operator.
        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        operator.setAllocationDelay(1);

        // 3. Allocate to operator set.
        IAllocationManagerTypes.AllocateParams memory allocateParams =
            operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );

========
    ) public rand(_random) {
>>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036)):src/test/integration/tests/SlashingWithdrawals.t.sol
        // 4. Queue withdrawal
        IDelegationManagerTypes.Withdrawal[] memory withdrawals =
>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036))
=======
>>>>>>> d6510cc7 (test(integration): implement registration and allocation invariants (#1042))
            staker.queueWithdrawals(strategies, _calculateExpectedShares(strategies, initTokenBalances));
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 5. Slash operator
<<<<<<< HEAD
<<<<<<< HEAD
        SlashingParams memory slashingParams;
=======
        IAllocationManagerTypes.SlashingParams memory slashingParams;
>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036))
=======
        SlashingParams memory slashingParams;
>>>>>>> d6510cc7 (test(integration): implement registration and allocation invariants (#1042))
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker deposit shares should be slashed");
        }

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State_AfterSlash(staker, operator, withdrawals[i], allocateParams, slashingParams);
        }

        // Check final state:
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker not have any underlying tokens");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_deallocate_slash_queue_completeAsTokens(
        uint24 _random
<<<<<<< HEAD
    ) public rand(_random) {
        // 4. Deallocate all.
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_DecrAlloc_State_Slashable(operator, deallocateParams);

        _rollForward_DeallocationDelay();
<<<<<<< HEAD

        // 5. Slash operator
        SlashingParams memory slashingParams;
=======
<<<<<<<< HEAD:src/test/integration/tests/Deposit_Delegate_Allocate.t.sol
    ) public {
        _configRand({_randomSeed: _random, _assetTypes: HOLDS_ALL, _userTypes: DEFAULT});
        _upgradeEigenLayerContracts(); // Upgrade contracts if forkType is not local

        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // Create an operator set and register an operator.
        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        operator.setAllocationDelay(1);

        console.log("block allocated", block.number);
        // 3. Allocate to operator set.
        IAllocationManagerTypes.AllocateParams memory allocateParams =
            operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );

========
    ) public rand(_random) {
>>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036)):src/test/integration/tests/SlashingWithdrawals.t.sol
        // 4. Deallocate all.
        IAllocationManagerTypes.AllocateParams memory deallocateParams = operator.deallocateAll(operatorSet);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // 5. Slash operator
        IAllocationManagerTypes.SlashingParams memory slashingParams;
>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036))
=======

        // 5. Slash operator
        SlashingParams memory slashingParams;
>>>>>>> d6510cc7 (test(integration): implement registration and allocation invariants (#1042))
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, deallocateParams, slashingParams, "staker deposit shares should be slashed");
        }
        
        // 6. Queue withdrawals
<<<<<<< HEAD
<<<<<<< HEAD
        Withdrawal[] memory withdrawals =
=======
        IDelegationManagerTypes.Withdrawal[] memory withdrawals =
>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036))
=======
        Withdrawal[] memory withdrawals =
>>>>>>> d6510cc7 (test(integration): implement registration and allocation invariants (#1042))
            staker.queueWithdrawals(strategies, _calculateExpectedShares(strategies, initTokenBalances));
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 7. Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens =
                _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State_AfterSlash(
                staker, operator, withdrawals[i], allocateParams, slashingParams, expectedTokens
            );
        }
        
        // Check Final State
        assert_HasUnderlyingTokenBalances(
            staker,
            allocateParams.strategies,
            initTokenBalances,
            "staker should have withdrawn all shares"
        );
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_deregister_slash(
        uint24 _random
<<<<<<< HEAD
    ) public rand(_random) {
        // 4. Deregister.
        operator.deregisterFromOperatorSet(operatorSet);
        check_Deregistration_State_PendingAllocation(operator, operatorSet);
<<<<<<< HEAD

        // 5. Slash operator
        SlashingParams memory slashingParams;
=======
<<<<<<<< HEAD:src/test/integration/tests/Deposit_Delegate_Allocate.t.sol
    ) public {
        _configRand({_randomSeed: _random, _assetTypes: HOLDS_ALL, _userTypes: DEFAULT});
        _upgradeEigenLayerContracts(); // Upgrade contracts if forkType is not local

        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // Create an operator set and register an operator.
        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);

        // 3. Allocate to operator set.
        IAllocationManagerTypes.AllocateParams memory allocateParams =
            operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );

========
    ) public rand(_random) {
>>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036)):src/test/integration/tests/SlashingWithdrawals.t.sol
        // 4. Deregister.
        operator.deregisterFromOperatorSet(operatorSet);

        // 5. Slash operator
        IAllocationManagerTypes.SlashingParams memory slashingParams;
>>>>>>> 6e2786c8 (test: enable shared setups for integration tests (#1036))
=======

        // 5. Slash operator
        SlashingParams memory slashingParams;
>>>>>>> d6510cc7 (test(integration): implement registration and allocation invariants (#1042))
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker deposit shares should be slashed");
        }
    }
}