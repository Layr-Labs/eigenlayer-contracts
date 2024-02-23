// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IBeaconChainOracle.sol";


contract BeaconChainOracle is IBeaconChainOracle 
{

    // The address of the native consensus layer block root oracle
    address constant systemContract = 0x000F3df6D732807Ef1319fB7B8bB8522d0Beac02;

    // emitted when the precompile is called
    event OracleBlockRootAdded(uint256 timestamp, bytes32 blockRoot);

    function timestampToBlockRoot(uint256 timestamp) external returns (bytes32){
        (bool success, bytes memory data) = address(systemContract).call(abi.encodePacked(timestamp));
        require(success, "BeaconChainOracle.timestampToBlockRoot: Precompile call failed");
        bytes32 blockRoot = abi.decode(data, (bytes32));
        emit OracleBlockRootAdded(timestamp, blockRoot);
        return blockRoot;
    }
}



