// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../utils/ExistingDeploymentParser.sol";

/**
 * @notice Script used for upgrading EigenPod and EPM Implementation for Holesky preprod
 * anvil --fork-url $RPC_HOLESKY
 * forge script script/deploy/holesky/v040-rc3-holesky-preprod-pepe.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * forge script script/deploy/holesky/v040-rc3-holesky-preprod-pepe.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --verify --broadcast -vvvv
 */
contract EigenPod_Checkpoint_Deploy_Preprod is ExistingDeploymentParser {

    address testAddress = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;
    address initOwner = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;

    function run() external virtual {
        _parseInitialDeploymentParams(
            "script/configs/holesky/eigenlayer_preprod.config.json"
        );
        _parseDeployedContracts(
            "script/configs/holesky/eigenlayer_addresses_preprod.config.json"
        );

        emit log_named_address("Deployer Address", msg.sender);

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        _upgradeEigenPods();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized();
        _verifyInitializationParams();

        logAndOutputContractAddresses("script/output/holesky/v040-rc3.output.json");
    }

    /**
     * @notice Deploy EigenPod Implementation for Holesky preprod and upgrade the beacon
     */
    function _upgradeEigenPods() internal {
        // Deploy implementation
        eigenPodImplementation = new EigenPod(
            IETHPOSDeposit(ETHPOSDepositAddress),
            eigenPodManager,
            EIGENPOD_GENESIS_TIME
        );

        // upgrade UpgradeableBeacon
        eigenPodBeacon.upgradeTo(address(eigenPodImplementation));
    }
}
