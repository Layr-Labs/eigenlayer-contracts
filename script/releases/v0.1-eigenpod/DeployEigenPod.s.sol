// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";
import "src/contracts/pods/EigenPod.sol";
import "src/contracts/pods/EigenPodManager.sol";
import "src/contracts/interfaces/IEigenPodManager.sol";

contract DeployEigenPodAndManager is EOADeployer {
    Deployment[] internal _deployments;

    /// @notice deploys an EigenPod and returns the deployed address
    function _deploy(Addresses memory, Environment memory, Params memory params) internal override returns (Deployment[] memory) {

        vm.startBroadcast();

        // create a defunct EigenPodManager
        address newEigenPodManager = address(new EigenPodManager(
            IETHPOSDeposit(address(0)),
            IBeacon(address(0)),
            IStrategyManager(address(0)),
            ISlasher(address(0)),
            IDelegationManager(address(0))
        ));

        // record new deployment
        _deployments.push(Deployment({
            name: type(EigenPodManager).name,
            deployedTo: newEigenPodManager
        }));

        // create and record new EigenPod pointing to defunct EigenPodManager
        _deployments.push(Deployment({
            name: type(EigenPod).name,
            deployedTo: address(new EigenPod(
                IETHPOSDeposit(params.ethPOS),
                IEigenPodManager(newEigenPodManager), // update EigenPodManager address
                params.EIGENPOD_GENESIS_TIME
            ))
        }));

        vm.stopBroadcast();

        return _deployments;
    }
}
