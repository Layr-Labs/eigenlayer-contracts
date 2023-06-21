// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title Interface for a `PaymentManager` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
interface IPaymentManager {
    /**
     * @notice deposit one-time fees by the `msg.sender` with this contract to pay for future tasks of this middleware
     * @param depositFor could be the `msg.sender` themselves, or a different address for whom `msg.sender` is depositing these future fees
     * @param amount is amount of futures fees being deposited
     */
    function depositFutureFees(address depositFor, uint256 amount) external;

    /// @notice Allows the `allowed` address to spend up to `amount` of the `msg.sender`'s funds that have been deposited in this contract
    function setAllowance(address allowed, uint256 amount) external;

    /// @notice Used for deducting the fees from the payer to the middleware
    function takeFee(address initiator, address payer, uint256 feeAmount) external;

    /// @notice the ERC20 token that will be used by the disperser to pay the service fees to middleware nodes.
    function paymentToken() external view returns (IERC20);
}
