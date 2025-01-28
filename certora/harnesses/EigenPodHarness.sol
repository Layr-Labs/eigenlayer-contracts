// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../src/contracts/pods/EigenPod.sol";

contract EigenPodHarness is EigenPod {

    constructor(
        IETHPOSDeposit _ethPOS,
        IEigenPodManager _eigenPodManager,
        //uint64 _MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
        uint64 _GENESIS_TIME
    )
        EigenPod(_ethPOS, _eigenPodManager, /* _MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR, */ _GENESIS_TIME) {}

    function get_validatorIndex(bytes32 pubkeyHash) public view returns (uint64) {
        return _validatorPubkeyHashToInfo[pubkeyHash].validatorIndex;
    }

    function get_validatorLastCheckpointed(bytes32 pubkeyHash) public view returns (uint64) {
        return _validatorPubkeyHashToInfo[pubkeyHash].lastCheckpointedAt;
    }

    function validatorIsActive(bytes32 pubkeyHash) public view returns (bool) {
        return _validatorPubkeyHashToInfo[pubkeyHash].status == IEigenPod.VALIDATOR_STATUS.ACTIVE;
    }
   
    function get_lastCheckpointTimestamp() public view returns (uint64)
    {
        return lastCheckpointTimestamp;
    }

    function get_currentCheckpointTimestamp() public view returns (uint64)
    {
        return currentCheckpointTimestamp;
    }    

    function get_restakedBalanceGwei(bytes32 pubkeyHash) public view returns (uint64) {
        return _validatorPubkeyHashToInfo[pubkeyHash].restakedBalanceGwei;
    }

    function get_mostRecentBalanceUpdateTimestamp(bytes32 pubkeyHash) public view returns (uint64) {
        return _validatorPubkeyHashToInfo[pubkeyHash].lastCheckpointedAt;
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
