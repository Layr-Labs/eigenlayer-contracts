// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

import "../../../src/contracts/strategies/StrategyFactory.sol";

// forge script script/deploy/mainnet/Deploy_Preprod_StrategyFactory.s.sol:Deploy_Preprod_StrategyFactory
//      --rpc-url $RPC_URL
//      --private-key $PRIVATE_KEY
//      --etherscan-api-key $ETHERSCAN_API_KEY
//      --verify
//      --broadcast
//      --vvvv
contract Deploy_Mainnet_StrategyFactory is Script, Test {
    function _addr(bytes memory b) internal pure returns (address) {
        return address(uint160(uint256(bytes32(b))));
    }

    function run() external {
        vm.startBroadcast();

        string memory json = "../../configs/mainnet/Deploy_Mainnet_StrategyFactory.config.json";

        address owner = _addr(vm.parseJson(json, "owner"));
        address strategyImplementation = _addr(vm.parseJson(json, "strategyImplementation"));
        address strategyManager = _addr(vm.parseJson(json, "strategyManager"));
        address eigenLayerProxyAdmin = _addr(vm.parseJson(json, "eigenLayerProxyAdmin"));
        address pauserRegistry = _addr(vm.parseJson(json, "pauserRegistry"));
        uint256 initialPausedStatus = uint256(bytes32(vm.parseJson(json, "initialPausedStatus")));

        // Deploy a Proxy Beacon for the strategy and transfer ownership.
        UpgradeableBeacon strategyBeacon = new UpgradeableBeacon(strategyImplementation);

        strategyBeacon.transferOwnership(owner);

        // Deploy a Strategy Factory implementation.
        StrategyFactory strategyFactoryImplementation = new StrategyFactory(IStrategyManager(strategyManager));
        // Deploy a Upgradeable Strategy Factory proxy.
        StrategyFactory strategyFactory = StrategyFactory(
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

        console.log("Strategy Proxy Beacon: ", address(strategyBeacon));
        console.log("Strategy Factory Implementation: ", address(strategyFactoryImplementation));
        console.log("Strategy Factory Proxy: ", address(strategyFactory));

        vm.stopBroadcast();
    }
}
