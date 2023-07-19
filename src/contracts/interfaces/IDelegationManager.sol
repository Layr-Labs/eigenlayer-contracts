// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./IStrategy.sol";

/**
 * @title The interface for the primary delegation contract for EigenLayer.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice  This is the contract for delegation in EigenLayer. The main functionalities of this contract are
 * - enabling anyone to register as an operator in EigenLayer
 * - allowing operators to specify parameters related to stakers who delegate to them
 * - enabling any staker to delegate its stake to the operator of its choice (a given staker can only delegate to a single operator at a time)
 * - enabling a staker to undelegate its assets from the operator it is delegated to (performed as part of the withdrawal process, initiated through the StrategyManager)
 */
interface IDelegationManager {
    struct OperatorDetails {
        // @notice address to receive the rewards that the operator earns via serving applications built on EigenLayer.
        address earningsReceiver;
        /**
         * @notice Address to verify signatures when a staker wishes to delegate to the operator, as well as controlling "forced undelegations".
         * @dev Signature verification follows these rules:
         * 1) If this address is left as address(0), then any staker will be free to delegate to the operator, i.e. no signature verification will be performed.
         * 2) If this address is an EOA (i.e. it has no code), then we follow standard ECDSA signature verification for delegations to the operator.
         * 3) If this address is a contract (i.e. it has code) then we forward a call to the contract and verify that it returns the correct EIP-1271 "magic value".
         */
        address delegationApprover;
        /**
         * @notice A minimum delay -- measured in blocks -- enforced between:
         * 1) the operator signalling their intent to register for a service, via calling `Slasher.optIntoSlashing`
         * and
         * 2) the operator completing registration for the service, via the service ultimately calling `Slasher.recordFirstStakeUpdate`
         * @dev note that for a specific operator, this value *cannot decrease*, i.e. if the operator wishes to modify their OperatorDetails,
         * then they are only allowed to either increase this value or keep it the same.
         */
        uint32 stakerOptOutWindowBlocks;
    }

    // TODO: documentation
    struct SignatureWithExpiry {
        bytes signature;
        uint256 expiry;
    }

    // @notice Emitted when a new operator registers in EigenLayer and provides their OperatorDetails.
    event OperatorRegistered(address indexed operator, OperatorDetails operatorDetails);

    // @notice Emitted when an operator updates their OperatorDetails to @param newOperatorDetails
    event OperatorDetailsModified(address indexed operator, OperatorDetails newOperatorDetails);

    // @notice Emitted when @param staker delegates to @param operator.
    event StakerDelegated(address indexed staker, address indexed operator);

    // @notice Emitted when @param staker undelegates from @param operator.
    event StakerUndelegated(address indexed staker, address indexed operator);

    /**
     * @notice Registers the `msg.sender` as an operator in EigenLayer, that stakers can choose to delegate to.
     * @param registeringOperatorDetails is the `OperatorDetails` for the operator.
     * @dev Note that once an operator is registered, they cannot 'deregister' as an operator, and they will forever be considered "delegated to themself".
     * @dev This function will revert if the caller attempts to set their `earningsReceiver` to address(0).
     */
    function registerAsOperator(OperatorDetails calldata registeringOperatorDetails) external;

    /**
     * @notice Updates the `msg.sender`'s stored `OperatorDetails`.
     * @param newOperatorDetails is the updated `OperatorDetails` for the operator, to replace their current OperatorDetails`.
     * @dev The `msg.sender` must have previously registered as an operator in EigenLayer via calling the `registerAsOperator` function.
     * @dev This function will revert if the caller attempts to set their `earningsReceiver` to address(0).
     */
    function modifyOperatorDetails(OperatorDetails calldata newOperatorDetails) external;

    /**
     * @notice Called by a staker to delegate its assets to the @param operator.
     * @param operator is the operator to whom the staker (`msg.sender`) is delegating its assets for use in serving applications built on EigenLayer.
     * @param approverSignatureAndExpiry is a parameter that will be used for verifying that the operator approves of this delegation action in the event that:
     * 1) the operator's `delegationApprover` address is set to a non-zero value.
     * AND
     * 2) neither the operator nor their `delegationApprover` is the `msg.sender`, since in the event that the operator or their delegationApprover
     * is the `msg.sender`, then approval is assumed.
     * @dev In the event that `approverSignatureAndExpiry` is not checked, its content is ignored entirely; it's recommended to use an empty input
     * in this case to save on complexity + gas costs
     */
    function delegateTo(address operator, SignatureWithExpiry memory approverSignatureAndExpiry) external;

    /**
     * @notice Delegates from @param staker to @param operator.
     * @notice This function will revert if the current `block.timestamp` is equal to or exceeds @param expiry
     * @dev The @param stakerSignature is used as follows:
     * 1) If `staker` is an EOA, then `stakerSignature` is verified to be a valid ECDSA stakerSignature from `staker`, indicating their intention for this action.
     * 2) If `staker` is a contract, then `stakerSignature` will be checked according to EIP-1271.
     * @param approverSignatureAndExpiry is a parameter that will be used for verifying that the operator approves of this delegation action in the event that:
     * 1) the operator's `delegationApprover` address is set to a non-zero value.
     * AND
     * 2) neither the operator nor their `delegationApprover` is the `msg.sender`, since in the event that the operator or their delegationApprover
     * is the `msg.sender`, then approval is assumed.
     * @dev In the event that `approverSignatureAndExpiry` is not checked, its content is ignored entirely; it's recommended to use an empty input
     * in this case to save on complexity + gas costs
     */
    function delegateToBySignature(
        address staker,
        address operator,
        SignatureWithExpiry memory stakerSignatureAndExpiry,
        SignatureWithExpiry memory approverSignatureAndExpiry
    ) external;

