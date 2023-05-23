// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./EigenLayerDeployer.t.sol";
import "./EigenLayerTestHelper.t.sol";
import "../contracts/operators/MerkleDelegationTerms.sol";

contract SlasherTests is EigenLayerTestHelper {
    ISlasher instance;
    uint256 constant HEAD = 0;
    address middleware = address(0xdeadbeef);
    address middleware_2 = address(0x009849);
    address middleware_3 = address(0x001000);
    address middleware_4 = address(0x002000);
    MerkleDelegationTerms delegationTerms;

     //performs basic deployment before each test
    function setUp() public override {
        super.setUp();
        delegationTerms = new MerkleDelegationTerms();
    }

    /**
     * @notice this function tests the slashing process by first freezing
     * the operator and then calling the strategyManager.slashShares()
     * to actually enforce the slashing conditions.
     */
    function testSlashing() public {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);

        // hardcoded inputs
        address[2] memory accounts = [acct_0, acct_1];
        uint256[2] memory depositAmounts;
        uint256 amountToDeposit = 1e18;
        address _operator = operator;
        strategyArray[0] = wethStrat;
        tokensArray[0] = weth;

        // have `_operator` make deposits in WETH strategy
        _testDepositWeth(_operator, amountToDeposit);
        // register `_operator` as an operator
        _testRegisterAsOperator(_operator, IDelegationTerms(_operator));

        // make deposit in WETH strategy from each of `accounts`, then delegate them to `_operator`
        for (uint256 i = 0; i < accounts.length; i++) {
            depositAmounts[i] = _testDepositWeth(accounts[i], amountToDeposit);
            _testDelegateToOperator(accounts[i], _operator);
        }

        uint256[] memory shareAmounts = new uint256[](1);
        shareAmounts[0] = depositAmounts[0];

        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;

        cheats.startPrank(_operator);
        slasher.optIntoSlashing(address(this));
        cheats.stopPrank();

        slasher.freezeOperator(_operator);

        uint256 prev_shares = delegation.operatorShares(_operator, strategyArray[0]);

        strategyManager.slashShares(_operator, acct_0, strategyArray, tokensArray, strategyIndexes, shareAmounts);

        require(
            delegation.operatorShares(_operator, strategyArray[0]) + shareAmounts[0] == prev_shares,
            "Malicious Operator slashed by incorrect amount"
        );
    }

    /**
     * @notice testing ownable permissions for slashing functions
     * addPermissionedContracts(), removePermissionedContracts()
     * and resetFrozenStatus().
     */
    function testOnlyOwnerFunctions(address incorrectCaller, address inputAddr)
        public
        fuzzedAddress(incorrectCaller)
        fuzzedAddress(inputAddr)
    {
        cheats.assume(incorrectCaller != slasher.owner());
        cheats.startPrank(incorrectCaller);
        address[] memory addressArray = new address[](1);
        addressArray[0] = inputAddr;
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        slasher.resetFrozenStatus(addressArray);
        cheats.stopPrank();
    }


    function testRecursiveCallRevert() public {
        //Register and opt into slashing with operator
        cheats.startPrank(operator);
        delegation.registerAsOperator(delegationTerms);
        slasher.optIntoSlashing(middleware);
        slasher.optIntoSlashing(middleware_2);
        slasher.optIntoSlashing(middleware_3);
        cheats.stopPrank();
        
        uint32 serveUntilBlock = uint32(block.number) + 1000;
        //these calls come from middlewares, we need more than 1 middleware to trigger the if clause on line 179
        // we need more than 2 middlewares to trigger the incorrect insertAfter being supplied by getCorrectValueForInsertAfter()
        cheats.startPrank(middleware);
        slasher.recordFirstStakeUpdate(operator, serveUntilBlock);
        cheats.stopPrank();

        cheats.startPrank(middleware_2);
        slasher.recordFirstStakeUpdate(operator, serveUntilBlock);
        cheats.stopPrank();
        
        cheats.startPrank(middleware_3);
        slasher.recordFirstStakeUpdate(operator, serveUntilBlock);
        cheats.stopPrank();
    
        cheats.startPrank(middleware);
        
        //we cannot add updateBlocks in the future so skip ahead to block 5.
        cheats.roll(5);
        //convert the middleware node address to a node number
        uint256 insertAfter = uint256(uint160(middleware));
        
        //Force fallbackRoutine to occur by specifying the wrong insertAfter. This then loops until we revert
        slasher.recordStakeUpdate(operator,5, serveUntilBlock, insertAfter);
        

    }

    function testRecordFirstStakeUpdate() public {
        
        //Register and opt into slashing with operator
        cheats.startPrank(operator);
        delegation.registerAsOperator(delegationTerms);
        slasher.optIntoSlashing(middleware);
        slasher.optIntoSlashing(middleware_3);
        cheats.stopPrank();

        uint32 serveUntilBlock = uint32(block.number) + 1000;

        cheats.startPrank(middleware_2);
        //unapproved slasher calls should fail
        cheats.expectRevert("Slasher.onlyRegisteredForService: Operator has not opted into slashing by caller");
        slasher.recordFirstStakeUpdate(operator, serveUntilBlock);
        cheats.stopPrank();

        cheats.startPrank(middleware);
        //valid conditions should succeed
        slasher.recordFirstStakeUpdate(operator, serveUntilBlock);

        //repeated calls to FirstStakeUpdate from the same middleware should fail.
        cheats.expectRevert("Slasher.recordFirstStakeUpdate: Appending middleware unsuccessful");
        slasher.recordFirstStakeUpdate(operator, serveUntilBlock);
        cheats.stopPrank();

        cheats.startPrank(middleware_3);
        //sequential calls from different approved slashing middlewares should succeed
        slasher.recordFirstStakeUpdate(operator, serveUntilBlock);
        cheats.stopPrank();
        
    }

    function testRecordStakeUpdate() public {
        ///Register and opt into slashing with operator
        cheats.startPrank(operator);
        delegation.registerAsOperator(delegationTerms);
        slasher.optIntoSlashing(middleware);
        slasher.optIntoSlashing(middleware_3);
        cheats.stopPrank();

        uint32 serveUntilBlock = uint32(block.number) + 1000;
        //convert the middleware node address to a node number
        uint256 insertAfter = uint256(uint160(middleware));

        cheats.startPrank(middleware);

        //calling before first stake update should fail
        cheats.expectRevert("Slasher.recordStakeUpdate: Removing middleware unsuccessful");
        slasher.recordStakeUpdate(operator,1, serveUntilBlock, insertAfter);

        //set up first stake update
        slasher.recordFirstStakeUpdate(operator, serveUntilBlock);

        cheats.expectRevert("Slasher.recordStakeUpdate: cannot provide update for future block");
        slasher.recordStakeUpdate(operator,5, serveUntilBlock, insertAfter);
        cheats.stopPrank();

        cheats.startPrank(middleware_2);
        //calling from unapproved middleware should fail
        cheats.expectRevert("Slasher.onlyRegisteredForService: Operator has not opted into slashing by caller");
        slasher.recordStakeUpdate(operator,1, serveUntilBlock, insertAfter);
        cheats.stopPrank();


        cheats.startPrank(middleware);
        //should succeed when given the correct settings
        cheats.roll(5);
        slasher.recordStakeUpdate(operator,3, serveUntilBlock, insertAfter);
        cheats.stopPrank();
    }

    function testOrderingRecordStakeUpdateVuln() public {
        ///Register and opt into slashing with operator
        cheats.startPrank(operator);
        delegation.registerAsOperator(delegationTerms);
        slasher.optIntoSlashing(middleware);
        slasher.optIntoSlashing(middleware_3);
        cheats.stopPrank();

        uint32 serveUntilBlock = uint32(block.number) + 1000;
        //convert the middleware node address to a node number
        uint256 insertAfter = uint256(uint160(middleware));

        cheats.startPrank(middleware);
        //set up first stake update so sizeOf check on line179 is equal to 1
        //uint256 sizeOf = slasher.sizeOfOperatorList(operator);
        uint256 sizeOf = slasher.operatorWhitelistedContractsLinkedListSize(operator);
        slasher.recordFirstStakeUpdate(operator, serveUntilBlock);
        require(sizeOf + 1 == slasher.operatorWhitelistedContractsLinkedListSize(operator));
        cheats.stopPrank();

        //calling recordStakeUpdate() before recordFirstStakeUpdate() for middleware_3 should fail 
        cheats.startPrank(middleware_3);
        sizeOf = slasher.operatorWhitelistedContractsLinkedListSize(operator);
        cheats.expectRevert("Slasher.recordStakeUpdate: Caller is not the list entrant");
        slasher.recordStakeUpdate(operator,1, serveUntilBlock, insertAfter);
        cheats.stopPrank();
        
    }

    function testOnlyRegisteredForService(address _slasher, uint32 _serveUntilBlock) public fuzzedAddress(_slasher) {
        cheats.prank(operator);
        delegation.registerAsOperator(delegationTerms);

        //slasher cannot call stake update unless operator has oped in
        cheats.prank(_slasher);
        cheats.expectRevert("Slasher.onlyRegisteredForService: Operator has not opted into slashing by caller");
        slasher.recordFirstStakeUpdate(operator, _serveUntilBlock);

        cheats.prank(_slasher);
        cheats.expectRevert("Slasher.onlyRegisteredForService: Operator has not opted into slashing by caller");
        slasher.recordStakeUpdate(operator, 1,_serveUntilBlock,1);

        cheats.prank(_slasher);
        cheats.expectRevert("Slasher.onlyRegisteredForService: Operator has not opted into slashing by caller");
        slasher.recordLastStakeUpdateAndRevokeSlashingAbility(operator, _serveUntilBlock);
    }

    function testOptIn(address _operator, address _slasher) public fuzzedAddress(_slasher) fuzzedAddress(_operator) {

        //cannot opt in until registered as operator
        cheats.prank(_operator);
        cheats.expectRevert("Slasher.optIntoSlashing: msg.sender is not a registered operator");
        slasher.optIntoSlashing(_slasher);

        //can opt in after registered as operator
        cheats.startPrank(_operator);
        delegation.registerAsOperator(delegationTerms);
        slasher.optIntoSlashing(_slasher);
        cheats.stopPrank();
    }

    function testFreezeOperator() public {
        cheats.prank(operator);
        delegation.registerAsOperator(delegationTerms);
        
        //cannot freeze until operator has oped in
        cheats.prank(middleware);
        cheats.expectRevert("Slasher.freezeOperator: msg.sender does not have permission to slash this operator");
        slasher.freezeOperator(operator);

        cheats.prank(operator);
        slasher.optIntoSlashing(middleware);

        //can freeze after operator has oped in
        cheats.prank(middleware);
        slasher.freezeOperator(operator);

        bool frozen = slasher.isFrozen(operator);
        require(frozen,"operator should be frozen");
    }

    function testResetFrozenOperator(address _attacker) public fuzzedAddress(_attacker) {
        cheats.assume(_attacker != slasher.owner());

        cheats.prank(operator);
        delegation.registerAsOperator(delegationTerms);

        cheats.prank(operator);
        slasher.optIntoSlashing(middleware);

        cheats.prank(middleware);
        slasher.freezeOperator(operator);

        address[] memory frozenAddresses = new address[](1);
        frozenAddresses[0] = operator;

        //no other address can unfreeze
        cheats.prank(_attacker);
        cheats.expectRevert("Ownable: caller is not the owner");
        slasher.resetFrozenStatus(frozenAddresses);

        //owner can unfreeze
        cheats.prank(slasher.owner());
        slasher.resetFrozenStatus(frozenAddresses);

        bool frozen = slasher.isFrozen(operator);
        require(!frozen,"operator should be unfrozen");
    }

    function testRecordLastStakeUpdateAndRevokeSlashingAbility() public {
        ///Register and opt into slashing with operator
        cheats.startPrank(operator);
        delegation.registerAsOperator(delegationTerms);
        slasher.optIntoSlashing(middleware);
        cheats.stopPrank();

        uint32 serveUntilBlock = 10;

        //stake update
        cheats.prank(middleware);
        slasher.recordFirstStakeUpdate(operator,serveUntilBlock);

        console.log("serveUntilBlock",slasher.getMiddlewareTimesIndexServeUntilBlock(operator,0));
        console.log("contractCanSlashOperatorUntil",slasher.contractCanSlashOperatorUntilBlock(operator,middleware));

        //middle can slash
        require(slasher.canSlash(operator,middleware),"middlewre should be able to slash");

        //revoke slashing
        cheats.prank(middleware);
        slasher.recordLastStakeUpdateAndRevokeSlashingAbility(operator,serveUntilBlock);

        cheats.roll(serveUntilBlock);
        //middleware can no longer slash
        require(!slasher.canSlash(operator,middleware),"middlewre should no longer be able to slash");
    }
}
