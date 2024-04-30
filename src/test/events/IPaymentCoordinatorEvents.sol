// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/contracts/interfaces/IPaymentCoordinator.sol";

interface IPaymentCoordinatorEvents {
    /// EVENTS ///

    /// @notice emitted when an AVS creates a valid RangePayment
    event RangePaymentCreated(
        address indexed avs,
        uint256 indexed paymentNonce,
        bytes32 indexed rangePaymentHash,
        IPaymentCoordinator.RangePayment rangePayment
    );
    /// @notice emitted when a valid RangePayment is created for all stakers by a valid submitter
    event RangePaymentForAllCreated(
        address indexed submitter,
        uint256 indexed paymentNonce,
        bytes32 indexed rangePaymentHash,
        IPaymentCoordinator.RangePayment rangePayment
    );
    /// @notice paymentUpdater is responsible for submiting DistributionRoots, only owner can set paymentUpdater
    event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater);
    event PayAllForRangeSubmitterSet(
        address indexed payAllForRangeSubmitter,
        bool indexed oldValue,
        bool indexed newValue
    );
    event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay);
    event CalculationIntervalSecondsSet(uint32 oldCalculationIntervalSeconds, uint32 newCalculationIntervalSeconds);
    event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips);
    event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer);
    /// @notice rootIndex is the specific array index of the newly created root in the storage array
    event DistributionRootSubmitted(
        uint32 indexed rootIndex,
        bytes32 indexed root,
        uint32 indexed paymentCalculationEndTimestamp,
        uint32 activatedAt
    );
    /// @notice root is one of the submitted distribution roots that was claimed against
    event PaymentClaimed(
        bytes32 root,
        address indexed earner,
        address indexed claimer,
        address indexed recipient,
        IERC20 token,
        uint256 claimedAmount
    );



    /// TOKEN EVENTS FOR TESTING ///
    /**
     * @dev Emitted when `value` tokens are moved from one account (`from`) to
     * another (`to`).
     *
     * Note that `value` may be zero.
     */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /**
     * @dev Emitted when the allowance of a `spender` for an `owner` is set by
     * a call to {approve}. `value` is the new allowance.
     */
    event Approval(address indexed owner, address indexed spender, uint256 value);
}
