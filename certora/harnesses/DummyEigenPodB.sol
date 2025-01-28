// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../src/contracts/pods/EigenPod.sol";
import "./EigenPodHarness.sol";

contract DummyEigenPodB is EigenPodHarness {

    constructor(
        IETHPOSDeposit _ethPOS,
        IEigenPodManager _eigenPodManager,
        uint64 _GENESIS_TIME
    )
    EigenPodHarness(_ethPOS, _eigenPodManager, _GENESIS_TIME) {}
}
