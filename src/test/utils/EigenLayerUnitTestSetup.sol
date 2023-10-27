// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "forge-std/Test.sol";

abstract contract EigenLayerUnitTestSetup is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    mapping(address => bool) public addressIsExcludedFromFuzzedInputs;

    address public constant pauser = address(555);
    address public constant unpauser = address(556);

    // Helper Functions/Modifiers
    modifier filterFuzzedAddressInputs(address fuzzedAddress) {
        cheats.assume(!addressIsExcludedFromFuzzedInputs[fuzzedAddress]);
        _;
    }
}