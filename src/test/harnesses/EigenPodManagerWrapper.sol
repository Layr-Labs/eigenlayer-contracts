// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../contracts/pods/EigenPodManager.sol";

///@notice This contract exposes a manual setter for podShares in order to initialize podShares as negative
contract EigenPodManagerWrapper is EigenPodManager {
    constructor(
        IETHPOSDeposit _ethPOS,
        IBeacon _eigenPodBeacon,
        IDelegationManager _delegationManager,
        IPauserRegistry _pauserRegistry
    ) EigenPodManager(_ethPOS, _eigenPodBeacon, _delegationManager, _pauserRegistry) {}

    function setPodOwnerShares(address owner, IEigenPod pod) external {
        ownerToPod[owner] = pod;
    }

    function setPodOwnerShares(address owner, int256 shares) external {
        podOwnerDepositShares[owner] = shares;
    }
}
