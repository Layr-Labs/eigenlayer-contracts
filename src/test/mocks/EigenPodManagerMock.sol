// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/permissions/Pausable.sol";

contract EigenPodManagerMock is Test, Pausable {
    receive() external payable {}
    fallback() external payable {}

    mapping(address => int256) public stakerDepositShares;

    constructor(IPauserRegistry _pauserRegistry) {
        _initializePauser(_pauserRegistry, 0);
    }

    function setStakerDepositShares(address podOwner, int256 shares) external {
        stakerDepositShares[podOwner] = shares;
    }

    function denebForkTimestamp() external pure returns (uint64) {
        return type(uint64).max;
    }
}