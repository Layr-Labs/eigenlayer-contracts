// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../contracts/pods/EigenPod.sol";
import "forge-std/Test.sol";

contract EigenPodHarness is EigenPod {

    constructor(
        IETHPOSDeposit _ethPOS,
        IEigenPodManager _eigenPodManager,
        uint64 _GENESIS_TIME,
        string memory _version
    ) EigenPod(
        _ethPOS,
        _eigenPodManager,
        _GENESIS_TIME,
        _version
    ) {}

    function getActiveValidatorCount() public view returns (uint256) {
        return activeValidatorCount;
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


}
