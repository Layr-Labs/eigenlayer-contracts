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
        IAVSDirectory _avsDirectory,
        IDelegationManager _delegationManager,
        IEigenPodManager _eigenPodManager,
        IStrategyManager _strategyManager,
        ISlasher _slasher,
        uint64 _maxPaymentDuration,
        uint64 _lowerBoundStartRange
    )
        PaymentCoordinatorStorage(
            _avsDirectory,
            _delegationManager,
            _eigenPodManager,
            _strategyManager,
            _slasher,
            _maxPaymentDuration,
            _lowerBoundStartRange
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
            emit RangePaymentCreated(msg.sender, rangePaymentHash, rangePayment);
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
            emit RangePaymentForAllCreated(msg.sender, rangePaymentHash, rangePayment);
        }
    }

    function _payForRange(RangePayment calldata rangePayment) internal returns (bytes32) {
        require(rangePayment.strategiesAndMultlipliers.length > 0, "PaymentCoordinator.payForRange: no strategies set");
        require(rangePayment.amount > 0, "PaymentCoordinator.payForRange: amount cannot be 0");
        require(
            rangePayment.duration <= MAX_PAYMENT_DURATION,
            "PaymentCoordinator.payForRange: duration exceeds MAX_PAYMENT_DURATION"
        );
        require(
            rangePayment.duration % calculationIntervalSeconds == 0,
            "PaymentCoordinator.payForRange: duration must be a multiple of calculationIntervalSeconds"
        );
        // If retroactive payments enabled must at least be past LOWER_BOUND_START_RANGE,
        // otherwise should start earliest at current block timestamp
        require(
            (retroactivePaymentsEnabled && rangePayment.startTimestamp >= LOWER_BOUND_START_RANGE) ||
                rangePayment.startTimestamp >= block.timestamp,
            "PaymentCoordinator.payForRange: invalid startTimestamp set"
        );

        // Require rangePayment is for whitelisted strategies
        for (uint256 j = 0; j < rangePayment.strategiesAndMultlipliers.length; j++) {
            IStrategy strategy = rangePayment.strategiesAndMultlipliers[j].strategy;
            require(
                strategyManager.strategyIsWhitelistedForDeposit(strategy),
                "PaymentCoordinator.payForRange: strategy not whitelisted for deposit"
            );
        }

        // Set hash of rangePayment in mapping
        bytes32 rangePaymentHash = keccak256(abi.encode(msg.sender, rangePayment));
        isRangePaymentHash[msg.sender][rangePaymentHash] = true;

        rangePayment.token.safeTransferFrom(msg.sender, address(this), rangePayment.amount);
        return rangePaymentHash;
    }

    /**
     * @notice Claim payments for the given claim
     * @param claim The claims to be processed
     */
    function processClaims(PaymentMerkleClaim calldata claim) external onlyWhenNotPaused(PAUSED_CLAIM_PAYMENTS) {}

    /**
     * @notice Sets the address of the entity that can claim payments on behalf of the account
     * @param account The account whose recipient is being set
     * @param recipient The address of the entity that can claim payments on behalf of the account
     * @dev Only callable by the `account`
     */
    function setRecipient(address account, address recipient) external {}

    /**
     * @notice Creates a new distribution root
     * @param root The merkle root of the distribution
     * @param paymentsCalculatedUntilTimestamp The timestamp until which payments have been calculated
     * @dev Only callable by the paymentUpdater
     */
    function submitRoot(bytes32 root, uint32 paymentsCalculatedUntilTimestamp) external onlyPaymentUpdater {}

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
