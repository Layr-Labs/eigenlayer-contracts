// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../test/EigenLayerDeployer.t.sol";
import "../contracts/interfaces/ISignatureUtils.sol";

contract EigenLayerTestHelper is EigenLayerDeployer {
    uint256[] sharesBefore;
    uint256[] balanceBefore;
    uint256[] priorTotalShares;
    uint256[] strategyTokenBalance;

    // takes in a Staker struct and gets the staker's updated state
    function _updateStakerState(Staker memory staker) internal view returns (Staker memory) {
        (IStrategy[] memory stakerStrategies, uint256[] memory stakerShares) = strategyManager.getDeposits(staker.staker);
        address delegatedTo = delegation.delegatedTo(staker.staker);
        Staker memory updatedState = Staker({
            // staker's address should never change
            staker: staker.staker,
            strategies: stakerStrategies,
            shares: stakerShares,
            delegatedTo: delegatedTo,
            // don't change this -- update the queued withdrawals elsewhere, when applicable
            queuedWithdrawals: staker.queuedWithdrawals
        });

        // TODO: add some invariant checks?
        return updatedState;
    }

    function _strategyInStakerList(Staker memory staker, IStrategy strategy) internal pure returns (bool) {
        for (uint256 i = 0; i < staker.strategies.length; ++i) {
            if (staker.strategies[i] == strategy) {
                return true;
            }
        }
        return false;
    }

    function _stakerShares(Staker memory staker, IStrategy strategy)  internal pure returns (uint256) {
        for (uint256 i = 0; i < staker.strategies.length; ++i) {
            if (staker.strategies[i] == strategy) {
                return staker.shares[i];
            }
        }
        return 0;
    }

    /**
     * @notice Register 'sender' as an operator, setting their 'OperatorDetails' in DelegationManager to 'operatorDetails', verifies
     * that the storage of DelegationManager contract is updated appropriately
     *
     * @param sender is the address being registered as an operator
     * @param operatorDetails is the `sender`'s OperatorDetails struct
     */
    function _testRegisterAsOperator(
        address sender,
        IDelegationManager.OperatorDetails memory operatorDetails
    ) internal returns (Operator memory) {
        cheats.startPrank(sender);
        string memory emptyStringForMetadataURI;
        delegation.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        assertTrue(delegation.isOperator(sender), "testRegisterAsOperator: sender is not a operator");

        assertTrue(
            keccak256(abi.encode(delegation.operatorDetails(sender))) == keccak256(abi.encode(operatorDetails)),
            "_testRegisterAsOperator: operatorDetails not set appropriately"
        );

        assertTrue(delegation.isDelegated(sender), "_testRegisterAsOperator: sender not marked as actively delegated");
        cheats.stopPrank();

        Operator memory operatorStruct = Operator({
            operator: sender,
            operatorDetails: operatorDetails,
            delegatedStakers: new address[](1)
        });
        operatorStruct.delegatedStakers[0] = sender;

        return operatorStruct;
    }

    // @notice Deposits `amountToDeposit` of `underlyingToken` from `staker` into `stratToDepositTo`.
    function _testDepositToStrategy(
        Staker memory staker,
        uint256 amountToDeposit,
        IERC20 underlyingToken,
        IStrategy stratToDepositTo
    ) internal returns (Staker memory stakerStateAfter) {
        // deposits will revert when amountToDeposit is 0
        require(amountToDeposit != 0, "bad usage of _testDepositToStrategy helper function");

        // whitelist the strategy for deposit, in the case where it wasn't before
        {
            if (!strategyManager.strategyIsWhitelistedForDeposit(stratToDepositTo)) {
                IStrategy[] memory _strategy = new IStrategy[](1);
                _strategy[0] = stratToDepositTo;
                cheats.prank(strategyManager.strategyWhitelister());
                strategyManager.addStrategiesToDepositWhitelist(_strategy);
            }
        }

        // assumes this contract already has the underlying token!
        uint256 contractBalance = underlyingToken.balanceOf(address(this));
        // calculate the expected output
        uint256 expectedSharesOut = stratToDepositTo.underlyingToShares(amountToDeposit);
        // logging and error for misusing this function (see assumption above)
        if (amountToDeposit > contractBalance) {
            emit log("amountToDeposit > contractBalance");
            emit log_named_uint("amountToDeposit is", amountToDeposit);
            emit log_named_uint("while contractBalance is", contractBalance);
            revert("_testDepositToStrategy failure");
        } else {
            underlyingToken.transfer(staker.staker, amountToDeposit);
            cheats.startPrank(staker.staker);
            underlyingToken.approve(address(strategyManager), type(uint256).max);
            strategyManager.depositIntoStrategy(stratToDepositTo, underlyingToken, amountToDeposit);
            cheats.stopPrank();

            stakerStateAfter = _updateStakerState(staker);

            // staker had zero shares in the strategy before, check that it is added correctly to stakerStrategyList array.
            if (_stakerShares(staker, stratToDepositTo) == 0) {
                // check that strategy is appropriately added to dynamic array of all of staker's strategies
                assertTrue(
                    stakerStateAfter.strategies[stakerStateAfter.strategies.length - 1] == stratToDepositTo,
                    "_testDepositToStrategy: stakerStrategyList array updated incorrectly"
                );
                assertEq(stakerStateAfter.strategies.length, staker.strategies.length + 1,
                    "strategy list did not update correctly");
            } else {
                assertTrue(_strategyInStakerList(staker, stratToDepositTo), "strategy somehow not in staker strategy array");
                assertTrue(_strategyInStakerList(stakerStateAfter, stratToDepositTo), "strategy somehow not in staker strategy array");
                assertEq(stakerStateAfter.strategies.length, staker.strategies.length,
                    "strategy list updated incorrectly");
            }

            // check that the shares difference matches the expected amount out
            assertEq(
                _stakerShares(stakerStateAfter, stratToDepositTo) - _stakerShares(staker, stratToDepositTo),
                expectedSharesOut,
                "_testDepositToStrategy: actual shares out should match expected shares out"
            );
        }
        return stakerStateAfter;
    }

    /**
     * @notice Deposits `amountToDeposit` of WETH from `staker` into `wethStrat`.
     * @param amountToDeposit Amount of WETH that is first *transferred from this contract to `staker`* and then deposited by `staker` into `wethStrat`
     */
    function _testDepositWeth(Staker memory staker, uint256 amountToDeposit) internal returns (Staker memory /*stakerStateAfter*/) {
        return _testDepositToStrategy(staker, amountToDeposit, weth, wethStrat);
    }






    /**
     * @notice Deposits `amountToDeposit` of WETH from address `sender` into `wethStrat`.
     * @param sender The address to spoof calls from using `cheats.startPrank(sender)`
     * @param amountToDeposit Amount of WETH that is first *transferred from this contract to `sender`* and then deposited by `sender` into `stratToDepositTo`
     */
    function _testDepositWeth(address sender, uint256 amountToDeposit) internal returns (uint256 amountDeposited) {
        cheats.assume(amountToDeposit <= wethInitialSupply);
        amountDeposited = _testDepositToStrategy(sender, amountToDeposit, weth, wethStrat);
    }

    /**
     * @notice Deposits `amountToDeposit` of EIGEN from address `sender` into `eigenStrat`.
     * @param sender The address to spoof calls from using `cheats.startPrank(sender)`
     * @param amountToDeposit Amount of EIGEN that is first *transferred from this contract to `sender`* and then deposited by `sender` into `stratToDepositTo`
     */
    function _testDepositEigen(address sender, uint256 amountToDeposit) internal returns (uint256 amountDeposited) {
        cheats.assume(amountToDeposit <= eigenTotalSupply);
        amountDeposited = _testDepositToStrategy(sender, amountToDeposit, eigenToken, eigenStrat);
    }

    /**
     * @notice Deposits `amountToDeposit` of `underlyingToken` from address `sender` into `stratToDepositTo`.
     * *If*  `sender` has zero shares prior to deposit, *then* checks that `stratToDepositTo` is correctly added to their `stakerStrategyList` array.
     *
     * @param sender The address to spoof calls from using `cheats.startPrank(sender)`
     * @param amountToDeposit Amount of WETH that is first *transferred from this contract to `sender`* and then deposited by `sender` into `stratToDepositTo`
     */
    function _testDepositToStrategy(
        address sender,
        uint256 amountToDeposit,
        IERC20 underlyingToken,
        IStrategy stratToDepositTo
    ) internal returns (uint256 amountDeposited) {
        // TODO: modify this to be a require statement instead? would probably be better to do the filtering on a higher level.
        // deposits will revert when amountToDeposit is 0
        cheats.assume(amountToDeposit > 0);

        // whitelist the strategy for deposit, in the case where it wasn't before
        {
            if (!strategyManager.strategyIsWhitelistedForDeposit(stratToDepositTo)) {
                IStrategy[] memory _strategy = new IStrategy[](1);
                _strategy[0] = stratToDepositTo;
                cheats.prank(strategyManager.strategyWhitelister());
                strategyManager.addStrategiesToDepositWhitelist(_strategy);
            }
        }

        uint256 stakerSharesBefore = strategyManager.stakerStrategyShares(sender, stratToDepositTo);
        // assumes this contract already has the underlying token!
        uint256 contractBalance = underlyingToken.balanceOf(address(this));
        // calculate the expected output
        uint256 expectedSharesOut = stratToDepositTo.underlyingToShares(amountToDeposit);
        // logging and error for misusing this function (see assumption above)
        if (amountToDeposit > contractBalance) {
            emit log("amountToDeposit > contractBalance");
            emit log_named_uint("amountToDeposit is", amountToDeposit);
            emit log_named_uint("while contractBalance is", contractBalance);
            revert("_testDepositToStrategy failure");
        } else {
            underlyingToken.transfer(sender, amountToDeposit);
            cheats.startPrank(sender);
            underlyingToken.approve(address(strategyManager), type(uint256).max);
            strategyManager.depositIntoStrategy(stratToDepositTo, underlyingToken, amountToDeposit);
            amountDeposited = amountToDeposit;
            cheats.stopPrank();

            // check if staker had zero shares before, that it is added correctly to stakerStrategyList array.
            if (stakerSharesBefore == 0) {
                // check that strategy is appropriately added to dynamic array of all of sender's strategies
                assertTrue(
                    strategyManager.stakerStrategyList(sender, strategyManager.stakerStrategyListLength(sender) - 1) ==
                        stratToDepositTo,
                    "_testDepositToStrategy: stakerStrategyList array updated incorrectly"
                );
            }

            // check that the shares output matches the expected amount out
            assertEq(
                strategyManager.stakerStrategyShares(sender, stratToDepositTo) - stakerSharesBefore,
                expectedSharesOut,
                "_testDepositToStrategy: actual shares out should match expected shares out"
            );
        }
    }

    /**
     * @notice tries to delegate from 'staker' to 'operator', verifies that staker has at least some shares
     * delegatedShares update correctly for 'operator' and delegated status is updated correctly for 'staker'
     * @param staker the staker address to delegate from
     * @param operator the operator address to delegate to
     */
    function _testDelegateToOperator(address staker, address operator) internal {
        //staker-specific information
        (IStrategy[] memory stakerStrategies, uint256[] memory stakerShares) = strategyManager.getDeposits(staker);

        uint256 numStrats = stakerShares.length;
        assertTrue(numStrats != 0, "_testDelegateToOperator: delegating from address with no deposits");
        uint256[] memory inititalSharesInStrats = new uint256[](numStrats);
        for (uint256 i = 0; i < numStrats; ++i) {
            inititalSharesInStrats[i] = delegation.operatorShares(operator, stakerStrategies[i]);
        }

        cheats.startPrank(staker);
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry;
        delegation.delegateTo(operator, signatureWithExpiry, bytes32(0));
        cheats.stopPrank();

        assertTrue(
            delegation.delegatedTo(staker) == operator,
            "_testDelegateToOperator: delegated address not set appropriately"
        );
        assertTrue(delegation.isDelegated(staker), "_testDelegateToOperator: delegated status not set appropriately");

        for (uint256 i = 0; i < numStrats; ++i) {
            uint256 stakerSharesBefore = inititalSharesInStrats[i];
            uint256 operatorSharesAfter = delegation.operatorShares(operator, stakerStrategies[i]);
            assertTrue(
                operatorSharesAfter == (stakerSharesBefore + stakerShares[i]),
                "_testDelegateToOperator: delegatedShares not increased correctly"
            );
        }
    }

    /**
     * @notice deploys 'numStratsToAdd' strategies contracts and initializes them to treat `underlyingToken` as their underlying token
     * and then deposits 'amountToDeposit' to each of them from 'sender'
     *
     * @param sender address that is depositing into the strategies
     * @param amountToDeposit amount being deposited
     * @param numStratsToAdd number of strategies that are being deployed and deposited into
     */
    function _testDepositStrategies(address sender, uint256 amountToDeposit, uint8 numStratsToAdd) internal {
        // hard-coded input
        IERC20 underlyingToken = weth;

        cheats.assume(numStratsToAdd > 0 && numStratsToAdd <= 20);
        IStrategy[] memory stratsToDepositTo = new IStrategy[](numStratsToAdd);
        for (uint8 i = 0; i < numStratsToAdd; ++i) {
            stratsToDepositTo[i] = StrategyBase(
                address(
                    new TransparentUpgradeableProxy(
                        address(baseStrategyImplementation),
                        address(eigenLayerProxyAdmin),
                        abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken, eigenLayerPauserReg)
                    )
                )
            );
            _testDepositToStrategy(sender, amountToDeposit, weth, StrategyBase(address(stratsToDepositTo[i])));
        }
        for (uint8 i = 0; i < numStratsToAdd; ++i) {
            // check that strategy is appropriately added to dynamic array of all of sender's strategies
            assertTrue(
                strategyManager.stakerStrategyList(sender, i) == stratsToDepositTo[i],
                "stakerStrategyList array updated incorrectly"
            );

            // TODO: perhaps remove this is we can. seems brittle if we don't track the number of strategies somewhere
            //store strategy in mapping of strategies
            strategies[i] = IStrategy(address(stratsToDepositTo[i]));
        }
    }

    /// @notice registers a fixed address as an operator, delegates to it from a second address,
    ///         and checks that the operator's voteWeights increase properly
    /// @param operator is the operator being delegated to.
    /// @param staker is the staker delegating stake to the operator.
    /// @param ethAmount is the amount of ETH to deposit into the operator's strategy.
    /// @param eigenAmount is the amount of EIGEN to deposit into the operator's strategy.
    function _testDelegation(
        address operator,
        address staker,
        uint256 ethAmount,
        uint256 eigenAmount
    ) internal {
        if (!delegation.isOperator(operator)) {
            IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
                earningsReceiver: operator,
                delegationApprover: address(0),
                stakerOptOutWindowBlocks: 0
            });
            _testRegisterAsOperator(operator, operatorDetails);
        }

        uint256 amountBefore = delegation.operatorShares(operator, wethStrat);

        //making additional deposits to the strategies
        assertTrue(!delegation.isDelegated(staker) == true, "testDelegation: staker is already delegated");
        _testDepositWeth(staker, ethAmount);
        _testDepositEigen(staker, eigenAmount);
        _testDelegateToOperator(staker, operator);
        assertTrue(delegation.isDelegated(staker) == true, "testDelegation: staker is not delegated");

        (/*IStrategy[] memory updatedStrategies*/, uint256[] memory updatedShares) = strategyManager.getDeposits(staker);

        IStrategy _strat = wethStrat;
        // IStrategy _strat = strategyManager.stakerStrategyList(staker, 0);
        assertTrue(address(_strat) != address(0), "stakerStrategyList not updated correctly");

        assertTrue(
            delegation.operatorShares(operator, _strat) - updatedShares[0] == amountBefore,
            "ETH operatorShares not updated correctly"
        );
    }

    /**
     * @notice Helper function to complete an existing queued withdrawal in shares
     * @param depositor is the address of the staker who queued a withdrawal
     * @param strategyArray is the array of strategies to withdraw from
     * @param tokensArray is the array of tokens to withdraw from said strategies
     * @param shareAmounts is the array of shares to be withdrawn from said strategies
     * @param delegatedTo is the address the staker has delegated their shares to
     * @param withdrawalStartBlock the block number of the original queued withdrawal
     * @param middlewareTimesIndex index in the middlewareTimes array used to queue this withdrawal
     */

    function _testCompleteQueuedWithdrawalShares(
        address depositor,
        IStrategy[] memory strategyArray,
        IERC20[] memory tokensArray,
        uint256[] memory shareAmounts,
        address delegatedTo,
        address withdrawer,
        uint256 nonce,
        uint32 withdrawalStartBlock,
        uint256 middlewareTimesIndex
    ) internal {
        cheats.startPrank(withdrawer);

        for (uint256 i = 0; i < strategyArray.length; i++) {
            sharesBefore.push(strategyManager.stakerStrategyShares(withdrawer, strategyArray[i]));
        }
        // emit log_named_uint("strategies", strategyArray.length);
        // emit log_named_uint("tokens", tokensArray.length);
        // emit log_named_uint("shares", shareAmounts.length);
        // emit log_named_address("depositor", depositor);
        // emit log_named_uint("withdrawalStartBlock", withdrawalStartBlock);
        // emit log_named_address("delegatedAddress", delegatedTo);
        // emit log("************************************************************************************************");

        IDelegationManager.Withdrawal memory queuedWithdrawal = IDelegationManager.Withdrawal({
            strategies: strategyArray,
            shares: shareAmounts,
            staker: depositor,
            withdrawer: withdrawer,
            nonce: nonce,
            startBlock: withdrawalStartBlock,
            delegatedTo: delegatedTo
        });

        // complete the queued withdrawal
        delegation.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, false);

        for (uint256 i = 0; i < strategyArray.length; i++) {
            require(
                strategyManager.stakerStrategyShares(withdrawer, strategyArray[i]) == sharesBefore[i] + shareAmounts[i],
                "_testCompleteQueuedWithdrawalShares: withdrawer shares not incremented"
            );
        }
        cheats.stopPrank();
    }

    /**
     * @notice Helper function to complete an existing queued withdrawal in tokens
     * @param depositor is the address of the staker who queued a withdrawal
     * @param strategyArray is the array of strategies to withdraw from
     * @param tokensArray is the array of tokens to withdraw from said strategies
     * @param shareAmounts is the array of shares to be withdrawn from said strategies
     * @param delegatedTo is the address the staker has delegated their shares to
     * @param withdrawalStartBlock the block number of the original queued withdrawal
     * @param middlewareTimesIndex index in the middlewareTimes array used to queue this withdrawal
     */
    function _testCompleteQueuedWithdrawalTokens(
        address depositor,
        IStrategy[] memory strategyArray,
        IERC20[] memory tokensArray,
        uint256[] memory shareAmounts,
        address delegatedTo,
        address withdrawer,
        uint256 nonce,
        uint32 withdrawalStartBlock,
        uint256 middlewareTimesIndex
    ) internal {
        cheats.startPrank(withdrawer);

        for (uint256 i = 0; i < strategyArray.length; i++) {
            balanceBefore.push(strategyArray[i].underlyingToken().balanceOf(withdrawer));
            priorTotalShares.push(strategyArray[i].totalShares());
            strategyTokenBalance.push(strategyArray[i].underlyingToken().balanceOf(address(strategyArray[i])));
        }

        IDelegationManager.Withdrawal memory queuedWithdrawal = IDelegationManager.Withdrawal({
            strategies: strategyArray,
            shares: shareAmounts,
            staker: depositor,
            withdrawer: withdrawer,
            nonce: nonce,
            startBlock: withdrawalStartBlock,
            delegatedTo: delegatedTo
        });
        // complete the queued withdrawal
        delegation.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, true);

        for (uint256 i = 0; i < strategyArray.length; i++) {
            //uint256 strategyTokenBalance = strategyArray[i].underlyingToken().balanceOf(address(strategyArray[i]));
            uint256 tokenBalanceDelta = (strategyTokenBalance[i] * shareAmounts[i]) / priorTotalShares[i];

            // filter out unrealistic case, where the withdrawer is the strategy contract itself
            cheats.assume(withdrawer != address(strategyArray[i]));
            require(
                strategyArray[i].underlyingToken().balanceOf(withdrawer) == balanceBefore[i] + tokenBalanceDelta,
                "_testCompleteQueuedWithdrawalTokens: withdrawer balance not incremented"
            );
        }
        cheats.stopPrank();
    }

    function _testQueueWithdrawal(
        address staker,
        IStrategy[] memory strategyArray,
        uint256[] memory shareAmounts,
        address withdrawer
    ) internal returns (bytes32) {
        cheats.startPrank(staker);

        IDelegationManager.QueuedWithdrawalParams[] memory params = new IDelegationManager.QueuedWithdrawalParams[](1);

        params[0] = IDelegationManager.QueuedWithdrawalParams({
            strategies: strategyArray,
            shares: shareAmounts,
            withdrawer: withdrawer
        });

        bytes32[] memory withdrawalRoots = new bytes32[](1);
        withdrawalRoots = delegation.queueWithdrawals(params);
        cheats.stopPrank();
        return withdrawalRoots[0];
    }
}
