// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../../contracts/permissions/Pausable.sol";
import "../../contracts/core/StrategyManagerStorage.sol";
import "../../contracts/interfaces/IEigenPodManager.sol";
import "../../contracts/interfaces/IDelegationManager.sol";

// import "forge-std/Test.sol";

contract StrategyManagerMock is
    Initializable,
    IStrategyManager,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable,
    Pausable
    // ,Test
{


    IDelegationManager public delegation;
    IEigenPodManager public eigenPodManager;
    ISlasher public slasher;
    address public strategyWhitelister;

    mapping(address => IStrategy[]) public strategiesToReturn;
    mapping(address => uint256[]) public sharesToReturn;

    mapping(IStrategy => bool) public strategyIsWhitelistedForDeposit;

    /// @notice Mapping: staker => cumulative number of queued withdrawals they have ever initiated. only increments (doesn't decrement)
    mapping(address => uint256) public cumulativeWithdrawalsQueued;
    
    mapping(IStrategy => bool) public thirdPartyTransfersForbidden;

    function setAddresses(IDelegationManager _delegation, IEigenPodManager _eigenPodManager, ISlasher _slasher) external
    {
       delegation = _delegation;
       slasher = _slasher;
       eigenPodManager = _eigenPodManager;
    }

    function depositIntoStrategy(IStrategy strategy, IERC20 token, uint256 amount)
        external
        returns (uint256) {}


    function depositBeaconChainETH(address staker, uint256 amount) external{}


    function recordBeaconChainETHBalanceUpdate(address overcommittedPodOwner, uint256 beaconChainETHStrategyIndex, int256 sharesDelta)
        external{}

    function depositIntoStrategyWithSignature(
        IStrategy strategy,
        IERC20 token,
        uint256 amount,
        address staker,
        uint256 expiry,
        bytes memory signature
    )
        external
        returns (uint256 shares) {}

    /// @notice Returns the current shares of `user` in `strategy`
    function stakerStrategyShares(address user, IStrategy strategy) external view returns (uint256 shares) {}

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

    function setThirdPartyTransfersForbidden(IStrategy strategy, bool value) external {
        emit UpdatedThirdPartyTransfersForbidden(strategy, value);
        thirdPartyTransfersForbidden[strategy] = value;
    }

    /**
     * @notice Get all details on the staker's deposits and corresponding shares
     * @return (staker's strategies, shares in these strategies)
     */
    function getDeposits(address staker) external view returns (IStrategy[] memory, uint256[] memory) {
        return (strategiesToReturn[staker], sharesToReturn[staker]);
    }

    /// @notice Returns the array of strategies in which `staker` has nonzero shares
    function stakerStrats(address staker) external view returns (IStrategy[] memory) {}

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

    function removeShares(address staker, IStrategy strategy, uint256 shares) external {}

    function addShares(address staker, IERC20 token, IStrategy strategy, uint256 shares) external {}
    
    function withdrawSharesAsTokens(address recipient, IStrategy strategy, uint256 shares, IERC20 token) external {}

    /// @notice returns the enshrined beaconChainETH Strategy
    function beaconChainETHStrategy() external view returns (IStrategy) {}

    // function withdrawalDelayBlocks() external view returns (uint256) {}

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
