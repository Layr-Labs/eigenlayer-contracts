// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

struct TUPInfo {
    address proxy;
    address impl;
    address pendingImpl;
}

struct BeaconInfo {
    address beacon;
    address impl;
    address pendingImpl;
}

struct TokenInfo {
    address proxy;
    address impl;
    address pendingImpl;
    address proxyAdmin;
}

struct Addresses {
    // admin
    address communityMultisig;
    address executorMultisig;
    address operationsMultisig;
    address pauserMultisig;
    address pauserRegistry;
    address proxyAdmin;
    address timelock;
    // core
    TUPInfo avsDirectory;
    TUPInfo delegationManager;
    TUPInfo rewardsCoordinator;
    TUPInfo slasher;
    TUPInfo strategyManager;
    // pods
    BeaconInfo eigenPod;
    TUPInfo eigenPodManager;
    TUPInfo delayedWithdrawalRouter;
    // strategies
    TUPInfo strategyFactory;
    BeaconInfo strategyBeacon;
    TUPInfo[] preLongtailStrats;
    // token
    TokenInfo EIGEN;
    TokenInfo bEIGEN;
    TUPInfo eigenStrategy;
}

library AddressUtils {

    function setPending(TUPInfo memory info, address pendingAddress) internal pure {
        info.pendingImpl = pendingAddress;
    }

    function setPending(BeaconInfo memory info, address pendingAddress) internal pure {
        info.pendingImpl = pendingAddress;
    }

    function setPending(TokenInfo memory info, address pendingAddress) internal pure {
        info.pendingImpl = pendingAddress;
    }

    function updateFromPending(TUPInfo memory info) internal pure {
        info.impl = getPending(info);
        info.pendingImpl = address(0);
    }

    function updateFromPending(BeaconInfo memory info) internal pure {
        info.impl = getPending(info);
        info.pendingImpl = address(0);
    }

    function updateFromPending(TokenInfo memory info) internal pure {
        info.impl = getPending(info);
        info.pendingImpl = address(0);
    }

    function getPending(TUPInfo memory info) internal pure returns (address) {
        if (info.pendingImpl == address(0)) {
            revert("no pending implementation found");
        }

        return info.pendingImpl;
    }

    function getPending(BeaconInfo memory info) internal pure returns (address) {
        if (info.pendingImpl == address(0)) {
            revert("no pending implementation found");
        }

        return info.pendingImpl;
    }

    function getPending(TokenInfo memory info) internal pure returns (address) {
        if (info.pendingImpl == address(0)) {
            revert("no pending implementation found");
        }

        return info.pendingImpl;
    }
}