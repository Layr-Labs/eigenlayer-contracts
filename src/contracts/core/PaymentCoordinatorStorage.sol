// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/contracts/interfaces/IAVSDirectory.sol";
import "src/contracts/interfaces/IStrategyManager.sol";
import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/ISlasher.sol";
import "src/contracts/interfaces/IEigenPodManager.sol";
import "src/contracts/interfaces/IPaymentCoordinator.sol";

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

    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");
    uint64 public immutable MAX_PAYMENT_DURATION;
    uint64 public immutable LOWER_BOUND_START_RANGE;

    IAVSDirectory public immutable avsDirectory;

    /// @notice The elegationManager contract for EigenLayer
    IDelegationManager public immutable delegationManager;

    /// @notice The EigenPodManager contract for EigenLayer
    IEigenPodManager public immutable eigenPodManager;

    /// @notice The StrategyManager contract for EigenLayer
    IStrategyManager public immutable strategyManager;

    /// @notice The Slasher contract for EigenLayer
    ISlasher public immutable slasher;

    /*******************************************************************************
                                       STORAGE 
    *******************************************************************************/

    /**
     * @notice Original EIP-712 Domain separator for this contract.
     * @dev The domain separator may change in the event of a fork that modifies the ChainID.
     * Use the getter function `domainSeparator` to get the current domain separator for this contract.
     */
    bytes32 internal _DOMAIN_SEPARATOR;

    /// Slot 2
    /// @notice The address of the entity that can update the contract with new merkle roots
    address public paymentUpdater;
    /// @notice The interval in seconds at which the calculation for range payment distribution is done.
    /// @dev Payment durations must be multiples of this interval.
    uint64 public calculationIntervalSeconds;

    /// Slot 3
    /// @notice Delay in timestamp before a posted root can be claimed against
    uint64 public activationDelay;
    /// @notice the commission for all operators across all avss
    uint16 public globalOperatorCommissionBips;
    /// @notice Payments must start at earliest now except unless retroactive payments enabled
    bool public retroactivePaymentsEnabled;

    /// @notice Mapping: account => the address of the entity to which new payments are directed on behalf of the account
    mapping(address => address) public claimerFor;

    /// @notice Mapping: claimer => token => total amount claimed
    mapping(address => mapping(IERC20 => uint256)) public cumulativeClaimed;
    /// @notice Used for unique rangePaymentHashes
    uint256 public paymentNonce;
    /// @notice Mapping: avs => rangePaymentHash => bool to check if range payment hash has been submitted
    mapping(address => mapping(bytes32 => bool)) public isRangePaymentHash;
    /// @notice Mapping: address => bool to check if the address is permissioned to submit payAllForRange
    mapping(address => bool) public isPayAllForRangeSubmitter;

    constructor(
        IAVSDirectory _avsDirectory,
        IDelegationManager _delegationManager,
        IEigenPodManager _eigenPodManager,
        IStrategyManager _strategyManager,
        ISlasher _slasher,
        uint64 _MAX_PAYMENT_DURATION,
        uint64 _LOWER_BOUND_START_RANGE
    ) {
        avsDirectory = _avsDirectory;
        delegationManager = _delegationManager;
        eigenPodManager = _eigenPodManager;
        strategyManager = _strategyManager;
        slasher = _slasher;
        MAX_PAYMENT_DURATION = _MAX_PAYMENT_DURATION;
        LOWER_BOUND_START_RANGE = _LOWER_BOUND_START_RANGE;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[43] private __gap;
}
