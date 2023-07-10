// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../interfaces/IPaymentCoordinator.sol";
import "../libraries/Merkle.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";


/**
 * @title Contract used to coordinate payments from AVSs to operators and in particular the subsequent splitting of earnings from operators to stakers
 * @author Layr Labs, Inc.
 */
contract PaymentCoordinator is 
    IPaymentCoordinator,
    Initializable,
    OwnableUpgradeable
{
    using SafeERC20 for IERC20;

    /// @notice address approved to post new Merkle roots
    address public rootPublisher;

    /// @notice delay in blocks before a Merkle root can be activated
    uint256 public constant MERKLE_ROOT_ACTIVATION_DELAY_BLOCKS = 7200;

    /// @notice maximum BIPS
    uint256 public constant MAX_BIPS = 10000;

    /// @notice maximum BIPS for eigenlayer
    uint256 public constant MAX_EIGENLAYER_SHARE_BIPS = 1500;

    /// @notice Array of roots of posted Merkle trees, as well as associated data like tree height
    MerkleRootPost[] internal _merkleRootPosts;

    /// @notice Mapping token => recipient => cumulative amount *claimed*
    mapping(IERC20 => mapping(address => uint256)) public cumulativeTokenAmountClaimedByRecipient;

    /// @notice Mapping token => accumulated amount earned by EigenLayer
    mapping(IERC20 => uint256) public accumulatedEigenLayerTokenEarnings;

     /// @notice variable that defines the share EigenLayer takes of all payments, in basis points
     uint256 public eigenLayerShareBIPs;

    // TODO: better define this event?
    event PaymentReceived(address indexed receivedFrom, Payment payment);

    // @notice Emitted when a new Merkle root is posted
    event NewMerkleRootPosted(MerkleRootPost merkleRootPost);

    event PaymentClaimed(MerkleLeaf merkleLeaf);

    /// @notice Emitted when the rootPublisher is changed
    event RootPublisherChanged(address indexed oldRootPublisher, address indexed newRootPublisher);

    /// @notice Emitted when the merkle root activiation delay is changed
    event MerkleRootActivationDelayBlocksChanged(uint256 newMerkleRootActivationDelayBlocks);


    /// @notice Emitted when the EigenLayer's percentage share is changed
    event EigenLayerShareBIPSChanged(uint256 newEigenLayerShareBIPS);

    modifier onlyRootPublisher {
        require(msg.sender == rootPublisher, "PaymentCoordinator: Only rootPublisher");
        _;
    }

    constructor() {
        _disableInitializers();
    }

    function initialize(address _initialOwner, address _rootPublisher, uint256 _eigenlayerShareBips)
        external
        initializer
    {
        _transferOwnership(_initialOwner);
        _setRootPublisher(_rootPublisher);
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
        require(payment.amounts.length == payment.quorums.length, "PaymentCoordinator.makePayment: payment amounts and quorums must have the same length");
        uint256 sumAmounts;
        for (uint256 i = 0; i < payment.amounts.length; i++) {
            sumAmounts += payment.amounts[i];
        }
        payment.token.safeTransferFrom(msg.sender, address(this), sumAmounts);

        accumulatedEigenLayerTokenEarnings[payment.token] += sumAmounts * eigenLayerShareBIPs / MAX_BIPS;
        emit PaymentReceived(msg.sender, payment);
    }

    // @notice Permissioned function which allows posting a new Merkle root
    function postMerkleRoot(bytes32 newRoot, uint256 height, uint256 calculatedUpToBlockNumber) external onlyRootPublisher {
        MerkleRootPost memory newMerkleRoot = MerkleRootPost({
                                                root: newRoot, 
                                                height: height, 
                                                confirmedAtBlockNumber: block.number + MERKLE_ROOT_ACTIVATION_DELAY_BLOCKS, 
                                                calculatedUpToBlockNumber: calculatedUpToBlockNumber
                                            });
        _merkleRootPosts.push(newMerkleRoot);
        emit NewMerkleRootPosted(newMerkleRoot);
    }

    /// @notice Permissioned function which allows rootPublisher to nullify a Merkle root
    function nullifyMerkleRoot(uint256 rootIndex) external onlyOwner {
        require(block.number <= _merkleRootPosts[rootIndex].confirmedAtBlockNumber, "PaymentCoordinator.nullifyMerkleRoot: Merkle root already confirmed");
        delete _merkleRootPosts[rootIndex];
    }

    // @notice Permissioned function which allows withdrawal of EigenLayer's share of `token` from all received payments
    function withdrawEigenLayerShare(IERC20 token, address recipient) external onlyOwner {
        uint256 amount = accumulatedEigenLayerTokenEarnings[token];
        accumulatedEigenLayerTokenEarnings[token] = 0;
        token.safeTransfer(recipient, amount); 
    }

    /**
    * @notice Called by a staker or operator to prove the inclusion of their earnings in a posted Merkle root and claim them.
    * @param proof Merkle proof showing that a leaf was included in the `rootIndex`-th
    * Merkle root posted for the `token`
    * @param rootIndex Specifies the node inside the Merkle tree corresponding to the specified root, `merkleRoots[rootIndex].root`.
    */
    function proveAndClaimEarnings(
        bytes memory proof,
        uint256 rootIndex,
        MerkleLeaf memory leaf,
        uint256 leafIndex
    ) external {
        require(leaf.amounts.length == leaf.tokens.length, "PaymentCoordinator.proveAndClaimEarnings: leaf amounts and tokens must be same length");
        require(_merkleRootPosts[rootIndex].confirmedAtBlockNumber < block.number, "PaymentCoordinator.proveAndClaimEarnings: Merkle root not yet confirmed");

        bytes32 root = _merkleRootPosts[rootIndex].root;
        require(root != bytes32(0), "PaymentCoordinator.proveAndClaimEarnings: Merkle root is null");

        bytes32 leafHash = computeLeafHash(leaf);
        require(Merkle.verifyInclusionKeccak(proof, root, leafHash, leafIndex), "PaymentCoordinator.proveAndClaimEarnings: Invalid proof");

        for(uint256 i = 0; i < leaf.amounts.length; i++) {
            uint256 amount = leaf.amounts[i] - cumulativeTokenAmountClaimedByRecipient[leaf.tokens[i]][leaf.recipient];
            cumulativeTokenAmountClaimedByRecipient[leaf.tokens[i]][leaf.recipient] = leaf.amounts[i];
            leaf.tokens[i].safeTransfer(leaf.recipient, amount);
            
        }
        emit PaymentClaimed(leaf);
    }

    function setRootPublisher(address _rootPublisher) external onlyOwner {
        _setRootPublisher(_rootPublisher);
    }

    function setEigenLayerShareBIPS(uint256 _eigenlayerShareBips) external onlyOwner {
        _setEigenLayerShareBIPS(_eigenlayerShareBips);
    }


    /// @notice Getter function for the length of the `merkleRoots` array
    function merkleRootPostsLength() external view returns (uint256) {
        return _merkleRootPosts.length;
    }

    /// @notice Getter function for a merkleRoot at a given index
    function merkleRootPosts(uint256 index) external view returns (MerkleRootPost memory) {
        return _merkleRootPosts[index];
    }

    function _setRootPublisher(address _rootPublisher) internal {
        address currentRootPublisher = rootPublisher;
        rootPublisher = _rootPublisher;
        emit RootPublisherChanged(currentRootPublisher, rootPublisher);
    }

    function _setEigenLayerShareBIPS(uint256 _eigenlayerShareBips) internal {
        require(_eigenlayerShareBips <= MAX_EIGENLAYER_SHARE_BIPS, "PaymentCoordinator: EigenLayer share cannot be greater than 100%");
        eigenLayerShareBIPs = _eigenlayerShareBips;

        emit EigenLayerShareBIPSChanged(eigenLayerShareBIPs);
    }

    function computeLeafHash(MerkleLeaf memory leaf) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(leaf.recipient, keccak256(abi.encodePacked(leaf.tokens)), keccak256(abi.encodePacked(leaf.amounts))));
    }
}