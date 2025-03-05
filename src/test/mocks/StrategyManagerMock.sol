// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "../../contracts/interfaces/IDelegationManager.sol";

contract StrategyManagerMock is Test {
    IDelegationManager public delegation;
    address public strategyWhitelister;

    mapping(address => IStrategy[]) public strategiesToReturn;
    mapping(address => uint[]) public sharesToReturn;

    /// @notice Mapping staker => strategy => shares withdrawn after a withdrawal has been completed
    mapping(address => mapping(IStrategy => uint)) public strategySharesWithdrawn;

    mapping(IStrategy => bool) public strategyIsWhitelistedForDeposit;

    /// @notice Mapping: staker => cumulative number of queued withdrawals they have ever initiated. only increments (doesn't decrement)
    mapping(address => uint) public cumulativeWithdrawalsQueued;

    constructor(IDelegationManager _delegation) {
        delegation = _delegation;
    }

    /**
     * @notice mocks the return value of getDeposits
     * @param staker staker whose deposits are being mocked
     * @param _strategiesToReturn strategies to return in getDeposits
     * @param _sharesToReturn shares to return in getDeposits
     */
    function setDeposits(address staker, IStrategy[] calldata _strategiesToReturn, uint[] calldata _sharesToReturn) external {
        require(_strategiesToReturn.length == _sharesToReturn.length, "StrategyManagerMock: length mismatch");
        strategiesToReturn[staker] = _strategiesToReturn;
        sharesToReturn[staker] = _sharesToReturn;
    }

    /**
     * @notice Adds deposit to the staker's deposits. Note that this function does not check if the staker
     * has already deposited for the strategy.
     */
    function addDeposit(address staker, IStrategy strategy, uint shares) external {
        strategiesToReturn[staker].push(strategy);
        sharesToReturn[staker].push(shares);
    }

    /**
     * @notice Get all details on the staker's deposits and corresponding shares
     * @return (staker's strategies, shares in these strategies)
     */
    function getDeposits(address staker) external view returns (IStrategy[] memory, uint[] memory) {
        return (strategiesToReturn[staker], sharesToReturn[staker]);
    }

    function stakerDepositShares(address staker, IStrategy strategy) public view returns (uint) {
        uint strategyIndex = _getStrategyIndex(staker, strategy);
        return sharesToReturn[staker][strategyIndex];
    }

    uint public stakerStrategyListLengthReturnValue;

    /// @notice Simple getter function that returns `stakerStrategyList[staker].length`.
    function stakerStrategyListLength(address /*staker*/ ) external view returns (uint) {
        return stakerStrategyListLengthReturnValue;
    }

    function setStakerStrategyListLengthReturnValue(uint valueToSet) public {
        stakerStrategyListLengthReturnValue = valueToSet;
    }

    function setStrategyWhitelist(IStrategy strategy, bool value) external {
        strategyIsWhitelistedForDeposit[strategy] = value;
    }

    function addStrategiesToDepositWhitelist(IStrategy[] calldata strategiesToWhitelist) external {
        for (uint i = 0; i < strategiesToWhitelist.length; ++i) {
            strategyIsWhitelistedForDeposit[strategiesToWhitelist[i]] = true;
        }
    }

    function removeDepositShares(address staker, IStrategy strategy, uint sharesToRemove) external returns (uint) {
        uint strategyIndex = _getStrategyIndex(staker, strategy);
        uint updatedShares = sharesToReturn[staker][strategyIndex] - sharesToRemove;
        sharesToReturn[staker][strategyIndex] = updatedShares;
        return updatedShares;
    }

    function removeStrategiesFromDepositWhitelist(IStrategy[] calldata /*strategiesToRemoveFromWhitelist*/ ) external pure {}

    function withdrawSharesAsTokens(
        address staker,
        IStrategy strategy,
        address, // token
        uint shares
    ) external {
        strategySharesWithdrawn[staker][strategy] += shares;
    }

    function addShares(address staker, IStrategy strategy, uint addedShares) external returns (uint, uint) {
        // Increase the staker's shares
        uint strategyIndex = _getStrategyIndex(staker, strategy);
        uint existingShares = sharesToReturn[staker][strategyIndex];
        sharesToReturn[staker][strategyIndex] += addedShares;

        return (existingShares, addedShares);
    }

    function burnShares(IStrategy strategy, uint sharesToBurn) external {}

    function _getStrategyIndex(address staker, IStrategy strategy) internal view returns (uint) {
        IStrategy[] memory strategies = strategiesToReturn[staker];
        uint strategyIndex = type(uint).max;
        for (uint i = 0; i < strategies.length; ++i) {
            if (strategies[i] == strategy) {
                strategyIndex = i;
                break;
            }
        }
        if (strategyIndex == type(uint).max) revert("StrategyManagerMock: strategy not found");

        return strategyIndex;
    }

    function setDelegationManager(IDelegationManager _delegation) external {
        delegation = _delegation;
    }

    fallback() external payable {}
    receive() external payable {}
}
