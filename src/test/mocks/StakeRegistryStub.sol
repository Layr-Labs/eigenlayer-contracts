// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/interfaces/IStakeRegistryStub.sol";

contract StakeRegistryStub is IStakeRegistryStub {
    function updateStakes(address[] memory) external {}
}
