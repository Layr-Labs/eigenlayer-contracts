// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IDelegationTerms.sol";

contract DelegationTermsMock is IDelegationTerms, Test {

    bool public shouldRevert;
    bool public shouldReturnData;

    function setShouldRevert(bool _shouldRevert) external {
        shouldRevert = _shouldRevert;
    }

    function setShouldReturnData(bool _shouldReturnData) external {
        shouldReturnData = _shouldReturnData;
    }

    function payForService(IERC20 /*token*/, uint256 /*amount*/) external payable {

    }

    function onDelegationWithdrawn(
        address /*delegator*/,
        IStrategy[] memory /*stakerStrategyList*/,
        uint256[] memory /*stakerShares*/
    ) external view returns (bytes memory) {
        if (shouldRevert) {
            revert("reverting as intended");
        }

        if (shouldReturnData) {
            bytes32[5] memory returnData = [bytes32(0), bytes32(0), bytes32(0), bytes32(0), bytes32(0)];
            return abi.encodePacked(returnData);
        }

        bytes memory emptyReturnData;
        return emptyReturnData;
    }

    function onDelegationReceived(
        address /*delegator*/,
        IStrategy[] memory /*stakerStrategyList*/,
        uint256[] memory /*stakerShares*/
    ) external view returns (bytes memory) {
        if (shouldRevert) {
            revert("reverting as intended");
        }
        if (shouldReturnData) {
            bytes32[5] memory returnData = [bytes32(0), bytes32(0), bytes32(0), bytes32(0), bytes32(0)];
            return abi.encodePacked(returnData);
        }

        bytes memory emptyReturnData;
        return emptyReturnData;
    }

}