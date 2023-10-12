// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";
import "openzeppelin-contracts/contracts/utils/Strings.sol";
import "../../../contracts/libraries/BN254.sol";

contract G2Operations is Test {
    using Strings for uint256;

    function mulGen(uint256 x) public returns (BN254.G2Point memory g2Point) {
        string[] memory inputs = new string[](5);
        inputs[0] = "go";
        inputs[1] = "run";
        inputs[2] = "src/test/ffi/util/g2mulGen.go";
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

    function mul(BN254.G2Point memory x, uint256 y) public returns (BN254.G2Point memory g2Point) {
        string[] memory inputs = new string[](9);
        inputs[0] = "go";
        inputs[1] = "run";
        inputs[2] = "src/test/ffi/util/g2mul.go";
        inputs[3] = x.X[1].toString();
        inputs[4] = x.X[0].toString();
        inputs[5] = x.Y[1].toString();
        inputs[6] = x.Y[0].toString();
        inputs[7] = y.toString(); 

        inputs[8] = "1";
        bytes memory res = vm.ffi(inputs);
        g2Point.X[1] = abi.decode(res, (uint256));

        inputs[8] = "2";
        res = vm.ffi(inputs);
        g2Point.X[0] = abi.decode(res, (uint256));

        inputs[8] = "3";
        res = vm.ffi(inputs);
        g2Point.Y[1] = abi.decode(res, (uint256));

        inputs[8] = "4";
        res = vm.ffi(inputs);
        g2Point.Y[0] = abi.decode(res, (uint256));
    }

    function add(BN254.G2Point memory x, BN254.G2Point memory y) public returns (BN254.G2Point memory g2Point) {
        string[] memory inputs = new string[](12);
        inputs[0] = "go";
        inputs[1] = "run";
        inputs[2] = "src/test/ffi/util/g2add.go";
        inputs[3] = x.X[1].toString(); 
        inputs[4] = x.X[0].toString(); 
        inputs[5] = x.Y[1].toString();
        inputs[6] = x.Y[0].toString();
        inputs[7] = y.X[1].toString();
        inputs[8] = y.X[0].toString();
        inputs[9] = y.Y[1].toString();
        inputs[10] = y.Y[0].toString();

        inputs[11] = "1";
        bytes memory res = vm.ffi(inputs);
        g2Point.X[1] = abi.decode(res, (uint256));

        inputs[11] = "2";
        res = vm.ffi(inputs);
        g2Point.X[0] = abi.decode(res, (uint256));

        inputs[11] = "3";
        res = vm.ffi(inputs);
        g2Point.Y[1] = abi.decode(res, (uint256));

        inputs[11] = "4";
        res = vm.ffi(inputs);
        g2Point.Y[0] = abi.decode(res, (uint256));
    }

}