// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/unit/StrategyManagerUnit/StrategyManagerUnit.t.sol";

contract StrategyManagerUnitTests_addShares is StrategyManagerUnitTests {
    function testAddSharesRevertsDelegationManagerModifier() external {
        DelegationManagerMock invalidDelegationManager = new DelegationManagerMock();
        cheats.expectRevert(bytes("StrategyManager.onlyDelegationManager: not the DelegationManager"));
        invalidDelegationManager.addShares(strategyManager, address(this), dummyStrat, 1);
    }

    function testAddSharesRevertsStakerZeroAddress(uint256 amount) external {
        cheats.expectRevert(bytes("StrategyManager._addShares: staker cannot be zero address"));
        delegationManagerMock.addShares(strategyManager, address(0), dummyStrat, amount);
    }

    function testAddSharesRevertsZeroShares(address staker) external {
        cheats.assume(staker != address(0));
        cheats.expectRevert(bytes("StrategyManager._addShares: shares should not be zero!"));
        delegationManagerMock.addShares(strategyManager, staker, dummyStrat, 0);
    }

    function testAddSharesAppendsStakerStrategyList(address staker, uint256 amount) external {
        cheats.assume(staker != address(0) && amount != 0);
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, dummyStrat);
        require(sharesBefore == 0, "Staker has already deposited into this strategy");
        require(!_isDepositedStrategy(staker, dummyStrat), "strategy shouldn't be deposited");

        delegationManagerMock.addShares(strategyManager, staker, dummyStrat, amount);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, dummyStrat);
        require(
            stakerStrategyListLengthAfter == stakerStrategyListLengthBefore + 1,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore + 1"
        );
        require(sharesAfter == amount, "sharesAfter != amount");
        require(_isDepositedStrategy(staker, dummyStrat), "strategy should be deposited");
    }

    function testAddSharesExistingShares(address staker, uint256 sharesAmount) external {
        cheats.assume(staker != address(0) && 0 < sharesAmount && sharesAmount <= dummyToken.totalSupply());
        uint256 initialAmount = 1e18;
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, initialAmount);
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, dummyStrat);
        require(sharesBefore == initialAmount, "Staker has not deposited into strategy");
        require(_isDepositedStrategy(staker, strategy), "strategy should be deposited");

        delegationManagerMock.addShares(strategyManager, staker, dummyStrat, sharesAmount);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, dummyStrat);
        require(
            stakerStrategyListLengthAfter == stakerStrategyListLengthBefore,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore"
        );
        require(sharesAfter == sharesBefore + sharesAmount, "sharesAfter != sharesBefore + amount");
        require(_isDepositedStrategy(staker, strategy), "strategy should be deposited");
    }

    function test_addSharesRevertsWhenSharesIsZero() external {
        // replace dummyStrat with Reenterer contract
        reenterer = new Reenterer();
        dummyStrat = StrategyBase(address(reenterer));

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(dummyStrat);
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        address staker = address(this);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        uint256 amount = 1e18;

        reenterer.prepareReturnData(abi.encode(uint256(0)));

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("StrategyManager._addShares: shares should not be zero!"));
        strategyManager.depositIntoStrategy(strategy, token, amount);
        cheats.stopPrank();
    }

    function test_addSharesRevertsWhenDepositWouldExeedMaxArrayLength() external {
        address staker = address(this);
        IERC20 token = dummyToken;
        uint256 amount = 1e18;
        IStrategy strategy = dummyStrat;

        // uint256 MAX_STAKER_STRATEGY_LIST_LENGTH = strategyManager.MAX_STAKER_STRATEGY_LIST_LENGTH();
        uint256 MAX_STAKER_STRATEGY_LIST_LENGTH = 32;

        // loop that deploys a new strategy and deposits into it
        for (uint256 i = 0; i < MAX_STAKER_STRATEGY_LIST_LENGTH; ++i) {
            cheats.startPrank(staker);
            strategyManager.depositIntoStrategy(strategy, token, amount);
            cheats.stopPrank();

            dummyStrat = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
            strategy = dummyStrat;

            // whitelist the strategy for deposit
            cheats.startPrank(strategyManager.owner());
            IStrategy[] memory _strategy = new IStrategy[](1);
            _strategy[0] = dummyStrat;
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(dummyStrat);
            strategyManager.addStrategiesToDepositWhitelist(_strategy);
            cheats.stopPrank();
        }

        require(
            strategyManager.stakerStrategyListLength(staker) == MAX_STAKER_STRATEGY_LIST_LENGTH,
            "strategyManager.stakerStrategyListLength(staker) != MAX_STAKER_STRATEGY_LIST_LENGTH"
        );

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("StrategyManager._addShares: deposit would exceed MAX_STAKER_STRATEGY_LIST_LENGTH"));
        strategyManager.depositIntoStrategy(strategy, token, amount);
        cheats.stopPrank();
    }
}
