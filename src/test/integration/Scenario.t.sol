// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/IntegrationTestRunner.t.sol";
import "src/test/integration/Staker.sol";
import "src/test/integration/GlobalRefs.sol";

contract TestScenario is IntegrationTestRunner {

    GlobalRefs globalRefs;
    Staker staker;

    function setUp() public override {
        // Setup parent
        IntegrationTestRunner.setUp();

        // Deploy Global Reference Contract
        globalRefs = new GlobalRefs(eigenPodManager, delegationManager, strategyManager);

        // Deploy Staker Contract
        staker = new Staker(globalRefs);
    }
    function test_flow1() public {
        assertTrue(true);
    }
}