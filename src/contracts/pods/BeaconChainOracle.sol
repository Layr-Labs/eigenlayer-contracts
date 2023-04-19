// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../interfaces/IBeaconChainOracle.sol";
import "../interfaces/ILightClientUpdater.sol";

/**
 * @title Oracle contract used for reading beacon state roots from the Telepathy protocol.
 * @author Layr Labs Inc.
 */
contract BeaconChainOracle is IBeaconChainOracle, Ownable {
    /// @notice The address of the contract to retrieve beacon state root.
    ILightClientUpdater public lightClientUpdater;

    /// @notice Mapping: prover address => whether or not the address is in the set of approved provers.
    mapping(address => bool) public isApprovedProver;

    /// @notice Emitted when addedProver is added to the set of approved provers.
    event ProverAdded(address addedProver);

    /// @notice Emitted when removedProver is removed from the set of approved provers.
    event ProverRemoved(address removedProver);

    constructor(address initialOwner, ILightClientUpdater _lightClientUpdater, address[] memory initialProvers) Ownable() {
        lightClientUpdater = _lightClientUpdater;
        for (uint256 i = 0; i < initialProvers.length; ) {
            _addApprovedProver(initialProvers[i]);
            unchecked {
                ++i;
            }
        }
        _transferOwnership(initialOwner);
    }

    /** 
     * @notice Function used to get the beacon state root at the specified blockNumber and make sure that it was proven by an approved prover.
     * @param blockNumber The blockNumber of the beacon state root to be retrieved.
     * @return The beacon state root at the specified blockNumber.
     * @dev Function will revert if the beacon state root at the specified blockNumber was not proven by an approved prover.
     */
    function beaconStateRootAtBlockNumber(uint256 blockNumber) external view returns (bytes32) {
        (address prover, bytes32 stateRoot) = lightClientUpdater.getBeaconStateRootAndProver(blockNumber);
        require(isApprovedProver[prover], "Prover is not approved");
        return stateRoot;
    }


    /**
     * @notice Owner-only function used to add provers to the set of approved provers.
     * @param provers Array of addresses to be added to the set.
     * @dev Function will have no effect on the i-th input address if provers[i] is already in the set of approved provers.
     */
    function addApprovedProvers(address[] memory provers) external onlyOwner {
        for (uint256 i = 0; i < provers.length; ) {
            _addApprovedProver(provers[i]);
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice Owner-only function used to remove provers from the set of approved provers.
     * @param provers Array of addresses to be removed from the set.
     * @dev Function will have no effect on the i-th input address if provers[i] is already not in the set of approved provers.
    */
    function removeApprovedProvers(address[] memory provers) external onlyOwner {
        for (uint256 i = 0; i < provers.length; ) {
            _removeApprovedProver(provers[i]);
            unchecked {
                ++i;
            }
        }
    }

    function _addApprovedProver(address prover) internal {
        if (!isApprovedProver[prover]) {
            isApprovedProver[prover] = true;
            emit ProverAdded(prover);
        }
    }

    function _removeApprovedProver(address prover) internal {
        if (isApprovedProver[prover]) {
            isApprovedProver[prover] = false;
            emit ProverRemoved(prover);
        }
    }
}