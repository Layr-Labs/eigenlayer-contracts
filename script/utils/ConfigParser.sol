// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Script.sol";
import "forge-std/Test.sol";

import "./StringUtils.sol";

struct Environment {
    uint chainid;
    string name;
    string lastUpdated;
}

struct Params {
    // admin
    address multiSendCallOnly;
    // pods
    address ethPOS;
    uint64 EIGENPOD_GENESIS_TIME;
    // rewards
    uint32 CALCULATION_INTERVAL_SECONDS;
    uint32 MAX_REWARDS_DURATION;
    uint32 MAX_RETROACTIVE_LENGTH;
    uint32 MAX_FUTURE_LENGTH;
    uint32 GENESIS_REWARDS_TIMESTAMP;
    address REWARDS_UPDATER_ADDRESS;
    uint32 ACTIVATION_DELAY;
    uint16 GLOBAL_OPERATOR_COMMISSION_BIPS;
}

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

contract ConfigParser is Script, Test {

    using StringUtils for *;

    string private _configPath;
    string private _configData;

    function _readConfigFile(string memory configPath) internal returns (Addresses memory, Environment memory, Params memory) {
        _configPath = configPath;
        _configData = vm.readFile(_configPath);
        emit log_named_string("Reading from config file", _configPath);

        Environment memory env = _readEnvironment();
        Params memory params = _readParams();
        Addresses memory addrs = _readAddresses();

        return (addrs, env, params);
    }

    function _writeConfigFile(Addresses memory addrs, Environment memory env, Params memory params) internal {
        emit log_named_string("Writing to config file", _configPath);

        (string memory configKey, string memory configObject) = _writeConfig(env, params);
        (string memory deploymentKey, string memory deploymentObject) = _writeDeployment(addrs);

        string memory parentObjectKey = "root";
        vm.serializeString(parentObjectKey, deploymentKey, deploymentObject);
        string memory parentObject = vm.serializeString(parentObjectKey, configKey, configObject);

        vm.writeJson(parentObject, _configPath);
    }

    function _printEnv(Environment memory env) internal {
        emit log("Config Environment Info:");
        emit log_named_string("- name", env.name);
        emit log_named_uint("- chainid", env.chainid);
        emit log_named_string("- last updated", env.lastUpdated);
    }

    /**
     *
     *                            READS
     *
     */

    function _readEnvironment() private returns (Environment memory) {
        return Environment({
            chainid: _readUint(".config.environment.chainid"),
            name: _readString(".config.environment.name"),
            lastUpdated: _readString(".config.environment.lastUpdated")
        });
    }

    function _readParams() private returns (Params memory) {
        return Params({
            multiSendCallOnly: _readAddress(".config.params.multiSendCallOnly"),
            ethPOS: _readAddress(".config.params.ethPOS"),
            EIGENPOD_GENESIS_TIME: uint64(_readUint(".config.params.EIGENPOD_GENESIS_TIME")),
            CALCULATION_INTERVAL_SECONDS: uint32(_readUint(".config.params.CALCULATION_INTERVAL_SECONDS")),
            MAX_REWARDS_DURATION: uint32(_readUint(".config.params.MAX_REWARDS_DURATION")),
            MAX_RETROACTIVE_LENGTH: uint32(_readUint(".config.params.MAX_RETROACTIVE_LENGTH")),
            MAX_FUTURE_LENGTH: uint32(_readUint(".config.params.MAX_FUTURE_LENGTH")),
            GENESIS_REWARDS_TIMESTAMP: uint32(_readUint(".config.params.GENESIS_REWARDS_TIMESTAMP")),
            REWARDS_UPDATER_ADDRESS: _readAddress(".config.params.REWARDS_UPDATER_ADDRESS"),
            ACTIVATION_DELAY: uint32(_readUint(".config.params.ACTIVATION_DELAY")),
            GLOBAL_OPERATOR_COMMISSION_BIPS: uint16(_readUint(".config.params.GLOBAL_OPERATOR_COMMISSION_BIPS"))
        });
    }

    function _readAddresses() private returns (Addresses memory) {
        return Addresses({
            // Admin
            communityMultisig: _readAddress(".deployment.admin.communityMultisig"),
            executorMultisig: _readAddress(".deployment.admin.executorMultisig"),
            operationsMultisig: _readAddress(".deployment.admin.operationsMultisig"),
            pauserMultisig: _readAddress(".deployment.admin.pauserMultisig"),
            pauserRegistry: _readAddress(".deployment.admin.pauserRegistry"),
            proxyAdmin: _readAddress(".deployment.admin.proxyAdmin"),
            timelock: _readAddress(".deployment.admin.timelock"),
            // Core
            avsDirectory: _readTUP(".deployment.core.avsDirectory"),
            delegationManager: _readTUP(".deployment.core.delegationManager"),
            rewardsCoordinator: _readTUP(".deployment.core.rewardsCoordinator"),
            slasher: _readTUP(".deployment.core.slasher"),
            strategyManager: _readTUP(".deployment.core.strategyManager"),
            // Pods
            eigenPod: _readBeacon(".deployment.pods.eigenPod"),
            eigenPodManager: _readTUP(".deployment.pods.eigenPodManager"),
            delayedWithdrawalRouter: _readTUP(".deployment.pods.delayedWithdrawalRouter"),
            // Strategies
            strategyFactory: _readTUP(".deployment.strategies.strategyFactory"),
            strategyBeacon: _readBeacon(".deployment.strategies.strategyBeacon"),
            preLongtailStrats: _readStrategies(".deployment.strategies.preLongtailStrats"),
            // Token
            EIGEN: _readToken(".deployment.token.EIGEN"),
            bEIGEN: _readToken(".deployment.token.bEIGEN"),
            eigenStrategy: _readTUP(".deployment.token.eigenStrategy")
        });
    }

    function _readTUP(string memory jsonLocation) private returns (TUPInfo memory) {
        return TUPInfo({
            proxy: _readAddress(jsonLocation.concat(".proxy")),
            impl: _readAddress(jsonLocation.concat(".impl")),
            pendingImpl: _readAddress(jsonLocation.concat(".pendingImpl"))
        });
    }

    function _readBeacon(string memory jsonLocation) private returns (BeaconInfo memory) {
        return BeaconInfo({
            beacon: _readAddress(jsonLocation.concat(".beacon")),
            impl: _readAddress(jsonLocation.concat(".impl")),
            pendingImpl: _readAddress(jsonLocation.concat(".pendingImpl"))
        });
    }

    function _readStrategies(string memory jsonLocation) private returns (TUPInfo[] memory) {
        address[] memory strategyProxies = stdJson.readAddressArray(_configData, jsonLocation.concat(".addrs"));
        address strategyImpl = _readAddress(jsonLocation.concat(".impl"));

        TUPInfo[] memory strategyInfos = new TUPInfo[](strategyProxies.length);

        for (uint i = 0; i < strategyInfos.length; i++) {
            strategyInfos[i] = TUPInfo({
                proxy: strategyProxies[i],
                impl: strategyImpl,
                pendingImpl: address(0)
            });
        }

        return strategyInfos;
    }

    function _readToken(string memory jsonLocation) private returns (TokenInfo memory) {
        return TokenInfo({
            proxy: _readAddress(jsonLocation.concat(".proxy")),
            impl: _readAddress(jsonLocation.concat(".impl")),
            pendingImpl: _readAddress(jsonLocation.concat(".pendingImpl")),
            proxyAdmin: _readAddress(jsonLocation.concat(".proxyAdmin"))
        });
    }

    function _readAddress(string memory jsonLocation) private returns (address) {
        return stdJson.readAddress(_configData, jsonLocation);
    }

    function _readUint(string memory jsonLocation) private returns (uint) {
        return stdJson.readUint(_configData, jsonLocation);
    }

    function _readString(string memory jsonLocation) private returns (string memory) {
        return stdJson.readString(_configData, jsonLocation);
    }

    /**
     *
     *                            WRITES
     *
     */

    function _writeDeployment(Addresses memory addrs) private returns (string memory, string memory) {
        // Admin
        string memory adminKey = "admin";
        vm.serializeAddress(adminKey, "communityMultisig", addrs.communityMultisig);
        vm.serializeAddress(adminKey, "executorMultisig", addrs.executorMultisig);
        vm.serializeAddress(adminKey, "operationsMultisig", addrs.operationsMultisig);
        vm.serializeAddress(adminKey, "pauserMultisig", addrs.pauserMultisig);
        vm.serializeAddress(adminKey, "pauserRegistry", addrs.pauserRegistry);
        vm.serializeAddress(adminKey, "proxyAdmin", addrs.proxyAdmin);
        string memory adminObject = vm.serializeAddress(adminKey, "timelock", addrs.timelock);

        // Core
        string memory coreKey = "core";
        _writeTUP(coreKey, "avsDirectory", addrs.avsDirectory);
        _writeTUP(coreKey, "delegationManager", addrs.delegationManager);
        _writeTUP(coreKey, "rewardsCoordinator", addrs.rewardsCoordinator);
        _writeTUP(coreKey, "slasher", addrs.slasher);
        string memory coreObject = _writeTUP(coreKey, "strategyManager", addrs.strategyManager);

        // Pods
        string memory podsKey = "pods";
        _writeBeacon(podsKey, "eigenPod", addrs.eigenPod);
        _writeTUP(podsKey, "eigenPodManager", addrs.eigenPodManager);
        string memory podsObject = _writeTUP(podsKey, "delayedWithdrawalRouter", addrs.delayedWithdrawalRouter);

        // Strategies
        string memory strategiesKey = "strategies";
        _writeTUP(strategiesKey, "strategyFactory", addrs.strategyFactory);
        _writeBeacon(strategiesKey, "strategyBeacon", addrs.strategyBeacon);
        string memory strategiesObject = _writeStrategies(strategiesKey, "preLongtailStrats", addrs.preLongtailStrats);

        // Token
        string memory tokenKey = "token";
        _writeToken(tokenKey, "EIGEN", addrs.EIGEN);
        _writeToken(tokenKey, "bEIGEN", addrs.bEIGEN);
        string memory tokenObject = _writeTUP(tokenKey, "eigenStrategy", addrs.eigenStrategy);

        // Deployment
        string memory deploymentKey = "deployment";
        vm.serializeString(deploymentKey, adminKey, adminObject);
        vm.serializeString(deploymentKey, coreKey, coreObject);
        vm.serializeString(deploymentKey, podsKey, podsObject);
        vm.serializeString(deploymentKey, strategiesKey, strategiesObject);
        string memory deploymentObject = vm.serializeString(deploymentKey, tokenKey, tokenObject);

        return (deploymentKey, deploymentObject);
    }

    function _writeConfig(Environment memory env, Params memory params) private returns (string memory, string memory) {
        // Environment
        string memory environmentKey = "environment";
        vm.serializeUint(environmentKey, "chainid", env.chainid);
        vm.serializeString(environmentKey, "name", env.name);
        string memory environmentObject = vm.serializeString(environmentKey, "lastUpdated", env.lastUpdated);

        // Params
        string memory paramsKey = "params";
        vm.serializeAddress(paramsKey, "multiSendCallOnly", params.multiSendCallOnly);
        vm.serializeAddress(paramsKey, "ethPOS", params.ethPOS);
        vm.serializeUint(paramsKey, "EIGENPOD_GENESIS_TIME", params.EIGENPOD_GENESIS_TIME);
        vm.serializeUint(paramsKey, "CALCULATION_INTERVAL_SECONDS", params.CALCULATION_INTERVAL_SECONDS);
        vm.serializeUint(paramsKey, "MAX_REWARDS_DURATION", params.MAX_REWARDS_DURATION);
        vm.serializeUint(paramsKey, "MAX_RETROACTIVE_LENGTH", params.MAX_RETROACTIVE_LENGTH);
        vm.serializeUint(paramsKey, "MAX_FUTURE_LENGTH", params.MAX_FUTURE_LENGTH);
        vm.serializeUint(paramsKey, "GENESIS_REWARDS_TIMESTAMP", params.GENESIS_REWARDS_TIMESTAMP);
        vm.serializeAddress(paramsKey, "REWARDS_UPDATER_ADDRESS", params.REWARDS_UPDATER_ADDRESS);
        vm.serializeUint(paramsKey, "ACTIVATION_DELAY", params.ACTIVATION_DELAY);
        string memory paramsObject = vm.serializeUint(paramsKey, "GLOBAL_OPERATOR_COMMISSION_BIPS", params.GLOBAL_OPERATOR_COMMISSION_BIPS);

        // Config
        string memory configKey = "config";
        vm.serializeString(configKey, environmentKey, environmentObject);
        string memory configObject = vm.serializeString(configKey, paramsKey, paramsObject);

        return (configKey, configObject);
    }

    function _writeTUP(string memory parentObjectKey, string memory tupKey, TUPInfo memory value) private returns (string memory) {
        vm.serializeAddress(tupKey, "proxy", value.proxy);
        vm.serializeAddress(tupKey, "impl", value.impl);
        string memory tupObject = vm.serializeAddress(tupKey, "pendingImpl", value.pendingImpl);

        return vm.serializeString(parentObjectKey, tupKey, tupObject);
    }

    function _writeBeacon(string memory parentObjectKey, string memory beaconKey, BeaconInfo memory value) private returns (string memory) {
        vm.serializeAddress(beaconKey, "beacon", value.beacon);
        vm.serializeAddress(beaconKey, "impl", value.impl);
        string memory beaconObject = vm.serializeAddress(beaconKey, "pendingImpl", value.pendingImpl);

        return vm.serializeString(parentObjectKey, beaconKey, beaconObject);
    }

    function _writeStrategies(string memory parentObjectKey, string memory strategiesKey, TUPInfo[] memory infos) private returns (string memory) {
        address impl = infos.length == 0 ? address(0) : infos[0].impl;

        address[] memory proxies = new address[](infos.length);
        for (uint i = 0; i < infos.length; i++) {
            proxies[i] = infos[i].proxy;
        }

        vm.serializeAddress(strategiesKey, "impl", impl);
        string memory strategiesObject = vm.serializeAddress(strategiesKey, "addrs", proxies);

        return vm.serializeString(parentObjectKey, strategiesKey, strategiesObject);
    }

    function _writeToken(string memory parentObjectKey, string memory tokenKey, TokenInfo memory value) private returns (string memory) {
        vm.serializeAddress(tokenKey, "proxy", value.proxy);
        vm.serializeAddress(tokenKey, "impl", value.impl);
        vm.serializeAddress(tokenKey, "pendingImpl", value.pendingImpl);
        string memory tokenObject = vm.serializeAddress(tokenKey, "proxyAdmin", value.proxyAdmin);

        return vm.serializeString(parentObjectKey, tokenKey, tokenObject);
    }
}