// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IStrategy.sol";

/**
 * @title Interface for the `IPaymentCoordinator` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
interface IPaymentCoordinator {

    /// STRUCTS ///
	struct RangePayment {
        IStrategy[] strategies;
        // relative weights of shares in the strategies
        uint256[] weights;
        IERC20 token;
        uint256 amount;
        uint256 startTimestamp;
        // duration in seconds
        uint256 duration;
	}

    struct DistributionRoot {
        // timestamp after which the root can be claimed against
        uint32 activatedAfter;
        // merkle root of the distribution
        bytes32 root;
    }

    struct MerkleLeaf {
        IERC20 token; 
        uint256 amount;
        // Explicit recipient of the claim
        address recipient;
    }

    struct PaymentMerkleClaim {
        MerkleLeaf leaf;
        // The index of the root in the list of roots
        uint32 rootIndex;
        // The index of the leaf in the merkle tree for this root
        uint32 leafIndex;
        bytes proof;
    }

    /// EVENTS ///

    event RangePaymentCreated(RangePayment rangePayment);
    event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater);
    event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay);
    event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips);
    event RecipientSet(address indexed account, address indexed recipient);
    event RootSubmitted(bytes32 root, uint32 paymentsCalculatedUntilTimestamp, uint32 activatedAfter);
    event PaymentClaimed(MerkleLeaf leaf);

    /// VIEW FUNCTIONS ///

    /// @notice The address of the entity that can update the contract with new merkle roots
    function paymentUpdater() external view returns (address);

    // @notice The owner of this contract
    function owner() external view returns (address);

    /**
     * @notice The interval in seconds at which the calculation for range payment distribution is done.
     * @dev Payment durations must be multiples of this interval.
     */
    function calculationIntervalSeconds() external view returns (uint32);

    /// @notice The maximum amount of time that a range payment can end in the future
    function MAX_PAYMENT_DURATION() external view returns (uint32);

    /// @notice The lower bound for the start of a range payment
    function LOWER_BOUND_START_RANGE() external view returns (uint32);

    /// @notice Delay in timestamp before a posted root can be claimed against
    function activationDelay() external view returns (uint32);

    /// @notice Mapping: account => the address of the entity to which new payments are directed on behalf of the account
    function recipientOf(address account) external view returns (address);

    /// @notice Mapping: recipient => token => total amount claimed
    function cumulativeClaimed(address recipient, IERC20 token) external view returns (uint256);

    /// @notice the commission for all operators across all avss
    function globalOperatorCommissionBips() external view returns (uint16);

    /// @notice returns the hash of the leaf
    function findLeafHash(MerkleLeaf calldata leaf) external view returns (bytes32);

    /// @notice returns 'true' if the claim would currently pass the check in `processClaims`
    function checkClaim(PaymentMerkleClaim calldata claim) external view returns (bool);

    /// EXTERNAL FUNCTIONS ///

    /**
     * @notice Initializes the contract
     * @param initialOwner The address of the initial owner of the contract
     * @param _activationDelay Delay in timestamp before a posted root can be claimed against
     * @param _globalCommissionBips The commission for all operators across all avss
     * @dev Only callable once
     */
    function initialize(address initialOwner, uint32 _activationDelay, uint16 _globalCommissionBips) external;
    
    /**
     * @notice Creates a new range payment on behalf of an AVS, to be split amongst the
     * set of stakers delegated to operators who are registered to the `avs`
     * @param rangePayment The range payment being created
     * @dev Expected to be called by the ServiceManager of the AVS on behalf of which the payment is being made
     * @dev The duration of the `rangePayment` cannot exceed `MAX_PAYMENT_DURATION`
     * @dev The tokens are sent to the `claimingManager` contract
     * @dev This function will revert if the `rangePayment` is malformed,
     * e.g. if the `strategies` and `weights` arrays are of non-equal lengths
     */
	function payForRange(RangePayment calldata rangePayment) external;

    /**
     * @notice similar to `payForRange` except the payment is split amongst *all* stakers
     * rather than just those delegated to operators who are registered to a single avs
     */
    function payAllForRange(RangePayment calldata rangePayment) external;


    /**
     * @notice Sets the delay in timestamp before a posted root can be claimed against
     * @param _activationDelay Delay in timestamp before a posted root can be claimed against
     * @dev Only callable by the contract owner
     */
    function setActivationDelay(uint32 _activationDelay) external;

    /**
     * @notice Sets the global commission for all operators across all avss
     * @param _globalCommissionBips The commission for all operators across all avss
     * @dev Only callable by the contract owner
     */
    function setGlobalOperatorCommission(uint16 _globalCommissionBips) external;

    /**
     * @notice Sets the address of the entity that can claim payments on behalf of the account
     * @param account The account whose recipient is being set
     * @param recipient The address of the entity that can claim payments on behalf of the account
     * @dev Only callable by the account or their paymentRecipient
     */
    function setRecipient(address account, address recipient) external;

    /**
     * @notice Creates a new distribution root
     * @param root The merkle root of the distribution
     * @param paymentsCalculatedUntilTimestamp The timestamp until which payments have been calculated
     * @dev Only callable by the paymentUpdater
     */
    function submitRoot(bytes32 root, uint32 paymentsCalculatedUntilTimestamp) external;

    /**
     * @notice Claims payments for the given claims
     * @param claims The claims to be processed
     */
    function processClaims(PaymentMerkleClaim[] calldata claims) external;
}
