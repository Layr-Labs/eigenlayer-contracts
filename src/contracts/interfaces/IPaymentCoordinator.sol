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
		IStrategy strategy;
		IERC20 token;
		uint256 amount;
		uint256 startRangeTimestamp;
        // duration in seconds
		uint256 duration;
	}

    /// EVENTS ///

    event RangePaymentCreated(RangePayment rangePayment);

    /// VIEW FUNCTIONS ///

    /// @notice The address of the entity that can update the contract with new merkle roots
    function claimingManager() external view returns (IClaimingManager);

    // @notice The owner of this contract
    function owner() external view returns (address);

    /// @notice The interval in seconds at which the calculation for range payment distribution is done. ranges must be aligned to multiples of this interval
    function calculationIntervalSeconds() external view returns (uint32);

    /// @notice The maximum amount of time that a range payment can end in the future
    function MAX_DURATION() external view returns (uint32);

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
     * @notice Creates a new range payment on behalf of an AVS, to be split amongst the
     * set of stakers delegated to operators who are registered to the `avs`
     * @param rangePayment The range payment being created
     * @param avs The ServiceManager of the AVS on behalf of which the payment is being made
     * @dev The duration of the `rangePayment` cannot exceed `MAX_DURATION`
     * @dev The tokens are sent to the `claimingManager` contract
     */
	function payForRange(RangePayment calldata rangePayment, address avs) external;

    /**
     * @notice similar to `payForRange` except the payment is split amongst *all* stakers
     * rather than just those delegated to operators who are registered to a single avs
     */
    function payAllForRange(RangePayment calldata rangePayment, address avs) external;
}
