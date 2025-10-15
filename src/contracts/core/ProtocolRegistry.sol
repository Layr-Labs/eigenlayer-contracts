// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/AccessControlEnumerableUpgradeable.sol";
import "../interfaces/IPausable.sol";
import "./ProtocolRegistryStorage.sol";

contract ProtocolRegistry is Initializable, AccessControlEnumerableUpgradeable, ProtocolRegistryStorage {
    using ShortStringsUpgradeable for *;
    using EnumerableMap for EnumerableMap.UintToAddressMap;

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */
    constructor(
        IProxyAdmin proxyAdmin
    ) ProtocolRegistryStorage(proxyAdmin) {
        _disableInitializers();
    }

    /// @inheritdoc IProtocolRegistry
    function initialize(address initialAdmin, address pauserMultisig) external initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, initialAdmin);
        _grantRole(PAUSER_ROLE, pauserMultisig);
    }

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /// @inheritdoc IProtocolRegistry
    function ship(
        address[] calldata addresses,
        DeploymentConfig[] calldata configs,
        string[] calldata names,
        string calldata semanticVersion
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        // Update the semantic version.
        _updateSemanticVersion(semanticVersion);
        for (uint256 i = 0; i < addresses.length; ++i) {
            // Append each provided
            _appendDeployment(addresses[i], configs[i], names[i]);
        }
    }

    /// @inheritdoc IProtocolRegistry
    function configure(address addr, DeploymentConfig calldata config) external onlyRole(DEFAULT_ADMIN_ROLE) {
        // Update the config
        _deploymentConfigs[addr] = config;
        // Emit the event.
        emit DeploymentConfigured(addr, config);
    }

    /// @inheritdoc IProtocolRegistry
    function pauseAll() external onlyRole(PAUSER_ROLE) {
        uint256 length = totalDeployments();
        // Iterate over all stored deployments.
        for (uint256 i = 0; i < length; ++i) {
            (, address addr) = _deployments.at(i);
            DeploymentConfig memory config = _deploymentConfigs[addr];
            // Only attempt to pause deployments marked as pausable.
            if (config.pausable && !config.deprecated) {
                IPausable(addr).pauseAll();
            }
        }
    }

    /**
     *
     *                             HELPER FUNCTIONS
     *
     */

    /// @dev Updates the semantic version of the protocol.
    function _updateSemanticVersion(
        string calldata semanticVersion
    ) internal {
        string memory previousSemanticVersion = _semanticVersion.toString();
        _semanticVersion = semanticVersion.toShortString();
        emit SemanticVersionUpdated(previousSemanticVersion, semanticVersion);
    }

    /// @dev Appends a deployment.
    function _appendDeployment(address addr, DeploymentConfig calldata config, string calldata name) internal {
        // Store name => address mapping
        _deployments.set({key: _unwrap(name.toShortString()), value: addr});
        // Store deployment config
        _deploymentConfigs[addr] = config;
        // Emit the events.
        emit DeploymentShipped(addr, config);
    }

    /// @dev Unwraps a ShortString to a uint256.
    function _unwrap(
        ShortString shortString
    ) internal pure returns (uint256) {
        return uint256(ShortString.unwrap(shortString));
    }

    /**
     *
     *                              VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IProtocolRegistry
    function version() external view virtual returns (string memory) {
        return _semanticVersion.toString();
    }

    /// @inheritdoc IProtocolRegistry
    function majorVersion() external view returns (string memory) {
        bytes memory v = bytes(_semanticVersion.toString());
        return string(abi.encodePacked(v[0]));
    }

    /// @inheritdoc IProtocolRegistry
    function getAddress(
        string calldata name
    ) external view returns (address) {
        return _deployments.get(_unwrap(name.toShortString()));
    }

    /// @inheritdoc IProtocolRegistry
    function getDeployment(
        string calldata name
    ) external view returns (address addr, DeploymentConfig memory config) {
        addr = _deployments.get(_unwrap(name.toShortString()));
        config = _deploymentConfigs[addr];
        return (addr, config);
    }

    /// @inheritdoc IProtocolRegistry
    function getAllDeployments()
        external
        view
        returns (string[] memory names, address[] memory addresses, DeploymentConfig[] memory configs)
    {
        uint256 length = totalDeployments();
        names = new string[](length);
        addresses = new address[](length);
        configs = new DeploymentConfig[](length);

        for (uint256 i = 0; i < length; ++i) {
            (uint256 nameShortString, address addr) = _deployments.at(i);
            names[i] = ShortString.wrap(bytes32(nameShortString)).toString();
            addresses[i] = addr;
            configs[i] = _deploymentConfigs[addr];
        }

        return (names, addresses, configs);
    }

    /// @inheritdoc IProtocolRegistry
    function totalDeployments() public view returns (uint256) {
        return _deployments.length();
    }
}
