// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

interface IProtocolRegistryErrors {
    /// @notice Thrown when two array parameters have mismatching lengths.
    error InputArrayLengthMismatch();
}

interface IProtocolRegistryTypes {}

interface IProtocolRegistryEvents is IProtocolRegistryTypes {
    /// @notice Emitted when the version is set for a given address.
    /// @param addr The address for which the version is set.
    /// @param semver The semantic version string set for the address.
    event VersionSet(address indexed addr, string semver);
}

interface IProtocolRegistry is IProtocolRegistryErrors, IProtocolRegistryEvents {
    /**
     * @notice Initializes the ProtocolRegistry with the initial owner.
     * @param initialOwner The address to set as the initial owner.
     */
    function initialize(
        address initialOwner
    ) external;

    /**
     * @notice Sets the semantic version for a single address.
     * @dev Only callable by the contract owner.
     * @param addr The address for which to set the version.
     * @param semver The semantic version string to set for the address.
     */
    function setVersion(address addr, string calldata semver) external;

    /**
     * @notice Sets the same semantic version for each address in the provided array.
     * @dev Only callable by the contract owner.
     * @param addresses The addresses for which to set the version.
     * @param semver The semantic version string to set for all addresses.
     */
    function setVersions(address[] calldata addresses, string calldata semver) external;

    /**
     * @notice Sets a distinct semantic version for each address in the provided array.
     * @dev Only callable by the contract owner.
     * @param addresses The addresses for which to set the version.
     * @param semvers The semantic version strings to set, one for each address.
     */
    function setVersions(address[] calldata addresses, string[] calldata semvers) external;

    /**
     * @notice Returns the semantic version string for a given address.
     * @param addr The address to query.
     * @return The semantic version string associated with the address.
     */
    function version(
        address addr
    ) external view returns (string memory);

    /**
     * @notice Returns the major version string for a given address.
     * @param addr The address to query.
     * @return The major version string associated with the address.
     */
    function majorVersion(
        address addr
    ) external view returns (string memory);
}
