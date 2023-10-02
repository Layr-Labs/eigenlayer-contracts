// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";
import "openzeppelin-contracts/contracts/utils/Strings.sol";
import "../../../contracts/libraries/BN254.sol";

contract G2Operations is Test {
    using Strings for uint256;

    function mul(uint256 x) public returns (BN254.G2Point memory g2Point) {
        string[] memory inputs = new string[](5);
        inputs[0] = "go";
        inputs[1] = "run";
        inputs[2] = "src/test/ffi/go/g2mul.go";
        inputs[3] = x.toString(); 

        inputs[4] = "1";
        bytes memory res = vm.ffi(inputs);
        g2Point.X[1] = abi.decode(res, (uint256));

        inputs[4] = "2";
        res = vm.ffi(inputs);
        g2Point.X[0] = abi.decode(res, (uint256));

        inputs[4] = "3";
        res = vm.ffi(inputs);
        g2Point.Y[1] = abi.decode(res, (uint256));

        inputs[4] = "4";
        res = vm.ffi(inputs);
        g2Point.Y[0] = abi.decode(res, (uint256));
    }

}