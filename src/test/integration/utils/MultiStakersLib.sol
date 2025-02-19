// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/users/User.t.sol";


/// @notice Contract that provides utility functions to reuse staker actions in User.t.sol
// for multiple stakers at once. 
library MultiStakersLib {

    function depositIntoEigenlayer(
        User[] memory stakers,
        IStrategy[][] memory strategies,
        uint[][] memory tokenBalances
    ) internal {
        for (uint i = 0; i < stakers.length; i++) {
            stakers[i].depositIntoEigenlayer(strategies[i], tokenBalances[i]);
        }
    }

    function delegateTo(
        User[] memory stakers,
        User operator
    ) internal {
        for (uint i = 0; i < stakers.length; i++) {
            stakers[i].delegateTo(operator);
        }
    }

    function undelegate(
        User[] memory stakers
    ) internal {
        for (uint i = 0; i < stakers.length; i++) {
            stakers[i].undelegate();
        }
    }

    function redelegate(
        User[] memory stakers,
        User newOperator
    ) internal {
        for (uint i = 0; i < stakers.length; i++) {
            stakers[i].redelegate(newOperator);
        }
    }

    function queueWithdrawals(
        User[] memory stakers,
        IStrategy[][] memory strategies,
        uint[][] memory shares
    ) internal returns (IDelegationManagerTypes.Withdrawal[][] memory withdrawals) {
        withdrawals = new IDelegationManagerTypes.Withdrawal[][](stakers.length);
        for (uint i = 0; i < stakers.length; i++) {
            withdrawals[i] = stakers[i].queueWithdrawals(strategies[i], shares[i]);
        }
    }

    function completeWithdrawalsAsTokens(
        User[] memory stakers,
        IDelegationManagerTypes.Withdrawal[][] memory withdrawals
    ) internal {
        for (uint i = 0; i < stakers.length; i++) {
            stakers[i].completeWithdrawalsAsTokens(withdrawals[i]);
        }
    }

    function completeWithdrawalsAsShares(
        User[] memory stakers,
        IDelegationManagerTypes.Withdrawal[][] memory withdrawals
    ) internal {
        for (uint i = 0; i < stakers.length; i++) {
            stakers[i].completeWithdrawalsAsShares(withdrawals[i]);
        }
    }
}
