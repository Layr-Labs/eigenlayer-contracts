// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../src/contracts/interfaces/IStrategyManager.sol";
import "../../src/contracts/interfaces/IStrategy.sol";
import "../../src/contracts/interfaces/IDelegationManager.sol";
import "../../src/contracts/strategies/StrategyBase.sol";
import "../../src/contracts/middleware/BLSRegistry.sol";

import "../../src/test/mocks/ServiceManagerMock.sol";
import "../../src/test/mocks/PublicKeyCompendiumMock.sol";
import "../../src/test/mocks/MiddlewareVoteWeigherMock.sol";



import "../../script/whitelist/ERC20PresetMinterPauser.sol";

import "../../script/whitelist/Staker.sol";
import "../../script/whitelist/Whitelister.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Create2.sol";

import "./EigenLayerTestHelper.t.sol";
import "./Delegation.t.sol";

import "forge-std/Test.sol";

contract WhitelisterTests is EigenLayerTestHelper {

    ERC20PresetMinterPauser dummyToken;
    IStrategy dummyStrat;
    IStrategy dummyStratImplementation;
    Whitelister whiteLister;

    BLSRegistry blsRegistry;
    BLSRegistry blsRegistryImplementation;


    ServiceManagerMock dummyServiceManager;
    BLSPublicKeyCompendiumMock dummyCompendium;
    MiddlewareRegistryMock dummyReg;

    

    modifier fuzzedAmounts(uint256 ethAmount, uint256 eigenAmount){
        cheats.assume(ethAmount >= 0 && ethAmount <= 1e18);
        cheats.assume(eigenAmount >= 0 && eigenAmount <= 1e18);
        _;
    }

    uint256 AMOUNT;

    // packed info used to help handle stack-too-deep errors
    struct DataForTestWithdrawal {
        IStrategy[] delegatorStrategies;
        uint256[] delegatorShares;
        IStrategyManager.WithdrawerAndNonce withdrawerAndNonce;
    }

    function setUp() public virtual override{
        EigenLayerDeployer.setUp();

        emptyContract = new EmptyContract();

        dummyCompendium = new BLSPublicKeyCompendiumMock();
        blsRegistry = BLSRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        dummyToken = new ERC20PresetMinterPauser("dummy staked ETH", "dsETH");
        dummyStratImplementation = new StrategyBase(strategyManager);
        dummyStrat = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                        address(dummyStratImplementation),
                        address(eigenLayerProxyAdmin),
                        abi.encodeWithSelector(StrategyBase.initialize.selector, dummyToken, eigenLayerPauserReg)
                    )
                )
        );

        whiteLister = new Whitelister(strategyManager, delegation, dummyToken, dummyStrat, blsRegistry);
        whiteLister.transferOwnership(theMultiSig);
        AMOUNT = whiteLister.DEFAULT_AMOUNT();

        dummyToken.grantRole(keccak256("MINTER_ROLE"), address(whiteLister));
        dummyToken.grantRole(keccak256("PAUSER_ROLE"), address(whiteLister));  

        dummyToken.grantRole(keccak256("MINTER_ROLE"), theMultiSig);
        dummyToken.grantRole(keccak256("PAUSER_ROLE"), theMultiSig);

        dummyToken.revokeRole(keccak256("MINTER_ROLE"), address(this));  
        dummyToken.revokeRole(keccak256("PAUSER_ROLE"), address(this));  


        dummyServiceManager  = new ServiceManagerMock(slasher);
        blsRegistryImplementation = new BLSRegistry(strategyManager, dummyServiceManager, 2, dummyCompendium);

        uint256[] memory _quorumBips = new uint256[](2);
        // split 60% ETH quorum, 40% EIGEN quorum
        _quorumBips[0] = 6000;
        _quorumBips[1] = 4000;

        VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] memory ethStratsAndMultipliers =
                new VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[](1);
            ethStratsAndMultipliers[0].strategy = wethStrat;
            ethStratsAndMultipliers[0].multiplier = 1e18;
        VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] memory eigenStratsAndMultipliers =
                new VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[](1);
            eigenStratsAndMultipliers[0].strategy = eigenStrat;
            eigenStratsAndMultipliers[0].multiplier = 1e18;

        eigenLayerProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(blsRegistry))),
                address(blsRegistryImplementation),
                abi.encodeWithSelector(BLSRegistry.initialize.selector, address(whiteLister), true, _quorumBips, ethStratsAndMultipliers, eigenStratsAndMultipliers)
            );
        
        dummyReg = new MiddlewareRegistryMock(
             dummyServiceManager,
             strategyManager
        );

        fuzzedAddressMapping[address(whiteLister)] = true;

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();
    }

    function testWhitelistingOperator(address operator) public fuzzedAddress(operator){
        cheats.startPrank(operator);
        IDelegationTerms dt = IDelegationTerms(address(89));
        delegation.registerAsOperator(dt);
        cheats.stopPrank();

        cheats.startPrank(theMultiSig);
        whiteLister.whitelist(operator);
        cheats.stopPrank();

        assertTrue(blsRegistry.whitelisted(operator) == true, "operator not added to whitelist");
    }

    function testWhitelistDepositIntoStrategy(address operator, uint256 depositAmount) external fuzzedAddress(operator){
        cheats.assume(depositAmount < AMOUNT);
        testWhitelistingOperator(operator);

        cheats.startPrank(theMultiSig);
        address staker = whiteLister.getStaker(operator);
        dummyToken.mint(staker, AMOUNT);

        whiteLister.depositIntoStrategy(staker, dummyStrat, dummyToken, depositAmount);
        cheats.stopPrank();
    }

    function testCallStakerFromNonWhitelisterAddress(address nonWhitelister, bytes memory data) external fuzzedAddress(nonWhitelister){
        testWhitelistingOperator(operator);
        address staker = whiteLister.getStaker(operator);

        cheats.startPrank(nonWhitelister);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        Staker(staker).callAddress(address(strategyManager), data);
    }

    function testNonWhitelistedOperatorRegistration(BN254.G1Point memory pk, string memory socket ) external {
        cheats.startPrank(operator);
        IDelegationTerms dt = IDelegationTerms(address(89));
        delegation.registerAsOperator(dt);
        cheats.stopPrank();

        cheats.expectRevert(bytes("BLSRegistry._registerOperator: not whitelisted"));
        blsRegistry.registerOperator(1, pk, socket);
    }

    

    function testWhitelistQueueWithdrawal(
            address operator 
        ) 
            public  fuzzedAddress(operator)
        {

        address staker = whiteLister.getStaker(operator);
        cheats.assume(staker!=operator);
        _testRegisterAsOperator(operator, IDelegationTerms(operator));

        {
            cheats.startPrank(theMultiSig);
            whiteLister.whitelist(operator);
            cheats.stopPrank();

            cheats.startPrank(operator);
            slasher.optIntoSlashing(address(dummyServiceManager));
            dummyReg.registerOperator(operator, uint32(block.timestamp) + 3 days);
            cheats.stopPrank();
        }

        // packed data structure to deal with stack-too-deep issues
        DataForTestWithdrawal memory dataForTestWithdrawal;

        // scoped block to deal with stack-too-deep issues
        {
            //delegator-specific information
            (IStrategy[] memory delegatorStrategies, uint256[] memory delegatorShares) =
                strategyManager.getDeposits(staker);
            dataForTestWithdrawal.delegatorStrategies = delegatorStrategies;
            dataForTestWithdrawal.delegatorShares = delegatorShares;

            IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce = 
                IStrategyManager.WithdrawerAndNonce({
                    withdrawer: staker,
                    // harcoded nonce value
                    nonce: 0
                }
            );
            dataForTestWithdrawal.withdrawerAndNonce = withdrawerAndNonce;
        }

        uint256[] memory strategyIndexes = new uint256[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        {
            // hardcoded values
            strategyIndexes[0] = 0;
            tokensArray[0] = dummyToken;
        }

        _testQueueWithdrawal(
            staker,
            dataForTestWithdrawal.delegatorStrategies,
            dataForTestWithdrawal.delegatorShares,
            strategyIndexes
        );

        {
            uint256 balanceBeforeWithdrawal = dummyToken.balanceOf(staker);

            _testCompleteQueuedWithdrawal(
                staker,
                dataForTestWithdrawal.delegatorStrategies,
                tokensArray,
                dataForTestWithdrawal.delegatorShares,
                operator,
                dataForTestWithdrawal.withdrawerAndNonce,
                uint32(block.number),
                1
            );
            emit log_named_uint("Balance Before Withdrawal", dummyToken.balanceOf(staker));
            emit log_named_uint("Balance After Withdrawal", balanceBeforeWithdrawal);
        
            require(dummyToken.balanceOf(staker) == balanceBeforeWithdrawal + AMOUNT, "balance not incremented");

        }        
    }

    function _testQueueWithdrawal(
        address staker,
        IStrategy[] memory strategyArray,
        uint256[] memory shareAmounts,
        uint256[] memory strategyIndexes
    )
        internal
    {
        cheats.startPrank(theMultiSig);
        whiteLister.queueWithdrawal(
            staker,
            strategyIndexes,
            strategyArray,
            shareAmounts,
            staker,
            true
        );
        cheats.stopPrank();
    }

     function _testCompleteQueuedWithdrawal(
        address staker,
        IStrategy[] memory strategyArray,
        IERC20[] memory tokensArray,
        uint256[] memory shareAmounts,
        address delegatedTo,
        IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce,
        uint32 withdrawalStartBlock,
        uint256 middlewareTimesIndex
    )
        internal
    {
        IStrategyManager.QueuedWithdrawal memory queuedWithdrawal = IStrategyManager.QueuedWithdrawal({
            strategies: strategyArray,
            shares: shareAmounts,
            depositor: staker,
            withdrawerAndNonce: withdrawerAndNonce,
            withdrawalStartBlock: withdrawalStartBlock,
            delegatedAddress: delegatedTo
        });

        // emit log("*******************COMPLETE***************************");
        // emit log_named_address("delegatedAddress", delegatedTo);
        // emit log_named_uint("withdrawalStartBlock", withdrawalStartBlock);
        // emit log_named_uint("withdrawerAndNonce.Nonce", withdrawerAndNonce.nonce);
        // emit log_named_address("withdrawerAndNonce.Adress", withdrawerAndNonce.withdrawer);
        // emit log_named_address("depositor", staker);
        // emit log("***********************************************************************");

        cheats.startPrank(theMultiSig);
        whiteLister.completeQueuedWithdrawal(staker, queuedWithdrawal, tokensArray, middlewareTimesIndex, true);
        cheats.stopPrank();
    }
    
    function testWhitelistTransfer(address operator, address receiver) public fuzzedAddress(receiver) {
        address staker = whiteLister.getStaker(operator);

        testWhitelistQueueWithdrawal(operator);

        cheats.startPrank(theMultiSig);

        whiteLister.transfer(staker, address(dummyToken), receiver, AMOUNT);
        cheats.stopPrank();
        require(dummyToken.balanceOf(receiver) == AMOUNT, "receiver hasn't received tokens");
    }
}