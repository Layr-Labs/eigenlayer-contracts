// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IStrategy.sol";
import "../../contracts/permissions/Pausable.sol";

contract EigenPodManagerMock is Test, Pausable {
    receive() external payable {}
    fallback() external payable {}

    mapping(address => int256) public podOwnerDepositShares;

    mapping(address => uint256) public podOwnerSharesWithdrawn;

    constructor(IPauserRegistry _pauserRegistry) Pausable(_pauserRegistry) {
        _initializePauser(0);
    }

    function podOwnerShares(address podOwner) external view returns (int256) {
        return podOwnerDepositShares[podOwner];
    }

    function stakerDepositShares(address user, address) public view returns (uint256 depositShares) {
        return podOwnerDepositShares[user] < 0 ? 0 : uint256(podOwnerDepositShares[user]);
    } 

    function setPodOwnerShares(address podOwner, int256 shares) external {
        podOwnerDepositShares[podOwner] = shares;
    }

    function removeDepositShares(address podOwner, IStrategy strategy, uint256 shares) external {
        podOwnerDepositShares[podOwner] -= int256(shares);
    }

    function denebForkTimestamp() external pure returns (uint64) {
        return type(uint64).max;
    }

    function withdrawSharesAsTokens(address podOwner, address /** strategy */, address /** token */, uint256 shares) external {
        podOwnerSharesWithdrawn[podOwner] += shares;
    }
}