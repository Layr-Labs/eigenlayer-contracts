// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/pods/EigenPodManager.sol";

contract EigenPodManagerHarness is EigenPodManager {

    constructor(
        IETHPOSDeposit _ethPOS,
        IBeacon _eigenPodBeacon,
        IStrategyManager _strategyManager,
        ISlasher _slasher,
        IDelegationManager _delegationManager
    )
        EigenPodManager(_ethPOS, _eigenPodBeacon, _strategyManager, _slasher, _delegationManager) {}

    function get_podOwnerShares(address podOwner) public view returns (int256) {
        return podOwnerShares[podOwner];
    }

    function get_podByOwner(address podOwner) public view returns (IEigenPod) {
        return ownerToPod[podOwner];
    }
}
