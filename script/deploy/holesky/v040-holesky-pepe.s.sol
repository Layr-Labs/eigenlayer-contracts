// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../utils/ExistingDeploymentParser.sol";

/**
 * @notice Script used for upgrading EigenPod and EPM Implementation for Holesky preprod
 * anvil --fork-url $RPC_HOLESKY
 * forge script script/deploy/holesky/v040-holesky-pepe.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * forge script script/deploy/holesky/v040-holesky-pepe.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --verify --broadcast -vvvv
 */
contract PEPE_Deploy_Preprod is ExistingDeploymentParser {

    address testAddress = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;
    address initOwner = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;

    function run() external virtual {
        _parseInitialDeploymentParams(
            "script/configs/holesky/Deploy_RewardsCoordinator.holesky.config.json"
        );
        _parseDeployedContracts(
            "script/configs/holesky/eigenlayer_addresses.config.json"
        );

        emit log_named_address("Deployer Address", msg.sender);

        emit log("PRIOR IMPLEMENTATION");
        emit log_named_address("current pod impl", address(eigenPodImplementation));
        emit log_named_address("- pod.ethPOS", address(eigenPodImplementation.ethPOS()));
        emit log_named_address("- pod.eigenPodManager", address(eigenPodImplementation.eigenPodManager()));
        emit log_named_uint("- pod.GENESIS_TIME", eigenPodImplementation.GENESIS_TIME());
        emit log_named_address("current manager impl", address(eigenPodManagerImplementation));
        emit log_named_address("- epm.ethPOS", address(eigenPodManagerImplementation.ethPOS()));
        emit log_named_address("- epm.eigenPodBeacon", address(eigenPodManagerImplementation.eigenPodBeacon()));
        emit log_named_address("- epm.strategyManager", address(eigenPodManagerImplementation.strategyManager()));
        emit log_named_address("- epm.slasher", address(eigenPodManagerImplementation.slasher()));
        emit log_named_address("- epm.delegationManager", address(eigenPodManagerImplementation.delegationManager()));

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        _deployPEPE();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        logAndOutputContractAddresses("script/output/holesky/v040.output.json");
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
        // upgrade UpgradeableBeacon
        eigenPodBeacon.upgradeTo(address(eigenPodImplementation));

        // upgrade TUPS
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation)
        );
    }
}
