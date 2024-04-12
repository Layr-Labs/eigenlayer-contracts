// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";

import "../../contracts/permissions/PauserRegistry.sol";

contract PauserRegistryUnitTests is Test {

    Vm cheats = Vm(VM_ADDRESS);

    PauserRegistry public pauserRegistry;

    address public pauser = address(555);
    address public unpauser = address(999);

    mapping(address => bool) public addressIsExcludedFromFuzzedInputs;

    event PauserStatusChanged(address pauser, bool canPause);

    event UnpauserChanged(address previousUnpauser, address newUnpauser);

    function setUp() virtual public {
        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);
    }

    function testSetIsPauserTrue(address newPauser) public {
        cheats.assume(newPauser != address(0));

        cheats.startPrank(pauserRegistry.unpauser());
        cheats.expectEmit(true, true, true, true, address(pauserRegistry));
        emit PauserStatusChanged(newPauser, true);
        pauserRegistry.setIsPauser(newPauser, true);
        cheats.stopPrank();

        require(pauserRegistry.isPauser(newPauser), "newPauser not set correctly");
    }

    function testSetIsPauserFalse() public {
        cheats.startPrank(pauserRegistry.unpauser());
        cheats.expectEmit(true, true, true, true, address(pauserRegistry));
        emit PauserStatusChanged(pauser, false);
        pauserRegistry.setIsPauser(pauser, false);
        cheats.stopPrank();

        require(!pauserRegistry.isPauser(pauser), "pauser not set correctly");
    }

    function testSetUnpauser(address newUnpauser) public {
        cheats.assume(newUnpauser != address(0));

        cheats.startPrank(pauserRegistry.unpauser());
        address oldAddress = pauserRegistry.unpauser();
        cheats.expectEmit(true, true, true, true, address(pauserRegistry));
        emit UnpauserChanged(oldAddress, newUnpauser);
        pauserRegistry.setUnpauser(newUnpauser);
        cheats.stopPrank();

        require(pauserRegistry.unpauser() == newUnpauser, "pauser not set correctly");
    }

    function testSetPauser_RevertsWhenCallingFromNotUnpauser(address notUnpauser, address newPauser) public {
        cheats.assume(notUnpauser != pauserRegistry.unpauser());
        cheats.assume(newPauser != address(0));

        cheats.startPrank(notUnpauser);
        cheats.expectRevert(bytes("msg.sender is not permissioned as unpauser"));
        pauserRegistry.setIsPauser(newPauser, true);
        cheats.stopPrank();
    }

    function testSetUnpauser_RevertsWhenCallingFromNotUnpauser(address notUnpauser, address newUnpauser) public {
        cheats.assume(notUnpauser != pauserRegistry.unpauser());
        cheats.assume(newUnpauser != address(0));

        cheats.startPrank(notUnpauser);
        cheats.expectRevert(bytes("msg.sender is not permissioned as unpauser"));
        pauserRegistry.setUnpauser(newUnpauser);
        cheats.stopPrank();
    }

    function testSetPauser_RevertsWhenSettingToZeroAddress() public {
        address newPauser = address(0);

        cheats.startPrank(pauserRegistry.unpauser());
        cheats.expectRevert(bytes("PauserRegistry._setPauser: zero address input"));
        pauserRegistry.setIsPauser(newPauser, true);
        cheats.stopPrank();
    }

    function testSetUnpauser_RevertsWhenSettingToZeroAddress() public {
        address newUnpauser = address(0);

        cheats.startPrank(pauserRegistry.unpauser());
        cheats.expectRevert(bytes("PauserRegistry._setUnpauser: zero address input"));
        pauserRegistry.setUnpauser(newUnpauser);
        cheats.stopPrank();
    }
}
