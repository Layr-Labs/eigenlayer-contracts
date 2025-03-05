// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/permissions/PermissionController.sol";
import "src/contracts/interfaces/IPermissionController.sol";

import "src/test/utils/EigenLayerUnitTestSetup.sol";

contract PermissionControllerUnitTests is EigenLayerUnitTestSetup, IPermissionControllerEvents, IPermissionControllerErrors {
    address account = address(0x1);
    address admin = address(0x2);
    address admin2 = address(0x3);
    address appointee1 = address(0x4);
    address appointee2 = address(0x5);

    address target1;
    address target2;
    bytes4 selector1 = IDelegationManager.updateOperatorMetadataURI.selector;
    bytes4 selector2 = IAllocationManager.modifyAllocations.selector;

    function setUp() public virtual override {
        // Setup - already deploys permissionController
        EigenLayerUnitTestSetup.setUp();

        // Set targets
        target1 = address(delegationManagerMock);
        target2 = address(allocationManagerMock);
    }
}

contract PermissionControllerUnitTests_AddPendingAdmin is PermissionControllerUnitTests {
    function testFuzz_getAdmin_notSet(address account) public view filterFuzzedAddressInputs(account) {
        address[] memory admins = permissionController.getAdmins(account);
        assertEq(admins[0], account, "Account is not admin");
        assertTrue(permissionController.isAdmin(account, account), "Account is not admin");
    }

    /// @notice Tests the setAdmin function when it has not been initialized
    function test_revert_caller_not_account() public {
        cheats.expectRevert(IPermissionControllerErrors.NotAdmin.selector);
        permissionController.addPendingAdmin(account, admin);
    }

    function test_revert_adminPending() public {
        // Set admin to be pending
        cheats.startPrank(account);
        permissionController.addPendingAdmin(account, admin);

        // Expect revert
        cheats.expectRevert(IPermissionControllerErrors.AdminAlreadyPending.selector);
        permissionController.addPendingAdmin(account, admin);
    }

    function test_addPendingAdmin_initialSet() public {
        // Expect emit
        cheats.expectEmit(true, true, true, true);
        emit PendingAdminAdded(account, admin);

        cheats.prank(account);
        permissionController.addPendingAdmin(account, admin);

        // Check storage
        address[] memory pendingAdmins = permissionController.getPendingAdmins(account);
        assertEq(pendingAdmins[0], admin, "Admin not set to pending");
        assertTrue(permissionController.isPendingAdmin(account, admin), "Pending admin not set correctly");
        assertFalse(permissionController.isAdmin(account, admin), "Pending admin not set correctly");
    }
}

contract PermissionControllerUnitTests_RemovePendingAdmin is PermissionControllerUnitTests {
    function setUp() public virtual override {
        // Setup
        PermissionControllerUnitTests.setUp();

        // Set admin to be pending
        cheats.prank(account);
        permissionController.addPendingAdmin(account, admin);
    }

    function test_revert_notAdmin() public {
        cheats.expectRevert(IPermissionControllerErrors.NotAdmin.selector);
        permissionController.removePendingAdmin(account, admin);
    }

    function test_revert_notPending() public {
        // Expect revert
        cheats.prank(account);
        cheats.expectRevert(IPermissionControllerErrors.AdminNotPending.selector);
        permissionController.removePendingAdmin(account, admin2);
    }

    function test_removePendingAdmin() public {
        // Expect emit
        cheats.expectEmit(true, true, true, true);
        emit PendingAdminRemoved(account, admin);

        cheats.prank(account);
        permissionController.removePendingAdmin(account, admin);

        // Check storage
        address[] memory pendingAdmins = permissionController.getPendingAdmins(account);
        assertEq(pendingAdmins.length, 0, "Pending admin not removed");
    }
}

