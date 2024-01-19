// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../src/contracts/interfaces/IETHPOSDeposit.sol";
import "../src/contracts/interfaces/IBeaconChainOracle.sol";

import "../src/contracts/core/StrategyManager.sol";
import "../src/contracts/core/Slasher.sol";
import "../src/contracts/core/DelegationManager.sol";

import "../src/contracts/strategies/StrategyBaseTVLLimits.sol";

import "../src/contracts/pods/EigenPod.sol";
import "../src/contracts/pods/EigenPodManager.sol";
import "../src/contracts/pods/DelayedWithdrawalRouter.sol";

import "../src/contracts/permissions/PauserRegistry.sol";

import "../src/test/mocks/EmptyContract.sol";
import "../src/test/mocks/ETHDepositMock.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";


contract EigenPodDeployer is Script, Test {


    IEigenPodManager public epm = IEigenPodManager(address(0x33e42d539abFe9b387B27b0e467374Bbb76cf925));
    function run() external {
        address hello = epm.createPod();
        emit log_named_address("THIS IS THE DEPLOYED ADDRESS", hello);
    }


}