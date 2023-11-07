// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IStrategyManager.sol";
import "../interfaces/IStrategy.sol";
import "../interfaces/IEigenPodManager.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/ISlasher.sol";
import "../permissions/Pausable.sol";
import "@openzeppelin/contracts/utils/cryptography/draft-EIP712.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

/**
 * @title Storage variables for the `Slasher` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract SlasherStorage is Initializable, OwnableUpgradeable, ISlasher, Pausable, EIP712 {
    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");
    /// @notice The EIP-712 typehash for the deposit struct used by the contract
    bytes32 public constant REGISTRATION_WITHOUT_SLASHING_TYPEHASH =
        keccak256("RegistrationWithoutSlashing(address operator,address avs,bytes32 salt,uint256 expiry)");

    /// @notice The central StrategyManager contract of EigenLayer
    IStrategyManager public immutable strategyManager;
    /// @notice The DelegationManager contract of EigenLayer
    IDelegationManager public immutable delegation;

    /// @notice operator => salt => whether that salt has been used
    mapping(address => mapping(bytes32 => bool)) public isOperatorSaltUsed;

    /// @notice avs => operator => bool whether the operator is registered without slashing for the avs
    mapping(address => mapping(address => bool)) public registeredWithoutSlashing;

    constructor(IStrategyManager _strategyManager, IDelegationManager _delegation, string memory eip712version) EIP712("SlashingManager", eip712version) {
        strategyManager = _strategyManager;
        delegation = _delegation;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}
