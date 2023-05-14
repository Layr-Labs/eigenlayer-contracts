// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../../contracts/interfaces/IStrategy.sol";

/**
 * @title Minimal interface for an `Strategy` contract.
 * @author Layr Labs, Inc.
 * @notice Custom `Strategy` implementations may expand extensively on this interface.
 */
 contract StrategyMock is IStrategy {

    bool public withdrawn;

    function deposit(IERC20 token, uint256 amount) external returns (uint256){}

    function withdraw(address depositor, IERC20 token, uint256 amountShares) external{
        withdrawn = true;
    }

    function sharesToUnderlying(uint256 amountShares) external returns (uint256){}

    function underlyingToShares(uint256 amountUnderlying) external returns (uint256){}

    function userUnderlying(address user) external returns (uint256){}

    function sharesToUnderlyingView(uint256 amountShares) external view returns (uint256){}

    function underlyingToSharesView(uint256 amountUnderlying) external view returns (uint256){}

    function userUnderlyingView(address user) external view returns (uint256){}

    function underlyingToken() external view returns (IERC20){}

    /// @notice The total number of extant shares in thie Strategy
    function totalShares() external view returns (uint256){}

    /// @notice Returns either a brief string explaining the strategy's goal & purpose, or a link to metadata that explains in more detail.
    function explanation() external view returns (string memory){}
}
