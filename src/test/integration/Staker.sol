// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";

import "src/contracts/interfaces/IDelegationManager.sol";

import "src/test/integration/GlobalRefs.sol";
import "src/test/integration/Operator.sol";

contract Staker is Test {
    // Pointer to global reference contract
    GlobalRefs globalRefs;

    // Local Storage
    IDelegationManager.Withdrawal[] public queuedWithdrawals;

    // State to help with helper function checks
    IStrategy[] internal _tempStakerStrategies; // array of strategies without duplicates
    mapping(IStrategy => bool) internal _tempStrategyRead;
    mapping(IStrategy => uint256) internal _tempStrategyShares;

    // Staker Address
    address staker = address(this);

    constructor(GlobalRefs _globalRefs) {
        globalRefs = _globalRefs;
    }

    function delegate(Operator operator) public {
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        bytes32 emptySalt;
        _delegate(operator, approverSignatureAndExpiry, emptySalt);

        // Assert that salt is not spent
        assertFalse(globalRefs.delegationManager().delegationApproverSaltIsSpent(
            globalRefs.delegationManager().delegationApprover(address(operator)),
            emptySalt
        ), "salt somehow spent too early");
    }

    function _delegate(Operator operator, ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry, bytes32 approverSalt) internal {
        // Save pre-delegation state of operator 
        (IStrategy[] memory stakerStrategies, uint256[] memory stakerShares) = globalRefs.delegationManager().getDelegatableShares(staker);
        uint256[] memory operatorSharesBefore = new uint256[](stakerShares.length);
        for (uint256 i = 0; i < stakerShares.length; i++) {
            operatorSharesBefore[i] = globalRefs.delegationManager().operatorShares(address(operator), stakerStrategies[i]);
        }

        // Delegate to operator
        globalRefs.delegationManager().delegateTo(address(operator), approverSignatureAndExpiry, approverSalt);

        // Checks
        assertEq(globalRefs.delegationManager().delegatedTo(staker), address(operator), "staker delegated to the wrong address");
        for(uint256 i = 0; i < stakerShares.length; i++) {
            assertEq(globalRefs.delegationManager().operatorShares(address(operator), stakerStrategies[i]), operatorSharesBefore[i] + stakerShares[i], "operator shares not updated correctly");
        }
        
        // Update operator's delegatedStaker array
        operator.addDelegatedStaker(address(this));
    }

    function depositIntoStrategy(IStrategy strategy, IERC20 token, uint256 amount) public returns (uint256 shares){
        // Save pre-deposit state
        uint256 stakerSharesBefore = _getStakerSharesForStrategy(strategy);
        uint256 operatorSharesBefore = _getOperatorSharesForStrategy(strategy);
        uint256 stakerStrategyListLengthBefore = globalRefs.strategyManager().stakerStrategyListLength(staker);

        // Deposit into strategy
        _dealTokenToStaker(token, amount);
        token.approve(address(globalRefs.strategyManager()), amount);
        shares = globalRefs.strategyManager().depositIntoStrategy(strategy, token, amount);

        // Checks
        assertEq(_getStakerSharesForStrategy(strategy), stakerSharesBefore + shares, "staker shares not updated correctly");
        assertEq(_getOperatorSharesForStrategy(strategy), operatorSharesBefore + shares, "operator shares not updated correctly");
        if (stakerSharesBefore == 0) {
            assertEq(globalRefs.strategyManager().stakerStrategyListLength(staker), stakerStrategyListLengthBefore + 1, "strategy not added to staker's strategy list");
            assertEq(
                address(globalRefs.strategyManager().stakerStrategyList(staker, stakerStrategyListLengthBefore)),
                address(strategy),
                "strategy not added to staker's strategy list"
            );        
        } else {
            assertEq(globalRefs.strategyManager().stakerStrategyListLength(staker), stakerStrategyListLengthBefore, "staker's strategy list length changed");
        }
    }

    function queueFullWithdrawal(IStrategy strategy) public {
        queueWithdrawal(strategy, _getStakerSharesForStrategy(strategy));
    }

    function queueWithdrawal(IStrategy strategy, uint256 sharesToWithdraw) public {
        // Initialize/populate strategy and share arrays
        IStrategy[] memory strategyArray = new IStrategy[](1);
        strategyArray[0] = strategy;
        uint256[] memory shareAmounts = new uint256[](1);
        shareAmounts[0] = sharesToWithdraw;

        // Create params array
        IDelegationManager.QueuedWithdrawalParams[] memory params = new IDelegationManager.QueuedWithdrawalParams[](1);
        params[0] = IDelegationManager.QueuedWithdrawalParams({
            strategies: strategyArray,
            shares: shareAmounts,
            withdrawer: address(this)
        });
        _queueWithdrawal(params);
    }

    function _queueWithdrawal(IDelegationManager.QueuedWithdrawalParams[] memory queueWithdrawalParams) internal {
        // Save pre-withdrawal state
        (uint256[] memory stakerSharesBefore, uint256[] memory operatorSharesBefore) = _setPreQueuedWithdrawalState(queueWithdrawalParams);

        // Queue withdrawal
        bytes32[] memory withdrawalRoots = globalRefs.delegationManager().queueWithdrawals(queueWithdrawalParams);

        // Checks
        assertEq(withdrawalRoots.length, queueWithdrawalParams.length, "withdrawal roots length mismatch");
        for (uint256 i = 0; i < withdrawalRoots.length; i++) {
            assertTrue(globalRefs.delegationManager().pendingWithdrawals(withdrawalRoots[i]), "withdrawal not pending");
        }
        assertEq(globalRefs.delegationManager().cumulativeWithdrawalsQueued(staker), queuedWithdrawals.length + queueWithdrawalParams.length, "cumulative withdrawals queued not updated correctly");
        uint256 stakerSharesAfter;
        uint256 operatorSharesAfter;
        for(uint256 i = 0; i < _tempStakerStrategies.length; i++) {
            stakerSharesAfter = _getStakerSharesForStrategy(_tempStakerStrategies[i]);
            operatorSharesAfter = _getOperatorSharesForStrategy(_tempStakerStrategies[i]);
            assertEq(stakerSharesAfter, stakerSharesBefore[i] - _tempStrategyShares[_tempStakerStrategies[i]], "staker shares not updated correctly");
            assertEq(operatorSharesAfter, operatorSharesBefore[i] - _tempStrategyShares[_tempStakerStrategies[i]], "operator shares not updated correctly");
        }

        // Update queuedWithdrawals array
        uint256 nonce;
        for(uint256 i = 0; i < queueWithdrawalParams.length; i++) {
            // Calculate nonce. Params.length is always <= cumulativeWithdrawalsQueued
            if (queueWithdrawalParams.length == globalRefs.delegationManager().cumulativeWithdrawalsQueued(staker)){
                nonce = i;
            } else {
                nonce = globalRefs.delegationManager().stakerNonce(staker) - queueWithdrawalParams.length - i;
            }

            IDelegationManager.Withdrawal memory withdrawal = IDelegationManager.Withdrawal({
                staker: staker,
                delegatedTo: globalRefs.delegationManager().delegatedTo(staker),
                withdrawer: queueWithdrawalParams[i].withdrawer,
                nonce: nonce,
                startBlock: uint32(block.number),
                strategies: queueWithdrawalParams[i].strategies,
                shares: queueWithdrawalParams[i].shares
            });
            queuedWithdrawals.push(withdrawal);
        }

        // Clear out temp state
        _clearTempState();
    }

    // Completes queued withdrawal as tokens at given index
    function completeQueuedWithdrawal(uint256 index) public {
        _completeQueuedWithdrawal(index);
    } 

    // Completes queued withdrawal as tokens
    function _completeQueuedWithdrawal(uint256 index) internal {
        // Setup withdrawal struct & store pre-withdrawal state
        IDelegationManager.Withdrawal memory withdrawal = queuedWithdrawals[index];
        IERC20[] memory tokens = new IERC20[](withdrawal.strategies.length);
        for (uint256 i = 0; i < withdrawal.strategies.length; i++) {
            tokens[i] = withdrawal.strategies[i].underlyingToken();
        }

        // Store pre withdrawal state for checks
        for(uint256 i = 0; i < withdrawal.strategies.length; i++) {
            if (!_tempStrategyRead[withdrawal.strategies[i]]) {
                _tempStrategyRead[withdrawal.strategies[i]]= true;
                _tempStakerStrategies.push(withdrawal.strategies[i]);
            }
            _tempStrategyShares[withdrawal.strategies[i]] += withdrawal.shares[i];
        }
        uint256[] memory balancesBefore = new uint256[](_tempStakerStrategies.length);
        uint256[] memory expectedTokensOut = new uint256[](_tempStakerStrategies.length);
        uint256[] memory totalSharesBefore = new uint256[](_tempStakerStrategies.length);
        for(uint256 i = 0; i < _tempStakerStrategies.length; i++) {
            balancesBefore[i] = _tempStakerStrategies[i].underlyingToken().balanceOf(staker);
            expectedTokensOut[i] = _tempStakerStrategies[i].sharesToUnderlying(_tempStrategyShares[_tempStakerStrategies[i]]);
            totalSharesBefore[i] = _tempStakerStrategies[i].totalShares();
        }

        // Complete queued withdrawal
        globalRefs.delegationManager().completeQueuedWithdrawal(withdrawal, tokens, 0, true);

        // Checks
        for(uint256 i = 0; i < _tempStakerStrategies.length; i++) {
            assertEq(_tempStakerStrategies[i].underlyingToken().balanceOf(staker), balancesBefore[i] + expectedTokensOut[i], "incorrect number of tokens received");
            assertEq(_tempStakerStrategies[i].totalShares(), totalSharesBefore[i] - _tempStrategyShares[_tempStakerStrategies[i]], "total shares not updated correctly");
        }

        // Clear temp state
        _clearTempState();

        // Delete ith element from queuedWithdrawals array
        queuedWithdrawals[index] = queuedWithdrawals[queuedWithdrawals.length - 1];
        queuedWithdrawals.pop();
    }

    // Helper Functions
    function _setPreQueuedWithdrawalState(IDelegationManager.QueuedWithdrawalParams[] memory queueWithdrawalParams) internal returns(uint256[] memory, uint256[] memory) {
        // There may be duplicate strategies to queue from multiple params, so we aggregate shares to queue for each strategy
        for (uint256 i = 0; i < queueWithdrawalParams.length; i++) {
            for (uint256 j = 0; j < queueWithdrawalParams[i].strategies.length; j++) {
                if (!_tempStrategyRead[queueWithdrawalParams[i].strategies[j]]) {
                    _tempStrategyRead[queueWithdrawalParams[i].strategies[j]] = true;
                    _tempStakerStrategies.push(queueWithdrawalParams[i].strategies[j]);
                }
                _tempStrategyShares[queueWithdrawalParams[i].strategies[j]] += queueWithdrawalParams[i].shares[j];
            }
        }
        uint256[] memory stakerSharesBefore = new uint256[](_tempStakerStrategies.length);
        uint256[] memory operatorSharesBefore = new uint256[](_tempStakerStrategies.length);
        for (uint256 i = 0; i < _tempStakerStrategies.length; i++) {
            stakerSharesBefore[i] = _getStakerSharesForStrategy(_tempStakerStrategies[i]);
            operatorSharesBefore[i] = _getOperatorSharesForStrategy(_tempStakerStrategies[i]);
        }
        return (stakerSharesBefore, operatorSharesBefore);
    }

    function _clearTempState() internal {
        for (uint256 i = 0; i < _tempStakerStrategies.length; i++) {
            delete _tempStrategyShares[_tempStakerStrategies[i]];
            delete _tempStrategyRead[_tempStakerStrategies[i]];
        }
        delete _tempStakerStrategies;
    }

    function _getOperatorSharesForStrategy(IStrategy strategy) internal view returns (uint256) {
        return globalRefs.delegationManager().operatorShares(globalRefs.delegationManager().delegatedTo(staker), strategy);
    }

    function _getStakerSharesForStrategy(IStrategy strategy) internal view returns (uint256) {
        return globalRefs.strategyManager().stakerStrategyShares(staker, strategy);
    }

    function _dealTokenToStaker(IERC20 token, uint256 amount) internal {
        StdCheats.deal(address(token), staker, amount);
    }
}