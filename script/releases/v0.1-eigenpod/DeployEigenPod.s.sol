// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";
import "src/contracts/pods/EigenPod.sol";

contract DeployEigenPod is EOABuilder {
    Deployment[] internal deployments;

    /// @notice deploys an EigenPod and returns the deployed address
    function _deploy(Addresses memory addrs, Environment memory, Params memory params) internal override returns (Deployment[] memory) {

        deployments.push(Deployment({
            name: type(EigenPod).name,
            deployedTo: address(new EigenPod(
                IETHPOSDeposit(params.ethPOS),
                IEigenPodManager(addrs.eigenPodManager.proxy),
                params.EIGENPOD_GENESIS_TIME
            ))
        }));

        return deployments;
    }
}