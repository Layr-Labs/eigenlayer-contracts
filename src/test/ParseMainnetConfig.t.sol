// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.19;

import "forge-std/Test.sol";

struct Governance {
    address communityMultisig;
    address executorMultisig;
    address operationsMultisig;
    address pauserMultisig;
    address timelock;
}

struct Protocol {
    address allocationManager;
    address avsDirectory;
    address beaconOracle;
    address bEIGEN;
    address delayedWithdrawalRouter;
    address delegationManager;
    address EIGEN;
    address eigenLayerPauserReg;
    address eigenLayerProxyAdmin;
    address eigenPodBeacon;
    address eigenPodManager;
    address eigenStrategy;
    address permissionController;
    address rewardsCoordinator;
    address strategyFactory;
    address strategyFactoryBeacon;
    address strategyManager;
}

struct Strategies {
    address[] addresses;
}

struct Config {
    Governance governance;
    Protocol protocol;
    Strategies strategies;
}

library ConfigParser {
    using stdToml for string;

    Vm constant vm = Vm(address(uint160(uint(keccak256("hevm cheat code")))));
    bytes32 constant ERC1967_IMPLEMENTATION_SLOT = bytes32(uint(keccak256("eip1967.proxy.implementation")) - 1);
    bytes32 constant ERC1967_BEACON_SLOT = bytes32(uint(keccak256("eip1967.proxy.beacon")) - 1);
    bytes32 constant ERC1967_ADMIN_SLOT = bytes32(uint(keccak256("eip1967.proxy.admin")) - 1);

    function impl(address proxy) internal view returns (bytes32) {
        return vm.load(proxy, ERC1967_IMPLEMENTATION_SLOT);
    }

    function beacon(address proxy) internal view returns (bytes32) {
        return vm.load(proxy, ERC1967_BEACON_SLOT);
    }

    function admin(address proxy) internal view returns (bytes32) {
        return vm.load(proxy, ERC1967_ADMIN_SLOT);
    }

    function parse(string memory path) internal returns (Config memory) {
        string memory toml = vm.readFile(path);
        return abi.decode(toml.parseRaw("."), (Config));
    }
}

contract ParseMainnetConfigTest is Test {
    function test() public {
        // Parse the config from the TOML file.
        Config memory config = ConfigParser.parse("./script/configs/mainnet/mainnet-addresses.config.toml");

        // Log governance addresses
        console2.log("\nGovernance Addresses");
        console2.log("communityMultisig:", config.governance.communityMultisig);
        console2.log("executorMultisig:", config.governance.executorMultisig);
        console2.log("operationsMultisig:", config.governance.operationsMultisig);
        console2.log("pauserMultisig:", config.governance.pauserMultisig);
        console2.log("timelock:", config.governance.timelock);

        // Log proxy addresses
        console2.log("\nProxy Addresses:");
        console2.log("allocationManager:", config.protocol.allocationManager);
        console2.log("avsDirectory:", config.protocol.avsDirectory);
        console2.log("beaconOracle:", config.protocol.beaconOracle);
        console2.log("bEIGEN:", config.protocol.bEIGEN);
        console2.log("delayedWithdrawalRouter:", config.protocol.delayedWithdrawalRouter);
        console2.log("delegationManager:", config.protocol.delegationManager);
        console2.log("EIGEN:", config.protocol.EIGEN);
        console2.log("eigenLayerPauserReg:", config.protocol.eigenLayerPauserReg);
        console2.log("eigenLayerProxyAdmin:", config.protocol.eigenLayerProxyAdmin);
        console2.log("eigenPodBeacon:", config.protocol.eigenPodBeacon);
        console2.log("eigenPodManager:", config.protocol.eigenPodManager);
        console2.log("eigenStrategy:", config.protocol.eigenStrategy);
        console2.log("permissionController:", config.protocol.permissionController);
        console2.log("rewardsCoordinator:", config.protocol.rewardsCoordinator);
        console2.log("strategyFactory:", config.protocol.strategyFactory);
        console2.log("strategyFactoryBeacon:", config.protocol.strategyFactoryBeacon);
        console2.log("strategyManager:", config.protocol.strategyManager);

        // Log strategy addresses
        console2.log("\nStrategy Addresses:");
        for (uint i = 0; i < config.strategies.addresses.length; i++) {
            console2.log("strategy", i + 1, ":", config.strategies.addresses[i]);
        }
    }
}
