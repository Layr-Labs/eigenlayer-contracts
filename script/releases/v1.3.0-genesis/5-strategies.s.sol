// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {console} from "forge-std/console.sol";
import "src/contracts/libraries/BeaconChainProofs.sol";
import {Execute, Queue} from "./3-execute.s.sol";
import {DeployFresh} from "./1-deploy.s.sol";


/**
 * Purpose: optionally enqueue a transaction that pauses all EigenPod functionality.
 */
contract Strategies is MultisigBuilder, Execute {
    using Env for *;

    function _runAsMultisig() prank(Env.opsMultisig()) internal virtual override(Execute, MultisigBuilder) {
        console.log("StrategyBeacon", address(Env.proxy.strategyFactory().strategyBeacon()));

        console.log("StrategyWhitelister", address(Env.proxy.strategyManager().strategyWhitelister()));
        console.log("OpsMultisig", Env.opsMultisig());

        if (Env.weth() != address(0)) {
            address wethStrategy = address(Env.proxy.strategyFactory().deployNewStrategy(IERC20(Env.weth())));
            deployInstance({
                name: type(StrategyBaseTVLLimits).name,
                deployedTo: wethStrategy
            });
        }

        if (Env.steth() != address(0)) {
            deployInstance({
                name: type(StrategyBaseTVLLimits).name,
                deployedTo: address(Env.proxy.strategyFactory().deployNewStrategy(IERC20(Env.steth())))
            });
        }
    }

    function testScript() public virtual override {
        DeployFresh._runAsEOA();
        _unsafeResetHasPranked();

        Queue._runAsMultisig();
        _unsafeResetHasPranked();

        TimelockController timelock = Env.timelockController();
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA

        Execute._runAsMultisig();
        _unsafeResetHasPranked();

        execute();
    }
}
