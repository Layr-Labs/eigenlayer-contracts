// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

import "../../../src/contracts/strategies/EigenStrategy.sol";
import "../../../src/contracts/strategies/StrategyFactory.sol";

// forge script script/Deploy_Preprod_StrategyFactory.s.sol:Deploy_Preprod_StrategyFactory
//      --rpc-url $RPC_URL
//      --private-key $PRIVATE_KEY
//      --etherscan-api-key $ETHERSCAN_API_KEY
//      --verify
//      --broadcast
//      --vvvv
contract Deploy_Preprod_StrategyManager is Script, Test {
    // https://docs.google.com/spreadsheets/d/1w8ckBfdVyv-Xh6iPwZ3hseDfX1gEptoJfomuwA8h3Jo/edit?gid=1046619329#gid=1046619329
    address owner = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;
    IPauserRegistry pauserRegistry = IPauserRegistry(0x9Ab2FEAf0465f0eD51Fc2b663eF228B418c9Dad1);
    IStrategyManager strategyManager = IStrategyManager(0xF9fbF2e35D8803273E214c99BF15174139f4E67a);
    address eigenLayerProxyAdmin = 0x1BEF05C7303d44e0E2FCD2A19d993eDEd4c51b5B;
    uint256 initialPausedStatus = 0;

    function run()
        external
        returns (
            UpgradeableBeacon strategyBeacon,
            EigenStrategy strategyImplementation,
            StrategyFactory strategyFactoryImplementation,
            StrategyFactory strategyFactory
        )
    {
        vm.startBroadcast();

        strategyFactoryImplementation = new StrategyFactory(strategyManager);
        strategyImplementation = new EigenStrategy(strategyManager);
        strategyBeacon = new UpgradeableBeacon(address(strategyImplementation));
        strategyFactory = StrategyFactory(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyFactoryImplementation),
                    eigenLayerProxyAdmin,
                    abi.encodeWithSelector(
                        StrategyFactory.initialize.selector, owner, pauserRegistry, initialPausedStatus, strategyBeacon
                    )
                )
            )
        );

        strategyBeacon.transferOwnership(owner);
        vm.stopBroadcast();
    }
}
