// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "src/contracts/interfaces/IDelegationManager.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IDelegationFaucet {
    function mintDepositAndDelegate(
        address _operator,
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt,
        uint256 _depositAmount
    ) external;

    function getStaker(address operator) external returns (address);

    function depositIntoStrategy(
        address staker,
        IStrategy strategy,
        IERC20 token,
        uint256 amount
    ) external returns (bytes memory);

    function queueWithdrawal(
        address staker,
        IDelegationManager.QueuedWithdrawalParams[] calldata queuedWithdrawalParams
    ) external returns (bytes memory);

    function completeQueuedWithdrawal(
        address staker,
        IDelegationManager.Withdrawal calldata queuedWithdrawal,
        IERC20[] calldata tokens,
        uint256 middlewareTimesIndex,
        bool receiveAsTokens
    ) external returns (bytes memory);

    function transfer(address staker, address token, address to, uint256 amount) external returns (bytes memory);

    function callAddress(address to, bytes memory data) external payable returns (bytes memory);
}
