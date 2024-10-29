// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/libraries/SlashingLib.sol";

contract SlashingLibUnitTests is Test {
    using SlashingLib for StakerScalingFactors;

    /// -----------------------------------------------------------------------
    /// Wad Math
    /// -----------------------------------------------------------------------

    // NOTE: 128 bit values are used to avoid overflow.

    function test_MulWad_Correctness(
        uint256 r
    ) public {
        uint256 x = bound(r, 0, type(uint128).max);
        uint256 y = bound(r, 0, type(uint128).max);

        assertEq(SlashingLib.mulWad(x, y), x * y / WAD);
    }

    function test_DivWad_Correctness(
        uint256 r
    ) public {
        uint256 x = bound(r, 0, type(uint128).max);
        uint256 y = bound(r, 0, type(uint128).max);

        if (y != 0) {
            assertEq(SlashingLib.divWad(x, y), x * WAD / y);
        } else {
            vm.expectRevert(); // div by zero
            SlashingLib.divWad(x, y);
        }
    }

    function test_MulWadRoundUp_Correctness(
        uint256 r
    ) public {
        uint256 x = bound(r, 0, type(uint128).max);
        uint256 y = bound(r, 0, type(uint128).max);

        uint256 result = SlashingLib.mulWadRoundUp(x, y);
        uint256 expected = x * y / WAD;

        if (mulmod(x, y, WAD) != 0) {
            assertEq(result, expected + 1);
        } else {
            assertEq(result, expected);
        }
    }

    /// -----------------------------------------------------------------------
    /// Getters
    /// -----------------------------------------------------------------------

    function test_GetDepositScalingFactor_InitiallyWad() public {
        StakerScalingFactors memory ssf = StakerScalingFactors({
            depositScalingFactor: 0,
            isBeaconChainScalingFactorSet: false,
            beaconChainScalingFactor: 0
        });
        assertEq(SlashingLib.getDepositScalingFactor(ssf), WAD);
    }

    function test_GetBeaconChainScalingFactor_InitiallyWad() public {
        StakerScalingFactors memory ssf = StakerScalingFactors({
            depositScalingFactor: 0,
            isBeaconChainScalingFactorSet: false,
            beaconChainScalingFactor: 0
        });
        assertEq(SlashingLib.getBeaconChainScalingFactor(ssf), WAD);
    }

    /// -----------------------------------------------------------------------
    /// Differential Test Helpers
    /// -----------------------------------------------------------------------

    function _differentialScaleSharesForQueuedWithdrawal(
        uint256 sharesToWithdraw,
        uint256 beaconChainScalingFactor,
        uint64 operatorMagnitude
    ) internal returns (uint256) {
        string memory args;
        args = vm.serializeString("args", "method", "scaleSharesForQueuedWithdrawal");
        args = vm.serializeUint("args", "sharesToWithdraw", sharesToWithdraw);
        args = vm.serializeUint("args", "beaconChainScalingFactor", beaconChainScalingFactor);
        args = vm.serializeUint("args", "operatorMagnitude", operatorMagnitude);

        string[] memory ffi = new string[](3);
        ffi[0] = "python3";
        ffi[1] = "src/test/ext/slashing-lib.t.py";
        ffi[2] = args;

        return abi.decode(vm.ffi(ffi), (uint256));
    }

    function _differentialScaleSharesForCompleteWithdrawal(
        uint256 scaledShares,
        uint256 beaconChainScalingFactor,
        uint64 operatorMagnitude
    ) internal returns (uint256) {
        string memory args;
        args = vm.serializeString("args", "method", "scaleSharesForCompleteWithdrawal");
        args = vm.serializeUint("args", "scaledShares", scaledShares);
        args = vm.serializeUint("args", "beaconChainScalingFactor", beaconChainScalingFactor);
        args = vm.serializeUint("args", "operatorMagnitude", operatorMagnitude);

        string[] memory ffi = new string[](3);
        ffi[0] = "python3";
        ffi[1] = "src/test/ext/slashing-lib.t.py";
        ffi[2] = args;

        return abi.decode(vm.ffi(ffi), (uint256));
    }

    /// -----------------------------------------------------------------------
    /// Differential Tests
    /// -----------------------------------------------------------------------

    function test_ScaleSharesForQueuedWithdrawal_Differential() public {
        // TODO: Document the constraints of each variable and fuzz them.
        uint256 sharesToWithdraw = 1e18;
        uint64 beaconChainScalingFactor = 1e18;
        uint64 operatorMagnitude = 1e18;

        uint256 result =
            _differentialScaleSharesForQueuedWithdrawal(sharesToWithdraw, beaconChainScalingFactor, operatorMagnitude);

        assertEq(
            SlashingLib.scaleSharesForQueuedWithdrawal(
                sharesToWithdraw,
                StakerScalingFactors({
                    isBeaconChainScalingFactorSet: false,
                    beaconChainScalingFactor: uint64(beaconChainScalingFactor),
                    depositScalingFactor: 0
                }),
                operatorMagnitude
            ),
            result
        );
    }
}
