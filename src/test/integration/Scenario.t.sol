// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/IntegrationTestRunner.t.sol";
import "src/test/integration/Operator.sol";
import "src/test/integration/Staker.sol";
import "src/test/integration/GlobalRefs.sol";

contract TestScenario is IntegrationTestRunner {
    GlobalRefs globalRefs;
    Staker staker;
    Operator operator;

    function setUp() public override {
        // Setup parent
        IntegrationTestRunner.setUp();

        // Deploy Global Reference Contract
        globalRefs = new GlobalRefs(eigenPodManager, delegationManager, strategyManager);

        // Deploy Staker and Operator contracts
        staker = new Staker(globalRefs);
        operator = new Operator(globalRefs);
    }

    function test_flow1() public {
        // Register Operator
        operator.register();

        // Delegate To Operator
        staker.delegate(operator);

        // Deposit into strategy
        staker.depositIntoStrategy(strategy1, strategy1Token, 1000);

        // Queue Withdrawal for entire strategy balance
        staker.queueFullWithdrawal(strategy1);

        // Complete first (0-indexed) queued withdrawal
        staker.completeQueuedWithdrawal(0);
    }
}