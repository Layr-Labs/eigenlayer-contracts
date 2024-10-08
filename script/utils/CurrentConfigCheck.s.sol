// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./ExistingDeploymentParser.sol";
import "./TimelockEncoding.sol";

/**
 * forge script script/utils/CurrentConfigCheck.s.sol:CurrentConfigCheck -vvv --sig "run(string)" $NETWORK_NAME
 * NETWORK_NAME options are currently preprod-holesky, testnet-holesky, mainnet, local
 */
contract CurrentConfigCheck is ExistingDeploymentParser, TimelockEncoding {
    string deployedContractsConfig;
    string intialDeploymentParams;
    string forkUrl;  
    string emptyString;

    function run(string memory networkName) public virtual {
        if (keccak256(abi.encodePacked(networkName)) == keccak256(abi.encodePacked("preprod-holesky"))) {
            deployedContractsConfig = "script/configs/holesky/eigenlayer_addresses_preprod.config.json";
            intialDeploymentParams = "script/configs/holesky/eigenlayer_preprod.config.json";
            forkUrl = vm.envString("RPC_HOLESKY");
            uint256 forkId = vm.createFork(forkUrl);
            vm.selectFork(forkId);
        } else if (keccak256(abi.encodePacked(networkName)) == keccak256(abi.encodePacked("testnet-holesky"))) {
            deployedContractsConfig = "script/configs/holesky/eigenlayer_addresses_testnet.config.json";
            intialDeploymentParams = "script/configs/holesky/eigenlayer_testnet.config.json";
            forkUrl = vm.envString("RPC_HOLESKY");
            uint256 forkId = vm.createFork(forkUrl);
            vm.selectFork(forkId);
        } else if (keccak256(abi.encodePacked(networkName)) == keccak256(abi.encodePacked("mainnet"))) {
            deployedContractsConfig = "script/configs/mainnet/mainnet-addresses.config.json";
            intialDeploymentParams = "script/configs/mainnet/mainnet-config.config.json"; 
            forkUrl = vm.envString("RPC_MAINNET");
            uint256 forkId = vm.createFork(forkUrl);
            vm.selectFork(forkId);
        }

        if (keccak256(abi.encodePacked(networkName)) == keccak256(abi.encodePacked("local"))) {
            deployedContractsConfig = "script/output/devnet/local_from_scratch_deployment_data.json";
            intialDeploymentParams = "script/configs/local/deploy_from_scratch.anvil.config.json";             
        }

        require(keccak256(abi.encodePacked(deployedContractsConfig)) != keccak256(abi.encodePacked(emptyString)),
            "deployedContractsConfig cannot be unset");
        require(keccak256(abi.encodePacked(intialDeploymentParams)) != keccak256(abi.encodePacked(emptyString)),
            "intialDeploymentParams cannot be unset");

        // read and log the chainID
        uint256 currentChainId = block.chainid;
        emit log_named_uint("You are parsing on ChainID", currentChainId);
        require(currentChainId == 1 || currentChainId == 17000 || currentChainId == 31337,
            "script is only for mainnet or holesky or local environment");

        _parseDeployedContracts(deployedContractsConfig);
        _parseInitialDeploymentParams(intialDeploymentParams);

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized(false);
        _verifyInitializationParams();

        // bytes memory data = abi.encodeWithSelector(ProxyAdmin.changeProxyAdmin.selector, address(bEIGEN), address(beigenTokenProxyAdmin));
        // bytes memory data = abi.encodeWithSignature("swapOwner(address,address,address)",
        //     0xCb8d2f9e55Bc7B1FA9d089f9aC80C583D2BDD5F7,
        //     0xcF19CE0561052a7A7Ff21156730285997B350A7D,
        //     0xFddd03C169E3FD9Ea4a9548dDC4BedC6502FE239
        // );
        // bytes memory callToExecutor = encodeForExecutor({
        //     from: communityMultisig,
        //     to: address(executorMultisig),
        //     value: 0,
        //     data: data,
        //     operation: ISafe.Operation.Call 
        // });

        // vm.prank(communityMultisig);
        // (bool success, /*bytes memory returndata*/) = executorMultisig.call(callToExecutor);
        // require(success, "call to executorMultisig failed");

        vm.startPrank(foundationMultisig);
        Ownable(address(bEIGEN)).transferOwnership(address(eigenTokenTimelockController));
        Ownable(address(EIGEN)).transferOwnership(address(eigenTokenTimelockController));
        vm.stopPrank();

        checkDesiredGovernanceConfiguration();
    }

    // check governance configuration
    function checkDesiredGovernanceConfiguration() public {
        assertEq(eigenLayerProxyAdmin.owner(), executorMultisig,
            "eigenLayerProxyAdmin.owner() != executorMultisig");
        assertEq(delegationManager.owner(), executorMultisig,
            "delegationManager.owner() != executorMultisig");
        assertEq(strategyManager.owner(), executorMultisig,
            "strategyManager.owner() != executorMultisig");
        assertEq(strategyManager.strategyWhitelister(), address(strategyFactory),
            "strategyManager.strategyWhitelister() != address(strategyFactory)");
        assertEq(strategyFactory.owner(), operationsMultisig,
            "strategyFactory.owner() != operationsMultisig");
        assertEq(avsDirectory.owner(), executorMultisig,
            "avsDirectory.owner() != executorMultisig");
        assertEq(rewardsCoordinator.owner(), operationsMultisig,
            "rewardsCoordinator.owner() != operationsMultisig");
        assertEq(eigenLayerPauserReg.unpauser(), executorMultisig,
            "eigenLayerPauserReg.unpauser() != operationsMultisig");
        require(eigenLayerPauserReg.isPauser(operationsMultisig),
            "operationsMultisig does not have pausing permissions");
        require(eigenLayerPauserReg.isPauser(executorMultisig),
            "executorMultisig does not have pausing permissions");
        require(eigenLayerPauserReg.isPauser(pauserMultisig),
            "pauserMultisig does not have pausing permissions");

        (bool success, bytes memory returndata) = timelock.staticcall(abi.encodeWithSignature("admin()"));
        require(success, "call to timelock.admin() failed");
        address timelockAdmin = abi.decode(returndata, (address));
        assertEq(timelockAdmin, operationsMultisig,
            "timelockAdmin != operationsMultisig");

        (success, returndata) = executorMultisig.staticcall(abi.encodeWithSignature("getOwners()"));
        require(success, "call to executorMultisig.getOwners() failed");
        address[] memory executorMultisigOwners = abi.decode(returndata, (address[]));
        require(executorMultisigOwners.length == 2,
            "executorMultisig owners wrong length");
        bool timelockInOwners;
        bool communityMultisigInOwners;
        for (uint256 i = 0; i < 2; ++i) {
            if (executorMultisigOwners[i] == timelock) {
                timelockInOwners = true;
            }
            if (executorMultisigOwners[i] == communityMultisig) {
                communityMultisigInOwners = true;
            }
        }
        require(timelockInOwners, "timelock not in executorMultisig owners");
        require(communityMultisigInOwners, "communityMultisig not in executorMultisig owners");

        require(eigenTokenProxyAdmin != beigenTokenProxyAdmin,
            "tokens must have different proxy admins to allow different timelock controllers");
        require(eigenTokenTimelockController != beigenTokenTimelockController,
            "tokens must have different timelock controllers");

        // note that proxy admin owners are different but _token_ owners are the same
        assertEq(Ownable(address(EIGEN)).owner(), address(eigenTokenTimelockController),
            "EIGEN.owner() != eigenTokenTimelockController");
        assertEq(Ownable(address(bEIGEN)).owner(), address(eigenTokenTimelockController),
            "bEIGEN.owner() != eigenTokenTimelockController");
        assertEq(eigenTokenProxyAdmin.owner(), address(eigenTokenTimelockController),
            "eigenTokenProxyAdmin.owner() != eigenTokenTimelockController");
        assertEq(beigenTokenProxyAdmin.owner(), address(beigenTokenTimelockController),
            "beigenTokenProxyAdmin.owner() != beigenTokenTimelockController");

        assertEq(eigenTokenProxyAdmin.getProxyAdmin(TransparentUpgradeableProxy(payable(address(EIGEN)))),
            address(eigenTokenProxyAdmin),
            "eigenTokenProxyAdmin is not actually the admin of the EIGEN token");
        assertEq(beigenTokenProxyAdmin.getProxyAdmin(TransparentUpgradeableProxy(payable(address(bEIGEN)))),
            address(beigenTokenProxyAdmin),
            "beigenTokenProxyAdmin is not actually the admin of the bEIGEN token");
    }
}
