// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./ILightClientUpdater.sol";

/**
 * @title IBeaconChainOracle interface for reading beacon state roots from the Telepathy protocol.
 * @author Layr Labs Inc.
 */
interface IBeaconChainOracle {
    /**
     * @notice Getter for the lightClientUpdater contract.
     * @return The address of the contract to retrieve beacon state roots.
     */
    function lightClientUpdater() external view returns (ILightClientUpdater);

    /**
     * @notice Check whether the prover address is in the set of approved provers.
     * @param prover The address to check.
     * @return True if the prover is approved, false otherwise.
     */
    function isApprovedProver(address prover) external view returns (bool);

    /**
     * @notice Function used to get the beacon state root at the specified blockNumber and make sure that it was proven by an approved prover.
     * @param blockNumber The blockNumber of the beacon state root to be retrieved.
     * @return The beacon state root at the specified blockNumber.
     * @dev Function will revert if the beacon state root at the specified blockNumber was not proven by an approved prover.
     */
    function beaconStateRootAtBlockNumber(uint256 blockNumber) external view returns (bytes32);

    /**
     * @notice Owner-only function used to add provers to the set of approved provers.
     * @param provers Array of addresses to be added to the set.
     * @dev Function will have no effect on the i-th input address if provers[i] is already in the set of approved provers.
     */
    function addApprovedProvers(address[] memory provers) external;

    /**
     * @notice Owner-only function used to remove provers from the set of approved provers.
     * @param provers Array of addresses to be removed from the set.
     * @dev Function will have no effect on the i-th input address if provers[i] is already not in the set of approved provers.
     */
    function removeApprovedProvers(address[] memory provers) external;
}
