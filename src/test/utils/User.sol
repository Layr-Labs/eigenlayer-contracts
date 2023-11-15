// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../EigenLayerDeployer.t.sol";

contract User is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    EigenLayerDeployer public immutable globalState;
    address public immutable user;


    // TODO: decide if the following storage is useful
    // list of all _uncompleted_ withdrawals created by the user
    IDelegationManager.Withdrawal[] public queuedWithdrawals;

    // list of all stakers who are delegated to the user, _if_ they are an operator. Should include the Operator themself.
    address[] public delegatedStakers;
    // TODO: info about AVS registrations

    // testing/mock contracts
    // IERC20 public eigenToken;
    // IERC20 public weth;
    // StrategyBase public wethStrat;
    // StrategyBase public eigenStrat;
    // StrategyBase public baseStrategyImplementation;
    // EmptyContract public emptyContract;

    constructor(EigenLayerDeployer globalStateContract) {
        globalState = globalStateContract;
        user = address(this);
    }

    function eigenLayerProxyAdmin() public view returns(ProxyAdmin) {
        return globalState.eigenLayerProxyAdmin();
    }
    function eigenLayerPauserReg() public view returns(PauserRegistry) {
        return globalState.eigenLayerPauserReg();
    }
    function slasher() public view returns(Slasher) {
        return globalState.slasher();
    }
    function delegation() public view returns(DelegationManager) {
        return globalState.delegation();
    }
    function strategyManager() public view returns(StrategyManager) {
        return globalState.strategyManager();
    }
    function eigenPodManager() public view returns(EigenPodManager) {
        return globalState.eigenPodManager();
    }
    function pod() public view returns(IEigenPod) {
        return globalState.pod();
    }
    function delayedWithdrawalRouter() public view returns(IDelayedWithdrawalRouter) {
        return globalState.delayedWithdrawalRouter();
    }
    function ethPOSDeposit() public view returns(IETHPOSDeposit) {
        return globalState.ethPOSDeposit();
    }
    function eigenPodBeacon() public view returns(IBeacon) {
        return globalState.eigenPodBeacon();
    }

    function strategyInStakerList(IStrategy strategy) public view returns (bool) {
        (IStrategy[] memory strategies, /* uint256[] memory shares */) = strategyManager().getDeposits(user);
        for (uint256 i = 0; i < strategies.length; ++i) {
            if (strategies[i] == strategy) {
                return true;
            }
        }
        return false;
    }

    function stakerShares(IStrategy strategy) public view returns (uint256) {
        (IStrategy[] memory strategies, uint256[] memory shares) = strategyManager().getDeposits(user);
        for (uint256 i = 0; i < strategies.length; ++i) {
            if (strategies[i] == strategy) {
                return shares[i];
            }
        }
        return 0;
    }

    /*
     * @notice Register this contract as an operator, setting their 'OperatorDetails' in DelegationManager to 'operatorDetails', verifies
     * that the storage of DelegationManager contract is updated appropriately
     */
    function registerAsOperator(
        IDelegationManager.OperatorDetails memory operatorDetails
    ) public {
        string memory emptyStringForMetadataURI;

        cheats.prank(user);
        delegation().registerAsOperator(operatorDetails, emptyStringForMetadataURI);

        assertTrue(delegation().isOperator(user), "testRegisterAsOperator: user is not a operator");
        assertTrue(
            keccak256(abi.encode(delegation().operatorDetails(user))) == keccak256(abi.encode(operatorDetails)),
            "_testRegisterAsOperator: operatorDetails not set appropriately"
        );
        assertTrue(delegation().isDelegated(user), "_testRegisterAsOperator: user not marked as actively delegated");
    }

    /**
     * @notice Default version of `registerAsOperator`. Does not set a `delegationApprover`.
     */
    function registerAsOperator() public {
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: user,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: uint32(0)
        });
        
        registerAsOperator(operatorDetails);
    }

    // @notice Deposits `amountToDeposit` of `underlyingToken` from this contract into `stratToDepositTo`.
    function depositToStrategy(
        uint256 amountToDeposit,
        IERC20 underlyingToken,
        IStrategy stratToDepositTo
    ) internal {
        // deposits will revert when amountToDeposit is 0
        require(amountToDeposit != 0, "bad usage of depositToStrategy helper function");

        // whitelist the strategy for deposit, in the case where it wasn't before
        {
            if (!strategyManager().strategyIsWhitelistedForDeposit(stratToDepositTo)) {
                IStrategy[] memory _strategy = new IStrategy[](1);
                _strategy[0] = stratToDepositTo;
                cheats.prank(strategyManager().strategyWhitelister());
                strategyManager().addStrategiesToDepositWhitelist(_strategy);
            }
        }

        // store prior state
        uint256 stakerSharesBefore = strategyManager().stakerStrategyShares(user, stratToDepositTo);
        uint256 stakerStrategyListLengthBefore = strategyManager().stakerStrategyListLength(user);
        if (stakerSharesBefore != 0) {
            assertTrue(strategyInStakerList(stratToDepositTo), "strategy somehow not in staker strategy array");
        }
        // calculate the expected output
        uint256 expectedSharesOut = stratToDepositTo.underlyingToShares(amountToDeposit);

        // perform the action itself
        // assumes this contract already has the underlying token!
        uint256 contractBalance = underlyingToken.balanceOf(address(this));
        // logging and error for misusing this function (see assumption above)
        if (amountToDeposit > contractBalance) {
            emit log("amountToDeposit > contractBalance");
            emit log_named_uint("amountToDeposit is", amountToDeposit);
            emit log_named_uint("while contractBalance is", contractBalance);
            revert("depositToStrategy test helper misuse");
        } else {
            underlyingToken.transfer(user, amountToDeposit);
            cheats.startPrank(user);
            underlyingToken.approve(address(strategyManager()), type(uint256).max);
            strategyManager().depositIntoStrategy(stratToDepositTo, underlyingToken, amountToDeposit);
            cheats.stopPrank();
        }

        // fetch post-action state
        uint256 stakerSharesAfter = strategyManager().stakerStrategyShares(user, stratToDepositTo);
        uint256 stakerStrategyListLengthAfter = strategyManager().stakerStrategyListLength(user);

        // check if staker had zero shares before, that it is added correctly to stakerStrategyList array.
        if (stakerSharesBefore == 0) {
            assertTrue(
                strategyManager().stakerStrategyList(user, stakerStrategyListLengthAfter - 1) ==
                    stratToDepositTo,
                "_testDepositToStrategy: stakerStrategyList array updated incorrectly"
            );
            assertEq(stakerStrategyListLengthAfter, stakerStrategyListLengthBefore + 1,
                "strategy list did not update correctly");
        // otherwise check that the list length was unmodified
        } else {
            assertEq(stakerStrategyListLengthAfter, stakerStrategyListLengthBefore,
                "strategy list updated incorrectly");
        }

        // check that the shares difference matches the expected amount out
        assertEq(
            stakerSharesAfter - stakerSharesBefore, expectedSharesOut,
            "_testDepositToStrategy: actual shares out should match expected shares out"
        );
        assertTrue(strategyInStakerList(stratToDepositTo), "strategy somehow not in staker strategy array");
    }

    // function _testQueueWithdrawal(
    //     Staker memory stakerStateBefore,
    //     IStrategy[] memory strategies,
    //     uint256[] memory shares,
    //     address withdrawer
    // ) internal returns (Staker memory stakerStateAfter) {
    //     // prepare struct for call
    //     IDelegationManager.QueuedWithdrawalParams[] memory withdrawalParams =
    //         new IDelegationManager.QueuedWithdrawalParams[](1);
    //     withdrawalParams[0] = IDelegationManager.QueuedWithdrawalParams({
    //         strategies: strategies,
    //         shares: shares,
    //         withdrawer: withdrawer
    //     });

    //     // actually make the call to queue the withdrawal
    //     cheats.prank(stakerStateBefore.staker);
    //     delegation.queueWithdrawals(withdrawalParams);

    //     // update the Staker struct appropriately
    //     stakerStateAfter = _updateStakerState(stakerStateBefore);
    //     stakerStateAfter.queuedWithdrawals =
    //         new IDelegationManager.Withdrawal[](stakerStateBefore.queuedWithdrawals.length + 1);
    //     for (uint256 i = 0; i < stakerStateBefore.queuedWithdrawals.length; ++i) {
    //         stakerStateAfter.queuedWithdrawals[i] = stakerStateBefore.queuedWithdrawals[i];
    //     }
    //     IDelegationManager.Withdrawal memory newWithdrawal = IDelegationManager.Withdrawal({
    //         staker: stakerStateBefore.staker,
    //         delegatedTo: stakerStateBefore.delegatedTo,
    //         withdrawer: withdrawer,
    //         nonce: delegation.stakerNonce(stakerStateBefore.staker),
    //         startBlock: uint32(block.number),
    //         strategies: strategies,
    //         shares: shares
    //     });
    //     stakerStateAfter.queuedWithdrawals[stakerStateAfter.queuedWithdrawals.length - 1] = newWithdrawal;

    //     // return the updated Staker struct
    //     return stakerStateAfter;
    // }

    // // @notice queue a withdrawal from the Staker starting with `stakerStateBefore`, for all of their deposits, to `withdrawer`
    // function _testQueueWithdraw_AllShares(Staker memory stakerStateBefore, address withdrawer)
    //     internal returns (Staker memory stakerStateAfter)
    // {
    //     return _testQueueWithdrawal({
    //         stakerStateBefore: stakerStateBefore,
    //         strategies: stakerStateBefore.strategies,
    //         shares: stakerStateBefore.shares,
    //         withdrawer: withdrawer
    //     });
    // }

}
