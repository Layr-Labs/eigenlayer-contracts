// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;


/**
 * IPermissionController
 *
 * This interface defines a delegated account feature that enables AVSs,
 * Operators, and potentially even Stakers to use role based access control to 
 * manage the delegation of their account for core protocol contract features.
 *
 * AVSs and Operators should be able to:
 *
 * 1) Create a delegated account if they dont already have one. Its critical
 *    that the message sender address becomes the ID of the account to ensure
 *    proxying or other issues don't enable people to shadow accounts for other
 *    actors.
 * 2) Produce multiple delegates of admin keys, to enable root key recovery
 *    in the event a seed phrase for an initial gets misplaced.
 * 3) Enable accounts to use fine-grained roles to lock down specific permissions
 *    such that delegates have the minimum access required to the account
 *    necessary to do their work.
 * 4) Enables any admins to add or revoke roles from anyone.
 * 5) Enables delegates to delegate across multiple AVSs or Operators should that
 *    need arise for composability purposes.
 *
 * An admin creates an account, and delegates permissions to other addresses by
 * signifying the core contract address and the function selector that is enabled
 * with that permission. This only works if the core contract function specified
 * supports delegated accounts by checking the permission controller in its code.
 *
 * Most core contract functions that support account delegation will take an account
 * address as input, so it should be obvious where the support is by looking at the ABIs.
 *
 * Instead of a bytes20, the account ID is simply the address of the msg.sender that
 * created the account. This remains secure, backwards compatible against AVS and Operators
 * whose "accounts" are their addresses, is less opaque, and prevents collisions easily.
 *
 */
interface IPermissionController {
    //////////////////////////////////////////////////
    // Errors 
    //////////////////////////////////////////////////

    /**
     * AccountAlreadyExists
     *
     * This error is thrown by the controller if a message sender
     * attempts to call #createAccount() when an account already
     * exists for that address.
     */
    error AccountAlreadyExists();

    /**
     * CallerNotAdmin
     *
     * This error this thrown by the controller if a message
     * sender calls a management interface without proper admin
     * permission for the declared account.
     */
    error CallerNotAdmin();

    /**
     * InvalidAccount
     *
     * This error is thrown by the controller if a message
     * sender attempts to change the state of a delegated
     * account this controller is unaware of.
     */
    error InvalidAccount();

    //////////////////////////////////////////////////
    // Events
    //////////////////////////////////////////////////
    
    /**
     * DelegatedAccountCreated
     *
     * This event is emitted when a message sender creates a delegated account.
     * 
     * @param owner the message sender, and the address of the first "admin" on this account.
     */
    event DelegatedAccountCreated(address indexed owner);
   
    /**
     * DelegatedAccountAdminChange
     *
     * This event is emitted when an account admin is added or removed by 
     * an existing admin.
     *
     * @param account  the address of the account an admin was changed on 
     * @param operator the message sender that operated as an admin to make the change 
     * @param delegate the address of the delegate whose admin rights were written
     * @param isAdmin  true if the delegate was added as admin, false if it was revoked. 
     */
    event DelegatedAccountAdminChange(
        address indexed account, 
        address operator,
        address delegate,
        bool isAdmin);

    /**
     * AccountDelegateStateChange
     *
     * This event is emitted when an account admin adds or removes a delegate
     * to or from their account for a given core protocol action.
     * 
     * @param account       the address of the account that delegation changed for
     * @param operator      the message sender acting as an admin who is changing permissions
     * @param target        the address of the contract that delegation is changing for
     * @param selector      the function selector on the contract that delegation is changing for
     * @param delegate      the address of the delegate you want to add or revoke perms for
     * @param hasPermission true if the delegation was added, false if it was revoked
     */
    event AccountDelegateStateChange(
        address indexed account,
        address operator,
        address target,
        bytes4 selector,
        address delegate,
        bool hasPermission);

    //////////////////////////////////////////////////
    // Public Management Interface
    //////////////////////////////////////////////////
    
    /**
     * createAccount()
     *
     * A legacy AVSServiceManager (or AVS EOA Wallet), or an Operator, or Operator Contract,
     * can call this to create an account. This contract does not ensure that the message
     * sender is registered as an AVS or Operator, such that this could presumably be done in
     * any order and composable along with onboarding work flows.
     *
     * The account ID will be set to the msg.sender, and no additional delegations or admin
     * keys will be created. This simply creates the stub account, and sets the message sender
     * as the default admin.
     *
     * This function will revert if the an account for that user already exists.
     */
    function createAccount() external;

