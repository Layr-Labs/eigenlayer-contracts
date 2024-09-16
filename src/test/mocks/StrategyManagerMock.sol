// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "../../contracts/interfaces/IDelegationManager.sol";

contract StrategyManagerMock is Test {
    IDelegationManager public delegation;
    address public strategyWhitelister;

    mapping(address => IStrategy[]) public strategiesToReturn;
    mapping(address => uint256[]) public sharesToReturn;

    /// @notice Mapping staker => strategy => shares withdrawn after a withdrawal has been completed
    mapping(address => mapping(IStrategy => uint256)) public strategySharesWithdrawn;

    mapping(IStrategy => bool) public strategyIsWhitelistedForDeposit;

    /// @notice Mapping: staker => cumulative number of queued withdrawals they have ever initiated. only increments (doesn't decrement)
    mapping(address => uint256) public cumulativeWithdrawalsQueued;

    constructor(IDelegationManager _delegation) {
        delegation = _delegation;
    }

    /**
     * @notice mocks the return value of getDeposits
     * @param staker staker whose deposits are being mocked
     * @param _strategiesToReturn strategies to return in getDeposits
     * @param _sharesToReturn shares to return in getDeposits
     */
    function setDeposits(address staker, IStrategy[] calldata _strategiesToReturn, uint256[] calldata _sharesToReturn) external {
        require(_strategiesToReturn.length == _sharesToReturn.length, "StrategyManagerMock: length mismatch");
        strategiesToReturn[staker] = _strategiesToReturn;
        sharesToReturn[staker] = _sharesToReturn;
    }

    /**
     * @notice Adds deposit to the staker's deposits. Note that this function does not check if the staker
     * has already deposited for the strategy.
     */
    function addDeposit(address staker, IStrategy strategy, uint256 shares) external {
        strategiesToReturn[staker].push(strategy);
        sharesToReturn[staker].push(shares);
    }

    /**
     * @notice Get all details on the staker's deposits and corresponding shares
     * @return (staker's strategies, shares in these strategies)
     */
    function getDeposits(address staker) external view returns (IStrategy[] memory, uint256[] memory) {
        return (strategiesToReturn[staker], sharesToReturn[staker]);
    }

    function stakerDepositShares(address staker, IStrategy strategy) public view returns (uint256) {
        uint256 strategyIndex = _getStrategyIndex(staker, strategy);
        return sharesToReturn[staker][strategyIndex];
    }

    uint256 public stakerStrategyListLengthReturnValue;

    /// @notice Simple getter function that returns `stakerStrategyList[staker].length`.
    function stakerStrategyListLength(address /*staker*/) external view returns (uint256) {
        return stakerStrategyListLengthReturnValue;
    }

    function setStakerStrategyListLengthReturnValue(uint256 valueToSet) public {
        stakerStrategyListLengthReturnValue = valueToSet;
    }

    function setStrategyWhitelist(IStrategy strategy, bool value) external {
        strategyIsWhitelistedForDeposit[strategy] = value;
    }

    function addStrategiesToDepositWhitelist(
        IStrategy[] calldata strategiesToWhitelist
    ) external {
        for (uint256 i = 0; i < strategiesToWhitelist.length; ++i) {
            strategyIsWhitelistedForDeposit[strategiesToWhitelist[i]] = true;
        }
    }

    function removeDepositShares(
        address staker, IStrategy strategy, uint256 sharesToRemove
    ) external {
        uint256 strategyIndex = _getStrategyIndex(staker, strategy);
        sharesToReturn[staker][strategyIndex] -= sharesToRemove;
    }

    function removeStrategiesFromDepositWhitelist(IStrategy[] calldata /*strategiesToRemoveFromWhitelist*/) external pure {}


    function withdrawSharesAsTokens(
        address staker, 
        IStrategy strategy, 
        address, // token 
        uint256 shares
    ) external {
        strategySharesWithdrawn[staker][strategy] += shares;
    }

    function addShares(
        address staker, 
        IStrategy strategy, 
        IERC20, // token 
        uint256 addedShares
    ) external returns (uint, uint) {
        // Increase the staker's shares
        uint256 strategyIndex = _getStrategyIndex(staker, strategy);
        uint256 existingShares = sharesToReturn[staker][strategyIndex];
        sharesToReturn[staker][strategyIndex] += addedShares;

        return (existingShares, addedShares);
    }

    function burnShares(IStrategy strategy, uint256 sharesToBurn) external {}

    function _getStrategyIndex(address staker, IStrategy strategy) internal view returns (uint256) {
        IStrategy[] memory strategies = strategiesToReturn[staker];
        uint256 strategyIndex = type(uint256).max;
        for (uint256 i = 0; i < strategies.length; ++i) {
            if (strategies[i] == strategy) {
                strategyIndex = i;
                break;
            }
        }
        if (strategyIndex == type(uint256).max) {
            revert ("StrategyManagerMock: strategy not found");
        }

        return strategyIndex;
    }

    function setDelegationManager(IDelegationManager _delegation) external {
        delegation = _delegation;
    }

    fallback() external payable {}
    receive() external payable {}
}
