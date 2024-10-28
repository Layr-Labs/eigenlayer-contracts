// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../libraries/Merkle.sol";
import "../permissions/Pausable.sol";
import "./RewardsCoordinatorStorage.sol";

/**
 * @title RewardsCoordinator
 * @author Eigen Labs Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice  This is the contract for rewards in EigenLayer. The main functionalities of this contract are
 * - enabling any ERC20 rewards from AVSs to their operators and stakers for a given time range
 * - allowing stakers and operators to claim their earnings including a commission bips for operators
 * - allowing the protocol to provide ERC20 tokens to stakers over a specified time range
 */
contract RewardsCoordinator is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    ReentrancyGuardUpgradeable,
    RewardsCoordinatorStorage
{
    using SafeERC20 for IERC20;

    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 internal constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");
    /// @dev Chain ID at the time of contract deployment
    uint256 internal immutable ORIGINAL_CHAIN_ID;
    /// @notice The maximum rewards token amount for a single rewards submission, constrained by off-chain calculation
    uint256 internal constant MAX_REWARDS_AMOUNT = 1e38 - 1;
    /// @notice Equivalent to 100%, but in basis points.
    uint16 internal constant ONE_HUNDRED_IN_BIPS = 10_000;
    /// @notice Minimum commission for operators for Programmatic Incentives in basis points
    uint16 internal constant MIN_PI_COMMISSION_BIPS = 1_000;

    /// @dev Index for flag that pauses calling createAVSRewardsSubmission
    uint8 internal constant PAUSED_AVS_REWARDS_SUBMISSION = 0;
    /// @dev Index for flag that pauses calling createRewardsForAllSubmission
    uint8 internal constant PAUSED_REWARDS_FOR_ALL_SUBMISSION = 1;
    /// @dev Index for flag that pauses calling processClaim
    uint8 internal constant PAUSED_PROCESS_CLAIM = 2;
    /// @dev Index for flag that pauses submitRoots and disableRoot
    uint8 internal constant PAUSED_SUBMIT_DISABLE_ROOTS = 3;
    /// @dev Index for flag that pauses calling rewardAllStakersAndOperators
    uint8 internal constant PAUSED_REWARD_ALL_STAKERS_AND_OPERATORS = 4;
    /// @dev Index for flag that pauses calling createAVSPerformanceRewardsSubmission
    uint8 internal constant PAUSED_AVS_PERFORMANCE_REWARDS_SUBMISSION = 5;
    /// @dev Index for flag that pauses calling setOperatorAVSCommission
    uint8 internal constant PAUSED_OPERATOR_AVS_COMMISSION = 6;
    /// @dev Index for flag that pauses calling setOperatorPICommission
    uint8 internal constant PAUSED_OPERATOR_PI_COMMISSION = 7;

    /// @dev Salt for the earner leaf, meant to distinguish from tokenLeaf since they have the same sized data
    uint8 internal constant EARNER_LEAF_SALT = 0;
    /// @dev Salt for the token leaf, meant to distinguish from earnerLeaf since they have the same sized data
    uint8 internal constant TOKEN_LEAF_SALT = 1;

    /// @notice Canonical, virtual beacon chain ETH strategy
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    modifier onlyRewardsUpdater() {
        require(msg.sender == rewardsUpdater, "RewardsCoordinator: caller is not the rewardsUpdater");
        _;
    }

    modifier onlyRewardsForAllSubmitter() {
        require(
            isRewardsForAllSubmitter[msg.sender],
            "RewardsCoordinator: caller is not a valid createRewardsForAllSubmission submitter"
        );
        _;
    }

    /// @dev Sets the immutable variables for the contract
    constructor(
        IDelegationManager _delegationManager,
        IStrategyManager _strategyManager,
        uint32 _CALCULATION_INTERVAL_SECONDS,
        uint32 _MAX_REWARDS_DURATION,
        uint32 _MAX_RETROACTIVE_LENGTH,
        uint32 _MAX_FUTURE_LENGTH,
        uint32 __GENESIS_REWARDS_TIMESTAMP
    )
        RewardsCoordinatorStorage(
            _delegationManager,
            _strategyManager,
            _CALCULATION_INTERVAL_SECONDS,
            _MAX_REWARDS_DURATION,
            _MAX_RETROACTIVE_LENGTH,
            _MAX_FUTURE_LENGTH,
            __GENESIS_REWARDS_TIMESTAMP
        )
    {
        _disableInitializers();
        ORIGINAL_CHAIN_ID = block.chainid;
    }

    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, rewardsUpdater and
     * configures the initial paused status, activationDelay, and globalOperatorCommissionBips.
     */
    function initialize(
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 initialPausedStatus,
        address _rewardsUpdater,
        uint32 _activationDelay,
        uint16 _globalCommissionBips
    ) external initializer {
        _DOMAIN_SEPARATOR = _calculateDomainSeparator();
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _transferOwnership(initialOwner);
        _setRewardsUpdater(_rewardsUpdater);
        _setActivationDelay(_activationDelay);
        _setGlobalOperatorCommission(_globalCommissionBips);
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /// @inheritdoc IRewardsCoordinator
    function createAVSRewardsSubmission(
        RewardsSubmission[] calldata rewardsSubmissions
    ) external onlyWhenNotPaused(PAUSED_AVS_REWARDS_SUBMISSION) nonReentrant {
        for (uint256 i = 0; i < rewardsSubmissions.length; i++) {
            RewardsSubmission calldata rewardsSubmission = rewardsSubmissions[i];
            uint256 nonce = submissionNonce[msg.sender];
            bytes32 rewardsSubmissionHash = keccak256(abi.encode(msg.sender, nonce, rewardsSubmission));

            _validateRewardsSubmission(rewardsSubmission);

            isAVSRewardsSubmissionHash[msg.sender][rewardsSubmissionHash] = true;
            submissionNonce[msg.sender] = nonce + 1;

            emit AVSRewardsSubmissionCreated(msg.sender, nonce, rewardsSubmissionHash, rewardsSubmission);
            rewardsSubmission.token.safeTransferFrom(msg.sender, address(this), rewardsSubmission.amount);
        }
    }

    /// @inheritdoc IRewardsCoordinator
    function createRewardsForAllSubmission(
        RewardsSubmission[] calldata rewardsSubmissions
    ) external onlyWhenNotPaused(PAUSED_REWARDS_FOR_ALL_SUBMISSION) onlyRewardsForAllSubmitter nonReentrant {
        for (uint256 i = 0; i < rewardsSubmissions.length; i++) {
            RewardsSubmission calldata rewardsSubmission = rewardsSubmissions[i];
            uint256 nonce = submissionNonce[msg.sender];
            bytes32 rewardsSubmissionForAllHash = keccak256(abi.encode(msg.sender, nonce, rewardsSubmission));

            _validateRewardsSubmission(rewardsSubmission);

            isRewardsSubmissionForAllHash[msg.sender][rewardsSubmissionForAllHash] = true;
            submissionNonce[msg.sender] = nonce + 1;

            emit RewardsSubmissionForAllCreated(msg.sender, nonce, rewardsSubmissionForAllHash, rewardsSubmission);
            rewardsSubmission.token.safeTransferFrom(msg.sender, address(this), rewardsSubmission.amount);
        }
    }

    /// @inheritdoc IRewardsCoordinator
    function createRewardsForAllEarners(
        RewardsSubmission[] calldata rewardsSubmissions
    ) external onlyWhenNotPaused(PAUSED_REWARD_ALL_STAKERS_AND_OPERATORS) onlyRewardsForAllSubmitter nonReentrant {
        for (uint256 i = 0; i < rewardsSubmissions.length; i++) {
            RewardsSubmission calldata rewardsSubmission = rewardsSubmissions[i];
            uint256 nonce = submissionNonce[msg.sender];
            bytes32 rewardsSubmissionForAllEarnersHash = keccak256(abi.encode(msg.sender, nonce, rewardsSubmission));

            _validateRewardsSubmission(rewardsSubmission);

            isRewardsSubmissionForAllEarnersHash[msg.sender][rewardsSubmissionForAllEarnersHash] = true;
            submissionNonce[msg.sender] = nonce + 1;

            emit RewardsSubmissionForAllEarnersCreated(
                msg.sender,
                nonce,
                rewardsSubmissionForAllEarnersHash,
                rewardsSubmission
            );
            rewardsSubmission.token.safeTransferFrom(msg.sender, address(this), rewardsSubmission.amount);
        }
    }

    /// @inheritdoc IRewardsCoordinator
    function createAVSPerformanceRewardsSubmission(
        address avs,
        PerformanceRewardsSubmission[] calldata performanceRewardsSubmissions
    ) external onlyWhenNotPaused(PAUSED_AVS_PERFORMANCE_REWARDS_SUBMISSION) nonReentrant {
        require(msg.sender == avs, "RewardsCoordinator.createAVSPerformanceRewardsSubmission: caller is not the AVS");

        for (uint256 i = 0; i < performanceRewardsSubmissions.length; i++) {
            PerformanceRewardsSubmission calldata performanceRewardsSubmission = performanceRewardsSubmissions[i];
            uint256 nonce = submissionNonce[avs];
            bytes32 performanceRewardsSubmissionHash = keccak256(abi.encode(avs, nonce, performanceRewardsSubmission));

            uint256 totalAmount = _validatePerformanceRewardsSubmission(performanceRewardsSubmission);

            isAVSPerformanceRewardsSubmissionHash[avs][performanceRewardsSubmissionHash] = true;
            submissionNonce[avs] = nonce + 1;

            emit AVSPerformanceRewardsSubmissionCreated(
                msg.sender,
                avs,
                nonce,
                performanceRewardsSubmissionHash,
                performanceRewardsSubmission
            );
            performanceRewardsSubmission.token.safeTransferFrom(msg.sender, address(this), totalAmount);
        }
    }

    /// @inheritdoc IRewardsCoordinator
    function processClaim(
        RewardsMerkleClaim calldata claim,
        address recipient
    ) external onlyWhenNotPaused(PAUSED_PROCESS_CLAIM) nonReentrant {
        _processClaim(claim, recipient);
    }

    /// @inheritdoc IRewardsCoordinator
    function processClaims(
        RewardsMerkleClaim[] calldata claims,
        address recipient
    ) external onlyWhenNotPaused(PAUSED_PROCESS_CLAIM) nonReentrant {
        for (uint256 i = 0; i < claims.length; i++) {
            _processClaim(claims[i], recipient);
        }
    }

    /// @inheritdoc IRewardsCoordinator
    function submitRoot(
        bytes32 root,
        uint32 rewardsCalculationEndTimestamp
    ) external onlyWhenNotPaused(PAUSED_SUBMIT_DISABLE_ROOTS) onlyRewardsUpdater {
        require(
            rewardsCalculationEndTimestamp > currRewardsCalculationEndTimestamp,
            "RewardsCoordinator.submitRoot: new root must be for newer calculated period"
        );
        require(
            rewardsCalculationEndTimestamp < block.timestamp,
            "RewardsCoordinator.submitRoot: rewardsCalculationEndTimestamp cannot be in the future"
        );
        uint32 rootIndex = uint32(_distributionRoots.length);
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;
        _distributionRoots.push(
            DistributionRoot({
                root: root,
                activatedAt: activatedAt,
                rewardsCalculationEndTimestamp: rewardsCalculationEndTimestamp,
                disabled: false
            })
        );
        currRewardsCalculationEndTimestamp = rewardsCalculationEndTimestamp;
        emit DistributionRootSubmitted(rootIndex, root, rewardsCalculationEndTimestamp, activatedAt);
    }

    /// @inheritdoc IRewardsCoordinator
    function disableRoot(uint32 rootIndex) external onlyWhenNotPaused(PAUSED_SUBMIT_DISABLE_ROOTS) onlyRewardsUpdater {
        require(rootIndex < _distributionRoots.length, "RewardsCoordinator.disableRoot: invalid rootIndex");
        DistributionRoot storage root = _distributionRoots[rootIndex];
        require(!root.disabled, "RewardsCoordinator.disableRoot: root already disabled");
        require(block.timestamp < root.activatedAt, "RewardsCoordinator.disableRoot: root already activated");
        root.disabled = true;
        emit DistributionRootDisabled(rootIndex);
    }

    /// @inheritdoc IRewardsCoordinator
    function setClaimerFor(address claimer) external {
        address earner = msg.sender;
        address prevClaimer = claimerFor[earner];
        claimerFor[earner] = claimer;
        emit ClaimerForSet(earner, prevClaimer, claimer);
    }

    /// @inheritdoc IRewardsCoordinator
    function setActivationDelay(uint32 _activationDelay) external onlyOwner {
        _setActivationDelay(_activationDelay);
    }

    /// @inheritdoc IRewardsCoordinator
    function setGlobalOperatorCommission(uint16 _globalCommissionBips) external onlyOwner {
        _setGlobalOperatorCommission(_globalCommissionBips);
    }

    /// @inheritdoc IRewardsCoordinator
    function setOperatorAVSCommission(
        address operator,
        address avs,
        uint16 commission
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_AVS_COMMISSION) {
        require(msg.sender == operator, "RewardsCoordinator.setOperatorAVSCommission: caller is not the operator");
        require(
            commission <= ONE_HUNDRED_IN_BIPS,
            "RewardsCoordinator.setOperatorAVSCommission: commission must be <= 10000 bips"
        );
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;

        OperatorCommission storage operatorCommission = operatorAVSCommissionBips[operator][avs];

        // If, the earlier 'new' commission is activated, we update the 'old' commission with the earlier 'new' commission.
        // Else, the earlier 'old' commission remains the same. This is essentially resetting the activation delay window
        // since the earlier commission setting didn't complete.
        if (block.timestamp >= operatorCommission.activatedAt) {
            operatorCommission.oldCommissionBips = operatorCommission.newCommissionBips;
        }
        operatorCommission.newCommissionBips = commission;
        operatorCommission.activatedAt = activatedAt;

        emit OperatorAVSCommissionBipsSet(
            msg.sender,
            operator,
            avs,
            activatedAt,
            operatorCommission.oldCommissionBips,
            commission
        );
    }

    /// @inheritdoc IRewardsCoordinator
    function setOperatorPICommission(
        address operator,
        uint16 commission
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_PI_COMMISSION) {
        require(msg.sender == operator, "RewardsCoordinator.setOperatorPICommission: caller is not the operator");
        require(
            commission <= ONE_HUNDRED_IN_BIPS,
            "RewardsCoordinator.setOperatorPICommission: commission must be <= 10000 bips"
        );
        require(
            commission >= MIN_PI_COMMISSION_BIPS,
            "RewardsCoordinator.setOperatorPICommission: commission must be >= 1000 bips"
        );
        uint32 activatedAt = uint32(block.timestamp) + activationDelay;

        OperatorCommission storage operatorCommission = operatorPICommissionBips[operator];

        // If, the earlier 'new' commission is activated, we update the 'old' commission with the earlier 'new' commission.
        // Else, the earlier 'old' commission remains the same. This is essentially resetting the activation delay window
        // since the earlier commission setting didn't complete.
        if (block.timestamp >= operatorCommission.activatedAt) {
            operatorCommission.oldCommissionBips = operatorCommission.newCommissionBips;
        }
        operatorCommission.newCommissionBips = commission;
        operatorCommission.activatedAt = activatedAt;

        emit OperatorPICommissionBipsSet(
            msg.sender,
            operator,
            activatedAt,
            operatorCommission.oldCommissionBips,
            commission
        );
    }

    /// @inheritdoc IRewardsCoordinator
    function setRewardsUpdater(address _rewardsUpdater) external onlyOwner {
        _setRewardsUpdater(_rewardsUpdater);
    }

    /// @inheritdoc IRewardsCoordinator
    function setRewardsForAllSubmitter(address _submitter, bool _newValue) external onlyOwner {
        bool prevValue = isRewardsForAllSubmitter[_submitter];
        emit RewardsForAllSubmitterSet(_submitter, prevValue, _newValue);
        isRewardsForAllSubmitter[_submitter] = _newValue;
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Validate a RewardsSubmission. Called from both `createAVSRewardsSubmission` and `createRewardsForAllSubmission`
     */
    function _validateRewardsSubmission(RewardsSubmission calldata rewardsSubmission) internal view {
        require(
            rewardsSubmission.strategiesAndMultipliers.length > 0,
            "RewardsCoordinator._validateRewardsSubmission: no strategies set"
        );
        require(rewardsSubmission.amount > 0, "RewardsCoordinator._validateRewardsSubmission: amount cannot be 0");
        require(
            rewardsSubmission.amount <= MAX_REWARDS_AMOUNT,
            "RewardsCoordinator._validateRewardsSubmission: amount too large"
        );
        require(
            rewardsSubmission.duration <= MAX_REWARDS_DURATION,
            "RewardsCoordinator._validateRewardsSubmission: duration exceeds MAX_REWARDS_DURATION"
        );
        require(
            rewardsSubmission.duration % CALCULATION_INTERVAL_SECONDS == 0,
            "RewardsCoordinator._validateRewardsSubmission: duration must be a multiple of CALCULATION_INTERVAL_SECONDS"
        );
        require(
            rewardsSubmission.startTimestamp % CALCULATION_INTERVAL_SECONDS == 0,
            "RewardsCoordinator._validateRewardsSubmission: startTimestamp must be a multiple of CALCULATION_INTERVAL_SECONDS"
        );
        require(
            block.timestamp - MAX_RETROACTIVE_LENGTH <= rewardsSubmission.startTimestamp &&
                GENESIS_REWARDS_TIMESTAMP <= rewardsSubmission.startTimestamp,
            "RewardsCoordinator._validateRewardsSubmission: startTimestamp too far in the past"
        );
        require(
            rewardsSubmission.startTimestamp <= block.timestamp + MAX_FUTURE_LENGTH,
            "RewardsCoordinator._validateRewardsSubmission: startTimestamp too far in the future"
        );

        // Require rewardsSubmission is for whitelisted strategy or beaconChainETHStrategy
        address currAddress = address(0);
        for (uint256 i = 0; i < rewardsSubmission.strategiesAndMultipliers.length; ++i) {
            IStrategy strategy = rewardsSubmission.strategiesAndMultipliers[i].strategy;
            require(
                strategyManager.strategyIsWhitelistedForDeposit(strategy) || strategy == beaconChainETHStrategy,
                "RewardsCoordinator._validateRewardsSubmission: invalid strategy considered"
            );
            require(
                currAddress < address(strategy),
                "RewardsCoordinator._validateRewardsSubmission: strategies must be in ascending order to handle duplicates"
            );
            currAddress = address(strategy);
        }
    }

    /**
     * @notice Validate a PerformanceRewardsSubmission. Called from `createAVSPerformanceRewardsSubmission`.
     * @dev Not checking for `MAX_FUTURE_LENGTH` (Since Performance based reward submissions are retroactive)
     * or `MAX_REWARDS_AMOUNT` (Since we no longer have the `1e38 - 1` limitation in the offchain rewards calculation)
     * @param performanceRewardsSubmission PerformanceRewardsSubmission to validate.
     * @return total amount to be transferred from the avs to the contract.
     */
    function _validatePerformanceRewardsSubmission(
        PerformanceRewardsSubmission calldata performanceRewardsSubmission
    ) internal view returns (uint256) {
        require(
            performanceRewardsSubmission.strategiesAndMultipliers.length > 0,
            "RewardsCoordinator._validatePerformanceRewardsSubmission: no strategies set"
        );
        require(
            performanceRewardsSubmission.operatorRewards.length > 0,
            "RewardsCoordinator._validatePerformanceRewardsSubmission: no operators rewarded"
        );

        uint256 totalAmount = 0;
        address currOperatorAddress = address(0);
        for (uint256 i = 0; i < performanceRewardsSubmission.operatorRewards.length; ++i) {
            OperatorReward calldata operatorReward = performanceRewardsSubmission.operatorRewards[i];
            require(
                operatorReward.operator != address(0),
                "RewardsCoordinator._validatePerformanceRewardsSubmission: operator cannot be 0 address"
            );
            require(
                currOperatorAddress < operatorReward.operator,
                "RewardsCoordinator._validatePerformanceRewardsSubmission: operators must be in ascending order to handle duplicates"
            );
            currOperatorAddress = operatorReward.operator;
            require(
                operatorReward.amount > 0,
                "RewardsCoordinator._validatePerformanceRewardsSubmission: operator reward amount cannot be 0"
            );
            totalAmount += operatorReward.amount;
        }

        require(
            performanceRewardsSubmission.duration <= MAX_REWARDS_DURATION,
            "RewardsCoordinator._validatePerformanceRewardsSubmission: duration exceeds MAX_REWARDS_DURATION"
        );
        require(
            performanceRewardsSubmission.duration % CALCULATION_INTERVAL_SECONDS == 0,
            "RewardsCoordinator._validatePerformanceRewardsSubmission: duration must be a multiple of CALCULATION_INTERVAL_SECONDS"
        );
        require(
            performanceRewardsSubmission.startTimestamp % CALCULATION_INTERVAL_SECONDS == 0,
            "RewardsCoordinator._validatePerformanceRewardsSubmission: startTimestamp must be a multiple of CALCULATION_INTERVAL_SECONDS"
        );
        require(
            block.timestamp - MAX_RETROACTIVE_LENGTH <= performanceRewardsSubmission.startTimestamp &&
                GENESIS_REWARDS_TIMESTAMP <= performanceRewardsSubmission.startTimestamp,
            "RewardsCoordinator._validatePerformanceRewardsSubmission: startTimestamp too far in the past"
        );
        require(
            performanceRewardsSubmission.startTimestamp + performanceRewardsSubmission.duration < block.timestamp,
            "RewardsCoordinator._validatePerformanceRewardsSubmission: performance rewards submission is not retroactive"
        );

        // Require performanceRewardsSubmission is for whitelisted strategy or beaconChainETHStrategy
        address currAddress = address(0);
        for (uint256 i = 0; i < performanceRewardsSubmission.strategiesAndMultipliers.length; ++i) {
            IStrategy strategy = performanceRewardsSubmission.strategiesAndMultipliers[i].strategy;
            require(
                strategyManager.strategyIsWhitelistedForDeposit(strategy) || strategy == beaconChainETHStrategy,
                "RewardsCoordinator._validatePerformanceRewardsSubmission: invalid strategy considered"
            );
            require(
                currAddress < address(strategy),
                "RewardsCoordinator._validatePerformanceRewardsSubmission: strategies must be in ascending order to handle duplicates"
            );
            currAddress = address(strategy);
        }

        return totalAmount;
    }

    function _checkClaim(RewardsMerkleClaim calldata claim, DistributionRoot memory root) internal view {
        require(!root.disabled, "RewardsCoordinator._checkClaim: root is disabled");
        require(block.timestamp >= root.activatedAt, "RewardsCoordinator._checkClaim: root not activated yet");
        require(
            claim.tokenIndices.length == claim.tokenTreeProofs.length,
            "RewardsCoordinator._checkClaim: tokenIndices and tokenProofs length mismatch"
        );
        require(
            claim.tokenTreeProofs.length == claim.tokenLeaves.length,
            "RewardsCoordinator._checkClaim: tokenTreeProofs and leaves length mismatch"
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
        // Validate index size so that there aren't multiple valid indices for the given proof
        // index can't be greater than 2**(tokenProof/32)
        require(
            tokenLeafIndex < (1 << (tokenProof.length / 32)),
            "RewardsCoordinator._verifyTokenClaim: invalid tokenLeafIndex"
        );

        // Verify inclusion of token leaf
        bytes32 tokenLeafHash = calculateTokenLeafHash(tokenLeaf);
        require(
            Merkle.verifyInclusionKeccak({
                root: earnerTokenRoot,
                index: tokenLeafIndex,
                proof: tokenProof,
                leaf: tokenLeafHash
            }),
            "RewardsCoordinator._verifyTokenClaim: invalid token claim proof"
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
        // Validate index size so that there aren't multiple valid indices for the given proof
        // index can't be greater than 2**(earnerProof/32)
        require(
            earnerLeafIndex < (1 << (earnerProof.length / 32)),
            "RewardsCoordinator._verifyEarnerClaimProof: invalid earnerLeafIndex"
        );
        // Verify inclusion of earner leaf
        bytes32 earnerLeafHash = calculateEarnerLeafHash(earnerLeaf);
        // forgefmt: disable-next-item
        require(
            Merkle.verifyInclusionKeccak({
                root: root,
                index: earnerLeafIndex,
                proof: earnerProof,
                leaf: earnerLeafHash
            }),
            "RewardsCoordinator._verifyEarnerClaimProof: invalid earner claim proof"
        );
    }

    /**
     * @notice Internal helper to process reward claims.
     * @param claim The RewardsMerkleClaims to be processed.
     * @param recipient The address recipient that receives the ERC20 rewards
     */
    function _processClaim(RewardsMerkleClaim calldata claim, address recipient) internal {
        DistributionRoot memory root = _distributionRoots[claim.rootIndex];
        _checkClaim(claim, root);
        // If claimerFor earner is not set, claimer is by default the earner. Else set to claimerFor
        address earner = claim.earnerLeaf.earner;
        address claimer = claimerFor[earner];
        if (claimer == address(0)) {
            claimer = earner;
        }
        require(msg.sender == claimer, "RewardsCoordinator.processClaim: caller is not valid claimer");
        for (uint256 i = 0; i < claim.tokenIndices.length; i++) {
            TokenTreeMerkleLeaf calldata tokenLeaf = claim.tokenLeaves[i];

            uint256 currCumulativeClaimed = cumulativeClaimed[earner][tokenLeaf.token];
            require(
                tokenLeaf.cumulativeEarnings > currCumulativeClaimed,
                "RewardsCoordinator.processClaim: cumulativeEarnings must be gt than cumulativeClaimed"
            );

            // Calculate amount to claim and update cumulativeClaimed
            uint256 claimAmount = tokenLeaf.cumulativeEarnings - currCumulativeClaimed;
            cumulativeClaimed[earner][tokenLeaf.token] = tokenLeaf.cumulativeEarnings;

            tokenLeaf.token.safeTransfer(recipient, claimAmount);
            emit RewardsClaimed(root.root, earner, claimer, recipient, tokenLeaf.token, claimAmount);
        }
    }

    function _setActivationDelay(uint32 _activationDelay) internal {
        emit ActivationDelaySet(activationDelay, _activationDelay);
        activationDelay = _activationDelay;
    }

    function _setGlobalOperatorCommission(uint16 _globalCommissionBips) internal {
        emit GlobalCommissionBipsSet(globalOperatorCommissionBips, _globalCommissionBips);
        globalOperatorCommissionBips = _globalCommissionBips;
    }

    function _setRewardsUpdater(address _rewardsUpdater) internal {
        emit RewardsUpdaterSet(rewardsUpdater, _rewardsUpdater);
        rewardsUpdater = _rewardsUpdater;
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IRewardsCoordinator
    function calculateEarnerLeafHash(EarnerTreeMerkleLeaf calldata leaf) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(EARNER_LEAF_SALT, leaf.earner, leaf.earnerTokenRoot));
    }

    /// @inheritdoc IRewardsCoordinator
    function calculateTokenLeafHash(TokenTreeMerkleLeaf calldata leaf) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(TOKEN_LEAF_SALT, leaf.token, leaf.cumulativeEarnings));
    }

    /// @inheritdoc IRewardsCoordinator
    function checkClaim(RewardsMerkleClaim calldata claim) public view returns (bool) {
        _checkClaim(claim, _distributionRoots[claim.rootIndex]);
        return true;
    }

    /// @inheritdoc IRewardsCoordinator
    function getOperatorAVSCommission(address operator, address avs) external view returns (uint16) {
        OperatorCommission memory operatorCommission = operatorAVSCommissionBips[operator][avs];

        if (operatorCommission.activatedAt == 0) {
            // Return the Global Operator Commission if the operator commission has not been initialized.
            return globalOperatorCommissionBips;
        } else {
            // Return the new commission if the new commission has been activated, else return the old commission.
            return
                (block.timestamp >= operatorCommission.activatedAt)
                    ? operatorCommission.newCommissionBips
                    : operatorCommission.oldCommissionBips;
        }
    }

    /// @inheritdoc IRewardsCoordinator
    function getOperatorPICommission(address operator) external view returns (uint16) {
        OperatorCommission memory operatorCommission = operatorPICommissionBips[operator];

        if (operatorCommission.activatedAt == 0) {
            // Return the Global Operator Commission if the operator commission has not been initialized.
            return globalOperatorCommissionBips;
        } else {
            // Return the new commission if the new commission has been activated, else return the old commission.
            return
                (block.timestamp >= operatorCommission.activatedAt)
                    ? operatorCommission.newCommissionBips
                    : operatorCommission.oldCommissionBips;
        }
    }

    /// @inheritdoc IRewardsCoordinator
    function getDistributionRootsLength() public view returns (uint256) {
        return _distributionRoots.length;
    }

    /// @inheritdoc IRewardsCoordinator
    function getDistributionRootAtIndex(uint256 index) external view returns (DistributionRoot memory) {
        return _distributionRoots[index];
    }

    /// @inheritdoc IRewardsCoordinator
    function getCurrentDistributionRoot() external view returns (DistributionRoot memory) {
        return _distributionRoots[_distributionRoots.length - 1];
    }

    /// @inheritdoc IRewardsCoordinator
    function getCurrentClaimableDistributionRoot() external view returns (DistributionRoot memory) {
        for (uint256 i = _distributionRoots.length; i > 0; i--) {
            DistributionRoot memory root = _distributionRoots[i - 1];
            if (!root.disabled && block.timestamp >= root.activatedAt) {
                return root;
            }
        }
        // Silence compiler warning.
        return DistributionRoot(bytes32(0), 0, 0, false);
    }

    /// @inheritdoc IRewardsCoordinator
    function getRootIndexFromHash(bytes32 rootHash) public view returns (uint32) {
        for (uint32 i = uint32(_distributionRoots.length); i > 0; i--) {
            if (_distributionRoots[i - 1].root == rootHash) {
                return i - 1;
            }
        }
        revert("RewardsCoordinator.getRootIndexFromHash: root not found");
    }

    /// @inheritdoc IRewardsCoordinator
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
