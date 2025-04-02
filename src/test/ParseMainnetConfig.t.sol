// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.19;

import "forge-std/Test.sol";
import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";

struct Governance {
    address communityMultisig;
    address executorMultisig;
    address operationsMultisig;
    address pauserMultisig;
    address pauserRegistry;
    address proxyAdmin;
    address timelock;
}

struct Protocol {
    address allocationManager;
    address avsDirectory;
    // address beaconOracle; // EigenLayerBeaconOracle
    address bEIGEN;
    address delayedWithdrawalRouter;
    address delegationManager;
    address EIGEN;
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
    using ConfigParser for *;
    using stdToml for string;
    using StdStyle for string;

    Vm constant vm = Vm(address(uint160(uint(keccak256("hevm cheat code")))));
    bytes32 constant ERC1967_IMPLEMENTATION_SLOT = bytes32(uint(keccak256("eip1967.proxy.implementation")) - 1);
    bytes32 constant ERC1967_BEACON_SLOT = bytes32(uint(keccak256("eip1967.proxy.beacon")) - 1);
    bytes32 constant ERC1967_ADMIN_SLOT = bytes32(uint(keccak256("eip1967.proxy.admin")) - 1);

    /// -----------------------------------------------------------------------
    /// Config Parsing
    /// -----------------------------------------------------------------------

    /// @dev Returns the config from a given TOML file.
    function parse(string memory path) internal returns (Config memory) {
        return abi.decode(vm.readFile(path).parseRaw("."), (Config));
    }

    /// -----------------------------------------------------------------------
    /// Proxy Storage
    /// -----------------------------------------------------------------------

    /// @dev Returns the implementation address of a proxy.
    function impl(address proxy) internal view returns (address) {
        return vm.load(proxy, ERC1967_IMPLEMENTATION_SLOT).toAddress();
    }

    /// @dev Returns the beacon address of a proxy.
    function beacon(address proxy) internal view returns (address) {
        return vm.load(proxy, ERC1967_BEACON_SLOT).toAddress();
    }

    /// @dev Returns the admin address of a proxy.
    function admin(address proxy) internal view returns (address) {
        return vm.load(proxy, ERC1967_ADMIN_SLOT).toAddress();
    }

    /// -----------------------------------------------------------------------
    /// Logging
    /// -----------------------------------------------------------------------

    function log(Config memory config) internal {
        console2.log("block.number", block.number);

        // Log governance addresses
        console2.log("\nGovernance Addresses");
        config.governance.communityMultisig.log("communityMultisig");
        config.governance.executorMultisig.log("executorMultisig");
        config.governance.operationsMultisig.log("operationsMultisig");
        config.governance.pauserMultisig.log("pauserMultisig");
        config.governance.pauserRegistry.log("pauserRegistry");
        config.governance.proxyAdmin.log("proxyAdmin");
        config.governance.timelock.log("timelock");

        // Log proxy addresses
        console2.log("\nProxy Addresses:");
        config.protocol.allocationManager.logProxy("allocationManager");
        config.protocol.avsDirectory.logProxy("avsDirectory");
        // config.protocol.beaconOracle.logProxy("beaconOracle");
        config.protocol.bEIGEN.logProxy("bEIGEN");
        config.protocol.delayedWithdrawalRouter.logProxy("delayedWithdrawalRouter");
        config.protocol.delegationManager.logProxy("delegationManager");
        config.protocol.EIGEN.logProxy("EIGEN");
        config.protocol.eigenPodBeacon.logBeaconImpl("eigenPodBeacon");
        config.protocol.eigenPodManager.logProxy("eigenPodManager");
        config.protocol.eigenStrategy.logProxy("eigenStrategy");
        config.protocol.permissionController.logProxy("permissionController");
        config.protocol.rewardsCoordinator.logProxy("rewardsCoordinator");
        config.protocol.strategyFactory.logProxy("strategyFactory");
        config.protocol.strategyFactoryBeacon.logBeaconImpl("strategyFactoryBeacon");
        config.protocol.strategyManager.logProxy("strategyManager");

        // Log strategy addresses
        console2.log("\nStrategy Addresses:");
        for (uint i = 0; i < config.strategies.addresses.length; i++) {
            config.strategies.addresses[i].logProxy(string.concat("strategy", vm.toString(i)));
        }
    }

    function log(address addr, string memory name) internal {
        console2.log(name.yellow(), addr);
    }

    /// @dev Logs the implementation address of a beacon with a name.
    function logBeaconImpl(address beacon, string memory name) internal {
        console2.log(name.yellow(), beacon, vm.toString(IBeacon(beacon).implementation()).dim());
    }

    /// @dev Logs the implementation address of a proxy with a name.
    function logProxy(address proxy, string memory name) internal {
        console2.log(name.yellow(), proxy, vm.toString(proxy.impl()).dim());
    }

    function toAddress(bytes32 x) internal pure returns (address r) {
        assembly {
            r := x
        }
    }
}

contract ParseMainnetConfigTest is Test {
    using ConfigParser for *;

    function test() public {
        vm.createSelectFork(vm.rpcUrl("mainnet"), 22_181_590);

        // Parse the config from the TOML file.
        Config memory config = ConfigParser.parse("./script/configs/mainnet/mainnet-addresses.config.toml");

        config.log();
    }
}
