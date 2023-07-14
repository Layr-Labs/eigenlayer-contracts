// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IDelegationManager.sol";


contract DelegationMock is IDelegationManager, Test {
    mapping(address => bool) public isOperator;

    function setIsOperator(address operator, bool _isOperatorReturnValue) external {
        isOperator[operator] = _isOperatorReturnValue;
    }

    mapping(address => address) public delegatedTo;

    function registerAsOperator(OperatorDetails calldata /*registeringOperatorDetails*/) external pure {}

    function delegateTo(address operator, SignatureWithExpiry memory /*approverSignatureAndExpiry*/) external {
        delegatedTo[msg.sender] = operator;
    }

    function modifyOperatorDetails(OperatorDetails calldata /*newOperatorDetails*/) external pure {}

    function delegateToBySignature(
        address /*staker*/,
        address /*operator*/,
        SignatureWithExpiry memory /*stakerSignatureAndExpiry*/,
        SignatureWithExpiry memory /*approverSignatureAndExpiry*/
    ) external pure {}

    function undelegate(address staker) external {
        delegatedTo[staker] == address(0);
    }

    function forceUndelegation(address /*staker*/, address /*operator*/) external pure {}

    function increaseDelegatedShares(address /*staker*/, IStrategy /*strategy*/, uint256 /*shares*/) external pure {}

    function decreaseDelegatedShares(address /*staker*/, IStrategy[] calldata /*strategies*/, uint256[] calldata /*shares*/) external pure {}

    function operatorDetails(address operator) external pure returns (OperatorDetails memory) {
        return (new OperatorDetails{
            earningsReceiver: operator,
            delegationApprover: operator,
            stakerOptOutWindowBlocks: 0
        });
    }

    function earningsReceiver(address operator) external pure returns (address) {
        return operator;
    }

    function delegationApprover(address operator) external pure returns (address) {
        return operator;
    }

    function stakerOptOutWindowBlocks(address operator) external pure returns (uint256) {
        return 0;
    }

    function operatorShares(address /*operator*/, IStrategy /*strategy*/) external pure returns (uint256) {}

    function isDelegated(address staker) external view returns (bool) {
        return (delegatedTo[staker] != address(0));
    }

    function isNotDelegated(address /*staker*/) external pure returns (bool) {}

    // function isOperator(address /*operator*/) external pure returns (bool) {}

    function stakerNonce(address /*staker*/) external pure returns (uint256) {}

    function delegationApproverNonce(address /*operator*/) external pure returns (uint256) {}
}