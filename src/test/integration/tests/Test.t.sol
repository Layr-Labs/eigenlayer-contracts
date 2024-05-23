// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

enum Network {
    LOCAL,
    HOLESKY,
    MAINNET
}

enum Version {
    LATEST,
    M2_0_2_5, // m2 + eigen + pod fix
    M2,       // m2 + eigen
    M1
}

contract Integration_Tester is IntegrationCheckUtils {

    // function setUp() public override {
    //     _select({
    //         network: Network.LOCAL,
    //         version: Version.LATEST
    //         // version: Version.0_2_5
    //     });

    //     super.setUp();
    // }

    // function test_Example() public {
    //     _config({
    //         _randomSeed: 0,
    //         _assetTypes: HOLDS_ETH,
    //         _userTypes: DEFAULT
    //     });

    //     // User staker = _newStaker(Assets.ETH);
    //     // User operator = _newOperator(Assets.NONE);

    //     User staker = _newRandomStaker();

    //     // staker.startValidators();        // start a random number of validators
    //     // staker.startNonFullValidators(); // start a random number of validators with < 32 ETH
    //     uint40[] memory validators = staker.startValidators(1);

    //     // staker.delegateTo(operator);
    //     // check_Delegation_State(staker, operator, ...);

    //     staker.verifyWithdrawalCredentials(validators);
    //     // check_VerifyWC_State(staker, validators);
    //     // assert_Snap_Added_OperatorShares(operator, BEACONCHAIN_ETH_STRAT, )
        
    //     // staker.startCheckpoint();
    //     // check_StartCheckpoint_State(staker, validators);

    //     // staker.completeCheckpoint();
    //     // check_CompleteCheckpoint_State(staker, validators);

    //     // beaconChain.rewardValidators();
        

    // }

    function test_Thing2() public {
        _configRand({
            _randomSeed: 0,
            _assetTypes: HOLDS_ETH,
            _userTypes: DEFAULT
        });

        emit log("1");

        (User staker, ,) = _newRandomStaker();

        uint40[] memory validators = staker.startValidators(4);

        for (uint i = 0; i < validators.length; i++) {
            _logValidator(validators[i]);
        }

        staker.verifyWithdrawalCredentials(validators);
    }

    // function test_Thing() public {
    //     beaconChain.newValidator(1 ether, "");
    //     beaconChain.newValidator(2 ether, "");

    //     ValidatorFieldsProof[] memory proofs = beaconChain.processEpoch();

    //     for (uint i = 0; i < proofs.length; i++) {
    //         ValidatorFieldsProof memory proof = proofs[i];

    //         bytes32 stateRoot = Merkle.processInclusionProofSha256({
    //             proof: proof.validatorFieldsProof,
    //             leaf: Merkle.merkleizeSha256(proof.validatorFields),
    //             index: uint(proof.validatorIndex)
    //         });

    //         emit log_named_bytes32("state root", stateRoot);
    //         _logValidator(proof.validatorIndex);
    //     }

    //     beaconChain.updateValidator(0, 1 ether);

    //     proofs = beaconChain.processEpoch();

    //     for (uint i = 0; i < proofs.length; i++) {
    //         ValidatorFieldsProof memory proof = proofs[i];

    //         bytes32 stateRoot = Merkle.processInclusionProofSha256({
    //             proof: proof.validatorFieldsProof,
    //             leaf: Merkle.merkleizeSha256(proof.validatorFields),
    //             index: uint(proof.validatorIndex)
    //         });

    //         emit log_named_bytes32("state root", stateRoot);
    //         _logValidator(proof.validatorIndex);
    //     }

    //     uint balanceExitedGwei = beaconChain.exitValidator(0) / GWEI_TO_WEI;
    //     emit log_named_uint("exited gwei", balanceExitedGwei);

    //     proofs = beaconChain.processEpoch();

    //     for (uint i = 0; i < proofs.length; i++) {
    //         ValidatorFieldsProof memory proof = proofs[i];

    //         bytes32 stateRoot = Merkle.processInclusionProofSha256({
    //             proof: proof.validatorFieldsProof,
    //             leaf: Merkle.merkleizeSha256(proof.validatorFields),
    //             index: uint(proof.validatorIndex)
    //         });

    //         emit log_named_bytes32("state root", stateRoot);
    //         _logValidator(proof.validatorIndex);
    //     }
    // }

    function _logValidator(uint40 validatorIndex) internal {
        emit log_named_uint("validator", validatorIndex);
        emit log_named_uint("current balance", beaconChain.currentBalance(validatorIndex));
        emit log_named_uint("effective balance", beaconChain.effectiveBalance(validatorIndex));
        emit log_named_uint("exit epoch", beaconChain.exitEpoch(validatorIndex));

        emit log("===");
    }

    // function test_Thing() public {
    //     uint40 aIdx = beaconChain.newValidator2(1 ether, "");

    //     bytes32 balanceRoot = beaconChain.getBalanceRoot(aIdx);
    //     emit log_named_bytes32("balance root a 1", balanceRoot);
    //     uint64 balanceA = beaconChain.getBalance(aIdx);
    //     emit log_named_uint("balance a 1", balanceA);

    //     uint40 bIdx = beaconChain.newValidator2(2 ether, "");

    //     balanceRoot = beaconChain.getBalanceRoot(bIdx);
    //     emit log_named_bytes32("balance root ab 12", balanceRoot);
    //     balanceA = beaconChain.getBalance(aIdx);
    //     emit log_named_uint("balance a 1", balanceA);
    //     uint64 balanceB = beaconChain.getBalance(bIdx);
    //     emit log_named_uint("balance b 2", balanceB);

    //     uint40 cIdx = beaconChain.newValidator2(3 ether, "");

    //     balanceRoot = beaconChain.getBalanceRoot(cIdx);
    //     emit log_named_bytes32("balance root abc 123", balanceRoot);
    //     balanceA = beaconChain.getBalance(aIdx);
    //     emit log_named_uint("balance a 1", balanceA);
    //     balanceB = beaconChain.getBalance(bIdx);
    //     emit log_named_uint("balance b 2", balanceB);
    //     uint64 balanceC = beaconChain.getBalance(cIdx);
    //     emit log_named_uint("balance c 3", balanceC);

    //     uint40 dIdx = beaconChain.newValidator2(4 ether, "");

    //     balanceRoot = beaconChain.getBalanceRoot(dIdx);
    //     emit log_named_bytes32("balance root abc 1234", balanceRoot);
    //     balanceA = beaconChain.getBalance(aIdx);
    //     emit log_named_uint("balance a 1", balanceA);
    //     balanceB = beaconChain.getBalance(bIdx);
    //     emit log_named_uint("balance b 2", balanceB);
    //     balanceC = beaconChain.getBalance(cIdx);
    //     emit log_named_uint("balance c 3", balanceC);
    //     uint64 balanceD = beaconChain.getBalance(dIdx);
    //     emit log_named_uint("balance d 4", balanceD);

    //     uint40 eIdx = beaconChain.newValidator2(5 ether, "");

    //     balanceRoot = beaconChain.getBalanceRoot(eIdx);
    //     emit log_named_bytes32("balance root e 5", balanceRoot);
    //     uint64 balanceE = beaconChain.getBalance(eIdx);
    //     emit log_named_uint("balance e 5", balanceE);
    // }
}