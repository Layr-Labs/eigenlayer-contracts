// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/unit/StrategyManagerUnit/StrategyManagerUnit.t.sol";

contract StrategyManagerUnitTests_removeShares is StrategyManagerUnitTests {
    function testRemoveSharesRevertsDelegationManagerModifier() external {
        DelegationManagerMock invalidDelegationManager = new DelegationManagerMock();
        cheats.expectRevert(bytes("StrategyManager.onlyDelegationManager: not the DelegationManager"));
        invalidDelegationManager.removeShares(strategyManager, address(this), dummyStrat, 1);
    }

    function testRemoveSharesRevertsShareAmountTooHigh(
        address staker,
        uint256 depositAmount,
        uint256 removeSharesAmount
    ) external {
        cheats.assume(staker != address(0));
        cheats.assume(depositAmount > 0 && depositAmount < dummyToken.totalSupply());
        cheats.assume(removeSharesAmount > depositAmount);
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        cheats.expectRevert(bytes("StrategyManager._removeShares: shareAmount too high"));
        delegationManagerMock.removeShares(strategyManager, staker, strategy, removeSharesAmount);
    }

    function testRemoveSharesRemovesStakerStrategyListSingleStrat(address staker, uint256 sharesAmount) external {
        cheats.assume(staker != address(0));
        cheats.assume(sharesAmount > 0 && sharesAmount < dummyToken.totalSupply());
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, sharesAmount);

        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
        require(sharesBefore == sharesAmount, "Staker has not deposited amount into strategy");

        delegationManagerMock.removeShares(strategyManager, staker, strategy, sharesAmount);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        require(
            stakerStrategyListLengthAfter == stakerStrategyListLengthBefore - 1,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore - 1"
        );
        require(sharesAfter == 0, "sharesAfter != 0");
        require(!_isDepositedStrategy(staker, strategy), "strategy should not be part of staker strategy list");
    }

    // Remove Strategy from staker strategy list with multiple strategies deposited
    function testRemoveSharesRemovesStakerStrategyListMultipleStrat(
        address staker,
        uint256[3] memory amounts,
        uint8 randStrategy
    ) external {
        cheats.assume(staker != address(0));
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = dummyStrat;
        strategies[1] = dummyStrat2;
        strategies[2] = dummyStrat3;
        for (uint256 i = 0; i < 3; ++i) {
            cheats.assume(amounts[i] > 0 && amounts[i] < dummyToken.totalSupply());
            _depositIntoStrategySuccessfully(strategies[i], staker, amounts[i]);
        }
        IStrategy removeStrategy = strategies[randStrategy % 3];
        uint256 removeAmount = amounts[randStrategy % 3];

        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256[] memory sharesBefore = new uint256[](3);
        for (uint256 i = 0; i < 3; ++i) {
            sharesBefore[i] = strategyManager.stakerStrategyShares(staker, strategies[i]);
            require(sharesBefore[i] == amounts[i], "Staker has not deposited amount into strategy");
            require(_isDepositedStrategy(staker, strategies[i]), "strategy should be deposited");
        }

        delegationManagerMock.removeShares(strategyManager, staker, removeStrategy, removeAmount);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, removeStrategy);
        require(
            stakerStrategyListLengthAfter == stakerStrategyListLengthBefore - 1,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore - 1"
        );
        require(sharesAfter == 0, "sharesAfter != 0");
        require(!_isDepositedStrategy(staker, removeStrategy), "strategy should not be part of staker strategy list");
    }

    // removeShares() from staker strategy list with multiple strategies deposited. Only callable by DelegationManager
    function testRemoveShares(uint256[3] memory depositAmounts, uint256[3] memory sharesAmounts) external {
        address staker = address(this);
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = dummyStrat;
        strategies[1] = dummyStrat2;
        strategies[2] = dummyStrat3;
        uint256[] memory sharesBefore = new uint256[](3);
        for (uint256 i = 0; i < 3; ++i) {
            cheats.assume(sharesAmounts[i] > 0 && sharesAmounts[i] <= depositAmounts[i]);
            _depositIntoStrategySuccessfully(strategies[i], staker, depositAmounts[i]);
            sharesBefore[i] = strategyManager.stakerStrategyShares(staker, strategies[i]);
            require(sharesBefore[i] == depositAmounts[i], "Staker has not deposited amount into strategy");
            require(_isDepositedStrategy(staker, strategies[i]), "strategy should be deposited");
        }
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);

        uint256 numPoppedStrategies = 0;
        uint256[] memory sharesAfter = new uint256[](3);
        for (uint256 i = 0; i < 3; ++i) {
            delegationManagerMock.removeShares(strategyManager, staker, strategies[i], sharesAmounts[i]);
            sharesAfter[i] = strategyManager.stakerStrategyShares(staker, strategies[i]);
            if (sharesAmounts[i] == depositAmounts[i]) {
                ++numPoppedStrategies;
                require(
                    !_isDepositedStrategy(staker, strategies[i]),
                    "strategy should not be part of staker strategy list"
                );
                require(sharesAfter[i] == 0, "sharesAfter != 0");
            } else {
                require(_isDepositedStrategy(staker, strategies[i]), "strategy should be part of staker strategy list");
                require(
                    sharesAfter[i] == sharesBefore[i] - sharesAmounts[i],
                    "sharesAfter != sharesBefore - sharesAmounts"
                );
            }
        }
        require(
            stakerStrategyListLengthBefore - numPoppedStrategies == strategyManager.stakerStrategyListLength(staker),
            "stakerStrategyListLengthBefore - numPoppedStrategies != strategyManager.stakerStrategyListLength(staker)"
        );
    }
}
