// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

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
    address public strategyWhitelister;

    mapping(address => IStrategy[]) public strategiesToReturn;
    mapping(address => Shares[]) public sharesToReturn;

    mapping(IStrategy => bool) public strategyIsWhitelistedForDeposit;

    /// @notice Mapping: staker => cumulative number of queued withdrawals they have ever initiated. only increments (doesn't decrement)
    mapping(address => uint256) public cumulativeWithdrawalsQueued;

    function setAddresses(IDelegationManager _delegation, IEigenPodManager _eigenPodManager) external
    {
       delegation = _delegation;
       eigenPodManager = _eigenPodManager;
    }

    function depositIntoStrategy(IStrategy strategy, IERC20 token, uint256 amount)
        external
        returns (OwnedShares) {}


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
        returns (OwnedShares shares) {}

    function getStakerStrategyList(
        address staker
    ) external view returns (IStrategy[] memory) {}

    /**
     * @notice mocks the return value of getDeposits
     * @param staker staker whose deposits are being mocked
     * @param _strategiesToReturn strategies to return in getDeposits
     * @param _sharesToReturn shares to return in getDeposits
     */
    function setDeposits(address staker, IStrategy[] calldata _strategiesToReturn, Shares[] calldata _sharesToReturn) external {
        require(_strategiesToReturn.length == _sharesToReturn.length, "StrategyManagerMock: length mismatch");
        strategiesToReturn[staker] = _strategiesToReturn;
        sharesToReturn[staker] = _sharesToReturn;
    }

    /**
     * @notice Get all details on the staker's deposits and corresponding shares
     * @return (staker's strategies, shares in these strategies)
     */
    function getDeposits(address staker) external view returns (IStrategy[] memory, Shares[] memory) {
        return (strategiesToReturn[staker], sharesToReturn[staker]);
    }

    /// @notice Returns the array of strategies in which `staker` has nonzero shares
    function stakerStrats(address staker) external view returns (IStrategy[] memory) {}

    uint256 public stakerStrategyListLengthReturnValue;

    /// @notice Returns the current shares of `user` in `strategy`
    function stakerStrategyShares(
        address user,
        IStrategy strategy
    ) external view returns (Shares shares) {}

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

    function removeShares(address staker, IStrategy strategy, Shares shares) external {}

    function addOwnedShares(address staker, IStrategy strategy, IERC20 token, OwnedShares shares) external {}
    
    function withdrawSharesAsTokens(address recipient, IStrategy strategy, IERC20 token, OwnedShares shares) external {}

    /// @notice returns the enshrined beaconChainETH Strategy
    function beaconChainETHStrategy() external view returns (IStrategy) {}

    // function withdrawalDelayBlocks() external view returns (uint256) {}

    function addStrategiesToDepositWhitelist(
        IStrategy[] calldata strategiesToWhitelist
    ) external {
        for (uint256 i = 0; i < strategiesToWhitelist.length; ++i) {
            strategyIsWhitelistedForDeposit[strategiesToWhitelist[i]] = true;
        }
    }

    function removeStrategiesFromDepositWhitelist(IStrategy[] calldata /*strategiesToRemoveFromWhitelist*/) external pure {}
}
