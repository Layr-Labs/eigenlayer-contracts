// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";


interface IOldPod {
    function GENESIS_TIME() external view returns (uint64);
}

contract Integration_Upgrade_MOOCOW_Base is UpgradeTest {
    AVS avs;
    OperatorSet operatorSet;

    User operator;
    AllocateParams allocateParams;
    SlashingParams slashParams;

    User staker;
    IStrategy[] strategies;
    IERC20[] tokens; // underlying token for each strategy
    uint[] initTokenBalances;
    uint[] initDepositShares;
    
    address constant pod_addr = 0xA6f93249580EC3F08016cD3d4154AADD70aC3C96;
    EigenPod constant pod = EigenPod(payable(pod_addr));
    bytes constant validatorPubkey = hex"9067f3fb35c5fc6aa9e05f102519c22050c447acaf6b88aeb16d92fc14e30891f4c2bc21ee95ef5779224f5ceef3a168";

    /// Shared setup:
    ///
    /// 1. Generate staker, operator, and AVS
    /// 2. Staker deposits and delegates to operator
    /// 3. AVS creates an operator set containing the strategies held by the staker
    /// 4. Operator allocates to operator set
    /// 5. Operator registers for operator set
    /// 6. Operator is randomly slashed by the operatorSet
    function _init() internal override {
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator_NoAssets();
        (avs,) = _newRandomAVS();
        tokens = _getUnderlyingTokens(strategies);

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);

        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // 3. Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);

        // 4. Operator allocates to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);

        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // 5. Operator registers for operator set
        operator.registerForOperatorSet(operatorSet);

        // 6. Operator is randomly slashed by the operatorSet
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator_PreRedistribution(slashParams);
    }
}

contract Integration_Upgrade_MOOCOW is Integration_Upgrade_MOOCOW_Base {
    function testFuzz_upgrade_checkGenesisTime(uint24 r) public rand(r) {

        // 1. Upgrade contracts
        _upgradeEigenLayerContracts();
        
        bool methodExists;
        uint64 genesisTime;
        try IOldPod(pod_addr).GENESIS_TIME() returns (uint64 gt) {
            methodExists = true;
            genesisTime = gt;
        } catch (bytes memory) {
            methodExists = false;
        }
        
        assertEq(methodExists, false);
    }
    
    function testFuzz_upgrade_checkValidatorPubkey(uint24 r) public rand(r) {
        // 1. Upgrade contracts
        _upgradeEigenLayerContracts();
        
        address podOwner = pod.podOwner();
        cheats.startPrank(podOwner);
        
        ConsolidationRequest[] memory conReqs = new ConsolidationRequest[](1);
        conReqs[0] = ConsolidationRequest({
            srcPubkey: validatorPubkey,
            targetPubkey: validatorPubkey
        });

        WithdrawalRequest[] memory witReqs = new WithdrawalRequest[](1);
        witReqs[0] = WithdrawalRequest({
            pubkey: validatorPubkey,
            amountGwei: 0
        });
        
        bool conMethodsExist;
        bool feeRequestSucceeds;
        bool conRequestSucceeds;
        uint conFee;
        try pod.getConsolidationRequestFee() returns (uint fee) {
            feeRequestSucceeds = true;
            conFee = fee;
        } catch (bytes memory) {
            feeRequestSucceeds = false;
        }

        try pod.requestConsolidation{ value: conFee }(conReqs) {
            conRequestSucceeds = true;
        } catch (bytes memory) {
            conRequestSucceeds = false;
        }
        
        cheats.stopPrank();
        
        require(feeRequestSucceeds, "fee request failed");

        // There should be no world where these disagree with each other
        require(feeRequestSucceeds == conRequestSucceeds, "consolidation method existence mismatch");
    }
    
    function testFuzz_upgrade_checkWithdrawal(uint24 r) public rand(r) {
        // 1. Upgrade contracts
        _upgradeEigenLayerContracts();
        
        address podOwner = pod.podOwner();
        cheats.startPrank(podOwner);
        
        bool feeRequestSucceeds;
        bool witRequestSucceeds;
        uint witFee;
        try pod.getWithdrawalRequestFee() returns (uint fee) {
            feeRequestSucceeds = true;
            witFee = fee;
        } catch (bytes memory) {
            feeRequestSucceeds = false;
        }

        WithdrawalRequest[] memory witReqs = new WithdrawalRequest[](1);
        witReqs[0] = WithdrawalRequest({
            pubkey: validatorPubkey,
            amountGwei: 0
        });

        try pod.requestWithdrawal{ value: witFee }(witReqs) {
            witRequestSucceeds = true;
        } catch (bytes memory) {
            witRequestSucceeds = false;
        }

        // There should be no world where these disagree with each other
        require(feeRequestSucceeds == witRequestSucceeds, "withdrawal method existence mismatch");

        cheats.stopPrank();
    }
}