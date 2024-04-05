// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../src/contracts/interfaces/IStrategyManager.sol";
import "../../src/contracts/interfaces/IStrategy.sol";
import "../../src/contracts/interfaces/IDelegationManager.sol";
import "../../src/contracts/interfaces/ISignatureUtils.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "forge-std/Test.sol";

contract Staker is Ownable {
    
    constructor(
        IStrategy strategy, 
        IStrategyManager strategyManager,
        IDelegationManager delegation,
        IERC20 token, 
        uint256 amount, 
        address operator
    ) Ownable() {
        token.approve(address(strategyManager), type(uint256).max);
        strategyManager.depositIntoStrategy(strategy, token, amount);
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry;
        delegation.delegateTo(operator, signatureWithExpiry, bytes32(0));
    }
    
    function callAddress(address implementation, bytes memory data) external onlyOwner returns(bytes memory) {
        uint256 length = data.length;
        bytes memory returndata;  
        assembly{
            let result := call(
                gas(),
                implementation,
                callvalue(),
                add(data, 32),
                length,
                0,
                0
            )
            mstore(returndata, returndatasize())
            returndatacopy(add(returndata, 32), 0, returndatasize())
        }


        return returndata;

    }

}
