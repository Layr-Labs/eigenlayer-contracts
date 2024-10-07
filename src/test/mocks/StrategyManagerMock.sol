// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "../../contracts/interfaces/IEigenPodManager.sol";
import "../../contracts/interfaces/IDelegationManager.sol";

contract StrategyManagerMock is Test {
    IDelegationManager public delegation;
    IEigenPodManager public eigenPodManager;
    address public strategyWhitelister;

    mapping(address => IStrategy[]) public strategiesToReturn;
    mapping(address => uint256[]) public sharesToReturn;

    mapping(IStrategy => bool) public strategyIsWhitelistedForDeposit;

    /// @notice Mapping: staker => cumulative number of queued withdrawals they have ever initiated. only increments (doesn't decrement)
    mapping(address => uint256) public cumulativeWithdrawalsQueued;
    mapping(IStrategy => bool) public thirdPartyTransfersForbidden;

    function setAddresses(IDelegationManager _delegation, IEigenPodManager _eigenPodManager) external
    {
       delegation = _delegation;
       eigenPodManager = _eigenPodManager;
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
     * @notice Get all details on the staker's deposits and corresponding shares
     * @return (staker's strategies, shares in these strategies)
     */
    function getDeposits(address staker) external view returns (IStrategy[] memory, uint256[] memory) {
        return (strategiesToReturn[staker], sharesToReturn[staker]);
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
        IStrategy[] calldata strategiesToWhitelist,
        bool[] calldata thirdPartyTransfersForbiddenValues
    ) external {
        for (uint256 i = 0; i < strategiesToWhitelist.length; ++i) {
            strategyIsWhitelistedForDeposit[strategiesToWhitelist[i]] = true;
            thirdPartyTransfersForbidden[strategiesToWhitelist[i]] = thirdPartyTransfersForbiddenValues[i];
        }
    }

    function removeStrategiesFromDepositWhitelist(IStrategy[] calldata /*strategiesToRemoveFromWhitelist*/) external pure {}
}
