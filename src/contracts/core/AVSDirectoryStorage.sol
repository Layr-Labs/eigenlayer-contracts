// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IStrategyManager.sol";
import "../interfaces/IDelegationManager.sol";

abstract contract AVSDirectoryStorage is IAVSDirectory {
    /// @notice Canonical, virtual beacon chain ETH strategy
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

    /// @notice The EIP-712 typehash for the `Registration` struct used by the contract
    bytes32 public constant OPERATOR_AVS_REGISTRATION_TYPEHASH =
        keccak256("OperatorAVSRegistration(address operator,address avs,bytes32 salt,uint256 expiry)");

    /// @notice The EIP-712 typehash for the `OperatorSetRegistration` struct used by the contract
    bytes32 public constant OPERATOR_SET_REGISTRATION_TYPEHASH =
        keccak256("OperatorSetRegistration(address avs,uint32[] operatorSetIds,bytes32 salt,uint256 expiry)");

    /// @notice The EIP-712 typehash for the `StandbyParams` struct used by the contract
    bytes32 public constant OPERATOR_STANDBY_UPDATE =
        keccak256("OperatorStandbyUpdate(StandbyParam[] standbyParams,bytes32 salt,uint256 expiry)");

    /// @dev Index for flag that pauses operator register/deregister to avs when set.
    uint8 internal constant PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS = 0;

    /// @dev The length of time in a given epoch.
    uint256 internal constant ONE_EPOCH = 7 days;

    /// @dev Chain ID at the time of contract deployment
    uint256 internal immutable ORIGINAL_CHAIN_ID = block.chainid;

    /// @dev The time that epochs start from.
    uint256 internal immutable EPOCH_GENESIS_TIME = block.timestamp;

    /// @notice The DelegationManager contract for EigenLayer
    IDelegationManager public immutable delegation;

    /**
     * @notice Original EIP-712 Domain separator for this contract.
     * @dev The domain separator may change in the event of a fork that modifies the ChainID.
     * Use the getter function `domainSeparator` to get the current domain separator for this contract.
     */
    bytes32 internal _DOMAIN_SEPARATOR;

    /// @notice Mapping: AVS => operator => enum of operator status to the AVS
    mapping(address => mapping(address => OperatorAVSRegistrationStatus)) public avsOperatorStatus;

    /// @notice Mapping: operator => 32-byte salt => whether or not the salt has already been used by the operator.
    /// @dev Salt is used in the `registerOperatorToAVS` and `registerOperatorToOperatorSet` function.
    mapping(address => mapping(bytes32 => bool)) public operatorSaltIsSpent;

    /// @notice Mapping: AVS => whether or not the AVS uses operator set
    mapping(address => bool) public isOperatorSetAVS;

    /// @notice Mapping: (avs, operator, operatorSetId) => (lastEpoch, onStandby)
    mapping(address => mapping(address => mapping(uint32 => OperatorInfo))) public operatorInfo;

    /// @notice Mapping: (avs, operator, operatorSetId, epochId) => (state)
    mapping(address => mapping(address => mapping(uint32 => mapping(uint256 => EpochStates)))) public operatorEpochState;

    /// @notice Mapping: avs => operator => number of operator sets the operator is registered for the AVS
    mapping(address => mapping(address => uint256)) public operatorAVSOperatorSetCount;

    constructor(IDelegationManager _delegation) {
        delegation = _delegation;
    }

    /// @notice Returns the current epoch.
    function currentEpoch() public view returns (uint256) {
        unchecked {
            // If `timestamp - EPOCH_GENESIS_TIME` is less than `ONE_EPOCH` the quotient is rounded down to zero.
            return (block.timestamp - EPOCH_GENESIS_TIME) / ONE_EPOCH;
        }
    }

    /// @dev Returns whether or not a given operator is registered in an operator set.
    function _isOperatorInOperatorSet(
        uint256 lastEpoch,
        EpochStates state
    ) internal view returns (bool) {
        // Overflow is not reasonably possible...
        unchecked {
            return
                !(state == EpochStates.NULL || (state == EpochStates.DEREGISTERED && currentEpoch() >= lastEpoch + 2));
        }
    }

    /// @notice Returns whether or not a given operator is registered in an operator set.
    function isOperatorInOperatorSet(address avs, address operator, uint32 operatorSetId) public view returns (bool) {
        uint256 lastEpoch = operatorInfo[avs][operator][operatorSetId].lastEpoch;
        return _isOperatorInOperatorSet(
            lastEpoch, operatorEpochState[avs][operator][operatorSetId][lastEpoch]
        );
    }

    /// @notice Returns whether or not a given operator is in standby mode.
    function onStandby(address avs, address operator, uint32 operatorSetId) public view returns (bool) {
        return operatorInfo[avs][operator][operatorSetId].onStandby;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[43] private __gap;
}
