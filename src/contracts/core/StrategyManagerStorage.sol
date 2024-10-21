// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IStrategyManager.sol";
import "../interfaces/IStrategy.sol";
import "../interfaces/IEigenPodManager.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IAVSDirectory.sol";

/**
 * @title Storage variables for the `StrategyManager` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract StrategyManagerStorage is IStrategyManager {
    // Constants

    /// @notice The EIP-712 typehash for the deposit struct used by the contract
    bytes32 public constant DEPOSIT_TYPEHASH =
        keccak256("Deposit(address staker,address strategy,address token,uint256 amount,uint256 nonce,uint256 expiry)");

    // maximum length of dynamic arrays in `stakerStrategyList` mapping, for sanity's sake
    uint8 internal constant MAX_STAKER_STRATEGY_LIST_LENGTH = 32;

    // index for flag that pauses deposits when set
    uint8 internal constant PAUSED_DEPOSITS = 0;

    // Immutables

    IDelegationManager public immutable delegation;

    // Mutatables

    bytes32 internal __deprecated_DOMAIN_SEPARATOR;

    // staker => number of signed deposit nonce (used in depositIntoStrategyWithSignature)
    mapping(address => uint256) public nonces;

    /// @notice Permissioned role, which can be changed by the contract owner. Has the ability to edit the strategy whitelist
    address public strategyWhitelister;

    /*
     * Reserved space previously used by the deprecated storage variable `withdrawalDelayBlocks.
     * This variable was migrated to the DelegationManager instead.
     */
    uint256 private __deprecated_withdrawalDelayBlocks;
    /// @notice Mapping: staker => Strategy => number of shares which they have deposited. All of these shares
    ///         may not be withdrawable if the staker has delegated to an operator that has been slashed.
    mapping(address staker => mapping(IStrategy strategy => uint256)) public stakerDepositShares;
    /// @notice Mapping: staker => array of strategies in which they have nonzero shares
    mapping(address staker => IStrategy[]) public stakerStrategyList;

    /// @notice *Deprecated* mapping: hash of withdrawal inputs, aka 'withdrawalRoot' => whether the withdrawal is pending
    /// @dev This mapping is preserved to allow the migration of withdrawals to the DelegationManager contract.
    mapping(bytes32 withdrawalRoot => bool) private __deprecated_withdrawalRootPending;
    /*
     * Reserved space previously used by the deprecated mapping(address => uint256) numWithdrawalsQueued.
     * This mapping tracked the cumulative number of queued withdrawals initiated by a staker.
     * Withdrawals are now initiated in the DlegationManager, so the mapping has moved to that contract.
     */
    mapping(address staker => uint256) private __deprecated_numWithdrawalsQueued;

    /// @notice Mapping: strategy => whether or not stakers are allowed to deposit into it
    mapping(IStrategy strategy => bool) public strategyIsWhitelistedForDeposit;
    /*
     * Reserved space previously used by the deprecated mapping(address => uint256) beaconChainETHSharesToDecrementOnWithdrawal.
     * This mapping tracked beaconChainETH "deficit" in cases where updates were made to shares retroactively.  However, this construction was
     * moved into the EigenPodManager contract itself.
     */
    mapping(address avs => uint256) internal beaconChainETHSharesToDecrementOnWithdrawal;

    /**
     * @notice Mapping: strategy => whether or not stakers are allowed to transfer strategy shares to another address
     * if true for a strategy, a user cannot depositIntoStrategyWithSignature into that strategy for another staker
     * and also when performing queueWithdrawals, a staker can only withdraw to themselves
     */
    mapping(IStrategy strategy => bool) private __deprecated_thirdPartyTransfersForbidden;

    // Construction

    /**
     * @param _delegation The delegation contract of EigenLayer.
     */
    constructor(
        IDelegationManager _delegation
    ) {
        delegation = _delegation;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[39] private __gap;
}
