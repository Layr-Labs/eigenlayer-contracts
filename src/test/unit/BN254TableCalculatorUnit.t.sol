// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/utils/EigenLayerMultichainUnitTestSetup.sol";
import "src/contracts/multichain/BN254TableCalculator.sol";

contract BN254TableCalculatorUnitTests is EigenLayerMultichainUnitTestSetup {
    using StdStyle for *;
    using ArrayLib for *;

    /// @notice Pointers to the BN254TableCalculator and its implementation
    BN254TableCalculator bn254TableCalculator;

    /// @notice The default operatorSet
    OperatorSet defaultOperatorSet = OperatorSet({avs: address(this), id: 0});

    /// @notice The default lookahead blocks for the slashable stake lookup
    uint LOOKAHEAD_BLOCKS = 50;

    function setUp() public override {
        EigenLayerMultichainUnitTestSetup.setUp();

        // Deploy BN254TableCalculator (immutable, non-upgradeable)
        bn254TableCalculator = new BN254TableCalculator(
            IKeyRegistrar(address(keyRegistrar)), IAllocationManager(address(allocationManagerMock)), LOOKAHEAD_BLOCKS
        );
    }
}