    /**
     * setAccountAdmin()
     *
     * An admin of an account can add or remove another admin to and from account. These admins
     * are capable of creating other admins, other delegations, and can potentially have
     * privledge escalation onto any delegated call. It is critical that you trust
     * admins, limit the number of admins to only what is necessary, and use role based access
     * control for all functions. In practice, multiple admins should only be used to
     * recover the main set of admin keys. 
     *
     * Setting an admin gives them the ability to revoke admin access, potentially even from
     * the admin that provided the permission - so use with caution and mainly for recovery
     * purposes only.
     *
     * This function will revert if:
     * - The specified account is not a valid account.
     * - The message sender is not a valid admin for the given account.
     *
     * This method will not revert if the operation results with a no-op. Attempting to
     * add an admin who already is one, or remove an admin that wasn't one to begin with,
     * will succeed without warning and complete normally. 
     *
     * WARNING: This method will allow an admin to revoke admin rights from themselves. This
     *          is enabled for obvious reasons but all admin management needs careful consideration.
     *          Calling this method as the only admin, setting yourself as the delegate, and
     *          setting isAdmin to false will lock your delegated permissions forever. In a way,
     *          this could be used by networks to go fully permissionless with no "owner" once
     *          all subordinate functions are properly delegated. It's also possible for an AVS
     *          to set the only admin as an immutable smart contract, potentially enabling DAO
     *          governance of the entire network composition.
     *
     * @param account  the address of the account you want to modify admin perms for 
     * @param delegate the address of the delegate you want to modify admin perms for
     * @param isAdmin  true if you want to make them an admin, false if you want to revoke it. 
     */
    function setAccountAdmin(address account, address delegate, bool isAdmin) external;

    /**
     * setDelegatedRole() 
     *
     * An admin of a valid account can call this method to enable another address to
     * act on its behalf when calling core contract methods. While an admin can pass in
     * any arbitrary target and selector, it is only if that contract and that method
     * support delegation by calling the PermissionController that the delegation will
     * be effective. Attempting to delegate a contract and selector that does not support
     * delegation will likely result in reversion or other errors.
     *
     * This function will revert if:
     * - The specified account is not a valid account
     * - The message sender is not a valid admin of the specified account
     *
     * The function will NOT revert if the delegate is already in the state specified by
     * the caller. For instance, setting a delegated role to true on a given target or selector
     * for a delegate that already has that permission will still SSTORE the value and execute.
     * The reason for this is we do not want to incur the incremental gas costs every time to
     * check. Frontends and tools should reasonably be able to determine existing state and
     * avoid no-ops.
     * 
     * @param account       the account you want to add a delegate to
     * @param target        the address of the contract that the delegate should be able to use
     * @param selector      the function selector on the target contract the delegate should be able to call
     * @param hasPermission true to give the delegate the permission, false to revoke it
     */ 
    function setDelegatedRole(address account, address target, bytes4 selector, address delegate, bool hasPermission) external;
    
    //////////////////////////////////////////////////
    // Public Introspection Interface
    //////////////////////////////////////////////////
    
    /**
     * isValidAccount()
     *
     * @param account the address of the account to inspect
     * @return true if the controller knows about this account, false otherwise.
     */
    function isValidAccount(address account) external view returns (bool);

    /**
     * hasDelegationOrAdmin()
     *
     * Anyone can call this method to check to see if a specific delegate is capable
     * of acting on contract function for a given account. Specifically, this method could
     * be used by core contracts to require() that the proper delegation is present to
     * support role based access control.
     *
     * This method will always return true or false, and never revert. Instead of
     * reverting on missing accounts, it will simply return false. This is because 
     * the answer "no" is valid for the question asked, whereas "what are the admins of
     * this empty account" being an empty array is misleading. 
     * 
     * This method will return true if the delegate provided is an admin of the account but does not have the
     * explicit delegated permission. The decision here is that root privledge escalation is
     * a convienence, and any admin could easily delegate any missing permissions to themselves.
     * While there are use cases for not providing escalation, they are rare and additional ABIs
     * for hasDelegation() (without admin escalation) can easily be added later).
     * 
     * @param account  the address of the account you want to check delegation for
     * @param target   the address of the contract the delegate is trying to use
     * @param selector the function selector on the contract the delegate is trying to use
     * @param delegate the address of the delegate in question
     * @return true if the delegate has permission on that account for the given contract and selector, false otherwise.
     */
    function hasDelegationOrAdmin(address account, address target, bytes4 selector, address delegate) external view returns (bool);

    /**
     * getAccountAdmins()
     *
     * Returns a list of all the active admins for a given account.
     *
     * This method will revert if the account is invalid. It could potentially
     * return an empty array if the account was completely revoked.
     *
     * @param account the account you want the admins for
     * @return an array of addresses that have admin and escalation perms
     */
    function getAccountAdmins(address account) external view returns (address[] memory);

    /**
     * getAccountPermissions 
     *
     * Returns a list of all selector delegations that a delegate may have for
     * a given account and contract. Calling this method on with all core contract
     * addresses passed in as 'target' enables anyone to introspect the full
     * set of protocol permissions for a given delegate and account.
     *
     * This method will revert if the account is not valid
     *
     * @param account    the account you want to introspect permissions for
     * @param delegate   the delegate you wish to inspect on that account
     * @param target     the contract address you want all of the permissions for 
     * @return selectors an array of function selectors for each contract address the perms are for 
     */
    function getAccountPermissions(address account, address delegate, address target) external view returns (bytes4[] memory);
}
