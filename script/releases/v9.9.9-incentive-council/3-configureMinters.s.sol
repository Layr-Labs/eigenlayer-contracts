// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import {UpgradeContracts} from "./2-upgradeContracts.s.sol";
import {DeployImplementations} from "./1-deployImplementations.s.sol";
import "../Env.sol";
import "../TestUtils.sol";

import {IBackingEigen} from "src/contracts/interfaces/IBackingEigen.sol";

/// Purpose: Configure minter permissions on BackingEigen
/// - Revoke minting rights from old hopper address
/// - Grant minting rights to EmissionsController
contract ConfigureMinters is UpgradeContracts {
    using Env for *;
    using Encode for *;

    // Old hopper address that needs minting rights revoked
    // TODO: Replace with actual hopper address
    address constant OLD_HOPPER_ADDRESS = 0x0000000000000000000000000000000000000000;

    function _runAsMultisig() internal virtual override prank(Env.executorMultisig()) {
        IBackingEigen backingEigen = Env.proxy.beigen();

        // 1. Revoke minting rights from old hopper address
        if (OLD_HOPPER_ADDRESS != address(0)) {
            backingEigen.setIsMinter(OLD_HOPPER_ADDRESS, false);
        }

        // 2. Grant minting rights to EmissionsController
        backingEigen.setIsMinter(address(Env.proxy.emissionsController()), true);
    }

    function testScript() public virtual override {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        // Check if already configured
        if (_areMintersConfigured()) {
            return;
        }

        // 1. Deploy implementations and proxy as EOA (from step 1)
        _mode = OperationalMode.EOA;
        DeployImplementations._runAsEOA();

        // 2. Upgrade contracts if not upgraded (from step 2)
        if (!_areContractsUpgraded()) {
            _mode = OperationalMode.Multisig;
            _unsafeResetHasPranked();
            UpgradeContracts._runAsMultisig();
            _unsafeResetHasPranked();
        }

        // 3. Configure minters via multisig
        execute();

        // 4. Validate minter configuration
        IBackingEigen backingEigen = Env.proxy.beigen();

        // Validate old hopper no longer has minting rights (if address is set)
        if (OLD_HOPPER_ADDRESS != address(0)) {
            assertFalse(backingEigen.isMinter(OLD_HOPPER_ADDRESS), "Old hopper should not have minting rights");
        }

        // Validate EmissionsController has minting rights
        assertTrue(
            backingEigen.isMinter(address(Env.proxy.emissionsController())),
            "EmissionsController should have minting rights"
        );
    }

    /// @dev Check if minters are already configured
    function _areMintersConfigured() internal view returns (bool) {
        IBackingEigen backingEigen = Env.proxy.beigen();

        // Check if EmissionsController is deployed
        if (address(Env.proxy.emissionsController()).code.length == 0) {
            return false;
        }

        // Check if EmissionsController has minting rights
        if (!backingEigen.isMinter(address(Env.proxy.emissionsController()))) {
            return false;
        }

        // Check if old hopper still has minting rights (if address is set)
        if (OLD_HOPPER_ADDRESS != address(0) && backingEigen.isMinter(OLD_HOPPER_ADDRESS)) {
            return false;
        }

        return true;
    }
}

