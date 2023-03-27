// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IStrategyManager.sol";
import "../interfaces/IDelegationTerms.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/ISlasher.sol";

/**
 * @title Storage variables for the `DelegationManager` contract.
 * @author Layr Labs, Inc.
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract DelegationManagerStorage is IDelegationManager {
    /// @notice Gas budget provided in calls to DelegationTerms contracts
    uint256 internal constant LOW_LEVEL_GAS_BUDGET = 1e5;

    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

    /// @notice The EIP-712 typehash for the delegation struct used by the contract
    bytes32 public constant DELEGATION_TYPEHASH =
        keccak256("Delegation(address delegator,address operator,uint256 nonce,uint256 expiry)");

    /// @notice EIP-712 Domain separator
    bytes32 public DOMAIN_SEPARATOR;

    /// @notice The StrategyManager contract for EigenLayer
    IStrategyManager public immutable strategyManager;

    /// @notice The Slasher contract for EigenLayer
    ISlasher public immutable slasher;

    /// @notice Mapping: operator => strategy => total number of shares in the strategy delegated to the operator
    mapping(address => mapping(IStrategy => uint256)) public operatorShares;

    /// @notice Mapping: operator => delegation terms contract
    mapping(address => IDelegationTerms) public delegationTerms;

    /// @notice Mapping: staker => operator whom the staker has delegated to
    mapping(address => address) public delegatedTo;

    /// @notice Mapping: delegator => number of signed delegation nonce (used in delegateToBySignature)
    mapping(address => uint256) public nonces;

    constructor(IStrategyManager _strategyManager, ISlasher _slasher) {
        strategyManager = _strategyManager;
        slasher = _slasher;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[46] private __gap;
}