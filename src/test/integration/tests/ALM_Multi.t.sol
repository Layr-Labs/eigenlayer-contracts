// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_ALM_Multi is IntegrationCheckUtils {
    using StdStyle for *;

    enum Action {
        REGISTER,
        DEREGISTER,
        INCR_ALLOC,
        INCR_ALLOC_FULL,
        DECR_ALLOC,
        DECR_ALLOC_FULL
    }

    enum State {
        NONE,
        REGISTERED,
        ALLOCATED,
        FULLY_ALLOCATED,
        REG_ALLOC,
        REG_FULLY_ALLOC
    }

    AVS avs;
    OperatorSet operatorSet;

    IStrategy[] strategies;

    /// iteration idx -> list of operators in each state
    mapping(uint => mapping(State => User[])) operators;

    /// operator -> list of strategies they have delegated assets in
    mapping(User => IStrategy[]) allocatedStrats;
    /// Last modifyAllocations params made by the operator
    mapping(User => AllocateParams) lastModifyParams;

    uint constant NUM_UNIQUE_ASSETS = 3;
    uint constant NUM_OPERATORS = 5;
    uint constant NUM_ITERATIONS = 20;

    function _init() internal virtual override {
        _configAssetAmounts(NUM_UNIQUE_ASSETS);

        (avs,) = _newRandomAVS();
        operatorSet = avs.createOperatorSet(allStrats);

        for (uint i = 0; i < NUM_OPERATORS; i++) {
            (User staker, IStrategy[] memory _strategies, uint[] memory initTokenBalances) = _newRandomStaker();

            User operator = _newRandomOperator_NoAssets();

            // 1. Deposit into strategies
            staker.depositIntoEigenlayer(_strategies, initTokenBalances);
            uint[] memory initDepositShares = _calculateExpectedShares(_strategies, initTokenBalances);
            check_Deposit_State(staker, _strategies, initDepositShares);

            // 2. Delegate to operator
            staker.delegateTo(operator);
            check_Delegation_State(staker, operator, _strategies, initDepositShares);

            allocatedStrats[operator] = _strategies;
            // Add operator to NONE state for the 0th iteration
            operators[0][State.NONE].push(operator);
        }
    }

    /// Reduce fuzz runs because this test is thiccc:
    ///
    /// forge-config: default.fuzz.runs = 10
    /// forge-config: forktest.fuzz.runs = 3
    function test_Multi(uint24 _r) public rand(_r) {
        // Do 20 iterations
        for (uint i = 1; i <= NUM_ITERATIONS; i++) {
            console.log("%s: %d", "iter".green().italic(), i - 1);

            _dispatchNone(i);
            _dispatchRegistered(i);
            _dispatchAllocated(i);
            _dispatchFullyAllocated(i);
            _dispatchRegAlloc(i);
            _dispatchRegFullyAlloc(i);

            // Ensure all pending actions are completed for the next iteration
            _rollForward_DeallocationDelay();
        }
    }

    /// @dev NONE operators can:
    /// [REGISTER, INCR_ALLOC, INCR_ALLOC_FULL]
    function _dispatchNone(uint iter) internal {
        // Fetch all NONE operators from previous iteration
        User[] memory _operators = operators[iter - 1][State.NONE];
        Action[3] memory actions = [Action.REGISTER, Action.INCR_ALLOC, Action.INCR_ALLOC_FULL];

        if (_operators.length == 0) return;

        console.log("%s: %d operators", "_dispatchNone".green(), _operators.length);

        for (uint i = 0; i < _operators.length; i++) {
            // Get operator and allocated strategies
            User operator = _operators[i];
            IStrategy[] memory _strats = allocatedStrats[operator];

            // Get action
            uint aI = _randUint(0, actions.length - 1);
            Action action = actions[aI];

            // Process action
            if (action == Action.REGISTER) {
                operator.registerForOperatorSet(operatorSet);
                check_Registration_State_NoAllocation(operator, operatorSet, allStrats);

                operators[iter][State.REGISTERED].push(operator);
            } else if (action == Action.INCR_ALLOC) {
                AllocateParams memory params = _genAllocation_HalfAvailable(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_IncrAlloc_State_NotSlashable(operator, params);

                lastModifyParams[operator] = params;
                operators[iter][State.ALLOCATED].push(operator);
            } else if (action == Action.INCR_ALLOC_FULL) {
                AllocateParams memory params = _genAllocation_AllAvailable(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_IncrAlloc_State_NotSlashable(operator, params);

                lastModifyParams[operator] = params;
                operators[iter][State.FULLY_ALLOCATED].push(operator);
            }
        }
    }

    /// @dev REGISTERED operators can:
    /// [DEREGISTER, INCR_ALLOC, INCR_ALLOC_FULL]
    function _dispatchRegistered(uint iter) internal {
        // Fetch all REGISTERED operators from previous iteration
        User[] memory _operators = operators[iter - 1][State.REGISTERED];
        Action[3] memory actions = [Action.DEREGISTER, Action.INCR_ALLOC, Action.INCR_ALLOC_FULL];

        if (_operators.length == 0) return;

        console.log("%s: %d operators", "_dispatchRegistered".green(), _operators.length);

        for (uint i = 0; i < _operators.length; i++) {
            // Get operator
            User operator = _operators[i];
            IStrategy[] memory _strats = allocatedStrats[operator];

            // Get action
            uint aI = _randUint(0, actions.length - 1);
            Action action = actions[aI];

            // Process action
            if (action == Action.DEREGISTER) {
                operator.deregisterFromOperatorSet(operatorSet);
                check_Deregistration_State_NoAllocation(operator, operatorSet);

                operators[iter][State.NONE].push(operator);
            } else if (action == Action.INCR_ALLOC) {
                AllocateParams memory params = _genAllocation_HalfAvailable(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_IncrAlloc_State_Slashable(operator, params);

                lastModifyParams[operator] = params;
                operators[iter][State.REG_ALLOC].push(operator);
            } else if (action == Action.INCR_ALLOC_FULL) {
                AllocateParams memory params = _genAllocation_AllAvailable(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_IncrAlloc_State_Slashable(operator, params);
                // check_FullyAllocated_State(operator, operatorSet, params.strategies); TODO

                lastModifyParams[operator] = params;
                operators[iter][State.REG_FULLY_ALLOC].push(operator);
            }
        }
    }

    /// @dev ALLOCATED operators can:
    /// [REGISTER, INCR_ALLOC, INCR_ALLOC_FULL, DECR_ALLOC, DECR_ALLOC_FULL]
    function _dispatchAllocated(uint iter) internal {
        // Fetch all ALLOCATED operators from previous iteration
        User[] memory _operators = operators[iter - 1][State.ALLOCATED];
        Action[5] memory actions = [Action.REGISTER, Action.INCR_ALLOC, Action.INCR_ALLOC_FULL, Action.DECR_ALLOC, Action.DECR_ALLOC_FULL];

        if (_operators.length == 0) return;

        console.log("%s: %d operators", "_dispatchAllocated".green(), _operators.length);

        for (uint i = 0; i < _operators.length; i++) {
            // Get operator
            User operator = _operators[i];
            IStrategy[] memory _strats = allocatedStrats[operator];

            // Get action
            uint aI = _randUint(0, actions.length - 1);
            Action action = actions[aI];

            // Process action
            if (action == Action.REGISTER) {
                operator.registerForOperatorSet(operatorSet);
                check_Registration_State_ActiveAllocation(operator, lastModifyParams[operator]);

                operators[iter][State.REG_ALLOC].push(operator);
            } else if (action == Action.INCR_ALLOC) {
                AllocateParams memory params = _genAllocation_HalfAvailable(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_IncrAlloc_State_NotSlashable(operator, params);

                lastModifyParams[operator] = params;
                operators[iter][State.ALLOCATED].push(operator);
            } else if (action == Action.INCR_ALLOC_FULL) {
                AllocateParams memory params = _genAllocation_AllAvailable(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_IncrAlloc_State_NotSlashable(operator, params);
                // check_FullyAllocated_State(operator); TODO

                lastModifyParams[operator] = params;
                operators[iter][State.FULLY_ALLOCATED].push(operator);
            } else if (action == Action.DECR_ALLOC) {
                AllocateParams memory params = _genDeallocation_HalfRemaining(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_DecrAlloc_State_NotSlashable(operator, params);

                lastModifyParams[operator] = params;
                operators[iter][State.ALLOCATED].push(operator);
            } else if (action == Action.DECR_ALLOC_FULL) {
                AllocateParams memory params = _genDeallocation_Full(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_DecrAlloc_State_NotSlashable(operator, params);
                // check_FullyDeallocated_State(operator); TODO

                lastModifyParams[operator] = params;
                operators[iter][State.NONE].push(operator);
            }
        }
    }

    /// @dev FULLY_ALLOCATED operators can:
    /// [REGISTER, DECR_ALLOC, DECR_ALLOC_FULL]
    function _dispatchFullyAllocated(uint iter) internal {
        // Fetch all FULLY_ALLOCATED operators from previous iteration
        User[] memory _operators = operators[iter - 1][State.FULLY_ALLOCATED];
        Action[3] memory actions = [Action.REGISTER, Action.DECR_ALLOC, Action.DECR_ALLOC_FULL];

        if (_operators.length == 0) return;

        console.log("%s: %d operators", "_dispatchFullyAllocated".green(), _operators.length);

        for (uint i = 0; i < _operators.length; i++) {
            // Get operator
            User operator = _operators[i];
            IStrategy[] memory _strats = allocatedStrats[operator];

            // Get action
            uint aI = _randUint(0, actions.length - 1);
            Action action = actions[aI];

            // Process action
            if (action == Action.REGISTER) {
                operator.registerForOperatorSet(operatorSet);
                check_Registration_State_ActiveAllocation(operator, lastModifyParams[operator]);

                operators[iter][State.REG_FULLY_ALLOC].push(operator);
            } else if (action == Action.DECR_ALLOC) {
                AllocateParams memory params = _genDeallocation_HalfRemaining(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_DecrAlloc_State_NotSlashable(operator, params);

                lastModifyParams[operator] = params;
                operators[iter][State.ALLOCATED].push(operator);
            } else if (action == Action.DECR_ALLOC_FULL) {
                AllocateParams memory params = _genDeallocation_Full(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_DecrAlloc_State_NotSlashable(operator, params);
                // check_FullyDeallocated_State(operator); TODO

                lastModifyParams[operator] = params;
                operators[iter][State.NONE].push(operator);
            }
        }
    }

    /// @dev REG_ALLOC operators can:
    /// [DEREGISTER, INCR_ALLOC, INCR_ALLOC_FULL, DECR_ALLOC, DECR_ALLOC_FULL]
    function _dispatchRegAlloc(uint iter) internal {
        // Fetch all REG_ALLOC operators from previous iteration
        User[] memory _operators = operators[iter - 1][State.REG_ALLOC];
        Action[5] memory actions = [Action.DEREGISTER, Action.INCR_ALLOC, Action.INCR_ALLOC_FULL, Action.DECR_ALLOC, Action.DECR_ALLOC_FULL];

        if (_operators.length == 0) return;

        console.log("%s: %d operators", "_dispatchRegAlloc".green(), _operators.length);

        for (uint i = 0; i < _operators.length; i++) {
            // Get operator
            User operator = _operators[i];
            IStrategy[] memory _strats = allocatedStrats[operator];

            // Get action
            uint aI = _randUint(0, actions.length - 1);
            Action action = actions[aI];

            // Process action
            if (action == Action.DEREGISTER) {
                operator.deregisterFromOperatorSet(operatorSet);
                check_Deregistration_State_ActiveAllocation(operator, operatorSet);

                operators[iter][State.ALLOCATED].push(operator);
            } else if (action == Action.INCR_ALLOC) {
                AllocateParams memory params = _genAllocation_HalfAvailable(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_IncrAlloc_State_Slashable(operator, params);

                lastModifyParams[operator] = params;
                operators[iter][State.REG_ALLOC].push(operator);
            } else if (action == Action.INCR_ALLOC_FULL) {
                AllocateParams memory params = _genAllocation_AllAvailable(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_IncrAlloc_State_Slashable(operator, params);
                // check_FullyAllocated_State(operator); TODO

                lastModifyParams[operator] = params;
                operators[iter][State.REG_FULLY_ALLOC].push(operator);
            } else if (action == Action.DECR_ALLOC) {
                AllocateParams memory params = _genDeallocation_HalfRemaining(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_DecrAlloc_State_Slashable(operator, params);

                lastModifyParams[operator] = params;
                operators[iter][State.REG_ALLOC].push(operator);
            } else if (action == Action.DECR_ALLOC_FULL) {
                AllocateParams memory params = _genDeallocation_Full(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_DecrAlloc_State_Slashable(operator, params);
                // check_FullyDeallocated_State(operator); TODO

                lastModifyParams[operator] = params;
                operators[iter][State.REGISTERED].push(operator);
            }
        }
    }

    /// @dev REG_FULLY_ALLOC operators can:
    /// [DEREGISTER, DECR_ALLOC, DECR_ALLOC_FULL]
    function _dispatchRegFullyAlloc(uint iter) internal {
        // Fetch all REG_FULLY_ALLOC operators from previous iteration
        User[] memory _operators = operators[iter - 1][State.REG_FULLY_ALLOC];
        Action[3] memory actions = [Action.DEREGISTER, Action.DECR_ALLOC, Action.DECR_ALLOC_FULL];

        if (_operators.length == 0) return;

        console.log("%s: %d operators", "_dispatchRegFullyAlloc".green(), _operators.length);

        for (uint i = 0; i < _operators.length; i++) {
            // Get operator
            User operator = _operators[i];
            IStrategy[] memory _strats = allocatedStrats[operator];

            // Get action
            uint aI = _randUint(0, actions.length - 1);
            Action action = actions[aI];

            // Process action
            if (action == Action.DEREGISTER) {
                operator.deregisterFromOperatorSet(operatorSet);
                check_Deregistration_State_ActiveAllocation(operator, operatorSet);

                operators[iter][State.FULLY_ALLOCATED].push(operator);
            } else if (action == Action.DECR_ALLOC) {
                AllocateParams memory params = _genDeallocation_HalfRemaining(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_DecrAlloc_State_Slashable(operator, params);

                lastModifyParams[operator] = params;
                operators[iter][State.REG_ALLOC].push(operator);
            } else if (action == Action.DECR_ALLOC_FULL) {
                AllocateParams memory params = _genDeallocation_Full(operator, operatorSet, _strats);
                operator.modifyAllocations(params);
                check_DecrAlloc_State_Slashable(operator, params);
                // check_FullyDeallocated_State(operator); TODO

                lastModifyParams[operator] = params;
                operators[iter][State.REGISTERED].push(operator);
            }
        }
    }
}
