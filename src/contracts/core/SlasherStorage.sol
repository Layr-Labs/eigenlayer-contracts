// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/ISlasher.sol";
import "../permissions/Pausable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

/**
 * @notice SlasherStorage
 */
abstract contract SlasherStorage is Initializable, OwnableUpgradeable, ISlasher, Pausable {
    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");
    /// @notice The EIP-712 typehash for the `QueueReallocation` struct used by the contract
    bytes32 public constant REALLOCATION_TYPEHASH =
        keccak256("Reallocation(address operator,MagnitudeAdjustmentsParam[],bytes32 salt,uint256 expiry)");
    /// @notice The EIP-712 typehash for the `MagnitudeDilution` struct used by the contract
    bytes32 public constant MAGNITUDE_DILUTION_TYPEHASH =
        keccak256("MagnitudeDilution(address operator,IStrategy[] strategies,uint64[] nonslashableToAdd,bytes32 salt,uint256 expiry)");
    /// @notice The EIP-712 typehash for the `MagnitudeConcentration` struct used by the contract
    bytes32 public constant MAGNITUDE_CONCENTRATION_TYPEHASH =
        keccak256("MagnitudeConcentration(address operator,IStrategy[] strategies,uint64[] nonslashableToDecrement,bytes32 salt,uint256 expiry)");

    /**
     * @notice Original EIP-712 Domain separator for this contract.
     * @dev The domain separator may change in the event of a fork that modifies the ChainID.
     * Use the getter function `domainSeparator` to get the current domain separator for this contract.
     */
    bytes32 internal _DOMAIN_SEPARATOR;

    // system contracts
    IStrategyManager public immutable strategyManager;
    IDelegationManager public immutable delegation;

    struct SlashingUpdate {
        uint64 slashingRate; // This is parts per (BIPS_FACTOR**2), i.e. parts per 1e8, pphm = parts per hundred million, to slash upon execution
        uint64 scalingFactor; // the scaling factor to apply to the operator's shares. this is only set upon execution
    }

    /// @notice Mapping: Operator => Strategy => epoch => SlashingUpdate[]
    mapping(address => mapping(IStrategy => mapping(uint32 => SlashingUpdate))) public slashingUpdates;

    /**
     * @notice Mapping: operator => strategy => share scalingFactor,
     * stored in the "SHARE_CONVERSION_SCALE", i.e. scalingFactor = 2 * SHARE_CONVERSION_SCALE indicates a scalingFactor of "2".
     * Note that a value of zero is treated as one, since this means that the operator has never been slashed
     */
    mapping(address => mapping(IStrategy => uint64)) internal _shareScalingFactor;

    /// @notice Mapping: operator => strategy => epochs in which the strategy was slashed for the operator
    // TODO: note that since default will be 0, we should probably make the "first epoch" actually be epoch 1 or something
    mapping(address => mapping(IStrategy => uint32[])) public slashingEpochHistory;

    /// @notice Mapping: allocator => 32-byte salt => whether or not the salt has already been used by the operator.
    /// @dev Salt is used in the `queueReallocation`, `queueMagnitudeConcentration`, `queueMagnitudeDilution` functions.
    mapping(address => mapping(bytes32 => bool)) public allocatorSaltIsSpent;

    constructor(
        IStrategyManager _strategyManager,
        IDelegationManager _delegationManager
    ) {
        strategyManager = _strategyManager;
        delegation = _delegationManager;
    }
}
