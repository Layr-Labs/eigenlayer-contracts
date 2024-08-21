// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/utils/Create2.sol";
import "../../utils/ExistingDeploymentParser.sol";

/**
 * FORK LOCAL
 * anvil --fork-url $RPC_MAINNET
 * forge script script/deploy/mainnet/v0.4.2-mainnet-pepe.s.sol:MainnetPEPEDeploy --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 *
 * MAINNET
 * forge script script/deploy/mainnet/v0.4.2-mainnet-pepe.s.sol:MainnetPEPEDeploy --rpc-url $RPC_MAINNET --private-key $PRIVATE_KEY --verify --broadcast -vvvv
 *
 */
contract MainnetPEPEDeploy is ExistingDeploymentParser {
    function run() external virtual {
        _parseInitialDeploymentParams(
            "script/configs/mainnet/mainnet-config.config.json"
        );
        _parseDeployedContracts(
            "script/configs/mainnet/mainnet-addresses.config.json"
        );

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        emit log_named_address("Deployer Address", msg.sender);

        _deployPEPE();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        _upgradePEPE();

        _testDeploy();

        // Post-upgrade sanity checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized();
        _verifyInitializationParams();

        logAndOutputContractAddresses("script/output/mainnet/v0.4.2-mainnet-pepe.output.json");
    }

    function _deployPEPE() internal {
        // Deploy EigenPod
        eigenPodImplementation = new EigenPod(
            IETHPOSDeposit(ETHPOSDepositAddress),
            eigenPodManager,
            EIGENPOD_GENESIS_TIME
        );

        // Deploy EigenPodManager
        eigenPodManagerImplementation = new EigenPodManager(
            IETHPOSDeposit(ETHPOSDepositAddress),
            eigenPodBeacon,
            strategyManager,
            slasher,
            delegationManager
        );
    }

    function _upgradePEPE() internal {
        vm.startPrank(address(executorMultisig));

        // upgrade UpgradeableBeacon
        eigenPodBeacon.upgradeTo(address(eigenPodImplementation));

        // upgrade TUPS
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation)
        );

        vm.stopPrank();
    }

    function _testDeploy() internal {
        require(eigenPodImplementation.activeValidatorCount() == 0, "unable to fetch activeValidatorCount");
    }
}
