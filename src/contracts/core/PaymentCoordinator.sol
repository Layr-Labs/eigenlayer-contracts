// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "src/contracts/libraries/Merkle.sol";
import "src/contracts/libraries/EIP1271SignatureUtils.sol";
import "src/contracts/permissions/Pausable.sol";
import "src/contracts/core/PaymentCoordinatorStorage.sol";

/**
 * @title PaymentCoordinator
 * @author Eigen Labs Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice  This is the contract for payments in EigenLayer. The main functionalities of this contract are
 * - enabling any ERC20 payments from AVSs to their operators and stakers for a given time range
 * - allowing stakers and operators to claim their earnings including a commission bips for operators
 * - allowing the protocol to provide ERC20 tokens to stakers over a specified time range
 */
contract PaymentCoordinator is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    ReentrancyGuardUpgradeable,
    PaymentCoordinatorStorage
{
    using SafeERC20 for IERC20;

    /// @dev Index for flag that pauses payForRange payments
    uint8 internal constant PAUSED_PAY_FOR_RANGE = 0;

    /// @dev Index for flag that pauses payAllForRange payments
    uint8 internal constant PAUSED_PAY_ALL_FOR_RANGE = 1;

    /// @dev Index for flag that pauses
    uint8 internal constant PAUSED_CLAIM_PAYMENTS = 2;

    /// @dev Chain ID at the time of contract deployment
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    modifier onlyPaymentUpdater() {
        require(msg.sender == paymentUpdater, "PaymentCoordinator: caller is not the paymentUpdater");
        _;
    }

    modifier onlyPayAllForRangeSubmitter() {
        require(
            isPayAllForRangeSubmitter[msg.sender],
            "PaymentCoordinator: caller is not a valid payAllForRange submitter"
        );
        _;
    }

    /// @dev Sets the immutable variables for the contract
    constructor(
        IDelegationManager _delegationManager,
        IStrategyManager _strategyManager,
        uint64 _maxPaymentDuration,
        uint64 _LOWER_BOUND_START_RANGE,
        uint64 _UPPER_BOUND_START_RANGE
    )
        PaymentCoordinatorStorage(
            _delegationManager,
            _strategyManager,
            _maxPaymentDuration,
            _LOWER_BOUND_START_RANGE,
            _UPPER_BOUND_START_RANGE
        )
    {
        _disableInitializers();
        ORIGINAL_CHAIN_ID = block.chainid;
    }

    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, paymentUpdater and
     * configures the initial paused status, activationDelay, and globalOperatorCommissionBips.
     */
    function initialize(
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 initialPausedStatus,
        address _paymentUpdater,
        uint64 _activationDelay,
        uint64 _calculationIntervalSeconds,
        uint16 _globalCommissionBips
    ) external initializer {
        _DOMAIN_SEPARATOR = _calculateDomainSeparator();
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _transferOwnership(initialOwner);
        _setPaymentUpdater(_paymentUpdater);
        _setActivationDelay(_activationDelay);
        _setCalculationIntervalSeconds(_calculationIntervalSeconds);
        _setGlobalOperatorCommission(_globalCommissionBips);
    }

    /**
     * @notice Creates a new range payment on behalf of an AVS, to be split amongst the
     * set of stakers delegated to operators who are registered to the AVS.
     * @param rangePayments The range payments being created
     */
    function payForRange(
        RangePayment[] calldata rangePayments
    ) external onlyWhenNotPaused(PAUSED_PAY_FOR_RANGE) nonReentrant {
        for (uint256 i = 0; i < rangePayments.length; i++) {
            RangePayment calldata rangePayment = rangePayments[i];
            bytes32 rangePaymentHash = _payForRange(rangePayment);
            emit RangePaymentCreated(msg.sender, paymentNonce, rangePaymentHash, rangePayment);
        }
    }

    /**
     * @notice similar to `payForRange` except the payment is split amongst *all* stakers
     * rather than just those delegated to operators who are registered to a single avs
     */
    function payAllForRange(
        RangePayment[] calldata rangePayments
    ) external onlyWhenNotPaused(PAUSED_PAY_ALL_FOR_RANGE) onlyPayAllForRangeSubmitter {
        for (uint256 i = 0; i < rangePayments.length; i++) {
            RangePayment calldata rangePayment = rangePayments[i];
            bytes32 rangePaymentHash = _payForRange(rangePayment);
            emit RangePaymentForAllCreated(msg.sender, paymentNonce, rangePaymentHash, rangePayment);
        }
    }

    /**
     * @notice Claim payments for the given claim
     * @param claim The claim to be processed
     */
    function processClaim(PaymentMerkleClaim calldata claim) external onlyWhenNotPaused(PAUSED_CLAIM_PAYMENTS) {
        _processClaim(claim);
    }

    function _payForRange(RangePayment calldata rangePayment) internal returns (bytes32) {
        bytes32 rangePaymentHash = keccak256(abi.encode(msg.sender, paymentNonce, rangePayment));
        require(
            !isRangePaymentHash[msg.sender][rangePaymentHash],
            "PaymentCoordinator._payForRange: range payment hash already submitted"
        );
        require(
            rangePayment.strategiesAndMultlipliers.length > 0,
            "PaymentCoordinator._payForRange: no strategies set"
        );
        require(rangePayment.amount > 0, "PaymentCoordinator._payForRange: amount cannot be 0");
        require(
            rangePayment.duration <= MAX_PAYMENT_DURATION,
            "PaymentCoordinator._payForRange: duration exceeds MAX_PAYMENT_DURATION"
        );
        require(
            rangePayment.duration % calculationIntervalSeconds == 0,
            "PaymentCoordinator._payForRange: duration must be a multiple of calculationIntervalSeconds"
        );
        if (retroactivePaymentsEnabled) {
            require(
                block.timestamp - LOWER_BOUND_START_RANGE <= rangePayment.startTimestamp,
                "PaymentCoordinator._payForRange: startTimestamp too far in the past"
            );
        } else {
            require(
                block.timestamp <= rangePayment.startTimestamp,
                "PaymentCoordinator._payForRange: startTimestamp cannot be in the past"
            );
        }
        require(
            rangePayment.startTimestamp <= block.timestamp + UPPER_BOUND_START_RANGE,
            "PaymentCoordinator._payForRange: startTimestamp too far in the future"
        );

        // Require rangePayment is for whitelisted strategy or beaconChainETHStrategy
        for (uint256 j = 0; j < rangePayment.strategiesAndMultlipliers.length; j++) {
            IStrategy strategy = rangePayment.strategiesAndMultlipliers[j].strategy;
            require(
                strategyManager.strategyIsWhitelistedForDeposit(strategy) ||
                    strategy == delegationManager.beaconChainETHStrategy(),
                "PaymentCoordinator._payForRange: invalid strategy considered"
            );
        }

        // Set hash of rangePayment in mapping
        isRangePaymentHash[msg.sender][rangePaymentHash] = true;
        paymentNonce++;

        rangePayment.token.safeTransferFrom(msg.sender, address(this), rangePayment.amount);
        return rangePaymentHash;
    }

    /**
     * @notice Process a payment claim.
     * @param claim PaymentMerkleClaim struct containing the claim details for proof verification
     * and payment transfer.
     * @dev only callable by claimerFor[claim.earner]
     */
    function _processClaim(PaymentMerkleClaim calldata claim) internal onlyWhenNotPaused(PAUSED_CLAIM_PAYMENTS) {
        DistributionRoot memory root = distributionRoots[claim.rootIndex];
        require(
            block.timestamp > root.activatedAt + activationDelay,
            "PaymentCoordinator._processClaim: root not activated yet"
        );

        // If claimerFor earner is not set, claimer is default the earner. Else set to claimerFor
        address claimer = claimerFor[claim.earner] == address(0) ? claim.earner : claimerFor[claim.earner];
        require(msg.sender == claimer, "PaymentCoordinator._processClaim: caller is not valid claimer");
        require(
            claim.leafIndices.length == claim.tokenTreeProofs.length,
            "PaymentCoordinator._processClaim: leafIndices and tokenProofs length mismatch"
        );
        require(
            claim.tokenTreeProofs.length == claim.leaves.length,
            "PaymentCoordinator._processClaim: tokenTreeProofs and leaves length mismatch"
        );

        // Verify inclusion of earners leaf (earner, earnerTokenRoot) in the distribution root
        _verifyEarnerClaimProof({
            root: root.root,
            earnerLeafIndex: claim.earnerIndex,
            earnerProof: claim.earnerTreeProof,
            earner: claim.earner,
            earnerTokenRoot: claim.earnerTokenRoot
        });
        // For each of the tokenLeaf proofs, verify inclusion of token tree leaf again the earnerTokenRoot
        for (uint256 i = 0; i < claim.leafIndices.length; ++i) {
            _processTokenClaim({
                earnerTokenRoot: claim.earnerTokenRoot,
                tokenLeafIndex: claim.leafIndices[i],
                tokenProof: claim.tokenTreeProofs[i],
                tokenLeaf: claim.leaves[i],
                earner: claim.earner,
                claimer: claimer,
                root: root.root
            });
        }
    }

    /**
     * @notice Verify a token claim and transfer claimable amount to the claimer
     * @param earnerTokenRoot root hash of the earner token subtree
     * @param tokenLeafIndex index of the token leaf
     * @param tokenProof proof of the token leaf in the earner token subtree
     * @param tokenLeaf token leaf to be verified
     * @param earner address of earner account
     * @param claimer address of the entity that can claim payments on behalf of the earner,
     * can be earner account itself or be set to a different address by the earner
     */
    function _processTokenClaim(
        bytes32 earnerTokenRoot,
        uint32 tokenLeafIndex,
        bytes calldata tokenProof,
        ClaimsTreeMerkleLeaf calldata tokenLeaf,
        address earner,
        address claimer,
        bytes32 root
    ) internal {
        // Verify token proof
        _verifyTokenClaimProof({
            earnerTokenRoot: earnerTokenRoot,
            tokenLeafIndex: tokenLeafIndex,
            tokenProof: tokenProof,
            tokenLeaf: tokenLeaf
        });

        // Calculate amount to claim and update cumulativeClaimed
        // Will revert if new leaf cumulativeEarnings is less than cumulative claimed
        uint256 claimAmount = tokenLeaf.cumulativeEarnings - cumulativeClaimed[earner][tokenLeaf.token];
        cumulativeClaimed[earner][tokenLeaf.token] = tokenLeaf.cumulativeEarnings;

        tokenLeaf.token.safeTransfer(claimer, claimAmount);
        emit PaymentClaimed(root, earnerTokenRoot, tokenLeaf);
    }

    /**
     * @notice verify inclusion of the token claim proof in the earner token root hash (earnerTokenRoot).
     * The token leaf comprises of the IERC20 token and cumulativeAmount of earnings.
     * @param earnerTokenRoot root hash of the earner token subtree
     * @param tokenLeafIndex index of the token leaf
     * @param tokenProof proof of the token leaf in the earner token subtree
     * @param tokenLeaf token leaf to be verified
     */
    function _verifyTokenClaimProof(
        bytes32 earnerTokenRoot,
        uint32 tokenLeafIndex,
        bytes calldata tokenProof,
        ClaimsTreeMerkleLeaf calldata tokenLeaf
    ) internal view {
        // Verify inclusion of token leaf
        bytes32 tokenLeafHash = sha256(abi.encode(tokenLeaf));
        require(
            Merkle.verifyInclusionSha256({
                root: earnerTokenRoot,
                index: tokenLeafIndex,
                proof: tokenProof,
                leaf: tokenLeafHash
            }),
            "PaymentCoordinator._verifyTokenClaim: invalid token claim proof"
        );
    }

    /**
     * @notice verify inclusion of earner claim proof in the distribution root. This verifies
     * the inclusion of the earner and earnerTokenRoot hash in the tree. The token claims are proven separately
     * against the earnerTokenRoot hash (see _verifyTokenClaimProof). The earner leaf comprises of (earner, earnerTokenRoot)
     * @param root distribution root that should be read from storage
     * @param earnerLeafIndex index of the earner leaf
     * @param earnerProof proof of the earners account root in the merkle tree
     * @param earner address of earner account.
     * @param earnerTokenRoot root hash of the token subtree for the earner.
     */
    function _verifyEarnerClaimProof(
        bytes32 root,
        uint32 earnerLeafIndex,
        bytes calldata earnerProof,
        address earner,
        bytes32 earnerTokenRoot
    ) internal view {
        // Verify inclusion of earner leaf
        bytes32 earnerLeafHash = sha256(abi.encodePacked(earner, earnerTokenRoot));
        require(
            Merkle.verifyInclusionSha256({
                root: root,
                index: earnerLeafIndex,
                proof: earnerProof,
                leaf: earnerLeafHash
            }),
            "PaymentCoordinator._verifyEarnerClaimProof: invalid earner claim proof"
        );
    }

    /**
     * @notice Sets the address of the entity that can claim payments on behalf of the earner
     * @param earner The earner whose claimer is being set
     * @param claimer The address of the entity that can claim payments on behalf of the earner
     * @dev Only callable by the `earner`
     */
    function setClaimer(address earner, address claimer) external {
        require(msg.sender == earner, "PaymentCoordinator.setClaimer: caller is not the earner");
        address prevClaimer = claimerFor[earner];
        claimerFor[earner] = claimer;
        emit ClaimerForSet(earner, prevClaimer, claimer);
    }

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
    ) external onlyPaymentUpdater {
        uint32 rootIndex = uint32(distributionRoots.length);
        distributionRoots.push(
            DistributionRoot({
                root: root,
                activatedAt: activatedAt,
                paymentCalculationStartTimestamp: paymentCalculationStartTimestamp,
                paymentCalculationEndTimestamp: paymentCalculationEndTimestamp
            })
        );
        emit DistributionRootSubmitted(
            rootIndex,
            root,
            paymentCalculationStartTimestamp,
            paymentCalculationEndTimestamp,
            activatedAt
        );
    }

    /// @notice returns the hash of the leaf
    function calculateLeafHash(ClaimsTreeMerkleLeaf calldata leaf) external view returns (bytes32) {}

    /// @notice returns 'true' if the claim would currently pass the check in `processClaims`
    function checkClaim(PaymentMerkleClaim calldata claim) external view returns (bool) {}

    /**
     * @notice Set a new value for calculationIntervalSeconds. Only callable by owner
     * Payment durations must be multiples of this interval
     * @param _calculationIntervalSeconds The new value for calculationIntervalSeconds
     */
    function setCalculationIntervalSeconds(uint64 _calculationIntervalSeconds) external onlyOwner {
        _setCalculationIntervalSeconds(_calculationIntervalSeconds);
    }

    /**
     * @notice Set a new value for retroactivePaymentsEnabled. Only callable by owner
     * @param _retroactivePaymentsEnabled The new value for retroactivePaymentsEnabled
     */
    function setRetroactivePaymentsEnabled(bool _retroactivePaymentsEnabled) external onlyOwner {
        retroactivePaymentsEnabled = _retroactivePaymentsEnabled;
    }

    /**
     * @notice Sets the delay in timestamp before a posted root can be claimed against
     * @dev Only callable by the contract owner
     * @param _activationDelay The new value for activationDelay
     */
    function setActivationDelay(uint64 _activationDelay) external onlyOwner {
        _setActivationDelay(_activationDelay);
    }

    /**
     * @notice Sets the global commission for all operators across all avss
     * @dev Only callable by the contract owner
     * @param _globalCommissionBips The commission for all operators across all avss
     */
    function setGlobalOperatorCommission(uint16 _globalCommissionBips) external onlyOwner {
        _setGlobalOperatorCommission(_globalCommissionBips);
    }

    /**
     * @notice Sets the permissioned `paymentUpdater` address which can post new roots
     * @dev Only callable by the contract owner
     * @param _paymentUpdater The address of the new paymentUpdater
     */
    function setPaymentUpdater(address _paymentUpdater) external onlyOwner {
        _setPaymentUpdater(_paymentUpdater);
    }

    function setPayAllForRangeSubmitter(address _submitter, bool _newValue) external onlyOwner {
        bool prevValue = isPayAllForRangeSubmitter[_submitter];
        emit PayAllForRangeSubmitterSet(_submitter, prevValue, _newValue);
        isPayAllForRangeSubmitter[_submitter] = _newValue;
    }

    function _setActivationDelay(uint64 _activationDelay) internal {
        emit ActivationDelaySet(activationDelay, _activationDelay);
        activationDelay = _activationDelay;
    }

    function _setCalculationIntervalSeconds(uint64 _calculationIntervalSeconds) internal {
        emit CalculationIntervalSecondsSet(calculationIntervalSeconds, _calculationIntervalSeconds);
        calculationIntervalSeconds = _calculationIntervalSeconds;
    }

    function _setGlobalOperatorCommission(uint16 _globalCommissionBips) internal {
        emit GlobalCommissionBipsSet(globalOperatorCommissionBips, _globalCommissionBips);
        globalOperatorCommissionBips = _globalCommissionBips;
    }

    function _setPaymentUpdater(address _paymentUpdater) internal {
        emit PaymentUpdaterSet(paymentUpdater, _paymentUpdater);
        paymentUpdater = _paymentUpdater;
    }

    /**
     * @notice Getter function for the current EIP-712 domain separator for this contract.
     *
     * @dev The domain separator will change in the event of a fork that changes the ChainID.
     * @dev By introducing a domain separator the DApp developers are guaranteed that there can be no signature collision.
     * for more detailed information please read EIP-712.
     */
    function domainSeparator() public view returns (bytes32) {
        if (block.chainid == ORIGINAL_CHAIN_ID) {
            return _DOMAIN_SEPARATOR;
        } else {
            return _calculateDomainSeparator();
        }
    }

    /**
     * @dev Recalculates the domain separator when the chainid changes due to a fork.
     */
    function _calculateDomainSeparator() internal view returns (bytes32) {
        return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("EigenLayer")), block.chainid, address(this)));
    }
}
