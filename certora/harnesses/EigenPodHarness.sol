// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/pods/EigenPod.sol";

contract EigenPodHarness is EigenPod {

    constructor(
        IETHPOSDeposit _ethPOS,
        IEigenPodManager _eigenPodManager,
        uint64 _MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
        uint64 _GENESIS_TIME
    )
        EigenPod(_ethPOS, _eigenPodManager, _MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR, _GENESIS_TIME) {}




    function get_podOwnerShares() public view returns (int256) {
        return eigenPodManager.podOwnerShares(podOwner);
    }


}
