// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../interfaces/IPaymentCoordinator.sol";
import "../libraries/Merkle.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "forge-std/Test.sol";



/**
 * @title Contract used to coordinate payments from AVSs to operators and in particular the subsequency splitting of earnings from operators to stakers
 * @author Layr Labs, Inc.
 */
contract PaymentCoordinator is 
    IPaymentCoordinator,
    Initializable,
    OwnableUpgradeable,
    Test
{
    using SafeERC20 for IERC20;

    /// @notice address approved to post new Merkle roots
    address public rootPublisher;

    uint256 public merkleRootActivationDelay;

    /// @notice maximum BIPS
    uint256 public constant MAX_BIPS = 10000;

    /// @notice Array of roots of posted Merkle trees, as well as associated data like tree height
    MerkleRootPost[] public merkleRootPosts;

    /// @notice Mapping token => recipient => cumulative amount *claimed*
    mapping(IERC20 => mapping(address => uint256)) public cumulativeTokenAmountClaimedByRecipient;

    /// @notice Mapping token => cumulative amount earned by EigenLayer
    mapping(IERC20 => uint256) public cumulativeEigenLayerTokeEarnings;

     /// @notice Constant that defines the share EigenLayer takes of all payments, in basis points
     uint256 public eigenLayerShareBIPs;

    // TODO: better define this event?
    event PaymentReceived(address indexed receivedFrom, Payment payment);

    // @notice Emitted when a new Merkle root is posted
    event NewMerkleRootPosted(MerkleRootPost merkleRootPost);

    event PaymentClaimed(MerkleLeaf merkleLeaf);

    modifier onlyRootPublisher {
        require(msg.sender == rootPublisher, "PaymentCoordinator: Only rootPublisher");
        _;
    }

    function initialize(address _initialOwner, address _rootPublisher, uint256 _eigenlayerShareBips, uint256 _merkleRootActivationDelay)
        external
        initializer
    {
        _transferOwnership(_initialOwner);
        _setRootPublisher(_rootPublisher);
        _setMerkleRootActivationDelay(_merkleRootActivationDelay);
        _setEigenLayerShareBIPS(_eigenlayerShareBips);
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

        cumulativeEigenLayerTokeEarnings[payment.token] += sumAmounts * eigenLayerShareBIPs / MAX_BIPS;
        emit PaymentReceived(msg.sender, payment);
    }

    // @notice Permissioned function which allows posting a new Merkle root
    function postMerkleRoot(bytes32 newRoot, uint256 height, uint256 calculatedUpToBlockNumber) external onlyRootPublisher{
        MerkleRootPost memory newMerkleRoot = MerkleRootPost(newRoot, height, block.number + merkleRootActivationDelay, calculatedUpToBlockNumber);
        merkleRootPosts.push(newMerkleRoot);
        emit NewMerkleRootPosted(newMerkleRoot);
    }

    // @notice Permissioned function which allows withdrawal of EigenLayer's share of `token` from all received payments
    function withdrawEigenlayerShare(IERC20 token, address recipient) external onlyRootPublisher{
        uint256 amount = cumulativeEigenLayerTokeEarnings[token];
        token.safeTransfer(recipient, amount);

        cumulativeEigenLayerTokeEarnings[token] = 0;
    }

    /**
    * @notice Called by a staker or operator to prove the inclusion of their earnings in a posted Merkle root and claim them.
    * @param proof Merkle proof showing that a leaf containing `(msg.sender, amount)` was included in the `rootIndex`-th
    * Merkle root posted for the `token`
    * @param rootIndex Specifies the node inside the Merkle tree corresponding to the specified root, `merkleRoots[rootIndex].root`.
    */
    function proveAndClaimEarnings(
        bytes memory proof,
        uint256 rootIndex,
        MerkleLeaf memory leaf
    ) external{
        require(leaf.amounts.length == leaf.tokens.length, "PaymentCoordinator.proveAndClaimEarnings: leaf amounts and tokens must be same length");
        require(merkleRootPosts[rootIndex].confirmedAtBlockNumber > block.number, "PaymentCoordinator.proveAndClaimEarnings: Merkle root not yet confirmed");
        bytes32 leafHash = keccak256(abi.encodePacked(leaf.recipient, keccak256(abi.encodePacked(leaf.tokens)), keccak256(abi.encodePacked(leaf.amounts)), leaf.index));
        bytes32 root = merkleRootPosts[rootIndex].root;
        require(root != bytes32(0), "PaymentCoordinator.proveAndClaimEarnings: Merkle root is null");
        require(Merkle.verifyInclusionKeccak(proof, root, leafHash, leaf.index), "PaymentCoordinator.proveAndClaimEarnings: Invalid proof");

        for(uint i = 0; i < leaf.amounts.length; i++) {
            leaf.tokens[i].safeTransfer(leaf.recipient, leaf.amounts[i]);
            cumulativeTokenAmountClaimedByRecipient[leaf.tokens[i]][leaf.recipient] += leaf.amounts[i];
        }
        emit PaymentClaimed(leaf);
    }

    function nullifyMerkleRoot(uint256 rootIndex) external onlyRootPublisher{
        require(block.number <= merkleRootPosts[rootIndex].confirmedAtBlockNumber, "PaymentCoordinator.nullifyMerkleRoot: Merkle root already confirmed");
        merkleRootPosts[rootIndex].root = bytes32(0);
    }

    function setRootPublisher(address _rootPublisher) external onlyOwner{
        _setRootPublisher(_rootPublisher);
    }

    function setMerkleRootActivationDelay(uint256 _merkleRootActivationDelay) external onlyOwner{
        _setMerkleRootActivationDelay(_merkleRootActivationDelay);
    }

    function setEigenLayerShareBIPS(uint256 _eigenlayerShareBips) external onlyOwner{
        _setEigenLayerShareBIPS(_eigenlayerShareBips);
    }


    /// @notice Getter function for the length of the `merkleRootPosts` array
    function merkleRootPostsLength() external view returns (uint256){
        return merkleRootPosts.length;
    }

    function _setRootPublisher(address _rootPublisher) internal {
        rootPublisher = _rootPublisher;
    }

    function _setMerkleRootActivationDelay(uint256 _merkleRootActivationDelay) internal {
        merkleRootActivationDelay = _merkleRootActivationDelay;
    }

    function _setEigenLayerShareBIPS(uint256 _eigenlayerShareBips) internal {
        require(_eigenlayerShareBips <= MAX_BIPS, "PaymentCoordinator: EigenLayer share cannot be greater than 100%");
        eigenLayerShareBIPs = _eigenlayerShareBips;
    }
}