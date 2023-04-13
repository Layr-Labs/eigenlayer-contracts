// SPDX-Identifier: UNLICENSED
pragma solidity =0.8.12;

import "src/contracts/interfaces/IDelegationTerms.sol";


contract DelegationTerms is IDelegationTerms {
    function payForService(IERC20 token, uint256 amount) external payable override {
        // do nothing
    }

    function onDelegationWithdrawn(
        address delegator,
        IStrategy[] memory stakerStrategyList,
        uint256[] memory stakerShares
    ) external override returns(bytes memory) {
        // do nothing
    }

    function onDelegationReceived(
        address delegator,
        IStrategy[] memory stakerStrategyList,
        uint256[] memory stakerShares
    ) external override returns(bytes memory) {
        // do nothing
    }
}