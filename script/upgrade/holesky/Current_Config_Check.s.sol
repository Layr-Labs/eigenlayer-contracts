// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../utils/ExistingDeploymentParser.sol";
import "../../utils/TimelockEncoding.sol";

/**
 * forge script script/upgrade/holesky/Current_Config_Check.s.sol:Current_Config_Check -vvv
 * 
 */
contract Current_Config_Check is ExistingDeploymentParser, TimelockEncoding {
    string deployedContractsConfig = "script/configs/holesky/eigenlayer_addresses_testnet.config.json";
    string intialDeploymentParams = "script/configs/holesky/eigenlayer_testnet.config.json";
    function run() external virtual {
        string memory forkUrl = vm.envString("RPC_HOLESKY");
        uint256 forkId = vm.createFork(forkUrl);
        vm.selectFork(forkId);

        // read and log the chainID
        uint256 currentChainId = block.chainid;
        emit log_named_uint("You are parsing on ChainID", currentChainId);
        require(currentChainId == 17000, "script is only for mainnet");

        _parseDeployedContracts(deployedContractsConfig);
        _parseInitialDeploymentParams(intialDeploymentParams);

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized(false);
        _verifyInitializationParams();

        // bytes memory data = abi.encodeWithSelector(strategyFactory.transferOwnership.selector, operationsMultisig);
        // bytes memory callToExecutor = encodeForExecutor({
        //     from: communityMultisig,
        //     to: address(strategyFactory),
        //     value: 0,
        //     data: data,
        //     operation: ISafe.Operation.Call 
        // });

        // vm.prank(communityMultisig);
        // (bool success, /*bytes memory returndata*/) = executorMultisig.call(callToExecutor);
        // require(success, "call to executorMultisig failed");

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
    }
}
