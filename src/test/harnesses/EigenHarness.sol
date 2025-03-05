// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "../../contracts/token/Eigen.sol";

contract EigenHarness is Eigen {
    constructor(IERC20 _bEIGEN) Eigen(_bEIGEN) {}

    /// expose internal mint function
    function mint(address to, uint amount) public {
        _mint(to, amount);
    }

    /**
     * @notice This function allows the owner to set the allowedFrom status of an address
     * @param from the address whose allowedFrom status is being set
     * @param isAllowedFrom the new allowedFrom status
     * @dev this function is callable by anoyone in the harness
     */
    function setAllowedFromPermissionless(address from, bool isAllowedFrom) external {
        allowedFrom[from] = isAllowedFrom;
        emit SetAllowedFrom(from, isAllowedFrom);
    }

    function setTransferRestrictionsDisabledAfterToMax() external {
        transferRestrictionsDisabledAfter = type(uint).max;
    }

    function transferOwnershipPermissionless(address newOwner) external {
        _transferOwnership(newOwner);
    }
}
