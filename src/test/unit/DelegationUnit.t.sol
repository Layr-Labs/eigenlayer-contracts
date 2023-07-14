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
    DelegationManager delegationManager;
    DelegationTermsMock delegationTermsMock;
    DelegationManager delegationManagerImplementation;
    StrategyBase strategyImplementation;
    StrategyBase strategyMock;


    uint256 GWEI_TO_WEI = 1e9;

    event OnDelegationReceivedCallFailure(IDelegationTerms indexed delegationTerms, bytes32 returnData);
    event OnDelegationWithdrawnCallFailure(IDelegationTerms indexed delegationTerms, bytes32 returnData);


    function setUp() override virtual public{
        EigenLayerDeployer.setUp();

        slasherMock = new SlasherMock();
        delegationTermsMock = new DelegationTermsMock();
        delegationManager = DelegationManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        strategyManagerMock = new StrategyManagerMock();

        delegationManagerImplementation = new DelegationManager(strategyManagerMock, slasherMock);

        cheats.startPrank(eigenLayerProxyAdmin.owner());
        eigenLayerProxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(delegationManager))), address(delegationManagerImplementation));
        cheats.stopPrank();
        
        delegationManager.initialize(address(this), eigenLayerPauserReg, 0);

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
        delegationManager.initialize(address(this), eigenLayerPauserReg, 0);
    }

    function testBadECDSASignatureExpiry(address staker, address operator, uint256 expiry, bytes memory signature) public{
        cheats.assume(expiry < block.timestamp);
        cheats.expectRevert(bytes("DelegationManager.delegateToBySignature: staker signature expired"));
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry = IDelegationManager.SignatureWithExpiry({
            signature: signature,
            expiry: expiry
        });
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry);
    }

    function testUndelegateFromNonStrategyManagerAddress(address undelegator) public fuzzedAddress(undelegator) {
        cheats.assume(undelegator != address(strategyManagerMock));
        cheats.expectRevert(bytes("onlyStrategyManager"));
        cheats.startPrank(undelegator);
        delegationManager.undelegate(address(this));
    }

    function testUndelegateByOperatorFromThemselves(address operator) public fuzzedAddress(operator) {
        cheats.startPrank(operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(operatorDetails);
        cheats.stopPrank();
        cheats.expectRevert(bytes("DelegationManager.undelegate: operators cannot undelegate from themselves"));
        
        cheats.startPrank(address(strategyManagerMock));
        delegationManager.undelegate(operator);
        cheats.stopPrank();
    }

    function testIncreaseDelegatedSharesFromNonStrategyManagerAddress(address operator, uint256 shares) public fuzzedAddress(operator) {
        cheats.assume(operator != address(strategyManagerMock));
        cheats.expectRevert(bytes("onlyStrategyManager"));
        cheats.startPrank(operator);
        delegationManager.increaseDelegatedShares(operator, strategyMock, shares);
    }

    function testDecreaseDelegatedSharesFromNonStrategyManagerAddress(
        address operator,  
        IStrategy[] memory strategies,  
        uint256[] memory shareAmounts
    ) public fuzzedAddress(operator) {
        cheats.assume(operator != address(strategyManagerMock));
        cheats.expectRevert(bytes("onlyStrategyManager"));
        cheats.startPrank(operator);
        delegationManager.decreaseDelegatedShares(operator, strategies, shareAmounts);
    }

    function testDelegateWhenOperatorIsFrozen(address operator, address staker) public fuzzedAddress(operator) fuzzedAddress(staker) {
        cheats.assume(operator != staker);
        
        cheats.startPrank(operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(operatorDetails);
        cheats.stopPrank();

        slasherMock.setOperatorFrozenStatus(operator, true);
        cheats.expectRevert(bytes("DelegationManager._delegate: cannot delegate to a frozen operator"));
        cheats.startPrank(staker);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry);
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
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(operatorDetails);
        cheats.stopPrank();

        cheats.startPrank(operator2);
        delegationManager.registerAsOperator(operatorDetails);
        cheats.stopPrank();

        cheats.startPrank(staker);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry);
        cheats.stopPrank();

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("DelegationManager._delegate: staker has existing delegation"));
        delegationManager.delegateTo(operator2, signatureWithExpiry);
        cheats.stopPrank();
    }

    function testDelegationToUnregisteredOperator(address operator) public{
        cheats.expectRevert(bytes("DelegationManager._delegate: operator is not registered in EigenLayer"));
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry);
    }

    function testDelegationWhenPausedNewDelegationIsSet(address operator, address staker) public fuzzedAddress(operator) fuzzedAddress(staker) {
        cheats.startPrank(pauser);
        delegationManager.pause(1);
        cheats.stopPrank();

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("Pausable: index is paused"));
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry);
        cheats.stopPrank();
    }
}