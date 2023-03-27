// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IPauserRegistry.sol";

/**
 * @title Defines pauser & unpauser roles + modifiers to be used elsewhere.
 * @author Layr Labs, Inc.
 */
contract PauserRegistry is IPauserRegistry {
    /// @notice Unique address that holds the pauser role.
    address public pauser;

    /// @notice Unique address that holds the unpauser role. Capable of changing *both* the pauser and unpauser addresses.
    address public unpauser;

    event PauserChanged(address previousPauser, address newPauser);

    event UnpauserChanged(address previousUnpauser, address newUnpauser);

    modifier onlyUnpauser() {
        require(msg.sender == unpauser, "msg.sender is not permissioned as unpauser");
        _;
    }

    constructor(address _pauser, address _unpauser) {
        _setPauser(_pauser);
        _setUnpauser(_unpauser);
    }

    /// @notice Sets new pauser - only callable by unpauser, as the unpauser is expected to be kept more secure, e.g. being a multisig with a higher threshold
    function setPauser(address newPauser) external onlyUnpauser {
        _setPauser(newPauser);
    }

    /// @notice Sets new unpauser - only callable by unpauser, as the unpauser is expected to be kept more secure, e.g. being a multisig with a higher threshold
    function setUnpauser(address newUnpauser) external onlyUnpauser {
        _setUnpauser(newUnpauser);
    }

    function _setPauser(address newPauser) internal {
        require(newPauser != address(0), "PauserRegistry._setPauser: zero address input");
        emit PauserChanged(pauser, newPauser);
        pauser = newPauser;
    }

    function _setUnpauser(address newUnpauser) internal {
        require(newUnpauser != address(0), "PauserRegistry._setUnpauser: zero address input");
        emit UnpauserChanged(unpauser, newUnpauser);
        unpauser = newUnpauser;
    }
}
