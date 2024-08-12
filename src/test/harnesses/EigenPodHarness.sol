// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../contracts/pods/EigenPod.sol";
import "forge-std/Test.sol";

contract EigenPodHarness is EigenPod {

    constructor(
        IETHPOSDeposit _ethPOS,
        IEigenPodManager _eigenPodManager,
        uint64 _GENESIS_TIME
    ) EigenPod(
        _ethPOS,
        _eigenPodManager,
        _GENESIS_TIME
    ) {}

    function getActiveValidatorCount() public view returns (uint256) {
        return activeValidatorCount;
    }

    function setActiveValidatorCount(uint _count) public {
        activeValidatorCount = _count;
    }

    function verifyWithdrawalCredentials(
        bytes32 beaconStateRoot,
        uint40 validatorIndex,
        bytes calldata validatorFieldsProof,
        bytes32[] calldata validatorFields
    ) public returns (uint256) {
        return _verifyWithdrawalCredentials(
            beaconStateRoot,
            validatorIndex,
            validatorFieldsProof,
            validatorFields
        );
    }

    function setValidatorStatus(bytes32 pkhash, VALIDATOR_STATUS status) public {
        _validatorPubkeyHashToInfo[pkhash].status = status;
    }

    function setValidatorRestakedBalance(bytes32 pkhash, uint64 restakedBalanceGwei) public {
        _validatorPubkeyHashToInfo[pkhash].restakedBalanceGwei = restakedBalanceGwei;
    }
}
