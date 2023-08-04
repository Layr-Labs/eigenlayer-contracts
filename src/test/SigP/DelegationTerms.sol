pragma solidity ^0.8.9;

import "../../contracts/strategies/StrategyBase.sol";
import "../../contracts/interfaces/IDelegationManager.sol";


contract SigPDelegationTerms {
    uint256 public paid;
    bytes public isDelegationWithdrawn;
    bytes public isDelegationReceived;


    function payForService(IERC20 /*token*/, uint256 /*amount*/) external payable {
        paid = 1;
    }

    function onDelegationWithdrawn(
        address /*delegator*/,
        IStrategy[] memory /*stakerStrats*/,
        uint256[] memory /*stakerShares*/
    ) external returns(bytes memory) {
        isDelegationWithdrawn = bytes("withdrawn");
        bytes memory _isDelegationWithdrawn = isDelegationWithdrawn;
        return _isDelegationWithdrawn;
    }

    // function onDelegationReceived(
    //     address delegator,
    //     uint256[] memory stakerShares
    // ) external;

    function onDelegationReceived(
        address /*delegator*/,
        IStrategy[] memory /*stakerStrats*/,
        uint256[] memory /*stakerShares*/
    ) external returns(bytes memory) {
        // revert("test");
        isDelegationReceived = bytes("received");
        bytes memory _isDelegationReceived = isDelegationReceived;
        return _isDelegationReceived;
    }

    function delegate() external {
        isDelegationReceived = bytes("received");
    }
}
