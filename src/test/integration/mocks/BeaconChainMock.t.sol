// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/libraries/BeaconChainProofs.sol";
import "src/contracts/libraries/Merkle.sol";
import "src/contracts/pods/EigenPodManager.sol";

import "src/test/mocks/ETHDepositMock.sol";
import "src/test/integration/mocks/EIP_4788_Oracle_Mock.t.sol";
import "src/test/integration/mocks/EIP_7002_Mock.t.sol";
import "src/test/integration/mocks/EIP_7251_Mock.t.sol";
import "src/test/integration/mocks/LibValidator.t.sol";
import "src/test/integration/mocks/LibProofGen.t.sol";
import "src/test/utils/Logger.t.sol";

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

// Min/max balances for valdiators
// see https://github.com/ethereum/consensus-specs/blob/dev/specs/electra/beacon-chain.md#gwei-values
uint constant MAX_EFFECTIVE_BALANCE_WEI = 2048 ether;
uint64 constant MAX_EFFECTIVE_BALANCE_GWEI = 2048 gwei;
uint constant MIN_ACTIVATION_BALANCE_WEI = 32 ether;
uint64 constant MIN_ACTIVATION_BALANCE_GWEI = 32 gwei;

/// @notice A Pectra Beacon Chain Mock Contract. For testing upgrades, use BeaconChainMock_Upgradeable
/// @notice This mock assumed the following
/**
 * @notice A Semi-Compatible Pectra Beacon Chain Mock Contract. For Testing Upgrades to Pectra use BeaconChainMock_Upgradeable
 * @dev This mock assumes the following:
 * - Ceiling is Max EB, at which sweeps will be triggered
 * - No support for consolidations or any execution layer triggerable actions (exits, partial withdrawals)
 */
