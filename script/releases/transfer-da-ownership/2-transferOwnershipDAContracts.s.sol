// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {TransferDAOwnershipConstants as C} from "./TransferDAOwnershipConstants.sol";
import "zeus-templates/utils/Encode.sol";

import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * Purpose:
 *      * enqueue a multisig transaction which;
 *             - transfers DA contract ownership to the EigenDA ops multisig
 *  This should be run via the core ops multisig.
 */
contract TransferOwnershipDAContracts is MultisigBuilder {
    using Env for *;

    function _runAsMultisig() internal override prank(Env.opsMultisig()) {
        Ownable(C.EIGEN_DA_EJECTION_MANAGER).transferOwnership(C.EIGEN_DA_OPS_MULTISIG);
        Ownable(C.EIGEN_DA_REGISTRY_COORDINATOR).transferOwnership(C.EIGEN_DA_OPS_MULTISIG);
        Ownable(C.EIGEN_DA_SERVICE_MANAGER).transferOwnership(C.EIGEN_DA_OPS_MULTISIG);
    }

    function testScript() public virtual {
        _runAsMultisig();

        assertTrue(
            Ownable(C.EIGEN_DA_EJECTION_MANAGER).owner() == C.EIGEN_DA_OPS_MULTISIG,
            "eigenDAEjectionManager owner invalid"
        );
        assertTrue(
            Ownable(C.EIGEN_DA_REGISTRY_COORDINATOR).owner() == C.EIGEN_DA_OPS_MULTISIG,
            "eigenDARegistryCoordinator owner invalid"
        );
        assertTrue(
            Ownable(C.EIGEN_DA_SERVICE_MANAGER).owner() == C.EIGEN_DA_OPS_MULTISIG,
            "eigenDAServiceManager owner invalid"
        );
    }
}
