// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "../mixins/SemVerMixin.sol";
import "./StrategyFactoryStorage.sol";
import "./StrategyBase.sol";
import "../permissions/Pausable.sol";

/**
 * @title Factory contract for deploying BeaconProxies of a Strategy contract implementation for arbitrary ERC20 tokens
 *        and automatically adding them to the StrategyWhitelist in EigenLayer.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @dev This may not be compatible with non-standard ERC20 tokens. Caution is warranted.
 */
contract StrategyFactory is StrategyFactoryStorage, OwnableUpgradeable, Pausable, SemVerMixin {
    uint8 internal constant PAUSED_NEW_STRATEGIES = 0;

    /// @notice EigenLayer's StrategyManager contract
    IStrategyManager public immutable strategyManager;

    /// @notice Since this contract is designed to be initializable, the constructor simply sets the immutable variables.
    constructor(
        IStrategyManager _strategyManager,
        IPauserRegistry _pauserRegistry,
        string memory _version
    ) Pausable(_pauserRegistry) SemVerMixin(_version) {
        strategyManager = _strategyManager;
        _disableInitializers();
    }

    function initialize(
        address _initialOwner,
        uint256 _initialPausedStatus,
        IBeacon _strategyBeacon
    ) public virtual initializer {
        _transferOwnership(_initialOwner);
        _setPausedStatus(_initialPausedStatus);
        _setStrategyBeacon(_strategyBeacon);
    }

    /**
     * @notice Deploy a new StrategyBase contract for the ERC20 token, using a beacon proxy
     * @dev A strategy contract must not yet exist for the token.
     * @dev Immense caution is warranted for non-standard ERC20 tokens, particularly "reentrant" tokens
     * like those that conform to ERC777.
     */
    function deployNewStrategy(
        IERC20 token
    ) external onlyWhenNotPaused(PAUSED_NEW_STRATEGIES) returns (IStrategy newStrategy) {
        require(!isBlacklisted[token], BlacklistedToken());
        require(deployedStrategies[token] == IStrategy(address(0)), StrategyAlreadyExists());
        IStrategy strategy = IStrategy(
            address(
                new BeaconProxy(
                    address(strategyBeacon), abi.encodeWithSelector(StrategyBase.initialize.selector, token)
                )
            )
        );
        _setStrategyForToken(token, strategy);
        IStrategy[] memory strategiesToWhitelist = new IStrategy[](1);
        strategiesToWhitelist[0] = strategy;
        strategyManager.addStrategiesToDepositWhitelist(strategiesToWhitelist);
        return strategy;
    }

    /**
     * @notice Owner-only function to prevent strategies from being created for given tokens.
     * @param tokens An array of token addresses to blacklist.
     */
    function blacklistTokens(
        IERC20[] calldata tokens
    ) external onlyOwner {
        IStrategy[] memory strategiesToRemove = new IStrategy[](tokens.length);
        uint256 removeIdx = 0;

        for (uint256 i; i < tokens.length; ++i) {
            require(!isBlacklisted[tokens[i]], AlreadyBlacklisted());
            isBlacklisted[tokens[i]] = true;
            emit TokenBlacklisted(tokens[i]);

            // If someone has already deployed a strategy for this token, add it
            // to the list of strategies to remove from the StrategyManager whitelist
            IStrategy deployed = deployedStrategies[tokens[i]];
            if (deployed != IStrategy(address(0))) {
                strategiesToRemove[removeIdx] = deployed;
                removeIdx++;
            }
        }

        // Manually adjust length to remove unused entries
        // New length == removeIdx
        assembly {
            mstore(strategiesToRemove, removeIdx)
        }

        if (strategiesToRemove.length > 0) {
            strategyManager.removeStrategiesFromDepositWhitelist(strategiesToRemove);
        }
    }

    /**
     * @notice Owner-only function to pass through a call to `StrategyManager.addStrategiesToDepositWhitelist`
     */
    function whitelistStrategies(
        IStrategy[] calldata strategiesToWhitelist
    ) external onlyOwner {
        strategyManager.addStrategiesToDepositWhitelist(strategiesToWhitelist);
    }

    /**
     * @notice Owner-only function to pass through a call to `StrategyManager.removeStrategiesFromDepositWhitelist`
     */
    function removeStrategiesFromWhitelist(
        IStrategy[] calldata strategiesToRemoveFromWhitelist
    ) external onlyOwner {
        strategyManager.removeStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist);
    }

    function _setStrategyForToken(IERC20 token, IStrategy strategy) internal {
        deployedStrategies[token] = strategy;
        emit StrategySetForToken(token, strategy);
    }

    function _setStrategyBeacon(
        IBeacon _strategyBeacon
    ) internal {
        emit StrategyBeaconModified(strategyBeacon, _strategyBeacon);
        strategyBeacon = _strategyBeacon;
    }
}
