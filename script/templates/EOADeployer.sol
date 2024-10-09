// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/utils/ConfigParser.sol";

/// @notice Deployment data struct
struct Deployment {
    string name;
    address deployedTo;
}

/// @notice template for an EOA script
abstract contract EOADeployer is ConfigParser {
    Deployment[] internal _deployments;

    function deploy(string memory envPath) public returns (Deployment[] memory) {
        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        return _deploy(addrs, env, params);
    }

    function _deploy(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (Deployment[] memory);
}