// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./ISlasher.sol";

/**
 * @title Interface for a `ServiceManager`-type contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
interface IServiceManager {

    // ServiceManager proxies to the slasher
    function slasher() external view returns (ISlasher);

    /// @notice Returns the current 'taskNumber' for the middleware
    function taskNumber() external view returns (uint32);

    /// @notice function that causes the ServiceManager to freeze the operator on EigenLayer, through a call to the Slasher contract
    /// @dev this function should contain slashing logic, to make sure operators are not needlessly being slashed
    function freezeOperator(address operator) external;

    /// @notice proxy call to the slasher, recording an initial stake update (on operator registration)
    function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) external;

    /// @notice proxy call to the slasher, recording a stake update
    function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 prevElement) external;

    /// @notice proxy call to the slasher, recording a final stake update (on operator deregistration)
    function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) external;

    /// @notice Returns the latest block until which operators must serve (could be in the past or future).
    /// @dev this should be called and the response passed to the recordStakeUpdate functionss' serveUntilBlock parameter
    function latestServeUntilBlock() external view returns (uint32);

    /// @notice required since the registry contract will call this function to permission its upgrades to be done by the same owner as the service manager
    function owner() external view returns (address);
}