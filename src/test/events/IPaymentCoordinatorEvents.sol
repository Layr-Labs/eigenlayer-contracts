// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/contracts/interfaces/IPaymentCoordinator.sol";

interface IPaymentCoordinatorEvents {
    /// EVENTS ///

    /// @notice emitted when an AVS creates a valid RangePayment
    event RangePaymentCreated(
        address indexed avs,
        uint256 indexed paymentNonce,
        bytes32 indexed rangePaymentHash,
        RangePayment rangePayment
    );
    /// @notice emitted when a valid RangePayment is created for all stakers by a valid submitter
    event RangePaymentForAllCreated(
        address indexed submitter,
        uint256 indexed paymentNonce,
        bytes32 indexed rangePaymentHash,
        RangePayment rangePayment
    );
    /// @notice paymentUpdater is responsible for submiting DistributionRoots, only owner can set paymentUpdater
    event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater);
    event PayAllForRangeSubmitterSet(
        address indexed payAllForRangeSubmitter,
        bool indexed oldValue,
        bool indexed newValue
    );
    event ActivationDelaySet(uint64 oldActivationDelay, uint64 newActivationDelay);
    event CalculationIntervalSecondsSet(uint64 oldCalculationIntervalSeconds, uint64 newCalculationIntervalSeconds);
    event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips);
    event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer);
    /// @notice rootIndex is the specific array index of the newly created root in the storage array 
    event DistributionRootSubmitted(
        uint32 indexed rootIndex,
        bytes32 indexed root,
        uint64 paymentCalculationStartTimestamp,
        uint64 paymentCalculationEndTimestamp,
        uint64 activatedAt
    );
    /// @notice earnerTokenRoot is the specific earner root hash that the claim is proven aganst.
    /// root is the DistributionRoot of that the earner subtree is included in.
    event PaymentClaimed(bytes32 indexed root, bytes32 indexed earnerTokenRoot, ClaimsTreeMerkleLeaf leaf);

}
