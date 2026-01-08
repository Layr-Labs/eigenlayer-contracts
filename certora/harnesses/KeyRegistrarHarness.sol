// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/permissions/KeyRegistrar.sol";

contract KeyRegistrarHarness is KeyRegistrar {
    using OperatorSetLib for OperatorSet;
    constructor(
        IPermissionController _permissionController,
        IAllocationManager _allocationManager,
        string memory _version
    )
        KeyRegistrar(_permissionController, _allocationManager, _version)
    {}

    /// @notice Returns the operatorSet key derived from avs and id
    function getOperatorSetKey(OperatorSet calldata os) external pure returns (bytes32) {
        return os.key(); // calls OperatorSetLib.key()
    }


    function getOperatorKeyDataHash(bytes32 key, address operator) external view returns (bytes32) {
        return keccak256(_operatorKeyInfo[key][operator].keyData);
    }
}