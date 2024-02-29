// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "../permissions/Pausable.sol";
import "../libraries/Merkle.sol";

import "../interfaces/IPaymentCoordinator.sol";

contract PaymentCoordinator is IPaymentCoordinator, Initializable, OwnableUpgradeable, Pausable {

    /// @notice The core claiming manager contract
    IClaimingManager public immutable claimingManager;

    /// @notice The interval in seconds at which the calculation for range payment distribution is done. ranges must be aligned to multiples of this interval
    uint32 public immutable CALCULATION_INTERVAL_SECONDS;

    /// @notice The maximum amount of time that a range payment can end in the future
    uint32 public immutable MAX_FUTURE_RANGE_END;

    /// @notice The lower bound for the start of a range payment
    uint32 public immutable LOWER_BOUND_START_RANGE; // todo

    /// EXTERNAL FUNCTIONS ///

    constructor(IClaimingManager _claimingManager, uint32 _CALCULATION_INTERVAL_SECONDS, uint32 _MAX_FUTURE_RANGE_END, uint32 _LOWER_BOUND_START_RANGE) {
        claimingManager = _claimingManager;
        CALCULATION_INTERVAL_SECONDS = _CALCULATION_INTERVAL_SECONDS;
        MAX_FUTURE_RANGE_END = _MAX_FUTURE_RANGE_END;
        LOWER_BOUND_START_RANGE = _LOWER_BOUND_START_RANGE;
    }

    /**
     * @notice Initializes the contract
     * @param initialOwner The address of the initial owner of the contract
     * @dev Only callable once
     */
    function initialize(address initialOwner) external initializer {
        __Ownable_init();
        transferOwnership(initialOwner);
    }

    /**
     * @notice Creates a new range payment on behalf of an AVS
     * @param rangePayment The range payment being created
     * @dev The end time of the range must be at most 3 months in the future
     * @dev The start time of the range must be after the configured lower bound
     * @dev The tokens are sent to the claiming manager contract
     */
	function payForRange(RangePayment calldata rangePayment) external {
        // Ensure the start and end timestamps are multiples of the granularity
        require(rangePayment.startRangeTimestamp % CALCULATION_INTERVAL_SECONDS == 0, "PaymentCoordinator.payForRange: startRangeTimestamp not aligned");
        require(rangePayment.endRangeTimestamp % CALCULATION_INTERVAL_SECONDS == 0, "PaymentCoordinator.payForRange: endRangeTimestamp not aligned");

        require(rangePayment.endRangeTimestamp >= rangePayment.startRangeTimestamp, "PaymentCoordinator.payForRange: endRangeTimestamp before startRangeTimestamp");
        require(rangePayment.startRangeTimestamp >= LOWER_BOUND_START_RANGE, "PaymentCoordinator.payForRange: startRangeTimestamp too early");
        require(rangePayment.endRangeTimestamp <= block.timestamp + MAX_FUTURE_RANGE_END, "PaymentCoordinator.payForRange: endRangeTimestamp too late");

        // TODO: check that the strategy is valid?

        // transfer the tokens to the claiming manager
        rangePayment.token.transferFrom(msg.sender, address(claimingManager), rangePayment.amount);

        emit RangePaymentCreated(rangePayment);
    }
}