    /**
     * @notice Undelegates `staker` from the operator who they are delegated to.
     * @notice Callable only by the StrategyManager
     * @dev Should only ever be called in the event that the `staker` has no active deposits in EigenLayer.
     * @dev Reverts if the `staker` is also an operator, since operators are not allowed to undelegate from themselves
     */
    function undelegate(address staker) external;

    /**
     * @notice Called by an operator's `delegationApprover` address, in order to forcibly undelegate a staker who is currently delegated to the operator.
     * @param operator The operator who the @param staker is currently delegated to.
     * @dev This function will revert if either:
     * A) The `msg.sender` does not match `operatorDetails(operator).delegationApprover`.
     * OR
     * B) The `staker` is not currently delegated to the `operator`.
     * @dev This function will also revert if the `staker` is the `operator`; operators are considered *permanently* delegated to themselves.
     * @return The root of the newly queued withdrawal.
     * @dev Note that it is assumed that a staker places some trust in an operator, in paricular for the operator to not get slashed; a malicious operator can use this function
     * to inconvenience a staker who is delegated to them, but the expectation is that the inconvenience is minor compared to the operator getting purposefully slashed.
     */
    function forceUndelegation(address staker, address operator) external returns (bytes32);

    /**
     * @notice *If the staker is actively delegated*, then increases the `staker`'s delegated shares in `strategy` by `shares`. Otherwise does nothing.
     * Called by the StrategyManager whenever new shares are added to a user's share balance.
     * @dev Callable only by the StrategyManager.
     */
    function increaseDelegatedShares(address staker, IStrategy strategy, uint256 shares) external;

    /**
     * @notice *If the staker is actively delegated*, then decreases the `staker`'s delegated shares in each entry of `strategies` by its respective `shares[i]`. Otherwise does nothing.
     * Called by the StrategyManager whenever shares are decremented from a user's share balance, for example when a new withdrawal is queued.
     * @dev Callable only by the StrategyManager.
     */
    function decreaseDelegatedShares(address staker, IStrategy[] calldata strategies, uint256[] calldata shares) external;

    /**
     * @notice returns the address of the operator that `staker` is delegated to.
     * @notice Mapping: staker => operator whom the staker is currently delegated to.
     * @dev Note that returning address(0) indicates that the staker is not actively delegated to any operator.
     */
    function delegatedTo(address staker) external view returns (address);

    /**
     * @notice returns the OperatorDetails of the `operator`.
     * @notice Mapping: operator => OperatorDetails struct
     */
    function operatorDetails(address operator) external view returns (OperatorDetails memory);

    // @notice Getter function for `_operatorDetails[operator].earningsReceiver`
    function earningsReceiver(address operator) external view returns (address);

    // @notice Getter function for `_operatorDetails[operator].delegationApprover`
    function delegationApprover(address operator) external view returns (address);

    // @notice Getter function for `_operatorDetails[operator].stakerOptOutWindowBlocks`
    function stakerOptOutWindowBlocks(address operator) external view returns (uint256);

    /**
     * @notice returns the total number of shares in `strategy` that are delegated to `operator`.
     * @notice Mapping: operator => strategy => total number of shares in the strategy delegated to the operator.
     */
    function operatorShares(address operator, IStrategy strategy) external view returns (uint256);

    /// @notice Returns 'true' if `staker` *is* actively delegated, and 'false' otherwise.
    function isDelegated(address staker) external view returns (bool);

    /// @notice Returns 'true' if `staker` is *not* actively delegated, and 'false' otherwise.
    function isNotDelegated(address staker) external view returns (bool);

    /// @notice Returns if an operator can be delegated to, i.e. the `operator` has previously called `registerAsOperator`.
    function isOperator(address operator) external view returns (bool);

    /// @notice Mapping: staker => number of signed delegation nonces (used in `delegateToBySignature`) from the staker that the contract has already checked
    function stakerNonce(address staker) external view returns (uint256);

    /**
     * @notice Mapping: operator => number of signed delegation nonces (used in `delegateTo` and `delegateToBySignature` if the operator
     * has specified a nonzero address as their `delegationApprover`)
     */
    function delegationApproverNonce(address operator) external view returns (uint256);

    /**
     * @notice External getter function that mirrors the staker signature hash calculation in the `delegateToBySignature` function
     * @param staker The signing staker
     * @param operator The operator who is being delegated
     * @param expiry The desired expiry time of the staker's signature
     */
    function calculateStakerDigestHash(address staker, address operator, uint256 expiry) external view returns (bytes32 stakerDigestHash);

    /**
     * @notice External getter function that mirrors the approver signature hash calculation in the `_delegate` function
     * @param operator The operator who is being delegated
     * @param expiry The desired expiry time of the approver's signature
     */
    function calculateApproverDigestHash(address operator, uint256 expiry) external view returns (bytes32 approverDigestHash);
}