// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";

library ShareScalingLib {
    /// @dev The initial total magnitude for an operator
    uint64 public constant INITIAL_TOTAL_MAGNITUDE = 1 ether;

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
            scaledShares[i] = shares[i] * INITIAL_TOTAL_MAGNITUDE / totalMagnitudes[i];
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
            descaledShares[i] = scaledShares[i] * totalMagnitudes[i] / INITIAL_TOTAL_MAGNITUDE;
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
        return shares * INITIAL_TOTAL_MAGNITUDE / totalMagnitudes[0];
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
        return scaledShares * totalMagnitudes[0] / INITIAL_TOTAL_MAGNITUDE;
    }
}