// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

interface IUpgradeableBeacon {
    function upgradeTo(
        address newImplementation
    ) external;
    function implementation() external returns (address);
}
