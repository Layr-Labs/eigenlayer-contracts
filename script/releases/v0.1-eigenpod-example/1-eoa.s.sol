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
}
