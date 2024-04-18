// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

interface IEigenPodManagerEvents {
    /// @notice Emitted to notify the update of the beaconChainOracle address
    event BeaconOracleUpdated(address indexed newOracleAddress);

    /// @notice Emitted to notify that the denebForkTimestamp has been set
    event DenebForkTimestampUpdated(uint64 denebForkTimestamp);

    /// @notice Emitted to notify the deployment of an EigenPod
    event PodDeployed(address indexed eigenPod, address indexed podOwner);

    /// @notice Emitted when the balance of an EigenPod is updated
    event PodSharesUpdated(address indexed podOwner, int256 sharesDelta);
}
