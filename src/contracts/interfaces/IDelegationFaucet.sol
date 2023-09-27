// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "src/contracts/interfaces/IStrategyManager.sol";
import "src/contracts/interfaces/IStrategy.sol";
import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IBLSRegistry.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Create2.sol";

interface IDelegationFaucet {
    function mintDepositAndDelegate(address operator, uint256 depositAmount) external;

    function getStaker(address operator) external returns (address);

    function depositIntoStrategy(
        address staker,
        IStrategy strategy,
        IERC20 token,
        uint256 amount
    ) external returns (bytes memory);

    function queueWithdrawal(
        address staker,
        uint256[] calldata strategyIndexes,
        IStrategy[] calldata strategies,
        uint256[] calldata shares,
        address withdrawer
    ) external returns (bytes memory);

    function completeQueuedWithdrawal(
        address staker,
        IStrategyManager.QueuedWithdrawal calldata queuedWithdrawal,
        IERC20[] calldata tokens,
        uint256 middlewareTimesIndex,
        bool receiveAsTokens
    ) external returns (bytes memory);

    function transfer(address staker, address token, address to, uint256 amount) external returns (bytes memory);

    function callAddress(address to, bytes memory data) external payable returns (bytes memory);
}
