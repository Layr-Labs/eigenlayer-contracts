// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../src/contracts/pods/EigenPod.sol";
import "./EigenPodHarness.sol";

contract DummyEigenPodB is EigenPodHarness {

    constructor(
        IETHPOSDeposit _ethPOS,
        IDelayedWithdrawalRouter _delayedWithdrawalRouter,
        IEigenPodManager _eigenPodManager,
        uint64 _MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
        uint64 _GENESIS_TIME
    )
    EigenPodHarness(_ethPOS, _delayedWithdrawalRouter, _eigenPodManager, _MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR, _GENESIS_TIME) {}
}
