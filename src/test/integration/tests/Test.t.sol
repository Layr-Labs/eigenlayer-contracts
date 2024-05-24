// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

enum Network {
    LOCAL,
    HOLESKY,
    MAINNET
}

enum Version {
    LATEST,
    M2_0_2_5, // m2 + eigen + pod fix
    M2,       // m2 + eigen
    M1
}

contract Integration_Tester is IntegrationCheckUtils {

    using Strings for *;
    using StdStyle for *;

    string constant HEADER_DELIMITER = "========================================================================";
    string constant SECTION_DELIMITER = "======";

    // function setUp() public override {
    //     _select({
    //         network: Network.LOCAL,
    //         version: Version.LATEST
    //         // version: Version.0_2_5
    //     });

    //     super.setUp();
    // }

    // function test_Example() public {
    //     _config({
    //         _randomSeed: 0,
    //         _assetTypes: HOLDS_ETH,
    //         _userTypes: DEFAULT
    //     });

    //     // User staker = _newStaker(Assets.ETH);
    //     // User operator = _newOperator(Assets.NONE);

    //     User staker = _newRandomStaker();

    //     // staker.startValidators();        // start a random number of validators
    //     // staker.startNonFullValidators(); // start a random number of validators with < 32 ETH
    //     uint40[] memory validators = staker.startValidators(1);

    //     // staker.delegateTo(operator);
    //     // check_Delegation_State(staker, operator, ...);

    //     staker.verifyWithdrawalCredentials(validators);
    //     // check_VerifyWC_State(staker, validators);
    //     // assert_Snap_Added_OperatorShares(operator, BEACONCHAIN_ETH_STRAT, )
        
    //     // staker.startCheckpoint();
    //     // check_StartCheckpoint_State(staker, validators);

    //     // staker.completeCheckpoint();
    //     // check_CompleteCheckpoint_State(staker, validators);

    //     // beaconChain.rewardValidators();
        

    // }

    function test_Thing() public {
        _configRand({
            _randomSeed: 0,
            _assetTypes: HOLDS_ETH,
            _userTypes: DEFAULT
        });

        (User staker, ,) = _newRandomStaker();

        uint numValidators = 10;
        uint40[] memory validators = staker.startValidators(numValidators);
        _logPod(staker);

        IStrategy[] memory strats = new IStrategy[](1);
        uint[] memory shares = new uint[](1);
        strats[0] = BEACONCHAIN_ETH_STRAT;
        shares[0] = 32 ether * numValidators;

        staker.verifyWithdrawalCredentials(validators);
        assert_Snap_Added_StakerShares(staker, strats, shares, "should have added shares");
        _logPod(staker);

        // Move forward one epoch and distribute consensus rewards to all validators
        // These rewards are not withdrawn to the execution layer!
        // beaconChain.advanceEpochWithRewards({ withdraw: false });

        staker.startCheckpoint();
        _logPod(staker);
        // check_StartCheckpoint_State(staker, validators, 1 gwei);

        staker.completeCheckpoint();
        _logPod(staker);
    }

    function test_Thing2() public {
        _configRand({
            _randomSeed: 0,
            _assetTypes: HOLDS_ETH,
            _userTypes: DEFAULT
        });

        (User staker, ,) = _newRandomStaker();
        cheats.pauseGasMetering();

        uint numValidators = 5;
        uint40[] memory validators = staker.startValidators(numValidators);
        _logPod(staker);

        IStrategy[] memory strats = new IStrategy[](1);
        uint[] memory shares = new uint[](1);
        strats[0] = BEACONCHAIN_ETH_STRAT;
        shares[0] = 32 ether * numValidators;

        CredentialProofs memory proofs = beaconChain.genCredentialProofs(validators);

        cheats.startPrank(address(staker));
        EigenPod pod = staker.pod();

        cheats.resumeGasMetering();
        uint gasBefore = gasleft();
        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });
        uint gasAfter = gasleft();
        uint gasConsumed = gasBefore - gasAfter;
        cheats.pauseGasMetering();

        emit log_named_uint("credential proofs for num validators", validators.length);
        emit log_named_uint("consumed gas", gasConsumed);

        pod.startCheckpoint(false);

        CheckpointProofs memory proofs2 = beaconChain.genCheckpointProofs(validators);

        cheats.resumeGasMetering();
        gasBefore = gasleft();
        pod.verifyCheckpointProofs({
            stateRootProof: proofs2.stateRootProof,
            proofs: proofs2.balanceProofs
        });
        gasAfter = gasleft();
        gasConsumed = gasBefore - gasAfter;
        cheats.pauseGasMetering();

        emit log_named_uint("checkpoint proofs for num validators", validators.length);
        emit log_named_uint("consumed gas", gasConsumed);

        // staker.verifyWithdrawalCredentials(validators);
        // assert_Snap_Added_StakerShares(staker, strats, shares, "should have added shares");
        // _logPod(staker);

        // // Move forward one epoch and distribute consensus rewards to all validators
        // // These rewards are not withdrawn to the execution layer!
        // // beaconChain.advanceEpochWithRewards({ withdraw: false });

        // staker.startCheckpoint();
        // _logPod(staker);
        // // check_StartCheckpoint_State(staker, validators, 1 gwei);

        // staker.completeCheckpoint();
        // _logPod(staker);
    }

    function _logValidator(uint40 validatorIndex) internal {
        emit log_named_uint("validator", validatorIndex);
        emit log_named_uint("current balance", beaconChain.currentBalance(validatorIndex));
        emit log_named_uint("effective balance", beaconChain.effectiveBalance(validatorIndex));
        emit log_named_uint("exit epoch", beaconChain.exitEpoch(validatorIndex));

        emit log("===");
    }

    function _logPod(User staker) internal {
        EigenPod pod = staker.pod();

        _logSection(string.concat(staker.NAME().cyan(), ": Pod Details"));

        _log("- hasRestaked", pod.hasRestaked());
        _log("- podOwnerShares", eigenPodManager.podOwnerShares(address(staker)));
        _log("- activeValidatorCount", pod.activeValidatorCount());
        _logU64("- withdrawableRestakedELGwei", pod.withdrawableRestakedExecutionLayerGwei());

        bool hasCheckpoint = pod.currentCheckpointTimestamp() != 0;
        _log("- has checkpoint", hasCheckpoint);
        if (hasCheckpoint) {
            IEigenPod.Checkpoint memory checkpoint = pod.currentCheckpoint();
            _log("-- beaconBlockRoot", checkpoint.beaconBlockRoot);
            _log("-- podBalanceGwei", checkpoint.podBalanceGwei);
            _log("-- balanceDeltasGwei", checkpoint.balanceDeltasGwei);
            _log("-- proofsRemaining", checkpoint.proofsRemaining);
        }

        _logSection("");
    }

    function _logSection(string memory name) internal {
        emit log(string.concat(
            SECTION_DELIMITER,
            name,
            SECTION_DELIMITER
        ));
    }

    function _log(string memory name, bool value) internal {
        emit log_named_string(name, value ? "true".green() : "false".magenta());
    }

    function _log(string memory name, bytes32 value) internal {
        emit log_named_string(name, value.dimBytes32());
    }

    function _log(string memory name, uint value) internal {
        emit log_named_uint(name, value);
    }

    function _logU64(string memory name, uint64 value) internal {
        emit log_named_uint(name, value);
    }

    function _log(string memory name, int value) internal {
        emit log_named_int(name, value);
    }

    // function test_Thing() public {
    //     uint40 aIdx = beaconChain.newValidator2(1 ether, "");

    //     bytes32 balanceRoot = beaconChain.getBalanceRoot(aIdx);
    //     emit log_named_bytes32("balance root a 1", balanceRoot);
    //     uint64 balanceA = beaconChain.getBalance(aIdx);
    //     emit log_named_uint("balance a 1", balanceA);

    //     uint40 bIdx = beaconChain.newValidator2(2 ether, "");

    //     balanceRoot = beaconChain.getBalanceRoot(bIdx);
    //     emit log_named_bytes32("balance root ab 12", balanceRoot);
    //     balanceA = beaconChain.getBalance(aIdx);
    //     emit log_named_uint("balance a 1", balanceA);
    //     uint64 balanceB = beaconChain.getBalance(bIdx);
    //     emit log_named_uint("balance b 2", balanceB);

    //     uint40 cIdx = beaconChain.newValidator2(3 ether, "");

    //     balanceRoot = beaconChain.getBalanceRoot(cIdx);
    //     emit log_named_bytes32("balance root abc 123", balanceRoot);
    //     balanceA = beaconChain.getBalance(aIdx);
    //     emit log_named_uint("balance a 1", balanceA);
    //     balanceB = beaconChain.getBalance(bIdx);
    //     emit log_named_uint("balance b 2", balanceB);
    //     uint64 balanceC = beaconChain.getBalance(cIdx);
    //     emit log_named_uint("balance c 3", balanceC);

    //     uint40 dIdx = beaconChain.newValidator2(4 ether, "");

    //     balanceRoot = beaconChain.getBalanceRoot(dIdx);
    //     emit log_named_bytes32("balance root abc 1234", balanceRoot);
    //     balanceA = beaconChain.getBalance(aIdx);
    //     emit log_named_uint("balance a 1", balanceA);
    //     balanceB = beaconChain.getBalance(bIdx);
    //     emit log_named_uint("balance b 2", balanceB);
    //     balanceC = beaconChain.getBalance(cIdx);
    //     emit log_named_uint("balance c 3", balanceC);
    //     uint64 balanceD = beaconChain.getBalance(dIdx);
    //     emit log_named_uint("balance d 4", balanceD);

    //     uint40 eIdx = beaconChain.newValidator2(5 ether, "");

    //     balanceRoot = beaconChain.getBalanceRoot(eIdx);
    //     emit log_named_bytes32("balance root e 5", balanceRoot);
    //     uint64 balanceE = beaconChain.getBalance(eIdx);
    //     emit log_named_uint("balance e 5", balanceE);
    // }
}