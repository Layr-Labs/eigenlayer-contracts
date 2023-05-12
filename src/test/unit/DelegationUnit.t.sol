// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";

import "../mocks/StrategyManagerMock.sol";

import "../mocks/SlasherMock.sol";
import "../EigenLayerTestHelper.t.sol";
import "../mocks/ERC20Mock.sol";
import "../mocks/DelegationTermsMock.sol";
import "../Delegation.t.sol";

contract DelegationUnitTests is EigenLayerTestHelper {

    StrategyManagerMock strategyManagerMock;
    SlasherMock slasherMock;
    DelegationManager delegationContract;
    DelegationTermsMock delegationTermsMock;
    DelegationManager delegationContractImplementation;
    StrategyBase strategyImplementation;
    StrategyBase strategyMock;


    uint256 GWEI_TO_WEI = 1e9;

    event OnDelegationReceivedCallFailure(IDelegationTerms indexed delegationTerms, bytes32 returnData);
    event OnDelegationWithdrawnCallFailure(IDelegationTerms indexed delegationTerms, bytes32 returnData);


    function setUp() override virtual public{
        EigenLayerDeployer.setUp();

        slasherMock = new SlasherMock();
        delegationTermsMock = new DelegationTermsMock();
        delegationContract = DelegationManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        strategyManagerMock = new StrategyManagerMock();

        delegationContractImplementation = new DelegationManager(strategyManagerMock, slasherMock);

        cheats.startPrank(eigenLayerProxyAdmin.owner());
        eigenLayerProxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(delegationContract))), address(delegationContractImplementation));
        cheats.stopPrank();
        
        delegationContract.initialize(address(this), eigenLayerPauserReg, 0);

        strategyImplementation = new StrategyBase(strategyManager);

        strategyMock = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, weth, eigenLayerPauserReg)
                )
            )
        );

    }

    function testReinitializeDelegation() public{
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        delegationContract.initialize(address(this), eigenLayerPauserReg, 0);
    }

    function testBadECDSASignatureExpiry(address staker, address operator, uint256 expiry, bytes memory signature) public{
        cheats.assume(expiry < block.timestamp);
        cheats.expectRevert(bytes("DelegationManager.delegateToBySignature: delegation signature expired"));
        delegationContract.delegateToBySignature(staker, operator, expiry, signature);
    }

    function testUndelegateFromNonStrategyManagerAddress(address undelegator) public fuzzedAddress(undelegator){
        cheats.assume(undelegator != address(strategyManagerMock));
        cheats.expectRevert(bytes("onlyStrategyManager"));
        cheats.startPrank(undelegator);
        delegationContract.undelegate(address(this));
    }

    function testUndelegateWithFrozenOperator(address undelegator) public fuzzedAddress(undelegator) {
        cheats.startPrank(undelegator);
        delegationContract.registerAsOperator(IDelegationTerms(address(this)));
        cheats.stopPrank();

        slasherMock.setOperatorFrozenStatus(address(this), true);
        cheats.expectRevert(bytes("StrategyManager.onlyNotFrozen: staker has been frozen and may be subject to slashing"));
        cheats.startPrank(address(strategyManagerMock));
        delegationContract.undelegate(address(this));
        cheats.stopPrank();
    }

    function testUndelegateByOperatorFromThemselves(address operator) public fuzzedAddress(operator){
        cheats.startPrank(operator);
        delegationContract.registerAsOperator(IDelegationTerms(address(this)));
        cheats.stopPrank();
        cheats.expectRevert(bytes("DelegationManager.undelegate: operators cannot undelegate from themselves"));
        
        cheats.startPrank(address(strategyManagerMock));
        delegationContract.undelegate(operator);
        cheats.stopPrank();
    }

    function testIncreaseDelegatedSharesFromNonStrategyManagerAddress(address operator, uint256 shares) public fuzzedAddress(operator){
        cheats.assume(operator != address(strategyManagerMock));
        cheats.expectRevert(bytes("onlyStrategyManager"));
        cheats.startPrank(operator);
        delegationContract.increaseDelegatedShares(operator, strategyMock, shares);
    }

    function testDecreaseDelegatedSharesFromNonStrategyManagerAddress(
        address operator,  
        IStrategy[] memory strategies,  
        uint256[] memory shareAmounts
    ) public fuzzedAddress(operator){
        cheats.assume(operator != address(strategyManagerMock));
        cheats.expectRevert(bytes("onlyStrategyManager"));
        cheats.startPrank(operator);
        delegationContract.decreaseDelegatedShares(operator, strategies, shareAmounts);
    }

    function testDelegateWhenOperatorIsFrozen(address operator, address staker) public fuzzedAddress(operator) fuzzedAddress(staker){
        cheats.assume(operator != staker);
        
        cheats.startPrank(operator);
        delegationContract.registerAsOperator(IDelegationTerms(address(this)));
        cheats.stopPrank();

        slasherMock.setOperatorFrozenStatus(operator, true);
        cheats.expectRevert(bytes("DelegationManager._delegate: cannot delegate to a frozen operator"));
        cheats.startPrank(staker);
        delegationContract.delegateTo(operator);
        cheats.stopPrank();
    }

    function testDelegateWhenStakerHasExistingDelegation(address staker, address operator, address operator2) public
        fuzzedAddress(staker)
        fuzzedAddress(operator)
        fuzzedAddress(operator2)
    {
        cheats.assume(operator != operator2);
        cheats.assume(staker != operator);
        cheats.assume(staker != operator2);

        cheats.startPrank(operator);
        delegationContract.registerAsOperator(IDelegationTerms(address(11)));
        cheats.stopPrank();

        cheats.startPrank(operator2);
        delegationContract.registerAsOperator(IDelegationTerms(address(10)));
        cheats.stopPrank();

        cheats.startPrank(staker);
        delegationContract.delegateTo(operator);
        cheats.stopPrank();

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("DelegationManager._delegate: staker has existing delegation"));
        delegationContract.delegateTo(operator2);
        cheats.stopPrank();
    }

    function testDelegationToUnregisteredOperator(address operator) public{
        cheats.expectRevert(bytes("DelegationManager._delegate: operator has not yet registered as a delegate"));
        delegationContract.delegateTo(operator);
    }

    function testDelegationWhenPausedNewDelegationIsSet(address operator, address staker) public fuzzedAddress(operator) fuzzedAddress(staker){
        cheats.startPrank(pauser);
        delegationContract.pause(1);
        cheats.stopPrank();

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("Pausable: index is paused"));
        delegationContract.delegateTo(operator);
        cheats.stopPrank();
    }

    function testRevertingDelegationReceivedHook(address operator, address staker) public fuzzedAddress(operator) fuzzedAddress(staker) {
        cheats.assume(operator != staker);

        delegationTermsMock.setShouldRevert(true);
        cheats.startPrank(operator);
        delegationContract.registerAsOperator(delegationTermsMock);
        cheats.stopPrank();

        cheats.startPrank(staker);
        cheats.expectEmit(true, false, false, false);
        emit OnDelegationReceivedCallFailure(delegationTermsMock, 0x0000000000000000000000000000000000000000000000000000000000000000);
        delegationContract.delegateTo(operator);
        cheats.stopPrank();
    }

    function testRevertingDelegationWithdrawnHook(
        address operator, 
        address staker
    ) public fuzzedAddress(operator) fuzzedAddress(staker){
        cheats.assume(operator != staker);
        delegationTermsMock.setShouldRevert(true);

        cheats.startPrank(operator);
        delegationContract.registerAsOperator(delegationTermsMock);
        cheats.stopPrank();

        cheats.startPrank(staker);
        delegationContract.delegateTo(operator);
        cheats.stopPrank();

        (IStrategy[] memory updatedStrategies, uint256[] memory updatedShares) =
            strategyManager.getDeposits(staker);

        cheats.startPrank(address(strategyManagerMock));
        cheats.expectEmit(true, false, false, false);
        emit OnDelegationWithdrawnCallFailure(delegationTermsMock, 0x0000000000000000000000000000000000000000000000000000000000000000);
        delegationContract.decreaseDelegatedShares(staker, updatedStrategies, updatedShares);
        cheats.stopPrank();
    }

    function testDelegationReceivedHookWithTooMuchReturnData(address operator, address staker) public fuzzedAddress(operator) fuzzedAddress(staker){
        cheats.assume(operator != staker);
        delegationTermsMock.setShouldReturnData(true);

        cheats.startPrank(operator);
        delegationContract.registerAsOperator(delegationTermsMock);
        cheats.stopPrank();

        cheats.startPrank(staker);
        delegationContract.delegateTo(operator);
        cheats.stopPrank();
    }

    function testDelegationWithdrawnHookWithTooMuchReturnData(
        address operator, 
        address staker
    ) public fuzzedAddress(operator) fuzzedAddress(staker){
        cheats.assume(operator != staker);

        delegationTermsMock.setShouldReturnData(true);


        cheats.startPrank(operator);
        delegationContract.registerAsOperator(delegationTermsMock);
        cheats.stopPrank();

        cheats.startPrank(staker);
        delegationContract.delegateTo(operator);
        cheats.stopPrank();

        (IStrategy[] memory updatedStrategies, uint256[] memory updatedShares) =
            strategyManager.getDeposits(staker);

        cheats.startPrank(address(strategyManagerMock));
        delegationContract.decreaseDelegatedShares(staker, updatedStrategies, updatedShares);
        cheats.stopPrank();
    }

    function testDelegationReceivedHookWithNoReturnData(address operator, address staker) public fuzzedAddress(operator) fuzzedAddress(staker){
        cheats.assume(operator != staker);

        cheats.startPrank(operator);
        delegationContract.registerAsOperator(delegationTermsMock);
        cheats.stopPrank();

        cheats.startPrank(staker);
        delegationContract.delegateTo(operator);
        cheats.stopPrank();
    }

    function testDelegationWithdrawnHookWithNoReturnData(
        address operator, 
        address staker
    ) public fuzzedAddress(operator) fuzzedAddress(staker){
        cheats.assume(operator != staker);

        cheats.startPrank(operator);
        delegationContract.registerAsOperator(delegationTermsMock);
        cheats.stopPrank();

        cheats.startPrank(staker);
        delegationContract.delegateTo(operator);
        cheats.stopPrank();

        (IStrategy[] memory updatedStrategies, uint256[] memory updatedShares) =
            strategyManager.getDeposits(staker);

        cheats.startPrank(address(strategyManagerMock));
        delegationContract.decreaseDelegatedShares(staker, updatedStrategies, updatedShares);
        cheats.stopPrank();
    }

}