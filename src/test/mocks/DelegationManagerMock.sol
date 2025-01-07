// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IStrategyManager.sol";
import "src/contracts/libraries/SlashingLib.sol";

contract DelegationManagerMock is Test {
    receive() external payable {}
    fallback() external payable {}

    mapping(address => bool) public isOperator;
    mapping (address => address) public delegatedTo;
    mapping(address => mapping(IStrategy => uint256)) public operatorShares;

    function getDelegatableShares(address staker) external view returns (IStrategy[] memory, uint256[] memory) {}

    function setMinWithdrawalDelayBlocks(uint256 newMinWithdrawalDelayBlocks) external {}

    function setStrategyWithdrawalDelayBlocks(IStrategy[] calldata strategies, uint256[] calldata withdrawalDelayBlocks) external {}

    function setIsOperator(address operator, bool _isOperatorReturnValue) external {
        isOperator[operator] = _isOperatorReturnValue;
    }

    function slashOperatorShares(
        address operator,
        IStrategy strategy,
        uint64 prevMaxMagnitude,
        uint64 newMaxMagnitude
    ) external {
        uint256 amountSlashed = SlashingLib.calcSlashedAmount({
            operatorShares: operatorShares[operator][strategy],
            prevMaxMagnitude: prevMaxMagnitude,
            newMaxMagnitude: newMaxMagnitude
        });

        operatorShares[operator][strategy] -= amountSlashed;
    }

    /// @notice returns the total number of shares in `strategy` that are delegated to `operator`.
    function setOperatorShares(address operator, IStrategy strategy, uint256 shares) external {
        operatorShares[operator][strategy] = shares;
    }

    /// @notice returns the total number of shares in `strategy` that are delegated to `operator`.
    function setOperatorsShares(address operator, IStrategy[] memory strategies, uint256 shares) external {
        for (uint i = 0; i < strategies.length; i++) {
            operatorShares[operator][strategies[i]] = shares;
        }
    }

    function delegateTo(address operator, ISignatureUtils.SignatureWithExpiry memory /*approverSignatureAndExpiry*/, bytes32 /*approverSalt*/) external {
        delegatedTo[msg.sender] = operator;
    }

    function undelegate(address staker) external returns (bytes32[] memory withdrawalRoot) {
        delegatedTo[staker] = address(0);
        return withdrawalRoot;
    }

    function getOperatorsShares(address[] memory operators, IStrategy[] memory strategies) external view returns (uint256[][] memory) {
        uint256[][] memory operatorSharesArray = new uint256[][](operators.length);
        for (uint256 i = 0; i < operators.length; i++) {
            operatorSharesArray[i] = new uint256[](strategies.length);
            for (uint256 j = 0; j < strategies.length; j++) {
                operatorSharesArray[i][j] = operatorShares[operators[i]][strategies[j]];
            }
        }
        return operatorSharesArray;
    }

    function operatorDetails(address operator) external pure returns (IDelegationManagerTypes.OperatorDetails memory) {
        IDelegationManagerTypes.OperatorDetails memory returnValue = IDelegationManagerTypes.OperatorDetails({
            __deprecated_earningsReceiver: operator,
            delegationApprover: operator,
            __deprecated_stakerOptOutWindowBlocks: 0
        });
        return returnValue;
    }
    
    function isDelegated(address staker) external view returns (bool) {
        return (delegatedTo[staker] != address(0));
    }

    // onlyDelegationManager functions in StrategyManager
    function addShares(
        IStrategyManager strategyManager,
        address staker,
        IERC20 token,
        IStrategy strategy,
        uint256 shares
    ) external {
        strategyManager.addShares(staker, strategy, token, shares);
    }

    function removeDepositShares(
        IStrategyManager strategyManager,
        address staker,
        IStrategy strategy,
        uint256 shares
    ) external {
        strategyManager.removeDepositShares(staker, strategy, shares);
    }

    function withdrawSharesAsTokens(
        IStrategyManager strategyManager,
        address recipient,
        IStrategy strategy,
        uint256 shares,
        IERC20 token
    ) external {
        strategyManager.withdrawSharesAsTokens(recipient, strategy, token, shares);
    }
}
