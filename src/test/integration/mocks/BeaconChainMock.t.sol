// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/libraries/BeaconChainProofs.sol";
import "src/contracts/libraries/Merkle.sol";
import "src/contracts/pods/EigenPodManager.sol";

import "src/test/mocks/ETHDepositMock.sol";
import "src/test/integration/mocks/EIP_4788_Oracle_Mock.t.sol";
import "src/test/utils/Logger.t.sol";

struct ValidatorFieldsProof {
    bytes32[] validatorFields;
    bytes validatorFieldsProof;
}

struct BalanceRootProof {
    bytes32 balanceRoot;
    bytes proof;
}

struct CheckpointProofs {
    BeaconChainProofs.BalanceContainerProof balanceContainerProof;
    BeaconChainProofs.BalanceProof[] balanceProofs;
}

struct CredentialProofs {
    uint64 beaconTimestamp;
    BeaconChainProofs.StateRootProof stateRootProof;
    bytes[] validatorFieldsProofs;
    bytes32[][] validatorFields;
}

struct StaleBalanceProofs {
    uint64 beaconTimestamp;
    BeaconChainProofs.StateRootProof stateRootProof;
    BeaconChainProofs.ValidatorProof validatorProof;
}

contract BeaconChainMock is Logger {
    using StdStyle for *;
    using print for *;

    struct Validator {
        bool isDummy;
        bool isSlashed;
        bytes32 pubkeyHash;
        bytes withdrawalCreds;
        uint64 effectiveBalanceGwei;
        uint64 activationEpoch;
        uint64 exitEpoch;
    }

    /// @dev The type of slash to apply to a validator
    enum SlashType {
        Minor, // `MINOR_SLASH_AMOUNT_GWEI`
        Half, // Half of the validator's balance
        Full // The validator's entire balance

    }

    /// @dev All withdrawals are processed with index == 0
    uint constant ZERO_NODES_LENGTH = 100;

    // Rewards given to each validator during epoch processing
    uint64 public constant CONSENSUS_REWARD_AMOUNT_GWEI = 1;
    uint64 public constant MINOR_SLASH_AMOUNT_GWEI = 10;

    /// PROOF CONSTANTS (PROOF LENGTHS, FIELD SIZES):

    // see https://eth2book.info/capella/part3/containers/state/#beaconstate
    uint constant BEACON_STATE_FIELDS = 32;
    // see https://eth2book.info/capella/part3/containers/blocks/#beaconblock
    uint constant BEACON_BLOCK_FIELDS = 5;

    uint immutable BLOCKROOT_PROOF_LEN = 32 * BeaconChainProofs.BEACON_BLOCK_HEADER_TREE_HEIGHT;
    uint immutable VAL_FIELDS_PROOF_LEN = 32 * ((BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1) + BeaconChainProofs.BEACON_STATE_TREE_HEIGHT);
    uint immutable BALANCE_CONTAINER_PROOF_LEN =
        32 * (BeaconChainProofs.BEACON_BLOCK_HEADER_TREE_HEIGHT + BeaconChainProofs.BEACON_STATE_TREE_HEIGHT);
    uint immutable BALANCE_PROOF_LEN = 32 * (BeaconChainProofs.BALANCE_TREE_HEIGHT + 1);

    uint64 genesisTime;
    uint64 public nextTimestamp;

    EigenPodManager eigenPodManager;
    IETHPOSDeposit constant DEPOSIT_CONTRACT = IETHPOSDeposit(0x00000000219ab540356cBB839Cbe05303d7705Fa);
    EIP_4788_Oracle_Mock constant EIP_4788_ORACLE = EIP_4788_Oracle_Mock(0x000F3df6D732807Ef1319fB7B8bB8522d0Beac02);

    /**
     * BeaconState containers, used for proofgen:
     * https://eth2book.info/capella/part3/containers/state/#beaconstate
     */

    // Validator container, references every validator created so far
    Validator[] validators;

    // Current balance container, keeps a balance for every validator
    //
    // Since balances are stored in groups of 4, it's easier to make
    // this a mapping rather than an array. If it were an array, its
    // length would be validators.length / 4;
    mapping(uint40 => bytes32) balances;

    /**
     * Generated proofs for each block timestamp:
     */

    // Maps block.timestamp -> calculated beacon block roots
    mapping(uint64 => bytes32) beaconBlockRoots;

    // Keeps track of the index of the last validator we've seen during epoch processing
    uint lastIndexProcessed;
    uint64 curTimestamp;

    // Maps block.timestamp -> beacon state root and proof
    mapping(uint64 => BeaconChainProofs.StateRootProof) stateRootProofs;

    // Maps block.timestamp -> balance container root and proof
    mapping(uint64 => BeaconChainProofs.BalanceContainerProof) balanceContainerProofs;

    // Maps block.timestamp -> validatorIndex -> credential proof for that timestamp
    mapping(uint64 => mapping(uint40 => ValidatorFieldsProof)) validatorFieldsProofs;

    // Maps block.timestamp -> balanceRootIndex -> balance proof for that timestamp
    mapping(uint64 => mapping(uint40 => BalanceRootProof)) balanceRootProofs;

    bytes32[] zeroNodes;

    constructor(EigenPodManager _eigenPodManager, uint64 _genesisTime) {
        genesisTime = _genesisTime;
        eigenPodManager = _eigenPodManager;

        // Create mock 4788 oracle
        cheats.etch(address(DEPOSIT_CONTRACT), type(ETHPOSDepositMock).runtimeCode);
        cheats.etch(address(EIP_4788_ORACLE), type(EIP_4788_Oracle_Mock).runtimeCode);

        // Calculate nodes of empty merkle tree
        bytes32 curNode = Merkle.merkleizeSha256(new bytes32[](8));
        zeroNodes = new bytes32[](ZERO_NODES_LENGTH);
        zeroNodes[0] = curNode;

        for (uint i = 1; i < zeroNodes.length; i++) {
            zeroNodes[i] = sha256(abi.encodePacked(curNode, curNode));
            curNode = zeroNodes[i];
        }
    }

    function NAME() public pure override returns (string memory) {
        return "BeaconChain";
    }

    /**
     *
     *                                 EXTERNAL METHODS
     *
     */

    /// @dev Creates a new validator by:
    /// - Creating the validator container
    /// - Setting their current/effective balance
    /// - Assigning them a new, unique index
    function newValidator(bytes memory withdrawalCreds) public payable returns (uint40) {
        print.method("newValidator");

        uint balanceWei = msg.value;

        // These checks mimic the checks made in the beacon chain deposit contract
        //
        // We sanity-check them here because this contract sorta acts like the
        // deposit contract and this ensures we only create validators that could
        // exist IRL
        require(balanceWei >= 1 ether, "BeaconChainMock.newValidator: deposit value too low");
        require(balanceWei % 1 gwei == 0, "BeaconChainMock.newValidator: value not multiple of gwei");
        uint depositAmount = balanceWei / GWEI_TO_WEI;
        require(depositAmount <= type(uint64).max, "BeaconChainMock.newValidator: deposit value too high");

        // Create new validator and return its unique index
        return _createValidator(withdrawalCreds, uint64(depositAmount));
    }

    /// @dev Initiate an exit by:
    /// - Updating the validator's exit epoch
    /// - Decreasing current balance to 0
    /// - Withdrawing the balance to the validator's withdrawal credentials
    /// NOTE that the validator's effective balance is unchanged until advanceEpoch is called
    /// @return exitedBalanceGwei The balance exited to the withdrawal address
    ///
    /// This partially mimics the beacon chain's behavior, which is:
    /// 1. when an exit is initiated, the validator's exit/withdrawable epochs are immediately set
    /// 2. in a future slot (as part of the withdrawal queue), the validator's current balance is set to 0
    ///    - at the end of this slot, the eth is withdrawn to the withdrawal credentials
    /// 3. when the epoch finalizes, the validator's effective balance is updated
    ///
    /// Because this mock beacon chain doesn't implement a withdrawal queue or per-slot processing,
    /// `exitValidator` combines steps 1 and 2 into this method.
    ///
    /// TODO we may need to advance a slot here to maintain the properties we want in startCheckpoint
    function exitValidator(uint40 validatorIndex) public returns (uint64 exitedBalanceGwei) {
        print.method("exitValidator");

        // Update validator.exitEpoch
        Validator storage v = validators[validatorIndex];
        require(!v.isDummy, "BeaconChainMock: attempting to exit dummy validator. We need those for proofgen >:(");
        require(v.exitEpoch == BeaconChainProofs.FAR_FUTURE_EPOCH, "BeaconChainMock: validator already exited");
        v.exitEpoch = currentEpoch() + 1;

        // Set current balance to 0
        exitedBalanceGwei = _currentBalanceGwei(validatorIndex);
        _setCurrentBalance(validatorIndex, 0);

        // Send current balance to pod
        address destination = _toAddress(validators[validatorIndex].withdrawalCreds);
        cheats.deal(destination, address(destination).balance + uint(uint(exitedBalanceGwei) * GWEI_TO_WEI));

        return exitedBalanceGwei;
    }

    function slashValidators(uint40[] memory _validators, SlashType _slashType) public returns (uint64 slashedBalanceGwei) {
        print.method("slashValidators");

        for (uint i = 0; i < _validators.length; i++) {
            uint40 validatorIndex = _validators[i];
            Validator storage v = validators[validatorIndex];
            require(!v.isDummy, "BeaconChainMock: attempting to exit dummy validator. We need those for proofgen >:(");

            // Mark slashed and initiate validator exit
            if (!v.isSlashed) {
                v.isSlashed = true;
                v.exitEpoch = currentEpoch() + 1;
            }

            // Calculate slash amount
            uint64 slashAmountGwei;
            uint64 curBalanceGwei = _currentBalanceGwei(validatorIndex);

            if (_slashType == SlashType.Minor) slashAmountGwei = MINOR_SLASH_AMOUNT_GWEI;
            else if (_slashType == SlashType.Half) slashAmountGwei = curBalanceGwei / 2;
            else if (_slashType == SlashType.Full) slashAmountGwei = curBalanceGwei;

            // Calculate slash amount
            if (slashAmountGwei > curBalanceGwei) {
                slashedBalanceGwei += curBalanceGwei;
                curBalanceGwei = 0;
            } else {
                slashedBalanceGwei += slashAmountGwei;
                curBalanceGwei -= slashAmountGwei;
            }

            // Decrease current balance (effective balance updated during epoch processing)
            _setCurrentBalance(validatorIndex, curBalanceGwei);

            console.log("   - Slashed validator %s by %s gwei", validatorIndex, slashAmountGwei);
        }

        return slashedBalanceGwei;
    }

    function slashValidators(uint40[] memory _validators, uint64 _slashAmountGwei) public {
        print.method("slashValidatorsAmountGwei");

        for (uint i = 0; i < _validators.length; i++) {
            uint40 validatorIndex = _validators[i];
            Validator storage v = validators[validatorIndex];
            require(!v.isDummy, "BeaconChainMock: attempting to exit dummy validator. We need those for proofgen >:(");

            // Mark slashed and initiate validator exit
            if (!v.isSlashed) {
                v.isSlashed = true;
                v.exitEpoch = currentEpoch() + 1;
            }

            // Calculate slash amount
            uint64 curBalanceGwei = _currentBalanceGwei(validatorIndex);

            // Calculate slash amount
            uint64 slashedAmountGwei;
            if (_slashAmountGwei > curBalanceGwei) {
                slashedAmountGwei = curBalanceGwei;
                _slashAmountGwei -= curBalanceGwei;
                curBalanceGwei = 0;
            } else {
                slashedAmountGwei = _slashAmountGwei;
                curBalanceGwei -= _slashAmountGwei;
            }

            // Decrease current balance (effective balance updated during epoch processing)
            _setCurrentBalance(validatorIndex, curBalanceGwei);

            console.log("   - Slashed validator %s by %s gwei", validatorIndex, slashedAmountGwei);
        }
    }

    /// @dev Move forward one epoch on the beacon chain, taking care of important epoch processing:
    /// - Award ALL validators CONSENSUS_REWARD_AMOUNT
    /// - Withdraw any balance over 32 ETH
    /// - Withdraw any balance for exited validators
    /// - Effective balances updated (NOTE: we do not use hysteresis!)
    /// - Move time forward one epoch
    /// - State root calculated and credential/balance proofs generated for all validators
    /// - Send state root to 4788 oracle
    ///
    /// Note:
    /// - DOES generate consensus rewards for ALL non-exited validators
    /// - DOES withdraw in excess of 32 ETH / if validator is exited
    function advanceEpoch() public {
        print.method("advanceEpoch");
        _generateRewards();
        _withdrawExcess();
        _advanceEpoch();
    }

    /// @dev Like `advanceEpoch`, but does NOT generate consensus rewards for validators.
    /// This amount is added to each validator's current balance before effective balances
    /// are updated.
    ///
    /// Note:
    /// - does NOT generate consensus rewards
    /// - DOES withdraw in excess of 32 ETH / if validator is exited
    function advanceEpoch_NoRewards() public {
        print.method("advanceEpoch_NoRewards");
        _withdrawExcess();
        _advanceEpoch();
    }

    /// @dev Like `advanceEpoch`, but explicitly does NOT withdraw if balances
    /// are over 32 ETH. This exists to support tests that check share increases solely
    /// due to beacon chain balance changes.
    ///
    /// Note:
    /// - DOES generate consensus rewards for ALL non-exited validators
    /// - does NOT withdraw in excess of 32 ETH
    /// - does NOT withdraw if validator is exited
    function advanceEpoch_NoWithdraw() public {
        print.method("advanceEpoch_NoWithdraw");
        _generateRewards();
        _advanceEpoch();
    }

    function advanceEpoch_NoWithdrawNoRewards() public {
        print.method("advanceEpoch_NoWithdrawNoRewards");
        _advanceEpoch();
    }

    /// @dev Iterate over all validators. If the validator is still active,
    /// add CONSENSUS_REWARD_AMOUNT_GWEI to its current balance
    function _generateRewards() internal {
        uint totalRewarded;
        for (uint i = 0; i < validators.length; i++) {
            Validator storage v = validators[i];
            if (v.isDummy) continue; // don't process dummy validators

            // If validator has not initiated exit, add rewards to their current balance
            if (v.exitEpoch == BeaconChainProofs.FAR_FUTURE_EPOCH) {
                uint64 balanceGwei = _currentBalanceGwei(uint40(i));
                balanceGwei += CONSENSUS_REWARD_AMOUNT_GWEI;
                totalRewarded++;

                _setCurrentBalance(uint40(i), balanceGwei);
            }
        }

        console.log("   - Generated rewards for %s of %s validators.", totalRewarded, validators.length);
    }

    /// @dev Iterate over all validators. If the validator has > 32 ETH current balance
    /// OR is exited, withdraw the excess to the validator's withdrawal address.
    function _withdrawExcess() internal {
        uint totalExcessWei;
        for (uint i = 0; i < validators.length; i++) {
            Validator storage v = validators[i];
            if (v.isDummy) continue; // don't process dummy validators

            uint balanceWei = uint(_currentBalanceGwei(uint40(i))) * GWEI_TO_WEI;
            address destination = _toAddress(v.withdrawalCreds);
            uint excessBalanceWei;
            uint64 newBalanceGwei = uint64(balanceWei / GWEI_TO_WEI);

            // If the validator has exited, withdraw any existing balance
            //
            // If the validator has > 32 ether, withdraw anything over that
            if (v.exitEpoch != BeaconChainProofs.FAR_FUTURE_EPOCH) {
                if (balanceWei == 0) continue;

                excessBalanceWei = balanceWei;
                newBalanceGwei = 0;
            } else if (balanceWei > 32 ether) {
                excessBalanceWei = balanceWei - 32 ether;
                newBalanceGwei = 32 gwei;
            }

            // Send ETH to withdrawal address
            cheats.deal(destination, address(destination).balance + excessBalanceWei);
            totalExcessWei += excessBalanceWei;

            // Update validator's current balance
            _setCurrentBalance(uint40(i), newBalanceGwei);
        }

        if (totalExcessWei != 0) console.log("- Withdrew excess balance:", totalExcessWei.asGwei());
    }

    function _advanceEpoch() public {
        cheats.pauseTracing();

        // Update effective balances for each validator
        for (uint i = 0; i < validators.length; i++) {
            Validator storage v = validators[i];
            if (v.isDummy) continue; // don't process dummy validators

            // Get current balance and trim anything over 32 ether
            uint64 balanceGwei = _currentBalanceGwei(uint40(i));
            if (balanceGwei > 32 gwei) balanceGwei = 32 gwei;

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
            treeHeight: BeaconChainProofs.BEACON_STATE_TREE_HEIGHT,
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

    /**
     *
     *                             INTERNAL FUNCTIONS
     *
     */
    function _createValidator(bytes memory withdrawalCreds, uint64 balanceGwei) internal returns (uint40) {
        cheats.pauseTracing();
        uint40 validatorIndex = uint40(validators.length);

        // HACK to make balance proofs work. Every 4 validators we create
        // a dummy validator with empty withdrawal credentials and a unique
        // balance value. This ensures that each balanceRoot is unique, which
        // allows our efficient beacon state builder to work.
        //
        // For more details on this hack see _buildMerkleTree
        if (validatorIndex % 4 == 0) {
            uint64 dummyBalanceGwei = type(uint64).max - uint64(validators.length);

            bytes memory _dummyPubkey = new bytes(48);
            assembly {
                mstore(add(48, _dummyPubkey), validatorIndex)
            }
            validators.push(
                Validator({
                    isDummy: true,
                    isSlashed: false,
                    pubkeyHash: sha256(abi.encodePacked(_dummyPubkey, bytes16(0))),
                    withdrawalCreds: "",
                    effectiveBalanceGwei: dummyBalanceGwei,
                    activationEpoch: BeaconChainProofs.FAR_FUTURE_EPOCH,
                    exitEpoch: BeaconChainProofs.FAR_FUTURE_EPOCH
                })
            );
            _setCurrentBalance(validatorIndex, dummyBalanceGwei);

            validatorIndex++;
        }

        // Use pubkey format from `EigenPod._calculateValidatorPubkeyHash`
        bytes memory _pubkey = new bytes(48);
        assembly {
            mstore(add(48, _pubkey), validatorIndex)
        }
        validators.push(
            Validator({
                isDummy: false,
                isSlashed: false,
                pubkeyHash: sha256(abi.encodePacked(_pubkey, bytes16(0))),
                withdrawalCreds: withdrawalCreds,
                effectiveBalanceGwei: balanceGwei,
                activationEpoch: currentEpoch(),
                exitEpoch: BeaconChainProofs.FAR_FUTURE_EPOCH
            })
        );
        _setCurrentBalance(validatorIndex, balanceGwei);

        cheats.resumeTracing();

        return validatorIndex;
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

    /// Timestamp -> merkle trees constructed at that timestamp
    /// Used to generate proofs
    mapping(uint64 => MerkleTrees) trees;

    /// @dev Builds a merkle tree using the given leaves and height
    /// -- if the leaves given are not complete (i.e. the depth should have more leaves),
    ///    a pre-calculated zero-node is used to complete the tree.
    /// -- each pair of nodes is stored in `siblings`, and their parent in `parents`.
    ///    These mappings are used to build proofs for any individual leaf
    /// @return The root of the merkle tree
    ///
    /// HACK: this sibling/parent method of tree construction relies on all passed-in leaves
    /// being unique, so that we don't overwrite siblings/parents. This is simple for trees
    /// like the validator tree, as each leaf is a validator's unique validatorFields.
    /// However, for the balances tree, the leaves may not be distinct. To get around this,
    /// _createValidator adds "dummy" validators every 4 validators created, with a unique
    /// balance value. This ensures each balance root is unique.
    function _buildMerkleTree(bytes32[] memory leaves, uint treeHeight, Tree storage tree) internal returns (bytes32) {
        for (uint depth = 0; depth < treeHeight; depth++) {
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
                else rightLeaf = _getZeroNode(depth);

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

        require(leaves.length == 1, "BeaconChainMock._buildMerkleTree: invalid tree somehow");
        return leaves[0];
    }

    function _genStateRootProof(bytes32 beaconStateRoot) internal {
        bytes memory proof = new bytes(BLOCKROOT_PROOF_LEN);
        bytes32 curNode = beaconStateRoot;

        uint depth = 0;
        for (uint i = 0; i < BeaconChainProofs.BEACON_BLOCK_HEADER_TREE_HEIGHT; i++) {
            bytes32 sibling = trees[curTimestamp].blockTree.siblings[curNode];

            // proof[j] = sibling;
            assembly {
                mstore(add(proof, add(32, mul(32, i))), sibling)
            }

            curNode = trees[curTimestamp].blockTree.parents[curNode];
            depth++;
        }

        stateRootProofs[curTimestamp] = BeaconChainProofs.StateRootProof({beaconStateRoot: beaconStateRoot, proof: proof});
    }

    function _genBalanceContainerProof(bytes32 balanceContainerRoot) internal {
        bytes memory proof = new bytes(BALANCE_CONTAINER_PROOF_LEN);
        bytes32 curNode = balanceContainerRoot;

        uint totalHeight = BALANCE_CONTAINER_PROOF_LEN / 32;
        uint depth = 0;
        for (uint i = 0; i < BeaconChainProofs.BEACON_STATE_TREE_HEIGHT; i++) {
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

    function _genCredentialProofs() internal {
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
            for (uint j = depth; j < 1 + BeaconChainProofs.VALIDATOR_TREE_HEIGHT + BeaconChainProofs.BEACON_STATE_TREE_HEIGHT; j++) {
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

    function _genBalanceProofs() internal {
        mapping(uint40 => BalanceRootProof) storage brProofs = balanceRootProofs[curTimestamp];

        // Calculate current balance proofs for each balance root
        uint numBalanceRoots = _numBalanceRoots();
        for (uint i = 0; i < numBalanceRoots; i++) {
            bytes memory proof = new bytes(BALANCE_PROOF_LEN);
            bytes32 balanceRoot = balances[uint40(i)];
            bytes32 curNode = balanceRoot;

            // Balance root leaf -> balances container root
            uint depth = 0;
            for (uint j = 0; j < 1 + BeaconChainProofs.BALANCE_TREE_HEIGHT; j++) {
                bytes32 sibling = trees[curTimestamp].balancesTree.siblings[curNode];

                // proof[j] = sibling;
                assembly {
                    mstore(add(proof, add(32, mul(32, j))), sibling)
                }

                curNode = trees[curTimestamp].balancesTree.parents[curNode];
                depth++;
            }

            brProofs[uint40(i)].balanceRoot = balanceRoot;
            brProofs[uint40(i)].proof = proof;
        }
    }

    function _getValidatorLeaves() internal view returns (bytes32[] memory) {
        bytes32[] memory leaves = new bytes32[](validators.length);

        // Place each validator's validatorFields into tree
        for (uint i = 0; i < validators.length; i++) {
            leaves[i] = Merkle.merkleizeSha256(_getValidatorFields(uint40(i)));
        }

        return leaves;
    }

    function _getBalanceLeaves() internal view returns (bytes32[] memory) {
        // Place each validator's current balance into tree
        bytes32[] memory leaves = new bytes32[](_numBalanceRoots());
        for (uint i = 0; i < leaves.length; i++) {
            leaves[i] = balances[uint40(i)];
        }

        return leaves;
    }

    function _numBalanceRoots() internal view returns (uint) {
        // Each balance leaf is shared by 4 validators. This uses div_ceil
        // to calculate the number of balance leaves
        return (validators.length == 0) ? 0 : ((validators.length - 1) / 4) + 1;
    }

    function _getBeaconStateLeaves(bytes32 validatorsRoot, bytes32 balancesRoot) internal pure returns (bytes32[] memory) {
        bytes32[] memory leaves = new bytes32[](BEACON_STATE_FIELDS);

        // Pre-populate leaves with dummy values so sibling/parent tracking is correct
        for (uint i = 0; i < leaves.length; i++) {
            leaves[i] = bytes32(i + 1);
        }

        // Place validatorsRoot and balancesRoot into tree
        leaves[BeaconChainProofs.VALIDATOR_CONTAINER_INDEX] = validatorsRoot;
        leaves[BeaconChainProofs.BALANCE_CONTAINER_INDEX] = balancesRoot;
        return leaves;
    }

    function _getBeaconBlockLeaves(bytes32 beaconStateRoot) internal pure returns (bytes32[] memory) {
        bytes32[] memory leaves = new bytes32[](BEACON_BLOCK_FIELDS);

        // Pre-populate leaves with dummy values so sibling/parent tracking is correct
        for (uint i = 0; i < leaves.length; i++) {
            leaves[i] = bytes32(i + 1);
        }

        // Place beaconStateRoot into tree
        leaves[BeaconChainProofs.STATE_ROOT_INDEX] = beaconStateRoot;
        return leaves;
    }

    function _currentBalanceGwei(uint40 validatorIndex) internal view returns (uint64) {
        return currentBalance(validatorIndex);
    }

    function currentEpoch() public view returns (uint64) {
        require(block.timestamp >= genesisTime, "BeaconChain.currentEpoch: current time is before genesis time");
        return uint64((block.timestamp - genesisTime) / BeaconChainProofs.SECONDS_PER_EPOCH);
    }

    /// @dev Returns the validator's exit epoch
    function exitEpoch(uint40 validatorIndex) public view returns (uint64) {
        return validators[validatorIndex].exitEpoch;
    }

    function totalEffectiveBalanceWei(uint40[] memory validatorIndices) public view returns (uint) {
        uint total;
        for (uint i = 0; i < validatorIndices.length; i++) {
            total += uint(validators[validatorIndices[i]].effectiveBalanceGwei * GWEI_TO_WEI);
        }

        return total;
    }

    /// @dev Returns the validator's effective balance, in gwei
    function effectiveBalance(uint40 validatorIndex) public view returns (uint64) {
        return validators[validatorIndex].effectiveBalanceGwei;
    }

    /// @dev Returns the validator's current balance, in gwei
    function currentBalance(uint40 validatorIndex) public view returns (uint64) {
        return BeaconChainProofs.getBalanceAtIndex(getBalanceRoot(validatorIndex), validatorIndex);
    }

    function getBalanceRoot(uint40 validatorIndex) public view returns (bytes32) {
        return balances[validatorIndex / 4];
    }

    function _getBalanceRootIndex(uint40 validatorIndex) internal pure returns (uint40) {
        return validatorIndex / 4;
    }

    function _getValidatorFields(uint40 validatorIndex) internal view returns (bytes32[] memory) {
        bytes32[] memory vFields = new bytes32[](8);
        Validator memory v = validators[validatorIndex];

        vFields[BeaconChainProofs.VALIDATOR_PUBKEY_INDEX] = v.pubkeyHash;
        vFields[BeaconChainProofs.VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX] = bytes32(v.withdrawalCreds);
        vFields[BeaconChainProofs.VALIDATOR_BALANCE_INDEX] = _toLittleEndianUint64(v.effectiveBalanceGwei);
        vFields[BeaconChainProofs.VALIDATOR_SLASHED_INDEX] = bytes32(abi.encode(v.isSlashed));
        vFields[BeaconChainProofs.VALIDATOR_ACTIVATION_EPOCH_INDEX] = _toLittleEndianUint64(v.activationEpoch);
        vFields[BeaconChainProofs.VALIDATOR_EXIT_EPOCH_INDEX] = _toLittleEndianUint64(v.exitEpoch);

        return vFields;
    }

    /// @dev Update the validator's current balance
    function _setCurrentBalance(uint40 validatorIndex, uint64 newBalanceGwei) internal {
        bytes32 balanceRoot = balances[validatorIndex / 4];
        balanceRoot = _calcBalanceRoot(balanceRoot, validatorIndex, newBalanceGwei);

        balances[validatorIndex / 4] = balanceRoot;
    }

    /// From EigenPod.sol
    function _nextEpochStartTimestamp(uint64 epoch) internal view returns (uint64) {
        return genesisTime + ((1 + epoch) * BeaconChainProofs.SECONDS_PER_EPOCH);
    }

    function _calcValProofIndex(uint40 validatorIndex) internal pure returns (uint) {
        return (BeaconChainProofs.VALIDATOR_CONTAINER_INDEX << (BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1)) | uint(validatorIndex);
    }

    function _calcBalanceProofIndex(uint40 balanceRootIndex) internal pure returns (uint) {
        return (BeaconChainProofs.BALANCE_CONTAINER_INDEX << (BeaconChainProofs.BALANCE_TREE_HEIGHT + 1)) | uint(balanceRootIndex);
    }

    function _getZeroNode(uint depth) internal view returns (bytes32) {
        require(depth < ZERO_NODES_LENGTH, "_getZeroNode: invalid depth");

        return zeroNodes[depth];
    }

    /// @dev Opposite of Endian.fromLittleEndianUint64
    function _toLittleEndianUint64(uint64 num) internal pure returns (bytes32) {
        uint lenum;

        // Rearrange the bytes from big-endian to little-endian format
        lenum |= uint((num & 0xFF) << 56);
        lenum |= uint((num & 0xFF00) << 40);
        lenum |= uint((num & 0xFF0000) << 24);
        lenum |= uint((num & 0xFF000000) << 8);
        lenum |= uint((num & 0xFF00000000) >> 8);
        lenum |= uint((num & 0xFF0000000000) >> 24);
        lenum |= uint((num & 0xFF000000000000) >> 40);
        lenum |= uint((num & 0xFF00000000000000) >> 56);

        // Shift the little-endian bytes to the end of the bytes32 value
        return bytes32(lenum << 192);
    }

    /// @dev Opposite of BeaconChainProofs.getBalanceAtIndex, calculates a new balance
    /// root by updating the balance at validatorIndex
    /// @return The new, updated balance root
    function _calcBalanceRoot(bytes32 balanceRoot, uint40 validatorIndex, uint64 newBalanceGwei) internal pure returns (bytes32) {
        // Clear out old balance
        uint bitShiftAmount = 256 - (64 * ((validatorIndex % 4) + 1));
        uint mask = ~(uint(0xFFFFFFFFFFFFFFFF) << bitShiftAmount);
        uint clearedRoot = uint(balanceRoot) & mask;

        // Convert validator balance to little endian and shift to correct position
        uint leBalance = uint(_toLittleEndianUint64(newBalanceGwei));
        uint shiftedBalance = leBalance >> (192 - bitShiftAmount);

        return bytes32(clearedRoot | shiftedBalance);
    }

    /// @dev Helper to convert 32-byte withdrawal credentials to an address
    function _toAddress(bytes memory withdrawalCreds) internal pure returns (address a) {
        bytes32 creds = bytes32(withdrawalCreds);
        uint160 mask = type(uint160).max;

        assembly {
            a := and(creds, mask)
        }
    }

    /**
     *
     *                               VIEW METHODS
     *
     */
    function getCredentialProofs(uint40[] memory _validators) public view returns (CredentialProofs memory) {
        // If we have not advanced an epoch since a validator was created, no proofs have been
        // generated for that validator. We check this here and revert early so we don't return
        // empty proofs.
        for (uint i = 0; i < _validators.length; i++) {
            require(
                _validators[i] <= lastIndexProcessed,
                "BeaconChain.getCredentialProofs: validator has not been included in beacon chain state (DID YOU CALL advanceEpoch YET?)"
            );
        }

        CredentialProofs memory proofs = CredentialProofs({
            beaconTimestamp: curTimestamp,
            stateRootProof: stateRootProofs[curTimestamp],
            validatorFieldsProofs: new bytes[](_validators.length),
            validatorFields: new bytes32[][](_validators.length)
        });

        // Get proofs for each validator
        for (uint i = 0; i < _validators.length; i++) {
            ValidatorFieldsProof memory proof = validatorFieldsProofs[curTimestamp][_validators[i]];
            proofs.validatorFieldsProofs[i] = proof.validatorFieldsProof;
            proofs.validatorFields[i] = proof.validatorFields;
        }

        return proofs;
    }

    function getCheckpointProofs(uint40[] memory _validators, uint64 timestamp) public view returns (CheckpointProofs memory) {
        // If we have not advanced an epoch since a validator was created, no proofs have been
        // generated for that validator. We check this here and revert early so we don't return
        // empty proofs.
        for (uint i = 0; i < _validators.length; i++) {
            require(
                _validators[i] <= lastIndexProcessed,
                "BeaconChain.getCredentialProofs: no checkpoint proof found (did you call advanceEpoch yet?)"
            );
        }

        CheckpointProofs memory proofs = CheckpointProofs({
            balanceContainerProof: balanceContainerProofs[timestamp],
            balanceProofs: new BeaconChainProofs.BalanceProof[](_validators.length)
        });

        // Get proofs for each validator
        for (uint i = 0; i < _validators.length; i++) {
            uint40 validatorIndex = _validators[i];
            uint40 balanceRootIndex = _getBalanceRootIndex(validatorIndex);
            BalanceRootProof memory proof = balanceRootProofs[timestamp][balanceRootIndex];

            proofs.balanceProofs[i] = BeaconChainProofs.BalanceProof({
                pubkeyHash: validators[validatorIndex].pubkeyHash,
                balanceRoot: proof.balanceRoot,
                proof: proof.proof
            });
        }

        return proofs;
    }

    function getStaleBalanceProofs(uint40 validatorIndex) public view returns (StaleBalanceProofs memory) {
        ValidatorFieldsProof memory vfProof = validatorFieldsProofs[curTimestamp][validatorIndex];
        return StaleBalanceProofs({
            beaconTimestamp: curTimestamp,
            stateRootProof: stateRootProofs[curTimestamp],
            validatorProof: BeaconChainProofs.ValidatorProof({validatorFields: vfProof.validatorFields, proof: vfProof.validatorFieldsProof})
        });
    }

    function balanceOfGwei(uint40 validatorIndex) public view returns (uint64) {
        return validators[validatorIndex].effectiveBalanceGwei;
    }

    function pubkeyHash(uint40 validatorIndex) public view returns (bytes32) {
        return validators[validatorIndex].pubkeyHash;
    }

    function pubkey(uint40 validatorIndex) public pure returns (bytes memory) {
        bytes memory _pubkey = new bytes(48);
        assembly {
            mstore(add(48, _pubkey), validatorIndex)
        }
        return _pubkey;
    }

    function getPubkeyHashes(uint40[] memory _validators) public view returns (bytes32[] memory) {
        bytes32[] memory pubkeyHashes = new bytes32[](_validators.length);

        for (uint i = 0; i < _validators.length; i++) {
            pubkeyHashes[i] = validators[_validators[i]].pubkeyHash;
        }

        return pubkeyHashes;
    }

    function isActive(uint40 validatorIndex) public view returns (bool) {
        return validators[validatorIndex].exitEpoch == BeaconChainProofs.FAR_FUTURE_EPOCH;
    }
}
