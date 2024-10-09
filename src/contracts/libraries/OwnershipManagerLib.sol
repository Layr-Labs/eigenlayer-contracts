// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "../interfaces/IOwnershipManager.sol";


library OwnershipManagerLib {
    function getIdentifier(address agent, address ownershipManager) internal view returns (bytes20) {
        return IOwnershipManager(ownershipManager).getIdentifier(agent, address(this), msg.sig);
    }
}