contract BeaconChainMock is Logger {
    using StdStyle for *;
    using print for *;
    using LibValidator for *;
    using LibProofGen for *;

    /// @dev The type of slash to apply to a validator
    enum SlashType {
        Minor, // `MINOR_SLASH_AMOUNT_GWEI`
        Half, // Half of the validator's balance
        Full // The validator's entire balance

    }

    // Rewards given to each validator during epoch processing
    uint64 public constant CONSENSUS_REWARD_AMOUNT_GWEI = 1;
    uint64 public constant MINOR_SLASH_AMOUNT_GWEI = 10;

    uint64 genesisTime;
    uint64 public nextTimestamp;

    EigenPodManager eigenPodManager;

    // Used to call predeploys as the system
    address constant SYSTEM_ADDRESS = 0xffffFFFfFFffffffffffffffFfFFFfffFFFfFFfE;

    /// Canonical beacon chain predeploy addresses
    IETHPOSDeposit constant DEPOSIT_CONTRACT = IETHPOSDeposit(0x00000000219ab540356cBB839Cbe05303d7705Fa);
    EIP_4788_Oracle_Mock constant EIP_4788_ORACLE = EIP_4788_Oracle_Mock(0x000F3df6D732807Ef1319fB7B8bB8522d0Beac02);
    EIP_7002_Mock constant WITHDRAWAL_PREDEPLOY = EIP_7002_Mock(payable(0x00000961Ef480Eb55e80D19ad83579A64c007002));
    EIP_7251_Mock constant CONSOLIDATION_PREDEPLOY = EIP_7251_Mock(payable(0x0000BBdDc7CE488642fb579F8B00f3a590007251));
    
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

    constructor(EigenPodManager _eigenPodManager, uint64 _genesisTime) {
        genesisTime = _genesisTime;
        eigenPodManager = _eigenPodManager;

        // Create mock 4788 oracle
        cheats.etch(address(DEPOSIT_CONTRACT), type(ETHPOSDepositMock).runtimeCode);
        cheats.etch(address(EIP_4788_ORACLE), type(EIP_4788_Oracle_Mock).runtimeCode);
        cheats.etch(address(CONSOLIDATION_PREDEPLOY), type(EIP_7251_Mock).runtimeCode);
        cheats.etch(address(WITHDRAWAL_PREDEPLOY), type(EIP_7002_Mock).runtimeCode);

        LibProofGen.usePectra();
    }

    function NAME() public pure virtual override returns (string memory) {
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
        address destination = v.withdrawalAddress();
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

    modifier onEpoch() {
        _updateCurrentEpoch();
        _;
        _advanceEpoch();
    }

    /// @dev Move forward one epoch on the beacon chain, taking care of important epoch processing:
    /// - Award ALL validators CONSENSUS_REWARD_AMOUNT
    /// - Process any pending partial withdrawals (new in Pectra)
    /// - Withdraw any balance over Max EB
    /// - Withdraw any balance for exited validators
    /// - Effective balances updated (NOTE: we do not use hysteresis!)
    /// - Move time forward one epoch
    /// - State root calculated and credential/balance proofs generated for all validators
    /// - Send state root to 4788 oracle
    ///
    /// Note:
    /// - DOES generate consensus rewards for ALL non-exited validators
    /// - DOES withdraw in excess of Max EB / if validator is exited
    function advanceEpoch() public onEpoch {
        print.method("advanceEpoch");

        _generateRewards();
        _processWithdrawals();
    }

    /// @dev Like `advanceEpoch`, but does NOT generate consensus rewards for validators.
    /// This amount is added to each validator's current balance before effective balances
    /// are updated.
    ///
    /// Note:
    /// - does NOT generate consensus rewards
    /// - DOES withdraw in excess of Max EB / if validator is exited
    function advanceEpoch_NoRewards() public onEpoch {
        print.method("advanceEpoch_NoRewards");

        _processWithdrawals();
    }

    /// @dev Like `advanceEpoch`, but explicitly does NOT withdraw if balances
    /// are over Max EB. This exists to support tests that check share increases solely
    /// due to beacon chain balance changes.
    ///
    /// Note:
    /// - DOES generate consensus rewards for ALL non-exited validators
    /// - does NOT withdraw in excess of Max EB
    /// - does NOT withdraw if validator is exited
    function advanceEpoch_NoWithdraw() public onEpoch {
        print.method("advanceEpoch_NoWithdraw");

        _generateRewards();
    }

    function advanceEpoch_NoWithdrawNoRewards() public onEpoch {
        print.method("advanceEpoch_NoWithdrawNoRewards");
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

    /// @dev Process EIP-7002 partial withdrawal requests and any withdrawals from the beacon chain
    function _processWithdrawals() internal {
        _process_EIP_7002_Requests();
        _process_EIP_7521_Requests();
        _withdrawFromBeaconChain();
    }

    /// @dev Handle consolidation requests
    function _process_EIP_7521_Requests() internal {
        // Call EIP-7521 predeploy and dequeue any consolidation requests
        cheats.prank(SYSTEM_ADDRESS);
        (bool ok, bytes memory data) = address(CONSOLIDATION_PREDEPLOY).call("");
        require(ok, "BeaconChainMock._process_EIP_7521_Requests: CONSOLIDATION_PREDEPLOY failed");

        ConsolidationReq[] memory requests = abi.decode(data, (ConsolidationReq[]));
        console.log("   - Dequeued %d consolidation requests.", requests.length);

        for (uint i = 0; i < requests.length; i++) {
            ConsolidationReq memory request = requests[i];
            uint40 sourceIndex = request.sourcePubkey.toIndex();
            uint40 targetIndex = request.targetPubkey.toIndex();

            if (sourceIndex >= validators.length || targetIndex >= validators.length) {
                _logSkip("source or target does not exist");
                continue;
            }
            
            Validator storage source = validators[sourceIndex];
            Validator storage target = validators[targetIndex];

            if (source.isDummy || target.isDummy) {
                _logSkip("dummy validator");
            } else if (request.source != source.withdrawalAddress()) {
                _logSkip("incorrect source address");
            } else if (source.hasBLSWC()) {
                _logSkip("source has BLS withdrawal credentials");
            } else if (!source.isActiveAt(currentEpoch())) {
                _logSkip("source is not active at current epoch");
            } else if (source.isExiting()) {
                _logSkip("source validator is exiting");
            } else if (sourceIndex == targetIndex) {
                // Handle switch to 0x02 request
                if (source.hasCompoundingWC()) {
                    _logSkip("switch request source already has 0x02 credentials");
                } else {
                    console.log("   -- Switching validator to 0x02 creds (idx: %d).", sourceIndex);
                
                    // The beacon chain would queue excess balance here as a "pending deposit", but
                    // we don't follow the spec that closely.
                    source.withdrawalCreds[0] = 0x02;
                }
            } else if (!target.hasCompoundingWC()) {
                _logSkip("target does not have compounding withdrawal credentials");
            } else if (!target.isActiveAt(currentEpoch())) {
                _logSkip("target is not active at current epoch");
            } else if (target.isExiting()) {
                _logSkip("target validator is exiting");
            } else if (source.pendingBalanceToWithdrawGwei != 0) {
                _logSkip("source has pending withdrawals in queue");
            } else if (source.isSlashed) {
                _logSkip("source validator is slashed");
            } else {
                console.log("   -- Consolidating source (%d) to target (%d).", sourceIndex, targetIndex);

                // Mark source validator exited
                source.exitEpoch = currentEpoch();

                // Transfer balance from source to target
                // Note the beacon chain would cap this transfer to the effective balance,
                // and withdraw the rest. Our effective balances aren't using hysteresis, so
                // we approximate this behavior by withdrawing anything over 32 ETH.
                uint64 transferAmtGwei = _currentBalanceGwei(sourceIndex);
                uint64 withdrawAmtGwei;
                if (transferAmtGwei > MIN_ACTIVATION_BALANCE_GWEI) {
                    withdrawAmtGwei = transferAmtGwei - MIN_ACTIVATION_BALANCE_GWEI;
                    transferAmtGwei = MIN_ACTIVATION_BALANCE_GWEI;
                }

                uint64 targetBalanceGwei = _currentBalanceGwei(targetIndex);
                _setCurrentBalance(sourceIndex, 0);
                _setCurrentBalance(targetIndex, targetBalanceGwei + transferAmtGwei);

                // Withdraw excess balance
                address destination = source.withdrawalAddress();
                cheats.deal(destination, address(destination).balance + uint(withdrawAmtGwei * GWEI_TO_WEI));
            }
        }
    }

    /// @dev Handle pending partial withdrawals AND withdraw excess validator balance
    function _process_EIP_7002_Requests() internal {
        // Call EIP-7002 predeploy and dequeue any withdrawal requests
        cheats.prank(SYSTEM_ADDRESS);
        (bool ok, bytes memory data) = address(WITHDRAWAL_PREDEPLOY).call("");
        require(ok, "BeaconChainMock._process_EIP_7002_Requests: WITHDRAWAL_PREDEPLOY failed");

        WithdrawalReq[] memory requests = abi.decode(data, (WithdrawalReq[]));
        console.log("   - Dequeued %d withdrawal requests.", requests.length);

        for (uint i = 0; i < requests.length; i++) {
            WithdrawalReq memory request = requests[i];
            uint40 validatorIndex = request.validatorPubkey.toIndex();

            if (validatorIndex >= validators.length) {
                _logSkip("validator does not exist");
                continue;
            }

            Validator storage v = validators[validatorIndex];

            bool isFullExitRequest = request.amountGwei == 0;

            uint64 balanceGwei = _currentBalanceGwei(validatorIndex);
            bool hasExcessBalance = balanceGwei > MIN_ACTIVATION_BALANCE_GWEI + v.pendingBalanceToWithdrawGwei;

            if (v.isDummy) {
                _logSkip("dummy validator");
            } else if (request.source != v.withdrawalAddress()) {
                _logSkip("incorrect source address");
            } else if (!v.isActiveAt(currentEpoch())) {
                _logSkip("inactive validator");
            } else if (v.isExiting()) {
                _logSkip("exit in progress");
            } else if (isFullExitRequest && v.pendingBalanceToWithdrawGwei != 0) {
                _logSkip("attempted full exit while pending withdrawal in queue");
            } else if (isFullExitRequest) {
                // TODO - swap to internal method
                exitValidator(validatorIndex);
            } else if (!v.hasCompoundingWC()) {
                _logSkip("attempted partial exit without 0x02 credentials");
            } else if (!hasExcessBalance) {
                _logSkip("validator does not have excess balance");
            } else {
                // Partial withdrawal - only withdraw down to 32 ETH
                uint64 toWithdrawGwei = balanceGwei - MIN_ACTIVATION_BALANCE_GWEI - v.pendingBalanceToWithdrawGwei;
                toWithdrawGwei = toWithdrawGwei > request.amountGwei ? request.amountGwei : toWithdrawGwei;

                v.pendingBalanceToWithdrawGwei += toWithdrawGwei;
            }
        }
    }

    function _logSkip(string memory reason) internal {
        console.log("   -- Skipping request with reason: %s.", reason);
    }

    /// @dev Iterate over all validators. If the validator:
    /// - has > than Max EB current balance
    /// - is exited
    /// - has a pending partial withdrawal
    ///
    /// ... withdraw the appropriate amount to the validator's withdrawal credentials
    function _withdrawFromBeaconChain() internal {
        uint64 totalWithdrawnGwei;
        for (uint i = 0; i < validators.length; i++) {
            Validator storage v = validators[i];
            if (v.isDummy) continue; // don't process dummy validators

            uint64 withdrawAmtGwei;
            uint64 curBalanceGwei = _currentBalanceGwei(uint40(i));
            // If the validator has nothing to withdraw, continue
            if (curBalanceGwei == 0) continue;

            if (v.exitEpoch != BeaconChainProofs.FAR_FUTURE_EPOCH) {
                // Process full withdrawal for exited validator
                withdrawAmtGwei = curBalanceGwei;
            } else {
                // If the validator has a pending withdrawal request, withdraw (but do not go below 32 eth)
                // Note that these reqs are only fulfilled for 0x02 validators (this is enforced elsewhere)
                if (v.pendingBalanceToWithdrawGwei != 0 && curBalanceGwei > MIN_ACTIVATION_BALANCE_GWEI) {
                    uint64 excessBalanceGwei = curBalanceGwei - MIN_ACTIVATION_BALANCE_GWEI;

                    withdrawAmtGwei = v.pendingBalanceToWithdrawGwei;
                    if (withdrawAmtGwei > excessBalanceGwei) withdrawAmtGwei = excessBalanceGwei;

                    // Reduce balance and set pending to 0
                    curBalanceGwei -= withdrawAmtGwei;
                    v.pendingBalanceToWithdrawGwei = 0;
                }
                 
                // Withdraw any amount above the max effective balance:
                // - For 0x01 validators, this is 32 ETH
                // - For 0x02 validators, this is 2048 ETH
                uint64 maxEBGwei = _getMaxEffectiveBalanceGwei(v);
                if (curBalanceGwei > maxEBGwei) {
                    uint64 excessBalanceGwei = curBalanceGwei - maxEBGwei;

                    withdrawAmtGwei += excessBalanceGwei;
                }
            }

            uint64 newBalanceGwei = curBalanceGwei - withdrawAmtGwei;

            // Send ETH to withdrawal address
            address destination = v.withdrawalAddress();
            uint withdrawAmtWei = withdrawAmtGwei * GWEI_TO_WEI;

            cheats.deal(destination, address(destination).balance + withdrawAmtWei);
            totalWithdrawnGwei += withdrawAmtGwei;

            // Update validator's current balance
            _setCurrentBalance(uint40(i), newBalanceGwei);
        }

        if (totalWithdrawnGwei != 0) console.log("- Total withdrawals from CL (gwei):", totalWithdrawnGwei);
    }

    function _updateCurrentEpoch() internal {
        uint64 curEpoch = currentEpoch();
        cheats.warp(_nextEpochStartTimestamp(curEpoch));
    }

    mapping(uint64 => StateProofs) proofs;

    function _advanceEpoch() internal virtual {
        cheats.pauseTracing();
        curTimestamp = uint64(block.timestamp);

        // Update effective balances for each validator
        for (uint i = 0; i < validators.length; i++) {
            Validator storage v = validators[i];
            if (v.isDummy) continue; // don't process dummy validators

            // Get current balance and trim anything over MAX EB
            uint64 balanceGwei = _currentBalanceGwei(uint40(i));
            uint64 maxEBGwei = _getMaxEffectiveBalanceGwei(v);
            if (balanceGwei > maxEBGwei) balanceGwei = maxEBGwei;

            v.effectiveBalanceGwei = balanceGwei;
        }

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

        bytes32 beaconBlockRoot = proofs[curTimestamp].generate({
            validators: validators,
            balances: balances
        });

        // Push new block root to oracle
        EIP_4788_ORACLE.setBlockRoot(curTimestamp, beaconBlockRoot);

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

            validators.push(
                Validator({
                    isDummy: true,
                    isSlashed: false,
                    pubkeyHash: validatorIndex.toPubkey().pubkeyHash(),
                    withdrawalCreds: new bytes(1),
                    effectiveBalanceGwei: dummyBalanceGwei,
                    activationEpoch: BeaconChainProofs.FAR_FUTURE_EPOCH,
                    exitEpoch: BeaconChainProofs.FAR_FUTURE_EPOCH,
                    pendingBalanceToWithdrawGwei: 0
                })
            );
            _setCurrentBalance(validatorIndex, dummyBalanceGwei);

            validatorIndex++;
        }

        validators.push(
            Validator({
                isDummy: false,
                isSlashed: false,
                pubkeyHash: validatorIndex.toPubkey().pubkeyHash(),
                withdrawalCreds: withdrawalCreds,
                effectiveBalanceGwei: balanceGwei,
                activationEpoch: currentEpoch(),
                exitEpoch: BeaconChainProofs.FAR_FUTURE_EPOCH,
                pendingBalanceToWithdrawGwei: 0
            })
        );
        _setCurrentBalance(validatorIndex, balanceGwei);

        cheats.resumeTracing();

        return validatorIndex;
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

    function totalEffectiveBalanceGwei(uint40[] memory validatorIndices) public view returns (uint64) {
        uint64 total;
        for (uint i = 0; i < validatorIndices.length; i++) {
            total += validators[validatorIndices[i]].effectiveBalanceGwei;
        }

        return total;
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

    /// @dev Opposite of BeaconChainProofs.getBalanceAtIndex, calculates a new balance
    /// root by updating the balance at validatorIndex
    /// @return The new, updated balance root
    function _calcBalanceRoot(bytes32 balanceRoot, uint40 validatorIndex, uint64 newBalanceGwei) internal pure returns (bytes32) {
        // Clear out old balance
        uint bitShiftAmount = 256 - (64 * ((validatorIndex % 4) + 1));
        uint mask = ~(uint(0xFFFFFFFFFFFFFFFF) << bitShiftAmount);
        uint clearedRoot = uint(balanceRoot) & mask;

        // Convert validator balance to little endian and shift to correct position
        uint leBalance = uint(newBalanceGwei.toLittleEndianUint64());
        uint shiftedBalance = leBalance >> (192 - bitShiftAmount);

        return bytes32(clearedRoot | shiftedBalance);
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

        StateProofs storage p = proofs[curTimestamp];

        CredentialProofs memory credentialProofs = CredentialProofs({
            beaconTimestamp: curTimestamp,
            stateRootProof: p.stateRootProof,
            validatorFieldsProofs: new bytes[](_validators.length),
            validatorFields: new bytes32[][](_validators.length)
        });

        // Get proofs for each validator
        for (uint i = 0; i < _validators.length; i++) {
            ValidatorFieldsProof memory proof = p.validatorFieldsProofs[_validators[i]];

            credentialProofs.validatorFieldsProofs[i] = proof.validatorFieldsProof;
            credentialProofs.validatorFields[i] = proof.validatorFields;
        }

        return credentialProofs;
    }

    function getCheckpointProofs(uint40[] memory _validators, uint64 timestamp) public view returns (CheckpointProofs memory) {
        // If we have not advanced an epoch since a validator was created, no proofs have been
        // generated for that validator. We check this here and revert early so we don't return
        // empty proofs.
        require(timestamp <= curTimestamp, "BeaconChain.getCheckpointProofs: future timestamp provided (did you call advanceEpoch yet?)");
        for (uint i = 0; i < _validators.length; i++) {
            require(
                _validators[i] <= lastIndexProcessed,
                "BeaconChain.getCheckpointProofs: no checkpoint proof found (did you call advanceEpoch yet?)"
            );
        }

        StateProofs storage p = proofs[timestamp];

        CheckpointProofs memory checkpointProofs = CheckpointProofs({
            balanceContainerProof: p.balanceContainerProof,
            balanceProofs: new BeaconChainProofs.BalanceProof[](_validators.length)
        });

        // Get proofs for each validator
        for (uint i = 0; i < _validators.length; i++) {
            uint40 validatorIndex = _validators[i];
            uint40 balanceRootIndex = _getBalanceRootIndex(validatorIndex);
            BalanceRootProof memory proof = p.balanceRootProofs[balanceRootIndex];

            checkpointProofs.balanceProofs[i] = BeaconChainProofs.BalanceProof({
                pubkeyHash: validators[validatorIndex].pubkeyHash,
                balanceRoot: proof.balanceRoot,
                proof: proof.proof
            });
        }

        return checkpointProofs;
    }

    function getStaleBalanceProofs(uint40 validatorIndex) public view returns (StaleBalanceProofs memory) {
        // If we have not advanced an epoch since a validator was created, no proofs have been
        // generated for that validator. We check this here and revert early so we don't return
        // empty proofs.
        require(validatorIndex <= lastIndexProcessed, "BeaconChain.getStaleBalanceProofs: no proof found (did you call advanceEpoch yet?)");

        StateProofs storage p = proofs[curTimestamp];

        ValidatorFieldsProof memory vfProof = p.validatorFieldsProofs[validatorIndex];
        return StaleBalanceProofs({
            beaconTimestamp: curTimestamp,
            stateRootProof: p.stateRootProof,
            validatorProof: BeaconChainProofs.ValidatorProof({validatorFields: vfProof.validatorFields, proof: vfProof.validatorFieldsProof})
        });
    }

    function balanceOfGwei(uint40 validatorIndex) public view returns (uint64) {
        return validators[validatorIndex].effectiveBalanceGwei;
    }

    function _getMaxEffectiveBalanceGwei(Validator storage v) internal virtual view returns (uint64) {
        return v.hasCompoundingWC() ? MAX_EFFECTIVE_BALANCE_GWEI : MIN_ACTIVATION_BALANCE_GWEI;
    }

    function hasCompoundingCreds(uint40 validatorIndex) public view returns (bool) {
        return validators[validatorIndex].hasCompoundingWC();
    }

    function pubkeyHash(uint40 validatorIndex) public view returns (bytes32) {
        return validators[validatorIndex].pubkeyHash;
    }

    function pubkey(uint40 validatorIndex) public pure returns (bytes memory) {
        return validatorIndex.toPubkey();
    }

    function getPubkeyHashes(uint40[] memory _validators) public view returns (bytes32[] memory) {
        bytes32[] memory pubkeyHashes = new bytes32[](_validators.length);

        for (uint i = 0; i < _validators.length; i++) {
            pubkeyHashes[i] = validators[_validators[i]].pubkeyHash;
        }

        return pubkeyHashes;
    }

    function isActive(uint40 validatorIndex) public view returns (bool) {
        return validators[validatorIndex].isActiveAt(currentEpoch());
    }
}