contract PermissionControllerUnitTests_AcceptAdmin is PermissionControllerUnitTests {
    function setUp() public virtual override {
        // Setup
        PermissionControllerUnitTests.setUp();

        // Set admin
        cheats.prank(account);
        permissionController.addPendingAdmin(account, admin);
    }

    function test_revert_adminNotPending() public {
        cheats.prank(admin2);
        cheats.expectRevert(IPermissionControllerErrors.AdminNotPending.selector);
        permissionController.acceptAdmin(account);
    }

    function test_acceptAdmin() public {
        // Expect emit
        cheats.expectEmit(true, true, true, true);
        emit AdminSet(account, admin);

        cheats.prank(admin);
        permissionController.acceptAdmin(account);

        // Check storage
        address[] memory admins = permissionController.getAdmins(account);
        address[] memory pendingAdmins = permissionController.getPendingAdmins(account);
        assertEq(pendingAdmins.length, 0, "Pending admin not removed");
        assertEq(admins.length, 1, "Additional admin not added");
        assertFalse(permissionController.isAdmin(account, account), "Account should no longer be admin");
        assertTrue(permissionController.isAdmin(account, admin), "Admin not set");
        assertTrue(permissionController.canCall(account, admin, address(0), bytes4(0)), "Admin cannot call");
    }

    function test_addMultipleAdmin_fromAccount() public {
        // Add another pending admin
        cheats.prank(account);
        permissionController.addPendingAdmin(account, admin2);

        // Assert pending admins length
        address[] memory pendingAdmins = permissionController.getPendingAdmins(account);
        assertEq(pendingAdmins.length, 2, "Additional pending admin not added");

        // Accept admins
        cheats.prank(admin);
        permissionController.acceptAdmin(account);
        cheats.prank(admin2);
        permissionController.acceptAdmin(account);

        // Check storage
        address[] memory admins = permissionController.getAdmins(account);
        assertEq(admins.length, 2, "Additional admin not added");
        assertTrue(permissionController.isAdmin(account, admin), "Old admin should still be admin");
        assertTrue(permissionController.isAdmin(account, admin2), "New admin not set correctly");
        assertTrue(permissionController.canCall(account, admin, address(0), bytes4(0)), "Admin cannot call");
        assertTrue(permissionController.canCall(account, admin2, address(0), bytes4(0)), "Admin cannot call");
    }

    function test_addMultipleAdmin_newAdminAdds() public {
        // Accept admin
        cheats.startPrank(admin);
        permissionController.acceptAdmin(account);

        // Add another pending admin
        permissionController.addPendingAdmin(account, admin2);
        cheats.stopPrank();

        // Accept admins
        cheats.prank(admin2);
        permissionController.acceptAdmin(account);

        // Check storage
        address[] memory admins = permissionController.getAdmins(account);
        assertEq(admins.length, 2, "Additional admin not added");
        assertTrue(permissionController.isAdmin(account, admin), "Old admin should still be admin");
        assertTrue(permissionController.isAdmin(account, admin2), "New admin not set correctly");
        assertTrue(permissionController.canCall(account, admin, address(0), bytes4(0)), "Admin cannot call");
        assertTrue(permissionController.canCall(account, admin2, address(0), bytes4(0)), "Admin cannot call");
    }
}

contract PermissionControllerUnitTests_RemoveAdmin is PermissionControllerUnitTests {
    function setUp() public virtual override {
        // Setup
        PermissionControllerUnitTests.setUp();

        // Set admins
        cheats.startPrank(account);
        permissionController.addPendingAdmin(account, admin);
        permissionController.addPendingAdmin(account, admin2);
        cheats.stopPrank();

        // Accept admins
        cheats.prank(admin);
        permissionController.acceptAdmin(account);
        cheats.prank(admin2);
        permissionController.acceptAdmin(account);
    }

    function test_revert_notAdmin() public {
        cheats.prank(appointee1);
        cheats.expectRevert(IPermissionControllerErrors.NotAdmin.selector);
        permissionController.removeAdmin(account, admin);
    }

    function test_revert_adminNotSet() public {
        cheats.prank(admin);
        cheats.expectRevert(IPermissionControllerErrors.AdminNotSet.selector);
        permissionController.removeAdmin(account, appointee1);
    }

    function test_revert_cannotHaveZeroAdmins() public {
        // Remove admin2
        cheats.startPrank(admin);
        permissionController.removeAdmin(account, admin2);

        // Remove admin
        cheats.expectRevert(IPermissionControllerErrors.CannotHaveZeroAdmins.selector);
        permissionController.removeAdmin(account, admin);
    }

    function test_removeAdmin() public {
        // Expect emit
        cheats.expectEmit(true, true, true, true);
        emit AdminRemoved(account, admin);

        cheats.prank(admin);
        permissionController.removeAdmin(account, admin);

        // Check storage
        assertFalse(permissionController.isAdmin(account, admin), "Admin not removed");
        assertFalse(permissionController.canCall(account, admin, address(0), bytes4(0)), "Admin can still call");

        // Assert that admin2 is now the only admin
        address[] memory admins = permissionController.getAdmins(account);
        assertEq(admins[0], admin2, "Account is not admin");
        assertTrue(permissionController.isAdmin(account, admin2), "Admin2 is not admin");
        assertFalse(permissionController.isAdmin(account, admin), "Account should not be admin");
    }
}

