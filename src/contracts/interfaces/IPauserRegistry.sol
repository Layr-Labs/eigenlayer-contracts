// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

/**
 * @title Interface for the `PauserRegistry` contract.
 * @author Layr Labs, Inc.
 */
interface IPauserRegistry {
    /// @notice Unique address that holds the pauser role.
    function pauser() external view returns (address);

    /// @notice Unique address that holds the unpauser role. Capable of changing *both* the pauser and unpauser addresses.
    function unpauser() external view returns (address);
}
