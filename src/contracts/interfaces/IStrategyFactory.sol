// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IStrategy.sol";
import "./IDurationVaultStrategy.sol";
import "./ISemVerMixin.sol";

/// @title Interface for the `StrategyFactory` contract.
/// @author Layr Labs, Inc.
/// @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
/// @dev This may not be compatible with non-standard ERC20 tokens. Caution is warranted.
interface IStrategyFactory {
    /// @dev Thrown when attempting to deploy a strategy for a blacklisted token.
    error BlacklistedToken();
    /// @dev Thrown when attempting to deploy a strategy that already exists.
    error StrategyAlreadyExists();
    /// @dev Thrown when attempting to blacklist a token that is already blacklisted
    error AlreadyBlacklisted();

    event TokenBlacklisted(IERC20 token);

    /// @notice Upgradeable beacon which new Strategies deployed by this contract point to
    function strategyBeacon() external view returns (IBeacon);

    /// @notice Upgradeable beacon which duration vault strategies deployed by this contract point to
    function durationVaultBeacon() external view returns (IBeacon);

    /// @notice Mapping token => Strategy contract for the token
    /// The strategies in this mapping are deployed by the StrategyFactory.
    /// The factory can only deploy a single strategy per token address
    /// These strategies MIGHT not be whitelisted in the StrategyManager,
    /// though deployNewStrategy does whitelist by default.
    /// These strategies MIGHT not be the only strategy for the underlying token
    /// as additional strategies can be whitelisted by the owner of the factory.
    function deployedStrategies(
        IERC20 token
    ) external view returns (IStrategy);

    /// @notice Deploy a new strategyBeacon contract for the ERC20 token.
    /// @param token the token to deploy a strategy for
    /// @dev A strategy contract must not yet exist for the token.
    /// $dev Immense caution is warranted for non-standard ERC20 tokens, particularly "reentrant" tokens
    /// like those that conform to ERC777.
    function deployNewStrategy(
        IERC20 token
    ) external returns (IStrategy newStrategy);

    /// @notice Deploys a new duration-bound vault strategy contract.
    /// @dev Enforces the same blacklist semantics as vanilla strategies.
    function deployDurationVaultStrategy(
        IDurationVaultStrategy.VaultConfig calldata config
    ) external returns (IDurationVaultStrategy newVault);

    /// @notice Returns all duration vaults that have ever been deployed for a given token.
    function getDurationVaults(
        IERC20 token
    ) external view returns (IDurationVaultStrategy[] memory);

    /// @notice Owner-only function to pass through a call to `StrategyManager.addStrategiesToDepositWhitelist`
    function whitelistStrategies(
        IStrategy[] calldata strategiesToWhitelist
    ) external;

    /// @notice Owner-only function to pass through a call to `StrategyManager.removeStrategiesFromDepositWhitelist`
    function removeStrategiesFromWhitelist(
        IStrategy[] calldata strategiesToRemoveFromWhitelist
    ) external;

    /// @notice Emitted whenever a slot is set in the `tokenStrategy` mapping
    event StrategySetForToken(IERC20 token, IStrategy strategy);

    /// @notice Emitted whenever a duration vault is deployed. The vault address uniquely identifies the deployment.
    event DurationVaultDeployed(
        IDurationVaultStrategy indexed vault,
        IERC20 indexed underlyingToken,
        address indexed vaultAdmin,
        uint32 duration,
        uint256 maxPerDeposit,
        uint256 stakeCap,
        string metadataURI,
        address operatorSetAVS,
        uint32 operatorSetId
    );
}
