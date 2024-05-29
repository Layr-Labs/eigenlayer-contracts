// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

contract Integration_Tester is IntegrationCheckUtils {

    using Strings for *;
    using StdStyle for *;
    
    string constant SECTION_DELIMITER = "======";

    function test_VerifyAll_Start_CompleteCP_WithRewardsWithdrawn(uint24 _rand) public {
        _configRand({
            _randomSeed: _rand,
            _assetTypes: HOLDS_ETH,
            _userTypes: DEFAULT
        });

        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint beaconBalanceWei) = staker.startValidators();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceWei);

        beaconChain.advanceEpoch();
        // check pod balances have increased

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    function test_VerifyAll_Start_CompleteCP_WithRewardsNotWithdrawn(uint24 _rand) public {
        _configRand({
            _randomSeed: _rand,
            _assetTypes: HOLDS_ETH,
            _userTypes: DEFAULT
        });

        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint beaconBalanceWei) = staker.startValidators();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceWei);

        beaconChain.advanceEpoch_NoWithdraw();
        // check pod balances have increased

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    function test_VerifyAll_Start_CompleteCP_NoRewards(uint24 _rand) public {
        _configRand({
            _randomSeed: _rand,
            _assetTypes: HOLDS_ETH,
            _userTypes: DEFAULT
        });

        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint beaconBalanceWei) = staker.startValidators();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceWei);

        beaconChain.advanceEpoch_NoRewards();
        // check pod balances have increased

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
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
}