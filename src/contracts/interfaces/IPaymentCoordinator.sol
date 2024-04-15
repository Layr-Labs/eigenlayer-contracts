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
        StrategyAndMultiplier[] strategiesAndMultipliers;
        IERC20 token;
        uint256 amount;
        uint64 startTimestamp;
        uint64 duration;
    }

    struct DistributionRoot {
        // merkle root of the distribution
        bytes32 root;
        // The timestamp until which payments have been calculated
        uint64 paymentCalculationEndTimestamp;
        // timestamp at which the root can be claimed against
        uint64 activatedAt;
    }

    struct EarnerTreeMerkleLeaf {
        address earner;
        bytes32 earnerTokenRoot;
    }

    struct TokenTreeMerkleLeaf {
        IERC20 token;
        uint256 cumulativeEarnings;
    }

    struct PaymentMerkleClaim {
        // The index of the root in the list of roots
        uint32 rootIndex;
        // proof of the earner's account root in the Earner Merkle tree
        uint32 earnerIndex;
        bytes earnerTreeProof;
        EarnerTreeMerkleLeaf earnerLeaf;
        // The indices and proofs of the leafs in the claimaint's merkle tree for this root
        uint32[] tokenIndices;
        bytes[] tokenTreeProofs;
        TokenTreeMerkleLeaf[] tokenLeaves;
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
        uint64 paymentCalculationEndTimestamp,
        uint64 activatedAt
    );
    /// @notice root is one of the submitted distribution roots that was claimed against
    event PaymentClaimed(bytes32 indexed root, TokenTreeMerkleLeaf leaf);

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

    /// @notice max amount of time that a payment can start in the past
    function MAX_RETROACTIVE_LENGTH() external view returns (uint64);

    /// @notice max amount of time that a payment can start in the future
    function MAX_FUTURE_LENGTH() external view returns (uint64);

    /// @notice absolute min timestamp that a payment can start at
    function GENESIS_PAYMENT_TIMESTAMP() external view returns (uint64);

    /// @notice Delay in timestamp before a posted root can be claimed against
    function activationDelay() external view returns (uint64);

    /// @notice Mapping: earner => the address of the entity to which new payments are directed on behalf of the earner
    function claimerFor(address earner) external view returns (address);

    /// @notice Mapping: claimer => token => total amount claimed
    function cumulativeClaimed(address claimer, IERC20 token) external view returns (uint256);

    /// @notice the commission for all operators across all avss
    function globalOperatorCommissionBips() external view returns (uint16);

    /// @notice return the hash of the earner's leaf
    function calculateEarnerLeafHash(EarnerTreeMerkleLeaf calldata leaf) external pure returns (bytes32);

    /// @notice returns the hash of the earner's token leaf
    function calculateTokenLeafHash(TokenTreeMerkleLeaf calldata leaf) external pure returns (bytes32);

    /// @notice returns 'true' if the claim would currently pass the check in `processClaims`
    function checkClaim(PaymentMerkleClaim calldata claim) external view returns (bool);

    /// @notice Timestamp for last submitted 
    function currPaymentCalculationEndTimestamp() external view returns (uint64);

    /// @notice loop through distribution roots from reverse and return hash
    function getRootIndexFromHash(bytes32 rootHash) external view returns (uint32);

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
     * @notice Claim payments against a given root (read from distributionRoots[claim.rootIndex])
     * @param claim The PaymentMerkleClaim to be processed.
     * Contains the root index, earner, payment leaves, and required proofs
     * @dev only callable by the valid claimer, that is
     * if claimerFor[claim.earner] is address(0) then only the earner can claim, otherwise only
     * claimerFor[claim.earner] can claim the payments.
     */
    function processClaim(PaymentMerkleClaim calldata claim) external;

    /**
     * @notice Creates a new distribution root
     * @param root The merkle root of the distribution
     * @param paymentCalculationEndTimestamp The timestamp until which payments have been calculated
     * @dev Only callable by the paymentUpdater
     */
    function submitRoot(
        bytes32 root,
        uint64 paymentCalculationEndTimestamp
    ) external;

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
     * @notice Sets the address of the entity that can claim payments on behalf of the earner (msg.sender)
     * @param claimer The address of the entity that can claim payments on behalf of the earner
     * @dev Only callable by the `earner`
     */
    function setClaimerFor(address claimer) external;
}