contract PermissionControllerUnitTests_SetAppointee is PermissionControllerUnitTests {
    function setUp() public virtual override {
        // Setup
        PermissionControllerUnitTests.setUp();

        // Set & accept admin
        cheats.prank(account);
        permissionController.addPendingAdmin(account, admin);
        cheats.prank(admin);
        permissionController.acceptAdmin(account);
    }

    function test_revert_notAdmin() public {
        cheats.expectRevert(IPermissionControllerErrors.NotAdmin.selector);
        permissionController.setAppointee(account, appointee1, address(0), bytes4(0));
    }

    function test_setAppointee() public {
        // Expect emit
        cheats.expectEmit(true, true, true, true);
        emit AppointeeSet(account, appointee1, target1, selector1);

        cheats.prank(admin);
        permissionController.setAppointee(account, appointee1, target1, selector1);

        // Validate Permissions
        _validateSetAppointee(account, appointee1, target1, selector1);
    }

    function test_revert_appointeeAlreadySet() public {
        // Set appointee
        cheats.startPrank(admin);
        permissionController.setAppointee(account, appointee1, target1, selector1);

        cheats.expectRevert(IPermissionControllerErrors.AppointeeAlreadySet.selector);
        permissionController.setAppointee(account, appointee1, target1, selector1);
        cheats.stopPrank();
    }

    function test_setMultipleAppointees() public {
        // Set appointees
        cheats.startPrank(admin);
        permissionController.setAppointee(account, appointee1, target1, selector1);
        permissionController.setAppointee(account, appointee1, target2, selector2);
        permissionController.setAppointee(account, appointee2, target1, selector1);
        cheats.stopPrank();

        // Validate Permissions
        _validateSetAppointee(account, appointee1, target1, selector1);
        _validateSetAppointee(account, appointee1, target2, selector2);
        _validateSetAppointee(account, appointee2, target1, selector1);
    }

    function _validateSetAppointee(address accountToCheck, address appointee, address target, bytes4 selector) internal view {
        assertTrue(permissionController.canCall(accountToCheck, appointee, target, selector));
        _validateAppointeePermissions(accountToCheck, appointee, target, selector);
        _validateGetAppointees(accountToCheck, appointee, target, selector);
    }

    function _validateAppointeePermissions(address accountToCheck, address appointee, address target, bytes4 selector) internal view {
        (address[] memory targets, bytes4[] memory selectors) = permissionController.getAppointeePermissions(accountToCheck, appointee);
        bool foundTargetSelector = false;
        for (uint i = 0; i < targets.length; ++i) {
            if (targets[i] == target) {
                assertEq(selectors[i], selector, "Selector does not match target");
                foundTargetSelector = true;
                break;
            }
        }
        assertTrue(foundTargetSelector, "Appointee does not have permission for given target and selector");
    }

    function _validateGetAppointees(address accountToCheck, address appointee, address target, bytes4 selector) internal view {
        (address[] memory appointees) = permissionController.getAppointees(accountToCheck, target, selector);
        bool foundAppointee = false;
        for (uint i = 0; i < appointees.length; ++i) {
            if (appointees[i] == appointee) {
                foundAppointee = true;
                break;
            }
        }
        assertTrue(foundAppointee, "Appointee not in list of appointees for given target and selector");
    }
}

