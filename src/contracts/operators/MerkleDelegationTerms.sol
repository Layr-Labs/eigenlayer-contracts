// SPDX-License-Identifier: BUSL-1.1

pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../interfaces/IDelegationTerms.sol";
import "../libraries/Merkle.sol";

/**
 * @title A 'Delegation Terms' contract that an operator can use to distribute earnings to stakers by periodically posting Merkle roots
 * @author Layr Labs, Inc.
 * @notice This contract specifies the delegation terms of a given operator. When a staker delegates its stake to an operator,
 * it has to agrees to the terms set in the operator's 'Delegation Terms' contract. Payments to an operator are routed through
 * their specified 'Delegation Terms' contract for subsequent distribution of earnings to individual stakers.
 * There are also hooks that call into an operator's DelegationTerms contract when a staker delegates to or undelegates from
 * the operator.
 * @dev This contract uses a system in which the operator posts roots of a *sparse Merkle tree*. Each leaf of the tree is expected
 * to contain the **cumulative** earnings of a staker. This will reduce the total number of actions that stakers who claim only rarely
 * have to take, while allowing stakers to claim their earnings as often as new Merkle roots are posted.
 */
contract MerkleDelegationTerms is Ownable, IDelegationTerms {
    using SafeERC20 for IERC20;

    struct TokenAndAmount {
        IERC20 token;
        uint256 amount;
    }

    struct MerkleRootAndTreeHeight {
        bytes32 root;
        uint256 height;
    }

    // sanity-check parameter on Merkle tree height
    uint256 internal constant MAX_HEIGHT = 256;

    /// @notice staker => token => cumulative amount *claimed*
    mapping(address => mapping(IERC20 => uint256)) public cumulativeClaimedByStakerOfToken;

    /// @notice Array of Merkle roots with heights, each posted by the operator (contract owner)
    MerkleRootAndTreeHeight[] public merkleRoots;

    // TODO: more events?
    event NewMerkleRootPosted(bytes32 newRoot, uint256 height);

    /**
     * @notice Used by the operator to withdraw tokens directly from this contract.
     * @param tokensAndAmounts ERC20 tokens to withdraw and the amount of each respective ERC20 token to withdraw.
     */ 
    function operatorWithdrawal(TokenAndAmount[] calldata tokensAndAmounts) external onlyOwner {
        uint256 tokensAndAmountsLength = tokensAndAmounts.length;
        for (uint256 i; i < tokensAndAmountsLength;) {
            tokensAndAmounts[i].token.safeTransfer(msg.sender, tokensAndAmounts[i].amount);
            cumulativeClaimedByStakerOfToken[msg.sender][tokensAndAmounts[i].token] += tokensAndAmounts[i].amount;
            unchecked {
                ++i;
            }
        }
    }

    /// @notice Used by the operator to post an updated root of the stakers' all-time earnings
    function postMerkleRoot(bytes32 newRoot, uint256 height) external onlyOwner {
        // sanity check
        require(height <= MAX_HEIGHT, "MerkleDelegationTerms.postMerkleRoot: height input too large");
        merkleRoots.push(
            MerkleRootAndTreeHeight({
                root: newRoot,
                height: height
            })
        );
        emit NewMerkleRootPosted(newRoot, height);
    }

    /** 
     * @notice Called by a staker to prove the inclusion of their earnings in a Merkle root (posted by the operator) and claim them.
     * @param tokensAndAmounts ERC20 tokens to withdraw and the amount of each respective ERC20 token to withdraw.
     * @param proof Merkle proof showing that a leaf containing `(msg.sender, tokensAndAmounts)` was included in the `rootIndex`-th
     * Merkle root posted by the operator.
     * @param nodeIndex Specifies the node inside the Merkle tree corresponding to the specified root, `merkleRoots[rootIndex].root`.
     * @param rootIndex Specifies the Merkle root to look up, using `merkleRoots[rootIndex]`
     */
    function proveEarningsAndWithdraw(
        TokenAndAmount[] calldata tokensAndAmounts,
        bytes memory proof,
        uint256 nodeIndex,
        uint256 rootIndex
    ) external {
        // calculate the leaf that the `msg.sender` is claiming
        bytes32 leafHash = calculateLeafHash(msg.sender, tokensAndAmounts);

        // verify that the proof length is appropriate for the chosen root
        require(proof.length == 32 * merkleRoots[rootIndex].height, "MerkleDelegationTerms.proveEarningsAndWithdraw: incorrect proof length");

        // check inclusion of the leafHash in the tree corresponding to `merkleRoots[rootIndex]`
        require(
            Merkle.verifyInclusionKeccak(
                proof,
                merkleRoots[rootIndex].root,
                leafHash,
                nodeIndex
            ),
            "MerkleDelegationTerms.proveEarningsAndWithdraw: proof of inclusion failed"
        );

        uint256 tokensAndAmountsLength = tokensAndAmounts.length;
        for (uint256 i; i < tokensAndAmountsLength;) {
            // calculate amount to send
            uint256 amountToSend = tokensAndAmounts[i].amount - cumulativeClaimedByStakerOfToken[msg.sender][tokensAndAmounts[i].token];

            if (amountToSend != 0) {
                // update claimed amount in storage
                cumulativeClaimedByStakerOfToken[msg.sender][tokensAndAmounts[i].token] = tokensAndAmounts[i].amount;

                // actually send the tokens
                tokensAndAmounts[i].token.safeTransfer(msg.sender, amountToSend);
            }
            unchecked {
                ++i;
            }
        }
    }

    /// @notice Helper function for calculating a leaf in a Merkle tree formatted as `(address staker, TokenAndAmount[] calldata tokensAndAmounts)`
    function calculateLeafHash(address staker, TokenAndAmount[] calldata tokensAndAmounts) public pure returns (bytes32) {
        return keccak256(abi.encode(staker, tokensAndAmounts));
    }

    // FUNCTIONS FROM INTERFACE
    function payForService(IERC20, uint256) external payable
    // solhint-disable-next-line no-empty-blocks
    {}

    /// @notice Hook for receiving new delegation   
    function onDelegationReceived(
        address,
        IStrategy[] memory,
        uint256[] memory
    ) external pure returns(bytes memory)
    // solhint-disable-next-line no-empty-blocks
    {}

    /// @notice Hook for withdrawing delegation   
    function onDelegationWithdrawn(
        address,
        IStrategy[] memory,
        uint256[] memory
    ) external pure returns(bytes memory)
    // solhint-disable-next-line no-empty-blocks
    {}
}