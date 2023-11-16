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

    // Staker Address
    address staker = address(this);

    constructor(GlobalRefs _globalRefs) {
        globalRefs = _globalRefs;
    }

    function delegate(Operator operator) public {
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        bytes32 emptySalt;
        _delegate(operator, approverSignatureAndExpiry, emptySalt);
    }

    function _delegate(Operator operator, ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry, bytes32 approverSalt) internal {
        //TODO: get operator shares before and operator shares after and compare
        globalRefs.delegationManager().delegateTo(address(operator), approverSignatureAndExpiry, approverSalt);

        // Checks
        assertEq(globalRefs.delegationManager().delegatedTo(staker), address(operator), "staker delegated to the wrong address");
        
        // Update operator's delegatedStaker array
        operator.addDelegatedStaker(address(this));
    }

    function delegateWithSig() public {}

    function depositIntoStrategy(IStrategy strategy, IERC20 token, uint256 amount) public {
        _dealTokenToStaker(token, amount);
        token.approve(address(globalRefs.strategyManager()), amount);

        uint256 shares = globalRefs.strategyManager().depositIntoStrategy(strategy, token, amount);

        // TODO: checks
    }

    function queueFullWithdrawal(IStrategy strategy) public {
        queueWithdrawal(strategy, globalRefs.strategyManager().stakerStrategyShares(staker, strategy));
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

    function queueWithdrawal(IDelegationManager.QueuedWithdrawalParams[] memory queueWithdrawalParams) public {}

    function _queueWithdrawal(IDelegationManager.QueuedWithdrawalParams[] memory queueWithdrawalParams) public {
        globalRefs.delegationManager().queueWithdrawals(queueWithdrawalParams);

        //TODO: checks

        // Update queuedWithdrawals array
        for(uint256 i = 0; i < queueWithdrawalParams.length; i++) {
            IDelegationManager.Withdrawal memory withdrawal = IDelegationManager.Withdrawal({
                staker: staker,
                delegatedTo: globalRefs.delegationManager().delegatedTo(staker),
                withdrawer: queueWithdrawalParams[i].withdrawer,
                nonce: globalRefs.delegationManager().stakerNonce(staker) - queueWithdrawalParams.length - i,
                startBlock: uint32(block.number),
                strategies: queueWithdrawalParams[i].strategies,
                shares: queueWithdrawalParams[i].shares
            });
            queuedWithdrawals.push(withdrawal);
        }
    }

    function completeQueuedWithdrawal() public {
        
    } 

    function _dealTokenToStaker(IERC20 token, uint256 amount) internal {
        StdCheats.deal(address(token), staker, amount);
    }
}