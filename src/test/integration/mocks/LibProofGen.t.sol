// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/libraries/BeaconChainProofs.sol";
import "src/contracts/libraries/Merkle.sol";

import "src/test/integration/mocks/LibValidator.t.sol";

struct ValidatorFieldsProof {
    bytes32[] validatorFields;
    bytes validatorFieldsProof;
}

struct BalanceRootProof {
    bytes32 balanceRoot;
    bytes proof;
}

struct Tree {
    mapping(bytes32 => bytes32) siblings;
    mapping(bytes32 => bytes32) parents;
}

struct MerkleTrees {
    Tree validatorTree;
    Tree balancesTree;
    Tree stateTree;
    Tree blockTree;
}

struct StateProofs {
    MerkleTrees trees;
    BeaconChainProofs.StateRootProof stateRootProof;
    BeaconChainProofs.BalanceContainerProof balanceContainerProof;
    mapping(uint40 => ValidatorFieldsProof) validatorFieldsProofs;
    mapping(uint40 => BalanceRootProof) balanceRootProofs;
}

struct Config {
    uint BEACON_STATE_TREE_HEIGHT;
    uint BEACON_STATE_FIELDS;
    uint BEACON_BLOCK_FIELDS;
    uint BLOCKROOT_PROOF_LEN;
    uint VAL_FIELDS_PROOF_LEN;
    uint BALANCE_CONTAINER_PROOF_LEN;
    uint BALANCE_PROOF_LEN;
}

