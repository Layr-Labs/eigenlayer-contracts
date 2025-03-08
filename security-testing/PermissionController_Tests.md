# PermissionController Security Tests

## Permission Boundary Tests

### Test Case: Permission Boundaries
- **Description**: Test that permissions are properly enforced
- **Implementation**:
  ```solidity
  function test_PermissionController_PermissionBoundaries() public {
      // Setup roles and permissions
      address unprivilegedUser = address(0x123);
      address privilegedUser = address(0x456);
      bytes32 role = keccak256("RESTRICTED_ROLE");
      
      // Grant permission to privileged user
      vm.startPrank(owner);
      permissionController.grantPermission(privilegedUser, role);
      vm.stopPrank();
      
      // Test access with insufficient permissions
      vm.startPrank(unprivilegedUser);
      vm.expectRevert("PermissionController.onlyPermitted: not permitted");
      permissionController.onlyPermitted(role);
      vm.stopPrank();
      
      // Test access with proper permissions
      vm.startPrank(privilegedUser);
      permissionController.onlyPermitted(role); // Should not revert
      vm.stopPrank();
  }
  ```

### Test Case: Permission Hierarchy
- **Description**: Test that permission hierarchy is properly enforced
- **Implementation**:
  ```solidity
  function test_PermissionController_PermissionHierarchy() public {
      // Setup roles and permissions
      address user = address(0x123);
      bytes32 parentRole = keccak256("PARENT_ROLE");
      bytes32 childRole = keccak256("CHILD_ROLE");
      
      // Setup role hierarchy
      vm.startPrank(owner);
      permissionController.setRoleAdmin(childRole, parentRole);
      
      // Grant parent role to user
      permissionController.grantPermission(user, parentRole);
      vm.stopPrank();
      
      // Verify user can grant child role
      vm.startPrank(user);
      address otherUser = address(0x456);
      permissionController.grantPermission(otherUser, childRole);
      vm.stopPrank();
      
      // Verify other user has child role
      bool hasPermission = permissionController.hasPermission(otherUser, childRole);
      assertTrue(hasPermission);
      
      // Verify user cannot grant unrelated role
      bytes32 unrelatedRole = keccak256("UNRELATED_ROLE");
      vm.startPrank(user);
      vm.expectRevert("PermissionController.onlyPermitted: not permitted");
      permissionController.grantPermission(otherUser, unrelatedRole);
      vm.stopPrank();
  }
  ```

### Test Case: Permission Scope
- **Description**: Test that permissions are properly scoped
- **Implementation**:
  ```solidity
  function test_PermissionController_PermissionScope() public {
      // Setup roles and permissions
      address user = address(0x123);
      bytes32 role1 = keccak256("ROLE_1");
      bytes32 role2 = keccak256("ROLE_2");
      
      // Grant role1 to user
      vm.startPrank(owner);
      permissionController.grantPermission(user, role1);
      vm.stopPrank();
      
      // Verify user has role1
      bool hasRole1 = permissionController.hasPermission(user, role1);
      assertTrue(hasRole1);
      
      // Verify user does not have role2
      bool hasRole2 = permissionController.hasPermission(user, role2);
      assertFalse(hasRole2);
      
      // Test access with role1
      vm.startPrank(user);
      permissionController.onlyPermitted(role1); // Should not revert
      
      // Test access with role2
      vm.expectRevert("PermissionController.onlyPermitted: not permitted");
      permissionController.onlyPermitted(role2);
      vm.stopPrank();
  }
  ```

## Admin Role Transition Tests

### Test Case: Admin Role Transitions
- **Description**: Test that admin role transitions are secure
- **Implementation**:
  ```solidity
  function test_PermissionController_AdminRoleTransitions() public {
      // Setup initial admin
      address initialAdmin = owner;
      address newAdmin = address(0x456);
      
      // Transfer admin role
      vm.startPrank(initialAdmin);
      permissionController.transferOwnership(newAdmin);
      vm.stopPrank();
      
      // Verify new admin has proper permissions
      vm.startPrank(newAdmin);
      address someUser = address(0x789);
      bytes32 someRole = keccak256("SOME_ROLE");
      permissionController.grantPermission(someUser, someRole);
      vm.stopPrank();
      
      // Verify old admin no longer has permissions
      vm.startPrank(initialAdmin);
      vm.expectRevert("PermissionController.onlyOwner: not owner");
      permissionController.grantPermission(someUser, someRole);
      vm.stopPrank();
  }
  ```

