// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/permissions/Pausable.sol";

contract EigenPodManagerMock is Test, Pausable {
    receive() external payable {}
    fallback() external payable {}

    mapping(address => int256) public podShares;

    constructor(IPauserRegistry _pauserRegistry) {
        _initializePauser(_pauserRegistry, 0);
    }

    function podOwnerShares(address podOwner) external view returns (int256) {
        return podShares[podOwner];
    }

    function setPodOwnerShares(address podOwner, int256 shares) external {
        podShares[podOwner] = shares;
    }

    function denebForkTimestamp() external pure returns (uint64) {
        return type(uint64).max;
    }
}