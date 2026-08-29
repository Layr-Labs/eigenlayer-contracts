// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {CrosschainDeployLib} from "../CrosschainDeployLib.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";

/// Purpose: precompute and register the DelegationManager implementation address without deploying its bytecode.
contract DeployImplementations is CoreContractsDeployer {
    using Env for *;

    address internal constant EXPECTED_MAINNET_IMPLEMENTATION = 0x6a8BEd4062C895130E2d09bA442D3eCEAd5Df6c2;

    function _runAsEOA() internal virtual override {
        address implementation = CrosschainDeployLib.computeCrosschainAddress(
            Env.protocolCouncilMultisig(), keccak256(_delegationManagerInitCode()), type(DelegationManager).name
        );
        if (keccak256(bytes(Env.env())) == keccak256("mainnet")) {
            require(implementation == EXPECTED_MAINNET_IMPLEMENTATION, "unexpected mainnet implementation");
        }
        deployImpl({name: type(DelegationManager).name, deployedTo: implementation});
    }

    function testScript() public virtual {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        runAsEOA();

        assertEq(
            address(Env.impl.delegationManager()).code.length,
            0,
            "implementation must remain undeployed until execution"
        );
    }

    function _delegationManagerInitCode() internal view returns (bytes memory) {
        return abi.encodePacked(
            type(DelegationManager).creationCode,
            abi.encode(
                Env.proxy.strategyManager(),
                Env.proxy.eigenPodManager(),
                Env.proxy.allocationManager(),
                Env.impl.pauserRegistry(),
                Env.proxy.permissionController(),
                Env.MIN_WITHDRAWAL_DELAY(),
                Env.deployVersion()
            )
        );
    }
}