### Test Case: Two-Step Ownership Transfer
- **Description**: Test that ownership transfer requires acceptance
- **Implementation**:
  ```solidity
  function test_PermissionController_TwoStepOwnershipTransfer() public {
      // Setup initial admin
      address initialAdmin = owner;
      address newAdmin = address(0x456);
      
      // Initiate ownership transfer
      vm.startPrank(initialAdmin);
      permissionController.transferOwnership(newAdmin);
      vm.stopPrank();
      
      // Verify new admin is not yet active
      address pendingOwner = permissionController.pendingOwner();
      assertEq(pendingOwner, newAdmin);
      address currentOwner = permissionController.owner();
      assertEq(currentOwner, initialAdmin);
      
      // Verify new admin cannot yet use admin functions
      vm.startPrank(newAdmin);
      address someUser = address(0x789);
      bytes32 someRole = keccak256("SOME_ROLE");
      vm.expectRevert("PermissionController.onlyOwner: not owner");
      permissionController.grantPermission(someUser, someRole);
      
      // Accept ownership
      permissionController.acceptOwnership();
      vm.stopPrank();
      
      // Verify new admin is now active
      currentOwner = permissionController.owner();
      assertEq(currentOwner, newAdmin);
      
      // Verify new admin can use admin functions
      vm.startPrank(newAdmin);
      permissionController.grantPermission(someUser, someRole);
      vm.stopPrank();
  }
  ```

### Test Case: Ownership Renouncement
- **Description**: Test that ownership can be renounced
- **Implementation**:
  ```solidity
  function test_PermissionController_OwnershipRenouncement() public {
      // Setup initial admin
      address initialAdmin = owner;
      
      // Renounce ownership
      vm.startPrank(initialAdmin);
      permissionController.renounceOwnership();
      vm.stopPrank();
      
      // Verify ownership is renounced
      address currentOwner = permissionController.owner();
      assertEq(currentOwner, address(0));
      
      // Verify no one can use admin functions
      vm.startPrank(initialAdmin);
      address someUser = address(0x789);
      bytes32 someRole = keccak256("SOME_ROLE");
      vm.expectRevert("PermissionController.onlyOwner: not owner");
      permissionController.grantPermission(someUser, someRole);
      vm.stopPrank();
  }
  ```

## Unauthorized Access Tests

### Test Case: Unauthorized Access Attempts
- **Description**: Test that unauthorized access attempts are properly rejected
- **Implementation**:
  ```solidity
  function test_PermissionController_UnauthorizedAccessAttempts() public {
      // Setup roles and permissions
      address unprivilegedUser = address(0x123);
      bytes32 role = keccak256("RESTRICTED_ROLE");
      
      // Test unauthorized access to grantPermission
      vm.startPrank(unprivilegedUser);
      vm.expectRevert("PermissionController.onlyOwner: not owner");
      permissionController.grantPermission(unprivilegedUser, role);
      vm.stopPrank();
      
      // Test unauthorized access to revokePermission
      vm.startPrank(unprivilegedUser);
      vm.expectRevert("PermissionController.onlyOwner: not owner");
      permissionController.revokePermission(unprivilegedUser, role);
      vm.stopPrank();
      
      // Test unauthorized access to setRoleAdmin
      vm.startPrank(unprivilegedUser);
      bytes32 otherRole = keccak256("OTHER_ROLE");
      vm.expectRevert("PermissionController.onlyOwner: not owner");
      permissionController.setRoleAdmin(role, otherRole);
      vm.stopPrank();
  }
  ```

### Test Case: Privilege Escalation Attempts
- **Description**: Test that privilege escalation attempts are properly rejected
- **Implementation**:
  ```solidity
  function test_PermissionController_PrivilegeEscalationAttempts() public {
      // Setup roles and permissions
      address user = address(0x123);
      bytes32 lowPrivilegeRole = keccak256("LOW_PRIVILEGE_ROLE");
      bytes32 highPrivilegeRole = keccak256("HIGH_PRIVILEGE_ROLE");
      
      // Grant low privilege role to user
      vm.startPrank(owner);
      permissionController.grantPermission(user, lowPrivilegeRole);
      vm.stopPrank();
      
      // Attempt privilege escalation
      vm.startPrank(user);
      vm.expectRevert("PermissionController.onlyPermitted: not permitted");
      permissionController.grantPermission(user, highPrivilegeRole);
      vm.stopPrank();
      
      // Verify user still only has low privilege role
      bool hasLowPrivilege = permissionController.hasPermission(user, lowPrivilegeRole);
      bool hasHighPrivilege = permissionController.hasPermission(user, highPrivilegeRole);
      assertTrue(hasLowPrivilege);
      assertFalse(hasHighPrivilege);
  }
  ```

