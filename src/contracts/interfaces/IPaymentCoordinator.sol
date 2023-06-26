// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";


/**
 * @title Contract used to coordinate payments from AVSs to operators and in particular the subsequency splitting of earnings from operators to stakers
 * @author Layr Labs, Inc.
 */
interface IPaymentCoordinator {
    /**
     * @notice Struct used by AVSs when informing EigenLayer of a payment made to an operator, but that could be earned at least
     * in part with funds from stakers who have delegated to the operator
     */
    struct Payment {
        IERC20 token;
        address operator;
        uint256[] amounts;
        uint256[] quorums;
        uint256 startBlockNumber;
        uint256 endBlockNumber;
    }

    // @notice Struct used when posting new Merkle roots
    struct MerkleRootPost {
        // the actual root of the tree
        bytes32 root;
        // the height of the tree
        uint256 height;
        // the block number after which the Merkle root can be proved against (to have a delay)
        uint256 confirmedAtBlockNumber;
        // the block number up to which the payment is calculated. Should be an already-finalized block that is sufficiently in the past.
        uint256 calculatedUpToBlockNumber;
    }

    // @notice Struct used for leaves of posted Merkle trees
    struct MerkleLeaf {
        address recipient;
        IERC20[] tokens;
        // cumulative all-time earnings in each token
        uint256[] amounts;
    }

    /// @notice Getter function for the length of the `merkleRootPosts` array
    function merkleRootPostsLength() external view returns (uint256);


    /// @notice getter cumulativeTokenAmountClaimedByRecipient (mapping(IERC20 => mapping(address => uint256))
    function cumulativeTokenAmountClaimedByRecipient(IERC20 token, address recipient) external view returns (uint256);

    /// @notice getter for merkleRootPosts 
    function merkleRootPosts(uint256 index) external view returns (MerkleRootPost memory);

    /**
     * @notice Makes a payment of sum(amounts) paid in `token`, for `operator`'s contributions to an AVS,
     * between `startBlockNumber` (inclusive) and `endBlockNumber` (inclusive)
     * @dev Transfers the total payment from the `msg.sender` to this contract, so the caller must have previously approved
     * this contract to transfer at least sum(`amounts`) of `token`
     * @notice Emits a `PaymentReceived` event
     */
    function makePayment(Payment calldata payment) external;

    // @notice Permissioned function which allows posting a new Merkle root
    function postMerkleRoot(bytes32 newRoot, uint256 height, uint256 calculatedUpToBlockNumber) external;

    // @notice Permissioned function which allows withdrawal of EigenLayer's share of `token` from all received payments
    function withdrawEigenlayerShare(IERC20 token, address recipient) external;

    /**
    * @notice Called by a staker or operator to prove the inclusion of their earnings in a posted Merkle root and claim them.
    * @param proof Merkle proof showing that a leaf containing `(msg.sender, amount)` was included in the `rootIndex`-th
    * Merkle root posted for the `token`
    * @param rootIndex Specifies the Merkle root to look up, using `merkleRootsByToken[token][rootIndex]`
    * @param leaf The leaf to be inserted into the Merkle tree
    */
    function proveAndClaimEarnings(
        bytes memory proof,
        uint256 rootIndex,
        MerkleLeaf memory leaf,
        uint256 leafIndex
    ) external;

}