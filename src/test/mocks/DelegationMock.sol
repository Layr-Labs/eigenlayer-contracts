// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IDelegationManager.sol";


contract DelegationMock is IDelegationManager, Test {
    mapping(address => bool) public isOperator;
    mapping(address => mapping(IStrategy => uint256)) public operatorShares;

    function setIsOperator(address operator, bool _isOperatorReturnValue) external {
        isOperator[operator] = _isOperatorReturnValue;
    }

    /// @notice returns the total number of shares in `strategy` that are delegated to `operator`.
    function setOperatorShares(address operator, IStrategy strategy, uint256 shares) external {
        operatorShares[operator][strategy] = shares;
    }

    mapping (address => address) public delegatedTo;

    function registerAsOperator(IDelegationTerms /*dt*/) external {}

    function delegateTo(address operator) external {
        delegatedTo[msg.sender] = operator;
    }

    function delegateToBySignature(address /*staker*/, address /*operator*/, uint256 /*expiry*/, bytes memory /*signature*/) external {}

    function undelegate(address staker) external {
        delegatedTo[staker] = address(0);
    }

    /// @notice returns the DelegationTerms of the `operator`, which may mediate their interactions with stakers who delegate to them.
    function delegationTerms(address /*operator*/) external view returns (IDelegationTerms) {}

    function increaseDelegatedShares(address /*staker*/, IStrategy /*strategy*/, uint256 /*shares*/) external {}

    function decreaseDelegatedShares(
        address /*staker*/,
        IStrategy[] calldata /*strategies*/,
        uint256[] calldata /*shares*/
    ) external {}

    function isDelegated(address staker) external view returns (bool) {
        return (delegatedTo[staker] != address(0));
    }

    function isNotDelegated(address /*staker*/) external pure returns (bool) {}
}