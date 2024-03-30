// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/core/PaymentCoordinator.sol";
import "src/contracts/strategies/StrategyBase.sol";

import "src/test/events/IPaymentCoordinatorEvents.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

/**
 * @notice Unit testing of the PaymentCoordinator contract
 * Contracts tested: PaymentCoordinator
 * Contracts not mocked: StrategyBase, PauserRegistry
 */
contract PaymentCoordinatorUnitTests is EigenLayerUnitTestSetup, IPaymentCoordinatorEvents {
    // Contract under test
    PaymentCoordinator public paymentCoordinator;
    PaymentCoordinator public paymentCoordinatorImplementation;

    // Mocks
    IERC20 token1;
    IERC20 token2;
    IERC20 token3;
    IStrategy strategyMock1;
    IStrategy strategyMock2;
    IStrategy strategyMock3;
    StrategyBase strategyImplementation;
    uint256 mockTokenInitialSupply = 10e50;

    // Config Variables
    /// @notice Max duration is 5 epochs (2 weeks * 5 = 10 weeks in seconds)
    uint64 MAX_PAYMENT_DURATION = 86400 * 70;
    
    /// @notice Lower bound start range is 3 months into the past
    uint64 LOWER_BOUND_START_RANGE = 86400 * 90;
    /// @notice Upper bound start range is 1 month into the future
    uint64 UPPER_BOUND_START_RANGE = 86400 * 30;

    /// @notice Delay in timestamp before a posted root can be claimed against
    uint64 activationDelay = 86400 * 7;
    /// @notice intervals(epochs) are 2 weeks
    uint64 calculationIntervalSeconds = 86400 * 14;
    /// @notice the commission for all operators across all avss
    uint16 globalCommissionBips = 1000;

    // PaymentCoordinator entities
    address paymentUpdater = address(1000);
    address defaultAVS = address(1001);
    address defaultClaimer = address(1002);
    address payAllSubmitter = address(1003);

    function setUp() public override {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        // Deploy PaymentCoordinator proxy and implementation
        paymentCoordinatorImplementation = new PaymentCoordinator(
            delegationManagerMock,
            strategyManagerMock,
            MAX_PAYMENT_DURATION,
            LOWER_BOUND_START_RANGE,
            UPPER_BOUND_START_RANGE
        );
        paymentCoordinator = PaymentCoordinator(
            address(
                new TransparentUpgradeableProxy(
                    address(paymentCoordinatorImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        PaymentCoordinator.initialize.selector,
                        address(this), // initOwner
                        pauserRegistry,
                        0, // 0 is initialPausedStatus
                        paymentUpdater,
                        activationDelay,
                        calculationIntervalSeconds,
                        globalCommissionBips
                    )
                )
            )
        );

        // Deploy mock token and strategy
        token1 = new ERC20PresetFixedSupply("dog wif hat", "MOCK1", mockTokenInitialSupply, address(this));
        token1 = new ERC20PresetFixedSupply("jeo boden", "MOCK2", mockTokenInitialSupply, address(this));
        token1 = new ERC20PresetFixedSupply("pepe wif avs", "MOCK3", mockTokenInitialSupply, address(this));

        strategyImplementation = new StrategyBase(strategyManagerMock);
        strategyMock1 = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, token1, pauserRegistry)
                )
            )
        );
        strategyMock2 = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, token2, pauserRegistry)
                )
            )
        );
        strategyMock3 = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, token3, pauserRegistry)
                )
            )
        );

        strategyManagerMock.setStrategyWhitelist(strategyMock1, true);
        strategyManagerMock.setStrategyWhitelist(strategyMock2, true);
        strategyManagerMock.setStrategyWhitelist(strategyMock3, true);

        // Exclude from fuzzed tests
        addressIsExcludedFromFuzzedInputs[address(paymentCoordinator)] = true;
    }
}

contract PaymentCoordinatorUnitTests_initializeAndSetters is PaymentCoordinatorUnitTests {}

contract PaymentCoordinatorUnitTests_payForRange is PaymentCoordinatorUnitTests {}

contract PaymentCoordinatorUnitTests_payAllForRange is PaymentCoordinatorUnitTests {}