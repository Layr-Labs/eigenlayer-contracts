// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../../contracts/permissions/Pausable.sol";
import "../../contracts/core/StrategyManagerStorage.sol";
import "../../contracts/interfaces/IServiceManager.sol";
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


    function recordOvercommittedBeaconChainETH(address overcommittedPodOwner, uint256 beaconChainETHStrategyIndex, uint256 amount)
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
     * @notice Get all details on the depositor's deposits and corresponding shares
     * @return (depositor's strategies, shares in these strategies)
     */
    function getDeposits(address depositor) external view returns (IStrategy[] memory, uint256[] memory) {}

    /// @notice Returns the array of strategies in which `staker` has nonzero shares
    function stakerStrats(address staker) external view returns (IStrategy[] memory) {}

    /// @notice Simple getter function that returns `stakerStrategyList[staker].length`.
    function stakerStrategyListLength(address staker) external view returns (uint256) {}


    function queueWithdrawal(
        uint256[] calldata strategyIndexes,
        IStrategy[] calldata strategies,
        uint256[] calldata shares,
        address withdrawer,
        bool undelegateIfPossible
    )
        external returns(bytes32) {}


    function completeQueuedWithdrawal(
        QueuedWithdrawal calldata queuedWithdrawal,
        IERC20[] calldata tokens,
        uint256 middlewareTimesIndex,
        bool receiveAsTokens
    )
        external{}

    function completeQueuedWithdrawals(
        QueuedWithdrawal[] calldata queuedWithdrawals,
        IERC20[][] calldata tokens,
        uint256[] calldata middlewareTimesIndexes,
        bool[] calldata receiveAsTokens
    )
        external{}


    function slashShares(
        address slashedAddress,
        address recipient,
        IStrategy[] calldata strategies,
        IERC20[] calldata tokens,
        uint256[] calldata strategyIndexes,
        uint256[] calldata shareAmounts
    )
        external{}

    /**
     * @notice Slashes an existing queued withdrawal that was created by a 'frozen' operator (or a staker delegated to one)
     * @param recipient The funds in the slashed withdrawal are withdrawn as tokens to this address.
     */
    function slashQueuedWithdrawal(
        address recipient,
        QueuedWithdrawal calldata queuedWithdrawal,
        IERC20[] calldata tokens,
        uint256[] calldata indicesToSkip
    )
        external{}

    /// @notice Returns the keccak256 hash of `queuedWithdrawal`.
    function calculateWithdrawalRoot(
        QueuedWithdrawal memory queuedWithdrawal
    )
        external
        pure
        returns (bytes32) {}

    /// @notice returns the enshrined beaconChainETH Strategy
    function beaconChainETHStrategy() external view returns (IStrategy) {}

    function withdrawalDelayBlocks() external view returns (uint256) {}

    function addStrategiesToDepositWhitelist(IStrategy[] calldata /*strategiesToWhitelist*/) external pure {}

    function removeStrategiesFromDepositWhitelist(IStrategy[] calldata /*strategiesToRemoveFromWhitelist*/) external pure {}   

    function undelegate() external pure {}

    function forceTotalWithdrawal(address /*staker*/) external pure returns (bytes32) {} 
}