library LibProofGen {
    using LibProofGen for *;
    using LibValidator for *;

    /**
     *
     *              CONSTANTS AND CONFIG
     *
     */
    bytes32 constant CONFIG_SLOT = keccak256("LibProofGen.config");

    function config() internal view returns (Config storage) {
        Config storage cfg;
        bytes32 _CONFIG_SLOT = CONFIG_SLOT;
        assembly {
            cfg.slot := _CONFIG_SLOT
        }

        return cfg;
    }

    /// @dev Dencun-specific config
    /// See (https://eth2book.info/capella/part3/containers/state/#beaconstate)
    function useDencun() internal {
        Config storage cfg = config();

        cfg.BEACON_STATE_TREE_HEIGHT = BeaconChainProofs.DENEB_BEACON_STATE_TREE_HEIGHT;
        cfg.BEACON_STATE_FIELDS = 32;
        cfg.BEACON_BLOCK_FIELDS = 5;
        cfg.BLOCKROOT_PROOF_LEN = 32 * BeaconChainProofs.BEACON_BLOCK_HEADER_TREE_HEIGHT;
        cfg.VAL_FIELDS_PROOF_LEN = 32 * ((BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1) + BeaconChainProofs.DENEB_BEACON_STATE_TREE_HEIGHT);
        cfg.BALANCE_CONTAINER_PROOF_LEN =
            32 * (BeaconChainProofs.BEACON_BLOCK_HEADER_TREE_HEIGHT + BeaconChainProofs.DENEB_BEACON_STATE_TREE_HEIGHT);
        cfg.BALANCE_PROOF_LEN = 32 * (BeaconChainProofs.BALANCE_TREE_HEIGHT + 1);
    }

    /// @dev Pectra-specific config
    /// See (see https://github.com/ethereum/consensus-specs/blob/dev/specs/electra/beacon-chain.md#beaconstate)
    function usePectra() internal {
        Config storage cfg = config();

        cfg.BEACON_STATE_TREE_HEIGHT = BeaconChainProofs.PECTRA_BEACON_STATE_TREE_HEIGHT;
        cfg.BEACON_STATE_FIELDS = 37;
        cfg.BEACON_BLOCK_FIELDS = 5;
        cfg.BLOCKROOT_PROOF_LEN = 32 * BeaconChainProofs.BEACON_BLOCK_HEADER_TREE_HEIGHT;
        cfg.VAL_FIELDS_PROOF_LEN = 32 * ((BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1) + BeaconChainProofs.PECTRA_BEACON_STATE_TREE_HEIGHT);
        cfg.BALANCE_CONTAINER_PROOF_LEN =
            32 * (BeaconChainProofs.BEACON_BLOCK_HEADER_TREE_HEIGHT + BeaconChainProofs.PECTRA_BEACON_STATE_TREE_HEIGHT);
        cfg.BALANCE_PROOF_LEN = 32 * (BeaconChainProofs.BALANCE_TREE_HEIGHT + 1);
    }

    /**
     *
     *             PROOFGEN: MAIN METHOD
     *
     */

    /// @dev This method is how our mock beacon chain provides EigenPods with valid proofs of beacon state.
    ///
    /// INPUT: the beacon state for the current timestamp (validators and balances)
    /// OUTPUT: the beacon block root calculated from the input. The mock beacon chain will inject this root
    /// into the EIP-4788 beacon block root oracle, as EigenPods query this oracle during proof verification.
    ///
    /// Proofs against the beacon block root are also generated and stored in StateProofs; these proofs can
    /// be fetched before calling an EigenPod method.
    ///
    /// This process is broken into 2 steps:
    ///
    /// 1. The merkle tree builder builds 4 merkle trees, one for each beacon state object we care about:
    ///
    ///         beaconBlockRoot
    ///                |                              (beacon block tree)
    ///         beaconStateRoot
    ///        /               \                      (beacon state tree)
    /// validatorContainerRoot, balanceContainerRoot
    ///          |                       |            (balances tree)
    ///          |              individual balances
    ///          |                                    (validators tree)
    /// individual validators
    ///
    /// Since the full beacon state is quite large, the merkle tree builder uses some hacks to generate
    /// only as much of the tree as we need for our tests. See `build` for details.
    ///
    /// 2. Once the merkle trees are built, we pre-generate proofs for EigenPod methods.
    function generate(StateProofs storage p, Validator[] memory validators, mapping(uint40 => bytes32) storage balances)
        internal
        returns (bytes32 beaconBlockRoot)
    {
        MerkleTrees storage trees = p.trees;

        // Build merkle tree for validators
        bytes32 validatorsRoot =
            trees.validatorTree.build({leaves: _getValidatorLeaves(validators), height: BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1});

        // Build merkle tree for current balances
        bytes32[] memory balanceLeaves = _getBalanceLeaves(balances, validators.length);
        bytes32 balanceContainerRoot = trees.balancesTree.build({leaves: balanceLeaves, height: BeaconChainProofs.BALANCE_TREE_HEIGHT + 1});

        // Build merkle tree for BeaconState
        bytes32 beaconStateRoot = trees.stateTree.build({
            leaves: _getBeaconStateLeaves(validatorsRoot, balanceContainerRoot),
            height: config().BEACON_STATE_TREE_HEIGHT
        });

        // Build merkle tree for BeaconBlock
        beaconBlockRoot = trees.blockTree.build({
            leaves: _getBeaconBlockLeaves(beaconStateRoot),
            height: BeaconChainProofs.BEACON_BLOCK_HEADER_TREE_HEIGHT
        });

        // Pre-generate proofs for EigenPod methods
        p.genStateRootProof(beaconStateRoot);
        p.genBalanceContainerProof(balanceContainerRoot);
        p.genCredentialProofs(validators);
        p.genBalanceProofs(balanceLeaves);

        return beaconBlockRoot;
    }

    /**
     *
     *              MERKLE TREE BUILDER
     *
     */

    /// @dev Builds a merkle tree in storage using the given leaves and height
    /// -- if the leaves given are not complete (i.e. the depth should have more leaves),
    ///    a pre-calculated zero-node is used to complete the tree.
    /// -- each pair of nodes is stored in `siblings`, and their parent in `parents`.
    ///    These mappings are used to build proofs for any individual leaf
    /// @return root the root of the merkle tree
    ///
    /// HACK: this sibling/parent method of tree construction relies on all passed-in leaves
    /// being unique, so that we don't overwrite siblings/parents. This is simple for trees
    /// like the validator tree, as each leaf is a validator's unique validatorFields.
    /// However, for the balances tree, the leaves may not be distinct. To get around this,
    /// BeaconChainMock._createValidator adds "dummy" validators every 4 validators created,
    /// with a unique balance value. This ensures each balance root is unique.
    function build(Tree storage tree, bytes32[] memory leaves, uint height) internal returns (bytes32 root) {
        for (uint depth = 0; depth < height; depth++) {
            uint newLength = (leaves.length + 1) / 2;
            bytes32[] memory newLeaves = new bytes32[](newLength);

            // Hash each pair of nodes in this level of the tree
            for (uint i = 0; i < newLength; i++) {
                uint leftIdx = 2 * i;
                uint rightIdx = leftIdx + 1;

                // Get left leaf
                bytes32 leftLeaf = leaves[leftIdx];

                // Calculate right leaf
                bytes32 rightLeaf;
                if (rightIdx < leaves.length) rightLeaf = leaves[rightIdx];
                else rightLeaf = _getFillerNode(depth);

                // Hash left and right
                bytes32 result = sha256(abi.encodePacked(leftLeaf, rightLeaf));
                newLeaves[i] = result;

                // Record results, used to generate individual proofs later:
                // Record left and right as siblings
                tree.siblings[leftLeaf] = rightLeaf;
                tree.siblings[rightLeaf] = leftLeaf;
                // Record the result as the parent of left and right
                tree.parents[leftLeaf] = result;
                tree.parents[rightLeaf] = result;
            }

            // Move up one level
            leaves = newLeaves;
        }

        require(leaves.length == 1, "LibProofGen.build: invalid tree somehow");
        return leaves[0];
    }

    /// @dev Fetch unique filler node for given depth
    function _getFillerNode(uint depth) internal view returns (bytes32) {
        bytes32 curNode = Merkle.merkleizeSha256(new bytes32[](8));

        for (uint i = 1; i < depth; i++) {
            curNode = sha256(abi.encodePacked(curNode, curNode));
        }

        return curNode;
    }

    /**
     *
     *              PROOF GENERATION
     *
     */

    /// @dev Generate global proof of beaconStateRoot -> beaconBlockRoot
    /// Used in verifyWithdrawalCredentials and verifyStaleBalance
    function genStateRootProof(StateProofs storage p, bytes32 beaconStateRoot) internal {
        bytes memory proof = new bytes(config().BLOCKROOT_PROOF_LEN);
        bytes32 curNode = beaconStateRoot;

        uint depth = 0;
        for (uint i = 0; i < BeaconChainProofs.BEACON_BLOCK_HEADER_TREE_HEIGHT; i++) {
            bytes32 sibling = p.trees.blockTree.siblings[curNode];

            // proof[j] = sibling;
            assembly {
                mstore(add(proof, add(32, mul(32, i))), sibling)
            }

            curNode = p.trees.blockTree.parents[curNode];
            depth++;
        }

        p.stateRootProof.beaconStateRoot = beaconStateRoot;
        p.stateRootProof.proof = proof;
    }

    /// @dev Generate global proof of balanceContainerRoot -> beaconStateRoot -> beaconBlockRoot
    /// Used in verifyCheckpointProofs
    function genBalanceContainerProof(StateProofs storage p, bytes32 balanceContainerRoot) internal {
        bytes memory proof = new bytes(config().BALANCE_CONTAINER_PROOF_LEN);
        bytes32 curNode = balanceContainerRoot;

        uint totalHeight = config().BALANCE_CONTAINER_PROOF_LEN / 32;
        uint depth = 0;
        for (uint i = 0; i < config().BEACON_STATE_TREE_HEIGHT; i++) {
            bytes32 sibling = p.trees.stateTree.siblings[curNode];

            // proof[j] = sibling;
            assembly {
                mstore(add(proof, add(32, mul(32, i))), sibling)
            }

            curNode = p.trees.stateTree.parents[curNode];
            depth++;
        }

        for (uint i = depth; i < totalHeight; i++) {
            bytes32 sibling = p.trees.blockTree.siblings[curNode];

            // proof[j] = sibling;
            assembly {
                mstore(add(proof, add(32, mul(32, i))), sibling)
            }

            curNode = p.trees.blockTree.parents[curNode];
            depth++;
        }

        p.balanceContainerProof.balanceContainerRoot = balanceContainerRoot;
        p.balanceContainerProof.proof = proof;
    }

    /// @dev Generate per-validator proofs of their validator -> validatorsRoot
    /// Used in verifyWithdrawalCredentials and verifyStaleBalance
    function genCredentialProofs(StateProofs storage p, Validator[] memory validators) internal {
        mapping(uint40 => ValidatorFieldsProof) storage vfProofs = p.validatorFieldsProofs;

        // Calculate credential proofs for each validator
        for (uint i = 0; i < validators.length; i++) {
            bytes memory proof = new bytes(config().VAL_FIELDS_PROOF_LEN);
            bytes32[] memory validatorFields = validators[i].getValidatorFields();
            bytes32 curNode = Merkle.merkleizeSha256(validatorFields);

            // Validator fields leaf -> validator container root
            uint depth = 0;
            for (uint j = 0; j < 1 + BeaconChainProofs.VALIDATOR_TREE_HEIGHT; j++) {
                bytes32 sibling = p.trees.validatorTree.siblings[curNode];

                // proof[j] = sibling;
                assembly {
                    mstore(add(proof, add(32, mul(32, j))), sibling)
                }

                curNode = p.trees.validatorTree.parents[curNode];
                depth++;
            }

            // Validator container root -> beacon state root
            for (uint j = depth; j < 1 + BeaconChainProofs.VALIDATOR_TREE_HEIGHT + config().BEACON_STATE_TREE_HEIGHT; j++) {
                bytes32 sibling = p.trees.stateTree.siblings[curNode];

                // proof[j] = sibling;
                assembly {
                    mstore(add(proof, add(32, mul(32, j))), sibling)
                }

                curNode = p.trees.stateTree.parents[curNode];
                depth++;
            }

            vfProofs[uint40(i)].validatorFields = validatorFields;
            vfProofs[uint40(i)].validatorFieldsProof = proof;
        }
    }

    /// @dev Generate per-balance root proofs of each balance root -> balanceContainerRoot
    /// Used in verifyCheckpointProofs
    function genBalanceProofs(StateProofs storage p, bytes32[] memory balanceLeaves) internal {
        mapping(uint40 => BalanceRootProof) storage brProofs = p.balanceRootProofs;

        // Calculate current balance proofs for each balance root
        for (uint i = 0; i < balanceLeaves.length; i++) {
            bytes memory proof = new bytes(config().BALANCE_PROOF_LEN);
            bytes32 balanceRoot = balanceLeaves[i];
            bytes32 curNode = balanceRoot;

            // Balance root leaf -> balances container root
            uint depth = 0;
            for (uint j = 0; j < 1 + BeaconChainProofs.BALANCE_TREE_HEIGHT; j++) {
                bytes32 sibling = p.trees.balancesTree.siblings[curNode];

                // proof[j] = sibling;
                assembly {
                    mstore(add(proof, add(32, mul(32, j))), sibling)
                }

                curNode = p.trees.balancesTree.parents[curNode];
                depth++;
            }

            brProofs[uint40(i)].balanceRoot = balanceRoot;
            brProofs[uint40(i)].proof = proof;
        }
    }

    /**
     *
     *              MERKLE LEAVES GETTERS
     *
     */

    /// @dev Get the leaves of the beacon block tree
    function _getBeaconBlockLeaves(bytes32 beaconStateRoot) internal view returns (bytes32[] memory) {
        bytes32[] memory leaves = new bytes32[](config().BEACON_BLOCK_FIELDS);

        // Pre-populate leaves with dummy values so sibling/parent tracking is correct
        for (uint i = 0; i < leaves.length; i++) {
            leaves[i] = bytes32(i + 1);
        }

        // Place beaconStateRoot into tree
        leaves[BeaconChainProofs.STATE_ROOT_INDEX] = beaconStateRoot;
        return leaves;
    }

    /// @dev Get the leaves of the beacon state tree
    function _getBeaconStateLeaves(bytes32 validatorsRoot, bytes32 balancesRoot) internal view returns (bytes32[] memory) {
        bytes32[] memory leaves = new bytes32[](config().BEACON_STATE_FIELDS);

        // Pre-populate leaves with dummy values so sibling/parent tracking is correct
        for (uint i = 0; i < leaves.length; i++) {
            leaves[i] = bytes32(i + 1);
        }

        // Place validatorsRoot and balancesRoot into tree
        leaves[BeaconChainProofs.VALIDATOR_CONTAINER_INDEX] = validatorsRoot;
        leaves[BeaconChainProofs.BALANCE_CONTAINER_INDEX] = balancesRoot;
        return leaves;
    }

    /// @dev Get the leaves of the validators merkle tree
    function _getValidatorLeaves(Validator[] memory validators) internal pure returns (bytes32[] memory) {
        bytes32[] memory leaves = new bytes32[](validators.length);

        // Place each validator's validatorFields into tree
        for (uint i = 0; i < validators.length; i++) {
            leaves[i] = Merkle.merkleizeSha256(validators[i].getValidatorFields());
        }

        return leaves;
    }

    /// @dev Get the leaves of the balances merkle tree
    function _getBalanceLeaves(mapping(uint40 => bytes32) storage balances, uint numValidators) internal view returns (bytes32[] memory) {
        // Each balance leaf is shared by 4 validators. This uses div_ceil
        // to calculate the number of balance leaves
        uint numBalanceRoots = numValidators == 0 ? 0 : ((numValidators - 1) / 4) + 1;

        // Place each validator's current balance into tree
        bytes32[] memory leaves = new bytes32[](numBalanceRoots);
        for (uint i = 0; i < leaves.length; i++) {
            leaves[i] = balances[uint40(i)];
        }

        return leaves;
    }
}
