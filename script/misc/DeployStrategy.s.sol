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

import "../src/contracts/strategies/StrategyBase.sol";

import "../src/contracts/pods/EigenPod.sol";
import "../src/contracts/pods/EigenPodManager.sol";
import "../src/contracts/pods/DelayedWithdrawalRouter.sol";

import "../src/contracts/permissions/PauserRegistry.sol";
import "../src/contracts/middleware/BLSPublicKeyCompendium.sol";

import "../src/contracts/libraries/BytesLib.sol";

import "../src/test/mocks/EmptyContract.sol";
import "../src/test/mocks/ETHDepositMock.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/misc/DeployStrategy.s.sol:DeployStrategy --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv

// NOTE: ONLY WORKS ON GOERLI
contract DeployStrategy is Script, Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    // EigenLayer Contracts
    ProxyAdmin public eigenLayerProxyAdmin = ProxyAdmin(0x28ceac2ff82B2E00166e46636e2A4818C29902e2);
    PauserRegistry public eigenLayerPauserReg = PauserRegistry(0x7cB9c5D6b9702f2f680e4d35cb1fC945D08208F6);
    StrategyBase public baseStrategyImplementation = StrategyBase(0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d);

    function run() external {
        // read and log the chainID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        address tokenAddress = 0x3ecCAdA3e11c1Cc3e9B5a53176A67cc3ABDD3E46;

        vm.startBroadcast();

        // create upgradeable proxies that each point to the implementation and initialize them
        address strategy = address(
                    new TransparentUpgradeableProxy(
                        address(baseStrategyImplementation),
                        address(eigenLayerProxyAdmin),
                        abi.encodeWithSelector(StrategyBase.initialize.selector, IERC20(tokenAddress), eigenLayerPauserReg)
                    )
                );

        vm.stopBroadcast();

        emit log_named_address("Strategy", strategy);

    }
}