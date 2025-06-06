// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/utils/EigenLayerMultichainUnitTestSetup.sol";
import "src/contracts/multichain/BN254TableCalculator.sol";
import "src/test/utils/OperatorWalletLib.sol";

contract BN254TableCalculatorUnitTests is EigenLayerMultichainUnitTestSetup, OperatorWalletLib {
    using OperatorWalletLib for *;
    using StdStyle for *;
    using ArrayLib for *;

    /// @notice Pointers to the BN254TableCalculator and its implementation
    BN254TableCalculator bn254TableCalculator;

    /// @notice The default operatorSet
    OperatorSet defaultOperatorSet = OperatorSet({avs: address(this), id: 0});

    /// @notice Mapping of salt to ensure uniqueness of operators
    mapping(uint => bool) operatorSaltTaken;

    function setUp() public override {
        EigenLayerMultichainUnitTestSetup.setUp();

        // Deploy BN254TableCalculator (immutable, non-upgradeable)
        bn254TableCalculator =
            new BN254TableCalculator(IKeyRegistrar(address(keyRegistrar)), IAllocationManager(address(allocationManagerMock)));

        // Setup operatorSet in mocks
        allocationManagerMock.setIsOperatorSet(defaultOperatorSet, true);

        // Set the CurveType in the keyRegistrar
        keyRegistrar.configureOperatorSet(defaultOperatorSet, IKeyRegistrarTypes.CurveType.BN254);
    }

    /// @dev Creates and registers operators in the operatorSet
    function _createAndRegisterOperators(Randomness r) internal returns (Operator[] memory) {
        uint operatorCount = 1; //r.Uint256(1, 10);

        Operator[] memory operators = new Operator[](operatorCount);
        address[] memory operatorAddresses = new address[](operatorCount);
        for (uint i = 0; i < operatorCount; ++i) {
            operators[i] = _createOperatorWallet(r);
            address operatorAddress = operators[i].key.addr;
            operatorAddresses[i] = operatorAddress;

            // Generate signature data for operator
            bytes32 messageHash = keyRegistrar.getBN254KeyRegistrationMessageHash(operatorAddress, defaultOperatorSet, );

            // Register each operator in the keyRegistrar
            cheats.prank(operatorAddress);
            keyRegistrar.registerKey(operatorAddress, defaultOperatorSet, );
        }

        allocationManagerMock.setMembersInOperatorSet(defaultOperatorSet, operatorAddresses);

        return operators;
    }

    /// @dev Adds strategies to the operatorSet
    function _addStrategiesToOperatorSet(Randomness r) internal returns (IStrategy[] memory) {
        uint strategyCount = r.Uint256(1, 10);
        IStrategy[] memory strategies = r.StrategyArray(strategyCount);

        allocationManagerMock.setStrategiesInOperatorSet(defaultOperatorSet, strategies);

        return strategies;
    }

    /// @dev Sets the minimum slashable stake for the operators in the operatorSet
    function _setMinimumSlashableStake(address[] memory operators, IStrategy[] memory strategies, Randomness r)
        internal
        returns (uint[][] memory)
    {
        uint[][] memory minimumSlashableStake = new uint[][](operators.length);

        for (uint i = 0; i < operators.length; ++i) {
            minimumSlashableStake[i] = r.Uint256Array(strategies.length, 1, 100_000_000 ether);
        }

        allocationManagerMock.setMinimumSlashableStake(defaultOperatorSet, operators, strategies, minimumSlashableStake);

        return minimumSlashableStake;
    }

    /// @dev Initializes the operatorSet with operators, strategies, and minimum slashable stake
    /// @dev Returns the operators, strategies, and minimum slashable stake
    function _initializeOperatorSet(Randomness r) internal returns (address[] memory, IStrategy[] memory, uint[][] memory) {
        address[] memory operators = _createAndRegisterOperators(r);
        IStrategy[] memory strategies = _addStrategiesToOperatorSet(r);
        uint[][] memory minimumSlashableStake = _setMinimumSlashableStake(operators, strategies, r);

        return (operators, strategies, minimumSlashableStake);
    }

    function _createOperatorWallet(Randomness r) internal returns (Operator memory) {
        // Generate a unique name for the operator
        uint salt = r.Uint256(1, type(uint).max);
        while (operatorSaltTaken[salt]) salt = r.Uint256(1, type(uint).max);
        operatorSaltTaken[salt] = true;
        string memory name = Strings.toString(salt);

        // Create the operator wallet
        return createOperator(name);
    }

    function testRegistration(Randomness r) public rand(r) {
        Operator[] memory operators = _createAndRegisterOperators(r);
    }
}