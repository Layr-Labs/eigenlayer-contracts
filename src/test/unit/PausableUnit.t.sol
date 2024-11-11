// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "../../contracts/permissions/PauserRegistry.sol";
import "../harnesses/PausableHarness.sol";

contract PausableUnitTests is Test {

    Vm cheats = Vm(VM_ADDRESS);

    PauserRegistry public pauserRegistry;
    PausableHarness public pausable;

    address public pauser = address(555);
    address public unpauser = address(999);
    uint256 public initPausedStatus = 0;

    mapping(address => bool) public isExcludedFuzzAddress;

    /// @notice Emitted when the pause is triggered by `account`, and changed to `newPausedStatus`.
    event Paused(address indexed account, uint256 newPausedStatus);

    /// @notice Emitted when the pause is lifted by `account`, and changed to `newPausedStatus`.
    event Unpaused(address indexed account, uint256 newPausedStatus);

    function setUp() virtual public {
        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);
        pausable = new PausableHarness(pauserRegistry);
        pausable.initializePauser(initPausedStatus);
    }
    
    function testPause(uint256 previousPausedStatus, uint256 newPausedStatus) public {
        // filter out any fuzzed inputs which would (improperly) flip any bits to '0'.
        cheats.assume(previousPausedStatus & newPausedStatus == previousPausedStatus);

        cheats.startPrank(pauser);
        cheats.expectEmit(true, true, true, true, address(pausable));
        emit Paused(pauser, previousPausedStatus);
        pausable.pause(previousPausedStatus);
        cheats.stopPrank();

        require(pausable.paused() == previousPausedStatus, "previousPausedStatus not set correctly");

        cheats.startPrank(pauser);
        cheats.expectEmit(true, true, true, true, address(pausable));
        emit Paused(pauser, newPausedStatus);
        pausable.pause(newPausedStatus);
        cheats.stopPrank();

        require(pausable.paused() == newPausedStatus, "newPausedStatus not set correctly");
    }

    function testPause_RevertsWhenCalledByNotPauser(address notPauser, uint256 newPausedStatus) public {
        cheats.assume(notPauser != pauser);

        cheats.startPrank(notPauser);
        cheats.expectRevert(IPausable.OnlyPauser.selector);
        pausable.pause(newPausedStatus);
        cheats.stopPrank();
    }

    function testPauseAll(uint256 previousPausedStatus) public {
        cheats.startPrank(pauser);
        cheats.expectEmit(true, true, true, true, address(pausable));
        emit Paused(pauser, previousPausedStatus);
        pausable.pause(previousPausedStatus);
        cheats.stopPrank();

        require(pausable.paused() == previousPausedStatus, "previousPausedStatus not set correctly");

        cheats.startPrank(pauser);
        cheats.expectEmit(true, true, true, true, address(pausable));
        emit Paused(pauser, type(uint256).max);
        pausable.pauseAll();
        cheats.stopPrank();

        require(pausable.paused() == type(uint256).max, "newPausedStatus not set correctly");
    }

    function testPauseAll_RevertsWhenCalledByNotPauser(address notPauser) public {
        cheats.assume(notPauser != pauser);

        cheats.startPrank(notPauser);
        cheats.expectRevert(IPausable.OnlyPauser.selector);
        pausable.pauseAll();
        cheats.stopPrank();
    }

    function testPause_RevertsWhenTryingToUnpause(uint256 previousPausedStatus, uint256 newPausedStatus) public {
        // filter to only fuzzed inputs which would (improperly) flip any bits to '0'.
        cheats.assume(previousPausedStatus & newPausedStatus != previousPausedStatus);

        cheats.startPrank(pauser);
        cheats.expectEmit(true, true, true, true, address(pausable));
        emit Paused(pauser, previousPausedStatus);
        pausable.pause(previousPausedStatus);
        cheats.stopPrank();

        require(pausable.paused() == previousPausedStatus, "previousPausedStatus not set correctly");

        cheats.startPrank(pauser);
        cheats.expectRevert(IPausable.InvalidNewPausedStatus.selector);
        pausable.pause(newPausedStatus);
        cheats.stopPrank();
    }

    function testPauseSingleIndex(uint256 previousPausedStatus, uint8 indexToPause) public {
        uint256 newPausedStatus = (2 ** indexToPause);
        testPause(previousPausedStatus, newPausedStatus);
        require(pausable.paused(indexToPause), "index is not paused");
    }

    function testUnpause(uint256 previousPausedStatus, uint256 newPausedStatus) public {
        // filter out any fuzzed inputs which would (improperly) flip any bits to '1'.
        cheats.assume(~previousPausedStatus & ~newPausedStatus == ~previousPausedStatus);

        cheats.startPrank(pauser);
        cheats.expectEmit(true, true, true, true, address(pausable));
        emit Paused(pauser, previousPausedStatus);
        pausable.pause(previousPausedStatus);
        cheats.stopPrank();

        require(pausable.paused() == previousPausedStatus, "previousPausedStatus not set correctly");

        cheats.startPrank(pausable.pauserRegistry().unpauser());
        cheats.expectEmit(true, true, true, true, address(pausable));
        emit Unpaused(pausable.pauserRegistry().unpauser(), newPausedStatus);
        pausable.unpause(newPausedStatus);
        cheats.stopPrank();

        require(pausable.paused() == newPausedStatus, "newPausedStatus not set correctly");
    }

    function testUnpause_RevertsWhenCalledByNotUnpauser(address notUnpauser, uint256 newPausedStatus) public {
        cheats.assume(notUnpauser != pausable.pauserRegistry().unpauser());

        cheats.startPrank(notUnpauser);
        cheats.expectRevert(IPausable.OnlyUnpauser.selector);
        pausable.unpause(newPausedStatus);
        cheats.stopPrank();
    }


    function testUnpause_RevertsWhenTryingToPause(uint256 previousPausedStatus, uint256 newPausedStatus) public {
        // filter to only fuzzed inputs which would (improperly) flip any bits to '1'.
        cheats.assume(~previousPausedStatus & ~newPausedStatus != ~previousPausedStatus);

        cheats.startPrank(pauser);
        pausable.pause(previousPausedStatus);
        cheats.stopPrank();

        require(pausable.paused() == previousPausedStatus, "previousPausedStatus not set correctly");

        cheats.startPrank(pausable.pauserRegistry().unpauser());
        cheats.expectRevert(IPausable.InvalidNewPausedStatus.selector);
        pausable.unpause(newPausedStatus);
        cheats.stopPrank();
    }
}