contract PermissionControllerUnitTests_RemoveAppointee is PermissionControllerUnitTests {
    function setUp() public virtual override {
        // Setup
        PermissionControllerUnitTests.setUp();

        // Set & accept admin
        cheats.prank(account);
        permissionController.addPendingAdmin(account, admin);
        cheats.prank(admin);
        permissionController.acceptAdmin(account);

        // Set appointees
        cheats.startPrank(admin);
        permissionController.setAppointee(account, appointee1, target1, selector1);
        permissionController.setAppointee(account, appointee1, target2, selector2);
        permissionController.setAppointee(account, appointee2, target1, selector1);
        cheats.stopPrank();
    }

    function test_revert_notAdmin() public {
        cheats.expectRevert(IPermissionControllerErrors.NotAdmin.selector);
        permissionController.removeAppointee(account, appointee1, target1, selector1);
    }

    function test_removeAppointee() public {
        // Expect emit
        cheats.expectEmit(true, true, true, true);
        emit AppointeeRemoved(account, appointee1, target1, selector1);

        cheats.prank(admin);
        permissionController.removeAppointee(account, appointee1, target1, selector1);

        // Validate Permissions
        _validateRemoveAppointee(account, appointee1, target1, selector1);
    }

    function test_revert_appointeeNotSet() public {
        cheats.expectRevert(IPermissionControllerErrors.AppointeeNotSet.selector);
        cheats.prank(admin);
        permissionController.removeAppointee(account, appointee2, target2, selector2);
    }

    function test_removeMultipleAppointees() public {
        // Remove appointees
        cheats.startPrank(admin);
        permissionController.removeAppointee(account, appointee1, target1, selector1);
        permissionController.removeAppointee(account, appointee1, target2, selector2);
        permissionController.removeAppointee(account, appointee2, target1, selector1);
        cheats.stopPrank();

        // Validate Permissions
        _validateRemoveAppointee(account, appointee1, target1, selector1);
        _validateRemoveAppointee(account, appointee1, target2, selector2);
        _validateRemoveAppointee(account, appointee2, target1, selector1);
    }

    // Tests that the encoding from adding an appointee is properly decoded
    function test_symmetricalTargetSelector() public view {
        // Test Decoding
        (address[] memory targets, bytes4[] memory selectors) = permissionController.getAppointeePermissions(account, appointee2);
        assertEq(targets.length, 1, "Incorrect number of targets");
        assertEq(selectors.length, 1, "Incorrect number of selectors");
        assertEq(targets[0], target1, "Incorrect target");
        assertEq(selectors[0], selector1, "Incorrect selector");
    }

    function _validateRemoveAppointee(address accountToCheck, address appointee, address target, bytes4 selector) internal view {
        assertFalse(permissionController.canCall(accountToCheck, appointee, target, selector));
        _validateAppointeePermissionsRemoved(accountToCheck, appointee, target, selector);
        _validateGetAppointeesRemoved(accountToCheck, appointee, target, selector);
    }

    function _validateAppointeePermissionsRemoved(address accountToCheck, address appointee, address target, bytes4 selector)
        internal
        view
    {
        (address[] memory targets, bytes4[] memory selectors) = permissionController.getAppointeePermissions(accountToCheck, appointee);
        bool foundTargetSelector = false;
        for (uint i = 0; i < targets.length; ++i) {
            if (targets[i] == target && selectors[i] == selector) {
                foundTargetSelector = true;
                break;
            }
        }
        assertFalse(foundTargetSelector, "Appointee still has permission for given target and selector");
    }

    function _validateGetAppointeesRemoved(address accountToCheck, address appointee, address target, bytes4 selector) internal view {
        (address[] memory appointees) = permissionController.getAppointees(accountToCheck, target, selector);
        bool foundAppointee = false;
        for (uint i = 0; i < appointees.length; ++i) {
            if (appointees[i] == appointee) {
                foundAppointee = true;
                break;
            }
        }
        assertFalse(foundAppointee, "Appointee still in list of appointees for given target and selector");
    }
}
