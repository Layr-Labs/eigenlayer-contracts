// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "script/whitelist/delegationFaucet/DelegationFaucet.sol";
import "script/whitelist/ERC20PresetMinterPauser.sol";

import "src/test/EigenLayerTestHelper.t.sol";

contract DelegationFaucetTests is EigenLayerTestHelper {
    // EigenLayer contracts
    DelegationFaucet delegationFaucet;

    // M2 testing/mock contracts
    ERC20PresetMinterPauser public stakeToken;
    StrategyBase public stakeTokenStrat;

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    uint256 public constant DEFAULT_AMOUNT = 100e18;
    address owner = cheats.addr(1000);

    function setUp() public virtual override {
        EigenLayerDeployer.setUp();

        // Deploy ERC20 stakeToken, StrategyBase, and add StrategyBase to whitelist
        stakeToken = new ERC20PresetMinterPauser("StakeToken", "STK");
        stakeTokenStrat = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(baseStrategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, stakeToken, eigenLayerPauserReg)
                )
            )
        );
        cheats.startPrank(strategyManager.strategyWhitelister());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = stakeTokenStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        // Deploy DelegationFaucet, grant it admin/mint/pauser roles, etc.
        delegationFaucet = new DelegationFaucet(strategyManager, delegation, stakeToken, stakeTokenStrat);
        targetContract(address(delegationFaucet));
        stakeToken.grantRole(MINTER_ROLE, address(delegationFaucet));
    }

    /**
     * @notice Assertions in test
     * - Checks staker contract is deployed
     * - Checks token supply before/after minting
     * - Checks token balances are updated correctly for staker and strategy contracts
     * @param _operatorIndex is the index of the operator to use from the test-data/operators.json file
     * @param _depositAmount is the amount of stakeToken to mint to the staker and deposit into the strategy
     */
    function test_mintDepositAndDelegate_CheckBalancesAndDeploys(uint8 _operatorIndex, uint256 _depositAmount) public {
        cheats.assume(_operatorIndex < 15 && _depositAmount < DEFAULT_AMOUNT);
        if (_depositAmount == 0) {
            // Passing 0 as amount param defaults the amount to DEFAULT_AMOUNT constant
            _depositAmount = DEFAULT_AMOUNT;
        }
        // Setup Operator
        address operator = getOperatorAddress(_operatorIndex);
        address stakerContract = delegationFaucet.getStaker(operator);
        _registerOperator(operator);

        // Mint token to Staker, deposit minted amount into strategy, and delegate to operator
        uint256 supplyBefore = stakeToken.totalSupply();
        uint256 stratBalanceBefore = stakeToken.balanceOf(address(stakeTokenStrat));
        assertTrue(
            !Address.isContract(stakerContract),
            "test_mintDepositAndDelegate_CheckBalancesAndDeploys: staker contract shouldn't be deployed"
        );
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationFaucet.mintDepositAndDelegate(operator, signatureWithExpiry, bytes32(0), _depositAmount);
        assertTrue(
            Address.isContract(stakerContract),
            "test_mintDepositAndDelegate_CheckBalancesAndDeploys: staker contract not deployed"
        );
        uint256 supplyAfter = stakeToken.totalSupply();
        uint256 stratBalanceAfter = stakeToken.balanceOf(address(stakeTokenStrat));
        // Check token supply and balances
        assertEq(
            supplyAfter,
            supplyBefore + _depositAmount,
            "test_mintDepositAndDelegate_CheckBalancesAndDeploys: token supply not updated correctly"
        );
        assertEq(
            stratBalanceAfter,
            stratBalanceBefore + _depositAmount,
            "test_mintDepositAndDelegate_CheckBalancesAndDeploys: strategy balance not updated correctly"
        );
        assertEq(
            stakeToken.balanceOf(stakerContract),
            0,
            "test_mintDepositAndDelegate_CheckBalancesAndDeploys: staker balance should be 0"
        );
    }

    /**
     * @notice Check the before/after values for strategy shares and operator shares
     * @param _operatorIndex is the index of the operator to use from the test-data/operators.json file
     * @param _depositAmount is the amount of stakeToken to mint to the staker and deposit into the strategy
     */
    function test_mintDepositAndDelegate_StrategyAndOperatorShares(
        uint8 _operatorIndex,
        uint256 _depositAmount
    ) public {
        cheats.assume(_operatorIndex < 15 && _depositAmount < DEFAULT_AMOUNT);
        if (_depositAmount == 0) {
            _depositAmount = DEFAULT_AMOUNT;
        }
        // Setup Operator
        address operator = getOperatorAddress(_operatorIndex);
        address stakerContract = delegationFaucet.getStaker(operator);
        _registerOperator(operator);

        // Mint token to Staker, deposit minted amount into strategy, and delegate to operator
        uint256 stakerSharesBefore = strategyManager.stakerStrategyShares(stakerContract, stakeTokenStrat);
        uint256 operatorSharesBefore = delegation.operatorShares(operator, stakeTokenStrat);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationFaucet.mintDepositAndDelegate(operator, signatureWithExpiry, bytes32(0), _depositAmount);

        uint256 stakerSharesAfter = strategyManager.stakerStrategyShares(stakerContract, stakeTokenStrat);
        uint256 operatorSharesAfter = delegation.operatorShares(operator, stakeTokenStrat);
        assertTrue(
            delegation.delegatedTo(stakerContract) == operator,
            "test_mintDepositAndDelegate_StrategyAndOperatorShares: delegated address not set appropriately"
        );
        assertTrue(
            delegation.isDelegated(stakerContract),
            "test_mintDepositAndDelegate_StrategyAndOperatorShares: delegated status not set appropriately"
        );
        assertEq(
            stakerSharesAfter,
            stakerSharesBefore + _depositAmount,
            "test_mintDepositAndDelegate_StrategyAndOperatorShares: staker shares not updated correctly"
        );
        assertEq(
            operatorSharesAfter,
            operatorSharesBefore + _depositAmount,
            "test_mintDepositAndDelegate_StrategyAndOperatorShares: operator shares not updated correctly"
        );
    }

    /**
     * @notice Invariant test the before/after values for strategy shares and operator shares from multiple runs
     */
    /// forge-config: default.invariant.runs = 5
    /// forge-config: default.invariant.depth = 20
    function invariant_test_mintDepositAndDelegate_StrategyAndOperatorShares() public {
        // cheats.assume(_operatorIndex < 15 && _depositAmount < DEFAULT_AMOUNT);
        // Setup Operator
        address operator = getOperatorAddress(0);
        address stakerContract = delegationFaucet.getStaker(operator);
        _registerOperator(operator);

        // Mint token to Staker, deposit minted amount into strategy, and delegate to operator
        uint256 stakerSharesBefore = strategyManager.stakerStrategyShares(stakerContract, stakeTokenStrat);
        uint256 operatorSharesBefore = delegation.operatorShares(operator, stakeTokenStrat);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationFaucet.mintDepositAndDelegate(operator, signatureWithExpiry, bytes32(0), DEFAULT_AMOUNT);

        uint256 stakerSharesAfter = strategyManager.stakerStrategyShares(stakerContract, stakeTokenStrat);
        uint256 operatorSharesAfter = delegation.operatorShares(operator, stakeTokenStrat);
        assertTrue(
            delegation.delegatedTo(stakerContract) == operator,
            "test_mintDepositAndDelegate_StrategyAndOperatorShares: delegated address not set appropriately"
        );
        assertTrue(
            delegation.isDelegated(stakerContract),
            "test_mintDepositAndDelegate_StrategyAndOperatorShares: delegated status not set appropriately"
        );
        assertEq(
            stakerSharesAfter,
            stakerSharesBefore + DEFAULT_AMOUNT,
            "test_mintDepositAndDelegate_StrategyAndOperatorShares: staker shares not updated correctly"
        );
        assertEq(
            operatorSharesAfter,
            operatorSharesBefore + DEFAULT_AMOUNT,
            "test_mintDepositAndDelegate_StrategyAndOperatorShares: operator shares not updated correctly"
        );
    }

    /**
     * @param _operatorIndex is the index of the operator to use from the test-data/operators.json file
     */
    function test_mintDepositAndDelegate_RevertsIf_UnregisteredOperater(uint8 _operatorIndex) public {
        cheats.assume(_operatorIndex < 15);
        address operator = getOperatorAddress(_operatorIndex);
        // Unregistered operator should revert
        cheats.expectRevert("DelegationFaucet: Operator not registered");
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationFaucet.mintDepositAndDelegate(operator, signatureWithExpiry, bytes32(0), DEFAULT_AMOUNT);
    }

    function test_depositIntoStrategy_IncreaseShares(uint8 _operatorIndex, uint256 _depositAmount) public {
        cheats.assume(_operatorIndex < 15 && _depositAmount < DEFAULT_AMOUNT);
        if (_depositAmount == 0) {
            _depositAmount = DEFAULT_AMOUNT;
        }
        // Setup Operator
        address operator = getOperatorAddress(_operatorIndex);
        address stakerContract = delegationFaucet.getStaker(operator);
        _registerOperator(operator);

        // Mint token to Staker, deposit minted amount into strategy, and delegate to operator
        uint256 stakerSharesBefore = strategyManager.stakerStrategyShares(stakerContract, stakeTokenStrat);
        uint256 operatorSharesBefore = delegation.operatorShares(operator, stakeTokenStrat);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationFaucet.mintDepositAndDelegate(operator, signatureWithExpiry, bytes32(0), DEFAULT_AMOUNT);

        uint256 stakerSharesAfter = strategyManager.stakerStrategyShares(stakerContract, stakeTokenStrat);
        uint256 operatorSharesAfter = delegation.operatorShares(operator, stakeTokenStrat);
        assertTrue(
            delegation.delegatedTo(stakerContract) == operator,
            "test_mintDepositAndDelegate_IncreaseShares: delegated address not set appropriately"
        );
        assertTrue(
            delegation.isDelegated(stakerContract),
            "test_mintDepositAndDelegate_IncreaseShares: delegated status not set appropriately"
        );
        assertEq(
            stakerSharesAfter,
            stakerSharesBefore + DEFAULT_AMOUNT,
            "test_mintDepositAndDelegate_IncreaseShares: staker shares not updated correctly"
        );
        assertEq(
            operatorSharesAfter,
            operatorSharesBefore + DEFAULT_AMOUNT,
            "test_mintDepositAndDelegate_IncreaseShares: operator shares not updated correctly"
        );

        // Deposit more into strategy
        stakerSharesBefore = stakerSharesAfter;
        operatorSharesBefore = operatorSharesAfter;
        delegationFaucet.depositIntoStrategy(stakerContract, stakeTokenStrat, stakeToken, _depositAmount);
        stakerSharesAfter = strategyManager.stakerStrategyShares(stakerContract, stakeTokenStrat);
        operatorSharesAfter = delegation.operatorShares(operator, stakeTokenStrat);

        assertEq(
            stakerSharesAfter,
            stakerSharesBefore + _depositAmount,
            "test_mintDepositAndDelegate_IncreasesShares: staker shares not updated correctly"
        );
        assertEq(
            operatorSharesAfter,
            operatorSharesBefore + _depositAmount,
            "test_mintDepositAndDelegate_IncreasesShares: operator shares not updated correctly"
        );
    }

    function test_transfer_TransfersERC20(uint8 _operatorIndex, address _to, uint256 _transferAmount) public {
        cheats.assume(_operatorIndex < 15);
        // Setup Operator
        address operator = getOperatorAddress(_operatorIndex);
        address stakerContract = delegationFaucet.getStaker(operator);
        _registerOperator(operator);

        // Mint token to Staker, deposit minted amount into strategy, and delegate to operator
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationFaucet.mintDepositAndDelegate(operator, signatureWithExpiry, bytes32(0), DEFAULT_AMOUNT);


        ERC20PresetMinterPauser mockToken = new ERC20PresetMinterPauser("MockToken", "MTK");
        mockToken.mint(stakerContract, _transferAmount);

        uint256 stakerBalanceBefore = mockToken.balanceOf(stakerContract);
        uint256 toBalanceBefore = mockToken.balanceOf(_to);
        delegationFaucet.transfer(stakerContract, address(mockToken), _to, _transferAmount);
        uint256 stakerBalanceAfter = mockToken.balanceOf(stakerContract);
        uint256 toBalanceAfter = mockToken.balanceOf(_to);
        assertEq(
            stakerBalanceBefore,
            stakerBalanceAfter + _transferAmount,
            "test_transfer_TransfersERC20: staker balance not updated correctly"
        );
        assertEq(
            toBalanceBefore + _transferAmount,
            toBalanceAfter,
            "test_transfer_TransfersERC20: to balance not updated correctly"
        );
    }

    

    function _registerOperator(address _operator) internal {
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: _operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        _testRegisterAsOperator(_operator, operatorDetails);
    }
}
