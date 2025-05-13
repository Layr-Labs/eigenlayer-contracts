// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/pods/EigenPod.sol";

contract EigenPodHarness is EigenPod {

    constructor(
        IETHPOSDeposit _ethPOS,
        IEigenPodManager _eigenPodManager,
        string memory _version
    )
        EigenPod(_ethPOS, _eigenPodManager, _version) {}

    function get_validatorIndex(bytes32 pubkeyHash) public view returns (uint64) {
        return _validatorPubkeyHashToInfo[pubkeyHash].validatorIndex;
    }

    function get_restakedBalanceGwei(bytes32 pubkeyHash) public view returns (uint64) {
        return _validatorPubkeyHashToInfo[pubkeyHash].restakedBalanceGwei;
    }

    function get_mostRecentBalanceUpdateTimestamp(bytes32 pubkeyHash) public view returns (uint64) {
        return _validatorPubkeyHashToInfo[pubkeyHash].mostRecentBalanceUpdateTimestamp;
    }

    function get_podOwnerShares() public view returns (int256) {
        return eigenPodManager.podOwnerShares(podOwner);
    }

    function get_withdrawableRestakedExecutionLayerGwei() public view returns (uint256) {
        return withdrawableRestakedExecutionLayerGwei;
    }

    function get_ETH_Balance() public view returns (uint256) {
        return address(this).balance;
    }
}
