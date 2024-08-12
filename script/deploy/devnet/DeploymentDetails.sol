// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.12;
// solhint-disable no-console

import {stdJson} from "forge-std/StdJson.sol";
import {console2} from "forge-std/console2.sol";
import {VmSafe} from "forge-std/Vm.sol";

struct ContractDeployment {
    address avsDirectory;
    address avsDirectoryImplementation;
    address baseStrategyImplementation;
    address delegationManager;
    address delegationManagerImplementation;
    address eigenPodBeacon;
    address eigenPodImplementation;
    address eigenPodManager;
    address eigenPodManagerImplementation;
    address emptyContract;
    address pauserRegistry;
    address proxyAdmin;
    address rewardsCoordinator;
    address rewardsCoordinatorImplementation;
    address slasher;
    address slasherImplementation;
    address strategyManager;
    address strategyManagerImplementation;
    address ethDepositAddress;
}

library DeploymentDetails {
    VmSafe private constant vm = VmSafe(address(uint160(uint256(keccak256("hevm cheat code")))));

    function write(
        ContractDeployment calldata deploymentDetails,
        address executorMultisig,
        address operationsMultisig,
        address nilContract
    ) external {
        // WRITE JSON DATA
        string memory parentObject = "parent object";

        string memory deployedAddresses = "address";
        vm.serializeAddress(deployedAddresses, "proxyAdmin", deploymentDetails.proxyAdmin);
        vm.serializeAddress(deployedAddresses, "pauserRegistry", deploymentDetails.pauserRegistry);
        vm.serializeAddress(deployedAddresses, "slasher", deploymentDetails.slasher);
        vm.serializeAddress(deployedAddresses, "slasherImplementation", deploymentDetails.slasherImplementation);
        vm.serializeAddress(deployedAddresses, "delegationManager", deploymentDetails.delegationManager);
        vm.serializeAddress(
            deployedAddresses,
            "delegationManagerImplementation",
            deploymentDetails.delegationManagerImplementation
        );
        vm.serializeAddress(deployedAddresses, "avsDirectory", deploymentDetails.avsDirectory);
        vm.serializeAddress(
            deployedAddresses,
            "avsDirectoryImplementation",
            deploymentDetails.avsDirectoryImplementation
        );
        vm.serializeAddress(deployedAddresses, "strategyManager", deploymentDetails.strategyManager);
        vm.serializeAddress(
            deployedAddresses,
            "strategyManagerImplementation",
            deploymentDetails.strategyManagerImplementation
        );
        vm.serializeAddress(deployedAddresses, "eigenPodManager", deploymentDetails.eigenPodManager);
        vm.serializeAddress(
            deployedAddresses,
            "eigenPodManagerImplementation",
            deploymentDetails.eigenPodManagerImplementation
        );
        vm.serializeAddress(deployedAddresses, "rewardsCoordinator", deploymentDetails.rewardsCoordinator);
        vm.serializeAddress(
            deployedAddresses,
            "rewardsCoordinatorImplementation",
            deploymentDetails.rewardsCoordinatorImplementation
        );
        vm.serializeAddress(deployedAddresses, "eigenPodBeacon", deploymentDetails.eigenPodBeacon);
        vm.serializeAddress(deployedAddresses, "eigenPodImplementation", deploymentDetails.eigenPodImplementation);
        vm.serializeAddress(
            deployedAddresses,
            "baseStrategyImplementation",
            deploymentDetails.baseStrategyImplementation
        );
        vm.serializeAddress(deployedAddresses, "emptyContract", nilContract);
        vm.serializeAddress(deployedAddresses, "ethDepositAddress", deploymentDetails.ethDepositAddress);

        string memory deployedAddressesOutput = vm.serializeString(deployedAddresses, "strategies", "");

        string memory parameters = "parameters";
        vm.serializeAddress(parameters, "executorMultisig", executorMultisig);
        string memory parametersOutput = vm.serializeAddress(parameters, "operationsMultisig", operationsMultisig);

        string memory chainInfo = "chainInfo";
        vm.serializeUint(chainInfo, "deploymentBlock", block.number);
        string memory chainInfo_output = vm.serializeUint(chainInfo, "chainId", block.chainid);

        // serialize all the data
        vm.serializeString(parentObject, deployedAddresses, deployedAddressesOutput);
        vm.serializeString(parentObject, chainInfo, chainInfo_output);
        string memory finalJson = vm.serializeString(parentObject, parameters, parametersOutput);
        // TODO: should output to different file depending on configFile passed to run()
        //       so that we don't override mainnet output by deploying to goerli for eg.
        vm.writeJson(finalJson, "script/output/devnet/M2_deploy_from_scratch_deployment_data.json");
        console2.log("Deployment data written to script/output/devnet/M2_deploy_from_scratch_deployment_data.json");
    }

    function read(string memory deploymentFileName) external returns (ContractDeployment memory) {
        string memory deployDetailsPath = string(bytes(string.concat("script/output/devnet/", deploymentFileName)));
        string memory deployData = vm.readFile(deployDetailsPath);
        ContractDeployment memory deploymentDetails;
        deploymentDetails.avsDirectory = stdJson.readAddress(deployData, ".address.avsDirectory");
        deploymentDetails.avsDirectoryImplementation = stdJson.readAddress(
            deployData,
            ".address.avsDirectoryImplementation"
        );
        deploymentDetails.baseStrategyImplementation = stdJson.readAddress(
            deployData,
            ".address.baseStrategyImplementation"
        );
        deploymentDetails.delegationManager = stdJson.readAddress(deployData, ".address.delegationManager");
        deploymentDetails.delegationManagerImplementation = stdJson.readAddress(
            deployData,
            ".address.delegationManagerImplementation"
        );
        deploymentDetails.eigenPodBeacon = stdJson.readAddress(deployData, ".address.eigenPodBeacon");
        deploymentDetails.eigenPodImplementation = stdJson.readAddress(deployData, ".address.eigenPodImplementation");
        deploymentDetails.eigenPodManager = stdJson.readAddress(deployData, ".address.eigenPodManager");
        deploymentDetails.eigenPodManagerImplementation = stdJson.readAddress(
            deployData,
            ".address.eigenPodManagerImplementation"
        );
        deploymentDetails.emptyContract = stdJson.readAddress(deployData, ".address.emptyContract");
        deploymentDetails.pauserRegistry = stdJson.readAddress(deployData, ".address.pauserRegistry");
        deploymentDetails.proxyAdmin = stdJson.readAddress(deployData, ".address.proxyAdmin");
        deploymentDetails.rewardsCoordinator = stdJson.readAddress(deployData, ".address.rewardsCoordinator");
        deploymentDetails.rewardsCoordinatorImplementation = stdJson.readAddress(
            deployData,
            ".address.rewardsCoordinatorImplementation"
        );
        deploymentDetails.slasher = stdJson.readAddress(deployData, ".address.slasher");
        deploymentDetails.slasherImplementation = stdJson.readAddress(deployData, ".address.slasherImplementation");
        deploymentDetails.strategyManager = stdJson.readAddress(deployData, ".address.strategyManager");
        deploymentDetails.strategyManagerImplementation = stdJson.readAddress(
            deployData,
            ".address.strategyManagerImplementation"
        );
        deploymentDetails.ethDepositAddress = stdJson.readAddress(deployData, ".address.ethDepositAddress");

        return deploymentDetails;
    }
}
