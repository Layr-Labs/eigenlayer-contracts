// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IStrategyManager.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IEigenPodManager.sol";
import "../interfaces/IPaymentCoordinator.sol";

/**
 * @title Storage variables for the `PaymentCoordinator` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract PaymentCoordinatorStorage is IPaymentCoordinator {

    /*******************************************************************************
                               CONSTANTS AND IMMUTABLES 
    *******************************************************************************/

    /// @notice The interval in seconds at which the calculation for range payment distribution is done.
    /// @dev Payment durations must be multiples of this interval. This is going to be configured to 1 week
    uint32 public immutable CALCULATION_INTERVAL_SECONDS;
    /// @notice The maximum amount of time (seconds) that a range payment can span over
    uint32 public immutable MAX_PAYMENT_DURATION;
    /// @notice max amount of time (seconds) that a payment can start in the past
    uint32 public immutable MAX_RETROACTIVE_LENGTH;
    /// @notice max amount of time (seconds) that a payment can start in the future
    uint32 public immutable MAX_FUTURE_LENGTH;
    /// @notice absolute min timestamp (seconds) that a payment can start at
    uint32 public immutable GENESIS_PAYMENT_TIMESTAMP;
    /// @notice The cadence at which a snapshot is taken offchain for calculating payment distributions
    uint32 internal constant SNAPSHOT_CADENCE = 1 days;   

    /// @notice The DelegationManager contract for EigenLayer
    IDelegationManager public immutable delegationManager;

    /// @notice The StrategyManager contract for EigenLayer
    IStrategyManager public immutable strategyManager;

    /*******************************************************************************
                                       STORAGE 
    *******************************************************************************/

    /**
     * @notice Original EIP-712 Domain separator for this contract.
     * @dev The domain separator may change in the event of a fork that modifies the ChainID.
     * Use the getter function `domainSeparator` to get the current domain separator for this contract.
     */
    bytes32 internal _DOMAIN_SEPARATOR;

    /// @notice list of roots submitted by the paymentUpdater
    DistributionRoot[] public distributionRoots;

    /// Slot 3
    /// @notice The address of the entity that can update the contract with new merkle roots
    address public paymentUpdater;
    /// @notice Delay in timestamp (seconds) before a posted root can be claimed against
    uint32 public activationDelay;
    /// @notice Timestamp for last submitted DistributionRoot
    uint32 public currPaymentCalculationEndTimestamp;

    /// Slot 4
    /// @notice the commission for all operators across all avss
    uint16 public globalOperatorCommissionBips;

    /// @notice Mapping: earner => the address of the entity who can call `processClaim` on behalf of the earner
    mapping(address => address) public claimerFor;

    /// @notice Mapping: earner => token => total amount claimed
    mapping(address => mapping(IERC20 => uint256)) public cumulativeClaimed;

    /// @notice Used for unique rangePaymentHashes per AVS and for PayAllForRangeSubmitters
    mapping(address => uint256) public paymentNonce;
    /// @notice Mapping: avs => rangePaymentHash => bool to check if range payment hash has been submitted
    mapping(address => mapping(bytes32 => bool)) public isRangePaymentHash;
    /// @notice Mapping: avs => rangePaymentForALlHash => bool to check if range payment hash for all has been submitted
    mapping(address => mapping(bytes32 => bool)) public isRangePaymentForAllHash;
    /// @notice Mapping: address => bool to check if the address is permissioned to submit payAllForRange
    mapping(address => bool) public isPayAllForRangeSubmitter;

    constructor(
        IDelegationManager _delegationManager,
        IStrategyManager _strategyManager,
        uint32 _CALCULATION_INTERVAL_SECONDS,
        uint32 _MAX_PAYMENT_DURATION,
        uint32 _MAX_RETROACTIVE_LENGTH,
        uint32 _MAX_FUTURE_LENGTH,
        uint32 _GENESIS_PAYMENT_TIMESTAMP
    ) {
        require(
            _GENESIS_PAYMENT_TIMESTAMP % _CALCULATION_INTERVAL_SECONDS == 0,
            "PaymentCoordinator: GENESIS_PAYMENT_TIMESTAMP must be a multiple of CALCULATION_INTERVAL_SECONDS"
        );
        require(
            _CALCULATION_INTERVAL_SECONDS % SNAPSHOT_CADENCE == 0,
            "PaymentCoordinator: CALCULATION_INTERVAL_SECONDS must be a multiple of SNAPSHOT_CADENCE"
        );
        delegationManager = _delegationManager;
        strategyManager = _strategyManager;
        CALCULATION_INTERVAL_SECONDS = _CALCULATION_INTERVAL_SECONDS;
        MAX_PAYMENT_DURATION = _MAX_PAYMENT_DURATION;
        MAX_RETROACTIVE_LENGTH = _MAX_RETROACTIVE_LENGTH;
        MAX_FUTURE_LENGTH = _MAX_FUTURE_LENGTH;
        GENESIS_PAYMENT_TIMESTAMP = _GENESIS_PAYMENT_TIMESTAMP;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[40] private __gap;
}
