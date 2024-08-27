// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";

library ShareScalingLib {
    /// @dev The initial total magnitude for an operator
    uint64 public constant INITIAL_TOTAL_MAGNITUDE = 1e18;

    /**
     * @notice Returns the scaled shares of for an operator for a particular set of strategies and shares delegated to an operator
     * @param avsDirectory The AVS directory
     * @param operator The operator to scale shares for
     * @param strategies The strategies of which shares are
     * @param shares The shares to scale
     */
    function scaleShares(
        IAVSDirectory avsDirectory,
        address operator,
        IStrategy[] memory strategies,
        uint256[] memory shares
    ) internal view returns (uint256[] memory) {
        uint64[] memory totalMagnitudes = avsDirectory.getTotalMagnitudes(operator, strategies);
        uint256[] memory scaledShares = new uint256[](shares.length);
        for (uint256 i = 0; i < shares.length; i++) {
            scaledShares[i] = _scaleShares(shares[i], totalMagnitudes[i]);
        }
        return scaledShares;
    }

    /**
     * @notice Returns the shares of for an operator for a particular set of strategies and scaled shares delegated to an operator
     * @param avsDirectory The AVS directory
     * @param operator The operator to scale shares for
     * @param strategies The strategies of which scaled shares are
     * @param scaledShares The scaled shares to descale
     */
    function descaleShares(
        IAVSDirectory avsDirectory,
        address operator,
        IStrategy[] memory strategies,
        uint256[] memory scaledShares
    ) internal view returns (uint256[] memory) {
        uint64[] memory totalMagnitudes = avsDirectory.getTotalMagnitudes(operator, strategies);
        uint256[] memory descaledShares = new uint256[](scaledShares.length);
        for (uint256 i = 0; i < scaledShares.length; i++) {
            descaledShares[i] = _descaleShares(scaledShares[i], totalMagnitudes[i]);
        }
        return descaledShares;
    }

    /**
     * @notice Returns the shares of for an operator for a particular set of strategies and scaled shares delegated to an operator
     * @param avsDirectory The AVS directory
     * @param operator The operator to scale shares for
     * @param strategies The strategies of which scaled shares are
     * @param scaledShares The scaled shares to descale
     */
    function descaleSharesAtTimestamp(
        IAVSDirectory avsDirectory,
        address operator,
        IStrategy[] memory strategies,
        uint256[] memory scaledShares,
        uint32 timestamp
    ) internal view returns (uint256[] memory) {
        uint64[] memory totalMagnitudes = avsDirectory.getTotalMagnitudesAtTimestamp(operator, strategies, timestamp);
        uint256[] memory descaledShares = new uint256[](scaledShares.length);
        for (uint256 i = 0; i < scaledShares.length; i++) {
            descaledShares[i] = _descaleShares(scaledShares[i], totalMagnitudes[i]);
        }
        return descaledShares;
    }

    /**
     * @notice Returns the scaled shares of for an operator for a particular strategy and shares delegated to an operator
     * @param avsDirectory The AVS directory
     * @param operator The operator to scale shares for
     * @param strategy The strategy of which shares are
     * @param shares The shares to scale
     */
    function scaleShares(
        IAVSDirectory avsDirectory,
        address operator,
        IStrategy strategy,
        uint256 shares
    ) internal view returns (uint256) {
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategy;
        uint64[] memory totalMagnitudes = avsDirectory.getTotalMagnitudes(operator, strategies);
        return _scaleShares(shares, totalMagnitudes[0]);
    }

    /**
     * @notice Returns the shares of for an operator for a particular strategy and scaled shares delegated to an operator
     * @param avsDirectory The AVS directory
     * @param operator The operator to scale shares for
     * @param strategy The strategy of which scaled shares are
     * @param scaledShares The scaled shares to descale
     */
    function descaleShares(
        IAVSDirectory avsDirectory,
        address operator,
        IStrategy strategy,
        uint256 scaledShares
    ) internal view returns (uint256) {
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategy;
        uint64[] memory totalMagnitudes = avsDirectory.getTotalMagnitudes(operator, strategies);
        return _descaleShares(scaledShares, totalMagnitudes[0]);
    }

    /// @notice helper pure to return scaledShares given shares and totalMagnitude
    function _scaleShares(uint256 shares, uint64 totalMagnitude) internal pure returns (uint256) {
        return shares * INITIAL_TOTAL_MAGNITUDE / totalMagnitude;
    }

    /// @notice helper pure to return shares(descaledShares) given scaledShares and totalMagnitude
    function _descaleShares(uint256 scaledShares, uint64 totalMagnitude) internal pure returns (uint256) {
        return scaledShares * totalMagnitude / INITIAL_TOTAL_MAGNITUDE;
    }
}
