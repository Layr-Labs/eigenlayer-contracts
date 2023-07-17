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
import "../src/contracts/middleware/BLSPublicKeyCompendium.sol";

import "../src/test/mocks/EmptyContract.sol";
import "../src/test/mocks/ETHDepositMock.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";



contract Withdraw is Script, Test {
    StrategyManager public strategyManager;
    StrategyManager public strategyManagerImplementation;

    function run() public {
        strategyManager = StrategyManager(0x858646372CC42E1A627fcE94aa7A7033e7CF075A);

        vm.startBroadcast();

        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;

        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = IStrategy(0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc);

        uint256[] memory shares = new uint256[](1);     
        shares[0] = 900000000000000000;   

        address withdawer = 0x336b4940f39b575893BAd2798b53E01EAecD3170;

        bool undelegateIfPossible = true;
    
        //strategyManager.queueWithdrawal(strategyIndexes, strategies, shares, withdawer, undelegateIfPossible);

        StrategyManager.WithdrawerAndNonce memory withdrawerAndNonce = IStrategyManager.WithdrawerAndNonce({
            withdrawer: withdawer,
            nonce: 1
        });

        StrategyManager.QueuedWithdrawal memory queuedWithdrawal = IStrategyManager.QueuedWithdrawal({
            strategies: strategies,
            shares: shares,
            depositor: 0x3Bda09943b6D0Eda1B4fdE3a7344897032b24061,
            withdrawerAndNonce: withdrawerAndNonce,
            withdrawalStartBlock: 17709941,
            delegatedAddress: address(0)
        });

        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = IERC20(0xBe9895146f7AF43049ca1c1AE358B0541Ea49704);

        uint256 middlewareTimesIndex = 0;
        bool receiveAsTokens = true;
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokens, middlewareTimesIndex, receiveAsTokens);
    }
}