// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

/**
 * @title Interface for a contract that updates the telepathy light client.
 */
interface ILightClientUpdater {
    /**
     * @notice Fetches the beacon state root and prover address for the given block number from the Telepathy Light Client contract.
     * @param blockNumber The block number to update the light client with.
     * @return The address of the prover for the state root.
     * @return The state root.
     */
    function getBeaconStateRootAndProver(uint256 blockNumber) external view returns(address, bytes32);
}