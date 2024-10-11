// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "./PermissionControllerStorage.sol";

contract PermissionController is PermissionControllerStorage {
    // enable library functions on enumerable set
    using EnumerableSet for EnumerableSet.AddressSet;
    using EnumerableBytes4Set for EnumerableBytes4Set.Bytes4Set;

    /**
     * onlyValidAccounts
     *
     * This modifier is used to ensure that account access only
     * occurs on valid accounts.
     */
    modifier onlyValidAccounts(address account) {
        require(accounts[account].isValid, InvalidAccount());
        _;
    }

    /**
     * onlyAccountAdmin
     *
     * This modifier is used to ensure that functions can only be accessed
     * by message senders that are admins for the given account.
     *
     */
    modifier onlyAccountAdmin(address account) {
        require(accounts[account].admins.contains(msg.sender), CallerNotAdmin());
        _;
    }

    //////////////////////////////////////////////////
    // Public Management Interface
    //////////////////////////////////////////////////
    
    /// @inheritdoc IPermissionController
    function createAccount() external {
        // we assume the message sender is an AVS, Operator, or Staker, but we don't
        // necessarily verify this. There is no harm in any address creating its own
        // delegated account. It's not meaningfully possible to "grief" account creation.
        DelegatedAccountInfo storage account = accounts[msg.sender];

        // check to make sure the account doesn't already exist
        require(!account.isValid, AccountAlreadyExists());

        // make the account valid by marking it as such, and by default add the
        // message sender as an admin. The account address is not a permanent
        // admin, so its inclusion in the set is critical.
        account.isValid = true;
        account.admins.add(msg.sender);

        emit DelegatedAccountCreated(msg.sender);
    }

    /// @inheritdoc IPermissionController
    function setAccountAdmin(address account, address delegate, bool isAdmin) onlyValidAccounts(account) onlyAccountAdmin(account) external {
        // at this point we know the caller is admin on a valid account, so
        // we can simply set the state and emit an event
        if (isAdmin) {
            accounts[account].admins.add(delegate);
        } else {
            accounts[account].admins.remove(delegate);
        }

        emit DelegatedAccountAdminChange(account, msg.sender, delegate, isAdmin);
    }

    /// @inheritdoc IPermissionController
    function setDelegatedRole(
        address account,
        address target,
        bytes4 selector,
        address delegate,
        bool hasPermission) onlyValidAccounts(account) onlyAccountAdmin(account) external {
        
        // at this point we know the caller is admin on a valid account,
        // so we can simply set the state, update the index and emit an event
        DelegatedAccountInfo storage info = accounts[account];
        info.delegations[delegate][target][selector] = hasPermission;
        if (hasPermission) {
            info.delegatedContractSelectors[delegate][target].add(selector);
        } else {
            info.delegatedContractSelectors[delegate][target].remove(selector);
        }

        emit AccountDelegateStateChange(account, msg.sender, target, selector, delegate, hasPermission); 
    }
    
    //////////////////////////////////////////////////
    // Public Introspection Interface
    //////////////////////////////////////////////////

    /// @inheritdoc IPermissionController
    function isValidAccount(address account) external view returns (bool) {
        return accounts[account].isValid;
    }

    /// @inheritdoc IPermissionController
    function hasDelegationOrAdmin(address account, address target, bytes4 selector, address delegate) external view returns (bool) {
        return accounts[account].delegations[delegate][target][selector] || accounts[account].admins.contains(delegate);
    }

    /// @inheritdoc IPermissionController
    function getAccountAdmins(address account) onlyValidAccounts(account) external view returns (address[] memory) {
        return accounts[account].admins.values(); 
    }

    /// @inheritdoc IPermissionController
    function getAccountPermissions(
        address account,
        address delegate,
        address target) onlyValidAccounts(account) external view returns (bytes4[] memory) {
        return accounts[account].delegatedContractSelectors[delegate][target].values();
    }
}
