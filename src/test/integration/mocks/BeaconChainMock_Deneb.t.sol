// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/mocks/BeaconChainMock.t.sol";

/// @notice A backwards-compatible BeaconChain Mock that updates containers & proofs from Deneb to Pectra
contract BeaconChainMock_DenebForkable is BeaconChainMock {
    using StdStyle for *;
    using print for *;

    // Denotes whether the beacon chain has been forked to Pectra
    bool isPectra;

    // The timestamp of the Pectra hard fork
    uint64 public pectraForkTimestamp;

    constructor(EigenPodManager _eigenPodManager, uint64 _genesisTime) BeaconChainMock(_eigenPodManager, _genesisTime) {
        /// DENEB SPECIFIC CONSTANTS (PROOF LENGTHS, FIELD SIZES):
        // see https://eth2book.info/capella/part3/containers/state/#beaconstate
        BEACON_STATE_FIELDS = 32;

        VAL_FIELDS_PROOF_LEN = 32 * ((BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1) + BeaconChainProofs.DENEB_BEACON_STATE_TREE_HEIGHT);

        BALANCE_CONTAINER_PROOF_LEN =
            32 * (BeaconChainProofs.BEACON_BLOCK_HEADER_TREE_HEIGHT + BeaconChainProofs.DENEB_BEACON_STATE_TREE_HEIGHT);

        MAX_EFFECTIVE_BALANCE_GWEI = 32 gwei;
        MAX_EFFECTIVE_BALANCE_WEI = 32 ether;
    }

    function NAME() public pure override returns (string memory) {
        return "BeaconChain_PectraForkable";
    }

    /**
     *
     *                             INTERNAL FUNCTIONS
     *
     */
    function _advanceEpoch() public override {
        cheats.pauseTracing();

        // Update effective balances for each validator
        for (uint i = 0; i < validators.length; i++) {
            Validator storage v = validators[i];
            if (v.isDummy) continue; // don't process dummy validators

            // Get current balance and trim anything over MaxEB
            uint64 balanceGwei = _currentBalanceGwei(uint40(i));
            if (balanceGwei > MAX_EFFECTIVE_BALANCE_GWEI) balanceGwei = MAX_EFFECTIVE_BALANCE_GWEI;

            v.effectiveBalanceGwei = balanceGwei;
        }

        // console.log("   Updated effective balances...".dim());
        // console.log("       timestamp:", block.timestamp);
        // console.log("       epoch:", currentEpoch());

        uint64 curEpoch = currentEpoch();
        cheats.warp(_nextEpochStartTimestamp(curEpoch));
        curTimestamp = uint64(block.timestamp);

        // console.log("   Jumping to next epoch...".dim());
        // console.log("       timestamp:", block.timestamp);
        // console.log("       epoch:", currentEpoch());

        // console.log("   Building beacon state trees...".dim());

        // Log total number of validators and number being processed for the first time
        if (validators.length > 0) {
            lastIndexProcessed = validators.length - 1;
        } else {
            // generate an empty root if we don't have any validators
            EIP_4788_ORACLE.setBlockRoot(curTimestamp, keccak256(""));

            // console.log("-- no validators; added empty block root");
            return;
        }

        // Build merkle tree for validators
        bytes32 validatorsRoot = _buildMerkleTree({
            leaves: _getValidatorLeaves(),
            treeHeight: BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1,
            tree: trees[curTimestamp].validatorTree
        });
        // console.log("-- validator container root", validatorsRoot);

        // Build merkle tree for current balances
        bytes32 balanceContainerRoot = _buildMerkleTree({
            leaves: _getBalanceLeaves(),
            treeHeight: BeaconChainProofs.BALANCE_TREE_HEIGHT + 1,
            tree: trees[curTimestamp].balancesTree
        });
        // console.log("-- balances container root", balanceContainerRoot);

        // Build merkle tree for BeaconState
        bytes32 beaconStateRoot = _buildMerkleTree({
            leaves: _getBeaconStateLeaves(validatorsRoot, balanceContainerRoot),
            treeHeight: getBeaconStateTreeHeight(),
            tree: trees[curTimestamp].stateTree
        });
        // console.log("-- beacon state root", beaconStateRoot);

        // Build merkle tree for BeaconBlock
        bytes32 beaconBlockRoot = _buildMerkleTree({
            leaves: _getBeaconBlockLeaves(beaconStateRoot),
            treeHeight: BeaconChainProofs.BEACON_BLOCK_HEADER_TREE_HEIGHT,
            tree: trees[curTimestamp].blockTree
        });

        // console.log("-- beacon block root", cheats.toString(beaconBlockRoot));

        // Push new block root to oracle
        EIP_4788_ORACLE.setBlockRoot(curTimestamp, beaconBlockRoot);

        // Pre-generate proofs to pass to EigenPod methods
        _genStateRootProof(beaconStateRoot);
        _genBalanceContainerProof(balanceContainerRoot);
        _genCredentialProofs();
        _genBalanceProofs();

        cheats.resumeTracing();
    }

    function _genCredentialProofs() internal override {
        mapping(uint40 => ValidatorFieldsProof) storage vfProofs = validatorFieldsProofs[curTimestamp];

        // Calculate credential proofs for each validator
        for (uint i = 0; i < validators.length; i++) {
            bytes memory proof = new bytes(VAL_FIELDS_PROOF_LEN);
            bytes32[] memory validatorFields = _getValidatorFields(uint40(i));
            bytes32 curNode = Merkle.merkleizeSha256(validatorFields);

            // Validator fields leaf -> validator container root
            uint depth = 0;
            for (uint j = 0; j < 1 + BeaconChainProofs.VALIDATOR_TREE_HEIGHT; j++) {
                bytes32 sibling = trees[curTimestamp].validatorTree.siblings[curNode];

                // proof[j] = sibling;
                assembly {
                    mstore(add(proof, add(32, mul(32, j))), sibling)
                }

                curNode = trees[curTimestamp].validatorTree.parents[curNode];
                depth++;
            }

            // Validator container root -> beacon state root
            for (uint j = depth; j < 1 + BeaconChainProofs.VALIDATOR_TREE_HEIGHT + getBeaconStateTreeHeight(); j++) {
                bytes32 sibling = trees[curTimestamp].stateTree.siblings[curNode];

                // proof[j] = sibling;
                assembly {
                    mstore(add(proof, add(32, mul(32, j))), sibling)
                }

                curNode = trees[curTimestamp].stateTree.parents[curNode];
                depth++;
            }

            vfProofs[uint40(i)].validatorFields = validatorFields;
            vfProofs[uint40(i)].validatorFieldsProof = proof;
        }
    }

    function _genBalanceContainerProof(bytes32 balanceContainerRoot) internal override {
        bytes memory proof = new bytes(BALANCE_CONTAINER_PROOF_LEN);
        bytes32 curNode = balanceContainerRoot;

        uint totalHeight = BALANCE_CONTAINER_PROOF_LEN / 32;
        uint depth = 0;
        for (uint i = 0; i < getBeaconStateTreeHeight(); i++) {
            bytes32 sibling = trees[curTimestamp].stateTree.siblings[curNode];

            // proof[j] = sibling;
            assembly {
                mstore(add(proof, add(32, mul(32, i))), sibling)
            }

            curNode = trees[curTimestamp].stateTree.parents[curNode];
            depth++;
        }

        for (uint i = depth; i < totalHeight; i++) {
            bytes32 sibling = trees[curTimestamp].blockTree.siblings[curNode];

            // proof[j] = sibling;
            assembly {
                mstore(add(proof, add(32, mul(32, i))), sibling)
            }

            curNode = trees[curTimestamp].blockTree.parents[curNode];
            depth++;
        }

        balanceContainerProofs[curTimestamp] =
            BeaconChainProofs.BalanceContainerProof({balanceContainerRoot: balanceContainerRoot, proof: proof});
    }

    /// @notice Forks the beacon chain to Pectra
    /// @dev Test battery should warp to the fork timestamp after calling this method
    function forkToPectra(uint64 _pectraForkTimestamp) public {
        // https://github.com/ethereum/consensus-specs/blob/dev/specs/electra/beacon-chain.md#beaconstate
        BEACON_STATE_FIELDS = 37;

        VAL_FIELDS_PROOF_LEN = 32 * ((BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1) + BeaconChainProofs.PECTRA_BEACON_STATE_TREE_HEIGHT);
        BALANCE_CONTAINER_PROOF_LEN =
            32 * (BeaconChainProofs.BEACON_BLOCK_HEADER_TREE_HEIGHT + BeaconChainProofs.PECTRA_BEACON_STATE_TREE_HEIGHT);

        MAX_EFFECTIVE_BALANCE_GWEI = 2048 gwei;
        MAX_EFFECTIVE_BALANCE_WEI = 2048 ether;

        isPectra = true;

        cheats.warp(_pectraForkTimestamp);
        pectraForkTimestamp = _pectraForkTimestamp;
    }

    function getBeaconStateTreeHeight() public view returns (uint) {
        return isPectra ? BeaconChainProofs.PECTRA_BEACON_STATE_TREE_HEIGHT : BeaconChainProofs.DENEB_BEACON_STATE_TREE_HEIGHT;
    }
}
