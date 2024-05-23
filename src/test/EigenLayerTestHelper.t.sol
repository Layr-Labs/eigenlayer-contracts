// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../test/EigenLayerDeployer.t.sol";
import "../contracts/interfaces/ISignatureUtils.sol";

contract EigenLayerTestHelper is EigenLayerDeployer {
    uint8 durationToInit = 2;
    uint256 public SECP256K1N_MODULUS = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141;
    uint256 public SECP256K1N_MODULUS_HALF = 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0;

    uint256[] sharesBefore;
    uint256[] balanceBefore;
    uint256[] priorTotalShares;
    uint256[] strategyTokenBalance;

    /**
     * @notice Helper function to test `initiateDelegation` functionality.  Handles registering as an operator, depositing tokens
     * into both Weth and Eigen strategies, as well as delegating assets from "stakers" to the operator.
     * @param operatorIndex is the index of the operator to use from the test-data/operators.json file
     * @param amountEigenToDeposit amount of eigen token to deposit
     * @param amountEthToDeposit amount of eth to deposit
     */

    function _testInitiateDelegation(
        uint8 operatorIndex,
        uint256 amountEigenToDeposit,
        uint256 amountEthToDeposit
    ) public returns (uint256 amountEthStaked, uint256 amountEigenStaked) {
        address operator = getOperatorAddress(operatorIndex);

        //setting up operator's delegation terms
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            __deprecated_earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        _testRegisterAsOperator(operator, operatorDetails);

        for (uint256 i; i < stakers.length; i++) {
            //initialize weth, eigen and eth balances for staker
            eigenToken.transfer(stakers[i], amountEigenToDeposit);
            weth.transfer(stakers[i], amountEthToDeposit);

            //deposit staker's eigen and weth into strategy manager
            _testDepositEigen(stakers[i], amountEigenToDeposit);
            _testDepositWeth(stakers[i], amountEthToDeposit);

            //delegate the staker's deposits to operator
            uint256 operatorEigenSharesBefore = delegation.operatorShares(operator, eigenStrat);
            uint256 operatorWETHSharesBefore = delegation.operatorShares(operator, wethStrat);
            _testDelegateToOperator(stakers[i], operator);
            //verify that `increaseOperatorShares` worked
            assertTrue(
                delegation.operatorShares(operator, eigenStrat) - operatorEigenSharesBefore == amountEigenToDeposit
            );
            assertTrue(delegation.operatorShares(operator, wethStrat) - operatorWETHSharesBefore == amountEthToDeposit);
        }
        amountEthStaked += delegation.operatorShares(operator, wethStrat);
        amountEigenStaked += delegation.operatorShares(operator, eigenStrat);

        return (amountEthStaked, amountEigenStaked);
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
    ) internal {
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
        // deposits will revert when amountToDeposit is 0
        cheats.assume(amountToDeposit > 0);

        // whitelist the strategy for deposit, in case it wasn't before
        {
            cheats.startPrank(strategyManager.strategyWhitelister());
            IStrategy[] memory _strategy = new IStrategy[](1);
            bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);
            _strategy[0] = stratToDepositTo;
            strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
            cheats.stopPrank();
        }

        uint256 operatorSharesBefore = strategyManager.stakerStrategyShares(sender, stratToDepositTo);
        // assumes this contract already has the underlying token!
        uint256 contractBalance = underlyingToken.balanceOf(address(this));
        // check the expected output
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

            //check if depositor has never used this strat, that it is added correctly to stakerStrategyList array.
            if (operatorSharesBefore == 0) {
                // check that strategy is appropriately added to dynamic array of all of sender's strategies
                assertTrue(
                    strategyManager.stakerStrategyList(sender, strategyManager.stakerStrategyListLength(sender) - 1) ==
                        stratToDepositTo,
                    "_testDepositToStrategy: stakerStrategyList array updated incorrectly"
                );
            }

            // check that the shares out match the expected amount out
            assertEq(
                strategyManager.stakerStrategyShares(sender, stratToDepositTo) - operatorSharesBefore,
                expectedSharesOut,
                "_testDepositToStrategy: actual shares out should match expected shares out"
            );
        }
        cheats.stopPrank();
    }

    /**
     * @notice tries to delegate from 'staker' to 'operator', verifies that staker has at least some shares
     * delegatedShares update correctly for 'operator' and delegated status is updated correctly for 'staker'
     * @param staker the staker address to delegate from
     * @param operator the operator address to delegate to
     */
    function _testDelegateToOperator(address staker, address operator) internal {
        //staker-specific information
        (IStrategy[] memory delegateStrategies, uint256[] memory delegateShares) = strategyManager.getDeposits(staker);

        uint256 numStrats = delegateShares.length;
        assertTrue(numStrats != 0, "_testDelegateToOperator: delegating from address with no deposits");
        uint256[] memory inititalSharesInStrats = new uint256[](numStrats);
        for (uint256 i = 0; i < numStrats; ++i) {
            inititalSharesInStrats[i] = delegation.operatorShares(operator, delegateStrategies[i]);
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
            uint256 operatorSharesBefore = inititalSharesInStrats[i];
            uint256 operatorSharesAfter = delegation.operatorShares(operator, delegateStrategies[i]);
            assertTrue(
                operatorSharesAfter == (operatorSharesBefore + delegateShares[i]),
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

    /**
     * @notice Creates a queued withdrawal from `staker`. Begins by registering the staker as a delegate (if specified), then deposits `amountToDeposit`
     * into the WETH strategy, and then queues a withdrawal using
     * `strategyManager.queueWithdrawal(strategyIndexes, strategyArray, tokensArray, shareAmounts, withdrawer)`
     * @notice After initiating a queued withdrawal, this test checks that `strategyManager.canCompleteQueuedWithdrawal` immediately returns the correct
     * response depending on whether `staker` is delegated or not.
     * @param staker The address to initiate the queued withdrawal
     * @param registerAsOperator If true, `staker` will also register as a delegate in the course of this function
     * @param amountToDeposit The amount of WETH to deposit
     */
    function _createQueuedWithdrawal(
        address staker,
        bool registerAsOperator,
        uint256 amountToDeposit,
        IStrategy[] memory strategyArray,
        uint256[] memory shareAmounts,
        uint256[] memory strategyIndexes,
        address withdrawer
    ) internal returns (bytes32 withdrawalRoot, IDelegationManager.Withdrawal memory queuedWithdrawal) {
        require(amountToDeposit >= shareAmounts[0], "_createQueuedWithdrawal: sanity check failed");

        // we do this here to ensure that `staker` is delegated if `registerAsOperator` is true
        if (registerAsOperator) {
            assertTrue(!delegation.isDelegated(staker), "_createQueuedWithdrawal: staker is already delegated");
            IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
                __deprecated_earningsReceiver: staker,
                delegationApprover: address(0),
                stakerOptOutWindowBlocks: 0
            });
            _testRegisterAsOperator(staker, operatorDetails);
            assertTrue(
                delegation.isDelegated(staker),
                "_createQueuedWithdrawal: staker isn't delegated when they should be"
            );
        }

        queuedWithdrawal = IDelegationManager.Withdrawal({
            strategies: strategyArray,
            shares: shareAmounts,
            staker: staker,
            withdrawer: withdrawer,
            nonce: delegation.cumulativeWithdrawalsQueued(staker),
            delegatedTo: delegation.delegatedTo(staker),
            startBlock: uint32(block.number)
        });

        {
            //make deposit in WETH strategy
            uint256 amountDeposited = _testDepositWeth(staker, amountToDeposit);
            // We can't withdraw more than we deposit
            if (shareAmounts[0] > amountDeposited) {
                cheats.expectRevert("StrategyManager._removeShares: shareAmount too high");
            }
        }

        //queue the withdrawal
        withdrawalRoot = _testQueueWithdrawal(staker, strategyIndexes, strategyArray, shareAmounts, withdrawer);
        return (withdrawalRoot, queuedWithdrawal);
    }

    /**
     * Helper for ECDSA signatures: combines V and S into VS - if S is greater than SECP256K1N_MODULUS_HALF, then we
     * get the modulus, so that the leading bit of s is always 0.  Then we set the leading
     * bit to be either 0 or 1 based on the value of v, which is either 27 or 28
     */
    function getVSfromVandS(uint8 v, bytes32 s) internal view returns (bytes32) {
        if (uint256(s) > SECP256K1N_MODULUS_HALF) {
            s = bytes32(SECP256K1N_MODULUS - uint256(s));
        }

        bytes32 vs = s;
        if (v == 28) {
            vs = bytes32(uint256(s) ^ (1 << 255));
        }

        return vs;
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
                __deprecated_earningsReceiver: operator,
                delegationApprover: address(0),
                stakerOptOutWindowBlocks: 0
            });
            _testRegisterAsOperator(operator, operatorDetails);
        }

        uint256 amountBefore = delegation.operatorShares(operator, wethStrat);

        //making additional deposits to the strategies
        assertTrue(!delegation.isDelegated(staker) == true, "testDelegation: staker is not delegate");
        _testDepositWeth(staker, ethAmount);
        _testDepositEigen(staker, eigenAmount);
        _testDelegateToOperator(staker, operator);
        assertTrue(delegation.isDelegated(staker) == true, "testDelegation: staker is not delegate");

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
        address depositor,
        uint256[] memory /*strategyIndexes*/,
        IStrategy[] memory strategyArray,
        uint256[] memory shareAmounts,
        address withdrawer
    ) internal returns (bytes32) {
        cheats.startPrank(depositor);

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
