// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IStrategyManager.sol";
import "../interfaces/IAllocatorManager.sol";
import "../interfaces/ISlasher.sol";
import "../interfaces/IEigenPodManager.sol";
import "../interfaces/IAVSDirectory.sol";

/**
 * @title Storage variables for the `AlloctionManager` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract AllocatorManagerStorage is IAllocatorManager {
    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

    /// @notice The EIP-712 typehash for the `StakerDelegation` struct used by the contract
    bytes32 public constant STAKER_DELEGATION_TYPEHASH =
        keccak256("StakerDelegation(address staker,address operator,uint256 nonce,uint256 expiry)");

    /// @notice The EIP-712 typehash for the `DelegationApproval` struct used by the contract
    bytes32 public constant DELEGATION_APPROVAL_TYPEHASH = keccak256(
        "DelegationApproval(address delegationApprover,address staker,address operator,bytes32 salt,uint256 expiry)"
    );

    uint256 internal constant BIG_NUMBER = 1e18;

    /**
     * @notice Original EIP-712 Domain separator for this contract.
     * @dev The domain separator may change in the event of a fork that modifies the ChainID.
     * Use the getter function `domainSeparator` to get the current domain separator for this contract.
     */
    bytes32 internal _DOMAIN_SEPARATOR;

    /// @notice The StrategyManager contract for EigenLayer
    IStrategyManager public immutable strategyManager;

    /// @notice The DelegateManager contract for EigenLayer
    IDelegationManager public immutable delegationManager;

    /// @notice The EigenPodManager contract for EigenLayer
    IEigenPodManager public immutable eigenPodManager;

    /// @notice The AVSDirectory contract for EigenLayer
    IAVSDirectory public immutable avsDirectory;

    // the number of 12-second blocks in 30 days (60 * 60 * 24 * 30 / 12 = 216,000)
    uint256 public constant MAX_WITHDRAWAL_DELAY_BLOCKS = 216_000;

    /// @notice Mapping: allocator => the number of shares assigned to the allocator scaled up by the allocator's total magnitude.
    mapping(address => mapping(IStrategy => uint256)) public scaledDelegatedShares;

    struct AllocatorDetails {
        address delegationApprover;
        bool isAllocator; // if true, then the address is an allocator. Needed beacause selectionApprover can be 0.
    }

    mapping(address => AllocatorDetails) internal _allocatorDetails;

    /// @notice Mapping: staker => whether the staker is a post SDA staker (i.e. they have no left over state unmigrated from the DelegationManager).
    mapping(address => bool) public isPostSDAStaker;

    /**
     * @notice Mapping: staker => allocator whom the staker is currently delegated to.
     * @dev Note that returning address(0) indicates that the staker is not actively delegated to any allocator.
     */
    mapping(address => address) internal _delegatedTo;

    /// @notice Mapping: staker => number of signed messages (used in `delegateToBySignature`) from the staker that this contract has already checked.
    mapping(address => uint256) public stakerNonce;

    /**
     * @notice Mapping: delegationApprover => 32-byte salt => whether or not the salt has already been used by the delegationApprover.
     * @dev Salts are used in the `delegateTo` and `delegateToBySignature` functions. Note that these functions only process the delegationApprover's
     * signature + the provided salt if the allocator being delegated to has specified a nonzero address as their `delegationApprover`.
     */
    mapping(address => mapping(bytes32 => bool)) public delegationApproverSaltIsSpent;

    /**
     * TODO: edit
     * @notice Global minimum withdrawal delay for all strategy withdrawals.
     * In a prior Goerli release, we only had a global min withdrawal delay across all strategies.
     * In addition, we now also configure withdrawal delays on a per-strategy basis.
     * To withdraw from a strategy, max(minWithdrawalDelayBlocks, strategyWithdrawalDelayBlocks[strategy]) number of blocks must have passed.
     * See mapping strategyWithdrawalDelayBlocks below for per-strategy withdrawal delays.
     */
    uint256 public minWithdrawalDelayBlocks;

    /// @notice Mapping: hash of withdrawal inputs, aka 'withdrawalRoot' => whether the withdrawal is pending
    mapping(bytes32 => bool) public pendingWithdrawals;

    /// @notice Mapping: staker => cumulative number of queued withdrawals they have ever initiated.
    /// @dev This only increments (doesn't decrement), and is used to help ensure that otherwise identical withdrawals have unique hashes.
    mapping(address => uint256) public cumulativeWithdrawalsQueued;


    /**
     * @notice Minimum delay enforced by this contract per Strategy for completing queued withdrawals. Measured in blocks, and adjustable by this contract's owner,
     * up to a maximum of `MAX_WITHDRAWAL_DELAY_BLOCKS`. Minimum value is 0 (i.e. no delay enforced).
     */
    mapping(IStrategy => uint256) public strategyWithdrawalDelayBlocks;

    constructor(IStrategyManager _strategyManager, IDelegationManager _delegationManager, IEigenPodManager _eigenPodManager, IAVSDirectory _avsDirectory) {
        strategyManager = _strategyManager;
        delegationManager = _delegationManager;
        eigenPodManager = _eigenPodManager;
        avsDirectory = _avsDirectory;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}
