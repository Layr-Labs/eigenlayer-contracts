// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IRegistryCoordinator.sol";
import "../interfaces/IBLSPublicKeyCompendium.sol";

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

abstract contract BLSPubkeyRegistryStorage is Initializable, IBLSPubkeyRegistry {
    /// @notice the hash of the zero pubkey aka BN254.G1Point(0,0)
    bytes32 internal constant ZERO_PK_HASH = hex"ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5";
    /// @notice the registry coordinator contract
    IRegistryCoordinator public immutable registryCoordinator;
    /// @notice the BLSPublicKeyCompendium contract against which pubkey ownership is checked
    IBLSPublicKeyCompendium public immutable pubkeyCompendium;

    /// @notice mapping of quorumNumber => ApkUpdate[], tracking the aggregate pubkey updates of every quorum
    mapping(uint8 => ApkUpdate[]) public quorumApkUpdates;
    /// @notice mapping of quorumNumber => current aggregate pubkey of quorum
    mapping(uint8 => BN254.G1Point) public quorumApk;

    constructor(IRegistryCoordinator _registryCoordinator, IBLSPublicKeyCompendium _pubkeyCompendium) {
        registryCoordinator = _registryCoordinator;
        pubkeyCompendium = _pubkeyCompendium;
        // disable initializers so that the implementation contract cannot be initialized
        _disableInitializers();
    }

    // storage gap for upgradeability
    uint256[48] private __GAP;
}
