// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

interface IProtocolRegistryErrors {}

interface IProtocolRegistryTypes {
    /**
     * @notice Configuration for a protocol deployment.
     * @param pausable Whether this deployment can be paused.
     * @param deprecated Whether this deployment is deprecated.
     */
    struct DeploymentConfig {
        bool pausable;
        bool deprecated;
    }
}

interface IProtocolRegistryEvents is IProtocolRegistryTypes {
    /**
     * @notice Emitted when a deployment is shipped.
     * @param addr The address of the deployment.
     * @param config The configuration for the deployment.
     */
    event DeploymentShipped(address indexed addr, DeploymentConfig config);

    /**
     * @notice Emitted when a deployment is configured.
     * @param addr The address of the deployment.
     * @param config The configuration for the deployment.
     */
    event DeploymentConfigured(address indexed addr, DeploymentConfig config);

    /**
     * @notice Emitted when the semantic version is updated.
     * @param previousSemanticVersion The previous semantic version.
     * @param semanticVersion The new semantic version.
     */
    event SemanticVersionUpdated(string previousSemanticVersion, string semanticVersion);
}

interface IProtocolRegistry is IProtocolRegistryErrors, IProtocolRegistryEvents {
    /**
     * @notice Initializes the ProtocolRegistry with the initial admin.
     * @param initialAdmin The address to set as the initial admin.
     * @param pauserMultisig The address to set as the pauser multisig.
     */
    function initialize(address initialAdmin, address pauserMultisig) external;

    /**
     * @notice Ships a list of deployments.
     * @dev Only callable by the admin.
     * @param addresses The addresses of the deployments to ship.
     * @param configs The configurations of the deployments to ship.
     * @param contractNames The names of the contracts to ship.
     * @param semanticVersion The semantic version to ship.
     * @dev Contract names must be in ALL CAPS and contain NO SPACES (e.g., "ALLOCATIONMANAGER").
     */
    function ship(
        address[] calldata addresses,
        DeploymentConfig[] calldata configs,
        string[] calldata contractNames,
        string calldata semanticVersion
    ) external;

    /**
     * @notice Configures a deployment.
     * @dev Only callable by the admin.
     * @param addr The address of the deployment to configure.
     * @param config The configuration to set.
     */
    function configure(address addr, DeploymentConfig calldata config) external;

    /**
     * @notice Pauses all deployments that support pausing.
     * @dev Loops over all deployments and attempts to invoke `pauseAll()` on each contract that is marked as pausable.
     *      Silently ignores errors during calls for rapid pausing in emergencies. Pauser role only.
     */
    function pauseAll() external;

    /**
     * @notice Returns the full semantic version string of the protocol (e.g. "1.2.3").
     * @dev Follows Semantic Versioning 2.0.0 (see https://semver.org/).
     * @return The SemVer-formatted version string of the protocol.
     */
    function version() external view returns (string memory);

    /**
     * @notice Returns the major version component of the protocol's semantic version.
     * @dev Extracts and returns only the major version number as a string (e.g. "1" for version "1.2.3").
     * @return The major version number as a string.
     */
    function majorVersion() external view returns (string memory);

    /**
     * @notice Returns a deployment by name.
     * @param name The name of the deployment to get.
     * @return address The address of the deployment.
     */
    function getAddress(
        string calldata name
    ) external view returns (address);

    /**
     * @notice Returns a deployment by name.
     * @param name The name of the deployment to get.
     * @return addr The address.
     * @return config The configuration.
     */
    function getDeployment(
        string calldata name
    ) external view returns (address addr, DeploymentConfig memory config);

    /**
     * @notice Returns all deployments.
     * @return names The names of the deployments.
     * @return addresses The addresses.
     * @return configs The configurations.
     */
    function getAllDeployments()
        external
        view
        returns (string[] memory names, address[] memory addresses, DeploymentConfig[] memory configs);

    /**
     * @notice Returns the total number of deployments.
     * @return The total number of deployments.
     */
    function totalDeployments() external view returns (uint256);

    /**
     * @notice Returns the pauser role for the protocol.
     * @return The pauser role for the protocol.
     */
    function PAUSER_ROLE() external view returns (bytes32);
}
