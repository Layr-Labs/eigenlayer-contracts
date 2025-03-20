// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

/// @notice DEPRECATED INTERFACE at commit hash https://github.com/Layr-Labs/eigenlayer-contracts/tree/0139d6213927c0a7812578899ddd3dda58051928
interface IDelayedWithdrawalRouter_DeprecatedM1 {
    // struct used to pack data into a single storage slot
    struct DelayedWithdrawal {
        uint224 amount;
        uint32 blockCreated;
    }

    // struct used to store a single users delayedWithdrawal data
    struct UserDelayedWithdrawals {
        uint delayedWithdrawalsCompleted;
        DelayedWithdrawal[] delayedWithdrawals;
    }

    /**
     * @notice Creates an delayed withdrawal for `msg.value` to the `recipient`.
     * @dev Only callable by the `podOwner`'s EigenPod contract.
     */
    function createDelayedWithdrawal(address podOwner, address recipient) external payable;

    /**
     * @notice Called in order to withdraw delayed withdrawals made to the `recipient` that have passed the `withdrawalDelayBlocks` period.
     * @param recipient The address to claim delayedWithdrawals for.
     * @param maxNumberOfWithdrawalsToClaim Used to limit the maximum number of withdrawals to loop through claiming.
     */
    function claimDelayedWithdrawals(address recipient, uint maxNumberOfWithdrawalsToClaim) external;

    /**
     * @notice Called in order to withdraw delayed withdrawals made to the caller that have passed the `withdrawalDelayBlocks` period.
     * @param maxNumberOfWithdrawalsToClaim Used to limit the maximum number of withdrawals to loop through claiming.
     */
    function claimDelayedWithdrawals(uint maxNumberOfWithdrawalsToClaim) external;

    /// @notice Owner-only function for modifying the value of the `withdrawalDelayBlocks` variable.
    function setWithdrawalDelayBlocks(uint newValue) external;

    /// @notice Getter function for the mapping `_userWithdrawals`
    function userWithdrawals(address user) external view returns (UserDelayedWithdrawals memory);

    /// @notice Getter function to get all delayedWithdrawals of the `user`
    function getUserDelayedWithdrawals(address user) external view returns (DelayedWithdrawal[] memory);

    /// @notice Getter function to get all delayedWithdrawals that are currently claimable by the `user`
    function getClaimableUserDelayedWithdrawals(address user) external view returns (DelayedWithdrawal[] memory);

    /// @notice Getter function for fetching the delayedWithdrawal at the `index`th entry from the `_userWithdrawals[user].delayedWithdrawals` array
    function userDelayedWithdrawalByIndex(address user, uint index) external view returns (DelayedWithdrawal memory);

    /// @notice Getter function for fetching the length of the delayedWithdrawals array of a specific user
    function userWithdrawalsLength(address user) external view returns (uint);

    /// @notice Convenience function for checking whether or not the delayedWithdrawal at the `index`th entry from the `_userWithdrawals[user].delayedWithdrawals` array is currently claimable
    function canClaimDelayedWithdrawal(address user, uint index) external view returns (bool);

    /**
     * @notice Delay enforced by this contract for completing any delayedWithdrawal. Measured in blocks, and adjustable by this contract's owner,
     * up to a maximum of `MAX_WITHDRAWAL_DELAY_BLOCKS`. Minimum value is 0 (i.e. no delay enforced).
     */
    function withdrawalDelayBlocks() external view returns (uint);
}
