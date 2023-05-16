// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "../../contracts/strategies/StrategyBase.sol";
import "../../contracts/strategies/StrategyBaseTVLLimits.sol";
import "../../contracts/permissions/PauserRegistry.sol";

import "../mocks/StrategyManagerMock.sol";
import "../mocks/ERC20_SetTransferReverting_Mock.sol";

import "forge-std/Test.sol";

contract StrategyBaseTVLLimitsUnitTests is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;
    IStrategyManager public strategyManager;
    IERC20 public underlyingToken;
    StrategyBaseTVLLimits public strategyBaseTVLLimitsImplementation;
    StrategyBaseTVLLimits public strategy;

    address public pauser = address(555);
    address public unpauser = address(999);

    uint256 initialSupply = 1e65;
    address initialOwner = address(this);

    uint256 maxDeposits = 100;
    uint256 maxPerDeposit = 5;

    function setUp() virtual public {
        proxyAdmin = new ProxyAdmin();

        pauserRegistry = new PauserRegistry(pauser, unpauser);

        strategyManager = new StrategyManagerMock();

        underlyingToken = new ERC20PresetFixedSupply("Test Token", "TEST", initialSupply, initialOwner);

        strategyBaseTVLLimitsImplementation = new StrategyBaseTVLLimits(strategyManager);

        strategy = StrategyBaseTVLLimits(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyBaseTVLLimitsImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken, pauserRegistry)
                )
            )
        );
    }

    function testSetTVLLimits(uint256 maxDepositsFuzzedInput, uint256 maxPerDepositFuzzedInput) public {
        cheats.startPrank(pauser);
        strategy.setTVLLimits(maxPerDepositFuzzedInput, maxDepositsFuzzedInput);
        cheats.stopPrank();
        (uint256 _maxPerDeposit, uint256 _maxDeposits) = strategy.getTVLLimits();

        assertEq(_maxPerDeposit, maxPerDepositFuzzedInput);
        assertEq(_maxDeposits, maxDepositsFuzzedInput);
    }

    function testSetTVLLimitsFailsWhenNotCalledByPauser(uint256 maxDepositsFuzzedInput, uint256 maxPerDepositFuzzedInput, address notPauser) public {
        cheats.assume(notPauser != address(proxyAdmin));
        cheats.assume(notPauser != pauser);
        cheats.startPrank(notPauser);
        cheats.expectRevert(bytes("msg.sender is not permissioned as pauser"));
        strategy.setTVLLimits(maxPerDepositFuzzedInput, maxDepositsFuzzedInput);
        cheats.stopPrank();
    }

    function testDepositMoreThanMaxPerDeposit(uint256 maxDepositsFuzzedInput, uint256 maxPerDepositFuzzedInput, uint256 amount) public {
        cheats.assume(amount > maxPerDepositFuzzedInput);
        cheats.startPrank(pauser);
        strategy.setTVLLimits(maxPerDepositFuzzedInput, maxDepositsFuzzedInput);
        cheats.stopPrank();

        cheats.startPrank(address(strategyManager));
        cheats.expectRevert(bytes("StrategyBaseTVLLimits: max per deposit exceeded"));
        strategy.deposit(underlyingToken, amount);
        cheats.stopPrank();
    }

    function testDepositMorethanMaxDeposits() public {
        maxDeposits = 1e12;
        maxPerDeposit = 3e11;
        uint256 numDeposits = maxDeposits / maxPerDeposit;

        underlyingToken.transfer(address(strategyManager), maxDeposits);
        cheats.startPrank(pauser);
        strategy.setTVLLimits(maxPerDeposit, maxDeposits);
        cheats.stopPrank();

        cheats.startPrank(address(strategyManager));
        for (uint256 i = 0; i < numDeposits; i++) {
            underlyingToken.transfer(address(strategy), maxPerDeposit);
            strategy.deposit(underlyingToken, maxPerDeposit);
        }
        cheats.stopPrank();

        underlyingToken.transfer(address(strategy), maxPerDeposit);
        require(underlyingToken.balanceOf(address(strategy)) > maxDeposits, "bad test setup");

        cheats.startPrank(address(strategyManager));
        cheats.expectRevert(bytes("StrategyBaseTVLLimits: max deposits exceeded"));
        strategy.deposit(underlyingToken, maxPerDeposit);
        cheats.stopPrank();
    }

    function testDepositValidAmount(uint256 depositAmount) public {
        maxDeposits = 1e12;
        maxPerDeposit = 3e11;
        cheats.assume(depositAmount > 0);
        cheats.assume(depositAmount < maxPerDeposit);

        cheats.startPrank(pauser);
        strategy.setTVLLimits(maxPerDeposit, maxDeposits);
        cheats.stopPrank();

        // we need to actually transfer the tokens to the strategy to avoid underflow in the `deposit` calculation
        underlyingToken.transfer(address(strategy), depositAmount);

        uint256 sharesBefore = strategy.totalShares();
        cheats.startPrank(address(strategyManager));
        strategy.deposit(underlyingToken, depositAmount);
        cheats.stopPrank();

        require(strategy.totalShares() == depositAmount + sharesBefore, "total shares not updated correctly");
    }

    /// @notice General-purpose test, re-useable, handles whether the deposit should revert or not and returns 'true' if it did revert.
    function testDeposit_WithTVLLimits(uint256 maxDepositsFuzzedInput, uint256 maxPerDepositFuzzedInput, uint256 depositAmount)
        public returns (bool depositReverted)
    {
        // need to filter this to make sure the deposit amounts can actually be transferred
        cheats.assume(depositAmount <= initialSupply);
        // set TVL limits
        cheats.startPrank(pauser);
        strategy.setTVLLimits(maxPerDepositFuzzedInput, maxDepositsFuzzedInput);
        cheats.stopPrank();

        // we need to calculate this before transferring tokens to the strategy
        uint256 expectedSharesOut = strategy.underlyingToShares(depositAmount);

        // we need to actually transfer the tokens to the strategy to avoid underflow in the `deposit` calculation
        underlyingToken.transfer(address(strategy), depositAmount);

        if (depositAmount > maxPerDepositFuzzedInput) {
            cheats.startPrank(address(strategyManager));
            cheats.expectRevert("StrategyBaseTVLLimits: max per deposit exceeded");
            strategy.deposit(underlyingToken, depositAmount);
            cheats.stopPrank();

            // transfer the tokens back from the strategy to not mess up the state
            cheats.startPrank(address(strategy));
            underlyingToken.transfer(address(this), depositAmount);
            cheats.stopPrank();

            // return 'true' since the call to `deposit` reverted
            return true;
        } else if (underlyingToken.balanceOf(address(strategy)) > maxDepositsFuzzedInput) {
            cheats.startPrank(address(strategyManager));
            cheats.expectRevert("StrategyBaseTVLLimits: max deposits exceeded");
            strategy.deposit(underlyingToken, depositAmount);
            cheats.stopPrank();

            // transfer the tokens back from the strategy to not mess up the state
            cheats.startPrank(address(strategy));
            underlyingToken.transfer(address(this), depositAmount);
            cheats.stopPrank();

            // return 'true' since the call to `deposit` reverted
            return true;
        } else {
            uint256 totalSharesBefore = strategy.totalShares();
            if (expectedSharesOut == 0) {
                cheats.startPrank(address(strategyManager));
                cheats.expectRevert("StrategyBase.deposit: newShares cannot be zero");
                strategy.deposit(underlyingToken, depositAmount);
                cheats.stopPrank();

                // transfer the tokens back from the strategy to not mess up the state
                cheats.startPrank(address(strategy));
                underlyingToken.transfer(address(this), depositAmount);
                cheats.stopPrank();

                // return 'true' since the call to `deposit` reverted
                return true;
            } else {
                cheats.startPrank(address(strategyManager));
                strategy.deposit(underlyingToken, depositAmount);
                cheats.stopPrank();

                require(strategy.totalShares() == expectedSharesOut + totalSharesBefore, "total shares not updated correctly");

                // return 'false' since the call to `deposit` succeeded
                return false;
            }
        }
    }

}