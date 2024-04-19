// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../libraries/Merkle.sol";
import "../libraries/EIP1271SignatureUtils.sol";
import "../permissions/Pausable.sol";
import "./PaymentCoordinatorStorage.sol";

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
        uint32 _maxPaymentDuration,
        uint32 _MAX_RETROACTIVE_LENGTH,
        uint32 _MAX_FUTURE_LENGTH,
        uint32 _GENESIS_PAYMENT_TIMESTAMP
    )
        PaymentCoordinatorStorage(
            _delegationManager,
            _strategyManager,
            _maxPaymentDuration,
            _MAX_RETROACTIVE_LENGTH,
            _MAX_FUTURE_LENGTH,
            _GENESIS_PAYMENT_TIMESTAMP
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
        uint32 _activationDelay,
        uint32 _calculationIntervalSeconds,
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
     * set of stakers delegated to operators who are registered to the `avs`
     * @param rangePayments The range payments being created
     * @dev Expected to be called by the ServiceManager of the AVS on behalf of which the payment is being made
     * @dev The duration of the `rangePayment` cannot exceed `MAX_PAYMENT_DURATION`
     * @dev The tokens are sent to the `PaymentCoordinator` contract
     * @dev Strategies must be in ascending order of addresses to check for duplicates
     * @dev This function will revert if the `rangePayment` is malformed,
     * e.g. if the `strategies` and `weights` arrays are of non-equal lengths
     */
    function payForRange(
        RangePayment[] calldata rangePayments
    ) external onlyWhenNotPaused(PAUSED_PAY_FOR_RANGE) nonReentrant {
        for (uint256 i = 0; i < rangePayments.length; i++) {
            RangePayment calldata rangePayment = rangePayments[i];
            uint256 nonce = paymentNonce[msg.sender];
            bytes32 rangePaymentHash = keccak256(abi.encode(msg.sender, nonce, rangePayment));
            require(
                !isRangePaymentHash[msg.sender][rangePaymentHash],
                "PaymentCoordinator._payForRange: range payment hash already submitted"
            );
            isRangePaymentHash[msg.sender][rangePaymentHash] = true;

            _payForRange(rangePayment);
            emit RangePaymentCreated(msg.sender, nonce, rangePaymentHash, rangePayment);
            paymentNonce[msg.sender] = nonce + 1;
        }
    }

    /**
     * @notice similar to `payForRange` except the payment is split amongst *all* stakers
     * rather than just those delegated to operators who are registered to a single avs and is
     * a permissioned call based on isPayAllForRangeSubmitter mapping.
     * @param rangePayments The range payments being created
     */
    function payAllForRange(
        RangePayment[] calldata rangePayments
    ) external onlyWhenNotPaused(PAUSED_PAY_ALL_FOR_RANGE) onlyPayAllForRangeSubmitter nonReentrant {
        for (uint256 i = 0; i < rangePayments.length; i++) {
            RangePayment calldata rangePayment = rangePayments[i];
            uint256 nonce = paymentNonce[msg.sender];
            bytes32 rangePaymentForAllHash = keccak256(abi.encode(msg.sender, nonce, rangePayment));
            require(
                !isRangePaymentForAllHash[msg.sender][rangePaymentForAllHash],
                "PaymentCoordinator._payForRange: range payment hash already submitted"
            );
            isRangePaymentForAllHash[msg.sender][rangePaymentForAllHash] = true;

            _payForRange(rangePayment);
            emit RangePaymentForAllCreated(msg.sender, nonce, rangePaymentForAllHash, rangePayment);
            paymentNonce[msg.sender] = nonce + 1;
        }
    }

    /**
     * @notice Claim payments against a given root (read from distributionRoots[claim.rootIndex]).
     * Earnings are cumulative so earners don't have to claim against all distribution roots they have earnings for,
     * they can simply claim against the latest root and the contract will calculate the difference between
     * their cumulativeEarnings and cumulativeClaimed. This difference is then transferred to claimerFor[claim.earner]
     * @param claim The PaymentMerkleClaim to be processed.
     * Contains the root index, earner, payment leaves, and required proofs
     * @dev only callable by the valid claimer, that is
     * if claimerFor[claim.earner] is address(0) then only the earner can claim, otherwise only
     * claimerFor[claim.earner] can claim the payments.
     */
    function processClaim(PaymentMerkleClaim calldata claim) external onlyWhenNotPaused(PAUSED_CLAIM_PAYMENTS) nonReentrant {
        DistributionRoot memory root = distributionRoots[claim.rootIndex];
        require(_checkClaim(claim, root), "PaymentCoordinator.processClaim: claim does not pass the check");
        // If claimerFor earner is not set, claimer is by default the earner. Else set to claimerFor
        address earner = claim.earnerLeaf.earner;
        address claimer = claimerFor[earner];
        if (claimer == address(0)) {
            claimer = earner;
        }
        require(msg.sender == claimer, "PaymentCoordinator.processClaim: caller is not valid claimer");
        for (uint256 i = 0; i < claim.tokenIndices.length; i++) {
            TokenTreeMerkleLeaf calldata tokenLeaf = claim.tokenLeaves[i];

            // Calculate amount to claim and update cumulativeClaimed
            // Will revert if new leaf cumulativeEarnings is less than cumulative claimed
            uint256 claimAmount = tokenLeaf.cumulativeEarnings - cumulativeClaimed[earner][tokenLeaf.token];
            cumulativeClaimed[earner][tokenLeaf.token] = tokenLeaf.cumulativeEarnings;

            tokenLeaf.token.safeTransfer(claimer, claimAmount);
            emit PaymentClaimed(root.root, earner, claimer, claim.tokenLeaves[i].token, claimAmount);
        }
    }

    /**
     * @notice Creates a new distribution root. activatedAt is set to block.timestamp + activationDelay
     * @param root The merkle root of the distribution
     * @param paymentCalculationEndTimestamp The timestamp until which payments have been calculated
     * @dev Only callable by the paymentUpdater
     */
    function submitRoot(
        bytes32 root,
        uint32 paymentCalculationEndTimestamp
    ) external onlyPaymentUpdater {
        require(
            paymentCalculationEndTimestamp > currPaymentCalculationEndTimestamp,
            "PaymentCoordinator.submitRoot: new root must be for newer calculated period"
        );
        uint32 rootIndex = uint32(distributionRoots.length);
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;
        distributionRoots.push(
            DistributionRoot({
                root: root,
                activatedAt: activatedAt,
                paymentCalculationEndTimestamp: paymentCalculationEndTimestamp
            })
        );
        currPaymentCalculationEndTimestamp = paymentCalculationEndTimestamp;
        emit DistributionRootSubmitted(rootIndex, root, paymentCalculationEndTimestamp, activatedAt);
    }

    /**
     * @notice Create a RangePayment. Called from both `payForRange` and `payAllForRange`
     */
    function _payForRange(RangePayment calldata rangePayment) internal {
        require(rangePayment.strategiesAndMultipliers.length > 0, "PaymentCoordinator._payForRange: no strategies set");
        require(rangePayment.amount > 0, "PaymentCoordinator._payForRange: amount cannot be 0");
        require(
            rangePayment.duration <= MAX_PAYMENT_DURATION,
            "PaymentCoordinator._payForRange: duration exceeds MAX_PAYMENT_DURATION"
        );
        require(
            rangePayment.duration % calculationIntervalSeconds == 0,
            "PaymentCoordinator._payForRange: duration must be a multiple of calculationIntervalSeconds"
        );
        require(
            rangePayment.startTimestamp % calculationIntervalSeconds == 0,
            "PaymentCoordinator._payForRange: startTimestamp must be a multiple of calculationIntervalSeconds"
        );
        require(
            block.timestamp - MAX_RETROACTIVE_LENGTH <= rangePayment.startTimestamp &&
                GENESIS_PAYMENT_TIMESTAMP <= rangePayment.startTimestamp,
            "PaymentCoordinator._payForRange: startTimestamp too far in the past"
        );
        require(
            rangePayment.startTimestamp <= block.timestamp + MAX_FUTURE_LENGTH,
            "PaymentCoordinator._payForRange: startTimestamp too far in the future"
        );

        // Require rangePayment is for whitelisted strategy or beaconChainETHStrategy
        address currAddress = address(0);
        for (uint256 i = 0; i < rangePayment.strategiesAndMultipliers.length; ++i) {
            IStrategy strategy = rangePayment.strategiesAndMultipliers[i].strategy;
            require(
                strategyManager.strategyIsWhitelistedForDeposit(strategy) ||
                    strategy == beaconChainETHStrategy,
                "PaymentCoordinator._payForRange: invalid strategy considered"
            );
            require(
                currAddress < address(strategy),
                "PaymentCoordinator._payForRange: strategies must be in ascending order to handle duplicates"
            );
            currAddress = address(strategy);
        }

        rangePayment.token.safeTransferFrom(msg.sender, address(this), rangePayment.amount);
    }

    function _checkClaim(PaymentMerkleClaim calldata claim, DistributionRoot memory root) internal view returns (bool) {
        require(block.timestamp >= root.activatedAt, "PaymentCoordinator._checkClaim: root not activated yet");
        require(
            claim.tokenIndices.length == claim.tokenTreeProofs.length,
            "PaymentCoordinator._checkClaim: tokenIndices and tokenProofs length mismatch"
        );
        require(
            claim.tokenTreeProofs.length == claim.tokenLeaves.length,
            "PaymentCoordinator._checkClaim: tokenTreeProofs and leaves length mismatch"
        );

        // Verify inclusion of earners leaf (earner, earnerTokenRoot) in the distribution root
        _verifyEarnerClaimProof({
            root: root.root,
            earnerLeafIndex: claim.earnerIndex,
            earnerProof: claim.earnerTreeProof,
            earnerLeaf: claim.earnerLeaf
        });
        // For each of the tokenLeaf proofs, verify inclusion of token tree leaf again the earnerTokenRoot
        for (uint256 i = 0; i < claim.tokenIndices.length; ++i) {
            _verifyTokenClaimProof({
                earnerTokenRoot: claim.earnerLeaf.earnerTokenRoot,
                tokenLeafIndex: claim.tokenIndices[i],
                tokenProof: claim.tokenTreeProofs[i],
                tokenLeaf: claim.tokenLeaves[i]
            });
        }

        return true;
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
        TokenTreeMerkleLeaf calldata tokenLeaf
    ) internal pure {
        // Verify inclusion of token leaf
        bytes32 tokenLeafHash = calculateTokenLeafHash(tokenLeaf);
        require(
            Merkle.verifyInclusionKeccak({
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
     * @param earnerLeaf leaf of earner merkle tree containing the earner address and earner's token root hash
     */
    function _verifyEarnerClaimProof(
        bytes32 root,
        uint32 earnerLeafIndex,
        bytes calldata earnerProof,
        EarnerTreeMerkleLeaf calldata earnerLeaf
    ) internal pure {
        // Verify inclusion of earner leaf
        bytes32 earnerLeafHash = calculateEarnerLeafHash(earnerLeaf);
        require(
            Merkle.verifyInclusionKeccak({
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
     * @param claimer The address of the entity that can claim payments on behalf of the earner
     * @dev Only callable by the `earner`
     */
    function setClaimerFor(address claimer) external {
        address earner = msg.sender;
        address prevClaimer = claimerFor[earner];
        claimerFor[earner] = claimer;
        emit ClaimerForSet(earner, prevClaimer, claimer);
    }

    /// @notice return the hash of the earner's leaf
    function calculateEarnerLeafHash(EarnerTreeMerkleLeaf calldata leaf) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(leaf.earner, leaf.earnerTokenRoot));
    }

    /// @notice returns the hash of the earner's token leaf
    function calculateTokenLeafHash(TokenTreeMerkleLeaf calldata leaf) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(leaf.token, leaf.cumulativeEarnings));
    }

    /**
     * @notice Set a new value for calculationIntervalSeconds. Only callable by owner
     * Payment durations must be multiples of this interval
     * @param _calculationIntervalSeconds The new value for calculationIntervalSeconds
     */
    function setCalculationIntervalSeconds(uint32 _calculationIntervalSeconds) external onlyOwner {
        _setCalculationIntervalSeconds(_calculationIntervalSeconds);
    }

    /**
     * @notice Sets the delay in timestamp before a posted root can be claimed against
     * @dev Only callable by the contract owner
     * @param _activationDelay The new value for activationDelay
     */
    function setActivationDelay(uint32 _activationDelay) external onlyOwner {
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

    function _setActivationDelay(uint32 _activationDelay) internal {
        emit ActivationDelaySet(activationDelay, _activationDelay);
        activationDelay = _activationDelay;
    }

    function _setCalculationIntervalSeconds(uint32 _calculationIntervalSeconds) internal {
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

    /// @notice returns 'true' if the claim would currently pass the check in `processClaims`
    /// but will revert if not valid
    function checkClaim(PaymentMerkleClaim calldata claim) public view returns (bool) {
        return _checkClaim(claim, distributionRoots[claim.rootIndex]);
    }

    function getDistributionRootsLength() public view returns (uint256) {
        return distributionRoots.length;
    }

    /// @notice loop through distribution roots from reverse and return hash
    function getRootIndexFromHash(bytes32 rootHash) public view returns (uint32) {
        for (uint32 i = uint32(distributionRoots.length); i > 0; i--) {
            if (distributionRoots[i - 1].root == rootHash) {
                return i - 1;
            }
        }
        revert("PaymentCoordinator.getRootIndexFromHash: root not found");
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
