// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/test/mocks/StrategyManagerMock.sol";
import "src/test/mocks/DelegationManagerMock.sol";
import "src/test/mocks/SlasherMock.sol";
import "src/test/mocks/EigenPodManagerMock.sol";
import "src/test/utils/Utils.sol";
import "src/test/utils/EigenLayerUnitTestBase.sol";

abstract contract EigenLayerUnitTestSetup is EigenLayerUnitTestBase, Utils {
    // Declare Mocks
    StrategyManagerMock strategyManagerMock;
    DelegationManagerMock public delegationManagerMock;
    SlasherMock public slasherMock;
    EigenPodManagerMock public eigenPodManagerMock;
    // strategies
    IERC20 public weth;
    IERC20 public eigenToken;
    StrategyBase public wethStrat;
    StrategyBase public eigenStrat;
    StrategyBase public baseStrategyImplementation;

    function setUp() public virtual override {
        EigenLayerUnitTestBase.setUp();
        strategyManagerMock = new StrategyManagerMock();
        delegationManagerMock = new DelegationManagerMock();
        slasherMock = new SlasherMock();
        eigenPodManagerMock = new EigenPodManagerMock();

        // deploy upgradeable proxy that points to StrategyBase implementation and initialize it
        baseStrategyImplementation = new StrategyBase(strategyManagerMock);

        addressIsExcludedFromFuzzedInputs[address(0)] = true;
        addressIsExcludedFromFuzzedInputs[address(strategyManagerMock)] = true;
        addressIsExcludedFromFuzzedInputs[address(delegationManagerMock)] = true;
        addressIsExcludedFromFuzzedInputs[address(slasherMock)] = true;
        addressIsExcludedFromFuzzedInputs[address(eigenPodManagerMock)] = true;
    }
}
