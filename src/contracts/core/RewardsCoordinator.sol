// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../libraries/Merkle.sol";
import "../interfaces/IStrategyManager.sol";
import "../permissions/Pausable.sol";
import "./RewardsCoordinatorStorage.sol";
import "../mixins/PermissionControllerMixin.sol";

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
    RewardsCoordinatorStorage,
    PermissionControllerMixin
{
    using SafeERC20 for IERC20;

    modifier onlyRewardsUpdater() {
        require(msg.sender == rewardsUpdater, UnauthorizedCaller());
        _;
    }

    modifier onlyRewardsForAllSubmitter() {
        require(isRewardsForAllSubmitter[msg.sender], UnauthorizedCaller());
        _;
    }

    /// @dev Sets the immutable variables for the contract
    constructor(
        IDelegationManager _delegationManager,
        IStrategyManager _strategyManager,
        IPauserRegistry _pauserRegistry,
        IPermissionController _permissionController,
        uint32 _CALCULATION_INTERVAL_SECONDS,
        uint32 _MAX_REWARDS_DURATION,
        uint32 _MAX_RETROACTIVE_LENGTH,
        uint32 _MAX_FUTURE_LENGTH,
        uint32 _GENESIS_REWARDS_TIMESTAMP
    )
        RewardsCoordinatorStorage(
            _delegationManager,
            _strategyManager,
            _CALCULATION_INTERVAL_SECONDS,
            _MAX_REWARDS_DURATION,
            _MAX_RETROACTIVE_LENGTH,
            _MAX_FUTURE_LENGTH,
            _GENESIS_REWARDS_TIMESTAMP
        )
        Pausable(_pauserRegistry)
        PermissionControllerMixin(_permissionController)
    {
        _disableInitializers();
    }

    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, rewardsUpdater and
     * configures the initial paused status, activationDelay, and globalOperatorCommissionBips.
     */
    function initialize(
        address initialOwner,
        uint256 initialPausedStatus,
        address _rewardsUpdater,
        uint32 _activationDelay,
        uint16 _globalCommissionBips
    ) external initializer {
        _setPausedStatus(initialPausedStatus);
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
        address avs,
        RewardsSubmission[] calldata rewardsSubmissions
    ) external onlyWhenNotPaused(PAUSED_AVS_REWARDS_SUBMISSION) checkCanCall(avs) nonReentrant {
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
                msg.sender, nonce, rewardsSubmissionForAllEarnersHash, rewardsSubmission
            );
            rewardsSubmission.token.safeTransferFrom(msg.sender, address(this), rewardsSubmission.amount);
        }
    }

    /// @inheritdoc IRewardsCoordinator
    function processClaim(
        RewardsMerkleClaim calldata claim,
        address recipient
    ) external onlyWhenNotPaused(PAUSED_PROCESS_CLAIM) nonReentrant {
        DistributionRoot memory root = _distributionRoots[claim.rootIndex];
        _checkClaim(claim, root);
        // If claimerFor earner is not set, claimer is by default the earner. Else set to claimerFor
        address earner = claim.earnerLeaf.earner;
        address claimer = claimerFor[earner];
        if (claimer == address(0)) {
            claimer = earner;
        }
        require(msg.sender == claimer, UnauthorizedCaller());
        for (uint256 i = 0; i < claim.tokenIndices.length; i++) {
            TokenTreeMerkleLeaf calldata tokenLeaf = claim.tokenLeaves[i];

            uint256 currCumulativeClaimed = cumulativeClaimed[earner][tokenLeaf.token];
            require(tokenLeaf.cumulativeEarnings > currCumulativeClaimed, EarningsNotGreaterThanClaimed());

            // Calculate amount to claim and update cumulativeClaimed
            uint256 claimAmount = tokenLeaf.cumulativeEarnings - currCumulativeClaimed;
            cumulativeClaimed[earner][tokenLeaf.token] = tokenLeaf.cumulativeEarnings;

            tokenLeaf.token.safeTransfer(recipient, claimAmount);
            emit RewardsClaimed(root.root, earner, claimer, recipient, tokenLeaf.token, claimAmount);
        }
    }

    /// @inheritdoc IRewardsCoordinator
    function submitRoot(
        bytes32 root,
        uint32 rewardsCalculationEndTimestamp
    ) external onlyWhenNotPaused(PAUSED_SUBMIT_DISABLE_ROOTS) onlyRewardsUpdater {
        require(
            rewardsCalculationEndTimestamp > currRewardsCalculationEndTimestamp, NewRootMustBeForNewCalculatedPeriod()
        );
        require(rewardsCalculationEndTimestamp < block.timestamp, RewardsEndTimestampNotElapsed());
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
    function disableRoot(
        uint32 rootIndex
    ) external onlyWhenNotPaused(PAUSED_SUBMIT_DISABLE_ROOTS) onlyRewardsUpdater {
        require(rootIndex < _distributionRoots.length, InvalidRootIndex());
        DistributionRoot storage root = _distributionRoots[rootIndex];
        require(!root.disabled, RootDisabled());
        require(block.timestamp < root.activatedAt, RootAlreadyActivated());
        root.disabled = true;
        emit DistributionRootDisabled(rootIndex);
    }

    /// @inheritdoc IRewardsCoordinator
    function setClaimerFor(
        address claimer
    ) external {
        address earner = msg.sender;
        _setClaimer(earner, claimer);
    }

    /// @inheritdoc IRewardsCoordinator
    function setClaimerFor(address earner, address claimer) external checkCanCall(earner) {
        require(delegationManager.isOperator(earner), EarnerNotOperator());
        _setClaimer(earner, claimer);
    }

    /// @inheritdoc IRewardsCoordinator
    function setActivationDelay(
        uint32 _activationDelay
    ) external onlyOwner {
        _setActivationDelay(_activationDelay);
    }

    /// @inheritdoc IRewardsCoordinator
    function setGlobalOperatorCommission(
        uint16 _globalCommissionBips
    ) external onlyOwner {
        _setGlobalOperatorCommission(_globalCommissionBips);
    }

    /// @inheritdoc IRewardsCoordinator
    function setRewardsUpdater(
        address _rewardsUpdater
    ) external onlyOwner {
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
    function _validateRewardsSubmission(
        RewardsSubmission calldata rewardsSubmission
    ) internal view {
        require(rewardsSubmission.strategiesAndMultipliers.length > 0, InputArrayLengthZero());
        require(rewardsSubmission.amount > 0, AmountIsZero());
        require(rewardsSubmission.amount <= MAX_REWARDS_AMOUNT, AmountExceedsMax());
        require(rewardsSubmission.duration <= MAX_REWARDS_DURATION, DurationExceedsMax());
        require(rewardsSubmission.duration % CALCULATION_INTERVAL_SECONDS == 0, InvalidDurationRemainder());
        require(rewardsSubmission.startTimestamp % CALCULATION_INTERVAL_SECONDS == 0, InvalidStartTimestampRemainder());
        require(
            block.timestamp - MAX_RETROACTIVE_LENGTH <= rewardsSubmission.startTimestamp
                && GENESIS_REWARDS_TIMESTAMP <= rewardsSubmission.startTimestamp,
            StartTimestampTooFarInPast()
        );
        require(rewardsSubmission.startTimestamp <= block.timestamp + MAX_FUTURE_LENGTH, StartTimestampTooFarInFuture());

        // Require rewardsSubmission is for whitelisted strategy or beaconChainETHStrategy
        address currAddress = address(0);
        for (uint256 i = 0; i < rewardsSubmission.strategiesAndMultipliers.length; ++i) {
            IStrategy strategy = rewardsSubmission.strategiesAndMultipliers[i].strategy;
            require(
                strategyManager.strategyIsWhitelistedForDeposit(strategy) || strategy == beaconChainETHStrategy,
                StrategyNotWhitelisted()
            );
            require(currAddress < address(strategy), StrategiesNotInAscendingOrder());
            currAddress = address(strategy);
        }
    }

    function _checkClaim(RewardsMerkleClaim calldata claim, DistributionRoot memory root) internal view {
        require(!root.disabled, RootDisabled());
        require(block.timestamp >= root.activatedAt, RootNotActivated());
        require(claim.tokenIndices.length == claim.tokenTreeProofs.length, InputArrayLengthMismatch());
        require(claim.tokenTreeProofs.length == claim.tokenLeaves.length, InputArrayLengthMismatch());

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
        require(tokenLeafIndex < (1 << (tokenProof.length / 32)), InvalidTokenLeafIndex());

        // Verify inclusion of token leaf
        bytes32 tokenLeafHash = calculateTokenLeafHash(tokenLeaf);
        require(
            Merkle.verifyInclusionKeccak({
                root: earnerTokenRoot,
                index: tokenLeafIndex,
                proof: tokenProof,
                leaf: tokenLeafHash
            }),
            InvalidClaimProof()
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
        require(earnerLeafIndex < (1 << (earnerProof.length / 32)), InvalidEarnerLeafIndex());
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
            InvalidClaimProof()
        );
    }

    function _setActivationDelay(
        uint32 _activationDelay
    ) internal {
        emit ActivationDelaySet(activationDelay, _activationDelay);
        activationDelay = _activationDelay;
    }

    function _setGlobalOperatorCommission(
        uint16 _globalCommissionBips
    ) internal {
        emit GlobalCommissionBipsSet(globalOperatorCommissionBips, _globalCommissionBips);
        globalOperatorCommissionBips = _globalCommissionBips;
    }

    function _setRewardsUpdater(
        address _rewardsUpdater
    ) internal {
        emit RewardsUpdaterSet(rewardsUpdater, _rewardsUpdater);
        rewardsUpdater = _rewardsUpdater;
    }

    function _setClaimer(address earner, address claimer) internal {
        address prevClaimer = claimerFor[earner];
        claimerFor[earner] = claimer;
        emit ClaimerForSet(earner, prevClaimer, claimer);
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IRewardsCoordinator
    function calculateEarnerLeafHash(
        EarnerTreeMerkleLeaf calldata leaf
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(EARNER_LEAF_SALT, leaf.earner, leaf.earnerTokenRoot));
    }

    /// @inheritdoc IRewardsCoordinator
    function calculateTokenLeafHash(
        TokenTreeMerkleLeaf calldata leaf
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(TOKEN_LEAF_SALT, leaf.token, leaf.cumulativeEarnings));
    }

    /// @inheritdoc IRewardsCoordinator
    function checkClaim(
        RewardsMerkleClaim calldata claim
    ) public view returns (bool) {
        _checkClaim(claim, _distributionRoots[claim.rootIndex]);
        return true;
    }

    /// @inheritdoc IRewardsCoordinator
    function operatorCommissionBips(
        address, // operator
        address // avs
    ) external view returns (uint16) {
        return globalOperatorCommissionBips;
    }

    /// @inheritdoc IRewardsCoordinator
    function getDistributionRootsLength() public view returns (uint256) {
        return _distributionRoots.length;
    }

    /// @inheritdoc IRewardsCoordinator
    function getDistributionRootAtIndex(
        uint256 index
    ) external view returns (DistributionRoot memory) {
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
    function getRootIndexFromHash(
        bytes32 rootHash
    ) public view returns (uint32) {
        for (uint32 i = uint32(_distributionRoots.length); i > 0; i--) {
            if (_distributionRoots[i - 1].root == rootHash) {
                return i - 1;
            }
        }
        revert InvalidRoot();
    }
}
