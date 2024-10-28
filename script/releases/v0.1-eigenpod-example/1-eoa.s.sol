// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "zeus-templates/templates/EOADeployer.sol";

import "src/contracts/pods/EigenPod.sol";
import "src/contracts/pods/EigenPodManager.sol";
import "src/contracts/interfaces/IEigenPodManager.sol";

import "zeus-templates/utils/StringUtils.sol";

contract DeployEigenPodAndManager is EOADeployer {
    using StringUtils for string;

    Deployment[] private _deployments;

    /// @notice deploys an EigenPod and returns the deployed address
    function _deploy(Addresses memory, Environment memory, Params memory params) internal override returns (Deployment[] memory) {

        vm.startBroadcast();

        // create a defunct EigenPodManager
        address newEigenPodManager = address(new EigenPodManager(
            IETHPOSDeposit(address(1)),
            IBeacon(address(2)),
            IStrategyManager(address(3)),
            ISlasher(address(4)),
            IDelegationManager(address(5))
        ));

        // record new deployment
        _deployments.push(Deployment({
            overrideName: "",
            deployedTo: newEigenPodManager,
            singleton: true
        }));

        // create and record new EigenPod pointing to defunct EigenPodManager
        _deployments.push(Deployment({
            overrideName: "",
            deployedTo: address(new EigenPod(
                IETHPOSDeposit(params.ethPOS),
                IEigenPodManager(newEigenPodManager), // update EigenPodManager address
                params.EIGENPOD_GENESIS_TIME
            )),
            singleton: true
        }));

        vm.stopBroadcast();

        return _deployments;
    }

    function setUp() public {
        (Addresses memory addrs, Environment memory env, Params memory params) = _readConfigFile("script/configs/zipzoop.json");
        _deploy(addrs, env, params);
    }

    function test_DeployEigenPodManager() public {
        Deployment memory eigenPodManager = _deployments[0];
        require(eigenPodManager.deployedTo != address(0), "eigenPodManager not deployed");
        require(eigenPodManager.overrideName.eq(""), "eigenPodManager name mismatch");

        EigenPodManager eigenPodManagerInstance = EigenPodManager(eigenPodManager.deployedTo);

        require(address(eigenPodManagerInstance.ethPOS()) == address(1), "ethPOS incorrectly set");
        require(address(eigenPodManagerInstance.eigenPodBeacon()) == address(2), "eigenPodBeacon incorrectly set");
        require(address(eigenPodManagerInstance.strategyManager()) == address(3), "strategyManager incorrectly set");
        require(address(eigenPodManagerInstance.slasher()) == address(4), "slasher incorrectly set");
        require(address(eigenPodManagerInstance.delegationManager()) == address(5), "delegation incorrectly set");
    }

    function test_DeployEigenPod() public {
        Deployment memory eigenPodManager = _deployments[0];
        Deployment memory eigenPod = _deployments[1];
        require(eigenPod.deployedTo != address(0), "eigenPod not deployed");
        require(eigenPod.overrideName.eq(""), "eigenPod name mismatch");

        EigenPod eigenPodInstance = EigenPod(payable(eigenPod.deployedTo));
        require(address(eigenPodInstance.eigenPodManager()) == eigenPodManager.deployedTo, "eigenPodManager not set");
    }
}
