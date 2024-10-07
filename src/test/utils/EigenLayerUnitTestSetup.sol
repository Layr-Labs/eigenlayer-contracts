// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/mocks/StrategyManagerMock.sol";
import "src/test/mocks/DelegationManagerMock.sol";
import "src/test/mocks/EigenPodManagerMock.sol";
import "src/test/utils/EigenLayerUnitTestBase.sol";

abstract contract EigenLayerUnitTestSetup is EigenLayerUnitTestBase {
    // Declare Mocks
    StrategyManagerMock public strategyManagerMock;
    DelegationManagerMock public delegationManagerMock;
    EigenPodManagerMock public eigenPodManagerMock;

    function setUp() public virtual override {
        EigenLayerUnitTestBase.setUp();
        strategyManagerMock = new StrategyManagerMock();
        delegationManagerMock = new DelegationManagerMock();
        eigenPodManagerMock = new EigenPodManagerMock(pauserRegistry);

        addressIsExcludedFromFuzzedInputs[address(0)] = true;
        addressIsExcludedFromFuzzedInputs[address(strategyManagerMock)] = true;
        addressIsExcludedFromFuzzedInputs[address(delegationManagerMock)] = true;
        addressIsExcludedFromFuzzedInputs[address(eigenPodManagerMock)] = true;
    }
}
