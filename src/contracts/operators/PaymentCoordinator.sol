// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../interfaces/IPaymentCoordinator.sol";

/**
 * @title Contract used to coordinate payments from AVSs to operators and in particular the subsequency splitting of earnings from operators to stakers
 * @author Layr Labs, Inc.
 */
contract PaymentCoordinator is IPaymentCoordinator{

    /// @notice address approved to post new Merkle roots
    address public admin;

    uint256 public merkleRootActivationDelay;

    /// @notice maximum BIPS
    uint256 public constant MAX_BIPS = 10000;

    /// @notice Array of roots of posted Merkle trees, as well as associated data like tree height
    MerkleRootPost[] public merkleRootPosts;

    /// @notice Mapping token => recipient => cumulative amount *claimed*
    mapping(IERC20 => mapping(address => uint256)) public cumulativeTokenAmountClaimedByRecipient;

     /// @notice Constant that defines the share EigenLayer takes of all payments, in basis points
     uint256 public EIGENLAYER_SHARE_BIPS;

    // TODO: better define this event?
    event PaymentReceived(address indexed receivedFrom, Payment payment);

    // @notice Emitted when a new Merkle root is posted
    event NewMerkleRootPosted(MerkleRootPost merkleRootPost);

    modifier onlyAdmin(address sender) {
        require(sender == admin, "PaymentCoordinator: Only admin");
        _;
    }

    constructor(uint256 _eigenlayerShareBips, address _admin, uint256 _merkleRootActivationDelay) {
        require(_eigenlayerShareBips <= MAX_BIPS, "PaymentCoordinator: EigenLayer share cannot be greater than 100%");
        EIGENLAYER_SHARE_BIPS = _eigenlayerShareBips;
        admin = _admin;
        merkleRootActivationDelay = _merkleRootActivationDelay;
    }


    /**
     * @notice Makes a payment of sum(amounts) paid in `token`, for `operator`'s contributions to an AVS,
     * between `startBlockNumber` (inclusive) and `endBlockNumber` (inclusive)
     * @dev Transfers the total payment from the `msg.sender` to this contract, so the caller must have previously approved
     * this contract to transfer at least sum(`amounts`) of `token`
     * @notice Emits a `PaymentReceived` event
     */
    function makePayment(Payment calldata payment) external{
        uint256 sumAmounts;
        for (uint256 i = 0; i < payment.amounts.length; i++) {
            sumAmounts += payment.amounts[i];
        }
        payment.token.safeTransferFrom(msg.sender, address(this), sumAmounts);
        emit PaymentReceived(msg.sender, payment);
    }

    // @notice Permissioned function which allows posting a new Merkle root
    function postMerkleRoot(bytes32 newRoot, uint256 height, uint256 calculatedUpToBlockNumber) external onlyAdmin(msg.sender){
        MerkleRootPost memory newMerkleRoot = MerkleRootPost(newRoot, height, block.number + merkleRootActivationDelay, calculatedUpToBlockNumber);
        merkleRootPosts.push(newMerkleRoot);
        emit NewMerkleRootPosted(newMerkleRoot);
    }

    // @notice Permissioned function which allows withdrawal of EigenLayer's share of `token` from all received payments
    function withdrawEigenlayerShare(IERC20 token, address recipient) external{
        
    }

    /**
    * @notice Called by a staker or operator to prove the inclusion of their earnings in a posted Merkle root and claim them.
    * @param token ERC20 token to claim
    * @param amount The `amount` contained in the leaf of the Merkle tree to be proved against the specified Merkle root
    * @param proof Merkle proof showing that a leaf containing `(msg.sender, amount)` was included in the `rootIndex`-th
    * Merkle root posted for the `token`
    * @param nodeIndex Specifies the node inside the Merkle tree corresponding to the specified root, `merkleRoots[rootIndex].root`.
    * @param rootIndex Specifies the Merkle root to look up, using `merkleRootsByToken[token][rootIndex]`
    */
    function proveAndClaimEarnings(
        IERC20 token,
        uint256 amount,
        bytes memory proof,
        uint256 nodeIndex,
        uint256 rootIndex
    ) external;


    /// @notice Getter function for the length of the `merkleRootPosts` array
    function merkleRootPostsLength() external view returns (uint256){
        return merkleRootPosts.length;
    }
}