### Test Case: Role Confusion Attacks
- **Description**: Test that role confusion attacks are properly rejected
- **Implementation**:
  ```solidity
  function test_PermissionController_RoleConfusionAttacks() public {
      // Setup roles and permissions
      address user = address(0x123);
      bytes32 role1 = keccak256("ROLE_1");
      bytes32 role2 = keccak256("ROLE_2");
      
      // Grant role1 to user
      vm.startPrank(owner);
      permissionController.grantPermission(user, role1);
      vm.stopPrank();
      
      // Attempt to use role1 to access role2 protected function
      vm.startPrank(user);
      vm.expectRevert("PermissionController.onlyPermitted: not permitted");
      permissionController.onlyPermitted(role2);
      vm.stopPrank();
      
      // Verify user still only has role1
      bool hasRole1 = permissionController.hasPermission(user, role1);
      bool hasRole2 = permissionController.hasPermission(user, role2);
      assertTrue(hasRole1);
      assertFalse(hasRole2);
  }
  ```

## Permission Delegation Tests

### Test Case: Permission Delegation Security
- **Description**: Test that permission delegation is secure
- **Implementation**:
  ```solidity
  function test_PermissionController_PermissionDelegationSecurity() public {
      // Setup roles and permissions
      address delegatorUser = address(0x123);
      address otherUser = address(0x456);
      address unprivilegedUser = address(0x789);
      bytes32 role = keccak256("DELEGATABLE_ROLE");
      bytes32 adminRole = keccak256("ADMIN_ROLE");
      
      // Setup role hierarchy
      vm.startPrank(owner);
      permissionController.setRoleAdmin(role, adminRole);
      
      // Grant admin role to delegator
      permissionController.grantPermission(delegatorUser, adminRole);
      vm.stopPrank();
      
      // Test that a user cannot delegate permissions they don't have
      vm.startPrank(unprivilegedUser);
      vm.expectRevert("PermissionController.onlyPermitted: not permitted");
      permissionController.grantPermission(otherUser, role);
      vm.stopPrank();
      
      // Test that a user with delegation permission can delegate
      vm.startPrank(delegatorUser);
      permissionController.grantPermission(otherUser, role);
      vm.stopPrank();
      
      // Verify the delegation worked
      bool hasPermission = permissionController.hasPermission(otherUser, role);
      assertTrue(hasPermission);
  }
  ```

### Test Case: Permission Revocation
- **Description**: Test that permission revocation is secure
- **Implementation**:
  ```solidity
  function test_PermissionController_PermissionRevocation() public {
      // Setup roles and permissions
      address delegatorUser = address(0x123);
      address otherUser = address(0x456);
      address unprivilegedUser = address(0x789);
      bytes32 role = keccak256("REVOCABLE_ROLE");
      bytes32 adminRole = keccak256("ADMIN_ROLE");
      
      // Setup role hierarchy
      vm.startPrank(owner);
      permissionController.setRoleAdmin(role, adminRole);
      
      // Grant admin role to delegator
      permissionController.grantPermission(delegatorUser, adminRole);
      
      // Grant role to other user
      permissionController.grantPermission(otherUser, role);
      vm.stopPrank();
      
      // Verify other user has role
      bool hasPermission = permissionController.hasPermission(otherUser, role);
      assertTrue(hasPermission);
      
      // Test that a user cannot revoke permissions without admin role
      vm.startPrank(unprivilegedUser);
      vm.expectRevert("PermissionController.onlyPermitted: not permitted");
      permissionController.revokePermission(otherUser, role);
      vm.stopPrank();
      
      // Test that a user with admin role can revoke
      vm.startPrank(delegatorUser);
      permissionController.revokePermission(otherUser, role);
      vm.stopPrank();
      
      // Verify the revocation worked
      hasPermission = permissionController.hasPermission(otherUser, role);
      assertFalse(hasPermission);
  }
  ```

### Test Case: Self-Revocation
- **Description**: Test that users can revoke their own permissions
- **Implementation**:
  ```solidity
  function test_PermissionController_SelfRevocation() public {
      // Setup roles and permissions
      address user = address(0x123);
      bytes32 role = keccak256("SELF_REVOCABLE_ROLE");
      
      // Grant role to user
      vm.startPrank(owner);
      permissionController.grantPermission(user, role);
      vm.stopPrank();
      
      // Verify user has role
      bool hasPermission = permissionController.hasPermission(user, role);
      assertTrue(hasPermission);
      
      // Test that user can revoke their own permission
      vm.startPrank(user);
      permissionController.renouncePermission(role);
      vm.stopPrank();
      
      // Verify the revocation worked
      hasPermission = permissionController.hasPermission(user, role);
      assertFalse(hasPermission);
  }
  ```
