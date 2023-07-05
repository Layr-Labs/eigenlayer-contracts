// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../utils/MockAVSDeployer.sol";

contract BLSOperatorStateRetrieverUnitTests is MockAVSDeployer {
    using BN254 for BN254.G1Point;

    function setUp() virtual public {
        _deployMockEigenLayerAndAVS();
    }

    function test_getOperatorState() public {
        emit log("hello");
    }
}