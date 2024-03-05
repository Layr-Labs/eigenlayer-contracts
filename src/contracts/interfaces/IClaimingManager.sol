// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title Interface for the `IClaimingManager` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
interface IClaimingManager {

    /// STRUCTS ///

    struct DistributionRoot {
        // timestamp after which the root can be claimed against
        uint32 activatedAfter;
        // merkle root of the distribution
        bytes32 root;
    }

    struct MerkleLeaf {
        IERC20 token; 
        uint256 amount;
        // Explicit recipient of the claim
        address recipient;
    }

    struct PaymentMerkleClaim {
        MerkleLeaf leaf;

        // The index of the root in the list of roots
        uint32 rootIndex;
        // The index of the leaf in the merkle tree for this root
        uint32 leafIndex;

        bytes proof;
    }

    /// EVENTS /// 

    event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater);
    event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay);
    event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips);
    event RecipientSet(address indexed account, address indexed recipient);
    event RootSubmitted(bytes32 root, uint32 paymentsCalculatedUntilTimestamp, uint32 activatedAfter);
    event PaymentClaimed(MerkleLeaf leaf);
    
    /// VIEW FUNCTIONS ///

    /// @notice The address to which payments are forwarded
    function paymentUpdater() external view returns (address);

    /// @notice Delay in timestamp before a posted root can be claimed against
    function activationDelay() external view returns (uint32);

    /// @notice Mapping: account => the address of the entity that can claim payments on behalf of the accounts
    function recipients(address account) external view returns (address);

    /// @notice Mapping: recipient => token => total amount claimed
    function cumulativeClaimed(address recipient, IERC20 token) external view returns (uint256);

    /// @notice the commission for all operators across all avss
    function globalCommissionBips() external view returns (uint16);

    /// @notice returns the hash of the leaf
    function findLeafHash(MerkleLeaf calldata leaf) external view returns (bytes32);

    /// @notice returns 'true' if the claim would currently pass the check in `processClaims`
    function checkClaim(PaymentMerkleClaim calldata claim) external view returns (bool);

    /// EXTERNAL FUNCTIONS ///

    /**
     * @notice Initializes the contract
     * @param initialOwner The address of the initial owner of the contract
     * @param _paymentUpdater The address of the entity that can update the payment merkle roots
     * @param _activationDelay Delay in timestamp before a posted root can be claimed against
     * @param _globalCommissionBips The commission for all operators across all avss
     * @dev Only callable once
     */
    function initialize(address initialOwner, address _paymentUpdater, uint32 _activationDelay, uint16 _globalCommissionBips) external;

    /**
     * @notice Sets the address of the entity that can update the payment merkle roots
     * @param _paymentUpdater The address of the entity that can update the payment merkle roots
     * @dev Only callable by the contract owner
     */
    function setPaymentUpdater(address _paymentUpdater) external;

    /**
     * @notice Sets the delay in timestamp before a posted root can be claimed against
     * @param _activationDelay Delay in timestamp before a posted root can be claimed against
     * @dev Only callable by the contract owner
     */
    function setActivationDelay(uint32 _activationDelay) external;

    /**
     * @notice Sets the global commission for all operators across all avss
     * @param _globalCommissionBips The commission for all operators across all avss
     * @dev Only callable by the contract owner
     */
    function setGlobalCommission(uint16 _globalCommissionBips) external;

    /**
     * @notice Sets the address of the entity that can claim payments on behalf of the account
     * @param account The account whose recipient is being set
     * @param recipient The address of the entity that can claim payments on behalf of the account
     * @dev Only callable by the account or their paymentRecipient
     */
    function setRecipient(address account, address recipient) external;

    /**
     * @notice Creates a new distribution root
     * @param root The merkle root of the distribution
     * @param paymentsCalculatedUntilTimestamp The timestamp until which payments have been calculated
     * @dev Only callable by the paymentUpdater
     */
    function submitRoot(bytes32 root, uint32 paymentsCalculatedUntilTimestamp) external;

    /**
     * @notice Claims payments for the given claims
     * @param claims The claims to be processed
     */
    function processClaims(PaymentMerkleClaim[] calldata claims) external;
}