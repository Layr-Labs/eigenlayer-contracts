// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "src/contracts/permissions/PauserRegistry.sol";
import "forge-std/Test.sol";

abstract contract EigenLayerUnitTestBase is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    PauserRegistry public pauserRegistry;
    ProxyAdmin public eigenLayerProxyAdmin;

    mapping(address => bool) public addressIsExcludedFromFuzzedInputs;

    address public constant pauser = address(555);
    address public constant unpauser = address(556);

    // Helper Functions/Modifiers
    modifier filterFuzzedAddressInputs(address fuzzedAddress) {
        cheats.assume(!addressIsExcludedFromFuzzedInputs[fuzzedAddress]);
        _;
    }

    function setUp() public virtual {
        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);
        eigenLayerProxyAdmin = new ProxyAdmin();

        addressIsExcludedFromFuzzedInputs[address(pauserRegistry)] = true;
        addressIsExcludedFromFuzzedInputs[address(eigenLayerProxyAdmin)] = true;
    }
}
