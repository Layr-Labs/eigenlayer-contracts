// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../contracts/pods/EigenPodManager.sol";

///@notice This contract exposed the internal `_calculateChangeInDelegatableShares` function for testing
contract EigenPodManagerWrapper is EigenPodManager {
    
    constructor(
        IETHPOSDeposit _ethPOS,
        IBeacon _eigenPodBeacon,
        IStrategyManager _strategyManager,
        IDelegationManager _delegationManager
    ) EigenPodManager(_ethPOS, _eigenPodBeacon, _strategyManager, _delegationManager) {}

    function calculateChangeInDelegatableShares(int256 sharesBefore, int256 sharesAfter) external pure returns (int256) {
        return _calculateChangeInDelegatableShares(sharesBefore, sharesAfter);
    }

    function setPodAddress(address owner, IEigenPod pod) external {
        ownerToPod[owner] = pod;
    }
}
