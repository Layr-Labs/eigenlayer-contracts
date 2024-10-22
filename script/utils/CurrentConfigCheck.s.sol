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

    address public protocolCouncilMultisig;
    TimelockController public protocolTimelockController;
    TimelockController public protocolTimelockController_BEIGEN;

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

        // vm.startPrank(foundationMultisig);
        // Ownable(address(bEIGEN)).transferOwnership(address(eigenTokenTimelockController));
        // Ownable(address(EIGEN)).transferOwnership(address(eigenTokenTimelockController));
        // vm.stopPrank();

        checkGovernanceConfiguration_Current();
    }

    // check governance configuration
    function checkGovernanceConfiguration_Current() public {
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

        require(eigenTokenTimelockController.hasRole(eigenTokenTimelockController.PROPOSER_ROLE(), foundationMultisig),
            "foundationMultisig does not have PROPOSER_ROLE on eigenTokenTimelockController");
        require(eigenTokenTimelockController.hasRole(eigenTokenTimelockController.EXECUTOR_ROLE(), foundationMultisig),
            "foundationMultisig does not have EXECUTOR_ROLE on eigenTokenTimelockController");
        require(eigenTokenTimelockController.hasRole(eigenTokenTimelockController.CANCELLER_ROLE(), operationsMultisig),
            "operationsMultisig does not have CANCELLER_ROLE on eigenTokenTimelockController");
        require(eigenTokenTimelockController.hasRole(eigenTokenTimelockController.TIMELOCK_ADMIN_ROLE(), executorMultisig),
            "executorMultisig does not have TIMELOCK_ADMIN_ROLE on eigenTokenTimelockController");
        require(eigenTokenTimelockController.hasRole(eigenTokenTimelockController.TIMELOCK_ADMIN_ROLE(), address(eigenTokenTimelockController)),
            "eigenTokenTimelockController does not have TIMELOCK_ADMIN_ROLE on itself");

        require(beigenTokenTimelockController.hasRole(beigenTokenTimelockController.PROPOSER_ROLE(), foundationMultisig),
            "foundationMultisig does not have PROPOSER_ROLE on beigenTokenTimelockController");
        require(beigenTokenTimelockController.hasRole(beigenTokenTimelockController.EXECUTOR_ROLE(), foundationMultisig),
            "foundationMultisig does not have EXECUTOR_ROLE on beigenTokenTimelockController");
        require(beigenTokenTimelockController.hasRole(beigenTokenTimelockController.CANCELLER_ROLE(), operationsMultisig),
            "operationsMultisig does not have CANCELLER_ROLE on beigenTokenTimelockController");
        require(beigenTokenTimelockController.hasRole(beigenTokenTimelockController.TIMELOCK_ADMIN_ROLE(), executorMultisig),
            "executorMultisig does not have TIMELOCK_ADMIN_ROLE on beigenTokenTimelockController");
        require(beigenTokenTimelockController.hasRole(beigenTokenTimelockController.TIMELOCK_ADMIN_ROLE(), address(beigenTokenTimelockController)),
            "beigenTokenTimelockController does not have TIMELOCK_ADMIN_ROLE on itself");
    }

    function checkGovernanceConfiguration_WithProtocolCouncil() public {
        assertEq(eigenLayerProxyAdmin.owner(), protocolCouncilMultisig,
            "eigenLayerProxyAdmin.owner() != protocolCouncilMultisig");
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
        bool protocolTimelockInOwners;
        bool communityMultisigInOwners;
        for (uint256 i = 0; i < 2; ++i) {
            if (executorMultisigOwners[i] == address(protocolTimelockController)) {
                protocolTimelockInOwners = true;
            }
            if (executorMultisigOwners[i] == communityMultisig) {
                communityMultisigInOwners = true;
            }
        }
        require(protocolTimelockInOwners, "timelock not in executorMultisig owners");
        require(communityMultisigInOwners, "communityMultisig not in executorMultisig owners");

        require(eigenTokenProxyAdmin != beigenTokenProxyAdmin,
            "tokens must have different proxy admins to allow different timelock controllers");
        require(protocolTimelockController != protocolTimelockController_BEIGEN,
            "tokens must have different timelock controllers");

        // note that proxy admin owners are different but _token_ owners are the same
        assertEq(Ownable(address(EIGEN)).owner(), address(protocolTimelockController),
            "EIGEN.owner() != protocolTimelockController");
        assertEq(Ownable(address(bEIGEN)).owner(), address(protocolTimelockController),
            "bEIGEN.owner() != protocolTimelockController");
        assertEq(eigenTokenProxyAdmin.owner(), address(protocolTimelockController),
            "eigenTokenProxyAdmin.owner() != protocolTimelockController");
        assertEq(beigenTokenProxyAdmin.owner(), address(protocolTimelockController_BEIGEN),
            "beigenTokenProxyAdmin.owner() != protocolTimelockController_BEIGEN");

        assertEq(eigenTokenProxyAdmin.getProxyAdmin(TransparentUpgradeableProxy(payable(address(EIGEN)))),
            address(eigenTokenProxyAdmin),
            "eigenTokenProxyAdmin is not actually the admin of the EIGEN token");
        assertEq(beigenTokenProxyAdmin.getProxyAdmin(TransparentUpgradeableProxy(payable(address(bEIGEN)))),
            address(beigenTokenProxyAdmin),
            "beigenTokenProxyAdmin is not actually the admin of the bEIGEN token");

        require(protocolTimelockController.hasRole(protocolTimelockController.PROPOSER_ROLE(), protocolCouncilMultisig),
            "protocolCouncilMultisig does not have PROPOSER_ROLE on protocolTimelockController");
        require(protocolTimelockController.hasRole(protocolTimelockController.EXECUTOR_ROLE(), protocolCouncilMultisig),
            "protocolCouncilMultisig does not have EXECUTOR_ROLE on protocolTimelockController");
        require(protocolTimelockController.hasRole(protocolTimelockController.PROPOSER_ROLE(), operationsMultisig),
            "operationsMultisig does not have PROPOSER_ROLE on protocolTimelockController");
        require(protocolTimelockController.hasRole(protocolTimelockController.CANCELLER_ROLE(), operationsMultisig),
            "operationsMultisig does not have CANCELLER_ROLE on protocolTimelockController");
        require(protocolTimelockController.hasRole(protocolTimelockController.TIMELOCK_ADMIN_ROLE(), executorMultisig),
            "executorMultisig does not have TIMELOCK_ADMIN_ROLE on protocolTimelockController");
        require(protocolTimelockController.hasRole(protocolTimelockController.TIMELOCK_ADMIN_ROLE(), address(protocolTimelockController)),
            "protocolTimelockController does not have TIMELOCK_ADMIN_ROLE on itself");

        require(protocolTimelockController_BEIGEN.hasRole(protocolTimelockController_BEIGEN.PROPOSER_ROLE(), protocolCouncilMultisig),
            "protocolCouncilMultisig does not have PROPOSER_ROLE on protocolTimelockController_BEIGEN");
        require(protocolTimelockController_BEIGEN.hasRole(protocolTimelockController_BEIGEN.EXECUTOR_ROLE(), protocolCouncilMultisig),
            "protocolCouncilMultisig does not have EXECUTOR_ROLE on protocolTimelockController_BEIGEN");
        require(protocolTimelockController_BEIGEN.hasRole(protocolTimelockController_BEIGEN.PROPOSER_ROLE(), operationsMultisig),
            "operationsMultisig does not have PROPOSER_ROLE on protocolTimelockController_BEIGEN");
        require(protocolTimelockController_BEIGEN.hasRole(protocolTimelockController_BEIGEN.CANCELLER_ROLE(), operationsMultisig),
            "operationsMultisig does not have CANCELLER_ROLE on protocolTimelockController_BEIGEN");
        require(protocolTimelockController_BEIGEN.hasRole(protocolTimelockController_BEIGEN.TIMELOCK_ADMIN_ROLE(), executorMultisig),
            "executorMultisig does not have TIMELOCK_ADMIN_ROLE on protocolTimelockController_BEIGEN");
        require(protocolTimelockController_BEIGEN.hasRole(protocolTimelockController_BEIGEN.TIMELOCK_ADMIN_ROLE(), address(protocolTimelockController_BEIGEN)),
            "beigenTokenTimelockController does not have TIMELOCK_ADMIN_ROLE on itself");
    }

    // forge script script/utils/CurrentConfigCheck.s.sol:CurrentConfigCheck -vvv --sig "simulateProtocolCouncilUpgrade(string)" $NETWORK_NAME
    function simulateProtocolCouncilUpgrade(string memory networkName) public virtual {
        run(networkName);
        simulateTimelockControllerDeploy();
        checkGovernanceConfiguration_WithProtocolCouncil();
    }

    function simulateTimelockControllerDeploy() public {
        // set up initially with sender and foundation being proposers/executor/canceller on timelock controller 
        // so that team can be added as only canceller
        address[] memory proposers = new address[](3);
        proposers[0] = protocolCouncilMultisig;
        proposers[1] = msg.sender;
        proposers[2] = operationsMultisig;

        address[] memory executors = new address[](2);
        executors[0] = protocolCouncilMultisig;
        executors[1] = msg.sender;

        vm.startBroadcast();
        protocolTimelockController = 
            new TimelockController(
                0, // no delay for setup
                proposers, 
                executors
            );
        protocolTimelockController_BEIGEN = 
            new TimelockController(
                0, // no delay for setup
                proposers, 
                executors
            );

        uint256 tx_array_length = 6;
        address[] memory targets = new address[](tx_array_length);
        for (uint256 i = 0; i < targets.length; ++i) {
            targets[i] = address(protocolTimelockController);
        }

        uint256[] memory values = new uint256[](tx_array_length);

        bytes[] memory payloads = new bytes[](tx_array_length);
        // 1. add operationsMultisig as canceller
        payloads[0] = abi.encodeWithSelector(AccessControl.grantRole.selector, protocolTimelockController.CANCELLER_ROLE(), operationsMultisig);
        // 2. remove sender as canceller
        payloads[1] = abi.encodeWithSelector(AccessControl.revokeRole.selector, protocolTimelockController.CANCELLER_ROLE(), msg.sender);
        // 3. remove sender as executor
        payloads[2] = abi.encodeWithSelector(AccessControl.revokeRole.selector, protocolTimelockController.EXECUTOR_ROLE(), msg.sender);
        // 4. remove sender as proposer
        payloads[3] = abi.encodeWithSelector(AccessControl.revokeRole.selector, protocolTimelockController.PROPOSER_ROLE(), msg.sender);
        // 5. remove sender as admin
        payloads[4] = abi.encodeWithSelector(AccessControl.revokeRole.selector, protocolTimelockController.PROPOSER_ROLE(), msg.sender);
        // 6. set min delay to appropriate length
        payloads[5] = abi.encodeWithSelector(protocolTimelockController.updateDelay.selector, 10 days);

        // schedule the batch
        protocolTimelockController.scheduleBatch(
            targets, 
            values, 
            payloads, 
            bytes32(0), // no predecessor needed
            bytes32(0), // no salt 
            0 // 0 enforced delay
        );

        // execute the batch
        protocolTimelockController.executeBatch(
            targets, 
            values, 
            payloads, 
            bytes32(0), // no predecessor needed
            bytes32(0) // no salt
        );

        for (uint256 i = 0; i < targets.length; ++i) {
            targets[i] = address(protocolTimelockController_BEIGEN);
        }

        // schedule the batch
        protocolTimelockController_BEIGEN.scheduleBatch(
            targets, 
            values, 
            payloads, 
            bytes32(0), // no predecessor needed
            bytes32(0), // no salt 
            0 // 0 enforced delay
        );

        // execute the batch
        protocolTimelockController_BEIGEN.executeBatch(
            targets, 
            values, 
            payloads, 
            bytes32(0), // no predecessor needed
            bytes32(0) // no salt
        );
        vm.stopBroadcast();

        // TODO
        // give proxy admin ownership to timelock controller
        // eigenTokenProxyAdmin.transferOwnership(address(timelockController));

        // // TODO: fix this
        // bytes memory data_transferOwnershipToExecutorMultisig = abi.encodeWithSelector(Ownable.transferOwnership.selector, address(executorMultisig));
        // bytes memory callToExecutor = encodeForExecutor({
        //     from: timelock,
        //     to: address(eigenTokenProxyAdmin),
        //     value: 0,
        //     data: data_transferOwnershipToExecutorMultisig,
        //     operation: ISafe.Operation.Call 
        // });
        // // uint256 timelockEta = block.timestamp + timelock.delay();
        // uint256 timelockEta = block.timestamp + 10 days;
        // (bytes memory calldata_to_timelock_queuing_action, bytes memory calldata_to_timelock_executing_action) =
        //     encodeForTimelock({
        //         to: address(executorMultisig),
        //         value: 0,
        //         data: callToExecutor,
        //         timelockEta: timelockEta
        //     });
        // vm.startPrank(operationsMultisig);

        // (bool success, /*bytes memory returndata*/) = timelock.call(calldata_to_timelock_queuing_action);
        // require(success, "call to timelock queuing action failed");
        // callToExecutor = encodeForExecutor({
        //     from: timelock,
        //     to: address(beigenTokenProxyAdmin),
        //     value: 0,
        //     data: data_transferOwnershipToExecutorMultisig,
        //     operation: ISafe.Operation.Call 
        // });
        // (bytes memory calldata_to_timelock_queuing_action_2, bytes memory calldata_to_timelock_executing_action_2) =
        //     encodeForTimelock({
        //         to: address(executorMultisig),
        //         value: 0,
        //         data: callToExecutor,
        //         timelockEta: timelockEta
        //     });
        // (success, /*bytes memory returndata*/) = timelock.call(calldata_to_timelock_queuing_action_2);
        // require(success, "call to timelock queuing action 2 failed");

        // vm.warp(timelockEta);
        // vm.prank(operationsMultisig);
        // (success, /*bytes memory returndata*/) = timelock.call(calldata_to_timelock_executing_action);
        // require(success, "call to timelock executing action failed");
        // (success, /*bytes memory returndata*/) = timelock.call(calldata_to_timelock_executing_action_2);
        // require(success, "call to timelock executing action 2 failed");
    }
}
