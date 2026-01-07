// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import {DeployImplementations} from "./1-deployImplementations.s.sol";
import {CoreUpgradeQueueBuilder} from "../CoreUpgradeQueueBuilder.sol";
import "../Env.sol";
import "../TestUtils.sol";

import {EmissionsController} from "src/contracts/core/EmissionsController.sol";

/// Purpose: Deploy EmissionsController proxy and upgrade both EmissionsController and RewardsCoordinator
contract UpgradeContracts is DeployImplementations, MultisigBuilder {
    using Env for *;
    using Encode for *;
    using CoreUpgradeQueueBuilder for MultisigCall[];

    function _runAsMultisig() internal virtual override prank(Env.executorMultisig()) {
        // 1. Deploy EmissionsController proxy with empty implementation
        ITransparentUpgradeableProxy emissionsControllerProxy = ITransparentUpgradeableProxy(
            payable(new TransparentUpgradeableProxy({
                    _logic: address(Env.impl.emptyContract()),
                    admin_: Env.proxyAdmin(),
                    _data: ""
                }))
        );

        // Save the proxy contract to the env
        _unsafeAddProxyContract(type(EmissionsController).name, address(emissionsControllerProxy));

        // Create multisig calls for upgrades
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls();

        // 2. Upgrade EmissionsController proxy to the new implementation and initialize
        executorCalls.upgradeAndInitializeEmissionsController({
            initialOwner: Env.executorMultisig(),
            initialIncentiveCouncil: Env.communityMultisig(),
            initialPausedStatus: 0
        });

        // 3. Upgrade RewardsCoordinator to the new implementation
        executorCalls.upgradeRewardsCoordinator();

        // Execute all calls via multisig
        _unsafeExecuteMultisigCalls(executorCalls);
    }

    /// @dev Helper to execute multisig calls directly
    function _unsafeExecuteMultisigCalls(
        MultisigCall[] storage calls
    ) internal {
        for (uint256 i = 0; i < calls.length; i++) {
            (bool success, bytes memory returnData) = calls[i].to.call(calls[i].data);
            require(success, string(abi.encodePacked("Multisig call failed: ", returnData)));
        }
    }

    function testScript() public virtual override {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        // Check if already upgraded
        if (_areContractsUpgraded()) {
            return;
        }

        // 1. Deploy implementations as EOA (from previous step)
        _mode = OperationalMode.EOA;
        DeployImplementations._runAsEOA();

        // 2. Deploy proxy and upgrade contracts via multisig
        execute();

        // 3. Validate proxy was deployed
        assertTrue(
            address(Env.proxy.emissionsController()).code.length > 0, "EmissionsController proxy should be deployed"
        );

        // Validate proxy admin is correct
        assertTrue(
            Env.getProxyAdminBySlot(address(Env.proxy.emissionsController())) == Env.proxyAdmin(),
            "EmissionsController proxyAdmin should be proxyAdmin"
        );

        // 4. Validate upgrades
        TestUtils.validateEmissionsControllerInitialized(Env.proxy.emissionsController());
        TestUtils.validateRewardsCoordinatorConstructor(Env.impl.rewardsCoordinator());

        // Validate EmissionsController implementation is correct
        address emissionsControllerImpl = Env.getProxyImplementationBySlot(address(Env.proxy.emissionsController()));
        assertTrue(
            emissionsControllerImpl == address(Env.impl.emissionsController()),
            "EmissionsController implementation should match"
        );

        // Validate RewardsCoordinator implementation is correct
        address rewardsCoordinatorImpl = Env.getProxyImplementationBySlot(address(Env.proxy.rewardsCoordinator()));
        assertTrue(
            rewardsCoordinatorImpl == address(Env.impl.rewardsCoordinator()),
            "RewardsCoordinator implementation should match"
        );
    }

    /// @dev Check if contracts are already upgraded
    function _areContractsUpgraded() internal view returns (bool) {
        // Check if EmissionsController is deployed and upgraded
        if (address(Env.proxy.emissionsController()).code.length == 0) {
            return false;
        }

        address emissionsControllerImpl = Env.getProxyImplementationBySlot(address(Env.proxy.emissionsController()));
        if (emissionsControllerImpl != address(Env.impl.emissionsController())) {
            return false;
        }

        // Check if RewardsCoordinator is upgraded
        address rewardsCoordinatorImpl = Env.getProxyImplementationBySlot(address(Env.proxy.rewardsCoordinator()));
        if (rewardsCoordinatorImpl != address(Env.impl.rewardsCoordinator())) {
            return false;
        }

        return true;
    }
}

