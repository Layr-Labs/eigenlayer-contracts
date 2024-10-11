// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;
import "../interfaces/IPermissionController.sol";

// We use Enumerable set to maintain an index. While we do
// spend a little more gas, the amount of money spent on gas
// maintaining this index is far less than the operations required
// to maintain off-chain indexes and serve this information over RPC
// to users. Storing a little bit more on-chain passes the "walk away test"
// which, for protocol account permissions, seems critical.
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

// we created a quick EnumerableSet for Bytes4 since OZ didn't have one
import "../libraries/EnumerableBytes4Set.sol";

abstract contract PermissionControllerStorage is IPermissionController {
    /**
     * DelegatedAccountInfo
     *
     * This stucture stores all of the information contained specifically
     * for a single delegated account. The account ID is the address used
     * in the key for the mapping this structure is in, and will not be
     * duplicated as part of the structure.
     */   
    struct DelegatedAccountInfo {
        // We store a sanity flag here to ensure that we can identify
        // an invalid account from one with no admins (fully revoked)
        bool isValid;   // EVM defaults this to false

        // There can be multiple admins for a given account, so
        // we will add and remove from this set as a way to determine
        // if they are an admin or not
        EnumerableSet.AddressSet admins;

        // delegate => contract => selector => state
        mapping(address => mapping(address => mapping(bytes4 => bool))) delegations;

        // we will also be maintaining a "walk away" index so introspecting
        // on existing permissions does not require off-chain indexing.
        // delegate => contract => selectors
        mapping(address => mapping(address => EnumerableBytes4Set.Bytes4Set)) delegatedContractSelectors; 
    }

    // The mapping of account address to all of the delegated account info for it.
    mapping(address => DelegatedAccountInfo) internal accounts;
}
