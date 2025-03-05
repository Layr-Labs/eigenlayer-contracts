// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {Test} from "forge-std/Test.sol";
import {SemVerMixin} from "src/contracts/mixins/SemVerMixin.sol";

// Helper contract to test the abstract SemVerMixin
contract SemVerMixinMock is SemVerMixin {
    constructor(string memory version) SemVerMixin(version) {}

    // Expose internal function for testing
    function majorVersion() public view returns (string memory) {
        return _majorVersion();
    }
}

contract SemVerMixinTest is Test {
    SemVerMixinMock public semVer;

    function test_version_returnsCorrectVersion() public {
        semVer = new SemVerMixinMock("v1.2.3");
        assertEq(semVer.version(), "v1.2.3");
    }

    function test_majorVersion_returnsCorrectMajorVersion() public {
        semVer = new SemVerMixinMock("v1.2.3");
        assertEq(semVer.majorVersion(), "v1");
    }
}
