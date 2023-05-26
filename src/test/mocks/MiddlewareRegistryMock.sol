// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/interfaces/IServiceManager.sol";
import "../../contracts/interfaces/IStrategyManager.sol";
import "../../contracts/interfaces/ISlasher.sol";

import "forge-std/Test.sol";




contract MiddlewareRegistryMock is  DSTest{
    IServiceManager public serviceManager;
    IStrategyManager public strategyManager;
    ISlasher public slasher;


    constructor(
        IServiceManager _serviceManager,
        IStrategyManager _strategyManager
    ){
        serviceManager = _serviceManager;
        strategyManager = _strategyManager;
        slasher = _strategyManager.slasher();
    }

    function registerOperator(address operator, uint32 serveUntil) public {        
        require(slasher.canSlash(operator, address(serviceManager)), "Not opted into slashing");
        serviceManager.recordFirstStakeUpdate(operator, serveUntil);

    }

    function deregisterOperator(address operator) public {
        uint32 latestServeUntilBlock = serviceManager.latestServeUntilBlock();
        serviceManager.recordLastStakeUpdateAndRevokeSlashingAbility(operator, latestServeUntilBlock);
    }

    function propagateStakeUpdate(address operator, uint32 blockNumber, uint256 prevElement) external {
        uint32 latestServeUntilBlock = serviceManager.latestServeUntilBlock();
        serviceManager.recordStakeUpdate(operator, blockNumber, latestServeUntilBlock, prevElement);
    }

    function isActiveOperator(address operator) external pure returns (bool) {
        if (operator != address(0)){
            return true;
        } else {
            return false;
        }
    }

}