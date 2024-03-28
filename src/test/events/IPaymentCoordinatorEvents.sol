// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/contracts/interfaces/IPaymentCoordinator.sol";

interface IPaymentCoordinatorEvents {
    event RangePaymentCreated(
        address indexed avs,
        bytes32 rangePaymentHash,
        IPaymentCoordinator.RangePayment rangePayment
    );
    event RangePaymentForAllCreated(
        address indexed submitter,
        bytes32 rangePaymentHash,
        IPaymentCoordinator.RangePayment rangePayment
    );
    event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater);
    event PayAllForRangeSubmitterSet(
        address indexed payAllForRangeSubmitter,
        bool indexed oldValue,
        bool indexed newValue
    );
    event ActivationDelaySet(uint64 oldActivationDelay, uint64 newActivationDelay);
    event CalculationIntervalSecondsSet(uint64 oldCalculationIntervalSeconds, uint64 newCalculationIntervalSeconds);
    event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips);
    event RecipientSet(address indexed account, address indexed recipient);
    event RootSubmitted(
        bytes32 root,
        uint32 paymentCalculationStartTimestamp,
        uint32 paymentCalculationEndTimestamp,
        uint32 activatedAt
    );
    event PaymentClaimed(IPaymentCoordinator.ClaimsTreeMerkleLeaf leaf);
}
