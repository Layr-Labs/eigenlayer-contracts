// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IStrategyManager.sol";
import "../interfaces/IStrategy.sol";
import "../interfaces/IEigenPodManager.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/ISlasher.sol";

/**
 * @title Storage variables for the `StrategyManager` contract.
 * @author Layr Labs, Inc.
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract StrategyManagerStorage is IStrategyManager {
    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");
    /// @notice The EIP-712 typehash for the deposit struct used by the contract
    bytes32 public constant DEPOSIT_TYPEHASH =
        keccak256("Deposit(address strategy,address token,uint256 amount,uint256 nonce,uint256 expiry)");
    /// @notice EIP-712 Domain separator
    bytes32 public DOMAIN_SEPARATOR;
    // staker => number of signed deposit nonce (used in depositIntoStrategyWithSignature)
    mapping(address => uint256) public nonces;

    // maximum length of dynamic arrays in `stakerStrategyList` mapping, for sanity's sake
    uint8 internal constant MAX_STAKER_STRATEGY_LIST_LENGTH = 32;

    // system contracts
    IDelegationManager public immutable delegation;
    IEigenPodManager public immutable eigenPodManager;
    ISlasher public immutable slasher;

    /// @notice Permissioned role, which can be changed by the contract owner. Has the ability to edit the strategy whitelist
    address public strategyWhitelister;

    /**
     * @notice Minimum delay enforced by this contract for completing queued withdrawals. Measured in blocks, and adjustable by this contract's owner,
     * up to a maximum of `MAX_WITHDRAWAL_DELAY_BLOCKS`. Minimum value is 0 (i.e. no delay enforced).
     * @dev Note that the withdrawal delay is not enforced on withdrawals of 'beaconChainETH', as the EigenPods have their own separate delay mechanic
     * and we want to avoid stacking multiple enforced delays onto a single withdrawal.
     */
    uint256 public withdrawalDelayBlocks;
    // the number of 12-second blocks in one week (60 * 60 * 24 * 7 / 12 = 50,400)
    uint256 public constant MAX_WITHDRAWAL_DELAY_BLOCKS = 50400;

    /// @notice Mapping: staker => Strategy => number of shares which they currently hold
    mapping(address => mapping(IStrategy => uint256)) public stakerStrategyShares;
    /// @notice Mapping: staker => array of strategies in which they have nonzero shares
    mapping(address => IStrategy[]) public stakerStrategyList;
    /// @notice Mapping: hash of withdrawal inputs, aka 'withdrawalRoot' => whether the withdrawal is pending
    mapping(bytes32 => bool) public withdrawalRootPending;
    /// @notice Mapping: staker => cumulative number of queued withdrawals they have ever initiated. only increments (doesn't decrement)
    mapping(address => uint256) public numWithdrawalsQueued;
    /// @notice Mapping: strategy => whether or not stakers are allowed to deposit into it
    mapping(IStrategy => bool) public strategyIsWhitelistedForDeposit;
    /*
     * @notice Mapping: staker => virtual 'beaconChainETH' shares that the staker 'owes' due to overcommitments of beacon chain ETH.
     * When overcommitment is proven, `StrategyManager.recordOvercommittedBeaconChainETH` is called. However, it is possible that the
     * staker already queued a withdrawal for more beaconChainETH shares than the `amount` input to this function. In this edge case,
     * the amount that cannot be decremented is added to the staker's `beaconChainETHSharesToDecrementOnWithdrawal` -- then when the staker completes a
     * withdrawal of beaconChainETH, the amount they are withdrawing is first decreased by their `beaconChainETHSharesToDecrementOnWithdrawal` amount.
     * In other words, a staker's `beaconChainETHSharesToDecrementOnWithdrawal` must be 'paid down' before they can "actually withdraw" beaconChainETH.
     * @dev In practice, this means not passing a call to `eigenPodManager.withdrawRestakedBeaconChainETH` until the staker's 
     * `beaconChainETHSharesToDecrementOnWithdrawal` has first been reduced to zero.
    */
    mapping(address => uint256) public beaconChainETHSharesToDecrementOnWithdrawal;

    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    constructor(IDelegationManager _delegation, IEigenPodManager _eigenPodManager, ISlasher _slasher) {
        delegation = _delegation;
        eigenPodManager = _eigenPodManager;
        slasher = _slasher;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[41] private __gap;
}
