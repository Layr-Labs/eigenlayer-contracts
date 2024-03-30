// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IStrategy.sol";
import "./IPauserRegistry.sol";

/**
 * @title Interface for the `IPaymentCoordinator` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice Allows AVSs to make "Range Payments", which get distributed amongst the AVSs' confirmed
 * Operators and the Stakers delegated to those Operators.
 * Calculations are performed based on the completed Range Payments, with the results posted in
 * a Merkle root against which Stakers & Operators can make claims.
 */
interface IPaymentCoordinator {
    /// STRUCTS ///
    struct StrategyAndMultiplier {
        IStrategy strategy;
        // weight used to compare shares in multiple strategies against one another
        uint96 multiplier;
    }

    struct RangePayment {
        // Strategies & relative weights of shares in the strategies
        StrategyAndMultiplier[] strategiesAndMultlipliers;
        IERC20 token;
        uint256 amount;
        uint64 startTimestamp;
        uint64 duration;
    }

    struct DistributionRoot {
        // merkle root of the distribution
        bytes32 root;
        // The timestamp which calculated payments started
        uint64 paymentCalculationStartTimestamp;
        // The timestamp until which payments have been calculated
        uint64 paymentCalculationEndTimestamp;
        // timestamp at which the root can be claimed against
        uint64 activatedAt;
    }

    struct ClaimsTreeMerkleLeaf {
        IERC20 token;
        uint256 cumulativeEarnings;
    }

    struct PaymentMerkleClaim {
        // The index of the root in the list of roots
        uint32 rootIndex;
        address earner;
        // proof of the earner's account root in the Merkle tree
        uint32 earnerIndex;
        bytes earnerTreeProof;
        // The indices and proofs of the leafs in the claimaint's merkle tree for this root
        bytes32 earnerTokenRoot;
        uint32[] leafIndices;
        bytes[] tokenTreeProofs;
        ClaimsTreeMerkleLeaf[] leaves;
    }

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

    /// VIEW FUNCTIONS ///

    /// @notice The address of the entity that can update the contract with new merkle roots
    function paymentUpdater() external view returns (address);

    /**
     * @notice The interval in seconds at which the calculation for range payment distribution is done.
     * @dev Payment durations must be multiples of this interval.
     */
    function calculationIntervalSeconds() external view returns (uint64);

    /// @notice The maximum amount of time that a range payment can end in the future
    function MAX_PAYMENT_DURATION() external view returns (uint64);

    /// @notice The lower bound for the start of a range payment
    function LOWER_BOUND_START_RANGE() external view returns (uint64);

    /// @notice Delay in timestamp before a posted root can be claimed against
    function activationDelay() external view returns (uint64);

    /// @notice Mapping: earner => the address of the entity to which new payments are directed on behalf of the earner
    function claimerFor(address earner) external view returns (address);

    /// @notice Mapping: claimer => token => total amount claimed
    function cumulativeClaimed(address claimer, IERC20 token) external view returns (uint256);

    /// @notice the commission for all operators across all avss
    function globalOperatorCommissionBips() external view returns (uint16);

    /// @notice returns the hash of the leaf
    function calculateLeafHash(ClaimsTreeMerkleLeaf calldata leaf) external view returns (bytes32);

    /// @notice returns 'true' if the claim would currently pass the check in `processClaims`
    function checkClaim(PaymentMerkleClaim calldata claim) external view returns (bool);

    /// EXTERNAL FUNCTIONS ///

    /**
     * @notice Creates a new range payment on behalf of an AVS, to be split amongst the
     * set of stakers delegated to operators who are registered to the `avs`
     * @param rangePayments The range payments being created
     * @dev Expected to be called by the ServiceManager of the AVS on behalf of which the payment is being made
     * @dev The duration of the `rangePayment` cannot exceed `MAX_PAYMENT_DURATION`
     * @dev The tokens are sent to the `claimingManager` contract
     * @dev This function will revert if the `rangePayment` is malformed,
     * e.g. if the `strategies` and `weights` arrays are of non-equal lengths
     */
    function payForRange(RangePayment[] calldata rangePayments) external;

    /**
     * @notice similar to `payForRange` except the payment is split amongst *all* stakers
     * rather than just those delegated to operators who are registered to a single avs
     */
    function payAllForRange(RangePayment[] calldata rangePayment) external;

    /**
     * @notice Sets the permissioned `paymentUpdater` address which can post new roots
     * @dev Only callable by the contract owner
     */
    function setPaymentUpdater(address _paymentUpdater) external;

    /**
     * @notice Sets the delay in timestamp before a posted root can be claimed against
     * @param _activationDelay Delay in timestamp before a posted root can be claimed against
     * @dev Only callable by the contract owner
     */
    function setActivationDelay(uint64 _activationDelay) external;

    /**
     * @notice Sets the global commission for all operators across all avss
     * @param _globalCommissionBips The commission for all operators across all avss
     * @dev Only callable by the contract owner
     */
    function setGlobalOperatorCommission(uint16 _globalCommissionBips) external;

    /**
     * @notice Sets the address of the entity that can claim payments on behalf of the earner
     * @param earner The earner whose claimer is being set
     * @param claimer The address of the entity that can claim payments on behalf of the earner
     * @dev Only callable by the `earner`
     */
    function setClaimer(address earner, address claimer) external;

    /**
     * @notice Creates a new distribution root
     * @param root The merkle root of the distribution
     * @param paymentCalculationStartTimestamp The start timestamp which payments have been calculated from
     * @param paymentCalculationEndTimestamp The timestamp until which payments have been calculated
     * @param activatedAt timestamp at which that the root can be claimed against
     * @dev Only callable by the paymentUpdater
     */
    function submitRoot(
        bytes32 root,
        uint64 paymentCalculationStartTimestamp,
        uint64 paymentCalculationEndTimestamp,
        uint64 activatedAt
    ) external;

    /**
     * @notice Claim payments against a given root (read from distributionRoots[claim.rootIndex])
     * @param claim The PaymentMerkleClaim to be processed.
     * Contains the root index, earner, payment leaves, and required proofs
     * @dev only callable by the valid claimer, that is
     * if claimerFor[claim.earner] is address(0) then only the earner can claim, otherwise only
     * claimerFor[claim.earner] can claim the payments.
     */
    function processClaim(PaymentMerkleClaim calldata claim) external;
}
