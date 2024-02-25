// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IClaimingManager.sol";
import "./IStrategy.sol";

/**
 * @title Interface for the `IPaymentCoordinator` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
interface IPaymentCoordinator {

    /// STRUCTS ///
	struct RangePayment {
        address avs;
		IStrategy strategy;
		IERC20 token;
		uint256 amount;
		uint256 startRangeTimestamp;
		uint256 endRangeTimestamp;
        bool toOperatorSet;
	}

    /// EVENTS ///

    event RangePaymentCreated(RangePayment rangePayment);

    /// VIEW FUNCTIONS ///

    /// @notice The address of the entity that can update the contract with new merkle roots
    function claimingManager() external view returns (IClaimingManager);

    /// @notice The interval in seconds at which the calculation for range payment distribution is done. ranges must be aligned to multiples of this interval
    function CALCULATION_INTERVAL_SECONDS() external view returns (uint32);

    /// @notice The maximum amount of time that a range payment can end in the future
    function MAX_FUTURE_RANGE_END() external view returns (uint32);

    /// @notice The lower bound for the start of a range payment
    function LOWER_BOUND_START_RANGE() external view returns (uint32);

    /// EXTERNAL FUNCTIONS ///

    /**
     * @notice Initializes the contract
     * @param initialOwner The address of the initial owner of the contract
     * @dev Only callable once
     */
    function initialize(address initialOwner) external;
    
    /**
     * @notice Creates a new range payment on behalf of an AVS
     * @param rangePayment The range payment being created
     * @dev The end time of the range must be at most 3 months in the future
     * @dev The start time of the range must be after the configured lower bound
     * @dev The tokens are sent to the claiming manager contract
     */
	function payForRange(RangePayment calldata rangePayment) external;
}
