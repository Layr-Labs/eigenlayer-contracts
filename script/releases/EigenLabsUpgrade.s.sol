// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {ZeusScript} from "zeus-templates/utils/ZeusScript.sol";
import {EncGnosisSafe} from "zeus-templates/utils/EncGnosisSafe.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

library EigenLabsUpgrade {
    using EncGnosisSafe for *;

    function _ethPos(
        ZeusScript self
    ) internal view returns (address) {
        return self.zAddress("ethPOS");
    }

    function _eigenpodGenesisTime(
        ZeusScript self
    ) internal view returns (uint64) {
        return self.zUint64("EIGENPOD_GENESIS_TIME");
    }

    function _eigenPodManagerPendingImpl(
        ZeusScript self
    ) internal view returns (address) {
        return self.zDeployedContract("EigenPodManager_pendingImpl");
    }

    function _operationsMultisig(
        ZeusScript self
    ) internal view returns (address) {
        return self.zAddress("operationsMultisig");
    }

    function _pauserRegistry(
        ZeusScript self
    ) internal view returns (address) {
        return self.zDeployedImpl("PauserRegistry");
    }

    function _proxyAdmin(
        ZeusScript self
    ) internal view returns (address) {
        return self.zAddress("proxyAdmin");
    }

    function _eigenPodManagerProxy(
        ZeusScript self
    ) internal view returns (address) {
        return self.zAddress("EigenPodManager_proxy");
    }

    function _eigenPodBeacon(
        ZeusScript self
    ) internal view returns (address) {
        return self.zAddress("EigenPod_beacon");
    }

    function _eigenPodPendingImpl(
        ZeusScript self
    ) internal view returns (address) {
        return self.zAddress("EigenPod_pendingImpl");
    }

    function _multiSendCallOnly(
        ZeusScript self
    ) internal view returns (address) {
        return self.zAddress("MultiSendCallOnly");
    }

    function _timelock(
        ZeusScript self
    ) internal view returns (TimelockController) {
        return TimelockController(payable(self.zAddress("timelockController")));
    }

    function _executorMultisig(
        ZeusScript self
    ) internal view returns (address) {
        return self.zAddress("executorMultisig");
    }

    function _protocolCouncilMultisig(
        ZeusScript self
    ) internal view returns (address) {
        return self.zAddress("protocolCouncilMultisig");
    }
}
