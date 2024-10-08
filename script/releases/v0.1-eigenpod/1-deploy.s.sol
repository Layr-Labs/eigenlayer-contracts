// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";
import "src/contracts/pods/EigenPod.sol";

//  EOA deployer -   deploy -> Deployment[]  (eoa deployer)
//  Multisig Deployer - execute -> Transaction[]


contract Deploy is EOABuilder {
    Deployment[] internal _deployments;

    /// @notice deploys an EigenPod and returns the deployed address
    function _deploy(Addresses memory addrs, Environment memory, Params memory params) internal override returns (Deployment[] memory) {
        vm.startBroadcast();
        _deployments.push(Deployment({
            name: type(EigenPod).name,
            deployedTo: address(new EigenPod(
                IETHPOSDeposit(params.ethPOS),
                IEigenPodManager(addrs.eigenPodManager.proxy),
                params.EIGENPOD_GENESIS_TIME
            ))
        }));
        vm.stopBroadcast();

        return _deployments;
    }
}