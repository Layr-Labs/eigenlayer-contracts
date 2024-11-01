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
    function _deploy() internal override returns (Deployment[] memory) {
        vm.startBroadcast();

        // create a defunct EigenPodManager
        address newEigenPodManager = address(
            new EigenPodManager(
                IETHPOSDeposit(address(1)),
                IBeacon(address(2)),
                IStrategyManager(address(3)),
                ISlasher(address(4)),
                IDelegationManager(address(5))
            )
        );

        _deployments.push(singleton(newEigenPodManager));

        address newEigenPod = address(
            new EigenPod(
                IETHPOSDeposit(_ethPos()),
                IEigenPodManager(newEigenPodManager), // update EigenPodManager address
                getUint64("EIGENPOD_GENESIS_TIME") // set genesis time
            )
        );

        // create and record new EigenPod pointing to defunct EigenPodManager
        _deployments.push(singleton(newEigenPod));

        vm.stopBroadcast();

        return _deployments;
    }

    function zeusTest() public override {
        _deploy();

        Deployment memory eigenPodManager = _deployments[0];
        Deployment memory eigenPod = _deployments[1];

        require(eigenPodManager.deployedTo != address(0), "EigenPodManager deployment failed");

        require(eigenPod.deployedTo != address(0), "EigenPod deployment failed");
    }
}
