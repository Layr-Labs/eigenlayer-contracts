// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "../../contracts/interfaces/IDelayedWithdrawalRouter.sol";

contract DelayedWithdrawalRouterMock is IDelayedWithdrawalRouter {
    /**
     * @notice Creates an delayed withdrawal for `msg.value` to the `recipient`.
     * @dev Only callable by the `podOwner`'s EigenPod contract.
     */
    function createDelayedWithdrawal(address podOwner, address recipient) external payable{}

    /**
     * @notice Called in order to withdraw delayed withdrawals made to the `recipient` that have passed the `withdrawalDelayBlocks` period.
     * @param recipient The address to claim delayedWithdrawals for.
     * @param maxNumberOfWithdrawalsToClaim Used to limit the maximum number of withdrawals to loop through claiming.
     */
    function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfWithdrawalsToClaim) external{}

    /**
     * @notice Called in order to withdraw delayed withdrawals made to the caller that have passed the `withdrawalDelayBlocks` period.
     * @param maxNumberOfWithdrawalsToClaim Used to limit the maximum number of withdrawals to loop through claiming.
     */
    function claimDelayedWithdrawals(uint256 maxNumberOfWithdrawalsToClaim) external{}

    /// @notice Owner-only function for modifying the value of the `withdrawalDelayBlocks` variable.
    function setWithdrawalDelayBlocks(uint256 newValue) external{}

    /// @notice Getter function for the mapping `_userWithdrawals`
    function userWithdrawals(address user) external view returns (UserDelayedWithdrawals memory){}

    /// @notice Getter function to get all delayedWithdrawals of the `user`
    function getUserDelayedWithdrawals(address user) external view returns (DelayedWithdrawal[] memory){}

    /// @notice Getter function to get all delayedWithdrawals that are currently claimable by the `user`
    function getClaimableUserDelayedWithdrawals(address user) external view returns (DelayedWithdrawal[] memory){}

    /// @notice Getter function for fetching the delayedWithdrawal at the `index`th entry from the `_userWithdrawals[user].delayedWithdrawals` array
    function userDelayedWithdrawalByIndex(address user, uint256 index) external view returns (DelayedWithdrawal memory){}

    /// @notice Getter function for fetching the length of the delayedWithdrawals array of a specific user
    function userWithdrawalsLength(address user) external view returns (uint256){}

    /// @notice Convenience function for checking whether or not the delayedWithdrawal at the `index`th entry from the `_userWithdrawals[user].delayedWithdrawals` array is currently claimable
    function canClaimDelayedWithdrawal(address user, uint256 index) external view returns (bool){}

    /**
     * @notice Delay enforced by this contract for completing any delayedWithdrawal. Measured in blocks, and adjustable by this contract's owner,
     * up to a maximum of `MAX_WITHDRAWAL_DELAY_BLOCKS`. Minimum value is 0 (i.e. no delay enforced).
     */
    function withdrawalDelayBlocks() external view returns (uint256){}
}
