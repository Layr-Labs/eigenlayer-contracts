// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./IVoteWeigher.sol";
import "./IPaymentManager.sol";

/**
 * @title Interface for a `ServiceManager`-type contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
interface IServiceManager {
    // @notice Event that must be emitted when the service's VoteWeigher contract changes
    event VoteWeigherChanged(IVoteWeigher previousVoteWeigher, IVoteWeigher newVoteWeigher);

    // @notice Event that must be emitted when the service's PaymentManager contract changes
    event PaymentManagerChanged(IPaymentManager previousPaymentManager, IPaymentManager newPaymentManager);

    /// @notice Returns the current 'taskNumber' for the middleware
    function taskNumber() external view returns (uint32);

    /// @notice Permissioned function that causes the ServiceManager to freeze the operator on EigenLayer, through a call to the Slasher contract
    function freezeOperator(address operator) external;

    /// @notice Permissioned function to have the ServiceManager forward a call to the slasher, recording an initial stake update (on operator registration)
    function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) external;

    /// @notice Permissioned function to have the ServiceManager forward a call to the slasher, recording a stake update
    function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 prevElement) external;

    /// @notice Permissioned function to have the ServiceManager forward a call to the slasher, recording a final stake update (on operator deregistration)
    function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) external;

    /// @notice Returns the latest block until which operators must serve.
    function latestServeUntilBlock() external view returns (uint32);

    function owner() external view returns (address);

    // @notice The service's VoteWeigher contract, which could be this contract itself
    function voteWeigher() external view returns (IVoteWeigher);

    // @notice The service's PaymentManager contract, which could be this contract itself
    function paymentManager() external view returns (IPaymentManager);
}